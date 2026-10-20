package dynamic

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户空间动态（桌面端）相关接口。响应模型见 desktop_feed_model.go。

// dynamicDesktopSpaceFeatures 是桌面端用户空间动态默认请求的功能开关串。
const dynamicDesktopSpaceFeatures = "itemOpusStyle,listOnlyfans,opusBigCover,onlyfansVote,forwardListHidden,decorationCard,commentsNewVersion,onlyfansAssetsV2,ugcDelete,onlyfansQaCard"

// dynamicDesktopSpaceWebLocation 是桌面端用户空间动态默认的页面位置标识。
const dynamicDesktopSpaceWebLocation = "333.1387"

// GetUserSpaceDynamicDesktopParam 指定桌面端用户空间动态的查询条件。
type GetUserSpaceDynamicDesktopParam struct {
	HostMid            string `json:"host_mid" request:"query"`                                   // 被查询用户 UID（mid），必填
	Offset             string `json:"offset,omitempty" request:"query,omitempty"`                 // 分页偏移量
	TimezoneOffset     int    `json:"timezone_offset,omitempty" request:"query,omitempty"`        // 时区偏移，默认 -480；传 0 时由库填入
	Platform           string `json:"platform,omitempty" request:"query,omitempty"`               // 平台，例如 web
	Features           string `json:"features,omitempty" request:"query,omitempty"`               // 功能开关，留空由库填入默认值
	WebLocation        string `json:"web_location,omitempty" request:"query,omitempty"`           // 页面位置标识，留空由库填入默认值
	DmImgSwitch        int    `json:"dm_img_switch,omitempty" request:"query,omitempty"`          // dm_img 风控开关，仅登录时存在
	DmImgList          string `json:"dm_img_list,omitempty" request:"query,omitempty"`            // dm_img 系列风控参数，JSON 字符串，仅未登录时存在
	DmImgStr           string `json:"dm_img_str,omitempty" request:"query,omitempty"`             // dm_img 系列风控参数，仅未登录时存在
	DmCoverImgStr      string `json:"dm_cover_img_str,omitempty" request:"query,omitempty"`       // dm_img 系列风控参数，仅未登录时存在
	DmImgInter         string `json:"dm_img_inter,omitempty" request:"query,omitempty"`           // dm_img 系列风控参数，JSON 字符串，仅未登录时存在
	XBiliDeviceReqJSON string `json:"x-bili-device-req-json,omitempty" request:"query,omitempty"` // 设备信息，JSON 字符串，例如 {"platform":"web","device":"pc"}
	XBiliWebReqJSON    string `json:"x-bili-web-req-json,omitempty" request:"query,omitempty"`    // 请求信息，JSON 字符串，例如 {"spm_id":"333.1387"}
}

// GetUserSpaceDynamicDesktop 获取用户空间动态（桌面端接口）
//
// 文档中的桌面端地址，参数与返回字段与 GetUserSpaceDynamic 相同，但结果类型
// 独立定义。未登录时需要 Cookie 含 buvid3，并可能需要 WBI 签名与 dm_img 风控参数；
// 登录后（SESSDATA）更稳定。
func (c *Client) GetUserSpaceDynamicDesktop(ctx context.Context, param GetUserSpaceDynamicDesktopParam) (*GetUserSpaceDynamicDesktopResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/desktop/v1/feed/space"
	)
	return bilibili.Execute[*GetUserSpaceDynamicDesktopResult](ctx, c.Client, method, url, normalizeGetUserSpaceDynamicDesktopParam(param))
}

// normalizeGetUserSpaceDynamicDesktopParam 补齐桌面端固定带、调用方通常不关心的那几个参数。
func normalizeGetUserSpaceDynamicDesktopParam(param GetUserSpaceDynamicDesktopParam) GetUserSpaceDynamicDesktopParam {
	if param.TimezoneOffset == 0 {
		param.TimezoneOffset = -480
	}
	if param.Features == "" {
		param.Features = dynamicDesktopSpaceFeatures
	}
	if param.WebLocation == "" {
		param.WebLocation = dynamicDesktopSpaceWebLocation
	}
	return param
}
