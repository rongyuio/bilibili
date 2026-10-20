package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 关注动态流（dynamic/all）相关接口。响应模型见 dynamic_all_model.go。

// checkDynamicUpdateWebLocation 是检测新动态接口默认的页面位置标识。
const checkDynamicUpdateWebLocation = "333.1365"

// CheckDynamicUpdateParam 指定检测新动态时使用的更新基线。
type CheckDynamicUpdateParam struct {
	Type           string `json:"type,omitempty" request:"query,omitempty"`         // 动态类型，默认为 all
	UpdateBaseline int64  `json:"update_baseline" request:"query"`                  // 更新基线，取上次 GetDynamicFeedAll 返回的 UpdateBaseline；填 0 也可正常获取
	WebLocation    string `json:"web_location,omitempty" request:"query,omitempty"` // 页面位置标识，留空由库填入默认值
}

// CheckDynamicUpdate 检测关注动态流是否有新动态。
//
// 认证方式：Cookie（SESSDATA）。
func (c *Client) CheckDynamicUpdate(ctx context.Context, param CheckDynamicUpdateParam) (*CheckDynamicUpdateResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/all/update"
	)
	return bilibili.Execute[*CheckDynamicUpdateResult](ctx, c.Client, method, url, normalizeCheckDynamicUpdateParam(param))
}

// normalizeCheckDynamicUpdateParam 补齐网页端固定带、调用方通常不关心的那几个参数。
func normalizeCheckDynamicUpdateParam(param CheckDynamicUpdateParam) CheckDynamicUpdateParam {
	if param.Type == "" {
		param.Type = "all"
	}
	if param.WebLocation == "" {
		param.WebLocation = checkDynamicUpdateWebLocation
	}
	return param
}
