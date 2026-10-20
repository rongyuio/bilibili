// Package clientinfo 封装客户端信息相关接口。
package clientinfo

import "github.com/rongyuio/bilibili"

// Client 是客户端信息接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造客户端信息接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
