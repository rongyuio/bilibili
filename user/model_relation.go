package user

import "encoding/json"

// 用户关注 / 粉丝 / 黑名单等关系相关响应模型。
//
// 各接口的列表项分别定义，即使字段集相同也不共用类型。

// GetUserFollowersResult 是查询用户粉丝明细的 data 对象。
type GetUserFollowersResult struct {
	List      []GetUserFollowersItem `json:"list"`       // 明细列表
	ReVersion json.Number            `json:"re_version"` // （？）（可能是number，可能是string）
	Total     int                    `json:"total"`      // 列表总数
}

// GetUserFollowersItem 是粉丝明细列表中的一项。
type GetUserFollowersItem struct {
	Mid            int                            `json:"mid"`             // 用户 mid
	Attribute      int                            `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                            `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                          `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                            `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetUserFollowersContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                         `json:"uname"`           // 用户昵称
	Face           string                         `json:"face"`            // 用户头像url
	FaceNft        int                            `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                         `json:"sign"`            // 用户签名
	OfficialVerify GetUserFollowersOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetUserFollowersVip            `json:"vip"`             // 会员信息
	NftIcon        string                         `json:"nft_icon"`        // （？）
	RecReason      string                         `json:"rec_reason"`      // （？）
	TrackID        string                         `json:"track_id"`        // （？）
}

// GetUserFollowersContractInfo 是粉丝明细中的契约计划信息。
type GetUserFollowersContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetUserFollowersOfficialVerify 是粉丝明细中的认证信息。
type GetUserFollowersOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserFollowersVip 是粉丝明细中的会员信息。
type GetUserFollowersVip struct {
	VipType            int                      `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                      `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                   `json:"dueRemark"`            // (?)
	AccessStatus       int                      `json:"accessStatus"`         // (?)
	VipStatus          int                      `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                   `json:"vipStatusWarn"`        // (?)
	ThemeType          int                      `json:"theme_type"`           // 会员样式 id
	Label              GetUserFollowersVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                      `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                   `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                   `json:"nickname_color"`       // 昵称颜色
}

// GetUserFollowersVipLabel 是粉丝明细中的会员铭牌样式。
type GetUserFollowersVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetUserFollowingsResult 是查询用户关注明细的 data 对象。
type GetUserFollowingsResult struct {
	List      []GetUserFollowingsItem `json:"list"`       // 明细列表
	ReVersion json.Number             `json:"re_version"` // （？）（可能是number，可能是string）
	Total     int                     `json:"total"`      // 列表总数
}

// GetUserFollowingsItem 是关注明细列表中的一项。
type GetUserFollowingsItem struct {
	Mid            int                             `json:"mid"`             // 用户 mid
	Attribute      int                             `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                             `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                           `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                             `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetUserFollowingsContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                          `json:"uname"`           // 用户昵称
	Face           string                          `json:"face"`            // 用户头像url
	FaceNft        int                             `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                          `json:"sign"`            // 用户签名
	OfficialVerify GetUserFollowingsOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetUserFollowingsVip            `json:"vip"`             // 会员信息
	NftIcon        string                          `json:"nft_icon"`        // （？）
	RecReason      string                          `json:"rec_reason"`      // （？）
	TrackID        string                          `json:"track_id"`        // （？）
}

// GetUserFollowingsContractInfo 是关注明细中的契约计划信息。
type GetUserFollowingsContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetUserFollowingsOfficialVerify 是关注明细中的认证信息。
type GetUserFollowingsOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserFollowingsVip 是关注明细中的会员信息。
type GetUserFollowingsVip struct {
	VipType            int                       `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                       `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                    `json:"dueRemark"`            // (?)
	AccessStatus       int                       `json:"accessStatus"`         // (?)
	VipStatus          int                       `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                    `json:"vipStatusWarn"`        // (?)
	ThemeType          int                       `json:"theme_type"`           // 会员样式 id
	Label              GetUserFollowingsVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                       `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                    `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                    `json:"nickname_color"`       // 昵称颜色
}

// GetUserFollowingsVipLabel 是关注明细中的会员铭牌样式。
type GetUserFollowingsVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetUserFollowings3Result 是查询用户关注明细3的 data 对象。
type GetUserFollowings3Result struct {
	List []GetUserFollowings3Item `json:"list"` // 明细列表
}

// GetUserFollowings3Item 是关注明细3列表中的一项。
type GetUserFollowings3Item struct {
	Mid       string `json:"mid"`       // 用户mid
	Attribute int    `json:"attribute"` // 关注属性。0：未关注。2：已关注。6：已互粉
	Uname     string `json:"uname"`     // 用户昵称
	Face      string `json:"face"`      // 用户头像url
}

// SearchUserFollowingsResult 是搜索关注明细的 data 对象。
type SearchUserFollowingsResult struct {
	List      []SearchUserFollowingsItem `json:"list"`       // 明细列表
	ReVersion json.Number                `json:"re_version"` // （？）（可能是number，可能是string）
	Total     int                        `json:"total"`      // 列表总数
}

// SearchUserFollowingsItem 是搜索关注明细列表中的一项。
type SearchUserFollowingsItem struct {
	Mid            int                                `json:"mid"`             // 用户 mid
	Attribute      int                                `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                                `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                              `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                                `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   SearchUserFollowingsContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                             `json:"uname"`           // 用户昵称
	Face           string                             `json:"face"`            // 用户头像url
	FaceNft        int                                `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                             `json:"sign"`            // 用户签名
	OfficialVerify SearchUserFollowingsOfficialVerify `json:"official_verify"` // 认证信息
	Vip            SearchUserFollowingsVip            `json:"vip"`             // 会员信息
	NftIcon        string                             `json:"nft_icon"`        // （？）
	RecReason      string                             `json:"rec_reason"`      // （？）
	TrackID        string                             `json:"track_id"`        // （？）
}

// SearchUserFollowingsContractInfo 是搜索关注明细中的契约计划信息。
type SearchUserFollowingsContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// SearchUserFollowingsOfficialVerify 是搜索关注明细中的认证信息。
type SearchUserFollowingsOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// SearchUserFollowingsVip 是搜索关注明细中的会员信息。
type SearchUserFollowingsVip struct {
	VipType            int                          `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                          `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                       `json:"dueRemark"`            // (?)
	AccessStatus       int                          `json:"accessStatus"`         // (?)
	VipStatus          int                          `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                       `json:"vipStatusWarn"`        // (?)
	ThemeType          int                          `json:"theme_type"`           // 会员样式 id
	Label              SearchUserFollowingsVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                          `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                       `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                       `json:"nickname_color"`       // 昵称颜色
}

// SearchUserFollowingsVipLabel 是搜索关注明细中的会员铭牌样式。
type SearchUserFollowingsVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetSameFollowingsResult 是查询共同关注明细的 data 对象。
type GetSameFollowingsResult struct {
	List      []GetSameFollowingsItem `json:"list"`       // 明细列表
	ReVersion json.Number             `json:"re_version"` // （？）（可能是number，可能是string）
	Total     int                     `json:"total"`      // 列表总数
}

// GetSameFollowingsItem 是共同关注明细列表中的一项。
type GetSameFollowingsItem struct {
	Mid            int                             `json:"mid"`             // 用户 mid
	Attribute      int                             `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                             `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                           `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                             `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetSameFollowingsContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                          `json:"uname"`           // 用户昵称
	Face           string                          `json:"face"`            // 用户头像url
	FaceNft        int                             `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                          `json:"sign"`            // 用户签名
	OfficialVerify GetSameFollowingsOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetSameFollowingsVip            `json:"vip"`             // 会员信息
	NftIcon        string                          `json:"nft_icon"`        // （？）
	RecReason      string                          `json:"rec_reason"`      // （？）
	TrackID        string                          `json:"track_id"`        // （？）
}

// GetSameFollowingsContractInfo 是共同关注明细中的契约计划信息。
type GetSameFollowingsContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetSameFollowingsOfficialVerify 是共同关注明细中的认证信息。
type GetSameFollowingsOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetSameFollowingsVip 是共同关注明细中的会员信息。
type GetSameFollowingsVip struct {
	VipType            int                       `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                       `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                    `json:"dueRemark"`            // (?)
	AccessStatus       int                       `json:"accessStatus"`         // (?)
	VipStatus          int                       `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                    `json:"vipStatusWarn"`        // (?)
	ThemeType          int                       `json:"theme_type"`           // 会员样式 id
	Label              GetSameFollowingsVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                       `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                    `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                    `json:"nickname_color"`       // 昵称颜色
}

// GetSameFollowingsVipLabel 是共同关注明细中的会员铭牌样式。
type GetSameFollowingsVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetWhispersResult 是查询悄悄关注明细的 data 对象。
type GetWhispersResult struct {
	List      []GetWhispersItem `json:"list"`       // 明细列表
	ReVersion json.Number       `json:"re_version"` // （？）（可能是number，可能是string）
}

// GetWhispersItem 是悄悄关注明细列表中的一项。
type GetWhispersItem struct {
	Mid            int                       `json:"mid"`             // 用户 mid
	Attribute      int                       `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                       `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                     `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                       `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetWhispersContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                    `json:"uname"`           // 用户昵称
	Face           string                    `json:"face"`            // 用户头像url
	FaceNft        int                       `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                    `json:"sign"`            // 用户签名
	OfficialVerify GetWhispersOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetWhispersVip            `json:"vip"`             // 会员信息
	NftIcon        string                    `json:"nft_icon"`        // （？）
	RecReason      string                    `json:"rec_reason"`      // （？）
	TrackID        string                    `json:"track_id"`        // （？）
}

// GetWhispersContractInfo 是悄悄关注明细中的契约计划信息。
type GetWhispersContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetWhispersOfficialVerify 是悄悄关注明细中的认证信息。
type GetWhispersOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetWhispersVip 是悄悄关注明细中的会员信息。
type GetWhispersVip struct {
	VipType            int                 `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                 `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string              `json:"dueRemark"`            // (?)
	AccessStatus       int                 `json:"accessStatus"`         // (?)
	VipStatus          int                 `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string              `json:"vipStatusWarn"`        // (?)
	ThemeType          int                 `json:"theme_type"`           // 会员样式 id
	Label              GetWhispersVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                 `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string              `json:"avatar_subscript_url"` // (?)
	NicknameColor      string              `json:"nickname_color"`       // 昵称颜色
}

// GetWhispersVipLabel 是悄悄关注明细中的会员铭牌样式。
type GetWhispersVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetFriendsResult 是查询互相关注明细的 data 对象。
type GetFriendsResult struct {
	List      []GetFriendsItem `json:"list"`       // 明细列表
	ReVersion json.Number      `json:"re_version"` // （？）（可能是number，可能是string）
}

// GetFriendsItem 是互相关注明细列表中的一项。
type GetFriendsItem struct {
	Mid            int                      `json:"mid"`             // 用户 mid
	Attribute      int                      `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                      `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                    `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                      `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetFriendsContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                   `json:"uname"`           // 用户昵称
	Face           string                   `json:"face"`            // 用户头像url
	FaceNft        int                      `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                   `json:"sign"`            // 用户签名
	OfficialVerify GetFriendsOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetFriendsVip            `json:"vip"`             // 会员信息
	NftIcon        string                   `json:"nft_icon"`        // （？）
	RecReason      string                   `json:"rec_reason"`      // （？）
	TrackID        string                   `json:"track_id"`        // （？）
}

// GetFriendsContractInfo 是互相关注明细中的契约计划信息。
type GetFriendsContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetFriendsOfficialVerify 是互相关注明细中的认证信息。
type GetFriendsOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetFriendsVip 是互相关注明细中的会员信息。
type GetFriendsVip struct {
	VipType            int                `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string             `json:"dueRemark"`            // (?)
	AccessStatus       int                `json:"accessStatus"`         // (?)
	VipStatus          int                `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string             `json:"vipStatusWarn"`        // (?)
	ThemeType          int                `json:"theme_type"`           // 会员样式 id
	Label              GetFriendsVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string             `json:"avatar_subscript_url"` // (?)
	NicknameColor      string             `json:"nickname_color"`       // 昵称颜色
}

// GetFriendsVipLabel 是互相关注明细中的会员铭牌样式。
type GetFriendsVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetBlacksResult 是查询黑名单明细的 data 对象。
type GetBlacksResult struct {
	List      []GetBlacksItem `json:"list"`       // 明细列表
	ReVersion json.Number     `json:"re_version"` // （？）（可能是number，可能是string）
}

// GetBlacksItem 是黑名单明细列表中的一项。
type GetBlacksItem struct {
	Mid            int                     `json:"mid"`             // 用户 mid
	Attribute      int                     `json:"attribute"`       // 对方对于自己的关系属性。0：未关注。1：悄悄关注（现已下线）。2：已关注。6：已互粉。128：已拉黑
	Mtime          int                     `json:"mtime"`           // 对方关注目标用户时间。秒级时间戳。互关后刷新
	Tag            []int                   `json:"tag"`             // 目标用户将对方分组到的 id
	Special        int                     `json:"special"`         // 目标用户特别关注对方标识。0：否。1：是
	ContractInfo   GetBlacksContractInfo   `json:"contract_info"`   // 契约计划相关信息
	Uname          string                  `json:"uname"`           // 用户昵称
	Face           string                  `json:"face"`            // 用户头像url
	FaceNft        int                     `json:"face_nft"`        // 是否为 NFT 头像。0：非 NFT 头像。1：NFT 头像
	Sign           string                  `json:"sign"`            // 用户签名
	OfficialVerify GetBlacksOfficialVerify `json:"official_verify"` // 认证信息
	Vip            GetBlacksVip            `json:"vip"`             // 会员信息
	NftIcon        string                  `json:"nft_icon"`        // （？）
	RecReason      string                  `json:"rec_reason"`      // （？）
	TrackID        string                  `json:"track_id"`        // （？）
}

// GetBlacksContractInfo 是黑名单明细中的契约计划信息。
type GetBlacksContractInfo struct {
	IsContract   bool `json:"is_contract"`   // 目标用户是否为对方的契约者
	IsContractor bool `json:"is_contractor"` // 对方是否为目标用户的契约者
	Ts           int  `json:"ts"`            // 对方成为目标用户的契约者的时间。秒级时间戳，仅当 is_contractor 项的值为 true 时才有此项
	UserAttr     int  `json:"user_attr"`     // 对方作为目标用户的契约者的属性。1：老粉。否则为原始粉丝。仅当有特殊属性时才有此项
}

// GetBlacksOfficialVerify 是黑名单明细中的认证信息。
type GetBlacksOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetBlacksVip 是黑名单明细中的会员信息。
type GetBlacksVip struct {
	VipType            int               `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int               `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string            `json:"dueRemark"`            // (?)
	AccessStatus       int               `json:"accessStatus"`         // (?)
	VipStatus          int               `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string            `json:"vipStatusWarn"`        // (?)
	ThemeType          int               `json:"theme_type"`           // 会员样式 id
	Label              GetBlacksVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int               `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string            `json:"avatar_subscript_url"` // (?)
	NicknameColor      string            `json:"nickname_color"`       // 昵称颜色
}

// GetBlacksVipLabel 是黑名单明细中的会员铭牌样式。
type GetBlacksVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// BatchModifyRelationResult 是批量操作用户关系的 data 对象。
type BatchModifyRelationResult struct {
	FailedFids []int `json:"failed_fids"` // 操作失败的 mid 列表
}

// GetUserRelationResult 是查询用户与自己关系（仅关注）的 data 对象。
type GetUserRelationResult struct {
	Mid       int   `json:"mid"`       // 目标用户 mid
	Attribute int   `json:"attribute"` // 关系属性。0：未关注。2：已关注。6：已互粉。128：已拉黑
	Mtime     int   `json:"mtime"`     // 关注对方时间。时间戳。未关注为 0
	Tag       []int `json:"tag"`       // 分组 id
	Special   int   `json:"special"`   // 特别关注标志。0：否。1：是
}

// GetUserRelation2Result 是查询用户与自己关系（互相关系）的 data 对象。
type GetUserRelation2Result struct {
	Relation   GetUserRelation2Relation `json:"relation"`    // 目标用户对于当前用户的关系
	BeRelation GetUserRelation2Relation `json:"be_relation"` // 当前用户对于目标用户的关系
}

// GetUserRelation2Relation 是互相关系接口中的关系对象。
type GetUserRelation2Relation struct {
	Mid       int   `json:"mid"`       // 目标用户 mid
	Attribute int   `json:"attribute"` // 关系属性。0：未关注。2：已关注。6：已互粉。128：已拉黑
	Mtime     int   `json:"mtime"`     // 关注对方时间。时间戳。未关注为 0
	Tag       []int `json:"tag"`       // 分组 id
	Special   int   `json:"special"`   // 特别关注标志。0：否。1：是
}

// BatchGetUserRelationDetail 是批量查询用户与自己关系 data 对象中的关系对象。
type BatchGetUserRelationDetail struct {
	Mid       int   `json:"mid"`       // 目标用户 mid
	Attribute int   `json:"attribute"` // 关系属性。0：未关注。2：已关注。6：已互粉。128：已拉黑
	Mtime     int   `json:"mtime"`     // 关注对方时间。时间戳。未关注为 0
	Tag       []int `json:"tag"`       // 分组 id
	Special   int   `json:"special"`   // 特别关注标志。0：否。1：是
}

// GetRelationTagsItem 是关注分组列表中的一项。
type GetRelationTagsItem struct {
	Tagid int    `json:"tagid"` // 分组 id。-10：特别关注。0：默认分组
	Name  string `json:"name"`  // 分组名称
	Count int    `json:"count"` // 分组成员数
	Tip   string `json:"tip"`   // 提示信息
}
