package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播流水相关接口。响应模型见 live_bill_model.go。
// 文档见 live/live_bill.md。

// GetLiveGiftTypes 获取所有礼物列表（含礼物 id、名称与价格）。
//
// 认证方式：Cookie（SESSDATA）。舰长、提督、总督等特殊礼物没有价格字段。
func (c *Client) GetLiveGiftTypes(ctx context.Context) ([]GetLiveGiftTypesItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/gift/v1/master/getGiftTypes"
	)
	return bilibili.Execute[[]GetLiveGiftTypesItem](ctx, c.Client, method, url, nil, liveMiscReferer())
}

// GetLiveGiftStreamParam 指定直播流水的查询条件。
type GetLiveGiftStreamParam struct {
	Limit     int    `json:"limit" request:"query"`                       // 一页有多少条目
	CoinType  int    `json:"coin_type" request:"query"`                   // 礼物类型。0-所有，1-电池礼物，2-银瓜子礼物
	BeginTime string `json:"begin_time" request:"query"`                  // 流水的日期，格式为 yyyy-MM-dd
	Uname     string `json:"uname,omitempty" request:"query,omitempty"`   // 筛选的用户名
	LastID    int64  `json:"last_id,omitempty" request:"query,omitempty"` // 上一页页末的礼物列表 id，翻页时必填
	GiftID    int64  `json:"gift_id,omitempty" request:"query,omitempty"` // 筛选的礼物 id
}

// GetLiveGiftStream 获取当前账号某一日的直播流水（收到礼物的记录），不自动翻页。
//
// 认证方式：Cookie（SESSDATA）。翻页时把上一页最后一条记录的 ID 填入 LastID。
func (c *Client) GetLiveGiftStream(ctx context.Context, param GetLiveGiftStreamParam) (*GetLiveGiftStreamResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/revenue/v1/giftStream/getReceivedGiftStreamNextList"
	)
	return bilibili.Execute[*GetLiveGiftStreamResult](ctx, c.Client, method, url, param, liveMiscReferer())
}
