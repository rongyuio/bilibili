package bilibili

import "encoding/json"

// 登录相关响应模型。

type Geetest struct {
	Gt        string `json:"gt"`        // 极验id。一般为固定值
	Challenge string `json:"challenge"` // 极验KEY。由B站后端产生用于人机验证
}

type CaptchaResult struct {
	Geetest Geetest `json:"geetest"` // 极验captcha数据
	Tencent any     `json:"tencent"` // (?)。**作用尚不明确**
	Token   string  `json:"token"`   // 登录 API token。与 captcha 无关，与登录接口有关
	Type    string  `json:"type"`    // 验证方式。用于判断使用哪一种验证方式，目前所见只有极验。geetest：极验
}

type LoginWithPasswordResult struct {
	Message      string `json:"message"`       // 扫码状态信息
	RefreshToken string `json:"refresh_token"` // 刷新refresh_token
	Status       int    `json:"status"`        // 成功为0
	Timestamp    int    `json:"timestamp"`     // 登录时间。未登录为0。时间戳 单位为毫秒
	URL          string `json:"url"`           // 游戏分站跨域登录 url
}

type CountryCrown struct {
	ID        int    `json:"id"`         // 国际代码值
	Cname     string `json:"cname"`      // 国家或地区名
	CountryID string `json:"country_id"` // 国家或地区区号
}

type GetCountryCrownResult struct {
	Common []CountryCrown `json:"common"` // 常用国家&地区
	Others []CountryCrown `json:"others"` // 其他国家&地区
}

type SendSMSResult struct {
	CaptchaKey string `json:"captcha_key"` // 短信登录 token。在下方传参时需要，请备用
}

type LoginWithSMSResult struct {
	IsNew  bool   `json:"is_new"` // 是否为新注册用户。false：非新注册用户。true：新注册用户
	Status int    `json:"status"` // 0。未知，可能0就是成功吧
	URL    string `json:"url"`    // 跳转 url。默认为 https://www.bilibili.com
}

type LoginWithQRCodeResult struct {
	URL          string `json:"url"`           // 游戏分站跨域登录 url。未登录为空
	RefreshToken string `json:"refresh_token"` // 刷新refresh_token。未登录为空
	Timestamp    int    `json:"timestamp"`     // 登录时间。未登录为0。时间戳 单位为毫秒
	Code         int    `json:"code"`          // 0：扫码登录成功。86038：二维码已失效。86090：二维码已扫码未确认。86101：未扫码
	Message      string `json:"message"`       // 扫码状态信息
}

type AccountInformation struct {
	Mid      json.Number `json:"mid"`       // 我的mid
	Uname    string      `json:"uname"`     // 我的昵称
	Userid   string      `json:"userid"`    // 我的用户名
	Sign     string      `json:"sign"`      // 我的签名
	Birthday string      `json:"birthday"`  // 我的生日。YYYY-MM-DD
	Sex      string      `json:"sex"`       // 我的性别。男 女 保密
	NickFree bool        `json:"nick_free"` // 是否未设置昵称。false：设置过昵称。true：未设置昵称
	Rank     string      `json:"rank"`      // 我的会员等级
}
type QRCode struct {
	URL       string `json:"url"`        // 二维码内容 (登录页面 url)
	QrcodeKey string `json:"qrcode_key"` // 扫码登录秘钥。恒为32字符
}
