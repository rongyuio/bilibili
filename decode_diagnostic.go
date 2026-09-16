package bilibili

import (
	"bytes"
	"encoding"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"unicode"
)

type jsonChild struct {
	key    string
	raw    json.RawMessage
	offset int64
}

// jsonChildren preserves document order, duplicate keys and byte offsets.
// It is only used after an unsuccessful decode, on syntactically valid JSON.
func jsonChildren(raw []byte) []jsonChild {
	d := json.NewDecoder(bytes.NewReader(raw))
	token, err := d.Token()
	if err != nil || (token != json.Delim('{') && token != json.Delim('[')) {
		return nil
	}
	var children []jsonChild
	for index := 0; d.More(); index++ {
		key := strconv.Itoa(index)
		if token == json.Delim('{') {
			k, err := d.Token()
			if err != nil {
				return children
			}
			key, _ = k.(string)
		}
		var value json.RawMessage
		if err := d.Decode(&value); err != nil {
			return children
		}
		children = append(children, jsonChild{key: key, raw: value, offset: d.InputOffset() - int64(len(value))})
	}
	return children
}

func jsonKind(raw []byte) string {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return "missing"
	}
	switch raw[0] {
	case '"':
		return "string"
	case '{':
		return "object"
	case '[':
		return "array"
	case 't', 'f':
		return "boolean"
	case 'n':
		return "null"
	default:
		return "number"
	}
}

func newDecodeError(method, endpoint string, raw []byte, typ reflect.Type, path string, base int64, cause error) *DecodeError {
	e := &DecodeError{Method: method, Endpoint: safeEndpoint(endpoint), RootType: diagnosticTypeName(typ),
		JSONPath: path, Expected: diagnosticTypeName(typ), Actual: jsonKind(raw), Err: cause}
	if syntax, ok := errors.AsType[*json.SyntaxError](cause); ok {
		e.Offset = base + syntax.Offset
		e.Actual = "invalid JSON"
		return e
	}
	if mismatch, ok := errors.AsType[*json.UnmarshalTypeError](cause); ok {
		e.Offset = base + mismatch.Offset
		e.Expected = diagnosticTypeName(mismatch.Type)
		e.GoField = mismatch.Field
	}
	if !json.Valid(raw) {
		return e
	}
	location, exact := diagnoseValue(raw, typ, path, "", base, false)
	if location != nil {
		e.JSONPath, e.GoField = location.JSONPath, location.GoField
		e.Expected, e.Actual, e.Offset = location.Expected, location.Actual, location.Offset
		e.Exact = exact
	}
	return e
}

var (
	jsonDecoderType    = reflect.TypeFor[json.Unmarshaler]()
	textDecoderType    = reflect.TypeFor[encoding.TextUnmarshaler]()
	rawMessageType     = reflect.TypeFor[json.RawMessage]()
	numberType         = reflect.TypeFor[json.Number]()
	numberOrStringType = reflect.TypeFor[NumberOrString]()
)

func hasDecoder(t reflect.Type) bool {
	return t.Implements(jsonDecoderType) || reflect.PointerTo(t).Implements(jsonDecoderType) ||
		t.Implements(textDecoderType) || reflect.PointerTo(t).Implements(textDecoderType)
}

// scalarTarget 判断该类型应当作为叶子交给 encoding/json 试探，而不是继续向下遍历。
// 带 ,string 的字段、json.Number、除 []byte 外的标量都属于这一类。
func scalarTarget(t reflect.Type, quoted bool) bool {
	if quoted || t == numberType {
		return true
	}
	if t.Kind() == reflect.Struct || t.Kind() == reflect.Map || t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		return t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 && !hasDecoder(t.Elem())
	}
	return true
}

// probeDecode 用 encoding/json 自身校验叶子值，覆盖 ,string 与数值溢出。
// 它只解码到一个临时值，不会写入调用方的目标。
func probeDecode(raw []byte, t reflect.Type, quoted bool) error {
	if quoted {
		wrapper := reflect.StructOf([]reflect.StructField{{Name: "Value", Type: t, Tag: `json:"value,string"`}})
		return json.Unmarshal(append(append([]byte(`{"value":`), raw...), '}'), reflect.New(wrapper).Interface())
	}
	return json.Unmarshal(raw, reflect.New(t).Interface())
}

// mapKeyMismatch 探测字符串键能否转换为 map 的键类型。
func mapKeyMismatch(key reflect.Type, member string) bool {
	keyJSON, _ := json.Marshal(member) //nolint:errchkjson // 键为字符串，Marshal 不会失败
	probe := append(append([]byte{'{'}, keyJSON...), []byte(":null}")...)
	return json.Unmarshal(probe, reflect.New(reflect.MapOf(key, rawMessageType)).Interface()) != nil
}

// diagnoseValue never invokes user-defined decoders. A custom decoder is an
// opaque boundary; if multiple boundaries could fail, retain their common parent.
func diagnoseValue(raw []byte, t reflect.Type, path, field string, offset int64, quoted bool) (*DecodeError, bool) {
	location := &DecodeError{JSONPath: path, GoField: field, Expected: diagnosticTypeName(t), Actual: jsonKind(raw), Offset: offset + 1}
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t == rawMessageType {
		return nil, false
	}
	if t == numberOrStringType {
		if numberOrStringAcceptsKind(jsonKind(raw)) {
			return nil, false
		}
		return location, true
	}
	if hasDecoder(t) {
		return location, false
	}
	if jsonKind(raw) == "null" {
		return nil, false
	}
	if scalarTarget(t, quoted) {
		if t.Kind() == reflect.Interface {
			return nil, false
		}
		if probeDecode(raw, t, quoted) != nil {
			return location, true
		}
		return nil, false
	}
	kind := jsonKind(raw)
	if ((t.Kind() == reflect.Struct || t.Kind() == reflect.Map) && kind != "object") ||
		((t.Kind() == reflect.Array || t.Kind() == reflect.Slice) && kind != "array") {
		return location, true
	}
	var fields []decodeField
	if t.Kind() == reflect.Struct {
		fields = decodeFields(t)
	}
	var boundary *DecodeError
	for i, child := range jsonChildren(raw) {
		childType, childField, childPath, childQuoted := t, field, path, false
		switch t.Kind() {
		case reflect.Struct:
			f, ok := matchDecodeField(fields, child.key)
			if !ok {
				continue
			}
			childType, childQuoted = f.typ, f.quoted
			childField = f.goPath
			if field != "" {
				childField = field + "." + childField
			}
			childPath = appendJSONKey(path, child.key)
		case reflect.Map:
			if hasDecoder(t.Key()) {
				return location, false
			}
			// Map-key conversion errors are located at their JSON member.
			if mapKeyMismatch(t.Key(), child.key) {
				location.JSONPath = appendJSONKey(path, child.key)
				location.Expected = diagnosticTypeName(t.Key())
				location.Actual = "object key"
				location.Offset = offset + child.offset + 1
				return location, true
			}
			childType, childPath = t.Elem(), appendJSONKey(path, child.key)
		case reflect.Array, reflect.Slice:
			if t.Kind() == reflect.Array && i >= t.Len() {
				continue
			}
			childType, childPath = t.Elem(), path+"["+child.key+"]"
		}
		found, exact := diagnoseValue(child.raw, childType, childPath, childField, offset+child.offset, childQuoted)
		if exact {
			return found, true
		}
		if found != nil {
			if boundary == nil {
				boundary = found
			} else {
				boundary = location
			}
		}
	}
	return boundary, false
}

func appendJSONKey(path, key string) string {
	if key != "" {
		identifier := true
		for i, r := range key {
			if r != '_' && !unicode.IsLetter(r) && (i <= 0 || !unicode.IsDigit(r)) {
				identifier = false
				break
			}
		}
		if identifier {
			return path + "." + key
		}
	}
	encoded, _ := json.Marshal(key) //nolint:errchkjson // key 为字符串，Marshal 不会失败
	return path + "[" + string(encoded) + "]"
}
