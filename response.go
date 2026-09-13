package bilibili

import (
	"encoding/json"
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

func decodeResponse(method, endpoint string, body []byte, out any) error {
	if err := validateOutput(out); err != nil {
		return err
	}
	var envelope commonResp[json.RawMessage]
	if err := json.Unmarshal(body, &envelope); err != nil {
		return newDecodeError(method, endpoint, body, reflect.TypeOf(envelope), "$", 0, err)
	}
	if envelope.Code != 0 {
		return Error{Code: envelope.Code, Message: envelope.Message}
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
	if err := json.Unmarshal(envelope.Data, value.Interface()); err != nil {
		var offset int64
		// Only failures need to locate data within the original response.
		for _, child := range jsonChildren(body) {
			if strings.EqualFold(child.key, "data") {
				offset = child.offset
			}
		}
		return newDecodeError(method, endpoint, envelope.Data, target.Type(), "$.data", offset, err)
	}
	target.Set(value.Elem())
	return nil
}
