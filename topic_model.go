package bilibili

// 话题相关响应模型。

// GetTopicFeedResult contains the data payload of one topic feed page.
type GetTopicFeedResult struct {
	RelatedTopics TopicRelatedTopics `json:"related_topics"`
	TopicCardList TopicCardList      `json:"topic_card_list"`
}

// TopicRelatedTopics 是话题 feed 的关联话题字段。当前接口返回的字段不稳定，
// 保留为空结构体以维持“忽略未知字段”的既有解码行为。
type TopicRelatedTopics struct{}

// TopicCardList contains topic cards, a continuation offset and sorting metadata.
type TopicCardList struct {
	HasMore         bool            `json:"has_more"`
	Items           []TopicFeedItem `json:"items"`
	Offset          string          `json:"offset"`
	TopicSortByConf TopicSortByConf `json:"topic_sort_by_conf"`
}

// TopicFeedItem wraps a dynamic card and its topic type.
type TopicFeedItem struct {
	DynamicCardItem TopicDynamicCard `json:"dynamic_card_item"`
	TopicType       string           `json:"topic_type"`
}

// TopicSortByConf describes the sorting choices returned by the topic feed.
type TopicSortByConf struct {
	AllSortBy     []TopicSortByItem `json:"all_sort_by"`
	DefaultSortBy int               `json:"default_sort_by"`
	ShowSortBy    int               `json:"show_sort_by"`
}

// TopicSortByItem 是一项可选排序方式。
type TopicSortByItem struct {
	SortBy   int    `json:"sort_by"`
	SortName string `json:"sort_name"`
}

// TopicDynamicCard is a dynamic card in a topic feed, using the topic response schema.
type TopicDynamicCard struct {
	Basic   TopicDynamicBasic   `json:"basic"`
	IDStr   string              `json:"id_str"`
	Modules TopicDynamicModules `json:"modules"`
	Type    string              `json:"type"`
	Visible bool                `json:"visible"`
}

// TopicDynamicBasic contains comment identifiers and actions for a topic card.
type TopicDynamicBasic struct {
	CommentIDStr string        `json:"comment_id_str"`
	CommentType  int           `json:"comment_type"`
	LikeIcon     TopicLikeIcon `json:"like_icon"`
	RidStr       string        `json:"rid_str"`
}

// TopicLikeIcon is the like icon metadata of a topic card.
type TopicLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int64  `json:"id"`
	StartURL  string `json:"start_url"`
}

// TopicDynamicModules groups the author, content, actions and counts of a topic card.
type TopicDynamicModules struct {
	ModuleAuthor  TopicModuleAuthor  `json:"module_author"`
	ModuleDynamic TopicModuleDynamic `json:"module_dynamic"`
	ModuleMore    DynamicModuleMore  `json:"module_more"`
	ModuleStat    TopicModuleStat    `json:"module_stat"`
}

// TopicModuleAuthor describes a topic card author; Following has no fixed schema yet.
type TopicModuleAuthor struct {
	Avatar          TopicAuthorAvatar `json:"avatar"`
	Face            string            `json:"face"`
	FaceNft         bool              `json:"face_nft"`
	Following       any               `json:"following"`
	JumpURL         string            `json:"jump_url"`
	Label           string            `json:"label"`
	Mid             int64             `json:"mid"`
	Name            string            `json:"name"`
	OfficialVerify  OfficialVerify    `json:"official_verify"`
	Pendant         TopicPendant      `json:"pendant"`
	PubAction       string            `json:"pub_action"`
	PubLocationText string            `json:"pub_location_text"`
	PubTime         string            `json:"pub_time"`
	PubTs           int               `json:"pub_ts"`
	Type            string            `json:"type"`
	Vip             TopicVip          `json:"vip"`
}

// TopicPendant 是话题动态作者的挂件信息，数值字段为 int。
type TopicPendant struct {
	Expire            int    `json:"expire"`
	Image             string `json:"image"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
	NPid              int    `json:"n_pid"`
	Name              string `json:"name"`
	Pid               int    `json:"pid"`
}

// TopicVip 是话题动态作者的会员信息；Label 复用共享的 Label 类型。
type TopicVip struct {
	AvatarSubscript    int    `json:"avatar_subscript"`
	AvatarSubscriptURL string `json:"avatar_subscript_url"`
	DueDate            int    `json:"due_date"`
	Label              Label  `json:"label"`
	NicknameColor      string `json:"nickname_color"`
	Status             int    `json:"status"`
	ThemeType          int    `json:"theme_type"`
	Type               int    `json:"type"`
}

// TopicAuthorAvatar contains topic author avatar rendering layers.
type TopicAuthorAvatar struct {
	ContainerSize  AvatarContainerSize `json:"container_size"`
	FallbackLayers struct {
		IsCriticalGroup bool `json:"is_critical_group"`
		Layers          []struct {
			GeneralSpec struct {
				PosSpec struct {
					AxisX         float64 `json:"axis_x"`
					AxisY         float64 `json:"axis_y"`
					CoordinatePos int     `json:"coordinate_pos"`
				} `json:"pos_spec"`
				RenderSpec struct {
					Opacity int `json:"opacity"`
				} `json:"render_spec"`
				SizeSpec struct {
					Height float64 `json:"height"`
					Width  float64 `json:"width"`
				} `json:"size_spec"`
			} `json:"general_spec"`
			LayerConfig struct {
				IsCritical bool `json:"is_critical"`
				Tags       struct {
					AVATARLAYER struct {
					} `json:"AVATAR_LAYER"`
					GENERALCFG struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCSSStyle struct {
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

// TopicModuleDynamic contains topic card content and its major media body.
type TopicModuleDynamic struct {
	Additional any               `json:"additional"`
	Desc       any               `json:"desc"`
	Major      TopicDynamicMajor `json:"major"`
	Topic      any               `json:"topic"`
}

// TopicDynamicMajor contains the video body of a topic card.
type TopicDynamicMajor struct {
	Archive TopicArchive `json:"archive"`
	Type    string       `json:"type"`
}

// TopicArchive describes video metadata using the existing topic field names.
type TopicArchive struct {
	Aid            string            `json:"aid"`
	Badge          TopicArchiveBadge `json:"badge"`
	Bvid           string            `json:"bvid"`
	Cover          string            `json:"cover"`
	Desc           string            `json:"desc"`
	DisablePreview int               `json:"disable_preview"`
	DurationText   string            `json:"duration_text"`
	JumpURL        string            `json:"jump_url"`
	Stat           TopicArchiveStat  `json:"stat"`
	Title          string            `json:"title"`
	Type           int               `json:"type"`
}

// TopicArchiveBadge 是话题稿件角标信息。
type TopicArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconURL any    `json:"icon_url"`
	Text    string `json:"text"`
}

// TopicArchiveStat 是话题稿件的播放与弹幕数。
type TopicArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// TopicModuleStat contains topic comment, forward and like counts.
type TopicModuleStat struct {
	Comment TopicStat     `json:"comment"`
	Forward TopicStat     `json:"forward"`
	Like    TopicLikeStat `json:"like"`
}

// TopicStat 是话题评论或转发的计数信息。
type TopicStat struct {
	Count     int  `json:"count"`
	Forbidden bool `json:"forbidden"`
}

// TopicLikeStat 是话题点赞计数信息，额外包含点赞状态。
type TopicLikeStat struct {
	Count     int  `json:"count"`
	Forbidden bool `json:"forbidden"`
	Status    bool `json:"status"`
}
