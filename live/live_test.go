package live

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// sendLiveDanmakuOKResponse 是发送弹幕成功的标准响应。
const sendLiveDanmakuOKResponse = `{"code":0,"message":"0","data":{"mode_info":{"mode":1,"show_player_type":0,"extra":""}}}`

// sendLiveDanmakuNavResponse 提供 WBI 签名所需的固定密钥。
const sendLiveDanmakuNavResponse = `{"code":0,"message":"0","data":{"wbi_img":{` +
	`"img_url":"https://i0.hdslb.com/bfs/wbi/7cd084941338484aae1ad9425b84077c.png",` +
	`"sub_url":"https://i0.hdslb.com/bfs/wbi/4932caff0ff746eab6f01bf08f87e641.png"}}}`

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
		dataExtend string
		rnd        int64
		wantDmType string
		wantData   string
	}{
		{name: "文本弹幕", wantData: "{}"},
		{name: "表情弹幕", dmType: 1, wantDmType: "1", wantData: "{}"},
		{name: "保留非 1 类型", dmType: 2, wantDmType: "2", wantData: "{}"},
		{name: "自定义附加上报", dataExtend: `{"room_category":"1"}`, wantData: `{"room_category":"1"}`},
		{name: "固定随机数", rnd: 1234567890, wantData: "{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capture := &sendLiveDanmakuCapture{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.URL.Path == "/x/web-interface/nav" {
					_, _ = w.Write([]byte(sendLiveDanmakuNavResponse))
					return
				}
				if err := r.ParseForm(); err != nil {
					t.Errorf("解析表单失败: %v", err)
				}
				capture.record(r.PostForm, r.URL.Query())
				_, _ = w.Write([]byte(sendLiveDanmakuOKResponse))
			}))
			defer server.Close()

			client := sendLiveDanmakuClientForTest(t, server.URL)
			result, err := client.SendLiveDanmaku(context.Background(), SendLiveDanmakuParam{
				RoomID: 196, Msg: "测试弹幕", DmType: tt.dmType,
				DataExtend: tt.dataExtend, Rnd: tt.rnd,
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
			if query.Get("web_location") != liveDanmakuWebLocation ||
				query.Get("w_rid") == "" || query.Get("wts") == "" {
				t.Fatalf("query 缺少 web_location 或 WBI 签名: %v", query)
			}
			if len(query) != 3 {
				t.Errorf("query 只应包含 web_location 与 WBI 字段: %v", query)
			}

			want := map[string]string{
				"msg": "测试弹幕", "roomid": "196", "bubble": "0",
				"color": "16777215", "fontsize": "25", "mode": "1",
				"data_extend": tt.wantData,
				"csrf":        "test-csrf", "csrf_token": "test-csrf",
			}
			for key, value := range want {
				if gotValue := form.Get(key); gotValue != value {
					t.Errorf("表单 %s = %q，期望 %q", key, gotValue, value)
				}
			}
			if tt.rnd != 0 {
				if gotValue := form.Get("rnd"); gotValue != strconv.FormatInt(tt.rnd, 10) {
					t.Errorf("rnd = %q，期望 %d", gotValue, tt.rnd)
				}
			} else if rnd, err := strconv.ParseInt(form.Get("rnd"), 10, 64); err != nil ||
				rnd < time.Now().Add(-time.Minute).Unix() {
				t.Errorf("默认 rnd 不是当前 Unix 秒: %q", form.Get("rnd"))
			}
			wantLen := len(want) + 1
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
