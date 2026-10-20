package live

// 用户关注的 UP 主直播情况相关响应模型（live-follow-up-live 域）。

// GetFollowingLiveStatusResult 是关注直播情况接口的 data 层。
type GetFollowingLiveStatusResult struct {
	Title           string                       `json:"title"`             // 标题
	PageSize        int                          `json:"pageSize"`          // 每页的数据数量
	TotalPage       int                          `json:"totalPage"`         // 分页数量
	List            []GetFollowingLiveStatusItem `json:"list"`              // UP 直播情况列表
	Count           int                          `json:"count"`             // 曾直播过的 UP 数量
	NeverLivedCount int                          `json:"never_lived_count"` // 未直播过的 UP 数量
	LiveCount       int                          `json:"live_count"`        // 正在直播的 UP 数量
	NeverLivedFaces []string                     `json:"never_lived_faces"` // 含义尚不明确
}

// GetFollowingLiveStatusItem 是关注直播情况列表中的一项。
type GetFollowingLiveStatusItem struct {
	RoomID           int64  `json:"roomid"`              // 房间号
	UID              int64  `json:"uid"`                 // 主播 uid
	Uname            string `json:"uname"`               // 主播名
	Title            string `json:"title"`               // 直播标题
	Face             string `json:"face"`                // 主播头像
	LiveStatus       int    `json:"live_status"`         // 是否正在直播：0 未直播，1 正在直播
	RecordNum        int    `json:"record_num"`          // 含义尚不明确，一般为 0
	RecentRecordID   string `json:"recent_record_id"`    // 含义尚不明确，一般为空
	IsAttention      int    `json:"is_attention"`        // 含义尚不明确，一般为 1
	ClipNum          int    `json:"clipnum"`             // 含义尚不明确，一般为 0
	FansNum          int    `json:"fans_num"`            // 含义尚不明确，一般为 0
	AreaName         string `json:"area_name"`           // 含义尚不明确，一般为空
	AreaValue        string `json:"area_value"`          // 含义尚不明确，一般为空
	Tags             string `json:"tags"`                // 含义尚不明确，一般为空
	RecentRecordIDV2 string `json:"recent_record_id_v2"` // 含义尚不明确，一般为空
	RecordNumV2      int    `json:"record_num_v2"`       // 含义尚不明确，一般为 0
	RecordLiveTime   int64  `json:"record_live_time"`    // 主播上一次直播结束的时间戳，正在直播时为 0
	AreaNameV2       string `json:"area_name_v2"`        // 频道的名称
	RoomNews         string `json:"room_news"`           // 房间公告
	Switch           bool   `json:"switch"`              // 含义尚不明确
	WatchIcon        string `json:"watch_icon"`          // 含义尚不明确
	TextSmall        string `json:"text_small"`          // 含义尚不明确，当主播正在直播时可能为在线人数
	RoomCover        string `json:"room_cover"`          // 房间封面图片的 URL
	ParentAreaID     int    `json:"parent_area_id"`      // 父分区 id
	AreaID           int    `json:"area_id"`             // 分区 id
}

// GetFollowingLiveWebListResult 是 PC 端正在直播列表接口的 data 层。
type GetFollowingLiveWebListResult struct {
	Rooms        []GetFollowingLiveWebListItem `json:"rooms"`          // 正在直播的房间列表；hit_ab 为 false 时只会获取前 10 个
	List         []GetFollowingLiveWebListItem `json:"list"`           // 正在直播的房间列表，疑似与 rooms 内容相同
	Count        int                           `json:"count"`          // 关注列表中正在直播的人数；hit_ab 为 true 时为 0
	NotLivingNum int                           `json:"not_living_num"` // 关注列表中未开播的人数；hit_ab 为 false 时为 0
}

// GetFollowingLiveWebListItem 是 PC 端正在直播列表中的房间信息。
type GetFollowingLiveWebListItem struct {
	Title            string `json:"title"`               // 直播间标题
	RoomID           int64  `json:"room_id"`             // 直播间真实 id
	UID              int64  `json:"uid"`                 // 目标用户 mid
	Online           int    `json:"online"`              // 观看人数；hit_ab 为 true 时为 0
	LiveTime         int    `json:"live_time"`           // 已经直播的时长（秒）；hit_ab 为 true 时为 0
	LiveStatus       int    `json:"live_status"`         // 开播状态：0 未开播，1 直播中，2 轮播中
	ShortID          int    `json:"short_id"`            // 直播间短 id；hit_ab 为 true 时为 0
	Area             int    `json:"area"`                // 分区 id；hit_ab 为 true 时为 0
	AreaName         string `json:"area_name"`           // 分区名称
	AreaV2ID         int    `json:"area_v2_id"`          // 二级分区 id
	AreaV2Name       string `json:"area_v2_name"`        // 二级分区名
	AreaV2ParentName string `json:"area_v2_parent_name"` // 二级父分区名
	AreaV2ParentID   int    `json:"area_v2_parent_id"`   // 二级父分区 id
	Uname            string `json:"uname"`               // 用户名
	Face             string `json:"face"`                // 用户头像图片链接
	TagName          string `json:"tag_name"`            // 标签名
	Tags             string `json:"tags"`                // 标签列表
	CoverFromUser    string `json:"cover_from_user"`     // 直播间封面图片链接；hit_ab 为 true 时为 0
	Keyframe         string `json:"keyframe"`            // 关键帧图片链接；hit_ab 为 true 时为 0
	LockTill         string `json:"lock_till"`           // 未知，格式 yyyy-MM-dd hh-mm-ss；hit_ab 为 true 时为空字符串
	HiddenTill       string `json:"hidden_till"`         // 未知，格式 yyyy-MM-dd hh-mm-ss；hit_ab 为 true 时为空字符串
	BroadcastType    int    `json:"broadcast_type"`      // 广播类型
	IsEncrypt        bool   `json:"is_encrypt"`          // 直播间是否加密
	Link             string `json:"link"`                // 直播间链接；hit_ab 为 true 时为空字符串
	Nickname         string `json:"nickname"`            // 用户昵称；hit_ab 为 true 时为空字符串
	Roomname         string `json:"roomname"`            // 直播间名称；hit_ab 为 true 时为空字符串
	Roomid           int64  `json:"roomid"`              // 直播间真实 id；hit_ab 为 true 时为 0
	LiveStartTime    int64  `json:"liveTime"`            // 开播时间；hit_ab 为 true 时为 0
}
