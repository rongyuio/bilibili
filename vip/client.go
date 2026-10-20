// Package vip 封装大会员相关接口。
package vip

import "github.com/rongyuio/bilibili"

// Client 是大会员接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造大会员接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
