package bilibili

import (
	"encoding/json"
	"net/http"
	"regexp"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestSypderSign(t *testing.T) {
	const (
		payload    = "hello world"
		secretKey  = "secret_key"
		md5Vec     = "61c95854c1cd8179128b54c19ac01c28"
		sha1Vec    = "15272f929f45d7f15e2bbfd7237741538847de8a"
		sha256Vec  = "cf1a418afaafc798df48fd804a2abf6970283afd8c40b41f818ad9b6ca4f8ca8"
		sha224Vec  = "84ec4f01a2e1a183a372856e28783a72ce64df7adbe42876e615b2aa"
		sha512Vec  = "bd6afeb7b0814f35093386b649c7f5065ad3e802ee22b5370ad379dc03532e613bf96e8778bc6a95ec4a550b8b2dee45615390b1788aa1610e116f83f013c458"
		sha384Vec  = "a5e1c5ab19af6f809c5d7a57b18144388497f9d41881461f8ae61370d993877f14029ebbf668e0470caf6695eedada87"
		cascadeVec = "f2dd8373dad611978f350343a10686e9f6355e2a78440674790a60e567305ae3"
		skipVec    = "41738fb460678a2c0c7ad54115010afc08b67032"
	)

	tests := []struct {
		name       string
		secretRule []int
		expect     string
	}{
		{"MD5", []int{0}, md5Vec},
		{"SHA1", []int{1}, sha1Vec},
		{"SHA256", []int{2}, sha256Vec},
		{"SHA224", []int{3}, sha224Vec},
		{"SHA512", []int{4}, sha512Vec},
		{"SHA384", []int{5}, sha384Vec},
		{"级联", []int{0, 1, 2}, cascadeVec},
		{"跳过未知规则", []int{0, 9, 1}, skipVec},
		{"空规则返回原文", nil, payload},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sypderSign(payload, tt.secretRule, secretKey)
			if got != tt.expect {
				t.Fatalf("sypderSign(%q, %v, %q) = %q, 期望 %q", payload, tt.secretRule, secretKey, got, tt.expect)
			}
			if len(tt.secretRule) > 0 {
				if matched, _ := regexp.MatchString(`^[0-9a-f]+$`, got); !matched {
					t.Fatalf("sypderSign 输出应为小写 hex，实际为 %q", got)
				}
			}
		})
	}
}

func TestLiveTraceIDJSON(t *testing.T) {
	got, err := liveTraceIDJSON(1, 2, 3, 4)
	if err != nil {
		t.Fatalf("liveTraceIDJSON 返回错误: %v", err)
	}
	if got != "[1,2,3,4]" {
		t.Fatalf("liveTraceIDJSON = %q, 期望 %q", got, "[1,2,3,4]")
	}
	var parsed []int64
	if err := json.Unmarshal([]byte(got), &parsed); err != nil {
		t.Fatalf("liveTraceIDJSON 结果不是合法 JSON: %v", err)
	}
	if len(parsed) != 4 || parsed[0] != 1 || parsed[3] != 4 {
		t.Fatalf("liveTraceIDJSON 反解结果不正确: %v", parsed)
	}
}

func TestLiveTraceDeviceJSON(t *testing.T) {
	got, err := liveTraceDeviceJSON("buvid", "uuid")
	if err != nil {
		t.Fatalf("liveTraceDeviceJSON 返回错误: %v", err)
	}
	if got != `["buvid","uuid"]` {
		t.Fatalf("liveTraceDeviceJSON = %q, 期望 %q", got, `["buvid","uuid"]`)
	}
	var parsed []string
	if err := json.Unmarshal([]byte(got), &parsed); err != nil {
		t.Fatalf("liveTraceDeviceJSON 结果不是合法 JSON: %v", err)
	}
	if len(parsed) != 2 || parsed[0] != "buvid" || parsed[1] != "uuid" {
		t.Fatalf("liveTraceDeviceJSON 反解结果不正确: %v", parsed)
	}
}

func TestLiveHeartBeatSignatureJSON(t *testing.T) {
	got, err := liveHeartBeatSignatureJSON(58, 211, 3, 1234567, "buvid", "uuid", 1000, 2000)
	if err != nil {
		t.Fatalf("liveHeartBeatSignatureJSON 返回错误: %v", err)
	}
	// 字段顺序是协议要求，必须精确冻结。
	expect := `{"platform":"web","parent_id":58,"area_id":211,"seq_id":3,"room_id":1234567,` +
		`"buvid":"buvid","uuid":"uuid","ets":1000,"time":60,"ts":2000}`
	if got != expect {
		t.Fatalf("liveHeartBeatSignatureJSON = %q, 期望 %q", got, expect)
	}
}

func TestRandomLiveUUID(t *testing.T) {
	uuid, err := randomLiveUUID()
	if err != nil {
		t.Fatalf("randomLiveUUID 返回错误: %v", err)
	}
	matched, _ := regexp.MatchString(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, uuid)
	if !matched {
		t.Fatalf("randomLiveUUID 格式不正确: %q", uuid)
	}
	again, err := randomLiveUUID()
	if err != nil {
		t.Fatalf("randomLiveUUID 第二次返回错误: %v", err)
	}
	if uuid == again {
		t.Fatalf("两次生成的 uuid 不应相同: %q", uuid)
	}
}

func TestMoveFormParams(t *testing.T) {
	type form struct {
		ID      string `json:"id"`
		RUID    int64  `json:"ruid"`
		Keep    string `json:"keep"`
		Deleted string `json:"deleted" request:"-"`
	}
	r := resty.New().R()
	if err := withParams(r, form{ID: "[1,2,3,4]", RUID: 9, Keep: "q", Deleted: "x"}); err != nil {
		t.Fatalf("withParams 返回错误: %v", err)
	}
	if err := moveFormParams("id", "ruid")(r); err != nil {
		t.Fatalf("moveFormParams 返回错误: %v", err)
	}
	if r.FormData.Get("id") != "[1,2,3,4]" || r.FormData.Get("ruid") != "9" {
		t.Fatalf("moveFormParams 未把字段移入表单: %v", r.FormData)
	}
	if r.QueryParam.Has("id") || r.QueryParam.Has("ruid") {
		t.Fatalf("moveFormParams 未从 query 删除字段: %v", r.QueryParam)
	}
	if r.QueryParam.Get("keep") != "q" {
		t.Fatalf("moveFormParams 不应影响未列出的字段: %v", r.QueryParam)
	}
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fatalf("withParams 默认分支应设置表单 Content-Type: %q", r.Header.Get("Content-Type"))
	}
}

func TestFillFormCsrf(t *testing.T) {
	r := resty.New().R().SetCookies([]*http.Cookie{{Name: "bili_jct", Value: "token123"}})
	if err := fillFormCsrf(nil)(r); err != nil {
		t.Fatalf("fillFormCsrf 返回错误: %v", err)
	}
	if r.FormData.Get("csrf") != "token123" || r.FormData.Get("csrf_token") != "token123" {
		t.Fatalf("fillFormCsrf 表单结果不正确: %v", r.FormData)
	}

	empty := resty.New().R()
	if err := fillFormCsrf(nil)(empty); err == nil {
		t.Fatal("缺少 bili_jct 时 fillFormCsrf 应返回错误")
	}
}
