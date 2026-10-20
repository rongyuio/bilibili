package appwidget

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 版本相关接口。响应模型见 version_model.go。

// GetAppVersionParam 是拼接最新版本下载地址的请求参数。
type GetAppVersionParam struct {
	Target string `json:"target,omitempty" request:"query,omitempty"`  // 目标平台。目测仅 pc_electron，错误的值会导致 HTTP 404
	CdnURL string `json:"cdn_url,omitempty" request:"query,omitempty"` // CDN URL，任意字符串
}

// GetAppVersion 拼接最新版本下载地址
func (c *Client) GetAppVersion(ctx context.Context, param GetAppVersionParam) (*GetAppVersionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-frontend/getappversion"
	)
	return bilibili.Execute[*GetAppVersionResult](ctx, c.Client, method, url, param)
}
