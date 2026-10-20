package live

// 直播间信息相关响应模型。
//
// 每个接口的返回体独立定义，嵌套对象均带方法前缀，不与其它接口共享类型。

// GetUserLiveRoomStatusResult 是用户对应的直播间状态。
type GetUserLiveRoomStatusResult struct {
	RoomStatus    int    `json:"roomStatus"`     // 直播间状态。0：无房间。1：有房间
	RoundStatus   int    `json:"roundStatus"`    // 轮播状态。0：未轮播。1：轮播
	LiveStatus    int    `json:"live_status"`    // 直播状态。0：未开播。1：直播中
	URL           string `json:"url"`            // 直播间网页url
	Title         string `json:"title"`          // 直播间标题
	Cover         string `json:"cover"`          // 直播间封面url
	Online        int64  `json:"online"`         // 直播间人气。值为上次直播时刷新
	RoomID        int64  `json:"roomid"`         // 直播间id（短号）
	BroadcastType int    `json:"broadcast_type"` // 0
	OnlineHidden  int    `json:"online_hidden"`  // 0
}

// GetRoomInitResult 是房间页初始化信息。
type GetRoomInitResult struct {
	RoomID      int64 `json:"room_id"`      // 直播间真实id
	ShortID     int64 `json:"short_id"`     // 直播间id（短号）
	UID         int64 `json:"uid"`          // 主播用户mid
	NeedP2P     int   `json:"need_p2p"`     // 是否p2p
	IsHidden    bool  `json:"is_hidden"`    // 是否隐藏
	IsLocked    bool  `json:"is_locked"`    // 是否锁定
	IsPortrait  bool  `json:"is_portrait"`  // 是否竖屏
	LiveStatus  int   `json:"live_status"`  // 直播状态。0：未开播。1：直播中。2：轮播中
	HiddenTill  int64 `json:"hidden_till"`  // 隐藏时间戳
	LockTill    int64 `json:"lock_till"`    // 锁定时间戳
	Encrypted   bool  `json:"encrypted"`    // 是否加密
	PwdVerified bool  `json:"pwd_verified"` // 加密房间是否通过密码验证。encrypted=true时才有意义
	LiveTime    int64 `json:"live_time"`    // 开播时间。未开播时为 -62170012800
	RoomShield  int   `json:"room_shield"`  // 未知
	IsSP        int   `json:"is_sp"`        // 是否为特殊直播间。0：普通直播间。1：付费直播间
	SpecialType int   `json:"special_type"` // 特殊直播间标志。0：普通直播间。1：付费直播间。2：拜年祭直播间
}

// GetMasterInfoResult 是主播信息的 data 对象。
type GetMasterInfoResult struct {
	Info         GetMasterInfoUserInfo `json:"info"`           // 主播信息
	Exp          GetMasterInfoExp      `json:"exp"`            // 经验等级
	FollowerNum  int64                 `json:"follower_num"`   // 主播粉丝数
	RoomID       int64                 `json:"room_id"`        // 直播间id（短号）
	MedalName    string                `json:"medal_name"`     // 粉丝勋章名
	GloryCount   int64                 `json:"glory_count"`    // 主播荣誉数
	Pendant      string                `json:"pendant"`        // 直播间头像框url
	LinkGroupNum int                   `json:"link_group_num"` // 0。作用尚不明确
	RoomNews     GetMasterInfoRoomNews `json:"room_news"`      // 主播公告
}

// GetMasterInfoUserInfo 是主播的基本信息。
type GetMasterInfoUserInfo struct {
	UID            int64                       `json:"uid"`             // 主播mid
	Uname          string                      `json:"uname"`           // 主播用户名
	Face           string                      `json:"face"`            // 主播头像url
	OfficialVerify GetMasterInfoOfficialVerify `json:"official_verify"` // 认证信息
	Gender         int                         `json:"gender"`          // 主播性别。-1：保密。0：女。1：男
}

// GetMasterInfoOfficialVerify 是主播的认证信息。
type GetMasterInfoOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetMasterInfoExp 是主播经验等级。
type GetMasterInfoExp struct {
	MasterLevel GetMasterInfoMasterLevel `json:"master_level"` // 主播等级
}

// GetMasterInfoMasterLevel 是主播等级信息。
type GetMasterInfoMasterLevel struct {
	Level   int     `json:"level"`   // 当前等级
	Color   int     `json:"color"`   // 等级框颜色
	Current []int64 `json:"current"` // 当前等级信息。下标0为升级积分，下标1为总积分
	Next    []int64 `json:"next"`    // 下一等级信息。下标0为升级积分，下标1为总积分
}

// GetMasterInfoRoomNews 是主播公告。
type GetMasterInfoRoomNews struct {
	Content   string `json:"content"`    // 公告内容
	Ctime     string `json:"ctime"`      // 公告时间
	CtimeText string `json:"ctime_text"` // 公告日期
}

// GetRoomBaseInfoResult 是直播间基本信息，by_room_ids 以直播间长ID为键。
type GetRoomBaseInfoResult struct {
	ByUIDs    map[string]any                 `json:"by_uids"`     // 空
	ByRoomIDs map[string]GetRoomBaseInfoItem `json:"by_room_ids"` // 直播间信息
}

// GetRoomBaseInfoItem 是单个直播间的基本信息。
type GetRoomBaseInfoItem struct {
	RoomID         int64  `json:"room_id"`          // 直播间长ID
	UID            int64  `json:"uid"`              // 主播用户mid
	AreaID         int    `json:"area_id"`          // 直播间分区ID
	LiveStatus     int    `json:"live_status"`      // 直播状态。0：未开播。1：直播中。2：轮播中
	LiveURL        string `json:"live_url"`         // 直播间网页url
	ParentAreaID   int    `json:"parent_area_id"`   // 直播间父分区ID
	Title          string `json:"title"`            // 直播间标题
	ParentAreaName string `json:"parent_area_name"` // 直播间父分区名称
	AreaName       string `json:"area_name"`        // 直播间分区名称
	LiveTime       string `json:"live_time"`        // 开播时间。yyyy-MM-dd HH:mm:ss
	Description    string `json:"description"`      // 直播间简介
	Tags           string `json:"tags"`             // 直播间标签。以 , 分隔
	Attention      int64  `json:"attention"`        // 关注数
	Online         int64  `json:"online"`           // 在线人数
	ShortID        int64  `json:"short_id"`         // 直播间短ID。为0是无短号
	Uname          string `json:"uname"`            // 主播用户名
	Cover          string `json:"cover"`            // 直播间封面url
	Background     string `json:"background"`       // 直播间背景url
	JoinSlide      int    `json:"join_slide"`       // 1
	LiveID         int64  `json:"live_id"`          // 0
	LiveIDStr      string `json:"live_id_str"`      // "0"
}

// GetRoomStatusInfoByUIDsItem 是批量查询结果中的单个直播间状态。
type GetRoomStatusInfoByUIDsItem struct {
	Title            string `json:"title"`               // 直播间标题
	RoomID           int64  `json:"room_id"`             // 直播间房间号。直播间实际房间号
	UID              int64  `json:"uid"`                 // 主播mid
	Online           int64  `json:"online"`              // 直播间在线人数
	LiveTime         int64  `json:"live_time"`           // 开播时间戳，单位秒，未开播时为0
	LiveStatus       int    `json:"live_status"`         // 直播间开播状态。0：未开播。1：正在直播。2：轮播中
	ShortID          int64  `json:"short_id"`            // 直播间房间号。直播间短房间号，常见于签约主播
	Area             int    `json:"area"`                // 直播间分区id
	AreaName         string `json:"area_name"`           // 直播间分区名
	AreaV2ID         int    `json:"area_v2_id"`          // 直播间新版分区id
	AreaV2Name       string `json:"area_v2_name"`        // 直播间新版分区名
	AreaV2ParentID   int    `json:"area_v2_parent_id"`   // 直播间父分区id
	AreaV2ParentName string `json:"area_v2_parent_name"` // 直播间父分区名
	Uname            string `json:"uname"`               // 主播用户名
	Face             string `json:"face"`                // 主播头像url
	TagName          string `json:"tag_name"`            // 直播间标签
	Tags             string `json:"tags"`                // 直播间自定标签
	CoverFromUser    string `json:"cover_from_user"`     // 直播间封面url
	Keyframe         string `json:"keyframe"`            // 直播间关键帧url
	LockTill         string `json:"lock_till"`           // 直播间封禁信息
	HiddenTill       string `json:"hidden_till"`         // 直播间隐藏信息
	BroadcastType    int    `json:"broadcast_type"`      // 直播类型。0：普通直播。1：手机直播
}

// GetRoomPlayInfoResult 是直播间信息（含直播流）的 data 对象。
type GetRoomPlayInfoResult struct {
	RoomID          int64                       `json:"room_id"`           // 直播间id
	ShortID         int64                       `json:"short_id"`          // 直播间短id
	UID             int64                       `json:"uid"`               // 主播uid
	IsHidden        bool                        `json:"is_hidden"`         // 直播间是否被隐藏
	IsLocked        bool                        `json:"is_locked"`         // 直播间是否被锁定
	IsPortrait      bool                        `json:"is_portrait"`       // 是否竖屏
	LiveStatus      int                         `json:"live_status"`       // 直播状态。0：未开播。1：直播中。2：轮播中
	HiddenTill      int64                       `json:"hidden_till"`       // 隐藏结束时间
	LockTill        int64                       `json:"lock_till"`         // 封禁结束时间。秒级时间戳
	Encrypted       bool                        `json:"encrypted"`         // 直播间为加密直播间
	PwdVerified     bool                        `json:"pwd_verified"`      // 是否通过密码验证。当 encrypted 为 true 时才有意义
	LiveTime        int64                       `json:"live_time"`         // 本次开播时间。秒级时间戳
	RoomShield      int                         `json:"room_shield"`       // 未知
	AllSpecialTypes []any                       `json:"all_special_types"` // 特殊直播间类型
	PlayURLInfo     *GetRoomPlayInfoPlayURLInfo `json:"playurl_info"`      // 直播流信息
}

// GetRoomPlayInfoPlayURLInfo 是直播流信息。
type GetRoomPlayInfoPlayURLInfo struct {
	ConfJSON string                  `json:"conf_json"` // 客户端配置
	PlayURL  *GetRoomPlayInfoPlayURL `json:"playurl"`   // 直播流
}

// GetRoomPlayInfoPlayURL 是直播流列表。
type GetRoomPlayInfoPlayURL struct {
	Cid     int64                    `json:"cid"`       // 直播间id
	GQNDesc []GetRoomPlayInfoQuality `json:"g_qn_desc"` // 清晰度列表
	Stream  []GetRoomPlayInfoStream  `json:"stream"`    // 直播流信息
	P2PData *GetRoomPlayInfoP2PData  `json:"p2p_data"`  // P2P信息
	DolbyQn any                      `json:"dolby_qn"`  // 杜比清晰度
}

// GetRoomPlayInfoQuality 是清晰度描述。
type GetRoomPlayInfoQuality struct {
	Qn       int     `json:"qn"`        // 清晰度代码
	Desc     string  `json:"desc"`      // 清晰度描述
	HdrDesc  string  `json:"hdr_desc"`  // HDR描述
	AttrDesc *string `json:"attr_desc"` // 属性描述
}

// GetRoomPlayInfoStream 是按协议划分的直播流。
type GetRoomPlayInfoStream struct {
	ProtocolName string                  `json:"protocol_name"` // 协议名
	Format       []GetRoomPlayInfoFormat `json:"format"`        // 格式列表
}

// GetRoomPlayInfoFormat 是直播流格式。
type GetRoomPlayInfoFormat struct {
	FormatName string                 `json:"format_name"` // 格式名
	Codec      []GetRoomPlayInfoCodec `json:"codec"`       // 编码列表
}

// GetRoomPlayInfoCodec 是直播流编码及其播放地址。
type GetRoomPlayInfoCodec struct {
	CodecName string                     `json:"codec_name"` // 编码名
	CurrentQn int                        `json:"current_qn"` // 当前清晰度编码
	AcceptQn  []int                      `json:"accept_qn"`  // 可用清晰度编码列表
	BaseURL   string                     `json:"base_url"`   // 播放源路径
	URLInfo   []GetRoomPlayInfoURLDetail `json:"url_info"`   // 域名信息列表
	HdrQn     *int                       `json:"hdr_qn"`     // HDR清晰度编码
	DolbyType int                        `json:"dolby_type"` // 杜比类型
	AttrName  string                     `json:"attr_name"`  // 属性名
}

// GetRoomPlayInfoURLDetail 是播放源域名信息。
type GetRoomPlayInfoURLDetail struct {
	Host      string `json:"host"`       // 域名
	Extra     string `json:"extra"`      // URL参数
	StreamTTL int64  `json:"stream_ttl"` // 流有效期
}

// GetRoomPlayInfoP2PData 是直播流 P2P 信息。
type GetRoomPlayInfoP2PData struct {
	P2P      bool `json:"p2p"`       // 是否启用P2P
	P2PType  int  `json:"p2p_type"`  // P2P类型
	MP2P     bool `json:"m_p2p"`     // 是否启用移动端P2P
	MServers any  `json:"m_servers"` // 移动端P2P服务器
}

// GetAnchorInRoomResult 是直播间主播信息的 data 对象。
type GetAnchorInRoomResult struct {
	Info  GetAnchorInRoomInfo  `json:"info"`  // 主播信息
	Level GetAnchorInRoomLevel `json:"level"` // 直播等级
	San   int                  `json:"san"`   // 主播san值。12满分
}

// GetAnchorInRoomInfo 是直播间主播的基本信息。
type GetAnchorInRoomInfo struct {
	UID               int64                         `json:"uid"`                 // 主播mid
	Uname             string                        `json:"uname"`               // 主播用户名
	Face              string                        `json:"face"`                // 主播头像URL
	Rank              string                        `json:"rank"`                // 主播排名
	PlatformUserLevel int                           `json:"platform_user_level"` // 平台用户等级
	MobileVerify      int                           `json:"mobile_verify"`       // 手机验证状态
	Identification    int                           `json:"identification"`      // 身份认证状态
	OfficialVerify    GetAnchorInRoomOfficialVerify `json:"official_verify"`     // 认证信息
	VipType           int                           `json:"vip_type"`            // VIP类型
	Gender            int                           `json:"gender"`              // 主播性别。-1：保密。0：女。1：男
}

// GetAnchorInRoomOfficialVerify 是主播认证信息（比普通认证多一个 role 字段）。
type GetAnchorInRoomOfficialVerify struct {
	Type int    `json:"type"` // 主播认证类型。-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 主播认证信息
	Role int    `json:"role"` // 未知
}

// GetAnchorInRoomLevel 是主播的直播等级信息。
type GetAnchorInRoomLevel struct {
	UID         int64                      `json:"uid"`          // 用户ID
	Cost        int64                      `json:"cost"`         // 消费金额
	Rcost       int64                      `json:"rcost"`        // 充值金额
	UserScore   string                     `json:"user_score"`   // 用户积分
	Vip         int                        `json:"vip"`          // VIP状态
	VipTime     string                     `json:"vip_time"`     // VIP到期时间
	Svip        int                        `json:"svip"`         // SVIP状态
	SvipTime    string                     `json:"svip_time"`    // SVIP到期时间
	UpdateTime  string                     `json:"update_time"`  // 更新时间
	MasterLevel GetAnchorInRoomMasterLevel `json:"master_level"` // 主播等级
	UserLevel   int                        `json:"user_level"`   // 用户等级
	Color       int                        `json:"color"`        // 颜色值
	AnchorScore int64                      `json:"anchor_score"` // 主播积分
}

// GetAnchorInRoomMasterLevel 是主播等级信息（含升级积分与颜色）。
type GetAnchorInRoomMasterLevel struct {
	Level            int     `json:"level"`              // 主播等级
	Color            int     `json:"color"`              // 颜色值
	Current          []int64 `json:"current"`            // 当前积分
	Next             []int64 `json:"next"`               // 下一等级积分
	AnchorScore      int64   `json:"anchor_score"`       // 主播积分
	UpgradeScore     int64   `json:"upgrade_score"`      // 升级积分
	MasterLevelColor int     `json:"master_level_color"` // 主播等级颜色值
	Sort             string  `json:"sort"`               // 排名
}
