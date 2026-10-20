// Package webwidget 封装网页展示位相关接口。
package webwidget

import "github.com/rongyuio/bilibili"

// Client 是网页展示位接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造网页展示位接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
