package videoranking

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 入站必刷视频相关接口。响应模型见 model_precious.go。

// GetPreciousVideoList 获取入站必刷视频
func (c *Client) GetPreciousVideoList(ctx context.Context) (*GetPreciousVideoListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/popular/precious"
	)
	return bilibili.Execute[*GetPreciousVideoListResult](ctx, c.Client, method, url, nil)
}
