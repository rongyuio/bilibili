package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频信息相关接口。响应模型见 model_info.go。

// GetVideoInfoParam 指定要查询的稿件。avid 与 bvid 任选一个。
type GetVideoInfoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoInfo 获取视频详细信息
func (c *Client) GetVideoInfo(ctx context.Context, param GetVideoInfoParam) (*GetVideoInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/view"
	)
	return bilibili.Execute[*GetVideoInfoResult](ctx, c.Client, method, url, param)
}

// GetVideoDetailInfoParam 指定要查询的稿件。avid 与 bvid 任选一个。
type GetVideoDetailInfoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoDetailInfo 获取视频超详细信息
func (c *Client) GetVideoDetailInfo(ctx context.Context, param GetVideoDetailInfoParam) (*GetVideoDetailInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/view/detail"
	)
	return bilibili.Execute[*GetVideoDetailInfoResult](ctx, c.Client, method, url, param)
}

// GetVideoRecommendListParam 指定要查询推荐列表的稿件。avid 与 bvid 任选一个。
type GetVideoRecommendListParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoRecommendList 获取单视频推荐列表
func (c *Client) GetVideoRecommendList(ctx context.Context, param GetVideoRecommendListParam) ([]GetVideoRecommendListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/related"
	)
	return bilibili.Execute[[]GetVideoRecommendListItem](ctx, c.Client, method, url, param)
}

// GetVideoPageListParam 指定要查询分P列表的稿件。avid 与 bvid 任选一个。
type GetVideoPageListParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoPageList 获取视频分P列表
func (c *Client) GetVideoPageList(ctx context.Context, param GetVideoPageListParam) ([]GetVideoPageListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/pagelist"
	)
	return bilibili.Execute[[]GetVideoPageListItem](ctx, c.Client, method, url, param)
}

// GetVideoTagsParam 指定要查询 TAG 的稿件。avid 与 bvid 任选一个。
type GetVideoTagsParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoTags 获取视频TAG
func (c *Client) GetVideoTags(ctx context.Context, param GetVideoTagsParam) ([]GetVideoTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/tag/archive/tags"
	)
	return bilibili.Execute[[]GetVideoTagsItem](ctx, c.Client, method, url, param)
}
