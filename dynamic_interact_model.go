package bilibili

// 动态点赞、直播中关注者与更新 UP 主响应模型。

type DynamicLikeList struct {
	ItemLikes  []DynamicLikeItem `json:"item_likes"`  // 点赞信息列表主体
	HasMore    int               `json:"has_more"`    // 是否还有下一页
	TotalCount int               `json:"total_count"` // 总计点赞数
	Gt         int               `json:"_gt_"`        // 固定值0
}

// DynamicLikeItem 是点赞列表中的一条点赞记录。
type DynamicLikeItem struct {
	UID      int                 `json:"uid"`
	Time     int                 `json:"time"`
	FaceURL  string              `json:"face_url"`
	Uname    string              `json:"uname"`
	UserInfo DynamicLikeUserInfo `json:"user_info"`
	Attend   int                 `json:"attend"`
}

// DynamicLikeUserInfo 是点赞用户的详细信息。
type DynamicLikeUserInfo struct {
	UID            int                  `json:"uid"`
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
	UID   int    `json:"uid"`   // 直播者id
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
	UID   int    `json:"uid"`
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
