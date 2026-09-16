package bilibili

import (
	"bytes"
	"errors"
	"reflect"
	"slices"
	"sync"
	"testing"
)

const spaceFeedEndpoint = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space"

// spaceFeedBody 组装一条只含单个动态的空间动态响应，author 是 module_author 对象的 JSON 内容。
func spaceFeedBody(author string) []byte {
	return []byte(`{"code":0,"message":"0","data":{"has_more":false,"offset":"","update_baseline":"","update_num":1,"items":[` +
		`{"id_str":1,"type":"DYNAMIC_TYPE_AV","visible":true,"modules":{"module_author":{` + author + `}}}]}}`)
}

func droppedPaths(dropped []DroppedField) []string {
	paths := make([]string, 0, len(dropped))
	for _, field := range dropped {
		paths = append(paths, field.JSONPath)
	}
	return paths
}

// TestDecodeResponseFollowingDrift 是 issue #20 的最小复现：following 在线上返回布尔。
func TestDecodeResponseFollowingDrift(t *testing.T) {
	body := spaceFeedBody(`"mid":123,"name":"作者","following":true,"pub_ts":1700000000`)
	var (
		out     DynamicInfo
		dropped []DroppedField
	)
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(field DroppedField) {
		dropped = append(dropped, field)
	})
	if err != nil {
		t.Fatalf("decodeResponse: %v", err)
	}
	if len(dropped) != 0 {
		t.Errorf("严格解码成功时不应上报丢弃，got %v", dropped)
	}
	if len(out.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(out.Items))
	}
	author := out.Items[0].Modules.ModuleAuthor
	if author.Following.Kind() != "boolean" || author.Following.String() != "true" {
		t.Errorf("Following = %q/%q, want boolean/true", author.Following.Kind(), author.Following.String())
	}
	if author.Name != "作者" {
		t.Errorf("Name = %q, want 作者", author.Name)
	}
	if author.Mid.String() != "123" {
		t.Errorf("Mid = %q, want 123", author.Mid.String())
	}
	if author.PubTs.String() != "1700000000" {
		t.Errorf("PubTs = %q, want 1700000000", author.PubTs.String())
	}
}

func TestDecodeResponseFollowingVariants(t *testing.T) {
	cases := []struct {
		in       string
		wantKind string
		wantText string
	}{
		{`null`, "null", ""},
		{`2`, "number", "2"},
		{`1`, "number", "1"},
		{`true`, "boolean", "true"},
		{`false`, "boolean", "false"},
		{`"2"`, "string", "2"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			var out DynamicInfo
			err := decodeResponse("GET", spaceFeedEndpoint, spaceFeedBody(`"following":`+c.in), &out, nil)
			if err != nil {
				t.Fatalf("decodeResponse(%s): %v", c.in, err)
			}
			following := out.Items[0].Modules.ModuleAuthor.Following
			if following.Kind() != c.wantKind || following.String() != c.wantText {
				t.Errorf("Following = %q/%q, want %q/%q", following.Kind(), following.String(), c.wantKind, c.wantText)
			}
		})
	}
}

func TestDecodeResponseDropsDriftedField(t *testing.T) {
	body := spaceFeedBody(`"mid":123,"name":"作者","following":null,"pub_ts":true`)
	var (
		out     DynamicInfo
		dropped []DroppedField
	)
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(field DroppedField) {
		dropped = append(dropped, field)
	})
	if err != nil {
		t.Fatalf("decodeResponse: %v", err)
	}
	author := out.Items[0].Modules.ModuleAuthor
	if author.PubTs.String() != "" {
		t.Errorf("PubTs = %q, want 被丢弃后的零值", author.PubTs.String())
	}
	// 其余字段必须保真。
	if author.Name != "作者" || author.Mid.String() != "123" {
		t.Errorf("漂移字段之外的字段未保真: %+v", author)
	}
	if len(out.Items) != 1 || out.UpdateNum.String() != "1" {
		t.Errorf("响应其余部分未保真: items=%d update_num=%q", len(out.Items), out.UpdateNum.String())
	}
	if len(dropped) != 1 {
		t.Fatalf("上报 %d 条丢弃，want 1: %v", len(dropped), dropped)
	}
	field := dropped[0]
	want := DroppedField{
		Method:   "GET",
		Endpoint: spaceFeedEndpoint,
		RootType: "bilibili.DynamicInfo",
		GoField:  "Items.Modules.ModuleAuthor.PubTs",
		JSONPath: "$.data.items[0].modules.module_author.pub_ts",
		Expected: "json.Number",
		Actual:   "boolean",
		Offset:   int64(bytes.Index(body, []byte(`"pub_ts":`))+len(`"pub_ts":`)) + 1,
	}
	if !reflect.DeepEqual(field, want) {
		t.Errorf("DroppedField = %+v, want %+v", field, want)
	}
}

func TestDecodeResponseDropsMultipleFields(t *testing.T) {
	body := []byte(`{"code":0,"message":"0","data":{"has_more":true,"update_num":1,"items":[` +
		`{"modules":{"module_author":{"mid":1,"name":"A","following":true,"pub_ts":true}}},` +
		`{"modules":{"module_author":{"mid":true,"name":"B","following":null,"pub_ts":1700000000}}}]}}`)
	var (
		out     DynamicInfo
		dropped []DroppedField
	)
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(field DroppedField) {
		dropped = append(dropped, field)
	})
	if err != nil {
		t.Fatalf("decodeResponse: %v", err)
	}
	want := []string{
		"$.data.items[0].modules.module_author.pub_ts",
		"$.data.items[1].modules.module_author.mid",
	}
	if paths := droppedPaths(dropped); !slices.Equal(paths, want) {
		t.Errorf("丢弃路径 = %v, want %v", paths, want)
	}
	if len(out.Items) != 2 || !out.HasMore {
		t.Errorf("响应其余部分未保真: items=%d has_more=%t", len(out.Items), out.HasMore)
	}
	if out.Items[0].Modules.ModuleAuthor.Name != "A" || out.Items[1].Modules.ModuleAuthor.Name != "B" {
		t.Errorf("未漂移的字段未保真: %+v", out.Items)
	}
}

func TestDecodeResponseNestedIndexPath(t *testing.T) {
	body := []byte(`{"code":0,"message":"0","data":{"items":[` +
		`{"modules":{"module_author":{"name":"A","pub_ts":1}}},` +
		`{"modules":{"module_author":{"name":"B","pub_ts":true}}},` +
		`{"modules":{"module_author":{"name":"C","pub_ts":3}}}]}}`)
	var (
		out     DynamicInfo
		dropped []DroppedField
	)
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(field DroppedField) {
		dropped = append(dropped, field)
	})
	if err != nil {
		t.Fatalf("decodeResponse: %v", err)
	}
	if paths := droppedPaths(dropped); !slices.Equal(paths, []string{"$.data.items[1].modules.module_author.pub_ts"}) {
		t.Errorf("丢弃路径 = %v", paths)
	}
	for i, want := range []string{"A", "B", "C"} {
		if got := out.Items[i].Modules.ModuleAuthor.Name; got != want {
			t.Errorf("items[%d].Name = %q, want %q", i, got, want)
		}
	}
	if out.Items[1].Modules.ModuleAuthor.PubTs.String() != "" {
		t.Error("漂移字段未被丢弃")
	}
	if out.Items[0].Modules.ModuleAuthor.PubTs.String() != "1" || out.Items[2].Modules.ModuleAuthor.PubTs.String() != "3" {
		t.Error("未漂移的元素被误改")
	}
}

// errAlwaysFailField 模拟自定义解码器内部失败：这是不透明边界，容错不应介入。
var errAlwaysFailField = errors.New("always fail field")

type alwaysFailField struct{}

func (*alwaysFailField) UnmarshalJSON([]byte) error { return errAlwaysFailField }

type toleranceProbe struct {
	Value alwaysFailField `json:"value"`
	Other string          `json:"other"`
}

func TestDecodeResponsePruneFailureUsesOriginalBytes(t *testing.T) {
	body := []byte(`{"code":0,"data":{"value":123,"other":"x"}}`)
	var (
		out     toleranceProbe
		dropped []DroppedField
	)
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(field DroppedField) {
		dropped = append(dropped, field)
	})
	var decodeErr *DecodeError
	if !errors.As(err, &decodeErr) {
		t.Fatalf("err = %v, want *DecodeError", err)
	}
	if !errors.Is(decodeErr, errAlwaysFailField) {
		t.Errorf("错误链丢失原始错误: %v", decodeErr)
	}
	if decodeErr.Exact {
		t.Error("自定义解码器边界应标记 Exact=false")
	}
	if decodeErr.JSONPath != "$.data.value" || decodeErr.GoField != "Value" {
		t.Errorf("定位 = %s / %s, want $.data.value / Value", decodeErr.JSONPath, decodeErr.GoField)
	}
	// 偏移必须指向原始响应体，而不是剪枝后的副本。
	if want := int64(bytes.Index(body, []byte("123"))) + 1; decodeErr.Offset != want {
		t.Errorf("Offset = %d, want %d", decodeErr.Offset, want)
	}
	if len(dropped) != 0 {
		t.Errorf("解码失败时不应上报丢弃，got %v", dropped)
	}
	if out.Value != (alwaysFailField{}) || out.Other != "" {
		t.Errorf("失败时不应写入 out: %+v", out)
	}
}

func TestDecodeResponseRootShapeMismatchIsFatal(t *testing.T) {
	body := []byte(`{"code":0,"data":[]}`)
	var (
		out     DynamicInfo
		dropped []DroppedField
	)
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(field DroppedField) {
		dropped = append(dropped, field)
	})
	var decodeErr *DecodeError
	if !errors.As(err, &decodeErr) {
		t.Fatalf("err = %v, want *DecodeError", err)
	}
	if decodeErr.JSONPath != "$.data" || decodeErr.Actual != "array" || !decodeErr.Exact {
		t.Errorf("定位 = %s / %s / exact=%t, want $.data / array / true", decodeErr.JSONPath, decodeErr.Actual, decodeErr.Exact)
	}
	if want := int64(bytes.Index(body, []byte("[]"))) + 1; decodeErr.Offset != want {
		t.Errorf("Offset = %d, want %d", decodeErr.Offset, want)
	}
	if len(dropped) != 0 {
		t.Errorf("根节点不匹配不应上报丢弃，got %v", dropped)
	}
}

func TestDecodeResponseNilHandler(t *testing.T) {
	body := spaceFeedBody(`"pub_ts":true`)
	var out DynamicInfo
	if err := decodeResponse("GET", spaceFeedEndpoint, body, &out, nil); err != nil {
		t.Fatalf("decodeResponse: %v", err)
	}
	if out.Items[0].Modules.ModuleAuthor.PubTs.String() != "" {
		t.Error("无回调时容错仍应生效")
	}
}

func TestDecodeResponseOutNil(t *testing.T) {
	body := spaceFeedBody(`"pub_ts":true`)
	called := false
	if err := decodeResponse("GET", spaceFeedEndpoint, body, nil, func(DroppedField) { called = true }); err != nil {
		t.Fatalf("decodeResponse: %v", err)
	}
	if called {
		t.Error("out=nil 时不应解码，也不应上报丢弃")
	}
}

func TestDecodeResponseBusinessErrorSkipsTolerance(t *testing.T) {
	body := spaceFeedBody(`"pub_ts":true`)
	body = bytes.Replace(body, []byte(`"code":0`), []byte(`"code":-404`), 1)
	called := false
	var out DynamicInfo
	err := decodeResponse("GET", spaceFeedEndpoint, body, &out, func(DroppedField) { called = true })
	var businessError Error
	if !errors.As(err, &businessError) || businessError.Code != -404 {
		t.Fatalf("err = %v, want 业务错误 -404", err)
	}
	if called {
		t.Error("业务错误时不应进入解码")
	}
	if len(out.Items) != 0 {
		t.Error("业务错误时不应写入 out")
	}
}

func TestDroppedFieldHandlerConcurrent(t *testing.T) {
	client := New()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range 200 {
			client.SetDroppedFieldHandler(func(DroppedField) {})
		}
	}()
	go func() {
		defer wg.Done()
		for range 200 {
			if handler := client.droppedFieldHandler(); handler != nil {
				handler(DroppedField{})
			}
		}
	}()
	wg.Wait()
	client.SetDroppedFieldHandler(nil)
	if client.droppedFieldHandler() != nil {
		t.Error("传 nil 应清除回调")
	}
}
