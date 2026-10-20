package user

import "encoding/json"

// 用户粉丝明细与关注分组的响应模型。
// 各接口的返回体独立定义，即使字段相同也不复用，便于后续接口变更互不影响。

// GetUserFansResult 查询用户粉丝明细（新）的响应。
type GetUserFansResult struct {
	List      []GetUserFansItem `json:"list"`       // 粉丝明细列表
	Offset    string            `json:"offset"`     // 偏移量。普通：{列表最后的粉丝的 mtime}:{mid}；智能推荐：rcmd
	ReVersion json.Number       `json:"re_version"` // （？）
	Total     int               `json:"total"`      // 粉丝总数
}

// GetUserFansItem 是查询用户粉丝明细（新）列表中的一名粉丝。
type GetUserFansItem struct {
	Mid            int                           `json:"mid"`             // 用户 mid
	Attribute      int                           `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                           `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                         `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                           `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetUserFansItemContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                        `json:"uname"`           // 用户昵称
	Face           string                        `json:"face"`            // 用户头像url
	FaceNft        int                           `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                        `json:"sign"`            // 用户签名
	OfficialVerify GetUserFansItemOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetUserFansItemVip            `json:"vip"`             // 会员信息
	NftIcon        string                        `json:"nft_icon"`        // （？）
	RecReason      string                        `json:"rec_reason"`      // （？）
	TrackID        string                        `json:"track_id"`        // （？）
}

// GetUserFansItemContractInfo 是粉丝明细中的契约计划信息。
type GetUserFansItemContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetUserFansItemOfficialVerify 是粉丝明细中的认证信息。
type GetUserFansItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserFansItemVip 是粉丝明细中的会员信息。
type GetUserFansItemVip struct {
	VipType            int                     `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                     `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                  `json:"dueRemark"`            // (?)
	AccessStatus       int                     `json:"accessStatus"`         // (?)
	VipStatus          int                     `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                  `json:"vipStatusWarn"`        // (?)
	ThemeType          int                     `json:"theme_type"`           // 会员样式 id
	Label              GetUserFansItemVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                     `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                  `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                  `json:"nickname_color"`       // 昵称颜色
}

// GetUserFansItemVipLabel 是粉丝明细中的会员铭牌样式。
type GetUserFansItemVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetUserFollowers3Result 查询用户粉丝明细（biligame 游戏中心 H5）的响应。
type GetUserFollowers3Result struct {
	List []GetUserFollowers3Item `json:"list"` // 粉丝明细列表
}

// GetUserFollowers3Item 是 biligame 游戏中心 H5 粉丝列表中的成员。
type GetUserFollowers3Item struct {
	Mid                string                                  `json:"mid"`                 // 用户 mid。该接口以字符串返回
	Attribute          int                                     `json:"attribute"`           // 关系属性
	Uname              string                                  `json:"uname"`               // 用户昵称
	Face               string                                  `json:"face"`                // 用户头像 url
	AttestationDisplay GetUserFollowers3ItemAttestationDisplay `json:"attestation_display"` // 认证信息
}

// GetUserFollowers3ItemAttestationDisplay 是 biligame 粉丝成员中的认证信息。
type GetUserFollowers3ItemAttestationDisplay struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserFollowersUnreadCountResult 获取自己粉丝列表未读状态的响应。
type GetUserFollowersUnreadCountResult struct {
	Count int   `json:"count"` // 自上次访问后新增的粉丝个数
	Time  int64 `json:"time"`  // 上次访问粉丝列表时间。秒级时间戳；没有新增粉丝时为 0
}

// GetRelationTagUsersItem 是关注分组明细中的成员。
type GetRelationTagUsersItem struct {
	Mid            int                                   `json:"mid"`             // 用户 mid
	Attribute      int                                   `json:"attribute"`       // 关系属性。恒为 0
	Tag            any                                   `json:"tag"`             // 关注分组 id。恒为 null
	Special        int                                   `json:"special"`         // 是否特别关注。恒为 0
	ContractInfo   GetRelationTagUsersItemContractInfo   `json:"contract_info"`   // 契约计划信息
	Uname          string                                `json:"uname"`           // 用户昵称
	Face           string                                `json:"face"`            // 用户头像 url
	Sign           string                                `json:"sign"`            // 用户签名
	FaceNft        int                                   `json:"face_nft"`        // 是否为 NFT 头像
	OfficialVerify GetRelationTagUsersItemOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetRelationTagUsersItemVip            `json:"vip"`             // 会员信息
	Live           GetRelationTagUsersItemLive           `json:"live"`            // 直播状态
	NftIcon        string                                `json:"nft_icon"`        // （？）
	RecReason      string                                `json:"rec_reason"`      // （？）
	TrackID        string                                `json:"track_id"`        // （？）
	FollowTime     string                                `json:"follow_time"`     // （？）
}

// GetRelationTagUsersItemContractInfo 是关注分组成员的契约计划信息。
type GetRelationTagUsersItemContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetRelationTagUsersItemOfficialVerify 是关注分组成员的认证信息。
type GetRelationTagUsersItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetRelationTagUsersItemVip 是关注分组成员的会员信息。
type GetRelationTagUsersItemVip struct {
	VipType            int                             `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                             `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                          `json:"dueRemark"`            // (?)
	AccessStatus       int                             `json:"accessStatus"`         // (?)
	VipStatus          int                             `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                          `json:"vipStatusWarn"`        // (?)
	ThemeType          int                             `json:"theme_type"`           // 会员样式 id
	Label              GetRelationTagUsersItemVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                             `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                          `json:"nickname_color"`       // 昵称颜色
}

// GetRelationTagUsersItemVipLabel 是关注分组成员的会员铭牌样式。
type GetRelationTagUsersItemVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetRelationTagUsersItemLive 是关注分组成员的直播状态。
type GetRelationTagUsersItemLive struct {
	LiveStatus int    `json:"live_status"` // 直播状态。0：未开播。1：已直播
	JumpURL    string `json:"jump_url"`    // 直播链接
}

// CreateRelationTagResult 创建分组的响应。
type CreateRelationTagResult struct {
	Tagid int `json:"tagid"` // 创建的分组的 id
}
