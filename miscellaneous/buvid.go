package miscellaneous

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// buvid 获取接口。

// GetBuvidResult 是仅获取 buvid3 接口的响应。
type GetBuvidResult struct {
	Buvid string `json:"buvid"` // buvid3。需手动存放至 cookie 中
}

// GetBuvid 仅获取 buvid3。
//
// 返回的 buvid 需要调用方自行保存到 Cookie 中。
func (c *Client) GetBuvid(ctx context.Context) (*GetBuvidResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-frontend/getbuvid"
	)
	return bilibili.Execute[*GetBuvidResult](ctx, c.Client, method, url, nil)
}

// GetBuvid34Result 是 buvid3 / buvid4 接口的响应。
type GetBuvid34Result struct {
	B3 string `json:"b_3"` // buvid3
	B4 string `json:"b_4"` // buvid4
}

// GetBuvid34 获取 buvid3 与 buvid4。
func (c *Client) GetBuvid34(ctx context.Context) (*GetBuvid34Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/frontend/finger/spi"
	)
	return bilibili.Execute[*GetBuvid34Result](ctx, c.Client, method, url, nil)
}
