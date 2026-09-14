package bilibili

// 动态卡片、门户与发布接口响应模型。

// DynamicCard 动态卡片内容。因为 ActivityInfos 、 Desc 、 Display 等字段会随着此动态类型不同发生一定的变化，无法统一，因此都转换成了 map[string]any ，请自行解析
type DynamicCard struct {
	ActivityInfos map[string]any `json:"activity_infos"` // 该条动态参与的活动
	Card          string         `json:"card"`           // 动态详细信息
	Desc          map[string]any `json:"desc"`           // 动态相关信息
	Display       map[string]any `json:"display"`        // 动态部分的可操作项
	ExtendJSON    string         `json:"extend_json"`    // 动态扩展项
}

type DynamicDetail struct {
	Card   *DynamicCard `json:"card"` // 动态卡片内容
	Result int          `json:"result"`
	Gt     int          `json:"_gt_"`
}

type DynamicPortal struct {
	MyInfo DynamicPortalMyInfo `json:"my_info"` // 个人关注的一些信息
	UpList []DynamicPortalUp   `json:"up_list"` // 最近更新的up主列表
}

// DynamicPortalMyInfo 是动态门户中的个人关注信息。
type DynamicPortalMyInfo struct {
	Dyns      int                    `json:"dyns"`      // 个人动态
	Face      string                 `json:"face"`      // 头像url
	FaceNft   int                    `json:"face_nft"`  // 含义尚不明确
	Follower  int                    `json:"follower"`  // 粉丝数量
	Following int                    `json:"following"` // 我的关注
	LevelInfo DynamicPortalLevelInfo `json:"level_info"`
	Mid       int                    `json:"mid"`      // 账户mid
	Name      string                 `json:"name"`     // 账户名称
	Official  Official               `json:"official"` // 认证信息
	SpaceBg   string                 `json:"space_bg"` // 账户个人中心的背景横幅url
	Vip       CardVip                `json:"vip"`      // vip信息
}

// DynamicPortalLevelInfo 是本人等级信息。
type DynamicPortalLevelInfo struct {
	CurrentExp   int   `json:"current_exp"`
	CurrentLevel int   `json:"current_level"` // 当前等级，0-6级
	CurrentMin   int   `json:"current_min"`
	LevelUp      int64 `json:"level_up"`
	NextExp      int   `json:"next_exp"`
}

// DynamicPortalUp 是一位最近更新的 UP 主。
type DynamicPortalUp struct {
	Face            string `json:"face"`       // UP主头像
	HasUpdate       bool   `json:"has_update"` // 最近是否有更新
	IsReserveRecall bool   `json:"is_reserve_recall"`
	Mid             int    `json:"mid"`   // UP主mid
	Uname           string `json:"uname"` // UP主昵称
}

// UploadDynamicBfsResult contains the uploaded image url and its pixel size.
type UploadDynamicBfsResult struct {
	ImageURL    string `json:"image_url"`    // 图片 url
	ImageWidth  int    `json:"image_width"`  // 图片宽度
	ImageHeight int    `json:"image_height"` // 图片高度
}

type CreateDynamicResult struct {
	Result       int    `json:"result"`         // 0
	Errmsg       string `json:"errmsg"`         // 像是服务器日志一样的东西
	DynamicID    int    `json:"dynamic_id"`     // 动态 id
	CreateResult int    `json:"create_result"`  // 1
	DynamicIDStr string `json:"dynamic_id_str"` // 动态 id。字符串格式
	Gt           int    `json:"_gt_"`           // 0
}

// DynamicList 包含置顶及热门的动态列表
//
// TODO: 因为不清楚 attentions 字段（关注列表）的格式，暂未对此字段进行解析
type DynamicList struct {
	Cards         *DynamicCard `json:"cards"` // 动态列表
	FounderUID    int          `json:"founder_uid,omitempty"`
	HasMore       int          `json:"has_more"` // 当前话题是否有额外的动态，0：无额外动态，1：有额外动态
	IsDrawerTopic int          `json:"is_drawer_topic,omitempty"`
	Offset        string       `json:"offset"` // 接下来获取列表时的偏移值，一般为当前获取的话题列表下最后一个动态id
	Gt            int          `json:"_gt_"`   // 固定值0，作用尚不明确
}
