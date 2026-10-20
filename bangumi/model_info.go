package bangumi

// 剧集基本信息、剧集明细与分集信息的响应模型。
//
// 本文件的类型只服务于 GetBangumiInfo / GetBangumiDetail / GetBangumiSection，
// 不与其它接口共用；嵌套对象与数组元素也各自独立定义。

// 剧集基本信息相关响应模型。

// GetBangumiInfoResult 获取剧集基本信息的结果。
type GetBangumiInfoResult struct {
	Media  GetBangumiInfoMedia   `json:"media"`            // 剧集信息
	Review *GetBangumiInfoReview `json:"review,omitempty"` // 用户操作信息。仅登录时存在
}

// GetBangumiInfoMedia 剧集基本信息。
type GetBangumiInfoMedia struct {
	Areas             []GetBangumiInfoArea `json:"areas"`              // 地区
	Cover             string               `json:"cover"`              // 封面图片 URL
	HorizontalPicture string               `json:"horizontal_picture"` // 横版封面图片 URL
	MediaID           int                  `json:"media_id"`           // 剧集 mdid
	NewEp             GetBangumiInfoNewEp  `json:"new_ep"`             // 最新一话信息
	Rating            GetBangumiInfoRating `json:"rating"`             // 评分信息
	SeasonID          int                  `json:"season_id"`          // 剧集 ssid
	ShareURL          string               `json:"share_url"`          // 剧集详情页链接
	Title             string               `json:"title"`              // 标题
	Type              int                  `json:"type"`               // 剧集类型 id。1：番剧，2：电影，3：纪录片，4：国创，5：电视剧，6：漫画，7：综艺
	TypeName          string               `json:"type_name"`          // 剧集类型名称
}

// GetBangumiInfoArea 剧集所属地区。
type GetBangumiInfoArea struct {
	ID   int    `json:"id"`   // 所属地区编号
	Name string `json:"name"` // 所属地区名称
}

// GetBangumiInfoNewEp 剧集最新一话信息。
type GetBangumiInfoNewEp struct {
	ID        int    `json:"id"`         // 最新一话的 epid
	Index     string `json:"index"`      // 最新一话名称
	IndexShow string `json:"index_show"` // 最新一话显示名称。例如「第1话」「更新至第4话」
}

// GetBangumiInfoRating 剧集评分信息。
type GetBangumiInfoRating struct {
	Count int     `json:"count"` // 总计评分人数
	Score float64 `json:"score"` // 评分
}

// GetBangumiInfoReview 当前用户对剧集的操作信息。
type GetBangumiInfoReview struct {
	IsCoin int `json:"is_coin"` // 作用尚不明确。实测为 0
	IsOpen int `json:"is_open"` // 作用尚不明确。实测为 1
}

// GetBangumiDetailResult 获取剧集明细的结果。
type GetBangumiDetailResult struct {
	Activity              GetBangumiDetailActivity      `json:"activity"`                // 参与的活动
	Actors                string                        `json:"actors"`                  // 声优
	Alias                 string                        `json:"alias"`                   // 作用尚不明确。通常为空
	Areas                 []GetBangumiDetailArea        `json:"areas"`                   // 地区
	BkgCover              string                        `json:"bkg_cover"`               // 网页背景图片 URL。无则为空
	Cover                 string                        `json:"cover"`                   // 剧集封面图片 URL
	DeliveryFragmentVideo bool                          `json:"delivery_fragment_video"` // 是否分段投递
	EnableVt              bool                          `json:"enable_vt"`               // 是否启用 VT
	Episodes              []GetBangumiDetailEpisode     `json:"episodes"`                // 正片剧集列表
	Evaluate              string                        `json:"evaluate"`                // 简介
	Freya                 GetBangumiDetailFreya         `json:"freya"`                   // Freya 相关配置
	HideEpVvVtDm          int                           `json:"hide_ep_vv_vt_dm"`        // 隐藏分集的 VV/VT/弹幕
	IconFont              GetBangumiDetailIconFont      `json:"icon_font"`               // 图标字体
	JpTitle               string                        `json:"jp_title"`                // 作用尚不明确。通常为空
	Link                  string                        `json:"link"`                    // 简介页面 URL
	MediaID               int                           `json:"media_id"`                // 剧集 mdid
	Mode                  int                           `json:"mode"`                    // 作用尚不明确。实测为 2
	MultiViewInfo         GetBangumiDetailMultiViewInfo `json:"multi_view_info"`         // 多视角信息
	NewEp                 GetBangumiDetailNewEp         `json:"new_ep"`                  // 更新信息
	Payment               *GetBangumiDetailPayment      `json:"payment,omitempty"`       // 会员与付费信息。无相关内容则无此项
	PayPack               GetBangumiDetailPayPack       `json:"payPack"`                 // 付费包信息。目前只有霹雳布袋戏付费套餐一种
	PlayStrategy          GetBangumiDetailPlayStrategy  `json:"play_strategy"`           // 播放策略
	Positive              GetBangumiDetailPositive      `json:"positive"`                // 正能量信息
	Publish               GetBangumiDetailPublish       `json:"publish"`                 // 发布信息
	Rating                *GetBangumiDetailRating       `json:"rating,omitempty"`        // 评分信息。无相关内容则无此项
	Record                string                        `json:"record"`                  // 备案号。无则为空
	Rights                GetBangumiDetailRights        `json:"rights"`                  // 属性标志信息
	SeasonID              int                           `json:"season_id"`               // 番剧 ssid
	SeasonTitle           string                        `json:"season_title"`            // 剧集标题
	Seasons               []GetBangumiDetailSeason      `json:"seasons"`                 // 同系列所有季信息
	Section               []GetBangumiDetailSection     `json:"section"`                 // 花絮、PV、番外等非正片内容。无相关内容则无此项
	SectionBottomDesc     string                        `json:"section_bottom_desc"`     // 板块底部说明
	Series                GetBangumiDetailSeries        `json:"series"`                  // 系列信息
	ShareCopy             string                        `json:"share_copy"`              // 分享文案。《{标题}》+{备注}
	ShareSubTitle         string                        `json:"share_sub_title"`         // 分享副标题
	ShareURL              string                        `json:"share_url"`               // 番剧播放页面 URL
	Show                  GetBangumiDetailShow          `json:"show"`                    // 网页全屏标志
	ShowSeasonType        int                           `json:"show_season_type"`        // 展示的剧集类型
	SquareCover           string                        `json:"square_cover"`            // 方形封面图片 URL
	Staff                 string                        `json:"staff"`                   // 制作人员信息
	Stat                  GetBangumiDetailStat          `json:"stat"`                    // 状态数
	Status                int                           `json:"status"`                  // 剧集状态
	Styles                []string                      `json:"styles"`                  // 标签
	Subtitle              string                        `json:"subtitle"`                // 剧集副标题
	Title                 string                        `json:"title"`                   // 剧集标题
	Total                 int                           `json:"total"`                   // 总计正片集数。未完结大多为 -1，已完结为正整数
	Type                  int                           `json:"type"`                    // 剧集类型。1：番剧，2：电影，3：纪录片，4：国创，5：电视剧，7：综艺
	UpInfo                *GetBangumiDetailUpInfo       `json:"up_info,omitempty"`       // UP 主信息。无相关内容则无此项
	UserStatus            GetBangumiDetailUserStatus    `json:"user_status"`             // 当前用户状态
}

// GetBangumiDetailActivity 剧集参与的活动。
type GetBangumiDetailActivity struct {
	Cover     string                    `json:"cover"`       // 活动封面 URL
	HeadBgURL string                    `json:"head_bg_url"` // 作用尚不明确。通常为空
	ID        int                       `json:"id"`          // 活动 id
	Link      string                    `json:"link"`        // 活动链接
	Pendants  []GetBangumiDetailPendant `json:"pendants"`    // 活动挂件
	Title     string                    `json:"title"`       // 活动标题
}

// GetBangumiDetailPendant 活动或 UP 主挂件。
type GetBangumiDetailPendant struct {
	Image string `json:"image"` // 挂件图片 URL
	Name  string `json:"name"`  // 挂件名称
	Pid   int    `json:"pid"`   // 挂件 id
}

// GetBangumiDetailArea 剧集所属地区。
type GetBangumiDetailArea struct {
	ID   int    `json:"id"`   // 所属地区编号
	Name string `json:"name"` // 所属地区名称
}

// GetBangumiDetailEpisode 正片剧集。
type GetBangumiDetailEpisode struct {
	Aid                int                           `json:"aid"`                // 单集稿件 avid
	ArchiveAttr        int                           `json:"archive_attr"`       // 稿件属性
	Badge              string                        `json:"badge"`              // 标签文字。例如「会员」「限免」
	BadgeInfo          GetBangumiDetailBadgeInfo     `json:"badge_info"`         // 标签信息
	BadgeType          int                           `json:"badge_type"`         // 标签类型
	Bvid               string                        `json:"bvid"`               // 单集稿件 bvid
	Cid                int                           `json:"cid"`                // 视频 cid
	Cover              string                        `json:"cover"`              // 单集封面 URL
	Dimension          GetBangumiDetailDimension     `json:"dimension"`          // 分辨率信息
	Duration           int                           `json:"duration"`           // 时长
	EnableVt           bool                          `json:"enable_vt"`          // 是否启用 VT
	EpID               int                           `json:"ep_id"`              // 单集 epid
	From               string                        `json:"from"`               // 来源
	ID                 int                           `json:"id"`                 // 单集 epid
	Interaction        GetBangumiDetailInteraction   `json:"interaction"`        // 互动信息
	IsSubView          bool                          `json:"is_sub_view"`        // 是否为子视角
	IsViewHide         bool                          `json:"is_view_hide"`       // 是否隐藏观看
	Link               string                        `json:"link"`               // 单集网页 URL
	LongTitle          string                        `json:"long_title"`         // 单集完整标题
	MultiViewEps       []GetBangumiDetailMultiViewEp `json:"multi_view_eps"`     // 多视角分集
	PubTime            int64                         `json:"pub_time"`           // 发布时间。时间戳
	Pv                 int                           `json:"pv"`                 // 作用尚不明确。实测为 0
	ReleaseDate        string                        `json:"release_date"`       // 作用尚不明确。通常为空
	Rights             GetBangumiDetailEpisodeRights `json:"rights"`             // 单集权限
	SectionType        int                           `json:"section_type"`       // 板块类型
	ShareCopy          string                        `json:"share_copy"`         // 分享文案
	ShareURL           string                        `json:"share_url"`          // 单集网页 URL
	ShortLink          string                        `json:"short_link"`         // 单集网页 URL 短链接
	ShowTitle          string                        `json:"show_title"`         // 展示标题
	ShowDrmLoginDialog bool                          `json:"showDrmLoginDialog"` // 是否展示 DRM 登录弹窗
	Skip               GetBangumiDetailSkip          `json:"skip"`               // 片头片尾信息
	Status             int                           `json:"status"`             // 状态
	Subtitle           string                        `json:"subtitle"`           // 单集副标题。观看次数文字
	Title              string                        `json:"title"`              // 单集标题
	ToastTitle         string                        `json:"toast_title"`        // 提示标题
	Vid                string                        `json:"vid"`                // 单集 vid。格式为 vupload_+{cid}
}

// GetBangumiDetailBadgeInfo 标签信息。
type GetBangumiDetailBadgeInfo struct {
	BgColor      string `json:"bg_color"`       // 标签背景色。日间模式
	BgColorNight string `json:"bg_color_night"` // 标签背景色。夜间模式
	Text         string `json:"text"`           // 标签文本
}

// GetBangumiDetailDimension 分辨率信息。
type GetBangumiDetailDimension struct {
	Height int `json:"height"` // 高度
	Rotate int `json:"rotate"` // 旋转角度
	Width  int `json:"width"`  // 宽度
}

// GetBangumiDetailInteraction 互动信息。
type GetBangumiDetailInteraction struct {
	GraphVersion int  `json:"graph_version"` // 互动视频版本
	Interaction  bool `json:"interaction"`   // 是否为互动视频
}

// GetBangumiDetailMultiViewEp 多视角分集。
type GetBangumiDetailMultiViewEp struct {
	EpID int `json:"ep_id"` // 分集 epid
}

// GetBangumiDetailEpisodeRights 单集权限。
type GetBangumiDetailEpisodeRights struct {
	AllowDemand   int `json:"allow_demand"`   // 是否允许点播
	AllowDm       int `json:"allow_dm"`       // 是否允许弹幕
	AllowDownload int `json:"allow_download"` // 是否允许下载
	AreaLimit     int `json:"area_limit"`     // 地区限制
}

// GetBangumiDetailSkip 片头片尾信息。
type GetBangumiDetailSkip struct {
	Ed GetBangumiDetailSkipTime `json:"ed"` // 片尾
	Op GetBangumiDetailSkipTime `json:"op"` // 片头
}

// GetBangumiDetailSkipTime 片头或片尾的时间范围。
type GetBangumiDetailSkipTime struct {
	End   int `json:"end"`   // 结束时间。秒
	Start int `json:"start"` // 开始时间。秒
}

// GetBangumiDetailFreya Freya 相关配置。
type GetBangumiDetailFreya struct {
	BubbleDesc    string `json:"bubble_desc"`     // 气泡描述
	BubbleShowCnt int    `json:"bubble_show_cnt"` // 气泡展示次数
	IconShow      int    `json:"icon_show"`       // 是否展示图标
}

// GetBangumiDetailIconFont 图标字体。
type GetBangumiDetailIconFont struct {
	Name string `json:"name"` // 图标名称
	Text string `json:"text"` // 图标文本
}

// GetBangumiDetailMultiViewInfo 多视角信息。
type GetBangumiDetailMultiViewInfo struct {
	ChangingDance     string `json:"changing_dance"`       // 切换舞蹈
	IsMultiViewSeason bool   `json:"is_multi_view_season"` // 是否为多视角剧集
}

// GetBangumiDetailNewEp 剧集更新信息。
type GetBangumiDetailNewEp struct {
	Desc  string `json:"desc"`   // 更新备注
	ID    int    `json:"id"`     // 最新一话 epid
	IsNew int    `json:"is_new"` // 是否最新发布。0：否，1：是
	Title string `json:"title"`  // 最新一话标题
}

// GetBangumiDetailPayment 会员与付费信息。
type GetBangumiDetailPayment struct {
	Discount          int                     `json:"discount"`            // 折扣。100 为原价
	PayType           GetBangumiDetailPayType `json:"pay_type"`            // 支付相关
	Price             string                  `json:"price"`               // 售价
	Promotion         string                  `json:"promotion"`           // 推广信息
	Tip               string                  `json:"tip"`                 // 提示
	ViewStartTime     int64                   `json:"view_start_time"`     // 观看开始时间。时间戳
	VipDiscount       int                     `json:"vip_discount"`        // 大会员折扣
	VipFirstPromotion string                  `json:"vip_first_promotion"` // 大会员首次推广信息
	VipPrice          string                  `json:"vip_price"`           // 大会员售价
	VipPromotion      string                  `json:"vip_promotion"`       // 大会员推广信息
}

// GetBangumiDetailPayType 支付方式开关。
type GetBangumiDetailPayType struct {
	AllowDiscount    int `json:"allow_discount"`     // 是否启用折扣。0：否，1：是
	AllowPack        int `json:"allow_pack"`         // 是否允许付费包
	AllowTicket      int `json:"allow_ticket"`       // 是否启用票券
	AllowTimeLimit   int `json:"allow_time_limit"`   // 是否启用时间限制
	AllowVipDiscount int `json:"allow_vip_discount"` // 是否启用大会员折扣
	ForbidBB         int `json:"forbid_bb"`          // 是否禁止使用 B 币券
}

// GetBangumiDetailPayPack 付费包信息。
type GetBangumiDetailPayPack struct {
	ID                int    `json:"id"`                    // 付费包 id
	NotPaidTextForApp string `json:"not_paid_text_for_app"` // 未付费时的 App 文案
	PaidTextForApp    string `json:"paid_text_for_app"`     // 已付费时的 App 文案
	PayPackURL        string `json:"pay_pack_url"`          // 付费包链接
	Status            int    `json:"status"`                // 状态
	Title             string `json:"title"`                 // 标题
}

// GetBangumiDetailPlayStrategy 播放策略。
type GetBangumiDetailPlayStrategy struct {
	Strategies []any `json:"strategies"` // 策略列表。文档未给出元素字段
}

// GetBangumiDetailPositive 正能量信息。
type GetBangumiDetailPositive struct {
	ID    int    `json:"id"`    // id
	Title string `json:"title"` // 标题
}

// GetBangumiDetailPublish 剧集发布信息。
type GetBangumiDetailPublish struct {
	IsFinish      int    `json:"is_finish"`       // 完结状态。0：未完结，1：已完结
	IsStarted     int    `json:"is_started"`      // 是否发布。0：未发布，1：已发布
	PubTime       string `json:"pub_time"`        // 发布时间。格式为 YYYY-MM-DD hh:mm:ss
	PubTimeShow   string `json:"pub_time_show"`   // 发布时间文字介绍
	UnknowPubDate int    `json:"unknow_pub_date"` // 作用尚不明确。实测为 0
	Weekday       int    `json:"weekday"`         // 作用尚不明确。实测为 0
}

// GetBangumiDetailRating 剧集评分信息。
type GetBangumiDetailRating struct {
	Count int     `json:"count"` // 总计评分人数
	Score float64 `json:"score"` // 评分
}

// GetBangumiDetailRights 剧集属性标志信息。
type GetBangumiDetailRights struct {
	AllowBp         int    `json:"allow_bp"`          // 是否允许 B 币
	AllowBpRank     int    `json:"allow_bp_rank"`     // 是否允许 B 币排行
	AllowDownload   int    `json:"allow_download"`    // 是否允许下载
	AllowReview     int    `json:"allow_review"`      // 是否允许评论
	AreaLimit       int    `json:"area_limit"`        // 地区限制
	BanAreaShow     int    `json:"ban_area_show"`     // 是否展示地区限制
	CanWatch        int    `json:"can_watch"`         // 是否可观看
	Copyright       string `json:"copyright"`         // 版权标志。bilibili：授权，dujia：独家
	ForbidPre       int    `json:"forbid_pre"`        // 是否禁止预览
	FreyaWhite      int    `json:"freya_white"`       // Freya 白名单
	IsCoverShow     int    `json:"is_cover_show"`     // 是否展示封面
	IsPreview       int    `json:"is_preview"`        // 是否为预览
	OnlyVipDownload int    `json:"only_vip_download"` // 是否仅大会员可下载
	Resource        string `json:"resource"`          // 资源
	WatchPlatform   int    `json:"watch_platform"`    // 观看平台
}

// GetBangumiDetailSeason 同系列剧集。
type GetBangumiDetailSeason struct {
	Badge               string                      `json:"badge"`                 // 标签文字
	BadgeInfo           GetBangumiDetailBadgeInfo   `json:"badge_info"`            // 标签信息
	BadgeType           int                         `json:"badge_type"`            // 标签类型
	Cover               string                      `json:"cover"`                 // 封面图片 URL
	EnableVt            bool                        `json:"enable_vt"`             // 是否启用 VT
	HorizontalCover1610 string                      `json:"horizontal_cover_1610"` // 16:10 横版封面 URL
	HorizontalCover169  string                      `json:"horizontal_cover_169"`  // 16:9 横版封面 URL
	IconFont            GetBangumiDetailIconFont    `json:"icon_font"`             // 图标字体
	MediaID             int                         `json:"media_id"`              // 剧集 mdid
	NewEp               GetBangumiDetailSeasonNewEp `json:"new_ep"`                // 最新一话信息
	SeasonID            int                         `json:"season_id"`             // 剧集 ssid
	SeasonTitle         string                      `json:"season_title"`          // 剧集标题
	SeasonType          int                         `json:"season_type"`           // 剧集类型
	Stat                GetBangumiDetailSeasonStat  `json:"stat"`                  // 状态数
}

// GetBangumiDetailSeasonNewEp 同系列剧集的最新一话信息。
type GetBangumiDetailSeasonNewEp struct {
	Cover     string `json:"cover"`      // 封面图片 URL
	ID        int    `json:"id"`         // 最新一话 epid
	IndexShow string `json:"index_show"` // 最新一话显示名称
}

// GetBangumiDetailSeasonStat 同系列剧集的状态数。
type GetBangumiDetailSeasonStat struct {
	Favorites    int `json:"favorites"`     // 收藏数
	SeriesFollow int `json:"series_follow"` // 系列追番数
	Views        int `json:"views"`         // 播放数
	Vt           int `json:"vt"`            // VT 数
}

// GetBangumiDetailSection 花絮、PV、番外等非正片板块。
type GetBangumiDetailSection struct {
	Attr       int                              `json:"attr"`        // 属性
	EpisodeID  int                              `json:"episode_id"`  // 单集 id。实测为 0
	EpisodeIDs []int                            `json:"episode_ids"` // 单集 id 列表
	Episodes   []GetBangumiDetailSectionEpisode `json:"episodes"`    // 板块内容
	ID         int                              `json:"id"`          // 板块 id
	Report     GetBangumiDetailSectionReport    `json:"report"`      // 举报信息
	Title      string                           `json:"title"`       // 板块标题
	Type       int                              `json:"type"`        // 板块类型
	Type2      int                              `json:"type2"`       // 板块子类型
}

// GetBangumiDetailSectionReport 板块的举报信息。
type GetBangumiDetailSectionReport struct {
	SeasonID    string `json:"season_id"`    // 剧集 ssid
	SeasonType  string `json:"season_type"`  // 剧集类型
	SecTitle    string `json:"sec_title"`    // 板块标题
	SectionID   string `json:"section_id"`   // 板块 id
	SectionType string `json:"section_type"` // 板块类型
}

// GetBangumiDetailSectionEpisode 非正片板块中的单集。
type GetBangumiDetailSectionEpisode struct {
	Aid                int                           `json:"aid"`                // 单集稿件 avid
	ArchiveAttr        int                           `json:"archive_attr"`       // 稿件属性
	Badge              string                        `json:"badge"`              // 标签文字
	BadgeInfo          GetBangumiDetailBadgeInfo     `json:"badge_info"`         // 标签信息
	BadgeType          int                           `json:"badge_type"`         // 标签类型
	Bvid               string                        `json:"bvid"`               // 单集稿件 bvid
	Cid                int                           `json:"cid"`                // 视频 cid
	Cover              string                        `json:"cover"`              // 单集封面 URL
	Dimension          GetBangumiDetailDimension     `json:"dimension"`          // 分辨率信息
	Duration           int                           `json:"duration"`           // 时长
	EnableVt           bool                          `json:"enable_vt"`          // 是否启用 VT
	EpID               int                           `json:"ep_id"`              // 单集 epid
	From               string                        `json:"from"`               // 来源
	IconFont           GetBangumiDetailIconFont      `json:"icon_font"`          // 图标字体
	ID                 int                           `json:"id"`                 // 单集 epid
	Interaction        GetBangumiDetailInteraction   `json:"interaction"`        // 互动信息
	IsViewHide         bool                          `json:"is_view_hide"`       // 是否隐藏观看
	Link               string                        `json:"link"`               // 单集网页 URL
	LinkType           string                        `json:"link_type"`          // 链接类型
	LongTitle          string                        `json:"long_title"`         // 单集完整标题
	PubTime            int64                         `json:"pub_time"`           // 发布时间。时间戳
	Pv                 int                           `json:"pv"`                 // 播放数
	ReleaseDate        string                        `json:"release_date"`       // 上线日期
	Report             GetBangumiDetailEpisodeReport `json:"report"`             // 举报信息
	Rights             GetBangumiDetailEpisodeRights `json:"rights"`             // 单集权限
	SectionType        int                           `json:"section_type"`       // 板块类型
	ShareCopy          string                        `json:"share_copy"`         // 分享文案
	ShareURL           string                        `json:"share_url"`          // 单集网页 URL
	ShortLink          string                        `json:"short_link"`         // 单集网页 URL 短链接
	ShowTitle          string                        `json:"show_title"`         // 展示标题
	ShowDrmLoginDialog bool                          `json:"showDrmLoginDialog"` // 是否展示 DRM 登录弹窗
	Skip               GetBangumiDetailSkip          `json:"skip"`               // 片头片尾信息
	Stat               GetBangumiDetailEpisodeStat   `json:"stat"`               // 状态数
	StatForUnity       GetBangumiDetailStatForUnity  `json:"stat_for_unity"`     // Unity 状态数
	Status             int                           `json:"status"`             // 状态
	Subtitle           string                        `json:"subtitle"`           // 单集副标题
	Title              string                        `json:"title"`              // 单集标题
	ToastTitle         string                        `json:"toast_title"`        // 提示标题
	UpInfo             GetBangumiDetailEpisodeUpInfo `json:"up_info"`            // UP 主信息
	Vid                string                        `json:"vid"`                // 单集 vid
}

// GetBangumiDetailEpisodeReport 单集的举报信息。
type GetBangumiDetailEpisodeReport struct {
	Aid         string `json:"aid"`          // 稿件 avid
	EpTitle     string `json:"ep_title"`     // 单集标题
	Position    string `json:"position"`     // 位置
	SeasonID    string `json:"season_id"`    // 剧集 ssid
	SeasonType  string `json:"season_type"`  // 剧集类型
	SectionID   string `json:"section_id"`   // 板块 id
	SectionType string `json:"section_type"` // 板块类型
}

// GetBangumiDetailEpisodeStat 单集的状态数。
type GetBangumiDetailEpisodeStat struct {
	Coin     int `json:"coin"`     // 投币数
	Danmakus int `json:"danmakus"` // 弹幕数
	Likes    int `json:"likes"`    // 点赞数
	Play     int `json:"play"`     // 播放数
	Reply    int `json:"reply"`    // 评论数
	Vt       int `json:"vt"`       // VT 数
}

// GetBangumiDetailStatForUnity 单集的 Unity 状态数。
type GetBangumiDetailStatForUnity struct {
	Coin    int                              `json:"coin"`    // 投币数
	Danmaku GetBangumiDetailStatForUnityItem `json:"danmaku"` // 弹幕信息
	Likes   int                              `json:"likes"`   // 点赞数
	Reply   int                              `json:"reply"`   // 评论数
	Vt      GetBangumiDetailStatForUnityItem `json:"vt"`      // VT 信息
}

// GetBangumiDetailStatForUnityItem Unity 状态数中的图标与文本。
type GetBangumiDetailStatForUnityItem struct {
	Icon     string `json:"icon"`      // 图标 URL
	PureText string `json:"pure_text"` // 纯文本
	Text     string `json:"text"`      // 文本
	Value    int    `json:"value"`     // 数值
}

// GetBangumiDetailEpisodeUpInfo 单集中的 UP 主信息。
type GetBangumiDetailEpisodeUpInfo struct {
	Avatar      string `json:"avatar"`       // 头像图片 URL
	Follower    int    `json:"follower"`     // 粉丝数
	IsFollow    int    `json:"is_follow"`    // 是否已关注
	Mid         int    `json:"mid"`          // UP 主 mid
	Uname       string `json:"uname"`        // UP 主昵称
	VerifyType  int    `json:"verify_type"`  // 认证类型
	VerifyType2 int    `json:"verify_type2"` // 第二认证类型
}

// GetBangumiDetailSeries 系列信息。
type GetBangumiDetailSeries struct {
	DisplayType int    `json:"display_type"` // 展示类型
	SeriesID    int    `json:"series_id"`    // 系列 id
	SeriesTitle string `json:"series_title"` // 系列名
}

// GetBangumiDetailShow 网页全屏标志。
type GetBangumiDetailShow struct {
	WideScreen int `json:"wide_screen"` // 是否全屏。0：正常，1：全屏
}

// GetBangumiDetailStat 剧集的状态数。
type GetBangumiDetailStat struct {
	Coins      int    `json:"coins"`       // 投币数
	Danmakus   int    `json:"danmakus"`    // 弹幕数
	Favorite   int    `json:"favorite"`    // 收藏数
	Favorites  int    `json:"favorites"`   // 收藏数
	FollowText string `json:"follow_text"` // 追番文本
	Hot        int    `json:"hot"`         // 热度
	Likes      int    `json:"likes"`       // 点赞数
	Reply      int    `json:"reply"`       // 评论数
	Share      int    `json:"share"`       // 分享数
	Views      int    `json:"views"`       // 播放数
	Vt         int    `json:"vt"`          // VT 数
}

// GetBangumiDetailUpInfo 剧集 UP 主信息。
type GetBangumiDetailUpInfo struct {
	Avatar             string                   `json:"avatar"`               // 头像图片 URL
	AvatarSubscriptURL string                   `json:"avatar_subscript_url"` // 头像角标图片 URL
	Follower           int                      `json:"follower"`             // 粉丝数
	IsFollow           int                      `json:"is_follow"`            // 是否已关注
	Mid                int                      `json:"mid"`                  // UP 主 mid
	NicknameColor      string                   `json:"nickname_color"`       // 昵称颜色
	Pendant            GetBangumiDetailPendant  `json:"pendant"`              // 挂件
	ThemeType          int                      `json:"theme_type"`           // 主题类型
	Uname              string                   `json:"uname"`                // UP 主昵称
	VerifyType         int                      `json:"verify_type"`          // 认证类型
	VipLabel           GetBangumiDetailVipLabel `json:"vip_label"`            // 大会员标签
	VipStatus          int                      `json:"vip_status"`           // 大会员状态
	VipType            int                      `json:"vip_type"`             // 大会员类型
}

// GetBangumiDetailVipLabel 大会员标签。
type GetBangumiDetailVipLabel struct {
	BgColor     string `json:"bg_color"`     // 背景色
	BgStyle     int    `json:"bg_style"`     // 背景样式
	BorderColor string `json:"border_color"` // 边框颜色
	Text        string `json:"text"`         // 标签文本
	TextColor   string `json:"text_color"`   // 文字颜色
}

// GetBangumiDetailUserStatus 当前用户对剧集的状态。
type GetBangumiDetailUserStatus struct {
	AreaLimit    int `json:"area_limit"`    // 地区限制
	BanAreaShow  int `json:"ban_area_show"` // 是否展示地区限制
	Follow       int `json:"follow"`        // 是否追番
	FollowStatus int `json:"follow_status"` // 追番状态
	Login        int `json:"login"`         // 是否登录
	Pay          int `json:"pay"`           // 是否付费
	PayPackPaid  int `json:"pay_pack_paid"` // 是否已购买付费包
	Sponsor      int `json:"sponsor"`       // 是否赞助
}

// GetBangumiSectionResult 获取剧集分集信息的结果。
type GetBangumiSectionResult struct {
	MainSection GetBangumiSectionMain   `json:"main_section"` // 正片信息
	Section     []GetBangumiSectionMain `json:"section"`      // 花絮、PV 等。结构同 main_section
}

// GetBangumiSectionMain 正片或非正片板块。
type GetBangumiSectionMain struct {
	Episodes []GetBangumiSectionEpisode `json:"episodes"` // 分集信息
	ID       int                        `json:"id"`       // 分组 id
	Type     int                        `json:"type"`     // 分组类型。0：正片，1：PV&其他，2：OP&ED
	Title    string                     `json:"title"`    // 分组标题
}

// GetBangumiSectionEpisode 分集信息。
type GetBangumiSectionEpisode struct {
	Aid        int                        `json:"aid"`         // 视频 av 号
	Badge      string                     `json:"badge"`       // 标签文字
	BadgeInfo  GetBangumiSectionBadgeInfo `json:"badge_info"`  // 标签信息
	BadgeType  int                        `json:"badge_type"`  // 标签类型
	Cid        int                        `json:"cid"`         // 分集 cid
	Cover      string                     `json:"cover"`       // 分集封面
	From       string                     `json:"from"`        // 来源
	ID         int                        `json:"id"`          // 分集 epid
	IsPremiere int                        `json:"is_premiere"` // 是否为首播
	LongTitle  string                     `json:"long_title"`  // 长标题
	ShareURL   string                     `json:"share_url"`   // 分集播放页 URL
	Status     int                        `json:"status"`      // 状态
	Title      string                     `json:"title"`       // 短标题
	Vid        string                     `json:"vid"`         // 分集 vid
}

// GetBangumiSectionBadgeInfo 分集的标签信息。
type GetBangumiSectionBadgeInfo struct {
	BgColor      string `json:"bg_color"`       // 标签背景色。日间模式
	BgColorNight string `json:"bg_color_night"` // 标签背景色。夜间模式
	Text         string `json:"text"`           // 标签文本
}
