package bilibili

import (
	"fmt"
	"net/url"
)

// HTTPError reports a response status rejected by an endpoint's success rules.
// Endpoint omits query parameters, user information and fragments in library errors.
// It does not contain response bodies, request headers or transport errors.
type HTTPError struct {
	Method     string
	Endpoint   string
	StatusCode int
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s %s: HTTP status %d", e.Method, e.Endpoint, e.StatusCode)
}

func newHTTPError(method, endpoint string, statusCode int) *HTTPError {
	return &HTTPError{Method: method, Endpoint: safeEndpoint(endpoint), StatusCode: statusCode}
}

// ParamError identifies an invalid request parameter. GoField includes a slice
// index when conversion of an element fails. Empty fields indicate a root error.
// Error omits parameter values and the cause text; Err may contain sensitive data.
type ParamError struct {
	RootType  string
	GoField   string
	Parameter string
	Location  string
	Err       error
	reason    string
}

func (e *ParamError) Error() string {
	reason := e.reason
	if reason == "" {
		reason = "parameter encoding failed"
	}
	return fmt.Sprintf("encode %s (Go field %s, parameter %s, location %s): %s",
		e.RootType, e.GoField, e.Parameter, e.Location, reason)
}

func (e *ParamError) Unwrap() error { return e.Err }

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
