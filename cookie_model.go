package bilibili

// Cookie 刷新与校验相关响应模型。

type GetWebCookieRefreshInfoResult struct {
	Refresh   bool  `json:"refresh"`   // 是否应该刷新 Cookie。true-需要刷新，false-不需要刷新
	Timestamp int64 `json:"timestamp"` // 用于获取 refresh_csrf 的毫秒时间戳
}

type GetWebCookieRefreshCsrfResult struct {
	RefreshCsrf string `json:"refresh_csrf"` // 实时刷新口令
}

type RefreshCookieResult struct {
	Status       int    `json:"status"`        // 未知
	Message      string `json:"message"`       // 未知
	RefreshToken string `json:"refresh_token"` // 新的持久化刷新口令
}
