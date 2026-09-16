package bilibili

import (
	"encoding/json"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestDeviceReqJSON(t *testing.T) {
	cases := []struct {
		name        string
		webLocation string
		want        string
	}{
		{
			name:        "默认动态页标识",
			webLocation: dynamicLotteryWebLocation,
			want:        `{"platform":"web","device":"pc","spmid":"333.1330"}`,
		},
		{
			name:        "自定义标识原样还原",
			webLocation: "444.41",
			want:        `{"platform":"web","device":"pc","spmid":"444.41"}`,
		},
		{
			name:        "空值仍产出合法 JSON",
			webLocation: "",
			want:        `{"platform":"web","device":"pc","spmid":""}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := deviceReqJSON(tc.webLocation)
			if err != nil {
				t.Fatalf("deviceReqJSON(%q) 返回错误: %v", tc.webLocation, err)
			}
			if got != tc.want {
				t.Fatalf("deviceReqJSON(%q) = %q，期望 %q", tc.webLocation, got, tc.want)
			}

			var parsed struct {
				Platform string `json:"platform"`
				Device   string `json:"device"`
				Spmid    string `json:"spmid"`
			}
			if err := json.Unmarshal([]byte(got), &parsed); err != nil {
				t.Fatalf("产出不是合法 JSON: %v", err)
			}
			if parsed.Spmid != tc.webLocation {
				t.Fatalf("spmid = %q，期望 %q", parsed.Spmid, tc.webLocation)
			}
			if parsed.Platform != "web" || parsed.Device != "pc" {
				t.Fatalf("platform/device = %q/%q，期望 web/pc", parsed.Platform, parsed.Device)
			}
		})
	}
}

func TestDeviceReqJSONEscapesValue(t *testing.T) {
	// 值含引号与反斜杠时，拼接结果仍须是合法 JSON，且能还原为原值。
	dirty := `333.1"33\0`
	got, err := deviceReqJSON(dirty)
	if err != nil {
		t.Fatalf("deviceReqJSON 返回错误: %v", err)
	}

	var parsed struct {
		Spmid string `json:"spmid"`
	}
	if err := json.Unmarshal([]byte(got), &parsed); err != nil {
		t.Fatalf("含特殊字符时产出非法 JSON: %v；实际输出 %q", err, got)
	}
	if parsed.Spmid != dirty {
		t.Fatalf("spmid = %q，期望原值 %q", parsed.Spmid, dirty)
	}
}

func TestDynamicLotteryDeviceHandler(t *testing.T) {
	const webLocation = "444.41"
	param := GetDynamicLotteryInfoParam{
		BusinessID:   "123456",
		BusinessType: 1,
		WebLocation:  webLocation,
	}

	r := resty.New().R()
	if err := withParams(r, param); err != nil {
		t.Fatalf("withParams 失败: %v", err)
	}
	if err := dynamicLotteryDeviceHandler(param.WebLocation)(r); err != nil {
		t.Fatalf("handler 失败: %v", err)
	}

	values, ok := r.QueryParam["x-bili-device-req-json"]
	if !ok {
		t.Fatal("query 缺少 x-bili-device-req-json")
	}
	if len(values) != 1 {
		t.Fatalf("x-bili-device-req-json 有 %d 个值，期望 1", len(values))
	}

	var parsed struct {
		Platform string `json:"platform"`
		Device   string `json:"device"`
		Spmid    string `json:"spmid"`
	}
	if err := json.Unmarshal([]byte(values[0]), &parsed); err != nil {
		t.Fatalf("注入的 JSON 非法: %v", err)
	}
	if parsed.Spmid != webLocation {
		t.Fatalf("spmid = %q，期望 %q", parsed.Spmid, webLocation)
	}
	if parsed.Platform != "web" || parsed.Device != "pc" {
		t.Fatalf("platform/device = %q/%q，期望 web/pc", parsed.Platform, parsed.Device)
	}

	// 业务字段仍在 query，且 spmid 与 web_location 保持一致。
	if got := r.QueryParam.Get("web_location"); got != webLocation {
		t.Fatalf("web_location = %q，期望 %q", got, webLocation)
	}
	if got := r.QueryParam.Get("business_id"); got != "123456" {
		t.Fatalf("business_id = %q，期望 123456", got)
	}
	if got := r.QueryParam.Get("business_type"); got != "1" {
		t.Fatalf("business_type = %q，期望 1", got)
	}
}

func TestWebLocationOrDefault(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		fallback string
		want     string
	}{
		{name: "留空取默认值", value: "", fallback: "333.1330", want: "333.1330"},
		{name: "非空原样返回", value: "444.41", fallback: "333.1330", want: "444.41"},
		{name: "默认值也为空时返回空", value: "", fallback: "", want: ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := webLocationOrDefault(tc.value, tc.fallback); got != tc.want {
				t.Fatalf("webLocationOrDefault(%q, %q) = %q，期望 %q", tc.value, tc.fallback, got, tc.want)
			}
		})
	}
}

func TestWebLocationConstants(t *testing.T) {
	// 冻结三个接口的默认值。直播与话题同值属巧合，一旦其中一个接口需要改值，
	// 应当只改对应常量，因此这里逐个断言而不是合并成一个。
	cases := []struct {
		name string
		got  string
		want string
	}{
		{name: "动态抽奖", got: dynamicLotteryWebLocation, want: "333.1330"},
		{name: "直播勋章", got: liveMedalWebLocation, want: "0.0"},
		{name: "话题动态", got: topicFeedWebLocation, want: "0.0"},
	}

	for _, tc := range cases {
		if tc.got != tc.want {
			t.Fatalf("%s 的默认页面标识 = %q，期望 %q", tc.name, tc.got, tc.want)
		}
	}
}
