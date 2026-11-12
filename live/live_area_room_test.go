package live

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

func TestGetLiveAreaRoomListParamEncoding(t *testing.T) {
	var query url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"success","data":[{"roomid":545068,"area_v2_id":86}]}`))
	}))
	defer server.Close()

	client := liveAreaRoomClientForTest(t, server.URL)
	rooms, err := client.GetLiveAreaRoomList(context.Background(), GetLiveAreaRoomListParam{
		ParentAreaID: 2,
		AreaID:       86,
	})
	if err != nil {
		t.Fatalf("GetLiveAreaRoomList 返回错误: %v", err)
	}
	if len(rooms) != 1 || rooms[0].RoomID != 545068 || rooms[0].AreaV2ID != 86 {
		t.Fatalf("解码结果不正确: %+v", rooms)
	}
	want := map[string]string{
		"parent_area_id": "2",
		"area_id":        "86",
		"sort_type":      "online",
		"page":           "1",
		"page_size":      "20",
	}
	for key, value := range want {
		if got := query.Get(key); got != value {
			t.Errorf("query %s = %q，期望 %q", key, got, value)
		}
	}
	if len(query) != len(want) {
		t.Errorf("query 包含额外字段: %v", query)
	}
}

func TestGetLiveAreaRoomListExplicitValues(t *testing.T) {
	var query url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"message":"success","data":[]}`))
	}))
	defer server.Close()

	client := liveAreaRoomClientForTest(t, server.URL)
	rooms, err := client.GetLiveAreaRoomList(context.Background(), GetLiveAreaRoomListParam{
		ParentAreaID: 2,
		AreaID:       0,
		SortType:     "dynamic",
		Page:         3,
		PageSize:     30,
	})
	if err != nil {
		t.Fatalf("GetLiveAreaRoomList 返回错误: %v", err)
	}
	if len(rooms) != 0 {
		t.Fatalf("空数组应解码为空切片: %+v", rooms)
	}
	for key, value := range map[string]string{
		"parent_area_id": "2",
		"area_id":        "0",
		"sort_type":      "dynamic",
		"page":           "3",
		"page_size":      "30",
	} {
		if got := query.Get(key); got != value {
			t.Errorf("query %s = %q，期望 %q", key, got, value)
		}
	}
}

func TestGetLiveAreaRoomListBusinessError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":404,"message":"这个分区不存在啦"}`))
	}))
	defer server.Close()

	client := liveAreaRoomClientForTest(t, server.URL)
	rooms, err := client.GetLiveAreaRoomList(context.Background(), GetLiveAreaRoomListParam{
		ParentAreaID: 999999,
		AreaID:       999999,
	})
	if rooms != nil {
		t.Fatalf("业务错误时不应返回部分结果: %+v", rooms)
	}
	var apiErr bilibili.Error
	if !errors.As(err, &apiErr) || apiErr.Code != 404 {
		t.Fatalf("错误类型或业务码不正确: %v", err)
	}
}

func liveAreaRoomClientForTest(t *testing.T, endpoint string) *Client {
	t.Helper()
	target, err := url.Parse(endpoint)
	if err != nil {
		t.Fatalf("解析测试服务地址失败: %v", err)
	}
	root := bilibili.NewWithClient(resty.New().SetTransport(&liveAreaRoomRoundTripper{target: target}))
	return New(root)
}

// liveAreaRoomRoundTripper 把方法内的线上 URL 指向本地测试服务，保留路径与 query。
type liveAreaRoomRoundTripper struct {
	target *url.URL
}

func (t *liveAreaRoomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.URL.Scheme = t.target.Scheme
	clone.URL.Host = t.target.Host
	clone.Host = t.target.Host
	return http.DefaultTransport.RoundTrip(clone)
}
