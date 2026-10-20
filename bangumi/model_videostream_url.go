package bangumi

// GetBangumiVideoStream 接口的响应模型。
//
// 本文件的类型只服务于 GetBangumiVideoStream，不与其它接口共用；
// durl / dash 及其嵌套对象也各自独立定义，不复用根包或其它接口的类型。

// 番剧视频流 URL 相关响应模型。
//
// 该接口的响应信封为 code/message/result（而非 code/message/data），
// 数据本体对应 GetBangumiVideoStreamResult。

// GetBangumiVideoStreamResult 描述番剧视频流接口 result 字段的数据本体。
type GetBangumiVideoStreamResult struct {
	Code              int                                  `json:"code"`               // 0。作用尚不明确
	IsPreview         int                                  `json:"is_preview"`         // 0。作用尚不明确
	Fnver             int                                  `json:"fnver"`              // 请求时提供的 fnver
	Fnval             int                                  `json:"fnval"`              // 请求时提供的 fnval
	VideoProject      bool                                 `json:"video_project"`      // true
	Type              string                               `json:"type"`               // 视频流类型。DASH / FLV / MP4
	Bp                int                                  `json:"bp"`                 // 是否可以承包。0
	VipType           int                                  `json:"vip_type"`           // 当前用户大会员类型
	VipStatus         int                                  `json:"vip_status"`         // 当前用户大会员状态
	IsDrm             bool                                 `json:"is_drm"`             // 是否为 DRM 内容。false
	NoRexcode         int                                  `json:"no_rexcode"`         // 0
	HasPaid           bool                                 `json:"has_paid"`           // 是否已付费。false
	Status            int                                  `json:"status"`             // 2
	From              string                               `json:"from"`               // local。作用尚不明确
	Result            string                               `json:"result"`             // suee。作用尚不明确
	Message           string                               `json:"message"`            // 空。作用尚不明确
	Quality           int                                  `json:"quality"`            // 当前的视频分辨率代码。含义见 [上表](#qn视频清晰度标识)
	Format            string                               `json:"format"`             // 视频格式
	TimeLength        int                                  `json:"timelength"`         // 视频长度。单位为毫秒，不同分辨率 / 格式可能有略微差异
	AcceptFormat      string                               `json:"accept_format"`      // 视频支持的全部格式。每项用 , 分隔
	AcceptDescription []string                             `json:"accept_description"` // 视频支持的分辨率列表（文字说明）
	AcceptQuality     []int                                `json:"accept_quality"`     // 视频支持的分辨率代码列表。含义见 [上表](#qn视频清晰度标识)
	VideoCodecID      int                                  `json:"video_codecid"`      // 默认选择视频流的编码 id。见「视频编码代码」
	SeekParam         string                               `json:"seek_param"`         // 固定值 start。作用尚不明确
	SeekType          string                               `json:"seek_type"`          // offset（dash、flv）/ second（mp4）。作用尚不明确
	Durl              []GetBangumiVideoStreamDurl          `json:"durl"`               // 视频分段。仅 flv / mp4 存在此项
	Dash              GetBangumiVideoStreamDash            `json:"dash"`               // DASH 音视频流信息。仅 dash 存在此项
	SupportFormats    []GetBangumiVideoStreamSupportFormat `json:"support_formats"`    // 支持格式的详细信息
	ClipInfoList      []GetBangumiVideoStreamClipInfo      `json:"clip_info_list"`     // 跳过片段列表
	RecordInfo        GetBangumiVideoStreamRecordInfo      `json:"record_info"`        // 备案登记信息
	CanWatchReason    int                                  `json:"can_watch_reason"`   // 可观看原因码。实测为 0
	Durls             []GetBangumiVideoStreamDurl          `json:"durls"`              // 视频分段列表（新版）。有效时为数组，无效时为 null，实测为空数组
	DyeID             string                               `json:"dyeid"`              // 染色 id。实测为空
	ErrorCode         int                                  `json:"error_code"`         // 错误码。实测为 0
}

// GetBangumiVideoStreamSupportFormat 描述 support_formats 数组中的一项。
type GetBangumiVideoStreamSupportFormat struct {
	Quality          int      `json:"quality"`             // 视频清晰度代码。含义见 [上表](#qn视频清晰度标识)
	Format           string   `json:"format"`              // 视频格式
	NewDescription   string   `json:"new_description"`     // 格式描述
	Description      string   `json:"description"`         // 格式描述
	DisplayDesc      string   `json:"display_desc"`        // 格式描述
	Superscript      string   `json:"superscript"`         // (?)
	Codecs           []string `json:"codecs"`              // 编码格式列表。如 avc1.640032（AVC 编码）、hev1.1.6.L153.90（HEVC 编码）
	NeedLogin        bool     `json:"need_login"`          // 需要登录
	NeedVip          bool     `json:"need_vip"`            // 需要大会员
	LimitWatchReason int      `json:"limit_watch_reason"`  // 限制观看原因码。实测为 0
	ReportParams     string   `json:"report_params"`       // 上报参数。JSON 字符串，实测为 "{}"
	SubDescription   string   `json:"sub_description"`     // 副描述。实测为空
	HasPreview       bool     `json:"has_preview"`         // 是否有预览。实测为 false
	CanWatchQnReason int      `json:"can_watch_qn_reason"` // 可观看清晰度原因码。实测为 0
}

// GetBangumiVideoStreamClipInfo 描述 clip_info_list 数组中的一项，即一个可跳过的片段。
type GetBangumiVideoStreamClipInfo struct {
	MaterialNo int    `json:"materialNo"` // 素材编号。实测为 0
	Start      int    `json:"start"`      // 开始时间。单位秒，实测为 65 / 1294
	End        int    `json:"end"`        // 结束时间。单位秒，实测为 156 / 1394
	ToastText  string `json:"toastText"`  // 提示文本。实测为「即将跳过片头」/「即将跳过片尾」
	ClipType   string `json:"clipType"`   // 片段类型。CLIP_TYPE_OP：片头；CLIP_TYPE_ED：片尾
}

// GetBangumiVideoStreamRecordInfo 描述备案登记信息。
type GetBangumiVideoStreamRecordInfo struct {
	RecordIcon string `json:"record_icon"` // 空串
	Record     string `json:"record"`      // 显示文案。如「登记号：10417060172092207」
}

// GetBangumiVideoStreamDurl 是 durl / durls 数组中的对象，即一个视频分段。
type GetBangumiVideoStreamDurl struct {
	Order     int      `json:"order"`      // 视频分段序号。某些视频会分为多个片段（从1顺序增长）
	Length    int      `json:"length"`     // 视频长度。单位为毫秒
	Size      int      `json:"size"`       // 视频大小。单位为 Byte
	Ahead     string   `json:"ahead"`      // （？）
	Vhead     string   `json:"vhead"`      // （？）
	URL       string   `json:"url"`        // 默认流 URL。注意 unicode 转义符。有效时间为120min
	BackupURL []string `json:"backup_url"` // 备用视频流 注意 unicode 转义符。有效时间为120min
}

// GetBangumiVideoStreamDash 是 dash 对象，即 DASH 音视频流信息。
type GetBangumiVideoStreamDash struct {
	Duration           int                                 `json:"duration"`        // 视频长度。秒值
	MinBufferTime      float64                             `json:"minBufferTime"`   // 1.5？
	MinBufferTimeSnake float64                             `json:"min_buffer_time"` // 1.5？
	Video              []GetBangumiVideoStreamAudioOrVideo `json:"video"`           // 视频流信息。同一清晰度可拥有 H.264 / H.265 / AV1 多种码流
	Audio              []GetBangumiVideoStreamAudioOrVideo `json:"audio"`           // 伴音流信息。当视频没有音轨时，此项为 null
	Dolby              GetBangumiVideoStreamDolby          `json:"dolby"`           // 杜比全景声伴音信息
	Flac               GetBangumiVideoStreamFlac           `json:"flac"`            // 无损音轨伴音信息。当视频没有无损音轨时，此项为 null
}

// GetBangumiVideoStreamDolby 是 dash.dolby 对象，即杜比全景声伴音信息。
type GetBangumiVideoStreamDolby struct {
	Type  int                                 `json:"type"`  // 杜比音效类型。1：普通杜比音效。2：全景杜比音效
	Audio []GetBangumiVideoStreamAudioOrVideo `json:"audio"` // 杜比伴音流列表
}

// GetBangumiVideoStreamFlac 是 dash.flac 对象，即无损音轨伴音信息。
type GetBangumiVideoStreamFlac struct {
	Display bool                              `json:"display"` // 是否在播放器显示切换 Hi-Res 无损音轨按钮
	Audio   GetBangumiVideoStreamAudioOrVideo `json:"audio"`   // 音频流信息
}

// GetBangumiVideoStreamAudioOrVideo 描述一条 DASH 音视频流。
//
// 该接口对同一份数据会同时返回 camelCase 与 snake_case 两种键（例如 baseUrl
// 与 base_url），因此每种键各保留一个字段：规范化命名对应 camelCase 键，
// 以 Snake 结尾的字段对应 snake_case 键。
type GetBangumiVideoStreamAudioOrVideo struct {
	ID                int                              `json:"id"`             // 音视频清晰度代码。参考 [qn视频清晰度标识](#qn视频清晰度标识)
	BaseURL           string                           `json:"baseUrl"`        // 默认流 URL。注意 unicode 转义符。有效时间为 120min
	BaseURLSnake      string                           `json:"base_url"`       // 同上（snake_case 键）
	BackupURL         []string                         `json:"backupUrl"`      // 备用流 URL
	BackupURLSnake    []string                         `json:"backup_url"`     // 同上（snake_case 键）
	Bandwidth         int                              `json:"bandwidth"`      // 所需最低带宽。单位为 Byte
	MimeType          string                           `json:"mimeType"`       // 格式 mimetype 类型
	MimeTypeSnake     string                           `json:"mime_type"`      // 同上（snake_case 键）
	Codecs            string                           `json:"codecs"`         // 编码/音频类型。eg：avc1.640032
	Width             int                              `json:"width"`          // 视频宽度。单位为像素。仅视频流存在该字段
	Height            int                              `json:"height"`         // 视频高度。单位为像素。仅视频流存在该字段
	FrameRate         string                           `json:"frameRate"`      // 视频帧率。仅视频流存在该字段
	FrameRateSnake    string                           `json:"frame_rate"`     // 同上（snake_case 键）
	Sar               string                           `json:"sar"`            // Sample Aspect Ratio（单个像素的宽高比）。音频流该值恒为空
	StartWithSap      int                              `json:"startWithSap"`   // Stream Access Point（流媒体访问位点）。音频流该值恒为空
	StartWithSapSnake int                              `json:"start_with_sap"` // 同上（snake_case 键）
	SegmentBase       GetBangumiVideoStreamSegmentBase `json:"SegmentBase"`    // 见下表。音频流该值恒为空
	SegmentBaseSnake  GetBangumiVideoStreamSegmentBase `json:"segment_base"`   // 同上（snake_case 键）
	CodecID           int                              `json:"codecid"`        // 码流编码标识代码。音频流该值恒为 0
}

// GetBangumiVideoStreamSegmentBase 是 DASH 音视频流的 SegmentBase 对象。
type GetBangumiVideoStreamSegmentBase struct {
	Initialization string `json:"initialization"` // ${init_first}-${init_last}。ftyp 加 moov box 在 m4s 文件中的范围（单位为 bytes）
	IndexRange     string `json:"index_range"`    // ${sidx_first}-${sidx_last}。sidx box 在 m4s 文件中的范围（单位为 bytes）
}
