package dynamic //nolint:dupl // 按接口隔离返回体，与 feed_all 动态结构重复是设计使然

import (
	"encoding/json"

	"github.com/rongyuio/bilibili"
)

// GetUserSpaceDynamicDesktop 的响应模型。
//
// 外层动态与原动态采用独立模型，不合并为递归类型；仅稿件、图集与 major
// 在同一接口内部复用。头像 fallback_layers 是随接口版本漂移的一次性渲染
// 配置树，没有复用价值且层级极深，保留为内联匿名结构体。

// GetUserSpaceDynamicDesktopResult 是用户空间动态的 data 对象。
type GetUserSpaceDynamicDesktopResult struct {
	HasMore        bool                             `json:"has_more"`        // 是否有更多数据
	Items          []GetUserSpaceDynamicDesktopItem `json:"items"`           // 数据数组
	Offset         string                           `json:"offset"`          // 偏移量，等于items中最后一条记录的id，获取下一页时使用
	UpdateBaseline string                           `json:"update_baseline"` // 更新基线，等于items中第一条记录的id
	UpdateNum      json.Number                      `json:"update_num"`      // 本次获取获取到了多少条新动态，在更新基线以上的动态条数
}

// GetUserSpaceDynamicDesktopItem 是空间动态中的一条动态，转发时包含原动态。
type GetUserSpaceDynamicDesktopItem struct {
	Basic   GetUserSpaceDynamicDesktopItemBasic    `json:"basic"`
	IDStr   json.Number                            `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules GetUserSpaceDynamicDesktopItemModules  `json:"modules"`
	Orig    GetUserSpaceDynamicDesktopOriginalItem `json:"orig,omitempty"`
	Type    string                                 `json:"type"`
	Visible bool                                   `json:"visible"`
}

// GetUserSpaceDynamicDesktopItemBasic 是动态的评论标识与操作信息。
type GetUserSpaceDynamicDesktopItemBasic struct {
	CommentIDStr string                             `json:"comment_id_str"`
	CommentType  int                                `json:"comment_type"`
	LikeIcon     GetUserSpaceDynamicDesktopLikeIcon `json:"like_icon"`
	RidStr       string                             `json:"rid_str"`
}

// GetUserSpaceDynamicDesktopLikeIcon 是空间动态的点赞图标元数据。
type GetUserSpaceDynamicDesktopLikeIcon struct {
	ActionURL string      `json:"action_url"`
	EndURL    string      `json:"end_url"`
	ID        json.Number `json:"id"`
	StartURL  string      `json:"start_url"`
}

// GetUserSpaceDynamicDesktopItemModules 汇总动态的作者、内容、操作与计数。
type GetUserSpaceDynamicDesktopItemModules struct {
	ModuleAuthor  GetUserSpaceDynamicDesktopModuleAuthor  `json:"module_author"`
	ModuleDynamic GetUserSpaceDynamicDesktopModuleDynamic `json:"module_dynamic"`
	ModuleMore    GetUserSpaceDynamicDesktopModuleMore    `json:"module_more"`
	ModuleStat    GetUserSpaceDynamicDesktopModuleStat    `json:"module_stat"`
}

// GetUserSpaceDynamicDesktopModuleAuthor 描述外层动态的作者。
type GetUserSpaceDynamicDesktopModuleAuthor struct {
	Avatar          GetUserSpaceDynamicDesktopAuthorAvatar   `json:"avatar"`
	Face            string                                   `json:"face"`
	FaceNft         bool                                     `json:"face_nft"`
	Following       bilibili.NumberOrString                  `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL         string                                   `json:"jump_url"`
	Label           string                                   `json:"label"`
	Mid             json.Number                              `json:"mid"`
	Name            string                                   `json:"name"`
	OfficialVerify  GetUserSpaceDynamicDesktopOfficialVerify `json:"official_verify"`
	Pendant         GetUserSpaceDynamicDesktopPendant        `json:"pendant"`
	PubAction       string                                   `json:"pub_action"`
	PubLocationText string                                   `json:"pub_location_text"`
	PubTime         string                                   `json:"pub_time"`
	PubTs           json.Number                              `json:"pub_ts"`
	Type            string                                   `json:"type"`
	Vip             GetUserSpaceDynamicDesktopAuthorVip      `json:"vip"`
}

// GetUserSpaceDynamicDesktopPendant 是外层动态作者的挂件信息，数值字段为 json.Number。
type GetUserSpaceDynamicDesktopPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetUserSpaceDynamicDesktopAuthorVip 是外层动态与原始动态的作者会员信息。
type GetUserSpaceDynamicDesktopAuthorVip struct {
	AvatarSubscript    int                             `json:"avatar_subscript"`
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"`
	DueDate            json.Number                     `json:"due_date"`
	Label              GetUserSpaceDynamicDesktopLabel `json:"label"`
	NicknameColor      string                          `json:"nickname_color"`
	Status             int                             `json:"status"`
	ThemeType          int                             `json:"theme_type"`
	Type               int                             `json:"type"`
}

// GetUserSpaceDynamicDesktopLabel 是外层动态作者的会员标签。
type GetUserSpaceDynamicDesktopLabel struct {
	Path                  string                              `json:"path"`                      // 空。作用尚不明确
	Text                  string                              `json:"text"`                      // 会员类型文案
	LabelTheme            string                              `json:"label_theme"`               // 会员标签
	TextColor             string                              `json:"text_color"`                // 会员标签
	BgStyle               int                                 `json:"bg_style"`                  // 1
	BgColor               string                              `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                              `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                                `json:"use_img_label"`             // true
	ImgLabelURIHans       string                              `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                              `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                              `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                              `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                                 `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                                 `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetUserSpaceDynamicDesktopLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                              `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                              `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetUserSpaceDynamicDesktopLabelGoto 是会员标签中的跳转信息。
type GetUserSpaceDynamicDesktopLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetUserSpaceDynamicDesktopOfficialVerify 是外层动态作者的认证标识。
type GetUserSpaceDynamicDesktopOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserSpaceDynamicDesktopAuthorAvatar 是外层动态作者的头像渲染信息。
type GetUserSpaceDynamicDesktopAuthorAvatar struct {
	ContainerSize  GetUserSpaceDynamicDesktopAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetUserSpaceDynamicDesktopAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetUserSpaceDynamicDesktopAvatarLayerWebCssStyle `json:"web_css_style"`
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

// GetUserSpaceDynamicDesktopAvatarContainerSize 是头像容器的尺寸。
type GetUserSpaceDynamicDesktopAvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetUserSpaceDynamicDesktopAvatarLayerGeneralSpec 是头像渲染层的通用规格。
type GetUserSpaceDynamicDesktopAvatarLayerGeneralSpec struct {
	PosSpec    GetUserSpaceDynamicDesktopAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetUserSpaceDynamicDesktopAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetUserSpaceDynamicDesktopAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetUserSpaceDynamicDesktopAvatarLayerPosSpec 是头像渲染层的位置规格。
type GetUserSpaceDynamicDesktopAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetUserSpaceDynamicDesktopAvatarLayerRenderSpec 是头像渲染层的渲染规格。
type GetUserSpaceDynamicDesktopAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetUserSpaceDynamicDesktopAvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type GetUserSpaceDynamicDesktopAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetUserSpaceDynamicDesktopAvatarLayerWebCssStyle 是头像渲染层的 CSS 样式。
type GetUserSpaceDynamicDesktopAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}

// GetUserSpaceDynamicDesktopModuleDynamic 包含外层动态的内容与可选描述、主体。
type GetUserSpaceDynamicDesktopModuleDynamic struct {
	Additional any                                    `json:"additional"`
	Desc       *GetUserSpaceDynamicDesktopDescription `json:"desc"`
	Major      *GetUserSpaceDynamicDesktopMajor       `json:"major"`
	Topic      any                                    `json:"topic"`
}

// GetUserSpaceDynamicDesktopDescription 是外层动态的文本与富文本节点。
type GetUserSpaceDynamicDesktopDescription struct {
	RichTextNodes []GetUserSpaceDynamicDesktopRichTextNode `json:"rich_text_nodes"`
	Text          string                                   `json:"text"`
}

// GetUserSpaceDynamicDesktopRichTextNode 是外层动态正文的一个富文本节点。
type GetUserSpaceDynamicDesktopRichTextNode struct {
	OrigText string                          `json:"orig_text"`
	Text     string                          `json:"text"`
	Type     string                          `json:"type"`
	JumpURL  string                          `json:"jump_url,omitempty"`
	Style    any                             `json:"style"`
	Emoji    GetUserSpaceDynamicDesktopEmoji `json:"emoji,omitempty"`
	Rid      string                          `json:"rid,omitempty"`
}

// GetUserSpaceDynamicDesktopEmoji 是富文本节点中的内联表情。
type GetUserSpaceDynamicDesktopEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetUserSpaceDynamicDesktopMajor 包含外层动态的媒体主体。
type GetUserSpaceDynamicDesktopMajor struct {
	Draw    GetUserSpaceDynamicDesktopDraw    `json:"draw,omitempty"`
	Type    string                            `json:"type"`
	Archive GetUserSpaceDynamicDesktopArchive `json:"archive,omitempty"`
}

// GetUserSpaceDynamicDesktopModuleMore 包含动态提供的附加操作。
type GetUserSpaceDynamicDesktopModuleMore struct {
	ThreePointItems []GetUserSpaceDynamicDesktopThreePointItem `json:"three_point_items"`
}

// GetUserSpaceDynamicDesktopThreePointItem 是附加操作菜单的一项。
type GetUserSpaceDynamicDesktopThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// GetUserSpaceDynamicDesktopModuleStat 包含动态的评论、转发与点赞计数。
type GetUserSpaceDynamicDesktopModuleStat struct {
	Comment GetUserSpaceDynamicDesktopStat     `json:"comment"`
	Forward GetUserSpaceDynamicDesktopStat     `json:"forward"`
	Like    GetUserSpaceDynamicDesktopLikeStat `json:"like"`
}

// GetUserSpaceDynamicDesktopStat 是评论或转发的计数信息。
type GetUserSpaceDynamicDesktopStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
}

// GetUserSpaceDynamicDesktopLikeStat 是点赞计数信息，额外包含当前用户的点赞状态。
type GetUserSpaceDynamicDesktopLikeStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
	Status    bool        `json:"status"`
}

// GetUserSpaceDynamicDesktopArchive 描述外层或原动态媒体主体中的视频稿件。
type GetUserSpaceDynamicDesktopArchive struct {
	Aid            string                                 `json:"aid"`
	Badge          GetUserSpaceDynamicDesktopArchiveBadge `json:"badge"`
	Bvid           string                                 `json:"bvid"`
	Cover          string                                 `json:"cover"`
	Desc           string                                 `json:"desc"`
	DisablePreview int                                    `json:"disable_preview"`
	DurationText   string                                 `json:"duration_text"`
	JumpURL        string                                 `json:"jump_url"`
	Stat           GetUserSpaceDynamicDesktopArchiveStat  `json:"stat"`
	Title          string                                 `json:"title"`
	Type           int                                    `json:"type"`
}

// GetUserSpaceDynamicDesktopArchiveBadge 是稿件角标信息。
type GetUserSpaceDynamicDesktopArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// GetUserSpaceDynamicDesktopArchiveStat 是稿件的播放与弹幕数。
type GetUserSpaceDynamicDesktopArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// GetUserSpaceDynamicDesktopDraw 描述外层或原动态媒体主体中的图片。
type GetUserSpaceDynamicDesktopDraw struct {
	ID    json.Number                          `json:"id"`
	Items []GetUserSpaceDynamicDesktopDrawItem `json:"items"`
}

// GetUserSpaceDynamicDesktopDrawItem 是图集中的一张图片。
type GetUserSpaceDynamicDesktopDrawItem struct {
	Height json.Number `json:"height"`
	Size   json.Number `json:"size"`
	Src    string      `json:"src"`
	Tags   []any       `json:"tags"`
	Width  json.Number `json:"width"`
}

// GetUserSpaceDynamicDesktopOriginalItem 是转发动态内嵌的原动态；它不是递归类型。
type GetUserSpaceDynamicDesktopOriginalItem struct {
	Basic   GetUserSpaceDynamicDesktopOriginalBasic   `json:"basic"`
	IDStr   json.Number                               `json:"id_str"`
	Modules GetUserSpaceDynamicDesktopOriginalModules `json:"modules"`
	Type    string                                    `json:"type"`
	Visible bool                                      `json:"visible"`
}

// GetUserSpaceDynamicDesktopOriginalBasic 是原动态的标识信息，其类型与外层不同。
type GetUserSpaceDynamicDesktopOriginalBasic struct {
	CommentIDStr string                                     `json:"comment_id_str"`
	CommentType  int                                        `json:"comment_type"`
	LikeIcon     GetUserSpaceDynamicDesktopOriginalLikeIcon `json:"like_icon"`
	RidStr       string                                     `json:"rid_str"`
}

// GetUserSpaceDynamicDesktopOriginalLikeIcon 是原动态的点赞图标；id 为整型，与外层不同。
type GetUserSpaceDynamicDesktopOriginalLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int    `json:"id"`
	StartURL  string `json:"start_url"`
}

// GetUserSpaceDynamicDesktopOriginalModules 汇总原动态的作者与内容。
type GetUserSpaceDynamicDesktopOriginalModules struct {
	ModuleAuthor  GetUserSpaceDynamicDesktopOriginalModuleAuthor  `json:"module_author"`
	ModuleDynamic GetUserSpaceDynamicDesktopOriginalModuleDynamic `json:"module_dynamic"`
}

// GetUserSpaceDynamicDesktopOriginalModuleAuthor 用原动态的响应结构描述作者。
type GetUserSpaceDynamicDesktopOriginalModuleAuthor struct {
	Avatar         GetUserSpaceDynamicDesktopOriginalAuthorAvatar `json:"avatar"`
	Decorate       GetUserSpaceDynamicDesktopDecorate             `json:"decorate,omitempty"`
	Face           string                                         `json:"face"`
	FaceNft        bool                                           `json:"face_nft"`
	Following      bilibili.NumberOrString                        `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL        string                                         `json:"jump_url"`
	Label          string                                         `json:"label"`
	Mid            json.Number                                    `json:"mid"`
	Name           string                                         `json:"name"`
	OfficialVerify GetUserSpaceDynamicDesktopOfficialVerify       `json:"official_verify"`
	Pendant        GetUserSpaceDynamicDesktopOriginalPendant      `json:"pendant"`
	PubAction      string                                         `json:"pub_action"`
	PubTime        string                                         `json:"pub_time"`
	PubTs          json.Number                                    `json:"pub_ts"`
	Type           string                                         `json:"type"`
	Vip            GetUserSpaceDynamicDesktopAuthorVip            `json:"vip"`
}

// GetUserSpaceDynamicDesktopDecorate 是原动态作者的装扮卡片信息。
type GetUserSpaceDynamicDesktopDecorate struct {
	CardURL string                                `json:"card_url"`
	Fan     GetUserSpaceDynamicDesktopDecorateFan `json:"fan"`
	ID      json.Number                           `json:"id"`
	JumpURL string                                `json:"jump_url"`
	Name    string                                `json:"name"`
	Type    int                                   `json:"type"`
}

// GetUserSpaceDynamicDesktopDecorateFan 是装扮卡片的粉丝专属信息。
type GetUserSpaceDynamicDesktopDecorateFan struct {
	Color  string      `json:"color"`
	IsFan  bool        `json:"is_fan"`
	NumStr string      `json:"num_str"`
	Number json.Number `json:"number"`
}

// GetUserSpaceDynamicDesktopOriginalPendant 是原动态作者的挂件信息；pid 为整型，与外层不同。
type GetUserSpaceDynamicDesktopOriginalPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               int         `json:"pid"`
}

// GetUserSpaceDynamicDesktopOriginalAuthorAvatar 是原动态作者的头像渲染信息。
type GetUserSpaceDynamicDesktopOriginalAuthorAvatar struct {
	ContainerSize  GetUserSpaceDynamicDesktopAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetUserSpaceDynamicDesktopAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetUserSpaceDynamicDesktopAvatarLayerWebCssStyle `json:"web_css_style"`
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

// GetUserSpaceDynamicDesktopOriginalModuleDynamic 包含原动态的内容。
type GetUserSpaceDynamicDesktopOriginalModuleDynamic struct {
	Additional any                                            `json:"additional"`
	Desc       *GetUserSpaceDynamicDesktopOriginalDescription `json:"desc"`
	Major      GetUserSpaceDynamicDesktopOriginalMajor        `json:"major"`
	Topic      any                                            `json:"topic"`
}

// GetUserSpaceDynamicDesktopOriginalDescription 是原动态的富文本内容，其节点与外层不同。
type GetUserSpaceDynamicDesktopOriginalDescription struct {
	RichTextNodes []GetUserSpaceDynamicDesktopOriginalRichTextNode `json:"rich_text_nodes"`
	Text          string                                           `json:"text"`
}

// GetUserSpaceDynamicDesktopOriginalRichTextNode 是原动态正文的一个富文本节点。
type GetUserSpaceDynamicDesktopOriginalRichTextNode struct {
	JumpURL  string                          `json:"jump_url,omitempty"`
	OrigText string                          `json:"orig_text"`
	Text     string                          `json:"text"`
	Type     string                          `json:"type"`
	Emoji    GetUserSpaceDynamicDesktopEmoji `json:"emoji,omitempty"`
}

// GetUserSpaceDynamicDesktopOriginalMajor 包含原动态的媒体主体，字段与外层 major 一致。
type GetUserSpaceDynamicDesktopOriginalMajor struct {
	Draw    GetUserSpaceDynamicDesktopDraw    `json:"draw,omitempty"`
	Type    string                            `json:"type"`
	Archive GetUserSpaceDynamicDesktopArchive `json:"archive,omitempty"`
}
