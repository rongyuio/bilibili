package bilibili

import (
	"errors"
	"testing"
)

func TestHTTPError(t *testing.T) {
	e := newHTTPError("GET", "https://api.bilibili.com/x?a=1&b=2#frag", 404)
	if e.Method != "GET" || e.StatusCode != 404 {
		t.Errorf("HTTPError fields wrong: %+v", e)
	}
	if e.Endpoint != "https://api.bilibili.com/x" {
		t.Errorf("Endpoint = %q, want https://api.bilibili.com/x", e.Endpoint)
	}
	if got := e.Error(); got != "GET https://api.bilibili.com/x: HTTP status 404" {
		t.Errorf("Error() = %q", got)
	}
}

func TestSafeEndpoint(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"https://api.bilibili.com/x?a=1&b=2", "https://api.bilibili.com/x"},
		{"https://user:pass@api.bilibili.com/x#frag", "https://api.bilibili.com/x"},
		{"%zz", "<invalid URL>"},
	}
	for _, c := range cases {
		if got := safeEndpoint(c.in); got != c.want {
			t.Errorf("safeEndpoint(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestParamError(t *testing.T) {
	inner := errors.New("inner")
	e := &ParamError{
		RootType:  "VideoParam",
		GoField:   "Bvid",
		Parameter: "bvid",
		Location:  "query",
		Err:       inner,
		reason:    "boom",
	}
	if !errors.Is(e, inner) {
		t.Error("ParamError should unwrap to inner")
	}
	if got := e.Error(); got != "encode VideoParam (Go field Bvid, parameter bvid, location query): boom" {
		t.Errorf("Error() = %q", got)
	}
}

func TestDecodeError(t *testing.T) {
	inner := errors.New("inner")
	e := &DecodeError{
		Method:   "GET",
		Endpoint: "https://api.bilibili.com/x",
		RootType: "VideoInfo",
		GoField:  "Bvid",
		JSONPath: "data.bvid",
		Expected: "string",
		Actual:   "number",
		Offset:   12,
		Exact:    true,
		Err:      inner,
	}
	if !errors.Is(e, inner) {
		t.Error("DecodeError should unwrap to inner")
	}
	want := "GET https://api.bilibili.com/x: decode VideoInfo at data.bvid (Go field Bvid, expected string, got number, offset 12, exact true)"
	if got := e.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestErrorString(t *testing.T) {
	e := Error{Code: -403, Message: "access denied"}
	if got := e.Error(); got != "错误码: -403, 错误信息: access denied" {
		t.Errorf("Error() = %q", got)
	}
}
