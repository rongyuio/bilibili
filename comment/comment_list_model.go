package comment

// GetCommentsDetailLazyLoad、GetCommentDialogTree、GetCommentInfo、GetCommentCount 的响应模型。
//
// 本文件的类型只服务于上述四个接口，不与其它接口共用；即使字段相同，
// 其它接口也在各自文件中单独定义。

// GetCommentsDetailLazyLoadResult 是「获取评论区明细（懒加载）」的 data 对象。
type GetCommentsDetailLazyLoadResult struct {
	Cursor           GetCommentsDetailLazyLoadCursor   `json:"cursor"`             // 游标信息
	Hots             []*GetCommentsDetailLazyLoadReply `json:"hots"`               // 热评列表。禁用时为 null
	Notice           *GetCommentsDetailLazyLoadNotice  `json:"notice"`             // 评论区公告信息。无效时为 null
	Replies          []*GetCommentsDetailLazyLoadReply `json:"replies"`            // 评论列表。禁用时为 null
	Top              GetCommentsDetailLazyLoadTop      `json:"top"`                // 置顶信息
	TopReplies       []*GetCommentsDetailLazyLoadReply `json:"top_replies"`        // 置顶评论
	LotteryCard      any                               `json:"lottery_card"`       // 抽奖评论。实测为 null
	Folder           GetCommentsDetailLazyLoadFolder   `json:"folder"`             // 评论折叠信息
	UpSelection      any                               `json:"up_selection"`       // (?) 文档未给出对象字段表
	Cm               any                               `json:"cm"`                 // 广告。文档未给出对象字段表
	CmInfo           any                               `json:"cm_info"`            // 广告控制。文档未给出对象字段表
	Effects          any                               `json:"effects"`            // (?)。文档未给出对象字段表
	Assist           int                               `json:"assist"`             // (?)
	Blacklist        int                               `json:"blacklist"`          // (?)
	Vote             int                               `json:"vote"`               // (?)
	Lottery          int                               `json:"lottery"`            // (?)
	Config           GetCommentsDetailLazyLoadConfig   `json:"config"`             // 评论区显示控制
	Upper            GetCommentsDetailLazyLoadUpper    `json:"upper"`              // UP 主信息
	ShowBvid         bool                              `json:"show_bvid"`          // 显示 bvid?
	Control          GetCommentsDetailLazyLoadControl  `json:"control"`            // 评论区输入属性
	Note             int                               `json:"note"`               // 1
	EsportsGradeCard any                               `json:"esports_grade_card"` // 实测为 null
	Callbacks        any                               `json:"callbacks"`          // 实测为 null
	ContextFeature   string                            `json:"context_feature"`    // (?)
}

// GetCommentsDetailLazyLoadCursor 是评论区懒加载的游标信息。
type GetCommentsDetailLazyLoadCursor struct {
	AllCount        int                                       `json:"all_count"`        // 全部评论条数
	IsBegin         bool                                      `json:"is_begin"`         // 是否为第一页。false：否。true：是
	Prev            int                                       `json:"prev"`             // 上页页码
	Next            int                                       `json:"next"`             // 下页页码
	IsEnd           bool                                      `json:"is_end"`           // 是否为最后页。false：否。true：是
	Mode            int                                       `json:"mode"`             // 排序方式
	SupportMode     []int                                     `json:"support_mode"`     // 支持的排序方式
	Name            string                                    `json:"name"`             // 评论区类型名
	PaginationReply *GetCommentsDetailLazyLoadPaginationReply `json:"pagination_reply"` // 用于下一次请求的偏移信息
	SessionID       string                                    `json:"session_id"`       // 空
}

// GetCommentsDetailLazyLoadPaginationReply 是评论区懒加载的翻页偏移信息。
type GetCommentsDetailLazyLoadPaginationReply struct {
	NextOffset string `json:"next_offset"` // 用于下一次请求的偏移信息
	PrevOffset string `json:"prev_offset"` // 用于本次请求的偏移信息
}

// GetCommentsDetailLazyLoadTop 是评论区置顶信息。
type GetCommentsDetailLazyLoadTop struct {
	Admin *GetCommentsDetailLazyLoadReply `json:"admin"` // 管理员置顶条目。无效时为 null
	Upper *GetCommentsDetailLazyLoadReply `json:"upper"` // UP 主置顶条目。无效时为 null
	Vote  *GetCommentsDetailLazyLoadReply `json:"vote"`  // 投票置顶条目。无效时为 null
}

// GetCommentsDetailLazyLoadConfig 是评论区显示控制。
type GetCommentsDetailLazyLoadConfig struct {
	ShowAdmin  int  `json:"showadmin"`    // 是否显示管理置顶
	ShowEntry  int  `json:"showentry"`    // (?)
	ShowFloor  int  `json:"showfloor"`    // 是否显示楼层号
	ShowTopic  int  `json:"showtopic"`    // 是否显示话题
	ShowUpFlag bool `json:"show_up_flag"` // 是否显示“UP 觉得很赞”标志
	ReadOnly   bool `json:"read_only"`    // 是否只读评论区
	ShowDelLog bool `json:"show_del_log"` // 是否显示删除记录
}

// GetCommentsDetailLazyLoadNotice 是评论区公告信息。
type GetCommentsDetailLazyLoadNotice struct {
	Content string `json:"content"` // 公告正文
	ID      int    `json:"id"`      // 公告 id
	Link    string `json:"link"`    // 公告页面链接 url
	Title   string `json:"title"`   // 公告标题
}

// GetCommentsDetailLazyLoadFolder 是折叠相关信息。
type GetCommentsDetailLazyLoadFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentsDetailLazyLoadUpper 是评论区的 UP 主信息。
type GetCommentsDetailLazyLoadUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// GetCommentsDetailLazyLoadControl 是评论区输入属性。
type GetCommentsDetailLazyLoadControl struct {
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

// GetCommentsDetailLazyLoadReply 是懒加载评论列表中的条目。
type GetCommentsDetailLazyLoadReply struct {
	Rpid         int                                        `json:"rpid"`          // 评论 rpid
	Oid          int                                        `json:"oid"`           // 评论区对象 id
	Type         int                                        `json:"type"`          // 评论区类型代码。
	Mid          int                                        `json:"mid"`           // 发送者 mid
	Root         int                                        `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                                        `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                                        `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                                        `json:"count"`         // 二级评论条数
	Rcount       int                                        `json:"rcount"`        // 回复评论条数
	Floor        int                                        `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                                        `json:"state"`         // (?)
	FansGrade    int                                        `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                                        `json:"attr"`          // 某属性位？
	Ctime        int                                        `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                                     `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                                     `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                                     `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                                        `json:"like"`          // 评论获赞数
	Action       int                                        `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       GetCommentsDetailLazyLoadReplyMember       `json:"member"`        // 评论发送者信息
	Content      GetCommentsDetailLazyLoadReplyContent      `json:"content"`       // 评论信息
	Replies      []*GetCommentsDetailLazyLoadReply          `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                                        `json:"assist"`        // (?)
	Folder       GetCommentsDetailLazyLoadReplyFolder       `json:"folder"`        // 折叠信息
	UpAction     GetCommentsDetailLazyLoadReplyUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                                       `json:"show_follow"`   // (?)
	Invisible    bool                                       `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetCommentsDetailLazyLoadReplyCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetCommentsDetailLazyLoadReplyReplyControl `json:"reply_control"` // 评论提示文案信息
}

// GetCommentsDetailLazyLoadReplyMember 是评论条目中的 member 对象。
type GetCommentsDetailLazyLoadReplyMember struct {
	Mid            string                                       `json:"mid"`             // 发送者 mid
	Uname          string                                       `json:"uname"`           // 发送者昵称
	Sex            string                                       `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                                       `json:"sign"`            // 发送者签名
	Avatar         string                                       `json:"avatar"`          // 发送者头像 url
	Rank           string                                       `json:"rank"`            // (?)
	DisplayRank    string                                       `json:"DisplayRank"`     // (?)
	LevelInfo      GetCommentsDetailLazyLoadReplyLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetCommentsDetailLazyLoadReplyPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetCommentsDetailLazyLoadReplyNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetCommentsDetailLazyLoadReplyOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetCommentsDetailLazyLoadReplyVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetCommentsDetailLazyLoadReplyFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                                          `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                                          `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    GetCommentsDetailLazyLoadReplyUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                                         `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                                       `json:"contract_desc"`   // 合作用户说明？
}

// GetCommentsDetailLazyLoadReplyLevelInfo 是评论条目中 member.level_info 对象。
type GetCommentsDetailLazyLoadReplyLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetCommentsDetailLazyLoadReplyPendant 是评论条目中 member.pendant 对象。
type GetCommentsDetailLazyLoadReplyPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetCommentsDetailLazyLoadReplyNameplate 是评论条目中 member.nameplate 对象。
type GetCommentsDetailLazyLoadReplyNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetCommentsDetailLazyLoadReplyOfficialVerify 是评论条目中 member.official_verify 对象。
type GetCommentsDetailLazyLoadReplyOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetCommentsDetailLazyLoadReplyVip 是评论条目中 member.vip 对象。
type GetCommentsDetailLazyLoadReplyVip struct {
	VipType            int                                    `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                                    `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                                 `json:"dueRemark"`            // (?)
	AccessStatus       int                                    `json:"accessStatus"`         // (?)
	VipStatus          int                                    `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                                 `json:"vipStatusWarn"`        // (?)
	ThemeType          int                                    `json:"theme_type"`           // 会员样式 id
	Label              GetCommentsDetailLazyLoadReplyVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                                    `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                                 `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                                 `json:"nickname_color"`       // 昵称颜色
}

// GetCommentsDetailLazyLoadReplyVipLabel 是评论条目中 member.vip.label 对象。
type GetCommentsDetailLazyLoadReplyVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetCommentsDetailLazyLoadReplyFansDetail 是评论条目中 member.fans_detail 对象。
type GetCommentsDetailLazyLoadReplyFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetCommentsDetailLazyLoadReplyUserSailing 是评论条目中 member.user_sailing 对象。
type GetCommentsDetailLazyLoadReplyUserSailing struct {
	Pendant         *GetCommentsDetailLazyLoadReplyPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetCommentsDetailLazyLoadReplyCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                                    `json:"cardbg_with_focus"` // (?)
}

// GetCommentsDetailLazyLoadReplyCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type GetCommentsDetailLazyLoadReplyCardbg struct {
	ID      int                               `json:"id"`       // 评论条目装扮 id
	Name    string                            `json:"name"`     // 评论条目装扮名称
	Image   string                            `json:"image"`    // 评论条目装扮图片 url
	JumpURL string                            `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetCommentsDetailLazyLoadReplyFan `json:"fan"`      // 粉丝专属信息
	Type    string                            `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetCommentsDetailLazyLoadReplyFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type GetCommentsDetailLazyLoadReplyFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetCommentsDetailLazyLoadReplyContent 是评论条目中的 content 对象。
type GetCommentsDetailLazyLoadReplyContent struct {
	Message  string                                  `json:"message"`  // 评论内容。**重要**
	Plat     int                                     `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                                  `json:"device"`   // 评论发送平台设备
	Members  []GetCommentsDetailLazyLoadReplyMember  `json:"members"`  // at 到的用户信息
	Emote    any                                     `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                                     `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                                     `json:"max_line"` // 6。收起最大行数
	Pictures []GetCommentsDetailLazyLoadReplyPicture `json:"pictures"` // 评论图片数组
}

// GetCommentsDetailLazyLoadReplyPicture 是评论条目中 content.pictures 数组中的对象。
type GetCommentsDetailLazyLoadReplyPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetCommentsDetailLazyLoadReplyFolder 是评论条目中的 folder 对象。
type GetCommentsDetailLazyLoadReplyFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentsDetailLazyLoadReplyUpAction 是评论条目中的 up_action 对象。
type GetCommentsDetailLazyLoadReplyUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// GetCommentsDetailLazyLoadReplyCardLabel 是评论条目中 card_label 数组中的对象。
type GetCommentsDetailLazyLoadReplyCardLabel struct {
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

// GetCommentsDetailLazyLoadReplyReplyControl 是评论条目中的 reply_control 对象。
type GetCommentsDetailLazyLoadReplyReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}

// GetCommentDialogTreeResult 是「获取指定评论对话树」的 data 对象。
type GetCommentDialogTreeResult struct {
	Cursor    GetCommentDialogTreeCursor   `json:"cursor"`    // 本页楼层信息
	Dialog    GetCommentDialogTreeDialog   `json:"dialog"`    // 对话楼层信息
	Replies   []*GetCommentDialogTreeReply `json:"replies"`   // 评论对话树列表
	Assist    int                          `json:"assist"`    // (?)
	Blacklist int                          `json:"blacklist"` // (?)
	Vote      int                          `json:"vote"`      // (?)
	Lottery   int                          `json:"lottery"`   // (?)
	Config    GetCommentDialogTreeConfig   `json:"config"`    // 评论区显示控制
	Upper     GetCommentDialogTreeUpper    `json:"upper"`     // UP 主 mid
	ShowBvid  bool                         `json:"show_bvid"` // 显示 bvid?
	Control   GetCommentDialogTreeControl  `json:"control"`   // 评论区输入属性
	Note      int                          `json:"note"`      // (?)
}

// GetCommentDialogTreeCursor 是评论对话树的本页楼层信息。
type GetCommentDialogTreeCursor struct {
	MinFloor int `json:"min_floor"` // 本页最低对话楼层
	MaxFloor int `json:"max_floor"` // 本页最高对话楼层
	Size     int `json:"size"`      // 本页项数
}

// GetCommentDialogTreeDialog 是评论对话树的二级评论楼层信息。
type GetCommentDialogTreeDialog struct {
	MinFloor int `json:"min_floor"` // 二级评论最低对话楼层
	MaxFloor int `json:"max_floor"` // 二级评论最高对话楼层
}

// GetCommentDialogTreeConfig 是评论区显示控制。
type GetCommentDialogTreeConfig struct {
	ShowAdmin  int  `json:"showadmin"`    // 是否显示管理置顶
	ShowEntry  int  `json:"showentry"`    // (?)
	ShowFloor  int  `json:"showfloor"`    // 是否显示楼层号
	ShowTopic  int  `json:"showtopic"`    // 是否显示话题
	ShowUpFlag bool `json:"show_up_flag"` // 是否显示“UP 觉得很赞”标志
	ReadOnly   bool `json:"read_only"`    // 是否只读评论区
	ShowDelLog bool `json:"show_del_log"` // 是否显示删除记录
}

// GetCommentDialogTreeUpper 是评论区的 UP 主信息。
type GetCommentDialogTreeUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// GetCommentDialogTreeControl 是评论区输入属性。
type GetCommentDialogTreeControl struct {
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

// GetCommentDialogTreeReply 是评论对话树中的条目。
type GetCommentDialogTreeReply struct {
	Rpid         int                                   `json:"rpid"`          // 评论 rpid
	Oid          int                                   `json:"oid"`           // 评论区对象 id
	Type         int                                   `json:"type"`          // 评论区类型代码。
	Mid          int                                   `json:"mid"`           // 发送者 mid
	Root         int                                   `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                                   `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                                   `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                                   `json:"count"`         // 二级评论条数
	Rcount       int                                   `json:"rcount"`        // 回复评论条数
	Floor        int                                   `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                                   `json:"state"`         // (?)
	FansGrade    int                                   `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                                   `json:"attr"`          // 某属性位？
	Ctime        int                                   `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                                `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                                `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                                `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                                   `json:"like"`          // 评论获赞数
	Action       int                                   `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       GetCommentDialogTreeReplyMember       `json:"member"`        // 评论发送者信息
	Content      GetCommentDialogTreeReplyContent      `json:"content"`       // 评论信息
	Replies      []*GetCommentDialogTreeReply          `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                                   `json:"assist"`        // (?)
	Folder       GetCommentDialogTreeReplyFolder       `json:"folder"`        // 折叠信息
	UpAction     GetCommentDialogTreeReplyUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                                  `json:"show_follow"`   // (?)
	Invisible    bool                                  `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetCommentDialogTreeReplyCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetCommentDialogTreeReplyReplyControl `json:"reply_control"` // 评论提示文案信息
}

// GetCommentDialogTreeReplyMember 是评论条目中的 member 对象。
type GetCommentDialogTreeReplyMember struct {
	Mid            string                                  `json:"mid"`             // 发送者 mid
	Uname          string                                  `json:"uname"`           // 发送者昵称
	Sex            string                                  `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                                  `json:"sign"`            // 发送者签名
	Avatar         string                                  `json:"avatar"`          // 发送者头像 url
	Rank           string                                  `json:"rank"`            // (?)
	DisplayRank    string                                  `json:"DisplayRank"`     // (?)
	LevelInfo      GetCommentDialogTreeReplyLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetCommentDialogTreeReplyPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetCommentDialogTreeReplyNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetCommentDialogTreeReplyOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetCommentDialogTreeReplyVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetCommentDialogTreeReplyFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                                     `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                                     `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    GetCommentDialogTreeReplyUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                                    `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                                  `json:"contract_desc"`   // 合作用户说明？
}

// GetCommentDialogTreeReplyLevelInfo 是评论条目中 member.level_info 对象。
type GetCommentDialogTreeReplyLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetCommentDialogTreeReplyPendant 是评论条目中 member.pendant 对象。
type GetCommentDialogTreeReplyPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetCommentDialogTreeReplyNameplate 是评论条目中 member.nameplate 对象。
type GetCommentDialogTreeReplyNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetCommentDialogTreeReplyOfficialVerify 是评论条目中 member.official_verify 对象。
type GetCommentDialogTreeReplyOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetCommentDialogTreeReplyVip 是评论条目中 member.vip 对象。
type GetCommentDialogTreeReplyVip struct {
	VipType            int                               `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                               `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                            `json:"dueRemark"`            // (?)
	AccessStatus       int                               `json:"accessStatus"`         // (?)
	VipStatus          int                               `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                            `json:"vipStatusWarn"`        // (?)
	ThemeType          int                               `json:"theme_type"`           // 会员样式 id
	Label              GetCommentDialogTreeReplyVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                               `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                            `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                            `json:"nickname_color"`       // 昵称颜色
}

// GetCommentDialogTreeReplyVipLabel 是评论条目中 member.vip.label 对象。
type GetCommentDialogTreeReplyVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetCommentDialogTreeReplyFansDetail 是评论条目中 member.fans_detail 对象。
type GetCommentDialogTreeReplyFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetCommentDialogTreeReplyUserSailing 是评论条目中 member.user_sailing 对象。
type GetCommentDialogTreeReplyUserSailing struct {
	Pendant         *GetCommentDialogTreeReplyPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetCommentDialogTreeReplyCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                               `json:"cardbg_with_focus"` // (?)
}

// GetCommentDialogTreeReplyCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type GetCommentDialogTreeReplyCardbg struct {
	ID      int                          `json:"id"`       // 评论条目装扮 id
	Name    string                       `json:"name"`     // 评论条目装扮名称
	Image   string                       `json:"image"`    // 评论条目装扮图片 url
	JumpURL string                       `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetCommentDialogTreeReplyFan `json:"fan"`      // 粉丝专属信息
	Type    string                       `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetCommentDialogTreeReplyFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type GetCommentDialogTreeReplyFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetCommentDialogTreeReplyContent 是评论条目中的 content 对象。
type GetCommentDialogTreeReplyContent struct {
	Message  string                             `json:"message"`  // 评论内容。**重要**
	Plat     int                                `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                             `json:"device"`   // 评论发送平台设备
	Members  []GetCommentDialogTreeReplyMember  `json:"members"`  // at 到的用户信息
	Emote    any                                `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                                `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                                `json:"max_line"` // 6。收起最大行数
	Pictures []GetCommentDialogTreeReplyPicture `json:"pictures"` // 评论图片数组
}

// GetCommentDialogTreeReplyPicture 是评论条目中 content.pictures 数组中的对象。
type GetCommentDialogTreeReplyPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetCommentDialogTreeReplyFolder 是评论条目中的 folder 对象。
type GetCommentDialogTreeReplyFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentDialogTreeReplyUpAction 是评论条目中的 up_action 对象。
type GetCommentDialogTreeReplyUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// GetCommentDialogTreeReplyCardLabel 是评论条目中 card_label 数组中的对象。
type GetCommentDialogTreeReplyCardLabel struct {
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

// GetCommentDialogTreeReplyReplyControl 是评论条目中的 reply_control 对象。
type GetCommentDialogTreeReplyReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}

// GetCommentInfoResult 是「获取指定评论信息」的 data 对象，即一条评论条目。
//
// 该接口已经弃用，仅保留以兼容旧调用。
type GetCommentInfoResult struct {
	Rpid         int                        `json:"rpid"`          // 评论 rpid
	Oid          int                        `json:"oid"`           // 评论区对象 id
	Type         int                        `json:"type"`          // 评论区类型代码。
	Mid          int                        `json:"mid"`           // 发送者 mid
	Root         int                        `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                        `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                        `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                        `json:"count"`         // 二级评论条数
	Rcount       int                        `json:"rcount"`        // 回复评论条数
	Floor        int                        `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                        `json:"state"`         // (?)
	FansGrade    int                        `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                        `json:"attr"`          // 某属性位？
	Ctime        int                        `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                     `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                     `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                     `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                        `json:"like"`          // 评论获赞数
	Action       int                        `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       GetCommentInfoMember       `json:"member"`        // 评论发送者信息
	Content      GetCommentInfoContent      `json:"content"`       // 评论信息
	Replies      []*GetCommentInfoResult    `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                        `json:"assist"`        // (?)
	Folder       GetCommentInfoFolder       `json:"folder"`        // 折叠信息
	UpAction     GetCommentInfoUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                       `json:"show_follow"`   // (?)
	Invisible    bool                       `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetCommentInfoCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetCommentInfoReplyControl `json:"reply_control"` // 评论提示文案信息
}

// GetCommentInfoMember 是评论条目中的 member 对象。
type GetCommentInfoMember struct {
	Mid            string                       `json:"mid"`             // 发送者 mid
	Uname          string                       `json:"uname"`           // 发送者昵称
	Sex            string                       `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                       `json:"sign"`            // 发送者签名
	Avatar         string                       `json:"avatar"`          // 发送者头像 url
	Rank           string                       `json:"rank"`            // (?)
	DisplayRank    string                       `json:"DisplayRank"`     // (?)
	LevelInfo      GetCommentInfoLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetCommentInfoPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetCommentInfoNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetCommentInfoOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetCommentInfoVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetCommentInfoFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                          `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                          `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    GetCommentInfoUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                         `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                       `json:"contract_desc"`   // 合作用户说明？
}

// GetCommentInfoLevelInfo 是评论条目中 member.level_info 对象。
type GetCommentInfoLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetCommentInfoPendant 是评论条目中 member.pendant 对象。
type GetCommentInfoPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetCommentInfoNameplate 是评论条目中 member.nameplate 对象。
type GetCommentInfoNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetCommentInfoOfficialVerify 是评论条目中 member.official_verify 对象。
type GetCommentInfoOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetCommentInfoVip 是评论条目中 member.vip 对象。
type GetCommentInfoVip struct {
	VipType            int                    `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                    `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                 `json:"dueRemark"`            // (?)
	AccessStatus       int                    `json:"accessStatus"`         // (?)
	VipStatus          int                    `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                 `json:"vipStatusWarn"`        // (?)
	ThemeType          int                    `json:"theme_type"`           // 会员样式 id
	Label              GetCommentInfoVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                    `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                 `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                 `json:"nickname_color"`       // 昵称颜色
}

// GetCommentInfoVipLabel 是评论条目中 member.vip.label 对象。
type GetCommentInfoVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetCommentInfoFansDetail 是评论条目中 member.fans_detail 对象。
type GetCommentInfoFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetCommentInfoUserSailing 是评论条目中 member.user_sailing 对象。
type GetCommentInfoUserSailing struct {
	Pendant         *GetCommentInfoPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetCommentInfoCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                    `json:"cardbg_with_focus"` // (?)
}

// GetCommentInfoCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type GetCommentInfoCardbg struct {
	ID      int               `json:"id"`       // 评论条目装扮 id
	Name    string            `json:"name"`     // 评论条目装扮名称
	Image   string            `json:"image"`    // 评论条目装扮图片 url
	JumpURL string            `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetCommentInfoFan `json:"fan"`      // 粉丝专属信息
	Type    string            `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetCommentInfoFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type GetCommentInfoFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetCommentInfoContent 是评论条目中的 content 对象。
type GetCommentInfoContent struct {
	Message  string                  `json:"message"`  // 评论内容。**重要**
	Plat     int                     `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                  `json:"device"`   // 评论发送平台设备
	Members  []GetCommentInfoMember  `json:"members"`  // at 到的用户信息
	Emote    any                     `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                     `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                     `json:"max_line"` // 6。收起最大行数
	Pictures []GetCommentInfoPicture `json:"pictures"` // 评论图片数组
}

// GetCommentInfoPicture 是评论条目中 content.pictures 数组中的对象。
type GetCommentInfoPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetCommentInfoFolder 是评论条目中的 folder 对象。
type GetCommentInfoFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetCommentInfoUpAction 是评论条目中的 up_action 对象。
type GetCommentInfoUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// GetCommentInfoCardLabel 是评论条目中 card_label 数组中的对象。
type GetCommentInfoCardLabel struct {
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

// GetCommentInfoReplyControl 是评论条目中的 reply_control 对象。
type GetCommentInfoReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}

// GetCommentCountResult 是「获取评论区评论总数」的 data 对象。
type GetCommentCountResult struct {
	Count int `json:"count"` // 评论条数
}
