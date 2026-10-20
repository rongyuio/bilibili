package login

// 登录相关响应模型（web 端旧版扫码登录）。

// GetQRCodeOldResult 是申请二维码（web 端旧版）的 data 对象。
type GetQRCodeOldResult struct {
	URL      string `json:"url"`      // 二维码内容 (登录页面 url)
	OAuthKey string `json:"oauthKey"` // 扫码登录秘钥。恒为32字符
}
