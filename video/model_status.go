package video

// 视频互动状态相关响应模型。按接口隔离，各接口的返回体独立定义，互不复用。

// GetVideoCoinStatusResult 是「判断视频是否被投币」的 data 对象。
type GetVideoCoinStatusResult struct {
	Multiply int `json:"multiply"` // 投币枚数。未投币为 0
}

// FavourVideoWebResult 是「收藏视频（Web 端）」的 data 对象。
type FavourVideoWebResult struct {
	Prompt     bool   `json:"prompt"`      // 是否为未关注用户收藏。false：否。true：是
	GaData     any    `json:"ga_data"`     // 作用尚不明确。可能为 null
	ToastMsg   string `json:"toast_msg"`   // 空。作用尚不明确
	SuccessNum int    `json:"success_num"` // 0。作用尚不明确
}

// GetVideoFavouredResult 是「判断视频是否被收藏」的 data 对象。
type GetVideoFavouredResult struct {
	Count    int  `json:"count"`    // 1。作用尚不明确
	Favoured bool `json:"favoured"` // 是否收藏。true：已收藏。false：未收藏
}
