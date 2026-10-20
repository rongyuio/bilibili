package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 稿件投诉相关接口。响应模型见 model_appeal.go。

// GetAppealTags 获取稿件投诉类型列表。
func (c *Client) GetAppealTags(ctx context.Context) ([]GetAppealTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/appeal/tags"
	)
	return bilibili.Execute[[]GetAppealTagsItem](ctx, c.Client, method, url, nil)
}
