package emoji

// 表情列表相关响应模型。
//
// 本文件只服务 GetMyEmojiList，类型不与其他接口共享：后续接口返回体
// 发生变更时，改动不会波及别的方法。

// GetMyEmojiListResult 是获取我的表情列表接口的 data 对象。
type GetMyEmojiListResult struct {
	Packages []GetMyEmojiListPackage `json:"packages"` // 表情包列表
	Setting  GetMyEmojiListSetting   `json:"setting"`  // 设置
}

// GetMyEmojiListSetting 是 data.setting 对象。
type GetMyEmojiListSetting struct {
	RecentLimit int    `json:"recent_limit"` // 最近使用表情上限。实测为 150
	Attr        int    `json:"attr"`         // 作用尚不明确。实测为 -1
	FocusPkgID  int    `json:"focus_pkg_id"` // 聚焦表情包 id。实测为 1
	Schema      string `json:"schema"`       // 表情包商城页面 url。实测为 h5 页面地址
}

// GetMyEmojiListPackage 是 data.packages 数组中的表情包。
type GetMyEmojiListPackage struct {
	ID    int                        `json:"id"`    // 表情包 ID
	Text  string                     `json:"text"`  // 表情包名称
	URL   string                     `json:"url"`   // 表情包标志图片 URL
	MTime int64                      `json:"mtime"` // 创建时间，时间戳
	Type  int                        `json:"type"`  // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr  int                        `json:"attr"`  // 未知作用
	Meta  GetMyEmojiListPackageMeta  `json:"meta"`  // 属性信息
	Emote []GetMyEmojiListItem       `json:"emote"` // 表情列表
	Flags GetMyEmojiListPackageFlags `json:"flags"` // 是否添加标志
}

// GetMyEmojiListPackageMeta 是 data.packages[].meta 对象。
type GetMyEmojiListPackageMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetMyEmojiListPackageFlags 是 data.packages[].flags 对象。
type GetMyEmojiListPackageFlags struct {
	Added bool `json:"added"` // 是否已添加,需要登录（SESSDATA）否则恒为false,true：已添加 false：未添加
}

// GetMyEmojiListItem 是 data.packages[].emote 数组中的表情。
type GetMyEmojiListItem struct {
	ID        int                     `json:"id"`         // 表情 ID
	PackageID int                     `json:"package_id"` // 表情包 ID
	Text      string                  `json:"text"`       // 表情转义符或颜文字
	URL       string                  `json:"url"`        // 表情图片 URL 或颜文字
	MTime     int64                   `json:"mtime"`      // 创建时间，时间戳
	Type      int                     `json:"type"`       // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr      int                     `json:"attr"`       // 未知作用
	Meta      GetMyEmojiListItemMeta  `json:"meta"`       // 属性信息
	Flags     GetMyEmojiListItemFlags `json:"flags"`      // 禁用标志
}

// GetMyEmojiListItemMeta 是 data.packages[].emote[].meta 对象。
type GetMyEmojiListItemMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetMyEmojiListItemFlags 是 data.packages[].emote[].flags 对象。
type GetMyEmojiListItemFlags struct {
	NoAccess bool `json:"no_access"` // 是否为禁用 true：禁用
}
