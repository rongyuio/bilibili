package article

// GetUserArticlesList 接口的响应模型。
//
// 本文件的类型只服务于 GetUserArticlesList，不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// GetUserArticlesListResult 是「获取用户专栏文集列表」的 data 对象。
type GetUserArticlesListResult struct {
	Lists []GetUserArticlesListItem `json:"lists"` // 文集信息列表
	Total int                       `json:"total"` // 文集总数
}

// GetUserArticlesListItem 是文集信息列表中的一项。
type GetUserArticlesListItem struct {
	List      GetUserArticlesListItemList      `json:"list"`      // 文集概览
	Articles  []GetUserArticlesListItemArticle `json:"articles"`  // 文集内的文章列表
	Author    GetUserArticlesListItemAuthor    `json:"author"`    // 文集作者信息
	Last      GetUserArticlesListItemArticle   `json:"last"`      // -。作用尚不明确。结构与data.articles[]中相似
	Attention bool                             `json:"attention"` // 是否关注文集作者。false：未关注。true：已关注。需要登录(Cookie) 。未登录为false
}

// GetUserArticlesListItemList 是文集概览。
type GetUserArticlesListItemList struct {
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

// GetUserArticlesListItemArticle 是文集内的文章。
type GetUserArticlesListItemArticle struct {
	ID              int                               `json:"id"`         // 专栏文章id
	Category        GetUserArticlesListItemCategory   `json:"category"`   // 分类
	Categories      []GetUserArticlesListItemCategory `json:"categories"` // 分类
	Title           string                            `json:"title"`      // 标题
	Summary         string                            `json:"summary"`    // 摘要
	BannerURL       string                            `json:"banner_url"` // 封面图
	AuthorUid       int                               `json:"author_uid"` // 作者mid。实测为 293793435
	OnlyFans        int                               `json:"only_fans"`  // 是否仅粉丝可见。实测为 0
	TemplateID      int                               `json:"template_id"`
	State           int                               `json:"state"`
	Author          *GetUserArticlesListItemAuthor    `json:"author"` // UP主信息
	Reprint         int                               `json:"reprint"`
	ImageURLs       []string                          `json:"image_urls"`
	PublishTime     int                               `json:"publish_time"` // 发布时间戳。单位：秒
	Ctime           int                               `json:"ctime"`        // 提交时间戳。单位：秒
	Stats           GetUserArticlesListItemStats      `json:"stats"`        // 专栏文章数据统计
	Tags            []GetUserArticlesListItemTag      `json:"tags"`         // 标签
	Words           int                               `json:"words"`
	Dynamic         string                            `json:"dynamic"` // 粉丝动态文案
	OriginImageURLs []string                          `json:"origin_image_urls"`
	IsLike          bool                              `json:"is_like"`
	Media           *GetUserArticlesListItemMedia     `json:"media"`
	ApplyTime       string                            `json:"apply_time"` // 空串
	CheckTime       string                            `json:"check_time"` // 空串
	Original        int                               `json:"original"`
	ActID           int                               `json:"act_id"`
	CoverAvid       int                               `json:"cover_avid"`
	Type            int                               `json:"type"`
	LikeState       int                               `json:"like_state"` // 是否点赞。0：未点赞。1：已点赞。需要登录(Cookie) 。未登录为0
}

// GetUserArticlesListItemCategory 是文章的分类。
type GetUserArticlesListItemCategory struct {
	ID       int    `json:"id"`        // 分类id
	ParentID int    `json:"parent_id"` // 父级分类id
	Name     string `json:"name"`      // 分类名称
}

// GetUserArticlesListItemAuthor 是文集作者信息。
type GetUserArticlesListItemAuthor struct {
	Mid            int                                   `json:"mid"`             // 作者mid
	Name           string                                `json:"name"`            // 作者昵称
	Face           string                                `json:"face"`            // 作者头像url
	OfficialVerify GetUserArticlesListItemOfficialVerify `json:"official_verify"` // 作者认证信息
	Nameplate      GetUserArticlesListItemNameplate      `json:"nameplate"`       // 作者勋章
	Vip            GetUserArticlesListItemVip            `json:"vip"`             // 作者大会员状态
}

// GetUserArticlesListItemOfficialVerify 是作者的认证信息。
type GetUserArticlesListItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserArticlesListItemNameplate 是作者的勋章。
type GetUserArticlesListItemNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetUserArticlesListItemVip 是作者的大会员信息，键名为 camelCase。
type GetUserArticlesListItemVip struct {
	VipType            int                             `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                             `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                          `json:"dueRemark"`            // (?)
	AccessStatus       int                             `json:"accessStatus"`         // (?)
	VipStatus          int                             `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                          `json:"vipStatusWarn"`        // (?)
	ThemeType          int                             `json:"theme_type"`           // 会员样式 id
	Label              GetUserArticlesListItemVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                             `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                          `json:"nickname_color"`       // 昵称颜色
}

// GetUserArticlesListItemVipLabel 是作者大会员的铭牌样式。
type GetUserArticlesListItemVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetUserArticlesListItemStats 是文章的状态数信息。
type GetUserArticlesListItemStats struct {
	View     int `json:"view"`     // 阅读数
	Favorite int `json:"favorite"` // 收藏数
	Like     int `json:"like"`     // 点赞数
	Dislike  int `json:"dislike"`  // 点踩数
	Reply    int `json:"reply"`    // 评论数
	Share    int `json:"share"`    // 分享数
	Coin     int `json:"coin"`     // 投币数
	Dynamic  int `json:"dynamic"`  // 动态转发数
}

// GetUserArticlesListItemTag 是文章标签。
type GetUserArticlesListItemTag struct {
	Tid  int    `json:"tid"`  // 标签id
	Name string `json:"name"` // 标签名称
}

// GetUserArticlesListItemMedia 是文章的影视信息。
type GetUserArticlesListItemMedia struct {
	Score    int    `json:"score"`     // 0
	MediaID  int    `json:"media_id"`  // 0
	Title    string `json:"title"`     // 空串
	Cover    string `json:"cover"`     // 空串
	Area     string `json:"area"`      // 空串
	TypeID   int    `json:"type_id"`   // 0
	TypeName string `json:"type_name"` // 空串
	Spoiler  int    `json:"spoiler"`   // 0
}
