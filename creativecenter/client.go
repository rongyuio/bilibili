// Package creativecenter 封装创作中心相关接口。
package creativecenter

import "github.com/rongyuio/bilibili"

// Client 是创作中心接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造创作中心接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
