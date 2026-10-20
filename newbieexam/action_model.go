package newbieexam

// 新用户答题操作相关响应模型。

// GetNewbieExamCaptchaResult 是获取验证码的响应 data。
type GetNewbieExamCaptchaResult struct {
	Type      string `json:"type"`      // 验证码类型。geetest：极验
	Gt        string `json:"gt"`        // 极验id
	Challenge string `json:"challenge"` // 极验key
	Token     string `json:"token"`     // (?)。作用尚不明确，实测为空
	URL       string `json:"url"`       // (?)。作用尚不明确，实测为空
}
