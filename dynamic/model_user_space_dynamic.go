package dynamic //nolint:dupl // 按接口隔离返回体，与 feed_all 动态结构重复是设计使然

import (
	"encoding/json"

	"github.com/rongyuio/bilibili"
)

// GetUserSpaceDynamic 的响应模型。
//
// 外层动态与原动态采用独立模型，不合并为递归类型；仅稿件、图集与 major
// 在同一接口内部复用。头像 fallback_layers 是随接口版本漂移的一次性渲染
// 配置树，没有复用价值且层级极深，保留为内联匿名结构体。

// GetUserSpaceDynamicResult 是用户空间动态的 data 对象。
type GetUserSpaceDynamicResult struct {
	HasMore        bool                      `json:"has_more"`        // 是否有更多数据
	Items          []GetUserSpaceDynamicItem `json:"items"`           // 数据数组
	Offset         string                    `json:"offset"`          // 偏移量，等于items中最后一条记录的id，获取下一页时使用
	UpdateBaseline string                    `json:"update_baseline"` // 更新基线，等于items中第一条记录的id
	UpdateNum      json.Number               `json:"update_num"`      // 本次获取获取到了多少条新动态，在更新基线以上的动态条数
}

// GetUserSpaceDynamicItem 是空间动态中的一条动态，转发时包含原动态。
type GetUserSpaceDynamicItem struct {
	Basic   GetUserSpaceDynamicItemBasic    `json:"basic"`
	IDStr   json.Number                     `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules GetUserSpaceDynamicItemModules  `json:"modules"`
	Orig    GetUserSpaceDynamicOriginalItem `json:"orig,omitempty"`
	Type    string                          `json:"type"`
	Visible bool                            `json:"visible"`
}

// GetUserSpaceDynamicItemBasic 是动态的评论标识与操作信息。
type GetUserSpaceDynamicItemBasic struct {
	CommentIDStr string                      `json:"comment_id_str"`
	CommentType  int                         `json:"comment_type"`
	LikeIcon     GetUserSpaceDynamicLikeIcon `json:"like_icon"`
	RidStr       string                      `json:"rid_str"`
}

// GetUserSpaceDynamicLikeIcon 是空间动态的点赞图标元数据。
type GetUserSpaceDynamicLikeIcon struct {
	ActionURL string      `json:"action_url"`
	EndURL    string      `json:"end_url"`
	ID        json.Number `json:"id"`
	StartURL  string      `json:"start_url"`
}

// GetUserSpaceDynamicItemModules 汇总动态的作者、内容、操作与计数。
type GetUserSpaceDynamicItemModules struct {
	ModuleAuthor  GetUserSpaceDynamicModuleAuthor  `json:"module_author"`
	ModuleDynamic GetUserSpaceDynamicModuleDynamic `json:"module_dynamic"`
	ModuleMore    GetUserSpaceDynamicModuleMore    `json:"module_more"`
	ModuleStat    GetUserSpaceDynamicModuleStat    `json:"module_stat"`
}

// GetUserSpaceDynamicModuleAuthor 描述外层动态的作者。
type GetUserSpaceDynamicModuleAuthor struct {
	Avatar          GetUserSpaceDynamicAuthorAvatar   `json:"avatar"`
	Face            string                            `json:"face"`
	FaceNft         bool                              `json:"face_nft"`
	Following       bilibili.NumberOrString           `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL         string                            `json:"jump_url"`
	Label           string                            `json:"label"`
	Mid             json.Number                       `json:"mid"`
	Name            string                            `json:"name"`
	OfficialVerify  GetUserSpaceDynamicOfficialVerify `json:"official_verify"`
	Pendant         GetUserSpaceDynamicPendant        `json:"pendant"`
	PubAction       string                            `json:"pub_action"`
	PubLocationText string                            `json:"pub_location_text"`
	PubTime         string                            `json:"pub_time"`
	PubTs           json.Number                       `json:"pub_ts"`
	Type            string                            `json:"type"`
	Vip             GetUserSpaceDynamicAuthorVip      `json:"vip"`
}

// GetUserSpaceDynamicPendant 是外层动态作者的挂件信息，数值字段为 json.Number。
type GetUserSpaceDynamicPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetUserSpaceDynamicAuthorVip 是外层动态与原始动态的作者会员信息。
type GetUserSpaceDynamicAuthorVip struct {
	AvatarSubscript    int                      `json:"avatar_subscript"`
	AvatarSubscriptURL string                   `json:"avatar_subscript_url"`
	DueDate            json.Number              `json:"due_date"`
	Label              GetUserSpaceDynamicLabel `json:"label"`
	NicknameColor      string                   `json:"nickname_color"`
	Status             int                      `json:"status"`
	ThemeType          int                      `json:"theme_type"`
	Type               int                      `json:"type"`
}

// GetUserSpaceDynamicLabel 是外层动态作者的会员标签。
type GetUserSpaceDynamicLabel struct {
	Path                  string                       `json:"path"`                      // 空。作用尚不明确
	Text                  string                       `json:"text"`                      // 会员类型文案
	LabelTheme            string                       `json:"label_theme"`               // 会员标签
	TextColor             string                       `json:"text_color"`                // 会员标签
	BgStyle               int                          `json:"bg_style"`                  // 1
	BgColor               string                       `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                       `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                         `json:"use_img_label"`             // true
	ImgLabelURIHans       string                       `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                       `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                       `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                       `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                          `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                          `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetUserSpaceDynamicLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                       `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                       `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetUserSpaceDynamicLabelGoto 是会员标签中的跳转信息。
type GetUserSpaceDynamicLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetUserSpaceDynamicOfficialVerify 是外层动态作者的认证标识。
type GetUserSpaceDynamicOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserSpaceDynamicAuthorAvatar 是外层动态作者的头像渲染信息。
type GetUserSpaceDynamicAuthorAvatar struct {
	ContainerSize  GetUserSpaceDynamicAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetUserSpaceDynamicAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetUserSpaceDynamicAvatarLayerWebCssStyle `json:"web_css_style"`
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

// GetUserSpaceDynamicAvatarContainerSize 是头像容器的尺寸。
type GetUserSpaceDynamicAvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetUserSpaceDynamicAvatarLayerGeneralSpec 是头像渲染层的通用规格。
type GetUserSpaceDynamicAvatarLayerGeneralSpec struct {
	PosSpec    GetUserSpaceDynamicAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetUserSpaceDynamicAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetUserSpaceDynamicAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetUserSpaceDynamicAvatarLayerPosSpec 是头像渲染层的位置规格。
type GetUserSpaceDynamicAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetUserSpaceDynamicAvatarLayerRenderSpec 是头像渲染层的渲染规格。
type GetUserSpaceDynamicAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetUserSpaceDynamicAvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type GetUserSpaceDynamicAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetUserSpaceDynamicAvatarLayerWebCssStyle 是头像渲染层的 CSS 样式。
type GetUserSpaceDynamicAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}

// GetUserSpaceDynamicModuleDynamic 包含外层动态的内容与可选描述、主体。
type GetUserSpaceDynamicModuleDynamic struct {
	Additional any                             `json:"additional"`
	Desc       *GetUserSpaceDynamicDescription `json:"desc"`
	Major      *GetUserSpaceDynamicMajor       `json:"major"`
	Topic      any                             `json:"topic"`
}

// GetUserSpaceDynamicDescription 是外层动态的文本与富文本节点。
type GetUserSpaceDynamicDescription struct {
	RichTextNodes []GetUserSpaceDynamicRichTextNode `json:"rich_text_nodes"`
	Text          string                            `json:"text"`
}

// GetUserSpaceDynamicRichTextNode 是外层动态正文的一个富文本节点。
type GetUserSpaceDynamicRichTextNode struct {
	OrigText string                   `json:"orig_text"`
	Text     string                   `json:"text"`
	Type     string                   `json:"type"`
	JumpURL  string                   `json:"jump_url,omitempty"`
	Style    any                      `json:"style"`
	Emoji    GetUserSpaceDynamicEmoji `json:"emoji,omitempty"`
	Rid      string                   `json:"rid,omitempty"`
}

// GetUserSpaceDynamicEmoji 是富文本节点中的内联表情。
type GetUserSpaceDynamicEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetUserSpaceDynamicMajor 包含外层动态的媒体主体。
type GetUserSpaceDynamicMajor struct {
	Draw    GetUserSpaceDynamicDraw    `json:"draw,omitempty"`
	Type    string                     `json:"type"`
	Archive GetUserSpaceDynamicArchive `json:"archive,omitempty"`
}

// GetUserSpaceDynamicModuleMore 包含动态提供的附加操作。
type GetUserSpaceDynamicModuleMore struct {
	ThreePointItems []GetUserSpaceDynamicThreePointItem `json:"three_point_items"`
}

// GetUserSpaceDynamicThreePointItem 是附加操作菜单的一项。
type GetUserSpaceDynamicThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// GetUserSpaceDynamicModuleStat 包含动态的评论、转发与点赞计数。
type GetUserSpaceDynamicModuleStat struct {
	Comment GetUserSpaceDynamicStat     `json:"comment"`
	Forward GetUserSpaceDynamicStat     `json:"forward"`
	Like    GetUserSpaceDynamicLikeStat `json:"like"`
}

// GetUserSpaceDynamicStat 是评论或转发的计数信息。
type GetUserSpaceDynamicStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
}

// GetUserSpaceDynamicLikeStat 是点赞计数信息，额外包含当前用户的点赞状态。
type GetUserSpaceDynamicLikeStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
	Status    bool        `json:"status"`
}

// GetUserSpaceDynamicArchive 描述外层或原动态媒体主体中的视频稿件。
type GetUserSpaceDynamicArchive struct {
	Aid            string                          `json:"aid"`
	Badge          GetUserSpaceDynamicArchiveBadge `json:"badge"`
	Bvid           string                          `json:"bvid"`
	Cover          string                          `json:"cover"`
	Desc           string                          `json:"desc"`
	DisablePreview int                             `json:"disable_preview"`
	DurationText   string                          `json:"duration_text"`
	JumpURL        string                          `json:"jump_url"`
	Stat           GetUserSpaceDynamicArchiveStat  `json:"stat"`
	Title          string                          `json:"title"`
	Type           int                             `json:"type"`
}

// GetUserSpaceDynamicArchiveBadge 是稿件角标信息。
type GetUserSpaceDynamicArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// GetUserSpaceDynamicArchiveStat 是稿件的播放与弹幕数。
type GetUserSpaceDynamicArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// GetUserSpaceDynamicDraw 描述外层或原动态媒体主体中的图片。
type GetUserSpaceDynamicDraw struct {
	ID    json.Number                   `json:"id"`
	Items []GetUserSpaceDynamicDrawItem `json:"items"`
}

// GetUserSpaceDynamicDrawItem 是图集中的一张图片。
type GetUserSpaceDynamicDrawItem struct {
	Height json.Number `json:"height"`
	Size   json.Number `json:"size"`
	Src    string      `json:"src"`
	Tags   []any       `json:"tags"`
	Width  json.Number `json:"width"`
}

// GetUserSpaceDynamicOriginalItem 是转发动态内嵌的原动态；它不是递归类型。
type GetUserSpaceDynamicOriginalItem struct {
	Basic   GetUserSpaceDynamicOriginalBasic   `json:"basic"`
	IDStr   json.Number                        `json:"id_str"`
	Modules GetUserSpaceDynamicOriginalModules `json:"modules"`
	Type    string                             `json:"type"`
	Visible bool                               `json:"visible"`
}

// GetUserSpaceDynamicOriginalBasic 是原动态的标识信息，其类型与外层不同。
type GetUserSpaceDynamicOriginalBasic struct {
	CommentIDStr string                              `json:"comment_id_str"`
	CommentType  int                                 `json:"comment_type"`
	LikeIcon     GetUserSpaceDynamicOriginalLikeIcon `json:"like_icon"`
	RidStr       string                              `json:"rid_str"`
}

// GetUserSpaceDynamicOriginalLikeIcon 是原动态的点赞图标；id 为整型，与外层不同。
type GetUserSpaceDynamicOriginalLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int    `json:"id"`
	StartURL  string `json:"start_url"`
}

// GetUserSpaceDynamicOriginalModules 汇总原动态的作者与内容。
type GetUserSpaceDynamicOriginalModules struct {
	ModuleAuthor  GetUserSpaceDynamicOriginalModuleAuthor  `json:"module_author"`
	ModuleDynamic GetUserSpaceDynamicOriginalModuleDynamic `json:"module_dynamic"`
}

// GetUserSpaceDynamicOriginalModuleAuthor 用原动态的响应结构描述作者。
type GetUserSpaceDynamicOriginalModuleAuthor struct {
	Avatar         GetUserSpaceDynamicOriginalAuthorAvatar `json:"avatar"`
	Decorate       GetUserSpaceDynamicDecorate             `json:"decorate,omitempty"`
	Face           string                                  `json:"face"`
	FaceNft        bool                                    `json:"face_nft"`
	Following      bilibili.NumberOrString                 `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL        string                                  `json:"jump_url"`
	Label          string                                  `json:"label"`
	Mid            json.Number                             `json:"mid"`
	Name           string                                  `json:"name"`
	OfficialVerify GetUserSpaceDynamicOfficialVerify       `json:"official_verify"`
	Pendant        GetUserSpaceDynamicOriginalPendant      `json:"pendant"`
	PubAction      string                                  `json:"pub_action"`
	PubTime        string                                  `json:"pub_time"`
	PubTs          json.Number                             `json:"pub_ts"`
	Type           string                                  `json:"type"`
	Vip            GetUserSpaceDynamicAuthorVip            `json:"vip"`
}

// GetUserSpaceDynamicDecorate 是原动态作者的装扮卡片信息。
type GetUserSpaceDynamicDecorate struct {
	CardURL string                         `json:"card_url"`
	Fan     GetUserSpaceDynamicDecorateFan `json:"fan"`
	ID      json.Number                    `json:"id"`
	JumpURL string                         `json:"jump_url"`
	Name    string                         `json:"name"`
	Type    int                            `json:"type"`
}

// GetUserSpaceDynamicDecorateFan 是装扮卡片的粉丝专属信息。
type GetUserSpaceDynamicDecorateFan struct {
	Color  string      `json:"color"`
	IsFan  bool        `json:"is_fan"`
	NumStr string      `json:"num_str"`
	Number json.Number `json:"number"`
}

// GetUserSpaceDynamicOriginalPendant 是原动态作者的挂件信息；pid 为整型，与外层不同。
type GetUserSpaceDynamicOriginalPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               int         `json:"pid"`
}

// GetUserSpaceDynamicOriginalAuthorAvatar 是原动态作者的头像渲染信息。
type GetUserSpaceDynamicOriginalAuthorAvatar struct {
	ContainerSize  GetUserSpaceDynamicAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetUserSpaceDynamicAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetUserSpaceDynamicAvatarLayerWebCssStyle `json:"web_css_style"`
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

// GetUserSpaceDynamicOriginalModuleDynamic 包含原动态的内容。
type GetUserSpaceDynamicOriginalModuleDynamic struct {
	Additional any                                     `json:"additional"`
	Desc       *GetUserSpaceDynamicOriginalDescription `json:"desc"`
	Major      GetUserSpaceDynamicOriginalMajor        `json:"major"`
	Topic      any                                     `json:"topic"`
}

// GetUserSpaceDynamicOriginalDescription 是原动态的富文本内容，其节点与外层不同。
type GetUserSpaceDynamicOriginalDescription struct {
	RichTextNodes []GetUserSpaceDynamicOriginalRichTextNode `json:"rich_text_nodes"`
	Text          string                                    `json:"text"`
}

// GetUserSpaceDynamicOriginalRichTextNode 是原动态正文的一个富文本节点。
type GetUserSpaceDynamicOriginalRichTextNode struct {
	JumpURL  string                   `json:"jump_url,omitempty"`
	OrigText string                   `json:"orig_text"`
	Text     string                   `json:"text"`
	Type     string                   `json:"type"`
	Emoji    GetUserSpaceDynamicEmoji `json:"emoji,omitempty"`
}

// GetUserSpaceDynamicOriginalMajor 包含原动态的媒体主体，字段与外层 major 一致。
type GetUserSpaceDynamicOriginalMajor struct {
	Draw    GetUserSpaceDynamicDraw    `json:"draw,omitempty"`
	Type    string                     `json:"type"`
	Archive GetUserSpaceDynamicArchive `json:"archive,omitempty"`
}
