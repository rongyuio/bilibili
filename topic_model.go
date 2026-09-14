package bilibili

// GetTopicFeedResult contains the data payload of one topic feed page.
type GetTopicFeedResult struct {
	RelatedTopics struct {
	} `json:"related_topics"`
	TopicCardList TopicCardList `json:"topic_card_list"`
}

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
	AllSortBy []struct {
		SortBy   int    `json:"sort_by"`
		SortName string `json:"sort_name"`
	} `json:"all_sort_by"`
	DefaultSortBy int `json:"default_sort_by"`
	ShowSortBy    int `json:"show_sort_by"`
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
	CommentIDStr string `json:"comment_id_str"`
	CommentType  int    `json:"comment_type"`
	LikeIcon     struct {
		ActionURL string `json:"action_url"`
		EndURL    string `json:"end_url"`
		ID        int64  `json:"id"`
		StartURL  string `json:"start_url"`
	} `json:"like_icon"`
	RidStr string `json:"rid_str"`
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
	Avatar         TopicAuthorAvatar `json:"avatar"`
	Face           string            `json:"face"`
	FaceNft        bool              `json:"face_nft"`
	Following      interface{}       `json:"following"`
	JumpURL        string            `json:"jump_url"`
	Label          string            `json:"label"`
	Mid            int64             `json:"mid"`
	Name           string            `json:"name"`
	OfficialVerify struct {
		Desc string `json:"desc"`
		Type int    `json:"type"`
	} `json:"official_verify"`
	Pendant struct {
		Expire            int    `json:"expire"`
		Image             string `json:"image"`
		ImageEnhance      string `json:"image_enhance"`
		ImageEnhanceFrame string `json:"image_enhance_frame"`
		NPid              int    `json:"n_pid"`
		Name              string `json:"name"`
		Pid               int    `json:"pid"`
	} `json:"pendant"`
	PubAction       string `json:"pub_action"`
	PubLocationText string `json:"pub_location_text"`
	PubTime         string `json:"pub_time"`
	PubTs           int    `json:"pub_ts"`
	Type            string `json:"type"`
	Vip             struct {
		AvatarSubscript    int    `json:"avatar_subscript"`
		AvatarSubscriptURL string `json:"avatar_subscript_url"`
		DueDate            int    `json:"due_date"`
		Label              struct {
			BgColor               string `json:"bg_color"`
			BgStyle               int    `json:"bg_style"`
			BorderColor           string `json:"border_color"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			LabelTheme            string `json:"label_theme"`
			Path                  string `json:"path"`
			Text                  string `json:"text"`
			TextColor             string `json:"text_color"`
			UseImgLabel           bool   `json:"use_img_label"`
		} `json:"label"`
		NicknameColor string `json:"nickname_color"`
		Status        int    `json:"status"`
		ThemeType     int    `json:"theme_type"`
		Type          int    `json:"type"`
	} `json:"vip"`
}

// TopicAuthorAvatar contains topic author avatar rendering layers.
type TopicAuthorAvatar struct {
	ContainerSize struct {
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
	} `json:"container_size"`
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
	Additional interface{}       `json:"additional"`
	Desc       interface{}       `json:"desc"`
	Major      TopicDynamicMajor `json:"major"`
	Topic      interface{}       `json:"topic"`
}

// TopicDynamicMajor contains the video body of a topic card.
type TopicDynamicMajor struct {
	Archive TopicArchive `json:"archive"`
	Type    string       `json:"type"`
}

// TopicArchive describes video metadata using the existing topic field names.
type TopicArchive struct {
	Aid   string `json:"aid"`
	Badge struct {
		BgColor string      `json:"bg_color"`
		Color   string      `json:"color"`
		IconURL interface{} `json:"icon_url"`
		Text    string      `json:"text"`
	} `json:"badge"`
	Bvid           string `json:"bvid"`
	Cover          string `json:"cover"`
	Desc           string `json:"desc"`
	DisablePreview int    `json:"disable_preview"`
	DurationText   string `json:"duration_text"`
	JumpURL        string `json:"jump_url"`
	Stat           struct {
		Danmaku string `json:"danmaku"`
		Play    string `json:"play"`
	} `json:"stat"`
	Title string `json:"title"`
	Type  int    `json:"type"`
}

// TopicModuleStat contains topic comment, forward and like counts.
type TopicModuleStat struct {
	Comment struct {
		Count     int  `json:"count"`
		Forbidden bool `json:"forbidden"`
	} `json:"comment"`
	Forward struct {
		Count     int  `json:"count"`
		Forbidden bool `json:"forbidden"`
	} `json:"forward"`
	Like struct {
		Count     int  `json:"count"`
		Forbidden bool `json:"forbidden"`
		Status    bool `json:"status"`
	} `json:"like"`
}
