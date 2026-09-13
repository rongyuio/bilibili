package bilibili

import (
	"fmt"
	"net/url"
)

// DecodeError describes a failed response decode without including response values.
// Offset is one-based in the response body; zero means unavailable.
// Exact is false when only a custom decoder boundary or the root can be located.
type DecodeError struct {
	Method   string
	Endpoint string
	RootType string
	GoField  string
	JSONPath string
	Expected string
	Actual   string
	Offset   int64
	Exact    bool
	Err      error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("%s %s: decode %s at %s (Go field %s, expected %s, got %s, offset %d, exact %t)",
		e.Method, e.Endpoint, e.RootType, e.JSONPath, e.GoField, e.Expected, e.Actual, e.Offset, e.Exact)
}

func (e *DecodeError) Unwrap() error { return e.Err }

func safeEndpoint(endpoint string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "<invalid URL>"
	}
	u.RawQuery, u.Fragment, u.User = "", "", nil
	return u.String()
}

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Message)
}
