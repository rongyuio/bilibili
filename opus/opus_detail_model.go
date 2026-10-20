package opus

import (
	"encoding/json"

	"github.com/rongyuio/bilibili"
)

// 图文详情相关响应模型。
//
// 本包不复用根包或其他子包的响应体：所有对象与数组元素都以 GetOpusDetail
// 为前缀独立定义，便于后续单个接口的字段变更不影响其他调用。
// modules 是带 module_type 的异构模块数组，每个元素只填充与自身类型对应的字段，
// 因此各模块字段使用指针。

// GetOpusDetailResult 是图文详情接口的 data。
type GetOpusDetailResult struct {
	Fallback *GetOpusDetailFallback `json:"fallback,omitempty"` // 旧版专栏的回退信息
	Item     *GetOpusDetailItem     `json:"item,omitempty"`     // 图文内容
}

// GetOpusDetailFallback 是旧版专栏未携带 htmlNewStyle 时的回退信息。
type GetOpusDetailFallback struct {
	ID   json.Number `json:"id"`   // 回退的内容 ID
	Type int         `json:"type"` // 回退的内容类型，已知 2 为专栏
}

// GetOpusDetailItem 是图文内容本体。
type GetOpusDetailItem struct {
	Basic    GetOpusDetailBasic    `json:"basic"`    // 基本信息
	Fallback json.Number           `json:"fallback"` // 回滚信息，请检查请求参数 features
	IDStr    string                `json:"id_str"`   // 动态 id
	Modules  []GetOpusDetailModule `json:"modules"`  // 模块信息，按 module_type 区分
	Type     int                   `json:"type"`     // 类型
}

// GetOpusDetailBasic 是图文的基本信息。
type GetOpusDetailBasic struct {
	CommentIDStr string                `json:"comment_id_str"` // 评论对象 id 字符串
	CommentType  int                   `json:"comment_type"`   // 评论区类型
	LikeIcon     GetOpusDetailLikeIcon `json:"like_icon"`      // 点赞图标
	RidStr       string                `json:"rid_str"`        // 关联 id 字符串
	Title        string                `json:"title"`          // 图文标题
	UID          json.Number           `json:"uid"`            // 作者 mid (UID)
}

// GetOpusDetailLikeIcon 是图文基本信息中的点赞图标。
type GetOpusDetailLikeIcon struct {
	ActionURL string      `json:"action_url"` // 点击后的图标 URL
	EndURL    string      `json:"end_url"`    // 结束状态图标 URL
	ID        json.Number `json:"id"`         // 图标 id
	StartURL  string      `json:"start_url"`  // 初始状态图标 URL
}

// GetOpusDetailModule 是图文的一个功能模块，由 ModuleType 决定哪个字段有值。
type GetOpusDetailModule struct {
	ModuleAuthor  *GetOpusDetailModuleAuthor  `json:"module_author,omitempty"`  // 作者模块
	ModuleBottom  *GetOpusDetailModuleBottom  `json:"module_bottom,omitempty"`  // 底部分享模块
	ModuleContent *GetOpusDetailModuleContent `json:"module_content,omitempty"` // 正文模块
	ModuleExtend  *GetOpusDetailModuleExtend  `json:"module_extend,omitempty"`  // 扩展信息模块
	ModuleStat    *GetOpusDetailModuleStat    `json:"module_stat,omitempty"`    // 计数模块
	ModuleTitle   *GetOpusDetailModuleTitle   `json:"module_title,omitempty"`   // 标题模块
	ModuleType    string                      `json:"module_type"`              // 模块类型
}

// GetOpusDetailModuleTitle 是标题模块。
type GetOpusDetailModuleTitle struct {
	Text string `json:"text"` // 标题文本
}

// GetOpusDetailModuleAuthor 是作者模块。
type GetOpusDetailModuleAuthor struct {
	Avatar          GetOpusDetailAuthorAvatar  `json:"avatar"`                  // 头像
	DecorateCard    *GetOpusDetailDecorateCard `json:"decorate_card,omitempty"` // 装扮卡片
	Face            string                     `json:"face"`                    // 头像 URL
	FaceNft         bool                       `json:"face_nft"`                // 是否为 NFT 头像
	Following       bilibili.NumberOrString    `json:"following"`               // 关注状态，可能为 null、布尔或数字
	JumpURL         string                     `json:"jump_url"`                // 主页跳转 URL
	Label           string                     `json:"label"`                   // 用户标签
	Mid             json.Number                `json:"mid"`                     // 作者 mid (UID)
	More            *GetOpusDetailModuleMore   `json:"more,omitempty"`          // 更多操作
	Name            string                     `json:"name"`                    // 用户名
	Official        GetOpusDetailOfficial      `json:"official"`                // 认证信息
	Pendant         GetOpusDetailPendant       `json:"pendant"`                 // 挂件
	PubLocationText string                     `json:"pub_location_text"`       // 发布地点
	PubTime         string                     `json:"pub_time"`                // 发布时间文案
	PubTs           json.Number                `json:"pub_ts"`                  // 发布时间戳，秒
	Type            string                     `json:"type"`                    // 作者类型
	ViewsText       string                     `json:"views_text"`              // 浏览数文案
	Vip             GetOpusDetailVip           `json:"vip"`                     // 大会员信息
}

// GetOpusDetailModuleMore 是作者模块中的更多操作。
type GetOpusDetailModuleMore struct {
	ThreePointItems []GetOpusDetailThreePointItem `json:"three_point_items"` // 三点菜单项
}

// GetOpusDetailThreePointItem 是三点菜单中的一项。
type GetOpusDetailThreePointItem struct {
	Label string `json:"label"` // 菜单项文案
	Type  string `json:"type"`  // 菜单项类型
}

// GetOpusDetailOfficial 是作者的认证信息。
type GetOpusDetailOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetOpusDetailPendant 是作者的挂件信息。
type GetOpusDetailPendant struct {
	Expire            json.Number `json:"expire"`              // 挂件到期时间
	Image             string      `json:"image"`               // 挂件图片 URL
	ImageEnhance      string      `json:"image_enhance"`       // 增强挂件图片 URL
	ImageEnhanceFrame string      `json:"image_enhance_frame"` // 增强挂件边框 URL
	NPid              json.Number `json:"n_pid"`               // 新挂件 id
	Name              string      `json:"name"`                // 挂件名称
	Pid               json.Number `json:"pid"`                 // 挂件 id
}

// GetOpusDetailVip 是作者的大会员信息。
type GetOpusDetailVip struct {
	AvatarIcon         GetOpusDetailAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
	AvatarSubscript    int                     `json:"avatar_subscript"`     // 是否显示会员图标：0=不显示，1=显示
	AvatarSubscriptURL string                  `json:"avatar_subscript_url"` // 大会员角标地址
	DueDate            json.Number             `json:"due_date"`             // 大会员到期时间，毫秒时间戳
	Label              GetOpusDetailVipLabel   `json:"label"`                // 大会员铭牌
	NicknameColor      string                  `json:"nickname_color"`       // 会员昵称颜色
	Role               int                     `json:"role"`                 // 大角色类型：1=月度，3=年度，7=十年，15=百年
	Status             int                     `json:"status"`               // 大会员状态：0=无，1=有
	ThemeType          int                     `json:"theme_type"`           // 会员样式 id
	TvDueDate          json.Number             `json:"tv_due_date"`          // 电视大会员到期时间，秒级时间戳
	TvVipPayType       int                     `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	TvVipStatus        int                     `json:"tv_vip_status"`        // 电视大会员状态：0=未开通
	Type               int                     `json:"type"`                 // 会员类型：0=无，1=月大会员，2=年度及以上大会员
	VipPayType         int                     `json:"vip_pay_type"`         // 支付类型：0=未支付，1=已支付
}

// GetOpusDetailAvatarIcon 是大会员角标信息。
type GetOpusDetailAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetOpusDetailVipLabel 是大会员铭牌。
type GetOpusDetailVipLabel struct {
	Path                  string                    `json:"path"`                      // 空。作用尚不明确
	Text                  string                    `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                    `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                    `json:"text_color"`                // 会员标签
	BgStyle               int                       `json:"bg_style"`                  // 1
	BgColor               string                    `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                    `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                      `json:"use_img_label"`             // true
	ImgLabelURIHans       string                    `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                    `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                    `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                    `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                       `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                       `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetOpusDetailVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                    `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                    `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetOpusDetailVipLabelGoto 是大会员铭牌中的 label_goto 对象。
type GetOpusDetailVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetOpusDetailAuthorAvatar 是图文作者的头像及渲染层配置。
type GetOpusDetailAuthorAvatar struct {
	ContainerSize  GetOpusDetailAvatarContainerSize `json:"container_size"`  // 容器尺寸
	FallbackLayers GetOpusDetailAvatarLayerGroup    `json:"fallback_layers"` // 降级渲染层
	Layers         []GetOpusDetailAvatarLayerGroup  `json:"layers"`          // 渲染层
	Mid            string                           `json:"mid"`             // 用户 mid，字符串形式
}

// GetOpusDetailAvatarContainerSize 是头像容器的尺寸。
type GetOpusDetailAvatarContainerSize struct {
	Height float64 `json:"height"` // 高度
	Width  float64 `json:"width"`  // 宽度
}

// GetOpusDetailAvatarLayerGroup 是一组头像渲染层。
type GetOpusDetailAvatarLayerGroup struct {
	IsCriticalGroup bool                       `json:"is_critical_group,omitempty"` // 是否为关键组
	Layers          []GetOpusDetailAvatarLayer `json:"layers"`                      // 层列表
}

// GetOpusDetailAvatarLayer 是头像渲染配置中的一层。
type GetOpusDetailAvatarLayer struct {
	GeneralSpec GetOpusDetailAvatarLayerGeneralSpec `json:"general_spec"` // 通用规格
	LayerConfig GetOpusDetailAvatarLayerConfig      `json:"layer_config"` // 层配置
	Resource    GetOpusDetailAvatarLayerResource    `json:"resource"`     // 资源
	Visible     bool                                `json:"visible"`      // 是否可见
}

// GetOpusDetailAvatarLayerGeneralSpec 是头像渲染层的通用规格。
type GetOpusDetailAvatarLayerGeneralSpec struct {
	PosSpec    GetOpusDetailAvatarLayerPosSpec    `json:"pos_spec"`    // 位置规格
	RenderSpec GetOpusDetailAvatarLayerRenderSpec `json:"render_spec"` // 渲染规格
	SizeSpec   GetOpusDetailAvatarLayerSizeSpec   `json:"size_spec"`   // 尺寸规格
}

// GetOpusDetailAvatarLayerPosSpec 是头像渲染层的位置规格。
type GetOpusDetailAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`         // X 轴位置
	AxisY         float64 `json:"axis_y"`         // Y 轴位置
	CoordinatePos int     `json:"coordinate_pos"` // 坐标位置
}

// GetOpusDetailAvatarLayerRenderSpec 是头像渲染层的渲染规格。
type GetOpusDetailAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"` // 不透明度
}

// GetOpusDetailAvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type GetOpusDetailAvatarLayerSizeSpec struct {
	Height float64 `json:"height"` // 高度
	Width  float64 `json:"width"`  // 宽度
}

// GetOpusDetailAvatarLayerConfig 是头像渲染层的配置。
type GetOpusDetailAvatarLayerConfig struct {
	IsCritical bool                         `json:"is_critical,omitempty"` // 是否为关键层
	Tags       GetOpusDetailAvatarLayerTags `json:"tags"`                  // 层标签
}

// GetOpusDetailAvatarLayerTags 是头像渲染层的标签集合，键名即标签名。
type GetOpusDetailAvatarLayerTags struct {
	AvatarLayer  struct{}                      `json:"AVATAR_LAYER"`  // 头像层标记
	GeneralCfg   GetOpusDetailAvatarGeneralCfg `json:"GENERAL_CFG"`   // 通用配置
	IconLayer    struct{}                      `json:"ICON_LAYER"`    // 图标层标记
	PendantLayer struct{}                      `json:"PENDENT_LAYER"` // 挂件层标记
}

// GetOpusDetailAvatarGeneralCfg 是头像渲染层的通用配置。
type GetOpusDetailAvatarGeneralCfg struct {
	ConfigType    int `json:"config_type"` // 配置类型
	GeneralConfig struct {
		WebCssStyle GetOpusDetailAvatarLayerWebCssStyle `json:"web_css_style"` // CSS 样式
	} `json:"general_config"` // 通用配置
}

// GetOpusDetailAvatarLayerWebCssStyle 是头像渲染层的 CSS 样式。
type GetOpusDetailAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`               // 圆角
	BackgroundColor string `json:"background-color,omitempty"` // 背景色
	Border          string `json:"border,omitempty"`           // 边框
	BoxSizing       string `json:"boxSizing,omitempty"`        // 盒模型
}

// GetOpusDetailAvatarLayerResource 是头像渲染层的资源，图片与动画二选一。
type GetOpusDetailAvatarLayerResource struct {
	ResAnimation *GetOpusDetailAvatarAnimation `json:"res_animation,omitempty"` // 动画资源
	ResImage     *GetOpusDetailAvatarImage     `json:"res_image,omitempty"`     // 图片资源
	ResType      int                           `json:"res_type"`                // 资源类型
}

// GetOpusDetailAvatarAnimation 是头像渲染层的动画资源。
type GetOpusDetailAvatarAnimation struct {
	WebpSrc GetOpusDetailAvatarImageSrc `json:"webp_src"` // WebP 资源
}

// GetOpusDetailAvatarImage 是头像渲染层的图片资源。
type GetOpusDetailAvatarImage struct {
	ImageSrc GetOpusDetailAvatarImageSrc `json:"image_src"` // 图片资源
}

// GetOpusDetailAvatarImageSrc 是头像渲染层的图片来源。
type GetOpusDetailAvatarImageSrc struct {
	Local       int                        `json:"local,omitempty"`       // 本地占位资源
	Placeholder int                        `json:"placeholder,omitempty"` // 占位类型
	Remote      *GetOpusDetailAvatarRemote `json:"remote,omitempty"`      // 远程资源
	SrcType     int                        `json:"src_type"`              // 资源来源类型
}

// GetOpusDetailAvatarRemote 是头像渲染层的远程资源。
type GetOpusDetailAvatarRemote struct {
	BfsStyle string `json:"bfs_style"` // BFS 样式
	URL      string `json:"url"`       // 资源 URL
}

// GetOpusDetailDecorateCard 是作者的装扮卡片信息。
type GetOpusDetailDecorateCard struct {
	BigCardURL   string                    `json:"big_card_url"`   // 大卡片 URL
	CardType     int                       `json:"card_type"`      // 卡片类型
	CardTypeName string                    `json:"card_type_name"` // 卡片类型名
	CardURL      string                    `json:"card_url"`       // 卡片 URL
	Fan          *GetOpusDetailDecorateFan `json:"fan,omitempty"`  // 粉丝专属信息
	ID           json.Number               `json:"id"`             // 卡片 id
	ImageEnhance string                    `json:"image_enhance"`  // 增强图片 URL
	ItemID       json.Number               `json:"item_id"`        // 装扮 id
	JumpURL      string                    `json:"jump_url"`       // 跳转 URL
	Name         string                    `json:"name"`           // 卡片名称
}

// GetOpusDetailDecorateFan 是装扮卡片的粉丝专属信息。
type GetOpusDetailDecorateFan struct {
	Color       string                    `json:"color"`                  // 粉丝颜色
	ColorFormat *GetOpusDetailColorFormat `json:"color_format,omitempty"` // 颜色格式
	IsFan       int                       `json:"is_fan"`                 // 是否为粉丝：1=是
	Name        string                    `json:"name"`                   // 粉丝铭牌文案
	NumDesc     string                    `json:"num_desc"`               // 粉丝编号文案
	Number      json.Number               `json:"number"`                 // 粉丝编号
}

// GetOpusDetailColorFormat 是粉丝铭牌的渐变色格式。
type GetOpusDetailColorFormat struct {
	Colors     []string `json:"colors"`      // 颜色列表
	EndPoint   string   `json:"end_point"`   // 渐变终点
	Gradients  []int    `json:"gradients"`   // 渐变位置
	StartPoint string   `json:"start_point"` // 渐变起点
}

// GetOpusDetailModuleContent 是正文模块。
type GetOpusDetailModuleContent struct {
	Paragraphs []GetOpusDetailParagraph `json:"paragraphs"` // 段落列表
}

// GetOpusDetailParagraph 是正文的一个段落，由 ParaType 决定哪个字段有值。
type GetOpusDetailParagraph struct {
	Align    int                         `json:"align"`          // 对齐方式
	Code     *GetOpusDetailParagraphCode `json:"code,omitempty"` // 代码块
	List     *GetOpusDetailParagraphList `json:"list,omitempty"` // 列表
	ParaType int                         `json:"para_type"`      // 段落类型
	Pic      *GetOpusDetailParagraphPic  `json:"pic,omitempty"`  // 图片
	Text     *GetOpusDetailParagraphText `json:"text,omitempty"` // 文本
}

// GetOpusDetailParagraphText 是段落中的文本节点集合。
type GetOpusDetailParagraphText struct {
	Nodes []GetOpusDetailTextNode `json:"nodes"` // 文本节点
}

// GetOpusDetailTextNode 是正文中的一个文本节点，普通文本与富文本二选一。
type GetOpusDetailTextNode struct {
	Rich *GetOpusDetailTextRich `json:"rich,omitempty"` // 富文本节点
	Type string                 `json:"type"`           // 节点类型
	Word *GetOpusDetailTextWord `json:"word,omitempty"` // 普通文本节点
}

// GetOpusDetailTextWord 是普通文本节点。
type GetOpusDetailTextWord struct {
	Color    string                 `json:"color,omitempty"` // 文字颜色
	FontSize int                    `json:"font_size"`       // 字号
	Style    GetOpusDetailTextStyle `json:"style"`           // 样式
	Words    string                 `json:"words"`           // 文字内容
}

// GetOpusDetailTextStyle 是文本节点的样式。
type GetOpusDetailTextStyle struct {
	Bold bool `json:"bold,omitempty"` // 是否加粗
}

// GetOpusDetailTextRich 是富文本节点，例如话题、提及等。
type GetOpusDetailTextRich struct {
	JumpURL  string `json:"jump_url"`  // 跳转 URL
	OrigText string `json:"orig_text"` // 原始文本
	Text     string `json:"text"`      // 显示文本
	Type     string `json:"type"`      // 富文本类型
}

// GetOpusDetailParagraphPic 是段落中的图片。
type GetOpusDetailParagraphPic struct {
	Pics  []GetOpusDetailParagraphPicItem `json:"pics"`  // 图片列表
	Style int                             `json:"style"` // 图片样式
}

// GetOpusDetailParagraphPicItem 是段落中的一张图片，尺寸字段可能为 null。
type GetOpusDetailParagraphPicItem struct {
	Height  *int     `json:"height"`   // 图片高度，可能为 null
	LiveURL *string  `json:"live_url"` // 动态图片 URL，可能为 null
	Size    *float64 `json:"size"`     // 图片大小，可能为 null
	URL     string   `json:"url"`      // 图片 URL
	Width   *int     `json:"width"`    // 图片宽度，可能为 null
}

// GetOpusDetailParagraphList 是段落中的列表。
type GetOpusDetailParagraphList struct {
	Items []GetOpusDetailListItem `json:"items"` // 列表项
	Style int                     `json:"style"` // 列表样式
}

// GetOpusDetailListItem 是列表中的一项。
type GetOpusDetailListItem struct {
	Level int                     `json:"level"` // 层级
	Nodes []GetOpusDetailTextNode `json:"nodes"` // 文本节点
	Order int                     `json:"order"` // 序号
}

// GetOpusDetailParagraphCode 是段落中的代码块。
type GetOpusDetailParagraphCode struct {
	Content string `json:"content"` // 代码内容
	Lang    string `json:"lang"`    // 语言标识，例如 language-bash
}

// GetOpusDetailModuleExtend 是扩展信息模块。
type GetOpusDetailModuleExtend struct {
	Items []GetOpusDetailExtendItem `json:"items"` // 扩展项
}

// GetOpusDetailExtendItem 是一个扩展项，通常为话题。
type GetOpusDetailExtendItem struct {
	Icon    any    `json:"icon"`     // 图标，可能为 null
	JumpURL string `json:"jump_url"` // 跳转 URL
	Text    string `json:"text"`     // 文本
}

// GetOpusDetailModuleBottom 是底部分享模块。
type GetOpusDetailModuleBottom struct {
	ShareInfo GetOpusDetailShareInfo `json:"share_info"` // 分享信息
}

// GetOpusDetailShareInfo 是分享信息。
type GetOpusDetailShareInfo struct {
	Pic     string `json:"pic"`     // 分享图 URL
	Summary string `json:"summary"` // 分享摘要
	Title   string `json:"title"`   // 分享标题
}

// GetOpusDetailModuleStat 是计数模块。
type GetOpusDetailModuleStat struct {
	Coin     GetOpusDetailLikeStat `json:"coin"`     // 投币数
	Comment  GetOpusDetailStat     `json:"comment"`  // 评论数
	Favorite GetOpusDetailLikeStat `json:"favorite"` // 收藏数
	Forward  GetOpusDetailStat     `json:"forward"`  // 转发数
	Like     GetOpusDetailLikeStat `json:"like"`     // 点赞数
}

// GetOpusDetailStat 是评论或转发的计数信息。
type GetOpusDetailStat struct {
	Count     json.Number `json:"count"`     // 计数
	Forbidden bool        `json:"forbidden"` // 是否被禁止显示
}

// GetOpusDetailLikeStat 是点赞计数信息，额外包含当前用户的点赞状态。
type GetOpusDetailLikeStat struct {
	Count     json.Number `json:"count"`     // 计数
	Forbidden bool        `json:"forbidden"` // 是否被禁止显示
	Status    bool        `json:"status"`    // 当前用户是否已点赞
}
