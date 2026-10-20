package bangumi

// GetBangumiTimeline 接口的响应模型。
//
// 本文件的类型只服务于 GetBangumiTimeline，不与其它接口共用。

// 番剧或影视时间线相关响应模型。

// GetBangumiTimelineDay 时间线中的一天。
type GetBangumiTimelineDay struct {
	Date      string                      `json:"date"`        // 当日日期。格式为 m-d
	DateTs    int64                       `json:"date_ts"`     // 当日日期时间戳。UNIX 秒级时间戳
	DayOfWeek int                         `json:"day_of_week"` // 一周中第几天。取值范围为 [1, 7]
	Episodes  []GetBangumiTimelineEpisode `json:"episodes"`    // 剧集列表
	IsToday   int                         `json:"is_today"`    // 是否今日。0：否。1：是
}

// GetBangumiTimelineEpisode 时间线中的一集。
type GetBangumiTimelineEpisode struct {
	Cover       string `json:"cover"`        // 封面图 url
	Delay       int    `json:"delay"`        // 是否推迟
	DelayID     int    `json:"delay_id"`     // 推迟一话 epid
	DelayIndex  string `json:"delay_index"`  // 推迟一话名称
	DelayReason string `json:"delay_reason"` // 推迟原因
	EpCover     string `json:"ep_cover"`     // 最新一话图 url
	EpisodeID   int    `json:"episode_id"`   // 最新一话的 epid
	PubIndex    string `json:"pub_index"`    // 最新一话名称
	PubTime     string `json:"pub_time"`     // 发布时间。格式为 HH:mm
	PubTs       int64  `json:"pub_ts"`       // 发布时间戳。UNIX 秒级时间戳
	Published   int    `json:"published"`    // 是否已发布。0：否。1：是
	Follows     string `json:"follows"`      // 追番人数。未公开时为 "-"
	Plays       string `json:"plays"`        // 播放数。未公开时为 "-"
	SeasonID    int    `json:"season_id"`    // 剧集 ssid
	SquareCover string `json:"square_cover"` // 缩略图 url
	Title       string `json:"title"`        // 剧集标题
}
