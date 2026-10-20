package appwidget

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// APP 端开屏相关接口。响应模型见 splash_model.go。

// GetSplashBrandListParam 是获取 APP 端开屏图片列表的请求参数。
type GetSplashBrandListParam struct {
	AppKey string `json:"appkey,omitempty" request:"query,omitempty"` // APP 密钥。APP 方式必要
	Ts     int64  `json:"ts,omitempty" request:"query,omitempty"`     // 当前时间戳。APP 方式必要
	Sign   string `json:"sign,omitempty" request:"query,omitempty"`   // APP 签名。APP 方式必要
}

// GetSplashBrandList 获取 APP 端开屏图片列表
//
// 该接口走 APP 域名，需按 APP 方式提供 appkey、ts 与 sign。
// 库只负责发请求，不计算 APP 签名；接口文档标注为未验证。
func (c *Client) GetSplashBrandList(ctx context.Context, param GetSplashBrandListParam) (*GetSplashBrandListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/splash/brand/list"
	)
	return bilibili.Execute[*GetSplashBrandListResult](ctx, c.Client, method, url, param)
}

// GetSplashListParam 是获取 APP 端开屏广告信息的请求参数。
type GetSplashListParam struct {
	Build    int    `json:"build" request:"query"`    // 客户端内部版本号
	MobiApp  string `json:"mobi_app" request:"query"` // android、iphone、ipad
	Platform string `json:"platform" request:"query"` // android、ios
	Height   int    `json:"height" request:"query"`   // 屏幕高度
	Width    int    `json:"width" request:"query"`    // 屏幕宽度
	Birth    string `json:"birth" request:"query"`    // 生日日期，四位数，例如 0101
}

// GetSplashList 获取 APP 端开屏广告信息
//
// 该接口走 APP 域名，参数按 APP 端填写；接口文档标注为未验证。
func (c *Client) GetSplashList(ctx context.Context, param GetSplashListParam) (*GetSplashListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/splash/list"
	)
	return bilibili.Execute[*GetSplashListResult](ctx, c.Client, method, url, param)
}
