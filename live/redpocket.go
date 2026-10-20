package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播间人气红包相关接口。响应模型见 redpocket_model.go。

// GetLiveRedPocketInfoParam 指定要查询红包信息的直播间。
type GetLiveRedPocketInfoParam struct {
	RoomID string `json:"roomid"` // 直播间 id
}

// GetLiveRedPocketInfo 获取指定直播间的红包信息。
//
// 认证方式为 Cookie（SESSDATA），可选：不携带 Cookie 时 user_status 恒为 2（未参与）。
// 直播间无红包时 popularity_red_pocket 为 null。
func (c *Client) GetLiveRedPocketInfo(ctx context.Context, param GetLiveRedPocketInfoParam) (*GetLiveRedPocketInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/lottery-interface/v1/lottery/getLotteryInfoWeb"
	)
	return bilibili.Execute[*GetLiveRedPocketInfoResult](ctx, c.Client, method, url, param, liveMiscReferer())
}
