package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播间表情相关接口。响应模型见 emoticons_model.go。

// GetLiveEmoticonsParam 指定要查询表情包的直播间与平台。
type GetLiveEmoticonsParam struct {
	Platform string `json:"platform" request:"query"` // 平台，已知：pc、android、ios
	RoomID   int    `json:"room_id" request:"query"`  // 直播间长号，只支持 room_id，不支持 short_id
}

// GetLiveEmoticons 获取指定直播间的表情包列表。
//
// 认证方式为 Cookie（SESSDATA）；room_id 必须是直播间长号。
func (c *Client) GetLiveEmoticons(ctx context.Context, param GetLiveEmoticonsParam) (*GetLiveEmoticonsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-ucenter/v2/emoticon/GetEmoticons"
	)
	return bilibili.Execute[*GetLiveEmoticonsResult](ctx, c.Client, method, url, param)
}
