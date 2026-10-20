package login

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 登录相关的非标准响应接口：公钥与盐、Cookie 刷新口令。

var (
	errLoginRawJSONDecode = errors.New("响应不是合法的 JSON")
	errRefreshCsrfFormat  = errors.New("refresh_csrf 页面格式无法解析")
	errCorrespondPath     = errors.New("correspondPath 不能为空或包含非法字符")
)

// GetLoginKeyResult 是密码登录所需的 RSA 公钥与密码盐值。
type GetLoginKeyResult struct {
	Hash string `json:"hash"` // 密码盐值。有效时间为 20s。恒为 16 字符。需要拼接在明文密码之前
	Key  string `json:"key"`  // rsa 公钥。PEM 格式编码。加密密码时需要使用
}

// GetLoginKey 获取 web 端密码登录的公钥与盐值（新版接口，标准信封）。
func (c *Client) GetLoginKey(ctx context.Context) (*GetLoginKeyResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/x/passport-login/web/key"
	)
	return bilibili.Execute[*GetLoginKeyResult](ctx, c.Client, method, url, nil)
}

// GetLoginKeyLegacyResult 是旧版公钥与盐值接口的响应。
//
// 与 GetLoginKeyResult 字段相同，但按「各接口返回体独立」的约定单独定义。
type GetLoginKeyLegacyResult struct {
	Hash string `json:"hash"` // 密码盐值。有效时间为 20s。恒为 16 字符。需要拼接在明文密码之前
	Key  string `json:"key"`  // rsa 公钥。PEM 格式编码。加密密码时需要使用
}

// GetLoginKeyLegacy 获取 web 端密码登录的公钥与盐值（旧版接口）。
//
// 响应直接是 {hash, key} 对象，没有 code/message/data 信封，因此走 ExecuteRaw。
func (c *Client) GetLoginKeyLegacy(ctx context.Context) (*GetLoginKeyLegacyResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/login?act=getkey"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, url, nil)
	if err != nil {
		return nil, err
	}
	var out GetLoginKeyLegacyResult
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, errLoginRawJSONDecode
	}
	return &out, nil
}

var legacyRefreshCsrfRe = regexp.MustCompile(`(?s)<div id="1-name">\s*([^<\s]+)\s*</div>`)

// GetRefreshCsrfParam 指定对应页面路径。
type GetRefreshCsrfParam struct {
	// CorrespondPath 使用当前毫秒时间戳生成的签名，由文档中的生成算法加密获得。
	CorrespondPath string
}

// GetRefreshCsrf 获取实时刷新口令 refresh_csrf。
//
// 接口返回一个 HTML 页面，口令位于 <div id="1-name"> 标签中；不是 JSON 信封。
func (c *Client) GetRefreshCsrf(ctx context.Context, param GetRefreshCsrfParam) (string, error) {
	path := strings.TrimSpace(param.CorrespondPath)
	if path == "" || strings.ContainsAny(path, "/?#\\") {
		return "", errCorrespondPath
	}
	const (
		method = resty.MethodGet
		base   = "https://www.bilibili.com/correspond/1/"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, base+path, nil)
	if err != nil {
		return "", err
	}
	m := legacyRefreshCsrfRe.FindSubmatch(body)
	if m == nil {
		return "", errRefreshCsrfFormat
	}
	return string(m[1]), nil
}
