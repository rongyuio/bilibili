package opus

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 图文详情相关接口。响应模型见 opus_detail_model.go。

// GetOpusDetailParam 指定要查询的图文及其可选参数。
type GetOpusDetailParam struct {
	ID             string `json:"id" request:"query"`                        // 动态 id，必填，纯数字字符串
	TimezoneOffset int    `json:"timezone_offset" request:"query,omitempty"` // 时区偏移，例如 -480
	Features       string `json:"features" request:"query,omitempty"`        // 功能开关，部分专栏要求包含 htmlNewStyle
}

// GetOpusDetail 获取图文详细信息。
//
// 鉴权方式：Cookie 中 buvid3 存在且不为空。
// 旧版专栏在 features 未包含 htmlNewStyle 时可能只返回 data.fallback，
// 此时 Item 为空而 Fallback 有值。
func (c *Client) GetOpusDetail(ctx context.Context, param GetOpusDetailParam) (*GetOpusDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/opus/detail"
	)
	return bilibili.Execute[*GetOpusDetailResult](ctx, c.Client, method, url, param)
}
