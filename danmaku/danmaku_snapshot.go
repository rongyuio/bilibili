package danmaku

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 弹幕快照接口。

// GetDanmakuSnapshotParam 指定要查询快照的稿件。
type GetDanmakuSnapshotParam struct {
	Aid string `json:"aid" request:"query"` // 稿件 avid 或稿件 bvid
}

// GetDanmakuSnapshot 获取弹幕快照。
//
// 返回最近产生的几条弹幕内容，最多 20 条。
func (c *Client) GetDanmakuSnapshot(ctx context.Context, param GetDanmakuSnapshotParam) ([]string, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/ajax"
	)
	return bilibili.Execute[[]string](ctx, c.Client, method, url, param)
}
