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
