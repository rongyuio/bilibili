package creativecenter

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 创作中心图文（动态 / 专栏）操作接口。文档见 creativecenter/opus.md。

// RemoveDynamicOpusParam 是删除动态（图文 opus）的请求参数。
type RemoveDynamicOpusParam struct {
	DynIDStr string `json:"dyn_id_str" request:"json"`                    // 动态的 opus id
	DynType  int    `json:"dyn_type,omitempty" request:"json,omitempty"`  // 动态类型。作用尚不明确
	RidStr   string `json:"rid_str,omitempty" request:"json,omitempty"`   // 文档未说明内容
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。如 web
}

// RemoveDynamicOpus 删除动态（图文 opus）。CSRF 走 query，动态 id 走 JSON 请求体。
//
// 与 dynamic 子包的 RemoveDynamic（旧版 rm_dynamic，参数为 dynamic_id）不是同一接口：
// 本方法对应创作中心的 x/dynamic/feed/operate/remove，使用 opus id 字符串。
func (c *Client) RemoveDynamicOpus(ctx context.Context, param RemoveDynamicOpusParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/dynamic/feed/operate/remove"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// DeleteArticleParam 是删除专栏的请求参数。
type DeleteArticleParam struct {
	Aid int `json:"aid"` // 文章cvid
}

// DeleteArticle 删除专栏。请求体为 x-www-form-urlencoded，aid 与 CSRF 均放在表单中。
func (c *Client) DeleteArticle(ctx context.Context, param DeleteArticleParam) error {
	const (
		method = resty.MethodPost
		url    = "https://member.bilibili.com/x/web/article/delete"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.MoveFormParams("aid"), bilibili.FillFormCsrf(c.Client))
	return err
}
