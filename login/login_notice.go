package login

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 登录记录相关接口。响应模型见 login_notice_model.go。
// 各接口的返回体独立定义，互不复用。

// loginLogWebLocation 是登录记录接口默认的页面位置标识。
const loginLogWebLocation = "333.33"

// NoticeParam 指定要查询的登录记录。
type NoticeParam struct {
	Mid   int    `json:"mid" request:"query"`                       // 用户mid。必须为自己的mid
	Buvid string `json:"buvid,omitempty" request:"query,omitempty"` // 设备虚拟id。web端为操作登录接口时Cookie中的buvid3，若登录设备无buvid则留空
}

// GetLoginNotice 查询指定设备的登录记录。
//
// 需要登录态（Cookie SESSDATA），mid 必须为当前登录用户自己的 mid。
func (c *Client) GetLoginNotice(ctx context.Context, param NoticeParam) (*GetLoginNoticeResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/safecenter/login_notice"
	)
	return bilibili.Execute[*GetLoginNoticeResult](ctx, c.Client, method, url, param)
}

// LogParam 指定最近一周登录情况的查询参数。
type LogParam struct {
	Jsonp       string `json:"jsonp,omitempty" request:"query,omitempty"`        // 回调函数名。默认 jsonp
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 网页位置。默认 333.33，留空由库填入默认值
}

// GetLoginLog 查询当前用户最近一周的登录情况。
//
// 需要登录态（Cookie SESSDATA）。
func (c *Client) GetLoginLog(ctx context.Context, param LogParam) (*GetLoginLogResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/web/login/log"
	)
	if param.WebLocation == "" {
		param.WebLocation = loginLogWebLocation
	}
	return bilibili.Execute[*GetLoginLogResult](ctx, c.Client, method, url, param)
}
