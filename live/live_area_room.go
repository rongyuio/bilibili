package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播分区房间列表相关接口。响应模型见 live_area_room_model.go。

// GetLiveAreaRoomListParam 指定获取直播分区下直播间的条件。
type GetLiveAreaRoomListParam struct {
	ParentAreaID int64  `json:"parent_area_id"`                     // 父分区 id；0 表示不限制父分区
	AreaID       int64  `json:"area_id"`                            // 子分区 id；0 表示全部子分区
	SortType     string `json:"sort_type" request:"default=online"` // 排序方式。留空时按人气排序；也可传 dynamic
	Page         int    `json:"page" request:"default=1"`           // 页码，从 1 开始。留空时自动填 1
	PageSize     int    `json:"page_size" request:"default=20"`     // 每页数量。留空时自动填 20
}

// GetLiveAreaRoomList 获取指定直播分区下的直播间列表。
// 传零值的 ParentAreaID 与 AreaID 可获取全站直播间；传父分区 id、AreaID 为 0 可获取该父分区的全部房间。
// 响应是房间数组，不提供总数或 has_more；翻页到末尾时返回空数组。
func (c *Client) GetLiveAreaRoomList(ctx context.Context, param GetLiveAreaRoomListParam) ([]GetLiveAreaRoomListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Area/getRoomList"
	)
	return bilibili.Execute[[]GetLiveAreaRoomListItem](ctx, c.Client, method, url, param)
}
