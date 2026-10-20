package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频 TAG（新版）相关接口。响应模型见 tags_model.go。

// GetVideoDetailTagsParam 是获取视频TAG信息的参数。
type GetVideoDetailTagsParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
	Cid  int    `json:"cid,omitempty" request:"query,omitempty"`  // 分P cid。提供此参数可返回对应分P的BGM信息
}

// GetVideoDetailTags 获取视频TAG信息（新）
//
// 与 GetVideoTags（旧接口 /x/tag/archive/tags）不同，本接口除普通标签外还会返回
// 话题与背景音乐类型的 TAG。
func (c *Client) GetVideoDetailTags(ctx context.Context, param GetVideoDetailTagsParam) ([]GetVideoDetailTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/view/detail/tag"
	)
	return bilibili.Execute[[]GetVideoDetailTagsItem](ctx, c.Client, method, url, param)
}
