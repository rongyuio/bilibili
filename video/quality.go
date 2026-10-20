package video

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

// VideoPlayedComplete 是 played_time 的特殊取值，表示视频已播放完成。
const VideoPlayedComplete = -1
