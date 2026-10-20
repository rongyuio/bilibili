package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频在线人数与状态数接口。响应模型见 model_online.go。

// GetVideoOnlineInfoParam 指定要查询在线人数的分P。avid 与 bvid 任选一个。
type GetVideoOnlineInfoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
	Cid  int    `json:"cid"`                                      // 视频cid。用于选择目标分P
}

// GetVideoOnlineInfo 获取视频在线人数
func (c *Client) GetVideoOnlineInfo(ctx context.Context, param GetVideoOnlineInfoParam) (*GetVideoOnlineInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/online/total"
	)
	return bilibili.Execute[*GetVideoOnlineInfoResult](ctx, c.Client, method, url, param)
}

// GetVideoStatusNumberParam 指定要查询状态数的稿件。avid 与 bvid 任选一个。
type GetVideoStatusNumberParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoStatusNumber 获取视频状态数
func (c *Client) GetVideoStatusNumber(ctx context.Context, param GetVideoStatusNumberParam) (*GetVideoStatusNumberResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/stat"
	)
	return bilibili.Execute[*GetVideoStatusNumberResult](ctx, c.Client, method, url, param)
}
