package vip

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 大会员相关操作接口。

// ReceiveVipExperienceResult 是领取大会员每日经验的结果。
type ReceiveVipExperienceResult struct {
	Type    int  `json:"type"`     // 0。作用尚不明确
	IsGrant bool `json:"is_grant"` // 是否领取成功
}

// ReceiveVipExperience 领取大会员每日经验。CSRF 自动填入表单。
func (c *Client) ReceiveVipExperience(ctx context.Context) (*ReceiveVipExperienceResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/vip/experience/add"
	)
	return bilibili.Execute[*ReceiveVipExperienceResult](ctx, c.Client, method, url, nil, bilibili.FillFormCsrf(c.Client))
}
