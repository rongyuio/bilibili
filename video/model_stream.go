package video

// GetVideoStream 接口的响应模型（DASH / Durl / Dolby / Flac 等）。

// GetVideoStreamResult 是「获取视频流地址_web端」的 data 对象。
type GetVideoStreamResult struct {
	From              string                        `json:"from"`               // local？
	Result            string                        `json:"result"`             // suee？
	Message           string                        `json:"message"`            // 空？
	Quality           int                           `json:"quality"`            // 清晰度标识
	Format            string                        `json:"format"`             // 视频格式。mp4/flv
	TimeLength        int                           `json:"timelength"`         // 视频长度。单位为毫秒
	AcceptFormat      string                        `json:"accept_format"`      // 支持的全部格式。每项用,分隔
	AcceptDescription []string                      `json:"accept_description"` // 支持的清晰度列表（文字说明）
	AcceptQuality     []int                         `json:"accept_quality"`     // 支持的清晰度列表（代码）
	VideoCodecID      int                           `json:"video_codecid"`      // 默认选择视频流的编码id
	SeekParam         string                        `json:"seek_param"`         // start？
	SeekType          string                        `json:"seek_type"`          // offset（DASH / FLV）？。 second（MP4）？
	Durl              []GetVideoStreamDurl          `json:"durl"`               // 视频分段流信息。注：仅 FLV / MP4 格式存在此字段
	Dash              GetVideoStreamDash            `json:"dash"`               // DASH 流信息。注：仅 DASH 格式存在此字段
	SupportFormats    []GetVideoStreamSupportFormat `json:"support_formats"`    // 支持格式的详细信息
	HighFormat        *string                       `json:"high_format"`        // （？）null
	LastPlayTime      int                           `json:"last_play_time"`     // 上次播放进度。毫秒值
	LastPlayCid       int                           `json:"last_play_cid"`      // 上次播放分P的 cid
}

// GetVideoStreamSupportFormat 是「获取视频流地址_web端」support_formats 数组中的对象。
type GetVideoStreamSupportFormat struct {
	Quality        int      `json:"quality"`         // 视频清晰度代码
	Format         string   `json:"format"`          // 视频格式
	NewDescription string   `json:"new_description"` // 格式描述
	DisplayDesc    string   `json:"display_desc"`    // 格式描述
	Superscript    string   `json:"superscript"`     // (?)
	Codecs         []string `json:"codecs"`          // 可用编码格式列表
}

// GetVideoStreamDurl 是「获取视频流地址_web端」durl 数组中的对象。
type GetVideoStreamDurl struct {
	Order     int      `json:"order"`      // 视频分段序号。某些视频会分为多个片段（从1顺序增长）
	Length    int      `json:"length"`     // 视频长度。单位为毫秒
	Size      int      `json:"size"`       // 视频大小。单位为 Byte
	Ahead     string   `json:"ahead"`      // （？）
	Vhead     string   `json:"vhead"`      // （？）
	URL       string   `json:"url"`        // 默认流 URL。有效时间为120min
	BackupURL []string `json:"backup_url"` // 备用视频流。有效时间为120min
}

// GetVideoStreamDash 是「获取视频流地址_web端」dash 对象。
type GetVideoStreamDash struct {
	Duration           int                          `json:"duration"`        // 视频长度。秒值
	MinBufferTime      float64                      `json:"minBufferTime"`   // 1.5？
	MinBufferTimeSnake float64                      `json:"min_buffer_time"` // 1.5？
	Video              []GetVideoStreamAudioOrVideo `json:"video"`           // 视频流信息。同一清晰度可拥有多种码流
	Audio              []GetVideoStreamAudioOrVideo `json:"audio"`           // 伴音流信息。当视频没有音轨时，此项为 null
	Dolby              GetVideoStreamDolby          `json:"dolby"`           // 杜比全景声伴音信息
	Flac               GetVideoStreamFlac           `json:"flac"`            // 无损音轨伴音信息。当视频没有无损音轨时，此项为 null
}

// GetVideoStreamDolby 是「获取视频流地址_web端」dash.dolby 对象。
type GetVideoStreamDolby struct {
	Type  int                          `json:"type"`  // 杜比音效类型。1：普通杜比音效。2：全景杜比音效
	Audio []GetVideoStreamAudioOrVideo `json:"audio"` // 杜比伴音流列表
}

// GetVideoStreamFlac 是「获取视频流地址_web端」dash.flac 对象。
type GetVideoStreamFlac struct {
	Display bool                       `json:"display"` // 是否在播放器显示切换Hi-Res无损音轨按钮
	Audio   GetVideoStreamAudioOrVideo `json:"audio"`   // 音频流信息
}

// GetVideoStreamAudioOrVideo 描述一条 DASH 音视频流。
//
// 该接口对同一份数据会同时返回 camelCase 与 snake_case 两种键（例如 baseUrl
// 与 base_url），因此每种键各保留一个字段：规范化命名对应 camelCase 键，
// 以 Snake 结尾的字段对应 snake_case 键。
type GetVideoStreamAudioOrVideo struct {
	ID                int                       `json:"id"`             // 音视频清晰度代码
	BaseURL           string                    `json:"baseUrl"`        // 默认流 URL
	BaseURLSnake      string                    `json:"base_url"`       // 同上（snake_case 键）
	BackupURL         []string                  `json:"backupUrl"`      // 备用流 URL
	BackupURLSnake    []string                  `json:"backup_url"`     // 同上（snake_case 键）
	Bandwidth         int                       `json:"bandwidth"`      // 所需最低带宽。单位为 Byte
	MimeType          string                    `json:"mimeType"`       // 格式 mimetype 类型
	MimeTypeSnake     string                    `json:"mime_type"`      // 同上（snake_case 键）
	Codecs            string                    `json:"codecs"`         // 编码/音频类型。eg：avc1.640032
	Width             int                       `json:"width"`          // 视频宽度。单位为像素。仅视频流存在该字段
	Height            int                       `json:"height"`         // 视频高度。单位为像素。仅视频流存在该字段
	FrameRate         string                    `json:"frameRate"`      // 视频帧率。仅视频流存在该字段
	FrameRateSnake    string                    `json:"frame_rate"`     // 同上（snake_case 键）
	Sar               string                    `json:"sar"`            // Sample Aspect Ratio
	StartWithSap      int                       `json:"startWithSap"`   // Stream Access Point
	StartWithSapSnake int                       `json:"start_with_sap"` // 同上（snake_case 键）
	SegmentBase       GetVideoStreamSegmentBase `json:"SegmentBase"`    // 见下表
	SegmentBaseSnake  GetVideoStreamSegmentBase `json:"segment_base"`   // 同上（snake_case 键）
	CodecID           int                       `json:"codecid"`        // 码流编码标识代码
}

// GetVideoStreamSegmentBase 是「获取视频流地址_web端」dash 音视频流的 SegmentBase 对象。
type GetVideoStreamSegmentBase struct {
	Initialization string `json:"initialization"` // ${init_first}-${init_last}。ftyp 加 moov box 在 m4s 文件中的范围（单位为 bytes）
	IndexRange     string `json:"index_range"`    // ${sidx_first}-${sidx_last}。sidx box 在 m4s 文件中的范围（单位为 bytes）
}
