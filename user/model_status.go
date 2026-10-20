package user

// 用户状态数相关响应模型。

// GetRelationStatResult 是查询关系状态数的 data 对象。
type GetRelationStatResult struct {
	Mid            int `json:"mid"`              // 目标用户mid
	Following      int `json:"following"`        // 关注数
	Whisper        int `json:"whisper"`          // 悄悄关注数。需要登录(Cookie或APP)，未登录或非自己恒为0
	Black          int `json:"black"`            // 黑名单数。需要登录(Cookie或APP)，未登录或非自己恒为0
	Follower       int `json:"follower"`         // 粉丝数
	FansMedalToast any `json:"fans_medal_toast"` // (?)。实测为 null
	FansEffect     any `json:"fans_effect"`      // (?)。实测为 null
}

// GetUpStatResult 是查询UP主状态数的 data 对象。
type GetUpStatResult struct {
	Archive *GetUpStatView `json:"archive"` // 视频播放量
	Article *GetUpStatView `json:"article"` // 专栏阅读量
	Likes   int            `json:"likes"`   // 获赞次数
}

// GetUpStatView 是 UP主状态数中 archive / article 对象的统计值。
type GetUpStatView struct {
	View int `json:"view"` // 视频播放量（archive）或专栏阅读量（article）
}

// GetNavNumResult 是查询用户导航栏状态数的 data 对象。
type GetNavNumResult struct {
	Video     int                 `json:"video"`      // 投稿视频数
	Bangumi   int                 `json:"bangumi"`    // 追番数。无视隐私设置
	Cinema    int                 `json:"cinema"`     // 追剧数。无视隐私设置
	Channel   *GetNavNumChannel   `json:"channel"`    // 视频列表数
	Favourite *GetNavNumFavourite `json:"favourite"`  // 收藏夹数
	Tag       int                 `json:"tag"`        // 关注TAG数。无视隐私设置
	Article   int                 `json:"article"`    // 投稿专栏数
	Playlist  int                 `json:"playlist"`   // 0。作用尚不明确
	Album     int                 `json:"album"`      // 投稿图文数
	Audio     int                 `json:"audio"`      // 投稿音频数
	Pugv      int                 `json:"pugv"`       // 投稿课程数
	Opus      int                 `json:"opus"`       // 动态数。文档字段表写作 upos，实测响应为 opus
	SeasonNum int                 `json:"season_num"` // 视频合集数
}

// GetNavNumChannel 是用户导航栏状态数中的视频列表数。
type GetNavNumChannel struct {
	Master int `json:"master"` // 视频列表数
	Guest  int `json:"guest"`  // 视频列表数
}

// GetNavNumFavourite 是用户导航栏状态数中的收藏夹数。
type GetNavNumFavourite struct {
	Master int `json:"master"` // 全部收藏夹数。需要登录(SESSDATA)，只能查看自己的
	Guest  int `json:"guest"`  // 公开收藏夹数
}

// GetAlbumUploadCountResult 是查询相簿投稿数的 data 对象。
type GetAlbumUploadCountResult struct {
	AllCount   int `json:"all_count"`   // 相簿总数。总数为以下三者之和
	DrawCount  int `json:"draw_count"`  // 发布绘画数
	PhotoCount int `json:"photo_count"` // 发布摄影数
	DailyCount int `json:"daily_count"` // 发布日常（图片动态）数
}
