package clientinfo

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// IP 归属地相关接口。响应模型见 ip_model.go。

// GetIPInfoNewParam 指定要查询归属地的 IP。
type GetIPInfoNewParam struct {
	IP string `json:"ip,omitempty" request:"query,omitempty"` // IP 地址。IPv4 或 IPv6 均可，留空则查询请求来源 IP
}

// GetIPInfoNew 查询任意 IP 地址的归属地（直播端 client 接口）
//
// 与 ip_service 接口等效，返回的 data 不含 zone_id 与 country_code。
func (c *Client) GetIPInfoNew(ctx context.Context, param GetIPInfoNewParam) (*GetIPInfoNewResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/client/v1/Ip/getInfoNew"
	)
	return bilibili.Execute[*GetIPInfoNewResult](ctx, c.Client, method, url, param)
}

// GetWebRoomIPInfo 根据请求 IP 确定属地（直播端 getIpInfo）
func (c *Client) GetWebRoomIPInfo(ctx context.Context) (*GetWebRoomIPInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/index/getIpInfo"
	)
	return bilibili.Execute[*GetWebRoomIPInfoResult](ctx, c.Client, method, url, nil)
}
