package note

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频笔记列表相关接口。各接口的返回体独立定义，互不复用。

// GetArchivePrivateNotesParam 是查询稿件私有笔记的请求参数。
type GetArchivePrivateNotesParam struct {
	Oid     int `json:"oid"`      // 目标 id
	OidType int `json:"oid_type"` // 目标 id 类型。0：视频（oid 为 avid）
}

// GetArchivePrivateNotesResult 是查询稿件私有笔记的 data 对象。
type GetArchivePrivateNotesResult struct {
	NoteIDs []string `json:"noteIds"` // 笔记 id 列表。无笔记则无此项
}

// GetArchivePrivateNotes 查询稿件私有笔记，返回该稿件下的笔记 id 列表
//
// 该接口只能查询私有笔记，无法查询公开笔记。需要 Cookie（SESSDATA）。
func (c *Client) GetArchivePrivateNotes(ctx context.Context, param GetArchivePrivateNotesParam) (*GetArchivePrivateNotesResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/list/archive"
	)
	return bilibili.Execute[*GetArchivePrivateNotesResult](ctx, c.Client, method, url, param)
}

// GetUserPrivateNotesParam 是查询当前用户私有笔记列表的请求参数。
type GetUserPrivateNotesParam struct {
	Ps int `json:"ps"` // 每页项数
	Pn int `json:"pn"` // 页码
}

// GetUserPrivateNotesResult 是查询用户私有笔记的 data 对象。
type GetUserPrivateNotesResult struct {
	List []GetUserPrivateNotesItem `json:"list"` // 明细列表，按创建顺序排列
	Page GetUserPrivateNotesPage   `json:"page"` // 页面信息
}

// GetUserPrivateNotesItem 是私有笔记列表中的一项。
type GetUserPrivateNotesItem struct {
	Title              string                 `json:"title"`                // 笔记标题
	Summary            string                 `json:"summary"`              // 笔记预览文本
	Mtime              string                 `json:"mtime"`                // 提交时间。格式 YYYY-MM-DD hh:mm
	Arc                GetUserPrivateNotesArc `json:"arc"`                  // 视频信息
	NoteID             int64                  `json:"note_id"`              // 笔记 id
	AuditStatus        int                    `json:"audit_status"`         // 0。作用尚不明确
	WebURL             string                 `json:"web_url"`              // 笔记 h5 页 url
	NoteIDStr          string                 `json:"note_id_str"`          // 笔记 id 字符串
	Message            string                 `json:"message"`              // 更新信息。"更新于 YYYY-MM-DD hh:mm"
	ForbidNoteEntrance bool                   `json:"forbid_note_entrance"` // (?)。作用尚不明确
	Likes              int                    `json:"likes"`                // 点赞数
	HasLike            bool                   `json:"has_like"`             // 是否点赞
}

// GetUserPrivateNotesArc 是私有笔记列表项关联稿件的视频信息。
type GetUserPrivateNotesArc struct {
	Oid     int    `json:"oid"`      // 目标 id
	Bvid    string `json:"bvid"`     // 稿件 bvid
	Pic     string `json:"pic"`      // 稿件封面图片 url
	Desc    string `json:"desc"`     // 视频简介
	Status  int    `json:"status"`   // 0
	OidType int    `json:"oid_type"` // 目标 id 类型。0：视频（oid 为 avid）
	Aid     int    `json:"aid"`      // 稿件 avid
}

// GetUserPrivateNotesPage 是私有笔记列表的分页信息。
type GetUserPrivateNotesPage struct {
	Total int `json:"total"` // 笔记总数
	Size  int `json:"size"`  // 每页项数
	Num   int `json:"num"`   // 页码
}

// GetUserPrivateNotes 查询当前用户的私有笔记列表
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetUserPrivateNotes(ctx context.Context, param GetUserPrivateNotesParam) (*GetUserPrivateNotesResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/list"
	)
	return bilibili.Execute[*GetUserPrivateNotesResult](ctx, c.Client, method, url, param)
}

// GetArchivePublicNotesParam 是查询稿件公开笔记列表的请求参数。
type GetArchivePublicNotesParam struct {
	Oid     int `json:"oid"`      // 目标 id
	OidType int `json:"oid_type"` // 目标 id 类型。0：视频（oid 为 avid）
	Ps      int `json:"ps"`       // 每页项数
	Pn      int `json:"pn"`       // 页码
}

// GetArchivePublicNotesResult 是查询稿件公开笔记的 data 对象。
type GetArchivePublicNotesResult struct {
	List           []GetArchivePublicNotesItem `json:"list"`             // 公开笔记列表，按创建顺序排列
	Page           GetArchivePublicNotesPage   `json:"page"`             // 页面信息
	ShowPublicNote bool                        `json:"show_public_note"` // 是否展示公开笔记
	Message        string                      `json:"message"`          // 提示信息
}

// GetArchivePublicNotesItem 是稿件公开笔记列表中的一项。
type GetArchivePublicNotesItem struct {
	Cvid    int64                       `json:"cvid"`     // 公开笔记对应的专栏 cvid
	Title   string                      `json:"title"`    // 笔记标题
	Summary string                      `json:"summary"`  // 笔记预览
	Pubtime string                      `json:"pubtime"`  // 发布时间。格式 YYYY-MM-DD hh:mm
	WebURL  string                      `json:"web_url"`  // 笔记 h5 页 url
	Message string                      `json:"message"`  // 更新信息。"更新于 YYYY-MM-DD hh:mm"
	Author  GetArchivePublicNotesAuthor `json:"author"`   // 作者信息
	Likes   int                         `json:"likes"`    // 点赞数
	HasLike bool                        `json:"has_like"` // 是否点赞。需要登录（Cookie）
}

// GetArchivePublicNotesAuthor 是稿件公开笔记的作者信息。
type GetArchivePublicNotesAuthor struct {
	Mid     int                                `json:"mid"`      // 用户 mid
	Name    string                             `json:"name"`     // 昵称
	Face    string                             `json:"face"`     // 头像 url
	Level   int                                `json:"level"`    // 用户等级
	VipInfo GetArchivePublicNotesAuthorVipInfo `json:"vip_info"` // 会员信息
	Pendant GetArchivePublicNotesAuthorPendant `json:"pendant"`  // 头像框信息
}

// GetArchivePublicNotesAuthorVipInfo 是稿件公开笔记作者的会员信息。
type GetArchivePublicNotesAuthorVipInfo struct {
	Type               int                                `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                                `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                                `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                                `json:"vip_pay_type"`         // 支付类型。0：未支付。1：已支付
	ThemeType          int                                `json:"theme_type"`           // 0。作用尚不明确
	Label              GetArchivePublicNotesVipLabel      `json:"label"`                // 会员标签
	AvatarSubscript    int                                `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                             `json:"nickname_color"`       // 会员昵称颜色
	Role               int                                `json:"role"`                 // 大角色类型。1：月度。3：年度。7：十年。15：百年
	AvatarSubscriptURL string                             `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                                `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                                `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetArchivePublicNotesVipOttInfo    `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetArchivePublicNotesVipSuperVip   `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                                `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetArchivePublicNotesVipAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
}

// GetArchivePublicNotesVipLabel 是稿件公开笔记作者的会员标签。
type GetArchivePublicNotesVipLabel struct {
	Path                  string                            `json:"path"`                      // 空。作用尚不明确
	Text                  string                            `json:"text"`                      // 会员类型文案
	LabelTheme            string                            `json:"label_theme"`               // 会员标签
	TextColor             string                            `json:"text_color"`                // 会员标签
	BgStyle               int                               `json:"bg_style"`                  // 1
	BgColor               string                            `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                            `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                              `json:"use_img_label"`             // true
	ImgLabelURIHans       string                            `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                            `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                            `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                            `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                               `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                               `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetArchivePublicNotesVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                            `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                            `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetArchivePublicNotesVipLabelGoto 是会员标签中的 label_goto 对象。
type GetArchivePublicNotesVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetArchivePublicNotesVipOttInfo 是会员信息中的 ott_info 对象。
type GetArchivePublicNotesVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetArchivePublicNotesVipSuperVip 是会员信息中的 super_vip 对象。
type GetArchivePublicNotesVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetArchivePublicNotesVipAvatarIcon 是会员信息中的 avatar_icon 对象。
type GetArchivePublicNotesVipAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetArchivePublicNotesAuthorPendant 是稿件公开笔记作者的头像框信息。
type GetArchivePublicNotesAuthorPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetArchivePublicNotesPage 是稿件公开笔记列表的分页信息。
type GetArchivePublicNotesPage struct {
	Total int `json:"total"` // 笔记总数
	Size  int `json:"size"`  // 每页项数
	Num   int `json:"num"`   // 页码
}

// GetArchivePublicNotes 查询稿件公开笔记列表
func (c *Client) GetArchivePublicNotes(ctx context.Context, param GetArchivePublicNotesParam) (*GetArchivePublicNotesResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/publish/list/archive"
	)
	return bilibili.Execute[*GetArchivePublicNotesResult](ctx, c.Client, method, url, param)
}

// GetUserPublicNotesParam 是查询当前用户公开笔记列表的请求参数。
type GetUserPublicNotesParam struct {
	Ps int `json:"ps"` // 每页项数
	Pn int `json:"pn"` // 页码
}

// GetUserPublicNotesResult 是查询用户公开笔记的 data 对象。
type GetUserPublicNotesResult struct {
	List []GetUserPublicNotesItem `json:"list"` // 明细列表，按创建顺序排列
	Page GetUserPublicNotesPage   `json:"page"` // 页面信息
}

// GetUserPublicNotesItem 是用户公开笔记列表中的一项。
//
// 接口文档未给出 list 元素的字段表，此处沿用稿件公开笔记的字段集。
type GetUserPublicNotesItem struct {
	Cvid    int64                    `json:"cvid"`     // 公开笔记对应的专栏 cvid
	Title   string                   `json:"title"`    // 笔记标题
	Summary string                   `json:"summary"`  // 笔记预览
	Pubtime string                   `json:"pubtime"`  // 发布时间。格式 YYYY-MM-DD hh:mm
	WebURL  string                   `json:"web_url"`  // 笔记 h5 页 url
	Message string                   `json:"message"`  // 更新信息。"更新于 YYYY-MM-DD hh:mm"
	Author  GetUserPublicNotesAuthor `json:"author"`   // 作者信息
	Likes   int                      `json:"likes"`    // 点赞数
	HasLike bool                     `json:"has_like"` // 是否点赞。需要登录（Cookie）
}

// GetUserPublicNotesAuthor 是用户公开笔记的作者信息。
type GetUserPublicNotesAuthor struct {
	Mid     int                             `json:"mid"`      // 用户 mid
	Name    string                          `json:"name"`     // 昵称
	Face    string                          `json:"face"`     // 头像 url
	Level   int                             `json:"level"`    // 用户等级
	VipInfo GetUserPublicNotesAuthorVipInfo `json:"vip_info"` // 会员信息
	Pendant GetUserPublicNotesAuthorPendant `json:"pendant"`  // 头像框信息
}

// GetUserPublicNotesAuthorVipInfo 是用户公开笔记作者的会员信息。
type GetUserPublicNotesAuthorVipInfo struct {
	Type               int                             `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                             `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                             `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                             `json:"vip_pay_type"`         // 支付类型。0：未支付。1：已支付
	ThemeType          int                             `json:"theme_type"`           // 0。作用尚不明确
	Label              GetUserPublicNotesVipLabel      `json:"label"`                // 会员标签
	AvatarSubscript    int                             `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                          `json:"nickname_color"`       // 会员昵称颜色
	Role               int                             `json:"role"`                 // 大角色类型。1：月度。3：年度。7：十年。15：百年
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                             `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                             `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetUserPublicNotesVipOttInfo    `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetUserPublicNotesVipSuperVip   `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                             `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetUserPublicNotesVipAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
}

// GetUserPublicNotesVipLabel 是用户公开笔记作者的会员标签。
type GetUserPublicNotesVipLabel struct {
	Path                  string                         `json:"path"`                      // 空。作用尚不明确
	Text                  string                         `json:"text"`                      // 会员类型文案
	LabelTheme            string                         `json:"label_theme"`               // 会员标签
	TextColor             string                         `json:"text_color"`                // 会员标签
	BgStyle               int                            `json:"bg_style"`                  // 1
	BgColor               string                         `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                         `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                           `json:"use_img_label"`             // true
	ImgLabelURIHans       string                         `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                         `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                         `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                         `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                            `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                            `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetUserPublicNotesVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                         `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                         `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetUserPublicNotesVipLabelGoto 是会员标签中的 label_goto 对象。
type GetUserPublicNotesVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetUserPublicNotesVipOttInfo 是会员信息中的 ott_info 对象。
type GetUserPublicNotesVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetUserPublicNotesVipSuperVip 是会员信息中的 super_vip 对象。
type GetUserPublicNotesVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetUserPublicNotesVipAvatarIcon 是会员信息中的 avatar_icon 对象。
type GetUserPublicNotesVipAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetUserPublicNotesAuthorPendant 是用户公开笔记作者的头像框信息。
type GetUserPublicNotesAuthorPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetUserPublicNotesPage 是用户公开笔记列表的分页信息。
type GetUserPublicNotesPage struct {
	Total int `json:"total"` // 笔记总数
	Size  int `json:"size"`  // 每页项数
	Num   int `json:"num"`   // 页码
}

// GetUserPublicNotes 查询当前用户的公开笔记列表
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetUserPublicNotes(ctx context.Context, param GetUserPublicNotesParam) (*GetUserPublicNotesResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/publish/list/user"
	)
	return bilibili.Execute[*GetUserPublicNotesResult](ctx, c.Client, method, url, param)
}
