package note

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 笔记详细信息相关接口。各接口的返回体独立定义，互不复用。

// GetNoteIsForbidParam 是查询稿件是否禁止笔记的请求参数。
type GetNoteIsForbidParam struct {
	Aid int `json:"aid" request:"query"` // 稿件avid
}

// GetNoteIsForbidResult 是查询稿件是否禁止笔记的 data 对象。
type GetNoteIsForbidResult struct {
	ForbidNoteEntrance bool `json:"forbid_note_entrance"` // 是否禁止笔记。true：禁止。false：允许
}

// GetNoteIsForbid 查询该稿件是否禁止笔记
func (c *Client) GetNoteIsForbid(ctx context.Context, param GetNoteIsForbidParam) (*GetNoteIsForbidResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/is_forbid"
	)
	return bilibili.Execute[*GetNoteIsForbidResult](ctx, c.Client, method, url, param)
}

// GetPrivateNoteInfoParam 是查询私有笔记内容的请求参数。
type GetPrivateNoteInfoParam struct {
	Oid     int   `json:"oid" request:"query"`      // 目标id
	OidType int   `json:"oid_type" request:"query"` // 目标id类型。0：视频（oid=avid）
	NoteID  int64 `json:"note_id" request:"query"`  // 笔记id
}

// GetPrivateNoteInfoResult 是私有笔记详情的 data 对象。
type GetPrivateNoteInfoResult struct {
	Arc                *GetPrivateNoteInfoArc  `json:"arc"`                  // 对应视频稿件信息
	AuditStatus        int                     `json:"audit_status"`         // 0。作用暂不明确
	CidCount           int                     `json:"cid_count"`            // 视频分P数
	Content            string                  `json:"content"`              // 笔记正文json序列
	ForbidNoteEntrance bool                    `json:"forbid_note_entrance"` // 是否禁止笔记入口
	PubReason          string                  `json:"pub_reason"`           // 公开笔记相关原因
	PubStatus          int                     `json:"pub_status"`           // 公开笔记状态。1：未公开。2：已公开。4：退回
	PubVersion         int                     `json:"pub_version"`          // 公开笔记版本
	Summary            string                  `json:"summary"`              // 笔记预览文本
	Tags               []GetPrivateNoteInfoTag `json:"tags"`                 // 笔记跳转标签列表，按笔记中位置排列
	Title              string                  `json:"title"`                // 笔记标题
}

// GetPrivateNoteInfoArc 是私有笔记对应的视频稿件信息。
type GetPrivateNoteInfoArc struct {
	Oid     int    `json:"oid"`      // 目标id
	OidType int    `json:"oid_type"` // 目标id类型。0：视频（oid=avid）
	Title   string `json:"title"`    // 视频标题
	Pic     string `json:"pic"`      // 视频封面
	Status  int    `json:"status"`   // 稿件状态
	Desc    string `json:"desc"`     // 视频简介
}

// GetPrivateNoteInfoTag 是私有笔记中的跳转标签。
type GetPrivateNoteInfoTag struct {
	Cid     int `json:"cid"`     // 视频cid
	Status  int `json:"status"`  // 0。作用尚不明确
	Index   int `json:"index"`   // 在稿件中的分P索引
	Seconds int `json:"seconds"` // 视频进度
	Pos     int `json:"pos"`     // 笔记中位置
}

// GetPrivateNoteInfo 查询私有笔记内容
//
// 需要登录（Cookie 中的 SESSDATA）。
func (c *Client) GetPrivateNoteInfo(ctx context.Context, param GetPrivateNoteInfoParam) (*GetPrivateNoteInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/info"
	)
	return bilibili.Execute[*GetPrivateNoteInfoResult](ctx, c.Client, method, url, param)
}

// GetPublicNoteInfoParam 是查询公开笔记内容的请求参数。
type GetPublicNoteInfoParam struct {
	Cvid int `json:"cvid" request:"query"` // 笔记对应的专栏cvid。非笔记id
}

// GetPublicNoteInfoResult 是公开笔记详情的 data 对象。
type GetPublicNoteInfoResult struct {
	Cvid               int                      `json:"cvid"`                 // 公开笔记对应的专栏cvid
	NoteID             int64                    `json:"note_id"`              // 笔记id
	Title              string                   `json:"title"`                // 笔记标题
	Summary            string                   `json:"summary"`              // 笔记预览
	Content            string                   `json:"content"`              // 笔记正文json序列
	CidCount           int                      `json:"cid_count"`            // (?)
	PubStatus          int                      `json:"pub_status"`           // 公开状态
	Tags               []GetPublicNoteInfoTag   `json:"tags"`                 // 标签列表
	Arc                *GetPublicNoteInfoArc    `json:"arc"`                  // 稿件信息
	Author             *GetPublicNoteInfoAuthor `json:"author"`               // 作者信息
	ForbidNoteEntrance bool                     `json:"forbid_note_entrance"` // 是否禁止笔记入口
	LastMtimeText      string                   `json:"last_mtime_text"`      // 最后修改时间，例如 "2022-02-07 23:44"
}

// GetPublicNoteInfoArc 是公开笔记对应的稿件信息。
type GetPublicNoteInfoArc struct {
	Oid       int    `json:"oid"`        // 目标id
	OidType   int    `json:"oid_type"`   // 目标id类型。0：视频（oid=avid）
	Title     string `json:"title"`      // 标题
	Status    int    `json:"status"`     // (?)
	Pic       string `json:"pic"`        // 封面url
	Desc      string `json:"desc"`       // 简介
	CvidCount int    `json:"cvid_count"` // (?)。实测为 4
}

// GetPublicNoteInfoTag 是公开笔记中的跳转标签。
type GetPublicNoteInfoTag struct {
	Cid     int `json:"cid"`     // 视频cid
	Status  int `json:"status"`  // 0。作用尚不明确
	Index   int `json:"index"`   // 在稿件中的分P索引
	Seconds int `json:"seconds"` // 视频进度
	Pos     int `json:"pos"`     // 笔记中位置
}

// GetPublicNoteInfoAuthor 是公开笔记的作者信息。
type GetPublicNoteInfoAuthor struct {
	Mid     int                            `json:"mid"`      // 用户mid
	Name    string                         `json:"name"`     // 昵称
	Face    string                         `json:"face"`     // 头像url
	Level   int                            `json:"level"`    // 用户等级
	VipInfo GetPublicNoteInfoAuthorVipInfo `json:"vip_info"` // 会员信息
	Pendant GetPublicNoteInfoAuthorPendant `json:"pendant"`  // 头像框信息
}

// GetPublicNoteInfoAuthorVipInfo 是公开笔记作者的会员信息。
type GetPublicNoteInfoAuthorVipInfo struct {
	Type               int                            `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                            `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                            `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                            `json:"vip_pay_type"`         // 支付类型。0：未支付。1：已支付
	ThemeType          int                            `json:"theme_type"`           // 0。作用尚不明确
	Label              GetPublicNoteInfoVipLabel      `json:"label"`                // 会员标签
	AvatarSubscript    int                            `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                         `json:"nickname_color"`       // 会员昵称颜色
	Role               int                            `json:"role"`                 // 大角色类型。1：月度。3：年度。7：十年。15：百年
	AvatarSubscriptURL string                         `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                            `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                            `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetPublicNoteInfoVipOttInfo    `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetPublicNoteInfoVipSuperVip   `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                            `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetPublicNoteInfoVipAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
}

// GetPublicNoteInfoVipLabel 是会员标签。
type GetPublicNoteInfoVipLabel struct {
	Path                  string                        `json:"path"`                      // 空。作用尚不明确
	Text                  string                        `json:"text"`                      // 会员类型文案
	LabelTheme            string                        `json:"label_theme"`               // 会员标签
	TextColor             string                        `json:"text_color"`                // 会员标签
	BgStyle               int                           `json:"bg_style"`                  // 1
	BgColor               string                        `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                        `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                          `json:"use_img_label"`             // true
	ImgLabelURIHans       string                        `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                        `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                        `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                        `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                           `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                           `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetPublicNoteInfoVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                        `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                        `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetPublicNoteInfoVipLabelGoto 是会员标签中的 label_goto 对象。
type GetPublicNoteInfoVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetPublicNoteInfoVipOttInfo 是会员信息中的 ott_info 对象。
type GetPublicNoteInfoVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetPublicNoteInfoVipSuperVip 是会员信息中的 super_vip 对象。
type GetPublicNoteInfoVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetPublicNoteInfoVipAvatarIcon 是会员信息中的 avatar_icon 对象。
type GetPublicNoteInfoVipAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetPublicNoteInfoAuthorPendant 是公开笔记作者的头像框信息。
type GetPublicNoteInfoAuthorPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetPublicNoteInfo 查询公开笔记内容
func (c *Client) GetPublicNoteInfo(ctx context.Context, param GetPublicNoteInfoParam) (*GetPublicNoteInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/note/publish/info"
	)
	return bilibili.Execute[*GetPublicNoteInfoResult](ctx, c.Client, method, url, param)
}
