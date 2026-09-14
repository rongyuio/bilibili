package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 表情相关接口。响应模型见 emote_model.go。

type EmoteActionParam struct {
	PackageID int    `json:"package_id"` // 表情包ID
	Business  string `json:"business"`   // 表情包使用场景
	IDs       []int  `json:"ids"`        // 表情包ID集合
}

// AddEmote 添加表情包
func (c *Client) AddEmote(ctx context.Context, param EmoteActionParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/emote/package/add"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// RemoveEmote 移除表情包
func (c *Client) RemoveEmote(ctx context.Context, param EmoteActionParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/emote/package/remove"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// GetMyEmoteList 获取我的表情包列表
func (c *Client) GetMyEmoteList(ctx context.Context, param EmoteActionParam) (*EmoteList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/user/panel"
	)
	return execute[*EmoteList](ctx, c, method, url, param)
}

// GetEmotePackageDetailInfo 获取指定ID表情包的详细信息
func (c *Client) GetEmotePackageDetailInfo(ctx context.Context, param EmoteActionParam) (*EmoteList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/package"
	)
	return execute[*EmoteList](ctx, c, method, url, param)
}

// GetAllEmoteList 获取所有表情列表
func (c *Client) GetAllEmoteList(ctx context.Context, param EmoteActionParam) (*AllEmoteList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/setting/panel"
	)
	return execute[*AllEmoteList](ctx, c, method, url, param)
}
