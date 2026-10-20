package dynamic //nolint:dupl // 按接口隔离返回体，与 user_space 动态结构重复是设计使然

import (
	"encoding/json"

	"github.com/rongyuio/bilibili"
)

// GetDynamicFeedAll 的响应模型。
//
// 外层动态与原动态采用独立模型，不合并为递归类型；仅稿件、图集与 major
// 在同一接口内部复用。头像 fallback_layers 是随接口版本漂移的一次性渲染
// 配置树，没有复用价值且层级极深，保留为内联匿名结构体。

// GetDynamicFeedAllResult 是关注动态流的 data 对象。
type GetDynamicFeedAllResult struct {
	HasMore        bool                    `json:"has_more"`        // 是否有更多数据
	Items          []GetDynamicFeedAllItem `json:"items"`           // 数据数组
	Offset         string                  `json:"offset"`          // 偏移量，等于items中最后一条记录的id，获取下一页时使用
	UpdateBaseline string                  `json:"update_baseline"` // 更新基线，等于items中第一条记录的id
	UpdateNum      json.Number             `json:"update_num"`      // 本次获取获取到了多少条新动态，在更新基线以上的动态条数
}

// GetDynamicFeedAllItem 是关注动态流中的一条动态，转发时包含原动态。
type GetDynamicFeedAllItem struct {
	Basic   GetDynamicFeedAllItemBasic    `json:"basic"`
	IDStr   json.Number                   `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules GetDynamicFeedAllItemModules  `json:"modules"`
	Orig    GetDynamicFeedAllOriginalItem `json:"orig,omitempty"`
	Type    string                        `json:"type"`
	Visible bool                          `json:"visible"`
}

// GetDynamicFeedAllItemBasic 是动态的评论标识与操作信息。
type GetDynamicFeedAllItemBasic struct {
	CommentIDStr string                    `json:"comment_id_str"`
	CommentType  int                       `json:"comment_type"`
	LikeIcon     GetDynamicFeedAllLikeIcon `json:"like_icon"`
	RidStr       string                    `json:"rid_str"`
}

// GetDynamicFeedAllLikeIcon 是关注动态流的点赞图标元数据。
type GetDynamicFeedAllLikeIcon struct {
	ActionURL string      `json:"action_url"`
	EndURL    string      `json:"end_url"`
	ID        json.Number `json:"id"`
	StartURL  string      `json:"start_url"`
}

// GetDynamicFeedAllItemModules 汇总动态的作者、内容、操作与计数。
type GetDynamicFeedAllItemModules struct {
	ModuleAuthor  GetDynamicFeedAllModuleAuthor  `json:"module_author"`
	ModuleDynamic GetDynamicFeedAllModuleDynamic `json:"module_dynamic"`
	ModuleMore    GetDynamicFeedAllModuleMore    `json:"module_more"`
	ModuleStat    GetDynamicFeedAllModuleStat    `json:"module_stat"`
}

// GetDynamicFeedAllModuleAuthor 描述外层动态的作者。
type GetDynamicFeedAllModuleAuthor struct {
	Avatar          GetDynamicFeedAllAuthorAvatar   `json:"avatar"`
	Face            string                          `json:"face"`
	FaceNft         bool                            `json:"face_nft"`
	Following       bilibili.NumberOrString         `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL         string                          `json:"jump_url"`
	Label           string                          `json:"label"`
	Mid             json.Number                     `json:"mid"`
	Name            string                          `json:"name"`
	OfficialVerify  GetDynamicFeedAllOfficialVerify `json:"official_verify"`
	Pendant         GetDynamicFeedAllPendant        `json:"pendant"`
	PubAction       string                          `json:"pub_action"`
	PubLocationText string                          `json:"pub_location_text"`
	PubTime         string                          `json:"pub_time"`
	PubTs           json.Number                     `json:"pub_ts"`
	Type            string                          `json:"type"`
	Vip             GetDynamicFeedAllAuthorVip      `json:"vip"`
}

// GetDynamicFeedAllPendant 是外层动态作者的挂件信息，数值字段为 json.Number。
type GetDynamicFeedAllPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// GetDynamicFeedAllAuthorVip 是外层动态与原始动态的作者会员信息。
type GetDynamicFeedAllAuthorVip struct {
	AvatarSubscript    int                    `json:"avatar_subscript"`
	AvatarSubscriptURL string                 `json:"avatar_subscript_url"`
	DueDate            json.Number            `json:"due_date"`
	Label              GetDynamicFeedAllLabel `json:"label"`
	NicknameColor      string                 `json:"nickname_color"`
	Status             int                    `json:"status"`
	ThemeType          int                    `json:"theme_type"`
	Type               int                    `json:"type"`
}

// GetDynamicFeedAllLabel 是外层动态作者的会员标签。
type GetDynamicFeedAllLabel struct {
	Path                  string                     `json:"path"`                      // 空。作用尚不明确
	Text                  string                     `json:"text"`                      // 会员类型文案
	LabelTheme            string                     `json:"label_theme"`               // 会员标签
	TextColor             string                     `json:"text_color"`                // 会员标签
	BgStyle               int                        `json:"bg_style"`                  // 1
	BgColor               string                     `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                     `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                       `json:"use_img_label"`             // true
	ImgLabelURIHans       string                     `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                     `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                     `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                     `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                        `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                        `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetDynamicFeedAllLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                     `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                     `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetDynamicFeedAllLabelGoto 是会员标签中的跳转信息。
type GetDynamicFeedAllLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetDynamicFeedAllOfficialVerify 是外层动态作者的认证标识。
type GetDynamicFeedAllOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetDynamicFeedAllAuthorAvatar 是外层动态作者的头像渲染信息。
type GetDynamicFeedAllAuthorAvatar struct {
	ContainerSize  GetDynamicFeedAllAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetDynamicFeedAllAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetDynamicFeedAllAvatarLayerWebCssStyle `json:"web_css_style"`
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

// GetDynamicFeedAllAvatarContainerSize 是头像容器的尺寸。
type GetDynamicFeedAllAvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetDynamicFeedAllAvatarLayerGeneralSpec 是头像渲染层的通用规格。
type GetDynamicFeedAllAvatarLayerGeneralSpec struct {
	PosSpec    GetDynamicFeedAllAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetDynamicFeedAllAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetDynamicFeedAllAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetDynamicFeedAllAvatarLayerPosSpec 是头像渲染层的位置规格。
type GetDynamicFeedAllAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetDynamicFeedAllAvatarLayerRenderSpec 是头像渲染层的渲染规格。
type GetDynamicFeedAllAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetDynamicFeedAllAvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type GetDynamicFeedAllAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetDynamicFeedAllAvatarLayerWebCssStyle 是头像渲染层的 CSS 样式。
type GetDynamicFeedAllAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}

// GetDynamicFeedAllModuleDynamic 包含外层动态的内容与可选描述、主体。
type GetDynamicFeedAllModuleDynamic struct {
	Additional any                           `json:"additional"`
	Desc       *GetDynamicFeedAllDescription `json:"desc"`
	Major      *GetDynamicFeedAllMajor       `json:"major"`
	Topic      any                           `json:"topic"`
}

// GetDynamicFeedAllDescription 是外层动态的文本与富文本节点。
type GetDynamicFeedAllDescription struct {
	RichTextNodes []GetDynamicFeedAllRichTextNode `json:"rich_text_nodes"`
	Text          string                          `json:"text"`
}

// GetDynamicFeedAllRichTextNode 是外层动态正文的一个富文本节点。
type GetDynamicFeedAllRichTextNode struct {
	OrigText string                 `json:"orig_text"`
	Text     string                 `json:"text"`
	Type     string                 `json:"type"`
	JumpURL  string                 `json:"jump_url,omitempty"`
	Style    any                    `json:"style"`
	Emoji    GetDynamicFeedAllEmoji `json:"emoji,omitempty"`
	Rid      string                 `json:"rid,omitempty"`
}

// GetDynamicFeedAllEmoji 是富文本节点中的内联表情。
type GetDynamicFeedAllEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// GetDynamicFeedAllMajor 包含外层动态的媒体主体。
type GetDynamicFeedAllMajor struct {
	Draw    GetDynamicFeedAllDraw    `json:"draw,omitempty"`
	Type    string                   `json:"type"`
	Archive GetDynamicFeedAllArchive `json:"archive,omitempty"`
}

// GetDynamicFeedAllModuleMore 包含动态提供的附加操作。
type GetDynamicFeedAllModuleMore struct {
	ThreePointItems []GetDynamicFeedAllThreePointItem `json:"three_point_items"`
}

// GetDynamicFeedAllThreePointItem 是附加操作菜单的一项。
type GetDynamicFeedAllThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// GetDynamicFeedAllModuleStat 包含动态的评论、转发与点赞计数。
type GetDynamicFeedAllModuleStat struct {
	Comment GetDynamicFeedAllStat     `json:"comment"`
	Forward GetDynamicFeedAllStat     `json:"forward"`
	Like    GetDynamicFeedAllLikeStat `json:"like"`
}

// GetDynamicFeedAllStat 是评论或转发的计数信息。
type GetDynamicFeedAllStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
}

// GetDynamicFeedAllLikeStat 是点赞计数信息，额外包含当前用户的点赞状态。
type GetDynamicFeedAllLikeStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
	Status    bool        `json:"status"`
}

// GetDynamicFeedAllArchive 描述外层或原动态媒体主体中的视频稿件。
type GetDynamicFeedAllArchive struct {
	Aid            string                        `json:"aid"`
	Badge          GetDynamicFeedAllArchiveBadge `json:"badge"`
	Bvid           string                        `json:"bvid"`
	Cover          string                        `json:"cover"`
	Desc           string                        `json:"desc"`
	DisablePreview int                           `json:"disable_preview"`
	DurationText   string                        `json:"duration_text"`
	JumpURL        string                        `json:"jump_url"`
	Stat           GetDynamicFeedAllArchiveStat  `json:"stat"`
	Title          string                        `json:"title"`
	Type           int                           `json:"type"`
}

// GetDynamicFeedAllArchiveBadge 是稿件角标信息。
type GetDynamicFeedAllArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// GetDynamicFeedAllArchiveStat 是稿件的播放与弹幕数。
type GetDynamicFeedAllArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// GetDynamicFeedAllDraw 描述外层或原动态媒体主体中的图片。
type GetDynamicFeedAllDraw struct {
	ID    json.Number                 `json:"id"`
	Items []GetDynamicFeedAllDrawItem `json:"items"`
}

// GetDynamicFeedAllDrawItem 是图集中的一张图片。
type GetDynamicFeedAllDrawItem struct {
	Height json.Number `json:"height"`
	Size   json.Number `json:"size"`
	Src    string      `json:"src"`
	Tags   []any       `json:"tags"`
	Width  json.Number `json:"width"`
}

// GetDynamicFeedAllOriginalItem 是转发动态内嵌的原动态；它不是递归类型。
type GetDynamicFeedAllOriginalItem struct {
	Basic   GetDynamicFeedAllOriginalBasic   `json:"basic"`
	IDStr   json.Number                      `json:"id_str"`
	Modules GetDynamicFeedAllOriginalModules `json:"modules"`
	Type    string                           `json:"type"`
	Visible bool                             `json:"visible"`
}

// GetDynamicFeedAllOriginalBasic 是原动态的标识信息，其类型与外层不同。
type GetDynamicFeedAllOriginalBasic struct {
	CommentIDStr string                            `json:"comment_id_str"`
	CommentType  int                               `json:"comment_type"`
	LikeIcon     GetDynamicFeedAllOriginalLikeIcon `json:"like_icon"`
	RidStr       string                            `json:"rid_str"`
}

// GetDynamicFeedAllOriginalLikeIcon 是原动态的点赞图标；id 为整型，与外层不同。
type GetDynamicFeedAllOriginalLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int    `json:"id"`
	StartURL  string `json:"start_url"`
}

// GetDynamicFeedAllOriginalModules 汇总原动态的作者与内容。
type GetDynamicFeedAllOriginalModules struct {
	ModuleAuthor  GetDynamicFeedAllOriginalModuleAuthor  `json:"module_author"`
	ModuleDynamic GetDynamicFeedAllOriginalModuleDynamic `json:"module_dynamic"`
}

// GetDynamicFeedAllOriginalModuleAuthor 用原动态的响应结构描述作者。
type GetDynamicFeedAllOriginalModuleAuthor struct {
	Avatar         GetDynamicFeedAllOriginalAuthorAvatar `json:"avatar"`
	Decorate       GetDynamicFeedAllDecorate             `json:"decorate,omitempty"`
	Face           string                                `json:"face"`
	FaceNft        bool                                  `json:"face_nft"`
	Following      bilibili.NumberOrString               `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL        string                                `json:"jump_url"`
	Label          string                                `json:"label"`
	Mid            json.Number                           `json:"mid"`
	Name           string                                `json:"name"`
	OfficialVerify GetDynamicFeedAllOfficialVerify       `json:"official_verify"`
	Pendant        GetDynamicFeedAllOriginalPendant      `json:"pendant"`
	PubAction      string                                `json:"pub_action"`
	PubTime        string                                `json:"pub_time"`
	PubTs          json.Number                           `json:"pub_ts"`
	Type           string                                `json:"type"`
	Vip            GetDynamicFeedAllAuthorVip            `json:"vip"`
}

// GetDynamicFeedAllDecorate 是原动态作者的装扮卡片信息。
type GetDynamicFeedAllDecorate struct {
	CardURL string                       `json:"card_url"`
	Fan     GetDynamicFeedAllDecorateFan `json:"fan"`
	ID      json.Number                  `json:"id"`
	JumpURL string                       `json:"jump_url"`
	Name    string                       `json:"name"`
	Type    int                          `json:"type"`
}

// GetDynamicFeedAllDecorateFan 是装扮卡片的粉丝专属信息。
type GetDynamicFeedAllDecorateFan struct {
	Color  string      `json:"color"`
	IsFan  bool        `json:"is_fan"`
	NumStr string      `json:"num_str"`
	Number json.Number `json:"number"`
}

// GetDynamicFeedAllOriginalPendant 是原动态作者的挂件信息；pid 为整型，与外层不同。
type GetDynamicFeedAllOriginalPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               int         `json:"pid"`
}

// GetDynamicFeedAllOriginalAuthorAvatar 是原动态作者的头像渲染信息。
type GetDynamicFeedAllOriginalAuthorAvatar struct {
	ContainerSize  GetDynamicFeedAllAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetDynamicFeedAllAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle GetDynamicFeedAllAvatarLayerWebCssStyle `json:"web_css_style"`
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

// GetDynamicFeedAllOriginalModuleDynamic 包含原动态的内容。
type GetDynamicFeedAllOriginalModuleDynamic struct {
	Additional any                                   `json:"additional"`
	Desc       *GetDynamicFeedAllOriginalDescription `json:"desc"`
	Major      GetDynamicFeedAllOriginalMajor        `json:"major"`
	Topic      any                                   `json:"topic"`
}

// GetDynamicFeedAllOriginalDescription 是原动态的富文本内容，其节点与外层不同。
type GetDynamicFeedAllOriginalDescription struct {
	RichTextNodes []GetDynamicFeedAllOriginalRichTextNode `json:"rich_text_nodes"`
	Text          string                                  `json:"text"`
}

// GetDynamicFeedAllOriginalRichTextNode 是原动态正文的一个富文本节点。
type GetDynamicFeedAllOriginalRichTextNode struct {
	JumpURL  string                 `json:"jump_url,omitempty"`
	OrigText string                 `json:"orig_text"`
	Text     string                 `json:"text"`
	Type     string                 `json:"type"`
	Emoji    GetDynamicFeedAllEmoji `json:"emoji,omitempty"`
}

// GetDynamicFeedAllOriginalMajor 包含原动态的媒体主体，字段与外层 major 一致。
type GetDynamicFeedAllOriginalMajor struct {
	Draw    GetDynamicFeedAllDraw    `json:"draw,omitempty"`
	Type    string                   `json:"type"`
	Archive GetDynamicFeedAllArchive `json:"archive,omitempty"`
}
