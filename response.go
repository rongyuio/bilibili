package bilibili

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type commonResp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// WBI nav may return a nonzero business code alongside usable signing keys.
func decodeWBIResponse(body []byte, out any) error {
	if err := json.Unmarshal(body, out); err != nil {
		return newDecodeError("GET", "https://api.bilibili.com/x/web-interface/nav", body, reflect.TypeOf(out).Elem(), "$", 0, err)
	}
	return nil
}

func decodeResponse(method, endpoint string, body []byte, out any, onDrop DroppedFieldHandler) error {
	if err := validateOutput(out); err != nil {
		return err
	}
	var envelope commonResp[json.RawMessage]
	if err := json.Unmarshal(body, &envelope); err != nil {
		return newDecodeError(method, endpoint, body, reflect.TypeOf(envelope), "$", 0, err)
	}
	if envelope.Code != 0 {
		return fmt.Errorf("%s %s: %w", method, safeEndpoint(endpoint), Error{Code: envelope.Code, Message: envelope.Message})
	}
	if out == nil {
		return nil
	}
	// Decode into a fresh value so callers never receive a partially decoded result.
	target := reflect.ValueOf(out).Elem()
	value := reflect.New(target.Type())
	if len(envelope.Data) == 0 {
		envelope.Data = json.RawMessage("null")
	}
	err := json.Unmarshal(envelope.Data, value.Interface())
	if err == nil {
		target.Set(value.Elem())
		return nil
	}
	// Only failures need to locate data within the original response.
	offset, found := dataOffset(body)
	// 单个字段的类型漂移不应作废整条响应：剪掉不兼容的叶子后重试。
	// 诊断始终基于未剪枝的字节，偏移量才能指向原始响应体。
	if pruned, dropped, changed := tolerateDecode(envelope.Data, target.Type(), offset, found, onDrop != nil); changed {
		retry := reflect.New(target.Type())
		if json.Unmarshal(pruned, retry.Interface()) == nil {
			target.Set(retry.Elem())
			endpoint, root := safeEndpoint(endpoint), diagnosticTypeName(target.Type())
			for _, field := range dropped {
				field.Method, field.Endpoint, field.RootType = method, endpoint, root
				onDrop(field)
			}
			return nil
		}
	}
	return newDecodeError(method, endpoint, envelope.Data, target.Type(), "$.data", offset, err)
}

// dataOffset 返回 data 值在响应体中的 0 起始偏移；响应缺少 data 时 found 为 false。
func dataOffset(body []byte) (int64, bool) {
	var offset int64
	found := false
	for _, child := range jsonChildren(body) {
		if strings.EqualFold(child.key, "data") {
			offset, found = child.offset, true
		}
	}
	return offset, found
}
