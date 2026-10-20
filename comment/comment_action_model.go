package comment

// AddComment 的响应模型。
//
// 本文件的类型只服务于 AddComment，不与其它接口共用；即使字段相同，
// 其它接口也在各自文件中单独定义。

// AddCommentResult 是「发表评论」的 data 对象。
type AddCommentResult struct {
	SuccessAction int                        `json:"success_action"` // 作用尚不明确
	SuccessToast  string                     `json:"success_toast"`  // 状态文字
	NeedCaptcha   bool                       `json:"need_captcha"`   // 评论需要验证码（未证实）
	URL           string                     `json:"url"`            // 作用尚不明确
	Rpid          int                        `json:"rpid"`           // 评论 rpid
	RpidStr       string                     `json:"rpid_str"`       // 评论 rpid。字串格式
	Dialog        int                        `json:"dialog"`         // 回复对方 rpid。一级评论为 0，二级评论为该评论 id，大于二级评论为上一级评论 id
	DialogStr     string                     `json:"dialog_str"`     // 回复对方 rpid。字串格式
	Root          int                        `json:"root"`           // 根评论 rpid。一级评论为 0，大于一级评论为根评论 id
	RootStr       string                     `json:"root_str"`       // 根评论 rpid。字串格式
	Parent        int                        `json:"parent"`         // 回复父评论 rpid。一级评论为 0，二级评论为根评论 id，大于二级评论为上一级评论 id
	ParentStr     string                     `json:"parent_str"`     // 回复父评论 rpid。字串格式
	Emote         map[string]AddCommentEmote `json:"emote"`          // 表情转义符信息
	Reply         *AddCommentReply           `json:"reply"`          // 评论条目。有效时为对象，无效时为 null
}

// AddCommentEmote 是发表评论结果中的表情转义符信息。
type AddCommentEmote struct {
	ID        int                 `json:"id"`         // 表情 id
	PackageID int                 `json:"package_id"` // 表情包 id
	State     int                 `json:"state"`      // 0
	Type      int                 `json:"type"`       // 表情类型。1：免费。2：会员专属。3：购买所得。4：颜文字
	Attr      int                 `json:"attr"`       // (?)
	Text      string              `json:"text"`       // 表情转义符
	URL       string              `json:"url"`        // 表情图片 url
	Meta      AddCommentEmoteMeta `json:"meta"`       // 属性信息
	MTime     int64               `json:"mtime"`      // 表情创建时间。时间戳
	JumpTitle string              `json:"jump_title"` // 表情名称
}

// AddCommentEmoteMeta 是发表评论结果中表情的属性信息。
type AddCommentEmoteMeta struct {
	Size    int      `json:"size"`    // 表情尺寸信息。1：小。2：大
	Suggest []string `json:"suggest"` // 联想词？实测为 [""]，作用尚不明确
	Alias   string   `json:"alias"`   // 简写名。无则无此项
}

// AddCommentReply 是发表评论结果中的评论条目。
type AddCommentReply struct {
	Rpid         int                         `json:"rpid"`          // 评论 rpid
	Oid          int                         `json:"oid"`           // 评论区对象 id
	Type         int                         `json:"type"`          // 评论区类型代码。
	Mid          int                         `json:"mid"`           // 发送者 mid
	Root         int                         `json:"root"`          // 根评论 rpid。若为一级评论则为 0。大于一级评论则为根评论 id
	Parent       int                         `json:"parent"`        // 回复父评论 rpid。若为一级评论则为 0。若为二级评论则为根评论 rpid。大于二级评论为上一级评 论 rpid
	Dialog       int                         `json:"dialog"`        // 回复对方 rpid。若为一级评论则为 0。若为二级评论则为该评论 rpid。大于二级评论为上一级评论 rpid
	Count        int                         `json:"count"`         // 二级评论条数
	Rcount       int                         `json:"rcount"`        // 回复评论条数
	Floor        int                         `json:"floor"`         // 评论楼层号。**注：若不支持楼层则无此项**
	State        int                         `json:"state"`         // (?)
	FansGrade    int                         `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                         `json:"attr"`          // 某属性位？
	Ctime        int                         `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                      `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                      `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                      `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                         `json:"like"`          // 评论获赞数
	Action       int                         `json:"action"`        // 当前用户操作状态。需要登录(Cookie 或 APP) 。否则恒为 0。0：无。1：已点赞。2：已点踩
	Member       AddCommentReplyMember       `json:"member"`        // 评论发送者信息
	Content      AddCommentReplyContent      `json:"content"`       // 评论信息
	Replies      []*AddCommentReply          `json:"replies"`       // 评论回复条目预览。**仅嵌套一层**。否则为 null
	Assist       int                         `json:"assist"`        // (?)
	Folder       AddCommentReplyFolder       `json:"folder"`        // 折叠信息
	UpAction     AddCommentReplyUpAction     `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                        `json:"show_follow"`   // (?)
	Invisible    bool                        `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []AddCommentReplyCardLabel  `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl AddCommentReplyReplyControl `json:"reply_control"` // 评论提示文案信息
}

// AddCommentReplyMember 是评论条目中的 member 对象。
type AddCommentReplyMember struct {
	Mid            string                        `json:"mid"`             // 发送者 mid
	Uname          string                        `json:"uname"`           // 发送者昵称
	Sex            string                        `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                        `json:"sign"`            // 发送者签名
	Avatar         string                        `json:"avatar"`          // 发送者头像 url
	Rank           string                        `json:"rank"`            // (?)
	DisplayRank    string                        `json:"DisplayRank"`     // (?)
	LevelInfo      AddCommentReplyLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        AddCommentReplyPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      AddCommentReplyNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify AddCommentReplyOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            AddCommentReplyVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *AddCommentReplyFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                           `json:"following"`       // 是否关注该用户。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	IsFollowed     int                           `json:"is_followed"`     // 是否被该用户关注。需要登录(Cookie或APP) 。否则恒为0。0：未关注。1：已关注
	UserSailing    AddCommentReplyUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                          `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                        `json:"contract_desc"`   // 合作用户说明？
}

// AddCommentReplyLevelInfo 是评论条目中 member.level_info 对象。
type AddCommentReplyLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// AddCommentReplyPendant 是评论条目中 member.pendant 对象。
type AddCommentReplyPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// AddCommentReplyNameplate 是评论条目中 member.nameplate 对象。
type AddCommentReplyNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// AddCommentReplyOfficialVerify 是评论条目中 member.official_verify 对象。
type AddCommentReplyOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// AddCommentReplyVip 是评论条目中 member.vip 对象。
type AddCommentReplyVip struct {
	VipType            int                     `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                     `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                  `json:"dueRemark"`            // (?)
	AccessStatus       int                     `json:"accessStatus"`         // (?)
	VipStatus          int                     `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                  `json:"vipStatusWarn"`        // (?)
	ThemeType          int                     `json:"theme_type"`           // 会员样式 id
	Label              AddCommentReplyVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                     `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                  `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                  `json:"nickname_color"`       // 昵称颜色
}

// AddCommentReplyVipLabel 是评论条目中 member.vip.label 对象。
type AddCommentReplyVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// AddCommentReplyFansDetail 是评论条目中 member.fans_detail 对象。
type AddCommentReplyFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// AddCommentReplyUserSailing 是评论条目中 member.user_sailing 对象。
type AddCommentReplyUserSailing struct {
	Pendant         *AddCommentReplyPendant `json:"pendant"`           // 头像框信息
	Cardbg          *AddCommentReplyCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                     `json:"cardbg_with_focus"` // (?)
}

// AddCommentReplyCardbg 是评论条目中 member.user_sailing.cardbg 对象。
type AddCommentReplyCardbg struct {
	ID      int                `json:"id"`       // 评论条目装扮 id
	Name    string             `json:"name"`     // 评论条目装扮名称
	Image   string             `json:"image"`    // 评论条目装扮图片 url
	JumpURL string             `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     AddCommentReplyFan `json:"fan"`      // 粉丝专属信息
	Type    string             `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// AddCommentReplyFan 是评论条目中 member.user_sailing.cardbg.fan 对象。
type AddCommentReplyFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// AddCommentReplyContent 是评论条目中的 content 对象。
type AddCommentReplyContent struct {
	Message  string                   `json:"message"`  // 评论内容。**重要**
	Plat     int                      `json:"plat"`     // 评论发送端。1：web端。2：安卓客户端。3：ios 客户端。4：wp 客户端
	Device   string                   `json:"device"`   // 评论发送平台设备
	Members  []AddCommentReplyMember  `json:"members"`  // at 到的用户信息
	Emote    any                      `json:"emote"`    // 需要渲染的表情转义。评论内容无表情则无此项
	JumpURL  any                      `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                      `json:"max_line"` // 6。收起最大行数
	Pictures []AddCommentReplyPicture `json:"pictures"` // 评论图片数组
}

// AddCommentReplyPicture 是评论条目中 content.pictures 数组中的对象。
type AddCommentReplyPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// AddCommentReplyFolder 是评论条目中的 folder 对象。
type AddCommentReplyFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// AddCommentReplyUpAction 是评论条目中的 up_action 对象。
type AddCommentReplyUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞。false：否。true：是
	Reply bool `json:"reply"` // 是否被UP主回复。false：否。true：是
}

// AddCommentReplyCardLabel 是评论条目中 card_label 数组中的对象。
type AddCommentReplyCardLabel struct {
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

// AddCommentReplyReplyControl 是评论条目中的 reply_control 对象。
type AddCommentReplyReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示。xx 天/小时 前发布
	Location          string `json:"location"`             // IP属地。IP属地：xx。评论者发送评论时的IP地址属地。仅对2022-07-25 11:00及以后发布的评论有效。需要登录
}
