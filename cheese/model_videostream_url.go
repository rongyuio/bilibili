package cheese

// GetCheeseVideoStream 接口的响应模型。
//
// 本文件的类型只服务于 GetCheeseVideoStream，不与其它接口共用；
// durl 数组元素也在本文件中独立定义。

// GetCheeseVideoStreamResult 是「获取课程视频流 URL」的 data 对象。
type GetCheeseVideoStreamResult struct {
	AcceptFormat      string                              `json:"accept_format"`      // 视频支持的分辨率格式。每项用 , 分隔
	Code              int                                 `json:"code"`               // 固定为 0
	Durl              []GetCheeseVideoStreamDurl          `json:"durl"`               // 视频分段
	SeekParam         string                              `json:"seek_param"`         // 固定值 start。作用尚不明确
	NoRexcode         int                                 `json:"no_rexcode"`         // 固定为 0
	Format            string                              `json:"format"`             // 视频格式
	Fnval             int                                 `json:"fnval"`              // 固定为 0
	VideoProject      bool                                `json:"video_project"`      // 固定为 true
	Fnver             int                                 `json:"fnver"`              // 固定为 0
	Message           string                              `json:"message"`            // 空。作用尚不明确
	Type              string                              `json:"type"`               // 视频格式。FLV 等
	AcceptQuality     []int                               `json:"accept_quality"`     // 视频支持的分辨率代码列表。含义见 [qn视频清晰度标识]
	Quality           int                                 `json:"quality"`            // 当前视频分辨率代码。含义见 [qn视频清晰度标识]
	TimeLength        int                                 `json:"timelength"`         // 视频长度。单位为毫秒，不同分辨率可能有略微差异
	Result            string                              `json:"result"`             // suee。作用尚不明确
	SeekType          string                              `json:"seek_type"`          // offset。作用尚不明确
	HasPaid           bool                                `json:"has_paid"`           // false。作用尚不明确
	SupportFormats    []GetCheeseVideoStreamSupportFormat `json:"supportFormats"`     // 视频分辨率详细列表
	From              string                              `json:"from"`               // local。作用尚不明确
	VideoCodecID      int                                 `json:"video_codecid"`      // 视频编码代码。作用尚不明确
	AcceptDescription []string                            `json:"accept_description"` // 视频支持的分辨率列表（文字说明）
	Status            int                                 `json:"status"`             // 固定为 0。作用尚不明确
}

// GetCheeseVideoStreamDurl 是「获取课程视频流 URL」durl 数组中的对象。
type GetCheeseVideoStreamDurl struct {
	Order     int      `json:"order"`      // 视频分段序号。某些视频会分为多个片段（从1顺序增长）
	Length    int      `json:"length"`     // 视频长度。单位为毫秒
	Size      int      `json:"size"`       // 视频大小。单位为 Byte
	Ahead     string   `json:"ahead"`      // （？）
	Vhead     string   `json:"vhead"`      // （？）
	URL       string   `json:"url"`        // 默认流 URL。注意 unicode 转义符。有效时间为120min
	BackupURL []string `json:"backup_url"` // 备用视频流 注意 unicode 转义符。有效时间为120min
}

// GetCheeseVideoStreamSupportFormat 是「获取课程视频流 URL」supportFormats 数组中的对象。
type GetCheeseVideoStreamSupportFormat struct {
	Format      string `json:"format"`      // 分辨率名称
	Description string `json:"description"` // 分辨率备注
	Quality     int    `json:"quality"`     // 分辨率代码。含义见 [qn视频清晰度标识]
}
