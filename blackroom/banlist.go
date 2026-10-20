package blackroom

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 小黑屋封禁公示相关接口。各接口的返回体独立定义，互不复用。

// GetBlockedUserListParam 是获取封禁用户公示列表的请求参数。
type GetBlockedUserListParam struct {
	Btype int `json:"btype,omitempty" request:"query,omitempty"` // 来源筛选。省略：全部类型。0：系统封禁。1：风纪仲裁（仲裁系统升级后此类型已不再更新）
	Otype int `json:"otype,omitempty" request:"query,omitempty"` // 类型筛选。默认为0：全部类型
	Pn    int `json:"pn,omitempty" request:"query,omitempty"`    // 页码。默认为1
}

// GetBlockedUserListItem 是封禁用户公示列表中的一项。
type GetBlockedUserListItem struct {
	ID                  int    `json:"id"`                  // 封禁公示id
	Uname               string `json:"uname"`               // 对象用户名
	Face                string `json:"face"`                // 对象用户头像url
	UID                 int    `json:"uid"`                 // 对象用户mid
	OriginTitle         string `json:"originTitle"`         // 来源标题
	BlockedRemark       string `json:"blockedRemark"`       // 封禁备注
	OriginURL           string `json:"originUrl"`           // 来源url
	OriginContentModify string `json:"originContentModify"` // 裁决正文
	OriginType          int    `json:"originType"`          // 来源类型
	OriginTypeName      string `json:"originTypeName"`      // 来源名
	PunishTitle         string `json:"punishTitle"`         // 公示标题
	PunishTime          int64  `json:"punishTime"`          // 处罚时间。时间戳
	PunishType          int    `json:"punishType"`          // 处理手段类型。2：封禁。3：永久封禁
	PunishTypeName      string `json:"punishTypeName"`      // 处理手段名
	MoralNum            int    `json:"moralNum"`            // 节操值。被封禁用户节操值均为0
	BlockedDays         int    `json:"blockedDays"`         // 封禁天数。永封为0
	PublishStatus       int    `json:"publishStatus"`       // 1。作用尚不明确
	BlockedType         int    `json:"blockedType"`         // 处理来源。0：系统封禁。1：风纪仲裁
	BlockedForever      int    `json:"blockedForever"`      // 是否永封。0：非永封。1：永封
	ReasonType          int    `json:"reasonType"`          // 封禁原因类型
	ReasonTypeName      string `json:"reasonTypeName"`      // 封禁原因类型名称
	OperatorName        string `json:"operatorName"`        // 空。作用尚不明确
	CaseID              int    `json:"caseId"`              // 仲裁信息id。系统封禁时固定为0
	Ctime               int64  `json:"ctime"`               // 创建时间。时间戳
	CommentSum          int    `json:"commentSum"`          // 该条目评论数
	OriginID            int    `json:"originID"`            // 原内容ID。实测为0
}

// GetBlockedUserList 获取封禁用户公示列表。每页最多10项
func (c *Client) GetBlockedUserList(ctx context.Context, param GetBlockedUserListParam) ([]GetBlockedUserListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/blocked/list"
	)
	return bilibili.Execute[[]GetBlockedUserListItem](ctx, c.Client, method, url, param)
}

// GetBlockedInfoParam 是获取封禁处理公示详情的请求参数。
type GetBlockedInfoParam struct {
	ID int `json:"id"` // 封禁公示id
}

// GetBlockedInfoResult 是封禁处理公示详情的 data 对象。
type GetBlockedInfoResult struct {
	ID                  int    `json:"id"`                  // 封禁公示id
	Uname               string `json:"uname"`               // 对象用户名
	Face                string `json:"face"`                // 对象用户头像url
	UID                 int    `json:"uid"`                 // 对象用户mid
	OriginTitle         string `json:"originTitle"`         // 来源标题
	BlockedRemark       string `json:"blockedRemark"`       // 封禁备注
	OriginURL           string `json:"originUrl"`           // 来源url
	OriginContentModify string `json:"originContentModify"` // 裁决正文
	OriginType          int    `json:"originType"`          // 来源类型
	OriginTypeName      string `json:"originTypeName"`      // 来源名
	PunishTitle         string `json:"punishTitle"`         // 公示标题
	PunishTime          int64  `json:"punishTime"`          // 处罚时间。时间戳
	PunishType          int    `json:"punishType"`          // 处理手段类型。2：封禁。3：永久封禁
	PunishTypeName      string `json:"punishTypeName"`      // 处理手段名
	MoralNum            int    `json:"moralNum"`            // 节操值。被封禁用户节操值均为0
	BlockedDays         int    `json:"blockedDays"`         // 封禁天数。永封为0
	PublishStatus       int    `json:"publishStatus"`       // 1。作用尚不明确
	BlockedType         int    `json:"blockedType"`         // 处理来源。0：系统封禁。1：风纪仲裁
	BlockedForever      int    `json:"blockedForever"`      // 是否永封。0：非永封。1：永封
	ReasonType          int    `json:"reasonType"`          // 封禁原因类型
	ReasonTypeName      string `json:"reasonTypeName"`      // 封禁原因类型名称
	OperatorName        string `json:"operatorName"`        // 空。作用尚不明确
	CaseID              int    `json:"caseId"`              // 仲裁信息id。系统封禁时固定为0
	Ctime               int64  `json:"ctime"`               // 创建时间。时间戳
	CommentSum          int    `json:"commentSum"`          // 该条目评论数
	OriginID            int    `json:"originID"`            // 原内容ID。实测为0
}

// GetBlockedInfo 封禁处理公示详情
func (c *Client) GetBlockedInfo(ctx context.Context, param GetBlockedInfoParam) (*GetBlockedInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/blocked/info"
	)
	return bilibili.Execute[*GetBlockedInfoResult](ctx, c.Client, method, url, param)
}
