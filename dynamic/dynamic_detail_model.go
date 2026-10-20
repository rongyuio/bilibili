package dynamic

import (
	"encoding/json"

	"github.com/rongyuio/bilibili"
)

// 动态详情相关响应模型。
//
// 本包不复用根包或其他子包的响应体：每个接口的 data 层以及所有嵌套对象、
// 数组元素都以接口方法名为前缀独立定义，即使字段完全相同也各自定义，
// 便于后续单个接口的字段变更不影响其他调用。

// GetDynamicDetailInfoResult 是动态详情接口的 data。
type GetDynamicDetailInfoResult struct {
	Item GetDynamicDetailInfoItem `json:"item"` // 动态详情
}

// GetDynamicDetailInfoItem is a space feed item, including its original item when reposted.
type GetDynamicDetailInfoItem struct {
	Basic   GetDynamicDetailInfoBasic        `json:"basic"`
	IDStr   json.Number                      `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules GetDynamicDetailInfoModules      `json:"modules"`
	Orig    GetDynamicDetailInfoOriginalItem `json:"orig,omitempty"`
	Type    string                           `json:"type"`
	Visible bool                             `json:"visible"`
}

// GetDynamicDetailInfoBasic contains comment identifiers and actions for a feed item.
type GetDynamicDetailInfoBasic struct {
	CommentIDStr string                       `json:"comment_id_str"`
	CommentType  int                          `json:"comment_type"`
	LikeIcon     GetDynamicDetailInfoLikeIcon `json:"like_icon"`
	RidStr       string                       `json:"rid_str"`
}

// GetDynamicDetailInfoLikeIcon is the like icon metadata of a space feed item.
type GetDynamicDetailInfoLikeIcon struct {
	ActionURL string      `json:"action_url"`
	EndURL    string      `json:"end_url"`
	ID        json.Number `json:"id"`
	StartURL  string      `json:"start_url"`
}

// GetDynamicDetailInfoModules groups the author, content, actions and counts of a feed item.
type GetDynamicDetailInfoModules struct {
	ModuleAuthor  GetDynamicDetailInfoModuleAuthor  `json:"module_author"`
	ModuleDynamic GetDynamicDetailInfoModuleDynamic `json:"module_dynamic"`
	ModuleMore    GetDynamicDetailInfoModuleMore    `json:"module_more"`
	ModuleStat    GetDynamicDetailInfoModuleStat    `json:"module_stat"`
}

// GetDynamicDetailInfoModuleAuthor describes the author of the outer feed item.
type GetDynamicDetailInfoModuleAuthor struct {
	Avatar          GetDynamicDetailInfoAuthorAvatar   `json:"avatar"`
	Face            string                             `json:"face"`
	FaceNft         bool                               `json:"face_nft"`
	Following       bilibili.NumberOrString            `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL         string                             `json:"jump_url"`
	Label           string                             `json:"label"`
	Mid             json.Number                        `json:"mid"`
	Name            string                             `json:"name"`
	OfficialVerify  GetDynamicDetailInfoOfficialVerify `json:"official_verify"`
	Pendant         GetDynamicDetailInfoPendant        `json:"pendant"`
	PubAction       string                             `json:"pub_action"`
	PubLocationText string                             `json:"pub_location_text"`
	PubTime         string                             `json:"pub_time"`
	PubTs           json.Number                        `json:"pub_ts"`
	Type            string                             `json:"type"`
	Vip             GetDynamicDetailInfoAuthorVip      `json:"vip"`
}

// GetDynamicDetailInfoPendant 外层动态作者的挂件信息，数值字段为 json.Number。
type GetDynamicDetailInfoPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetDynamicDetailInfoAuthorVip 外层动态与原始动态共用的作者会员信息。
type GetDynamicDetailInfoAuthorVip struct {
	AvatarSubscript    int                          `json:"avatar_subscript"`
	AvatarSubscriptURL string                       `json:"avatar_subscript_url"`
	DueDate            json.Number                  `json:"due_date"`
	Label              GetDynamicDetailInfoVipLabel `json:"label"`
	NicknameColor      string                       `json:"nickname_color"`
	Status             int                          `json:"status"`
	ThemeType          int                          `json:"theme_type"`
	Type               int                          `json:"type"`
}

// GetDynamicDetailInfoAuthorAvatar contains the outer author avatar rendering layers.
type GetDynamicDetailInfoAuthorAvatar struct {
	ContainerSize  GetDynamicDetailInfoAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetDynamicDetailInfoAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetDynamicDetailInfoAvatarLayerWebCssStyle `json:"web_css_style"`
						} `json:"general_config"`
					} `json:"GENERAL_CFG"`
					IconLayer struct{} `json:"ICON_LAYER,omitempty"`
				} `json:"tags"`
			} `json:"layer_config"`
			Resource struct {
				ResAnimation struct {
					WebpSrc struct {
						Placeholder int `json:"placeholder"`
						Remote      struct {
							BfsStyle string `json:"bfs_style"`
							URL      string `json:"url"`
						} `json:"remote"`
						SrcType int `json:"src_type"`
					} `json:"webp_src"`
				} `json:"res_animation,omitempty"`
				ResType  int `json:"res_type"`
				ResImage struct {
					ImageSrc struct {
						Local   int `json:"local"`
						SrcType int `json:"src_type"`
					} `json:"image_src"`
				} `json:"res_image,omitempty"`
			} `json:"resource"`
			Visible bool `json:"visible"`
		} `json:"layers"`
	} `json:"fallback_layers"`
	Mid string `json:"mid"`
}

// GetDynamicDetailInfoAvatarContainerSize 头像容器的尺寸，外层动态、原始动态与话题动态共用。
type GetDynamicDetailInfoAvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetDynamicDetailInfoModuleDynamic contains the outer item content and optional description and major body.
type GetDynamicDetailInfoModuleDynamic struct {
	Additional any                              `json:"additional"`
	Desc       *GetDynamicDetailInfoDescription `json:"desc"`
	Major      *GetDynamicDetailInfoMajor       `json:"major"`
	Topic      any                              `json:"topic"`
}

// GetDynamicDetailInfoDescription contains text and rich text nodes for the outer item.
type GetDynamicDetailInfoDescription struct {
	RichTextNodes []GetDynamicDetailInfoRichTextNode `json:"rich_text_nodes"`
	Text          string                             `json:"text"`
}

// GetDynamicDetailInfoRichTextNode is one rich text node of an outer item description.
type GetDynamicDetailInfoRichTextNode struct {
	OrigText string                    `json:"orig_text"`
	Text     string                    `json:"text"`
	Type     string                    `json:"type"`
	JumpURL  string                    `json:"jump_url,omitempty"`
	Style    any                       `json:"style"`
	Emoji    GetDynamicDetailInfoEmoji `json:"emoji,omitempty"`
	Rid      string                    `json:"rid,omitempty"`
}

// GetDynamicDetailInfoEmoji is an inline emoji of a rich text node.
type GetDynamicDetailInfoEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetDynamicDetailInfoMajor contains the outer item media body.
type GetDynamicDetailInfoMajor struct {
	Draw    GetDynamicDetailInfoDraw    `json:"draw,omitempty"`
	Type    string                      `json:"type"`
	Archive GetDynamicDetailInfoArchive `json:"archive,omitempty"`
}

// GetDynamicDetailInfoModuleMore contains the additional actions offered for a feed item.
type GetDynamicDetailInfoModuleMore struct {
	ThreePointItems []GetDynamicDetailInfoThreePointItem `json:"three_point_items"`
}

// GetDynamicDetailInfoThreePointItem is one entry of the additional actions menu.
type GetDynamicDetailInfoThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// GetDynamicDetailInfoModuleStat contains comment, forward and like counts for a feed item.
type GetDynamicDetailInfoModuleStat struct {
	Comment GetDynamicDetailInfoStat     `json:"comment"`
	Forward GetDynamicDetailInfoStat     `json:"forward"`
	Like    GetDynamicDetailInfoLikeStat `json:"like"`
}

// GetDynamicDetailInfoStat 评论或转发的计数信息。
type GetDynamicDetailInfoStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
}

// GetDynamicDetailInfoLikeStat 点赞计数信息，额外包含当前用户的点赞状态。
type GetDynamicDetailInfoLikeStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
	Status    bool        `json:"status"`
}

// GetDynamicDetailInfoArchive describes a video archive in either an outer or original media body.
type GetDynamicDetailInfoArchive struct {
	Aid            string                           `json:"aid"`
	Badge          GetDynamicDetailInfoArchiveBadge `json:"badge"`
	Bvid           string                           `json:"bvid"`
	Cover          string                           `json:"cover"`
	Desc           string                           `json:"desc"`
	DisablePreview int                              `json:"disable_preview"`
	DurationText   string                           `json:"duration_text"`
	JumpURL        string                           `json:"jump_url"`
	Stat           GetDynamicDetailInfoArchiveStat  `json:"stat"`
	Title          string                           `json:"title"`
	Type           int                              `json:"type"`
}

// GetDynamicDetailInfoArchiveBadge 稿件角标信息。
type GetDynamicDetailInfoArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// GetDynamicDetailInfoArchiveStat 稿件的播放与弹幕数。
type GetDynamicDetailInfoArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// GetDynamicDetailInfoDraw describes images in either an outer or original media body.
type GetDynamicDetailInfoDraw struct {
	ID    json.Number                    `json:"id"`
	Items []GetDynamicDetailInfoDrawItem `json:"items"`
}

// GetDynamicDetailInfoDrawItem 图集中的一张图片。
type GetDynamicDetailInfoDrawItem struct {
	Height json.Number `json:"height"`
	Size   json.Number `json:"size"`
	Src    string      `json:"src"`
	Tags   []any       `json:"tags"`
	Width  json.Number `json:"width"`
}

// GetDynamicDetailInfoOriginalItem is the original item embedded in a repost; it is not recursive.
type GetDynamicDetailInfoOriginalItem struct {
	Basic   GetDynamicDetailInfoOriginalBasic   `json:"basic"`
	IDStr   json.Number                         `json:"id_str"`
	Modules GetDynamicDetailInfoOriginalModules `json:"modules"`
	Type    string                              `json:"type"`
	Visible bool                                `json:"visible"`
}

// GetDynamicDetailInfoOriginalBasic contains the original item identifiers, whose types differ from the outer item.
type GetDynamicDetailInfoOriginalBasic struct {
	CommentIDStr string                               `json:"comment_id_str"`
	CommentType  int                                  `json:"comment_type"`
	LikeIcon     GetDynamicDetailInfoOriginalLikeIcon `json:"like_icon"`
	RidStr       string                               `json:"rid_str"`
}

// GetDynamicDetailInfoOriginalLikeIcon 原始动态的点赞图标；id 为整型，与外层不同。
type GetDynamicDetailInfoOriginalLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int    `json:"id"`
	StartURL  string `json:"start_url"`
}

// GetDynamicDetailInfoOriginalModules groups the author and content of an original item.
type GetDynamicDetailInfoOriginalModules struct {
	ModuleAuthor  GetDynamicDetailInfoOriginalModuleAuthor  `json:"module_author"`
	ModuleDynamic GetDynamicDetailInfoOriginalModuleDynamic `json:"module_dynamic"`
}

// GetDynamicDetailInfoOriginalModuleAuthor describes the original author using the original response schema.
type GetDynamicDetailInfoOriginalModuleAuthor struct {
	Avatar         GetDynamicDetailInfoOriginalAuthorAvatar `json:"avatar"`
	Decorate       GetDynamicDetailInfoDecorate             `json:"decorate,omitempty"`
	Face           string                                   `json:"face"`
	FaceNft        bool                                     `json:"face_nft"`
	Following      bilibili.NumberOrString                  `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL        string                                   `json:"jump_url"`
	Label          string                                   `json:"label"`
	Mid            json.Number                              `json:"mid"`
	Name           string                                   `json:"name"`
	OfficialVerify GetDynamicDetailInfoOfficialVerify       `json:"official_verify"`
	Pendant        GetDynamicDetailInfoOriginalPendant      `json:"pendant"`
	PubAction      string                                   `json:"pub_action"`
	PubTime        string                                   `json:"pub_time"`
	PubTs          json.Number                              `json:"pub_ts"`
	Type           string                                   `json:"type"`
	Vip            GetDynamicDetailInfoAuthorVip            `json:"vip"`
}

// GetDynamicDetailInfoDecorate 原始动态作者的装扮卡片信息。
type GetDynamicDetailInfoDecorate struct {
	CardURL string                          `json:"card_url"`
	Fan     GetDynamicDetailInfoDecorateFan `json:"fan"`
	ID      json.Number                     `json:"id"`
	JumpURL string                          `json:"jump_url"`
	Name    string                          `json:"name"`
	Type    int                             `json:"type"`
}

// GetDynamicDetailInfoDecorateFan 装扮卡片的粉丝专属信息。
type GetDynamicDetailInfoDecorateFan struct {
	Color  string      `json:"color"`
	IsFan  bool        `json:"is_fan"`
	NumStr string      `json:"num_str"`
	Number json.Number `json:"number"`
}

// GetDynamicDetailInfoOriginalPendant 原始动态作者的挂件信息；pid 为整型，与外层不同。
type GetDynamicDetailInfoOriginalPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               int         `json:"pid"`
}

// GetDynamicDetailInfoOriginalAuthorAvatar contains avatar rendering layers from the original item schema.
type GetDynamicDetailInfoOriginalAuthorAvatar struct {
	ContainerSize  GetDynamicDetailInfoAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetDynamicDetailInfoAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetDynamicDetailInfoAvatarLayerWebCssStyle `json:"web_css_style"`
						} `json:"general_config"`
					} `json:"GENERAL_CFG,omitempty"`
					PendantLayer struct {
					} `json:"PENDENT_LAYER,omitempty"`
					IconLayer struct {
					} `json:"ICON_LAYER,omitempty"`
				} `json:"tags"`
			} `json:"layer_config"`
			Resource struct {
				ResImage struct {
					ImageSrc struct {
						Placeholder int `json:"placeholder,omitempty"`
						Remote      struct {
							BfsStyle string `json:"bfs_style"`
							URL      string `json:"url"`
						} `json:"remote,omitempty"`
						SrcType int `json:"src_type"`
						Local   int `json:"local,omitempty"`
					} `json:"image_src"`
				} `json:"res_image"`
				ResType int `json:"res_type"`
			} `json:"resource"`
			Visible bool `json:"visible"`
		} `json:"layers"`
	} `json:"fallback_layers"`
	Mid string `json:"mid"`
}

// GetDynamicDetailInfoOriginalModuleDynamic contains the original content; Major retains its value semantics.
type GetDynamicDetailInfoOriginalModuleDynamic struct {
	Additional any                                      `json:"additional"`
	Desc       *GetDynamicDetailInfoOriginalDescription `json:"desc"`
	Major      GetDynamicDetailInfoOriginalMajor        `json:"major"`
	Topic      any                                      `json:"topic"`
}

// GetDynamicDetailInfoOriginalDescription contains original rich text nodes, which differ from outer item nodes.
type GetDynamicDetailInfoOriginalDescription struct {
	RichTextNodes []GetDynamicDetailInfoOriginalRichTextNode `json:"rich_text_nodes"`
	Text          string                                     `json:"text"`
}

// GetDynamicDetailInfoOriginalRichTextNode is one rich text node of an original item description.
type GetDynamicDetailInfoOriginalRichTextNode struct {
	JumpURL  string                    `json:"jump_url,omitempty"`
	OrigText string                    `json:"orig_text"`
	Text     string                    `json:"text"`
	Type     string                    `json:"type"`
	Emoji    GetDynamicDetailInfoEmoji `json:"emoji,omitempty"`
}

// GetDynamicDetailInfoAvatarLayerPosSpec 头像渲染层的位置规格。
type GetDynamicDetailInfoAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetDynamicDetailInfoAvatarLayerRenderSpec 头像渲染层的渲染规格。
type GetDynamicDetailInfoAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetDynamicDetailInfoAvatarLayerSizeSpec 头像渲染层的尺寸规格。
type GetDynamicDetailInfoAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetDynamicDetailInfoAvatarLayerGeneralSpec 头像渲染层的通用规格，三个头像模型共用。
type GetDynamicDetailInfoAvatarLayerGeneralSpec struct {
	PosSpec    GetDynamicDetailInfoAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetDynamicDetailInfoAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetDynamicDetailInfoAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetDynamicDetailInfoAvatarLayerWebCssStyle 头像渲染层的 CSS 样式（整体动态与原动态共用；话题仅返回 borderRadius）。
type GetDynamicDetailInfoAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"background-color,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}

// GetDynamicDetailInfoOfficialVerify 认证标识（仅 type/desc，用于动态、话题、视频等处的 official_verify 字段）。
type GetDynamicDetailInfoOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

type GetDynamicDetailInfoVipLabel struct {
	Path                  string                           `json:"path"`                      // 空。作用尚不明确
	Text                  string                           `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                           `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                           `json:"text_color"`                // 会员标签
	BgStyle               int                              `json:"bg_style"`                  // 1
	BgColor               string                           `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                           `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                             `json:"use_img_label"`             // true
	ImgLabelURIHans       string                           `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                           `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                           `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                           `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                              `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                              `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetDynamicDetailInfoVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                           `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                           `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetDynamicDetailInfoVipLabelGoto `label` 中的 `label_goto` 对象
type GetDynamicDetailInfoVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetDynamicDetailInfoOriginalMajor 是原动态的媒体主体；字段与整体动态的 Major 一致，仍独立定义。
type GetDynamicDetailInfoOriginalMajor struct {
	Draw    GetDynamicDetailInfoDraw    `json:"draw,omitempty"`
	Type    string                      `json:"type"`
	Archive GetDynamicDetailInfoArchive `json:"archive,omitempty"`
}

// GetDynamicDetailDesktopResult 是动态详情（桌面端）接口的 data。
type GetDynamicDetailDesktopResult struct {
	Item GetDynamicDetailDesktopItem `json:"item"` // 动态详情
}

// GetDynamicDetailDesktopItem is a space feed item, including its original item when reposted.
type GetDynamicDetailDesktopItem struct {
	Basic   GetDynamicDetailDesktopBasic        `json:"basic"`
	IDStr   json.Number                         `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules GetDynamicDetailDesktopModules      `json:"modules"`
	Orig    GetDynamicDetailDesktopOriginalItem `json:"orig,omitempty"`
	Type    string                              `json:"type"`
	Visible bool                                `json:"visible"`
}

// GetDynamicDetailDesktopBasic contains comment identifiers and actions for a feed item.
type GetDynamicDetailDesktopBasic struct {
	CommentIDStr string                          `json:"comment_id_str"`
	CommentType  int                             `json:"comment_type"`
	LikeIcon     GetDynamicDetailDesktopLikeIcon `json:"like_icon"`
	RidStr       string                          `json:"rid_str"`
}

// GetDynamicDetailDesktopLikeIcon is the like icon metadata of a space feed item.
type GetDynamicDetailDesktopLikeIcon struct {
	ActionURL string      `json:"action_url"`
	EndURL    string      `json:"end_url"`
	ID        json.Number `json:"id"`
	StartURL  string      `json:"start_url"`
}

// GetDynamicDetailDesktopModules groups the author, content, actions and counts of a feed item.
type GetDynamicDetailDesktopModules struct {
	ModuleAuthor  GetDynamicDetailDesktopModuleAuthor  `json:"module_author"`
	ModuleDynamic GetDynamicDetailDesktopModuleDynamic `json:"module_dynamic"`
	ModuleMore    GetDynamicDetailDesktopModuleMore    `json:"module_more"`
	ModuleStat    GetDynamicDetailDesktopModuleStat    `json:"module_stat"`
}

// GetDynamicDetailDesktopModuleAuthor describes the author of the outer feed item.
type GetDynamicDetailDesktopModuleAuthor struct {
	Avatar          GetDynamicDetailDesktopAuthorAvatar   `json:"avatar"`
	Face            string                                `json:"face"`
	FaceNft         bool                                  `json:"face_nft"`
	Following       bilibili.NumberOrString               `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL         string                                `json:"jump_url"`
	Label           string                                `json:"label"`
	Mid             json.Number                           `json:"mid"`
	Name            string                                `json:"name"`
	OfficialVerify  GetDynamicDetailDesktopOfficialVerify `json:"official_verify"`
	Pendant         GetDynamicDetailDesktopPendant        `json:"pendant"`
	PubAction       string                                `json:"pub_action"`
	PubLocationText string                                `json:"pub_location_text"`
	PubTime         string                                `json:"pub_time"`
	PubTs           json.Number                           `json:"pub_ts"`
	Type            string                                `json:"type"`
	Vip             GetDynamicDetailDesktopAuthorVip      `json:"vip"`
}

// GetDynamicDetailDesktopPendant 外层动态作者的挂件信息，数值字段为 json.Number。
type GetDynamicDetailDesktopPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetDynamicDetailDesktopAuthorVip 外层动态与原始动态共用的作者会员信息。
type GetDynamicDetailDesktopAuthorVip struct {
	AvatarSubscript    int                             `json:"avatar_subscript"`
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"`
	DueDate            json.Number                     `json:"due_date"`
	Label              GetDynamicDetailDesktopVipLabel `json:"label"`
	NicknameColor      string                          `json:"nickname_color"`
	Status             int                             `json:"status"`
	ThemeType          int                             `json:"theme_type"`
	Type               int                             `json:"type"`
}

// GetDynamicDetailDesktopAuthorAvatar contains the outer author avatar rendering layers.
type GetDynamicDetailDesktopAuthorAvatar struct {
	ContainerSize  GetDynamicDetailDesktopAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetDynamicDetailDesktopAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetDynamicDetailDesktopAvatarLayerWebCssStyle `json:"web_css_style"`
						} `json:"general_config"`
					} `json:"GENERAL_CFG"`
					IconLayer struct{} `json:"ICON_LAYER,omitempty"`
				} `json:"tags"`
			} `json:"layer_config"`
			Resource struct {
				ResAnimation struct {
					WebpSrc struct {
						Placeholder int `json:"placeholder"`
						Remote      struct {
							BfsStyle string `json:"bfs_style"`
							URL      string `json:"url"`
						} `json:"remote"`
						SrcType int `json:"src_type"`
					} `json:"webp_src"`
				} `json:"res_animation,omitempty"`
				ResType  int `json:"res_type"`
				ResImage struct {
					ImageSrc struct {
						Local   int `json:"local"`
						SrcType int `json:"src_type"`
					} `json:"image_src"`
				} `json:"res_image,omitempty"`
			} `json:"resource"`
			Visible bool `json:"visible"`
		} `json:"layers"`
	} `json:"fallback_layers"`
	Mid string `json:"mid"`
}

// GetDynamicDetailDesktopAvatarContainerSize 头像容器的尺寸，外层动态、原始动态与话题动态共用。
type GetDynamicDetailDesktopAvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetDynamicDetailDesktopModuleDynamic contains the outer item content and optional description and major body.
type GetDynamicDetailDesktopModuleDynamic struct {
	Additional any                                 `json:"additional"`
	Desc       *GetDynamicDetailDesktopDescription `json:"desc"`
	Major      *GetDynamicDetailDesktopMajor       `json:"major"`
	Topic      any                                 `json:"topic"`
}

// GetDynamicDetailDesktopDescription contains text and rich text nodes for the outer item.
type GetDynamicDetailDesktopDescription struct {
	RichTextNodes []GetDynamicDetailDesktopRichTextNode `json:"rich_text_nodes"`
	Text          string                                `json:"text"`
}

// GetDynamicDetailDesktopRichTextNode is one rich text node of an outer item description.
type GetDynamicDetailDesktopRichTextNode struct {
	OrigText string                       `json:"orig_text"`
	Text     string                       `json:"text"`
	Type     string                       `json:"type"`
	JumpURL  string                       `json:"jump_url,omitempty"`
	Style    any                          `json:"style"`
	Emoji    GetDynamicDetailDesktopEmoji `json:"emoji,omitempty"`
	Rid      string                       `json:"rid,omitempty"`
}

// GetDynamicDetailDesktopEmoji is an inline emoji of a rich text node.
type GetDynamicDetailDesktopEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetDynamicDetailDesktopMajor contains the outer item media body.
type GetDynamicDetailDesktopMajor struct {
	Draw    GetDynamicDetailDesktopDraw    `json:"draw,omitempty"`
	Type    string                         `json:"type"`
	Archive GetDynamicDetailDesktopArchive `json:"archive,omitempty"`
}

// GetDynamicDetailDesktopModuleMore contains the additional actions offered for a feed item.
type GetDynamicDetailDesktopModuleMore struct {
	ThreePointItems []GetDynamicDetailDesktopThreePointItem `json:"three_point_items"`
}

// GetDynamicDetailDesktopThreePointItem is one entry of the additional actions menu.
type GetDynamicDetailDesktopThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// GetDynamicDetailDesktopModuleStat contains comment, forward and like counts for a feed item.
type GetDynamicDetailDesktopModuleStat struct {
	Comment GetDynamicDetailDesktopStat     `json:"comment"`
	Forward GetDynamicDetailDesktopStat     `json:"forward"`
	Like    GetDynamicDetailDesktopLikeStat `json:"like"`
}

// GetDynamicDetailDesktopStat 评论或转发的计数信息。
type GetDynamicDetailDesktopStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
}

// GetDynamicDetailDesktopLikeStat 点赞计数信息，额外包含当前用户的点赞状态。
type GetDynamicDetailDesktopLikeStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
	Status    bool        `json:"status"`
}

// GetDynamicDetailDesktopArchive describes a video archive in either an outer or original media body.
type GetDynamicDetailDesktopArchive struct {
	Aid            string                              `json:"aid"`
	Badge          GetDynamicDetailDesktopArchiveBadge `json:"badge"`
	Bvid           string                              `json:"bvid"`
	Cover          string                              `json:"cover"`
	Desc           string                              `json:"desc"`
	DisablePreview int                                 `json:"disable_preview"`
	DurationText   string                              `json:"duration_text"`
	JumpURL        string                              `json:"jump_url"`
	Stat           GetDynamicDetailDesktopArchiveStat  `json:"stat"`
	Title          string                              `json:"title"`
	Type           int                                 `json:"type"`
}

// GetDynamicDetailDesktopArchiveBadge 稿件角标信息。
type GetDynamicDetailDesktopArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// GetDynamicDetailDesktopArchiveStat 稿件的播放与弹幕数。
type GetDynamicDetailDesktopArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// GetDynamicDetailDesktopDraw describes images in either an outer or original media body.
type GetDynamicDetailDesktopDraw struct {
	ID    json.Number                       `json:"id"`
	Items []GetDynamicDetailDesktopDrawItem `json:"items"`
}

// GetDynamicDetailDesktopDrawItem 图集中的一张图片。
type GetDynamicDetailDesktopDrawItem struct {
	Height json.Number `json:"height"`
	Size   json.Number `json:"size"`
	Src    string      `json:"src"`
	Tags   []any       `json:"tags"`
	Width  json.Number `json:"width"`
}

// GetDynamicDetailDesktopOriginalItem is the original item embedded in a repost; it is not recursive.
type GetDynamicDetailDesktopOriginalItem struct {
	Basic   GetDynamicDetailDesktopOriginalBasic   `json:"basic"`
	IDStr   json.Number                            `json:"id_str"`
	Modules GetDynamicDetailDesktopOriginalModules `json:"modules"`
	Type    string                                 `json:"type"`
	Visible bool                                   `json:"visible"`
}

// GetDynamicDetailDesktopOriginalBasic contains the original item identifiers, whose types differ from the outer item.
type GetDynamicDetailDesktopOriginalBasic struct {
	CommentIDStr string                                  `json:"comment_id_str"`
	CommentType  int                                     `json:"comment_type"`
	LikeIcon     GetDynamicDetailDesktopOriginalLikeIcon `json:"like_icon"`
	RidStr       string                                  `json:"rid_str"`
}

// GetDynamicDetailDesktopOriginalLikeIcon 原始动态的点赞图标；id 为整型，与外层不同。
type GetDynamicDetailDesktopOriginalLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int    `json:"id"`
	StartURL  string `json:"start_url"`
}

// GetDynamicDetailDesktopOriginalModules groups the author and content of an original item.
type GetDynamicDetailDesktopOriginalModules struct {
	ModuleAuthor  GetDynamicDetailDesktopOriginalModuleAuthor  `json:"module_author"`
	ModuleDynamic GetDynamicDetailDesktopOriginalModuleDynamic `json:"module_dynamic"`
}

// GetDynamicDetailDesktopOriginalModuleAuthor describes the original author using the original response schema.
type GetDynamicDetailDesktopOriginalModuleAuthor struct {
	Avatar         GetDynamicDetailDesktopOriginalAuthorAvatar `json:"avatar"`
	Decorate       GetDynamicDetailDesktopDecorate             `json:"decorate,omitempty"`
	Face           string                                      `json:"face"`
	FaceNft        bool                                        `json:"face_nft"`
	Following      bilibili.NumberOrString                     `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL        string                                      `json:"jump_url"`
	Label          string                                      `json:"label"`
	Mid            json.Number                                 `json:"mid"`
	Name           string                                      `json:"name"`
	OfficialVerify GetDynamicDetailDesktopOfficialVerify       `json:"official_verify"`
	Pendant        GetDynamicDetailDesktopOriginalPendant      `json:"pendant"`
	PubAction      string                                      `json:"pub_action"`
	PubTime        string                                      `json:"pub_time"`
	PubTs          json.Number                                 `json:"pub_ts"`
	Type           string                                      `json:"type"`
	Vip            GetDynamicDetailDesktopAuthorVip            `json:"vip"`
}

// GetDynamicDetailDesktopDecorate 原始动态作者的装扮卡片信息。
type GetDynamicDetailDesktopDecorate struct {
	CardURL string                             `json:"card_url"`
	Fan     GetDynamicDetailDesktopDecorateFan `json:"fan"`
	ID      json.Number                        `json:"id"`
	JumpURL string                             `json:"jump_url"`
	Name    string                             `json:"name"`
	Type    int                                `json:"type"`
}

// GetDynamicDetailDesktopDecorateFan 装扮卡片的粉丝专属信息。
type GetDynamicDetailDesktopDecorateFan struct {
	Color  string      `json:"color"`
	IsFan  bool        `json:"is_fan"`
	NumStr string      `json:"num_str"`
	Number json.Number `json:"number"`
}

// GetDynamicDetailDesktopOriginalPendant 原始动态作者的挂件信息；pid 为整型，与外层不同。
type GetDynamicDetailDesktopOriginalPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               int         `json:"pid"`
}

// GetDynamicDetailDesktopOriginalAuthorAvatar contains avatar rendering layers from the original item schema.
type GetDynamicDetailDesktopOriginalAuthorAvatar struct {
	ContainerSize  GetDynamicDetailDesktopAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetDynamicDetailDesktopAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetDynamicDetailDesktopAvatarLayerWebCssStyle `json:"web_css_style"`
						} `json:"general_config"`
					} `json:"GENERAL_CFG,omitempty"`
					PendantLayer struct {
					} `json:"PENDENT_LAYER,omitempty"`
					IconLayer struct {
					} `json:"ICON_LAYER,omitempty"`
				} `json:"tags"`
			} `json:"layer_config"`
			Resource struct {
				ResImage struct {
					ImageSrc struct {
						Placeholder int `json:"placeholder,omitempty"`
						Remote      struct {
							BfsStyle string `json:"bfs_style"`
							URL      string `json:"url"`
						} `json:"remote,omitempty"`
						SrcType int `json:"src_type"`
						Local   int `json:"local,omitempty"`
					} `json:"image_src"`
				} `json:"res_image"`
				ResType int `json:"res_type"`
			} `json:"resource"`
			Visible bool `json:"visible"`
		} `json:"layers"`
	} `json:"fallback_layers"`
	Mid string `json:"mid"`
}

// GetDynamicDetailDesktopOriginalModuleDynamic contains the original content; Major retains its value semantics.
type GetDynamicDetailDesktopOriginalModuleDynamic struct {
	Additional any                                         `json:"additional"`
	Desc       *GetDynamicDetailDesktopOriginalDescription `json:"desc"`
	Major      GetDynamicDetailDesktopOriginalMajor        `json:"major"`
	Topic      any                                         `json:"topic"`
}

// GetDynamicDetailDesktopOriginalDescription contains original rich text nodes, which differ from outer item nodes.
type GetDynamicDetailDesktopOriginalDescription struct {
	RichTextNodes []GetDynamicDetailDesktopOriginalRichTextNode `json:"rich_text_nodes"`
	Text          string                                        `json:"text"`
}

// GetDynamicDetailDesktopOriginalRichTextNode is one rich text node of an original item description.
type GetDynamicDetailDesktopOriginalRichTextNode struct {
	JumpURL  string                       `json:"jump_url,omitempty"`
	OrigText string                       `json:"orig_text"`
	Text     string                       `json:"text"`
	Type     string                       `json:"type"`
	Emoji    GetDynamicDetailDesktopEmoji `json:"emoji,omitempty"`
}

// GetDynamicDetailDesktopAvatarLayerPosSpec 头像渲染层的位置规格。
type GetDynamicDetailDesktopAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetDynamicDetailDesktopAvatarLayerRenderSpec 头像渲染层的渲染规格。
type GetDynamicDetailDesktopAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetDynamicDetailDesktopAvatarLayerSizeSpec 头像渲染层的尺寸规格。
type GetDynamicDetailDesktopAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetDynamicDetailDesktopAvatarLayerGeneralSpec 头像渲染层的通用规格，三个头像模型共用。
type GetDynamicDetailDesktopAvatarLayerGeneralSpec struct {
	PosSpec    GetDynamicDetailDesktopAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetDynamicDetailDesktopAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetDynamicDetailDesktopAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetDynamicDetailDesktopAvatarLayerWebCssStyle 头像渲染层的 CSS 样式（整体动态与原动态共用；话题仅返回 borderRadius）。
type GetDynamicDetailDesktopAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"background-color,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}

// GetDynamicDetailDesktopOfficialVerify 认证标识（仅 type/desc，用于动态、话题、视频等处的 official_verify 字段）。
type GetDynamicDetailDesktopOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

type GetDynamicDetailDesktopVipLabel struct {
	Path                  string                              `json:"path"`                      // 空。作用尚不明确
	Text                  string                              `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                              `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                              `json:"text_color"`                // 会员标签
	BgStyle               int                                 `json:"bg_style"`                  // 1
	BgColor               string                              `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                              `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                                `json:"use_img_label"`             // true
	ImgLabelURIHans       string                              `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                              `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                              `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                              `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                                 `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                                 `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetDynamicDetailDesktopVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                              `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                              `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetDynamicDetailDesktopVipLabelGoto `label` 中的 `label_goto` 对象
type GetDynamicDetailDesktopVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetDynamicDetailDesktopOriginalMajor 是原动态的媒体主体；字段与整体动态的 Major 一致，仍独立定义。
type GetDynamicDetailDesktopOriginalMajor struct {
	Draw    GetDynamicDetailDesktopDraw    `json:"draw,omitempty"`
	Type    string                         `json:"type"`
	Archive GetDynamicDetailDesktopArchive `json:"archive,omitempty"`
}

// GetDynamicReactionListResult 是动态赞与转发列表接口的 data。
type GetDynamicReactionListResult struct {
	HasMore bool                         `json:"has_more"` // 是否有更多
	Items   []GetDynamicReactionListItem `json:"items"`    // 赞与转发列表
	Offset  string                       `json:"offset"`   // 偏移量，套了一层字符串的 JSON 对象，用于下次请求
	Total   int                          `json:"total"`    // 总数
}

// GetDynamicReactionListItem 是赞与转发列表中的一条记录。
type GetDynamicReactionListItem struct {
	Action string      `json:"action"` // 操作类型：赞了 / 转发了
	Attend int         `json:"attend"` // 参与数：1=对方仅关注了发送者，2=发送者关注了对方
	Desc   string      `json:"desc"`   // 描述，例如「共同关注了3个Up主」
	Face   string      `json:"face"`   // 头像 URL
	Mid    json.Number `json:"mid"`    // 用户 mid；接口可能返回数字或数字字符串
	Name   string      `json:"name"`   // 用户名
}

// GetDynamicForwardListResult 动态转发列表接口的 data。
type GetDynamicForwardListResult struct {
	HasMore bool                        `json:"has_more"` // 是否有更多数据
	Items   []GetDynamicForwardListItem `json:"items"`    // 转发列表
	Offset  string                      `json:"offset"`   // 下一页偏移，用于下一次请求
	Total   int                         `json:"total"`    // 总数
}

// GetDynamicForwardListItem 一条转发动态，转发列表与转发动态信息接口共用。
type GetDynamicForwardListItem struct {
	Desc    GetDynamicForwardListDescription `json:"desc"`     // 转发描述
	IDStr   string                           `json:"id_str"`   // 进行转发的动态 id
	PubTime string                           `json:"pub_time"` // 发布时间，例如「1分钟前」
	User    GetDynamicForwardListUser        `json:"user"`     // 转发用户
}

// GetDynamicForwardListUser 转发动态的发布用户信息。
type GetDynamicForwardListUser struct {
	Face     string                        `json:"face"`     // 头像 URL
	FaceNft  bool                          `json:"face_nft"` // 是否为 NFT 头像
	Mid      json.Number                   `json:"mid"`      // 用户 mid (UID)；接口可能返回数字或数字字符串
	Name     string                        `json:"name"`     // 用户名
	Official GetDynamicForwardListOfficial `json:"official"` // 认证信息
	Pendant  GetDynamicForwardListPendant  `json:"pendant"`  // 装扮信息
	Vip      GetDynamicForwardListVip      `json:"vip"`      // 大会员信息
}

// GetDynamicForwardListVip 转发动态用户的大会员信息。
type GetDynamicForwardListVip struct {
	AvatarIcon         GetDynamicForwardListAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
	AvatarSubscript    int                             `json:"avatar_subscript"`     // 是否显示会员图标：0=不显示，1=显示
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // 大会员角标地址
	DueDate            json.Number                     `json:"due_date"`             // 大会员到期时间，毫秒时间戳
	Label              GetDynamicForwardListVipLabel   `json:"label"`                // 大会员铭牌
	NicknameColor      string                          `json:"nickname_color"`       // 会员昵称颜色
	Role               int                             `json:"role"`                 // 大角色类型：1=月度，3=年度，7=十年，15=百年
	Status             int                             `json:"status"`               // 大会员状态：0=无，1=有
	ThemeType          int                             `json:"theme_type"`           // 会员样式 id
	TvDueDate          json.Number                     `json:"tv_due_date"`          // 电视大会员到期时间，秒级时间戳
	TvVipPayType       int                             `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	TvVipStatus        int                             `json:"tv_vip_status"`        // 电视大会员状态：0=未开通
	Type               int                             `json:"type"`                 // 会员类型：0=无，1=月大会员，2=年度及以上大会员
	VipPayType         int                             `json:"vip_pay_type"`         // 支付类型：0=未支付，1=已支付
}

// GetDynamicForwardListDescription contains text and rich text nodes for the outer item.
type GetDynamicForwardListDescription struct {
	RichTextNodes []GetDynamicForwardListRichTextNode `json:"rich_text_nodes"`
	Text          string                              `json:"text"`
}

// GetDynamicForwardListRichTextNode is one rich text node of an outer item description.
type GetDynamicForwardListRichTextNode struct {
	OrigText string                     `json:"orig_text"`
	Text     string                     `json:"text"`
	Type     string                     `json:"type"`
	JumpURL  string                     `json:"jump_url,omitempty"`
	Style    any                        `json:"style"`
	Emoji    GetDynamicForwardListEmoji `json:"emoji,omitempty"`
	Rid      string                     `json:"rid,omitempty"`
}

// GetDynamicForwardListEmoji is an inline emoji of a rich text node.
type GetDynamicForwardListEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetDynamicForwardListPendant 外层动态作者的挂件信息，数值字段为 json.Number。
type GetDynamicForwardListPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetDynamicForwardListOfficial 是成员认证信息（role/title/desc/type 四字段）。
//
// 与 OfficialVerify（仅 type/desc）语义相近但字段集不同，两者保留独立类型。
type GetDynamicForwardListOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetDynamicForwardListAvatarIcon 是大会员角标信息，对应 vip 中的 avatar_icon 对象。
type GetDynamicForwardListAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

type GetDynamicForwardListVipLabel struct {
	Path                  string                            `json:"path"`                      // 空。作用尚不明确
	Text                  string                            `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                            `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                            `json:"text_color"`                // 会员标签
	BgStyle               int                               `json:"bg_style"`                  // 1
	BgColor               string                            `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                            `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                              `json:"use_img_label"`             // true
	ImgLabelURIHans       string                            `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                            `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                            `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                            `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                               `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                               `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetDynamicForwardListVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                            `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                            `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetDynamicForwardListVipLabelGoto `label` 中的 `label_goto` 对象
type GetDynamicForwardListVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetDynamicPicturesResult 是动态图片列表接口的 data（data 本身即为数组）。
type GetDynamicPicturesResult []GetDynamicPicturesItem

// GetDynamicPicturesItem 是动态中的一张图片。
type GetDynamicPicturesItem struct {
	Height int     `json:"height"` // 图片高度
	Size   float64 `json:"size"`   // 图片大小，单位 KB
	Src    string  `json:"src"`    // 图片 URL
	Width  int     `json:"width"`  // 图片宽度
}

// GetDynamicForwardItemResult 获取转发动态信息接口的 data。
type GetDynamicForwardItemResult struct {
	Item GetDynamicForwardItemItem `json:"item"` // 转发动态信息
}

// GetDynamicForwardItemItem 一条转发动态，转发列表与转发动态信息接口共用。
type GetDynamicForwardItemItem struct {
	Desc    GetDynamicForwardItemDescription `json:"desc"`     // 转发描述
	IDStr   string                           `json:"id_str"`   // 进行转发的动态 id
	PubTime string                           `json:"pub_time"` // 发布时间，例如「1分钟前」
	User    GetDynamicForwardItemUser        `json:"user"`     // 转发用户
}

// GetDynamicForwardItemUser 转发动态的发布用户信息。
type GetDynamicForwardItemUser struct {
	Face     string                        `json:"face"`     // 头像 URL
	FaceNft  bool                          `json:"face_nft"` // 是否为 NFT 头像
	Mid      json.Number                   `json:"mid"`      // 用户 mid (UID)；接口可能返回数字或数字字符串
	Name     string                        `json:"name"`     // 用户名
	Official GetDynamicForwardItemOfficial `json:"official"` // 认证信息
	Pendant  GetDynamicForwardItemPendant  `json:"pendant"`  // 装扮信息
	Vip      GetDynamicForwardItemVip      `json:"vip"`      // 大会员信息
}

// GetDynamicForwardItemVip 转发动态用户的大会员信息。
type GetDynamicForwardItemVip struct {
	AvatarIcon         GetDynamicForwardItemAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
	AvatarSubscript    int                             `json:"avatar_subscript"`     // 是否显示会员图标：0=不显示，1=显示
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // 大会员角标地址
	DueDate            json.Number                     `json:"due_date"`             // 大会员到期时间，毫秒时间戳
	Label              GetDynamicForwardItemVipLabel   `json:"label"`                // 大会员铭牌
	NicknameColor      string                          `json:"nickname_color"`       // 会员昵称颜色
	Role               int                             `json:"role"`                 // 大角色类型：1=月度，3=年度，7=十年，15=百年
	Status             int                             `json:"status"`               // 大会员状态：0=无，1=有
	ThemeType          int                             `json:"theme_type"`           // 会员样式 id
	TvDueDate          json.Number                     `json:"tv_due_date"`          // 电视大会员到期时间，秒级时间戳
	TvVipPayType       int                             `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	TvVipStatus        int                             `json:"tv_vip_status"`        // 电视大会员状态：0=未开通
	Type               int                             `json:"type"`                 // 会员类型：0=无，1=月大会员，2=年度及以上大会员
	VipPayType         int                             `json:"vip_pay_type"`         // 支付类型：0=未支付，1=已支付
}

// GetDynamicForwardItemDescription contains text and rich text nodes for the outer item.
type GetDynamicForwardItemDescription struct {
	RichTextNodes []GetDynamicForwardItemRichTextNode `json:"rich_text_nodes"`
	Text          string                              `json:"text"`
}

// GetDynamicForwardItemRichTextNode is one rich text node of an outer item description.
type GetDynamicForwardItemRichTextNode struct {
	OrigText string                     `json:"orig_text"`
	Text     string                     `json:"text"`
	Type     string                     `json:"type"`
	JumpURL  string                     `json:"jump_url,omitempty"`
	Style    any                        `json:"style"`
	Emoji    GetDynamicForwardItemEmoji `json:"emoji,omitempty"`
	Rid      string                     `json:"rid,omitempty"`
}

// GetDynamicForwardItemEmoji is an inline emoji of a rich text node.
type GetDynamicForwardItemEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetDynamicForwardItemPendant 外层动态作者的挂件信息，数值字段为 json.Number。
type GetDynamicForwardItemPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetDynamicForwardItemOfficial 是成员认证信息（role/title/desc/type 四字段）。
//
// 与 OfficialVerify（仅 type/desc）语义相近但字段集不同，两者保留独立类型。
type GetDynamicForwardItemOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetDynamicForwardItemAvatarIcon 是大会员角标信息，对应 vip 中的 avatar_icon 对象。
type GetDynamicForwardItemAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

type GetDynamicForwardItemVipLabel struct {
	Path                  string                            `json:"path"`                      // 空。作用尚不明确
	Text                  string                            `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                            `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                            `json:"text_color"`                // 会员标签
	BgStyle               int                               `json:"bg_style"`                  // 1
	BgColor               string                            `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                            `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                              `json:"use_img_label"`             // true
	ImgLabelURIHans       string                            `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                            `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                            `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                            `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                               `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                               `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetDynamicForwardItemVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                            `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                            `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetDynamicForwardItemVipLabelGoto `label` 中的 `label_goto` 对象
type GetDynamicForwardItemVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}
