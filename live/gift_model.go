package live

// 直播礼物相关响应模型。字段定义依据接口文档 live/gift.md。

// GetLiveRoomGiftListResult 是直播间礼物列表接口的响应本体。
type GetLiveRoomGiftListResult struct {
	GiftConfig GetLiveRoomGiftConfig `json:"gift_config"` // 礼物面板配置
}

// GetLiveRoomGiftConfig 是礼物面板的配置信息。
type GetLiveRoomGiftConfig struct {
	BaseConfig GetLiveRoomGiftBaseConfig `json:"base_config"` // 基础礼物配置
}

// GetLiveRoomGiftBaseConfig 是基础礼物配置，礼物列表位于 list 中。
type GetLiveRoomGiftBaseConfig struct {
	List []GetLiveRoomGiftListItem `json:"list"` // 礼物列表
}

// GetLiveRoomGiftListItem 是直播间内一个可投喂礼物的信息。
type GetLiveRoomGiftListItem struct {
	ID                int    `json:"id"`                  // 礼物id
	Name              string `json:"name"`                // 礼物名字
	Price             int    `json:"price"`               // 价格，该值/1000 的单位为元
	Type              int    `json:"type"`                // 礼物类型
	CoinType          string `json:"coin_type"`           // 货币类型，一般为 gold，即电池
	Effect            int    `json:"effect"`              // 特效类型，观察到的值可能为 0、2、3
	StayTime          int    `json:"stay_time"`           // 礼物展示的时间，观察到的值均为 3
	AnimationFrameNum int    `json:"animation_frame_num"` // 礼物动画帧数
	Desc              string `json:"desc"`                // 礼物描述
	ImgBasic          string `json:"img_basic"`           // 礼物图片
	Gif               string `json:"gif"`                 // 礼物gif动画
}

// GetLiveBlindBoxProbabilityResult 是盲盒概率接口的响应本体。
type GetLiveBlindBoxProbabilityResult struct {
	NoteText      string                            `json:"note_text"`       // 描述
	BlindPrice    int                               `json:"blind_price"`     // 盲盒价格
	BlindGiftName string                            `json:"blind_gift_name"` // 盲盒名字
	Gifts         []GetLiveBlindBoxProbabilityGift  `json:"gifts"`           // 各档奖品的爆出概率
	FridayYqID    int                               `json:"friday_yq_id"`    // (?)，实测为 109038
	AbRes         int                               `json:"ab_res"`          // (?)，实测为 1
	PreImgs       []GetLiveBlindBoxProbabilityImage `json:"pre_imgs"`        // 预览图列表
	IsFirst       bool                              `json:"is_first"`        // 是否为首抽
	UID           int64                             `json:"uid"`             // 用户 UID
	ConfID        int                               `json:"conf_id"`         // 盲盒配置id
}

// GetLiveBlindBoxProbabilityGift 是盲盒可能爆出的一个礼物及其概率。
type GetLiveBlindBoxProbabilityGift struct {
	GiftID    int    `json:"gift_id"`     // 爆出的礼物id
	Price     int    `json:"price"`       // 爆出的礼物价格
	GiftName  string `json:"gift_name"`   // 礼物名字
	GiftImg   string `json:"gift_img"`    // 礼物图片
	Chance    string `json:"chance"`      // 概率
	IsWinGift int    `json:"is_win_gift"` // 是否为中奖礼物，实测 1×5 / 0×2
}

// GetLiveBlindBoxProbabilityImage 是盲盒奖品的预览图。
type GetLiveBlindBoxProbabilityImage struct {
	GiftID     int    `json:"gift_id"`     // 礼物id
	PreviewURL string `json:"preview_url"` // 预览图url
}
