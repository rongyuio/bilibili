package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频观看进度上报接口。

// ReportVideoProgressParam 指定视频观看进度上报参数。
// 字段默认编码进 query，再由内部处理函数移入 URL 编码表单。
type ReportVideoProgressParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必填
	Aid       int    `json:"aid"`                                            // 稿件 avid
	Cid       int    `json:"cid"`                                            // 视频 cid，用于识别分P
	Progress  int    `json:"progress,omitempty" request:"query,omitempty"`   // 观看进度，单位秒，默认为 0
	Platform  string `json:"platform,omitempty" request:"query,omitempty"`   // 平台标识，可为 android
}

// ReportVideoProgress 上报视频观看进度（APP / Cookie 双端）。
//
// 请求体为 x-www-form-urlencoded，业务参数放在表单中。
// 使用 Cookie 登录时，CSRF 取自 Cookie bili_jct 并同样写入表单；
// 填写 AccessKey 时按 APP 方式认证，不再附带 CSRF，因此无 Cookie 的客户端也能调用。
func (c *Client) ReportVideoProgress(ctx context.Context, param ReportVideoProgressParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/report"
	)
	handlers := []bilibili.Handler{
		bilibili.MoveFormParams("access_key", "aid", "cid", "progress", "platform"),
	}
	if param.AccessKey == "" {
		handlers = append(handlers, bilibili.FillFormCsrf(c.Client))
	}
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, handlers...)
	return err
}
