package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播礼物相关接口。响应模型见 gift_model.go。

// GetLiveRoomGiftListParam 指定要查询礼物面板的直播间及其分区。
type GetLiveRoomGiftListParam struct {
	Platform     string `json:"platform"`                                 // 平台标识，例如 web
	RoomID       int    `json:"room_id"`                                  // 主播房间号
	AreaParentID int    `json:"area_parent_id" request:"query,omitempty"` // 直播分区；不填写可能会获取不到部分活动礼物
	AreaID       int    `json:"area_id" request:"query,omitempty"`        // 直播子分区；不填写可能会获取不到部分活动礼物
}

// GetLiveRoomGiftList 获取直播间内可投喂的礼物列表。
func (c *Client) GetLiveRoomGiftList(ctx context.Context, param GetLiveRoomGiftListParam) (*GetLiveRoomGiftListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/giftPanel/roomGiftList"
	)
	return bilibili.Execute[*GetLiveRoomGiftListResult](ctx, c.Client, method, url, param)
}

// GetLiveBlindBoxProbabilityParam 指定要查询概率的盲盒礼物。
type GetLiveBlindBoxProbabilityParam struct {
	GiftID int `json:"gift_id"` // 盲盒对应的礼物id
}

// GetLiveBlindBoxProbability 获取盲盒各档奖品的爆出概率。
func (c *Client) GetLiveBlindBoxProbability(ctx context.Context, param GetLiveBlindBoxProbabilityParam) (*GetLiveBlindBoxProbabilityResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/general-interface/v1/blindFirstWin/getInfo"
	)
	return bilibili.Execute[*GetLiveBlindBoxProbabilityResult](ctx, c.Client, method, url, param)
}
