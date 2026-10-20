package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态详情相关接口（web-dynamic v1 detail）。响应模型见 dynamic_detail_model.go。

// dynamicDetailFeatures 是动态详情接口网页端默认请求的功能开关串。
//
// 与动态流的开关串不是同一份：少几个开关只会让卡片少几项可选内容，不影响返回结构，
// 因此留成常量而不是暴露给调用方。
const dynamicDetailFeatures = "itemOpusStyle,opusBigCover,onlyfansVote,endFooterHidden,decorationCard,onlyfansAssetsV2,ugcDelete,onlyfansQaCard,commentsNewVersion"

// dynamicDetailWebLocation 是动态详情接口默认的页面位置标识。
const dynamicDetailWebLocation = "333.1368"

// dynamicReactionWebLocation 是动态赞与转发列表接口默认的页面位置标识。
const dynamicReactionWebLocation = "333.1369"

// GetDynamicDetailInfoParam 指定要查询的动态与网页端可选参数。
type GetDynamicDetailInfoParam struct {
	ID                 int64  `json:"id" request:"query"`                               // 动态 ID，必填
	TimezoneOffset     int    `json:"timezone_offset" request:"query,omitempty"`        // 时区偏移，网页端为 -480；传 0 时由库填入
	Platform           string `json:"platform" request:"query,omitempty"`               // 平台，网页端为 web
	GaiaSource         string `json:"gaia_source" request:"query,omitempty"`            // 来源，网页端为 main_web
	Features           string `json:"features" request:"query,omitempty"`               // 功能开关，留空由库填入网页端默认值
	WebLocation        string `json:"web_location" request:"query,omitempty"`           // 页面位置标识，留空由库填入默认值
	XBiliDeviceReqJSON string `json:"x-bili-device-req-json" request:"query,omitempty"` // 设备信息，JSON 字符串，例如 {"platform":"web","device":"pc"}
	XBiliWebReqJSON    string `json:"x-bili-web-req-json" request:"query,omitempty"`    // 请求信息，JSON 字符串，例如 {"spm_id":"333.1368"}
}

// GetDynamicDetailInfo 获取动态详情。
//
// 认证方式：Cookie（SESSDATA，非必要）；Wbi 签名非必要，此处未启用。
// data.item 的字段结构独立定义在 GetDynamicDetailInfoItem 中。
func (c *Client) GetDynamicDetailInfo(ctx context.Context, param GetDynamicDetailInfoParam) (*GetDynamicDetailInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/detail"
	)
	return bilibili.Execute[*GetDynamicDetailInfoResult](ctx, c.Client, method, url, normalizeGetDynamicDetailInfoParam(param))
}

// GetDynamicDetailDesktop 获取动态详情（桌面端接口）。
//
// 参数与返回字段目前与 GetDynamicDetailInfo 相同，但结果类型独立定义，
// 桌面端响应变化时不会影响移动端调用方。
func (c *Client) GetDynamicDetailDesktop(ctx context.Context, param GetDynamicDetailInfoParam) (*GetDynamicDetailDesktopResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/desktop/v1/detail"
	)
	return bilibili.Execute[*GetDynamicDetailDesktopResult](ctx, c.Client, method, url, normalizeGetDynamicDetailInfoParam(param))
}

// normalizeGetDynamicDetailInfoParam 补齐网页端固定带、调用方通常不关心的那几个参数。
func normalizeGetDynamicDetailInfoParam(param GetDynamicDetailInfoParam) GetDynamicDetailInfoParam {
	if param.TimezoneOffset == 0 {
		param.TimezoneOffset = -480
	}
	if param.Features == "" {
		param.Features = dynamicDetailFeatures
	}
	if param.WebLocation == "" {
		param.WebLocation = dynamicDetailWebLocation
	}
	return param
}

// GetDynamicReactionListParam 指定要查询的动态与翻页偏移。
type GetDynamicReactionListParam struct {
	ID          int64  `json:"id" request:"query"`                     // 动态 ID，必填
	Offset      string `json:"offset" request:"query,omitempty"`       // 翻页偏移，取上一页返回的 Offset
	WebLocation string `json:"web_location" request:"query,omitempty"` // 页面位置标识，留空由库填入默认值
}

// GetDynamicReactionList 获取动态的赞与转发列表。
//
// 认证方式：Cookie（SESSDATA）；不登录时总数与列表均为空。
func (c *Client) GetDynamicReactionList(ctx context.Context, param GetDynamicReactionListParam) (*GetDynamicReactionListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/detail/reaction"
	)
	if param.WebLocation == "" {
		param.WebLocation = dynamicReactionWebLocation
	}
	return bilibili.Execute[*GetDynamicReactionListResult](ctx, c.Client, method, url, param)
}

// GetDynamicForwardListParam 指定要查询的动态与翻页偏移。
type GetDynamicForwardListParam struct {
	ID     int64  `json:"id" request:"query"`               // 动态 ID，必填
	Offset string `json:"offset" request:"query,omitempty"` // 翻页偏移，取上一页返回的 Offset
}

// GetDynamicForwardList 获取动态的转发列表。
//
// 鉴权方式：请求头 User-Agent 需为有效浏览器。
func (c *Client) GetDynamicForwardList(ctx context.Context, param GetDynamicForwardListParam) (*GetDynamicForwardListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/detail/forward"
	)
	return bilibili.Execute[*GetDynamicForwardListResult](ctx, c.Client, method, url, param)
}

// GetDynamicPicturesParam 指定要查询的动态。
type GetDynamicPicturesParam struct {
	ID int64 `json:"id" request:"query"` // 动态 ID，必填
}

// GetDynamicPictures 获取动态中的图片列表。
func (c *Client) GetDynamicPictures(ctx context.Context, param GetDynamicPicturesParam) (*GetDynamicPicturesResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/detail/pic"
	)
	return bilibili.Execute[*GetDynamicPicturesResult](ctx, c.Client, method, url, param)
}

// GetDynamicForwardItemParam 指定要查询的转发动态与时区偏移。
type GetDynamicForwardItemParam struct {
	ID             int64 `json:"id" request:"query"`                        // 动态 ID，必填
	TimezoneOffset int   `json:"timezone_offset" request:"query,omitempty"` // 时区偏移，网页端取 (new Date).getTimezoneOffset()
}

// GetDynamicForwardItem 获取转发动态的信息（即某条转发动态所转发的原动态条目）。
//
// 认证方式：Cookie（SESSDATA）。
func (c *Client) GetDynamicForwardItem(ctx context.Context, param GetDynamicForwardItemParam) (*GetDynamicForwardItemResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/detail/forward/item"
	)
	return bilibili.Execute[*GetDynamicForwardItemResult](ctx, c.Client, method, url, param)
}
