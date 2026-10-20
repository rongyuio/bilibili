package login

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// Cookie 刷新与首页访问接口。响应模型见 cookie_model.go。

// errLoginNilContext、errLoginMissingHTTPResponse 复刻根包请求链路的错误语义；
// 包级静态变量满足 err113。
var (
	errLoginNilContext          = errors.New("request context is nil")
	errLoginMissingHTTPResponse = errors.New("missing HTTP response")
	errRefreshCsrfNotFound      = errors.New("refresh CSRF not found in correspond page")
)

// loginCheckContext 在请求开始前拒绝 nil 或已结束的 context。
func loginCheckContext(ctx context.Context) error {
	if ctx == nil {
		return errLoginNilContext
	}
	return ctx.Err()
}

// loginSafeEndpoint 去掉 URL 中的 query、fragment 与 user 信息，避免错误信息泄漏。
func loginSafeEndpoint(endpoint string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "<invalid URL>"
	}
	u.RawQuery, u.Fragment, u.User = "", "", nil
	return u.String()
}

// loginCookieValue 从 Cookie 列表中按名称取值。
func loginCookieValue(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

// GetWebCookieRefreshInfo 获取web端cookie刷新信息
func (c *Client) GetWebCookieRefreshInfo(ctx context.Context) (*GetWebCookieRefreshInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/x/passport-login/web/cookie/info"
	)

	return bilibili.Execute[*GetWebCookieRefreshInfoResult](ctx, c.Client, method, url, nil)
}

// GetWebCookieRefreshCsrfParam 指定刷新口令所用的毫秒时间戳。
type GetWebCookieRefreshCsrfParam struct {
	Timestamp int64 `json:"timestamp"` // 毫秒时间戳
}

// 正则匹配 <div id="1-name">RefreshCsrf</div> 中的刷新口令
var refreshCsrfRegex = regexp.MustCompile(`<div\s+id="1-name"\s*>(.*?)</div>`)

// GetWebCookieRefreshCsrf 获取web端cookie刷新口令。
//
// 该接口返回 HTML 页面而非 code/message/data 结构，无法复用 Execute；
// 这里走 Resty 逃生通道，并自行复刻根包的 Cookie 快照与响应 Cookie 合并语义。
func (c *Client) GetWebCookieRefreshCsrf(ctx context.Context, param GetWebCookieRefreshCsrfParam) (*GetWebCookieRefreshCsrfResult, error) {
	if err := loginCheckContext(ctx); err != nil {
		return nil, err
	}
	correspondPath, err := loginGetCorrespondPath(param.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("getCorrespondPath: %w", err)
	}

	url := "https://www.bilibili.com/correspond/1/" + correspondPath
	r := c.Resty().R().SetContext(ctx).SetCookies(c.GetCookies())
	response, err := r.Execute(resty.MethodGet, url)
	if response != nil && response.RawResponse != nil {
		c.SetCookies(response.Cookies())
	}
	if err != nil {
		return nil, fmt.Errorf("request refresh CSRF: %s %s: %w", resty.MethodGet, loginSafeEndpoint(url), err)
	}
	if response == nil || response.RawResponse == nil {
		return nil, fmt.Errorf("request refresh CSRF: %s %s: %w", resty.MethodGet, loginSafeEndpoint(url), errLoginMissingHTTPResponse)
	}
	if !response.IsSuccess() {
		return nil, fmt.Errorf("request refresh CSRF: %w", &bilibili.HTTPError{
			Method:     resty.MethodGet,
			Endpoint:   loginSafeEndpoint(url),
			StatusCode: response.StatusCode(),
		})
	}

	matches := refreshCsrfRegex.FindStringSubmatch(response.String())
	if len(matches) < 2 {
		return nil, errRefreshCsrfNotFound
	}

	return &GetWebCookieRefreshCsrfResult{RefreshCsrf: matches[1]}, nil
}

// RefreshCookieParam 是刷新 Cookie 的请求参数。
type RefreshCookieParam struct {
	Csrf         string `json:"csrf,omitempty" request:"query,omitempty"`                    // 位于 Cookie 中的bili_jct字段，不传将当前 client 中获取
	RefreshCsrf  string `json:"refresh_csrf" request:"query"`                                // 实时刷新口令
	Source       string `json:"source,omitempty" request:"query,omitempty,default=main_web"` // 访问来源，一般为：main_web
	RefreshToken string `json:"refresh_token" request:"query"`                               // 在登录成功时返回的持久化刷新口令，localStorage 中的ac_time_value字段
}

// RefreshCookie 刷新Cookie
func (c *Client) RefreshCookie(ctx context.Context, param RefreshCookieParam) (*RefreshCookieResult, error) {
	// Csrf falls back to the cookie snapshot when the caller omits it; unlike
	// fillCsrf this interface tolerates an empty value and lets the server decide.
	if param.Csrf == "" {
		param.Csrf = loginCookieValue(c.GetCookies(), "bili_jct")
	}
	const (
		method = resty.MethodPost
		url    = "https://passport.bilibili.com/x/passport-login/web/cookie/refresh"
	)

	return bilibili.Execute[*RefreshCookieResult](ctx, c.Client, method, url, param)
}

func init() {
	const publicKeyPEM = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLgd2OAkcGVtoE3ThUREbio0Eg
Uc/prcajMKXvkCKFCWhJYJcLkcM2DKKcSeFpD/j6Boy538YXnR6VhcuUJOhH2x71
nzPjfdTcqMz7djHum0qSZA0AyCBDABUqCrfNgCiJ00Ra7GmRj+YCK1NJEuewlb40
JNrRuoEUXpabUzGB8QIDAQAB
-----END PUBLIC KEY-----
`
	pubKeyBlock, _ := pem.Decode([]byte(publicKeyPEM))
	if pubKeyBlock == nil {
		panic("rsa public key decode error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	var ok bool
	loginCorrespondPathPublicKey, ok = pubInterface.(*rsa.PublicKey)
	if !ok {
		panic("rsa public key type error")
	}
}

var loginCorrespondPathPublicKey *rsa.PublicKey

// loginGetCorrespondPath 生成 CorrespondPath 算法，参数：GetWebCookieRefreshInfoResult.Timestamp
func loginGetCorrespondPath(timestamp int64) (string, error) {
	var (
		hash   = sha256.New()
		random = rand.Reader
		msg    = fmt.Appendf(nil, "refresh_%d", timestamp)
	)
	encryptedData, err := rsa.EncryptOAEP(hash, random, loginCorrespondPathPublicKey, msg, nil)
	if err != nil {
		return "", fmt.Errorf("加密 CorrespondPath: %w", err)
	}

	return hex.EncodeToString(encryptedData), nil
}
