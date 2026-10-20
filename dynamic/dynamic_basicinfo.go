package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态草稿相关接口。响应模型见 dynamic_basicinfo_model.go。

// GetDynamicDraftList 获取当前账号的动态草稿列表。
//
// 认证方式：Cookie（SESSDATA）。接口无请求参数；每条草稿的正文以 JSON 字符串形式
// 放在 GetDynamicDraftListItem.Request 中，需要调用方自行解析。
func (c *Client) GetDynamicDraftList(ctx context.Context) (*GetDynamicDraftListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_draft/v1/dynamic_draft/get_drafts"
	)
	return bilibili.Execute[*GetDynamicDraftListResult](ctx, c.Client, method, url, nil)
}
