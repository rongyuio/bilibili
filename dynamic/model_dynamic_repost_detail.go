package dynamic

// GetDynamicRepostDetail 的响应模型。

// GetDynamicRepostDetailResult 是动态转发列表的 data 对象。
type GetDynamicRepostDetailResult struct {
	HasMore int                          `json:"has_more"` // 是否还有下一页
	Total   int                          `json:"total"`    // 总计包含
	Items   []GetDynamicRepostDetailItem `json:"items"`
	Gt      int                          `json:"_gt_"` // 固定值0
}

// GetDynamicRepostDetailItem 是转发列表中的一条转发。
type GetDynamicRepostDetailItem struct {
	Desc       GetDynamicRepostDetailDesc    `json:"desc"`
	Card       string                        `json:"card"`
	ExtendJSON string                        `json:"extend_json"`
	Display    GetDynamicRepostDetailDisplay `json:"display"`
}

// GetDynamicRepostDetailDesc 是转发条目的动态元信息。
type GetDynamicRepostDetailDesc struct {
	UID          int                               `json:"uid"`
	Type         int                               `json:"type"`
	Rid          int64                             `json:"rid"`
	Acl          int                               `json:"acl"`
	View         int                               `json:"view"`
	Repost       int                               `json:"repost"`
	Like         int                               `json:"like"`
	IsLiked      int                               `json:"is_liked"`
	DynamicID    int64                             `json:"dynamic_id"`
	Timestamp    int                               `json:"timestamp"`
	PreDyID      int64                             `json:"pre_dy_id"`
	OrigDyID     int64                             `json:"orig_dy_id"`
	OrigType     int                               `json:"orig_type"`
	UserProfile  GetDynamicRepostDetailUserProfile `json:"user_profile"`
	UIDType      int                               `json:"uid_type"`
	Stype        int                               `json:"stype"`
	RType        int                               `json:"r_type"`
	InnerID      int                               `json:"inner_id"`
	Status       int                               `json:"status"`
	DynamicIDStr string                            `json:"dynamic_id_str"`
	PreDyIDStr   string                            `json:"pre_dy_id_str"`
	OrigDyIDStr  string                            `json:"orig_dy_id_str"`
	RidStr       string                            `json:"rid_str"`
	Origin       GetDynamicRepostDetailOrigin      `json:"origin"`
	Previous     GetDynamicRepostDetailPrevious    `json:"previous"`
}

// GetDynamicRepostDetailUserProfile 是转发条目作者的用户信息。
type GetDynamicRepostDetailUserProfile struct {
	Info      GetDynamicRepostDetailUserInfo  `json:"info"`
	Card      GetDynamicRepostDetailCard      `json:"card"`
	Vip       GetDynamicRepostDetailUserVip   `json:"vip"`
	Pendant   GetDynamicRepostDetailPendant   `json:"pendant"`
	Rank      string                          `json:"rank"`
	Sign      string                          `json:"sign"`
	LevelInfo GetDynamicRepostDetailLevelInfo `json:"level_info"`
}

// GetDynamicRepostDetailUserInfo 是转发条目作者的基础信息。
type GetDynamicRepostDetailUserInfo struct {
	UID     int    `json:"uid"`
	Uname   string `json:"uname"`
	Face    string `json:"face"`
	FaceNft int    `json:"face_nft"`
}

// GetDynamicRepostDetailCard 是转发条目作者的卡片信息。
type GetDynamicRepostDetailCard struct {
	OfficialVerify GetDynamicRepostDetailOfficialVerify `json:"official_verify"`
}

// GetDynamicRepostDetailOfficialVerify 是转发条目作者的认证标识。
type GetDynamicRepostDetailOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetDynamicRepostDetailUserVip 是转发条目作者的会员信息（camelCase 变体）。
type GetDynamicRepostDetailUserVip struct {
	VipType            int                            `json:"vipType"`
	VipDueDate         int64                          `json:"vipDueDate"`
	VipStatus          int                            `json:"vipStatus"`
	ThemeType          int                            `json:"themeType"`
	Label              GetDynamicRepostDetailVipLabel `json:"label"`
	AvatarSubscript    int                            `json:"avatar_subscript"`
	NicknameColor      string                         `json:"nickname_color"`
	Role               int                            `json:"role"`
	AvatarSubscriptURL string                         `json:"avatar_subscript_url"`
}

// GetDynamicRepostDetailVipLabel 是转发条目作者的会员铭牌。
type GetDynamicRepostDetailVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetDynamicRepostDetailPendant 是转发条目作者的挂件信息。
type GetDynamicRepostDetailPendant struct {
	Pid               int    `json:"pid"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	Expire            int    `json:"expire"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
}

// GetDynamicRepostDetailLevelInfo 是转发条目作者的等级信息。
type GetDynamicRepostDetailLevelInfo struct {
	CurrentLevel int `json:"current_level"`
}

// GetDynamicRepostDetailOrigin 是转发条目的原动态元信息。
type GetDynamicRepostDetailOrigin struct {
	UID          int    `json:"uid"`
	Type         int    `json:"type"`
	Rid          int    `json:"rid"`
	Acl          int    `json:"acl"`
	View         int    `json:"view"`
	Repost       int    `json:"repost"`
	Like         int    `json:"like"`
	DynamicID    int64  `json:"dynamic_id"`
	Timestamp    int    `json:"timestamp"`
	PreDyID      int    `json:"pre_dy_id"`
	OrigDyID     int    `json:"orig_dy_id"`
	UIDType      int    `json:"uid_type"`
	Stype        int    `json:"stype"`
	RType        int    `json:"r_type"`
	InnerID      int    `json:"inner_id"`
	Status       int    `json:"status"`
	DynamicIDStr string `json:"dynamic_id_str"`
	PreDyIDStr   string `json:"pre_dy_id_str"`
	OrigDyIDStr  string `json:"orig_dy_id_str"`
	RidStr       string `json:"rid_str"`
}

// GetDynamicRepostDetailPrevious 是转发条目的上一条动态元信息。
type GetDynamicRepostDetailPrevious struct {
	UID          int    `json:"uid"`
	Type         int    `json:"type"`
	Rid          int64  `json:"rid"`
	Acl          int    `json:"acl"`
	View         int    `json:"view"`
	Repost       int    `json:"repost"`
	Like         int    `json:"like"`
	DynamicID    int64  `json:"dynamic_id"`
	Timestamp    int    `json:"timestamp"`
	PreDyID      int64  `json:"pre_dy_id"`
	OrigDyID     int64  `json:"orig_dy_id"`
	UIDType      int    `json:"uid_type"`
	Stype        int    `json:"stype"`
	RType        int    `json:"r_type"`
	InnerID      int    `json:"inner_id"`
	Status       int    `json:"status"`
	DynamicIDStr string `json:"dynamic_id_str"`
	PreDyIDStr   string `json:"pre_dy_id_str"`
	OrigDyIDStr  string `json:"orig_dy_id_str"`
	RidStr       string `json:"rid_str"`
}

// GetDynamicRepostDetailDisplay 是转发条目的可操作项。
type GetDynamicRepostDetailDisplay struct {
	Origin   GetDynamicRepostDetailDisplayOrigin `json:"origin"`
	Relation GetDynamicRepostDetailRelation      `json:"relation"`
}

// GetDynamicRepostDetailDisplayOrigin 是转发条目原动态的可操作项。
type GetDynamicRepostDetailDisplayOrigin struct {
	EmojiInfo GetDynamicRepostDetailEmojiInfo `json:"emoji_info"`
	Relation  GetDynamicRepostDetailRelation  `json:"relation"`
}

// GetDynamicRepostDetailRelation 是转发条目的关注关系。
type GetDynamicRepostDetailRelation struct {
	Status     int `json:"status"`
	IsFollow   int `json:"is_follow"`
	IsFollowed int `json:"is_followed"`
}

// GetDynamicRepostDetailEmojiInfo 是转发条目的表情列表。
type GetDynamicRepostDetailEmojiInfo struct {
	EmojiDetails []GetDynamicRepostDetailEmojiDetail `json:"emoji_details"`
}

// GetDynamicRepostDetailEmojiDetail 是转发条目表情列表的一项。
type GetDynamicRepostDetailEmojiDetail struct {
	EmojiName string                          `json:"emoji_name"`
	ID        int                             `json:"id"`
	PackageID int                             `json:"package_id"`
	State     int                             `json:"state"`
	Type      int                             `json:"type"`
	Attr      int                             `json:"attr"`
	Text      string                          `json:"text"`
	URL       string                          `json:"url"`
	Meta      GetDynamicRepostDetailEmojiMeta `json:"meta"`
	Mtime     int                             `json:"mtime"`
}

// GetDynamicRepostDetailEmojiMeta 是转发条目表情的元信息。
type GetDynamicRepostDetailEmojiMeta struct {
	Size int `json:"size"`
}
