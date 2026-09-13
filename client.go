package bilibili

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	wbi      *WBI
	resty    *resty.Client
	cookieMu sync.Mutex
	cookies  []*http.Cookie
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

// NewAnonymousClient 返回一个带有游客cookie的 bilibili.Client
func NewAnonymousClient() *Client {
	url := "https://www.bilibili.com/"
	method := resty.MethodGet

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Priority", "u=0, i")
	req.Header.Add("Sec-Ch-Ua", "Not")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0")
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer func() { _ = res.Body.Close() }()

	bili_client := New()
	bili_client.SetCookies(res.Cookies())
	return bili_client
}

// NewWithClient 接收一个自定义的*resty.Client为参数
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
