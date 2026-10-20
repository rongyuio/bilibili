package live

// 直播间用户相关响应模型（live/user.md）。信封 code/message/ttl 由库统一处理。
//
// 每个接口的返回体独立定义，嵌套对象均带方法前缀，不与其它接口共享类型。

// GetMyLiveMedalsResult 是当前账号持有的粉丝勋章列表的响应。
type GetMyLiveMedalsResult struct {
	Count    int                     `json:"count"`     // 勋章数量
	Items    []GetMyLiveMedalsItem   `json:"items"`     // 粉丝勋章信息本体
	PageInfo GetMyLiveMedalsPageInfo `json:"page_info"` // 页码信息
}

// GetMyLiveMedalsItem 是一枚粉丝勋章。
type GetMyLiveMedalsItem struct {
	CanDeleted       bool   `json:"can_deleted"`        // 可否删除
	DayLimit         int    `json:"day_limit"`          // 日经验上限（原力值）
	GuardLevel       int    `json:"guard_level"`        // 大航海类型
	GuardMedalTitle  string `json:"guard_medal_title"`  // 加成状态
	Intimacy         int    `json:"intimacy"`           // 当前已得亲密度
	IsLighted        int    `json:"is_lighted"`         // 是否点亮。0：未点亮，1：点亮
	Level            int    `json:"level"`              // 勋章等级
	MedalName        string `json:"medal_name"`         // 勋章名
	MedalColorBorder int    `json:"medal_color_border"` // 勋章边框颜色，十进制表示的 16 进制值
	MedalColorStart  int    `json:"medal_color_start"`  // 勋章起始颜色，从右往左渐变（20 级+勋章）
	MedalColorEnd    int    `json:"medal_color_end"`    // 勋章结束颜色，从右往左渐变（20 级+勋章）
	MedalID          int    `json:"medal_id"`           // 粉丝勋章 id
	NextIntimacy     int    `json:"next_intimacy"`      // 升级所需经验
	TodayFeed        int    `json:"today_feed"`         // 本日亲密度
	RoomID           int64  `json:"roomid"`             // 直播间房间号。注意字段名没有下划线
	Status           int    `json:"status"`             // 状态
	TargetID         int64  `json:"target_id"`          // up 主 mid
	TargetName       string `json:"target_name"`        // up 主用户名
	Uname            string `json:"uname"`              // up 主用户名
}

// GetMyLiveMedalsPageInfo 是粉丝勋章列表的页码信息。
type GetMyLiveMedalsPageInfo struct {
	TotalPage int `json:"total_page"` // 页码总长度
	CurPage   int `json:"cur_page"`   // 当前返回的页码
}

// GetLiveSignInfoResult 是本月直播签到信息的响应。
type GetLiveSignInfoResult struct {
	Text              string `json:"text"`              // 今日签到奖励信息，默认为空
	SpecialText       string `json:"specialText"`       // 连续签到奖励信息，默认为空
	Status            int    `json:"status"`            // 签到状态。0：未签到，1：已签到
	AllDays           int    `json:"allDays"`           // 当月天数
	CurMonth          int    `json:"curMonth"`          // 当前月
	CurYear           int    `json:"curYear"`           // 当前年
	CurDay            int    `json:"curDay"`            // 当前日
	CurDate           string `json:"curDate"`           // 当前日期（格式化），例如 2023-2-19
	HadSignDays       int    `json:"hadSignDays"`       // 当月已签到天数，默认为 0
	NewTask           int    `json:"newTask"`           // 作用未知，默认为 0
	SignDaysList      []int  `json:"signDaysList"`      // 当月已签到日列表，默认为空
	SignBonusDaysList []int  `json:"signBonusDaysList"` // 当月已签到且有特殊奖励日列表，默认为空
}

// GetLastMonthSignDaysResult 是上月直播签到信息的响应。
type GetLastMonthSignDaysResult struct {
	Days              int   `json:"days"`              // 上月天数
	Month             int   `json:"month"`             // 上月月份值
	HadSignDays       int   `json:"hadSignDays"`       // 上月已签到天数，默认为 0
	SignDaysList      []int `json:"signDaysList"`      // 上月已签到日列表，默认为空
	SignBonusDaysList []int `json:"signBonusDaysList"` // 上月已签到且有特殊奖励日列表，默认为空
}

// GetOnlineGoldRankResult 是直播间贡献榜的响应。
type GetOnlineGoldRankResult struct {
	OnlineNum      int                      `json:"onlineNum"`      // 在线观众数量
	OnlineRankItem []GetOnlineGoldRankItem  `json:"OnlineRankItem"` // 贡献榜结果
	OwnInfo        GetOnlineGoldRankOwnInfo `json:"ownInfo"`        // 当前用户自己的信息
	TipsText       string                   `json:"tips_text"`      // 提示文字
	ValueText      string                   `json:"value_text"`     // 贡献值文案
	AB             GetOnlineGoldRankAB      `json:"ab"`             // AB 测试
	OnlineNumText  string                   `json:"onlineNumText"`  // 在线观众数量（字符串）
}

// GetOnlineGoldRankItem 是贡献榜中的一名用户。
type GetOnlineGoldRankItem struct {
	UserRank    int                     `json:"userRank"`     // 贡献值排名
	UID         int64                   `json:"uid"`          // 用户 id
	Name        string                  `json:"name"`         // 用户名字
	Face        string                  `json:"face"`         // 用户头像
	Score       int64                   `json:"score"`        // 贡献值
	MedalInfo   *GetOnlineGoldRankMedal `json:"medalInfo"`    // 粉丝牌对象，可能为 null
	GuardLevel  int                     `json:"guard_level"`  // 大航海类型
	WealthLevel int                     `json:"wealth_level"` // 荣耀等级
	IsMystery   bool                    `json:"is_mystery"`   // 是否为神秘人
	Uinfo       *GetOnlineGoldRankUinfo `json:"uinfo"`        // 用户详细信息，可能为 null
}

// GetOnlineGoldRankMedal 是贡献榜用户的粉丝牌信息。
type GetOnlineGoldRankMedal struct {
	GuardLevel       int    `json:"guardLevel"`       // 大航海类型
	MedalColorStart  int    `json:"medalColorStart"`  // 粉丝牌渐变起始色
	MedalColorEnd    int    `json:"medalColorEnd"`    // 粉丝牌渐变结束色
	MedalColorBorder int    `json:"medalColorBorder"` // 粉丝牌边框颜色
	MedalName        string `json:"medalName"`        // 粉丝牌名字
	Level            int    `json:"level"`            // 粉丝牌等级
	TargetID         int64  `json:"targetId"`         // 主播 uid
	IsLight          int    `json:"isLight"`          // 是否点亮。0：未点亮，1：点亮
}

// GetOnlineGoldRankUinfo 是贡献榜用户的详细信息。
type GetOnlineGoldRankUinfo struct {
	UID       int64                        `json:"uid"`        // 用户 uid
	Base      GetOnlineGoldRankBase        `json:"base"`       // 用户基本信息
	Medal     *GetOnlineGoldRankUinfoMedal `json:"medal"`      // 用户粉丝牌，可能为 null
	Face      string                       `json:"face"`       // 用户头像
	Guard     *GetOnlineGoldRankGuard      `json:"guard"`      // 大航海信息，可能为 null
	Wealth    *GetOnlineGoldRankWealth     `json:"wealth"`     // 荣耀等级
	Anon      any                          `json:"anon"`       // 匿名信息。实测为 null，作用尚不明确
	BubbleBox any                          `json:"bubble_box"` // 气泡框。实测为 null，作用尚不明确
	DmConfig  any                          `json:"dm_config"`  // 弹幕配置。实测为 null，作用尚不明确
}

// GetOnlineGoldRankBase 是贡献榜用户的基本信息。
type GetOnlineGoldRankBase struct {
	Name         string                        `json:"name"`           // 用户名
	Face         string                        `json:"face"`           // 用户头像
	NameColor    int                           `json:"name_color"`     // 昵称颜色
	IsMystery    bool                          `json:"is_mystery"`     // 是否为神秘人
	RiskCtrlInfo any                           `json:"risk_ctrl_info"` // 风控信息，可能为 null
	OriginInfo   GetOnlineGoldRankOriginInfo   `json:"origin_info"`    // 原始用户信息
	OfficialInfo GetOnlineGoldRankOfficialInfo `json:"official_info"`  // 官方认证信息
	NameColorStr string                        `json:"name_color_str"` // 昵称颜色字符串
}

// GetOnlineGoldRankOriginInfo 是贡献榜用户未隐藏身份时的原始信息。
type GetOnlineGoldRankOriginInfo struct {
	Name string `json:"name"` // 用户名
	Face string `json:"face"` // 用户头像
}

// GetOnlineGoldRankOfficialInfo 是贡献榜用户的官方认证信息。
type GetOnlineGoldRankOfficialInfo struct {
	Role  int    `json:"role"`  // 认证角色
	Title string `json:"title"` // 认证标题
	Desc  string `json:"desc"`  // 认证描述
	Type  int    `json:"type"`  // 认证类型。-1：无
}

// GetOnlineGoldRankUinfoMedal 是贡献榜用户的用户勋章信息。
type GetOnlineGoldRankUinfoMedal struct {
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

// GetOnlineGoldRankGuard 是贡献榜用户的大航海信息。
type GetOnlineGoldRankGuard struct {
	Level      int    `json:"level"`       // 大航海类型。1、2、3 分别为总督、提督、舰长
	ExpiredStr string `json:"expired_str"` // 到期时间文案
}

// GetOnlineGoldRankWealth 是贡献榜用户的荣耀等级。
type GetOnlineGoldRankWealth struct {
	Level     int    `json:"level"`       // 荣耀等级
	DmIconKey string `json:"dm_icon_key"` // 荣耀等级图标 key，实测为空
}

// GetOnlineGoldRankOwnInfo 是当前用户在贡献榜中的信息。
type GetOnlineGoldRankOwnInfo struct {
	UID         int64                   `json:"uid"`          // 用户 id
	Name        string                  `json:"name"`         // 用户名
	Face        string                  `json:"face"`         // 用户头像
	Rank        int                     `json:"rank"`         // 排名。未上榜时为 -1
	NeedScore   int64                   `json:"needScore"`    // 还差多少贡献值上榜
	Score       int64                   `json:"score"`        // 贡献值
	GuardLevel  int                     `json:"guard_level"`  // 大航海类型
	WealthLevel int                     `json:"wealth_level"` // 荣耀等级
	IsMystery   bool                    `json:"is_mystery"`   // 是否为神秘人
	Uinfo       *GetOnlineGoldRankUinfo `json:"uinfo"`        // 用户详细信息，可能为 null
}

// GetOnlineGoldRankAB 是贡献榜的 AB 测试信息。
type GetOnlineGoldRankAB struct {
	GuardAccompanyList int `json:"guard_accompany_list"` // 是否展示大航海陪伴列表
}

// GetLiveGuardActiveResult 是查询自己在某直播间观看时长的响应。
type GetLiveGuardActiveResult struct {
	Ruid      int64  `json:"ruid"`        // 主播 uid
	Rusername string `json:"rusername"`   // 主播用户名
	Rface     string `json:"rface"`       // 主播头像
	Username  string `json:"username"`    // 自己的用户名
	Accompany int    `json:"accomany"`    // 大航海陪伴天数。接口字段拼写为 accomany（少一个 p）
	WatchTime int64  `json:"watch_time"`  // 观看时长，单位为秒
	UpMedal   any    `json:"up_medal"`    // 粉丝牌信息
	GuardNum3 int    `json:"guard_num_3"` // 主播舰长数量
	GuardNum2 int    `json:"guard_num_2"` // 主播提督数量
	GuardNum1 int    `json:"guard_num_1"` // 主播总督数量
	IsLive    int    `json:"is_live"`     // 直播状态
}

// GetLiveUserCardResult 是查询用户在直播间内信息的响应。
type GetLiveUserCardResult struct {
	UID              int64                          `json:"uid"`                 // 目标用户 id
	Uname            string                         `json:"uname"`               // 目标用户名
	Face             string                         `json:"face"`                // 目标用户头像
	VerifyType       int                            `json:"verify_type"`         // 认证类型。0：无
	Desc             string                         `json:"desc"`                // 目标用户认证信息
	UnameColor       int                            `json:"uname_color"`         // 目标用户名颜色
	RoomID           int64                          `json:"room_id"`             // 目标用户直播间房间号
	Pendant          string                         `json:"pendant"`             // 目标用户佩戴的头像框 url
	PendantFrom      int                            `json:"pendant_from"`        // 头像框来源。作用尚不明确
	FollowNum        int64                          `json:"follow_num"`          // 目标用户粉丝数
	AttentionNum     int64                          `json:"attention_num"`       // 目标用户关注数
	RelationStatus   int                            `json:"relation_status"`     // 与目标用户的关系。1：已关注
	PrivilegeType    int                            `json:"privilege_type"`      // 目标用户大航海类型。0：无，3：舰长
	FansMedal        any                            `json:"fans_medal"`          // 粉丝牌。目标用户隐藏粉丝牌时为 null
	TitleSum         int                            `json:"title_sum"`           // 目标用户拥有的头衔数
	WearingTitle     string                         `json:"wearing_title"`       // 目标用户佩戴的头衔
	MainVip          int                            `json:"main_vip"`            // 目标用户大会员状态。0：无，2：大会员
	IsBlock          int                            `json:"is_block"`            // 是否被拉黑。始终为 0？
	IsAdmin          int                            `json:"is_admin"`            // 是否房管
	FansMedalListURL string                         `json:"fans_medal_list_url"` // 目标用户粉丝牌墙 url
	Wearing          any                            `json:"wearing"`             // 目标用户佩戴的装扮列表。文档示例中为 {}
	IsBlack          int                            `json:"is_black"`            // 是否被关小黑屋
	AdminLevel       int                            `json:"admin_level"`         // 房管等级
	HeadPicture      string                         `json:"head_picture"`        // 目标用户头像框 url
	HeadText         string                         `json:"head_text"`           // 头像框文字
	HeadURL          string                         `json:"head_url"`            // 头像框 url。作用尚不明确
	HeadBusiness     int                            `json:"head_business"`       // 头像框业务类型。作用尚不明确
	HeadSkinIcon     string                         `json:"head_skin_icon"`      // 头像框皮肤图标 url
	PrivilegeCenter  GetLiveUserCardPrivilegeCenter `json:"privilege_center"`    // 会员中心信息
	IsNFT            int                            `json:"is_nft"`              // 是否为 NFT 头像。0：否
	NFTDmark         string                         `json:"nft_dmark"`           // NFT 头像水印 url
	IsRealFans       bool                           `json:"is_real_fans"`        // 是否为真实粉丝
	WealthInfo       GetLiveUserCardWealthInfo      `json:"wealth_info"`         // 荣耀等级
	Guard            GetLiveUserCardGuard           `json:"guard"`               // 大航海信息
	ChatURLAndroid   string                         `json:"chat_url_android"`    // Android 端聊天跳转 url
	ChatURLIOS       string                         `json:"chat_url_ios"`        // iOS 端聊天跳转 url
	GiftStar         any                            `json:"gift_star"`           // 礼物星球。实测为 null，作用尚不明确
	IsMystery        bool                           `json:"is_mystery"`          // 是否为神秘人
	TextControl      any                            `json:"text_control"`        // 文本控制。实测为 null，作用尚不明确
	UinfoMedal       any                            `json:"uinfo_medal"`         // 用户信息勋章。实测为 null，作用尚不明确
	GuardAttire      any                            `json:"guard_attire"`        // 大航海装扮。实测为 null，作用尚不明确
	FlashBuy         int                            `json:"flash_buy"`           // 是否开启闪购。作用尚不明确
	FlashBuyURL      string                         `json:"flash_buy_url"`       // 闪购链接。作用尚不明确
}

// GetLiveUserCardPrivilegeCenter 是直播间用户卡片中的会员中心信息。
type GetLiveUserCardPrivilegeCenter struct {
	LinkURL              string                           `json:"link_url"`                // 跳转 url
	Title                string                           `json:"title"`                   // 标题
	RedDotCopywriting    string                           `json:"red_dot_copywriting"`     // 红点文案
	LinkType             GetLiveUserCardPrivilegeLinkType `json:"link_type"`               // 各类型的跳转 url
	NewRedDotCopywriting string                           `json:"new_red_dot_copywriting"` // 新红点文案
	GloryDressOn         bool                             `json:"glory_dress_on"`          // 是否开启荣耀装扮。作用尚不明确
}

// GetLiveUserCardPrivilegeLinkType 是会员中心各类型的跳转 url。
type GetLiveUserCardPrivilegeLinkType struct {
	FansMedalType string `json:"fans_medal_type"` // 粉丝牌类型的跳转 url
	TitleType     string `json:"title_type"`      // 头衔类型的跳转 url
	NewType       string `json:"new_type"`        // 新类型的跳转 url
}

// GetLiveUserCardWealthInfo 是直播间用户卡片中的荣耀等级信息。
type GetLiveUserCardWealthInfo struct {
	UID              int64  `json:"uid"`                // 目标用户 id
	Level            int    `json:"level"`              // 荣耀等级
	LevelTotalScore  int64  `json:"level_total_score"`  // 下一等级的荣耀值
	CurScore         int64  `json:"cur_score"`          // 当前荣耀值
	UpgradeNeedScore int64  `json:"upgrade_need_score"` // 升到下一等级还差多少荣耀值
	Status           int    `json:"status"`             // 荣耀等级状态。作用尚不明确
	DmIconKey        string `json:"dm_icon_key"`        // 荣耀等级图标 key，实测为空
}

// GetLiveUserCardGuard 是直播间用户卡片中的大航海信息。
type GetLiveUserCardGuard struct {
	Accompany      int `json:"accompany"`       // 陪伴人数。作用尚不明确
	AccompanySlake int `json:"accompany_slake"` // 陪伴人数（slake）。作用尚不明确
}
