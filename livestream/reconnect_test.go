package livestream

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rongyuio/bilibili"
)

// 连接层测试：心跳体、断线重连、认证码分流与关闭语义。
// 全部使用本地 httptest，不访问真实 B 站接口。

type reconnectTestServer struct {
	mu          sync.Mutex
	upgrader    websocket.Upgrader
	replies     []int
	connCount   int
	danmuInfo   int
	authBodies  []danmakuAuthBody
	heartbeat   []string
	t           *testing.T
	dropAfterNo int
}

func newReconnectTestServer(t *testing.T, replies ...int) *reconnectTestServer {
	return &reconnectTestServer{
		upgrader: websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }},
		replies:  replies,
		t:        t,
	}
}

func (s *reconnectTestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasSuffix(r.URL.Path, "/Room/room_init"):
		writeTestEnvelope(w, `{"room_id":23058,"short_id":1}`)
	case strings.HasSuffix(r.URL.Path, "/index/getDanmuInfo"):
		s.mu.Lock()
		s.danmuInfo++
		token := "token-" + itoa(s.danmuInfo)
		s.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"0","data":{"token":"` + token +
			`","host_list":[{"host":"example.test","port":2243,"ws_port":2244,"wss_port":2245}]}}`))
	case strings.HasSuffix(r.URL.Path, "/frontend/finger/spi"):
		writeTestEnvelope(w, `{"b_3":"buvid3","b_4":"buvid4"}`)
	case r.URL.Path == "/x/web-interface/nav":
		writeWbiNav(s.t, w)
	case strings.HasSuffix(r.URL.Path, "/sub"):
		s.serveWebSocket(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (s *reconnectTestServer) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer func() { _ = conn.Close() }()
	s.mu.Lock()
	index := s.connCount
	s.connCount++
	code := 0
	if index < len(s.replies) {
		code = s.replies[index]
	}
	s.mu.Unlock()
	for {
		_, packet, err := conn.ReadMessage()
		if err != nil {
			return
		}
		decoded, err := decodeDanmakuPacket(packet)
		if err != nil {
			return
		}
		switch decoded.Operation {
		case danmakuOpAuth:
			s.handleTestAuth(conn, decoded, code)
			if code != 0 {
				return
			}
		case danmakuOpHeartbeat:
			s.handleTestHeartbeat(conn, decoded)
			return
		}
	}
}

func (s *reconnectTestServer) handleTestAuth(conn *websocket.Conn, packet danmakuPacket, code int) {
	var body danmakuAuthBody
	_ = json.Unmarshal(packet.Body, &body)
	s.mu.Lock()
	s.authBodies = append(s.authBodies, body)
	s.mu.Unlock()
	reply, _ := json.Marshal(map[string]int{"code": code})
	if err := conn.WriteMessage(websocket.BinaryMessage,
		encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpAuthReply, reply)); err != nil {
		return
	}
	if code != 0 {
		return
	}
	_ = conn.WriteMessage(websocket.BinaryMessage,
		encodeDanmakuPacket(danmakuProtoPlain, danmakuOpMessage,
			[]byte(`{"cmd":"DANMU_MSG","info":[[0,1],"hello",[42,"ry"]]}`)))
}

func (s *reconnectTestServer) handleTestHeartbeat(conn *websocket.Conn, packet danmakuPacket) {
	s.mu.Lock()
	s.heartbeat = append(s.heartbeat, string(packet.Body))
	s.mu.Unlock()
	_ = conn.WriteMessage(websocket.BinaryMessage,
		encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpHeartbeatReply, []byte(`[1]`)))
}

func (s *reconnectTestServer) snapshot() (int, int, []danmakuAuthBody, []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.connCount, s.danmuInfo, append([]danmakuAuthBody(nil), s.authBodies...), append([]string(nil), s.heartbeat...)
}

func TestConnectHeartbeatAndReconnect(t *testing.T) {
	server := newReconnectTestServer(t, 0, 0, 0)
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	stream, err := connectTestStream(t, httpServer.URL, time.Hour, 80*time.Millisecond)
	if err != nil {
		t.Fatalf("Connect 返回错误: %v", err)
	}
	defer func() { _ = stream.Close() }()
	waitTestEvent(t, stream)
	waitTestReconnects(t, server, 2)
	_, _, auths, beats := server.snapshot()
	if beats[0] != "{}" {
		t.Fatalf("心跳体 = %q，期望空对象", beats[0])
	}
	for _, body := range auths {
		if body.Protover != danmakuProtoPlain || body.RoomID != 23058 || body.UID != 42 {
			t.Fatalf("认证包不正确: %+v", body)
		}
	}
	closeTestStream(t, stream)
}

func TestConnectAuthTokenExpiredOnReconnect(t *testing.T) {
	server := newReconnectTestServer(t, 0, danmakuAuthCodeTokenExpired, 0)
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	stream, err := connectTestStream(t, httpServer.URL, time.Hour, 60*time.Millisecond)
	if err != nil {
		t.Fatalf("Connect 返回错误: %v", err)
	}
	defer func() { _ = stream.Close() }()
	waitTestEvent(t, stream)
	waitTestReconnects(t, server, 3)
	_, danmuInfo, auths, _ := server.snapshot()
	if danmuInfo < 2 {
		t.Fatalf("token 过期后应重新拉取 getDanmuInfo，实际次数 %d", danmuInfo)
	}
	if len(auths) >= 3 && auths[1].Key == auths[2].Key {
		t.Fatalf("重连应使用刷新后的 token: %q", auths[2].Key)
	}
	closeTestStream(t, stream)
}

func TestConnectAuthRejectedOnFirstConnection(t *testing.T) {
	server := newReconnectTestServer(t, danmakuAuthCodeTokenExpired)
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	_, err := connectTestStream(t, httpServer.URL, time.Hour, time.Second)
	if !errors.Is(err, errDanmakuAuthTokenExpired) {
		t.Fatalf("首次认证失败的错误 = %v，期望 token 过期错误", err)
	}
}

func TestConnectTokenRefreshLimit(t *testing.T) {
	replies := []int{0}
	for range danmakuMaxTokenRefresh + 2 {
		replies = append(replies, danmakuAuthCodeTokenExpired)
	}
	server := newReconnectTestServer(t, replies...)
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	stream, err := connectTestStream(t, httpServer.URL, time.Hour, 40*time.Millisecond)
	if err != nil {
		t.Fatalf("Connect 返回错误: %v", err)
	}
	defer func() { _ = stream.Close() }()
	select {
	case <-stream.Done():
	case <-time.After(5 * time.Second):
		t.Fatal("token 持续过期时应结束信息流")
	}
	if err := stream.Err(); !errors.Is(err, errDanmakuAuthTokenExpired) {
		t.Fatalf("结束原因 = %v，期望 token 过期错误", err)
	}
}

func connectTestStream(t *testing.T, endpoint string, heartbeat, idle time.Duration) (*Stream, error) {
	t.Helper()
	root := bilibili.NewWithClient(httpClientForTest(t, endpoint))
	root.SetCookiesString("DedeUserID=42")
	return New(root).Connect(context.Background(), ConnectOptions{
		RoomID:            1,
		UseTLS:            boolPointer(false),
		AuthProtover:      protoverPointer(danmakuProtoPlain),
		HeartbeatInterval: heartbeat,
		ReadIdleTimeout:   idle,
		ReconnectDelay:    5 * time.Millisecond,
		JoinTimeout:       time.Second,
		Dialer:            testWSDialer(t, endpoint),
	})
}

func waitTestEvent(t *testing.T, stream *Stream) {
	t.Helper()
	select {
	case event := <-stream.Events():
		if event.Cmd != "DANMU_MSG" || event.Room != 23058 {
			t.Fatalf("事件结果不正确: %+v", event)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("等待事件超时")
	}
}

func waitTestReconnects(t *testing.T, server *reconnectTestServer, want int) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for {
		count, _, _, _ := server.snapshot()
		if count >= want {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("重连次数不足: %d，期望 %d", count, want)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func closeTestStream(t *testing.T, stream *Stream) {
	t.Helper()
	_ = stream.Close()
	select {
	case <-stream.Done():
	case <-time.After(time.Second):
		t.Fatal("Done 未关闭")
	}
	// Close 后 channel 可能还剩缓冲事件，继续读取直到关闭。
	deadline := time.After(time.Second)
	for {
		select {
		case _, open := <-stream.Events():
			if !open {
				return
			}
		case <-deadline:
			t.Fatal("关闭后 Events 仍未关闭")
		}
	}
}

func boolPointer(value bool) *bool { return &value }

func itoa(value int) string {
	return strconv.Itoa(value)
}
