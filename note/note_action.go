package note

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频笔记操作相关接口（保存、删除）。各接口的返回体独立定义，互不复用。

// SaveVideoNoteParam 是保存视频笔记的请求参数，所有字段进入 URL 编码表单。
type SaveVideoNoteParam struct {
	Oid         int    `json:"oid"`                                              // 目标 id
	OidType     int    `json:"oid_type"`                                         // 目标 id 类型。0：视频（oid 为 avid）
	NoteID      int64  `json:"note_id,omitempty" request:"query,omitempty"`      // 笔记 id。创建时无需此项
	Title       string `json:"title"`                                            // 笔记标题
	Summary     string `json:"summary"`                                          // 笔记预览文本
	Content     string `json:"content"`                                          // 笔记正文 json 序列
	Tags        string `json:"tags,omitempty" request:"query,omitempty"`         // 笔记跳转标签列表
	Cls         int    `json:"cls,omitempty" request:"query,omitempty"`          // 1。作用尚不明确
	From        string `json:"from,omitempty" request:"query,omitempty"`         // 提交类型。auto：自动提交。save：手动提交。close：关闭时自动提交
	ContLen     int    `json:"cont_len,omitempty" request:"query,omitempty"`     // 正文字数
	Platform    string `json:"platform,omitempty" request:"query,omitempty"`     // 平台。可为 web
	Publish     int    `json:"publish,omitempty" request:"query,omitempty"`      // 是否公开笔记。0：不公开。1：公开
	AutoComment int    `json:"auto_comment,omitempty" request:"query,omitempty"` // 是否添加到评论区。0：不添加。1：添加
}

// SaveVideoNoteResult 是保存视频笔记的 data 对象。
type SaveVideoNoteResult struct {
	NoteID int64 `json:"note_id"` // 笔记 id
}

// SaveVideoNote 保存视频笔记。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
// 稿件已存在笔记时返回 79508，笔记所属视频不合法时返回 79513。
func (c *Client) SaveVideoNote(ctx context.Context, param SaveVideoNoteParam) (*SaveVideoNoteResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/note/add"
	)
	return bilibili.Execute[*SaveVideoNoteResult](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("oid", "oid_type", "note_id", "title", "summary", "content", "tags",
			"cls", "from", "cont_len", "platform", "publish", "auto_comment"),
		bilibili.FillFormCsrf(c.Client))
}

// DeleteVideoNoteParam 是删除视频笔记的请求参数，所有字段进入 URL 编码表单。
type DeleteVideoNoteParam struct {
	Oid    int   `json:"oid"`                                         // 目标稿件 avid
	NoteID int64 `json:"note_id,omitempty" request:"query,omitempty"` // 笔记 id
}

// DeleteVideoNote 删除视频笔记。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) DeleteVideoNote(ctx context.Context, param DeleteVideoNoteParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/note/del"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("oid", "note_id"),
		bilibili.FillFormCsrf(c.Client))
	return err
}
