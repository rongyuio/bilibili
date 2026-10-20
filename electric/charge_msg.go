package electric

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 充电留言相关接口。各接口的返回体独立定义，互不复用。
// 文档见 electric/charge_msg.md。

// GetElectricChargeRemarkListParam 是查询我收到的充电留言的请求参数。
type GetElectricChargeRemarkListParam struct {
	Begin string `json:"begin,omitempty" request:"query,omitempty"` // 起始日期。默认2016-01-01
	End   string `json:"end,omitempty" request:"query,omitempty"`   // 结束日期。默认2050-01-01
	Pn    string `json:"pn,omitempty" request:"query,omitempty"`    // 页数
	Ps    string `json:"ps,omitempty" request:"query,omitempty"`    // 分页大小。取值范围[1,12]
}

// GetElectricChargeRemarkListItem 是充电留言列表 list 数组中的对象。
type GetElectricChargeRemarkListItem struct {
	Aid         int    `json:"aid"`          // 稿件avid
	Bvid        string `json:"bvid"`         // 稿件bvid
	ID          int    `json:"id"`           // 留言记录id
	Mid         int    `json:"mid"`          // 留言者mid（充电用户）。列表接口中为0
	ReplyMid    int    `json:"reply_mid"`    // UP主mid。列表接口中为0
	ElecNum     int    `json:"elec_num"`     // 充电电池数。列表接口中为0
	State       int    `json:"state"`        // UP是否已经回复这条留言。0：未回复，1：已回复
	Msg         string `json:"msg"`          // 留言内容
	Aname       string `json:"aname"`        // 空
	Uname       string `json:"uname"`        // 留言者用户名。列表接口中为空
	Avator      string `json:"avator"`       // 留言者头像。列表接口中为空
	ReplyName   string `json:"reply_name"`   // UP主用户名。列表接口中为空
	ReplyAvator string `json:"reply_avator"` // UP主头像。列表接口中为空
	ReplyMsg    string `json:"reply_msg"`    // 回复内容。列表接口中为空
	Ctime       int64  `json:"ctime"`        // 留言时间。毫秒级时间戳
	ReplyTime   int64  `json:"reply_time"`   // 回复时间。毫秒级时间戳
}

// GetElectricChargeRemarkListPager 是收到的充电留言列表的分页信息。
type GetElectricChargeRemarkListPager struct {
	Current int `json:"current"` // 当前页数
	Size    int `json:"size"`    // 当前分页大小
	Total   int `json:"total"`   // 记录总数
}

// GetElectricChargeRemarkListResult 是查询我收到的充电留言的 data 对象。
type GetElectricChargeRemarkListResult struct {
	List  []GetElectricChargeRemarkListItem `json:"list"`  // 留言列表
	Pager GetElectricChargeRemarkListPager  `json:"pager"` // 分页信息
}

// GetElectricChargeRemarkList 查询我收到的充电留言。
func (c *Client) GetElectricChargeRemarkList(ctx context.Context, param GetElectricChargeRemarkListParam) (*GetElectricChargeRemarkListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/elec/remark/list"
	)
	return bilibili.Execute[*GetElectricChargeRemarkListResult](ctx, c.Client, method, url, param)
}

// GetElectricChargeRemarkDetailParam 是查询充电留言详情的请求参数。
type GetElectricChargeRemarkDetailParam struct {
	ID int `json:"id"` // 留言id
}

// GetElectricChargeRemarkDetailResult 是查询充电留言详情的 data 对象。
type GetElectricChargeRemarkDetailResult struct {
	Aid         int    `json:"aid"`          // 稿件avid
	Bvid        string `json:"bvid"`         // 稿件bvid
	ID          int    `json:"id"`           // 留言id
	Mid         int    `json:"mid"`          // 留言者mid（充电用户）
	ReplyMid    int    `json:"reply_mid"`    // UP主mid
	ElecNum     int    `json:"elec_num"`     // 充电电池数
	State       int    `json:"state"`        // UP是否已经回复这条留言。0：未回复，1：已回复
	Msg         string `json:"msg"`          // 留言内容
	Aname       string `json:"aname"`        // 空
	Uname       string `json:"uname"`        // 留言者用户名
	Avator      string `json:"avator"`       // 留言者头像
	ReplyName   string `json:"reply_name"`   // UP主用户名
	ReplyAvator string `json:"reply_avator"` // UP主头像
	ReplyMsg    string `json:"reply_msg"`    // 回复内容
	Ctime       int64  `json:"ctime"`        // 留言时间。毫秒级时间戳
	ReplyTime   int64  `json:"reply_time"`   // 回复时间。毫秒级时间戳
}

// GetElectricChargeRemarkDetail 查询充电留言详情。
func (c *Client) GetElectricChargeRemarkDetail(ctx context.Context, param GetElectricChargeRemarkDetailParam) (*GetElectricChargeRemarkDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/elec/remark/detail"
	)
	return bilibili.Execute[*GetElectricChargeRemarkDetailResult](ctx, c.Client, method, url, param)
}
