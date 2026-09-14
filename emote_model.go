package bilibili

// 表情相关响应模型。

type EmotePackageFlags struct {
	Added bool `json:"added"` // 是否已添加,需要登录（SESSDATA）否则恒为false,true：已添加 false：未添加
}

type Meta struct {
	Size    int    `json:"size"`               // 表情尺寸信息 1:小 2:大
	ItemID  int    `json:"item_id,omitempty"`  // 购买物品 ID，可能为空
	ItemURL string `json:"item_url,omitempty"` // 购买物品页面 URL，可能为空
	Alias   string `json:"alias"`              // 简写名
}

type EmoteFlags struct {
	NoAccess bool `json:"no_access"` // 是否为禁用 true：禁用
}

type Emote struct {
	ID        int        `json:"id"`         // 表情 ID
	PackageID int        `json:"package_id"` // 表情包 ID
	Text      string     `json:"text"`       // 表情转义符或颜文字
	URL       string     `json:"url"`        // 表情图片 URL 或颜文字
	MTime     int64      `json:"mtime"`      // 创建时间，时间戳
	Type      int        `json:"type"`       // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr      int        `json:"attr"`       // 未知作用
	Meta      Meta       `json:"meta"`       // 属性信息
	Flags     EmoteFlags `json:"flags"`      // 禁用标志
}

type EmotePackage struct {
	ID    int               `json:"id"`    // 表情包 ID
	Text  string            `json:"text"`  // 表情包名称
	URL   string            `json:"url"`   // 表情包标志图片 URL
	MTime int64             `json:"mtime"` // 创建时间，时间戳
	Type  int               `json:"type"`  // 表情类型 1：普通	2：会员专属	3：购买所得	4：颜文字
	Attr  int               `json:"attr"`  // 未知作用
	Meta  Meta              `json:"meta"`  // 属性信息
	Emote []Emote           `json:"emote"` // 表情列表
	Flags EmotePackageFlags `json:"flags"` // 是否添加标志
}

type EmoteList struct {
	EmotePackages []EmotePackage `json:"packages"` // 表情包包列表
}

type Mall struct {
	Title string `json:"title"` // 商城名称
	Url   string `json:"url"`   // 商城页面url
}

type AllEmoteList struct {
	User_panel_packages []EmotePackage `json:"user_panel_packages"`
	All_packages        []EmotePackage `json:"all_packages"`
	Mall                Mall           `json:"mall"`
}
