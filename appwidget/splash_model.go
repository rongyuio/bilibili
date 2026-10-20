package appwidget

import "github.com/rongyuio/bilibili"

// APP 端开屏相关响应模型。各接口的返回体独立定义，互不复用。

// GetSplashBrandListResult 是获取 APP 端开屏图片列表接口的 data 对象。
type GetSplashBrandListResult struct {
	PullInterval    int                      `json:"pull_interval"`      // 作用尚不明确
	Forcibly        bool                     `json:"forcibly"`           // 作用尚不明确
	Rule            string                   `json:"rule"`               // 作用尚不明确
	List            []GetSplashBrandListItem `json:"list"`               // 开屏图片列表
	Show            []GetSplashBrandListShow `json:"show"`               // 默认显示项
	HasNewSplashSet bool                     `json:"has_new_splash_set"` // 是否有新的开屏设置
	NewSplashHash   string                   `json:"new_splash_hash"`    // 新开屏资源哈希
	ForceShowTimes  int                      `json:"force_show_times"`   // 强制展示次数
	ShowHash        string                   `json:"show_hash"`          // 开屏资源哈希
	BadgeFrom       string                   `json:"badge_from"`         // 角标来源
	QueryList       []any                    `json:"query_list"`         // 查询列表。无效时为 null
}

// GetSplashBrandListItem 是开屏图片列表中的一项。
type GetSplashBrandListItem struct {
	ID        int64  `json:"id"`         // 开屏图片 id
	Thumb     string `json:"thumb"`      // 开屏图片 url
	LogoURL   string `json:"logo_url"`   // “bilibili” logo 图片 url
	ThumbHash string `json:"thumb_hash"` // 开屏图片缩略图哈希。32 位 MD5
	ThumbSize int64  `json:"thumb_size"` // 开屏图片缩略图大小，单位字节
	LogoSize  int64  `json:"logo_size"`  // logo 图片大小，单位字节
	ThumbName string `json:"thumb_name"` // 开屏图片名称
	ShowLogo  bool   `json:"show_logo"`  // 是否展示 logo
}

// GetSplashBrandListShow 是默认显示项中的一项。
type GetSplashBrandListShow struct {
	ID              int64  `json:"id"`                // 开屏图片 id
	BeginTime       int64  `json:"begin_time"`        // 起始时间，时间戳
	EndTime         int64  `json:"end_time"`          // 结束时间，时间戳
	Probability     int    `json:"probability"`       // 作用尚不明确
	Duration        int    `json:"duration"`          // 显示时间，单位毫秒
	ShowLogo        bool   `json:"show_logo"`         // 是否展示 logo
	LogoPosition    string `json:"logo_position"`     // logo 位置
	NewSplash       bool   `json:"new_splash"`        // 是否为新开屏
	InitialPushTime int64  `json:"initial_push_time"` // 首次推送时间，Unix 秒级时间戳
}

// GetSplashListResult 是获取 APP 端开屏广告信息接口的 data 对象。
type GetSplashListResult struct {
	SplashRequestID string              `json:"splash_request_id"` // 当前请求 id。长度为 33，前 13 位是当前秒级时间戳
	MaxTime         int                 `json:"max_time"`          // 最长显示时间
	MinInterval     int                 `json:"min_interval"`      // 最少显示时间，千分之一秒单位
	PullInterval    int                 `json:"pull_interval"`     // 显示时间
	KeepIDs         any                 `json:"keep_ids"`          // 显示时间。文档标注为 int，示例返回 null
	List            []GetSplashListItem `json:"list"`              // 广告列表
}

// GetSplashListItem 是开屏广告列表中的一项。
type GetSplashListItem struct {
	ID                       int64                      `json:"id"`                         // 广告 id
	Type                     int                        `json:"type"`                       // 1
	CardType                 int                        `json:"card_type"`                  // 39
	Duration                 int                        `json:"duration"`                   // 持续时间
	BeginTime                int64                      `json:"begin_time"`                 // 开始展示时间，秒级时间戳
	EndTime                  int64                      `json:"end_time"`                   // 停止展示时间，毫秒级时间戳
	Thumb                    string                     `json:"thumb"`                      // 广告图片
	Hash                     string                     `json:"hash"`                       // 作用尚不明确
	LogoURL                  string                     `json:"logo_url"`                   // 空值
	LogoHash                 string                     `json:"logo_hash"`                  // 空值
	Skip                     int                        `json:"skip"`                       // 是否跳转
	URI                      string                     `json:"uri"`                        // 跳转的 url
	VideoURL                 string                     `json:"video_url"`                  // 视频 url
	VideoHash                string                     `json:"video_hash"`                 // 视频 hash
	VideoWidth               int                        `json:"video_width"`                // 视频宽度
	VideoHeight              int                        `json:"video_height"`               // 视频高度
	URITitle                 string                     `json:"uri_title"`                  // 空值
	Source                   int                        `json:"source"`                     // 929
	CmMark                   int                        `json:"cm_mark"`                    // 1
	AdCb                     string                     `json:"ad_cb"`                      // 作用尚不明确
	ResourceID               int64                      `json:"resource_id"`                // 926
	RequestID                string                     `json:"request_id"`                 // 作用尚不明确
	ClientIP                 string                     `json:"client_ip"`                  // 客户端 IP
	IsAd                     bool                       `json:"is_ad"`                      // 作用尚不明确
	IsAdLoc                  bool                       `json:"is_ad_loc"`                  // 作用尚不明确
	SchemaTitle              string                     `json:"schema_title"`               // 滑动进入第三方应用
	SchemaCallupWhiteList    []string                   `json:"schema_callup_white_list"`   // 跳转的 APP 的白名单
	Extra                    GetSplashListExtra         `json:"extra"`                      // 显示信息
	EnablePreDownload        bool                       `json:"enable_pre_download"`        // 是否可以当场下载 APP
	EnableBackgroundDownload bool                       `json:"enable_background_download"` // 是否可以静默下载 APP
	InteractType             int                        `json:"interact_type"`              // 是否可以跳转（互动）
	InteractURL              string                     `json:"interact_url"`               // 跳转（互动）url
	InteractDistance         int                        `json:"interact_distance"`          // 作用尚不明确
	GuideButtonList          []GetSplashListGuideButton `json:"guide_button_list"`          // 所有的按钮
	MarkWithSkipStyle        int                        `json:"mark_with_skip_style"`       // 作用尚不明确
	SkipButtonHeight         float64                    `json:"skip_button_height"`         // 作用尚不明确
}

// GetSplashListExtra 是开屏广告的显示信息。
type GetSplashListExtra struct {
	UseAdWebV2               bool                   `json:"use_ad_web_v2"`
	ShowURLs                 []string               `json:"show_urls"`
	ClickURLs                []string               `json:"click_urls"`
	Show1sURLs               []string               `json:"show_1s_urls"`
	DownloadWhitelist        []string               `json:"download_whitelist"`
	OpenWhitelist            []string               `json:"open_whitelist"`
	Card                     GetSplashListExtraCard `json:"card"`
	ReportTime               int64                  `json:"report_time"`
	SalesType                int                    `json:"sales_type"`
	SpecialIndustry          bool                   `json:"special_industry"`
	PreloadLandingpage       int                    `json:"preload_landingpage"`
	ShareInfo                map[string]any         `json:"share_info"`
	UpzoneEntranceType       int                    `json:"upzone_entrance_type"`
	UpzoneEntranceReportID   int64                  `json:"upzone_entrance_report_id"`
	TopviewPicURL            string                 `json:"topview_pic_url"`
	TopviewVideoURL          string                 `json:"topview_video_url"`
	ClickArea                int                    `json:"click_area"`
	ShopID                   int64                  `json:"shop_id"`
	UpMid                    int64                  `json:"up_mid"`
	TrackID                  string                 `json:"track_id"`
	EnableStoreDirectLaunch  int                    `json:"enable_store_direct_launch"`
	EnableDoubleJump         bool                   `json:"enable_double_jump"`
	StoreCallupCard          bool                   `json:"store_callup_card"`
	EnableH5Alert            bool                   `json:"enable_h5_alert"`
	SpecialIndustryStyle     int                    `json:"special_industry_style"`
	MacroReplacePriority     int                    `json:"macro_replace_priority"`
	FeedbackPanelStyle       int                    `json:"feedback_panel_style"`
	AdContentType            int                    `json:"ad_content_type"`
	EnableH5PreLoad          int                    `json:"enable_h5_pre_load"`
	HotActivityID            int64                  `json:"hot_activity_id"`
	ProductID                int64                  `json:"product_id"`
	LandingpageDownloadStyle int                    `json:"landingpage_download_style"`
}

// GetSplashListExtraCard 是显示信息中的卡片配置。
type GetSplashListExtraCard struct {
	CardType                       int  `json:"card_type"`
	ExtremeTeamStatus              bool `json:"extreme_team_status"`
	SupportTransition              bool `json:"support_transition"`
	UnderPlayerInteractionStyle    int  `json:"under_player_interaction_style"`
	GradeDenominator               int  `json:"grade_denominator"`
	StarLevel                      int  `json:"star_level"`
	LiveBookingPopulationThreshold int  `json:"live_booking_population_threshold"`
	OriMarkHidden                  int  `json:"ori_mark_hidden"`
	UseMultiCover                  bool `json:"use_multi_cover"`
	FoldTime                       int  `json:"fold_time"`
	LiveRoomPopularity             int  `json:"live_room_popularity"`
	LiveTagShow                    bool `json:"live_tag_show"`
	GradeLevel                     int  `json:"grade_level"`
}

// GetSplashListGuideButton 是开屏广告中的一个引导按钮。
type GetSplashListGuideButton struct {
	ID                         int64                   `json:"id"`
	X                          int                     `json:"x"`
	Y                          int                     `json:"y"`
	BgColor                    string                  `json:"bg_color"`
	BgColorNight               string                  `json:"bg_color_night"`
	TextColor                  string                  `json:"text_color"`
	TextColorNight             string                  `json:"text_color_night"`
	GuideInstructions          string                  `json:"guide_instructions"`
	GuideInstructionsNew       string                  `json:"guide_instructions_new"`
	Width                      bilibili.NumberOrString `json:"width"` // 文档标注为 str，示例返回数字
	JumpURL                    string                  `json:"jump_url"`
	SchemaPackageName          string                  `json:"schema_package_name"`
	SchemaTitle                string                  `json:"schema_title"`
	SchemaTitleNew             string                  `json:"schema_title_new"`
	Height                     bilibili.NumberOrString `json:"height"` // 文档标注为 str，示例返回数字
	SlideThresholdValue        bilibili.NumberOrString `json:"slide_threshold_value"`
	SlideBorderColor           string                  `json:"slide_border_color"`
	GuideImageURL              string                  `json:"guide_image_url"`
	GuideImageMd5              string                  `json:"guide_image_md5"`
	InteractStyle              bilibili.NumberOrString `json:"interact_style"`
	JumpImageURL               string                  `json:"jump_image_url"`
	JumpImageMd5               string                  `json:"jump_image_md5"`
	SchemaImageURL             string                  `json:"schema_image_url"`
	SchemaImageMd5             string                  `json:"schema_image_md5"`
	ClickExpandRatio           bilibili.NumberOrString `json:"click_expand_ratio"`
	LogoImageURL               string                  `json:"logo_image_url"`
	LogoImageMd5               string                  `json:"logo_image_md5"`
	RelatedIDs                 []int64                 `json:"related_ids"`
	SchemaList                 []any                   `json:"schema_list"`
	FontRatio                  float64                 `json:"font_ratio"`
	Seq                        int                     `json:"seq"`
	DegradeType                int                     `json:"degrade_type"`
	TwistAngle                 int                     `json:"twist_angle"`
	TwistSpeed                 int                     `json:"twist_speed"`
	SecondaryGuideInstructions string                  `json:"secondary_guide_instructions"`
	SecondaryFontRatio         float64                 `json:"secondary_font_ratio"`
	SecondaryTextColor         string                  `json:"secondary_text_color"`
	SecondaryTextColorNight    string                  `json:"secondary_text_color_night"`
	ActivityTime               int64                   `json:"activity_time"` // 秒级时间戳
	TimeShowType               int                     `json:"time_show_type"`
}
