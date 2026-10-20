package article

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 专栏显示卡片相关接口。响应模型见 model_article_cards.go。

// articleCardWebLocation 是专栏卡片接口默认的页面位置标识。
const articleCardWebLocation = "333.1305"

// GetArticleCardsParam 是获取专栏显示卡片信息的参数。
type GetArticleCardsParam struct {
	IDs         string `json:"ids,omitempty" request:"query,omitempty"`          // 被查询的 id 列表，以 , 分隔。可填视频完整 AV/BV 号、专栏 CV 号、直播间长/短号；可能会合并同类项
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面位置标识。留空时由库填入 333.1305
}

// GetArticleCards 获取专栏显示卡片信息
//
// 接口按内容类型返回视频、专栏或直播间对象，结果见 GetArticleCardsResult。
// 请求头 Referer 需在 .bilibili.com 下，使用 New() 构造的客户端已默认带上。
func (c *Client) GetArticleCards(ctx context.Context, param GetArticleCardsParam) (*GetArticleCardsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/article/cards"
	)
	if param.WebLocation == "" {
		param.WebLocation = articleCardWebLocation
	}
	return bilibili.Execute[*GetArticleCardsResult](ctx, c.Client, method, url, param)
}
