package search

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 分类搜索相关接口。响应模型见 model_category.go。

// CategorySearchParam 是分类搜索的请求参数。
type CategorySearchParam struct {
	SearchType string `json:"search_type" request:"query"`                     // 搜索目标类型。视频：video，番剧：media_bangumi，影视：media_ft，直播间及主播：live，直播间：live_room，主播：live_user，专栏：article，话题：topic，用户：bili_user，相簿：photo
	Keyword    string `json:"keyword" request:"query"`                         // 需要搜索的关键词
	Order      string `json:"order,omitempty" request:"query,omitempty"`       // 结果排序方式。视频、专栏及相簿默认为 totalrank；直播间默认为 online；用户默认为 0
	OrderSort  int    `json:"order_sort,omitempty" request:"query,omitempty"`  // 用户粉丝数及等级排序顺序，仅用于搜索用户。由高到低：0，由低到高：1
	UserType   int    `json:"user_type,omitempty" request:"query,omitempty"`   // 用户分类筛选，仅用于搜索用户。全部用户：0，up主：1，普通用户：2，认证用户：3
	Duration   int    `json:"duration,omitempty" request:"query,omitempty"`    // 视频时长筛选，仅用于搜索视频。全部时长：0，10分钟以下：1，10-30分钟：2，30-60分钟：3，60分钟以上：4
	Tids       int    `json:"tids,omitempty" request:"query,omitempty"`        // 视频分区筛选，仅用于搜索视频。全部分区：0，筛选分区：目标分区tid
	CategoryID int    `json:"category_id,omitempty" request:"query,omitempty"` // 专栏及相簿分区筛选。全部分区：0
	Page       int    `json:"page,omitempty" request:"query,omitempty"`        // 页码，默认为1
}

// CategorySearch 分类搜索（web端）。接口需要 WBI 签名。
func (c *Client) CategorySearch(ctx context.Context, param CategorySearchParam) (*CategorySearchResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/search/type"
	)
	return bilibili.Execute[*CategorySearchResult](ctx, c.Client, method, url, param, c.FillWbi())
}
