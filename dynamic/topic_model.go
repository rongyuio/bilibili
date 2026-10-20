package dynamic

// 话题动态相关响应模型。
//
// 本包不复用根包或其他子包的响应体：每个接口的 data 层以及所有嵌套对象、
// 数组元素都以接口方法名为前缀独立定义，即使字段完全相同也各自定义，
// 便于后续单个接口的字段变更不影响其他调用。

// GetTopicFeedResult contains the data payload of one topic feed page.
type GetTopicFeedResult struct {
	RelatedTopics GetTopicFeedRelatedTopics `json:"related_topics"`
	TopicCardList GetTopicFeedCardList      `json:"topic_card_list"`
}

// GetTopicFeedRelatedTopics 是话题 feed 的关联话题字段。当前接口返回的字段不稳定。
type GetTopicFeedRelatedTopics struct{}

// GetTopicFeedCardList contains topic cards, a continuation offset and sorting metadata.
type GetTopicFeedCardList struct {
	HasMore         bool                   `json:"has_more"`
	Items           []GetTopicFeedItem     `json:"items"`
	Offset          string                 `json:"offset"`
	TopicSortByConf GetTopicFeedSortByConf `json:"topic_sort_by_conf"`
}

// GetTopicFeedItem wraps a dynamic card and its topic type.
type GetTopicFeedItem struct {
	DynamicCardItem GetTopicFeedDynamicCard `json:"dynamic_card_item"`
	TopicType       string                  `json:"topic_type"`
}

// GetTopicFeedSortByConf describes the sorting choices returned by the topic feed.
type GetTopicFeedSortByConf struct {
	AllSortBy     []GetTopicFeedSortByItem `json:"all_sort_by"`
	DefaultSortBy int                      `json:"default_sort_by"`
	ShowSortBy    int                      `json:"show_sort_by"`
}

// GetTopicFeedSortByItem 是一项可选排序方式。
type GetTopicFeedSortByItem struct {
	SortBy   int    `json:"sort_by"`
	SortName string `json:"sort_name"`
}

// GetTopicFeedDynamicCard is a dynamic card in a topic feed, using the topic response schema.
type GetTopicFeedDynamicCard struct {
	Basic   GetTopicFeedDynamicBasic   `json:"basic"`
	IDStr   string                     `json:"id_str"`
	Modules GetTopicFeedDynamicModules `json:"modules"`
	Type    string                     `json:"type"`
	Visible bool                       `json:"visible"`
}

// GetTopicFeedDynamicBasic contains comment identifiers and actions for a topic card.
type GetTopicFeedDynamicBasic struct {
	CommentIDStr string               `json:"comment_id_str"`
	CommentType  int                  `json:"comment_type"`
	LikeIcon     GetTopicFeedLikeIcon `json:"like_icon"`
	RidStr       string               `json:"rid_str"`
}

// GetTopicFeedLikeIcon is the like icon metadata of a topic card.
type GetTopicFeedLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int64  `json:"id"`
	StartURL  string `json:"start_url"`
}

// GetTopicFeedDynamicModules groups the author, content, actions and counts of a topic card.
type GetTopicFeedDynamicModules struct {
	ModuleAuthor  GetTopicFeedModuleAuthor  `json:"module_author"`
	ModuleDynamic GetTopicFeedModuleDynamic `json:"module_dynamic"`
	ModuleMore    GetTopicFeedModuleMore    `json:"module_more"`
	ModuleStat    GetTopicFeedModuleStat    `json:"module_stat"`
}

// GetTopicFeedModuleAuthor describes a topic card author; Following has no fixed schema yet.
type GetTopicFeedModuleAuthor struct {
	Avatar          GetTopicFeedAuthorAvatar   `json:"avatar"`
	Face            string                     `json:"face"`
	FaceNft         bool                       `json:"face_nft"`
	Following       any                        `json:"following"`
	JumpURL         string                     `json:"jump_url"`
	Label           string                     `json:"label"`
	Mid             int64                      `json:"mid"`
	Name            string                     `json:"name"`
	OfficialVerify  GetTopicFeedOfficialVerify `json:"official_verify"`
	Pendant         GetTopicFeedPendant        `json:"pendant"`
	PubAction       string                     `json:"pub_action"`
	PubLocationText string                     `json:"pub_location_text"`
	PubTime         string                     `json:"pub_time"`
	PubTs           int                        `json:"pub_ts"`
	Type            string                     `json:"type"`
	Vip             GetTopicFeedVip            `json:"vip"`
}

// GetTopicFeedPendant 是话题动态作者的挂件信息，数值字段为 int。
type GetTopicFeedPendant struct {
	Expire            int    `json:"expire"`
	Image             string `json:"image"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
	NPid              int    `json:"n_pid"`
	Name              string `json:"name"`
	Pid               int    `json:"pid"`
}

// GetTopicFeedVip 是话题动态作者的会员信息；Label 复用共享的 Label 类型。
type GetTopicFeedVip struct {
	AvatarSubscript    int                  `json:"avatar_subscript"`
	AvatarSubscriptURL string               `json:"avatar_subscript_url"`
	DueDate            int                  `json:"due_date"`
	Label              GetTopicFeedVipLabel `json:"label"`
	NicknameColor      string               `json:"nickname_color"`
	Status             int                  `json:"status"`
	ThemeType          int                  `json:"theme_type"`
	Type               int                  `json:"type"`
}

// GetTopicFeedAuthorAvatar contains topic author avatar rendering layers.
type GetTopicFeedAuthorAvatar struct {
	ContainerSize  GetTopicFeedAvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec GetTopicFeedAvatarLayerGeneralSpec `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle struct {
								BorderRadius string `json:"borderRadius"`
							} `json:"web_css_style"`
						} `json:"general_config"`
					} `json:"GENERAL_CFG"`
				} `json:"tags"`
			} `json:"layer_config"`
			Resource struct {
				ResImage struct {
					ImageSrc struct {
						Placeholder int `json:"placeholder"`
						Remote      struct {
							BfsStyle string `json:"bfs_style"`
							URL      string `json:"url"`
						} `json:"remote"`
						SrcType int `json:"src_type"`
					} `json:"image_src"`
				} `json:"res_image"`
				ResType int `json:"res_type"`
			} `json:"resource"`
			Visible bool `json:"visible"`
		} `json:"layers"`
	} `json:"fallback_layers"`
	Mid string `json:"mid"`
}

// GetTopicFeedModuleDynamic contains topic card content and its major media body.
type GetTopicFeedModuleDynamic struct {
	Additional any                      `json:"additional"`
	Desc       any                      `json:"desc"`
	Major      GetTopicFeedDynamicMajor `json:"major"`
	Topic      any                      `json:"topic"`
}

// GetTopicFeedDynamicMajor contains the video body of a topic card.
type GetTopicFeedDynamicMajor struct {
	Archive GetTopicFeedArchive `json:"archive"`
	Type    string              `json:"type"`
}

// GetTopicFeedArchive describes video metadata using the existing topic field names.
type GetTopicFeedArchive struct {
	Aid            string                   `json:"aid"`
	Badge          GetTopicFeedArchiveBadge `json:"badge"`
	Bvid           string                   `json:"bvid"`
	Cover          string                   `json:"cover"`
	Desc           string                   `json:"desc"`
	DisablePreview int                      `json:"disable_preview"`
	DurationText   string                   `json:"duration_text"`
	JumpURL        string                   `json:"jump_url"`
	Stat           GetTopicFeedArchiveStat  `json:"stat"`
	Title          string                   `json:"title"`
	Type           int                      `json:"type"`
}

// GetTopicFeedModuleStat contains topic comment, forward and like counts.
type GetTopicFeedModuleStat struct {
	Comment GetTopicFeedStat     `json:"comment"`
	Forward GetTopicFeedStat     `json:"forward"`
	Like    GetTopicFeedLikeStat `json:"like"`
}

// GetTopicFeedStat 是话题评论或转发的计数信息。
type GetTopicFeedStat struct {
	Count     int  `json:"count"`
	Forbidden bool `json:"forbidden"`
}

// GetTopicFeedLikeStat 是话题点赞计数信息，额外包含点赞状态。
type GetTopicFeedLikeStat struct {
	Count     int  `json:"count"`
	Forbidden bool `json:"forbidden"`
	Status    bool `json:"status"`
}

// GetTopicFeedModuleMore contains the additional actions offered for a feed item.
type GetTopicFeedModuleMore struct {
	ThreePointItems []GetTopicFeedThreePointItem `json:"three_point_items"`
}

// GetTopicFeedThreePointItem is one entry of the additional actions menu.
type GetTopicFeedThreePointItem struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

// GetTopicFeedAvatarContainerSize 是头像容器的尺寸，外层动态、原始动态与话题动态共用。
type GetTopicFeedAvatarContainerSize struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetTopicFeedAvatarLayerGeneralSpec 是头像渲染层的通用规格，三个头像模型共用。
type GetTopicFeedAvatarLayerGeneralSpec struct {
	PosSpec    GetTopicFeedAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetTopicFeedAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetTopicFeedAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetTopicFeedAvatarLayerPosSpec 是头像渲染层的位置规格。
type GetTopicFeedAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetTopicFeedAvatarLayerRenderSpec 是头像渲染层的渲染规格。
type GetTopicFeedAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetTopicFeedAvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type GetTopicFeedAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetTopicFeedAvatarLayerWebCssStyle 是头像渲染层的 CSS 样式（整体动态与原动态共用；话题仅返回 borderRadius）。
type GetTopicFeedAvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"background-color,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}

// GetTopicFeedOfficialVerify 是认证标识（仅 type/desc，用于动态、话题、视频等处的 official_verify 字段）。
type GetTopicFeedOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

type GetTopicFeedVipLabel struct {
	Path                  string                   `json:"path"`                      // 空。作用尚不明确
	Text                  string                   `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                   `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                   `json:"text_color"`                // 会员标签
	BgStyle               int                      `json:"bg_style"`                  // 1
	BgColor               string                   `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                   `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                     `json:"use_img_label"`             // true
	ImgLabelURIHans       string                   `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                   `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                   `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                   `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                      `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                      `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetTopicFeedVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                   `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                   `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetTopicFeedVipLabelGoto `label` 中的 `label_goto` 对象
type GetTopicFeedVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetTopicFeedArchiveBadge 是话题稿件角标信息。
type GetTopicFeedArchiveBadge struct {
	BgColor string `json:"bg_color"` // 背景色
	Color   string `json:"color"`    // 文字颜色
	IconURL any    `json:"icon_url"` // 图标 URL
	Text    string `json:"text"`     // 文本
}

// GetTopicFeedArchiveStat 是话题稿件的播放与弹幕数。
type GetTopicFeedArchiveStat struct {
	Danmaku string `json:"danmaku"` // 弹幕数
	Play    string `json:"play"`    // 播放数
}
