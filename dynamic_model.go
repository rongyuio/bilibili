package bilibili

import "encoding/json"

// 动态相关响应模型，包含空间动态（space feed）与其他动态接口。
//
// 外层动态与 DynamicOriginalItem 采用独立模型，不合并为递归类型；
// 仅 DynamicArchive、DynamicDraw 在两处完全一致而复用。
//
// 说明：头像 fallback_layers 是随接口版本漂移的一次性渲染配置树，
// 没有复用价值且层级极深，保留为内联匿名结构体。

// DynamicItem is a space feed item, including its original item when reposted.
type DynamicItem struct {
	Basic   DynamicItemBasic    `json:"basic"`
	IDStr   json.Number         `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
	Modules DynamicItemModules  `json:"modules"`
	Orig    DynamicOriginalItem `json:"orig,omitempty"`
	Type    string              `json:"type"`
	Visible bool                `json:"visible"`
}

// DynamicInfo is a page of space feed items and its pagination metadata.
type DynamicInfo struct {
	HasMore        bool          `json:"has_more"`        // 是否有更多数据
	Items          []DynamicItem `json:"items"`           // 数据数组
	Offset         string        `json:"offset"`          // 偏移量，等于items中最后一条记录的id，获取下一页时使用
	UpdateBaseline string        `json:"update_baseline"` // 更新基线，等于items中第一条记录的id
	UpdateNum      json.Number   `json:"update_num"`      // 本次获取获取到了多少条新动态，在更新基线以上的动态条数
}

// DynamicItemBasic contains comment identifiers and actions for a feed item.
type DynamicItemBasic struct { // 见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/all.md#data%E5%AF%B9%E8%B1%A1---items%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84%E5%AF%B9%E8%B1%A1---basic%E5%AF%B9%E8%B1%A1
	CommentIDStr string          `json:"comment_id_str"`
	CommentType  int             `json:"comment_type"`
	LikeIcon     DynamicLikeIcon `json:"like_icon"`
	RidStr       string          `json:"rid_str"`
}

// DynamicLikeIcon is the like icon metadata of a space feed item.
type DynamicLikeIcon struct {
	ActionURL string      `json:"action_url"`
	EndURL    string      `json:"end_url"`
	ID        json.Number `json:"id"`
	StartURL  string      `json:"start_url"`
}

// DynamicItemModules groups the author, content, actions and counts of a feed item.
type DynamicItemModules struct {
	ModuleAuthor  DynamicModuleAuthor  `json:"module_author"`
	ModuleDynamic DynamicModuleDynamic `json:"module_dynamic"`
	ModuleMore    DynamicModuleMore    `json:"module_more"`
	ModuleStat    DynamicModuleStat    `json:"module_stat"`
}

// DynamicModuleAuthor describes the author of the outer feed item.
type DynamicModuleAuthor struct {
	Avatar          DynamicAuthorAvatar `json:"avatar"`
	Face            string              `json:"face"`
	FaceNft         bool                `json:"face_nft"`
	Following       NumberOrString      `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL         string              `json:"jump_url"`
	Label           string              `json:"label"`
	Mid             json.Number         `json:"mid"`
	Name            string              `json:"name"`
	OfficialVerify  OfficialVerify      `json:"official_verify"`
	Pendant         DynamicPendant      `json:"pendant"`
	PubAction       string              `json:"pub_action"`
	PubLocationText string              `json:"pub_location_text"`
	PubTime         string              `json:"pub_time"`
	PubTs           json.Number         `json:"pub_ts"`
	Type            string              `json:"type"`
	Vip             DynamicAuthorVip    `json:"vip"`
}

// DynamicPendant 是外层动态作者的挂件信息，数值字段为 json.Number。
type DynamicPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               json.Number `json:"pid"`
}

// DynamicAuthorVip 是外层动态与原始动态共用的作者会员信息。
type DynamicAuthorVip struct {
	AvatarSubscript    int         `json:"avatar_subscript"`
	AvatarSubscriptURL string      `json:"avatar_subscript_url"`
	DueDate            json.Number `json:"due_date"`
	Label              Label       `json:"label"`
	NicknameColor      string      `json:"nickname_color"`
	Status             int         `json:"status"`
	ThemeType          int         `json:"theme_type"`
	Type               int         `json:"type"`
}

// DynamicAuthorAvatar contains the outer author avatar rendering layers.
type DynamicAuthorAvatar struct {
	ContainerSize  AvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec AvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle AvatarLayerWebCssStyle `json:"web_css_style"`
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

// DynamicModuleDynamic contains the outer item content and optional description and major body.
type DynamicModuleDynamic struct {
	Additional any                 `json:"additional"`
	Desc       *DynamicDescription `json:"desc"`
	Major      *DynamicMajor       `json:"major"`
	Topic      any                 `json:"topic"`
}

// DynamicDescription contains text and rich text nodes for the outer item.
type DynamicDescription struct {
	RichTextNodes []DynamicRichTextNode `json:"rich_text_nodes"`
	Text          string                `json:"text"`
}

// DynamicRichTextNode is one rich text node of an outer item description.
type DynamicRichTextNode struct {
	OrigText string       `json:"orig_text"`
	Text     string       `json:"text"`
	Type     string       `json:"type"`
	JumpURL  string       `json:"jump_url,omitempty"`
	Style    any          `json:"style"`
	Emoji    DynamicEmoji `json:"emoji,omitempty"`
	Rid      string       `json:"rid,omitempty"`
}

// DynamicEmoji is an inline emoji of a rich text node.
type DynamicEmoji struct {
	IconURL string `json:"icon_url"`
	Size    int    `json:"size"`
	Text    string `json:"text"`
	Type    int    `json:"type"`
}

// DynamicMajor contains the outer item media body.
type DynamicMajor struct {
	Draw    DynamicDraw    `json:"draw,omitempty"`
	Type    string         `json:"type"`
	Archive DynamicArchive `json:"archive,omitempty"`
}

// DynamicModuleMore contains the additional actions offered for a feed item.
type DynamicModuleMore struct {
	ThreePointItems []DynamicThreePointItem `json:"three_point_items"`
}

// DynamicThreePointItem is one entry of the additional actions menu.
type DynamicThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// DynamicModuleStat contains comment, forward and like counts for a feed item.
type DynamicModuleStat struct {
	Comment DynamicStat     `json:"comment"`
	Forward DynamicStat     `json:"forward"`
	Like    DynamicLikeStat `json:"like"`
}

// DynamicStat 是评论或转发的计数信息。
type DynamicStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
}

// DynamicLikeStat 是点赞计数信息，额外包含当前用户的点赞状态。
type DynamicLikeStat struct {
	Count     json.Number `json:"count"`
	Forbidden bool        `json:"forbidden"`
	Status    bool        `json:"status"`
}

// DynamicArchive describes a video archive in either an outer or original media body.
type DynamicArchive struct {
	Aid            string              `json:"aid"`
	Badge          DynamicArchiveBadge `json:"badge"`
	Bvid           string              `json:"bvid"`
	Cover          string              `json:"cover"`
	Desc           string              `json:"desc"`
	DisablePreview int                 `json:"disable_preview"`
	DurationText   string              `json:"duration_text"`
	JumpURL        string              `json:"jump_url"`
	Stat           DynamicArchiveStat  `json:"stat"`
	Title          string              `json:"title"`
	Type           int                 `json:"type"`
}

// DynamicArchiveBadge 是稿件角标信息。
type DynamicArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// DynamicArchiveStat 是稿件的播放与弹幕数。
type DynamicArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// DynamicDraw describes images in either an outer or original media body.
type DynamicDraw struct {
	ID    json.Number       `json:"id"`
	Items []DynamicDrawItem `json:"items"`
}

// DynamicDrawItem 是图集中的一张图片。
type DynamicDrawItem struct {
	Height json.Number `json:"height"`
	Size   json.Number `json:"size"`
	Src    string      `json:"src"`
	Tags   []any       `json:"tags"`
	Width  json.Number `json:"width"`
}

// AvatarContainerSize 是头像容器的尺寸，外层动态、原始动态与话题动态共用。
type AvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}
