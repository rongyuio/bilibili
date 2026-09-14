package bilibili

// 直播间信息、开播与分区相关响应模型。

type StudioInfo struct {
	Status     int   `json:"status"`
	MasterList []any `json:"master_list"`
}

type LiveRoomInfo struct {
	UID                  int         `json:"uid"`                // 主播mid
	RoomID               int         `json:"room_id"`            // 直播间长号
	ShortID              int         `json:"short_id"`           // 直播间短号。为0是无短号
	Attention            int         `json:"attention"`          // 关注数量
	Online               int         `json:"online"`             // 观看人数
	IsPortrait           bool        `json:"is_portrait"`        // 是否竖屏
	Description          string      `json:"description"`        // 描述
	LiveStatus           int         `json:"live_status"`        // 直播状态。0：未开播。1：直播中。2：轮播中
	AreaID               int         `json:"area_id"`            // 分区id
	ParentAreaID         int         `json:"parent_area_id"`     // 父分区id
	ParentAreaName       string      `json:"parent_area_name"`   // 父分区名称
	OldAreaID            int         `json:"old_area_id"`        // 旧版分区id
	Background           string      `json:"background"`         // 背景图片链接
	Title                string      `json:"title"`              // 标题
	UserCover            string      `json:"user_cover"`         // 封面
	Keyframe             string      `json:"keyframe"`           // 关键帧。用于网页端悬浮展示
	IsStrictRoom         bool        `json:"is_strict_room"`     // 未知。未知
	LiveTime             string      `json:"live_time"`          // 直播开始时间。YYYY-MM-DD HH:mm:ss
	Tags                 string      `json:"tags"`               // 标签。','分隔
	IsAnchor             int         `json:"is_anchor"`          // 未知。未知
	RoomSilentType       string      `json:"room_silent_type"`   // 禁言状态
	RoomSilentLevel      int         `json:"room_silent_level"`  // 禁言等级
	RoomSilentSecond     int         `json:"room_silent_second"` // 禁言时间。单位是秒
	AreaName             string      `json:"area_name"`          // 分区名称
	Pardants             string      `json:"pardants"`           // 未知。未知
	AreaPardants         string      `json:"area_pardants"`      // 未知。未知
	HotWords             []string    `json:"hot_words"`          // 热词
	HotWordsStatus       int         `json:"hot_words_status"`   // 热词状态
	Verify               string      `json:"verify"`             // 未知。未知
	NewPendants          NewPendants `json:"new_pendants"`       // 头像框\大v
	UpSession            string      `json:"up_session"`         // 未知
	PkStatus             int         `json:"pk_status"`          // pk状态
	PkID                 int         `json:"pk_id"`              // pk id
	BattleID             int         `json:"battle_id"`          // 未知
	AllowChangeAreaTime  int         `json:"allow_change_area_time"`
	AllowUploadCoverTime int         `json:"allow_upload_cover_time"`
	StudioInfo           StudioInfo  `json:"studio_info"`
}

type UpdateLiveRoomTitleResult struct {
	SubSessionKey string `json:"sub_session_key"` // 信息变动标识
	AuditInfo     any    `json:"audit_info"`      // 标题审核信息（不一定有值，因此在这里不进行解析）
}

type Rtmp struct {
	Addr     string `json:"addr"`     // RTMP推流（发送）地址。**重要**
	Code     string `json:"code"`     // RTMP推流参数（密钥）。**重要**
	NewLink  string `json:"new_link"` // 获取CDN推流ip地址重定向信息的url。没啥用
	Provider string `json:"provider"` // ？？？。作用尚不明确
}

type Protocol struct {
	Protocol string `json:"protocol"` // rtmp。作用尚不明确
	Addr     string `json:"addr"`     // RTMP推流（发送）地址
	Code     string `json:"code"`     // RTMP推流参数（密钥）
	NewLink  string `json:"new_link"` // 获取CDN推流ip地址重定向信息的url
	Provider string `json:"provider"` // txy。作用尚不明确
}

type StartLiveResult struct {
	Change    int        `json:"change"`    // 是否改变状态。0：未改变。1：改变
	Status    string     `json:"status"`    // LIVE
	RoomType  int        `json:"room_type"` // 0。作用尚不明确
	Rtmp      Rtmp       `json:"rtmp"`      // RTMP推流地址信息
	Protocols []Protocol `json:"protocols"` // ？？？。作用尚不明确
	TryTime   string     `json:"try_time"`  // ？？？。作用尚不明确
	LiveKey   string     `json:"live_key"`  // ？？？。作用尚不明确
	Notice    Notice     `json:"notice"`    // ？？？。作用尚不明确
}

type StopLiveResult struct {
	Change int    `json:"change"` // 是否改变状态。0：未改变。1：改变
	Status string `json:"status"` // PREPARING
}

type SubLiveArea struct {
	ID         string `json:"id"`          // 子分区id
	ParentID   string `json:"parent_id"`   // 父分区id
	OldAreaID  string `json:"old_area_id"` // 旧分区id
	Name       string `json:"name"`        // 子分区名
	ActID      string `json:"act_id"`      // 0。**作用尚不明确**
	PkStatus   string `json:"pk_status"`   // ？？？。**作用尚不明确**
	HotStatus  int    `json:"hot_status"`  // 是否为热门分区。0：否。1：是
	LockStatus string `json:"lock_status"` // 0。**作用尚不明确**
	Pic        string `json:"pic"`         // 子分区标志图片url
	ParentName string `json:"parent_name"` // 父分区名
	AreaType   int    `json:"area_type"`
}

type LiveAreaList struct {
	ID   int           `json:"id"`   // 父分区id
	Name string        `json:"name"` // 父分区名
	List []SubLiveArea `json:"list"` // 子分区列表
}

type HomePageLiveVersion struct {
	CurrVersion      string `json:"curr_version,omitempty" request:"query,omitempty"`      // 直播姬最新版本号
	Build            int    `json:"build,omitempty" request:"query,omitempty"`             // 直播姬构建号
	Instruction      string `json:"instruction,omitempty" request:"query,omitempty"`       // 更新说明（简要）
	FileSize         string `json:"file_size,omitempty" request:"query,omitempty"`         // 文件大小（字节）
	FileMd5          string `json:"file_md5,omitempty" request:"query,omitempty"`          // 安装包文件MD5
	Content          string `json:"content,omitempty" request:"query,omitempty"`           // HTML格式的更新内容
	DownloadURL      string `json:"download_url,omitempty" request:"query,omitempty"`      // 安装包下载链接
	HdiffpatchSwitch int    `json:"hdiffpatch_switch,omitempty" request:"query,omitempty"` // 增量更新开关?
}
