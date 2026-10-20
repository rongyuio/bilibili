// Package blackroom 封装小黑屋与风纪相关接口。
package blackroom

import "github.com/rongyuio/bilibili"

// Client 是小黑屋与风纪接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造小黑屋与风纪接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
