package emoji

// 表情包相关响应模型。
//
// 每个接口的返回体独立定义，互不复用：后续某个接口的返回结构变更时，
// 不会影响其他方法。同一接口内部字段形状完全相同的两个数组
// （GetAllEmoteList 的 user_panel_packages 与 all_packages）共用元素类型，
// 因为它们同属一个返回体、只会一起变化。

// GetMyEmoteListResult 是获取我的表情包列表接口的 data 对象。
type GetMyEmoteListResult struct {
	EmotePackages []GetMyEmoteListPackage `json:"packages"` // 表情包包列表
}

// GetMyEmoteListPackage 是 data.packages 数组中的表情包。
type GetMyEmoteListPackage struct {
	ID    int                        `json:"id"`    // 表情包 ID
	Text  string                     `json:"text"`  // 表情包名称
	URL   string                     `json:"url"`   // 表情包标志图片 URL
	MTime int64                      `json:"mtime"` // 创建时间，时间戳
	Type  int                        `json:"type"`  // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr  int                        `json:"attr"`  // 未知作用
	Meta  GetMyEmoteListPackageMeta  `json:"meta"`  // 属性信息
	Emote []GetMyEmoteListItem       `json:"emote"` // 表情列表
	Flags GetMyEmoteListPackageFlags `json:"flags"` // 是否添加标志
}

// GetMyEmoteListPackageMeta 是 data.packages[].meta 对象。
type GetMyEmoteListPackageMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetMyEmoteListPackageFlags 是 data.packages[].flags 对象。
type GetMyEmoteListPackageFlags struct {
	Added bool `json:"added"` // 是否已添加,需要登录（SESSDATA）否则恒为false,true：已添加 false：未添加
}

// GetMyEmoteListItem 是 data.packages[].emote 数组中的表情。
type GetMyEmoteListItem struct {
	ID        int                     `json:"id"`         // 表情 ID
	PackageID int                     `json:"package_id"` // 表情包 ID
	Text      string                  `json:"text"`       // 表情转义符或颜文字
	URL       string                  `json:"url"`        // 表情图片 URL 或颜文字
	MTime     int64                   `json:"mtime"`      // 创建时间，时间戳
	Type      int                     `json:"type"`       // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr      int                     `json:"attr"`       // 未知作用
	Meta      GetMyEmoteListItemMeta  `json:"meta"`       // 属性信息
	Flags     GetMyEmoteListItemFlags `json:"flags"`      // 禁用标志
}

// GetMyEmoteListItemMeta 是 data.packages[].emote[].meta 对象。
type GetMyEmoteListItemMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetMyEmoteListItemFlags 是 data.packages[].emote[].flags 对象。
type GetMyEmoteListItemFlags struct {
	NoAccess bool `json:"no_access"` // 是否为禁用 true：禁用
}

// GetEmotePackageDetailInfoResult 是获取指定表情包详细信息接口的 data 对象。
//
// 文档标注 data 在参数无效时为 null，此时 Execute 返回 (nil, nil)。
type GetEmotePackageDetailInfoResult struct {
	Packages []GetEmotePackageDetailInfoPackage `json:"packages"` // 表情包列表
}

// GetEmotePackageDetailInfoPackage 是 data.packages 数组中的表情包。
type GetEmotePackageDetailInfoPackage struct {
	ID    int                                   `json:"id"`    // 表情包 ID
	Text  string                                `json:"text"`  // 表情包名称
	URL   string                                `json:"url"`   // 表情包标志图片 URL
	MTime int64                                 `json:"mtime"` // 创建时间，时间戳
	Type  int                                   `json:"type"`  // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr  int                                   `json:"attr"`  // 未知作用
	Meta  GetEmotePackageDetailInfoPackageMeta  `json:"meta"`  // 属性信息
	Emote []GetEmotePackageDetailInfoItem       `json:"emote"` // 表情列表
	Flags GetEmotePackageDetailInfoPackageFlags `json:"flags"` // 是否添加标志
}

// GetEmotePackageDetailInfoPackageMeta 是 data.packages[].meta 对象。
type GetEmotePackageDetailInfoPackageMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetEmotePackageDetailInfoPackageFlags 是 data.packages[].flags 对象。
type GetEmotePackageDetailInfoPackageFlags struct {
	Added bool `json:"added"` // 是否已添加,需要登录（SESSDATA）否则恒为false,true：已添加 false：未添加
}

// GetEmotePackageDetailInfoItem 是 data.packages[].emote 数组中的表情。
type GetEmotePackageDetailInfoItem struct {
	ID        int                                `json:"id"`         // 表情 ID
	PackageID int                                `json:"package_id"` // 表情包 ID
	Text      string                             `json:"text"`       // 表情转义符或颜文字
	URL       string                             `json:"url"`        // 表情图片 URL 或颜文字
	MTime     int64                              `json:"mtime"`      // 创建时间，时间戳
	Type      int                                `json:"type"`       // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr      int                                `json:"attr"`       // 未知作用
	Meta      GetEmotePackageDetailInfoItemMeta  `json:"meta"`       // 属性信息
	Flags     GetEmotePackageDetailInfoItemFlags `json:"flags"`      // 禁用标志
}

// GetEmotePackageDetailInfoItemMeta 是 data.packages[].emote[].meta 对象。
type GetEmotePackageDetailInfoItemMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetEmotePackageDetailInfoItemFlags 是 data.packages[].emote[].flags 对象。
type GetEmotePackageDetailInfoItemFlags struct {
	NoAccess bool `json:"no_access"` // 是否为禁用 true：禁用
}

// GetAllEmoteListResult 是获取所有表情列表接口的 data 对象。
type GetAllEmoteListResult struct {
	UserPanelPackages []GetAllEmoteListPackage `json:"user_panel_packages"` // 用户拥有的表情包
	AllPackages       []GetAllEmoteListPackage `json:"all_packages"`        // 所有表情包
	Mall              GetAllEmoteListMall      `json:"mall"`                // 商城页面
}

// GetAllEmoteListPackage 是 user_panel_packages 与 all_packages 数组中的表情包。
type GetAllEmoteListPackage struct {
	ID    int                         `json:"id"`    // 表情包 ID
	Text  string                      `json:"text"`  // 表情包名称
	URL   string                      `json:"url"`   // 表情包标志图片 URL
	MTime int64                       `json:"mtime"` // 创建时间，时间戳
	Type  int                         `json:"type"`  // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr  int                         `json:"attr"`  // 未知作用
	Meta  GetAllEmoteListPackageMeta  `json:"meta"`  // 属性信息
	Emote []GetAllEmoteListItem       `json:"emote"` // 表情列表
	Flags GetAllEmoteListPackageFlags `json:"flags"` // 是否添加标志
}

// GetAllEmoteListPackageMeta 是 packages[].meta 对象。
type GetAllEmoteListPackageMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetAllEmoteListPackageFlags 是 packages[].flags 对象。
type GetAllEmoteListPackageFlags struct {
	Added bool `json:"added"` // 是否已添加,需要登录（SESSDATA）否则恒为false,true：已添加 false：未添加
}

// GetAllEmoteListItem 是 packages[].emote 数组中的表情。
type GetAllEmoteListItem struct {
	ID        int                      `json:"id"`         // 表情 ID
	PackageID int                      `json:"package_id"` // 表情包 ID
	Text      string                   `json:"text"`       // 表情转义符或颜文字
	URL       string                   `json:"url"`        // 表情图片 URL 或颜文字
	MTime     int64                    `json:"mtime"`      // 创建时间，时间戳
	Type      int                      `json:"type"`       // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr      int                      `json:"attr"`       // 未知作用
	Meta      GetAllEmoteListItemMeta  `json:"meta"`       // 属性信息
	Flags     GetAllEmoteListItemFlags `json:"flags"`      // 禁用标志
}

// GetAllEmoteListItemMeta 是 packages[].emote[].meta 对象。
type GetAllEmoteListItemMeta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

// GetAllEmoteListItemFlags 是 packages[].emote[].flags 对象。
type GetAllEmoteListItemFlags struct {
	NoAccess bool `json:"no_access"` // 是否为禁用 true：禁用
}

// GetAllEmoteListMall 是 data.mall 对象。
type GetAllEmoteListMall struct {
	Title string `json:"title"` // 商城名称
	URL   string `json:"url"`   // 商城页面url
}
