package bilibili

import (
	"bytes"
	"encoding"
	"encoding/json"
	"errors"
	"reflect"
	"sort"
	"strconv"
	"strings"
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
	e := &DecodeError{Method: method, Endpoint: safeEndpoint(endpoint), RootType: typ.String(),
		JSONPath: path, Expected: typ.String(), Actual: jsonKind(raw), Err: cause}
	var syntax *json.SyntaxError
	if errors.As(cause, &syntax) {
		e.Offset = base + syntax.Offset
		e.Actual = "invalid JSON"
		return e
	}
	var mismatch *json.UnmarshalTypeError
	if errors.As(cause, &mismatch) {
		e.Offset = base + mismatch.Offset
		e.Expected = mismatch.Type.String()
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

// diagnoseValue never invokes user-defined decoders. A custom decoder is an
// opaque boundary; if multiple boundaries could fail, retain their common parent.
func diagnoseValue(raw []byte, t reflect.Type, path, field string, offset int64, quoted bool) (*DecodeError, bool) {
	location := &DecodeError{JSONPath: path, GoField: field, Expected: t.String(), Actual: jsonKind(raw), Offset: offset + 1}
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t == rawMessageType {
		return nil, false
	}
	if t == numberOrStringType {
		if kind := jsonKind(raw); kind == "number" || kind == "string" || kind == "null" {
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
	// Validate leaves through encoding/json itself, including ,string and overflow.
	if quoted || t == numberType || (t.Kind() != reflect.Struct && t.Kind() != reflect.Map && t.Kind() != reflect.Array && t.Kind() != reflect.Slice) ||
		(t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8) {
		if t.Kind() == reflect.Interface {
			return nil, false
		}
		var err error
		if quoted {
			wrapper := reflect.StructOf([]reflect.StructField{{Name: "Value", Type: t, Tag: `json:"value,string"`}})
			err = json.Unmarshal(append(append([]byte(`{"value":`), raw...), '}'), reflect.New(wrapper).Interface())
		} else {
			err = json.Unmarshal(raw, reflect.New(t).Interface())
		}
		if err != nil {
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
			keyJSON, _ := json.Marshal(child.key)
			probe := append(append([]byte{'{'}, keyJSON...), []byte(":null}")...)
			if err := json.Unmarshal(probe, reflect.New(reflect.MapOf(t.Key(), rawMessageType)).Interface()); err != nil {
				location.JSONPath = appendJSONKey(path, child.key)
				location.Expected = t.Key().String()
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
			if !(r == '_' || unicode.IsLetter(r) || (i > 0 && unicode.IsDigit(r))) {
				identifier = false
				break
			}
		}
		if identifier {
			return path + "." + key
		}
	}
	encoded, _ := json.Marshal(key)
	return path + "[" + string(encoded) + "]"
}

type decodeField struct {
	name   string
	goPath string
	typ    reflect.Type
	index  []int
	tagged bool
	quoted bool
}

// Match encoding/json's promoted-field dominance, then its exact/folded lookup.
func decodeFields(root reflect.Type) []decodeField {
	var candidates []decodeField
	var visit func(reflect.Type, string, []int, map[reflect.Type]bool)
	visit = func(t reflect.Type, prefix string, index []int, parents map[reflect.Type]bool) {
		if parents[t] {
			return
		}
		parents[t] = true
		defer delete(parents, t)
		for i := range t.NumField() {
			f := t.Field(i)
			ft := f.Type
			if ft.Kind() == reflect.Pointer {
				ft = ft.Elem()
			}
			if !f.IsExported() && (!f.Anonymous || ft.Kind() != reflect.Struct) {
				continue
			}
			tag := f.Tag.Get("json")
			if tag == "-" {
				continue
			}
			parts := strings.Split(tag, ",")
			name := parts[0]
			if !validJSONTag(name) {
				name = ""
			}
			idx := append(append([]int(nil), index...), i)
			goPath := prefix + f.Name
			if name == "" && f.Anonymous && ft.Kind() == reflect.Struct {
				visit(ft, goPath+".", idx, parents)
				continue
			}
			tagged := name != ""
			if name == "" {
				name = f.Name
			}
			quoted := false
			for _, option := range parts[1:] {
				if option == "string" {
					quoted = true
				}
			}
			switch ft.Kind() {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.String:
			default:
				quoted = false
			}
			candidates = append(candidates, decodeField{name: name, goPath: goPath, typ: f.Type, index: idx, tagged: tagged, quoted: quoted})
		}
	}
	visit(root, "", nil, make(map[reflect.Type]bool))
	groups := make(map[string][]decodeField)
	for _, f := range candidates {
		groups[f.name] = append(groups[f.name], f)
	}
	var fields []decodeField
	for _, group := range groups {
		sort.SliceStable(group, func(i, j int) bool {
			if len(group[i].index) != len(group[j].index) {
				return len(group[i].index) < len(group[j].index)
			}
			return group[i].tagged && !group[j].tagged
		})
		if len(group) > 1 && len(group[0].index) == len(group[1].index) && group[0].tagged == group[1].tagged {
			continue
		}
		fields = append(fields, group[0])
	}
	sort.Slice(fields, func(i, j int) bool {
		for k := 0; k < len(fields[i].index) && k < len(fields[j].index); k++ {
			if fields[i].index[k] != fields[j].index[k] {
				return fields[i].index[k] < fields[j].index[k]
			}
		}
		return len(fields[i].index) < len(fields[j].index)
	})
	return fields
}

func matchDecodeField(fields []decodeField, key string) (decodeField, bool) {
	for _, f := range fields {
		if f.name == key {
			return f, true
		}
	}
	for _, f := range fields {
		if strings.EqualFold(f.name, key) {
			return f, true
		}
	}
	return decodeField{}, false
}

func validJSONTag(tag string) bool {
	if tag == "" {
		return false
	}
	for _, r := range tag {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", r) {
			return false
		}
	}
	return true
}
