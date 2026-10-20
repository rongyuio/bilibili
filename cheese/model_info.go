package cheese

// GetCheeseInfo 接口的响应模型。
//
// 本文件的类型只服务于 GetCheeseInfo，不与其它接口共用；即使字段相同，
// 其它接口也在各自 model_*.go 中单独定义。

// GetCheeseInfoResult 是「获取课程基本信息」的 data 对象。
type GetCheeseInfoResult struct {
	ABTestInfo                       GetCheeseInfoABTestInfo            `json:"abtest_info"`                          // 样式A/B测试信息。作用尚不明确
	ActiveMarket                     []any                              `json:"active_market"`                        // 已选市场？作用尚不明确
	ActivityList                     []any                              `json:"activity_list"`                        // 活动列表。实测为空数组
	BeSubscription                   bool                               `json:"be_subscription"`                      // 是否已订阅？作用尚不明确
	Brief                            GetCheeseInfoBrief                 `json:"brief"`                                // 课程概述信息
	CatalogueTop                     GetCheeseInfoCatalogueTop          `json:"catalogue_top"`                        // 目录顶部引导信息
	Coach                            GetCheeseInfoCoach                 `json:"coach"`                                // 课程顾问（加群）信息
	Consulting                       GetCheeseInfoConsulting            `json:"consulting"`                           // 咨询服务信息
	Cooperation                      map[string]any                     `json:"cooperation"`                          // 合作信息
	Cooperators                      []any                              `json:"cooperators"`                          // 合作方列表。实测为空数组
	Coupon                           *GetCheeseInfoCoupon               `json:"coupon"`                               // 优惠信息。若无优惠则无此项
	CourseContent                    string                             `json:"course_content"`                       // 课程内容提示文字
	Courses                          []any                              `json:"courses"`                              // 课程列表。实测为空数组
	Cover                            string                             `json:"cover"`                                // 课程封面url
	EpCatalogue                      []any                              `json:"ep_catalogue"`                         // 分集目录。实测为空数组
	EpCount                          int                                `json:"ep_count"`                             // 总计集数
	EpisodePage                      GetCheeseInfoEpisodePage           `json:"episode_page"`                         // 课程分集信息
	EpisodeSort                      int                                `json:"episode_sort"`                         // 1。作用尚不明确
	EpisodeTag                       GetCheeseInfoEpisodeTag            `json:"episode_tag"`                          // 分集标签信息
	Episodes                         []GetCheeseInfoEpisode             `json:"episodes"`                             // 课程分集列表
	ExpiryDay                        int                                `json:"expiry_day"`                           // 课程有效期天数。0表示长期有效
	ExpiryInfoContent                string                             `json:"expiry_info_content"`                  // 有效期说明文字。如「长期有效」
	Faq                              GetCheeseInfoFAQ                   `json:"faq"`                                  // 常见问题信息1。合并格式
	Faq1                             GetCheeseInfoFAQ1                  `json:"faq1"`                                 // 常见问题信息2。分离格式
	IsEnableCash                     bool                               `json:"is_enable_cash"`                       // 是否允许现金支付
	IsSeries                         bool                               `json:"is_series"`                            // 是否为系列课程
	LiveEpCount                      int                                `json:"live_ep_count"`                        // 直播集数
	Notice                           map[string]any                     `json:"notice"`                               // 公告信息。实测为空对象
	OpenedEpCount                    int                                `json:"opened_ep_count"`                      // 已开放集数
	PackInfo                         *GetCheeseInfoPackInfo             `json:"pack_info"`                            // 联报套餐信息。无联报则无此项
	PaidJump                         GetCheeseInfoPaidJump              `json:"paid_jump"`                            // 付费跳转信息
	PaidView                         bool                               `json:"paid_view"`                            // 是否已付费观看
	Payment                          GetCheeseInfoPayment               `json:"payment"`                              // 付费信息
	PreviewedPurchaseNote            GetCheeseInfoPreviewedPurchaseNote `json:"previewed_purchase_note"`              // 试看状态下的购买提示文案
	PurchaseFormatNote               GetCheeseInfoPurchaseFormatNote    `json:"purchase_format_note"`                 // 购买须知信息（分离格式）
	PurchaseNote                     GetCheeseInfoPurchaseNote          `json:"purchase_note"`                        // 购买须知信息。合并格式
	PurchaseProtocol                 GetCheeseInfoPurchaseProtocol      `json:"purchase_protocol"`                    // 付费内容协议信息
	RecommendSeasons                 []GetCheeseInfoRecommendSeason     `json:"recommend_seasons"`                    // 推荐课程列表
	ReleaseBottomInfo                string                             `json:"release_bottom_info"`                  // 底部更新状态文字
	ReleaseInfo                      string                             `json:"release_info"`                         // 更新状态文字+更新连载状态文字
	ReleaseInfo2                     string                             `json:"release_info2"`                        // 更新连载状态文字
	ReleaseStatus                    string                             `json:"release_status"`                       // 更新状态文字
	SeasonID                         int                                `json:"season_id"`                            // 课程ssid。与番剧ssid不互通
	SeasonSelectedRanks              []GetCheeseInfoSeasonSelectedRank  `json:"season_selected_ranks"`                // 精选课程榜单信息
	SeasonTag                        int                                `json:"season_tag"`                           // 课程标签。作用尚不明确
	ShareURL                         string                             `json:"share_url"`                            // 页面url
	ShortLink                        string                             `json:"short_link"`                           // bilibili uri链接
	ShowWatermark                    bool                               `json:"show_watermark"`                       // 是否显示水印
	Stat                             GetCheeseInfoStat                  `json:"stat"`                                 // 播放数信息
	Status                           int                                `json:"status"`                               // 10。作用尚不明确
	StopSell                         bool                               `json:"stop_sell"`                            // 是否停止售卖
	SubscriptionUpdateCountCycleText string                             `json:"subscription_update_count_cycle_text"` // 订阅更新周期文案。作用尚不明确
	Subtitle                         string                             `json:"subtitle"`                             // 课程简介信息
	Title                            string                             `json:"title"`                                // 课程标题
	UpInfo                           GetCheeseInfoUpInfo                `json:"up_info"`                              // UP主信息
	UpdateStatus                     int                                `json:"update_status"`                        // 更新状态。作用尚不明确
	UserStatus                       *GetCheeseInfoUserStatus           `json:"user_status"`                          // 用户状态信息。需要登录（SESSDATA）且referer为 https://www.bilibili.com
	WatermarkInterval                int                                `json:"watermark_interval"`                   // 水印间隔。单位为秒
}

// GetCheeseInfoABTestInfo 是「获取课程基本信息」abtest_info 对象。
type GetCheeseInfoABTestInfo struct {
	StyleABTest int `json:"style_abtest"` // AB测试样式值。作用尚不明确
}

// GetCheeseInfoBrief 是「获取课程基本信息」brief 对象。
type GetCheeseInfoBrief struct {
	Content string                    `json:"content"` // 空
	Img     []GetCheeseInfoBriefImage `json:"img"`     // 简介图片
	Title   string                    `json:"title"`   // 课程概述
	Type    int                       `json:"type"`    // 2
}

// GetCheeseInfoBriefImage 是「获取课程基本信息」brief.img 数组中的对象。
type GetCheeseInfoBriefImage struct {
	AspectRatio float64 `json:"aspect_ratio"` // ？？？作用尚不明确
	URL         string  `json:"url"`          // 简介图片url
}

// GetCheeseInfoCatalogueTop 是「获取课程基本信息」catalogue_top 对象。
type GetCheeseInfoCatalogueTop struct {
	BtnText      string `json:"btn_text"`       // 按钮文字。如「点击保存」
	GuideText    string `json:"guide_text"`     // 引导文字
	ImgURL       string `json:"img_url"`        // 二维码图片url
	SubGuideText string `json:"sub_guide_text"` // 副引导文字
	Type         int    `json:"type"`           // ???作用尚不明确
}

// GetCheeseInfoCoach 是「获取课程基本信息」coach 对象。
type GetCheeseInfoCoach struct {
	DetailURL    string `json:"detail_url"`    // 详情页url
	DirectoryURL string `json:"directory_url"` // 课程目录页url
	H5URL        string `json:"h5_url"`        // H5页面url
}

// GetCheeseInfoConsulting 是「获取课程基本信息」consulting 对象。
type GetCheeseInfoConsulting struct {
	ConsultingFlag bool   `json:"consulting_flag"` // 是否开启咨询
	ConsultingURL  string `json:"consulting_url"`  // 咨询页面url。实测为空
}

// GetCheeseInfoCoupon 是「获取课程基本信息」coupon 对象。
type GetCheeseInfoCoupon struct {
	Amount             float64 `json:"amount"`               // 折扣乘数
	CouponType         int     `json:"coupon_type"`          // 优惠券类型。作用尚不明确
	DiscountAmount     string  `json:"discount_amount"`      // 折扣金额文字。如「领券立减115.42」
	DiscountMessage    string  `json:"discount_message"`     // 折扣提示文字
	ExpireMinute       string  `json:"expire_minute"`        // 结束时间（分钟精度）。YYYY-MM-DD HH:MM
	ExpireTime         string  `json:"expire_time"`          // 结束时间。YYYY-MM-DD HH:MM:SS
	IsFrequenter       bool    `json:"is_frequenter"`        // 是否为老用户？作用尚不明确
	ReceiveExpireTime  int64   `json:"receive_expire_time"`  // 领取过期时间。时间戳
	SceneBackgroundImg string  `json:"scene_background_img"` // 场景背景图url。实测为空
	SceneBenefitImg    string  `json:"scene_benefit_img"`    // 场景优惠图url。实测为空
	SceneCountdown     bool    `json:"scene_countdown"`      // 是否显示倒计时
	SceneMark          string  `json:"scene_mark"`           // 场景标记。实测为空
	ShortTitle         string  `json:"short_title"`          // 优惠短标题
	ShowAmount         string  `json:"show_amount"`          // 展示折扣。如「7.1」
	StartTime          string  `json:"start_time"`           // 起始时间。YYYY-MM-DD HH:MM:SS
	Status             int     `json:"status"`               // ？？？作用尚不明确
	Title              string  `json:"title"`                // 优惠标题
	Token              string  `json:"token"`                // 领取token
	UseExpireTime      int64   `json:"use_expire_time"`      // 使用过期时间。时间戳
	UseScope           string  `json:"use_scope"`            // 使用范围说明。如「仅限当前课程使用」
}

// GetCheeseInfoEpisodePage 是「获取课程基本信息」episode_page 对象。
type GetCheeseInfoEpisodePage struct {
	Next  bool `json:"next"`  // 是否存在下一页
	Num   int  `json:"num"`   // 当前页码
	Size  int  `json:"size"`  // 每页项数
	Total int  `json:"total"` // 总计项数
}

// GetCheeseInfoEpisodeTag 是「获取课程基本信息」episode_tag 对象。
type GetCheeseInfoEpisodeTag struct {
	PartPreviewTag string `json:"part_preview_tag"` // 部分试看标签。如「部分试看」
	PayTag         string `json:"pay_tag"`          // 付费标签。如「付费」
	PreviewTag     string `json:"preview_tag"`      // 试看标签。如「全集试看」
}

// GetCheeseInfoEpisode 是「获取课程基本信息」episodes 数组中的对象。
type GetCheeseInfoEpisode struct {
	Aid            int    `json:"aid"`              // 课程分集avid。与普通稿件avid部分不互通
	CatalogueIndex int    `json:"catalogue_index"`  // 分集在目录中的序号
	Cid            int64  `json:"cid"`              // 课程分集cid。与普通视频cid部分不互通
	Cover          string `json:"cover"`            // 分集封面url
	Duration       int    `json:"duration"`         // 课程分集时间长度。单位为秒
	EpStatus       int    `json:"ep_status"`        // 分集状态？作用尚不明确
	EpisodeCanView bool   `json:"episode_can_view"` // 分集是否可观看
	From           string `json:"from"`             // pugv
	ID             int    `json:"id"`               // 课程分集epid。与番剧epid不互通
	Index          int    `json:"index"`            // 课程分集数
	Label          string `json:"label"`            // 分集标签。如「全集试看」
	Page           int    `json:"page"`             // 1
	Play           int    `json:"play"`             // 课程分集播放量
	PlayWay        int    `json:"play_way"`         // 播放方式？作用尚不明确
	Playable       bool   `json:"playable"`         // 是否可播放
	PreviewToast   string `json:"preview_toast"`    // 试看提示文字
	ReleaseDate    int64  `json:"release_date"`     // 课程分集发布时间。时间戳
	ShowVt         bool   `json:"show_vt"`          // 是否显示播放量？作用尚不明确
	Status         int    `json:"status"`           // 分集权限属性。1：可观看，2：不可观看
	Subtitle       string `json:"subtitle"`         // 分集时长文字信息。如「视频课 · 37分3秒」
	Title          string `json:"title"`            // 课程分集标题
	Watched        bool   `json:"watched"`          // 是否观看该集。需要登录（SESSDATA）且referer为 https://www.bilibili.com
	WatchedHistory int    `json:"watchedHistory"`   // 该集观看历史。需要登录（SESSDATA）且referer为 https://www.bilibili.com
}

// GetCheeseInfoFAQ 是「获取课程基本信息」faq 对象（合并格式）。
type GetCheeseInfoFAQ struct {
	Content string `json:"content"` // 常见问题信息
	Link    string `json:"link"`    // 常见问题页面url
	Title   string `json:"title"`   // 常见问题
}

// GetCheeseInfoFAQ1 是「获取课程基本信息」faq1 对象（分离格式）。
type GetCheeseInfoFAQ1 struct {
	Items []GetCheeseInfoFAQItem `json:"items"` // 常见问题列表
	Title string                 `json:"title"` // 常见问题
}

// GetCheeseInfoFAQItem 是「获取课程基本信息」faq1.items 数组中的对象。
type GetCheeseInfoFAQItem struct {
	Answer   string `json:"answer"`   // 回答信息
	Question string `json:"question"` // 问题信息
}

// GetCheeseInfoPackInfo 是「获取课程基本信息」pack_info 对象。
type GetCheeseInfoPackInfo struct {
	PackItemList   []GetCheeseInfoPackItem `json:"pack_item_list"`   // 套餐项目列表
	PackNotice2    string                  `json:"pack_notice2"`     // 套餐提示文字。如「联报限时折上折」
	ShowPacksRight bool                    `json:"show_packs_right"` // 是否显示右侧套餐
	Title          string                  `json:"title"`            // 套餐区块标题。如「套餐优惠」
}

// GetCheeseInfoPackItem 是「获取课程基本信息」pack_info.pack_item_list 数组中的对象。
type GetCheeseInfoPackItem struct {
	ABTestInfo              GetCheeseInfoPackABTestInfo `json:"abtest_info"`                // AB测试信息
	Brief                   GetCheeseInfoBrief          `json:"brief"`                      // 套餐简介。结构与 data 中的 brief 对象相同
	BtnNotice               string                      `json:"btn_notice"`                 // 按钮文字。如「立即购买」
	BtnOriginalPriceNotice  string                      `json:"btn_original_price_notice"`  // 按钮原价提示。如「价格1194」
	BtnShowAmountUnit       string                      `json:"btn_show_amount_unit"`       // 按钮价格单位。如「元」
	CouponBatchToken        string                      `json:"coupon_batch_token"`         // 优惠券批次token。实测为空
	IsEnableCash            bool                        `json:"is_enable_cash"`             // 是否支持现金支付
	Name                    string                      `json:"name"`                       // 套餐名称
	OriginalPrice           string                      `json:"original_price"`             // 原价。单位为元
	OriginalPricePrefix     string                      `json:"original_price_prefix"`      // 原价前缀。如「价值」
	OriginalPriceUnit       string                      `json:"original_price_unit"`        // 原价单位。如「元」
	ProductAvailable        bool                        `json:"product_available"`          // 套餐是否可购买
	ProductCover            string                      `json:"product_cover"`              // 套餐封面url
	ProductID               int                         `json:"product_id"`                 // 套餐商品id
	ProductNotice           string                      `json:"product_notice"`             // 套餐提示文字。实测为空
	ProductPaid             bool                        `json:"product_paid"`               // 是否已购买套餐
	ProductStatus           int                         `json:"product_status"`             // 套餐状态。作用尚不明确
	ProductType             int                         `json:"product_type"`               // 套餐类型。作用尚不明确
	RelativeItemCountNotice string                      `json:"relative_item_count_notice"` // 关联课程数量提示。如「共3门课程」
	RelativeItemList        []GetCheeseInfoRelativeItem `json:"relative_item_list"`         // 关联课程列表
	SeasonEpCountNotice     string                      `json:"season_ep_count_notice"`     // 关联课程集数提示。如「含3门课，共242课时」
	ShowAmount              string                      `json:"show_amount"`                // 展示价格
	ShowAmountNotice        string                      `json:"show_amount_notice"`         // 展示价格提示。实测为空
	VisibleDetail           bool                        `json:"visible_detail"`             // 是否显示详情
}

// GetCheeseInfoPackABTestInfo 是「获取课程基本信息」pack_item_list 中的 AB 测试信息。
type GetCheeseInfoPackABTestInfo struct {
	UpgradeShowAB int `json:"upgrade_show_ab"` // 升级展示AB测试值。作用尚不明确
}

// GetCheeseInfoRelativeItem 是「获取课程基本信息」relative_item_list 数组中的对象。
type GetCheeseInfoRelativeItem struct {
	Cover               string `json:"cover"`                 // 关联课程封面url
	EpCountNotice       string `json:"ep_count_notice"`       // 关联课程集数提示。如「共116课时」
	ExpiryNotice        string `json:"expiry_notice"`         // 有效期提示。如「长期有效」
	OriginalPrice       string `json:"original_price"`        // 关联课程原价。单位为元
	OriginalPriceNotice string `json:"original_price_notice"` // 关联课程原价提示。如「价值398元」
	OriginalPricePrefix string `json:"original_price_prefix"` // 原价前缀。如「价值」
	OriginalPriceUnit   string `json:"original_price_unit"`   // 原价单位。如「元」
	SeasonID            int    `json:"season_id"`             // 关联课程season_id
	SeasonPaid          bool   `json:"season_paid"`           // 关联课程是否已购买
	Sort                int    `json:"sort"`                  // 排序值。从1开始
	Title               string `json:"title"`                 // 关联课程标题
}

// GetCheeseInfoPaidJump 是「获取课程基本信息」paid_jump 对象。
type GetCheeseInfoPaidJump struct {
	JumpURLForApp string `json:"jump_url_for_app"` // App内跳转url。bilibili:// 协议的深链接
	URL           string `json:"url"`              // 跳转url。实测为空
}

// GetCheeseInfoPayment 是「获取课程基本信息」payment 对象。
type GetCheeseInfoPayment struct {
	BpEnough       int     `json:"bp_enough"`       // B币是否充足？作用尚不明确
	Desc           string  `json:"desc"`            // 付费价格说明
	DiscountDesc   string  `json:"discount_desc"`   // 折扣说明信息。若无优惠则无此项
	DiscountPrefix string  `json:"discount_prefix"` // 折扣类型前缀。若无优惠则无此项
	MyBp           int     `json:"my_bp"`           // 当前B币折扣？作用尚不明确
	PayShade       string  `json:"pay_shade"`       // 观看所有视频付费说明
	Price          float64 `json:"price"`           // 课程价格。单位为B币（折扣后价格）
	PriceFormat    string  `json:"price_format"`    // 课程价格。单位为B币（折扣后价格）
	PriceUnit      string  `json:"price_unit"`      // 价格单位。如「元」
	RefreshText    string  `json:"refresh_text"`    // 刷新按钮文字
	SelectText     string  `json:"select_text"`     // 选择按钮文字
}

// GetCheeseInfoPreviewedPurchaseNote 是「获取课程基本信息」previewed_purchase_note 对象。
type GetCheeseInfoPreviewedPurchaseNote struct {
	LongWatchText string `json:"long_watch_text"` // 长时间观看提示文字。如「试看5分钟，购买后即可观看完整内容」
	PayText       string `json:"pay_text"`        // 支付按钮文字。如「券后支付」
	PriceFormat   string `json:"price_format"`    // 价格文字。如「 282.58 元」
	WatchText     string `json:"watch_text"`      // 观看提示文字。如「即可观看所有视频」
	WatchingText  string `json:"watching_text"`   // 试看提示文字。如「试看中，购买即可观看完整内容」
}

// GetCheeseInfoPurchaseFormatNote 是「获取课程基本信息」purchase_format_note 对象（分离格式）。
type GetCheeseInfoPurchaseFormatNote struct {
	ContentList []GetCheeseInfoPurchaseNoteItem `json:"content_list"` // 购买须知内容列表
	Link        string                          `json:"link"`         // 购买须知页面url
	Title       string                          `json:"title"`        // 购买须知标题
}

// GetCheeseInfoPurchaseNoteItem 是「获取课程基本信息」purchase_format_note.content_list 数组中的对象。
type GetCheeseInfoPurchaseNoteItem struct {
	Bold    bool   `json:"bold"`    // 是否加粗
	Content string `json:"content"` // 须知内容
	Number  string `json:"number"`  // 条目序号。如「1.」
}

// GetCheeseInfoPurchaseNote 是「获取课程基本信息」purchase_note 对象（合并格式）。
type GetCheeseInfoPurchaseNote struct {
	Content string `json:"content"` // 购买须知信息文字内容
	Link    string `json:"link"`    // 购买须知页面url
	Title   string `json:"title"`   // 购买须知
}

// GetCheeseInfoPurchaseProtocol 是「获取课程基本信息」purchase_protocol 对象。
type GetCheeseInfoPurchaseProtocol struct {
	Link  string `json:"link"`  // 付费内容协议页面url
	Title string `json:"title"` // 付费内容协议
}

// GetCheeseInfoRecommendSeason 是「获取课程基本信息」recommend_seasons 数组中的对象。
type GetCheeseInfoRecommendSeason struct {
	Cover     string `json:"cover"`      // 推荐课程封面url
	EpCount   string `json:"ep_count"`   // 集数文字信息。如「共7课时」
	ID        int    `json:"id"`         // 推荐课程season_id
	SeasonURL string `json:"season_url"` // 推荐课程url
	Subtitle  string `json:"subtitle"`   // 推荐课程副标题
	Title     string `json:"title"`      // 推荐课程标题
	View      int    `json:"view"`       // 推荐课程播放量
}

// GetCheeseInfoSeasonSelectedRank 是「获取课程基本信息」season_selected_ranks 数组中的对象。
type GetCheeseInfoSeasonSelectedRank struct {
	NightShowSeasonSelectedMarkURL string `json:"night_show_season_selected_mark_url"` // 夜间模式精选角标url。实测为空
	SeasonSelectedRankLink         string `json:"season_selected_rank_link"`           // 精选榜单页面url
	ShowSeasonSelectedMarkURL      string `json:"show_season_selected_mark_url"`       // 精选角标url
}

// GetCheeseInfoStat 是「获取课程基本信息」stat 对象。
type GetCheeseInfoStat struct {
	Play     int    `json:"play"`      // 播放量
	PlayDesc string `json:"play_desc"` // 播放量文字信息
	ShowVt   bool   `json:"show_vt"`   // 是否显示播放量？作用尚不明确
}

// GetCheeseInfoUpInfo 是「获取课程基本信息」up_info 对象。
type GetCheeseInfoUpInfo struct {
	Avatar      string               `json:"avatar"`       // UP主头像url
	Brief       string               `json:"brief"`        // UP主备注。非个人签名
	Follower    int                  `json:"follower"`     // UP主粉丝数
	IsFollow    int                  `json:"is_follow"`    // 是否关注UP主。需要登录（SESSDATA）且referer为 https://www.bilibili.com，0：未关注，1：已关注
	IsLiving    bool                 `json:"is_living"`    // UP主是否正在直播
	Link        string               `json:"link"`         // UP主空间url
	Mid         int                  `json:"mid"`          // UP主uid
	Pendant     GetCheeseInfoPendant `json:"pendant"`      // ???作用尚不明确
	SeasonCount int                  `json:"season_count"` // UP主开设课程数
	Uname       string               `json:"uname"`        // UP主昵称
}

// GetCheeseInfoPendant 是「获取课程基本信息」up_info.pendant 对象。
type GetCheeseInfoPendant struct {
	Image    string `json:"image"`    // 空。作用尚不明确
	Name     string `json:"name"`     // 空。作用尚不明确
	Pid      int    `json:"pid"`      // 挂件id。0表示无挂件
	Follower int    `json:"follower"` // 0。作用尚不明确
}

// GetCheeseInfoUserStatus 是「获取课程基本信息」user_status 对象。
type GetCheeseInfoUserStatus struct {
	Bp           int   `json:"bp"`            // B币数？作用尚不明确
	ExpireAt     int64 `json:"expire_at"`     // 到期时间。实测为0，作用尚不明确
	Favored      int   `json:"favored"`       // 是否收藏。0：未收藏，1：已收藏
	FavoredCount int   `json:"favored_count"` // 1。作用尚不明确
	IsExpired    bool  `json:"is_expired"`    // 是否已过期
	IsFirstPaid  bool  `json:"is_first_paid"` // 是否首次付费
	//nolint:misspell // 接口字段名就是 payed
	Paid              int                    `json:"payed"`               // 是否购买。0：未购买，1：已购买
	Progress          *GetCheeseInfoProgress `json:"progress"`            // 课程进度。未登录无此项
	UserExpiryContent string                 `json:"user_expiry_content"` // 有效期说明文字。如「长期有效」
}

// GetCheeseInfoProgress 是「获取课程基本信息」user_status.progress 对象。
type GetCheeseInfoProgress struct {
	LastEpID    int    `json:"last_ep_id"`    // 最后观看的epid
	LastEpIndex string `json:"last_ep_index"` // 最后观看的标题
	LastTime    int    `json:"last_time"`     // 最后观看的时间进度。单位为秒
}

// GetCheeseEpisodeList 接口的响应模型。
//
// 本文件的类型只服务于 GetCheeseEpisodeList，不与 GetCheeseInfo 共用。

// GetCheeseEpisodeListResult 是「获取课程分集列表」的 data 对象。
type GetCheeseEpisodeListResult struct {
	Items []GetCheeseEpisodeListItem `json:"items"` // 课程分集列表
	Page  GetCheeseEpisodeListPage   `json:"page"`  // 列表分页信息
}

// GetCheeseEpisodeListItem 是「获取课程分集列表」items 数组中的对象。
type GetCheeseEpisodeListItem struct {
	Aid            int    `json:"aid"`              // 课程分集avid。与普通稿件avid部分不互通
	CatalogueIndex int    `json:"catalogue_index"`  // 分集在目录中的序号
	Cid            int64  `json:"cid"`              // 课程分集cid。与普通视频cid部分不互通
	Cover          string `json:"cover"`            // 分集封面url
	Duration       int    `json:"duration"`         // 课程分集时间长度。单位为秒
	EpStatus       int    `json:"ep_status"`        // 分集状态？作用尚不明确
	EpisodeCanView bool   `json:"episode_can_view"` // 分集是否可观看
	From           string `json:"from"`             // pugv
	ID             int    `json:"id"`               // 课程分集epid。与番剧epid不互通
	Index          int    `json:"index"`            // 课程分集数
	Label          string `json:"label"`            // 分集标签。如「全集试看」
	Page           int    `json:"page"`             // 1
	Play           int    `json:"play"`             // 课程分集播放量
	PlayWay        int    `json:"play_way"`         // 播放方式？作用尚不明确
	Playable       bool   `json:"playable"`         // 是否可播放
	PreviewToast   string `json:"preview_toast"`    // 试看提示文字
	ReleaseDate    int64  `json:"release_date"`     // 课程分集发布时间。时间戳
	ShowVt         bool   `json:"show_vt"`          // 是否显示播放量？作用尚不明确
	Status         int    `json:"status"`           // 分集权限属性。1：可观看，2：不可观看
	Subtitle       string `json:"subtitle"`         // 分集时长文字信息。如「视频课 · 37分3秒」
	Title          string `json:"title"`            // 课程分集标题
	Watched        bool   `json:"watched"`          // 是否观看该集。需要登录（SESSDATA）且referer为 https://www.bilibili.com
	WatchedHistory int    `json:"watchedHistory"`   // 该集观看历史。需要登录（SESSDATA）且referer为 https://www.bilibili.com
}

// GetCheeseEpisodeListPage 是「获取课程分集列表」page 对象。
type GetCheeseEpisodeListPage struct {
	Next  bool `json:"next"`  // 是否存在下一页
	Num   int  `json:"num"`   // 当前页码
	Size  int  `json:"size"`  // 每页项数
	Total int  `json:"total"` // 总计项数
}
