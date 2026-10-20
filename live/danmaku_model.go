package live

// 直播间弹幕相关响应模型（data 层）。信封字段由库统一处理。

// GetLiveDanmakuConfigResult 是当前用户对应直播间的可发弹幕配置。
type GetLiveDanmakuConfigResult struct {
	Group []GetLiveDanmakuConfigGroup `json:"group"` // 弹幕颜色组
	Mode  []GetLiveDanmakuConfigMode  `json:"mode"`  // 弹幕显示模式
}

// GetLiveDanmakuConfigGroup 是一组弹幕颜色。
type GetLiveDanmakuConfigGroup struct {
	Name  string                      `json:"name"`  // 显示名称
	Sort  int                         `json:"sort"`  // 用于排序
	Color []GetLiveDanmakuConfigColor `json:"color"` // 该组颜色列表
}

// GetLiveDanmakuConfigColor 是单个弹幕颜色。
type GetLiveDanmakuConfigColor struct {
	Name     string `json:"name"`      // 颜色名
	Color    string `json:"color"`     // 十进制颜色值，发送弹幕用
	ColorHex string `json:"color_hex"` // 十六进制颜色值，发送弹幕用
	Status   int    `json:"status"`    // 可用状态。0：不可用；1：可用
	Weight   int    `json:"weight"`    // (?)。作用尚不明确
	ColorID  int    `json:"color_id"`  // 颜色 id
	Origin   int    `json:"origin"`    // (?)。作用尚不明确
}

// GetLiveDanmakuConfigMode 是一种弹幕显示模式。
type GetLiveDanmakuConfigMode struct {
	Name   string `json:"name"`   // 模式名称
	Mode   int    `json:"mode"`   // 模式值，发送弹幕用
	Type   string `json:"type"`   // 模式类型
	Status int    `json:"status"` // 可用状态。0：不可用；1：可用
}

// GetLiveRoomHistoryDanmakuResult 是直播间最近的历史弹幕。
type GetLiveRoomHistoryDanmakuResult struct {
	Admin []GetLiveRoomHistoryDanmakuItem `json:"admin"` // 管理员最新的 10 条弹幕
	Room  []GetLiveRoomHistoryDanmakuItem `json:"room"`  // 普通用户最新的 10 条弹幕
}

// GetLiveRoomHistoryDanmakuItem 是一条历史弹幕。
//
// medal、title、user_level 在文档中标注为格式不明，保留为 any 切片。
type GetLiveRoomHistoryDanmakuItem struct {
	Text         string                                   `json:"text"`           // 弹幕内容
	DMType       int                                      `json:"dm_type"`        // 弹幕类型
	UID          int64                                    `json:"uid"`            // 弹幕发送者的 UID
	Nickname     string                                   `json:"nickname"`       // 弹幕发送者的昵称
	UnameColor   string                                   `json:"uname_color"`    // 弹幕发送者的昵称颜色
	Timeline     string                                   `json:"timeline"`       // 发送时间，格式为 yyyy-MM-dd HH:mm:ss
	IsAdmin      int                                      `json:"isadmin"`        // 是否为管理员
	VIP          int                                      `json:"vip"`            // 是否为 VIP
	SVIP         int                                      `json:"svip"`           // 是否为 SVIP
	Medal        []any                                    `json:"medal"`          // 粉丝勋章信息，格式不明
	Title        []any                                    `json:"title"`          // 标题，格式不明
	UserLevel    []any                                    `json:"user_level"`     // 用户等级信息，格式不明
	Rank         int                                      `json:"rank"`           // 排名
	TeamID       int                                      `json:"teamid"`         // 所属战队 id
	Rnd          string                                   `json:"rnd"`            // 随机数种子
	UserTitle    string                                   `json:"user_title"`     // 用户头衔
	GuardLevel   int                                      `json:"guard_level"`    // 大航海等级
	Bubble       int                                      `json:"bubble"`         // 气泡 id
	BubbleColor  string                                   `json:"bubble_color"`   // 气泡颜色
	LPL          int                                      `json:"lpl"`            // (?)
	YeahSpaceURL string                                   `json:"yeah_space_url"` // (?)
	JumpToURL    string                                   `json:"jump_to_url"`    // (?)
	CheckInfo    *GetLiveRoomHistoryDanmakuCheckInfo      `json:"check_info"`     // 弹幕审核信息
	VoiceDMInfo  *GetLiveRoomHistoryDanmakuVoiceDMInfo    `json:"voice_dm_info"`  // 语音弹幕信息
	Emoticon     *GetLiveRoomHistoryDanmakuEmoticon       `json:"emoticon"`       // 房间独有表情信息
	Emots        map[string]GetLiveRoomHistoryDanmakuEmot `json:"emots"`          // 默认表情信息，键为表情名；不含默认表情时为 null
	IDStr        string                                   `json:"id_str"`         // 弹幕 id
	WealthLevel  int                                      `json:"wealth_level"`   // 财富等级
	BubbleIDV2   int                                      `json:"bubble_id_v2"`   // 气泡 id（新版）
	Reply        *GetLiveRoomHistoryDanmakuReply          `json:"reply"`          // 回复的弹幕
	GroupMedal   any                                      `json:"group_medal"`    // 粉丝团勋章（通常为 null）
	User         *GetLiveRoomHistoryDanmakuUser           `json:"user"`           // 该用户信息
}

// GetLiveRoomHistoryDanmakuCheckInfo 是弹幕审核信息。
type GetLiveRoomHistoryDanmakuCheckInfo struct {
	Ts int64  `json:"ts"` // 审核时间戳
	Ct string `json:"ct"` // 审核校验值
}

// GetLiveRoomHistoryDanmakuVoiceDMInfo 是语音弹幕信息。
type GetLiveRoomHistoryDanmakuVoiceDMInfo struct {
	VoiceURL     string `json:"voice_url"`     // 语音链接
	FileFormat   string `json:"file_format"`   // 文件格式
	Text         string `json:"text"`          // 语音转写文本
	FileDuration int    `json:"file_duration"` // 文件时长
	FileID       string `json:"file_id"`       // 文件 id
}

// GetLiveRoomHistoryDanmakuEmoticon 是房间独有表情信息。
type GetLiveRoomHistoryDanmakuEmoticon struct {
	ID             int    `json:"id"`              // 表情 id
	EmoticonUnique string `json:"emoticon_unique"` // 表情的独特标识，格式为 room_房间号_表情id
	Text           string `json:"text"`            // 表情的触发词
	Perm           int    `json:"perm"`            // 发送权限。1：所有人都可发送
	URL            string `json:"url"`             // 表情的图像链接
	InPlayerArea   int    `json:"in_player_area"`  // 是否显示在直播画面区域
	BulgeDisplay   int    `json:"bulge_display"`   // 是否高亮显示
	IsDynamic      int    `json:"is_dynamic"`      // 是否为动态表情。0：静态图像；1：动态图像
	Height         int    `json:"height"`          // 表情的高度
	Width          int    `json:"width"`           // 表情的宽度
}

// GetLiveRoomHistoryDanmakuEmot 是默认表情信息（emots 对象中的值）。
type GetLiveRoomHistoryDanmakuEmot struct {
	Count         int    `json:"count"`          // 重复发送数量
	Descript      string `json:"descript"`       // 表情描述
	Emoji         string `json:"emoji"`          // 表情描述
	EmotionID     int    `json:"emotion_id"`     // 表情 id
	EmotionUnique string `json:"emotion_unique"` // 表情的独特标识符，格式可能为 emoji_表情id
	Height        int    `json:"height"`         // 表情的高度
	URL           string `json:"url"`            // 表情的图像链接
	Width         int    `json:"width"`          // 表情的宽度
}

// GetLiveRoomHistoryDanmakuReply 是被回复的弹幕信息。
type GetLiveRoomHistoryDanmakuReply struct {
	ShowReply       bool   `json:"show_reply"`        // 是否展示回复
	ReplyMid        int64  `json:"reply_mid"`         // 被回复用户 mid
	ReplyUname      string `json:"reply_uname"`       // 被回复用户昵称
	ReplyUnameColor string `json:"reply_uname_color"` // 被回复用户昵称颜色
	ReplyIsMystery  bool   `json:"reply_is_mystery"`  // 被回复用户是否为神秘用户
}

// GetLiveRoomHistoryDanmakuUser 是弹幕发送者的用户信息。
type GetLiveRoomHistoryDanmakuUser struct {
	UID         int64                                 `json:"uid"`          // 用户 mid
	Base        *GetLiveRoomHistoryDanmakuUserBase    `json:"base"`         // 用户基本信息
	Medal       *GetLiveRoomHistoryDanmakuUserMedal   `json:"medal"`        // 粉丝勋章信息
	Wealth      any                                   `json:"wealth"`       // 财富信息（通常为 null）
	Title       *GetLiveRoomHistoryDanmakuUserTitle   `json:"title"`        // 用户标题
	Guard       any                                   `json:"guard"`        // 大航海信息（通常为 null）
	UheadFrame  any                                   `json:"uhead_frame"`  // 头像框（通常为 null）
	GuardLeader *GetLiveRoomHistoryDanmakuGuardLeader `json:"guard_leader"` // 大航海首领信息
}

// GetLiveRoomHistoryDanmakuUserBase 是用户基本信息。
type GetLiveRoomHistoryDanmakuUserBase struct {
	Name         string                                 `json:"name"`           // 用户名
	Face         string                                 `json:"face"`           // 用户头像 url
	NameColor    int                                    `json:"name_color"`     // 用户名颜色
	IsMystery    bool                                   `json:"is_mystery"`     // 是否为神秘用户
	RiskCtrlInfo any                                    `json:"risk_ctrl_info"` // 风控信息（通常为 null）
	OriginInfo   *GetLiveRoomHistoryDanmakuOriginInfo   `json:"origin_info"`    // 原始信息
	OfficialInfo *GetLiveRoomHistoryDanmakuOfficialInfo `json:"official_info"`  // 认证信息
	NameColorStr string                                 `json:"name_color_str"` // 用户名颜色字符串
}

// GetLiveRoomHistoryDanmakuOriginInfo 是用户的原始信息。
type GetLiveRoomHistoryDanmakuOriginInfo struct {
	Name string `json:"name"` // 用户名
	Face string `json:"face"` // 用户头像 url
}

// GetLiveRoomHistoryDanmakuOfficialInfo 是用户的认证信息。
type GetLiveRoomHistoryDanmakuOfficialInfo struct {
	Role  int    `json:"role"`  // 认证角色
	Title string `json:"title"` // 认证头衔
	Desc  string `json:"desc"`  // 认证描述
	Type  int    `json:"type"`  // 认证类型
}

// GetLiveRoomHistoryDanmakuUserMedal 是弹幕发送者佩戴的粉丝勋章。
type GetLiveRoomHistoryDanmakuUserMedal struct {
	Name               string `json:"name"`                  // 勋章名
	Level              int    `json:"level"`                 // 勋章等级
	ColorStart         int    `json:"color_start"`           // 渐变起始颜色
	ColorEnd           int    `json:"color_end"`             // 渐变结束颜色
	ColorBorder        int    `json:"color_border"`          // 边框颜色
	Color              int    `json:"color"`                 // 主颜色
	ID                 int    `json:"id"`                    // 勋章 id
	Typ                int    `json:"typ"`                   // 勋章类型
	IsLight            int    `json:"is_light"`              // 是否点亮
	Ruid               int64  `json:"ruid"`                  // 勋章所属主播 uid
	GuardLevel         int    `json:"guard_level"`           // 大航海等级
	Score              int64  `json:"score"`                 // 亲密度
	GuardIcon          string `json:"guard_icon"`            // 大航海图标 url
	HonorIcon          string `json:"honor_icon"`            // 荣誉图标 url
	V2MedalColorStart  string `json:"v2_medal_color_start"`  // 新版渐变起始颜色
	V2MedalColorEnd    string `json:"v2_medal_color_end"`    // 新版渐变结束颜色
	V2MedalColorBorder string `json:"v2_medal_color_border"` // 新版边框颜色
	V2MedalColorText   string `json:"v2_medal_color_text"`   // 新版文字颜色
	V2MedalColorLevel  string `json:"v2_medal_color_level"`  // 新版等级颜色
	UserReceiveCount   int    `json:"user_receive_count"`    // 用户已领取数量
}

// GetLiveRoomHistoryDanmakuUserTitle 是用户头衔。
type GetLiveRoomHistoryDanmakuUserTitle struct {
	OldTitleCSSID string `json:"old_title_css_id"` // 旧版头衔样式 id
	TitleCSSID    string `json:"title_css_id"`     // 头衔样式 id
}

// GetLiveRoomHistoryDanmakuGuardLeader 是大航海首领信息。
type GetLiveRoomHistoryDanmakuGuardLeader struct {
	IsGuardLeader bool `json:"is_guard_leader"` // 是否为舰队首领
}
