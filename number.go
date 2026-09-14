package bilibili

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

// NumberOrString preserves a JSON number, string (including "--"), or null.
// Its zero value represents null. Conversions never silently replace invalid values.
type NumberOrString struct {
	text string
	kind string
}

// errInvalidNumberOrString 表示值既不是数字也不是字符串。
var errInvalidNumberOrString = errors.New("NumberOrString requires a number, string, or null")

// String returns the raw text form.
func (n NumberOrString) String() string { return n.text }

// Kind returns "number", "string", or "null".
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

// UnmarshalJSON accepts a JSON number, string, or null.
func (n *NumberOrString) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if bytes.Equal(data, []byte("null")) {
		*n = NumberOrString{}
		return nil
	}
	if len(data) > 0 && data[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		*n = NumberOrString{text: text, kind: "string"}
		return nil
	}
	if !json.Valid(data) || jsonKind(data) != "number" {
		return errInvalidNumberOrString
	}
	*n = NumberOrString{text: string(data), kind: "number"}
	return nil
}

// MarshalJSON emits null, the original string, or the original number text.
func (n NumberOrString) MarshalJSON() ([]byte, error) {
	switch n.kind {
	case "string":
		return json.Marshal(n.text)
	case "number":
		return []byte(n.text), nil
	default:
		return []byte("null"), nil
	}
}
