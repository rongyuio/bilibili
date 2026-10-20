package comment

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 评论区明细（懒加载）、评论对话树、指定评论信息与评论区评论总数接口。
// 响应模型见 comment_list_model.go。

// GetCommentsDetailLazyLoadParam 是获取评论区明细（懒加载）的请求参数。
type GetCommentsDetailLazyLoadParam struct {
	AccessKey     string `json:"access_key,omitempty" request:"query,omitempty"`     // APP 登录 Token
	Type          int    `json:"type"`                                               // 评论区类型代码
	Oid           int    `json:"oid"`                                                // 目标评论区 id
	Mode          int    `json:"mode,omitempty" request:"query,omitempty"`           // 排序方式。默认为 3。0、3：仅按热度。1：按热度+按时间。2：仅按时间
	Next          int    `json:"next,omitempty" request:"query,omitempty"`           // 翻页。不推荐，已弃用，优先级比 pagination_str 高
	PaginationStr string `json:"pagination_str,omitempty" request:"query,omitempty"` // 分页信息。JSON 字符串，取上次响应 data.cursor.pagination_reply.next_offset
	Plat          int    `json:"plat,omitempty" request:"query,omitempty"`           // 平台类型。如 1
	SeekRpid      string `json:"seek_rpid,omitempty" request:"query,omitempty"`      // 空。当获取第一页评论时存在
	WebLocation   string `json:"web_location,omitempty" request:"query,omitempty"`   // 页面标识。如 1315875
}

// GetCommentsDetailLazyLoad 获取评论区明细（懒加载）
func (c *Client) GetCommentsDetailLazyLoad(ctx context.Context, param GetCommentsDetailLazyLoadParam) (*GetCommentsDetailLazyLoadResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply/wbi/main"
	)
	return bilibili.Execute[*GetCommentsDetailLazyLoadResult](ctx, c.Client, method, url, param, c.FillWbi())
}

// GetCommentDialogTreeParam 是获取指定评论对话树的请求参数。
type GetCommentDialogTreeParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Root      int    `json:"root"`                                           // 根回复 rpid
	Dialog    int    `json:"dialog"`                                         // 对话树根 rpid
	Size      int    `json:"size"`                                           // 每页最大项数
	MinFloor  int    `json:"min_floor,omitempty" request:"query,omitempty"`  // (?)
}

// GetCommentDialogTree 获取指定评论对话树，按照对话链排列
func (c *Client) GetCommentDialogTree(ctx context.Context, param GetCommentDialogTreeParam) (*GetCommentDialogTreeResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply/dialog/cursor"
	)
	return bilibili.Execute[*GetCommentDialogTreeResult](ctx, c.Client, method, url, param)
}

// GetCommentInfoParam 是获取指定评论信息的请求参数。
type GetCommentInfoParam struct {
	Type int `json:"type,omitempty" request:"query,omitempty"` // 评论区类型代码
	Oid  int `json:"oid,omitempty" request:"query,omitempty"`  // 目标评论区 id
	Rpid int `json:"rpid"`                                     // 目标评论 rpid
}

// GetCommentInfo 获取指定评论信息。
//
// 该接口已经弃用，仅保留以兼容旧调用。
func (c *Client) GetCommentInfo(ctx context.Context, param GetCommentInfoParam) (*GetCommentInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply/info"
	)
	return bilibili.Execute[*GetCommentInfoResult](ctx, c.Client, method, url, param)
}

// GetCommentCountParam 是获取评论区评论总数的请求参数。
type GetCommentCountParam struct {
	Type int `json:"type"` // 评论区类型代码
	Oid  int `json:"oid"`  // 目标评论区 id
}

// GetCommentCount 获取评论区评论总数
func (c *Client) GetCommentCount(ctx context.Context, param GetCommentCountParam) (*GetCommentCountResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply/count"
	)
	return bilibili.Execute[*GetCommentCountResult](ctx, c.Client, method, url, param)
}
