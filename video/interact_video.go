package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 互动视频相关接口。响应模型见 model_interact_video.go。

// GetInteractVideoEdgeInfoParam 指定要查询的互动视频模块。avid 与 bvid 任选一个。
type GetInteractVideoEdgeInfoParam struct {
	Aid          int    `json:"aid,omitempty" request:"query,omitempty"`     // 稿件avid。avid与bvid任选一个
	Bvid         string `json:"bvid,omitempty" request:"query,omitempty"`    // 稿件bvid。avid与bvid任选一个
	GraphVersion int    `json:"graph_version"`                               // 剧情图id。可于播放器接口的 interaction 对象取得
	EdgeID       int    `json:"edge_id,omitempty" request:"query,omitempty"` // 模块编号。0或留空为起始模块
}

// GetInteractVideoEdgeInfo 获取互动视频模块详细信息
func (c *Client) GetInteractVideoEdgeInfo(ctx context.Context, param GetInteractVideoEdgeInfoParam) (*GetInteractVideoEdgeInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/stein/edgeinfo_v2"
	)
	return bilibili.Execute[*GetInteractVideoEdgeInfoResult](ctx, c.Client, method, url, param)
}
