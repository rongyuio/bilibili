package creativecenter

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 创作中心电磁力接口。文档见 creativecenter/railgun.md。

// GetElectricLevelResult 是获取电磁力等级（web 端）的 data 对象。
type GetElectricLevelResult struct {
	Mid        int   `json:"mid"`         // 当前用户mid
	Level      int   `json:"level"`       // 电磁力等级
	Score      int   `json:"score"`       // 电磁力分数
	Credit     int   `json:"credit"`      // 信用分
	State      int   `json:"state"`       // 作用尚不明确，一直是2
	UpdateDate int64 `json:"update_date"` // 作用尚不明确，实测为0
}

// GetElectricLevel 获取电磁力等级（web端）。
//
// 仅支持 Cookie（SESSDATA）认证，无请求参数。电磁力每周日下午刷新。
func (c *Client) GetElectricLevel(ctx context.Context) (*GetElectricLevelResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/studio/up-rating/v3/rating/info"
	)
	return bilibili.Execute[*GetElectricLevelResult](ctx, c.Client, method, url, nil)
}
