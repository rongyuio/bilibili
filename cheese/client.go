// Package cheese 封装课程相关接口。
package cheese

import "github.com/rongyuio/bilibili"

// Client 是课程接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造课程接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
