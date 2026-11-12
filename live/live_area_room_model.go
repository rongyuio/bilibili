package live

// 直播分区房间列表接口的响应模型。本接口的类型只服务本接口。

// GetLiveAreaRoomListItem 是直播分区房间列表中的一个直播间。
type GetLiveAreaRoomListItem struct {
	RoomID           int64  `json:"roomid"`              // 直播间号（长号）。注意字段名没有下划线
	UID              int64  `json:"uid"`                 // 主播 UID
	Title            string `json:"title"`               // 直播间标题
	Uname            string `json:"uname"`               // 主播昵称
	Online           int    `json:"online"`              // 人气值
	UserCover        string `json:"user_cover"`          // 主播自定义封面 URL
	UserCoverFlag    int    `json:"user_cover_flag"`     // 封面标记
	SystemCover      string `json:"system_cover"`        // 系统关键帧 URL
	Cover            string `json:"cover"`               // 实际展示的封面 URL
	ShowCover        string `json:"show_cover"`          // 展示封面 URL
	Link             string `json:"link"`                // 直播间链接，仅 pathname 与 query 部分
	Face             string `json:"face"`                // 主播头像 URL
	ParentID         int    `json:"parent_id"`           // 父分区 id
	ParentName       string `json:"parent_name"`         // 父分区名称
	AreaID           int    `json:"area_id"`             // 子分区 id
	AreaName         string `json:"area_name"`           // 子分区名称
	AreaV2ParentID   int    `json:"area_v2_parent_id"`   // 二级父分区 id
	AreaV2ParentName string `json:"area_v2_parent_name"` // 二级父分区名称
	AreaV2ID         int    `json:"area_v2_id"`          // 二级子分区 id
	AreaV2Name       string `json:"area_v2_name"`        // 二级子分区名称
	WebPendent       string `json:"web_pendent"`         // 网页端挂件信息
	PkID             int64  `json:"pk_id"`               // PK id
	SessionID        string `json:"session_id"`          // 推荐会话 id
	GroupID          int64  `json:"group_id"`            // 推荐分组 id
}
