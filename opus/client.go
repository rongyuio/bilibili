// Package opus 封装图文相关接口。
package opus

import "github.com/rongyuio/bilibili"

// Client 是图文接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造图文接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
