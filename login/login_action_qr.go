package login

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 登录相关接口（web 端旧版扫码登录）。响应模型见 login_action_qr_model.go。

// GetQRCodeOld 申请二维码（web 端旧版）。
// 密钥超时为 180 秒。返回值中的 URL 用于生成二维码，OAuthKey 需保存备用。
func (c *Client) GetQRCodeOld(ctx context.Context) (*GetQRCodeOldResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/qrcode/getLoginUrl"
	)
	return bilibili.Execute[*GetQRCodeOldResult](ctx, c.Client, method, url, nil)
}
