package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 专栏相关接口。响应模型见 article_model.go。

type GetArticlesInfoParam struct {
	ID int `json:"id"` // 文集rlid
}

// GetArticlesInfo 获取文集基本信息
func (c *Client) GetArticlesInfo(ctx context.Context, param GetArticlesInfoParam) (*ArticlesInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/list/web/articles"
	)
	return execute[*ArticlesInfo](ctx, c, method, url, param)
}

// GetArticleInfoParam 是 GetArticlesInfoParam 的别名。
type GetArticleInfoParam = GetArticlesInfoParam

// GetArticleInfo 获取专栏文章基本信息
func (c *Client) GetArticleInfo(ctx context.Context, param GetArticleInfoParam) (*ArticleInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/viewinfo"
	)
	return execute[*ArticleInfo](ctx, c, method, url, param)
}

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
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

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
	return execute[*CoinArticleResult](ctx, c, method, url, param, fillCsrf(c))
}

// FavoritesArticleParam 是 GetArticlesInfoParam 的别名。
type FavoritesArticleParam = GetArticlesInfoParam

// FavoritesArticle 收藏文章
func (c *Client) FavoritesArticle(ctx context.Context, param FavoritesArticleParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/article/favorites/add"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type GetUserArticleListParam struct {
	Mid  int    `json:"mid"`                                      // 用户uid
	Pn   int    `json:"pn,omitempty" request:"query,omitempty"`   // 默认：1
	Ps   int    `json:"ps,omitempty" request:"query,omitempty"`   // 默认：30。范围：[1,30]
	Sort string `json:"sort,omitempty" request:"query,omitempty"` // publish_time：最新发布。view：最多阅读。fav：最多收藏。默认：publish_time
}

// GetUserArticleList 获取用户专栏文章列表
func (c *Client) GetUserArticleList(ctx context.Context, param GetUserArticleListParam) (*UserArticleList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/article"
	)
	return execute[*UserArticleList](ctx, c, method, url, param, fillCsrf(c))
}

type GetUserArticlesListParam struct {
	Mid      int    `json:"mid"`                                      // 用户uid
	Sort     int    `json:"sort,omitempty" request:"query,omitempty"` // 排序方式。0：最近更新。1：最多阅读
	Jsonp    string `json:"jsonp,omitempty" request:"query,omitempty"`
	Callback string `json:"callback,omitempty" request:"query,omitempty"`
}

// GetUserArticlesList 获取用户专栏文集列表
func (c *Client) GetUserArticlesList(ctx context.Context, param GetUserArticlesListParam) (*UserArticlesList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/up/lists"
	)
	return execute[*UserArticlesList](ctx, c, method, url, param, fillCsrf(c))
}
