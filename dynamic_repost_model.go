package bilibili

// 动态转发列表与转发用户响应模型。

type DynamicGroupItem struct {
	UID                int    `json:"uid"`                  // 用户id
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
	ExtendJSON string               `json:"extend_json"`
	Display    DynamicRepostDisplay `json:"display"`
}

// DynamicRepostDesc 是转发条目的动态元信息。
type DynamicRepostDesc struct {
	UID          int                      `json:"uid"`
	Type         int                      `json:"type"`
	Rid          int64                    `json:"rid"`
	Acl          int                      `json:"acl"`
	View         int                      `json:"view"`
	Repost       int                      `json:"repost"`
	Like         int                      `json:"like"`
	IsLiked      int                      `json:"is_liked"`
	DynamicID    int64                    `json:"dynamic_id"`
	Timestamp    int                      `json:"timestamp"`
	PreDyID      int64                    `json:"pre_dy_id"`
	OrigDyID     int64                    `json:"orig_dy_id"`
	OrigType     int                      `json:"orig_type"`
	UserProfile  DynamicRepostUserProfile `json:"user_profile"`
	UIDType      int                      `json:"uid_type"`
	Stype        int                      `json:"stype"`
	RType        int                      `json:"r_type"`
	InnerID      int                      `json:"inner_id"`
	Status       int                      `json:"status"`
	DynamicIDStr string                   `json:"dynamic_id_str"`
	PreDyIDStr   string                   `json:"pre_dy_id_str"`
	OrigDyIDStr  string                   `json:"orig_dy_id_str"`
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
	UID     int    `json:"uid"`
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
	AvatarSubscriptURL string   `json:"avatar_subscript_url"`
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

// DynamicRepostPrevious 是转发条目的上一条动态元信息。
type DynamicRepostPrevious struct {
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
	ID        int                    `json:"id"`
	PackageID int                    `json:"package_id"`
	State     int                    `json:"state"`
	Type      int                    `json:"type"`
	Attr      int                    `json:"attr"`
	Text      string                 `json:"text"`
	URL       string                 `json:"url"`
	Meta      DynamicRepostEmojiMeta `json:"meta"`
	Mtime     int                    `json:"mtime"`
}

// DynamicRepostEmojiMeta 是转发条目表情的元信息。
type DynamicRepostEmojiMeta struct {
	Size int `json:"size"`
}
