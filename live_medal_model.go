package bilibili

// Models preserve the fields declared by local tools; upstream types have not been verified live.

// GetLiveMedalWallResult contains the corresponding live medal response fields.
type GetLiveMedalWallResult struct {
	List            []LiveMedalWallItem `json:"list"`
	Count           int                 `json:"count"`
	CloseSpaceMedal int                 `json:"close_space_medal"`
	OnlyShowWearing int                 `json:"only_show_wearing"`
	Name            string              `json:"name"`
	Icon            string              `json:"icon"`
	Uid             int                 `json:"uid"`
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
	TargetId         int64  `json:"target_id"`
	Level            int    `json:"level"`
	MedalName        string `json:"medal_name"`
	MedalColorStart  int    `json:"medal_color_start"`
	MedalColorEnd    int    `json:"medal_color_end"`
	MedalColorBorder int    `json:"medal_color_border"`
	GuardLevel       int    `json:"guard_level"`
	WearingStatus    int    `json:"wearing_status"`
	MedalId          int    `json:"medal_id"`
	Intimacy         int    `json:"intimacy"`
	NextIntimacy     int    `json:"next_intimacy"`
	TodayFeed        int    `json:"today_feed"`
	DayLimit         int    `json:"day_limit"`
	GuardIcon        string `json:"guard_icon"`
	HonorIcon        string `json:"honor_icon"`
}

// LiveMedalWallItemUinfoMedal contains the corresponding live medal response fields.
type LiveMedalWallItemUinfoMedal struct {
	Name               string `json:"name"`
	Level              int    `json:"level"`
	ColorStart         int    `json:"color_start"`
	ColorEnd           int    `json:"color_end"`
	ColorBorder        int    `json:"color_border"`
	Color              int    `json:"color"`
	Id                 int    `json:"id"`
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
	GiftId           int         `json:"gift_id"`
	Price            int         `json:"price"`
	GiftDiscountInfo interface{} `json:"gift_discount_info"`
}

// GetLiveFansMedalPanelResult contains the corresponding live medal response fields.
type GetLiveFansMedalPanelResult struct {
	List        []LiveFansMedalPanelItem   `json:"list"`
	SpecialList []LiveFansMedalPanelItem   `json:"special_list"`
	BottomBar   interface{}                `json:"bottom_bar"`
	PageInfo    LiveFansMedalPanelPageInfo `json:"page_info"`
	TotalNumber int                        `json:"total_number"`
	HasMedal    int                        `json:"has_medal"`
	GroupMedal  interface{}                `json:"group_medal"`
}

// LiveFansMedalPanelPageInfo contains the corresponding live medal response fields.
type LiveFansMedalPanelPageInfo struct {
	Number          int  `json:"number"`
	CurrentPage     int  `json:"current_page"`
	HasMore         bool `json:"has_more"`
	NextPage        int  `json:"next_page"`
	NextLightStatus int  `json:"next_light_status"`
	TotalPage       int  `json:"total_page"`
}

// LiveFansMedalPanelItem contains the corresponding live medal response fields.
type LiveFansMedalPanelItem struct {
	Medal       LiveFansMedalPanelItemMedal      `json:"medal"`
	AnchorInfo  LiveFansMedalPanelItemAnchorInfo `json:"anchor_info"`
	Superscript interface{}                      `json:"superscript"`
	RoomInfo    LiveFansMedalPanelItemRoomInfo   `json:"room_info"`
	UinfoMedal  LiveFansMedalPanelItemUinfoMedal `json:"uinfo_medal"`
}

// LiveFansMedalPanelItemMedal contains the corresponding live medal response fields.
type LiveFansMedalPanelItemMedal struct {
	Uid                int         `json:"uid"`
	TargetId           int         `json:"target_id"`
	TargetName         string      `json:"target_name"`
	MedalId            int         `json:"medal_id"`
	Level              int         `json:"level"`
	MedalName          string      `json:"medal_name"`
	MedalColor         int         `json:"medal_color"`
	Intimacy           int         `json:"intimacy"`
	NextIntimacy       int         `json:"next_intimacy"`
	DayLimit           int         `json:"day_limit"`
	TodayFeed          int         `json:"today_feed"`
	MedalColorStart    int         `json:"medal_color_start"`
	MedalColorEnd      int         `json:"medal_color_end"`
	MedalColorBorder   int         `json:"medal_color_border"`
	IsLighted          int         `json:"is_lighted"`
	GuardLevel         int         `json:"guard_level"`
	WearingStatus      int         `json:"wearing_status"`
	MedalIconId        int         `json:"medal_icon_id"`
	MedalIconUrl       string      `json:"medal_icon_url"`
	GuardIcon          string      `json:"guard_icon"`
	HonorIcon          string      `json:"honor_icon"`
	CanDelete          bool        `json:"can_delete"`
	V2MedalColorStart  string      `json:"v2_medal_color_start"`
	V2MedalColorEnd    string      `json:"v2_medal_color_end"`
	V2MedalColorBorder string      `json:"v2_medal_color_border"`
	V2MedalColorText   string      `json:"v2_medal_color_text"`
	V2MedalColorLevel  string      `json:"v2_medal_color_level"`
	DayLimitExtra      interface{} `json:"day_limit_extra"`
}

// LiveFansMedalPanelItemAnchorInfo contains the corresponding live medal response fields.
type LiveFansMedalPanelItemAnchorInfo struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Verify   int    `json:"verify"`
}

// LiveFansMedalPanelItemRoomInfo contains the corresponding live medal response fields.
type LiveFansMedalPanelItemRoomInfo struct {
	RoomId       int    `json:"room_id"`
	LivingStatus int    `json:"living_status"`
	Url          string `json:"url"`
}

// LiveFansMedalPanelItemUinfoMedal contains the corresponding live medal response fields.
type LiveFansMedalPanelItemUinfoMedal struct {
	Name               string `json:"name"`
	Level              int    `json:"level"`
	ColorStart         int    `json:"color_start"`
	ColorEnd           int    `json:"color_end"`
	ColorBorder        int    `json:"color_border"`
	Color              int    `json:"color"`
	Id                 int    `json:"id"`
	Typ                int    `json:"typ"`
	IsLight            int    `json:"is_light"`
	Ruid               int    `json:"ruid"`
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
