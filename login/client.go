// Package login 封装登录与会话相关接口。
package login

import "github.com/rongyuio/bilibili"

// Client 是登录与会话接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造登录与会话接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
