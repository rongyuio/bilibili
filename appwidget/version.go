package appwidget

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// APP 版本相关接口。响应模型见 version_model.go。

// GetVersionInfoParam 是获取版本信息的请求参数。
type GetVersionInfoParam struct {
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 目标客户端类型。留空为最新有更新的平台
}

// GetVersionInfo 获取版本信息
func (c *Client) GetVersionInfo(ctx context.Context, param GetVersionInfoParam) ([]GetVersionInfoItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/version"
	)
	return bilibili.Execute[[]GetVersionInfoItem](ctx, c.Client, method, url, param)
}

// GetVersionUpgradeParam 是版本更新接口的请求参数。
type GetVersionUpgradeParam struct {
	Abi      string `json:"abi,omitempty" request:"query,omitempty"`      // 设备 CPU ABI，例如 arm64-v8a
	AppID    string `json:"appid,omitempty" request:"query,omitempty"`    // 客户端类型，例如 tv.danmaku.bili
	Brand    string `json:"brand,omitempty" request:"query,omitempty"`    // 设备品牌，例如 google
	Build    int    `json:"build" request:"query"`                        // 当前 APP 版本号
	Channel  string `json:"channel" request:"query"`                      // APP 分发渠道
	DeviceID string `json:"deviceid,omitempty" request:"query,omitempty"` // 设备唯一 ID，即 XU-Prefixed BUVID
	Env      string `json:"env,omitempty" request:"query,omitempty"`      // APP 分发环境，默认 prod
	Iv       int    `json:"iv,omitempty" request:"query,omitempty"`       // 当前 APP innerVer
	MobiApp  string `json:"mobi_app" request:"query"`                     // APP mobi_app
	Model    string `json:"model,omitempty" request:"query,omitempty"`    // 设备型号
	Nt       string `json:"nt" request:"query"`                           // 默认 1
	Ov       string `json:"ov" request:"query"`                           // 设备 OS 版本
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 设备平台
	Screen   string `json:"screen,omitempty" request:"query,omitempty"`   // 设备屏幕参数
	Sn       int    `json:"sn" request:"query"`                           // 当前 APP SN
	Vn       string `json:"vn" request:"query"`                           // 当前 APP 版本名称
	AppKey   string `json:"-" request:"-"`                                // 请求头 app-key。APPKey 名称，大体等同于 mobi_app
	Buvid    string `json:"-" request:"-"`                                // 请求头 buvid。XU-Prefixed BUVID
}

// GetVersionUpgrade 版本更新
//
// 除 URL 参数外，接口还要求 app-key 与 buvid 两个请求头，分别由
// GetVersionUpgradeParam.AppKey 与 Buvid 提供。业务错误码：-304 木有改动。
func (c *Client) GetVersionUpgrade(ctx context.Context, param GetVersionUpgradeParam) (*GetVersionUpgradeResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/version/fawkes/upgrade"
	)
	return bilibili.Execute[*GetVersionUpgradeResult](ctx, c.Client, method, url, param, versionUpgradeHeaders(param))
}

// versionUpgradeHeaders 把 app-key 与 buvid 写入请求头；未提供时不覆盖客户端默认值。
func versionUpgradeHeaders(param GetVersionUpgradeParam) bilibili.Handler {
	return func(r *resty.Request) error {
		if param.AppKey != "" {
			r.SetHeader("app-key", param.AppKey)
		}
		if param.Buvid != "" {
			r.SetHeader("buvid", param.Buvid)
		}
		return nil
	}
}

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
