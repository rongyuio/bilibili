package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态操作相关接口。响应模型见 dynamic_action_model.go。

// LikeDynamicParam 指定要点赞的动态与点赞状态。
type LikeDynamicParam struct {
	DynIDStr  string `json:"dyn_id_str" request:"json"`                     // 动态id
	Up        int    `json:"up" request:"json"`                             // 点赞状态。0：切换。1：点赞。2：取消点赞
	Spmid     string `json:"spmid,omitempty" request:"json,omitempty"`      // 页面标识。文档示例为 333.1369.0.0
	FromSpmid string `json:"from_spmid,omitempty" request:"json,omitempty"` // 来源页面标识。文档示例为 333.999.0.0
}

// LikeDynamic 点赞动态
func (c *Client) LikeDynamic(ctx context.Context, param LikeDynamicParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/dynamic/feed/dyn/thumb"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}
