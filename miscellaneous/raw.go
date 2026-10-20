package miscellaneous

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 非 JSON 信封的杂项接口：服务器时间与 LaTeX 渲染。

var (
	errServerDateFormat = errors.New("服务器时间响应格式无法解析")
	errRawJSONDecode    = errors.New("响应不是合法的 JSON")
)

var serverDateRe = regexp.MustCompile(`Date\.UTC\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)`)

// GetServerUTCTime 获取服务器端 UTC 时间。
//
// 接口返回一段 JS（window.serverdate = Date.UTC(Y, M, D, h, m, s)），
// 不是标准的 code/message/data 信封，因此走 ExecuteRaw。
func (c *Client) GetServerUTCTime(ctx context.Context) (time.Time, error) {
	const (
		method = resty.MethodGet
		url    = "https://interface.bilibili.com/serverdate.js"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, url, nil)
	if err != nil {
		return time.Time{}, err
	}
	m := serverDateRe.FindSubmatch(body)
	if m == nil {
		return time.Time{}, errServerDateFormat
	}
	parts := make([]int, 6)
	for i := range parts {
		parts[i], err = strconv.Atoi(string(m[i+1]))
		if err != nil {
			return time.Time{}, errServerDateFormat
		}
	}
	// Date.UTC 的月份从 0 开始。
	return time.Date(parts[0], time.Month(parts[1]+1), parts[2], parts[3], parts[4], parts[5], 0, time.UTC), nil
}

// LatexParam 是 LaTeX 渲染的请求参数。
type LatexParam struct {
	Formula string `json:"formula,omitempty" request:"query,omitempty"` // LaTeX 公式
}

// RenderLatex 将 LaTeX 公式渲染为 SVG 图片。
//
// 响应是完整的 image/svg+xml 文件，可直接用于 HTML 的 <img> 或 Markdown。
func (c *Client) RenderLatex(ctx context.Context, param LatexParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-frontend/mathjax/tex"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
}

// GetVideoPbpParam 指定要高能进度条的稿件。
type GetVideoPbpParam struct {
	Aid  int    `json:"aid" request:"query"`                      // 稿件avid
	Cid  int    `json:"cid" request:"query"`                      // 视频cid
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid
	R    string `json:"r" request:"query"`                        // 加载来源。固定为 loader
}

// GetVideoPbpModule 是高能进度条响应中的一个模块。
type GetVideoPbpModule struct {
	Name      string            `json:"name"`       // 模块名。固定为 pbp
	Version   string            `json:"version"`    // 脚本版本
	ScriptSrc string            `json:"script_src"` // 脚本地址
	LoadMode  string            `json:"load_mode"`  // 加载模式。固定为 pbp
	Params    GetVideoPbpParams `json:"params"`     // 参数本体
}

// GetVideoPbpParams 是高能进度条模块的参数本体。
type GetVideoPbpParams struct {
	Data GetVideoPbpData `json:"data"` // 数据本体
}

// GetVideoPbpData 是高能进度条的数据本体。
type GetVideoPbpData struct {
	StepSec int               `json:"step_sec"` // 采样间隔时间。单位为秒，由视频时长决定
	Tagstr  string            `json:"tagstr"`   // 作用尚不明确
	Events  GetVideoPbpEvents `json:"events"`   // 数据本体
	Debug   string            `json:"debug"`    // 调试信息。json 字串
}

// GetVideoPbpEvents 是高能进度条的弹幕趋势顶点列表。
type GetVideoPbpEvents struct {
	Default []int `json:"default"` // 顶点值列表
}

// GetVideoPbpResult 是高能进度条接口的根响应，没有 code/message/data 信封。
type GetVideoPbpResult struct {
	Modules []GetVideoPbpModule `json:"modules"` // 模块列表。仅含1个元素
}

// GetVideoPbp 获取视频高能进度条的弹幕趋势顶点列表。
func (c *Client) GetVideoPbp(ctx context.Context, param GetVideoPbpParam) (*GetVideoPbpResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://bvc.bilivideo.com/pbp/data"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
	if err != nil {
		return nil, err
	}
	var out GetVideoPbpResult
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, errRawJSONDecode
	}
	return &out, nil
}

// GetLiveTimestampResult 是 RTC 时间戳接口的响应。
type GetLiveTimestampResult struct {
	Timestamp int64 `json:"timestamp"` // 服务器端UTC时间戳
	Microtime int64 `json:"microtime"` // 服务器端UTC时间戳。毫秒
}

// GetLiveTimestamp 获取适用于 RTC 的服务器时间戳。
func (c *Client) GetLiveTimestamp(ctx context.Context) (*GetLiveTimestampResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/open-interface/v1/rtc/getTimestamp"
	)
	return bilibili.Execute[*GetLiveTimestampResult](ctx, c.Client, method, url, nil)
}
