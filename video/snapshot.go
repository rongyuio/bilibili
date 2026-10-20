package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频快照相关接口。响应模型见 snapshot_model.go。

// GetVideoSnapshotParam 是获取视频快照（web 端）的参数。
type GetVideoSnapshotParam struct {
	Aid   int    `json:"aid,omitempty" request:"query,omitempty"`   // 稿件avid。avid 与 bvid 任选一个
	Bvid  string `json:"bvid,omitempty" request:"query,omitempty"`  // 稿件bvid。avid 与 bvid 任选一个
	Cid   int    `json:"cid,omitempty" request:"query,omitempty"`   // 分P cid。默认为1P
	Index int    `json:"index,omitempty" request:"query,omitempty"` // 是否返回 json 数组格式的截取时间表。1：需要。0：不需要。默认为0
}

// GetVideoSnapshot 获取视频快照（web端）。
//
// 快照的截取时间根据视频画面变化程度决定，各视频不相同；截取时间表与快照一一对应，
// 按从左到右、从上到下的顺序排布。
func (c *Client) GetVideoSnapshot(ctx context.Context, param GetVideoSnapshotParam) (*GetVideoSnapshotResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/videoshot"
	)
	return bilibili.Execute[*GetVideoSnapshotResult](ctx, c.Client, method, url, param)
}
