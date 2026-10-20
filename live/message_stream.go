package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 信息流认证秘钥接口。响应模型见 message_stream_model.go。

// GetDanmuInfoParam 指定获取信息流认证秘钥的参数。
type GetDanmuInfoParam struct {
	ID          int64  `json:"id"`                                               // 直播间真实 id
	Type        int    `json:"type,omitempty" request:"query,omitempty"`         // (?)。作用尚不明确
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)。作用尚不明确
}

// GetDanmuInfo 获取信息流认证秘钥。
// 返回信息流服务器节点列表与 token，用于建立直播弹幕的 WebSocket / TCP 连接。
// 该接口自 2025-05-26 起强制 WBI 签名，自 2025-06-27 起要求 Cookie 中存在 buvid3；
// 未认证视作未登录，部分消息会受到隐私保护限制。
func (c *Client) GetDanmuInfo(ctx context.Context, param GetDanmuInfoParam) (*GetDanmuInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
	)
	return bilibili.Execute[*GetDanmuInfoResult](ctx, c.Client, method, url, param, c.FillWbi())
}
