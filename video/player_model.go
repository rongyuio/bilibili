package video

// web 播放器信息响应模型。

// GetVideoPlayerInfoResult 是 web 播放器信息的 data 对象。
type GetVideoPlayerInfoResult struct {
	Aid               int64                           `json:"aid"`                  // 视频 aid
	Bvid              string                          `json:"bvid"`                 // 视频 bvid
	AllowBp           bool                            `json:"allow_bp"`             // （？）。作用尚不明确
	NoShare           bool                            `json:"no_share"`             // 是否禁止分享
	Cid               int64                           `json:"cid"`                  // 视频 cid
	MaxLimit          int                             `json:"max_limit"`            // （？）。文档未列出，示例中存在
	PageNo            int                             `json:"page_no"`              // （？）。文档未列出，示例中存在
	HasNext           bool                            `json:"has_next"`             // （？）。文档未列出，示例中存在
	DmMask            *GetVideoPlayerInfoDmMask       `json:"dm_mask"`              // webmask 防挡字幕信息。若无则没有防挡功能
	Subtitle          *GetVideoPlayerInfoSubtitle     `json:"subtitle"`             // 字幕信息。若无则没有字幕，未登录时为空
	ViewPoints        []GetVideoPlayerInfoViewPoint   `json:"view_points"`          // 分段章节信息
	IPInfo            GetVideoPlayerInfoIPInfo        `json:"ip_info"`              // 请求 IP 信息
	LoginMid          int64                           `json:"login_mid"`            // 登录用户 mid
	LoginMidHash      string                          `json:"login_mid_hash"`       // 登录用户 mid 的哈希
	IsOwner           bool                            `json:"is_owner"`             // 是否为该视频 UP 主
	Name              string                          `json:"name"`                 // 登录用户昵称
	Permission        string                          `json:"permission"`           // 权限。文档标为 num，实际响应为逗号分隔的字符串
	LevelInfo         GetVideoPlayerInfoLevelInfo     `json:"level_info"`           // 登录用户等级信息
	Vip               GetVideoPlayerInfoVip           `json:"vip"`                  // 登录用户 VIP 信息
	AnswerStatus      int                             `json:"answer_status"`        // 答题状态
	BlockTime         int64                           `json:"block_time"`           // 封禁时间？
	Role              string                          `json:"role"`                 // 用户角色
	LastPlayTime      int64                           `json:"last_play_time"`       // 上次观看时间？
	LastPlayCid       int64                           `json:"last_play_cid"`        // 上次观看 cid？
	NowTime           int64                           `json:"now_time"`             // 当前 UNIX 秒级时间戳
	OnlineCount       int                             `json:"online_count"`         // 在线人数
	NeedLoginSubtitle bool                            `json:"need_login_subtitle"`  // 是否必须登录才能查看字幕
	PreviewToast      string                          `json:"preview_toast"`        // 付费预览提示。如 `为创作付费，购买观看完整视频|购买观看`
	Interaction       *GetVideoPlayerInfoInteraction  `json:"interaction"`          // 互动视频资讯。非互动视频没有该键
	Options           GetVideoPlayerInfoOptions       `json:"options"`              // 播放器选项
	GuideAttention    []any                           `json:"guide_attention"`      // 引导关注信息。作用尚不明确
	JumpCard          []any                           `json:"jump_card"`            // 跳转卡片信息。作用尚不明确
	OperationCard     []any                           `json:"operation_card"`       // 操作卡片信息。作用尚不明确
	OnlineSwitch      GetVideoPlayerInfoOnlineSwitch  `json:"online_switch"`        // 在线状态开关
	Fawkes            GetVideoPlayerInfoFawkes        `json:"fawkes"`               // 播放器相关信息？
	ShowSwitch        GetVideoPlayerInfoShowSwitch    `json:"show_switch"`          // 显示开关
	BgmInfo           *GetVideoPlayerInfoBgmInfo      `json:"bgm_info"`             // 背景音乐信息。可能为 null
	ToastBlock        bool                            `json:"toast_block"`          // （？）。作用尚不明确
	IsUpowerExclusive bool                            `json:"is_upower_exclusive"`  // 是否为充电专属视频
	IsUpowerPlay      bool                            `json:"is_upower_play"`       // （？）。作用尚不明确
	IsUgcPayPreview   bool                            `json:"is_ugc_pay_preview"`   // 是否为 UGC 付费预览
	ElecHighLevel     GetVideoPlayerInfoElecHighLevel `json:"elec_high_level"`      // 充电专属视频信息
	DisableShowUpInfo bool                            `json:"disable_show_up_info"` // （？）。作用尚不明确
}

// GetVideoPlayerInfoLevelInfo 是 web 播放器信息中的登录用户等级信息。
type GetVideoPlayerInfoLevelInfo struct {
	CurrentLevel int   `json:"current_level"` // 当前等级
	CurrentMin   int   `json:"current_min"`   // 当前等级经验最低值
	CurrentExp   int   `json:"current_exp"`   // 当前经验
	NextExp      int   `json:"next_exp"`      // 升级下一等级需达到的经验
	LevelUp      int64 `json:"level_up"`      // 升级时间？
}

// GetVideoPlayerInfoDmMask 是 webmask 防挡字幕信息。
type GetVideoPlayerInfoDmMask struct {
	Cid     int64  `json:"cid"`      // 视频 cid
	Plat    int    `json:"plat"`     // 未知
	Fps     int    `json:"fps"`      // webmask 取样 fps
	Time    int64  `json:"time"`     // 未知
	MaskURL string `json:"mask_url"` // webmask 资源 url
}

// GetVideoPlayerInfoSubtitle 是字幕信息。
type GetVideoPlayerInfoSubtitle struct {
	AllowSubmit bool                             `json:"allow_submit"` // 是否允许提交字幕
	Lan         string                           `json:"lan"`          // 默认语言类型
	LanDoc      string                           `json:"lan_doc"`      // 默认语言中文名称
	Subtitles   []GetVideoPlayerInfoSubtitleItem `json:"subtitles"`    // 字幕列表。未登录为 []
}

// GetVideoPlayerInfoSubtitleItem 是字幕列表中的一项。
type GetVideoPlayerInfoSubtitleItem struct {
	AiStatus    int    `json:"ai_status"`    // AI 字幕状态
	AiType      int    `json:"ai_type"`      // AI 字幕类型
	ID          int64  `json:"id"`           // 字幕 id
	IDStr       string `json:"id_str"`       // 字幕 id 字符串。和 id 不一样
	IsLock      bool   `json:"is_lock"`      // 是否锁定
	Lan         string `json:"lan"`          // 语言类型英文字母缩写
	LanDoc      string `json:"lan_doc"`      // 语言类型中文名称
	SubtitleURL string `json:"subtitle_url"` // 字幕资源 url
	Type        int    `json:"type"`         // 字幕类型。0
}

// GetVideoPlayerInfoViewPoint 是分段章节信息中的一项。
type GetVideoPlayerInfoViewPoint struct {
	Content  string `json:"content"`   // 分段章节名
	From     int    `json:"from"`      // 分段章节起始秒数
	To       int    `json:"to"`        // 分段章节结束秒数
	Type     int    `json:"type"`      // 章节类型
	ImgURL   string `json:"imgUrl"`    // 图片资源地址
	LogoURL  string `json:"logoUrl"`   // 图标资源地址
	TeamType string `json:"team_type"` // 团队类型
	TeamName string `json:"team_name"` // 团队名称
}

// GetVideoPlayerInfoIPInfo 是请求 IP 信息。
type GetVideoPlayerInfoIPInfo struct {
	IP       string `json:"ip"`       // 请求 IP
	ZoneIP   string `json:"zone_ip"`  // 区域 IP
	ZoneID   int64  `json:"zone_id"`  // 区域 id
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
}

// GetVideoPlayerInfoInteraction 是互动视频资讯。非互动视频没有该键。
type GetVideoPlayerInfoInteraction struct {
	GraphVersion int64  `json:"graph_version"` // 剧情图 id
	Msg          string `json:"msg"`           // 提示信息。未登录有机会返回 `登录后才能体验全部结局哦～`
	ErrorToast   string `json:"error_toast"`   // 错误信息？
	Mark         int    `json:"mark"`          // 0？
	NeedReload   int    `json:"need_reload"`   // 0？
}

// GetVideoPlayerInfoOptions 是播放器选项。
type GetVideoPlayerInfoOptions struct {
	Is360      bool `json:"is_360"`      // 是否 360 全景视频
	WithoutVip bool `json:"without_vip"` // （？）。作用尚不明确
}

// GetVideoPlayerInfoOnlineSwitch 是在线状态开关。字段值在响应中为字符串。
type GetVideoPlayerInfoOnlineSwitch struct {
	EnableGrayDashPlayback string `json:"enable_gray_dash_playback"` // 是否启用灰度 dash 播放
	NewBroadcast           string `json:"new_broadcast"`             // 是否启用新版广播
	RealtimeDm             string `json:"realtime_dm"`               // 是否启用实时弹幕
	SubtitleSubmitSwitch   string `json:"subtitle_submit_switch"`    // 字幕提交开关
}

// GetVideoPlayerInfoFawkes 是播放器相关信息。
type GetVideoPlayerInfoFawkes struct {
	ConfigVersion int `json:"config_version"` // 配置版本
	FfVersion     int `json:"ff_version"`     // 特性开关版本
}

// GetVideoPlayerInfoShowSwitch 是显示开关。
type GetVideoPlayerInfoShowSwitch struct {
	LongProgress bool `json:"long_progress"` // 是否显示长进度条
}

// GetVideoPlayerInfoBgmInfo 是背景音乐信息。
type GetVideoPlayerInfoBgmInfo struct {
	MusicID    string `json:"music_id"`    // 音乐 id
	MusicTitle string `json:"music_title"` // 音乐标题
	JumpURL    string `json:"jump_url"`    // 跳转 URL
}

// GetVideoPlayerInfoElecHighLevel 是充电专属视频信息。
type GetVideoPlayerInfoElecHighLevel struct {
	PrivilegeType int    `json:"privilege_type"` // 解锁视频所需最低定价档位的代码
	Title         string `json:"title"`          // 提示标题
	SubTitle      string `json:"sub_title"`      // 提示子标题
	ShowButton    bool   `json:"show_button"`    // 是否显示按钮
	ButtonText    string `json:"button_text"`    // 按钮文本
	JumpURL       any    `json:"jump_url"`       // 跳转 url 信息。文档标为 obj，示例中为空字符串
	Intro         string `json:"intro"`          // 充电介绍语
	Open          bool   `json:"open"`           // （？）。作用尚不明确
	New           bool   `json:"new"`            // （？）。作用尚不明确
	QuestionText  string `json:"question_text"`  // （？）。作用尚不明确
	QaDetailLink  string `json:"qa_detail_link"` // （？）。作用尚不明确
}

// GetVideoPlayerInfoVip 是登录用户的大会员信息。
type GetVideoPlayerInfoVip struct {
	Type               int                             `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                             `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                             `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                             `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int                             `json:"theme_type"`           // 0。作用尚不明确
	Label              GetVideoPlayerInfoVipLabel      `json:"label"`                // 会员标签
	AvatarSubscript    int                             `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                          `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int                             `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                             `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                             `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetVideoPlayerInfoVipOttInfo    `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetVideoPlayerInfoVipSuperVip   `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                             `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetVideoPlayerInfoVipAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
}

// GetVideoPlayerInfoVipLabel 是大会员标签。
type GetVideoPlayerInfoVipLabel struct {
	Path                  string                         `json:"path"`                      // 空。作用尚不明确
	Text                  string                         `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                         `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                         `json:"text_color"`                // 会员标签
	BgStyle               int                            `json:"bg_style"`                  // 1
	BgColor               string                         `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                         `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                           `json:"use_img_label"`             // true
	ImgLabelURIHans       string                         `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                         `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                         `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                         `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                            `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                            `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetVideoPlayerInfoVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                         `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                         `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetVideoPlayerInfoVipLabelGoto 是大会员标签的跳转信息。
type GetVideoPlayerInfoVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetVideoPlayerInfoVipOttInfo 是大会员信息中的 ott_info 对象。
type GetVideoPlayerInfoVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetVideoPlayerInfoVipSuperVip 是大会员信息中的 super_vip 对象。
type GetVideoPlayerInfoVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetVideoPlayerInfoVipAvatarIcon 是大会员角标信息。
type GetVideoPlayerInfoVipAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}
