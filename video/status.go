package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频互动状态相关接口（点赞、投币、收藏的查询与收藏操作）。响应模型见 model_status.go。

// GetVideoLikeStatusParam 指定要查询点赞状态的稿件。avid 与 bvid 任选一个。
type GetVideoLikeStatusParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Aid       int    `json:"aid,omitempty" request:"query,omitempty"`        // 稿件 avid。avid 与 bvid 任选一个
	Bvid      string `json:"bvid,omitempty" request:"query,omitempty"`       // 稿件 bvid。avid 与 bvid 任选一个
}

// GetVideoLikeStatus 判断视频近期是否被点赞。
// 该接口只能判断近期内的点赞状态，不能反映全部历史点赞，详见接口文档。
// 返回 0：未点赞，1：已点赞。
func (c *Client) GetVideoLikeStatus(ctx context.Context, param GetVideoLikeStatusParam) (int, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/has/like"
	)
	return bilibili.Execute[int](ctx, c.Client, method, url, param)
}

// GetVideoCoinStatusParam 指定要查询投币状态的稿件。avid 与 bvid 任选一个。
type GetVideoCoinStatusParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Aid       int    `json:"aid,omitempty" request:"query,omitempty"`        // 稿件 avid。avid 与 bvid 任选一个
	Bvid      string `json:"bvid,omitempty" request:"query,omitempty"`       // 稿件 bvid。avid 与 bvid 任选一个
}

// GetVideoCoinStatus 判断视频是否被投币
func (c *Client) GetVideoCoinStatus(ctx context.Context, param GetVideoCoinStatusParam) (*GetVideoCoinStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/coins"
	)
	return bilibili.Execute[*GetVideoCoinStatusResult](ctx, c.Client, method, url, param)
}

// FavourVideoWebParam 指定 Web 端收藏操作。
type FavourVideoWebParam struct {
	Rid         int    `json:"rid"`                                               // 稿件 avid
	Type        int    `json:"type"`                                              // 必须为 2
	AddMediaIDs []int  `json:"add_media_ids,omitempty" request:"query,omitempty"` // 需要加入的收藏夹 mlid。同时添加多个，用 ,（%2C）分隔
	DelMediaIDs []int  `json:"del_media_ids,omitempty" request:"query,omitempty"` // 需要取消的收藏夹 mlid。同时取消多个，用 ,（%2C）分隔
	Platform    string `json:"platform,omitempty" request:"query,omitempty"`      // 平台标识。web 端：web
	EabX        int    `json:"eab_x,omitempty" request:"query,omitempty"`         // 1。作用尚不明确
	Ramval      int    `json:"ramval,omitempty" request:"query,omitempty"`        // 正整数。可能与在该页面的停留时间相关
	Ga          int    `json:"ga,omitempty" request:"query,omitempty"`            // 1。作用尚不明确
	GaiaSource  string `json:"gaia_source,omitempty" request:"query,omitempty"`   // web 端：web_normal
}

// FavourVideoWeb 收藏视频（Web 端）。
// 请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) FavourVideoWeb(ctx context.Context, param FavourVideoWebParam) (*FavourVideoWebResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/deal"
	)
	return bilibili.Execute[*FavourVideoWebResult](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("rid", "type", "add_media_ids", "del_media_ids", "platform", "eab_x", "ramval", "ga", "gaia_source"),
		bilibili.FillFormCsrf(c.Client))
}

// GetVideoFavouredParam 指定要查询收藏状态的稿件。
type GetVideoFavouredParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Aid       string `json:"aid"`                                            // 稿件 avid 或稿件 bvid
}

// GetVideoFavoured 判断视频是否被收藏
func (c *Client) GetVideoFavoured(ctx context.Context, param GetVideoFavouredParam) (*GetVideoFavouredResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/fav/video/favoured"
	)
	return bilibili.Execute[*GetVideoFavouredResult](ctx, c.Client, method, url, param)
}
