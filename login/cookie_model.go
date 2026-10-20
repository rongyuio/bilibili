package login

// Cookie 刷新与校验相关响应模型。各接口的返回体独立定义，互不复用。

// GetWebCookieRefreshInfoResult 是获取 cookie 刷新信息接口的 data 对象。
type GetWebCookieRefreshInfoResult struct {
	Refresh   bool  `json:"refresh"`   // 是否应该刷新 Cookie。true-需要刷新，false-不需要刷新
	Timestamp int64 `json:"timestamp"` // 用于获取 refresh_csrf 的毫秒时间戳
}

// GetWebCookieRefreshCsrfResult 是获取 cookie 刷新口令接口的响应。
type GetWebCookieRefreshCsrfResult struct {
	RefreshCsrf string `json:"refresh_csrf"` // 实时刷新口令
}

// RefreshCookieResult 是刷新 Cookie 接口的 data 对象。
type RefreshCookieResult struct {
	Status       int    `json:"status"`        // 未知
	Message      string `json:"message"`       // 未知
	RefreshToken string `json:"refresh_token"` // 新的持久化刷新口令
}
