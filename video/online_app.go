package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频在线人数相关接口。响应模型见 model_online_app.go。

// GetVideoOnlineInfoAppParam 是获取视频在线人数（APP 端）的参数。
type GetVideoOnlineInfoAppParam struct {
	Aid    int    `json:"aid" request:"query"`    // 稿件avid
	Appkey string `json:"appkey" request:"query"` // APP密钥
	Cid    int    `json:"cid" request:"query"`    // 视频cid。用于选择目标分P
	Ts     int64  `json:"ts" request:"query"`     // 当前时间戳
	Sign   string `json:"sign" request:"query"`   // APP签名
}

// GetVideoOnlineInfoApp 获取视频在线人数（APP端）。
//
// 该接口走 APP 域名，除 aid 与 cid 外还需按 APP 方式提供 appkey、ts 与 sign。
// 库只负责发请求，不计算 APP 签名；接口文档标注为未验证。
func (c *Client) GetVideoOnlineInfoApp(ctx context.Context, param GetVideoOnlineInfoAppParam) (*GetVideoOnlineInfoAppResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/view/video/online"
	)
	return bilibili.Execute[*GetVideoOnlineInfoAppResult](ctx, c.Client, method, url, param)
}
