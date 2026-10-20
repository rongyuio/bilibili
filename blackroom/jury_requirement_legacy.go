package blackroom

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetJuryRequirementLegacyParam 是检查是否具有申请风纪委员会资格（旧版 API）的请求参数。
type GetJuryRequirementLegacyParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必要
}

// GetJuryRequirementLegacy 检查是否具有申请风纪委员会的资格（旧版 API）
//
// 旧版地址，比新版多返回 level 字段。需要登录（Cookie 或 APP）。
func (c *Client) GetJuryRequirementLegacy(ctx context.Context, param GetJuryRequirementLegacyParam) (*GetJuryRequirementLegacyResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/jury/requirement"
	)
	return bilibili.Execute[*GetJuryRequirementLegacyResult](ctx, c.Client, method, url, param)
}
