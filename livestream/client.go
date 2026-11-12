// Package livestream 封装直播弹幕 WebSocket 信息流。
package livestream

import "github.com/rongyuio/bilibili"

// Client 是直播弹幕信息流的调用入口，复用根包的会话与 WBI 签名。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造直播弹幕信息流客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
