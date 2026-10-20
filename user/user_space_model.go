package user

// 用户空间相关响应模型。
//
// 本文件的类型只服务于用户空间各接口，不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// GetUserSpaceTopVideoResult 是「查询用户置顶视频」的 data 对象。
type GetUserSpaceTopVideoResult struct {
	Aid          int                           `json:"aid"`            // 稿件avid
	Videos       int                           `json:"videos"`         // 视频分P总数。默认为1
	Tid          int                           `json:"tid"`            // 分区tid
	Tname        string                        `json:"tname"`          // 子分区名称
	Copyright    int                           `json:"copyright"`      // 是否转载。1：原创。2：转载。3：未填写
	Pic          string                        `json:"pic"`            // 视频封面图片url
	Title        string                        `json:"title"`          // 稿件标题
	PubDate      int                           `json:"pubdate"`        // 稿件发布时间。时间戳
	Ctime        int                           `json:"ctime"`          // 用户提交稿件的时间。时间戳
	Desc         string                        `json:"desc"`           // 视频简介
	State        int                           `json:"state"`          // 视频状态
	Attribute    int                           `json:"attribute"`      // 稿件属性位配置
	Duration     int                           `json:"duration"`       // 视频总计持续时长（所有分P）。单位为秒
	Rights       GetUserSpaceTopVideoRights    `json:"rights"`         // 视频属性标志
	Owner        GetUserSpaceTopVideoOwner     `json:"owner"`          // 视频UP主信息
	Stat         GetUserSpaceTopVideoStat      `json:"stat"`           // 视频状态数
	Dynamic      string                        `json:"dynamic"`        // 视频同步发布的动态的文字内容。无为空
	Cid          int                           `json:"cid"`            // 视频1P cid
	Dimension    GetUserSpaceTopVideoDimension `json:"dimension"`      // 视频1P分辨率
	Bvid         string                        `json:"bvid"`           // 稿件bvid
	Reason       string                        `json:"reason"`         // 置顶视频备注
	InterVideo   bool                          `json:"inter_video"`    // 是否为合作视频
	ElecArcBadge string                        `json:"elec_arc_badge"` // (?)。实测为空
	ViewSelfType int                           `json:"view_self_type"` // (?)。实测为0
}

// GetUserSpaceTopVideoRights 是置顶视频的 rights 对象。
type GetUserSpaceTopVideoRights struct {
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

// GetUserSpaceTopVideoOwner 是置顶视频的 owner 对象。
type GetUserSpaceTopVideoOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetUserSpaceTopVideoStat 是置顶视频的 stat 对象。
type GetUserSpaceTopVideoStat struct {
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

// GetUserSpaceTopVideoDimension 是置顶视频的 dimension 对象。
type GetUserSpaceTopVideoDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetUserSpaceMasterpieceListItem 是代表作视频列表中的一项。
type GetUserSpaceMasterpieceListItem struct {
	Aid          int                                      `json:"aid"`            // 稿件avid
	Videos       int                                      `json:"videos"`         // 视频分P总数。默认为1
	Tid          int                                      `json:"tid"`            // 分区tid
	Tname        string                                   `json:"tname"`          // 子分区名称
	Copyright    int                                      `json:"copyright"`      // 是否转载。1：原创。2：转载。3：未填写
	Pic          string                                   `json:"pic"`            // 视频封面图片url
	Title        string                                   `json:"title"`          // 稿件标题
	PubDate      int                                      `json:"pubdate"`        // 稿件发布时间。时间戳
	Ctime        int                                      `json:"ctime"`          // 用户提交稿件的时间。时间戳
	Desc         string                                   `json:"desc"`           // 视频简介
	State        int                                      `json:"state"`          // 视频状态
	Attribute    int                                      `json:"attribute"`      // 稿件属性位配置
	Duration     int                                      `json:"duration"`       // 视频总计持续时长（所有分P）。单位为秒
	Rights       GetUserSpaceMasterpieceListItemRights    `json:"rights"`         // 视频属性标志
	Owner        GetUserSpaceMasterpieceListItemOwner     `json:"owner"`          // 视频UP主信息
	Stat         GetUserSpaceMasterpieceListItemStat      `json:"stat"`           // 视频状态数
	Dynamic      string                                   `json:"dynamic"`        // 视频同步发布的动态的文字内容。无为空
	Cid          int                                      `json:"cid"`            // 视频1P cid
	Dimension    GetUserSpaceMasterpieceListItemDimension `json:"dimension"`      // 视频1P分辨率
	Bvid         string                                   `json:"bvid"`           // 稿件bvid
	Reason       string                                   `json:"reason"`         // 置顶视频备注
	InterVideo   bool                                     `json:"inter_video"`    // 是否为合作视频
	ElecArcBadge string                                   `json:"elec_arc_badge"` // (?)。实测为空
	ViewSelfType int                                      `json:"view_self_type"` // (?)。实测为0
}

// GetUserSpaceMasterpieceListItemRights 是代表作视频项的 rights 对象。
type GetUserSpaceMasterpieceListItemRights struct {
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

// GetUserSpaceMasterpieceListItemOwner 是代表作视频项的 owner 对象。
type GetUserSpaceMasterpieceListItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetUserSpaceMasterpieceListItemStat 是代表作视频项的 stat 对象。
type GetUserSpaceMasterpieceListItemStat struct {
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

// GetUserSpaceMasterpieceListItemDimension 是代表作视频项的 dimension 对象。
type GetUserSpaceMasterpieceListItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetUserSpaceTagsItem 是用户个人TAG列表中的一项。
type GetUserSpaceTagsItem struct {
	Mid  int      `json:"mid"`  // 目标用户mid
	Tags []string `json:"tags"` // TAG名称。上限5条
}

// GetUserSpaceSettingsResult 是「查询空间设置」的 data 对象。
type GetUserSpaceSettingsResult struct {
	Privacy             GetUserSpaceSettingsPrivacy      `json:"privacy"`                // 空间隐私权限
	IndexOrder          []GetUserSpaceSettingsIndexOrder `json:"index_order"`            // 空间板块布局
	Theme               string                           `json:"theme"`                  // default主题？作用尚不明确
	ThemePreviewImgPath string                           `json:"theme_preview_img_path"` // 主题预览图路径？作用尚不明确
	Toutu               GetUserSpaceSettingsToutu        `json:"toutu"`                  // 空间头图
}

// GetUserSpaceSettingsPrivacy 是空间设置中的隐私权限。0：隐藏。1：公开
type GetUserSpaceSettingsPrivacy struct {
	Bangumi    int `json:"bangumi"`     // 追番及追剧
	Bbq        int `json:"bbq"`         // 轻视频
	Channel    int `json:"channel"`     // 频道
	CoinsVideo int `json:"coins_video"` // 最近投币的视频
	Comic      int `json:"comic"`       // 追漫
	DressUp    int `json:"dress_up"`    // 装扮
	FavVideo   int `json:"fav_video"`   // 收藏夹
	Groups     int `json:"groups"`      // 圈子？作用尚不明确
	LikesVideo int `json:"likes_video"` // 最近点赞的视频
	PlayedGame int `json:"played_game"` // 最近玩过的游戏
	Tags       int `json:"tags"`        // 订阅标签
	UserInfo   int `json:"user_info"`   // 个人资料
}

// GetUserSpaceSettingsIndexOrder 是空间设置中的板块布局项。按先左侧后右侧的顺序排列
type GetUserSpaceSettingsIndexOrder struct {
	ID   int    `json:"id"`   // 板块编号
	Name string `json:"name"` // 板块名称
}

// GetUserSpaceSettingsToutu 是空间设置中的空间头图。
type GetUserSpaceSettingsToutu struct {
	Sid          int    `json:"sid"`           // 空间头图ID
	Expire       int64  `json:"expire"`        // 到期时间？时间戳？作用尚不明确
	SImg         string `json:"s_img"`         // 空间头图小图相对路径
	LImg         string `json:"l_img"`         // 空间头图相对路径
	AndroidImg   string `json:"android_img"`   // 空。手机端头图与web端不同，需用另一个api获取
	IphoneImg    string `json:"iphone_img"`    // 空
	IpadImg      string `json:"ipad_img"`      // 空
	ThumbnailImg string `json:"thumbnail_img"` // 缩略图
	Platform     int    `json:"platform"`      // 0。作用尚不明确
}

// GetUserSpaceTopPhotoListItem 是可用空间头图列表中的一项。
type GetUserSpaceTopPhotoListItem struct {
	ID           int    `json:"id"`            // 空间头图ID
	ProductName  string `json:"product_name"`  // 显示名称
	Price        int    `json:"price"`         // 价格
	CoinType     int    `json:"coin_type"`     // 支付类型？
	VipFree      int    `json:"vip_free"`      // 大会员免费
	SImg         string `json:"s_img"`         // 小图URI。需自行与 i0.hdslb.com 拼接成URL
	LImg         string `json:"l_img"`         // 大图URI。同 s_img
	ThumbnailImg string `json:"thumbnail_img"` // 空。并不存在的缩略图？
	SortNum      int    `json:"sort_num"`      // 排序编号
	IsDisable    int    `json:"is_disable"`    // 已禁用。0：未禁用
	Expire       int64  `json:"expire"`        // 过期时间？UNIX时间戳，或0为永不过期
	Had          int    `json:"had"`           // 是否拥有？当 expire 不为0时为1，否则为0
}

// GetUserSpaceLastPlayGameResult 是「查询用户最近玩过的游戏V2」的 data 对象。
type GetUserSpaceLastPlayGameResult struct {
	PageNum    int                            `json:"page_num"`    // 当前页码
	PageSize   int                            `json:"page_size"`   // 每页项数
	TotalCount int                            `json:"total_count"` // 游戏总数
	List       []GetUserSpaceLastPlayGameItem `json:"list"`        // 游戏列表
}

// GetUserSpaceLastPlayGameItem 是最近玩过的游戏列表中的一项。
type GetUserSpaceLastPlayGameItem struct {
	GameBaseID   int      `json:"game_base_id"`   // 游戏id。biligame 的游戏 id
	GameName     string   `json:"game_name"`      // 游戏名
	GameIcon     string   `json:"game_icon"`      // 游戏图标url
	Grade        float64  `json:"grade"`          // 游戏评分。如8.2、6.8
	DetailURL    string   `json:"detail_url"`     // 游戏主页url
	GameTags     []string `json:"game_tags"`      // 游戏标签
	Notice       string   `json:"notice"`         // 游戏简介文案
	GiftTitle    string   `json:"gift_title"`     // 游戏礼物文案
	GameStatusV2 int      `json:"game_status_v2"` // (?)
}

// GetUserSpaceCoinVideoListAppResult 是「查询用户最近投币视频（APP）」的 data 对象。
type GetUserSpaceCoinVideoListAppResult struct {
	Count int                                `json:"count"` // 投币的视频数
	Item  []GetUserSpaceCoinVideoListAppItem `json:"item"`  // 投币视频列表
}

// GetUserSpaceCoinVideoListAppItem 是 APP 端投币视频列表中的一项。
type GetUserSpaceCoinVideoListAppItem struct {
	Title           string `json:"title"`             // 稿件标题
	Subtitle        string `json:"subtitle"`          // 空
	Tname           string `json:"tname"`             // 空
	Cover           string `json:"cover"`             // 封面图片url
	CoverIcon       string `json:"cover_icon"`        // 空
	URI             string `json:"uri"`               // APP跳转uri。如 bilibili://video/2
	Param           string `json:"param"`             // 稿件avid
	Goto            string `json:"goto"`              // av
	Length          string `json:"length"`            // 空。文档标为 num，示例返回空字符串
	Duration        int    `json:"duration"`          // 稿件视频长度。单位为秒
	IsPopular       bool   `json:"is_popular"`        // (?)
	IsSteins        bool   `json:"is_steins"`         // (?)
	IsUgcpay        bool   `json:"is_ugcpay"`         // (?)
	IsCooperation   bool   `json:"is_cooperation"`    // (?)。文档标为 str，示例为 bool
	IsPgc           bool   `json:"is_pgc"`            // (?)。文档标为 str，示例为 bool
	IsLivePlayback  bool   `json:"is_live_playback"`  // (?)。文档标为 str，示例为 bool
	IsPugv          bool   `json:"is_pugv"`           // (?)。文档标为 str，示例为 bool
	IsFold          bool   `json:"is_fold"`           // (?)。文档标为 num，示例为 bool
	Play            int    `json:"play"`              // 播放量
	Danmaku         int    `json:"danmaku"`           // 弹幕量
	Ctime           int    `json:"ctime"`             // 发布时间。时间戳
	UgcPay          int    `json:"ugc_pay"`           // 0
	Author          string `json:"author"`            // UP主昵称
	State           bool   `json:"state"`             // true
	Videos          int    `json:"videos"`            // 0
	ViewContent     string `json:"view_content"`      // 稿件播放量文案
	IconType        int    `json:"icon_type"`         // 0
	PublishTimeText string `json:"publish_time_text"` // 空
}

// GetUserSpaceLikeVideoListResult 是「查询用户最近点赞视频（Web）」的 data 对象。
//
// 文档字段表把 data 写成数组，但响应示例为 {"list": [...]}，这里按示例建模。
type GetUserSpaceLikeVideoListResult struct {
	List []GetUserSpaceLikeVideoListItem `json:"list"` // 点赞视频列表
}

// GetUserSpaceLikeVideoListAppResult 是「查询用户最近点赞视频（APP）」的 data 对象。
type GetUserSpaceLikeVideoListAppResult struct {
	Count int                                `json:"count"` // 点赞的视频数
	Item  []GetUserSpaceLikeVideoListAppItem `json:"item"`  // 点赞视频列表
}

// GetUserSpaceLikeVideoListAppItem 是 APP 端点赞视频列表中的一项。
type GetUserSpaceLikeVideoListAppItem struct {
	Title           string `json:"title"`             // 稿件标题
	Subtitle        string `json:"subtitle"`          // 空
	Tname           string `json:"tname"`             // 空
	Cover           string `json:"cover"`             // 封面图片url
	CoverIcon       string `json:"cover_icon"`        // 空
	URI             string `json:"uri"`               // APP跳转uri。如 bilibili://video/2
	Param           string `json:"param"`             // 稿件avid
	Goto            string `json:"goto"`              // av
	Length          string `json:"length"`            // 空。文档标为 num，示例返回空字符串
	Duration        int    `json:"duration"`          // 稿件视频长度。单位为秒
	IsPopular       bool   `json:"is_popular"`        // (?)
	IsSteins        bool   `json:"is_steins"`         // (?)
	IsUgcpay        bool   `json:"is_ugcpay"`         // (?)
	IsCooperation   bool   `json:"is_cooperation"`    // (?)。文档标为 str，示例为 bool
	IsPgc           bool   `json:"is_pgc"`            // (?)。文档标为 str，示例为 bool
	IsLivePlayback  bool   `json:"is_live_playback"`  // (?)。文档标为 str，示例为 bool
	IsPugv          bool   `json:"is_pugv"`           // (?)。文档标为 str，示例为 bool
	IsFold          bool   `json:"is_fold"`           // (?)。文档标为 num，示例为 bool
	Play            int    `json:"play"`              // 播放量
	Danmaku         int    `json:"danmaku"`           // 弹幕量
	Ctime           int    `json:"ctime"`             // 发布时间。时间戳
	UgcPay          int    `json:"ugc_pay"`           // 0
	Author          string `json:"author"`            // UP主昵称
	State           bool   `json:"state"`             // true
	Videos          int    `json:"videos"`            // 0
	ViewContent     string `json:"view_content"`      // 稿件播放量文案
	IconType        int    `json:"icon_type"`         // 0
	PublishTimeText string `json:"publish_time_text"` // 空
}

// GetUserSpaceArchiveCursorResult 是「查询用户投稿明细（APP）」的 data 对象。
type GetUserSpaceArchiveCursorResult struct {
	Count              int                              `json:"count"`                // 总数
	EpisodicButton     GetUserSpaceArchiveCursorButton  `json:"episodic_button"`      // 播放按钮
	HasNext            bool                             `json:"has_next"`             // 有无后续
	HasPrev            bool                             `json:"has_prev"`             // 有无前置
	Item               []GetUserSpaceArchiveCursorItem  `json:"item"`                 // 视频列表。注意是 item 没有复数
	LastWatchedLocator GetUserSpaceArchiveCursorLocator `json:"last_watched_locator"` // 上次播放定位
	Order              []GetUserSpaceArchiveCursorOrder `json:"order"`                // 排列方式
}

// GetUserSpaceArchiveCursorButton 是投稿明细中播放按钮的显示文字与跳转链接。
type GetUserSpaceArchiveCursorButton struct {
	Text string `json:"text"` // 四个大字“播放全部”
	URI  string `json:"uri"`  // 跳转链接
}

// GetUserSpaceArchiveCursorItem 是投稿明细中的视频项。
type GetUserSpaceArchiveCursorItem struct {
	Author          string                                    `json:"author"`            // 作者昵称
	Bvid            string                                    `json:"bvid"`              // 视频bv号
	Cover           string                                    `json:"cover"`             // 封面url
	CoverIcon       string                                    `json:"cover_icon"`        // 封面icon？未知
	Ctime           int                                       `json:"ctime"`             // 发布时间。秒级时间戳
	CursorAttr      GetUserSpaceArchiveCursorItemAttr         `json:"cursor_attr"`       // 上次观看标记
	Danmaku         int                                       `json:"danmaku"`           // 弹幕数
	Duration        int                                       `json:"duration"`          // 视频时长
	FirstCid        int                                       `json:"first_cid"`         // 首个频道cid
	Goto            string                                    `json:"goto"`              // 跳转去向类型。已知av为视频
	IconType        int                                       `json:"icon_type"`         // icon类型？未知
	IsCooperation   bool                                      `json:"is_cooperation"`    // 是否为合作视频
	IsFold          bool                                      `json:"is_fold"`           // 是否折叠？未知
	IsLivePlayback  bool                                      `json:"is_live_playback"`  // 是否为直播回放
	IsOnself        bool                                      `json:"is_onself"`         // 是否自己？未知。文档字段表的写法
	IsOneself       bool                                      `json:"is_oneself"`        // 是否自己？未知。响应示例的写法
	IsPgc           bool                                      `json:"is_pgc"`            // 是否开pgc
	IsPopular       bool                                      `json:"is_popular"`        // 是否上热门
	IsPugv          bool                                      `json:"is_pugv"`           // 未知
	IsSteins        bool                                      `json:"is_steins"`         // 未知
	IsUgcpay        bool                                      `json:"is_ugcpay"`         // 未知
	Length          string                                    `json:"length"`            // 长度？未知。文档标为 num，示例返回空字符串
	Param           string                                    `json:"param"`             // 参数。实际为视频aid
	Play            int                                       `json:"play"`              // 播放量
	PublishTimeText string                                    `json:"publish_time_text"` // 发布时间汉字字符串版
	State           bool                                      `json:"state"`             // 状态？未知
	Subtitle        string                                    `json:"subtitle"`          // 子标题
	ThreePoints     []GetUserSpaceArchiveCursorItemThreePoint `json:"three_points"`      // 三点按钮内容。文档字段表的写法
	ThreePoint      []GetUserSpaceArchiveCursorItemThreePoint `json:"three_point"`       // 三点按钮内容。响应示例的写法
	Title           string                                    `json:"title"`             // 标题
	Tname           string                                    `json:"tname"`             // tag的name
	UgcPay          int                                       `json:"ugc_pay"`           // ugc支付？未知
	URI             string                                    `json:"uri"`               // APP跳转uri
	Video           int                                       `json:"video"`             // 未知
	ViewContent     string                                    `json:"view_content"`      // 播放量字符串版
}

// GetUserSpaceArchiveCursorItemThreePoint 是投稿明细项的三点按钮内容（稍后再看、分享）。
type GetUserSpaceArchiveCursorItemThreePoint struct {
	Type           string `json:"type"`             // 类型。如 addtoview、share
	Icon           string `json:"icon"`             // 图标url
	Text           string `json:"text"`             // 显示文字
	ShareSuccToast string `json:"share_succ_toast"` // 分享成功提示
	ShareFailToast string `json:"share_fail_toast"` // 分享失败提示
	SharePath      string `json:"share_path"`       // 分享路径
	ShortLink      string `json:"short_link"`       // 短链接
}

// GetUserSpaceArchiveCursorItemAttr 是投稿明细项的上次观看标记。
type GetUserSpaceArchiveCursorItemAttr struct {
	IsLastWatchedArc bool `json:"is_last_watched_arc"` // 是否为上次播放
	Rank             int  `json:"rank"`                // 视频所处该列表排名
}

// GetUserSpaceArchiveCursorLocator 是投稿明细的上次播放定位。
type GetUserSpaceArchiveCursorLocator struct {
	DisplayThreshold int    `json:"display_threshold"` // 显示阈值
	InsertRanking    int    `json:"insert_ranking"`    // 嵌入排名
	Text             string `json:"text"`              // 定位上次观看字样
}

// GetUserSpaceArchiveCursorOrder 是投稿明细的排列方式。
type GetUserSpaceArchiveCursorOrder struct {
	Title string `json:"title"` // 显示文字。如最新发布、最多播放
	Value string `json:"value"` // 值。如 pubdate、click
}

// GetUserSpaceAudioListResult 是「查询用户投稿音频明细」的 data 对象。字段名沿用接口的驼峰命名
type GetUserSpaceAudioListResult struct {
	CurPage   int                         `json:"curPage"`   // 当前页码
	PageCount int                         `json:"pageCount"` // 总页数
	TotalSize int                         `json:"totalSize"` // 总计数
	PageSize  int                         `json:"pageSize"`  // 每页项数
	Data      []GetUserSpaceAudioListItem `json:"data"`      // 音频列表
}

// GetUserSpaceAudioListItem 是用户投稿音频列表中的一项。
type GetUserSpaceAudioListItem struct {
	ID         int                                `json:"id"`         // 音频AU号
	UID        int                                `json:"uid"`        // 用户mid
	Uname      string                             `json:"uname"`      // 用户名
	Author     any                                `json:"author"`     // null
	Title      string                             `json:"title"`      // 音频标题
	Cover      string                             `json:"cover"`      // 音频封面图片URL
	Intro      any                                `json:"intro"`      // null
	Lyric      string                             `json:"lyric"`      // 音频歌词URL
	Crtype     int                                `json:"crtype"`     // 1
	Duration   int                                `json:"duration"`   // 音频时长。单位为秒
	Passtime   int                                `json:"passtime"`   // 发布时间。UNIX秒级时间戳
	Curtime    int                                `json:"curtime"`    // 0
	Aid        any                                `json:"aid"`        // null
	Cid        any                                `json:"cid"`        // null
	Msid       any                                `json:"msid"`       // null
	Attr       any                                `json:"attr"`       // null
	Limit      any                                `json:"limit"`      // null
	ActivityID any                                `json:"activityId"` // null
	Limitdesc  any                                `json:"limitdesc"`  // null
	Ctime      int64                              `json:"ctime"`      // 发布时间。UNIX毫秒时间戳
	Statistic  GetUserSpaceAudioListItemStatistic `json:"statistic"`  // 统计数
	CoinNum    int                                `json:"coin_num"`   // 硬币数量
	IsCooper   any                                `json:"is_cooper"`  // null
}

// GetUserSpaceAudioListItemStatistic 是音频项的统计数。
type GetUserSpaceAudioListItemStatistic struct {
	Sid     int `json:"sid"`     // 音频AU号
	Play    int `json:"play"`    // 播放数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

// GetUserSpaceCourseListResult 是「查询用户发布的课程列表」的 data 对象。
type GetUserSpaceCourseListResult struct {
	Items []GetUserSpaceCourseListItem `json:"items"` // 课程列表
	Page  GetUserSpaceCourseListPage   `json:"page"`  // 页面信息
}

// GetUserSpaceCourseListPage 是课程列表的页面信息。
type GetUserSpaceCourseListPage struct {
	Next  bool `json:"next"`  // 是否存在下一页。false：最后一页。true：存在下一页
	Num   int  `json:"num"`   // 当前页码
	Size  int  `json:"size"`  // 每页项数
	Total int  `json:"total"` // 总计项数
}

// GetUserSpaceCourseListItem 是课程列表中的一项。
type GetUserSpaceCourseListItem struct {
	Cover              string                                 `json:"cover"`                // 课程封面url
	EpCount            int                                    `json:"ep_count"`             // 课程集数
	Link               string                                 `json:"link"`                 // 课程页面url。手机版页面url
	Page               int                                    `json:"page"`                 // 1。作用尚不明确
	Play               int                                    `json:"play"`                 // 课程播放数
	SeasonID           int                                    `json:"season_id"`            // 课程ssid
	Status             string                                 `json:"status"`               // 更新状态说明
	Subtitle           string                                 `json:"subtitle"`             // 课程副标题
	Title              string                                 `json:"title"`                // 课程标题
	Cooperated         bool                                   `json:"cooperated"`           // 是否为合作视频。实测为false
	CooperationMark    string                                 `json:"cooperation_mark"`     // 合作标记。实测为空
	Marks              []any                                  `json:"marks"`                // 标记。实测为空数组
	SeasonSelectedRank GetUserSpaceCourseListItemSelectedRank `json:"season_selected_rank"` // 精选排名信息
}

// GetUserSpaceCourseListItemSelectedRank 是课程项的精选排名信息。
type GetUserSpaceCourseListItemSelectedRank struct {
	BadgeType                 int    `json:"badge_type"`                    // 徽章类型。实测为2
	SeasonID                  int    `json:"season_id"`                     // 合集ssid。实测为51
	ShowSeasonSelectedMarkURL string `json:"show_season_selected_mark_url"` // 精选标记图片url
}

// GetUserSpaceBangumiPreviewListResult 是「查询用户追番预览列表」的 data 对象。
type GetUserSpaceBangumiPreviewListResult struct {
	Count  int                                  `json:"count"`  // 总计追番数
	Pages  int                                  `json:"pages"`  // 0。作用尚不明确
	Result []GetUserSpaceBangumiPreviewListItem `json:"result"` // 追番预览列表
}

// GetUserSpaceBangumiPreviewListItem 是追番预览列表中的一项。
type GetUserSpaceBangumiPreviewListItem struct {
	Brief         string `json:"brief"`           // 简介
	Cover         string `json:"cover"`           // 封面图片url
	Evaluate      string `json:"evaluate"`        // 空
	Favorites     int    `json:"favorites"`       // 追番数
	IsFinish      int    `json:"is_finish"`       // 是否已完结。0：未完结。1：已完结
	LastEpIndex   int    `json:"last_ep_index"`   // 0。作用尚不明确
	NewestEpIndex int    `json:"newest_ep_index"` // 最新一话。可能为0
	SeasonID      string `json:"season_id"`       // 番剧ssid
	ShareURL      string `json:"share_url"`       // 播放页面链接url
	Title         string `json:"title"`           // 标题
	TotalCount    int    `json:"total_count"`     // 总计集数。未完结：-1。已完结：非0正整数
}

// GetUserSpaceBangumiFollowListResult 是「查询用户追番（追剧）明细」的 data 对象。
type GetUserSpaceBangumiFollowListResult struct {
	List  []GetUserSpaceBangumiFollowListItem `json:"list"`  // 追番列表
	Pn    int                                 `json:"pn"`    // 当前页码
	Ps    int                                 `json:"ps"`    // 每页项数
	Total int                                 `json:"total"` // 总计追番数
}

// GetUserSpaceBangumiFollowListItem 是追番明细列表中的一项。基本同番剧详细信息中的 result 对象
type GetUserSpaceBangumiFollowListItem struct {
	SeasonID       int                                        `json:"season_id"`        // 番剧ssid
	MediaID        int                                        `json:"media_id"`         // 番剧mdid
	SeasonType     int                                        `json:"season_type"`      // 剧集类型
	SeasonTypeName string                                     `json:"season_type_name"` // 剧集类型名称
	Title          string                                     `json:"title"`            // 标题
	Cover          string                                     `json:"cover"`            // 封面图片url
	TotalCount     int                                        `json:"total_count"`      // 总计集数
	IsFinish       int                                        `json:"is_finish"`        // 是否已完结
	IsStarted      int                                        `json:"is_started"`       // 是否已开播
	IsPlay         int                                        `json:"is_play"`          // 是否可播放
	Badge          string                                     `json:"badge"`            // 角标文案。如会员专享
	BadgeType      int                                        `json:"badge_type"`       // 角标类型
	Rights         GetUserSpaceBangumiFollowListItemRights    `json:"rights"`           // 属性标志
	Stat           GetUserSpaceBangumiFollowListItemStat      `json:"stat"`             // 状态数
	NewEp          GetUserSpaceBangumiFollowListItemNewEp     `json:"new_ep"`           // 最新一话信息
	Rating         GetUserSpaceBangumiFollowListItemRating    `json:"rating"`           // 评分信息
	SquareCover    string                                     `json:"square_cover"`     // 方形封面图片url
	SeasonStatus   int                                        `json:"season_status"`    // 剧集状态
	SeasonTitle    string                                     `json:"season_title"`     // 剧集标题
	BadgeEp        string                                     `json:"badge_ep"`         // 单集角标。如会员
	MediaAttr      int                                        `json:"media_attr"`       // 媒体属性
	SeasonAttr     int                                        `json:"season_attr"`      // 剧集属性
	Evaluate       string                                     `json:"evaluate"`         // 简介
	Areas          []GetUserSpaceBangumiFollowListItemArea    `json:"areas"`            // 地区
	Subtitle       string                                     `json:"subtitle"`         // 副标题
	FirstEp        int                                        `json:"first_ep"`         // 第一话epid
	CanWatch       int                                        `json:"can_watch"`        // 是否可观看
	Series         GetUserSpaceBangumiFollowListItemSeries    `json:"series"`           // 系列信息
	Publish        GetUserSpaceBangumiFollowListItemPublish   `json:"publish"`          // 发布信息
	Mode           int                                        `json:"mode"`             // 作用尚不明确。实测为2
	Section        []GetUserSpaceBangumiFollowListItemSection `json:"section"`          // 花絮、PV、番外等非正片内容
	URL            string                                     `json:"url"`              // 播放页面url
	FollowStatus   int                                        `json:"follow_status"`    // 追番状态
	IsNew          int                                        `json:"is_new"`           // 是否有更新
	Progress       string                                     `json:"progress"`         // 观看进度
	BothFollow     bool                                       `json:"both_follow"`      // 是否同时追番
}

// GetUserSpaceBangumiFollowListItemRights 是追番明细项的属性标志。
type GetUserSpaceBangumiFollowListItemRights struct {
	AllowReview    int `json:"allow_review"`    // 是否允许评论
	IsSelection    int `json:"is_selection"`    // 是否为精选
	SelectionStyle int `json:"selection_style"` // 精选样式
}

// GetUserSpaceBangumiFollowListItemStat 是追番明细项的状态数。
type GetUserSpaceBangumiFollowListItemStat struct {
	Follow       int `json:"follow"`        // 追番数
	View         int `json:"view"`          // 播放数
	Danmaku      int `json:"danmaku"`       // 弹幕数
	Reply        int `json:"reply"`         // 评论数
	Coin         int `json:"coin"`          // 投币数
	SeriesFollow int `json:"series_follow"` // 系列追番数
	SeriesView   int `json:"series_view"`   // 系列播放数
}

// GetUserSpaceBangumiFollowListItemNewEp 是追番明细项的最新一话信息。
type GetUserSpaceBangumiFollowListItemNewEp struct {
	ID        int    `json:"id"`         // 最新一话epid
	IndexShow string `json:"index_show"` // 最新一话显示名称。如全13话
	Cover     string `json:"cover"`      // 封面图片url
	Title     string `json:"title"`      // 标题
	LongTitle string `json:"long_title"` // 长标题
	PubTime   string `json:"pub_time"`   // 发布时间
	Duration  int    `json:"duration"`   // 时长
}

// GetUserSpaceBangumiFollowListItemRating 是追番明细项的评分信息。
type GetUserSpaceBangumiFollowListItemRating struct {
	Score float64 `json:"score"` // 评分
	Count int     `json:"count"` // 评分人数
}

// GetUserSpaceBangumiFollowListItemArea 是追番明细项所属地区。
type GetUserSpaceBangumiFollowListItemArea struct {
	ID   int    `json:"id"`   // 所属地区编号
	Name string `json:"name"` // 所属地区名称
}

// GetUserSpaceBangumiFollowListItemSeries 是追番明细项的系列信息。
type GetUserSpaceBangumiFollowListItemSeries struct {
	SeriesID    int    `json:"series_id"`     // 系列id
	Title       string `json:"title"`         // 系列名
	SeasonCount int    `json:"season_count"`  // 系列季数
	NewSeasonID int    `json:"new_season_id"` // 最新一季ssid
}

// GetUserSpaceBangumiFollowListItemPublish 是追番明细项的发布信息。
type GetUserSpaceBangumiFollowListItemPublish struct {
	PubTime         string `json:"pub_time"`          // 发布时间
	PubTimeShow     string `json:"pub_time_show"`     // 发布时间文字介绍
	ReleaseDate     string `json:"release_date"`      // 放送日期
	ReleaseDateShow string `json:"release_date_show"` // 放送日期文字介绍
}

// GetUserSpaceBangumiFollowListItemSection 是追番明细项的非正片板块。
type GetUserSpaceBangumiFollowListItemSection struct {
	SectionID     int    `json:"section_id"`     // 板块id
	SeasonID      int    `json:"season_id"`      // 剧集ssid
	LimitGroup    int    `json:"limit_group"`    // 限制组
	WatchPlatform int    `json:"watch_platform"` // 观看平台
	Type          int    `json:"type"`           // 板块类型
	Copyright     string `json:"copyright"`      // 版权标志
	Title         string `json:"title"`          // 板块标题
	BanAreaShow   int    `json:"ban_area_show"`  // 是否展示地区限制
}

// GetUserSpaceSubscribedTagsResult 是「查询用户关注的TAG（话题）」的 data 对象。
type GetUserSpaceSubscribedTagsResult struct {
	Tags  []GetUserSpaceSubscribedTagsItem `json:"tags"`  // 关注TAG列表
	Count int                              `json:"count"` // 关注TAG的数
}

// GetUserSpaceSubscribedTagsItem 是关注TAG列表中的一项。
type GetUserSpaceSubscribedTagsItem struct {
	ArchiveCount int    `json:"archive_count"` // 0。作用尚不明确
	Cover        string `json:"cover"`         // TAG图片url。无则为空
	Name         string `json:"name"`          // TAG名
	Notify       int    `json:"notify"`        // 1。作用尚不明确
	TagID        int    `json:"tag_id"`        // tag_id
	UpdatedTs    string `json:"updated_ts"`    // 1970-01-01 08:00:00。作用尚不明确
}
