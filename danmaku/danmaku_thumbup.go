package danmaku

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 弹幕点赞查询接口。返回体独立定义，不与其他接口复用。

// GetDanmakuThumbupStatsParam 指定要查询点赞数的弹幕。
type GetDanmakuThumbupStatsParam struct {
	Oid int64   `json:"oid" request:"query"` // 视频 cid
	IDs []int64 `json:"ids" request:"query"` // 欲查询的弹幕 dmid，多个之间用 , 分隔
}

// GetDanmakuThumbupStats 查询指定 dmid 的点赞数
func (c *Client) GetDanmakuThumbupStats(ctx context.Context, param GetDanmakuThumbupStatsParam) (map[string]GetDanmakuThumbupStatsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/thumbup/stats"
	)
	return bilibili.Execute[map[string]GetDanmakuThumbupStatsItem](ctx, c.Client, method, url, param)
}
