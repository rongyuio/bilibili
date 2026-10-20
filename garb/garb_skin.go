package garb

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// APP 端主题相关接口。各接口的返回体独立定义，互不复用。

// GetGarbSkinParam 是获取主题及加载动画的请求参数。
type GetGarbSkinParam struct {
	AccessKey string `json:"access_key" request:"query"` // APP登录Token
	Appkey    string `json:"appkey" request:"query"`     // APP密钥
	Build     int    `json:"build" request:"query"`      // 版本。可为 6082000
	Ts        int64  `json:"ts" request:"query"`         // 当前时间戳。可为 0
	Sign      string `json:"sign" request:"query"`       // APP签名
}

// GetGarbSkinResult 是获取主题及加载动画的 data 对象。
type GetGarbSkinResult struct {
	UserEquip  *GetGarbSkinUserEquip `json:"user_equip"`  // 个性主题信息。有效时有此项
	SkinColors []GetGarbSkinColor    `json:"skin_colors"` // 主题颜色列表
	LoadEquip  *GetGarbSkinLoadEquip `json:"load_equip"`  // 加载动画信息。有效时有此项
}

// GetGarbSkinUserEquip 是个性主题信息。
type GetGarbSkinUserEquip struct {
	ID         int64                `json:"id"`          // 装扮id
	Name       string               `json:"name"`        // 装扮名称
	Preview    string               `json:"preview"`     // 装扮封面url
	Ver        int64                `json:"ver"`         // 装扮版本。时间戳
	PackageURL string               `json:"package_url"` // 装扮包url
	PackageMD5 string               `json:"package_md5"` // 装扮包md5校验值
	Data       GetGarbSkinEquipData `json:"data"`        // 装扮配置
}

// GetGarbSkinEquipData 是装扮配置。
type GetGarbSkinEquipData struct {
	ColorMode                 string `json:"color_mode"`                    // 颜色模式。light：亮色，dark：暗色
	Color                     string `json:"color"`                         // 前景色。十六进制颜色值，例：#ffffff
	ColorSecondPage           string `json:"color_second_page"`             // 背景色。十六进制颜色值
	TailColor                 string `json:"tail_color"`                    // 底边栏颜色
	TailColorSelected         string `json:"tail_color_selected"`           // 底边栏颜色（选择时）
	TailIconAni               bool   `json:"tail_icon_ani"`                 // 有无底边栏动画。false：无，true：有
	TailIconAniMode           string `json:"tail_icon_ani_mode"`            // 底边栏动画循环播放。once：播放一次
	HeadMyselfMp4Play         string `json:"head_myself_mp4_play"`          // 我的页面头图（视频）循环。once：播放一次，loop：循环播放
	TailIconMode              string `json:"tail_icon_mode"`                // 底栏改图标或改颜色。img：图标，color：颜色
	SideBgColor               string `json:"side_bg_color"`                 // 侧边栏颜色。5.x 版本客户端的侧边栏
	SideLineColor             string `json:"side_line_color"`               // 侧边栏线条颜色。5.x 版本客户端的侧边栏
	TailIconColor             string `json:"tail_icon_color"`               // 底栏图标颜色。tail_icon_mode=color 时必要
	TailIconColorDark         string `json:"tail_icon_color_dark"`          // 底栏图标颜色（夜间模式）。tail_icon_mode=color 时必要
	TailIconColorSelected     string `json:"tail_icon_color_selected"`      // 底栏图标颜色（选择时）。tail_icon_mode=color 时必要
	TailIconColorSelectedDark string `json:"tail_icon_color_selected_dark"` // 底栏图标颜色（选择时）（夜间模式）。tail_icon_mode=color 时必要
}

// GetGarbSkinColor 是一种主题颜色。
type GetGarbSkinColor struct {
	ID        int    `json:"id"`         // 颜色id
	Name      string `json:"name"`       // 颜色名称
	IsFree    bool   `json:"is_free"`    // 是否免费。false：收费，true：免费
	Price     int    `json:"price"`      // 价格。单位为硬币
	IsBought  bool   `json:"is_bought"`  // 是否已购买。false：未购买，true：已购买
	Status    int    `json:"status"`     // 状态。1：自动续费，4：已退订
	BuyTime   int64  `json:"buy_time"`   // 购买时间。毫秒时间戳
	DueTime   int64  `json:"due_time"`   // 到期时间。毫秒时间戳
	ColorName string `json:"color_name"` // 颜色类型名称
	IsOverdue bool   `json:"is_overdue"` // 是否已到期。false：未到期，true：已到期
}

// GetGarbSkinLoadEquip 是加载动画信息。
type GetGarbSkinLoadEquip struct {
	ID         int64  `json:"id"`          // 装扮id
	Name       string `json:"name"`        // 装扮名称
	Ver        int64  `json:"ver"`         // 装扮版本。时间戳
	LoadingURL string `json:"loading_url"` // 加载动画图标url
}

// GetGarbSkin 获取主题及加载动画。
//
// 仅可 APP 认证：appkey、ts、sign 等参数由调用方按 APP 签名规则生成，库不计算签名。
// 接口文档标注为未验证。
func (c *Client) GetGarbSkin(ctx context.Context, param GetGarbSkinParam) (*GetGarbSkinResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/resource/show/skin"
	)
	return bilibili.Execute[*GetGarbSkinResult](ctx, c.Client, method, url, param)
}
