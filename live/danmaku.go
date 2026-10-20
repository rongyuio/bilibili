package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播间弹幕接口。响应模型见 danmaku_model.go。

// GetLiveDanmakuConfigParam 指定要查询弹幕配置的直播间。
type GetLiveDanmakuConfigParam struct {
	RoomID      int64  `json:"room_id" request:"query"`                          // 直播间 id
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面位置标识。作用尚不明确
}

// GetLiveDanmakuConfig 获取当前用户对应直播间可发弹幕的配置。
//
// 未登录也可调用，此时颜色只有白色可用、模式只有滚动。文档中的 w_rid / wts
// 为 WBI 签名字段且非必要，本方法不主动签名。
func (c *Client) GetLiveDanmakuConfig(ctx context.Context, param GetLiveDanmakuConfigParam) (*GetLiveDanmakuConfigResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/dM/GetDMConfigByGroup"
	)
	return bilibili.Execute[*GetLiveDanmakuConfigResult](ctx, c.Client, method, url, param)
}

// GetLiveRoomHistoryDanmakuParam 指定要查询历史弹幕的直播间。
type GetLiveRoomHistoryDanmakuParam struct {
	RoomID int64 `json:"roomid" request:"query"` // 直播间短 id
}

// GetLiveRoomHistoryDanmaku 获取直播间最近的历史弹幕。
//
// 响应中管理员与普通用户各返回最近 10 条。
func (c *Client) GetLiveRoomHistoryDanmaku(ctx context.Context, param GetLiveRoomHistoryDanmakuParam) (*GetLiveRoomHistoryDanmakuResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/dM/gethistory"
	)
	return bilibili.Execute[*GetLiveRoomHistoryDanmakuResult](ctx, c.Client, method, url, param)
}
