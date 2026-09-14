package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
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
		return params, parameterError(root, "", "", "", "invalid parameter type", errors.New("parameters must be a struct or a pointer to a struct"))
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
				return params, parameterError(root, field.Name, name, location, "conflicting body encodings", errors.New("JSON and multipart fields cannot share a request body"))
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
		return strings.Join(locations, ","), errors.New("a field must declare at most one request location")
	}
	if len(locations) == 1 {
		return locations[0], nil
	}
	return "query", nil
}

func parameterString(value reflect.Value, realVal any, field string) (string, string, error) {
	// Preserve the existing query slice convention, including nil slices and defaults.
	if value.Kind() == reflect.Slice {
		values := make([]string, value.Len())
		for i := range value.Len() {
			text, err := cast.ToStringE(value.Index(i).Interface())
			if err != nil {
				return "", fmt.Sprintf("%s[%d]", field, i), err
			}
			values[i] = text
		}
		return strings.Join(values, ","), field, nil
	}
	text, err := cast.ToStringE(realVal)
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
