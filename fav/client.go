// Package fav 封装收藏夹相关接口。
package fav

import "github.com/rongyuio/bilibili"

// Client 是收藏夹接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造收藏夹接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
