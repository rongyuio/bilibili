package message

import (
	"context"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 消息设置相关接口。文档见 message/settings.md，各接口的返回体独立定义，互不复用。

// splitQueryParam 把 encodeParams 以逗号拼接的单值切片参数还原为同名重复参数。
// 自动回复接口要求 type[] 重复出现，而不是逗号分隔。
func splitQueryParam(key string) bilibili.Handler {
	return func(r *resty.Request) error {
		values := r.QueryParam[key]
		if len(values) == 1 && values[0] != "" {
			r.QueryParam[key] = strings.Split(values[0], ",")
		}
		return nil
	}
}

// GetUserPreferenceSettingsParam 是获取用户偏好设置的请求参数。
type GetUserPreferenceSettingsParam struct {
	MsgNotify         int    `json:"msg_notify,omitempty" request:"query,omitempty"`          // 是否显示“消息提醒”设置。为任意整数时返回数据中包含 msg_notify
	ShowUnfollowedMsg int    `json:"show_unfollowed_msg,omitempty" request:"query,omitempty"` // 是否显示“收起未关注人消息”设置。为任意整数时返回数据中包含 show_unfollowed_msg
	Build             int    `json:"build,omitempty" request:"query,omitempty"`               // 客户端内部版本号。默认为 0
	MobiApp           string `json:"mobi_app,omitempty" request:"query,omitempty"`            // 平台标识。可为 web 等
}

// GetUserPreferenceSettingsResult 是获取用户偏好设置接口的 data 对象。
type GetUserPreferenceSettingsResult struct {
	ShowUnfollowedMsg      int `json:"show_unfollowed_msg"`      // 收起未关注人消息。0：不收起。1：收起。仅请求参数 show_unfollowed_msg 为整数时返回
	MsgNotify              int `json:"msg_notify"`               // 消息提醒。1：接收。3：不接收。仅请求参数 msg_notify 为整数时返回
	SetLike                int `json:"set_like"`                 // 收到的赞提醒（旧）。0：接收。5：不接收
	SetComment             int `json:"set_comment"`              // 回复我的提醒（旧）。0：所有人。1：关注的人。2：不接受任何消息提醒
	SetAt                  int `json:"set_at"`                   // @ 我的提醒（旧）。0：所有人。1：关注的人。2：不接受任何消息提醒
	IsGroupFold            int `json:"is_group_fold"`            // 收起应援团消息。0：不收起。1：收起
	ShouldReceiveGroup     int `json:"should_receive_group"`     // 接收应援团消息。0：不接收。1：接收
	ReceiveUnfollowMsg     int `json:"receive_unfollow_msg"`     // 接收未关注人消息。前端隐藏此开关。0：不接收。1：接收
	FollowedReply          int `json:"followed_reply"`           // 被关注回复。0：关闭。1：开启
	KeysReply              int `json:"keys_reply"`               // 关键词回复。0：关闭。1：开启
	RecvReply              int `json:"recv_reply"`               // 收到消息回复。0：关闭。1：开启
	VoyageReply            int `json:"voyage_reply"`             // 大航海上船回复。0：关闭。1：开启
	RecommendFollowedReply int `json:"recommend_followed_reply"` // 被关注时自动推送作品。0：关闭。1：开启
	AiIntercept            int `json:"ai_intercept"`             // 私信智能拦截。0：关闭。1：开启
	AntiHarassment         any `json:"anti_harassment"`          // 防骚扰和互动人群设置。此接口恒返回 null
	SetRecvReply           int `json:"set_recv_reply"`           // 回复与@提醒（新）。0：所有人。1：关注的人。2：不接受任何消息提醒
	SetRecvLike            int `json:"set_recv_like"`            // 收到喜欢提醒（新）。0：接收。2：不接收
	SetNewFollow           int `json:"set_new_follow"`           // 新增粉丝提醒（新）。0：接收。2：不接收
}

// GetUserPreferenceSettings 获取用户偏好设置。文档允许 GET 或 POST，这里使用 GET，因此不需要 CSRF。
func (c *Client) GetUserPreferenceSettings(ctx context.Context, param GetUserPreferenceSettingsParam) (*GetUserPreferenceSettingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_setting/v1/link_setting/get"
	)
	return bilibili.Execute[*GetUserPreferenceSettingsResult](ctx, c.Client, method, url, param)
}

// GetMessageBlockWordsParam 是获取消息屏蔽词的请求参数。
type GetMessageBlockWordsParam struct {
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetMessageBlockWordsItem 是一条消息屏蔽词。
type GetMessageBlockWordsItem struct {
	Content string `json:"content"` // 屏蔽词内容
	Ctime   int64  `json:"ctime"`   // 屏蔽词添加时间。秒级时间戳
}

// GetMessageBlockWordsResult 是获取消息屏蔽词接口的 data 对象。
type GetMessageBlockWordsResult struct {
	Words         []GetMessageBlockWordsItem `json:"words"`           // 消息屏蔽词列表。无效时为 null
	Message       string                     `json:"message"`         // 作用尚不明确
	MaxWordLength int                        `json:"max_word_length"` // 单个屏蔽词最大长度。目前为 20
	MaxWordsSize  int                        `json:"max_words_size"`  // 屏蔽词最大数量。目前为 200
}

// GetMessageBlockWords 获取消息屏蔽词
func (c *Client) GetMessageBlockWords(ctx context.Context, param GetMessageBlockWordsParam) (*GetMessageBlockWordsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/x/im/link_setting/get_block_words"
	)
	return bilibili.Execute[*GetMessageBlockWordsResult](ctx, c.Client, method, url, param)
}

// AddMessageBlockWordParam 是添加消息屏蔽词的请求参数。
type AddMessageBlockWordParam struct {
	Content string `json:"content"`                                      // 屏蔽词内容。最多 20 个字符
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// AddMessageBlockWord 添加消息屏蔽词。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) AddMessageBlockWord(ctx context.Context, param AddMessageBlockWordParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/x/im/link_setting/add_block_word"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("content", "build", "mobi_app"),
		bilibili.FillFormCsrf(c.Client))
	return err
}

// DeleteMessageBlockWordParam 是删除消息屏蔽词的请求参数。
type DeleteMessageBlockWordParam struct {
	Content string `json:"content"`                                      // 屏蔽词内容
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// DeleteMessageBlockWord 删除消息屏蔽词。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) DeleteMessageBlockWord(ctx context.Context, param DeleteMessageBlockWordParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/x/im/link_setting/del_block_word"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("content", "build", "mobi_app"),
		bilibili.FillFormCsrf(c.Client))
	return err
}

// GetAutoReplyTextsParam 是获取自动回复文本的请求参数。
type GetAutoReplyTextsParam struct {
	Type    []int  `json:"type[]" request:"query"`                       // 自动回复类型。1：被关注回复。2：关键词回复。3：收到消息回复。5：大航海上船回复
	UID     int64  `json:"uid,omitempty" request:"query,omitempty"`      // 自己的 mid
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetAutoReplyTextsItem 是一条自动回复规则。
type GetAutoReplyTextsItem struct {
	ID    int64  `json:"id"`    // 自动回复规则 id
	UID   int64  `json:"uid"`   // 自己的 mid
	Type  int    `json:"type"`  // 自动回复类型。同请求参数 type[]
	Title string `json:"title"` // 规则名称。仅关键词回复类型有效，否则为空文本
	Key1  string `json:"key1"`  // 精确匹配关键词。仅关键词回复类型有效，以中文逗号分隔每一个关键词
	Key2  string `json:"key2"`  // 模糊匹配关键词。仅关键词回复类型有效，以中文逗号分隔每一个关键词
	Reply string `json:"reply"` // 回复内容
}

// GetAutoReplyTextsResult 是获取自动回复文本接口的 data 对象。
type GetAutoReplyTextsResult struct {
	Texts []GetAutoReplyTextsItem `json:"texts"` // 自动回复文本。未设置该类型自动回复文本时没有此项
}

// GetAutoReplyTexts 获取自动回复文本/关键词回复规则。
// type[] 是重复参数，用 splitQueryParam 把编码后的逗号分隔值拆回多个同名 query 参数。
func (c *Client) GetAutoReplyTexts(ctx context.Context, param GetAutoReplyTextsParam) (*GetAutoReplyTextsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/x/im/auto_reply/get_reply_text"
	)
	return bilibili.Execute[*GetAutoReplyTextsResult](ctx, c.Client, method, url, param, splitQueryParam("type[]"))
}

// GetSystemSettingsParam 是获取系统设置的请求参数。
type GetSystemSettingsParam struct {
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetSystemSettingsOldUpAssistantDoor 是原“UP主小助手”入口相关信息。
type GetSystemSettingsOldUpAssistantDoor struct {
	ShowOldUpDoor int    `json:"show_old_up_door"` // 是否显示原“UP主小助手”入口。目前恒为 0
	Title         string `json:"title"`            // 原“UP主小助手”入口标题。目前为空文本
	SubTitle      string `json:"sub_title"`        // 原“UP主小助手”入口子标题。目前为空文本
}

// GetSystemSettingsSessionCfg 是某个私信会话的配置信息。
type GetSystemSettingsSessionCfg struct {
	HiddenEmoteBtn bool   `json:"hidden_emote_btn"` // 是否隐藏表情按钮
	HiddenPicBtn   bool   `json:"hidden_pic_btn"`   // 是否隐藏发送图片按钮
	MaxHeight      int    `json:"max_height"`       // 最大高度。作用尚不明确
	HintText       string `json:"hint_text"`        // 私信输入框的提示内容
	HiddenTopHint  bool   `json:"hidden_top_hint"`  // 作用尚不明确
	IsGptAccount   bool   `json:"is_gpt_account"`   // 是否为 AI 会话
}

// GetSystemSettingsGameMsg 是游戏相关消息配置。仅当请求参数 mobi_app 不为 web 时存在其中的项目。
type GetSystemSettingsGameMsg struct {
	ShowInstall bool `json:"show_install"` // 作用尚不明确
}

// GetSystemSettingsResult 是获取系统设置接口的 data 对象。
type GetSystemSettingsResult struct {
	IsCreateGroupAvailable        int                                    `json:"is_create_group_available"`         // 是否可创建粉丝团。粉丝数达到 1000 后开启
	IsAutoReplyAvailable          int                                    `json:"is_auto_reply_available"`           // 是否可使用自动回复功能。粉丝数达到 1000 后开启
	VcHintTitle                   string                                 `json:"vc_hint_title"`                     // 稿件自动推送小灰条内容
	VcHintTitleButton             string                                 `json:"vc_hint_title_button"`              // 稿件自动推送小灰条旁的按钮内容
	VcHintDetail                  string                                 `json:"vc_hint_detail"`                    // 稿件自动推送详细信息小灰条内容
	VcHintDetailButton            string                                 `json:"vc_hint_detail_button"`             // 关闭稿件自动推送按钮内容
	AutoReplyHTML                 string                                 `json:"auto_reply_html"`                   // 自动回复 H5 页面 url
	IsReceiveUnfollowWL           int                                    `json:"is_receive_unfollow_wl"`            // 是否显示限制接收未关注人消息功能
	IsVoyage                      int                                    `json:"is_voyage"`                         // 是否在自动回复页面显示“大航海自动回复”选项
	IsAutoReplyRecommendAvailable int                                    `json:"is_auto_reply_recommend_available"` // 是否在自动回复的“被关注回复”页面显示推送往期作品复选框
	IsDiscussStyleIMPage          int                                    `json:"is_discuss_style_im_page"`          // 作用尚不明确
	DiscussUnreadStyleIMPage      int                                    `json:"discuss_unread_style_im_page"`      // 作用尚不明确
	OldUpAssistantDoor            GetSystemSettingsOldUpAssistantDoor    `json:"old_up_assistant_door"`             // 原“UP主小助手”入口相关信息
	IsNewUpAssistantEffective     int                                    `json:"is_new_up_assistant_effective"`     // 新“UP主小助手”是否启用
	IsArchiveGray                 bool                                   `json:"is_archive_gray"`                   // 是否不显示“私信存档”按钮
	SessionCfg                    map[string]GetSystemSettingsSessionCfg `json:"session_cfg"`                       // 特定私信会话配置信息。键为用户 mid
	MigrateSessionAPI             bool                                   `json:"migrate_session_api"`               // 作用尚不明确
	GameMsg                       *GetSystemSettingsGameMsg              `json:"game_msg"`                          // 作用尚不明确。仅当请求参数 mobi_app 不为 web 时存在其中的项目
	AutoReplyMsgDesc              string                                 `json:"auto_reply_msg_desc"`               // 指示消息为自动回复消息的提示内容
	HuahuoGroupIconNew            string                                 `json:"huahuo_group_icon_new"`             // 浅色模式下的花火图标 url
	HuahuoGroupIconDarkNew        string                                 `json:"huahuo_group_icon_dark_new"`        // 深色模式下的花火图标 url
	ImDisabledInputHint           string                                 `json:"im_disabled_input_hint"`            // 某 UP 主禁用私信功能时的提示内容
}

// GetSystemSettings 获取系统设置。此接口中的设置一般不能由用户随意更改。
func (c *Client) GetSystemSettings(ctx context.Context, param GetSystemSettingsParam) (*GetSystemSettingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_setting/v1/link_setting/get_sys_setting"
	)
	return bilibili.Execute[*GetSystemSettingsResult](ctx, c.Client, method, url, param)
}

// GetCreateGroupIconVisibilityParam 是获取创建粉丝团按钮可见性的请求参数。
type GetCreateGroupIconVisibilityParam struct {
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetCreateGroupIconVisibilityResult 是获取创建粉丝团按钮可见性接口的 data 对象。
type GetCreateGroupIconVisibilityResult struct {
	Show int `json:"show"` // 是否显示创建粉丝团按钮。粉丝数达到 1000 后显示
}

// GetCreateGroupIconVisibility 获取创建粉丝团按钮可见性
func (c *Client) GetCreateGroupIconVisibility(ctx context.Context, param GetCreateGroupIconVisibilityParam) (*GetCreateGroupIconVisibilityResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_group/v1/member/show_create_group_icon"
	)
	return bilibili.Execute[*GetCreateGroupIconVisibilityResult](ctx, c.Client, method, url, param)
}
