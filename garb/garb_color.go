package garb

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 主题色相关接口。各接口的返回体独立定义，互不复用。

// GetGarbColorInfoFromAppParam 是获取主题色基本信息（APP 方式）的请求参数。
type GetGarbColorInfoFromAppParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token
	Appkey    string `json:"appkey" request:"query"`                         // APP密钥
	Build     int    `json:"build" request:"query"`                          // 版本。可为 6082000
	Sign      string `json:"sign" request:"query"`                           // APP签名
}

// GetGarbColorInfoFromAppItem 是获取主题色基本信息（APP 方式）返回的主题色。
type GetGarbColorInfoFromAppItem struct {
	ID        int    `json:"id"`         // 颜色id
	Name      string `json:"name"`       // 颜色名称
	IsFree    bool   `json:"is_free"`    // 是否免费。false：收费。true：免费
	Price     int    `json:"price"`      // 价格。单位为硬币
	IsBought  bool   `json:"is_bought"`  // 是否已购买。false：未购买。true：已购买
	Status    int    `json:"status"`     // 状态。1：自动续费。4：已退订
	BuyTime   int64  `json:"buy_time"`   // 购买时间。毫秒时间戳
	DueTime   int64  `json:"due_time"`   // 到期时间。毫秒时间戳
	ColorName string `json:"color_name"` // 颜色类型名称
	IsOverdue bool   `json:"is_overdue"` // 是否已到期。false：未到期。true：已到期
}

// GetGarbColorInfoFromApp 获取主题色基本信息（APP方式，appkey 鉴权）。
func (c *Client) GetGarbColorInfoFromApp(ctx context.Context, param GetGarbColorInfoFromAppParam) ([]GetGarbColorInfoFromAppItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://club.bilibili.com/api/query.skin.list.do"
	)
	return bilibili.Execute[[]GetGarbColorInfoFromAppItem](ctx, c.Client, method, url, param)
}

// GetGarbColorInfoParam 是获取主题色基本信息（Cookie 方式）的请求参数。
type GetGarbColorInfoParam struct {
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 android
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 版本。可为 6082000
}

// GetGarbColorInfoItem 是获取主题色基本信息（Cookie 方式）返回的主题色。
type GetGarbColorInfoItem struct {
	ID        int    `json:"id"`         // 颜色id
	Name      string `json:"name"`       // 颜色名称
	IsFree    bool   `json:"is_free"`    // 是否免费。false：收费。true：免费
	Price     int    `json:"price"`      // 价格。单位为硬币
	IsBought  bool   `json:"is_bought"`  // 是否已购买。false：未购买。true：已购买
	Status    int    `json:"status"`     // 状态。1：自动续费。4：已退订
	BuyTime   int64  `json:"buy_time"`   // 购买时间。毫秒时间戳
	DueTime   int64  `json:"due_time"`   // 到期时间。毫秒时间戳
	ColorName string `json:"color_name"` // 颜色类型名称
	IsOverdue bool   `json:"is_overdue"` // 是否已到期。false：未到期。true：已到期
}

// GetGarbColorInfo 获取主题色基本信息（Cookie 方式）。
func (c *Client) GetGarbColorInfo(ctx context.Context, param GetGarbColorInfoParam) ([]GetGarbColorInfoItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/garb/skin/color/list"
	)
	return bilibili.Execute[[]GetGarbColorInfoItem](ctx, c.Client, method, url, param)
}
