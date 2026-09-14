package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 搜索相关接口。响应模型见 search_model.go。
// https://socialsisteryi.github.io/bilibili-API-collect/docs/search/*

type SearchParam struct {
	Keyword string `json:"keyword" request:"query"` //	需要搜索的关键词
	// Page     int    `json:"page" request:"query"`      // 从1开始
	// PageSize int    `json:"page_size" request:"query"` // 默认42
}

// IntegratedSearch 综合搜索（web端）
// https://socialsisteryi.github.io/bilibili-API-collect/docs/search/search_request.html#%E7%BB%BC%E5%90%88%E6%90%9C%E7%B4%A2-web%E7%AB%AF
func (c *Client) IntegratedSearch(ctx context.Context, param SearchParam) (*SearchRespData, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/search/all/v2"
	)
	return execute[*SearchRespData](ctx, c, method, url, param, c.fillWbi())
}
