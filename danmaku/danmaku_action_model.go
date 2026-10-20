package danmaku

import "encoding/json"

// 弹幕操作接口的响应模型。各接口的返回体独立定义，互不复用。

// SendVideoDanmakuResult 是发送视频弹幕的响应（data 部分）。
type SendVideoDanmakuResult struct {
	Action      string          `json:"action"`       // （？）
	Animation   string          `json:"animation"`    // 弹幕样式 JSON？
	ColorfulSrc json.RawMessage `json:"colorful_src"` // 渐变彩色样式信息，类型为字符串或对象。请求参数 colorful=60001 时有效
	DmContent   string          `json:"dm_content"`   // （？）
	Dmid        int64           `json:"dmid"`         // 弹幕 dmid
	DmidStr     string          `json:"dmid_str"`     // 弹幕 dmid 的字串形式
	Visible     bool            `json:"visible"`      // （？）
}

// CheckDanmakuAdvPermissionResult 是视频高级弹幕发送权限的当前状态。
type CheckDanmakuAdvPermissionResult struct {
	Coins   int  `json:"coins"`   // 需要支付的硬币数
	Confirm int  `json:"confirm"` // 是否同意。1：同意；2：未同意。未购买时无此项
	Accept  bool `json:"accept"`  // 是否允许申请
	HasBuy  bool `json:"hasBuy"`  // 是否已购买。未购买时无此项
}
