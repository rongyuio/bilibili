package login

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 登录操作（人机验证）相关接口。响应模型见 login_action_readme_model.go。
//
// 人机验证流程：先申请 captcha 拿到 gt 与 challenge，完成极验验证后取得
// validate 与 seccode，再用于密码登录或短信验证码登录。

// GetOldCaptchaParam 指定申请旧版 captcha 验证码的参数。
type GetOldCaptchaParam struct {
	Plat int `json:"plat" request:"query"` // 平台类型。默认为 6
}

// GetOldCaptcha 申请 captcha 验证码（旧版）。
//
// 该接口曾从文档移除，实测仍可正常使用。新代码优先使用 Captcha。
func (c *Client) GetOldCaptcha(ctx context.Context, param GetOldCaptchaParam) (*GetOldCaptchaResult, error) {
	const (
		method = resty.MethodGet
		url    = "http://passport.bilibili.com/web/captcha/combine" //nolint:revive // 接口地址原样保留
	)
	return bilibili.Execute[*GetOldCaptchaResult](ctx, c.Client, method, url, param)
}
