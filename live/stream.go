package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播间视频流相关接口。响应模型见 stream_model.go。

// GetLiveStreamURLParam 指定要获取视频流的真实直播间。
type GetLiveStreamURLParam struct {
	Cid      int64  `json:"cid" request:"query"`                          // 目标真实直播间号，即直播间的 room_id（非短号）
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 直播流格式。h5：hls 方式；web：http-flv 方式；默认为 http-flv 方式
	Quality  int    `json:"quality,omitempty" request:"query,omitempty"`  // 画质。与 qn 任选其一。2：流畅；3：高清；4：原画
	Qn       string `json:"qn,omitempty" request:"query,omitempty"`       // 画质。与 quality 任选其一。80：流畅；150：高清；250：超清；400：蓝光；10000：原画；20000：4K；25000：默认；30000：杜比
}

// GetLiveStreamURL 根据真实直播间号获取直播视频流。
//
// cid 必须传真实直播间号（room_id），短号需要先通过 GetLiveRoomInfo 换算。
// 接口属于直播域，请求会附带直播间页面的 Referer / Origin。
func (c *Client) GetLiveStreamURL(ctx context.Context, param GetLiveStreamURLParam) (*GetLiveStreamURLResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Room/playUrl"
	)
	return bilibili.Execute[*GetLiveStreamURLResult](ctx, c.Client, method, url, param, liveReferer(param.Cid))
}
