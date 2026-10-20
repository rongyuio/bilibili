package broadcast

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetBroadcastServersParam 指定广播服务器地址的返回形式。
type GetBroadcastServersParam struct {
	Platform string `json:"platform" request:"query"` // 平台选择。为 web 时输出域名，其他时输出 ip
}

// GetBroadcastServersBackoff 是广播服务器的退避配置。
type GetBroadcastServersBackoff struct {
	MaxDelay  int     `json:"max_delay"`  // 作用尚不明确
	BaseDelay int     `json:"base_delay"` // 作用尚不明确
	Factor    float64 `json:"factor"`     // 作用尚不明确
	Jitter    float64 `json:"jitter"`     // 作用尚不明确
}

// GetBroadcastServersResult 是广播服务器地址的 data 对象。
type GetBroadcastServersResult struct {
	Domain       string                      `json:"domain"`        // 广播服务器 url
	TcpPort      int                         `json:"tcp_port"`      // tcp 端口
	WsPort       int                         `json:"ws_port"`       // websocket 端口
	WssPort      int                         `json:"wss_port"`      // websocket ssl 端口
	Heartbeat    int                         `json:"heartbeat"`     // 最大心跳包间隔时间
	Nodes        []string                    `json:"nodes"`         // 服务节点地址列表
	Backoff      *GetBroadcastServersBackoff `json:"backoff"`       // 作用尚不明确
	HeartbeatMax int                         `json:"heartbeat_max"` // 作用尚不明确
	WssPortV2    int                         `json:"wss_port_v2"`   // 作用尚不明确。实测为 0
}

// GetBroadcastServers 获取广播服务器地址。
func (c *Client) GetBroadcastServers(ctx context.Context, param GetBroadcastServersParam) (*GetBroadcastServersResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/broadcast/servers"
	)
	return bilibili.Execute[*GetBroadcastServersResult](ctx, c.Client, method, url, param)
}
