package login

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 会员中心相关接口：每日奖励、大会员状态、账号安全、实名认证与硬币/经验/节操记录。
// 响应模型见 member_center_model.go。各接口的返回体独立定义，互不复用。

// memberCenterWebLocation 是会员中心日志接口默认的页面位置标识。
const memberCenterWebLocation = "333.33"

// webLocationOrDefault 在调用方未指定页面标识时返回接口默认值。
func webLocationOrDefault(webLocation, fallback string) string {
	if webLocation == "" {
		return fallback
	}
	return webLocation
}

// GetDailyRewardStatusParam 指定每日奖励状态的查询参数。
type GetDailyRewardStatusParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
}

// GetDailyRewardStatus 查询每日奖励状态
//
// 认证方式：Cookie（SESSDATA）或 APP
func (c *Client) GetDailyRewardStatus(ctx context.Context, param GetDailyRewardStatusParam) (*GetDailyRewardStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/web/exp/reward"
	)
	return bilibili.Execute[*GetDailyRewardStatusResult](ctx, c.Client, method, url, param)
}

// GetDailyCoinExp 查询每日投币获得经验数
//
// 该接口实时更新。认证方式：仅可 Cookie（SESSDATA）
func (c *Client) GetDailyCoinExp(ctx context.Context) (int, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/coin/today/exp"
	)
	return bilibili.Execute[int](ctx, c.Client, method, url, nil)
}

// GetVipStatus 查询大会员状态
//
// 认证方式：仅可 Cookie（SESSDATA）
func (c *Client) GetVipStatus(ctx context.Context) (*GetVipStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vip/web/user/info"
	)
	return bilibili.Execute[*GetVipStatusResult](ctx, c.Client, method, url, nil)
}

// GetAccountSecurityInfoParam 指定账号安全情况的查询参数。
type GetAccountSecurityInfoParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
}

// GetAccountSecurityInfo 查询账号安全情况
//
// 认证方式：Cookie（SESSDATA）或 APP。Cookie 方式下 Cookie 中需存在且不为 0 的 DedeUserID
func (c *Client) GetAccountSecurityInfo(ctx context.Context, param GetAccountSecurityInfoParam) (*GetAccountSecurityInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/web/site/user/info"
	)
	return bilibili.Execute[*GetAccountSecurityInfoResult](ctx, c.Client, method, url, param)
}

// GetRealnameStatusParam 指定实名认证状态的查询参数。
type GetRealnameStatusParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
}

// GetRealnameStatus 查询账号实名认证状态
//
// 认证方式：Cookie（SESSDATA）或 APP
func (c *Client) GetRealnameStatus(ctx context.Context, param GetRealnameStatusParam) (*GetRealnameStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/realname/status"
	)
	return bilibili.Execute[*GetRealnameStatusResult](ctx, c.Client, method, url, param)
}

// GetRealnameApplyStatusParam 指定实名认证详细信息的查询参数。
type GetRealnameApplyStatusParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
}

// GetRealnameApplyStatus 查询实名认证详细信息
//
// 认证方式：Cookie（SESSDATA）或 APP
func (c *Client) GetRealnameApplyStatus(ctx context.Context, param GetRealnameApplyStatusParam) (*GetRealnameApplyStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/realname/apply/status"
	)
	return bilibili.Execute[*GetRealnameApplyStatusResult](ctx, c.Client, method, url, param)
}

// GetCoinLogParam 指定硬币变化情况的查询参数。
type GetCoinLogParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
}

// GetCoinLog 查询硬币变化情况
//
// 仅能查询最近一周的记录。认证方式：Cookie（SESSDATA）或 APP
func (c *Client) GetCoinLog(ctx context.Context, param GetCoinLogParam) (*GetCoinLogResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/web/coin/log"
	)
	return bilibili.Execute[*GetCoinLogResult](ctx, c.Client, method, url, param)
}

// GetExpLogParam 指定经验记录的查询参数。
type GetExpLogParam struct {
	Jsonp       string `json:"jsonp,omitempty" request:"query,omitempty"`        // 回调函数名。默认 jsonp
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 网页位置。默认 333.33，留空由库填入默认值
}

// GetExpLog 查询最近一周的经验记录
//
// 认证方式：Cookie（SESSDATA）
func (c *Client) GetExpLog(ctx context.Context, param GetExpLogParam) (*GetExpLogResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/web/exp/log"
	)
	param.WebLocation = webLocationOrDefault(param.WebLocation, memberCenterWebLocation)
	return bilibili.Execute[*GetExpLogResult](ctx, c.Client, method, url, param)
}

// GetMoralLogParam 指定节操记录的查询参数。
type GetMoralLogParam struct {
	Jsonp       string `json:"jsonp,omitempty" request:"query,omitempty"`        // 回调函数名。默认 jsonp
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 网页位置。默认 333.33，留空由库填入默认值
}

// GetMoralLog 查询最近一周的节操记录
//
// 认证方式：Cookie（SESSDATA）
func (c *Client) GetMoralLog(ctx context.Context, param GetMoralLogParam) (*GetMoralLogResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/web/moral/log"
	)
	param.WebLocation = webLocationOrDefault(param.WebLocation, memberCenterWebLocation)
	return bilibili.Execute[*GetMoralLogResult](ctx, c.Client, method, url, param)
}
