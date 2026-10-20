package emoji

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 表情列表相关接口。响应模型见 emoji_list_model.go。

// ListParam 指定要查询的表情面板使用场景。
type ListParam struct {
	Business string `json:"business"` // 使用场景。reply：评论区。dynamic：动态
}

// GetMyEmojiList 获取我的表情列表
//
// 使用登录（Cookie）进行会员专属及已购买表情包的分发，否则全为免费表情包。
func (c *Client) GetMyEmojiList(ctx context.Context, param ListParam) (*GetMyEmojiListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/emote/user/panel/web"
	)
	return bilibili.Execute[*GetMyEmojiListResult](ctx, c.Client, method, url, param)
}
