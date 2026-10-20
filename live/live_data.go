package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播数据相关接口。响应模型见 live_data_model.go。

// GetLiveSessionDataParam 指定要查询的直播场次。
type GetLiveSessionDataParam struct {
	LiveKey string `json:"live_key"` // 标记直播场次的 key；不提供将获得无效数据
}

// GetLiveSessionData 获取某场直播的直播数据，只能查询自己的场次。
//
// 建议在直播结束当下立即调用，否则直播时长可能不准确。
func (c *Client) GetLiveSessionData(ctx context.Context, param GetLiveSessionDataParam) (*GetLiveSessionDataResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/live/StopLiveData"
	)
	return bilibili.Execute[*GetLiveSessionDataResult](ctx, c.Client, method, url, param, liveMiscReferer())
}

// GetLivePerformance 获取当前账号的直播表现（界面以雷达图展示）。
func (c *Client) GetLivePerformance(ctx context.Context) (*GetLivePerformanceResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/date/Overview"
	)
	return bilibili.Execute[*GetLivePerformanceResult](ctx, c.Client, method, url, nil, liveMiscReferer())
}
