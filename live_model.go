package bilibili

// 直播勋章相关响应模型。字段定义沿用本地工具，未与线上返回逐字段核对；
// 勋章墙与面板的用户勋章字段完全一致，合并为 LiveUinfoMedal 并保留别名。

// GetLiveMedalWallResult contains the corresponding live medal response fields.
type GetLiveMedalWallResult struct {
	List            []LiveMedalWallItem `json:"list"`
	Count           int                 `json:"count"`
	CloseSpaceMedal int                 `json:"close_space_medal"`
	OnlyShowWearing int                 `json:"only_show_wearing"`
	Name            string              `json:"name"`
	Icon            string              `json:"icon"`
	UID             int                 `json:"uid"`
	Level           int                 `json:"level"`
}

// LiveMedalWallItem contains the corresponding live medal response fields.
type LiveMedalWallItem struct {
	MedalInfo  LiveMedalWallItemMedalInfo  `json:"medal_info"`
	TargetName string                      `json:"target_name"`
	TargetIcon string                      `json:"target_icon"`
	Link       string                      `json:"link"`
	LiveStatus int                         `json:"live_status"`
	Official   int                         `json:"official"`
	UinfoMedal LiveMedalWallItemUinfoMedal `json:"uinfo_medal"`
}

// LiveMedalWallItemMedalInfo contains the corresponding live medal response fields.
type LiveMedalWallItemMedalInfo struct {
	TargetID         int64  `json:"target_id"`
	Level            int    `json:"level"`
	MedalName        string `json:"medal_name"`
	MedalColorStart  int    `json:"medal_color_start"`
	MedalColorEnd    int    `json:"medal_color_end"`
	MedalColorBorder int    `json:"medal_color_border"`
	GuardLevel       int    `json:"guard_level"`
	WearingStatus    int    `json:"wearing_status"`
	MedalID          int    `json:"medal_id"`
	Intimacy         int    `json:"intimacy"`
	NextIntimacy     int    `json:"next_intimacy"`
	TodayFeed        int    `json:"today_feed"`
	DayLimit         int    `json:"day_limit"`
	GuardIcon        string `json:"guard_icon"`
	HonorIcon        string `json:"honor_icon"`
}

// LiveUinfoMedal 是直播勋章墙与勋章面板共用的用户勋章信息。
//
// 勋章墙的 LiveMedalWallItemUinfoMedal 与面板的
// LiveFansMedalPanelItemUinfoMedal 字段完全一致，合并到本类型并保留原名别名。
type LiveUinfoMedal struct {
	Name               string `json:"name"`
	Level              int    `json:"level"`
	ColorStart         int    `json:"color_start"`
	ColorEnd           int    `json:"color_end"`
	ColorBorder        int    `json:"color_border"`
	Color              int    `json:"color"`
	ID                 int    `json:"id"`
	Typ                int    `json:"typ"`
	IsLight            int    `json:"is_light"`
	Ruid               int64  `json:"ruid"`
	GuardLevel         int    `json:"guard_level"`
	Score              int    `json:"score"`
	GuardIcon          string `json:"guard_icon"`
	HonorIcon          string `json:"honor_icon"`
	V2MedalColorStart  string `json:"v2_medal_color_start"`
	V2MedalColorEnd    string `json:"v2_medal_color_end"`
	V2MedalColorBorder string `json:"v2_medal_color_border"`
	V2MedalColorText   string `json:"v2_medal_color_text"`
	V2MedalColorLevel  string `json:"v2_medal_color_level"`
	UserReceiveCount   int    `json:"user_receive_count"`
}

// LiveMedalWallItemUinfoMedal 是 LiveUinfoMedal 的别名。
type LiveMedalWallItemUinfoMedal = LiveUinfoMedal

// GetLiveActivatedMedalInfoResult contains the corresponding live medal response fields.
type GetLiveActivatedMedalInfoResult struct {
	Face                   string                   `json:"face"`
	Name                   string                   `json:"name"`
	MedalName              string                   `json:"medal_name"`
	FansMedalCount         int                      `json:"fans_medal_count"`
	Level                  int                      `json:"level"`
	IsLighted              bool                     `json:"is_lighted"`
	Intimacy               int                      `json:"intimacy"`
	NextIntimacy           int                      `json:"next_intimacy"`
	TaskLightDays          int                      `json:"task_light_days"`
	TaskInfo               []LiveActivatedMedalTask `json:"task_info"`
	FansClubGiftInfo       LiveActivatedMedalGift   `json:"fans_club_gift_info"`
	MedalColorBorder       string                   `json:"medal_color_border"`
	MedalColor             string                   `json:"medal_color"`
	MedalColorText         string                   `json:"medal_color_text"`
	MedalColorLevel        string                   `json:"medal_color_level"`
	GuardLevel             int                      `json:"guard_level"`
	LightSource            int                      `json:"light_source"`
	FreeIntimacy           int                      `json:"free_intimacy"`
	ReachFreeIntimacyLimit bool                     `json:"reach_free_intimacy_limit"`
}

// LiveActivatedMedalTask contains the corresponding live medal response fields.
type LiveActivatedMedalTask struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	AddText     string `json:"add_text"`
	JumpType    string `json:"jump_type"`
	IsDone      bool   `json:"is_done"`
	IconGuard   string `json:"icon_guard"`
	IconAdmiral string `json:"icon_admiral"`
	IconCaptain string `json:"icon_captain"`
}

// LiveActivatedMedalGift contains the corresponding live medal response fields.
type LiveActivatedMedalGift struct {
	GiftID           int `json:"gift_id"`
	Price            int `json:"price"`
	GiftDiscountInfo any `json:"gift_discount_info"`
}

// GetLiveFansMedalPanelResult contains the corresponding live medal response fields.
type GetLiveFansMedalPanelResult struct {
	List        []LiveFansMedalPanelItem   `json:"list"`
	SpecialList []LiveFansMedalPanelItem   `json:"special_list"`
	BottomBar   any                        `json:"bottom_bar"`
	PageInfo    LiveFansMedalPanelPageInfo `json:"page_info"`
	TotalNumber int                        `json:"total_number"`
	HasMedal    int                        `json:"has_medal"`
	GroupMedal  any                        `json:"group_medal"`
}

// LiveFansMedalPanelPageInfo 直播勋章面板分页信息。
//
// 字段集（当前页/下一页/总页数等）与其他业务分页差异较大，保留独立类型。
type LiveFansMedalPanelPageInfo struct {
	Number          int  `json:"number"`
	CurrentPage     int  `json:"current_page"` // 当前页码
	HasMore         bool `json:"has_more"`     // 始终为true,不能依靠它判断
	NextPage        int  `json:"next_page"`
	NextLightStatus int  `json:"next_light_status"`
	TotalPage       int  `json:"total_page"` // 总页数
}

// LiveFansMedalPanelItem contains the corresponding live medal response fields.
type LiveFansMedalPanelItem struct {
	Medal       LiveFansMedalPanelItemMedal      `json:"medal"`
	AnchorInfo  LiveFansMedalPanelItemAnchorInfo `json:"anchor_info"`
	Superscript any                              `json:"superscript"`
	RoomInfo    LiveFansMedalPanelItemRoomInfo   `json:"room_info"`
	UinfoMedal  LiveFansMedalPanelItemUinfoMedal `json:"uinfo_medal"`
}

// LiveFansMedalPanelItemMedal contains the corresponding live medal response fields.
type LiveFansMedalPanelItemMedal struct {
	UID                int    `json:"uid"`
	TargetID           int    `json:"target_id"`
	TargetName         string `json:"target_name"`
	MedalID            int    `json:"medal_id"`
	Level              int    `json:"level"`
	MedalName          string `json:"medal_name"`
	MedalColor         int    `json:"medal_color"`
	Intimacy           int    `json:"intimacy"`
	NextIntimacy       int    `json:"next_intimacy"`
	DayLimit           int    `json:"day_limit"`
	TodayFeed          int    `json:"today_feed"`
	MedalColorStart    int    `json:"medal_color_start"`
	MedalColorEnd      int    `json:"medal_color_end"`
	MedalColorBorder   int    `json:"medal_color_border"`
	IsLighted          int    `json:"is_lighted"`
	GuardLevel         int    `json:"guard_level"`
	WearingStatus      int    `json:"wearing_status"`
	MedalIconID        int    `json:"medal_icon_id"`
	MedalIconURL       string `json:"medal_icon_url"`
	GuardIcon          string `json:"guard_icon"`
	HonorIcon          string `json:"honor_icon"`
	CanDelete          bool   `json:"can_delete"`
	V2MedalColorStart  string `json:"v2_medal_color_start"`
	V2MedalColorEnd    string `json:"v2_medal_color_end"`
	V2MedalColorBorder string `json:"v2_medal_color_border"`
	V2MedalColorText   string `json:"v2_medal_color_text"`
	V2MedalColorLevel  string `json:"v2_medal_color_level"`
	DayLimitExtra      any    `json:"day_limit_extra"`
}

// LiveFansMedalPanelItemAnchorInfo contains the corresponding live medal response fields.
type LiveFansMedalPanelItemAnchorInfo struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Verify   int    `json:"verify"`
}

// LiveFansMedalPanelItemRoomInfo contains the corresponding live medal response fields.
type LiveFansMedalPanelItemRoomInfo struct {
	RoomID       int    `json:"room_id"`
	LivingStatus int    `json:"living_status"`
	URL          string `json:"url"`
}

// LiveFansMedalPanelItemUinfoMedal 是 LiveUinfoMedal 的别名。
type LiveFansMedalPanelItemUinfoMedal = LiveUinfoMedal

type Frame struct {
	Name       string `json:"name"`         // 名称
	Value      string `json:"value"`        // 值
	Position   int    `json:"position"`     // 位置
	Desc       string `json:"desc"`         // 描述
	Area       int    `json:"area"`         // 分区
	AreaOld    int    `json:"area_old"`     // 旧分区
	BgColor    string `json:"bg_color"`     // 背景色
	BgPic      string `json:"bg_pic"`       // 背景图
	UseOldArea bool   `json:"use_old_area"` // 是否旧分区号
}

type Badge struct {
	Name     string `json:"name"`     // 类型。v_person: 个人认证(黄) 。 v_company: 企业认证(蓝)
	Position int    `json:"position"` // 位置
	Value    string `json:"value"`    // 值
	Desc     string `json:"desc"`     // 描述
}

type NewPendants struct {
	Frame       *Frame `json:"frame"`        // 头像框
	MobileFrame *Frame `json:"mobile_frame"` // 同上。手机版, 结构一致, 可能null
	Badge       *Badge `json:"badge"`        // 大v
	MobileBadge *Badge `json:"mobile_badge"` // 同上。手机版, 结构一致, 可能null
}

type StudioInfo struct {
	Status     int   `json:"status"`
	MasterList []any `json:"master_list"`
}

type LiveRoomInfo struct {
	UID                  int         `json:"uid"`                // 主播mid
	RoomID               int         `json:"room_id"`            // 直播间长号
	ShortID              int         `json:"short_id"`           // 直播间短号。为0是无短号
	Attention            int         `json:"attention"`          // 关注数量
	Online               int         `json:"online"`             // 观看人数
	IsPortrait           bool        `json:"is_portrait"`        // 是否竖屏
	Description          string      `json:"description"`        // 描述
	LiveStatus           int         `json:"live_status"`        // 直播状态。0：未开播。1：直播中。2：轮播中
	AreaID               int         `json:"area_id"`            // 分区id
	ParentAreaID         int         `json:"parent_area_id"`     // 父分区id
	ParentAreaName       string      `json:"parent_area_name"`   // 父分区名称
	OldAreaID            int         `json:"old_area_id"`        // 旧版分区id
	Background           string      `json:"background"`         // 背景图片链接
	Title                string      `json:"title"`              // 标题
	UserCover            string      `json:"user_cover"`         // 封面
	Keyframe             string      `json:"keyframe"`           // 关键帧。用于网页端悬浮展示
	IsStrictRoom         bool        `json:"is_strict_room"`     // 未知。未知
	LiveTime             string      `json:"live_time"`          // 直播开始时间。YYYY-MM-DD HH:mm:ss
	Tags                 string      `json:"tags"`               // 标签。','分隔
	IsAnchor             int         `json:"is_anchor"`          // 未知。未知
	RoomSilentType       string      `json:"room_silent_type"`   // 禁言状态
	RoomSilentLevel      int         `json:"room_silent_level"`  // 禁言等级
	RoomSilentSecond     int         `json:"room_silent_second"` // 禁言时间。单位是秒
	AreaName             string      `json:"area_name"`          // 分区名称
	Pardants             string      `json:"pardants"`           // 未知。未知
	AreaPardants         string      `json:"area_pardants"`      // 未知。未知
	HotWords             []string    `json:"hot_words"`          // 热词
	HotWordsStatus       int         `json:"hot_words_status"`   // 热词状态
	Verify               string      `json:"verify"`             // 未知。未知
	NewPendants          NewPendants `json:"new_pendants"`       // 头像框\大v
	UpSession            string      `json:"up_session"`         // 未知
	PkStatus             int         `json:"pk_status"`          // pk状态
	PkID                 int         `json:"pk_id"`              // pk id
	BattleID             int         `json:"battle_id"`          // 未知
	AllowChangeAreaTime  int         `json:"allow_change_area_time"`
	AllowUploadCoverTime int         `json:"allow_upload_cover_time"`
	StudioInfo           StudioInfo  `json:"studio_info"`
}

type UpdateLiveRoomTitleResult struct {
	SubSessionKey string `json:"sub_session_key"` // 信息变动标识
	AuditInfo     any    `json:"audit_info"`      // 标题审核信息（不一定有值，因此在这里不进行解析）
}

type Rtmp struct {
	Addr     string `json:"addr"`     // RTMP推流（发送）地址。**重要**
	Code     string `json:"code"`     // RTMP推流参数（密钥）。**重要**
	NewLink  string `json:"new_link"` // 获取CDN推流ip地址重定向信息的url。没啥用
	Provider string `json:"provider"` // ？？？。作用尚不明确
}

type Protocol struct {
	Protocol string `json:"protocol"` // rtmp。作用尚不明确
	Addr     string `json:"addr"`     // RTMP推流（发送）地址
	Code     string `json:"code"`     // RTMP推流参数（密钥）
	NewLink  string `json:"new_link"` // 获取CDN推流ip地址重定向信息的url
	Provider string `json:"provider"` // txy。作用尚不明确
}

type StartLiveResult struct {
	Change    int        `json:"change"`    // 是否改变状态。0：未改变。1：改变
	Status    string     `json:"status"`    // LIVE
	RoomType  int        `json:"room_type"` // 0。作用尚不明确
	Rtmp      Rtmp       `json:"rtmp"`      // RTMP推流地址信息
	Protocols []Protocol `json:"protocols"` // ？？？。作用尚不明确
	TryTime   string     `json:"try_time"`  // ？？？。作用尚不明确
	LiveKey   string     `json:"live_key"`  // ？？？。作用尚不明确
	Notice    Notice     `json:"notice"`    // ？？？。作用尚不明确
}

type StopLiveResult struct {
	Change int    `json:"change"` // 是否改变状态。0：未改变。1：改变
	Status string `json:"status"` // PREPARING
}

type SubLiveArea struct {
	ID         string `json:"id"`          // 子分区id
	ParentID   string `json:"parent_id"`   // 父分区id
	OldAreaID  string `json:"old_area_id"` // 旧分区id
	Name       string `json:"name"`        // 子分区名
	ActID      string `json:"act_id"`      // 0。**作用尚不明确**
	PkStatus   string `json:"pk_status"`   // ？？？。**作用尚不明确**
	HotStatus  int    `json:"hot_status"`  // 是否为热门分区。0：否。1：是
	LockStatus string `json:"lock_status"` // 0。**作用尚不明确**
	Pic        string `json:"pic"`         // 子分区标志图片url
	ParentName string `json:"parent_name"` // 父分区名
	AreaType   int    `json:"area_type"`
}

type LiveAreaList struct {
	ID   int           `json:"id"`   // 父分区id
	Name string        `json:"name"` // 父分区名
	List []SubLiveArea `json:"list"` // 子分区列表
}

type HomePageLiveVersion struct {
	CurrVersion      string `json:"curr_version,omitempty" request:"query,omitempty"`      // 直播姬最新版本号
	Build            int    `json:"build,omitempty" request:"query,omitempty"`             // 直播姬构建号
	Instruction      string `json:"instruction,omitempty" request:"query,omitempty"`       // 更新说明（简要）
	FileSize         string `json:"file_size,omitempty" request:"query,omitempty"`         // 文件大小（字节）
	FileMd5          string `json:"file_md5,omitempty" request:"query,omitempty"`          // 安装包文件MD5
	Content          string `json:"content,omitempty" request:"query,omitempty"`           // HTML格式的更新内容
	DownloadURL      string `json:"download_url,omitempty" request:"query,omitempty"`      // 安装包下载链接
	HdiffpatchSwitch int    `json:"hdiffpatch_switch,omitempty" request:"query,omitempty"` // 增量更新开关?
}
