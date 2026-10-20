package blackroom

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 风纪委员仲裁操作相关接口。各接口的返回体独立定义，互不复用。

// GetJuryCaseNextResult 是拉取新案件接口的 data 对象。
type GetJuryCaseNextResult struct {
	CaseID string `json:"case_id"` // 仲裁案件 id
}

// GetJuryCaseNext 拉取新案件
//
// 需要登录（Cookie）。业务错误码：25006 风纪委员资格已过期，25008 没有案件，25014 已审满。
func (c *Client) GetJuryCaseNext(ctx context.Context) (*GetJuryCaseNextResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/case/next"
	)
	return bilibili.Execute[*GetJuryCaseNextResult](ctx, c.Client, method, url, nil)
}
