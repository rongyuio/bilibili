package video

// GetTopRecommendVideo 接口的响应模型。

// GetTopRecommendVideoResult 是「获取首页视频推荐列表」的 data 对象。
type GetTopRecommendVideoResult struct {
	BusinessCard          any                        `json:"business_card"`            // 无意义
	FloorInfo             any                        `json:"floor_info"`               // 无意义
	Item                  []GetTopRecommendVideoItem `json:"item"`                     // 推荐列表
	Mid                   int                        `json:"mid"`                      // 用户mid,未登录为0
	PreloadExposePct      float64                    `json:"preload_expose_pct"`       // 用于预加载?
	PreloadFloorExposePct float64                    `json:"preload_floor_expose_pct"` // 用于预加载?
	SideBarColumn         []any                      `json:"side_bar_column"`          // 边栏列表?	可参考字段 item 及对应功能文档
	UserFeature           any                        `json:"user_feature"`             // 无意义
}

// GetTopRecommendVideoItem 是「获取首页视频推荐列表」item 数组中的对象。
type GetTopRecommendVideoItem struct {
	AvFeature       any                            `json:"av_feature"`        // 暂无参考意义
	BusinessInfo    any                            `json:"business_info"`     // 商业推广信息，通常为 null
	Bvid            string                         `json:"bvid"`              // 视频 bvid
	Cid             int                            `json:"cid"`               // 视频 cid
	DislikeSwitch   int                            `json:"dislike_switch"`    // 不感兴趣开关
	DislikeSwitchPC int                            `json:"dislike_switch_pc"` // PC端不感兴趣开关
	Duration        int                            `json:"duration"`          // 视频时长
	EnableVt        int                            `json:"enable_vt"`         // 未知作用
	Goto            string                         `json:"goto"`              // 目标类型 (av, ogv, live)
	ID              int                            `json:"id"`                // 视频 avid / 直播间 id
	IsFollowed      int                            `json:"is_followed"`       // 是否已关注
	IsStock         int                            `json:"is_stock"`          // 未知作用
	OgvInfo         any                            `json:"ogv_info"`          // 通常为 null
	Owner           GetTopRecommendVideoOwner      `json:"owner"`             // 视频 UP 主信息
	Pic             string                         `json:"pic"`               // 视频封面
	Pic43           string                         `json:"pic_4_3"`           // 4:3 比例封面
	Pos             int                            `json:"pos"`               // 位置
	PubDate         int                            `json:"pubdate"`           // 发布时间（秒级时间戳）
	RcmdReason      GetTopRecommendVideoRcmdReason `json:"rcmd_reason"`       // 推荐理由
	RoomInfo        any                            `json:"room_info"`         // 通常为 null
	ShowInfo        int                            `json:"show_info"`         // 展示信息（1: 普通视频, 0: 直播）
	Stat            any                            `json:"stat"`              // 视频状态信息
	Title           string                         `json:"title"`             // 视频标题
	TrackID         string                         `json:"track_id"`          // 跟踪标识
	URI             string                         `json:"uri"`               // 目标页 URI
	VtDisplay       string                         `json:"vt_display"`        // 未知作用
}

// GetTopRecommendVideoOwner 是「获取首页视频推荐列表」item[].owner 对象。
type GetTopRecommendVideoOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP昵称
	Face string `json:"face"` // 头像URL
}

// GetTopRecommendVideoRcmdReason 是「获取首页视频推荐列表」item[].rcmd_reason 对象。
type GetTopRecommendVideoRcmdReason struct {
	ReasonType int    `json:"reason_type"` // 原因类型
	Content    string `json:"content"`     // 原因描述（仅当 reason_type 为 3 时存在）
}
