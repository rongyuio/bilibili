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

// 天选时刻相关响应模型。

// LiveWebAreaList 是 Web 端直播间分区列表的响应（官方在 data 内又嵌套了一层 data）。
type LiveWebAreaList struct {
	Data []LiveWebArea `json:"data"` // 一级分区列表
}

// LiveWebArea 是 Web 端直播间分区列表中的一级分区。
type LiveWebArea struct {
	ID   int    `json:"id"`   // 分区 id
	Name string `json:"name"` // 分区名称
}

// LiveAreaRoomList 是直播二级分区房间列表的响应。
type LiveAreaRoomList struct {
	NewTags []LiveAreaRoomListTag  `json:"new_tags"` // 排序方式列表。翻页时一般把第一项的 SortType 作为下一页的排序方式
	List    []LiveAreaRoomListItem `json:"list"`     // 房间列表
	HasMore int                    `json:"has_more"` // 是否还有更多。1-有
}

// LiveAreaRoomListTag 是直播二级分区房间列表的排序方式。
type LiveAreaRoomListTag struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SortType string `json:"sort_type"`
}

// LiveAreaRoomListItem 是直播二级分区房间列表中的房间。
type LiveAreaRoomListItem struct {
	RoomID      int64                          `json:"roomid"`       // 直播间号。注意字段名没有下划线
	UID         int64                          `json:"uid"`          // 主播 UID
	Title       string                         `json:"title"`        // 直播间标题
	Uname       string                         `json:"uname"`        // 主播昵称
	ParentID    int                            `json:"parent_id"`    // 父分区 id
	ParentName  string                         `json:"parent_name"`  // 父分区名称
	AreaID      int                            `json:"area_id"`      // 子分区 id
	AreaName    string                         `json:"area_name"`    // 子分区名称
	PendantInfo map[string]LiveAreaRoomPendant `json:"pendant_info"` // 挂件信息。键为 source_id（如 "2"）
}

// LiveAreaRoomPendant 是直播二级分区房间列表中的挂件信息。官方字段拼写为 pendent（无 i）。
// PendentID 为 504 表示天选时刻房间。
type LiveAreaRoomPendant struct {
	PendentID int64 `json:"pendent_id"` // 挂件 id
	Content   any   `json:"content"`    // 挂件内容
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

// LiveWebRoomList 是 Web 端直播间列表（新接口）的响应。
type LiveWebRoomList struct {
	RoomList          []LiveWebRoomModule `json:"room_list"`           // 房间模块列表
	RecommendRoomList []LiveWebRoom       `json:"recommend_room_list"` // 推荐房间列表
}

// LiveWebRoomModule 是直播间列表中的一个模块。
type LiveWebRoomModule struct {
	ModuleInfo struct {
		ID    int    `json:"id"`    // 模块 id
		Title string `json:"title"` // 模块标题
		Count int    `json:"count"` // 房间总数
	} `json:"module_info"`
	List []LiveWebRoom `json:"list"` // 房间列表
}

// LiveWebRoom 是直播间列表中的房间。
type LiveWebRoom struct {
	RoomID     int64  `json:"roomid"`              // 直播间号。注意字段名没有下划线
	UID        int64  `json:"uid"`                 // 主播 UID
	Uname      string `json:"uname"`               // 主播昵称
	Title      string `json:"title"`               // 直播间标题
	AreaV2Name string `json:"area_v2_name"`        // 子分区名
	ParentName string `json:"area_v2_parent_name"` // 父分区名
	Online     int    `json:"online"`              // 人气值
	Cover      string `json:"cover"`               // 封面
}
