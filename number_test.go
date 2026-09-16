package bilibili

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestNumberOrStringNumber(t *testing.T) {
	var n NumberOrString
	if err := json.Unmarshal([]byte("123"), &n); err != nil {
		t.Fatal(err)
	}
	if n.Kind() != "number" {
		t.Errorf("Kind = %q, want number", n.Kind())
	}
	if n.String() != "123" {
		t.Errorf("String = %q, want 123", n.String())
	}
	if v, err := n.Int64(); err != nil || v != 123 {
		t.Errorf("Int64 = %d, %v; want 123, nil", v, err)
	}
	if v, err := n.Float64(); err != nil || v != 123.0 {
		t.Errorf("Float64 = %v, %v; want 123, nil", v, err)
	}
}

func TestNumberOrStringString(t *testing.T) {
	var n NumberOrString
	if err := json.Unmarshal([]byte(`"--"`), &n); err != nil {
		t.Fatal(err)
	}
	if n.Kind() != "string" {
		t.Errorf("Kind = %q, want string", n.Kind())
	}
	if n.String() != "--" {
		t.Errorf("String = %q, want --", n.String())
	}
	if _, err := n.Int64(); err == nil {
		t.Error("Int64 should fail for --")
	}
}

func TestNumberOrStringNull(t *testing.T) {
	var n NumberOrString
	if err := json.Unmarshal([]byte("null"), &n); err != nil {
		t.Fatal(err)
	}
	if n.Kind() != "null" {
		t.Errorf("Kind = %q, want null", n.Kind())
	}
	if n.String() != "" {
		t.Errorf("String = %q, want empty", n.String())
	}
}

func TestNumberOrStringBool(t *testing.T) {
	for in, want := range map[string]string{"true": "true", "false": "false"} {
		var n NumberOrString
		if err := json.Unmarshal([]byte(in), &n); err != nil {
			t.Fatal(err)
		}
		if n.Kind() != "boolean" {
			t.Errorf("Kind = %q, want boolean", n.Kind())
		}
		if n.String() != want {
			t.Errorf("String = %q, want %s", n.String(), want)
		}
		if _, err := n.Int64(); err == nil {
			t.Errorf("Int64 should fail for %s", in)
		}
		if _, err := n.Float64(); err == nil {
			t.Errorf("Float64 should fail for %s", in)
		}
	}
}

func TestNumberOrStringInvalid(t *testing.T) {
	// 语法非法的输入由 encoding/json 在调用 UnmarshalJSON 之前拒绝，不会返回本错误。
	for _, in := range []string{`{"a":1}`, `[1]`} {
		var n NumberOrString
		if err := json.Unmarshal([]byte(in), &n); !errors.Is(err, errInvalidNumberOrString) {
			t.Errorf("Unmarshal(%s) err = %v, want errInvalidNumberOrString", in, err)
		}
	}
}

// TestNumberOrStringKindAgreement 锁住 UnmarshalJSON 与共享判定的一致性，
// 避免解码层与诊断/容错层对同一个 kind 得出不同结论。
func TestNumberOrStringKindAgreement(t *testing.T) {
	for _, in := range []string{"null", "123", `"x"`, "true", "{}", "[]"} {
		var n NumberOrString
		err := json.Unmarshal([]byte(in), &n)
		if accepted := numberOrStringAcceptsKind(jsonKind([]byte(in))); accepted == (err != nil) {
			t.Errorf("kind %q: acceptsKind = %t, err = %v; want 两者一致", in, accepted, err)
		}
	}
}

func TestNumberOrStringMarshalRoundTrip(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"123", "123"},
		{`"hello"`, `"hello"`},
		{`"--"`, `"--"`},
		{"true", "true"},
		{"false", "false"},
		{"null", "null"},
	}
	for _, c := range cases {
		var n NumberOrString
		if err := json.Unmarshal([]byte(c.in), &n); err != nil {
			t.Fatalf("Unmarshal(%s): %v", c.in, err)
		}
		out, err := json.Marshal(n)
		if err != nil {
			t.Fatalf("Marshal(%s): %v", c.in, err)
		}
		if string(out) != c.want {
			t.Errorf("Marshal(%s) = %s, want %s", c.in, out, c.want)
		}
	}
}
