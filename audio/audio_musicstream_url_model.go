package audio

// 音频流URL相关响应模型。
//
// 两个接口的返回体结构一致，但按接口各自独立定义，避免一方变更影响另一方。

// GetAudioStreamURLQuality 是获取音频流URL（web 端）返回的一种音质。
type GetAudioStreamURLQuality struct {
	Type        int    `json:"type"`        // 音质代码。0：128K。1：192K。2：320K。3：无损FLAC
	Desc        string `json:"desc"`        // 音质名称
	Size        int64  `json:"size"`        // 该音质的文件大小。单位为字节
	Bps         string `json:"bps"`         // 比特率标签
	Tag         string `json:"tag"`         // 音质标签
	Require     int    `json:"require"`     // 是否需要会员权限。0：不需要。1：需要
	Requiredesc string `json:"requiredesc"` // 会员权限标签
}

// GetAudioStreamURLResult 是获取音频流URL（web 端）接口的 data 对象。
type GetAudioStreamURLResult struct {
	Sid       int                        `json:"sid"`       // 音频auid
	Type      int                        `json:"type"`      // 音质标识。-1：试听片段（192K）。0：128K。1：192K。2：320K。3：FLAC
	Info      string                     `json:"info"`      // 空。作用尚不明确
	Timeout   int                        `json:"timeout"`   // 有效时长。单位为秒，一般为3小时
	Size      int64                      `json:"size"`      // 文件大小。单位为字节，type为-1时为0
	Cdns      []string                   `json:"cdns"`      // 音频流url。可能包含备用地址
	Qualities []GetAudioStreamURLQuality `json:"qualities"` // 音质列表。web端接口返回null
	Title     string                     `json:"title"`     // 音频标题。web端接口为空
	Cover     string                     `json:"cover"`     // 音频封面url。web端接口为空
}

// GetPaidAudioStreamURLQuality 是获取音频流URL（可获取付费音频）返回的一种音质。
type GetPaidAudioStreamURLQuality struct {
	Type        int    `json:"type"`        // 音质代码。0：128K。1：192K。2：320K。3：无损FLAC
	Desc        string `json:"desc"`        // 音质名称
	Size        int64  `json:"size"`        // 该音质的文件大小。单位为字节
	Bps         string `json:"bps"`         // 比特率标签
	Tag         string `json:"tag"`         // 音质标签
	Require     int    `json:"require"`     // 是否需要会员权限。0：不需要。1：需要
	Requiredesc string `json:"requiredesc"` // 会员权限标签
}

// GetPaidAudioStreamURLResult 是获取音频流URL（可获取付费音频）接口的 data 对象。
type GetPaidAudioStreamURLResult struct {
	Sid       int                            `json:"sid"`       // 音频auid
	Type      int                            `json:"type"`      // 音质标识。-1：试听片段（192K）。0：128K。1：192K。2：320K。3：FLAC
	Info      string                         `json:"info"`      // 空。作用尚不明确
	Timeout   int                            `json:"timeout"`   // 有效时长。单位为秒，一般为3小时
	Size      int64                          `json:"size"`      // 文件大小。单位为字节，type为-1时为0
	Cdns      []string                       `json:"cdns"`      // 音频流url。可能包含备用地址
	Qualities []GetPaidAudioStreamURLQuality `json:"qualities"` // 音质列表
	Title     string                         `json:"title"`     // 音频标题
	Cover     string                         `json:"cover"`     // 音频封面url
}
