package bilibili

// 活动抽奖相关响应模型。

// GetActivityLotteryTimesResult contains the corresponding lottery response fields.
type GetActivityLotteryTimesResult struct {
	Times         int `json:"times"`
	LotteryType   int `json:"lottery_type"`
	Points        int `json:"points"`
	PointsPerTime int `json:"points_per_time"`
	Intergral     any `json:"intergral"`
	Stime         int `json:"stime"`
	Etime         int `json:"etime"`
}

// GetDynamicLotteryInfoResult contains the corresponding lottery response fields.
type GetDynamicLotteryInfoResult struct {
	LotteryID         int                          `json:"lottery_id"`
	SenderUID         int                          `json:"sender_uid"`
	BusinessType      int                          `json:"business_type"`
	BusinessID        int64                        `json:"business_id"`
	Status            int                          `json:"status"` // 2=已开奖
	LotteryTime       int                          `json:"lottery_time"`
	LotteryAtNum      int                          `json:"lottery_at_num"`
	LotteryFeedLimit  int                          `json:"lottery_feed_limit"`
	NeedPost          int                          `json:"need_post"`
	FirstPrize        int                          `json:"first_prize"`
	SecondPrize       int                          `json:"second_prize"`
	ThirdPrize        int                          `json:"third_prize"`
	Ts                int                          `json:"ts"`
	Participants      int                          `json:"participants"`
	HasChargeRight    bool                         `json:"has_charge_right"`
	Participated      bool                         `json:"participated"`
	Followed          bool                         `json:"followed"`
	Reposted          bool                         `json:"reposted"`
	LotteryDetailURL  string                       `json:"lottery_detail_url"`
	FirstPrizeCmt     string                       `json:"first_prize_cmt"`
	ThirdPrizeCmt     string                       `json:"third_prize_cmt"`
	FirstPrizePic     string                       `json:"first_prize_pic"`
	SecondPrizePic    string                       `json:"second_prize_pic"`
	ThirdPrizePic     string                       `json:"third_prize_pic"`
	VipBatchSign      string                       `json:"vip_batch_sign"`
	VipRedirectURL    string                       `json:"vip_redirect_url"`
	UpowerRedirectURL string                       `json:"upower_redirect_url"`
	PrizeTypeFirst    DynamicLotteryFirstPrizeType `json:"prize_type_first"`
	LotteryResult     DynamicLotteryResult         `json:"lottery_result"`
}

// DynamicLotteryFirstPrizeType contains the corresponding lottery response fields.
type DynamicLotteryFirstPrizeType struct {
	Type  int                      `json:"type"`
	Value DynamicLotteryPrizeValue `json:"value"`
}

// DynamicLotteryPrizeValue contains the corresponding lottery response fields.
type DynamicLotteryPrizeValue struct {
	Stype int `json:"stype"`
	Count int `json:"count"`
}

// DynamicLotteryResult contains the corresponding lottery response fields.
type DynamicLotteryResult struct {
	FirstPrizeResult []DynamicLotteryWinner `json:"first_prize_result"`
}

// DynamicLotteryWinner contains the corresponding lottery response fields.
type DynamicLotteryWinner struct {
	UID          int    `json:"uid"`
	Name         string `json:"name"`
	Face         string `json:"face"`
	HongbaoMoney int    `json:"hongbao_money"`
}
