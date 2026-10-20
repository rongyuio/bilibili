// Package bili 提供聚合入口：一个对象持有根客户端与全部业务域。
//
// 所有域共享同一个根会话。只需要单个业务域时，直接用根包与对应子包即可。
package bili

import (
	"context"

	"github.com/rongyuio/bilibili"
)

// New 创建根客户端并构造聚合入口。
func New() *Client {
	return newClient(bilibili.New())
}

// NewWith 用已有的根客户端构造聚合入口，不新建会话。root 为 nil 时等同 New。
func NewWith(root *bilibili.Client) *Client {
	if root == nil {
		return New()
	}
	return newClient(root)
}

// NewAnonymous 获取游客 Cookie 并构造聚合入口。
func NewAnonymous(ctx context.Context) (*Client, error) {
	root, err := bilibili.NewAnonymousClient(ctx)
	if err != nil {
		return nil, err
	}
	return newClient(root), nil
}
