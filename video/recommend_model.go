package video

// 视频推荐相关响应模型。

// GetStoryVideoListResult 根据点击视频获取的短视频播放列表。
type GetStoryVideoListResult struct {
	Config any                     `json:"config"` // 配置
	Items  []GetStoryVideoListItem `json:"items"`  // 视频信息列表
}

// GetStoryVideoListItem 短视频播放列表中的视频信息。
type GetStoryVideoListItem struct {
	Bvid               string                          `json:"bvid"`                // 视频bv号
	CardGoto           string                          `json:"card_goto"`           // 卡片跳转
	Copyright          int                             `json:"copyright"`           // 版权
	Cover              string                          `json:"cover"`               // 封面url
	Desc               string                          `json:"desc"`                // 视频描述
	Dimension          any                             `json:"dimension"`           // 包括视频尺寸、旋转角度。文档标注为num，实际可能为对象
	DislikeReasonsV2   any                             `json:"dislike_reasons_v2"`  // 不喜欢原因v2的各种显示
	DislikeReasonsV3   any                             `json:"dislike_reasons_v3"`  // 不喜欢原因v3的各种显示
	Duration           int                             `json:"duration"`            // 视频时长。单位为秒
	FfCover            string                          `json:"ff_cover"`            // 短视频封面原图
	Goto               string                          `json:"goto"`                // 跳转去向类型。此处为vertical_av
	Owner              GetStoryVideoListItemOwner      `json:"owner"`               // 拥有者，也就是up信息
	Param              string                          `json:"param"`               // 参数，实际为视频aid
	PlayerArgs         GetStoryVideoListItemPlayerArgs `json:"player_args"`         // 播放器参数
	Pubdate            int64                           `json:"pubdate"`             // 发布时间。秒级时间戳
	ReportFlowData     string                          `json:"report_flow_data"`    // 报告流数据。是个伪装成obj的str
	ReqUser            any                             `json:"req_user"`            // 未知。为空
	Rights             any                             `json:"rights"`              // 未知
	ShareBottomButton  any                             `json:"share_bottom_button"` // 分享下方按钮
	ShareGuide         any                             `json:"share_guide"`         // 未知
	ShortLink          string                          `json:"short_link"`          // 视频短链接
	ShowReport         any                             `json:"show_report"`         // 显示举报
	Stat               GetStoryVideoListItemStat       `json:"stat"`                // 视频信息
	SubTitle           string                          `json:"sub_title"`           // 子标题。显示的是描述播放量的文字
	SubmissionEntrance any                             `json:"submission_entrance"` // 提交入口。icon的uri
	ThreePointButton   any                             `json:"three_point_button"`  // 三点按钮要显示的内容。里面有两个array
	ThumbUpAnimation   string                          `json:"thumb_up_animation"`  // 未知
	Title              string                          `json:"title"`               // 视频标题
	TopSearchBar       any                             `json:"top_search_bar"`      // 搜索栏。内含一个跳转的uri
	TrackID            string                          `json:"track_id"`            // 路由track
	URI                string                          `json:"uri"`                 // uri
	ViewContent        string                          `json:"view_content"`        // 用于显示的文本形式播放量
	Vip                any                             `json:"vip"`                 // vip相关的信息
}

// GetStoryVideoListItemOwner 短视频播放列表中的UP主信息。
type GetStoryVideoListItemOwner struct {
	Attention      int                                `json:"attention"`       // 未知。包含一个mid的int字段
	Avatar         any                                `json:"avatar"`          // 一些显示设置
	Face           string                             `json:"face"`            // 头像url
	Fans           int                                `json:"fans"`            // 粉丝量
	LikeNum        int                                `json:"like_num"`        // 获赞数
	Mid            int                                `json:"mid"`             // up主的mid
	Name           string                             `json:"name"`            // up主的昵称
	OfficialVerify GetStoryVideoListItemOwnerOfficial `json:"official_verify"` // 官方认证。type=-1为无认证；type不为-1时title为称号
	Relation       int                                `json:"relation"`        // 关系？未知
	SubAvatar      any                                `json:"sub_avatar"`      // 子形象。包含一个mid的int字段
	Upower         any                                `json:"upower"`          // 充电相关。包含一个button_uri字段类型为str，是充电跳转链接
}

// GetStoryVideoListItemOwnerOfficial 是 UP 主的官方认证信息。
type GetStoryVideoListItemOwnerOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetStoryVideoListItemPlayerArgs 短视频播放列表中的播放器参数。
type GetStoryVideoListItemPlayerArgs struct {
	Aid  int64  `json:"aid"`  // 视频作者的aid
	Cid  int64  `json:"cid"`  // 所属频道的cid
	Type string `json:"type"` // 内容类型。视频一般都是av
}

// GetStoryVideoListItemStat 短视频播放列表中的视频状态数。
type GetStoryVideoListItemStat struct {
	Aid      int64 `json:"aid"`      // 视频作者的aid
	Coin     int   `json:"coin"`     // 视频硬币数
	Danmaku  int   `json:"danmaku"`  // 视频弹幕数
	Favorite int   `json:"favorite"` // 视频收藏数
	Follow   int   `json:"follow"`   // 未知。很多都是0
	Like     int   `json:"like"`     // 视频点赞数
	Reply    int   `json:"reply"`    // 视频评论与回复总数
	Share    int   `json:"share"`    // 视频分享数
	View     int   `json:"view"`     // 视频播放
}

// GetShortVideoListResult 短视频模式视频列表。
type GetShortVideoListResult struct {
	Config         GetShortVideoListConfig `json:"config"`          // 一些界面相关的内容
	Items          []GetShortVideoListItem `json:"items"`           // 视频列表
	InterestChoose any                     `json:"interest_choose"` // 兴趣选择？实测为null，作用尚不明确
}

// GetShortVideoListConfig 短视频模式下的界面配置。
type GetShortVideoListConfig struct {
	Column                     int     `json:"column"`                         // 列数？实测为2
	AutoplayCard               int     `json:"autoplay_card"`                  // 自动播放卡片？实测为2
	FeedCleanAbtest            int     `json:"feed_clean_abtest"`              // 信息流清理AB测试分组？实测为0
	HomeTransferTest           int     `json:"home_transfer_test"`             // 首页转场测试？实测为0
	AutoRefreshTime            int     `json:"auto_refresh_time"`              // 自动刷新时间。单位毫秒，实测为1200
	ShowInlineDanmaku          int     `json:"show_inline_danmaku"`            // 是否显示内嵌弹幕？实测为1
	Toast                      any     `json:"toast"`                          // 未知。响应示例中为空对象
	SingleAutoplayFlag         int     `json:"single_autoplay_flag"`           // 未知。响应示例中为1
	IsBackToHomepage           bool    `json:"is_back_to_homepage"`            // 是否返回首页。实测为true
	EnableRcmdGuide            bool    `json:"enable_rcmd_guide"`              // 是否启用推荐引导。实测为true
	InlineSound                int     `json:"inline_sound"`                   // 内嵌声音？实测为2
	AutoRefreshTimeByAppear    int     `json:"auto_refresh_time_by_appear"`    // 可见时的自动刷新时间。单位毫秒
	AutoRefreshTimeByActive    int     `json:"auto_refresh_time_by_active"`    // 活跃时的自动刷新时间。单位毫秒
	TriggerLoadmoreLeftLineNum int     `json:"trigger_loadmore_left_line_num"` // 未知。响应示例中为-1
	HistoryCacheSize           int     `json:"history_cache_size"`             // 未知。响应示例中为10
	VisibleArea                int     `json:"visible_area"`                   // 可见区域？实测为80
	CardDensityExp             int     `json:"card_density_exp"`               // 卡片密度实验分组？实测为1
	SmallCoverWHRatio          float64 `json:"small_cover_wh_ratio"`           // 小封面宽高比。实测为1.33
	VideoMode                  int     `json:"video_mode"`                     // 未知。响应示例中为1
	SpaceEnlargeExp            int     `json:"space_enlarge_exp"`              // 空间扩展实验分组？实测为1
	AutoRefreshTimeByBehavior  int     `json:"auto_refresh_time_by_behavior"`  // 未知。响应示例中为5
	StoryModeV2GuideExp        int     `json:"story_mode_v2_guide_exp"`        // 故事模式v2引导实验分组？实测为6
	AutoRefreshByBehavior      int     `json:"auto_refresh_by_behavior"`       // 未知。响应示例中为1
	ThreePointStyle            int     `json:"three_point_style"`              // 未知。响应示例中为1
	ExposureDurationStartRatio float64 `json:"exposure_duration_start_ratio"`  // 未知。响应示例中为0.8
	ExposureDurationEndRatio   float64 `json:"exposure_duration_end_ratio"`    // 未知。响应示例中为0.8
	ExposureDurationMinMs      int     `json:"exposure_duration_min_ms"`       // 未知。响应示例中为1
	PegasusCode                int     `json:"pegasus_code"`                   // 飞马实验代号？实测为-1200
	RcmdLabelMngEntrance       int     `json:"rcmd_label_mng_entrance"`        // 推荐标签管理入口？实测为0
}

// GetShortVideoListItem 短视频模式视频列表中的一项。
//
// 列表项既可能是普通视频，也可能是横幅（static/ad），因此同时包含
// 视频字段与横幅条目字段。
type GetShortVideoListItem struct {
	CanPlay                      int                                   `json:"can_play"`                         // 是否可播放。实测为1
	CardGoto                     string                                `json:"card_goto"`                        // 卡片跳转类型。视频为av，横幅为banner
	CardType                     string                                `json:"card_type"`                        // 卡片类型。视频为small_cover_v2，横幅为banner_v8
	Cover                        string                                `json:"cover"`                            // 封面url
	CoverLeft1ContentDescription string                                `json:"cover_left_1_content_description"` // 播放量。如8.9万观看
	CoverLeft2ContentDescription string                                `json:"cover_left_2_content_description"` // 弹幕数。如250弹幕
	CoverLeftText1               string                                `json:"cover_left_text_1"`                // 播放量。如8.9万
	CoverLeftText2               string                                `json:"cover_left_text_2"`                // 弹幕数。如250
	CoverRightContentDescription string                                `json:"cover_right_content_description"`  // 视频长度。如1分钟20秒
	CoverRightText               string                                `json:"cover_right_text"`                 // 视频长度。如1:20
	DescButton                   GetShortVideoListItemDescButton       `json:"desc_button"`                      // up主信息
	Param                        string                                `json:"param"`                            // 视频aid
	PlayerArgs                   GetShortVideoListItemPlayerArgs       `json:"player_args"`                      // 视频信息
	TalkBack                     string                                `json:"talk_back"`                        // 语音播报文本
	Title                        string                                `json:"title"`                            // 标题
	URI                          string                                `json:"uri"`                              // 跳转链接
	GotoIcon                     GetShortVideoListItemGotoIcon         `json:"goto_icon"`                        // 视频卡片左下角图标。普通视频时出现
	RcmdReasonStyle              *GetShortVideoListItemRcmdReasonStyle `json:"rcmd_reason_style"`                // 左下角的'竖屏'or'2千点赞'。热门视频时出现
	BannerItem                   []GetShortVideoListItem               `json:"banner_item"`                      // 横幅列表。登录后使用accessKey出现
	Goto                         string                                `json:"goto"`                             // 跳转类型。实测为av
	ThreePoint                   *GetShortVideoListItemThreePoint      `json:"three_point"`                      // 三点菜单
	Args                         GetShortVideoListItemArgs             `json:"args"`                             // 视频参数
	ThreePointV2                 []GetShortVideoListItemThreePointV2   `json:"three_point_v2"`                   // 三点菜单v2
	ThreePointV                  string                                `json:"three_point_v"`                    // 三点菜单版本。实测为v2
	CoverLeftIcon1               int                                   `json:"cover_left_icon_1"`                // 封面左上角图标1。实测为1
	CoverLeftIcon2               int                                   `json:"cover_left_icon_2"`                // 封面左上角图标2。实测为3
	OfficialIcon                 int                                   `json:"official_icon"`                    // 官方图标。实测为16
	CoverInfoPriority            int                                   `json:"cover_info_priority"`              // 封面信息优先级。实测为123
	Idx                          int64                                 `json:"idx"`                              // 未知。响应示例中为秒级时间戳
	RcmdReason                   string                                `json:"rcmd_reason"`                      // 推荐理由文字。如3万点赞
	ReportFlowData               string                                `json:"report_flow_data"`                 // 报告流数据。伪装成obj的str
	TrackID                      string                                `json:"track_id"`                         // 路由track
	Type                         string                                `json:"type"`                             // 横幅条目类型。static：静态横幅。ad：广告
	ResourceID                   int                                   `json:"resource_id"`                      // 横幅资源id
	ID                           int                                   `json:"id"`                               // 横幅id
	Index                        int                                   `json:"index"`                            // 横幅在列表中的索引
	StaticBanner                 *GetShortVideoListItemStaticBanner    `json:"static_banner"`                    // 静态横幅信息。type为static时出现
	AdBanner                     any                                   `json:"ad_banner"`                        // 广告横幅信息。type为ad时出现，结构复杂
}

// GetShortVideoListItemDescButton 短视频卡片中的UP主信息按钮。
type GetShortVideoListItemDescButton struct {
	Event string `json:"event"` // 事件名
	Text  string `json:"text"`  // up名称
	Type  int    `json:"type"`  // 类型。实测为1
	URI   string `json:"uri"`   // 跳转链接
}

// GetShortVideoListItemPlayerArgs 短视频卡片的播放器参数。
type GetShortVideoListItemPlayerArgs struct {
	Aid      int64  `json:"aid"`      // 视频aid
	Cid      int64  `json:"cid"`      // 视频cid
	Duration int    `json:"duration"` // 视频长度。单位为秒
	Type     string `json:"type"`     // 内容类型
}

// GetShortVideoListItemThreePoint 三点菜单。
type GetShortVideoListItemThreePoint struct {
	DislikeReasons []GetShortVideoListItemReason `json:"dislike_reasons"` // 不感兴趣理由列表。每项结构与three_point_v2中的reasons相同
	Feedbacks      []GetShortVideoListItemReason `json:"feedbacks"`       // 反馈理由列表。每项结构与three_point_v2中的reasons相同
	WatchLater     int                           `json:"watch_later"`     // 稍后再看。1：显示。实测为1
}

// GetShortVideoListItemArgs 短视频卡片中的视频参数。
type GetShortVideoListItemArgs struct {
	UpID   int    `json:"up_id"`   // UP主mid
	UpName string `json:"up_name"` // UP主昵称
	Rid    int    `json:"rid"`     // 分区id。响应示例中出现
	Rname  string `json:"rname"`   // 分区名。响应示例中出现
	Tid    int    `json:"tid"`     // 分区id
	Tname  string `json:"tname"`   // 分区名
	Aid    int64  `json:"aid"`     // 视频aid
}

// GetShortVideoListItemThreePointV2 三点菜单v2中的一个按钮。
type GetShortVideoListItemThreePointV2 struct {
	Title     string                        `json:"title"`      // 按钮文字。如添加至稍后再看、反馈、不感兴趣
	Type      string                        `json:"type"`       // 按钮类型。watch_later：稍后再看。feedback：反馈。dislike：不感兴趣
	Icon      string                        `json:"icon"`       // 图标url。仅watch_later类型出现
	IconNight string                        `json:"icon_night"` // 夜间图标url。仅watch_later类型出现
	Subtitle  string                        `json:"subtitle"`   // 按钮副标题。仅feedback与dislike类型出现
	Reasons   []GetShortVideoListItemReason `json:"reasons"`    // 可选择的原因列表。仅feedback与dislike类型出现
}

// GetShortVideoListItemReason 三点菜单中的不感兴趣或反馈理由。
type GetShortVideoListItemReason struct {
	ID     int    `json:"id"`     // 原因id
	Name   string `json:"name"`   // 原因名称。如恐怖血腥
	Toast  string `json:"toast"`  // 提示文字。如将减少相似内容推荐
	Extend string `json:"extend"` // 附加信息。JSON字符串，实测含tid
}

// GetShortVideoListItemGotoIcon 视频卡片左下角图标。
type GetShortVideoListItemGotoIcon struct {
	IconURL      string `json:"icon_url"`       // 图标链接
	IconNightURL string `json:"icon_night_url"` // 夜间图标链接。应该跟上面那个一样
	IconWidth    int    `json:"icon_width"`     // 宽度。如16
	IconHeight   int    `json:"icon_height"`    // 高度。如16
}

// GetShortVideoListItemRcmdReasonStyle 推荐理由的样式。
type GetShortVideoListItemRcmdReasonStyle struct {
	Text             string `json:"text"`               // 文字信息。如1万点赞、竖屏
	TextColor        string `json:"text_color"`         // 字体颜色
	BgColor          string `json:"bg_color"`           // 文本的背景颜色
	BorderColor      string `json:"border_color"`       // 边框颜色
	TextColorNight   string `json:"text_color_night"`   // 夜间字体颜色
	BgColorNight     string `json:"bg_color_night"`     // 夜间背景颜色
	BorderColorNight string `json:"border_color_night"` // 夜间边框颜色
	BgStyle          int    `json:"bg_style"`           // 背景样式。响应示例中出现
}

// GetShortVideoListItemStaticBanner 静态横幅信息。
type GetShortVideoListItemStaticBanner struct {
	ID         int    `json:"id"`          // 横幅id
	Title      string `json:"title"`       // 横幅标题
	Image      string `json:"image"`       // 横幅图片url
	Hash       string `json:"hash"`        // 图片hash
	URI        string `json:"uri"`         // 跳转链接
	RequestID  string `json:"request_id"`  // 请求id
	SrcID      int    `json:"src_id"`      // 资源id
	IsAdLoc    bool   `json:"is_ad_loc"`   // 是否为广告位
	ClientIP   string `json:"client_ip"`   // 客户端ip
	ServerType int    `json:"server_type"` // 服务端类型
	ResourceID int    `json:"resource_id"` // 资源id
	Index      int    `json:"index"`       // 在列表中的索引
	CmMark     int    `json:"cm_mark"`     // 广告标记
}
