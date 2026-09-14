package bilibili

import (
	"net/url"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type encodedParams struct {
	query       url.Values
	body        map[string]any
	contentType string
	present     bool
}

func withParams(r *resty.Request, in any) error {
	params, err := encodeParams(in)
	if err != nil {
		return err
	}
	if !params.present {
		return nil
	}
	for key, values := range params.query {
		r.QueryParam[key] = values
	}
	r.SetHeader("Content-Type", params.contentType)
	if len(params.body) > 0 {
		r.SetBody(params.body)
	}
	return nil
}

func encodeParams(in any) (params encodedParams, err error) {
	if in == nil {
		return params, nil
	}

	inType := reflect.TypeOf(in)
	inValue := reflect.ValueOf(in)

	switch inType.Kind() {
	case reflect.Ptr:
		// 如果是空指针，直接返回
		if inValue.IsNil() {
			return params, nil
		}
		inType = inType.Elem()
		inValue = inValue.Elem()
		if inType.Kind() != reflect.Struct {
			return params, errors.New("参数类型错误")
		}
	case reflect.Struct:
	default:
		return params, errors.New("参数类型错误")
	}

	params.present = true
	params.query = make(url.Values)
	bodyMap := make(map[string]any, 4)
	contentType := ""
	for i := range inType.NumField() {
		fieldType := inType.Field(i)
		if !fieldType.IsExported() {
			continue
		}
		fieldValue := inValue.Field(i)
		tValue := fieldType.Tag.Get("request")

		if tValue == "-" {
			continue
		}

		// 获取字段名
		var fieldName string

		tagMap := parseTag(tValue)
		if name, ok := tagMap["field"]; ok {
			fieldName = name
		} else if jsonValue := fieldType.Tag.Get("json"); jsonValue != "" && jsonValue != "-" {
			if index := strings.Index(jsonValue, ","); index != -1 {
				jsonValue = jsonValue[:index]
			}
			fieldName = jsonValue
		} else {
			fieldName = toSnakeCase(fieldType.Name)
		}

		var realVal any
		if !fieldValue.IsZero() {
			realVal = fieldValue.Interface()
		} else {
			// 设置了 omitempty 代表不传
			if _, ok := tagMap["omitempty"]; ok {
				continue
			}
			// 设置了 default 代表使用默认值
			if v, ok := tagMap["default"]; ok {
				realVal = v
			} else {
				// 否则使用零值
				realVal = fieldValue.Interface()
			}
		}

		contentType = "application/x-www-form-urlencoded"
		for name := range tagMap {
			switch name {
			case "query":
				contentType = "application/x-www-form-urlencoded"
			case "json":
				contentType = "application/json"
			case "form-data":
				contentType = "multipart/form-data"
			}
		}
		if contentType == "application/x-www-form-urlencoded" {
			_, ok1 := tagMap["json"]
			_, ok2 := tagMap["form-data"]
			if !ok1 && !ok2 {
				// 对query类型的字段进行特殊处理
				if fieldType.Type.Kind() == reflect.Slice {
					strSlice := make([]string, 0, 4)
					for i := range fieldValue.Len() {
						strSlice = append(strSlice, cast.ToString(fieldValue.Index(i).Interface()))
					}
					realVal = strings.Join(strSlice, ",")
				}
			}
			params.query.Set(fieldName, cast.ToString(realVal))
		} else {
			bodyMap[fieldName] = realVal
		}
	}

	params.contentType = contentType
	params.body = bodyMap

	return params, nil
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

func toSnakeCase(s string) string {
	var result strings.Builder
	result.Grow(len(s) * 2)

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
