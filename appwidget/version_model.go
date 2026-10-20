package appwidget

// APP 版本相关响应模型。各接口的返回体独立定义，互不复用。

// GetVersionInfoItem 是获取版本信息接口 data 数组中的一项。
type GetVersionInfoItem struct {
	Plat    int    `json:"plat"`    // 平台。0：Android，1：iPhone，2：iPad HD，3：UWP，12：TV 版，35：车机版，40：PC 客户端
	Desc    string `json:"desc"`    // 版本描述
	Version string `json:"version"` // 版本号
	Build   int    `json:"build"`   // 内部版本号
	Ptime   int64  `json:"ptime"`   // 发布时间。文档表格未列出，响应示例中存在
}

// GetVersionUpgradeResult 是版本更新接口的 data 对象。
type GetVersionUpgradeResult struct {
	Title       string `json:"title"`        // 标题
	Content     string `json:"content"`      // 内容
	Version     string `json:"version"`      // 最新版本名称
	VersionCode int    `json:"version_code"` // 最新版本号，即 build
	URL         string `json:"url"`          // 下载地址
	Size        int64  `json:"size"`         // 大小，单位字节
	Md5         string `json:"md5"`          // MD5 值
	Silent      int    `json:"silent"`       // 是否静默安装
	UpgradeType int    `json:"upgrade_type"` // 升级类型
	Cycle       int    `json:"cycle"`        // 升级周期
	Policy      int    `json:"policy"`       // 升级策略
	PolicyURL   string `json:"policy_url"`   // 升级策略链接
	IsGray      int    `json:"is_gray"`      // 是否灰度。文档表格未列出，响应示例中存在
	Ptime       int64  `json:"ptime"`        // 发布时间
}

// GetAppVersionResult 是拼接最新版本下载地址接口的 data 对象。
type GetAppVersionResult struct {
	CdnURL     string `json:"cdn_url"`      // CDN URL，内容同请求参数 cdn_url
	SoleCdnURL string `json:"sole_cdn_url"` // 带版本的 CDN URL，即 cdn_url + "?v=" + 最新版本
}
