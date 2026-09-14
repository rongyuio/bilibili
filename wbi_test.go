package bilibili

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"testing"
	"time"
)

// seqTab 返回 [0,1,2,...,n-1] 的 mixin key 加密表，让 GenerateMixinKey 的结果可预测。
func seqTab(n int) []int {
	tab := make([]int, n)
	for i := range tab {
		tab[i] = i
	}
	return tab
}

func TestGenerateMixinKey(t *testing.T) {
	wbi := NewDefaultWbi().WithMixinKeyEncTab(seqTab(32))

	orig := "0123456789abcdefghijklmnopqrstuvwxyzABCDEF"
	if got := wbi.GenerateMixinKey(orig); got != "0123456789abcdefghijklmnopqrstuv" {
		t.Errorf("GenerateMixinKey = %q, want first 32 chars", got)
	}

	// 输入不足 32 字符时返回空字符串
	if got := wbi.GenerateMixinKey("short"); got != "" {
		t.Errorf("GenerateMixinKey(short) = %q, want empty", got)
	}
}

func TestWbiSanitizeString(t *testing.T) {
	if got := wbiSanitizeString("a!b'c(d)e*f"); got != "abcdef" {
		t.Errorf("wbiSanitizeString = %q, want abcdef", got)
	}
	if got := wbiSanitizeString("hello"); got != "hello" {
		t.Errorf("wbiSanitizeString = %q, want hello", got)
	}
}

func TestSignMap(t *testing.T) {
	wbi := NewDefaultWbi().WithMixinKeyEncTab(seqTab(32))
	const imgKey = "01234567890123456789012345678901"
	const subKey = "23456789012345678901234567890123"
	wbi.SetKeys(imgKey, subKey)

	ctx := context.Background()
	ts := time.Unix(1700000000, 0)
	payload := map[string]string{"b": "2", "a": "1", "w_rid": "stale"}

	result, err := wbi.SignMap(ctx, payload, ts)
	if err != nil {
		t.Fatal(err)
	}
	if result["wts"] != "1700000000" {
		t.Errorf("wts = %q, want 1700000000", result["wts"])
	}
	if result["a"] != "1" || result["b"] != "2" {
		t.Errorf("original params not preserved: %v", result)
	}
	if result["w_rid"] == "stale" {
		t.Error("w_rid should be recomputed")
	}

	// 期望 w_rid = md5("a=1&b=2&wts=1700000000" + mixinKey)
	mixinKey := wbi.GenerateMixinKey(imgKey + subKey)
	hash := md5.Sum([]byte("a=1&b=2&wts=1700000000" + mixinKey))
	want := hex.EncodeToString(hash[:])
	if result["w_rid"] != want {
		t.Errorf("w_rid = %q, want %q", result["w_rid"], want)
	}
}

func TestSignQuery(t *testing.T) {
	wbi := NewDefaultWbi().WithMixinKeyEncTab(seqTab(32))
	wbi.SetKeys("01234567890123456789012345678901", "23456789012345678901234567890123")

	ctx := context.Background()
	ts := time.Unix(1700000000, 0)
	query := url.Values{"b": {"2"}, "a": {"1"}}

	result, err := wbi.SignQuery(ctx, query, ts)
	if err != nil {
		t.Fatal(err)
	}
	if result.Get("a") != "1" || result.Get("b") != "2" {
		t.Errorf("original query not preserved: %v", result)
	}
	if result.Get("wts") != "1700000000" {
		t.Errorf("wts = %q, want 1700000000", result.Get("wts"))
	}
	if result.Get("w_rid") == "" {
		t.Error("w_rid missing")
	}
}

func TestSignMapNilContext(t *testing.T) {
	wbi := NewDefaultWbi()
	if _, err := wbi.SignMap(nil, map[string]string{"a": "1"}, time.Now()); err == nil {
		t.Error("SignMap with nil ctx should return error")
	}
}
