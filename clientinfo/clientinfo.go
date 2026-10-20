package clientinfo

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetIPAddrLocationParam 指定要查询归属地的 IP。
type GetIPAddrLocationParam struct {
	IP string `json:"ip,omitempty" request:"query,omitempty"` // IP 地址。IPv4 或 IPv6 均可，留空则查询请求来源 IP
}

// GetIPAddrLocationResult 是 IP 归属地接口的 data 对象。
type GetIPAddrLocationResult struct {
	Addr      string                  `json:"addr"`      // 公网 IP 地址
	Country   string                  `json:"country"`   // 国家/地区名
	Province  string                  `json:"province"`  // 省/州。非必须存在项
	City      string                  `json:"city"`      // 城市。非必须存在项
	ISP       string                  `json:"isp"`       // 运营商名
	Latitude  bilibili.NumberOrString `json:"latitude"`  // 纬度。文档标注为 num，实际返回也可能是字符串
	Longitude bilibili.NumberOrString `json:"longitude"` // 经度。文档标注为 num，实际返回也可能是字符串
}

// GetIPAddrLocation 查询任意 IP 地址的归属地。
//
// 文档给出的两个地址等效，此处使用 ip_service 接口。
func (c *Client) GetIPAddrLocation(ctx context.Context, param GetIPAddrLocationParam) (*GetIPAddrLocationResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/ip_service/v1/ip_service/get_ip_addr"
	)
	return bilibili.Execute[*GetIPAddrLocationResult](ctx, c.Client, method, url, param)
}
