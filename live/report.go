package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播心跳上报相关接口。响应模型见 report_model.go。

// ReportLiveWebHeartBeatParam 指定 Web 端直播心跳上报的参数，字段均进入 URL 查询串。
type ReportLiveWebHeartBeatParam struct {
	// heartbeat 正文，使用 base64 编码。解码后为以竖线分隔的四个数字：
	// "上次返回的 next_interval|真实直播间号|1|0"，例如 "60|26863308|1|0"。
	Hb string `json:"hb,omitempty" request:"query,omitempty"`
	Pf string `json:"pf,omitempty" request:"query,omitempty"` // 平台名称，可为 web
}

// ReportLiveWebHeartBeat 上报 Web 端直播心跳，返回下一次心跳的间隔。
// 调用方需按 NextInterval 等待后再次上报。
func (c *Client) ReportLiveWebHeartBeat(ctx context.Context, param ReportLiveWebHeartBeatParam) (*ReportLiveWebHeartBeatResult, error) {
	return bilibili.Execute[*ReportLiveWebHeartBeatResult](ctx, c.Client, resty.MethodGet,
		"https://live-trace.bilibili.com/xlive/rdata-interface/v1/heartbeat/webHeartBeat", param)
}
