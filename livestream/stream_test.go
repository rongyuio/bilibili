package livestream

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/rongyuio/bilibili"
)

func TestConnectLiveDanmakuStream(t *testing.T) {
	var gotAuth danmakuAuthBody
	var gotHeartbeats int
	upgrader := websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/Room/room_init"):
			writeTestEnvelope(w, `{"room_id":23058,"short_id":1}`)
		case strings.HasSuffix(r.URL.Path, "/index/getDanmuInfo"):
			writeTestEnvelope(w, `{"token":"token","host_list":[{"host":"example.test","port":2243,"ws_port":2244,"wss_port":2245}]}`)
		case strings.HasSuffix(r.URL.Path, "/frontend/finger/spi"):
			writeTestEnvelope(w, `{"b_3":"buvid3","b_4":"buvid4"}`)
		case r.URL.Path == "/x/web-interface/nav":
			writeWbiNav(t, w)
		case r.URL.Path == "/sub" && websocket.IsWebSocketUpgrade(r):
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				t.Errorf("WebSocket 升级失败: %v", err)
				return
			}
			defer func() { _ = conn.Close() }()
			for {
				_, packet, err := conn.ReadMessage()
				if err != nil {
					return
				}
				decoded, err := decodeDanmakuPacket(packet)
				if err != nil {
					t.Errorf("客户端帧不合法: %v", err)
					return
				}
				switch decoded.Operation {
				case danmakuOpAuth:
					if err := json.Unmarshal(decoded.Body, &gotAuth); err != nil {
						t.Errorf("认证包不是 JSON: %v", err)
						return
					}
					_ = conn.WriteMessage(websocket.BinaryMessage,
						encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpAuthReply, []byte(`{"code":0}`)))
				case danmakuOpHeartbeat:
					gotHeartbeats++
					_ = conn.WriteMessage(websocket.BinaryMessage,
						encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpHeartbeatReply, []byte(`[1]`)))
					_ = conn.WriteMessage(websocket.BinaryMessage,
						encodeDanmakuPacket(danmakuProtoPlain, danmakuOpMessage, []byte(`{"cmd":"DANMU_MSG","info":["hello"]}`)))
				}
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	root := bilibili.NewWithClient(httpClientForTest(t, server.URL))
	root.SetCookiesString("DedeUserID=42")
	client := New(root)
	useWS := false
	stream, err := client.Connect(context.Background(), ConnectOptions{
		RoomID:            1,
		UseTLS:            &useWS,
		AuthProtover:      protoverPointer(danmakuProtoPlain),
		HeartbeatInterval: 40 * time.Millisecond,
		Dialer:            testWSDialer(t, server.URL),
	})
	if err != nil {
		t.Fatalf("Connect 返回错误: %v", err)
	}
	defer func() { _ = stream.Close() }()

	event := <-stream.Events()
	if event.Cmd != "DANMU_MSG" || event.Room != 23058 {
		t.Fatalf("事件结果不正确: %+v", event)
	}
	var payload map[string]any
	if err := json.Unmarshal(event.Raw, &payload); err != nil {
		t.Fatalf("Raw 不是 JSON: %v", err)
	}
	time.Sleep(140 * time.Millisecond)
	_ = stream.Close()
	<-stream.Done()
	if gotAuth.RoomID != 23058 || gotAuth.UID != 42 || gotAuth.Key != "token" {
		t.Fatalf("服务端收到的认证包不正确: %+v", gotAuth)
	}
	if gotHeartbeats < 2 {
		t.Fatalf("心跳次数不足: %d", gotHeartbeats)
	}
	if err := stream.Err(); err != nil && !errors.Is(err, websocket.ErrCloseSent) {
		t.Fatalf("关闭后的错误不符合预期: %v", err)
	}
}

func writeTestEnvelope(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"code":0,"message":"0","data":` + data + `}`))
}

func writeWbiNav(t *testing.T, w http.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"code":0,"message":"0","data":{"wbi_img":{` +
		`"img_url":"https://i0.hdslb.com/bfs/wbi/7cd084941338484aae1ad9425b84077c.png",` +
		`"sub_url":"https://i0.hdslb.com/bfs/wbi/4932caff0ff746eab6f01bf08f87e641.png"}}}`))
}

func httpClientForTest(t *testing.T, endpoint string) *resty.Client {
	t.Helper()
	target, err := url.Parse(endpoint)
	if err != nil {
		t.Fatalf("解析测试服务地址失败: %v", err)
	}
	client := resty.New().SetTransport(&danmakuRoundTripper{target: target})
	return client
}

// danmakuRoundTripper 把方法内的线上 URL 指向本地测试服务，保留路径与 query。
type danmakuRoundTripper struct {
	target *url.URL
}

func (t *danmakuRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.URL.Scheme = t.target.Scheme
	clone.URL.Host = t.target.Host
	clone.Host = t.target.Host
	return http.DefaultTransport.RoundTrip(clone)
}

type testDialer struct {
	address string
	dialer  *websocket.Dialer
}

func (d testDialer) DialContext(ctx context.Context, _, _ string, header http.Header) (*websocket.Conn, *http.Response, error) {
	return d.dialer.DialContext(ctx, d.address, header)
}

func testWSDialer(t *testing.T, endpoint string) WebSocketDialer {
	t.Helper()
	dialer := *websocket.DefaultDialer
	return testDialer{
		address: "ws" + strings.TrimPrefix(endpoint, "http") + "/sub",
		dialer:  &dialer,
	}
}

// protoverPointer 便于测试指定认证协议版本。
func protoverPointer(value uint16) *uint16 { return &value }
