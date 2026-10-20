package article

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 专栏相关接口。每个接口的响应模型见对应的 model_*.go。

// GetArticlesInfoParam 是获取文集基本信息的参数。
type GetArticlesInfoParam struct {
	ID int `json:"id"` // 文集rlid
}

// GetArticlesInfo 获取文集基本信息
func (c *Client) GetArticlesInfo(ctx context.Context, param GetArticlesInfoParam) (*GetArticlesInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/list/web/articles"
	)
	return bilibili.Execute[*GetArticlesInfoResult](ctx, c.Client, method, url, param)
}

// GetArticleInfoParam 是获取专栏文章基本信息的参数。
type GetArticleInfoParam struct {
	ID int `json:"id"` // 专栏cvid
}

// GetArticleInfo 获取专栏文章基本信息
func (c *Client) GetArticleInfo(ctx context.Context, param GetArticleInfoParam) (*GetArticleInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/viewinfo"
	)
	return bilibili.Execute[*GetArticleInfoResult](ctx, c.Client, method, url, param)
}

// LikeArticleParam 是点赞文章的参数。
type LikeArticleParam struct {
	ID   int `json:"id"`   // 文章cvid
	Type int `json:"type"` // 操作方式。1：点赞。2：取消赞
}

// LikeArticle 点赞文章
func (c *Client) LikeArticle(ctx context.Context, param LikeArticleParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/article/like"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// CoinArticleParam 是投币文章的参数。
type CoinArticleParam struct {
	Aid      int `json:"aid"`      // 文章cvid
	Upid     int `json:"upid"`     // 文章作者mid
	Multiply int `json:"multiply"` // 投币数量。上限为2
	Avtype   int `json:"avtype"`   // 2。必须为2
}

// CoinArticle 投币文章
func (c *Client) CoinArticle(ctx context.Context, param CoinArticleParam) (*CoinArticleResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/coin/add"
	)
	return bilibili.Execute[*CoinArticleResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// FavoritesArticleParam 是收藏文章的参数。
type FavoritesArticleParam struct {
	ID int `json:"id"` // 文章cvid
}

// FavoritesArticle 收藏文章
func (c *Client) FavoritesArticle(ctx context.Context, param FavoritesArticleParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/article/favorites/add"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// GetUserArticleListParam 是获取用户专栏文章列表的参数。
type GetUserArticleListParam struct {
	Mid  int    `json:"mid"`                                      // 用户uid
	Pn   int    `json:"pn,omitempty" request:"query,omitempty"`   // 默认：1
	Ps   int    `json:"ps,omitempty" request:"query,omitempty"`   // 默认：30。范围：[1,30]
	Sort string `json:"sort,omitempty" request:"query,omitempty"` // publish_time：最新发布。view：最多阅读。fav：最多收藏。默认：publish_time
}

// GetUserArticleList 获取用户专栏文章列表
func (c *Client) GetUserArticleList(ctx context.Context, param GetUserArticleListParam) (*GetUserArticleListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/article"
	)
	return bilibili.Execute[*GetUserArticleListResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// GetUserArticlesListParam 是获取用户专栏文集列表的参数。
type GetUserArticlesListParam struct {
	Mid      int    `json:"mid"`                                      // 用户uid
	Sort     int    `json:"sort,omitempty" request:"query,omitempty"` // 排序方式。0：最近更新。1：最多阅读
	Jsonp    string `json:"jsonp,omitempty" request:"query,omitempty"`
	Callback string `json:"callback,omitempty" request:"query,omitempty"`
}

// GetUserArticlesList 获取用户专栏文集列表
func (c *Client) GetUserArticlesList(ctx context.Context, param GetUserArticlesListParam) (*GetUserArticlesListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/up/lists"
	)
	return bilibili.Execute[*GetUserArticlesListResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}
