package live

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// sendLiveDanmakuOKResponse 是发送弹幕成功的标准响应。
const sendLiveDanmakuOKResponse = `{"code":0,"message":"0","data":{"mode_info":{"mode":1,"show_player_type":0,"extra":""}}}`

// sendLiveDanmakuCapture 记录测试服务收到的表单与 query。
type sendLiveDanmakuCapture struct {
	mu    sync.Mutex
	form  url.Values
	query url.Values
}

func (c *sendLiveDanmakuCapture) record(form, query url.Values) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.form = form
	c.query = query
}

func (c *sendLiveDanmakuCapture) snapshot() (url.Values, url.Values) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.form, c.query
}

func TestSendLiveDanmakuDmType(t *testing.T) {
	tests := []struct {
		name       string
		dmType     int
		wantDmType string // 空字符串表示表单里不应出现 dm_type
	}{
		{name: "文本弹幕不带 dm_type", dmType: 0, wantDmType: ""},
		{name: "表情弹幕带 dm_type=1", dmType: 1, wantDmType: "1"},
		{name: "非 1 的类型保留原值", dmType: 2, wantDmType: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capture := &sendLiveDanmakuCapture{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					t.Errorf("解析表单失败: %v", err)
				}
				capture.record(r.PostForm, r.URL.Query())
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(sendLiveDanmakuOKResponse))
			}))
			defer server.Close()

			client := sendLiveDanmakuClientForTest(t, server.URL)
			result, err := client.SendLiveDanmaku(context.Background(), SendLiveDanmakuParam{
				RoomID: 196,
				Msg:    "测试弹幕",
				DmType: tt.dmType,
			})
			if err != nil {
				t.Fatalf("SendLiveDanmaku 返回错误: %v", err)
			}
			if result == nil || result.ModeInfo.Mode != 1 {
				t.Fatalf("解码结果不正确: %+v", result)
			}

			form, query := capture.snapshot()
			if form == nil {
				t.Fatal("未捕获到请求表单")
			}
			got, present := form["dm_type"]
			switch {
			case tt.wantDmType == "" && present:
				t.Fatalf("文本弹幕表单不应包含 dm_type: %v", form)
			case tt.wantDmType != "" && (len(got) != 1 || got[0] != tt.wantDmType):
				t.Fatalf("dm_type = %v，期望 %q", got, tt.wantDmType)
			}
			if len(query) != 0 {
				t.Errorf("表单字段应全部从 query 移出，残留: %v", query)
			}

			want := map[string]string{
				"msg":        "测试弹幕",
				"roomid":     "196",
				"bubble":     "0",
				"color":      "16777215",
				"fontsize":   "25",
				"mode":       "1",
				"rnd":        "1",
				"csrf":       "test-csrf",
				"csrf_token": "test-csrf",
			}
			for key, value := range want {
				if gotValue := form.Get(key); gotValue != value {
					t.Errorf("表单 %s = %q，期望 %q", key, gotValue, value)
				}
			}
			wantLen := len(want)
			if tt.wantDmType != "" {
				wantLen++
			}
			if len(form) != wantLen {
				t.Errorf("表单字段数量不符: %v", form)
			}
		})
	}
}

func sendLiveDanmakuClientForTest(t *testing.T, endpoint string) *Client {
	t.Helper()
	target, err := url.Parse(endpoint)
	if err != nil {
		t.Fatalf("解析测试服务地址失败: %v", err)
	}
	root := bilibili.NewWithClient(resty.New().SetTransport(&sendLiveDanmakuRoundTripper{target: target}))
	root.SetCookie(&http.Cookie{Name: "bili_jct", Value: "test-csrf"})
	return New(root)
}

// sendLiveDanmakuRoundTripper 把方法内的线上 URL 指向本地测试服务，保留路径与表单。
type sendLiveDanmakuRoundTripper struct {
	target *url.URL
}

func (t *sendLiveDanmakuRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.URL.Scheme = t.target.Scheme
	clone.URL.Host = t.target.Host
	clone.Host = t.target.Host
	return http.DefaultTransport.RoundTrip(clone)
}
