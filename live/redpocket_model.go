package live

// 直播间人气红包响应模型。信封 code/message/ttl 由库统一处理。

// GetLiveRedPocketInfoResult 是查询直播间红包信息的响应（data 部分）。
type GetLiveRedPocketInfoResult struct {
	PK                  any                                       `json:"pk"`                    // 可能为 null
	Guard               any                                       `json:"guard"`                 // 可能为 null
	Gift                any                                       `json:"gift"`                  // 可能为 null
	Storm               any                                       `json:"storm"`                 // 可能为 null
	Silver              any                                       `json:"silver"`                // 可能为 null
	ActivityBox         GetLiveRedPocketInfoActivityBox           `json:"activity_box"`          // 活动宝箱信息
	Danmu               any                                       `json:"danmu"`                 // 可能为 null
	Anchor              any                                       `json:"anchor"`                // 可能为 null
	RedPocket           any                                       `json:"red_pocket"`            // 可能为 null
	PopularityRedPocket []GetLiveRedPocketInfoPopularityRedPocket `json:"popularity_red_pocket"` // 人气红包信息。无红包时为 null
	ActivityBoxInfo     any                                       `json:"activity_box_info"`     // 可能为 null
}

// GetLiveRedPocketInfoActivityBox 是直播间活动宝箱信息。
type GetLiveRedPocketInfoActivityBox struct {
	ActivityID  int64  `json:"ACTIVITY_ID"`  // 活动 id
	ActivityPic string `json:"ACTIVITY_PIC"` // 活动图片
}

// GetLiveRedPocketInfoPopularityRedPocket 是直播间人气红包的一条记录。
type GetLiveRedPocketInfoPopularityRedPocket struct {
	LotID           int64                       `json:"lot_id"`           // 红包 id
	SenderUID       int64                       `json:"sender_uid"`       // 红包发送者 uid
	SenderName      string                      `json:"sender_name"`      // 红包发送者昵称
	SenderFace      string                      `json:"sender_face"`      // 红包发送者头像
	JoinRequirement int                         `json:"join_requirement"` // 参与条件
	Danmu           string                      `json:"danmu"`            // 参与红包时自动发送的弹幕内容
	Awards          []GetLiveRedPocketInfoAward `json:"awards"`           // 红包内容
	LotStatus       int                         `json:"lot_status"`       // 抽奖状态。作用尚不明确
	H5URL           string                      `json:"h5_url"`           // 红包界面
	UserStatus      int                         `json:"user_status"`      // 用户是否已参与。1：已参与；2：未参与
	LotConfigID     int                         `json:"lot_config_id"`    // 抽奖配置 id。作用尚不明确
	TotalPrice      int                         `json:"total_price"`      // 红包总计价格
	StartTime       int64                       `json:"start_time"`       // 红包开始时间（秒级时间戳）
	EndTime         int64                       `json:"end_time"`         // 红包结束时间（秒级时间戳）
	LastTime        int                         `json:"last_time"`        // 红包持续时长（秒）
	RemoveTime      int64                       `json:"remove_time"`      // 红包移除时间（秒级时间戳）
	ReplaceTime     int64                       `json:"replace_time"`     // 红包替换时间（秒级时间戳）
	CurrentTime     int64                       `json:"current_time"`     // 服务器当前时间（秒级时间戳）
}

// GetLiveRedPocketInfoAward 是人气红包中的一个奖品。
type GetLiveRedPocketInfoAward struct {
	GiftID   int64  `json:"gift_id"`   // 礼物 id
	Num      int    `json:"num"`       // 礼物数量
	GiftName string `json:"gift_name"` // 礼物名称
	GiftPic  string `json:"gift_pic"`  // 礼物图片
}
