package bilibili

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// 话题相关接口。响应模型见 topic_model.go。

// GetTopicFeedParam 指定话题动态列表的查询条件；可选零值不发送，不自动填入默认值。
type GetTopicFeedParam struct {
	TopicID     string `json:"topic_id" request:"query"`               // 话题 ID
	SortBy      int    `json:"sort_by" request:"query,omitempty"`      // 排序；现有工具使用 3（最新），2 为最热
	PageSize    int    `json:"page_size" request:"query,omitempty"`    // 每页数量
	Offset      string `json:"offset" request:"query,omitempty"`       // 下一页使用上次返回的 offset
	Features    string `json:"features" request:"query,omitempty"`     // 接口扩展能力，按需显式传入
	WebLocation string `json:"web_location" request:"query,omitempty"` // 页面来源标识
}

// GetTopicFeed 获取一页话题动态；分页由调用方根据 HasMore 和 Offset 控制。
// 返回响应的 data，复用统一的 Cookie、context 和错误处理，不额外启用 WBI。
func (c *Client) GetTopicFeed(ctx context.Context, param GetTopicFeedParam) (*GetTopicFeedResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/topic"
	)
	return execute[*GetTopicFeedResult](ctx, c, method, url, param)
}
