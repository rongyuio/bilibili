package bilibili

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// 话题相关接口。响应模型见 topic_model.go。

// topicFeedWebLocation 是话题动态接口默认的页面来源标识。
const topicFeedWebLocation = "0.0"

// 话题动态列表的排序方式。
const (
	TopicFeedSortByLatest = 3 // 最新
	TopicFeedSortByHot    = 2 // 最热
)

// GetTopicFeedParam 指定话题动态列表的查询条件；可选零值不发送，WebLocation 留空时由库填默认值。
type GetTopicFeedParam struct {
	TopicID     string `json:"topic_id" request:"query"`               // 话题 ID
	SortBy      int    `json:"sort_by" request:"query,omitempty"`      // 排序方式，见 TopicFeedSortBy* 常量
	PageSize    int    `json:"page_size" request:"query,omitempty"`    // 每页数量
	Offset      string `json:"offset" request:"query,omitempty"`       // 下一页使用上次返回的 offset
	Features    string `json:"features" request:"query,omitempty"`     // 接口扩展能力，按需显式传入
	WebLocation string `json:"web_location" request:"query,omitempty"` // 页面来源标识；留空时由库填入 topicFeedWebLocation
}

// GetTopicFeed 获取一页话题动态；分页由调用方根据 HasMore 和 Offset 控制。
// 返回响应的 data，复用统一的 Cookie、context 和错误处理，不额外启用 WBI。
func (c *Client) GetTopicFeed(ctx context.Context, param GetTopicFeedParam) (*GetTopicFeedResult, error) {
	param.WebLocation = webLocationOrDefault(param.WebLocation, topicFeedWebLocation)
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/topic"
	)
	return execute[*GetTopicFeedResult](ctx, c, method, url, param)
}
