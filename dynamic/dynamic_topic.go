package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态发布相关的话题接口。响应模型见 dynamic_topic_model.go。

// SearchPublishTopicParam 指定发布话题搜索的关键字与分页条件。
type SearchPublishTopicParam struct {
	Keywords    string `json:"keywords,omitempty" request:"query,omitempty"`     // 关键词
	Content     string `json:"content,omitempty" request:"query,omitempty"`      // 空
	UploadID    string `json:"upload_id,omitempty" request:"query,omitempty"`    // 上传 id。页面刷新第一次请求时为空，后均非空，格式为 `${mid}_${ts}_${dig4}`
	PageSize    int    `json:"page_size,omitempty" request:"query,omitempty"`    // 页大小。默认为 20
	PageNum     int    `json:"page_num,omitempty" request:"query,omitempty"`     // 固定为 1，不用于翻页
	Offset      int    `json:"offset,omitempty" request:"query,omitempty"`       // 偏移。可从响应 data.page_info.offset 中获取
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面来源标识。如 333.1365
}

// SearchPublishTopic 搜索发布话题
func (c *Client) SearchPublishTopic(ctx context.Context, param SearchPublishTopicParam) (*SearchPublishTopicResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/topic/pub/search"
	)
	return bilibili.Execute[*SearchPublishTopicResult](ctx, c.Client, method, url, param)
}

// GetRecommendSearchTopicParam 指定推荐搜索话题的查询条件。
type GetRecommendSearchTopicParam struct {
	Keywords    string `json:"keywords,omitempty" request:"query,omitempty"`     // 关键词。文档标注为空
	UploadID    string `json:"upload_id,omitempty" request:"query,omitempty"`    // 上传 id。含义同 SearchPublishTopic
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面来源标识。如 333.1365，有时请求不带该参数
}

// GetRecommendSearchTopic 推荐搜索话题。接口可能不返回数据，但网页端会请求它
func (c *Client) GetRecommendSearchTopic(ctx context.Context, param GetRecommendSearchTopicParam) (*GetRecommendSearchTopicResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/topic/pub/rcmd/search"
	)
	return bilibili.Execute[*GetRecommendSearchTopicResult](ctx, c.Client, method, url, param)
}

// CheckTopicExistedParam 指定要检查的话题名。
type CheckTopicExistedParam struct {
	Topic string `json:"topic" request:"query"` // 话题名。大小写敏感
}

// CheckTopicExisted 检查话题是否存在。需要 Cookie（SESSDATA）
func (c *Client) CheckTopicExisted(ctx context.Context, param CheckTopicExistedParam) (*CheckTopicExistedResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/topic/pub/is_existed"
	)
	return bilibili.Execute[*CheckTopicExistedResult](ctx, c.Client, method, url, param)
}
