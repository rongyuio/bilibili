package video

// 互动视频相关响应模型。

// GetInteractVideoEdgeInfoResult 是「获取互动视频模块详细信息」的 data 对象。
type GetInteractVideoEdgeInfoResult struct {
	Title          string                              `json:"title"`           // 视频模块（分P）标题
	EdgeID         int                                 `json:"edge_id"`         // 当前模块id
	StoryList      []GetInteractVideoEdgeInfoStoryItem `json:"story_list"`      // 进度回溯信息。未登录仅有起始模块
	Edges          GetInteractVideoEdgeInfoEdges       `json:"edges"`           // 当前模块信息
	Preload        GetInteractVideoEdgeInfoPreload     `json:"preload"`         // 预加载的分P
	HiddenVars     []GetInteractVideoEdgeInfoHiddenVar `json:"hidden_vars"`     // 变量列表。无变量时不存在此项
	IsLeaf         int                                 `json:"is_leaf"`         // 是否为结束模块。0：普通模块。1：结束模块
	NoTutorial     int                                 `json:"no_tutorial"`     // 禁止记录选择。1：禁止。非禁止时无此项
	NoBacktracking int                                 `json:"no_backtracking"` // 禁止进度回溯。1：禁止。非禁止时无此项
	NoEvaluation   int                                 `json:"no_evaluation"`   // 禁止结尾评分。1：禁止。非禁止时无此项
}

// GetInteractVideoEdgeInfoStoryItem 是进度回溯信息中的模块。
type GetInteractVideoEdgeInfoStoryItem struct {
	NodeID    int    `json:"node_id"`    // 模块编号
	EdgeID    int    `json:"edge_id"`    // 模块编号
	Title     string `json:"title"`      // 模块（分P）标题
	Cid       int    `json:"cid"`        // 模块（分P）cid
	StartPos  int    `json:"start_pos"`  // 记录播放开始位置。单位为毫秒
	Cover     string `json:"cover"`      // 分P封面url
	IsCurrent int    `json:"is_current"` // 是否为当前模块。1：是。仅为当前模块时存在
	Cursor    int    `json:"cursor"`     // 进度序号。从0开始向上增长
}

// GetInteractVideoEdgeInfoEdges 是当前模块信息。
type GetInteractVideoEdgeInfoEdges struct {
	Dimension GetInteractVideoEdgeInfoDimension  `json:"dimension"` // 当前分P分辨率。有部分视频无法获取分辨率
	Questions []GetInteractVideoEdgeInfoQuestion `json:"questions"` // 问题。结束模块无此项
	Skin      GetInteractVideoEdgeInfoSkin       `json:"skin"`      // 问题外观
}

// GetInteractVideoEdgeInfoDimension 是当前分P分辨率。
type GetInteractVideoEdgeInfoDimension struct {
	Width  int    `json:"width"`  // 当前分P宽度
	Height int    `json:"height"` // 当前分P高度
	Rotate int    `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
	Sar    string `json:"sar"`    // 作用尚不明确
}

// GetInteractVideoEdgeInfoQuestion 是互动视频中的问题。
type GetInteractVideoEdgeInfoQuestion struct {
	ID          int                              `json:"id"`            // 作用尚不明确
	Type        int                              `json:"type"`          // 选项显示模式。0：不显示选项。1：底部选项模式。2：坐标定点模式
	StartTimeR  int                              `json:"start_time_r"`  // 300 或 duration。作用尚不明确
	Duration    int                              `json:"duration"`      // 回答限时。单位为毫秒。不限时为-1
	PauseVideo  int                              `json:"pause_video"`   // 是否暂停播放视频。0：不暂停。1：暂停播放
	Title       string                           `json:"title"`         // 空。作用尚不明确
	Choices     []GetInteractVideoEdgeInfoChoice `json:"choices"`       // 选项列表
	FadeInTime  int                              `json:"fade_in_time"`  // 选项淡入时间。毫秒
	FadeOutTime int                              `json:"fade_out_time"` // 选项淡出时间。毫秒
}

// GetInteractVideoEdgeInfoChoice 是问题中的选项。
type GetInteractVideoEdgeInfoChoice struct {
	ID             int    `json:"id"`              // 选项所跳转的模块id
	PlatformAction string `json:"platform_action"` // 跳转信息文字。JUMP+{所跳转的模块编号}+{所跳转的cid}
	NativeAction   string `json:"native_action"`   // 点击后对变量运算语句。每项间用分号隔开。无为空
	Condition      string `json:"condition"`       // 选项出现条件判断语句。无为空
	Cid            int    `json:"cid"`             // 选项所跳转分P的cid
	X              int    `json:"x"`               // 选项出现的x坐标。仅坐标模式有此项
	Y              int    `json:"y"`               // 选项出现的y坐标。仅坐标模式有此项
	TextAlign      int    `json:"text_align"`      // 选项文本对齐方式
	Option         string `json:"option"`          // 选项文字
	Selected       any    `json:"selected"`        // 选择动画信息。文档未给出字段
	Submited       any    `json:"submited"`        // 提交动画信息。文档未给出字段
	IsDefault      int    `json:"is_default"`      // 是否为默认选项。1：是。非默认选项无此项
	IsHidden       int    `json:"is_hidden"`       // 是否为隐藏选项。1：是。非隐藏选项无此项
}

// GetInteractVideoEdgeInfoSkin 是问题外观。
type GetInteractVideoEdgeInfoSkin struct {
	ChoiceImage            string `json:"choice_image"`             // 选项组件外观图片url
	TitleTextColor         string `json:"title_text_color"`         // 文字颜色。RGBA格式
	TitleShadowColor       string `json:"title_shadow_color"`       // 文字阴影颜色。RGBA格式
	TitleShadowOffsetX     int    `json:"title_shadow_offset_x"`    // 文字阴影x偏移
	TitleShadowOffsetY     int    `json:"title_shadow_offset_y"`    // 文字阴影y偏移
	TitleShadowRadius      int    `json:"title_shadow_radius"`      // 文字阴影半径
	ProgressbarColor       string `json:"progressbar_color"`        // 倒计时条颜色。RGBA格式
	ProgressbarShadowColor string `json:"progressbar_shadow_color"` // 倒计时条阴影颜色。RGBA格式
}

// GetInteractVideoEdgeInfoPreload 是预加载的分P。
type GetInteractVideoEdgeInfoPreload struct {
	Video []GetInteractVideoEdgeInfoPreloadVideo `json:"video"` // 预加载的分P。内容为当前所有选项的跳转视频
}

// GetInteractVideoEdgeInfoPreloadVideo 是预加载的分P。
type GetInteractVideoEdgeInfoPreloadVideo struct {
	Aid int `json:"aid"` // 稿件avid
	Cid int `json:"cid"` // 分P cid
}

// GetInteractVideoEdgeInfoHiddenVar 是互动视频变量。
type GetInteractVideoEdgeInfoHiddenVar struct {
	Value         float64 `json:"value"`          // 变量值。随机值为随机整数
	ID            string  `json:"id"`             // 变量编号
	IDV2          string  `json:"id_v2"`          // 变量编号。语句中一般使用这种
	Type          int     `json:"type"`           // 变量类型。1：普通变量。2：随机值
	IsShow        int     `json:"is_show"`        // 是否展示变量。0：否。1：是
	Name          string  `json:"name"`           // 变量名
	SkipOverwrite int     `json:"skip_overwrite"` // 0。作用尚不明确
}
