package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 大航海与粉丝团成员相关接口。响应模型见 live_guard_model.go。

// GetLiveGuardMemberListParam 指定要查询大航海成员的直播间与分页条件。
type GetLiveGuardMemberListParam struct {
	RoomID   int64 `json:"roomid"`                                        // 直播间号
	Page     int   `json:"page"`                                          // 页数
	RUID     int64 `json:"ruid"`                                          // 主播 UID
	PageSize int   `json:"page_size,omitempty" request:"query,omitempty"` // 页大小。默认 20，最大 30，超过则按 10 处理
	Typ      int   `json:"typ,omitempty" request:"query,omitempty"`       // 排序方式。3、4、5 分别为按周/月/总航海亲密度排序
}

// GetLiveGuardMemberList 查询直播间的大航海成员。
//
// page=1 时返回的 list 从榜单第四名开始，前三名在 top3 中。
func (c *Client) GetLiveGuardMemberList(ctx context.Context, param GetLiveGuardMemberListParam) (*GetLiveGuardMemberListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-room/v2/guardTab/topListNew"
	)
	return bilibili.Execute[*GetLiveGuardMemberListResult](ctx, c.Client, method, url, param,
		liveReferer(param.RoomID))
}

// GetLiveFansMemberListParam 指定要查询粉丝团成员的主播与分页条件。
type GetLiveFansMemberListParam struct {
	Page     int   `json:"page"`                                          // 页数
	RUID     int64 `json:"ruid"`                                          // 主播 UID
	PageSize int   `json:"page_size"`                                     // 每页返回的数量。最大 30，超过则按 10 处理
	RankType int   `json:"rank_type,omitempty" request:"query,omitempty"` // 排序方式。1：按粉丝牌仍亮着的成员亲密度排序；2：按所有未上舰成员亲密度排序
	Ts       int64 `json:"ts,omitempty" request:"query,omitempty"`        // 13 位时间戳。rank_type=2 时必要，该值 >=1000 即可
}

// GetLiveFansMemberList 查询主播的粉丝团成员榜单。
func (c *Client) GetLiveFansMemberList(ctx context.Context, param GetLiveFansMemberListParam) (*GetLiveFansMemberListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/general-interface/v1/rank/getFansMembersRank"
	)
	return bilibili.Execute[*GetLiveFansMemberListResult](ctx, c.Client, method, url, param,
		liveReferer(0))
}
