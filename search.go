package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 搜索相关接口。响应模型见 search_model.go。

type SearchParam struct {
	Keyword string `json:"keyword" request:"query"` // 需要搜索的关键词
	// Page     int    `json:"page" request:"query"`      // 从1开始
	// PageSize int    `json:"page_size" request:"query"` // 默认42
}

// IntegratedSearch 综合搜索（web端）
func (c *Client) IntegratedSearch(ctx context.Context, param SearchParam) (*SearchRespData, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/search/all/v2"
	)
	return execute[*SearchRespData](ctx, c, method, url, param, c.fillWbi())
}
