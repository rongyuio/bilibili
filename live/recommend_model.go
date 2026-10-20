package live

// 直播推荐相关响应模型。

// GetLiveRecommendListResult 是主页直播推荐接口的 data 层。
type GetLiveRecommendListResult struct {
	RecommendRoomList []GetLiveRecommendListItem `json:"recommend_room_list"` // 推荐房间列表
	TopRoomID         int64                      `json:"top_room_id"`         // 置顶直播间号
}

// GetLiveRecommendListItem 是推荐房间列表中的一项。
type GetLiveRecommendListItem struct {
	HeadBox              *GetLiveRecommendListHeadBox               `json:"head_box"`               // 头像框，无则为 null
	AreaV2ID             int                                        `json:"area_v2_id"`             // 分区 ID
	AreaV2ParentID       int                                        `json:"area_v2_parent_id"`      // 父分区 ID
	AreaV2Name           string                                     `json:"area_v2_name"`           // 分区名称
	AreaV2ParentName     string                                     `json:"area_v2_parent_name"`    // 父分区名称
	BroadcastType        int                                        `json:"broadcast_type"`         // 广播类型。作用尚不明确
	Cover                string                                     `json:"cover"`                  // 封面 URL
	Link                 string                                     `json:"link"`                   // 直播间链接，仅 pathname 与 query 部分
	Online               int                                        `json:"online"`                 // 观看人数
	PendantInfo          map[string]GetLiveRecommendListPendantInfo `json:"pendant_Info"`           // 挂件信息，键为位置编号。作用尚不明确
	RoomID               int64                                      `json:"roomid"`                 // 直播间 ID
	Title                string                                     `json:"title"`                  // 直播间标题
	Uname                string                                     `json:"uname"`                  // 主播用户名
	Face                 string                                     `json:"face"`                   // 主播头像 URL
	Verify               GetLiveRecommendListVerify                 `json:"verify"`                 // 认证信息
	UID                  int64                                      `json:"uid"`                    // 主播用户 mid
	Keyframe             string                                     `json:"keyframe"`               // 关键帧 URL
	IsAutoPlay           int                                        `json:"is_auto_play"`           // 是否自动播放。作用尚不明确
	HeadBoxType          int                                        `json:"head_box_type"`          // 头像框类型。作用尚不明确
	Flag                 int                                        `json:"flag"`                   // 标记。作用尚不明确
	GroupID              int64                                      `json:"group_id"`               // 推荐分组 ID。作用尚不明确
	SessionID            string                                     `json:"session_id"`             // 会话 ID。格式为小写无分隔 UUID + 下划线 + 大写连字符 UUID
	ShowCallback         string                                     `json:"show_callback"`          // 展示回调 URL
	ClickCallback        string                                     `json:"click_callback"`         // 点击回调 URL
	SpecialID            int                                        `json:"special_id"`             // 特殊 ID。作用尚不明确
	WatchedShow          GetLiveRecommendListWatchedShow            `json:"watched_show"`           // 观看展示
	IsNFT                int                                        `json:"is_nft"`                 // 是否为 NFT 头像
	NFTDmark             string                                     `json:"nft_dmark"`              // 作用尚不明确
	IsAd                 bool                                       `json:"is_ad"`                  // 是否为广告
	AdTransparentContent any                                        `json:"ad_transparent_content"` // 作用尚不明确，通常为 null
	ShowAdIcon           bool                                       `json:"show_ad_icon"`           // 是否显示广告图标
	Status               bool                                       `json:"status"`                 // 状态。作用尚不明确
	Followers            int                                        `json:"followers"`              // 作用尚不明确，通常为 0
}

// GetLiveRecommendListHeadBox 是推荐房间的头像框信息。
type GetLiveRecommendListHeadBox struct {
	Name  string `json:"name"`  // 头像框名称
	Value string `json:"value"` // 头像框图片 URL
	Desc  string `json:"desc"`  // 头像框描述
}

// GetLiveRecommendListPendantInfo 是推荐房间的挂件信息。
type GetLiveRecommendListPendantInfo struct {
	Type      string `json:"type"`       // 挂件类型
	Name      string `json:"name"`       // 挂件名称
	Position  int    `json:"position"`   // 展示位置
	Text      string `json:"text"`       // 挂件文案
	BgColor   string `json:"bg_color"`   // 背景颜色
	BgPic     string `json:"bg_pic"`     // 背景图片 URL
	PendantID int64  `json:"pendant_id"` // 挂件 ID
	Priority  int    `json:"priority"`   // 优先级
	CreatedAt int64  `json:"created_at"` // 创建时间戳
}

// GetLiveRecommendListVerify 是推荐房间主播的认证信息。
type GetLiveRecommendListVerify struct {
	Role int    `json:"role"` // 认证角色
	Desc string `json:"desc"` // 认证描述
	Type int    `json:"type"` // 认证类型。-1：无。0：个人认证。1：机构认证
}

// GetLiveRecommendListWatchedShow 是推荐房间的观看展示信息。
type GetLiveRecommendListWatchedShow struct {
	Switch       bool   `json:"switch"`        // 作用尚不明确
	Num          int    `json:"num"`           // 看过人数
	TextSmall    string `json:"text_small"`    // 小文本
	TextLarge    string `json:"text_large"`    // 大文本，形如 xxx人看过
	Icon         string `json:"icon"`          // 浅色线条眼睛图标 URL
	IconLocation int    `json:"icon_location"` // 图标位置
	IconWeb      string `json:"icon_web"`      // 深色线条眼睛图标 URL（Web 端）
}
