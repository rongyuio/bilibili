package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 导航栏动态接口。响应模型见 dynamic_nav_model.go。

// GetDynamicNavListParam 指定导航栏动态的查询条件。
type GetDynamicNavListParam struct {
	UpdateBaseline int64 `json:"update_baseline,omitempty" request:"query,omitempty"` // 更新基线。获取新动态时使用
	Offset         int64 `json:"offset,omitempty" request:"query,omitempty"`          // 分页偏移量。翻页时使用
}

// GetDynamicNavList 获取导航栏动态列表。
//
// 需要登录（Cookie 中的 SESSDATA）。
func (c *Client) GetDynamicNavList(ctx context.Context, param GetDynamicNavListParam) (*GetDynamicNavListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/nav"
	)
	return bilibili.Execute[*GetDynamicNavListResult](ctx, c.Client, method, url, param)
}
