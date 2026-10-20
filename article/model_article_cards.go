package article

import "encoding/json"

// GetArticleCards 接口的响应模型。
//
// 卡片接口的 data 是以请求 ids 中每一项为键的对象，值随内容类型变化，
// 因此 GetArticleCardsResult 按内容归类，而不是照搬键名。

// 卡片对象类型。仅用于 GetArticleCardsResult.UnmarshalJSON 内部分类。
const (
	articleCardKindVideo   = "video"
	articleCardKindArticle = "article"
	articleCardKindLive    = "live"
)

// GetArticleCardsResult 是「获取专栏显示卡片信息」的 data 对象。
//
// 键为请求 ids 中的原始写法（如 av2、BV1xx411c7mD、cv1、lv5440），
// 值可能是视频、专栏或直播间对象，分别归入 Videos、Articles、LiveRooms。
// 无法识别的对象（含 null）保留原始 JSON 放在 Other 中，由调用方自行解析。
type GetArticleCardsResult struct {
	Videos    map[string]GetArticleCardsVideo    // 视频对象，键为 AV 号或 BV 号
	Articles  map[string]GetArticleCardsArticle  // 专栏对象，键为 CV 号
	LiveRooms map[string]GetArticleCardsLiveRoom // 直播间对象，键为直播间长/短号
	Other     map[string]json.RawMessage         // 未识别的对象，保留原始 JSON
}

// UnmarshalJSON 按特征字段把 data 中每一项分到对应类型。
//
// 单个卡片解析失败时不会作废整条响应，而是退化为 Other 中的原始 JSON。
func (c *GetArticleCardsResult) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	cards := GetArticleCardsResult{}
	for key, value := range raw {
		handled := false
		switch articleCardKindOf(value) {
		case articleCardKindVideo:
			var card GetArticleCardsVideo
			if json.Unmarshal(value, &card) == nil {
				if cards.Videos == nil {
					cards.Videos = make(map[string]GetArticleCardsVideo)
				}
				cards.Videos[key] = card
				handled = true
			}
		case articleCardKindArticle:
			var card GetArticleCardsArticle
			if json.Unmarshal(value, &card) == nil {
				if cards.Articles == nil {
					cards.Articles = make(map[string]GetArticleCardsArticle)
				}
				cards.Articles[key] = card
				handled = true
			}
		case articleCardKindLive:
			var card GetArticleCardsLiveRoom
			if json.Unmarshal(value, &card) == nil {
				if cards.LiveRooms == nil {
					cards.LiveRooms = make(map[string]GetArticleCardsLiveRoom)
				}
				cards.LiveRooms[key] = card
				handled = true
			}
		}
		if !handled {
			if cards.Other == nil {
				cards.Other = make(map[string]json.RawMessage)
			}
			cards.Other[key] = value
		}
	}
	*c = cards
	return nil
}

// articleCardKindOf 通过对象中的特征字段判断卡片类型。
func articleCardKindOf(raw json.RawMessage) string {
	var probe map[string]json.RawMessage
	if err := json.Unmarshal(raw, &probe); err != nil {
		return ""
	}
	switch {
	case probe["room_id"] != nil:
		return articleCardKindLive
	case probe["aid"] != nil:
		return articleCardKindVideo
	case probe["id"] != nil:
		return articleCardKindArticle
	default:
		return ""
	}
}

// GetArticleCardsVideo 是 data 中代表视频的对象（键名为 avxxx 或 Bvxxxxxxxxxx）。
//
// 详细字段可参见视频基本信息接口。
type GetArticleCardsVideo struct {
	Aid         int                      `json:"aid"`           // 稿件 avid
	Bvid        string                   `json:"bvid"`          // 稿件 bvid
	Cid         int                      `json:"cid"`           // 当前分P cid
	Copyright   int                      `json:"copyright"`     // 版权标志。1：自制。2：转载
	Cover43     string                   `json:"cover43"`       // 43 比例封面 url
	Ctime       int                      `json:"ctime"`         // 创建时间戳。单位：秒
	Desc        string                   `json:"desc"`          // 视频简介
	Dimension   GetArticleCardsDimension `json:"dimension"`     // 视频尺寸
	Duration    int                      `json:"duration"`      // 视频时长。单位：秒
	Dynamic     string                   `json:"dynamic"`       // 粉丝动态
	Owner       GetArticleCardsOwner     `json:"owner"`         // UP主信息
	Pic         string                   `json:"pic"`           // 封面 url
	Pubdate     int                      `json:"pubdate"`       // 发布时间戳。单位：秒
	Rights      GetArticleCardsRights    `json:"rights"`        // 视频权限信息
	ShortLinkV2 string                   `json:"short_link_v2"` // 短链接
	Stat        GetArticleCardsStat      `json:"stat"`          // 视频状态数
	State       int                      `json:"state"`         // 视频状态
	Tid         int                      `json:"tid"`           // 分区 id
	Title       string                   `json:"title"`         // 视频标题
	Tname       string                   `json:"tname"`         // 分区名称
	Videos      int                      `json:"videos"`        // 分P 数
	VtSwitch    bool                     `json:"vt_switch"`     // 是否支持互动视频
}

// GetArticleCardsDimension 是视频的尺寸。
type GetArticleCardsDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetArticleCardsOwner 是视频的 UP 主信息。
type GetArticleCardsOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetArticleCardsRights 是视频的权限信息。
type GetArticleCardsRights struct {
	Bp            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	Hd5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UgcPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UgcPayPreview int `json:"ugc_pay_preview"` // 0。作用尚不明确
	NoBackground  int `json:"no_background"`   // 0。作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0。作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0。作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0。作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0。作用尚不明确
}

// GetArticleCardsStat 是视频的状态数。
type GetArticleCardsStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetArticleCardsArticle 是 data 中代表专栏的对象（键名为 cvxxx）。
//
// 详细字段可参见专栏内容接口。
type GetArticleCardsArticle struct {
	ActID            int                       `json:"act_id"`             // 活动 id
	ApplyTime        string                    `json:"apply_time"`         // 申请时间。空串
	Attributes       int                       `json:"attributes"`         // 文章属性
	AuthenMark       any                       `json:"authenMark"`         // 认证标记。固定为 null
	Author           GetArticleCardsAuthor     `json:"author"`             // 作者信息
	BannerURL        string                    `json:"banner_url"`         // 头图 url
	Categories       []GetArticleCardsCategory `json:"categories"`         // 所属分类列表
	Category         GetArticleCardsCategory   `json:"category"`           // 所属分类
	CheckState       int                       `json:"check_state"`        // 审核状态
	CheckTime        string                    `json:"check_time"`         // 审核时间。空串
	ContentPicList   any                       `json:"content_pic_list"`   // 正文图片列表。固定为 null
	CoverAvid        int                       `json:"cover_avid"`         // 封面视频 avid
	Ctime            int                       `json:"ctime"`              // 创建时间戳。单位：秒
	Dispute          any                       `json:"dispute"`            // 争议信息。固定为 null
	Dynamic          string                    `json:"dynamic"`            // 粉丝动态
	ID               int                       `json:"id"`                 // 专栏文章 cvid
	ImageURLs        []string                  `json:"image_urls"`         // 动态封面图片 url 列表
	IsLike           bool                      `json:"is_like"`            // 是否点赞。需要登录(Cookie)
	List             *GetArticleCardsList      `json:"list"`               // 所属文集。可能为 null
	Media            *GetArticleCardsMedia     `json:"media"`              // 影视信息
	Mtime            int                       `json:"mtime"`              // 修改时间戳。单位：秒
	OriginImageURLs  []string                  `json:"origin_image_urls"`  // 封面原图 url 列表
	OriginTemplateID int                       `json:"origin_template_id"` // 原始模板 id
	Original         int                       `json:"original"`           // 是否为原创。0：否。1：是
	PrivatePub       int                       `json:"private_pub"`        // 是否私密发布
	PublishTime      int                       `json:"publish_time"`       // 发布时间戳。单位：秒
	Reprint          int                       `json:"reprint"`            // 转载状态
	State            int                       `json:"state"`              // 文章状态
	Stats            GetArticleCardsStats      `json:"stats"`              // 状态数信息
	Summary          string                    `json:"summary"`            // 摘要
	TemplateID       int                       `json:"template_id"`        // 模板 id
	Title            string                    `json:"title"`              // 标题
	TopVideoInfo     any                       `json:"top_video_info"`     // 置顶视频信息。固定为 null
	Type             int                       `json:"type"`               // 文章类别。0：文章。2：笔记
	Words            int                       `json:"words"`              // 字数
}

// GetArticleCardsCategory 是专栏所属分类。
type GetArticleCardsCategory struct {
	ID       int    `json:"id"`        // 分类id
	ParentID int    `json:"parent_id"` // 父级分类id
	Name     string `json:"name"`      // 分类名称
}

// GetArticleCardsList 是专栏所属文集。
type GetArticleCardsList struct {
	ID            int    `json:"id"`             // 文集rlid
	Mid           int    `json:"mid"`            // 文集作者mid
	Name          string `json:"name"`           // 文集名称
	ImageURL      string `json:"image_url"`      // 文集封面图片url
	UpdateTime    int    `json:"update_time"`    // 文集更新时间。时间戳
	Ctime         int    `json:"ctime"`          // 文集创建时间。时间戳
	PublishTime   int    `json:"publish_time"`   // 文集发布时间。时间戳
	Summary       string `json:"summary"`        // 文集简介
	Words         int    `json:"words"`          // 文集字数
	Read          int    `json:"read"`           // 文集阅读量
	ArticlesCount int    `json:"articles_count"` // 文集内文章数量
	State         int    `json:"state"`          // 1或3。作用尚不明确
	Reason        string `json:"reason"`         // 空。作用尚不明确
	ApplyTime     string `json:"apply_time"`     // 空。作用尚不明确
	CheckTime     string `json:"check_time"`     // 空。作用尚不明确
}

// GetArticleCardsMedia 是专栏的影视信息。
type GetArticleCardsMedia struct {
	Score    int    `json:"score"`     // 0
	MediaID  int    `json:"media_id"`  // 0
	Title    string `json:"title"`     // 空串
	Cover    string `json:"cover"`     // 空串
	Area     string `json:"area"`      // 空串
	TypeID   int    `json:"type_id"`   // 0
	TypeName string `json:"type_name"` // 空串
	Spoiler  int    `json:"spoiler"`   // 0
}

// GetArticleCardsStats 是专栏的状态数信息。
type GetArticleCardsStats struct {
	View     int `json:"view"`     // 阅读数
	Favorite int `json:"favorite"` // 收藏数
	Like     int `json:"like"`     // 点赞数
	Dislike  int `json:"dislike"`  // 点踩数
	Reply    int `json:"reply"`    // 评论数
	Share    int `json:"share"`    // 分享数
	Coin     int `json:"coin"`     // 投币数
	Dynamic  int `json:"dynamic"`  // 动态转发数
}

// GetArticleCardsAuthor 是专栏卡片中的作者信息。
type GetArticleCardsAuthor struct {
	Mid            int                           `json:"mid"`             // 作者 mid
	Name           string                        `json:"name"`            // 作者昵称
	Face           string                        `json:"face"`            // 作者头像 url
	Fans           int                           `json:"fans"`            // 粉丝数。卡片中实测为 0
	Level          int                           `json:"level"`           // 等级。卡片中实测为 0
	Nameplate      GetArticleCardsNameplate      `json:"nameplate"`       // 作者勋章
	OfficialVerify GetArticleCardsOfficialVerify `json:"official_verify"` // 作者认证信息
	Pendant        GetArticleCardsPendant        `json:"pendant"`         // 作者挂件
	Vip            GetArticleCardsVip            `json:"vip"`             // 作者大会员状态
}

// GetArticleCardsNameplate 是作者的勋章。
type GetArticleCardsNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetArticleCardsOfficialVerify 是作者的认证信息。
type GetArticleCardsOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetArticleCardsPendant 是作者的头像挂件。
type GetArticleCardsPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetArticleCardsVip 是作者的大会员信息。
type GetArticleCardsVip struct {
	Type               int                        `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                        `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                        `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                        `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int                        `json:"theme_type"`           // 0。作用尚不明确
	Label              GetArticleCardsVipLabel    `json:"label"`                // 会员标签
	AvatarSubscript    int                        `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                     `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int                        `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptURL string                     `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                        `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                        `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetArticleCardsVipOttInfo  `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetArticleCardsVipSuperVip `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                        `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetArticleCardsAvatarIcon  `json:"avatar_icon"`          // 大会员角标信息
}

// GetArticleCardsVipLabel 是作者大会员的标签。
type GetArticleCardsVipLabel struct {
	Path                  string                      `json:"path"`                      // 空。作用尚不明确
	Text                  string                      `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                      `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                      `json:"text_color"`                // 会员标签
	BgStyle               int                         `json:"bg_style"`                  // 1
	BgColor               string                      `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                      `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                        `json:"use_img_label"`             // true
	ImgLabelURIHans       string                      `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                      `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                      `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                      `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                         `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                         `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetArticleCardsVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                      `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                      `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetArticleCardsVipLabelGoto 是大会员标签的跳转信息。
type GetArticleCardsVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetArticleCardsVipOttInfo 是大会员信息中的 ott_info 对象。
type GetArticleCardsVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetArticleCardsVipSuperVip 是大会员信息中的 super_vip 对象。
type GetArticleCardsVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetArticleCardsAvatarIcon 是大会员角标信息。
type GetArticleCardsAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetArticleCardsLiveRoom 是 data 中代表直播间的对象。
//
// 详细字段可参见直播间基本信息接口。
type GetArticleCardsLiveRoom struct {
	AreaV2Name     string `json:"area_v2_name"`     // 分区完整名称
	Cover          string `json:"cover"`            // 直播封面
	Face           string `json:"face"`             // 主播头像
	LiveStatus     int    `json:"live_status"`      // 直播状态。0：未开播。1：直播中。2：轮播中
	Online         int    `json:"online"`           // 在线人数
	PendentRu      string `json:"pendent_ru"`       // 右侧挂件 url
	PendentRuColor string `json:"pendent_ru_color"` // 右侧挂件文字颜色
	PendentRuPic   string `json:"pendent_ru_pic"`   // 右侧挂件图片
	Role           int    `json:"role"`             // 角色
	RoomID         int    `json:"room_id"`          // 直播间长 id
	Title          string `json:"title"`            // 直播间标题
	UID            int    `json:"uid"`              // 主播 UID (mid)
	Uname          string `json:"uname"`            // 主播用户名
}
