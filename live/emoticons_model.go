package live

// 直播间表情相关响应模型。

// GetLiveEmoticonsResult 是获取直播间表情包的响应 data。
type GetLiveEmoticonsResult struct {
	Data        []GetLiveEmoticonsPackage `json:"data"`         // 表情包数据
	FansBrand   int                       `json:"fans_brand"`   // 品牌标识?
	PurchaseURL *string                   `json:"purchase_url"` // 购买链接?，可能为 null
}

// GetLiveEmoticonsPackage 是一个表情包。
type GetLiveEmoticonsPackage struct {
	CurrentCover          string                     `json:"current_cover"`           // 封面 URL
	Emoticons             []GetLiveEmoticonsEmoticon `json:"emoticons"`               // 表情列表
	PkgDescript           string                     `json:"pkg_descript"`            // 文字描述
	PkgID                 int                        `json:"pkg_id"`                  // 包 ID
	PkgName               string                     `json:"pkg_name"`                // 包名称
	PkgPerm               int                        `json:"pkg_perm"`                // 使用权限?
	PkgType               int                        `json:"pkg_type"`                // 包类型
	RecentlyUsedEmoticons []any                      `json:"recently_used_emoticons"` // 最近使用的表情
	TopShow               GetLiveEmoticonsTopShow    `json:"top_show"`                // 顶部展示信息?
	TopShowRecent         GetLiveEmoticonsTopShow    `json:"top_show_recent"`         // 最近使用的顶部展示信息?
	UnlockIdentity        int                        `json:"unlock_identity"`         // 解锁所需身份标识?
	UnlockNeedGift        int                        `json:"unlock_need_gift"`        // 解锁所需礼物?
}

// GetLiveEmoticonsEmoticon 是表情包中的单个表情。
type GetLiveEmoticonsEmoticon struct {
	BulgeDisplay      int    `json:"bulge_display"`       // 突出展示?
	Descript          string `json:"descript"`            // 描述
	Emoji             string `json:"emoji"`               // 触发关键词?
	EmoticonID        int    `json:"emoticon_id"`         // 表情 ID
	EmoticonUnique    string `json:"emoticon_unique"`     // 表情唯一标识
	EmoticonValueType int    `json:"emoticon_value_type"` // 表情值类型?
	Height            int    `json:"height"`              // 表情图片高度
	Identity          int    `json:"identity"`            // 身份限制标识?
	InPlayerArea      int    `json:"in_player_area"`      // 播放器区域内展示?
	IsDynamic         int    `json:"is_dynamic"`          // 是否为动态表情?
	Perm              int    `json:"perm"`                // 使用权限?
	UnlockNeedGift    int    `json:"unlock_need_gift"`    // 解锁需求礼物?
	UnlockNeedLevel   int    `json:"unlock_need_level"`   // 解锁需求等级?
	UnlockShowColor   string `json:"unlock_show_color"`   // 解锁展示颜色?
	UnlockShowImage   string `json:"unlock_show_image"`   // 解锁展示图片?
	UnlockShowText    string `json:"unlock_show_text"`    // 解锁展示文字?
	URL               string `json:"url"`                 // 表情图片 URL
	Width             int    `json:"width"`               // 表情图片宽度
}

// GetLiveEmoticonsTopShow 是表情包顶部展示信息。
type GetLiveEmoticonsTopShow struct {
	TopLeft  GetLiveEmoticonsTopShowItem `json:"top_left"`  // 左上
	TopRight GetLiveEmoticonsTopShowItem `json:"top_right"` // 右上
}

// GetLiveEmoticonsTopShowItem 是顶部展示信息中的一格。
type GetLiveEmoticonsTopShowItem struct {
	Image string `json:"image"` // 图片?
	Text  string `json:"text"`  // 文字
}
