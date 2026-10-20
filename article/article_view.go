package article

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 专栏正文相关接口。响应模型见 model_article_content.go。

// GetArticleContentParam 是获取专栏正文内容的参数。
type GetArticleContentParam struct {
	ID         int    `json:"id" request:"query"`                              // 专栏文章 ID
	GaiaSource string `json:"gaia_source,omitempty" request:"query,omitempty"` // 来源，固定为 main_web
}

// GetArticleContent 获取专栏正文内容
func (c *Client) GetArticleContent(ctx context.Context, param GetArticleContentParam) (*GetArticleContentResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/view"
	)
	return bilibili.Execute[*GetArticleContentResult](ctx, c.Client, method, url, param)
}
