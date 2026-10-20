package comment

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 评论操作相关接口（发表、点赞、点踩、删除）。响应模型见 comment_action_model.go。

// AddCommentParam 是发表评论的请求参数。
type AddCommentParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Root      int    `json:"root,omitempty" request:"query,omitempty"`       // 根评论 rpid。二级评论以上使用
	Parent    int    `json:"parent,omitempty" request:"query,omitempty"`     // 父评论 rpid。二级评论同根评论 id，大于二级评论为要回复的评论 id
	Message   string `json:"message"`                                        // 发送评论内容。最大 1000 字符，表情使用表情转义符
	Plat      int    `json:"plat,omitempty" request:"query,omitempty"`       // 发送平台标识。1：web 端。2：安卓客户端。3：ios 客户端。4：wp 客户端。默认为 1
}

// AddComment 发表评论。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) AddComment(ctx context.Context, param AddCommentParam) (*AddCommentResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/reply/add"
	)
	return bilibili.Execute[*AddCommentResult](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("access_key", "type", "oid", "root", "parent", "message", "plat"),
		bilibili.FillFormCsrf(c.Client))
}

// LikeCommentParam 是点赞评论的请求参数。
type LikeCommentParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Rpid      int    `json:"rpid"`                                           // 目标评论 rpid
	Action    int    `json:"action,omitempty" request:"query,omitempty"`     // 操作代码。默认为 0。0：取消赞。1：点赞
}

// LikeComment 点赞评论。点赞成功后会同时消去该评论的点踩。
func (c *Client) LikeComment(ctx context.Context, param LikeCommentParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/reply/action"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("access_key", "type", "oid", "rpid", "action"),
		bilibili.FillFormCsrf(c.Client))
	return err
}

// HateCommentParam 是点踩评论的请求参数。
type HateCommentParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Rpid      int    `json:"rpid"`                                           // 目标评论 rpid
	Action    int    `json:"action,omitempty" request:"query,omitempty"`     // 操作代码。默认为 0。0：取消踩。1：点踩
}

// HateComment 点踩评论。点踩成功后会同时消去该评论的点赞。
func (c *Client) HateComment(ctx context.Context, param HateCommentParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/reply/hate"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("access_key", "type", "oid", "rpid", "action"),
		bilibili.FillFormCsrf(c.Client))
	return err
}

// DeleteCommentParam 是删除评论的请求参数。
type DeleteCommentParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token
	Type      int    `json:"type"`                                           // 评论区类型代码
	Oid       int    `json:"oid"`                                            // 目标评论区 id
	Rpid      int    `json:"rpid"`                                           // 目标评论 rpid
}

// DeleteComment 删除评论。只能删除自己的评论，或自己管理的评论区下的评论。
func (c *Client) DeleteComment(ctx context.Context, param DeleteCommentParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/reply/del"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("access_key", "type", "oid", "rpid"),
		bilibili.FillFormCsrf(c.Client))
	return err
}
