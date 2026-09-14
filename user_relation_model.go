package bilibili

import "encoding/json"

// 用户关注 / 粉丝 / 黑名单等关系相关响应模型。

// RelationUserPage 是带总数的关系明细列表（粉丝、关注、共同关注、搜索关注）。
type RelationUserPage struct {
	List      []RelationUser `json:"list"`       // 明细列表
	ReVersion json.Number    `json:"re_version"` // （？）（可能是number，可能是string）
	Total     int            `json:"total"`      // 列表总数
}

// GetUserFollowersResult 是 RelationUserPage 的别名。
type GetUserFollowersResult = RelationUserPage

// GetUserFollowingsResult 是 RelationUserPage 的别名。
type GetUserFollowingsResult = RelationUserPage

type UserFollowingsDetail2 struct {
	Mid            int            `json:"mid"`             // 用户 mid
	Attribute      int            `json:"attribute"`       // 关注属性。0：未关注。2：已关注。6：已互粉
	Mtime          int            `json:"mtime"`           // 关注对方时间。时间戳。互关后刷新
	Tag            []int          `json:"tag"`             // 分组 id
	Special        int            `json:"special"`         // 特别关注标志。0：否。1：是
	Uname          string         `json:"uname"`           // 用户昵称
	Face           string         `json:"face"`            // 用户头像 url
	Sign           string         `json:"sign"`            // 用户签名
	OfficialVerify OfficialVerify `json:"official_verify"` // 认证信息
	Vip            Vip            `json:"vip"`             // 会员信息
	Live           int            `json:"live"`            // 是否直播。0：未直播。1：直播中
}

type GetUserFollowings2Result struct {
	List      []UserFollowingsDetail2 `json:"list"`       // 明细列表
	ReVersion json.Number             `json:"re_version"` // （？）（可能是number，可能是string）
	Total     int                     `json:"total"`      // 关注总数
}

type UserFollowingsDetail3 struct {
	Mid       string `json:"mid"`       // 用户mid
	Attribute int    `json:"attribute"` // 关注属性。0：未关注。2：已关注。6：已互粉
	Uname     string `json:"uname"`     // 用户昵称
	Face      string `json:"face"`      // 用户头像url
}

type GetUserFollowings3Result struct {
	List []UserFollowingsDetail3 `json:"list"` // 明细列表
}

// SearchUserFollowingsResult 是 RelationUserPage 的别名。
type SearchUserFollowingsResult = RelationUserPage

// GetSameFollowingsResult 是 RelationUserPage 的别名。
type GetSameFollowingsResult = RelationUserPage

// RelationUserList 是不带总数的关系名单（悄悄关注、互关、黑名单）。
type RelationUserList struct {
	List      []RelationUser `json:"list"`       // 明细列表
	ReVersion json.Number    `json:"re_version"` // （？）（可能是number，可能是string）
}

// GetWhispersResult 是 RelationUserList 的别名。
type GetWhispersResult = RelationUserList

// GetFriendsResult 是 RelationUserList 的别名。
type GetFriendsResult = RelationUserList

// GetBlacksResult 是 RelationUserList 的别名。
type GetBlacksResult = RelationUserList

type BatchModifyRelationResult struct {
	FailedFids []int `json:"failed_fids"` // 操作失败的 mid 列表
}

type RelationDetail struct {
	Mid       int   `json:"mid"`       // 目标用户 mid
	Attribute int   `json:"attribute"` // 关系属性。0：未关注。2：已关注。6：已互粉。128：已拉黑
	Mtime     int   `json:"mtime"`     // 关注对方时间。时间戳。未关注为 0
	Tag       []int `json:"tag"`       // 分组 id
	Special   int   `json:"special"`   // 特别关注标志。0：否。1：是
}

type GetUserRelation2Result struct {
	Relation   RelationDetail `json:"relation"`    // 目标用户对于当前用户的关系
	BeRelation RelationDetail `json:"be_relation"` // 当前用户对于目标用户的关系
}

type RelationTag struct {
	Tagid int    `json:"tagid"` // 分组 id。-10：特别关注。0：默认分组
	Name  string `json:"name"`  // 分组名称
	Count int    `json:"count"` // 分组成员数
	Tip   string `json:"tip"`   // 提示信息
}
