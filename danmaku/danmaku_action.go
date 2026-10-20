package danmaku

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 弹幕操作接口。各接口的返回体独立定义，不与其他接口复用。

// danmakuActionWebLocation 是普通视频弹幕操作默认的页面位置标识。
const danmakuActionWebLocation = "1315873"

// danmakuWebLocationOrDefault 在调用方未指定页面标识时返回接口默认值。
func danmakuWebLocationOrDefault(webLocation, fallback string) string {
	if webLocation == "" {
		return fallback
	}
	return webLocation
}

// SendVideoDanmakuParam 指定要发送的视频弹幕。
// 字段默认编码进 query，再由内部处理函数移入 URL 编码表单。
type SendVideoDanmakuParam struct {
	AccessKey    string `json:"access_key,omitempty" request:"query,omitempty"`    // APP 登录 Token
	Type         int    `json:"type" request:"query"`                              // 弹幕类选择。1：视频弹幕；2：漫画弹幕
	Oid          int64  `json:"oid" request:"query"`                               // 视频 cid
	Msg          string `json:"msg" request:"query"`                               // 弹幕内容，长度小于 100 字符
	Bvid         string `json:"bvid,omitempty" request:"query,omitempty"`          // 稿件 bvid。avid 与 bvid 任选一个
	Aid          int64  `json:"aid,omitempty" request:"query,omitempty"`           // 稿件 avid。avid 与 bvid 任选一个
	Progress     int64  `json:"progress,omitempty" request:"query,omitempty"`      // 弹幕出现在视频内的时间，单位毫秒，默认为 0
	Color        int    `json:"color,omitempty" request:"query,omitempty"`         // 十进制 RGB888 颜色值，如 16777215 为白色
	Fontsize     int    `json:"fontsize,omitempty" request:"query,omitempty"`      // 字号。12：极小；16：超小；18：小；25：标准；36：大；45：超大；64：极大。默认 25
	Pool         int    `json:"pool,omitempty" request:"query,omitempty"`          // 弹幕池选择。0：普通池；1：字幕池；2：特殊池（代码/BAS 弹幕）。默认 0
	Mode         int    `json:"mode" request:"query"`                              // 弹幕类型。1：普通弹幕；4：底部弹幕；5：顶部弹幕；7：高级弹幕；9：BAS 弹幕（pool 必须为 2）
	Rnd          int64  `json:"rnd,omitempty" request:"query,omitempty"`           // 当前时间戳 * 1000000。省略时发送冷却时间为 90s，带上则为 5s
	Colorful     int    `json:"colorful,omitempty" request:"query,omitempty"`      // 彩色弹幕。60001：专属渐变彩色（需要会员）
	CheckboxType int    `json:"checkbox_type,omitempty" request:"query,omitempty"` // 是否带 UP 身份标识。0：普通；4：带有标识
	WebLocation  string `json:"web_location,omitempty" request:"query,omitempty"`  // URL 参数，普通视频固定 1315873，留空由库填入默认值
}

// SendVideoDanmaku 发送一条视频弹幕。
//
// 需要登录态：CSRF 取自 Cookie bili_jct。请求使用 WBI 签名，只对 query 参数签名，
// 业务参数走 URL 编码表单；web_location 保留在 query 中参与签名。
// mode=6 的逆向弹幕与 mode=8 的代码弹幕不可发送。
func (c *Client) SendVideoDanmaku(ctx context.Context, param SendVideoDanmakuParam) (*SendVideoDanmakuResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/dm/post"
	)
	param.WebLocation = danmakuWebLocationOrDefault(param.WebLocation, danmakuActionWebLocation)
	return bilibili.Execute[*SendVideoDanmakuResult](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("access_key", "type", "oid", "msg", "bvid", "aid", "progress",
			"color", "fontsize", "pool", "mode", "rnd", "colorful", "checkbox_type"),
		bilibili.FillFormCsrf(c.Client), c.FillWbi())
}

// RecallDanmakuParam 指定要撤回的弹幕。
type RecallDanmakuParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Dmid      int64  `json:"dmid" request:"query"`                           // 弹幕 dmid
	Cid       int64  `json:"cid" request:"query"`                            // 视频 cid
}

// RecallDanmaku 撤回自己发送的弹幕。
//
// 仅能撤回两分钟内的弹幕，且每天只有 3 次机会。需要登录态：CSRF 取自 Cookie bili_jct。
func (c *Client) RecallDanmaku(ctx context.Context, param RecallDanmakuParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/dm/recall"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("access_key", "dmid", "cid"), bilibili.FillFormCsrf(c.Client))
	return err
}

// CheckDanmakuAdvPermissionParam 指定要查询高级弹幕发送权限的视频。
type CheckDanmakuAdvPermissionParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Mode      string `json:"mode" request:"query"`                           // 固定为 sp
	Cid       int64  `json:"cid" request:"query"`                            // 视频 cid
}

// CheckDanmakuAdvPermission 检测视频高级弹幕发送权限。
//
// 响应 data 在有效时是状态对象，无效时为 null，此时返回 nil。
func (c *Client) CheckDanmakuAdvPermission(ctx context.Context, param CheckDanmakuAdvPermissionParam) (*CheckDanmakuAdvPermissionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/dm/adv/state"
	)
	return bilibili.Execute[*CheckDanmakuAdvPermissionResult](ctx, c.Client, method, url, param)
}
