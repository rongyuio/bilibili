package vip

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 大会员中心与大积分相关接口。响应模型见 vip_center_model.go。

// GetVipCenterCombineParam 是获取大会员中心信息的请求参数。
type GetVipCenterCombineParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
	Platform  string `json:"platform,omitempty" request:"query,omitempty"`   // 平台表示。web端：web。安卓APP：android
	MobiApp   string `json:"mobi_app,omitempty" request:"query,omitempty"`   // APP 名称。安卓APP：android
	Build     int    `json:"build,omitempty" request:"query,omitempty"`      // 构建 id
}

// GetVipCenterCombine 获取大会员中心信息
func (c *Client) GetVipCenterCombine(ctx context.Context, param GetVipCenterCombineParam) (*GetVipCenterCombineResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vip/web/vip_center/combine"
	)
	return bilibili.Execute[*GetVipCenterCombineResult](ctx, c.Client, method, url, param)
}

// GetVipPointHomepageCombineParam 是获取大积分中心信息的请求参数。
type GetVipPointHomepageCombineParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
}

// GetVipPointHomepageCombine 获取大积分中心信息
func (c *Client) GetVipPointHomepageCombine(ctx context.Context, param GetVipPointHomepageCombineParam) (*GetVipPointHomepageCombineResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.biliapi.com/x/vip_point/homepage/combine"
	)
	return bilibili.Execute[*GetVipPointHomepageCombineResult](ctx, c.Client, method, url, param)
}

// GetVipPointListParam 是获取大积分改变记录的请求参数。
type GetVipPointListParam struct {
	AccessKey  string `json:"access_key,omitempty" request:"query,omitempty"`  // APP登录Token。APP方式必要
	ChangeType int    `json:"change_type,omitempty" request:"query,omitempty"` // 改变类型。0：所有类型。1：获取记录。2：消耗记录。默认0
	Pn         int    `json:"pn,omitempty" request:"query,omitempty"`          // 分页页数。默认1
	Ps         int    `json:"ps,omitempty" request:"query,omitempty"`          // 分页大小。默认20
}

// GetVipPointList 获取大积分改变记录
func (c *Client) GetVipPointList(ctx context.Context, param GetVipPointListParam) (*GetVipPointListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vip_point/list"
	)
	return bilibili.Execute[*GetVipPointListResult](ctx, c.Client, method, url, param)
}
