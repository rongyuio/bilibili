package bilibili

// 视频流与播放地址相关响应模型（DASH / Durl / Dolby / Flac 等）。

// 视频清晰度代码（qn）。用于 GetVideoStreamParam.Qn 与 ReportVideoWatchTime 的 quality 参数，
// 取值含义见接口文档的「qn视频清晰度标识」表。
const (
	VideoQuality240P      = 6   // 240P 极速。仅 MP4 格式、platform=html5 时有效
	VideoQuality360P      = 16  // 360P 流畅
	VideoQuality480P      = 32  // 480P 清晰
	VideoQuality720P      = 64  // 720P 高清。WEB 端默认值
	VideoQuality720P60    = 74  // 720P60 高帧率
	VideoQuality1080P     = 80  // 1080P 高清。TV 端与 APP 端默认值
	VideoQualitySmart     = 100 // 智能修复
	VideoQuality1080PPlus = 112 // 1080P+ 高码率
	VideoQuality1080P60   = 116 // 1080P60 高帧率
	VideoQuality4K        = 120 // 4K 超清
	VideoQualityHDR       = 125 // HDR 真彩色。仅 DASH 格式
	VideoQualityDolby     = 126 // 杜比视界。仅 DASH 格式
	VideoQuality8K        = 127 // 8K 超高清。仅 DASH 格式
	VideoQualityHDRVivid  = 129 // HDR Vivid
)

type SupportFormat struct {
	Quality        int      `json:"quality"`         // 视频清晰度代码。含义见 [上表](#qn视频清晰度标识)
	Format         string   `json:"format"`          // 视频格式
	NewDescription string   `json:"new_description"` // 格式描述
	DisplayDesc    string   `json:"display_desc"`    // 格式描述
	Superscript    string   `json:"superscript"`     // (?)
	Codecs         []string `json:"codecs"`          // 可用编码格式列表 例：av01.0.13M.08.0.110.01.01.01.0 使用AV1编码, avc1.640034 使用AVC编码, hev1.1.6.L153.90 使用HEVC编码
}

type Durl struct {
	Order     int      `json:"order"`      // 视频分段序号。某些视频会分为多个片段（从1顺序增长）
	Length    int      `json:"length"`     // 视频长度。单位为毫秒
	Size      int      `json:"size"`       // 视频大小。单位为 Byte
	Ahead     string   `json:"ahead"`      // （？）
	Vhead     string   `json:"vhead"`      // （？）
	URL       string   `json:"url"`        // 默认流 URL。注意 unicode 转义符。有效时间为120min
	BackupURL []string `json:"backup_url"` // 备用视频流 注意 unicode 转义符。有效时间为120min
}

type Dash struct {
	Duration           int            `json:"duration"`        // 视频长度。秒值
	MinBufferTime      float64        `json:"minBufferTime"`   // 1.5？
	MinBufferTimeSnake float64        `json:"min_buffer_time"` // 1.5？
	Video              []AudioOrVideo `json:"video"`           // 视频流信息 同一清晰度可拥有 H.264 / H.265 / AV1 多种码流<br />HDR 仅支持 H.265 |
	Audio              []AudioOrVideo `json:"audio"`           // 伴音流信息。当视频没有音轨时，此项为 null
	Dolby              Dolby          `json:"dolby"`           // 杜比全景声伴音信息
	Flac               Flac           `json:"flac"`            // 无损音轨伴音信息。当视频没有无损音轨时，此项为 null
}

type Dolby struct {
	Type  int            `json:"type"`  // 杜比音效类型。1：普通杜比音效。2：全景杜比音效
	Audio []AudioOrVideo `json:"audio"` // 杜比伴音流列表
}

type Flac struct {
	Display bool         `json:"display"` // 是否在播放器显示切换Hi-Res无损音轨按钮
	Audio   AudioOrVideo `json:"audio"`   // 音频流信息。同上文 DASH 流中video及audio数组中的对象
}

// AudioOrVideo 描述一条 DASH 音视频流。
//
// 该接口对同一份数据会同时返回 camelCase 与 snake_case 两种键（例如 baseUrl
// 与 base_url），因此每种键各保留一个字段：规范化命名对应 camelCase 键，
// 以 Snake 结尾的字段对应 snake_case 键。
type AudioOrVideo struct {
	ID                int         `json:"id"`             // 音视频清晰度代码。参考上表。[qn视频清晰度标识](#qn视频清晰度标识)。[视频伴音音质代码](#视 频伴音音质代码)
	BaseURL           string      `json:"baseUrl"`        // 默认流 URL。注意 unicode 转义符。有效时间为 120min
	BaseURLSnake      string      `json:"base_url"`       // 同上（snake_case 键）
	BackupURL         []string    `json:"backupUrl"`      // 备用流 URL
	BackupURLSnake    []string    `json:"backup_url"`     // 同上（snake_case 键）
	Bandwidth         int         `json:"bandwidth"`      // 所需最低带宽。单位为 Byte
	MimeType          string      `json:"mimeType"`       // 格式 mimetype 类型
	MimeTypeSnake     string      `json:"mime_type"`      // 同上（snake_case 键）
	Codecs            string      `json:"codecs"`         // 编码/音频类型。eg：avc1.640032
	Width             int         `json:"width"`          // 视频宽度。单位为像素。仅视频流存在该字段
	Height            int         `json:"height"`         // 视频高度。单位为像素。仅视频流存在该字段
	FrameRate         string      `json:"frameRate"`      // 视频帧率。仅视频流存在该字段
	FrameRateSnake    string      `json:"frame_rate"`     // 同上（snake_case 键）
	Sar               string      `json:"sar"`            // Sample Aspect Ratio（单个像素的宽高比）。音频流该值恒为空
	StartWithSap      int         `json:"startWithSap"`   // Stream Access Point（流媒体访问位点）。音频流该值恒为空
	StartWithSapSnake int         `json:"start_with_sap"` // 同上（snake_case 键）
	SegmentBase       SegmentBase `json:"SegmentBase"`    // 见下表。url 对应 m4s 文件中，头部的位置。音频流该值恒为空
	SegmentBaseSnake  SegmentBase `json:"segment_base"`   // 同上（snake_case 键）
	CodecID           int         `json:"codecid"`        // 码流编码标识代码。含义见 [上表](#视频编码代码)。音频流该值恒为0
}

type SegmentBase struct {
	Initialization string `json:"initialization"` // ${init_first}-${init_last}。eg：0-821。ftyp (file type) box 加 上 moov box 在 m4s 文件中的范围（单位为 bytes）。如 0-821 表示开头 820 个字节
	IndexRange     string `json:"index_range"`    // ${sidx_first}-${sidx_last}。eg：822-1309。sidx (segment index) box 在 m4s 文件中的范围（单位为 bytes）。sidx 的核心是一个数组，记录了各关键帧的时间戳及其在文件中的位置，。其作用是索引 (拖进 度条)
}

type GetVideoStreamResult struct {
	From              string          `json:"from"`               // local？
	Result            string          `json:"result"`             // suee？
	Message           string          `json:"message"`            // 空？
	Quality           int             `json:"quality"`            // 清晰度标识。含义见 [上表](#qn视频清晰度标识)
	Format            string          `json:"format"`             // 视频格式。mp4/flv
	TimeLength        int             `json:"timelength"`         // 视频长度。单位为毫秒。不同分辨率 / 格式可能有略微差异
	AcceptFormat      string          `json:"accept_format"`      // 支持的全部格式。每项用,分隔
	AcceptDescription []string        `json:"accept_description"` // 支持的清晰度列表（文字说明）
	AcceptQuality     []int           `json:"accept_quality"`     // 支持的清晰度列表（代码）。含义见 [上表](#qn视频清晰度标识)
	VideoCodecID      int             `json:"video_codecid"`      // 默认选择视频流的编码id。含义见 [上表](#视频编码代码)
	SeekParam         string          `json:"seek_param"`         // start？
	SeekType          string          `json:"seek_type"`          // offset（DASH / FLV）？。 second（MP4）？
	Durl              []Durl          `json:"durl"`               // 视频分段流信息。注：仅 FLV / MP4 格式存在此字段
	Dash              Dash            `json:"dash"`               // DASH 流信息。注：仅 DASH 格式存在此字段
	SupportFormats    []SupportFormat `json:"support_formats"`    // 支持格式的详细信息
	HighFormat        *string         `json:"high_format"`        // （？）null
	LastPlayTime      int             `json:"last_play_time"`     // 上次播放进度。毫秒值
	LastPlayCid       int             `json:"last_play_cid"`      // 上次播放分P的 cid
}
