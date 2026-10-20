package live

// 直播流水的响应模型。接口见 live_bill.go。

// GetLiveGiftTypesItem 是礼物列表中的一项。
type GetLiveGiftTypesItem struct {
	GiftID   int64  `json:"gift_id"`   // 礼物 id
	GiftName string `json:"gift_name"` // 礼物名称
	Price    *int64 `json:"price"`     // 瓜子数量。电池礼物为金瓜子数量，银瓜子礼物为银瓜子数量（金瓜子数量 / 100 = 电池数量）；舰长、提督、总督等特殊礼物没有该字段
}

// GetLiveGiftStreamResult 是直播流水的查询结果。
type GetLiveGiftStreamResult struct {
	HasMore      int                     `json:"has_more"`      // 是否有下一页。1-是，0-否
	TotalHamster int64                   `json:"total_hamster"` // 总的金仓鼠收益
	List         []GetLiveGiftStreamItem `json:"list"`          // 礼物列表
}

// GetLiveGiftStreamItem 是直播流水中的一条收礼记录。
type GetLiveGiftStreamItem struct {
	UID           int64  `json:"uid"`            // 送礼用户的 uid
	Uname         string `json:"uname"`          // 用户名
	Time          string `json:"time"`           // 送礼时间
	GiftID        int64  `json:"gift_id"`        // 礼物 id
	GiftName      string `json:"gift_name"`      // 礼物名字
	GiftImg       string `json:"gift_img"`       // 礼物图片链接
	GiftNum       int    `json:"gift_num"`       // 礼物数量
	Hamster       int64  `json:"hamster"`        // 金仓鼠数量
	Gold          int64  `json:"gold"`           // 礼物价值（金瓜子）
	Silver        int64  `json:"silver"`         // 礼物价值（银瓜子）
	IosHamster    int64  `json:"ios_hamster"`    // 由 iOS 端送出的礼物所收到的金仓鼠
	NormalHamster int64  `json:"normal_hamster"` // 一般情况下收到的金仓鼠
	IosGold       int64  `json:"ios_gold"`       // 由 iOS 端送出的礼物所收到的金瓜子数量
	NormalGold    int64  `json:"normal_gold"`    // 一般情况下收到的金瓜子数量
	IsHybrid      bool   `json:"is_hybrid"`      // 是否混合。作用不明
	ID            int64  `json:"id"`             // 此项 id，用于翻页
	//nolint:misspell // 接口字段名原样保留
	IsOpenPlatfrom int `json:"is_open_platfrom"` // 是否开放平台。作用不明
	//nolint:misspell // 接口字段名原样保留
	OpenPlatfromRate int    `json:"open_platfrom_rate"` // 开放平台比率 (?)。作用不明
	ReceiveTitle     string `json:"receive_title"`      // 作用不明，一般为「主播」
	RoomID           int64  `json:"room_id"`            // 送礼房间 id。如果礼物为上舰，此项为 0
}
