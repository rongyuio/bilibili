package newbieexam

// 新用户答题题目拉取相关响应模型。
//
// 基础题、附加题、自选题三个接口的 data 结构相同，但按本包「返回体不复用」
// 的约定，各接口的模型独立定义，互不共享类型。

// GetNewbieExamBaseQuestionResult 是拉取基础题的响应 data。
type GetNewbieExamBaseQuestionResult struct {
	Question *GetNewbieExamBaseQuestionQuestion `json:"question"` // 题目内容
}

// GetNewbieExamBaseQuestionQuestion 是基础题的题目内容。
type GetNewbieExamBaseQuestionQuestion struct {
	ID          int                               `json:"id"`           // 问题id。问题id不代表题图，因为选项是打乱的
	Number      int                               `json:"number"`       // 当前题号
	QHeight     float64                           `json:"q_height"`     // 问题部分总高度
	QCoordY     float64                           `json:"q_coord_y"`    // 问题部分Y裁剪起始。当然是0
	Image       string                            `json:"image"`        // 题图url。题目文字+所有的选项文字排版成一张图，存储在bfs中
	From        string                            `json:"from"`         // 问题来源页面url。如「xx弹幕是否违规」所在的视频页
	Options     []GetNewbieExamBaseQuestionOption `json:"options"`      // 选项列表。至少2个
	TypeID      int                               `json:"type_id"`      // 题目父类型id。36为基础题
	TypeName    string                            `json:"type_name"`    // 分院小电视提示文案。如：xx小电视
	TypeImage   string                            `json:"type_image"`   // 分院小电视图标url
	DecisionCtx map[string]any                    `json:"decision_ctx"` // 决策上下文。实测为空对象
	Question    *GetNewbieExamBaseQuestionContent `json:"question"`     // 题目文字与答案列表。接口文档给出了该嵌套对象，但响应示例中未出现
}

// GetNewbieExamBaseQuestionOption 是基础题的题目选项。
type GetNewbieExamBaseQuestionOption struct {
	Number int     `json:"number"`  // 选项序号。如：1为A 2为B
	High   float64 `json:"high"`    // 选项部分总高度
	CoordY float64 `json:"coord_y"` // 选项部分Y裁剪起始
	Hash   string  `json:"hash"`    // 选项hash。作为提交对应选项时应传参
}

// GetNewbieExamBaseQuestionContent 是基础题的题目文字与答案。
type GetNewbieExamBaseQuestionContent struct {
	Title string                            `json:"title"` // 问题标题。实测为加密字符串
	Ans   []GetNewbieExamBaseQuestionAnswer `json:"ans"`   // 答案列表
}

// GetNewbieExamBaseQuestionAnswer 是基础题的题目答案。
type GetNewbieExamBaseQuestionAnswer struct {
	Key   string `json:"key"`   // 选项hash。实测为 md5 字符串
	Title string `json:"title"` // 选项内容。实测为加密字符串
}

// GetNewbieExamExtraQuestionResult 是拉取附加题的响应 data。
type GetNewbieExamExtraQuestionResult struct {
	Question *GetNewbieExamExtraQuestionQuestion `json:"question"` // 题目内容
}

// GetNewbieExamExtraQuestionQuestion 是附加题的题目内容。
type GetNewbieExamExtraQuestionQuestion struct {
	ID          int                                `json:"id"`           // 问题id。问题id不代表题图，因为选项是打乱的
	Number      int                                `json:"number"`       // 当前题号
	QHeight     float64                            `json:"q_height"`     // 问题部分总高度
	QCoordY     float64                            `json:"q_coord_y"`    // 问题部分Y裁剪起始。当然是0
	Image       string                             `json:"image"`        // 题图url。题目文字+所有的选项文字排版成一张图，存储在bfs中
	From        string                             `json:"from"`         // 问题来源页面url。如「xx弹幕是否违规」所在的视频页
	Options     []GetNewbieExamExtraQuestionOption `json:"options"`      // 选项列表。至少2个
	TypeID      int                                `json:"type_id"`      // 题目父类型id。36为基础题
	TypeName    string                             `json:"type_name"`    // 分院小电视提示文案。如：xx小电视
	TypeImage   string                             `json:"type_image"`   // 分院小电视图标url
	DecisionCtx map[string]any                     `json:"decision_ctx"` // 决策上下文。实测为空对象
	Question    *GetNewbieExamExtraQuestionContent `json:"question"`     // 题目文字与答案列表。接口文档给出了该嵌套对象，但响应示例中未出现
}

// GetNewbieExamExtraQuestionOption 是附加题的题目选项。
type GetNewbieExamExtraQuestionOption struct {
	Number int     `json:"number"`  // 选项序号。如：1为A 2为B
	High   float64 `json:"high"`    // 选项部分总高度
	CoordY float64 `json:"coord_y"` // 选项部分Y裁剪起始
	Hash   string  `json:"hash"`    // 选项hash。作为提交对应选项时应传参
}

// GetNewbieExamExtraQuestionContent 是附加题的题目文字与答案。
type GetNewbieExamExtraQuestionContent struct {
	Title string                             `json:"title"` // 问题标题。实测为加密字符串
	Ans   []GetNewbieExamExtraQuestionAnswer `json:"ans"`   // 答案列表
}

// GetNewbieExamExtraQuestionAnswer 是附加题的题目答案。
type GetNewbieExamExtraQuestionAnswer struct {
	Key   string `json:"key"`   // 选项hash。实测为 md5 字符串
	Title string `json:"title"` // 选项内容。实测为加密字符串
}

// GetNewbieExamProQuestionResult 是拉取自选题的响应 data。
type GetNewbieExamProQuestionResult struct {
	Question *GetNewbieExamProQuestionQuestion `json:"question"` // 题目内容
}

// GetNewbieExamProQuestionQuestion 是自选题的题目内容。
type GetNewbieExamProQuestionQuestion struct {
	ID          int                              `json:"id"`           // 问题id。问题id不代表题图，因为选项是打乱的
	Number      int                              `json:"number"`       // 当前题号
	QHeight     float64                          `json:"q_height"`     // 问题部分总高度
	QCoordY     float64                          `json:"q_coord_y"`    // 问题部分Y裁剪起始。当然是0
	Image       string                           `json:"image"`        // 题图url。题目文字+所有的选项文字排版成一张图，存储在bfs中
	From        string                           `json:"from"`         // 问题来源页面url。如「xx弹幕是否违规」所在的视频页
	Options     []GetNewbieExamProQuestionOption `json:"options"`      // 选项列表。至少2个
	TypeID      int                              `json:"type_id"`      // 题目父类型id。36为基础题
	TypeName    string                           `json:"type_name"`    // 分院小电视提示文案。如：xx小电视
	TypeImage   string                           `json:"type_image"`   // 分院小电视图标url
	DecisionCtx map[string]any                   `json:"decision_ctx"` // 决策上下文。实测为空对象
	Question    *GetNewbieExamProQuestionContent `json:"question"`     // 题目文字与答案列表。接口文档给出了该嵌套对象，但响应示例中未出现
}

// GetNewbieExamProQuestionOption 是自选题的题目选项。
type GetNewbieExamProQuestionOption struct {
	Number int     `json:"number"`  // 选项序号。如：1为A 2为B
	High   float64 `json:"high"`    // 选项部分总高度
	CoordY float64 `json:"coord_y"` // 选项部分Y裁剪起始
	Hash   string  `json:"hash"`    // 选项hash。作为提交对应选项时应传参
}

// GetNewbieExamProQuestionContent 是自选题的题目文字与答案。
type GetNewbieExamProQuestionContent struct {
	Title string                           `json:"title"` // 问题标题。实测为加密字符串
	Ans   []GetNewbieExamProQuestionAnswer `json:"ans"`   // 答案列表
}

// GetNewbieExamProQuestionAnswer 是自选题的题目答案。
type GetNewbieExamProQuestionAnswer struct {
	Key   string `json:"key"`   // 选项hash。实测为 md5 字符串
	Title string `json:"title"` // 选项内容。实测为加密字符串
}
