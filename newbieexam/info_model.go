package newbieexam

// 新用户答题信息查询相关响应模型。

// GetNewbieExamStatusResult 是查询答题状态的响应 data。
type GetNewbieExamStatusResult struct {
	Hid              int64  `json:"hid"`                // 答题会话id
	Mid              int    `json:"mid"`                // 答题用户mid
	Score            int    `json:"score"`              // 当前得分
	Status           int    `json:"status"`             // 答题状态。0：未答题，2：答题中，3：已通过
	Number           int    `json:"number"`             // 当前题号
	Result           string `json:"result"`             // 是否通过答题。failed：未通过，succeed：已通过
	Stage            string `json:"stage"`              // 当前答题阶段。base：基础题，extra：附加题，pro_type：等待选择自选题类型，pro：自选题，complete：已完成
	Version          string `json:"version"`            // 答题版本。当前为 v4
	StartTime        int64  `json:"start_time"`         // 本次答题开始时间。时间戳
	FirstAnswer      int    `json:"first_answer"`       // (?)。作用尚不明确
	Progress         string `json:"progress"`           // 当前答题进度。百分比，60分为100%
	Text             string `json:"text"`               // 提示文案
	URL              string `json:"url"`                // 答题页面url
	InRegAudit       bool   `json:"in_reg_audit"`       // 是否为第一次答题
	Edition          int    `json:"edition"`            // 答题版本。0：旧版（40+10+50），2：新版（40+30+30）
	Rewards          any    `json:"rewards"`            // (?)。实测为 null
	Captcha          int    `json:"captcha"`            // 是否已经提交验证码。1：已提交，仅新版提交验证码后存在
	MustRightNum     int    `json:"must_right_num"`     // 必答题需答对题数。实测为 60
	BaseErrNum       int    `json:"base_err_num"`       // 基础题错题数。实测为 0
	OnlyMemberPassed bool   `json:"only_member_passed"` // (?)。实测为 true
	ScoreTip         string `json:"score_tip"`          // (?)。实测为空
}

// GetNewbieExamProTypeListItem 是自选题父分类。
type GetNewbieExamProTypeListItem struct {
	ID     int                             `json:"id"`     // 父分类id
	Name   string                          `json:"name"`   // 父分类名
	Fields []GetNewbieExamProTypeListField `json:"fields"` // 子分类列表
}

// GetNewbieExamProTypeListField 是自选题父分类下的子分类。
type GetNewbieExamProTypeListField struct {
	ID   int    `json:"id"`   // 分类id
	Name string `json:"name"` // 分类名
}

// GetNewbieExamResultResult 是查询答题结果的响应 data。
//
// 类型名中重复的 Result 来自「方法名 + Result」的命名规则。
type GetNewbieExamResultResult struct {
	Hid            int64                             `json:"hid"`               // 答题会话id
	Mid            int                               `json:"mid"`               // 答题用户mid
	Member         int                               `json:"member"`            // (?)
	Score          int                               `json:"score"`             // 得分
	Level          int                               `json:"level"`             // 用户等级
	FirstPass      int                               `json:"first_pass"`        // (?)
	Uname          string                            `json:"uname"`             // 用户昵称
	Face           string                            `json:"face"`              // 用户头像url
	Status         string                            `json:"status"`            // 答题结果。failed：未通过，succeed：已通过
	QuestionTypes  []GetNewbieExamResultQuestionType `json:"question_types"`    // 已选择的自选题列表
	Power          []GetNewbieExamResultPower        `json:"power"`             // 自选题分类得分情况
	StartTime      int64                             `json:"start_time"`        // 开始答题时间。时间戳
	Share          GetNewbieExamResultShare          `json:"share"`             // (?)
	CanShowRankBtn bool                              `json:"can_show_rank_btn"` // (?)
	IsSameUser     bool                              `json:"is_same_user"`      // (?)
	ViewMore       string                            `json:"view_more"`         // (?)
	VideoInfo      GetNewbieExamResultVideoInfo      `json:"video_info"`        // (?)
	MainTids       []int                             `json:"main_tids"`         // (?)。文档未给出元素类型
	SubTids        []int                             `json:"sub_tids"`          // (?)。文档未给出元素类型
	PowerResult    any                               `json:"power_result"`      // (?)。实测为 null
	ScoreRate      int                               `json:"score_rate"`        // 分数超过平均用户的百分比
	Permission     GetNewbieExamResultPermission     `json:"permission"`        // (?)
	Rewards        any                               `json:"rewards"`           // (?)。实测为 null
}

// GetNewbieExamResultQuestionType 是答题结果中已选择的自选题分类。
type GetNewbieExamResultQuestionType struct {
	ID   int    `json:"id"`   // 分类id
	Name string `json:"name"` // 分类名
}

// GetNewbieExamResultPower 是答题结果中的自选题分类得分。
type GetNewbieExamResultPower struct {
	Score int    `json:"score"` // 该分类得分
	Name  string `json:"name"`  // 分类名
}

// GetNewbieExamResultShare 是答题结果的分享文案。
type GetNewbieExamResultShare struct {
	Content      string `json:"content"`       // (?)
	ShortContent string `json:"short_content"` // (?)
}

// GetNewbieExamResultVideoInfo 是答题结果中的视频信息。实测各字段均为空。
type GetNewbieExamResultVideoInfo struct {
	URL      string `json:"url"`       // (?)
	Name     string `json:"name"`      // (?)
	Img      string `json:"img"`       // (?)
	WatchNum string `json:"watch_num"` // (?)
	UpNum    string `json:"up_num"`    // (?)
}

// GetNewbieExamResultPermission 是答题结果中的权限信息。
type GetNewbieExamResultPermission struct {
	Reply   bool `json:"reply"`    // 是否允许评论。实测为 true
	ColorDm bool `json:"color_dm"` // 是否允许彩色弹幕。实测为 true
}
