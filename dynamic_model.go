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
	IdStr   json.Number         `json:"id_str"` // 这个字段，B站返回的数据有时是number，有时是string
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
	CommentIdStr string          `json:"comment_id_str"`
	CommentType  int             `json:"comment_type"`
	LikeIcon     DynamicLikeIcon `json:"like_icon"`
	RidStr       string          `json:"rid_str"`
}

// DynamicLikeIcon is the like icon metadata of a space feed item.
type DynamicLikeIcon struct {
	ActionUrl string      `json:"action_url"`
	EndUrl    string      `json:"end_url"`
	Id        json.Number `json:"id"`
	StartUrl  string      `json:"start_url"`
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
	Following       json.Number         `json:"following"` // 关注状态；调试样本返回数字 2，不能按布尔值处理，完整状态含义待确认。
	JumpUrl         string              `json:"jump_url"`
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
	AvatarSubscriptUrl string      `json:"avatar_subscript_url"`
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
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AvatarLayer struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GeneralCfg struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle struct {
								BorderRadius    string `json:"borderRadius"`
								BackgroundColor string `json:"background-color,omitempty"`
								Border          string `json:"border,omitempty"`
								BoxSizing       string `json:"boxSizing,omitempty"`
							} `json:"web_css_style"`
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
							Url      string `json:"url"`
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
	JumpUrl  string       `json:"jump_url,omitempty"`
	Style    any          `json:"style"`
	Emoji    DynamicEmoji `json:"emoji,omitempty"`
	Rid      string       `json:"rid,omitempty"`
}

// DynamicEmoji is an inline emoji of a rich text node.
type DynamicEmoji struct {
	IconUrl string `json:"icon_url"`
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

// DynamicOriginalItem is the original item embedded in a repost; it is not recursive.
type DynamicOriginalItem struct {
	Basic   DynamicOriginalBasic   `json:"basic"`
	IdStr   json.Number            `json:"id_str"`
	Modules DynamicOriginalModules `json:"modules"`
	Type    string                 `json:"type"`
	Visible bool                   `json:"visible"`
}

// DynamicOriginalBasic contains the original item identifiers, whose types differ from the outer item.
type DynamicOriginalBasic struct {
	CommentIdStr string                  `json:"comment_id_str"`
	CommentType  int                     `json:"comment_type"`
	LikeIcon     DynamicOriginalLikeIcon `json:"like_icon"`
	RidStr       string                  `json:"rid_str"`
}

// DynamicOriginalLikeIcon 是原始动态的点赞图标；id 为整型，与外层不同。
type DynamicOriginalLikeIcon struct {
	ActionUrl string `json:"action_url"`
	EndUrl    string `json:"end_url"`
	Id        int    `json:"id"`
	StartUrl  string `json:"start_url"`
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
	Following      json.Number                 `json:"following"` // 关注状态；调试样本返回数字 1，不能按布尔值处理，完整状态含义待确认。
	JumpUrl        string                      `json:"jump_url"`
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
	CardUrl string             `json:"card_url"`
	Fan     DynamicDecorateFan `json:"fan"`
	Id      json.Number        `json:"id"`
	JumpUrl string             `json:"jump_url"`
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
				IsCritical bool `json:"is_critical,omitempty"`
				Tags       struct {
					AVATARLAYER struct {
					} `json:"AVATAR_LAYER,omitempty"`
					GENERALCFG struct {
						ConfigType    int `json:"config_type"`
						GeneralConfig struct {
							WebCssStyle struct {
								BorderRadius    string `json:"borderRadius"`
								BackgroundColor string `json:"background-color,omitempty"`
								Border          string `json:"border,omitempty"`
								BoxSizing       string `json:"boxSizing,omitempty"`
							} `json:"web_css_style"`
						} `json:"general_config"`
					} `json:"GENERAL_CFG,omitempty"`
					PENDENTLAYER struct {
					} `json:"PENDENT_LAYER,omitempty"`
					ICONLAYER struct {
					} `json:"ICON_LAYER,omitempty"`
				} `json:"tags"`
			} `json:"layer_config"`
			Resource struct {
				ResImage struct {
					ImageSrc struct {
						Placeholder int `json:"placeholder,omitempty"`
						Remote      struct {
							BfsStyle string `json:"bfs_style"`
							Url      string `json:"url"`
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
	JumpUrl  string       `json:"jump_url,omitempty"`
	OrigText string       `json:"orig_text"`
	Text     string       `json:"text"`
	Type     string       `json:"type"`
	Emoji    DynamicEmoji `json:"emoji,omitempty"`
}

// DynamicOriginalMajor contains the original media body, retaining its field order.
type DynamicOriginalMajor struct {
	Archive DynamicArchive `json:"archive,omitempty"`
	Type    string         `json:"type"`
	Draw    DynamicDraw    `json:"draw,omitempty"`
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
	JumpUrl        string              `json:"jump_url"`
	Stat           DynamicArchiveStat  `json:"stat"`
	Title          string              `json:"title"`
	Type           int                 `json:"type"`
}

// DynamicArchiveBadge 是稿件角标信息。
type DynamicArchiveBadge struct {
	BgColor string `json:"bg_color"`
	Color   string `json:"color"`
	IconUrl any    `json:"icon_url"`
	Text    string `json:"text"`
}

// DynamicArchiveStat 是稿件的播放与弹幕数。
type DynamicArchiveStat struct {
	Danmaku string `json:"danmaku"`
	Play    string `json:"play"`
}

// DynamicDraw describes images in either an outer or original media body.
type DynamicDraw struct {
	Id    json.Number       `json:"id"`
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

type DynamicGroupItem struct {
	Uid                int    `json:"uid"`                  // 用户id
	Uname              string `json:"uname"`                // 用户昵称
	Face               string `json:"face"`                 // 用户头像url
	Fans               int    `json:"fans"`                 // 用户粉丝数
	OfficialVerifyType int    `json:"official_verify_type"` // 认证信息?
}

type DynamicGroup struct {
	GroupType int                `json:"group_type"` // 2:我的关注。4:其他
	GroupName string             `json:"group_name"` // 分组名字
	Items     []DynamicGroupItem `json:"items"`      // 用户信息
}

type SearchDynamicAtResult struct {
	Groups []DynamicGroup `json:"groups"` // 内容分组
	Gt     int            `json:"_gt_"`   // 固定值0
}

// DynamicRepostDetail 是动态转发列表。
type DynamicRepostDetail struct {
	HasMore int                 `json:"has_more"` // 是否还有下一页
	Total   int                 `json:"total"`    // 总计包含
	Items   []DynamicRepostItem `json:"items"`
	Gt      int                 `json:"_gt_"` // 固定值0
}

// DynamicRepostItem 是转发列表中的一条转发。
type DynamicRepostItem struct {
	Desc       DynamicRepostDesc    `json:"desc"`
	Card       string               `json:"card"`
	ExtendJson string               `json:"extend_json"`
	Display    DynamicRepostDisplay `json:"display"`
}

// DynamicRepostDesc 是转发条目的动态元信息。
type DynamicRepostDesc struct {
	Uid          int                      `json:"uid"`
	Type         int                      `json:"type"`
	Rid          int64                    `json:"rid"`
	Acl          int                      `json:"acl"`
	View         int                      `json:"view"`
	Repost       int                      `json:"repost"`
	Like         int                      `json:"like"`
	IsLiked      int                      `json:"is_liked"`
	DynamicId    int64                    `json:"dynamic_id"`
	Timestamp    int                      `json:"timestamp"`
	PreDyId      int64                    `json:"pre_dy_id"`
	OrigDyId     int64                    `json:"orig_dy_id"`
	OrigType     int                      `json:"orig_type"`
	UserProfile  DynamicRepostUserProfile `json:"user_profile"`
	UidType      int                      `json:"uid_type"`
	Stype        int                      `json:"stype"`
	RType        int                      `json:"r_type"`
	InnerId      int                      `json:"inner_id"`
	Status       int                      `json:"status"`
	DynamicIdStr string                   `json:"dynamic_id_str"`
	PreDyIdStr   string                   `json:"pre_dy_id_str"`
	OrigDyIdStr  string                   `json:"orig_dy_id_str"`
	RidStr       string                   `json:"rid_str"`
	Origin       DynamicRepostOrigin      `json:"origin"`
	Previous     DynamicRepostPrevious    `json:"previous"`
}

// DynamicRepostUserProfile 是转发条目作者的用户信息。
type DynamicRepostUserProfile struct {
	Info      DynamicRepostUserInfo `json:"info"`
	Card      DynamicRepostCard     `json:"card"`
	Vip       DynamicUserVip        `json:"vip"`
	Pendant   DynamicUserPendant    `json:"pendant"`
	Rank      string                `json:"rank"`
	Sign      string                `json:"sign"`
	LevelInfo DynamicUserLevelInfo  `json:"level_info"`
}

// DynamicRepostUserInfo 是转发条目作者的基础信息。
type DynamicRepostUserInfo struct {
	Uid     int    `json:"uid"`
	Uname   string `json:"uname"`
	Face    string `json:"face"`
	FaceNft int    `json:"face_nft"`
}

// DynamicRepostCard 是转发条目作者的卡片信息。
type DynamicRepostCard struct {
	OfficialVerify OfficialVerify `json:"official_verify"`
}

// DynamicUserVip 是转发与点赞用户的会员信息（camelCase 变体）。
type DynamicUserVip struct {
	VipType            int      `json:"vipType"`
	VipDueDate         int64    `json:"vipDueDate"`
	VipStatus          int      `json:"vipStatus"`
	ThemeType          int      `json:"themeType"`
	Label              VipLabel `json:"label"`
	AvatarSubscript    int      `json:"avatar_subscript"`
	NicknameColor      string   `json:"nickname_color"`
	Role               int      `json:"role"`
	AvatarSubscriptUrl string   `json:"avatar_subscript_url"`
}

// DynamicUserPendant 是转发与点赞用户的挂件信息。
type DynamicUserPendant struct {
	Pid               int    `json:"pid"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	Expire            int    `json:"expire"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
}

// DynamicUserLevelInfo 是转发与点赞用户的等级信息。
type DynamicUserLevelInfo struct {
	CurrentLevel int `json:"current_level"`
}

// DynamicRepostOrigin 是转发条目的原动态元信息。
type DynamicRepostOrigin struct {
	Uid          int    `json:"uid"`
	Type         int    `json:"type"`
	Rid          int    `json:"rid"`
	Acl          int    `json:"acl"`
	View         int    `json:"view"`
	Repost       int    `json:"repost"`
	Like         int    `json:"like"`
	DynamicId    int64  `json:"dynamic_id"`
	Timestamp    int    `json:"timestamp"`
	PreDyId      int    `json:"pre_dy_id"`
	OrigDyId     int    `json:"orig_dy_id"`
	UidType      int    `json:"uid_type"`
	Stype        int    `json:"stype"`
	RType        int    `json:"r_type"`
	InnerId      int    `json:"inner_id"`
	Status       int    `json:"status"`
	DynamicIdStr string `json:"dynamic_id_str"`
	PreDyIdStr   string `json:"pre_dy_id_str"`
	OrigDyIdStr  string `json:"orig_dy_id_str"`
	RidStr       string `json:"rid_str"`
}

// DynamicRepostPrevious 是转发条目的上一条动态元信息。
type DynamicRepostPrevious struct {
	Uid          int    `json:"uid"`
	Type         int    `json:"type"`
	Rid          int64  `json:"rid"`
	Acl          int    `json:"acl"`
	View         int    `json:"view"`
	Repost       int    `json:"repost"`
	Like         int    `json:"like"`
	DynamicId    int64  `json:"dynamic_id"`
	Timestamp    int    `json:"timestamp"`
	PreDyId      int64  `json:"pre_dy_id"`
	OrigDyId     int64  `json:"orig_dy_id"`
	UidType      int    `json:"uid_type"`
	Stype        int    `json:"stype"`
	RType        int    `json:"r_type"`
	InnerId      int    `json:"inner_id"`
	Status       int    `json:"status"`
	DynamicIdStr string `json:"dynamic_id_str"`
	PreDyIdStr   string `json:"pre_dy_id_str"`
	OrigDyIdStr  string `json:"orig_dy_id_str"`
	RidStr       string `json:"rid_str"`
}

// DynamicRepostDisplay 是转发条目的可操作项。
type DynamicRepostDisplay struct {
	Origin   DynamicRepostDisplayOrigin `json:"origin"`
	Relation DynamicRepostRelation      `json:"relation"`
}

// DynamicRepostDisplayOrigin 是转发条目原动态的可操作项。
type DynamicRepostDisplayOrigin struct {
	EmojiInfo DynamicRepostEmojiInfo `json:"emoji_info"`
	Relation  DynamicRepostRelation  `json:"relation"`
}

// DynamicRepostRelation 是转发条目的关注关系。
type DynamicRepostRelation struct {
	Status     int `json:"status"`
	IsFollow   int `json:"is_follow"`
	IsFollowed int `json:"is_followed"`
}

// DynamicRepostEmojiInfo 是转发条目的表情列表。
type DynamicRepostEmojiInfo struct {
	EmojiDetails []DynamicRepostEmojiDetail `json:"emoji_details"`
}

// DynamicRepostEmojiDetail 是转发条目表情列表的一项。
type DynamicRepostEmojiDetail struct {
	EmojiName string                 `json:"emoji_name"`
	Id        int                    `json:"id"`
	PackageId int                    `json:"package_id"`
	State     int                    `json:"state"`
	Type      int                    `json:"type"`
	Attr      int                    `json:"attr"`
	Text      string                 `json:"text"`
	Url       string                 `json:"url"`
	Meta      DynamicRepostEmojiMeta `json:"meta"`
	Mtime     int                    `json:"mtime"`
}

// DynamicRepostEmojiMeta 是转发条目表情的元信息。
type DynamicRepostEmojiMeta struct {
	Size int `json:"size"`
}

type DynamicLikeList struct {
	ItemLikes  []DynamicLikeItem `json:"item_likes"`  // 点赞信息列表主体
	HasMore    int               `json:"has_more"`    // 是否还有下一页
	TotalCount int               `json:"total_count"` // 总计点赞数
	Gt         int               `json:"_gt_"`        // 固定值0
}

// DynamicLikeItem 是点赞列表中的一条点赞记录。
type DynamicLikeItem struct {
	Uid      int                 `json:"uid"`
	Time     int                 `json:"time"`
	FaceUrl  string              `json:"face_url"`
	Uname    string              `json:"uname"`
	UserInfo DynamicLikeUserInfo `json:"user_info"`
	Attend   int                 `json:"attend"`
}

// DynamicLikeUserInfo 是点赞用户的详细信息。
type DynamicLikeUserInfo struct {
	Uid            int                  `json:"uid"`
	Uname          string               `json:"uname"`
	Face           string               `json:"face"`
	Rank           string               `json:"rank"`
	OfficialVerify OfficialVerify       `json:"official_verify"`
	Vip            DynamicUserVip       `json:"vip"`
	Pendant        DynamicUserPendant   `json:"pendant"`
	Sign           string               `json:"sign"`
	LevelInfo      DynamicUserLevelInfo `json:"level_info"`
}

type DynamicLiveUserList struct {
	Count int               `json:"count"` // 直播者数量
	Group string            `json:"group"` // 固定值"default"，作用尚不明确
	Items []DynamicLiveUser `json:"items"` // 直播者列表
	Gt    int               `json:"_gt_"`  // 固定值0，作用尚不明确
}

// DynamicLiveUser 是一个正在直播的已关注者。
type DynamicLiveUser struct {
	Uid   int    `json:"uid"`   // 直播者id
	Uname string `json:"uname"` // 直播者昵称
	Face  string `json:"face"`  // 直播者头像
	Link  string `json:"link"`  // 直播链接
	Title string `json:"title"` // 直播标题
}

type DynamicUpList struct {
	ButtonStatement string              `json:"button_statement"` // 固定值空，作用尚不明确
	Items           []DynamicUpListItem `json:"items"`            // 更新者列表
	Gt              int                 `json:"_gt_"`             // 固定值0，作用尚不明确
}

// DynamicUpListItem 是一条发布新动态的已关注者记录。
type DynamicUpListItem struct {
	UserProfile DynamicUpUserProfile `json:"user_profile"`
	HasUpdate   int                  `json:"has_update"`
}

// DynamicUpUserProfile 是更新者的用户信息。
type DynamicUpUserProfile struct {
	Info      DynamicUpUserInfo  `json:"info"`
	Card      DynamicRepostCard  `json:"card"`
	Vip       DynamicUpVip       `json:"vip"`
	Pendant   DynamicUpPendant   `json:"pendant"`
	Rank      string             `json:"rank"`
	Sign      string             `json:"sign"`
	LevelInfo DynamicUpLevelInfo `json:"level_info"`
}

// DynamicUpUserInfo 是更新者的基础信息。
type DynamicUpUserInfo struct {
	Uid   int    `json:"uid"`
	Uname string `json:"uname"`
	Face  string `json:"face"`
}

// DynamicUpVip 是更新者的会员信息，其标签仅有 path 字段。
type DynamicUpVip struct {
	VipType       int               `json:"vipType"`
	VipDueDate    int64             `json:"vipDueDate"`
	DueRemark     string            `json:"dueRemark"`
	AccessStatus  int               `json:"accessStatus"`
	VipStatus     int               `json:"vipStatus"`
	VipStatusWarn string            `json:"vipStatusWarn"`
	ThemeType     int               `json:"themeType"`
	Label         DynamicUpVipLabel `json:"label"`
}

// DynamicUpVipLabel 是更新者的会员标签（仅含 path）。
type DynamicUpVipLabel struct {
	Path string `json:"path"`
}

// DynamicUpPendant 是更新者的挂件信息。
type DynamicUpPendant struct {
	Pid          int    `json:"pid"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Expire       int    `json:"expire"`
	ImageEnhance string `json:"image_enhance"`
}

// DynamicUpLevelInfo 是更新者的等级信息。
type DynamicUpLevelInfo struct {
	CurrentLevel int    `json:"current_level"`
	CurrentMin   int    `json:"current_min"`
	CurrentExp   int    `json:"current_exp"`
	NextExp      string `json:"next_exp"`
}

// DynamicCard 动态卡片内容。因为 ActivityInfos 、 Desc 、 Display 等字段会随着此动态类型不同发生一定的变化，无法统一，因此都转换成了 map[string]any ，请自行解析
type DynamicCard struct {
	ActivityInfos map[string]any `json:"activity_infos"` // 该条动态参与的活动
	Card          string         `json:"card"`           // 动态详细信息
	Desc          map[string]any `json:"desc"`           // 动态相关信息
	Display       map[string]any `json:"display"`        // 动态部分的可操作项
	ExtendJson    string         `json:"extend_json"`    // 动态扩展项
}

type DynamicDetail struct {
	Card   *DynamicCard `json:"card"` // 动态卡片内容
	Result int          `json:"result"`
	Gt     int          `json:"_gt_"`
}

type DynamicPortal struct {
	MyInfo DynamicPortalMyInfo `json:"my_info"` // 个人关注的一些信息
	UpList []DynamicPortalUp   `json:"up_list"` // 最近更新的up主列表
}

// DynamicPortalMyInfo 是动态门户中的个人关注信息。
type DynamicPortalMyInfo struct {
	Dyns      int                    `json:"dyns"`      // 个人动态
	Face      string                 `json:"face"`      // 头像url
	FaceNft   int                    `json:"face_nft"`  // 含义尚不明确
	Follower  int                    `json:"follower"`  // 粉丝数量
	Following int                    `json:"following"` // 我的关注
	LevelInfo DynamicPortalLevelInfo `json:"level_info"`
	Mid       int                    `json:"mid"`      // 账户mid
	Name      string                 `json:"name"`     // 账户名称
	Official  Official               `json:"official"` // 认证信息
	SpaceBg   string                 `json:"space_bg"` // 账户个人中心的背景横幅url
	Vip       CardVip                `json:"vip"`      // vip信息
}

// DynamicPortalLevelInfo 是本人等级信息。
type DynamicPortalLevelInfo struct {
	CurrentExp   int   `json:"current_exp"`
	CurrentLevel int   `json:"current_level"` // 当前等级，0-6级
	CurrentMin   int   `json:"current_min"`
	LevelUp      int64 `json:"level_up"`
	NextExp      int   `json:"next_exp"`
}

// DynamicPortalUp 是一位最近更新的 UP 主。
type DynamicPortalUp struct {
	Face            string `json:"face"`       // UP主头像
	HasUpdate       bool   `json:"has_update"` // 最近是否有更新
	IsReserveRecall bool   `json:"is_reserve_recall"`
	Mid             int    `json:"mid"`   // UP主mid
	Uname           string `json:"uname"` // UP主昵称
}

// UploadDynamicBfsResult contains the uploaded image url and its pixel size.
type UploadDynamicBfsResult struct {
	ImageUrl    string `json:"image_url"`    // 图片 url
	ImageWidth  int    `json:"image_width"`  // 图片宽度
	ImageHeight int    `json:"image_height"` // 图片高度
}

type CreateDynamicResult struct {
	Result       int    `json:"result"`         // 0
	Errmsg       string `json:"errmsg"`         // 像是服务器日志一样的东西
	DynamicId    int    `json:"dynamic_id"`     // 动态 id
	CreateResult int    `json:"create_result"`  // 1
	DynamicIdStr string `json:"dynamic_id_str"` // 动态 id。字符串格式
	Gt           int    `json:"_gt_"`           // 0
}

// DynamicList 包含置顶及热门的动态列表
//
// TODO: 因为不清楚 attentions 字段（关注列表）的格式，暂未对此字段进行解析
type DynamicList struct {
	Cards         *DynamicCard `json:"cards"` // 动态列表
	FounderUid    int          `json:"founder_uid,omitempty"`
	HasMore       int          `json:"has_more"` // 当前话题是否有额外的动态，0：无额外动态，1：有额外动态
	IsDrawerTopic int          `json:"is_drawer_topic,omitempty"`
	Offset        string       `json:"offset"` // 接下来获取列表时的偏移值，一般为当前获取的话题列表下最后一个动态id
	Gt            int          `json:"_gt_"`   // 固定值0，作用尚不明确
}
