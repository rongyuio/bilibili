package electric

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 包月充电相关接口。各接口的返回体独立定义，互不复用。
// 文档见 electric/monthly.md。

// liveReferer 给**直播域**的接口设 Referer / Origin。
//
// 包月充电列表走 `api.live.bilibili.com`，而真实调用发生在直播间页面
// （`https://live.bilibili.com/`）—— 与 `New()` 给的主站头对不上，
// 风控会回业务码 -352。拿不到房间号的接口退成直播首页，域对得上。
func liveReferer() bilibili.Handler {
	return func(r *resty.Request) error {
		r.SetHeader("Referer", "https://live.bilibili.com/")
		r.SetHeader("Origin", "https://live.bilibili.com")
		return nil
	}
}

// GetElectricMonthlyListParam 是获取包月充电列表的请求参数。
type GetElectricMonthlyListParam struct {
	Page int `json:"page" request:"query"` // 页码
	Type int `json:"type" request:"query"` // 充电状态。1：使用中，2：已过期
}

// GetElectricMonthlyListRenew 是包月充电列表中的充电自动续费详情。
type GetElectricMonthlyListRenew struct {
	UID             int    `json:"uid"`               // 自己的mid
	RUID            int    `json:"ruid"`              // UP主的mid
	GoodsID         int    `json:"goods_id"`          // 充电类型。172：一个月，173：连续包月，174：连续包年
	Status          int    `json:"status"`            // 充电状态。1
	NextExecuteTime int64  `json:"next_execute_time"` // 下次续费时间。秒级时间戳
	SignedTime      int64  `json:"signed_time"`       // 签约时间。秒级时间戳
	SignedPrice     int64  `json:"signed_price"`      // 下次续费金额。单位为千分之一元人民币
	PayChannel      int    `json:"pay_channel"`       // 签约平台。2：微信支付，4：支付宝
	Period          int    `json:"period"`            // 下次充电天数
	MobileApp       string `json:"mobile_app"`        // 充电渠道。可为 android 等
}

// GetElectricMonthlyListUpItem 是包月充电列表 item 数组中的充电档位对象。
type GetElectricMonthlyListUpItem struct {
	PrivilegeType int                           `json:"privilege_type"` // 充电档位代码
	Icon          string                        `json:"icon"`           // 充电图标
	Name          string                        `json:"name"`           // 充电档位名称
	ExpireTime    int64                         `json:"expire_time"`    // 该档位过期时间。秒级时间戳
	Renew         *GetElectricMonthlyListRenew  `json:"renew"`          // 充电自动续费详情。关闭自动续费时为 null
	StartTime     int64                         `json:"start_time"`     // 该档位生效时间。秒级时间戳
	RenewList     []GetElectricMonthlyListRenew `json:"renew_list"`     // 充电自动续费列表。关闭自动续费时为 null
}

// GetElectricMonthlyListUp 是包月充电列表 list 数组中的对象。
type GetElectricMonthlyListUp struct {
	UpUID          int                            `json:"up_uid"`           // 充电UP主mid
	UserName       string                         `json:"user_name"`        // 充电UP主昵称
	UserFace       string                         `json:"user_face"`        // 充电UP主头像url
	Item           []GetElectricMonthlyListUpItem `json:"item"`             // 充电详情
	Start          int64                          `json:"start"`            // 开始充电时间。秒级时间戳
	HighLevelState int                            `json:"high_level_state"` // 是否可对UP主进行高档充电
	ElecReplyState int                            `json:"elec_reply_state"` // 是否可对UP主进行专属问答。0：否，1：是，2：状态未知
}

// GetElectricMonthlyListResult 是获取包月充电列表的 data 对象。
type GetElectricMonthlyListResult struct {
	List      []GetElectricMonthlyListUp `json:"list"`       // 包月充电UP主列表。最多10个，无内容时为 null
	Page      int                        `json:"page"`       // 当前页数
	PageSize  int                        `json:"page_size"`  // 当前分页大小。一般为10
	TotalPage int                        `json:"total_page"` // 总页数
	TotalNum  int                        `json:"total_num"`  // 用户总数
	IsMore    int                        `json:"is_more"`    // 是否有更多用户。0：否，1：是
}

// GetElectricMonthlyList 获取自己正在使用或已过期的包月充电列表。
func (c *Client) GetElectricMonthlyList(ctx context.Context, param GetElectricMonthlyListParam) (*GetElectricMonthlyListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/revenue/v1/guard/getChargeRecord"
	)
	return bilibili.Execute[*GetElectricMonthlyListResult](ctx, c.Client, method, url, param, liveReferer())
}

// GetElectricMonthlyDetailParam 是获取UP主包月充电详情的请求参数。
type GetElectricMonthlyDetailParam struct {
	UpMid int `json:"up_mid" request:"query"` // 目标用户mid
}

// GetElectricMonthlyDetailRankUser 是包月充电详情中充电用户列表的对象。
type GetElectricMonthlyDetailRankUser struct {
	Rank     int    `json:"rank"`     // 充电用户索引
	Mid      int    `json:"mid"`      // 充电用户mid
	Nickname string `json:"nickname"` // 充电用户昵称
	Avatar   string `json:"avatar"`   // 充电用户头像url
}

// GetElectricMonthlyDetailUPowerRank 是包月充电详情中的充电详情对象。
type GetElectricMonthlyDetailUPowerRank struct {
	Total     int                                `json:"total"`      // 充电用户总数
	TotalDesc string                             `json:"total_desc"` // 充电总数文字说明。示例："1+"
	List      []GetElectricMonthlyDetailRankUser `json:"list"`       // 充电用户列表。最多展示30个
}

// GetElectricMonthlyDetailItem 是包月充电详情中的充电欢迎语信息。
type GetElectricMonthlyDetailItem struct {
	IntroVideoAid string `json:"intro_video_aid"` // 充电介绍视频AV号
	Welcomes      string `json:"welcomes"`        // 充电介绍语
}

// GetElectricMonthlyDetailUserCard 是包月充电详情中的UP主信息。
type GetElectricMonthlyDetailUserCard struct {
	Avatar   string `json:"avatar"`   // 用户头像url
	Nickname string `json:"nickname"` // 用户昵称
}

// GetElectricMonthlyDetailVoucherState 是包月充电券信息。
type GetElectricMonthlyDetailVoucherState struct {
	VoucherOrderID   string `json:"voucher_order_id"`   // 包月充电券订单id。实测为空
	UserVoucherState int    `json:"user_voucher_state"` // 用户包月充电券状态。实测为 0
	VipAction        int    `json:"vip_action"`         // 大会员相关动作。实测为 0
	VoucherName      string `json:"voucher_name"`       // 包月充电券名称。实测为空
	ValidatePrice    int    `json:"validate_price"`     // 生效价格。实测为 0
	ValidateLevel    int    `json:"validate_level"`     // 生效档位。实测为 0
	Activity         string `json:"activity"`           // 活动。实测为空
	CommonTime       int    `json:"common_time"`        // 普通时长。实测为 0
	HighLevelTime    int    `json:"high_level_time"`    // 高档时长。实测为 0
	SpecifyUp        int    `json:"specify_up"`         // 指定UP主。实测为 0
	SpecifyLevel     int    `json:"specify_level"`      // 指定档位。实测为 0
}

// GetElectricMonthlyDetailResult 是获取UP主包月充电详情的 data 对象。
type GetElectricMonthlyDetailResult struct {
	UPowerRank       GetElectricMonthlyDetailUPowerRank   `json:"upower_rank"`        // 充电详情
	Item             GetElectricMonthlyDetailItem         `json:"item"`               // 充电欢迎语信息
	UserCard         GetElectricMonthlyDetailUserCard     `json:"user_card"`          // UP主信息
	UPowerLevel      int                                  `json:"upower_level"`       // UP主开通的充电等级。1：非高档充电，2：高档充电
	ElecReplyState   int                                  `json:"elec_reply_state"`   // 是否可对UP主进行专属问答
	VoucherState     GetElectricMonthlyDetailVoucherState `json:"voucher_state"`      // 包月充电券信息
	UPowerRightCount map[string]int                       `json:"upower_right_count"` // 不同充电档位下的充电权益数。键为充电档位代码
	OnlyContainMedal bool                                 `json:"only_contain_medal"` // 享有的权益仅为粉丝勋章
	PrivilegeType    int                                  `json:"privilege_type"`     // 当前给该UP主包月充电的档位。从未包月充电过则为0
}

// GetElectricMonthlyDetail 获取UP主的包月充电详情。
func (c *Client) GetElectricMonthlyDetail(ctx context.Context, param GetElectricMonthlyDetailParam) (*GetElectricMonthlyDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/upower/item/detail"
	)
	return bilibili.Execute[*GetElectricMonthlyDetailResult](ctx, c.Client, method, url, param)
}

// GetElectricMonthlyRelationParam 是获取与UP主包月充电关系的请求参数。
type GetElectricMonthlyRelationParam struct {
	UpMid int `json:"up_mid" request:"query"` // 目标用户mid
}

// GetElectricMonthlyRelationUpCard 是包月充电关系中的UP主信息。
type GetElectricMonthlyRelationUpCard struct {
	Mid           int    `json:"mid"`            // UP主mid
	Nickname      string `json:"nickname"`       // UP主昵称
	OfficialTitle string `json:"official_title"` // UP主认证信息
	Avatar        string `json:"avatar"`         // UP主头像url
}

// GetElectricMonthlyRelationUserCard 是包月充电关系中的自己的信息。
type GetElectricMonthlyRelationUserCard struct {
	Avatar   string `json:"avatar"`   // 用户头像url
	Nickname string `json:"nickname"` // 用户昵称
}

// GetElectricMonthlyRelationRankUser 是包月充电关系中充电用户列表的对象。
type GetElectricMonthlyRelationRankUser struct {
	Rank     int    `json:"rank"`     // 充电用户索引
	Mid      int    `json:"mid"`      // 充电用户mid
	Nickname string `json:"nickname"` // 充电用户昵称
	Avatar   string `json:"avatar"`   // 充电用户头像url
}

// GetElectricMonthlyRelationUPowerRank 是包月充电关系中的充电详情对象。
type GetElectricMonthlyRelationUPowerRank struct {
	Total     int                                  `json:"total"`      // 充电用户总数
	TotalDesc string                               `json:"total_desc"` // 充电总数文字说明。示例："1+"
	List      []GetElectricMonthlyRelationRankUser `json:"list"`       // 充电用户列表。最多展示6个
}

// GetElectricMonthlyRelationChallengeInfo 是充电挑战信息。
type GetElectricMonthlyRelationChallengeInfo struct {
	ChallengeID         string `json:"challenge_id"`          // 充电挑战id。实测为空
	Description         string `json:"description"`           // 充电挑战描述。实测为空
	ChallengeType       int    `json:"challenge_type"`        // 充电挑战类型。实测为 0
	RemainingDays       int    `json:"remaining_days"`        // 剩余天数。实测为 0
	EndTime             string `json:"end_time"`              // 结束时间。实测为空
	Progress            int    `json:"progress"`              // 进度。实测为 0
	Targets             []any  `json:"targets"`               // 目标列表。文档未给出元素结构，实测为空数组
	State               int    `json:"state"`                 // 状态。实测为 0
	EndTimeUnix         int64  `json:"end_time_unix"`         // 结束时间。Unix秒级时间戳，实测为 0
	PubDyn              int    `json:"pub_dyn"`               // 是否发布动态。实测为 0
	DynContent          string `json:"dyn_content"`           // 动态内容。实测为空
	ChallengeLevel      int    `json:"challenge_level"`       // 充电挑战档位。实测为 0
	ChallengeLevelPrice int    `json:"challenge_level_price"` // 充电挑战档位价格。实测为 0
}

// GetElectricMonthlyRelationResult 是获取与UP主包月充电关系的 data 对象。
type GetElectricMonthlyRelationResult struct {
	Days             int                                     `json:"days"`               // 已保持多少天包月充电状态
	UpCard           GetElectricMonthlyRelationUpCard        `json:"up_card"`            // UP主信息
	UserCard         GetElectricMonthlyRelationUserCard      `json:"user_card"`          // 自己的信息
	RemainDays       int                                     `json:"remain_days"`        // 剩余天数。未处于包月充电状态为-1
	RemainLess1Day   int                                     `json:"remain_less_1day"`   // 剩余的天数是否小于1天。0：否，1：是，未处于包月充电状态为0
	UPowerRank       GetElectricMonthlyRelationUPowerRank    `json:"upower_rank"`        // 充电详情
	UPowerIcon       string                                  `json:"upower_icon"`        // 充电图标url。仅在处于包月充电状态时有内容
	UPowerRightCount int                                     `json:"upower_right_count"` // 当前自己享有该UP主的充电权益数
	OnlyContainMedal bool                                    `json:"only_contain_medal"` // 享有的权益仅为粉丝勋章
	PrivilegeType    int                                     `json:"privilege_type"`     // 当前给该UP主包月充电的档位代码。未处于包月充电状态为0
	ChallengeInfo    GetElectricMonthlyRelationChallengeInfo `json:"challenge_info"`     // 充电挑战信息
}

// GetElectricMonthlyRelation 获取自己与UP主的包月充电关系。
func (c *Client) GetElectricMonthlyRelation(ctx context.Context, param GetElectricMonthlyRelationParam) (*GetElectricMonthlyRelationResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/upower/charge/follow/info"
	)
	return bilibili.Execute[*GetElectricMonthlyRelationResult](ctx, c.Client, method, url, param)
}

// GetElectricMonthlyRankParam 是获取包月充电用户排名的请求参数。
type GetElectricMonthlyRankParam struct {
	UpMid         int `json:"up_mid" request:"query"`                             // 目标用户mid
	Ps            int `json:"ps" request:"query"`                                 // 每页项数。最大为101
	Pn            int `json:"pn" request:"query"`                                 // 页码
	PrivilegeType int `json:"privilege_type,omitempty" request:"query,omitempty"` // 充电档位代码。默认为可显示排名的档位中最高定价的档位
}

// GetElectricMonthlyRankUpInfo 是排名接口中的UP主信息。
type GetElectricMonthlyRankUpInfo struct {
	Mid            int    `json:"mid"`              // UP主mid
	Nickname       string `json:"nickname"`         // UP主昵称
	Avatar         string `json:"avatar"`           // UP主头像url
	Type           int    `json:"type"`             // UP主认证类型。-1：无，0：UP主认证，1：机构认证
	Title          string `json:"title"`            // UP主认证文字
	UPowerState    int    `json:"upower_state"`     // UP主充电功能开启状态。0：未开通充电功能，1：已开通自定义充电，2：已开通包月、自定义充电，3：已开通包月高档、自定义充电
	UPowerNumShow  int    `json:"upower_num_show"`  // 是否显示充电人数。实测为 1
	UPowerRankShow int    `json:"upower_rank_show"` // 是否显示充电排名。实测为 1
}

// GetElectricMonthlyRankItem 是排名接口 rank_info 数组中的充电用户。
type GetElectricMonthlyRankItem struct {
	Mid        int    `json:"mid"`         // 充电用户mid
	Nickname   string `json:"nickname"`    // 充电用户昵称
	Avatar     string `json:"avatar"`      // 充电用户头像url
	Rank       int    `json:"rank"`        // 包月充电排名。不在该充电档位用户列表里为-1
	Day        int    `json:"day"`         // 包月充电天数
	ExpireAt   int64  `json:"expire_at"`   // 包月充电过期时间。秒级时间戳，rank_info 中恒为0
	RemainDays int    `json:"remain_days"` // 剩余天数。rank_info 中恒为0，该档位充电状态已失效为0
}

// GetElectricMonthlyRankUserInfo 是排名接口中自己与UP主的充电关系。
type GetElectricMonthlyRankUserInfo struct {
	Mid        int    `json:"mid"`         // 用户mid
	Nickname   string `json:"nickname"`    // 用户昵称
	Avatar     string `json:"avatar"`      // 用户头像url
	Rank       int    `json:"rank"`        // 包月充电排名。不在该充电档位用户列表里为-1
	Day        int    `json:"day"`         // 包月充电天数
	ExpireAt   int64  `json:"expire_at"`   // 包月充电过期时间。秒级时间戳，若从未给UP主在该档位下包月充电过为0
	RemainDays int    `json:"remain_days"` // 剩余天数。若该档位充电状态已失效为0
}

// GetElectricMonthlyRankLevelInfo 是可显示排名的充电档位信息。
type GetElectricMonthlyRankLevelInfo struct {
	PrivilegeType int    `json:"privilege_type"` // 当前充电档位代码
	Name          string `json:"name"`           // 档位名称
	Price         int    `json:"price"`          // 档位价格。单位为百分之一元人民币
	MemberTotal   int    `json:"member_total"`   // 当前档位的用户总数
}

// GetElectricMonthlyRankResult 是获取包月充电用户排名的 data 对象。
type GetElectricMonthlyRankResult struct {
	UpInfo        GetElectricMonthlyRankUpInfo      `json:"up_info"`        // UP主信息
	RankInfo      []GetElectricMonthlyRankItem      `json:"rank_info"`      // 当前档位的充电用户排名
	UserInfo      GetElectricMonthlyRankUserInfo    `json:"user_info"`      // 自己在该档位下与UP主的充电关系
	MemberTotal   int                               `json:"member_total"`   // 当前档位的充电用户总数
	PrivilegeType int                               `json:"privilege_type"` // 当前充电档位代码
	IsCharge      bool                              `json:"is_charge"`      // 自己是否给该UP主包月充电过。无论档位
	Tabs          []int                             `json:"tabs"`           // 可显示排名的充电档位代码列表
	LevelInfo     []GetElectricMonthlyRankLevelInfo `json:"level_info"`     // 可显示排名的充电档位信息
}

// GetElectricMonthlyRank 获取给UP主包月充电的用户排名。
func (c *Client) GetElectricMonthlyRank(ctx context.Context, param GetElectricMonthlyRankParam) (*GetElectricMonthlyRankResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/upower/up/member/rank/v2"
	)
	return bilibili.Execute[*GetElectricMonthlyRankResult](ctx, c.Client, method, url, param)
}
