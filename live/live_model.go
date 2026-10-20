package live

// 直播勋章、分区、天选时刻、直播间列表与弹幕相关响应模型。
//
// 每个接口的返回体独立定义：即使字段与其它接口完全相同，也不共享类型，
// 便于后续单个接口的返回体变更不影响其它调用。

// GetLiveMedalWallResult 是用户直播勋章墙的 data 对象。
type GetLiveMedalWallResult struct {
	List            []GetLiveMedalWallItem `json:"list"`
	Count           int                    `json:"count"`
	CloseSpaceMedal int                    `json:"close_space_medal"`
	OnlyShowWearing int                    `json:"only_show_wearing"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
	UID             int                    `json:"uid"`
	Level           int                    `json:"level"`
}

// GetLiveMedalWallItem 是勋章墙中的一枚勋章。
type GetLiveMedalWallItem struct {
	MedalInfo  GetLiveMedalWallItemMedalInfo  `json:"medal_info"`
	TargetName string                         `json:"target_name"`
	TargetIcon string                         `json:"target_icon"`
	Link       string                         `json:"link"`
	LiveStatus int                            `json:"live_status"`
	Official   int                            `json:"official"`
	UinfoMedal GetLiveMedalWallItemUinfoMedal `json:"uinfo_medal"`
}

// GetLiveMedalWallItemMedalInfo 是勋章墙中勋章的牌面信息。
type GetLiveMedalWallItemMedalInfo struct {
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

// GetLiveMedalWallItemUinfoMedal 是勋章墙中的用户勋章信息。
type GetLiveMedalWallItemUinfoMedal struct {
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

// GetLiveActivatedMedalInfoResult 是已激活勋章及任务信息的 data 对象。
type GetLiveActivatedMedalInfoResult struct {
	Face                   string                          `json:"face"`
	Name                   string                          `json:"name"`
	MedalName              string                          `json:"medal_name"`
	FansMedalCount         int                             `json:"fans_medal_count"`
	Level                  int                             `json:"level"`
	IsLighted              bool                            `json:"is_lighted"`
	Intimacy               int                             `json:"intimacy"`
	NextIntimacy           int                             `json:"next_intimacy"`
	TaskLightDays          int                             `json:"task_light_days"`
	TaskInfo               []GetLiveActivatedMedalInfoTask `json:"task_info"`
	FansClubGiftInfo       GetLiveActivatedMedalInfoGift   `json:"fans_club_gift_info"`
	MedalColorBorder       string                          `json:"medal_color_border"`
	MedalColor             string                          `json:"medal_color"`
	MedalColorText         string                          `json:"medal_color_text"`
	MedalColorLevel        string                          `json:"medal_color_level"`
	GuardLevel             int                             `json:"guard_level"`
	LightSource            int                             `json:"light_source"`
	FreeIntimacy           int                             `json:"free_intimacy"`
	ReachFreeIntimacyLimit bool                            `json:"reach_free_intimacy_limit"`
}

// GetLiveActivatedMedalInfoTask 是已激活勋章的任务信息。
type GetLiveActivatedMedalInfoTask struct {
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

// GetLiveActivatedMedalInfoGift 是已激活勋章的粉丝团礼物信息。
type GetLiveActivatedMedalInfoGift struct {
	GiftID           int `json:"gift_id"`
	Price            int `json:"price"`
	GiftDiscountInfo any `json:"gift_discount_info"`
}

// GetLiveFansMedalPanelResult 是当前账号直播勋章面板的 data 对象。
type GetLiveFansMedalPanelResult struct {
	List        []GetLiveFansMedalPanelItem   `json:"list"`
	SpecialList []GetLiveFansMedalPanelItem   `json:"special_list"`
	BottomBar   any                           `json:"bottom_bar"`
	PageInfo    GetLiveFansMedalPanelPageInfo `json:"page_info"`
	TotalNumber int                           `json:"total_number"`
	HasMedal    int                           `json:"has_medal"`
	GroupMedal  any                           `json:"group_medal"`
}

// GetLiveFansMedalPanelPageInfo 是直播勋章面板的分页信息。
//
// 字段集（当前页/下一页/总页数等）与其他业务分页差异较大，保留独立类型。
type GetLiveFansMedalPanelPageInfo struct {
	Number          int  `json:"number"`
	CurrentPage     int  `json:"current_page"` // 当前页码
	HasMore         bool `json:"has_more"`     // 始终为true,不能依靠它判断
	NextPage        int  `json:"next_page"`
	NextLightStatus int  `json:"next_light_status"`
	TotalPage       int  `json:"total_page"` // 总页数
}

// GetLiveFansMedalPanelItem 是勋章面板中的一枚勋章。
type GetLiveFansMedalPanelItem struct {
	Medal       GetLiveFansMedalPanelItemMedal      `json:"medal"`
	AnchorInfo  GetLiveFansMedalPanelItemAnchorInfo `json:"anchor_info"`
	Superscript any                                 `json:"superscript"`
	RoomInfo    GetLiveFansMedalPanelItemRoomInfo   `json:"room_info"`
	UinfoMedal  GetLiveFansMedalPanelItemUinfoMedal `json:"uinfo_medal"`
}

// GetLiveFansMedalPanelItemMedal 是勋章面板中勋章的牌面信息。
type GetLiveFansMedalPanelItemMedal struct {
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

// GetLiveFansMedalPanelItemAnchorInfo 是勋章面板中的主播信息。
type GetLiveFansMedalPanelItemAnchorInfo struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Verify   int    `json:"verify"`
}

// GetLiveFansMedalPanelItemRoomInfo 是勋章面板中的直播间信息。
type GetLiveFansMedalPanelItemRoomInfo struct {
	RoomID       int    `json:"room_id"`
	LivingStatus int    `json:"living_status"`
	URL          string `json:"url"`
}

// GetLiveFansMedalPanelItemUinfoMedal 是勋章面板中的用户勋章信息。
type GetLiveFansMedalPanelItemUinfoMedal struct {
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

// GetLiveWebAreaListResult 是 Web 端直播间分区列表的响应（官方在 data 内又嵌套了一层 data）。
type GetLiveWebAreaListResult struct {
	Data []GetLiveWebAreaListItem `json:"data"` // 一级分区列表
}

// GetLiveWebAreaListItem 是 Web 端直播间分区列表中的一级分区。
type GetLiveWebAreaListItem struct {
	ID   int    `json:"id"`   // 分区 id
	Name string `json:"name"` // 分区名称
}

// CheckLiveAnchorLotteryResult 是直播间天选时刻的查询结果。
type CheckLiveAnchorLotteryResult struct {
	ID             int64  `json:"id"`               // 天选抽奖 id，参与时原样传入 JoinLiveAnchorLottery
	RoomID         int64  `json:"room_id"`          // 直播间号
	Status         int    `json:"status"`           // 状态。1-进行中，2-已结束
	AwardName      string `json:"award_name"`       // 奖品名称
	AwardNum       int    `json:"award_num"`        // 奖品数量
	Danmu          string `json:"danmu"`            // 参与弹幕口令
	JoinType       int    `json:"join_type"`        // 参与方式
	RequireType    int    `json:"require_type"`     // 参与条件类型。0-无条件，1-关注主播，2-粉丝勋章等级，3-提督/舰长
	RequireValue   int    `json:"require_value"`    // 参与条件数值
	RequireText    string `json:"require_text"`     // 参与条件文案
	GiftID         int64  `json:"gift_id"`          // 需要赠送的礼物 id
	GiftName       string `json:"gift_name"`        // 礼物名称
	GiftNum        int    `json:"gift_num"`         // 礼物数量
	GiftPrice      int    `json:"gift_price"`       // 礼物单价。大于 0 表示参与需要付费赠礼
	CurGiftNum     int    `json:"cur_gift_num"`     // 当前已送数量
	SendGiftEnsure int    `json:"send_gift_ensure"` // 是否确认送出礼物
}

// JoinLiveAnchorLotteryResult 是参与天选时刻抽奖的响应。
type JoinLiveAnchorLotteryResult struct {
	DiscountID int64  `json:"discount_id"`  // (?)。作用尚不明确
	Gold       int64  `json:"gold"`         // 金瓜子余额
	Silver     int64  `json:"silver"`       // 银瓜子余额
	CurGiftNum int64  `json:"cur_gift_num"` // 本次消耗的礼物数量
	GoodsID    int64  `json:"goods_id"`     // (?)。作用尚不明确
	NewOrderID string `json:"new_order_id"` // (?)。订单号
}

// GetLiveWebRoomListResult 是 Web 端直播间列表（新接口）的响应。
type GetLiveWebRoomListResult struct {
	RoomList          []GetLiveWebRoomListModule `json:"room_list"`           // 房间模块列表
	RecommendRoomList []GetLiveWebRoomListItem   `json:"recommend_room_list"` // 推荐房间列表
}

// GetLiveWebRoomListModule 是直播间列表中的一个模块。
type GetLiveWebRoomListModule struct {
	ModuleInfo GetLiveWebRoomListModuleInfo `json:"module_info"`
	List       []GetLiveWebRoomListItem     `json:"list"` // 房间列表
}

// GetLiveWebRoomListModuleInfo 是直播间列表模块的信息。
type GetLiveWebRoomListModuleInfo struct {
	ID    int    `json:"id"`    // 模块 id
	Title string `json:"title"` // 模块标题
	Count int    `json:"count"` // 房间总数
}

// GetLiveWebRoomListItem 是直播间列表中的房间。
type GetLiveWebRoomListItem struct {
	RoomID     int64  `json:"roomid"`              // 直播间号。注意字段名没有下划线
	UID        int64  `json:"uid"`                 // 主播 UID
	Uname      string `json:"uname"`               // 主播昵称
	Title      string `json:"title"`               // 直播间标题
	AreaV2Name string `json:"area_v2_name"`        // 子分区名
	ParentName string `json:"area_v2_parent_name"` // 父分区名
	Online     int    `json:"online"`              // 人气值
	Cover      string `json:"cover"`               // 封面
}

// GetLiveHotRankListResult 是首页人气榜直播间列表的响应。
type GetLiveHotRankListResult struct {
	List []GetLiveHotRankListItem `json:"list"` // 榜单房间列表
}

// GetLiveHotRankListItem 是人气榜中的一个直播间。
type GetLiveHotRankListItem struct {
	RoomID          int64  `json:"roomid"`              // 直播间号。注意字段名没有下划线
	UID             int64  `json:"uid"`                 // 主播 UID
	Uname           string `json:"uname"`               // 主播昵称
	Face            string `json:"face"`                // 主播头像
	Title           string `json:"title"`               // 直播间标题
	AreaV2ID        int    `json:"area_v2_id"`          // 子分区 id
	AreaV2Name      string `json:"area_v2_name"`        // 子分区名
	ParentID        int    `json:"area_v2_parent_id"`   // 父分区 id
	ParentName      string `json:"area_v2_parent_name"` // 父分区名
	Online          int    `json:"online"`              // 人气值
	OfficialVerify  int    `json:"official_verify"`     // 官方认证。-1：无
	LotStatus       int    `json:"lot_status"`          // 抽奖活动状态。0：无
	RedPocketStatus int    `json:"red_pocket_status"`   // 红包状态。0：无
	UserNum         int    `json:"user_num"`            // 助力人数
	UserNumText     string `json:"user_num_text"`       // 助力人数文案
	Score           int    `json:"score"`               // 榜单分数
	ScoreText       string `json:"score_text"`          // 榜单分数文案
	LiveStatus      int    `json:"live_status"`         // 开播状态。1：直播中
	Rank            int    `json:"rank"`                // 榜单排名
}

// SendLiveDanmakuResult 是发送直播弹幕的响应（data 部分）。
type SendLiveDanmakuResult struct {
	ModeInfo GetSendLiveDanmakuModeInfo `json:"mode_info"` // 弹幕模式信息
}

// GetSendLiveDanmakuModeInfo 是弹幕下发后的模式信息。
type GetSendLiveDanmakuModeInfo struct {
	Mode           int    `json:"mode"`             // 实际生效的弹幕模式
	ShowPlayerType int    `json:"show_player_type"` // 播放器展示类型
	Extra          string `json:"extra"`            // 附加信息
}
