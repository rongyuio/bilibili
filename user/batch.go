package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户批量查询相关接口。响应模型见 model_batch.go。

// GetNameToUIDParam 是用户名查 mid 的请求参数。
type GetNameToUIDParam struct {
	Names string `json:"names" request:"query"` // 用户名。多个用户名以英文逗号分隔
}

// GetNameToUID 用户名查 mid。
func (c *Client) GetNameToUID(ctx context.Context, param GetNameToUIDParam) (*NameToUIDResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/name-to-uid"
	)
	return bilibili.Execute[*NameToUIDResult](ctx, c.Client, method, url, param)
}
