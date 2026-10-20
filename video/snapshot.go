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

// GetVideoSnapshotAppParam 是获取视频快照（APP 端）的参数。
type GetVideoSnapshotAppParam struct {
	Aid    int    `json:"aid" request:"query"`    // 稿件avid
	Appkey string `json:"appkey" request:"query"` // APP密钥
	Cid    int    `json:"cid" request:"query"`    // 分P cid
	Ts     int64  `json:"ts" request:"query"`     // 当前时间戳
	Sign   string `json:"sign" request:"query"`   // APP签名
}

// GetVideoSnapshotApp 获取视频快照（APP端）。
//
// 该接口走 APP 域名，除 aid 与 cid 外还需按 APP 方式提供 appkey、ts 与 sign。
func (c *Client) GetVideoSnapshotApp(ctx context.Context, param GetVideoSnapshotAppParam) (*GetVideoSnapshotAppResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/view/video/shot"
	)
	return bilibili.Execute[*GetVideoSnapshotAppResult](ctx, c.Client, method, url, param)
}
