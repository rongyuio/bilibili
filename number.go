package bilibili

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

// NumberOrString preserves a JSON number, string (including "--"), boolean, or null.
// Its zero value represents null. Conversions never silently replace invalid values.
// 布尔值来自线上真实漂移（动态 following 在登录态返回布尔），不是通用的宽松入口。
type NumberOrString struct {
	text string
	kind string
}

// errInvalidNumberOrString 表示值不是数字、字符串、布尔值或 null。
var errInvalidNumberOrString = errors.New("NumberOrString requires a number, string, boolean, or null")

// numberOrStringAcceptsKind 是接受集合的唯一判定来源。
// 它同时被 UnmarshalJSON 与 decode_diagnostic.go / decode_tolerate.go 使用，
// 修改时必须保证三处仍然一致。
func numberOrStringAcceptsKind(kind string) bool {
	switch kind {
	case "number", "string", "boolean", "null":
		return true
	default:
		return false
	}
}

// String returns the raw text form.
func (n NumberOrString) String() string { return n.text }

// Kind returns "number", "string", "boolean", or "null".
func (n NumberOrString) Kind() string {
	if n.kind == "" {
		return "null"
	}
	return n.kind
}

// Int64 parses the value as a decimal int64.
func (n NumberOrString) Int64() (int64, error) { return strconv.ParseInt(n.text, 10, 64) }

// Float64 parses the value as a float64.
func (n NumberOrString) Float64() (float64, error) { return strconv.ParseFloat(n.text, 64) }

// UnmarshalJSON accepts a JSON number, string, boolean, or null.
func (n *NumberOrString) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	kind := jsonKind(data)
	if !numberOrStringAcceptsKind(kind) {
		return errInvalidNumberOrString
	}
	switch kind {
	case "null":
		*n = NumberOrString{}
		return nil
	case "string":
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		*n = NumberOrString{text: text, kind: "string"}
		return nil
	default: // number、boolean：原样保留文本
		if !json.Valid(data) {
			return errInvalidNumberOrString
		}
		*n = NumberOrString{text: string(data), kind: kind}
		return nil
	}
}

// MarshalJSON emits null, the original string, number, or boolean text.
func (n NumberOrString) MarshalJSON() ([]byte, error) {
	switch n.kind {
	case "string":
		return json.Marshal(n.text)
	case "number", "boolean":
		return []byte(n.text), nil
	default:
		return []byte("null"), nil
	}
}
