package bilibili

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"regexp"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// GetWebCookieRefreshInfo 获取web端cookie刷新信息
func (c *Client) GetWebCookieRefreshInfo(ctx context.Context) (*GetWebCookieRefreshInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/x/passport-login/web/cookie/info"
	)

	return execute[*GetWebCookieRefreshInfoResult](ctx, c, method, url, nil)
}

type GetWebCookieRefreshCsrfParam struct {
	Timestamp int64 `json:"timestamp"` // 毫秒时间戳
}

// 正则匹配 <div id="1-name">RefreshCsrf</div> 中的刷新口令
var refreshCsrfRegex = regexp.MustCompile(`<div\s+id="1-name"\s*>(.*?)</div>`)

// GetWebCookieRefreshCsrf 获取web端cookie刷新口令。
//
// 该接口返回 HTML 页面而非 code/message/data 结构，无法复用 execute；
// 仍通过 newRequest/sendRaw 保持 Cookie 快照与响应 Cookie 合并语义。
func (c *Client) GetWebCookieRefreshCsrf(ctx context.Context, param GetWebCookieRefreshCsrfParam) (*GetWebCookieRefreshCsrfResult, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	correspondPath, err := getCorrespondPath(param.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("getCorrespondPath: %w", err)
	}

	url := "https://www.bilibili.com/correspond/1/" + correspondPath
	response, err := c.sendRaw(c.newRequest(ctx), resty.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("request refresh CSRF: %w", err)
	}
	if !response.IsSuccess() {
		return nil, fmt.Errorf("request refresh CSRF: %w", newHTTPError(resty.MethodGet, url, response.StatusCode()))
	}

	matches := refreshCsrfRegex.FindStringSubmatch(response.String())
	if len(matches) < 2 {
		return nil, errors.New("refresh CSRF not found in correspond page")
	}

	return &GetWebCookieRefreshCsrfResult{RefreshCsrf: matches[1]}, nil
}

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
		r := c.newRequest(ctx)
		param.Csrf = cookieValue(r.Cookies, "bili_jct")
	}
	const (
		method = resty.MethodPost
		url    = "https://passport.bilibili.com/x/passport-login/web/cookie/refresh"
	)

	return execute[*RefreshCookieResult](ctx, c, method, url, param)
}

// RefreshWebCookie 一条龙完成 web 端 Cookie 刷新：依次获取刷新信息、获取刷新口令、刷新 Cookie。
//
// refreshToken 为登录成功或上次刷新得到的持久化刷新口令（localStorage 中的 ac_time_value）。
// 不需要刷新（刷新信息接口返回 refresh=false）时返回 (nil, nil)。
//
// 刷新成功后，返回值中的 RefreshToken 是新的持久化刷新口令，调用方必须自行持久化，
// 否则之后将无法再次刷新。刷新产生的新 Cookie（含新的 bili_jct）已自动合并进客户端。
func (c *Client) RefreshWebCookie(ctx context.Context, refreshToken string) (*RefreshCookieResult, error) {
	info, err := c.GetWebCookieRefreshInfo(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil || !info.Refresh {
		return nil, nil //nolint:nilnil // 不需要刷新是正常分支，文档已说明返回 (nil, nil)
	}
	csrf, err := c.GetWebCookieRefreshCsrf(ctx, GetWebCookieRefreshCsrfParam{Timestamp: info.Timestamp})
	if err != nil {
		return nil, err
	}
	return c.RefreshCookie(ctx, RefreshCookieParam{
		RefreshCsrf:  csrf.RefreshCsrf,
		RefreshToken: refreshToken,
	})
}

// GetHomePage 访问 B 站首页。响应中的 Set-Cookie 已由请求链路自动合并进客户端，
// 可用于补全或更新 Cookie（例如 buvid3 等设备字段）。页面内容为 HTML，本方法不解析，仅校验 HTTP 状态码。
func (c *Client) GetHomePage(ctx context.Context) error {
	if err := checkContext(ctx); err != nil {
		return err
	}
	r := c.newRequest(ctx).
		SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7").
		SetHeader("Sec-Fetch-Dest", "document").
		SetHeader("Sec-Fetch-Mode", "navigate").
		SetHeader("Sec-Fetch-Site", "none").
		SetHeader("Sec-Fetch-User", "?1").
		SetHeader("Upgrade-Insecure-Requests", "1")
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/"
	)
	response, err := c.sendRaw(r, method, url)
	if err != nil {
		return err
	}
	if !response.IsSuccess() {
		return newHTTPError(method, url, response.StatusCode())
	}
	return nil
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
	pubInterface, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	var ok bool
	correspondPathPublicKey, ok = pubInterface.(*rsa.PublicKey)
	if !ok {
		panic("rsa public key type error")
	}
}

var correspondPathPublicKey *rsa.PublicKey

// 生成CorrespondPath 算法，参数：GetWebCookieRefreshInfoResult.Timestamp
func getCorrespondPath(timestamp int64) (string, error) {
	var (
		hash   = sha256.New()
		random = rand.Reader
		msg    = fmt.Appendf(nil, "refresh_%d", timestamp)
	)
	encryptedData, err := rsa.EncryptOAEP(hash, random, correspondPathPublicKey, msg, nil)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return hex.EncodeToString(encryptedData), nil
}
