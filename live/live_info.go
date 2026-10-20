package live

import (
	"context"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播间信息相关接口。响应模型见 live_info_model.go。

// splitQueryParam 把 encodeParams 以逗号拼接的单值切片参数还原为同名重复参数。
// 个别直播接口要求 uids[]、room_ids 这类参数重复出现，而不是逗号分隔。
func splitQueryParam(key string) bilibili.Handler {
	return func(r *resty.Request) error {
		values := r.QueryParam[key]
		if len(values) == 1 && values[0] != "" {
			r.QueryParam[key] = strings.Split(values[0], ",")
		}
		return nil
	}
}

// GetUserLiveRoomStatusParam 指定要查询的用户。
type GetUserLiveRoomStatusParam struct {
	Mid int64 `json:"mid"` // 目标用户mid
}

// GetUserLiveRoomStatus 获取用户对应的直播间状态
func (c *Client) GetUserLiveRoomStatus(ctx context.Context, param GetUserLiveRoomStatusParam) (*GetUserLiveRoomStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Room/getRoomInfoOld"
	)
	return bilibili.Execute[*GetUserLiveRoomStatusResult](ctx, c.Client, method, url, param)
}

// GetRoomInitParam 指定要查询的房间页初始化信息。
type GetRoomInitParam struct {
	ID int64 `json:"id"` // 目标直播间号（短号）
}

// GetRoomInit 获取房间页初始化信息
func (c *Client) GetRoomInit(ctx context.Context, param GetRoomInitParam) (*GetRoomInitResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Room/room_init"
	)
	return bilibili.Execute[*GetRoomInitResult](ctx, c.Client, method, url, param)
}

// GetMasterInfoParam 指定要查询的主播。
type GetMasterInfoParam struct {
	UID int64 `json:"uid"` // 目标用户mid
}

// GetMasterInfo 获取主播信息
func (c *Client) GetMasterInfo(ctx context.Context, param GetMasterInfoParam) (*GetMasterInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/live_user/v1/Master/info"
	)
	return bilibili.Execute[*GetMasterInfoResult](ctx, c.Client, method, url, param)
}

// GetRoomBaseInfoParam 指定要查询的直播间。
type GetRoomBaseInfoParam struct {
	//nolint:misspell // 接口固定值原样保留
	ReqBiz  string  `json:"req_biz"`                                      // 固定为 web_room_componet
	RoomIDs []int64 `json:"room_ids,omitempty" request:"query,omitempty"` // 直播间短ID。多个重复该参数即可
}

// GetRoomBaseInfo 获取直播间基本信息，亦可用于批量获取
func (c *Client) GetRoomBaseInfo(ctx context.Context, param GetRoomBaseInfoParam) (*GetRoomBaseInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/index/getRoomBaseInfo"
	)
	return bilibili.Execute[*GetRoomBaseInfoResult](ctx, c.Client, method, url, param,
		splitQueryParam("room_ids"))
}

// GetRoomStatusInfoByUIDsParam 指定要查询的主播列表。
type GetRoomStatusInfoByUIDsParam struct {
	UIDs []int64 `json:"uids[]" request:"query"` // 要查询的主播mid
}

// GetRoomStatusInfoByUIDs 批量查询直播间状态
func (c *Client) GetRoomStatusInfoByUIDs(ctx context.Context, param GetRoomStatusInfoByUIDsParam) (map[string]GetRoomStatusInfoByUIDsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids"
	)
	return bilibili.Execute[map[string]GetRoomStatusInfoByUIDsItem](ctx, c.Client, method, url, param,
		splitQueryParam("uids[]"))
}

// GetRoomPlayInfoParam 指定要查询的直播间与直播流格式。
type GetRoomPlayInfoParam struct {
	RoomID    int64  `json:"room_id"`                                        // 直播间id
	Protocol  string `json:"protocol"`                                       // 直播协议。0：http_stream。1：http_hls。可多选，英文逗号分隔
	Format    string `json:"format"`                                         // 格式。0：flv。1：ts。2：fmp4。可多选，英文逗号分隔
	Codec     string `json:"codec"`                                          // 编码格式。0：AVC。1：HEVC。可多选，英文逗号分隔
	Qn        int    `json:"qn,omitempty" request:"query,omitempty"`         // 清晰度编码。默认150
	Platform  string `json:"platform,omitempty" request:"query,omitempty"`   // 平台。网页端为 web
	Ptype     int    `json:"ptype,omitempty" request:"query,omitempty"`      // 固定为 8
	Dolby     int    `json:"dolby,omitempty" request:"query,omitempty"`      // 固定为 5
	Panorama  int    `json:"panorama,omitempty" request:"query,omitempty"`   // 固定为 1
	OnlyAudio int    `json:"only_audio,omitempty" request:"query,omitempty"` // 是否为音频流。默认为视频流，1 为音频流
}

// GetRoomPlayInfo 获取直播间信息
func (c *Client) GetRoomPlayInfo(ctx context.Context, param GetRoomPlayInfoParam) (*GetRoomPlayInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo"
	)
	return bilibili.Execute[*GetRoomPlayInfoResult](ctx, c.Client, method, url, param)
}

// GetAnchorInRoomParam 指定要查询的直播间。
type GetAnchorInRoomParam struct {
	RoomID int64 `json:"roomid"` // 直播间号。可以为短号
}

// GetAnchorInRoom 获取直播间主播信息
func (c *Client) GetAnchorInRoom(ctx context.Context, param GetAnchorInRoomParam) (*GetAnchorInRoomResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/live_user/v1/UserInfo/get_anchor_in_room"
	)
	return bilibili.Execute[*GetAnchorInRoomResult](ctx, c.Client, method, url, param)
}
