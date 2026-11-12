package livestream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 直播弹幕信息流的连接、认证、心跳、自动重连与事件分发。

var (
	errDanmakuNilContext            = errors.New("live danmaku context is nil")
	errDanmakuAuthRejected          = errors.New("live danmaku auth was rejected")
	errDanmakuAuthTokenExpired      = errors.New("live danmaku auth token expired")
	errDanmakuAuthRiskControl       = errors.New("live danmaku auth rejected by risk control")
	errDanmakuClosed                = errors.New("live danmaku stream is closed")
	errDanmakuNoUsablePort          = errors.New("live danmaku host has no usable port")
	errDanmakuAuthReplyMissing      = errors.New("live danmaku auth reply is missing")
	errDanmakuExpectedBinaryMessage = errors.New("live danmaku expected binary message")
	errDanmakuAllHostsFailed        = errors.New("live danmaku all hosts failed")
	errDanmakuJoinTimeout           = errors.New("live danmaku join timed out")
	errDanmakuReadIdleTimeout       = errors.New("live danmaku connection read idle timed out")
	errDanmakuRefreshTokenFailed    = errors.New("live danmaku token refresh failed")
)

// Event 是信息流上的一条事件。Cmd 为服务端命令名，Raw 是未加工的 JSON 消息体。
type Event struct {
	Cmd  string          `json:"cmd"`
	Raw  json.RawMessage `json:"raw"`
	Room int64           `json:"room"`
}

// Stream 是一条已建立的直播弹幕信息流。断线后自动轮换节点重连，
// 生命周期由 Close 控制，调用方传入的 context 只约束连接准备阶段。
type Stream struct {
	room     int64
	events   chan Event
	done     chan struct{}
	closeErr error
	once     sync.Once
	mu       sync.RWMutex
	conn     *websocket.Conn
	//nolint:containedctx // 信息流生命周期由 Close 控制，context 用于取消后台重连与读取
	ctx    context.Context
	cancel context.CancelFunc

	client          *Client
	info            danmakuConnectionInfo
	dialer          WebSocketDialer
	heartbeat       time.Duration
	joinTimeout     time.Duration
	readIdleTimeout time.Duration
	reconnectDelay  time.Duration
	attempt         int
}

// Events 返回事件 channel。连接结束或主动关闭后 channel 关闭。
func (s *Stream) Events() <-chan Event { return s.events }

// Done 在连接结束后关闭。
func (s *Stream) Done() <-chan struct{} { return s.done }

// Close 主动关闭连接，幂等。关闭后不会再重连。
func (s *Stream) Close() error {
	var closeErr error
	s.once.Do(func() {
		s.cancel()
		s.mu.Lock()
		if s.conn != nil {
			closeErr = s.conn.Close()
			if s.closeErr == nil {
				s.closeErr = closeErr
			}
		}
		s.mu.Unlock()
		close(s.done)
	})
	return closeErr
}

// Err 返回导致信息流结束的致命错误；主动关闭时为底层连接的关闭错误。
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

func durationOrDefault(value time.Duration, fallback time.Duration) time.Duration {
	if value <= 0 {
		return fallback
	}
	return value
}

// Connect 获取弹幕 token、建立 WebSocket、完成认证并返回事件流。
// 首个连接同步建立，认证失败会立即返回；之后断线由后台自动重连。
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
	runCtx, cancel := context.WithCancel(context.Background())
	stream := &Stream{
		room:            info.RoomID,
		events:          make(chan Event, eventBufferSize(options.EventBufferSize)),
		done:            make(chan struct{}),
		ctx:             runCtx,
		cancel:          cancel,
		client:          c,
		info:            info,
		dialer:          dialer,
		heartbeat:       heartbeatInterval(options.HeartbeatInterval),
		joinTimeout:     durationOrDefault(options.JoinTimeout, danmakuJoinTimeout*time.Second),
		readIdleTimeout: durationOrDefault(options.ReadIdleTimeout, danmakuReadIdleTimeout*time.Second),
		reconnectDelay:  durationOrDefault(options.ReconnectDelay, danmakuReconnectDelay*time.Second),
	}
	attempt := 0
	conn, initial, err := stream.openSession(&attempt)
	// 首次建连失败一律返回错误，由调用方决定是否重试。自动重连只处理
	// 已经建立过会话之后的断线，避免 Connect 长时间阻塞在不可达节点上。
	if err != nil {
		cancel()
		close(stream.events)
		close(stream.done)
		return nil, err
	}
	stream.attempt = attempt
	go stream.run(conn, initial)
	return stream, nil
}

func heartbeatInterval(value time.Duration) time.Duration {
	return durationOrDefault(value, danmakuHeartbeatPeriod*time.Second)
}

func eventBufferSize(value int) int {
	if value <= 0 {
		return danmakuEventBufferSize
	}
	return value
}

// run 在会话断开后轮换节点重连，直到 Close 或遇到致命错误。
func (s *Stream) run(conn *websocket.Conn, initial []danmakuPacket) {
	defer func() {
		s.closeConn()
		close(s.events)
		s.once.Do(func() { s.cancel(); close(s.done) })
	}()
	refreshFailures := 0
	riskRetries := 0
	delay := s.reconnectDelay
	sessionErr := s.runSession(conn, initial)
	for s.ctx.Err() == nil {
		s.closeConn()
		switch {
		case errors.Is(sessionErr, errDanmakuAuthTokenExpired):
			refreshFailures++
			if refreshFailures > danmakuMaxTokenRefresh {
				s.setError(sessionErr)
				return
			}
			if err := s.refreshToken(); err != nil {
				s.setError(err)
				return
			}
		case errors.Is(sessionErr, errDanmakuAuthRejected):
			// 除 -101 / ±352 外的认证拒绝通常是房间关闭或权限问题，持续重连没有意义。
			s.setError(sessionErr)
			return
		case errors.Is(sessionErr, errDanmakuAuthRiskControl):
			riskRetries++
			if riskRetries > danmakuMaxRiskRetries {
				s.setError(sessionErr)
				return
			}
			// 公共节点固定是候选列表最后一项，风控时直接切过去。
			s.attempt = len(s.info.Hosts) - 1
		}
		if err := waitRetry(s.ctx, delay); err != nil {
			return
		}
		delay = nextReconnectDelay(delay)
		var nextInitial []danmakuPacket
		conn, nextInitial, sessionErr = s.openSession(&s.attempt)
		if sessionErr != nil {
			if s.ctx.Err() != nil {
				return
			}
			continue
		}
		refreshFailures = 0
		riskRetries = 0
		delay = s.reconnectDelay
		sessionErr = s.runSession(conn, nextInitial)
	}
}

// nextReconnectDelay 让连续失败的重连逐步退避，上限 30 秒，避免网络中断时空转。
func nextReconnectDelay(delay time.Duration) time.Duration {
	const maxDelay = 30 * time.Second
	if delay <= 0 || delay >= maxDelay {
		return maxDelay
	}
	next := delay * 2
	if next > maxDelay {
		return maxDelay
	}
	return next
}

func waitRetry(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// refreshToken 重新拉取 getDanmuInfo，只更新 token 与节点列表，
// 不重复解析房间号或改写会话 Cookie。
func (s *Stream) refreshToken() error {
	result, err := s.client.danmuInfo(s.ctx, s.info.RoomID)
	if err != nil {
		return fmt.Errorf("%w: %w", errDanmakuRefreshTokenFailed, err)
	}
	if result == nil || result.Token == "" {
		return errDanmakuRefreshTokenFailed
	}
	s.info.Token = result.Token
	if len(result.HostList) > 0 {
		hosts := append([]getDanmuInfoHost(nil), result.HostList...)
		s.info.Hosts = append(hosts, fallbackDanmakuHost())
		s.attempt = 0
	}
	return nil
}

// openSession 依次尝试节点，完成拨号、入房认证并取回首帧的所有包。
func (s *Stream) openSession(attempt *int) (*websocket.Conn, []danmakuPacket, error) {
	var lastErr error
	for range s.info.Hosts {
		if s.ctx.Err() != nil {
			return nil, nil, s.ctx.Err()
		}
		index := *attempt % len(s.info.Hosts)
		*attempt = *attempt + 1
		conn, err := s.dialHost(s.info.Hosts[index])
		if err != nil {
			lastErr = err
			continue
		}
		auth, err := authPacket(s.info)
		if err != nil {
			_ = conn.Close()
			return nil, nil, err
		}
		if err = writePacket(conn, auth); err != nil {
			_ = conn.Close()
			lastErr = err
			continue
		}
		// 认证回复可能与首批消息分在不同 WebSocket 帧，入房时限内持续读取。
		_ = conn.SetReadDeadline(time.Now().Add(s.joinTimeout))
		var initial []danmakuPacket
		for {
			frame, err := readPacket(conn)
			if err != nil {
				_ = conn.Close()
				var netErr net.Error
				if errors.As(err, &netErr) && netErr.Timeout() {
					lastErr = fmt.Errorf("%w: %w", errDanmakuJoinTimeout, err)
				} else {
					lastErr = err
				}
				break
			}
			packets, err := decodeDanmakuPackets(frame)
			if err != nil {
				_ = conn.Close()
				lastErr = err
				break
			}
			initial = append(initial, packets...)
			authErr := validateAuthPacket(packets)
			if authErr == nil {
				_ = conn.SetReadDeadline(time.Time{})
				s.setConn(conn)
				return conn, initial, nil
			}
			if !errors.Is(authErr, errDanmakuAuthReplyMissing) {
				_ = conn.SetReadDeadline(time.Time{})
				_ = conn.Close()
				if errors.Is(authErr, errDanmakuAuthTokenExpired) || errors.Is(authErr, errDanmakuAuthRiskControl) {
					return nil, nil, authErr
				}
				lastErr = authErr
				break
			}
		}
	}
	if lastErr == nil {
		lastErr = errDanmakuInfoNoHost
	}
	return nil, nil, fmt.Errorf("%w: %w", errDanmakuAllHostsFailed, lastErr)
}

func (s *Stream) dialHost(host getDanmuInfoHost) (*websocket.Conn, error) {
	scheme, port := "ws", host.WSPort
	if s.info.UseTLS {
		scheme, port = "wss", host.WSSPort
	}
	if port <= 0 {
		return nil, errDanmakuNoUsablePort
	}
	address := scheme + "://" + host.Host + ":" + strconv.Itoa(port) + "/sub"
	header := http.Header{}
	header.Set("User-Agent", danmakuUserAgent)
	header.Set("Origin", "https://live.bilibili.com")
	conn, response, err := s.dialer.DialContext(s.ctx, networkForTLS(s.info.UseTLS), address, header)
	if response != nil && response.Body != nil {
		_ = response.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("connect live danmaku websocket: %w", err)
	}
	return conn, nil
}

// runSession 维持单个连接的心跳、读取与事件分发。任何返回都会触发重连。
func (s *Stream) runSession(conn *websocket.Conn, initial []danmakuPacket) error {
	sessionDone := make(chan struct{})
	defer close(sessionDone)
	results := make(chan danmakuReadResult, 1)
	go s.readLoop(conn, results, sessionDone)
	heartbeat := time.NewTicker(s.heartbeat)
	defer heartbeat.Stop()
	idle := time.NewTimer(s.readIdleTimeout)
	defer idle.Stop()
	if err := writePacket(conn, heartbeatPacket()); err != nil {
		return err
	}
	for _, packet := range initial {
		if err := s.dispatch(packet); err != nil {
			return err
		}
		if isActiveOperation(packet.Operation) {
			resetTimer(idle, s.readIdleTimeout)
		}
	}
	for {
		select {
		case <-s.done:
			return nil
		case <-s.ctx.Done():
			return nil
		case <-idle.C:
			return errDanmakuReadIdleTimeout
		case <-heartbeat.C:
			if err := writePacket(conn, heartbeatPacket()); err != nil {
				return err
			}
		case result := <-results:
			if result.err != nil {
				return result.err
			}
			packets, err := decodeDanmakuPackets(result.data)
			if err != nil {
				return err
			}
			for _, packet := range packets {
				if err = s.dispatch(packet); err != nil {
					return err
				}
				if isActiveOperation(packet.Operation) {
					resetTimer(idle, s.readIdleTimeout)
				}
			}
		}
	}
}

func isActiveOperation(operation uint32) bool {
	switch operation {
	case danmakuOpHeartbeatReply, danmakuOpAuthReply, danmakuOpMessage:
		return true
	default:
		return false
	}
}

func resetTimer(timer *time.Timer, delay time.Duration) {
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}
	_ = timer.Reset(delay)
}

// readLoop 是连接上唯一的读取者；连接关闭或会话结束时退出。
func (s *Stream) readLoop(conn *websocket.Conn, results chan<- danmakuReadResult, sessionDone <-chan struct{}) {
	for {
		result := danmakuReadResult{}
		messageType, reader, err := conn.NextReader()
		switch {
		case err != nil:
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				result.err = errDanmakuClosed
			} else {
				result.err = err
			}
		case messageType != websocket.BinaryMessage:
			result.err = fmt.Errorf("%w: %d", errDanmakuExpectedBinaryMessage, messageType)
		default:
			result.data, result.err = readAllLimited(reader)
		}
		select {
		case results <- result:
		case <-sessionDone:
			return
		case <-s.ctx.Done():
			return
		case <-s.done:
			return
		}
		if result.err != nil {
			return
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

func (s *Stream) dispatch(packet danmakuPacket) error {
	switch packet.Operation {
	case danmakuOpAuthReply:
		return validateAuthReply(packet.Body)
	case danmakuOpMessage:
		event := Event{
			Cmd:  commandName(danmakuCommandOf(packet.Body)),
			Raw:  append([]byte(nil), packet.Body...),
			Room: s.room,
		}
		select {
		case s.events <- event:
		case <-s.done:
		case <-s.ctx.Done():
		}
	}
	return nil
}

func validateAuthPacket(packets []danmakuPacket) error {
	for _, packet := range packets {
		if packet.Operation != danmakuOpAuthReply {
			continue
		}
		return validateAuthReply(packet.Body)
	}
	return errDanmakuAuthReplyMissing
}

// validateAuthReply 把认证业务码映射为可判定错误，供重连逻辑选择刷新 token 或切换节点。
// 非 JSON 的认证回复按接受处理，避免服务端变更回复形态时误判断线。
func validateAuthReply(body []byte) error {
	var reply struct {
		Code int `json:"code"`
	}
	if err := json.Unmarshal(body, &reply); err != nil {
		return nil //nolint:nilerr // 非 JSON 认证回复按接受处理，服务端形态变化不应断开连接
	}
	switch reply.Code {
	case 0:
		return nil
	case danmakuAuthCodeTokenExpired:
		return errDanmakuAuthTokenExpired
	case danmakuAuthCodeRiskControl, -danmakuAuthCodeRiskControl:
		return errDanmakuAuthRiskControl
	default:
		return fmt.Errorf("%w: code %d", errDanmakuAuthRejected, reply.Code)
	}
}

func (s *Stream) setConn(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.conn = conn
	if conn != nil && s.ctx.Err() != nil {
		_ = conn.Close()
	}
}

func (s *Stream) closeConn() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.conn != nil {
		_ = s.conn.Close()
		s.conn = nil
	}
}

func (s *Stream) setError(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closeErr == nil {
		s.closeErr = err
	}
}
