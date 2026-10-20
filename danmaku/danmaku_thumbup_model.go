package danmaku

// 弹幕点赞查询响应模型。返回体独立定义，不与其他接口复用。

// GetDanmakuThumbupStatsItem 单条弹幕的点赞信息。
type GetDanmakuThumbupStatsItem struct {
	Likes              int                                       `json:"likes"`                // 对应弹幕所获得的点赞数
	UserLike           int                                       `json:"user_like"`            // 当前用户是否点赞。0：未点赞，1：已点赞。需要登录（Cookie 或 APP）
	IDStr              string                                    `json:"id_str"`               // 弹幕 dmid 字符串形式
	ShowThreshold      int                                       `json:"show_threshold"`       // 作用尚不明确
	ReplyCount         int                                       `json:"reply_count"`          // 作用尚不明确
	GrayReleaseControl *GetDanmakuThumbupStatsGrayReleaseControl `json:"gray_release_control"` // 灰度发布控制信息
}

// GetDanmakuThumbupStatsGrayReleaseControl 弹幕点赞相关的灰度发布控制信息。
type GetDanmakuThumbupStatsGrayReleaseControl struct {
	ShowDmReply bool `json:"show_dm_reply"` // 作用尚不明确
}
