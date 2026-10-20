package article

import (
	"bytes"
	"encoding/json"

	"github.com/rongyuio/bilibili"
)

// GetArticleContent 接口的响应模型。
//
// 本文件的类型只服务于 GetArticleContent，不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// GetArticleContentResult 是「获取专栏正文内容」的 data 对象。
type GetArticleContentResult struct {
	ActID            int                         `json:"act_id"`             // 操作 ID?
	ApplyTime        string                      `json:"apply_time"`         // 应用时间?
	Attributes       int                         `json:"attributes"`         // 属性位? 可能不存在
	AuthenMark       any                         `json:"authenMark"`         // 授权码? 恒为 null
	Author           GetArticleContentAuthor     `json:"author"`             // 作者信息
	BannerURL        string                      `json:"banner_url"`         // 文章头图 URL，空则为无
	Categories       []GetArticleContentCategory `json:"categories"`         // 专栏分类信息，首项为主分区，第二项为子分区
	Category         GetArticleContentCategory   `json:"category"`           // 子分区
	CheckState       int                         `json:"check_state"`        // 检查状态?
	CheckTime        string                      `json:"check_time"`         // 检查时间?
	Content          string                      `json:"content"`            // 文章内容。type 为 0 时为 HTML，为 3 时为 JSON 字符串，可解析为 GetArticleContentOps
	ContentPicList   json.RawMessage             `json:"content_pic_list"`   // 内容图片列表?
	CoverAvid        int                         `json:"cover_avid"`         // 封面视频 AV 号，0 为无视频
	Ctime            int                         `json:"ctime"`              // 创建时间，UNIX 秒级时间戳
	Dispute          json.RawMessage             `json:"dispute"`            // 争议信息?
	DynIDStr         string                      `json:"dyn_id_str"`         // 动态 opus id
	Dynamic          string                      `json:"dynamic"`            // 动态信息? 可能不存在
	ID               int                         `json:"id"`                 // 专栏文章 ID
	ImageURLs        []string                    `json:"image_urls"`         // 图片 URL
	IsLike           bool                        `json:"is_like"`            // 是否喜欢?
	Keywords         string                      `json:"keywords"`           // 关键词，以 , 分隔
	List             GetArticleContentList       `json:"list"`               // 文集信息
	Media            *GetArticleContentMedia     `json:"media"`              // 媒体信息? 可能为 null
	Mtime            int                         `json:"mtime"`              // 修改时间，UNIX 秒级时间戳
	Opus             json.RawMessage             `json:"opus"`               // opus 信息，type 为 3 时存在，结构见接口文档示例
	OriginImageURLs  []string                    `json:"origin_image_urls"`  // 原始图片 URL
	OriginTemplateID int                         `json:"origin_template_id"` // 原始模板 ID?
	Original         int                         `json:"original"`           // 是否原创。0：非原创。1：原创
	PrivatePub       int                         `json:"private_pub"`        // 仅自己可见
	PublishTime      int                         `json:"publish_time"`       // 发布时间，UNIX 秒级时间戳
	Reprint          int                         `json:"reprint"`            // 是否允许转载。0：不允许。1：允许规范转载
	State            int                         `json:"state"`              // 专栏状态
	Stats            GetArticleContentStats      `json:"stats"`              // 统计数据
	Summary          string                      `json:"summary"`            // 专栏开头部分内容，纯文本
	Tags             []GetArticleContentTag      `json:"tags"`               // 专栏标签
	TemplateID       int                         `json:"template_id"`        // 模板 ID?
	Title            string                      `json:"title"`              // 专栏标题
	TopVideoInfo     json.RawMessage             `json:"top_video_info"`     // 封面视频信息? 恒为 null
	TotalArtNum      int                         `json:"total_art_num"`      // 作者总文章数
	Type             int                         `json:"type"`               // 正文类型。0：HTML。3：JSON
	VersionID        int                         `json:"version_id"`         // 版本 ID?
	Words            int                         `json:"words"`              // 文章总词数
}

// GetArticleContentAuthor 是专栏正文中的作者信息，字段集比 GetArticlesInfoAuthor 多出 fans、level 与 pendant。
type GetArticleContentAuthor struct {
	Mid            int                             `json:"mid"`             // 作者 mid
	Name           string                          `json:"name"`            // 作者昵称
	Face           string                          `json:"face"`            // 作者头像 URL
	Pendant        GetArticleContentPendant        `json:"pendant"`         // 作者头像挂件
	OfficialVerify GetArticleContentOfficialVerify `json:"official_verify"` // 作者认证信息
	Nameplate      GetArticleContentNameplate      `json:"nameplate"`       // 作者勋章
	Vip            GetArticleContentAuthorVip      `json:"vip"`             // 作者大会员状态
	Fans           int                             `json:"fans"`            // 作者粉丝数
	Level          int                             `json:"level"`           // 作者等级
}

// GetArticleContentPendant 是作者的头像挂件。
type GetArticleContentPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetArticleContentOfficialVerify 是作者的认证信息。
type GetArticleContentOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetArticleContentNameplate 是作者的勋章。
type GetArticleContentNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetArticleContentAuthorVip 是专栏正文作者的大会员信息，键名为 snake_case。
type GetArticleContentAuthorVip struct {
	Type            int                       `json:"type"`             // 大会员类型。0：无。1：月会员。2：年以上会员
	Status          int                       `json:"status"`           // 大会员状态。0：无。1：有
	DueDate         int                       `json:"due_date"`         // 大会员到期时间，毫秒时间戳
	VipPayType      int                       `json:"vip_pay_type"`     // 是否自动续费?
	ThemeType       int                       `json:"theme_type"`       // 会员样式 id
	Label           GetArticleContentVipLabel `json:"label"`            // 会员铭牌样式
	AvatarSubscript int                       `json:"avatar_subscript"` // 是否显示会员图标?
	NicknameColor   string                    `json:"nickname_color"`   // 昵称颜色
}

// GetArticleContentVipLabel 是作者大会员的铭牌样式。
type GetArticleContentVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetArticleContentCategory 是专栏分类。
type GetArticleContentCategory struct {
	ID       int    `json:"id"`        // 分类id
	ParentID int    `json:"parent_id"` // 父级分类id
	Name     string `json:"name"`      // 分类名称
}

// GetArticleContentList 是专栏所属文集。
type GetArticleContentList struct {
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

// GetArticleContentMedia 是专栏正文中的媒体信息，比 GetArticleCardsMedia 多出 season_id。
type GetArticleContentMedia struct {
	Area     string `json:"area"`      // 地区
	Cover    string `json:"cover"`     // 封面 URL
	MediaID  int    `json:"media_id"`  // 媒体 id
	Score    int    `json:"score"`     // 评分
	SeasonID int    `json:"season_id"` // 剧集 id
	Spoiler  int    `json:"spoiler"`   // 是否包含剧透
	Title    string `json:"title"`     // 媒体标题
	TypeID   int    `json:"type_id"`   // 媒体类型 id
	TypeName string `json:"type_name"` // 媒体类型名称
}

// GetArticleContentStats 是专栏的统计数据。
type GetArticleContentStats struct {
	View     int `json:"view"`     // 阅读数
	Favorite int `json:"favorite"` // 收藏数
	Like     int `json:"like"`     // 点赞数
	Dislike  int `json:"dislike"`  // 点踩数
	Reply    int `json:"reply"`    // 评论数
	Share    int `json:"share"`    // 分享数
	Coin     int `json:"coin"`     // 投币数
	Dynamic  int `json:"dynamic"`  // 动态转发数
}

// GetArticleContentTag 是专栏标签。
type GetArticleContentTag struct {
	Tid  int    `json:"tid"`  // 标签id
	Name string `json:"name"` // 标签名称
}

// GetArticleContentOps 是 data.content 在 type 为 3 时的 JSON 结构。
// 调用方可用 json.Unmarshal([]byte(content.Content), &ops) 解析。
type GetArticleContentOps struct {
	Ops []GetArticleContentOp `json:"ops"` // 以 JSON 呈现的文本内容
}

// GetArticleContentOp 是 JSON 正文中的一个文本片段。
// insert 既可能是字符串也可能是卡片对象，二者互斥：字符串写入 Text，对象写入 Insert。
type GetArticleContentOp struct {
	Attributes *GetArticleContentOpAttribute `json:"-"` // 属性，可能不存在
	Text       string                        `json:"-"` // insert 为字符串时的文本
	Insert     *GetArticleContentInsert      `json:"-"` // insert 为对象时的卡片内容
}

// UnmarshalJSON 同时接受文档表格中的 attribute 与响应示例中的 attributes 两种键名。
func (op *GetArticleContentOp) UnmarshalJSON(data []byte) error {
	var raw struct {
		Attributes *GetArticleContentOpAttribute `json:"attributes"`
		Attribute  *GetArticleContentOpAttribute `json:"attribute"`
		Insert     json.RawMessage               `json:"insert"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	op.Attributes = raw.Attributes
	if op.Attributes == nil {
		op.Attributes = raw.Attribute
	}
	op.Text = ""
	op.Insert = nil
	insert := bytes.TrimSpace(raw.Insert)
	if len(insert) == 0 || bytes.Equal(insert, []byte("null")) {
		return nil
	}
	if insert[0] == '"' {
		return json.Unmarshal(insert, &op.Text)
	}
	var card GetArticleContentInsert
	if err := json.Unmarshal(insert, &card); err != nil {
		return err
	}
	op.Insert = &card
	return nil
}

// GetArticleContentOpAttribute 是文本片段的渲染属性，全部字段可能不存在。
type GetArticleContentOpAttribute struct {
	Align      *string `json:"align"`      // 文字对齐，参见 CSS text-align
	Blockquote *bool   `json:"blockquote"` // 块级引用
	Bold       *bool   `json:"bold"`       // 加粗
	Class      *string `json:"class"`      // 类名
	Color      *string `json:"color"`      // 颜色，十六进制颜色值
	Header     *int    `json:"header"`     // 标题级别
	Strike     *bool   `json:"strike"`     // 删除线
	Link       *string `json:"link"`       // 站内链接
	Italic     *bool   `json:"italic"`     // 斜体
	List       *string `json:"list"`       // 列表。bullet：无序列表。ordered：有序列表
}

// GetArticleContentInsert 是 insert 为对象时的卡片内容，同一次 insert 只会出现其中一种。
type GetArticleContentInsert struct {
	NativeImage *GetArticleContentNativeImage `json:"native-image"` // 原生图片
	CutOff      *GetArticleContentCutOff      `json:"cut-off"`      // 分割线
	VideoCard   *GetArticleContentCard        `json:"video-card"`   // 视频卡片
	ArticleCard *GetArticleContentCard        `json:"article-card"` // 专栏卡片
	VoteCard    *GetArticleContentCard        `json:"vote-card"`    // 投票卡片
	LiveCard    *GetArticleContentCard        `json:"live-card"`    // 直播间卡片
}

// GetArticleContentNativeImage 是原生图片，字段对应 <img> 的 data-* 属性。
type GetArticleContentNativeImage struct {
	Alt    string `json:"alt"`    // 备用文本描述，即 <img> 的 alt 属性
	URL    string `json:"url"`    // 图片 URL，即 <img> 的 src 属性
	Width  int    `json:"width"`  // 图片宽度，即 data-w 属性
	Height int    `json:"height"` // 图片高度，即 data-h 属性
	Size   int    `json:"size"`   // 图片文件大小，即 data-size 属性
	Status string `json:"status"` // 固定为 loaded，即 data-status 属性
}

// GetArticleContentCutOff 是分割线。
type GetArticleContentCutOff struct {
	Type string `json:"type"` // 类型，文档标注为字符串
	URL  string `json:"url"`  // 分割线图片 URL
}

// GetArticleContentCard 是视频、专栏、投票、直播间四类卡片的公共结构，四者的字段集完全一致。
type GetArticleContentCard struct {
	Alt    string                  `json:"alt"`    // 备用文本描述
	Height int                     `json:"height"` // 卡片高度
	ID     string                  `json:"id"`     // 卡片目标 id，如 av99999999、cv1、lv5440
	Size   any                     `json:"size"`   // 恒为 null
	Status string                  `json:"status"` // 固定为 loaded
	Tid    bilibili.NumberOrString `json:"tid"`    // 卡片类型。文档标注为 number，取值如 1.1、2、7、8
	URL    string                  `json:"url"`    // 卡片图片 URL
	Width  int                     `json:"width"`  // 卡片宽度
}
