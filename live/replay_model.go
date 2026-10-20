package live

// 直播回放相关响应模型。
//
// 每个接口的 data 与其中嵌套的对象各自定义，即使字段相同也不复用，
// 以免某个接口的返回体变更影响其他接口的调用方。

// GetLiveReplayListResult 是获取自己的直播回放列表的 data 对象。
type GetLiveReplayListResult struct {
	ReplayInfo           []GetLiveReplayListItem     `json:"replay_info"`             // 回放信息列表；无结果时为 null
	Pagination           GetLiveReplayListPagination `json:"pagination"`              // 分页信息
	ArchiveFlag          bool                        `json:"archive_flag"`            // (?)。作用尚不明确
	CanEdit              int                         `json:"can_edit"`                // (?)。作用尚不明确
	CanUpload            bool                        `json:"can_upload"`              // (?)。作用尚不明确
	HasThirdPlatformLive bool                        `json:"has_third_platform_live"` // (?)。作用尚不明确
}

// GetLiveReplayListPagination 是获取自己的直播回放列表的分页信息。
type GetLiveReplayListPagination struct {
	Page     int `json:"page"`      // 请求的页码
	PageSize int `json:"page_size"` // 内容数量
	Total    int `json:"total"`     // 总计内容数量
}

// GetLiveReplayListItem 是一场直播回放的信息。
type GetLiveReplayListItem struct {
	ReplayID  int64                          `json:"replay_id"`  // 直播回放 id
	LiveInfo  GetLiveReplayListItemLiveInfo  `json:"live_info"`  // 直播信息
	VideoInfo GetLiveReplayListItemVideoInfo `json:"video_info"` // 回放视频信息
	AlarmInfo GetLiveReplayListItemAlarmInfo `json:"alarm_info"` // 警报信息
	RoomID    int64                          `json:"room_id"`    // 直播间 id
	LiveKey   string                         `json:"live_key"`   // 标记直播场次的 key
	StartTime int64                          `json:"start_time"` // 直播开始秒时间戳；即调用「开始直播」接口的时间
	EndTime   int64                          `json:"end_time"`   // 直播结束秒时间戳；即调用「关闭直播」接口的时间
}

// GetLiveReplayListItemLiveInfo 是回放对应的直播信息。
type GetLiveReplayListItemLiveInfo struct {
	Title    string `json:"title"`     // 直播标题；直播结束时的标题
	Cover    string `json:"cover"`     // 直播封面
	LiveTime int64  `json:"live_time"` // 直播时间；同 GetLiveReplayListItem.StartTime
	LiveType int    `json:"live_type"` // 直播类型?。作用尚不明确
	Platform string `json:"platform"`  // 直播平台
}

// GetLiveReplayListItemVideoInfo 是回放视频的信息。
type GetLiveReplayListItemVideoInfo struct {
	ReplayStatus  int    `json:"replay_status"`  // 回放状态。作用尚不明确
	EstimatedTime string `json:"estimated_time"` // 直播回放合成结束时间；未合成时为 "1970-01-01 08:00:00"
	Duration      int64  `json:"duration"`       // 直播时长。单位秒
	DownloadURL   string `json:"download_url"`   // 下载链接片段；整场直播回放合成成功时存在
	AlertCode     int    `json:"alert_code"`     // 快速检查警告代码；整场直播回放合成失败时不存在
	AlertMessage  string `json:"alert_message"`  // 快速检查警告信息；整场直播回放合成失败时不存在
}

// GetLiveReplayListItemAlarmInfo 是回放合成的警报信息。
type GetLiveReplayListItemAlarmInfo struct {
	Code         int    `json:"code"`           // 回放合成警报代码
	Message      string `json:"message"`        // 回放合成错误信息
	CurTime      int64  `json:"cur_time"`       // 当前时间戳。Unix 秒时间戳
	IsBanPublish bool   `json:"is_ban_publish"` // 是否禁止发布?
}

// GetOtherLiveReplayListResult 是获取某位主播回放列表的 data 对象。
type GetOtherLiveReplayListResult struct {
	ReplayInfo           []GetOtherLiveReplayListItem     `json:"replay_info"`             // 回放信息列表；无结果时为 null
	Pagination           GetOtherLiveReplayListPagination `json:"pagination"`              // 分页信息
	ArchiveFlag          bool                             `json:"archive_flag"`            // (?)。作用尚不明确
	CanEdit              int                              `json:"can_edit"`                // (?)。作用尚不明确
	CanUpload            bool                             `json:"can_upload"`              // (?)。作用尚不明确
	HasThirdPlatformLive bool                             `json:"has_third_platform_live"` // (?)。作用尚不明确
}

// GetOtherLiveReplayListPagination 是获取某位主播回放列表的分页信息。
type GetOtherLiveReplayListPagination struct {
	Page     int `json:"page"`      // 请求的页码
	PageSize int `json:"page_size"` // 内容数量
	Total    int `json:"total"`     // 总计内容数量
}

// GetOtherLiveReplayListItem 是某位主播的一场直播回放信息。
type GetOtherLiveReplayListItem struct {
	ReplayID  int64                               `json:"replay_id"`  // 直播回放 id
	LiveInfo  GetOtherLiveReplayListItemLiveInfo  `json:"live_info"`  // 直播信息
	VideoInfo GetOtherLiveReplayListItemVideoInfo `json:"video_info"` // 回放视频信息
	AlarmInfo GetOtherLiveReplayListItemAlarmInfo `json:"alarm_info"` // 警报信息
	RoomID    int64                               `json:"room_id"`    // 直播间 id
	LiveKey   string                              `json:"live_key"`   // 标记直播场次的 key
	StartTime int64                               `json:"start_time"` // 直播开始秒时间戳；即调用「开始直播」接口的时间
	EndTime   int64                               `json:"end_time"`   // 直播结束秒时间戳；即调用「关闭直播」接口的时间
}

// GetOtherLiveReplayListItemLiveInfo 是回放对应的直播信息。
type GetOtherLiveReplayListItemLiveInfo struct {
	Title    string `json:"title"`     // 直播标题；直播结束时的标题
	Cover    string `json:"cover"`     // 直播封面
	LiveTime int64  `json:"live_time"` // 直播时间；同 GetOtherLiveReplayListItem.StartTime
	LiveType int    `json:"live_type"` // 直播类型?。作用尚不明确
	Platform string `json:"platform"`  // 直播平台
}

// GetOtherLiveReplayListItemVideoInfo 是回放视频的信息。
type GetOtherLiveReplayListItemVideoInfo struct {
	ReplayStatus  int    `json:"replay_status"`  // 回放状态。作用尚不明确
	EstimatedTime string `json:"estimated_time"` // 直播回放合成结束时间；未合成时为 "1970-01-01 08:00:00"
	Duration      int64  `json:"duration"`       // 直播时长。单位秒
	DownloadURL   string `json:"download_url"`   // 下载链接片段；整场直播回放合成成功时存在
	AlertCode     int    `json:"alert_code"`     // 快速检查警告代码；整场直播回放合成失败时不存在
	AlertMessage  string `json:"alert_message"`  // 快速检查警告信息；整场直播回放合成失败时不存在
}

// GetOtherLiveReplayListItemAlarmInfo 是回放合成的警报信息。
type GetOtherLiveReplayListItemAlarmInfo struct {
	Code         int    `json:"code"`           // 回放合成警报代码
	Message      string `json:"message"`        // 回放合成错误信息
	CurTime      int64  `json:"cur_time"`       // 当前时间戳。Unix 秒时间戳
	IsBanPublish bool   `json:"is_ban_publish"` // 是否禁止发布?
}

// GetLivePublishedSliceListResult 是自己已发布的回放片段列表（data 部分）。
type GetLivePublishedSliceListResult struct {
	List     []GetLivePublishedSliceListItem `json:"list"`      // 切片信息
	Page     int                             `json:"page"`      // 请求的页码
	PageSize int                             `json:"page_size"` // 内容数量
	Total    int                             `json:"total"`     // 总计内容数量
}

// GetLivePublishedSliceListItem 是自己已发布的一个回放片段。
type GetLivePublishedSliceListItem struct {
	SliceID      int64  `json:"slice_id"`      // 切片 id。文档字段名写作 silce_id，实际响应为 slice_id
	AvTitle      string `json:"av_title"`      // 切片标题
	AvCover      string `json:"av_cover"`      // 切片封面
	AvStatus     int    `json:"av_status"`     // 切片状态。1：发布中。2：已投稿。3：投稿失败
	Avid         int64  `json:"avid"`          // 切片视频的 avid。状态为 2 时存在
	Ctime        string `json:"ctime"`         // 切片创建时间
	StartTm      string `json:"start_tm"`      // 切片开始时间
	EndTm        string `json:"end_tm"`        // 切片结束时间
	AvDuration   int64  `json:"av_duration"`   // 切片时长。状态为 2 且创作中心出现有效视频时长时存在
	FailedReason string `json:"failed_reason"` // 失败原因。状态为 3 时存在，2024-09-01 前发布失败的切片可能不存在
	LiveType     int    `json:"live_type"`     // (?)。作用尚不明确
	CntPlay      int64  `json:"cnt_play"`      // 播放数。视频有播放时存在；若该页出现任意状态不为 2 的项也会不存在
	CntDanmaku   int64  `json:"cnt_danmaku"`   // 弹幕数。视频有弹幕时存在；若该页出现任意状态不为 2 的项也会不存在
}

// GetOtherPublishedSliceListResult 是你为某位主播剪辑的已发布片段列表（data 部分）。
type GetOtherPublishedSliceListResult struct {
	SliceInfo  []GetOtherPublishedSliceListItem     `json:"slice_info"` // 切片信息；无内容时为 null
	Pagination GetOtherPublishedSliceListPagination `json:"pagination"` // 分页信息
}

// GetOtherPublishedSliceListPagination 是你为某位主播剪辑的已发布片段列表的分页信息。
type GetOtherPublishedSliceListPagination struct {
	Page     int `json:"page"`      // 请求的页码
	PageSize int `json:"page_size"` // 内容数量
	Total    int `json:"total"`     // 总计内容数量
}

// GetOtherPublishedSliceListItem 是你为某位主播剪辑的一个已发布片段。
type GetOtherPublishedSliceListItem struct {
	SliceID     int64  `json:"slice_id"`      // 切片 id
	UID         int64  `json:"uid"`           // 发布者 uid?
	LiveUID     int64  `json:"live_uid"`      // 主播 uid
	LiveKey     string `json:"live_key"`      // 标记直播场次的 key
	Title       string `json:"title"`         // 切片标题
	Cover       string `json:"cover"`         // 切片封面。投稿时未传递封面将无此字段
	HighLightID int64  `json:"high_light_id"` // 绑定的高光时刻。投稿时提供相关参数时存在
	StartTime   string `json:"start_time"`    // 片段开始时间
	EndTime     string `json:"end_time"`      // 片段结束时间
	Status      int    `json:"status"`        // 切片状态。参见 GetLivePublishedSliceListItem.AvStatus
	FailReason  string `json:"fail_reason"`   // 失败提示。状态为 3 且有提示信息时存在
	Filename    string `json:"filename"`      // 切片视频文件名。内部使用；视频合成完成时存在
	Avid        int64  `json:"avid"`          // 切片视频的 avid。状态为 2 时存在
	Ctime       string `json:"ctime"`         // 切片创建时间
	AvDuration  int64  `json:"av_duration"`   // 切片时长。状态为 2 且创作中心出现有效视频时长时存在
	LiveType    int    `json:"live_type"`     // (?)。作用尚不明确
}

// GetLiveReplayDraftListResult 是回放剪辑草稿列表（data 部分）。
type GetLiveReplayDraftListResult struct {
	DraftInfo  []GetLiveReplayDraftListItem     `json:"draft_info"` // 草稿信息；无结果时为 null
	Pagination GetLiveReplayDraftListPagination `json:"pagination"` // 分页信息
}

// GetLiveReplayDraftListPagination 是回放剪辑草稿列表的分页信息。
type GetLiveReplayDraftListPagination struct {
	Page     int `json:"page"`      // 请求的页码
	PageSize int `json:"page_size"` // 内容数量
	Total    int `json:"total"`     // 总计内容数量
}

// GetLiveReplayDraftListItem 是一条回放剪辑草稿。
type GetLiveReplayDraftListItem struct {
	ID            int64  `json:"id"`              // 回放剪辑 id
	Title         string `json:"title"`           // 回放剪辑标题
	Cover         string `json:"cover"`           // 回放剪辑封面；有封面时存在
	LiveKey       string `json:"live_key"`        // 标记直播场次的 key
	Ctime         string `json:"ctime"`           // 回放剪辑创建时间
	LiveStartTime string `json:"live_start_time"` // 直播开始时间
	LiveEndTime   string `json:"live_end_time"`   // 直播结束时间
	LiveType      int    `json:"live_type"`       // (?)。作用尚不明确
}

// GetOtherReplayDraftListResult 是你为某位主播剪辑的回放草稿列表（data 部分）。
type GetOtherReplayDraftListResult struct {
	DraftInfo  []GetOtherReplayDraftListItem     `json:"draft_info"` // 草稿信息；无结果时为 null
	Pagination GetOtherReplayDraftListPagination `json:"pagination"` // 分页信息
}

// GetOtherReplayDraftListPagination 是你为某位主播剪辑的回放草稿列表的分页信息。
type GetOtherReplayDraftListPagination struct {
	Page     int `json:"page"`      // 请求的页码
	PageSize int `json:"page_size"` // 内容数量
	Total    int `json:"total"`     // 总计内容数量
}

// GetOtherReplayDraftListItem 是你为某位主播剪辑的一条回放草稿。
type GetOtherReplayDraftListItem struct {
	ID            int64  `json:"id"`              // 回放剪辑 id
	Title         string `json:"title"`           // 回放剪辑标题
	Cover         string `json:"cover"`           // 回放剪辑封面；有封面时存在
	LiveKey       string `json:"live_key"`        // 标记直播场次的 key
	Ctime         string `json:"ctime"`           // 回放剪辑创建时间
	LiveStartTime string `json:"live_start_time"` // 直播开始时间
	LiveEndTime   string `json:"live_end_time"`   // 直播结束时间
	LiveType      int    `json:"live_type"`       // (?)。作用尚不明确
}

// GetLiveReplayInfoResult 是指定回放的信息（data 部分）。
type GetLiveReplayInfoResult struct {
	List []GetLiveReplayInfoItem `json:"list"` // 回放信息列表
}

// GetLiveReplayInfoItem 是一条回放的基础信息。
type GetLiveReplayInfoItem struct {
	UID       int64  `json:"uid"`        // 用户 mid
	RecordID  int64  `json:"record_id"`  // 直播回放 id
	Title     string `json:"title"`      // 直播标题
	Cover     string `json:"cover"`      // 直播封面
	Status    int    `json:"status"`     // 回放状态
	StartTime int64  `json:"start_time"` // 直播开始时间。Unix 秒时间戳
	EndTime   int64  `json:"end_time"`   // 直播结束时间。Unix 秒时间戳
}

// GetLiveSliceStreamResult 是自己某场直播的切片视频流（data 部分）。
type GetLiveSliceStreamResult struct {
	List    []GetLiveSliceStreamItem    `json:"list"`     // 直播回放视频列表；该场回放没有视频流时为 null
	BanList []GetLiveSliceStreamBanItem `json:"ban_list"` // 不可发布的回放时间；没有不可发布的时间时为 null
}

// GetLiveSliceStreamItem 是一段回放视频流。
type GetLiveSliceStreamItem struct {
	StartTime int64  `json:"start_time"` // 片段开始时间戳。Unix 秒时间戳
	EndTime   int64  `json:"end_time"`   // 片段结束时间戳。Unix 秒时间戳
	Stream    string `json:"stream"`     // 直播回放视频流
	Type      int    `json:"type"`       // 类型?。2：一般回放?
}

// GetLiveSliceStreamBanItem 是一段不可发布的回放时间。
type GetLiveSliceStreamBanItem struct {
	StartTime int64 `json:"start_time"` // 不可发布片段的开始时间戳。Unix 秒时间戳
	EndTime   int64 `json:"end_time"`   // 不可发布片段的结束时间戳。Unix 秒时间戳
}

// GetOtherLiveSliceStreamResult 是某位主播某场直播的切片视频流（data 部分）。
type GetOtherLiveSliceStreamResult struct {
	List []GetOtherLiveSliceStreamItem `json:"list"` // 直播回放视频列表
}

// GetOtherLiveSliceStreamItem 是某位主播某场直播的一段回放视频流。
type GetOtherLiveSliceStreamItem struct {
	StartTime int64  `json:"start_time"` // 片段开始时间戳。Unix 秒时间戳
	EndTime   int64  `json:"end_time"`   // 片段结束时间戳。Unix 秒时间戳
	Stream    string `json:"stream"`     // 直播回放视频流
	Type      int    `json:"type"`       // 类型?。2：一般回放?
}

// GetLiveReplaySessionDataResult 是一场直播的会话数据（data 部分）。
type GetLiveReplaySessionDataResult struct {
	SessionData   []GetLiveReplaySessionDataSample      `json:"session_data"`    // 会话数据
	MaxDanmaku    int64                                 `json:"max_danmaku"`     // 弹幕最多的时间戳。Unix 秒时间戳，没有则为 0
	MaxPCU        int64                                 `json:"max_pcu"`         // 进房最多的时间戳。Unix 秒时间戳，没有则为 0
	MaxValue      int                                   `json:"max_value"`       // (?)。效果未知
	HighLightData []GetLiveReplaySessionDataHighLight   `json:"high_light_data"` // 高光时刻数据
	AssURL        string                                `json:"ass_url"`         // ASS 字幕链接，用作弹幕显示
	HighLightStat GetLiveReplaySessionDataHighLightStat `json:"high_light_stat"` // 高亮状态信息
	ExtraList     any                                   `json:"extra_list"`      // (?)。实测为 null
}

// GetLiveReplaySessionDataSample 是一次会话数据采样。
type GetLiveReplaySessionDataSample struct {
	TS    int64 `json:"ts"`    // 采样时间。Unix 秒时间戳
	Value int   `json:"value"` // 弹幕数量
}

// GetLiveReplaySessionDataHighLight 是一个高光时刻。
type GetLiveReplaySessionDataHighLight struct {
	ID        int64  `json:"id"`         // 高光 id
	Type      int    `json:"type"`       // 高光类型。1：弹幕。2：进房
	StartTime int64  `json:"start_time"` // 高光开始时间戳。Unix 秒时间戳
	EndTime   int64  `json:"end_time"`   // 高光结束时间戳。Unix 秒时间戳
	Title     string `json:"title"`      // 高光提示标题
	Cover     string `json:"cover"`      // (?)。目前为 ""
	Extra     string `json:"extra"`      // (?)。目前为 ""
}

// GetLiveReplaySessionDataHighLightStat 是高亮状态信息。
type GetLiveReplaySessionDataHighLightStat struct {
	OriginalHL int `json:"original_hl"` // 「高光时刻」数量；有「高光时刻」时存在
	ManualHL   int `json:"manual_hl"`   // 「已保存」数量；有「已保存」时存在
}

// GetOtherLiveReplaySessionDataResult 是某位主播某场直播的会话数据（data 部分）。
type GetOtherLiveReplaySessionDataResult struct {
	SessionData   []GetOtherLiveReplaySessionDataSample      `json:"session_data"`    // 会话数据
	MaxDanmaku    int64                                      `json:"max_danmaku"`     // 弹幕最多的时间戳。Unix 秒时间戳，没有则为 0
	MaxPCU        int64                                      `json:"max_pcu"`         // 进房最多的时间戳。Unix 秒时间戳，没有则为 0
	MaxValue      int                                        `json:"max_value"`       // (?)。效果未知
	HighLightData []GetOtherLiveReplaySessionDataHighLight   `json:"high_light_data"` // 高光时刻数据
	AssURL        string                                     `json:"ass_url"`         // ASS 字幕链接，用作弹幕显示
	HighLightStat GetOtherLiveReplaySessionDataHighLightStat `json:"high_light_stat"` // 高亮状态信息
	ExtraList     any                                        `json:"extra_list"`      // (?)。实测为 null
}

// GetOtherLiveReplaySessionDataSample 是一次会话数据采样。
type GetOtherLiveReplaySessionDataSample struct {
	TS    int64 `json:"ts"`    // 采样时间。Unix 秒时间戳
	Value int   `json:"value"` // 弹幕数量
}

// GetOtherLiveReplaySessionDataHighLight 是一个高光时刻。
type GetOtherLiveReplaySessionDataHighLight struct {
	ID        int64  `json:"id"`         // 高光 id
	Type      int    `json:"type"`       // 高光类型。1：弹幕。2：进房
	StartTime int64  `json:"start_time"` // 高光开始时间戳。Unix 秒时间戳
	EndTime   int64  `json:"end_time"`   // 高光结束时间戳。Unix 秒时间戳
	Title     string `json:"title"`      // 高光提示标题
	Cover     string `json:"cover"`      // (?)。目前为 ""
	Extra     string `json:"extra"`      // (?)。目前为 ""
}

// GetOtherLiveReplaySessionDataHighLightStat 是高亮状态信息。
type GetOtherLiveReplaySessionDataHighLightStat struct {
	OriginalHL int `json:"original_hl"` // 「高光时刻」数量；有「高光时刻」时存在
	ManualHL   int `json:"manual_hl"`   // 「已保存」数量；有「已保存」时存在
}

// GetLiveManualHighlightListResult 是某场直播的用户高光列表（data 部分）。
type GetLiveManualHighlightListResult struct {
	HighLightStat       GetLiveManualHighlightListHighLightStat `json:"high_light_stat"`       // 高亮状态信息
	ManualHighlightList []GetLiveManualHighlightListItem        `json:"manual_highlight_list"` // 高光列表；有内容时存在
}

// GetLiveManualHighlightListHighLightStat 是高亮状态信息。
type GetLiveManualHighlightListHighLightStat struct {
	OriginalHL int `json:"original_hl"` // 「高光时刻」数量；有「高光时刻」时存在
	ManualHL   int `json:"manual_hl"`   // 「已保存」数量；有「已保存」时存在
}

// GetLiveManualHighlightListItem 是一条手动保存的高光片段。
type GetLiveManualHighlightListItem struct {
	ManualID       int64  `json:"manual_id"`       // 片段 id
	HighlightTitle string `json:"highlight_title"` // 片段标题
	StartTS        int64  `json:"start_ts"`        // 开始时间戳。Unix 秒时间戳
	EndTS          int64  `json:"end_ts"`          // 结束时间戳。Unix 秒时间戳
	UID            int64  `json:"uid"`             // 保存者 uid?
	RUID           int64  `json:"ruid"`            // 主播 uid
	LiveKey        string `json:"live_key"`        // 标记直播场次的 key
}
