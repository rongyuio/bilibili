package video

// 视频AI总结相关响应模型。

// GetVideoAISummaryResult 是 AI 总结接口的 data 对象。
type GetVideoAISummaryResult struct {
	Code        int                           `json:"code"`         // 返回值。-1：不支持AI摘要（敏感内容等）或其他因素导致请求异常；0：有摘要；1：无摘要（未识别到语音）
	ModelResult *GetVideoAISummaryModelResult `json:"model_result"` // 摘要内容。无摘要时可能为 null
	Stid        string                        `json:"stid"`         // 摘要 id。code=1 且该字段为 "0" 时已加入总结队列；code=1 且该字段为空时未识别到语音
	Status      int                           `json:"status"`       // 作用尚不明确
	LikeNum     int                           `json:"like_num"`     // 点赞数。默认为 0
	DislikeNum  int                           `json:"dislike_num"`  // 点踩数。默认为 0
}

// GetVideoAISummaryModelResult 是 AI 总结的摘要内容。
type GetVideoAISummaryModelResult struct {
	ResultType int                                    `json:"result_type"` // 数据类型。0：没有摘要；1：仅有摘要总结；2：摘要以及提纲
	Summary    string                                 `json:"summary"`     // 视频摘要。通常为一段概括整个视频内容的文本
	Outline    []GetVideoAISummaryModelResultOutline  `json:"outline"`     // 分段提纲。无数据时为空数组
	Subtitle   []GetVideoAISummaryModelResultSubtitle `json:"subtitle"`    // AI字幕。无数据时为空数组
}

// GetVideoAISummaryModelResultOutline 是摘要的一个分段提纲。
type GetVideoAISummaryModelResultOutline struct {
	Title       string                                           `json:"title"`        // 分段标题。段落内容的概括
	PartOutline []GetVideoAISummaryModelResultOutlinePartOutline `json:"part_outline"` // 分段要点
	Timestamp   int64                                            `json:"timestamp"`    // 分段起始时间。单位为秒
}

// GetVideoAISummaryModelResultOutlinePartOutline 是分段提纲中的一个要点。
type GetVideoAISummaryModelResultOutlinePartOutline struct {
	Timestamp int64  `json:"timestamp"` // 要点起始时间。单位为秒
	Content   string `json:"content"`   // 小结内容
}

// GetVideoAISummaryModelResultSubtitle 是一段 AI 字幕。
type GetVideoAISummaryModelResultSubtitle struct {
	PartSubtitle []GetVideoAISummaryModelResultSubtitlePartSubtitle `json:"part_subtitle"` // 字幕分段
	Timestamp    int64                                              `json:"timestamp"`     // 字幕识别起始时间。单位为秒
	Title        string                                             `json:"title"`         // 字幕标题。固定为空字符串
}

// GetVideoAISummaryModelResultSubtitlePartSubtitle 是 AI 字幕中的一个分段。
type GetVideoAISummaryModelResultSubtitlePartSubtitle struct {
	Content        string `json:"content"`         // 字幕内容
	StartTimestamp int64  `json:"start_timestamp"` // 分段开始时间。单位为秒
	EndTimestamp   int64  `json:"end_timestamp"`   // 分段结束时间。单位为秒
}
