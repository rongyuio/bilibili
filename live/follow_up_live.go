package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户关注的 UP 主直播情况相关接口（live-follow-up-live 域）。
// 响应模型见 follow_up_live_model.go。

// GetFollowingLiveStatusParam 指定关注直播情况列表的分页与筛选参数。
type GetFollowingLiveStatusParam struct {
	Page         int  `json:"page,omitempty" request:"query,omitempty"`         // 页码，默认为 1
	PageSize     int  `json:"page_size,omitempty" request:"query,omitempty"`    // 每页的数据数量，有效值 1-10，默认为 10
	IgnoreRecord int  `json:"ignoreRecord,omitempty" request:"query,omitempty"` // 含义尚不明确，一般为 1
	HitAB        bool `json:"hit_ab,omitempty" request:"query,omitempty"`       // 含义尚不明确，建议用 true，不填为 false
}

// GetFollowingLiveStatus 获取当前账号关注的所有 UP 主的直播情况，需要登录（SESSDATA）。
func (c *Client) GetFollowingLiveStatus(ctx context.Context, param GetFollowingLiveStatusParam) (*GetFollowingLiveStatusResult, error) {
	return bilibili.Execute[*GetFollowingLiveStatusResult](ctx, c.Client, resty.MethodGet,
		"https://api.live.bilibili.com/xlive/web-ucenter/user/following", param)
}

// GetFollowingLiveWebListParam 指定 PC 端正在直播列表的筛选参数。
type GetFollowingLiveWebListParam struct {
	HitAB bool `json:"hit_ab,omitempty" request:"query,omitempty"` // 会影响回复中部分字段的取值，不填为 false
}

// GetFollowingLiveWebList 获取当前账号关注的、正在直播的 UP 主列表（PC 端），需要登录（SESSDATA）。
func (c *Client) GetFollowingLiveWebList(ctx context.Context, param GetFollowingLiveWebListParam) (*GetFollowingLiveWebListResult, error) {
	return bilibili.Execute[*GetFollowingLiveWebListResult](ctx, c.Client, resty.MethodGet,
		"https://api.live.bilibili.com/xlive/web-ucenter/v1/xfetter/GetWebList", param)
}
