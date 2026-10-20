package cheese

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 课程基本信息相关接口。响应模型见 model_info.go。

// GetCheeseInfoParam 获取课程基本信息的参数。
type GetCheeseInfoParam struct {
	SeasonID int `json:"season_id,omitempty" request:"query,omitempty"` // 课程ssid。与ep_id任选其一，与番剧ssid不互通
	EpID     int `json:"ep_id,omitempty" request:"query,omitempty"`     // 课程epid。与season_id任选其一，与番剧epid不互通
}

// GetCheeseInfo 获取课程基本信息
func (c *Client) GetCheeseInfo(ctx context.Context, param GetCheeseInfoParam) (*GetCheeseInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pugv/view/web/season"
	)
	return bilibili.Execute[*GetCheeseInfoResult](ctx, c.Client, method, url, param)
}

// GetCheeseEpisodeListParam 获取课程分集列表的参数。
type GetCheeseEpisodeListParam struct {
	SeasonID int `json:"season_id" request:"query"`              // 课程ssid。与番剧ssid不互通
	Ps       int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为50
	Pn       int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
}

// GetCheeseEpisodeList 获取课程分集列表
func (c *Client) GetCheeseEpisodeList(ctx context.Context, param GetCheeseEpisodeListParam) (*GetCheeseEpisodeListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pugv/view/web/ep/list"
	)
	return bilibili.Execute[*GetCheeseEpisodeListResult](ctx, c.Client, method, url, param)
}
