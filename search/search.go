package search

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 搜索相关接口。响应模型见 model_search.go。

// IntegratedSearchParam 是综合搜索的请求参数。
type IntegratedSearchParam struct {
	Keyword string `json:"keyword" request:"query"` // 需要搜索的关键词
	// Page     int    `json:"page" request:"query"`      // 从1开始
	// PageSize int    `json:"page_size" request:"query"` // 默认42
}

// IntegratedSearch 综合搜索（web端）。接口需要 WBI 签名。
func (c *Client) IntegratedSearch(ctx context.Context, param IntegratedSearchParam) (*IntegratedSearchResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/search/all/v2"
	)
	return bilibili.Execute[*IntegratedSearchResult](ctx, c.Client, method, url, param, c.FillWbi())
}
