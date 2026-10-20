package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播推荐相关接口。响应模型见 recommend_model.go。

// RecommendParam 是主页直播推荐接口的查询参数。
type RecommendParam struct {
	Platform    string `json:"platform"`                                         // 平台类型，默认为 web，实测可为任意非空串
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面标识，一般为 333.1007
}

// GetLiveRecommendList 获取主页直播推荐列表
func (c *Client) GetLiveRecommendList(ctx context.Context, param RecommendParam) (*GetLiveRecommendListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-interface/v1/webMain/getMoreRecList"
	)
	return bilibili.Execute[*GetLiveRecommendListResult](ctx, c.Client, method, url, param)
}
