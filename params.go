package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/url"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
)

type encodedParams struct {
	query        url.Values
	jsonBody     []byte
	multipart    map[string]string
	bodyLocation string
	present      bool
}

var (
	errParamNotStruct       = errors.New("parameters must be a struct or a pointer to a struct")
	errConflictingLocations = errors.New("a field must declare at most one request location")
	errConflictingBodies    = errors.New("JSON and multipart fields cannot share a request body")
)

func withParams(r *resty.Request, in any) error {
	params, err := encodeParams(in)
	if err != nil {
		return err
	}
	if !params.present {
		return nil
	}
	maps.Copy(r.QueryParam, params.query)
	switch params.bodyLocation {
	case "json":
		r.SetHeader("Content-Type", "application/json")
		r.SetBody(params.jsonBody)
	case "form-data":
		// Resty constructs the body and sets a matching boundary when sending.
		r.SetMultipartFormData(params.multipart)
	default:
		r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	}
	return nil
}

func encodeParams(in any) (params encodedParams, err error) {
	if in == nil {
		return params, nil
	}
	inType := reflect.TypeOf(in)
	inValue := reflect.ValueOf(in)
	if inType.Kind() == reflect.Pointer {
		if inValue.IsNil() {
			return params, nil
		}
		inType = inType.Elem()
		inValue = inValue.Elem()
	}
	root := diagnosticTypeName(inType)
	if inType.Kind() != reflect.Struct {
		return params, parameterError(root, "", "", "", "invalid parameter type", errParamNotStruct)
	}
	params.query = make(url.Values)
	params.multipart = make(map[string]string)
	body := make(map[string]json.RawMessage)
	for i := range inType.NumField() {
		field := inType.Field(i)
		if !field.IsExported() || field.Tag.Get("request") == "-" {
			continue
		}
		value := inValue.Field(i)
		tags := parseTag(field.Tag.Get("request"))
		name := parameterName(field, tags)
		realVal := value.Interface()
		if value.IsZero() {
			if _, omit := tags["omitempty"]; omit {
				continue
			}
			if fallback, ok := tags["default"]; ok {
				realVal = fallback
			}
		}
		location, err := parameterLocation(tags)
		if err != nil {
			return params, parameterError(root, field.Name, name, location, "conflicting request locations", err)
		}
		if location != "query" {
			if params.bodyLocation != "" && params.bodyLocation != location {
				return params, parameterError(root, field.Name, name, location, "conflicting body encodings", errConflictingBodies)
			}
			params.bodyLocation = location
		}
		params.present = true
		if location == "json" {
			// Encode each value once so custom marshalers are not executed again by Resty.
			raw, err := json.Marshal(realVal)
			if err != nil {
				return params, parameterError(root, field.Name, name, location, "JSON encoding failed", err)
			}
			body[name] = raw
			continue
		}
		text, fieldPath, err := parameterString(value, realVal, field.Name)
		if err != nil {
			return params, parameterError(root, fieldPath, name, location, "string conversion failed", err)
		}
		if location == "query" {
			params.query.Set(name, text)
		} else {
			params.multipart[name] = text
		}
	}
	if params.bodyLocation == "json" {
		params.jsonBody, err = json.Marshal(body)
		if err != nil {
			return params, parameterError(root, "", "", "json", "JSON encoding failed", err)
		}
	}
	return params, nil
}

func parameterName(field reflect.StructField, tags map[string]string) string {
	if name, ok := tags["field"]; ok {
		return name
	}
	if name := field.Tag.Get("json"); name != "" && name != "-" {
		name, _, _ = strings.Cut(name, ",")
		return name
	}
	return toSnakeCase(field.Name)
}

func parameterLocation(tags map[string]string) (string, error) {
	var locations []string
	for _, name := range []string{"query", "json", "form-data"} {
		if _, ok := tags[name]; ok {
			locations = append(locations, name)
		}
	}
	if len(locations) > 1 {
		return strings.Join(locations, ","), errConflictingLocations
	}
	if len(locations) == 1 {
		return locations[0], nil
	}
	return "query", nil
}

func parameterString(value reflect.Value, realVal any, field string) (text string, fieldPath string, err error) {
	// Preserve the existing query slice convention, including nil slices and defaults.
	if value.Kind() == reflect.Slice {
		values := make([]string, value.Len())
		for i := range value.Len() {
			s, convErr := cast.ToStringE(value.Index(i).Interface())
			if convErr != nil {
				return "", fmt.Sprintf("%s[%d]", field, i), convErr
			}
			values[i] = s
		}
		return strings.Join(values, ","), field, nil
	}
	text, err = cast.ToStringE(realVal)
	return text, field, err
}

func parameterError(root, field, name, location, reason string, cause error) *ParamError {
	return &ParamError{RootType: root, GoField: field, Parameter: name, Location: location, Err: cause, reason: reason}
}

func parseTag(tag string) map[string]string {
	parts := strings.Split(tag, ",")

	pMap := make(map[string]string, 10)
	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 1 {
			pMap[kv[0]] = ""
		} else {
			pMap[kv[0]] = kv[1]
		}
	}

	return pMap
}

// snakeCaseInitialisms 是推导请求参数名时需要整体视作一个词的初始缩写，
// 与导出标识符的命名约定保持一致。复数形式排在单数前面，匹配时长的优先。
var snakeCaseInitialisms = []string{
	"UUIDs", "UUID", "UIDs", "UID", "URLs", "URL", "URIs", "URI",
	"JSON", "HTTP", "APIs", "API", "IDs", "ID",
}

// toSnakeCase 把导出字段名转成请求参数名，仅在字段既没有 field 标签也没有 json 标签时使用。
//
// 初始缩写整体作为一个词，因此 IDs 转为 ids、IDsA 转为 ids_a、UIDType 转为 uid_type、
// DynamicID 转为 dynamic_id、BaseURL 转为 base_url。按 Go 惯例写成 ID/URL/UID 的字段，
// 推导出的参数名与旧写法（Ids/Url/Uid）保持一致。
func toSnakeCase(s string) string {
	runes := []rune(s)

	var words []string
	for i := 0; i < len(runes); {
		if initialism, next, ok := matchInitialism(runes, i); ok {
			words = append(words, initialism)
			i = next
			continue
		}

		// 普通单词：从当前位置收到底，遇到下一个大写字母另起一词。
		j := i + 1
		for j < len(runes) && !unicode.IsUpper(runes[j]) {
			j++
		}
		words = append(words, string(runes[i:j]))
		i = j
	}

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "_")
}

// matchInitialism 尝试在 runes[i] 处匹配一个初始缩写。只有缩写后面是字符串结尾、
// 大写字母或非字母字符时才成立，避免把 Initialization 这样的普通单词误当成缩写。
func matchInitialism(runes []rune, i int) (string, int, bool) {
	for _, initialism := range snakeCaseInitialisms {
		letters := []rune(initialism)

		next := i + len(letters)
		if next > len(runes) || string(runes[i:next]) != initialism {
			continue
		}

		if next == len(runes) || unicode.IsUpper(runes[next]) || !unicode.IsLetter(runes[next]) {
			return initialism, next, true
		}
	}

	return "", 0, false
}
