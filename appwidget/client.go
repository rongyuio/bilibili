// Package appwidget 封装APP 组件相关接口。
package appwidget

import "github.com/rongyuio/bilibili"

// Client 是APP 组件接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造APP 组件接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
