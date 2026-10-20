// Package newbieexam 封装新用户答题相关接口。
package newbieexam

import "github.com/rongyuio/bilibili"

// Client 是新用户答题接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造新用户答题接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
