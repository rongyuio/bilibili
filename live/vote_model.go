package live

// 直播投票相关响应模型。

// GetLiveVotePanelResult 是查询投票信息的 data 部分。
type GetLiveVotePanelResult struct {
	VoteInfo  GetLiveVotePanelVoteInfo   `json:"vote_info"` // 当前活动的投票信息。当前没有活动的投票时为空对象
	Templates []GetLiveVotePanelTemplate `json:"templates"` // 投票模板
}

// GetLiveVotePanelVoteInfo 是一场投票的信息。
type GetLiveVotePanelVoteInfo struct {
	Status        int                      `json:"status"`         // 投票状态。0：NONE，1：等待审核，2：审核失败，4：投票进行中，5：投票结束，6：STOP
	Question      string                   `json:"question"`       // 投票问题
	Options       []GetLiveVotePanelOption `json:"options"`        // 投票选项
	Duration      int                      `json:"duration"`       // 投票持续时间，单位毫秒
	Result        int                      `json:"result"`         // 投票结果偏向哪一方状态
	ResultText    string                   `json:"result_text"`    // 投票结果偏向哪一方
	EtimeStr      string                   `json:"etime_str"`      // 投票结束时间
	LeftDuration  int                      `json:"left_duration"`  // 投票剩余时间，单位毫秒。投票结束后不存在
	InteractionID int64                    `json:"interaction_id"` // 投票互动 id
	TemplateID    int64                    `json:"template_id"`    // 投票模板 id。使用模板时存在
}

// GetLiveVotePanelOption 是投票中的一个选项。
type GetLiveVotePanelOption struct {
	Idx     int     `json:"idx"`     // 选项 id
	Desc    string  `json:"desc"`    // 投票选项描述
	Percent float64 `json:"percent"` // 投票选项占比显示
}

// GetLiveVotePanelTemplate 是一个投票模板。
type GetLiveVotePanelTemplate struct {
	TemplateID int64  `json:"template_id"` // 模板 id
	Question   string `json:"question"`    // 投票问题
	OptionA    string `json:"option_a"`    // 投票选项 A
	OptionB    string `json:"option_b"`    // 投票选项 B
}

// GetLiveVoteHistoryResult 是查询投票历史的 data 部分。
type GetLiveVoteHistoryResult struct {
	History []GetLiveVoteHistoryItem `json:"history"` // 投票历史列表
}

// GetLiveVoteHistoryItem 是一场投票的信息。
type GetLiveVoteHistoryItem struct {
	Status        int                        `json:"status"`         // 投票状态。0：NONE，1：等待审核，2：审核失败，4：投票进行中，5：投票结束，6：STOP
	Question      string                     `json:"question"`       // 投票问题
	Options       []GetLiveVoteHistoryOption `json:"options"`        // 投票选项
	Duration      int                        `json:"duration"`       // 投票持续时间，单位毫秒
	Result        int                        `json:"result"`         // 投票结果偏向哪一方状态
	ResultText    string                     `json:"result_text"`    // 投票结果偏向哪一方
	EtimeStr      string                     `json:"etime_str"`      // 投票结束时间
	LeftDuration  int                        `json:"left_duration"`  // 投票剩余时间，单位毫秒。投票结束后不存在
	InteractionID int64                      `json:"interaction_id"` // 投票互动 id
	TemplateID    int64                      `json:"template_id"`    // 投票模板 id。使用模板时存在
}

// GetLiveVoteHistoryOption 是投票中的一个选项。
type GetLiveVoteHistoryOption struct {
	Idx     int     `json:"idx"`     // 选项 id
	Desc    string  `json:"desc"`    // 投票选项描述
	Percent float64 `json:"percent"` // 投票选项占比显示
}
