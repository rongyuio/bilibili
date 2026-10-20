// Package lottery 封装活动抽奖与动态抽奖相关接口。
package lottery

import "github.com/rongyuio/bilibili"

// Client 是抽奖接口的调用入口，复用根包的请求链路与会话。
type Client struct {
	*bilibili.Client
}

// New 用根包客户端构造抽奖接口客户端。
func New(c *bilibili.Client) *Client {
	return &Client{Client: c}
}
