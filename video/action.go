package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频交互相关接口。响应模型见 model_action.go。

// TagParam 指定要点赞或点踩的视频 TAG。
type TagParam struct {
	Aid   int `json:"aid"`    // 稿件avid
	TagID int `json:"tag_id"` // tag_id
}

// LikeVideoTag 点赞视频TAG，重复请求为取消
func (c *Client) LikeVideoTag(ctx context.Context, param TagParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/tag/archive/like2"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// HateVideoTag 点踩视频TAG，重复访问为取消
func (c *Client) HateVideoTag(ctx context.Context, param TagParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/tag/archive/hate2"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// LikeVideoParam 指定点赞操作。avid 与 bvid 任选一个。
type LikeVideoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件 avid。avid 与 bvid 任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件 bvid。avid 与 bvid 任选一个
	Like int    `json:"like"`                                     // 操作方式。1：点赞。2：取消赞
}

// LikeVideo 点赞视频
func (c *Client) LikeVideo(ctx context.Context, param LikeVideoParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/archive/like"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// CoinVideoParam 指定投币操作。avid 与 bvid 任选一个。
type CoinVideoParam struct {
	Aid        int    `json:"aid,omitempty" request:"query,omitempty"`         // 稿件 avid。avid 与 bvid 任选一个
	Bvid       string `json:"bvid,omitempty" request:"query,omitempty"`        // 稿件 bvid。avid 与 bvid 任选一个
	Multiply   int    `json:"multiply"`                                        // 投币数量。上限为2
	SelectLike int    `json:"select_like,omitempty" request:"query,omitempty"` // 是否附加点赞。0：不点赞。1：同时点赞。默认为0
}

// CoinVideo 投币视频
func (c *Client) CoinVideo(ctx context.Context, param CoinVideoParam) (*CoinVideoResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/coin/add"
	)
	return bilibili.Execute[*CoinVideoResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// FavourVideoParam 指定收藏操作。
type FavourVideoParam struct {
	Rid         int   `json:"rid"`                                               // 稿件 avid
	Type        int   `json:"type"`                                              // 必须为2
	AddMediaIDs []int `json:"add_media_ids,omitempty" request:"query,omitempty"` // 需要加入的收藏夹 mlid。同时添加多个，用,（%2C）分隔
	DelMediaIDs []int `json:"del_media_ids,omitempty" request:"query,omitempty"` // 需要取消的收藏夹 mlid。同时取消多个，用,（%2C）分隔
}

// FavourVideo 收藏视频
func (c *Client) FavourVideo(ctx context.Context, param FavourVideoParam) (*FavourVideoResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/medialist/gateway/coll/resource/deal"
	)
	return bilibili.Execute[*FavourVideoResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// LikeCoinFavourVideoParam 指定一键三连的稿件。avid 与 bvid 任选一个。
type LikeCoinFavourVideoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// LikeCoinFavourVideo 一键三连视频
func (c *Client) LikeCoinFavourVideo(ctx context.Context, param LikeCoinFavourVideoParam) (*LikeCoinFavourVideoResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/archive/like/triple"
	)
	return bilibili.Execute[*LikeCoinFavourVideoResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}
