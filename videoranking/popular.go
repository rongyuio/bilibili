package videoranking

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 热门视频与每周必看相关接口。响应模型见 model_popular.go。

// GetPopularVideoListParam 是获取当前热门视频列表的请求参数。
type GetPopularVideoListParam struct {
	Pn int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
	Ps int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 20
}

// GetPopularVideoList 获取当前热门视频列表
//
// 携带 Cookie 时返回数据的排行是个性化的。
func (c *Client) GetPopularVideoList(ctx context.Context, param GetPopularVideoListParam) (*GetPopularVideoListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/popular"
	)
	return bilibili.Execute[*GetPopularVideoListResult](ctx, c.Client, method, url, param)
}

// GetWeeklyPopularList 获取每周必看全部列表
func (c *Client) GetWeeklyPopularList(ctx context.Context) (*GetWeeklyPopularListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/popular/series/list"
	)
	return bilibili.Execute[*GetWeeklyPopularListResult](ctx, c.Client, method, url, nil)
}

// GetWeeklyPopularDetailParam 是获取每周必看选期详细信息的请求参数。
type GetWeeklyPopularDetailParam struct {
	Number int `json:"number"` // 期数
}

// GetWeeklyPopularDetail 获取每周必看选期详细信息
func (c *Client) GetWeeklyPopularDetail(ctx context.Context, param GetWeeklyPopularDetailParam) (*GetWeeklyPopularDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/popular/series/one"
	)
	return bilibili.Execute[*GetWeeklyPopularDetailResult](ctx, c.Client, method, url, param)
}
