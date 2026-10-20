package login

// 登录记录相关响应模型。各接口的返回体独立定义，互不复用。

// GetLoginNoticeResult 是查询登录记录接口的 data 对象。
type GetLoginNoticeResult struct {
	Mid        int    `json:"mid"`         // 登录用户mid
	DeviceName string `json:"device_name"` // 登录设备。依靠操作登录接口时的UA决定
	LoginType  string `json:"login_type"`  // 登录方式。根据登录接口决定
	LoginTime  string `json:"login_time"`  // 登录时间。格式为 YYYY-MM-DD hh:mm:ss
	Location   string `json:"location"`    // 登录位置。依靠ip决定
	IP         string `json:"ip"`          // 登录ip。部分用*打码
}

// GetLoginLogItem 是一条登录记录。
type GetLoginLogItem struct {
	IP     string `json:"ip"`      // 登录IP。IPv4末2位以*打码，IPv6末5位以*打码
	Time   int64  `json:"time"`    // 登录时间。UNIX秒级时间戳
	TimeAt string `json:"time_at"` // 登录时间。格式为 yyyy-MM-dd HH:mm:ss
	Status bool   `json:"status"`  // 是否登录成功
	Type   int    `json:"type"`    // 登录方式
	Geo    string `json:"geo"`     // 登录地理位置
}

// GetLoginLogResult 是最近一周登录情况接口的 data 对象。
type GetLoginLogResult struct {
	Count int               `json:"count"` // 记录总数
	List  []GetLoginLogItem `json:"list"`  // 登录记录列表
}
