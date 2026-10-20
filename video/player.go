package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// web 播放器信息接口。响应模型见 player_model.go。

// GetVideoPlayerInfoParam 是获取 web 播放器信息的参数。
type GetVideoPlayerInfoParam struct {
	Aid      int64  `json:"aid,omitempty" request:"query,omitempty"`       // 稿件 avid。aid 与 bvid 任选一个
	Bvid     string `json:"bvid,omitempty" request:"query,omitempty"`      // 稿件 bvid。aid 与 bvid 任选一个
	Cid      int64  `json:"cid"`                                           // 稿件 cid
	SeasonID int64  `json:"season_id,omitempty" request:"query,omitempty"` // 番剧 season_id
	EpID     int64  `json:"ep_id,omitempty" request:"query,omitempty"`     // 剧集 ep_id
}

// GetVideoPlayerInfo 获取 web 播放器信息。
//
// 提供正常播放需要的元数据，包括智能防挡弹幕、字幕、章节看点等。
// 接口需要 WBI 签名，w_rid 与 wts 由客户端自动追加。
func (c *Client) GetVideoPlayerInfo(ctx context.Context, param GetVideoPlayerInfoParam) (*GetVideoPlayerInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/wbi/v2"
	)
	return bilibili.Execute[*GetVideoPlayerInfoResult](ctx, c.Client, method, url, param, c.FillWbi())
}
