package customerservice

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// HeartBeatParam 是客服消息心跳的请求参数。
type HeartBeatParam struct {
	ShopID       int    `json:"shop_id,omitempty" request:"query,omitempty"`        // 店铺 id，默认为 0
	ShopFatherID int    `json:"shop_father_id,omitempty" request:"query,omitempty"` // 父店铺 id，默认为 0
	Build        int    `json:"build,omitempty" request:"query,omitempty"`          // 客户端版本，web 端为 0
	MobiApp      string `json:"mobi_app,omitempty" request:"query,omitempty"`       // 客户端类型，如 web
}

// CustomerServiceHeartBeat 客服消息心跳。响应没有 data 字段。
func (c *Client) CustomerServiceHeartBeat(ctx context.Context, param HeartBeatParam) error {
	const (
		method = resty.MethodGet
		url    = "https://customerservice.bilibili.com/x/custom/session_svr/v1/heart_beat"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param)
	return err
}
