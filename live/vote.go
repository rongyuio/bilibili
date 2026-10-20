package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播投票相关接口。响应模型见 vote_model.go。

// GetLiveVotePanelParam 指定要查询投票信息的直播间。
type GetLiveVotePanelParam struct {
	RoomID int64 `json:"room_id" request:"query"` // 直播间 id。必须为登录信息对应的直播间
}

// GetLiveVotePanel 查询直播间当前活动的投票信息与可用投票模板。
//
// 认证方式为 Cookie（SESSDATA）。非登录信息对应的直播间时 data 为 null。
func (c *Client) GetLiveVotePanel(ctx context.Context, param GetLiveVotePanelParam) (*GetLiveVotePanelResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-room/v1/dm/interaction/votePanel"
	)
	return bilibili.Execute[*GetLiveVotePanelResult](ctx, c.Client, method, url, param)
}

// GetLiveVoteHistoryParam 指定要查询投票历史的直播间。
type GetLiveVoteHistoryParam struct {
	RoomID int64 `json:"room_id" request:"query"` // 直播间 id。必须为登录信息对应的直播间
}

// GetLiveVoteHistory 查询直播间最近 10 条投票历史。
//
// 认证方式为 Cookie（SESSDATA）。非登录信息对应的直播间时 data 为空对象。
func (c *Client) GetLiveVoteHistory(ctx context.Context, param GetLiveVoteHistoryParam) (*GetLiveVoteHistoryResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-room/v1/dm/interaction/voteHistory"
	)
	return bilibili.Execute[*GetLiveVoteHistoryResult](ctx, c.Client, method, url, param)
}
