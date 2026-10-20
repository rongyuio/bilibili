package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态 @ 用户时的搜索填充列表相关接口。响应模型见 dynamic_atlist_model.go。

// SearchDynamicMentionUserParam 指定 @ 用户搜索的关键字。
type SearchDynamicMentionUserParam struct {
	Keyword string `json:"keyword,omitempty" request:"query,omitempty"` // 搜索关键字；为空时返回所有关注用户
}

// SearchDynamicMentionUser 根据关键字搜索用户（at 别人时的填充列表）。
//
// 关键字不要求顺序匹配。认证方式：Cookie（SESSDATA）。
func (c *Client) SearchDynamicMentionUser(ctx context.Context, param SearchDynamicMentionUserParam) (*SearchDynamicMentionUserResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/mention/search"
	)
	return bilibili.Execute[*SearchDynamicMentionUserResult](ctx, c.Client, method, url, param)
}
