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

func TestNumberOrStringInvalid(t *testing.T) {
	for _, in := range []string{`{"a":1}`, `[1]`, `true`} {
		var n NumberOrString
		if err := json.Unmarshal([]byte(in), &n); !errors.Is(err, errInvalidNumberOrString) {
			t.Errorf("Unmarshal(%s) err = %v, want errInvalidNumberOrString", in, err)
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
