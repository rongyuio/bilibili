package newbieexam

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 新用户答题操作相关接口。响应模型见 action_model.go。

// GetNewbieExamCaptcha 获取新用户答题验证码。
//
// 该接口没有请求参数，需要登录（Cookie SESSDATA 或 APP 认证）。
// 返回的 gt 与 challenge 用于完成极验人机验证，再提交验证码。
func (c *Client) GetNewbieExamCaptcha(ctx context.Context) (*GetNewbieExamCaptchaResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/captcha"
	)
	return bilibili.Execute[*GetNewbieExamCaptchaResult](ctx, c.Client, method, url, nil)
}
