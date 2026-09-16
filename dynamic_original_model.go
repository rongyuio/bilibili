package bilibili

import "encoding/json"

// 原动态（转发动态内嵌的原动态）响应模型。

// DynamicOriginalItem is the original item embedded in a repost; it is not recursive.
type DynamicOriginalItem struct {
	Basic   DynamicOriginalBasic   `json:"basic"`
	IDStr   json.Number            `json:"id_str"`
	Modules DynamicOriginalModules `json:"modules"`
	Type    string                 `json:"type"`
	Visible bool                   `json:"visible"`
}

// DynamicOriginalBasic contains the original item identifiers, whose types differ from the outer item.
type DynamicOriginalBasic struct {
	CommentIDStr string                  `json:"comment_id_str"`
	CommentType  int                     `json:"comment_type"`
	LikeIcon     DynamicOriginalLikeIcon `json:"like_icon"`
	RidStr       string                  `json:"rid_str"`
}

// DynamicOriginalLikeIcon 是原始动态的点赞图标；id 为整型，与外层不同。
type DynamicOriginalLikeIcon struct {
	ActionURL string `json:"action_url"`
	EndURL    string `json:"end_url"`
	ID        int    `json:"id"`
	StartURL  string `json:"start_url"`
}

// DynamicOriginalModules groups the author and content of an original item.
type DynamicOriginalModules struct {
	ModuleAuthor  DynamicOriginalModuleAuthor  `json:"module_author"`
	ModuleDynamic DynamicOriginalModuleDynamic `json:"module_dynamic"`
}

// DynamicOriginalModuleAuthor describes the original author using the original response schema.
type DynamicOriginalModuleAuthor struct {
	Avatar         DynamicOriginalAuthorAvatar `json:"avatar"`
	Decorate       DynamicDecorate             `json:"decorate,omitempty"`
	Face           string                      `json:"face"`
	FaceNft        bool                        `json:"face_nft"`
	Following      NumberOrString              `json:"following"` // 关注状态漂移：未登录返回 null、登录态返回布尔、早期调试样本返回数字 1/2；完整状态含义待确认，不定义状态常量。
	JumpURL        string                      `json:"jump_url"`
	Label          string                      `json:"label"`
	Mid            json.Number                 `json:"mid"`
	Name           string                      `json:"name"`
	OfficialVerify OfficialVerify              `json:"official_verify"`
	Pendant        DynamicOriginalPendant      `json:"pendant"`
	PubAction      string                      `json:"pub_action"`
	PubTime        string                      `json:"pub_time"`
	PubTs          json.Number                 `json:"pub_ts"`
	Type           string                      `json:"type"`
	Vip            DynamicAuthorVip            `json:"vip"`
}

// DynamicDecorate 是原始动态作者的装扮卡片信息。
type DynamicDecorate struct {
	CardURL string             `json:"card_url"`
	Fan     DynamicDecorateFan `json:"fan"`
	ID      json.Number        `json:"id"`
	JumpURL string             `json:"jump_url"`
	Name    string             `json:"name"`
	Type    int                `json:"type"`
}

// DynamicDecorateFan 是装扮卡片的粉丝专属信息。
type DynamicDecorateFan struct {
	Color  string      `json:"color"`
	IsFan  bool        `json:"is_fan"`
	NumStr string      `json:"num_str"`
	Number json.Number `json:"number"`
}

// DynamicOriginalPendant 是原始动态作者的挂件信息；pid 为整型，与外层不同。
type DynamicOriginalPendant struct {
	Expire            json.Number `json:"expire"`
	Image             string      `json:"image"`
	ImageEnhance      string      `json:"image_enhance"`
	ImageEnhanceFrame string      `json:"image_enhance_frame"`
	NPid              json.Number `json:"n_pid"`
	Name              string      `json:"name"`
	Pid               int         `json:"pid"`
}

// DynamicOriginalAuthorAvatar contains avatar rendering layers from the original item schema.
type DynamicOriginalAuthorAvatar struct {
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

// DynamicOriginalModuleDynamic contains the original content; Major retains its value semantics.
type DynamicOriginalModuleDynamic struct {
	Additional any                         `json:"additional"`
	Desc       *DynamicOriginalDescription `json:"desc"`
	Major      DynamicOriginalMajor        `json:"major"`
	Topic      any                         `json:"topic"`
}

// DynamicOriginalDescription contains original rich text nodes, which differ from outer item nodes.
type DynamicOriginalDescription struct {
	RichTextNodes []DynamicOriginalRichTextNode `json:"rich_text_nodes"`
	Text          string                        `json:"text"`
}

// DynamicOriginalRichTextNode is one rich text node of an original item description.
type DynamicOriginalRichTextNode struct {
	JumpURL  string       `json:"jump_url,omitempty"`
	OrigText string       `json:"orig_text"`
	Text     string       `json:"text"`
	Type     string       `json:"type"`
	Emoji    DynamicEmoji `json:"emoji,omitempty"`
}

// DynamicOriginalMajor contains the original media body, retaining its field order.
// DynamicOriginalMajor 是 DynamicMajor 的别名（原动态的 major 与整体动态字段完全一致）。
type DynamicOriginalMajor = DynamicMajor
