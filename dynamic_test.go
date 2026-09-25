package bilibili

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
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

// ── 响应解码（按接口真实形状）──────────────────────────────────────

// 关注流的响应里 id_str 是**带引号的字符串**（见 all.md 的字段表），而 DynamicItem.IDStr
// 声明的是 json.Number —— 两者能不能对上，只能拿真实形状试。
//
// 对不上的后果很隐蔽：解码会走容错剪枝、把 id_str 变成空，于是「每条动态都拿不到 id」、
// 整条转发链路静默不干活，而接口全都回 200。
func TestDecodeDynamicFeedAllResponse(t *testing.T) {
	const body = `{
		"code": 0, "message": "0", "ttl": 1,
		"data": {
			"has_more": true,
			"offset": "755402172521250838",
			"update_baseline": "755402172521250839",
			"update_num": 2,
			"items": [{
				"id_str": "755402172521250838",
				"type": "DYNAMIC_TYPE_DRAW",
				"visible": true,
				"modules": {
					"module_author": { "mid": 3494379073309365, "name": "无限大" },
					"module_dynamic": { "desc": { "text": "▼转发本条动态+关注" } }
				}
			}]
		}
	}`

	var out *DynamicInfo
	err := decodeResponse("GET", "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/all",
		[]byte(body), &out, nil)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}
	if out == nil || len(out.Items) != 1 {
		t.Fatalf("条目数不对: %+v", out)
	}
	if got := out.Items[0].IDStr.String(); got != "755402172521250838" {
		t.Errorf("id_str = %q，期望 755402172521250838", got)
	}
	if !out.HasMore || out.Offset != "755402172521250838" {
		t.Errorf("分页字段不对: has_more=%v offset=%q", out.HasMore, out.Offset)
	}
	if mid, err := out.Items[0].Modules.ModuleAuthor.Mid.Int64(); err != nil || mid != 3494379073309365 {
		t.Errorf("作者 mid = %d err=%v，期望 3494379073309365", mid, err)
	}
}

// 转发接口的响应形状。示例里还带一个 share_window（我们没建模）—— 宽松解码要能吃掉它，
// 不能因为多了个字段就整条解不出来。
func TestDecodeRepostDynamicResult(t *testing.T) {
	const body = `{
		"code": 0, "message": "0", "ttl": 1,
		"data": {
			"dyn_id": 755402172521250838,
			"dyn_id_str": "755402172521250838",
			"dyn_type": 1,
			"dyn_rid": 221621773,
			"share_window": { "main_title": "分享后会获得更多曝光，快去分享吧" }
		}
	}`

	var out *RepostDynamicResult
	err := decodeResponse("POST", "https://api.bilibili.com/x/dynamic/feed/create/dyn",
		[]byte(body), &out, nil)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}
	if out.DynID != 755402172521250838 || out.DynIDStr != "755402172521250838" {
		t.Errorf("动态 id 不对: %+v", out)
	}
	if out.DynType != 1 {
		t.Errorf("dyn_type = %d，期望 1（转发）", out.DynType)
	}
	if out.DynRid != 221621773 {
		t.Errorf("dyn_rid = %d，期望 221621773", out.DynRid)
	}
}

// 业务码非 0 → bilibili.Error。调用方靠 errors.As 分流：业务错误表示「这条动态不是抽奖」
// 或「登录过期」，都不是网络故障 —— 别把它和 HTTPError 混了。
func TestRepostDynamicBusinessError(t *testing.T) {
	const body = `{"code":-101,"message":"账号未登录","ttl":1}`

	var out *RepostDynamicResult
	err := decodeResponse("POST", "https://api.bilibili.com/x/dynamic/feed/create/dyn",
		[]byte(body), &out, nil)
	if err == nil {
		t.Fatal("业务码非 0 应当报错")
	}

	var apiErr Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("期望 bilibili.Error，实际 %T: %v", err, err)
	}
	if apiErr.Code != -101 || apiErr.Message != "账号未登录" {
		t.Errorf("错误内容不对: code=%d message=%q", apiErr.Code, apiErr.Message)
	}
}

// ── 请求形状（假 transport，不发出真实请求）────────────────────────

// roundTripFunc 让测试能在**不发出真实请求**的前提下走完整条请求链路。
//
// 库里其他用例都是纯逻辑（见 AGENTS.md 的「常用命令」一节）；这里破例，是因为
// RepostDynamic 是**不可逆的写操作** —— URL、方法、CSRF 的位置值得一次性钉死，而
// handler 单测覆盖不到前两项（URL 与方法是写死在 execute 调用里的常量）。
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// stubTransport 把客户端发出的请求换成给的那份响应，并把请求本身交给 seen。
func stubTransport(t *testing.T, c *Client, status int, body string, seen func(*http.Request, []byte)) {
	t.Helper()

	c.Resty().SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if seen != nil {
			var raw []byte
			if req.Body != nil {
				raw, _ = io.ReadAll(req.Body)
			}

			seen(req, raw)
		}

		return &http.Response{
			StatusCode: status,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(body)),
			Request:    req,
		}, nil
	}))
}

// 转发的完整请求：POST /x/dynamic/feed/create/dyn?platform=web&csrf=…，JSON body，
// 原动态 ID 在外层的 web_repost_src。
func TestRepostDynamicSendsExpectedRequest(t *testing.T) {
	const (
		csrf = "csrf-value"
		mid  = "10001"
		dyn  = int64(755402172521250838)
	)

	client := New()
	client.SetRawCookies("SESSDATA=x; bili_jct=" + csrf + "; DedeUserID=" + mid)

	var (
		gotMethod string
		gotPath   string
		gotQuery  url.Values
		gotCT     string
		gotBody   []byte
	)
	stubTransport(t, client, http.StatusOK,
		`{"code":0,"message":"0","data":{"dyn_id":1,"dyn_id_str":"1","dyn_type":1}}`,
		func(req *http.Request, raw []byte) {
			gotMethod, gotPath, gotQuery = req.Method, req.URL.Path, req.URL.Query()
			gotCT, gotBody = req.Header.Get("Content-Type"), raw
		})

	res, err := client.RepostDynamic(context.Background(), RepostDynamicParam{DynamicID: dyn})
	if err != nil {
		t.Fatalf("RepostDynamic 失败: %v", err)
	}
	if res == nil || res.DynID != 1 {
		t.Fatalf("响应没解出来: %+v", res)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("方法 = %q，期望 POST", gotMethod)
	}
	if gotPath != "/x/dynamic/feed/create/dyn" {
		t.Errorf("路径 = %q，期望 /x/dynamic/feed/create/dyn", gotPath)
	}
	if got := gotQuery.Get("csrf"); got != csrf {
		t.Errorf("query csrf = %q，期望 %q", got, csrf)
	}
	if got := gotQuery.Get("platform"); got != "web" {
		t.Errorf("query platform = %q，期望 web", got)
	}
	if gotCT != "application/json" {
		t.Errorf("Content-Type = %q，期望 application/json", gotCT)
	}

	var sent struct {
		DynReq struct {
			Scene int `json:"scene"`
		} `json:"dyn_req"`
		WebRepostSrc struct {
			DynIDStr string `json:"dyn_id_str"`
		} `json:"web_repost_src"`
	}
	if err := json.Unmarshal(gotBody, &sent); err != nil {
		t.Fatalf("发出去的 body 不是合法 JSON: %v\n%s", err, gotBody)
	}
	if sent.DynReq.Scene != dynamicRepostScene {
		t.Errorf("scene = %d，期望 %d", sent.DynReq.Scene, dynamicRepostScene)
	}
	if sent.WebRepostSrc.DynIDStr != strconv.FormatInt(dyn, 10) {
		t.Errorf("web_repost_src.dyn_id_str = %q，期望 %d", sent.WebRepostSrc.DynIDStr, dyn)
	}
}

// 关注流的完整请求：GET，platform=web 由库补上。
func TestGetDynamicFeedAllSendsExpectedRequest(t *testing.T) {
	client := New()

	var (
		gotMethod string
		gotPath   string
		gotQuery  url.Values
	)
	stubTransport(t, client, http.StatusOK,
		`{"code":0,"message":"0","data":{"has_more":false,"items":[]}}`,
		func(req *http.Request, _ []byte) {
			gotMethod, gotPath, gotQuery = req.Method, req.URL.Path, req.URL.Query()
		})

	out, err := client.GetDynamicFeedAll(context.Background(), GetDynamicFeedAllParam{})
	if err != nil {
		t.Fatalf("GetDynamicFeedAll 失败: %v", err)
	}
	if out == nil {
		t.Fatal("响应没解出来")
	}

	if gotMethod != http.MethodGet {
		t.Errorf("方法 = %q，期望 GET", gotMethod)
	}
	if gotPath != "/x/polymer/web-dynamic/v1/feed/all" {
		t.Errorf("路径 = %q，期望 /x/polymer/web-dynamic/v1/feed/all", gotPath)
	}
	for k, v := range map[string]string{
		"platform":        "web",
		"type":            "all",
		"timezone_offset": "-480",
	} {
		if got := gotQuery.Get(k); got != v {
			t.Errorf("query %s = %q，期望 %q", k, got, v)
		}
	}
}
