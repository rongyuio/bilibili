package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 评论相关接口。响应模型见 comment_model.go。

type GetCommentsDetailParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码，见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/comment/readme.md
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Sort      int    `json:"sort,omitempty" request:"query,omitempty"`       // 排序方式。默认为0。0：按时间。1：按点赞数。2：按回复数
	Nohot     int    `json:"nohot,omitempty" request:"query,omitempty"`      // 是否不显示热评。默认为0。1：不显示。0：显示
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为20。定义域：1-20
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为1
}

// GetCommentsDetail 获取评论区明细
func (c *Client) GetCommentsDetail(ctx context.Context, param GetCommentsDetailParam) (*CommentsDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply"
	)
	return execute[*CommentsDetail](ctx, c, method, url, param)
}

type GetCommentReplyParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码，见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/comment/readme.md
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Root      int    `json:"root"`                                           // 根回复 rpid
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为20。定义域：1-49 。 但 data_replies 的最大内容数为20,因此设置为49其实也只会有20条回复被返回
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为1
}

// GetCommentReply 获取指定评论的回复，按照回复顺序排序
func (c *Client) GetCommentReply(ctx context.Context, param GetCommentReplyParam) (*CommentReply, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply/reply"
	)
	return execute[*CommentReply](ctx, c, method, url, param)
}

type GetCommentsHotReplyParam struct {
	Type int `json:"type"`                                   // 评论区类型代码
	Oid  int `json:"oid"`                                    // 目标评论区 id
	Root int `json:"root"`                                   // 根回复 rpid
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为20。定义域：1-49
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
}

// GetCommentsHotReply 获取评论区热评
func (c *Client) GetCommentsHotReply(ctx context.Context, param GetCommentsHotReplyParam) (*CommentsHotReply, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/reply/hot"
	)
	return execute[*CommentsHotReply](ctx, c, method, url, param)
}
