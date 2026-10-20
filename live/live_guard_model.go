package live

// 大航海与粉丝团成员响应模型。信封 code/message/ttl 由库统一处理。
//
// 每个接口的返回体独立定义，嵌套对象均带方法前缀，不与其它接口共享类型。

// GetLiveGuardMemberListResult 是查询大航海成员的响应。
type GetLiveGuardMemberListResult struct {
	Info          GetLiveGuardMemberListInfo         `json:"info"`           // 榜单信息
	Top3          []GetLiveGuardMemberListItem       `json:"top3"`           // 整个列表的 top3
	List          []GetLiveGuardMemberListItem       `json:"list"`           // 大航海成员。page=1 时 list[0] 为榜单第四名
	MyFollowInfo  GetLiveGuardMemberListMyFollowInfo `json:"my_follow_info"` // 当前账号的陪伴信息
	GuardWarn     GetLiveGuardMemberListGuardWarn    `json:"guard_warn"`     // 大航海到期提醒
	ExistBenefit  bool                               `json:"exist_benefit"`  // 是否存在大航海权益
	RemindBenefit string                             `json:"remind_benefit"` // 权益提醒文案
	AB            GetLiveGuardMemberListAB           `json:"ab"`             // AB 测试信息
	RemindMsg     string                             `json:"remind_msg"`     // 提醒文案
	Typ           int                                `json:"typ"`            // 类型
	Extop         any                                `json:"extop"`          // (?)。作用尚不明确
	GuardLeader   any                                `json:"guard_leader"`   // 大航海首领
	MainText      string                             `json:"main_text"`      // 主文案
	SubText       string                             `json:"sub_text"`       // 副文案
	BtnType       int                                `json:"btn_type"`       // 按钮类型
	PromptText    string                             `json:"prompt_text"`    // 提示文案
}

// GetLiveGuardMemberListInfo 是大航海成员榜单的信息。
type GetLiveGuardMemberListInfo struct {
	Num                     int      `json:"num"`                        // 成员数量
	Page                    int      `json:"page"`                       // 总页数
	Now                     int      `json:"now"`                        // 当前页数
	AchievementLevel        int      `json:"achievement_level"`          // 成就等级
	AnchorGuardAchieveLevel int      `json:"anchor_guard_achieve_level"` // 主播大航海成就等级
	AchievementIconSrc      string   `json:"achievement_icon_src"`       // 成就图标 URL
	BuyGuardIconSrc         string   `json:"buy_guard_icon_src"`         // 购买大航海图标 URL
	RuleDocSrc              string   `json:"rule_doc_src"`               // 规则文档 URL
	ExBackgroundSrc         string   `json:"ex_background_src"`          // 额外背景图 URL
	ColorStart              string   `json:"color_start"`                // 渐变色起始
	ColorEnd                string   `json:"color_end"`                  // 渐变色结束
	TabColor                []string `json:"tab_color"`                  // 标签颜色
	TitleColor              []string `json:"title_color"`                // 标题颜色
}

// GetLiveGuardMemberListItem 是大航海成员榜单中的一名成员。
type GetLiveGuardMemberListItem struct {
	Ruid      int64                           `json:"ruid"`      // 主播 UID
	Rank      int                             `json:"rank"`      // 榜单排名
	Accompany int                             `json:"accompany"` // 陪伴天数
	Uinfo     GetLiveGuardMemberListItemUinfo `json:"uinfo"`     // 用户信息
	Score     int                             `json:"score"`     // 亲密度。恒为 0
}

// GetLiveGuardMemberListItemUinfo 是大航海成员的用户信息。
type GetLiveGuardMemberListItemUinfo struct {
	UID         int64                           `json:"uid"`          // 用户 UID
	Base        GetLiveGuardMemberListItemBase  `json:"base"`         // 用户基本信息
	Medal       GetLiveGuardMemberListItemMedal `json:"medal"`        // 粉丝牌
	Wealth      any                             `json:"wealth"`       // 荣耀等级相关信息，可能为 null
	Title       any                             `json:"title"`        // 头衔信息，可能为 null
	Guard       GetLiveGuardMemberListItemGuard `json:"guard"`        // 大航海信息
	UheadFrame  any                             `json:"uhead_frame"`  // 头像框，可能为 null
	GuardLeader any                             `json:"guard_leader"` // 大航海首领，可能为 null
}

// GetLiveGuardMemberListItemBase 是大航海成员的用户基本信息。
type GetLiveGuardMemberListItemBase struct {
	Name         string                                 `json:"name"`           // 用户名
	Face         string                                 `json:"face"`           // 用户头像
	NameColor    int                                    `json:"name_color"`     // 昵称颜色
	IsMystery    bool                                   `json:"is_mystery"`     // 是否为神秘人
	RiskCtrlInfo any                                    `json:"risk_ctrl_info"` // 风控信息，可能为 null
	OriginInfo   GetLiveGuardMemberListItemOriginInfo   `json:"origin_info"`    // 原始用户信息
	OfficialInfo GetLiveGuardMemberListItemOfficialInfo `json:"official_info"`  // 官方认证信息
	NameColorStr string                                 `json:"name_color_str"` // 昵称颜色字符串
}

// GetLiveGuardMemberListItemOriginInfo 是大航海成员未隐藏身份时的原始信息。
type GetLiveGuardMemberListItemOriginInfo struct {
	Name string `json:"name"` // 用户名
	Face string `json:"face"` // 用户头像
}

// GetLiveGuardMemberListItemOfficialInfo 是大航海成员的官方认证信息。
type GetLiveGuardMemberListItemOfficialInfo struct {
	Role  int    `json:"role"`  // 认证角色
	Title string `json:"title"` // 认证标题
	Desc  string `json:"desc"`  // 认证描述
	Type  int    `json:"type"`  // 认证类型。-1：无
}

// GetLiveGuardMemberListItemMedal 是大航海成员的用户勋章信息。
type GetLiveGuardMemberListItemMedal struct {
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

// GetLiveGuardMemberListItemGuard 是大航海成员的舰队信息。
type GetLiveGuardMemberListItemGuard struct {
	Level      int    `json:"level"`       // 大航海类型。1、2、3 分别为总督、提督、舰长
	ExpiredStr string `json:"expired_str"` // 到期时间文案
}

// GetLiveGuardMemberListMyFollowInfo 是当前账号在榜单中的陪伴信息。
type GetLiveGuardMemberListMyFollowInfo struct {
	AccompanyDays int                               `json:"accompany_days"` // 陪伴天数
	AutoRenew     int                               `json:"auto_renew"`     // 是否自动续费
	RenewRemind   GetLiveGuardMemberListRenewRemind `json:"renew_remind"`   // 续费提醒
	Rank          int                               `json:"rank"`           // 榜单排名
	Ruid          int64                             `json:"ruid"`           // 主播 UID
	Uinfo         any                               `json:"uinfo"`          // 用户信息，可能为 null
	ExpiredTime   string                            `json:"expired_time"`   // 到期时间
}

// GetLiveGuardMemberListRenewRemind 是大航海续费提醒。
type GetLiveGuardMemberListRenewRemind struct {
	Content string `json:"content"` // 提醒内容
	Type    int    `json:"type"`    // 提醒类型
	Hint    string `json:"hint"`    // 提示
}

// GetLiveGuardMemberListGuardWarn 是大航海到期提醒。
type GetLiveGuardMemberListGuardWarn struct {
	IsWarn      int    `json:"is_warn"`      // 是否提醒
	Warn        string `json:"warn"`         // 提醒内容
	Expired     int    `json:"expired"`      // 是否已到期
	WillExpired int    `json:"will_expired"` // 是否即将到期
	Address     string `json:"address"`      // 跳转地址
}

// GetLiveGuardMemberListAB 是大航海榜单的 AB 测试信息。
type GetLiveGuardMemberListAB struct {
	GuardAccompanyList int `json:"guard_accompany_list"` // 是否启用陪伴列表
}

// GetLiveFansMemberListResult 是查询粉丝团成员的响应。
type GetLiveFansMemberListResult struct {
	Item        []GetLiveFansMemberListItem `json:"item"`         // 粉丝团成员列表
	Num         int                         `json:"num"`          // 粉丝团成员数量
	MedalStatus int                         `json:"medal_status"` // 粉丝牌状态
}

// GetLiveFansMemberListItem 是粉丝团成员榜单中的一名成员。
type GetLiveFansMemberListItem struct {
	UserRank         int                                 `json:"user_rank"`          // 排名
	UID              int64                               `json:"uid"`                // 用户 UID
	Name             string                              `json:"name"`               // 用户名
	Face             string                              `json:"face"`               // 用户头像
	Score            int                                 `json:"score"`              // 亲密度
	MedalName        string                              `json:"medal_name"`         // 粉丝牌名字
	Level            int                                 `json:"level"`              // 粉丝牌等级
	TargetID         int64                               `json:"target_id"`          // 主播 UID
	Special          string                              `json:"special"`            // 特殊标识
	GuardLevel       int                                 `json:"guard_level"`        // 大航海类型。1、2、3 分别为总督、提督、舰长
	MedalColorStart  int                                 `json:"medal_color_start"`  // 粉丝牌渐变起始色
	MedalColorEnd    int                                 `json:"medal_color_end"`    // 粉丝牌渐变结束色
	MedalColorBorder int                                 `json:"medal_color_border"` // 粉丝牌边框颜色
	GuardIcon        string                              `json:"guard_icon"`         // 大航海图标 URL
	HonorIcon        string                              `json:"honor_icon"`         // 荣誉图标 URL
	UinfoMedal       GetLiveFansMemberListItemUinfoMedal `json:"uinfo_medal"`        // 粉丝牌
	Tag              any                                 `json:"tag"`                // 标签，可能为 null
	IsPokeable       bool                                `json:"is_pokeable"`        // 是否可被戳一戳
}

// GetLiveFansMemberListItemUinfoMedal 是粉丝团成员的用户勋章信息。
type GetLiveFansMemberListItemUinfoMedal struct {
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
