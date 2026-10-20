package miscellaneous

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

var (
	regBv   = regexp.MustCompile(`(?i)bv([\dA-Za-z]{10})`)
	regLive = regexp.MustCompile(`^https://live.bilibili.com/(\d+)`)
)

// errNilContext、errMissingHTTPResponse、errUnresolvedLink 复刻根包请求链路的
// 错误语义；包级静态变量满足 err113。
var (
	errNilContext          = errors.New("request context is nil")
	errMissingHTTPResponse = errors.New("missing HTTP response")
	errUnresolvedLink      = errors.New("无法解析链接")
)

// checkContext 在请求开始前拒绝 nil 或已结束的 context。
func checkContext(ctx context.Context) error {
	if ctx == nil {
		return errNilContext
	}
	return ctx.Err()
}

// safeEndpoint 去掉 URL 中的 query、fragment 与 user 信息，避免错误信息泄漏。
func safeEndpoint(endpoint string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "<invalid URL>"
	}
	u.RawQuery, u.Fragment, u.User = "", "", nil
	return u.String()
}

// UnwrapShortURL 解析短链接，传入一个完整的短链接。
//
// 第一个返回值如果是"bvid"，则第二个返回值是视频的bvid (string)。
// 第一个返回值如果是"live"，则第二个返回值是直播间id (int)。
//
// 该接口的目标地址由调用方提供且只读取 302 Location 头，不返回
// code/message/data 结构：Execute 会拒绝非 200 状态，ExecuteRaw 会把
// 302 当成 *HTTPError，因此这里走 Resty 逃生通道，并自行复刻根包的
// Cookie 快照与响应 Cookie 合并。
func (c *Client) UnwrapShortURL(ctx context.Context, shortURL string) (string, any, error) {
	if err := checkContext(ctx); err != nil {
		return "", nil, err
	}
	r := c.Resty().R().SetContext(ctx).SetCookies(c.GetCookies())
	resp, err := r.Execute(resty.MethodGet, shortURL)
	if resp != nil && resp.RawResponse != nil {
		c.SetCookies(resp.Cookies())
	}
	if err != nil {
		return "", nil, fmt.Errorf("%s %s: %w", resty.MethodGet, safeEndpoint(shortURL), err)
	}
	if resp == nil || resp.RawResponse == nil {
		return "", nil, fmt.Errorf("%s %s: %w", resty.MethodGet, safeEndpoint(shortURL), errMissingHTTPResponse)
	}
	if resp.StatusCode() != http.StatusFound {
		return "", nil, fmt.Errorf("解析短链接失败: %w", &bilibili.HTTPError{
			Method:     resty.MethodGet,
			Endpoint:   safeEndpoint(shortURL),
			StatusCode: resp.StatusCode(),
		})
	}
	location := resp.Header().Get("Location")
	{
		ret := regBv.FindString(location)
		if len(ret) > 0 {
			return "bvid", ret, nil
		}
	}
	{
		ret := regLive.FindStringSubmatch(location)
		if len(ret) > 0 {
			rid, err := strconv.Atoi(ret[1])
			if err != nil {
				return "", nil, err
			}
			return "live", rid, nil
		}
	}
	return "", nil, fmt.Errorf("%w：%s", errUnresolvedLink, location)
}

// nowResult 是服务器当前时间接口的 data 对象。
type nowResult struct {
	Now int64 `json:"now"` // 当前的时间戳
}

// Now 获取服务器当前时间
func (c *Client) Now(ctx context.Context) (time.Time, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/report/click/now"
	)
	result, err := bilibili.Execute[*nowResult](ctx, c.Client, method, url, nil)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(result.Now, 0), nil
}

// Av2Bv 将av号转换为bv号，返回格式为"BV1xxxxxxxxx"。
func Av2Bv(aid int) string {
	const (
		xorCode  = 0x1552356C4CDB
		maxAid   = 1 << 51
		alphabet = "FcwAPNKTMug3GV5Lj7EJnHpWsx4tb8haYeviqBz6rkCy12mUSDQX9RdoZf"
	)
	bvid := []byte("BV1000000000")
	tmp := (maxAid | aid) ^ xorCode
	for _, e := range []int{11, 10, 3, 8, 4, 6, 5, 7, 9} {
		bvid[e] = alphabet[tmp%len(alphabet)]
		tmp /= len(alphabet)
	}
	return string(bvid)
}

// Bv2Av 将bv号转换为av号，传入的bv号格式为"BV1xxxxxxxxx"，前面的"BV"不区分大小写。
func Bv2Av(bvid string) int {
	if len(bvid) != 12 {
		panic("bvid 格式错误: " + bvid)
	}
	const (
		xorCode  = 0x1552356C4CDB
		maskCode = 1<<51 - 1
		alphabet = "FcwAPNKTMug3GV5Lj7EJnHpWsx4tb8haYeviqBz6rkCy12mUSDQX9RdoZf"
	)
	tmp := 0
	for _, e := range []int{9, 7, 5, 6, 4, 8, 3, 10, 11} {
		idx := strings.IndexByte(alphabet, bvid[e])
		tmp = tmp*len(alphabet) + idx
	}
	return (tmp & maskCode) ^ xorCode
}

// GetZoneLocationResult 是通过 IP 确定地理位置接口的 data 对象。
type GetZoneLocationResult struct {
	Addr        string `json:"addr"`         // 公网IP地址
	Country     string `json:"country"`      // 国家/地区名
	Province    string `json:"province"`     // 省/州。非必须存在项
	City        string `json:"city"`         // 城市。非必须存在项
	Isp         string `json:"isp"`          // 运营商名
	Latitude    int    `json:"latitude"`     // 纬度
	Longitude   int    `json:"longitude"`    // 经度
	ZoneID      int    `json:"zone_id"`      // ip数据库id
	CountryCode int    `json:"country_code"` // 国家/地区代码
}

// GetZoneLocation 通过ip确定地理位置
func (c *Client) GetZoneLocation(ctx context.Context) (*GetZoneLocationResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/zone"
	)
	return bilibili.Execute[*GetZoneLocationResult](ctx, c.Client, method, url, nil)
}

// GetRegionDailyCountResult 是分区当日投稿稿件数接口的 data 对象。
type GetRegionDailyCountResult struct {
	RegionCount map[int]int `json:"region_count"` // 分区当日投稿稿件数信息
}

// GetRegionDailyCount 获取分区当日投稿稿件数
func (c *Client) GetRegionDailyCount(ctx context.Context) (*GetRegionDailyCountResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/online"
	)
	return bilibili.Execute[*GetRegionDailyCountResult](ctx, c.Client, method, url, nil)
}
