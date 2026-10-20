package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 批量获取用户详细信息相关接口。响应模型见 model_info.go。

// GetUserCardsParam 是批量获取用户详细信息（PC 端用户卡片）的请求参数。
type GetUserCardsParam struct {
	UIDs []int `json:"uids" request:"query"` // 目标用户的 mid 列表。最多 50 个
}

// GetUserCards 批量获取用户详细信息（PC 端用户卡片）。
func (c *Client) GetUserCards(ctx context.Context, param GetUserCardsParam) (map[string]*GetUserCardsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/pc-electron/v1/user/cards"
	)
	return bilibili.Execute[map[string]*GetUserCardsItem](ctx, c.Client, method, url, param)
}

// GetUserInfosParam 是批量获取用户详细信息（IM 用户信息）的请求参数。
type GetUserInfosParam struct {
	UIDs []int `json:"uids" request:"query"` // 目标用户的 mid 列表
}

// GetUserInfos 批量获取用户详细信息（IM 用户信息）。
func (c *Client) GetUserInfos(ctx context.Context, param GetUserInfosParam) ([]*GetUserInfosItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/x/im/user_infos"
	)
	return bilibili.Execute[[]*GetUserInfosItem](ctx, c.Client, method, url, param)
}
