// Package danmaku 封装弹幕相关接口。
package danmaku

import "github.com/rongyuio/bilibili"

// Client 是弹幕接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造弹幕接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
