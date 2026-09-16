package bilibili

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

// errNoCookies 表示匿名客户端初始化未取得任何 Cookie。
var errNoCookies = errors.New("initialize anonymous client: response contains no valid cookies")

// Client supports concurrent ordinary requests after configuration.
// Login, session replacement and configuration changes must happen between requests.
// A Client must not be copied after first use.
type Client struct {
	wbi         *WBI
	resty       *resty.Client
	cookieMu    sync.Mutex
	cookies     []*http.Cookie
	dropMu      sync.RWMutex
	dropHandler DroppedFieldHandler
}

// New 返回一个默认的 bilibili.Client
func New() *Client {
	restyClient := resty.New().
		SetRedirectPolicy(resty.NoRedirectPolicy()).
		SetTimeout(20*time.Second).
		SetHeader("Accept", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9").
		SetHeader("Origin", "https://www.bilibili.com").
		SetHeader("Referer", "https://www.bilibili.com/").
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	return NewWithClient(restyClient)
}

// NewAnonymousClient fetches guest cookies using the default client configuration.
// It returns initialization failures, including cancellation, to the caller.
func NewAnonymousClient(ctx context.Context) (*Client, error) {
	if err := checkContext(ctx); err != nil {
		return nil, fmt.Errorf("initialize anonymous client: %w", err)
	}
	client := New()
	r := client.newRequest(ctx).
		SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9").
		SetHeader("Pragma", "no-cache").
		SetHeader("Priority", "u=0, i").
		SetHeader("Sec-Ch-Ua", "Not").
		SetHeader("Sec-Ch-Ua-Mobile", "?0").
		SetHeader("Sec-Ch-Ua-Platform", "Windows").
		SetHeader("Sec-Fetch-Dest", "document").
		SetHeader("Sec-Fetch-Mode", "navigate").
		SetHeader("Sec-Fetch-Site", "none").
		SetHeader("Sec-Fetch-User", "?1").
		SetHeader("Upgrade-Insecure-Requests", "1").
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0")
	const endpoint = "https://www.bilibili.com/"
	resp, err := client.sendRaw(r, resty.MethodGet, endpoint)
	if err != nil {
		return nil, fmt.Errorf("initialize anonymous client: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("initialize anonymous client: %w", newHTTPError(resty.MethodGet, endpoint, resp.StatusCode()))
	}
	if len(client.GetCookies()) == 0 {
		return nil, errNoCookies
	}
	return client, nil
}

// NewWithClient takes exclusive ownership of restyClient; nil creates a default client.
// Explicit Cookies are copied into Client storage and removed from Resty. Its Jar
// is disabled; callers must export any existing Jar session before this call.
func NewWithClient(restyClient *resty.Client) *Client {
	if restyClient == nil {
		return New()
	}
	client := &Client{
		wbi:   NewDefaultWbi(),
		resty: restyClient,
	}
	client.SetCookies(restyClient.Cookies)
	restyClient.Cookies = nil
	restyClient.SetCookieJar(nil)
	client.wbi.owner = client
	return client
}

// Resty exposes configuration and a low-level escape hatch. Configure it only
// between requests. Direct Resty requests do not use Client cookie storage.
// Do not re-enable its Jar or set its Cookies; use Client.SetCookies instead.
func (c *Client) Resty() *resty.Client {
	return c.resty
}

// GetCookiesString 获取字符串格式的cookies，方便自行存储后下次使用。配合下面的 SetCookiesString 使用。
func (c *Client) GetCookiesString() string {
	cookies := c.GetCookies()
	cookieStrings := make([]string, 0, len(cookies))
	for _, cookie := range cookies {
		cookieStrings = append(cookieStrings, cookie.String())
	}
	return strings.Join(cookieStrings, "\n")
}

// SetCookiesString 设置Cookies，但是是字符串格式，配合 GetCookiesString 使用。有些功能必须登录或设置Cookies后才能使用。
func (c *Client) SetCookiesString(cookiesString string) {
	c.SetCookies((&resty.Response{RawResponse: &http.Response{Header: http.Header{
		"Set-Cookie": strings.Split(cookiesString, "\n"),
	}}}).Cookies())
}

// SetRawCookies 如果你是从浏览器request的header中直接复制出来的cookies，调用这个函数。
func (c *Client) SetRawCookies(rawCookies string) {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}

	c.SetCookies(req.Cookies())
}

// SetCookie 设置单个cookie
func (c *Client) SetCookie(cookie *http.Cookie) {
	c.SetCookies([]*http.Cookie{cookie})
}

// SetCookies merges independent copies by name. Call manually only between requests.
func (c *Client) SetCookies(cookies []*http.Cookie) {
	incoming := cloneCookies(cookies)
	c.cookieMu.Lock()
	defer c.cookieMu.Unlock()
	now := time.Now()
	for _, cookie := range incoming {
		remove := cookie.MaxAge < 0 || (cookie.MaxAge == 0 && !cookie.Expires.IsZero() && !cookie.Expires.After(now))
		if cookie.MaxAge > 0 {
			// Normalize once to an absolute deadline; snapshots must not renew MaxAge.
			seconds := min(int64(cookie.MaxAge), int64((1<<63-1)/time.Second))
			cookie.Expires = now.Add(time.Duration(seconds) * time.Second)
			cookie.MaxAge = 0
		}
		found := false
		for i, current := range c.cookies {
			if current.Name != cookie.Name {
				continue
			}
			if remove {
				c.cookies = append(c.cookies[:i], c.cookies[i+1:]...)
			} else {
				c.cookies[i] = cookie
			}
			found = true
			break
		}
		if !found && !remove {
			c.cookies = append(c.cookies, cookie)
		}
	}
}

// GetCookies returns a deep copy of the currently valid cookies.
func (c *Client) GetCookies() []*http.Cookie {
	c.cookieMu.Lock()
	defer c.cookieMu.Unlock()
	now := time.Now()
	valid := c.cookies[:0]
	for _, cookie := range c.cookies {
		if cookie.Expires.IsZero() || cookie.Expires.After(now) {
			valid = append(valid, cookie)
		}
	}
	clear(c.cookies[len(valid):])
	c.cookies = valid
	return cloneCookies(valid)
}

func cookieValue(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie != nil && cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
