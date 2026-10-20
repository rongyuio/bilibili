package message

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetIMWebUnreadParam 是获取未读消息数（IM web 接口）的请求参数。
type GetIMWebUnreadParam struct {
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号，默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识，可为 web 等
}

// GetIMWebUnread 获取未读消息数（IM web 接口）
//
// 认证方式：Cookie（SESSDATA）。本接口不返回未读私信数。
func (c *Client) GetIMWebUnread(ctx context.Context, param GetIMWebUnreadParam) (*GetIMWebUnreadResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/x/im/web/msgfeed/unread"
	)
	return bilibili.Execute[*GetIMWebUnreadResult](ctx, c.Client, method, url, param)
}
