package login

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 登录基本信息相关接口。响应模型见 login_info_model.go。
// 各接口的返回体独立定义，互不复用。

// GetNavInfo 获取导航栏用户信息（原始数据）。
//
// 仅可使用 Cookie（SESSDATA）认证。未登录时返回 -101 业务错误，
// 但响应中的 wbi_img 即使未登录也会返回，库的 WBI 签名内部同样依赖该接口。
func (c *Client) GetNavInfo(ctx context.Context) (*GetNavInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/nav"
	)
	return bilibili.Execute[*GetNavInfoResult](ctx, c.Client, method, url, nil)
}

// GetHomeUserInfo 获取登录用户信息（仅部分字段）。
//
// Deprecated: 该接口已被官方弃用，字段也不完整，请改用 GetNavInfo。
func (c *Client) GetHomeUserInfo(ctx context.Context) (*GetHomeUserInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://account.bilibili.com/home/userInfo"
	)
	return bilibili.Execute[*GetHomeUserInfoResult](ctx, c.Client, method, url, nil)
}

// NavStatParam 是登录用户状态数的请求参数。
type NavStatParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必填，Cookie方式留空
}

// GetNavStat 获取当前登录用户的关注数、粉丝数与动态数（双端）。
//
// 可使用 Cookie（SESSDATA）或 APP 方式认证。
func (c *Client) GetNavStat(ctx context.Context, param NavStatParam) (*GetNavStatResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/nav/stat"
	)
	return bilibili.Execute[*GetNavStatResult](ctx, c.Client, method, url, param)
}

// GetCoin 获取当前登录用户的硬币数。
//
// 仅可使用 Cookie（SESSDATA）认证，且要求 Cookie 中 DedeUserID 存在且不为 0。
func (c *Client) GetCoin(ctx context.Context) (*GetCoinResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://account.bilibili.com/site/getCoin"
	)
	return bilibili.Execute[*GetCoinResult](ctx, c.Client, method, url, nil)
}
