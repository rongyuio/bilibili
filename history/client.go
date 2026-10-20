// Package history 封装历史记录相关接口。
package history

import "github.com/rongyuio/bilibili"

// Client 是历史记录接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造历史记录接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
