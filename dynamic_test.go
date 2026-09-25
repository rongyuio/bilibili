package bilibili

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestNormalizeDynamicFeedAllParamDefaults(t *testing.T) {
	got := normalizeDynamicFeedAllParam(GetDynamicFeedAllParam{})

	if got.FeedType != "all" {
		t.Errorf("FeedType = %q，期望 all", got.FeedType)
	}
	if got.TimezoneOffset != -480 {
		t.Errorf("TimezoneOffset = %d，期望 -480", got.TimezoneOffset)
	}
	if got.Features != dynamicFeedAllFeatures {
		t.Errorf("Features = %q，期望默认开关串", got.Features)
	}
	if got.WebLocation != dynamicWebLocation {
		t.Errorf("WebLocation = %q，期望 %q", got.WebLocation, dynamicWebLocation)
	}
}

func TestNormalizeDynamicFeedAllParamKeepsExplicitValues(t *testing.T) {
	want := GetDynamicFeedAllParam{
		FeedType:       "video",
		Offset:         "123",
		TimezoneOffset: 480,
		Features:       "custom",
		WebLocation:    "1.2.3",
	}

	if got := normalizeDynamicFeedAllParam(want); got != want {
		t.Errorf("显式给的值被覆盖: got %+v，want %+v", got, want)
	}
}

// 关注流的参数名靠 json 标签纠正：字段名 FeedType 推导出来是 feed_type，
// 而接口要的是 type。这条用例钉住标签没被写错。
func TestDynamicFeedAllParamEncoding(t *testing.T) {
	r := resty.New().R()
	if err := withParams(r, normalizeDynamicFeedAllParam(GetDynamicFeedAllParam{Offset: "123"})); err != nil {
		t.Fatalf("withParams 失败: %v", err)
	}

	want := map[string]string{
		"type":            "all",
		"offset":          "123",
		"timezone_offset": "-480",
		"features":        dynamicFeedAllFeatures,
		"web_location":    dynamicWebLocation,
	}
	for key, value := range want {
		if got := r.QueryParam.Get(key); got != value {
			t.Errorf("query %s = %q，期望 %q", key, got, value)
		}
	}
	if _, ok := r.QueryParam["feed_type"]; ok {
		t.Errorf("出现了不该有的 feed_type 参数: %v", r.QueryParam)
	}
}

func TestRepostDynamicRejectsInvalidDynamicID(t *testing.T) {
	// 参数校验发生在发请求之前，所以这里不会联网。
	for _, id := range []int64{0, -1} {
		_, err := New().RepostDynamic(context.Background(), RepostDynamicParam{DynamicID: id})
		var paramErr *ParamError
		if !errors.As(err, &paramErr) {
			t.Fatalf("DynamicID=%d：期望 *ParamError，实际 %v", id, err)
		}
	}
}

func TestRepostDynamicHandlerRequiresLogin(t *testing.T) {
	r := resty.New().R().SetCookies([]*http.Cookie{{Name: "bili_jct", Value: "csrf"}})
	if err := repostDynamicHandler(RepostDynamicParam{DynamicID: 1})(r); err == nil {
		t.Fatal("缺少 DedeUserID 时期望返回错误")
	}
}

func TestRepostDynamicHandlerBody(t *testing.T) {
	const (
		mid       = "10001"
		dynamicID = int64(755402172521250838)
	)

	r := resty.New().R().SetCookies([]*http.Cookie{
		{Name: "DedeUserID", Value: mid},
		{Name: "bili_jct", Value: "csrf"},
	})

	before := time.Now().Unix()
	if err := repostDynamicHandler(RepostDynamicParam{DynamicID: dynamicID, Content: "转发文案"})(r); err != nil {
		t.Fatalf("handler 失败: %v", err)
	}
	after := time.Now().Unix()

	if got := r.QueryParam.Get("platform"); got != "web" {
		t.Errorf("query platform = %q，期望 web", got)
	}
	if got := r.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q，期望 application/json", got)
	}

	var body struct {
		DynReq struct {
			Content struct {
				Contents []struct {
					RawText string `json:"raw_text"`
					Type    int    `json:"type"`
					BizID   string `json:"biz_id"`
				} `json:"contents"`
			} `json:"content"`
			Scene    int    `json:"scene"`
			UploadID string `json:"upload_id"`
			Meta     struct {
				AppMeta struct {
					From    string `json:"from"`
					MobiApp string `json:"mobi_app"`
				} `json:"app_meta"`
			} `json:"meta"`
			Option struct {
				Aigc int `json:"aigc"`
			} `json:"option"`
			RepostSrc json.RawMessage `json:"repost_src"`
		} `json:"dyn_req"`
		WebRepostSrc struct {
			DynIDStr string `json:"dyn_id_str"`
		} `json:"web_repost_src"`
	}

	raw, ok := r.Body.([]byte)
	if !ok {
		t.Fatalf("body 类型 = %T，期望 []byte", r.Body)
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("body 不是合法 JSON: %v\n%s", err, raw)
	}

	if body.DynReq.Scene != dynamicRepostScene {
		t.Errorf("scene = %d，期望 %d", body.DynReq.Scene, dynamicRepostScene)
	}
	// 原动态 ID 必须在外层；放在 dyn_req.repost_src 里服务端不认。
	if got := body.WebRepostSrc.DynIDStr; got != strconv.FormatInt(dynamicID, 10) {
		t.Errorf("web_repost_src.dyn_id_str = %q，期望 %d", got, dynamicID)
	}
	if len(body.DynReq.RepostSrc) != 0 {
		t.Errorf("dyn_req.repost_src 不该出现: %s", body.DynReq.RepostSrc)
	}

	if len(body.DynReq.Content.Contents) != 1 {
		t.Fatalf("contents 长度 = %d，期望 1", len(body.DynReq.Content.Contents))
	}
	node := body.DynReq.Content.Contents[0]
	if node.RawText != "转发文案" || node.Type != 1 || node.BizID != "" {
		t.Errorf("文本节点不符: %+v", node)
	}

	if body.DynReq.Meta.AppMeta.From != "create.dynamic.web" || body.DynReq.Meta.AppMeta.MobiApp != "web" {
		t.Errorf("meta 不符: %+v", body.DynReq.Meta.AppMeta)
	}
	if body.DynReq.Option.Aigc != 2 {
		t.Errorf("option.aigc = %d，期望 2", body.DynReq.Option.Aigc)
	}

	// upload_id 形如「mid_秒级时间戳_四位随机」
	parts := strings.Split(body.DynReq.UploadID, "_")
	if len(parts) != 3 {
		t.Fatalf("upload_id 格式不对: %q", body.DynReq.UploadID)
	}
	if parts[0] != mid {
		t.Errorf("upload_id 里的 mid = %q，期望 %q", parts[0], mid)
	}
	ts, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		t.Fatalf("upload_id 时间戳不是数字: %q", parts[1])
	}
	if ts < before || ts > after {
		t.Errorf("upload_id 时间戳 %d 不在 [%d, %d] 内", ts, before, after)
	}
	n, err := strconv.Atoi(parts[2])
	if err != nil || n < 0 || n > 9999 {
		t.Errorf("upload_id 随机段不合法: %q", parts[2])
	}
}

// 没有文案时 contents 必须是空数组：nil 切片会被序列化成 null。
func TestRepostDynamicHandlerEmptyContent(t *testing.T) {
	r := resty.New().R().SetCookies([]*http.Cookie{{Name: "DedeUserID", Value: "1"}})
	if err := repostDynamicHandler(RepostDynamicParam{DynamicID: 1})(r); err != nil {
		t.Fatalf("handler 失败: %v", err)
	}

	raw := string(r.Body.([]byte))
	if !strings.Contains(raw, `"contents":[]`) {
		t.Fatalf("空文案时 contents 不是空数组: %s", raw)
	}
}

func TestRepostDynamicHandlerKeepsCallerUploadID(t *testing.T) {
	r := resty.New().R().SetCookies([]*http.Cookie{{Name: "DedeUserID", Value: "1"}})
	if err := repostDynamicHandler(RepostDynamicParam{DynamicID: 1, UploadID: "fixed_1_2"})(r); err != nil {
		t.Fatalf("handler 失败: %v", err)
	}

	if raw := string(r.Body.([]byte)); !strings.Contains(raw, `"upload_id":"fixed_1_2"`) {
		t.Fatalf("调用方给的 upload_id 没被采用: %s", raw)
	}
}

func TestNewDynamicUploadIDFormat(t *testing.T) {
	before := time.Now().Unix()
	got := newDynamicUploadID("10001")
	after := time.Now().Unix()

	parts := strings.Split(got, "_")
	if len(parts) != 3 {
		t.Fatalf("upload_id 格式不对: %q", got)
	}
	if parts[0] != "10001" {
		t.Errorf("mid = %q，期望 10001", parts[0])
	}
	ts, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || ts < before || ts > after {
		t.Errorf("时间戳 %q 不在 [%d, %d] 内", parts[1], before, after)
	}
	n, err := strconv.Atoi(parts[2])
	if err != nil || n < 0 || n > 9999 {
		t.Errorf("随机段不合法: %q", parts[2])
	}
}
