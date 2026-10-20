// Package emoji 封装表情相关接口。
package emoji

import "github.com/rongyuio/bilibili"

// Client 是表情接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造表情接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
