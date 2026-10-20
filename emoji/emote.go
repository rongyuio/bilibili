package emoji

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 表情包相关接口。响应模型见 emote_model.go。

// EmoteActionParam 是表情包操作与查询共用的请求参数。
//
// 各接口只取自己需要的字段，其余字段按零值照常发送（与迁移前一致）。
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
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// RemoveEmote 移除表情包
func (c *Client) RemoveEmote(ctx context.Context, param EmoteActionParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/emote/package/remove"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// GetMyEmoteList 获取我的表情包列表
func (c *Client) GetMyEmoteList(ctx context.Context, param EmoteActionParam) (*GetMyEmoteListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/user/panel"
	)
	return bilibili.Execute[*GetMyEmoteListResult](ctx, c.Client, method, url, param)
}

// GetEmotePackageDetailInfo 获取指定ID表情包的详细信息
func (c *Client) GetEmotePackageDetailInfo(ctx context.Context, param EmoteActionParam) (*GetEmotePackageDetailInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/package"
	)
	return bilibili.Execute[*GetEmotePackageDetailInfoResult](ctx, c.Client, method, url, param)
}

// GetAllEmoteList 获取所有表情列表
func (c *Client) GetAllEmoteList(ctx context.Context, param EmoteActionParam) (*GetAllEmoteListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/setting/panel"
	)
	return bilibili.Execute[*GetAllEmoteListResult](ctx, c.Client, method, url, param)
}
