package appwidget

// 版本相关响应模型。各接口的返回体独立定义，互不复用。

// GetAppVersionResult 是拼接最新版本下载地址接口的 data 对象。
type GetAppVersionResult struct {
	CdnURL     string `json:"cdn_url"`      // CDN URL，内容同请求参数 cdn_url
	SoleCdnURL string `json:"sole_cdn_url"` // 带版本的 CDN URL，即 cdn_url + "?v=" + 最新版本
}
