package livestream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 直播弹幕信息流的连接、认证、心跳与事件分发。

var (
	errDanmakuNilContext            = errors.New("live danmaku context is nil")
	errDanmakuAuthRejected          = errors.New("live danmaku auth was rejected")
	errDanmakuClosed                = errors.New("live danmaku stream is closed")
	errDanmakuNoUsablePort          = errors.New("live danmaku host has no usable port")
	errDanmakuAuthReplyMissing      = errors.New("live danmaku auth reply is missing")
	errDanmakuExpectedBinaryMessage = errors.New("live danmaku expected binary message")
)

// Event 是信息流上的一条事件。Cmd 为服务端命令名，Raw 是未加工的 JSON 消息体。
type Event struct {
	Cmd  string          `json:"cmd"`
	Raw  json.RawMessage `json:"raw"`
	Room int64           `json:"room"`
}

// Stream 是一条已建立的直播弹幕信息流。
type Stream struct {
	room     int64
	events   chan Event
	done     chan struct{}
	closeErr error
	once     sync.Once
	mu       sync.RWMutex
	conn     *websocket.Conn
}

// Events 返回事件 channel。连接结束或主动关闭后 channel 关闭。
func (s *Stream) Events() <-chan Event { return s.events }

// Done 在连接结束后关闭。
func (s *Stream) Done() <-chan struct{} { return s.done }

// Close 主动关闭连接，幂等。
func (s *Stream) Close() error {
	var closeErr error
	s.once.Do(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.conn != nil {
			closeErr = s.conn.Close()
			if s.closeErr == nil {
				s.closeErr = closeErr
			}
		}
		close(s.done)
	})
	return closeErr
}

// Err 返回连接结束原因；主动关闭时可能为底层连接关闭错误。
func (s *Stream) Err() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.closeErr
}

// WebSocketDialer 建立信息流 WebSocket 连接；可由调用方提供代理或测试实现。
type WebSocketDialer interface {
	DialContext(ctx context.Context, network, address string, header http.Header) (*websocket.Conn, *http.Response, error)
}

// GorillaDialer 用 gorilla/websocket 的默认拨号器实现 WebSocketDialer。
type GorillaDialer struct {
	Dialer *websocket.Dialer
}

// DialContext 建立连接。network 只用于标识是否使用 TLS，实际 URL 由 address 决定。
func (d GorillaDialer) DialContext(ctx context.Context, _, address string, header http.Header) (*websocket.Conn, *http.Response, error) {
	dialer := d.Dialer
	if dialer == nil {
		dialer = websocket.DefaultDialer
	}
	return dialer.DialContext(ctx, address, header)
}

func defaultWSDialer() WebSocketDialer {
	return GorillaDialer{Dialer: &websocket.Dialer{
		HandshakeTimeout: 15 * time.Second,
		ReadBufferSize:   danmakuReadBufferSize,
		WriteBufferSize:  danmakuReadBufferSize,
		Proxy:            http.ProxyFromEnvironment,
	}}
}

func networkForTLS(useTLS bool) string {
	if useTLS {
		return "tcp-tls"
	}
	return "tcp"
}

// Connect 获取弹幕 token、建立 WebSocket、完成认证并返回事件流。
// ctx 只约束连接准备阶段；连接生命周期由 Stream.Close 控制。
func (c *Client) Connect(ctx context.Context, options ConnectOptions) (*Stream, error) {
	if ctx == nil {
		return nil, errDanmakuNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	info, err := c.connectInfo(ctx, options)
	if err != nil {
		return nil, err
	}
	dialer := options.Dialer
	if dialer == nil {
		dialer = defaultWSDialer()
	}
	scheme, port := "ws", info.Host.WSPort
	if info.UseTLS {
		scheme, port = "wss", info.Host.WSSPort
	}
	if port <= 0 {
		return nil, errDanmakuNoUsablePort
	}
	address := scheme + "://" + info.Host.Host + ":" + strconv.Itoa(port) + "/sub"
	header := http.Header{}
	header.Set("User-Agent", danmakuUserAgent)
	header.Set("Origin", "https://live.bilibili.com")
	conn, _, err := dialer.DialContext(ctx, networkForTLS(info.UseTLS), address, header)
	if err != nil {
		return nil, fmt.Errorf("connect live danmaku websocket: %w", err)
	}
	auth, err := authPacket(info)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	stream := &Stream{
		room:   info.RoomID,
		events: make(chan Event, eventBufferSize(options.EventBufferSize)),
		done:   make(chan struct{}),
		conn:   conn,
	}
	go stream.run(conn, auth, heartbeatInterval(options.HeartbeatInterval))
	return stream, nil
}

func heartbeatInterval(value time.Duration) time.Duration {
	if value <= 0 {
		return danmakuHeartbeatPeriod * time.Second
	}
	return value
}

func eventBufferSize(value int) int {
	if value <= 0 {
		return danmakuEventBufferSize
	}
	return value
}

func (s *Stream) run(conn *websocket.Conn, auth []byte, heartbeat time.Duration) {
	defer func() {
		_ = conn.Close()
		close(s.events)
		s.once.Do(func() { close(s.done) })
	}()

	if err := writePacket(conn, auth); err != nil {
		s.setError(err)
		return
	}
	first, err := readPacket(conn)
	if err != nil {
		s.setError(err)
		return
	}
	if err := validateAuthPacket(first); err != nil {
		s.setError(err)
		return
	}
	if err := writePacket(conn, heartbeatPacket()); err != nil {
		s.setError(err)
		return
	}

	readResult := make(chan danmakuReadResult, 1)
	go readDanmakuPacket(conn, readResult)
	ticker := time.NewTicker(heartbeat)
	defer ticker.Stop()
	for {
		select {
		case <-s.done:
			return
		case <-ticker.C:
			if err := writePacket(conn, heartbeatPacket()); err != nil {
				s.setError(err)
				return
			}
		case result := <-readResult:
			if result.err != nil {
				if !errors.Is(result.err, errDanmakuClosed) {
					s.setError(result.err)
				}
				return
			}
			packets, err := decodeDanmakuPackets(result.data)
			if err != nil {
				s.setError(err)
				return
			}
			for _, item := range packets {
				if !s.dispatch(item) {
					return
				}
			}
			go readDanmakuPacket(conn, readResult)
		}
	}
}

func writePacket(conn *websocket.Conn, packet []byte) error {
	_ = conn.SetWriteDeadline(time.Now().Add(danmakuWriteTimeout * time.Second))
	return conn.WriteMessage(websocket.BinaryMessage, packet)
}

func readPacket(conn *websocket.Conn) ([]byte, error) {
	messageType, reader, err := conn.NextReader()
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			return nil, errDanmakuClosed
		}
		return nil, err
	}
	if messageType != websocket.BinaryMessage {
		return nil, fmt.Errorf("%w: %d", errDanmakuExpectedBinaryMessage, messageType)
	}
	return readAllLimited(reader)
}

func (s *Stream) dispatch(packet danmakuPacket) bool {
	switch packet.Operation {
	case danmakuOpHeartbeatReply:
		return true
	case danmakuOpAuthReply:
		if err := validateAuthReply(packet.Body); err != nil {
			s.setError(err)
			return false
		}
		return true
	case danmakuOpMessage:
		event := Event{
			Cmd:  danmakuCommandOf(packet.Body),
			Raw:  append([]byte(nil), packet.Body...),
			Room: s.room,
		}
		select {
		case s.events <- event:
			return true
		case <-s.done:
			return false
		}
	default:
		return true
	}
}

func validateAuthPacket(frame []byte) error {
	packets, err := decodeDanmakuPackets(frame)
	if err != nil {
		return err
	}
	for _, packet := range packets {
		if packet.Operation != danmakuOpAuthReply {
			continue
		}
		return validateAuthReply(packet.Body)
	}
	return errDanmakuAuthReplyMissing
}

func validateAuthReply(body []byte) error {
	var reply struct {
		Code int `json:"code"`
	}
	if err := json.Unmarshal(body, &reply); err == nil && reply.Code != 0 {
		return fmt.Errorf("%w: code %d", errDanmakuAuthRejected, reply.Code)
	}
	return nil
}

func (s *Stream) setError(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closeErr == nil {
		s.closeErr = err
	}
}
