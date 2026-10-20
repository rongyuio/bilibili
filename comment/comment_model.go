package comment

// GetCommentsDetail、GetCommentReply、GetCommentsHotReply 的响应模型。
//
// 本文件的类型只服务于上述三个接口，不与其它接口共用；即使字段相同，
// 其它接口也在各自文件中单独定义。

// GetCommentsDetailResult 是「获取评论区明细」的 data 对象。
type GetCommentsDetailResult struct {
	Page        GetCommentsDetailPage     `json:"page"`         // 页信息
	Config      GetCommentsDetailConfig   `json:"config"`       // 评论区显示控制
	Replies     []*GetCommentsDetailReply `json:"replies"`      // 评论列表
	Hots        []*GetCommentsDetailReply `json:"hots"`         // 热评列表
	Upper       *GetCommentsDetailReply   `json:"upper"`        // 置顶评论
	Top         any                       `json:"top"`          // (?)
	Notice      *GetCommentsDetailNotice  `json:"notice"`       // 评论区公告信息
	Vote        int                       `json:"vote"`         // 投票评论?
	Blacklist   int                       `json:"blacklist"`    // (?)
	Assist      int                       `json:"assist"`       // (?)
	Mode        int                       `json:"mode"`         // 评论区类型id
	SupportMode []int                     `json:"support_mode"` // 评论区支持的类型id
	Folder      GetCommentsDetailFolder   `json:"folder"`       // 折叠相关信息
	LotteryCard any                       `json:"lottery_card"` // (?)
	ShowBvid    bool                      `json:"show_bvid"`    // 显示bvid?
	Control     GetCommentsDetailControl  `json:"control"`      // 评论区输入属性
}

// GetCommentsDetailPage 是评论区页码信息。
type GetCommentsDetailPage struct {
	Num    int `json:"num"`    // 当前页码
	Size   int `json:"size"`   // 每页项数
	Count  int `json:"count"`  // 根评论条数
	Acount int `json:"acount"` // 总计评论条数
}

// GetCommentsDetailConfig 是评论区显示控制。
type GetCommentsDetailConfig struct {
	ShowAdmin  int  `json:"showadmin"`    // 是否显示管理置顶
	ShowEntry  int  `json:"showentry"`    // (?)
	ShowFloor  int  `json:"showfloor"`    // 是否显示楼层号
	ShowTopic  int  `json:"showtopic"`    // 是否显示话题
	ShowUpFlag bool `json:"show_up_flag"` // 是否显示“UP 觉得很赞”标志
	ReadOnly   bool `json:"read_only"`    // 是否只读评论区
	ShowDelLog bool `json:"show_del_log"` // 是否显示删除记录
}

// GetCommentsDetailReply 是评论条目。
type GetCommentsDetailReply struct {
	Rpid         int                                `json:"rpid"`          // 评论 rpid
	Oid          int                                `json:"oid"`           // 评论区对象 id
	Type         int                                `json:"type"`          // 评论区类型代码。
	Mid          int                                `json:"mid"`           // 发送者 mid
	Root         int                                `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                                `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                                `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                                `json:"count"`         // 二级评论条数
	Rcount       int                                `json:"rcount"`        // 回复评论条数
	Floor        int                                `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                                `json:"state"`         // (?)
	FansGrade    int                                `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                                `json:"attr"`          // 某属性位？
	Ctime        int                                `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                             `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                             `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                             `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                                `json:"like"`          // 评论获赞数
	Action       int                                `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       GetCommentsDetailReplyMember       `json:"member"`        // 评论发送者信息
	Content      GetCommentsDetailReplyContent      `json:"content"`       // 评论信息
	Replies      []*GetCommentsDetailReply          `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                                `json:"assist"`        // (?)
	Folder       GetCommentsDetailReplyFolder       `json:"folder"`        // 折叠信息
	UpAction     GetCommentsDetailReplyUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                               `json:"show_follow"`   // (?)
	Invisible    bool                               `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetCommentsDetailReplyCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetCommentsDetailReplyReplyControl `json:"reply_control"` // 评论提示文案信息
}

// GetCommentsDetailReplyMember 是评论条目中的 member 对象。
type GetCommentsDetailReplyMember struct {
	Mid            string                               `json:"mid"`             // 发送者 mid
	Uname          string                               `json:"uname"`           // 发送者昵称
	Sex            string                               `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                               `json:"sign"`            // 发送者签名
	Avatar         string                               `json:"avatar"`          // 发送者头像 url
	Rank           string                               `json:"rank"`            // (?)
	DisplayRank    string                               `json:"DisplayRank"`     // (?)
	LevelInfo      GetCommentsDetailReplyLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetCommentsDetailReplyPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetCommentsDetailReplyNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetCommentsDetailReplyOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetCommentsDetailReplyVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetCommentsDetailReplyFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                                  `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                                  `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    GetCommentsDetailReplyUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                                 `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                               `json:"contract_desc"`   // 合作用户说明？
}

// GetCommentsDetailReplyLevelInfo 是评论条目中 member.level_info 对象。
type GetCommentsDetailReplyLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetCommentsDetailReplyPendant 是评论条目中 member.pendant 对象。
type GetCommentsDetailReplyPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetCommentsDetailReplyNameplate 是评论条目中 member.nameplate 对象。
type GetCommentsDetailReplyNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetCommentsDetailReplyOfficialVerify 是评论条目中 member.official_verify 对象。
type GetCommentsDetailReplyOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetCommentsDetailReplyVip 是评论条目中 member.vip 对象。
type GetCommentsDetailReplyVip struct {
	VipType            int                            `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                            `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                         `json:"dueRemark"`            // (?)
	AccessStatus       int                            `json:"accessStatus"`         // (?)
	VipStatus          int                            `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                         `json:"vipStatusWarn"`        // (?)
	ThemeType          int                            `json:"theme_type"`           // 会员样式 id
	Label              GetCommentsDetailReplyVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                            `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                         `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                         `json:"nickname_color"`       // 昵称颜色
}

// GetCommentsDetailReplyVipLabel 是评论条目中 member.vip.label 对象。
type GetCommentsDetailReplyVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetCommentsDetailReplyFansDetail 是评论条目中 member.fans_detail 对象。
type GetCommentsDetailReplyFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetCommentsDetailReplyUserSailing 是评论条目中 member.user_sailing 对象。
type GetCommentsDetailReplyUserSailing struct {
	Pendant         *GetCommentsDetailReplyPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetCommentsDetailReplyCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                            `json:"cardbg_with_focus"` // (?)
}

// GetCommentsDetailReplyCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type GetCommentsDetailReplyCardbg struct {
	ID      int                       `json:"id"`       // 评论条目装扮 id
	Name    string                    `json:"name"`     // 评论条目装扮名称
	Image   string                    `json:"image"`    // 评论条目装扮图片 url
	JumpURL string                    `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetCommentsDetailReplyFan `json:"fan"`      // 粉丝专属信息
	Type    string                    `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetCommentsDetailReplyFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type GetCommentsDetailReplyFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetCommentsDetailReplyContent 是评论条目中的 content 对象。
type GetCommentsDetailReplyContent struct {
	Message  string                          `json:"message"`  // 评论内容。**重要**
	Plat     int                             `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                          `json:"device"`   // 评论发送平台设备
	Members  []GetCommentsDetailReplyMember  `json:"members"`  // at 到的用户信息
	Emote    any                             `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                             `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                             `json:"max_line"` // 6。收起最大行数
	Pictures []GetCommentsDetailReplyPicture `json:"pictures"` // 评论图片数组
}

// GetCommentsDetailReplyPicture 是评论条目中 content.pictures 数组中的对象。
type GetCommentsDetailReplyPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetCommentsDetailReplyFolder 是评论条目中的 folder 对象。
type GetCommentsDetailReplyFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentsDetailReplyUpAction 是评论条目中的 up_action 对象。
type GetCommentsDetailReplyUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// GetCommentsDetailReplyCardLabel 是评论条目中 card_label 数组中的对象。
type GetCommentsDetailReplyCardLabel struct {
	Rpid             int    `json:"rpid"`              // 评论 rpid
	TextContent      string `json:"text_content"`      // 标签文本。已知有妙评
	TextColorDay     string `json:"text_color_day"`    // 日间文本颜色。十六进制颜色值，下同
	TextColorNight   string `json:"text_color_night"`  // 夜间文本颜色
	LabelColorDay    string `json:"label_color_day"`   // 日间标签颜色
	LabelColorNight  string `json:"label_color_night"` // 夜间标签颜色
	Image            string `json:"image"`             // 作用不明
	Type             int    `json:"type"`              // 1。作用不明
	Background       string `json:"background"`        // 背景图片 url
	BackgroundWidth  int    `json:"background_width"`  // 背景图片宽度
	BackgroundHeight int    `json:"background_height"` // 背景图片高度
	JumpURL          string `json:"jump_url"`          // 跳转链接
	Effect           int    `json:"effect"`            // 0。作用不明，可能用于控制动画，下同
	EffectStartTime  int    `json:"effect_start_time"` // 0
}

// GetCommentsDetailReplyReplyControl 是评论条目中的 reply_control 对象。
type GetCommentsDetailReplyReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}

// GetCommentsDetailNotice 是评论区公告信息。
type GetCommentsDetailNotice struct {
	Type       int    `json:"type"`        // 1。作用尚不明确
	Status     int    `json:"status"`      // 0。作用尚不明确
	Title      string `json:"title"`       // 空。作用尚不明确
	Msg        string `json:"msg"`         // 空。作用尚不明确
	ButtonText string `json:"button_text"` // 空。作用尚不明确
	ButtonURL  string `json:"button_url"`  // 空。作用尚不明确
}

// GetCommentsDetailFolder 是折叠相关信息。
type GetCommentsDetailFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentsDetailControl 是评论区输入属性。
type GetCommentsDetailControl struct {
	InputDisable           bool   `json:"input_disable"`             // 是否禁止新增评论。用户涉及合约争议，锁定该用户所有稿件、动态的评论区，不允许新增评论，root_input_text和child_input_text值为“当前评论区不可新增评论”
	RootInputText          string `json:"root_input_text"`           // 评论框文字
	ChildInputText         string `json:"child_input_text"`          // 评论框文字
	BgText                 string `json:"bg_text"`                   // 空评论区文字
	WebSelection           bool   `json:"web_selection"`             // 评论是否筛选后可见。false：无需筛选。true：需要筛选
	AnswerGuideText        string `json:"answer_guide_text"`         // 答题页面链接文字
	AnswerGuideIconURL     string `json:"answer_guide_icon_url"`     // 答题页面图标 url
	AnswerGuideIosURL      string `json:"answer_guide_ios_url"`      // 答题页面 ios url
	AnswerGuideAndroidURL  string `json:"answer_guide_android_url"`  // 答题页面安卓 url
	GiveupInputText        string `json:"giveup_input_text"`         // 放弃发表评论时的提示文字。如 不发没关系，请继续友善哦~
	ScreenshotIconState    int    `json:"screenshot_icon_state"`     // 截图图标状态。0：隐藏。1：显示
	UploadPictureIconState int    `json:"upload_picture_icon_state"` // 上传图片图标状态。0：隐藏。1：显示
	DisableJumpEmote       bool   `json:"disable_jump_emote"`        // 是否禁止点击表情跳转。实测为 false
	EnableCharged          bool   `json:"enable_charged"`            // 是否开启充电评论？。实测为 false，作用尚不明确
	EnableCmBizHelper      bool   `json:"enable_cm_biz_helper"`      // 是否开启商业评论助手？。实测为 false，作用尚不明确
	// 空评论区页面配置？（？）。文档只给了字段名与类型，没有对象表；实测为 null
	EmptyPage any `json:"empty_page"`
	// 预加载资源？（？）。文档只给了字段名与类型，没有对象表；实测为 null
	PreloadResources []any `json:"preload_resources"`
}

// GetCommentReplyResult 是「获取指定评论的回复」的 data 对象。
type GetCommentReplyResult struct {
	Config   GetCommentReplyConfig  `json:"config"`    // 评论区显示控制
	Control  GetCommentReplyControl `json:"control"`   // 评论区输入属性
	Page     GetCommentReplyPage    `json:"page"`      // 页面信息
	Replies  []*GetCommentReplyItem `json:"replies"`   // 评论对话树列表。最大内容数为20
	Root     *GetCommentReplyItem   `json:"root"`      // 根评论信息
	ShowBvid bool                   `json:"show_bvid"` // 显示 bvid?
	ShowText string                 `json:"show_text"` // (?)
	ShowType int                    `json:"show_type"` // (?)
	Upper    GetCommentReplyUpper   `json:"upper"`     // UP主 mid
}

// GetCommentReplyConfig 是评论区显示控制。
type GetCommentReplyConfig struct {
	ShowAdmin  int  `json:"showadmin"`    // 是否显示管理置顶
	ShowEntry  int  `json:"showentry"`    // (?)
	ShowFloor  int  `json:"showfloor"`    // 是否显示楼层号
	ShowTopic  int  `json:"showtopic"`    // 是否显示话题
	ShowUpFlag bool `json:"show_up_flag"` // 是否显示“UP 觉得很赞”标志
	ReadOnly   bool `json:"read_only"`    // 是否只读评论区
	ShowDelLog bool `json:"show_del_log"` // 是否显示删除记录
}

// GetCommentReplyControl 是评论区输入属性。
type GetCommentReplyControl struct {
	InputDisable           bool   `json:"input_disable"`             // 是否禁止新增评论。用户涉及合约争议，锁定该用户所有稿件、动态的评论区，不允许新增评论，root_input_text和child_input_text值为“当前评论区不可新增评论”
	RootInputText          string `json:"root_input_text"`           // 评论框文字
	ChildInputText         string `json:"child_input_text"`          // 评论框文字
	BgText                 string `json:"bg_text"`                   // 空评论区文字
	WebSelection           bool   `json:"web_selection"`             // 评论是否筛选后可见。false：无需筛选。true：需要筛选
	AnswerGuideText        string `json:"answer_guide_text"`         // 答题页面链接文字
	AnswerGuideIconURL     string `json:"answer_guide_icon_url"`     // 答题页面图标 url
	AnswerGuideIosURL      string `json:"answer_guide_ios_url"`      // 答题页面 ios url
	AnswerGuideAndroidURL  string `json:"answer_guide_android_url"`  // 答题页面安卓 url
	GiveupInputText        string `json:"giveup_input_text"`         // 放弃发表评论时的提示文字。如 不发没关系，请继续友善哦~
	ScreenshotIconState    int    `json:"screenshot_icon_state"`     // 截图图标状态。0：隐藏。1：显示
	UploadPictureIconState int    `json:"upload_picture_icon_state"` // 上传图片图标状态。0：隐藏。1：显示
	DisableJumpEmote       bool   `json:"disable_jump_emote"`        // 是否禁止点击表情跳转。实测为 false
	EnableCharged          bool   `json:"enable_charged"`            // 是否开启充电评论？。实测为 false，作用尚不明确
	EnableCmBizHelper      bool   `json:"enable_cm_biz_helper"`      // 是否开启商业评论助手？。实测为 false，作用尚不明确
	// 空评论区页面配置？（？）。文档只给了字段名与类型，没有对象表；实测为 null
	EmptyPage any `json:"empty_page"`
	// 预加载资源？（？）。文档只给了字段名与类型，没有对象表；实测为 null
	PreloadResources []any `json:"preload_resources"`
}

// GetCommentReplyPage 是评论区页码信息。
type GetCommentReplyPage struct {
	Num    int `json:"num"`    // 当前页码
	Size   int `json:"size"`   // 每页项数
	Count  int `json:"count"`  // 根评论条数
	Acount int `json:"acount"` // 总计评论条数
}

// GetCommentReplyItem 是评论对话树中的条目。
type GetCommentReplyItem struct {
	Rpid         int                             `json:"rpid"`          // 评论 rpid
	Oid          int                             `json:"oid"`           // 评论区对象 id
	Type         int                             `json:"type"`          // 评论区类型代码。
	Mid          int                             `json:"mid"`           // 发送者 mid
	Root         int                             `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                             `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                             `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                             `json:"count"`         // 二级评论条数
	Rcount       int                             `json:"rcount"`        // 回复评论条数
	Floor        int                             `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                             `json:"state"`         // (?)
	FansGrade    int                             `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                             `json:"attr"`          // 某属性位？
	Ctime        int                             `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                          `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                          `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                          `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                             `json:"like"`          // 评论获赞数
	Action       int                             `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       GetCommentReplyItemMember       `json:"member"`        // 评论发送者信息
	Content      GetCommentReplyItemContent      `json:"content"`       // 评论信息
	Replies      []*GetCommentReplyItem          `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                             `json:"assist"`        // (?)
	Folder       GetCommentReplyItemFolder       `json:"folder"`        // 折叠信息
	UpAction     GetCommentReplyItemUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                            `json:"show_follow"`   // (?)
	Invisible    bool                            `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetCommentReplyItemCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetCommentReplyItemReplyControl `json:"reply_control"` // 评论提示文案信息
}

// GetCommentReplyItemMember 是评论条目中的 member 对象。
type GetCommentReplyItemMember struct {
	Mid            string                            `json:"mid"`             // 发送者 mid
	Uname          string                            `json:"uname"`           // 发送者昵称
	Sex            string                            `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                            `json:"sign"`            // 发送者签名
	Avatar         string                            `json:"avatar"`          // 发送者头像 url
	Rank           string                            `json:"rank"`            // (?)
	DisplayRank    string                            `json:"DisplayRank"`     // (?)
	LevelInfo      GetCommentReplyItemLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetCommentReplyItemPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetCommentReplyItemNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetCommentReplyItemOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetCommentReplyItemVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetCommentReplyItemFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                               `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                               `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    GetCommentReplyItemUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                              `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                            `json:"contract_desc"`   // 合作用户说明？
}

// GetCommentReplyItemLevelInfo 是评论条目中 member.level_info 对象。
type GetCommentReplyItemLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetCommentReplyItemPendant 是评论条目中 member.pendant 对象。
type GetCommentReplyItemPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetCommentReplyItemNameplate 是评论条目中 member.nameplate 对象。
type GetCommentReplyItemNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetCommentReplyItemOfficialVerify 是评论条目中 member.official_verify 对象。
type GetCommentReplyItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetCommentReplyItemVip 是评论条目中 member.vip 对象。
type GetCommentReplyItemVip struct {
	VipType            int                         `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                         `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                      `json:"dueRemark"`            // (?)
	AccessStatus       int                         `json:"accessStatus"`         // (?)
	VipStatus          int                         `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                      `json:"vipStatusWarn"`        // (?)
	ThemeType          int                         `json:"theme_type"`           // 会员样式 id
	Label              GetCommentReplyItemVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                         `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                      `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                      `json:"nickname_color"`       // 昵称颜色
}

// GetCommentReplyItemVipLabel 是评论条目中 member.vip.label 对象。
type GetCommentReplyItemVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetCommentReplyItemFansDetail 是评论条目中 member.fans_detail 对象。
type GetCommentReplyItemFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetCommentReplyItemUserSailing 是评论条目中 member.user_sailing 对象。
type GetCommentReplyItemUserSailing struct {
	Pendant         *GetCommentReplyItemPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetCommentReplyItemCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                         `json:"cardbg_with_focus"` // (?)
}

// GetCommentReplyItemCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type GetCommentReplyItemCardbg struct {
	ID      int                    `json:"id"`       // 评论条目装扮 id
	Name    string                 `json:"name"`     // 评论条目装扮名称
	Image   string                 `json:"image"`    // 评论条目装扮图片 url
	JumpURL string                 `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetCommentReplyItemFan `json:"fan"`      // 粉丝专属信息
	Type    string                 `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetCommentReplyItemFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type GetCommentReplyItemFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetCommentReplyItemContent 是评论条目中的 content 对象。
type GetCommentReplyItemContent struct {
	Message  string                       `json:"message"`  // 评论内容。**重要**
	Plat     int                          `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                       `json:"device"`   // 评论发送平台设备
	Members  []GetCommentReplyItemMember  `json:"members"`  // at 到的用户信息
	Emote    any                          `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                          `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                          `json:"max_line"` // 6。收起最大行数
	Pictures []GetCommentReplyItemPicture `json:"pictures"` // 评论图片数组
}

// GetCommentReplyItemPicture 是评论条目中 content.pictures 数组中的对象。
type GetCommentReplyItemPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetCommentReplyItemFolder 是评论条目中的 folder 对象。
type GetCommentReplyItemFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentReplyItemUpAction 是评论条目中的 up_action 对象。
type GetCommentReplyItemUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// GetCommentReplyItemCardLabel 是评论条目中 card_label 数组中的对象。
type GetCommentReplyItemCardLabel struct {
	Rpid             int    `json:"rpid"`              // 评论 rpid
	TextContent      string `json:"text_content"`      // 标签文本。已知有妙评
	TextColorDay     string `json:"text_color_day"`    // 日间文本颜色。十六进制颜色值，下同
	TextColorNight   string `json:"text_color_night"`  // 夜间文本颜色
	LabelColorDay    string `json:"label_color_day"`   // 日间标签颜色
	LabelColorNight  string `json:"label_color_night"` // 夜间标签颜色
	Image            string `json:"image"`             // 作用不明
	Type             int    `json:"type"`              // 1。作用不明
	Background       string `json:"background"`        // 背景图片 url
	BackgroundWidth  int    `json:"background_width"`  // 背景图片宽度
	BackgroundHeight int    `json:"background_height"` // 背景图片高度
	JumpURL          string `json:"jump_url"`          // 跳转链接
	Effect           int    `json:"effect"`            // 0。作用不明，可能用于控制动画，下同
	EffectStartTime  int    `json:"effect_start_time"` // 0
}

// GetCommentReplyItemReplyControl 是评论条目中的 reply_control 对象。
type GetCommentReplyItemReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}

// GetCommentReplyUpper 是评论区的 UP 主信息。
type GetCommentReplyUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// GetCommentsHotReplyResult 是「获取评论区热评」的 data 对象。
type GetCommentsHotReplyResult struct {
	Page    GetCommentsHotReplyPage   `json:"page"`    // 页面信息
	Replies []GetCommentsHotReplyItem `json:"replies"` // 热评列表
}

// GetCommentsHotReplyPage 是评论区页码信息。
type GetCommentsHotReplyPage struct {
	Num    int `json:"num"`    // 当前页码
	Size   int `json:"size"`   // 每页项数
	Count  int `json:"count"`  // 根评论条数
	Acount int `json:"acount"` // 总计评论条数
}

// GetCommentsHotReplyItem 是热评列表中的条目。
type GetCommentsHotReplyItem struct {
	Rpid         int                                 `json:"rpid"`          // 评论 rpid
	Oid          int                                 `json:"oid"`           // 评论区对象 id
	Type         int                                 `json:"type"`          // 评论区类型代码。
	Mid          int                                 `json:"mid"`           // 发送者 mid
	Root         int                                 `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                                 `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                                 `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                                 `json:"count"`         // 二级评论条数
	Rcount       int                                 `json:"rcount"`        // 回复评论条数
	Floor        int                                 `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                                 `json:"state"`         // (?)
	FansGrade    int                                 `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                                 `json:"attr"`          // 某属性位？
	Ctime        int                                 `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                              `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                              `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                              `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                                 `json:"like"`          // 评论获赞数
	Action       int                                 `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       GetCommentsHotReplyItemMember       `json:"member"`        // 评论发送者信息
	Content      GetCommentsHotReplyItemContent      `json:"content"`       // 评论信息
	Replies      []*GetCommentsHotReplyItem          `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                                 `json:"assist"`        // (?)
	Folder       GetCommentsHotReplyItemFolder       `json:"folder"`        // 折叠信息
	UpAction     GetCommentsHotReplyItemUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                                `json:"show_follow"`   // (?)
	Invisible    bool                                `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetCommentsHotReplyItemCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetCommentsHotReplyItemReplyControl `json:"reply_control"` // 评论提示文案信息
}

// GetCommentsHotReplyItemMember 是评论条目中的 member 对象。
type GetCommentsHotReplyItemMember struct {
	Mid            string                                `json:"mid"`             // 发送者 mid
	Uname          string                                `json:"uname"`           // 发送者昵称
	Sex            string                                `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                                `json:"sign"`            // 发送者签名
	Avatar         string                                `json:"avatar"`          // 发送者头像 url
	Rank           string                                `json:"rank"`            // (?)
	DisplayRank    string                                `json:"DisplayRank"`     // (?)
	LevelInfo      GetCommentsHotReplyItemLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetCommentsHotReplyItemPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetCommentsHotReplyItemNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetCommentsHotReplyItemOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetCommentsHotReplyItemVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetCommentsHotReplyItemFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                                   `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                                   `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    GetCommentsHotReplyItemUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                                  `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                                `json:"contract_desc"`   // 合作用户说明？
}

// GetCommentsHotReplyItemLevelInfo 是评论条目中 member.level_info 对象。
type GetCommentsHotReplyItemLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetCommentsHotReplyItemPendant 是评论条目中 member.pendant 对象。
type GetCommentsHotReplyItemPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetCommentsHotReplyItemNameplate 是评论条目中 member.nameplate 对象。
type GetCommentsHotReplyItemNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetCommentsHotReplyItemOfficialVerify 是评论条目中 member.official_verify 对象。
type GetCommentsHotReplyItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetCommentsHotReplyItemVip 是评论条目中 member.vip 对象。
type GetCommentsHotReplyItemVip struct {
	VipType            int                             `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                             `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                          `json:"dueRemark"`            // (?)
	AccessStatus       int                             `json:"accessStatus"`         // (?)
	VipStatus          int                             `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                          `json:"vipStatusWarn"`        // (?)
	ThemeType          int                             `json:"theme_type"`           // 会员样式 id
	Label              GetCommentsHotReplyItemVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                             `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                          `json:"nickname_color"`       // 昵称颜色
}

// GetCommentsHotReplyItemVipLabel 是评论条目中 member.vip.label 对象。
type GetCommentsHotReplyItemVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetCommentsHotReplyItemFansDetail 是评论条目中 member.fans_detail 对象。
type GetCommentsHotReplyItemFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetCommentsHotReplyItemUserSailing 是评论条目中 member.user_sailing 对象。
type GetCommentsHotReplyItemUserSailing struct {
	Pendant         *GetCommentsHotReplyItemPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetCommentsHotReplyItemCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                             `json:"cardbg_with_focus"` // (?)
}

// GetCommentsHotReplyItemCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type GetCommentsHotReplyItemCardbg struct {
	ID      int                        `json:"id"`       // 评论条目装扮 id
	Name    string                     `json:"name"`     // 评论条目装扮名称
	Image   string                     `json:"image"`    // 评论条目装扮图片 url
	JumpURL string                     `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetCommentsHotReplyItemFan `json:"fan"`      // 粉丝专属信息
	Type    string                     `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetCommentsHotReplyItemFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type GetCommentsHotReplyItemFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetCommentsHotReplyItemContent 是评论条目中的 content 对象。
type GetCommentsHotReplyItemContent struct {
	Message  string                           `json:"message"`  // 评论内容。**重要**
	Plat     int                              `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                           `json:"device"`   // 评论发送平台设备
	Members  []GetCommentsHotReplyItemMember  `json:"members"`  // at 到的用户信息
	Emote    any                              `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                              `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                              `json:"max_line"` // 6。收起最大行数
	Pictures []GetCommentsHotReplyItemPicture `json:"pictures"` // 评论图片数组
}

// GetCommentsHotReplyItemPicture 是评论条目中 content.pictures 数组中的对象。
type GetCommentsHotReplyItemPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetCommentsHotReplyItemFolder 是评论条目中的 folder 对象。
type GetCommentsHotReplyItemFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentsHotReplyItemUpAction 是评论条目中的 up_action 对象。
type GetCommentsHotReplyItemUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// GetCommentsHotReplyItemCardLabel 是评论条目中 card_label 数组中的对象。
type GetCommentsHotReplyItemCardLabel struct {
	Rpid             int    `json:"rpid"`              // 评论 rpid
	TextContent      string `json:"text_content"`      // 标签文本。已知有妙评
	TextColorDay     string `json:"text_color_day"`    // 日间文本颜色。十六进制颜色值，下同
	TextColorNight   string `json:"text_color_night"`  // 夜间文本颜色
	LabelColorDay    string `json:"label_color_day"`   // 日间标签颜色
	LabelColorNight  string `json:"label_color_night"` // 夜间标签颜色
	Image            string `json:"image"`             // 作用不明
	Type             int    `json:"type"`              // 1。作用不明
	Background       string `json:"background"`        // 背景图片 url
	BackgroundWidth  int    `json:"background_width"`  // 背景图片宽度
	BackgroundHeight int    `json:"background_height"` // 背景图片高度
	JumpURL          string `json:"jump_url"`          // 跳转链接
	Effect           int    `json:"effect"`            // 0。作用不明，可能用于控制动画，下同
	EffectStartTime  int    `json:"effect_start_time"` // 0
}

// GetCommentsHotReplyItemReplyControl 是评论条目中的 reply_control 对象。
type GetCommentsHotReplyItemReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}
