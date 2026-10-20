package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态首页公告栏接口。响应模型见 dynamic_banner_model.go。

// GetDynamicBannerParam 指定新版公告栏的查询条件。
type GetDynamicBannerParam struct {
	Platform    int    `json:"platform,omitempty" request:"query,omitempty"`     // 平台。默认为 1，留空无
	Position    string `json:"position,omitempty" request:"query,omitempty"`     // 位置。默认为 web动态，留空无
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面标识。固定为 333.1365
}

// GetDynamicBanner 获取动态首页公告栏（新版）
func (c *Client) GetDynamicBanner(ctx context.Context, param GetDynamicBannerParam) (*GetDynamicBannerResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/dynamic/feed/dyn/banner"
	)
	return bilibili.Execute[*GetDynamicBannerResult](ctx, c.Client, method, url, param)
}

// GetDynamicBannerLegacyParam 指定旧版公告栏的查询条件。
type GetDynamicBannerLegacyParam struct {
	Platform int    `json:"platform,omitempty" request:"query,omitempty"` // 平台。默认为 1
	Position string `json:"position,omitempty" request:"query,omitempty"` // 位置。默认为 web动态
}

// GetDynamicBannerLegacy 获取动态首页公告栏（旧版）
func (c *Client) GetDynamicBannerLegacy(ctx context.Context, param GetDynamicBannerLegacyParam) (*GetDynamicBannerLegacyResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/banner_svr/v1/banner_svr/show"
	)
	return bilibili.Execute[*GetDynamicBannerLegacyResult](ctx, c.Client, method, url, param)
}
