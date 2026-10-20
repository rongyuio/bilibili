package login

// 登录操作（人机验证）相关响应模型。

// GetOldCaptchaResult 是申请 captcha 验证码（旧版）的 data 对象。
type GetOldCaptchaResult struct {
	Result GetOldCaptchaResultData `json:"result"` // 套了个娃
	Type   int                     `json:"type"`   // 1。作用尚不明确
}

// GetOldCaptchaResultData 是申请 captcha 验证码（旧版）data.result 中的对象。
type GetOldCaptchaResultData struct {
	Success   int    `json:"success"`   // 1。作用尚不明确
	Gt        string `json:"gt"`        // 极验id。一般为固定值
	Challenge string `json:"challenge"` // 极验KEY。由B站后端产生用于人机验证
	Key       string `json:"key"`       // 登录秘钥。与 captcha 无关，与登录接口有关，亦作 token
}
