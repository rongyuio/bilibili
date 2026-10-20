package electric

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 充电列表相关接口。各接口的返回体独立定义，互不复用。
// 文档见 electric/charge_list.md。

// GetSpaceElectricChargeListParam 是获取空间充电公示列表的请求参数。
type GetSpaceElectricChargeListParam struct {
	UpMid int `json:"up_mid"` // 目标用户mid
}

// GetSpaceElectricChargeListVipInfo 是空间充电公示列表中充电用户的会员信息。
type GetSpaceElectricChargeListVipInfo struct {
	VipDueMsec int64 `json:"vipDueMsec"` // 大会员过期时间。恒为0
	VipStatus  int   `json:"vipStatus"`  // 大会员状态。包月充电时恒为0；自定义充电：0：无，1：有
	VipType    int   `json:"vipType"`    // 大会员类型。包月充电时恒为0；自定义充电：0：无，1：月大会员，2：年度及以上大会员
}

// GetSpaceElectricChargeListItem 是空间充电公示列表 list 数组中的对象。
type GetSpaceElectricChargeListItem struct {
	Uname         string                            `json:"uname"`          // 充电用户昵称
	Avatar        string                            `json:"avatar"`         // 充电用户头像url
	Mid           int                               `json:"mid"`            // 充电对象mid
	PayMid        int                               `json:"pay_mid"`        // 充电用户mid
	Rank          int                               `json:"rank"`           // 充电用户排名。取决于充电的多少
	TrendType     int                               `json:"trend_type"`     // 0。作用尚不明确
	VipInfo       GetSpaceElectricChargeListVipInfo `json:"vip_info"`       // 充电用户会员信息
	Message       string                            `json:"message"`        // 充电留言。无为空
	MsgHidden     int                               `json:"msg_hidden"`     // 0。作用尚不明确
	MessageHidden int                               `json:"message_hidden"` // 是否隐藏留言。实测为 0
}

// GetSpaceElectricChargeListResult 是获取空间充电公示列表的 data 对象。
type GetSpaceElectricChargeListResult struct {
	Count      int                              `json:"count"`       // 本月充电人数
	List       []GetSpaceElectricChargeListItem `json:"list"`        // 本月充电用户列表。按照充电排名排列
	TotalCount int                              `json:"total_count"` // 总计充电次数
	Total      int                              `json:"total"`       // 总计充电次数。同 total_count
	SpecialDay int                              `json:"special_day"` // 0。作用尚不明确
}

// GetSpaceElectricChargeList 获取空间充电公示列表。
func (c *Client) GetSpaceElectricChargeList(ctx context.Context, param GetSpaceElectricChargeListParam) (*GetSpaceElectricChargeListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/ugcpay-rank/elec/month/up"
	)
	return bilibili.Execute[*GetSpaceElectricChargeListResult](ctx, c.Client, method, url, param)
}

// GetVideoElectricChargeListParam 是获取视频充电鸣谢名单的请求参数。
type GetVideoElectricChargeListParam struct {
	Mid  int    `json:"mid"`                                      // 目标用户mid
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 目标稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 目标稿件bvid。avid与bvid任选一个
}

// GetVideoElectricChargeListHighLevelJumpURL 是充电专属视频的跳转url信息。
type GetVideoElectricChargeListHighLevelJumpURL struct {
	UpLink         string `json:"up_link"`         // UP主包月充电购买页面url
	PaywallLink    string `json:"paywall_link"`    // 视频试看后的提示页面url
	PreviewbarLink string `json:"previewbar_link"` // 视频预览页面url
}

// GetVideoElectricChargeListHighLevel 是视频充电鸣谢名单中的充电专属视频信息。
type GetVideoElectricChargeListHighLevel struct {
	PrivilegeType int                                        `json:"privilege_type"` // 解锁视频所需最低定价档位的代码
	Title         string                                     `json:"title"`          // 提示标题。「该视频为『{充电档位名称}』专属视频」
	SubTitle      string                                     `json:"sub_title"`      // 提示子标题。「开通『{充电档位定价}元档包月充电』即可观看」
	ShowButton    bool                                       `json:"show_button"`    // 是否显示按钮
	ButtonText    string                                     `json:"button_text"`    // 按钮文本。「去开通」
	JumpURL       GetVideoElectricChargeListHighLevelJumpURL `json:"jump_url"`       // 跳转url信息。详细信息有待补充
	Intro         string                                     `json:"intro"`          // 充电介绍语
	Open          bool                                       `json:"open"`           // (?)
	New           bool                                       `json:"new"`            // (?)
	QuestionText  string                                     `json:"question_text"`  // (?)
	QaDetailLink  string                                     `json:"qa_detail_link"` // (?)
}

// GetVideoElectricChargeListShowInfo 是视频充电鸣谢名单的展示选项。
type GetVideoElectricChargeListShowInfo struct {
	Show      bool                                `json:"show"`       // 是否展示视频充电鸣谢名单
	State     int                                 `json:"state"`      // 充电功能开启状态。-1：未开通充电功能。1：已开通自定义充电。2：已开通包月、自定义充电。3：已开通包月高档、自定义充电
	Title     string                              `json:"title"`      // 充电按钮显示文字。空字符串或「充电」或「充电中」
	JumpURL   string                              `json:"jump_url"`   // 跳转url
	Icon      string                              `json:"icon"`       // 充电图标
	HighLevel GetVideoElectricChargeListHighLevel `json:"high_level"` // 充电专属视频信息
	WithQaID  int                                 `json:"with_qa_id"` // 充电问答id
}

// GetVideoElectricChargeListVipInfo 是视频充电鸣谢名单中充电用户的会员信息。
type GetVideoElectricChargeListVipInfo struct {
	VipDueMsec int64 `json:"vipDueMsec"` // 大会员过期时间。恒为0
	VipStatus  int   `json:"vipStatus"`  // 大会员状态。包月充电时恒为0；自定义充电：0：无，1：有
	VipType    int   `json:"vipType"`    // 大会员类型。包月充电时恒为0；自定义充电：0：无，1：月大会员，2：年度及以上大会员
}

// GetVideoElectricChargeListItem 是视频充电鸣谢名单 list 数组中的对象。
type GetVideoElectricChargeListItem struct {
	Mid        int                               `json:"mid"`         // 充电对象mid
	PayMid     int                               `json:"pay_mid"`     // 充电用户mid
	Rank       int                               `json:"rank"`        // 充电用户排名。取决于充电的多少
	Uname      string                            `json:"uname"`       // 充电用户昵称
	Avatar     string                            `json:"avatar"`      // 充电用户头像url
	Message    string                            `json:"message"`     // 充电留言。无为空
	MsgDeleted int                               `json:"msg_deleted"` // 0。作用尚不明确
	VipInfo    GetVideoElectricChargeListVipInfo `json:"vip_info"`    // 充电用户会员信息
	TrendType  int                               `json:"trend_type"`  // 0。作用尚不明确
}

// GetVideoElectricChargeListResult 是获取视频充电鸣谢名单的 data 对象。
type GetVideoElectricChargeListResult struct {
	ShowInfo    GetVideoElectricChargeListShowInfo `json:"show_info"`     // 展示选项
	AvCount     int                                `json:"av_count"`      // 目标视频充电人数
	Count       int                                `json:"count"`         // 本月充电人数
	TotalCount  int                                `json:"total_count"`   // 总计充电人数
	SpecialDay  int                                `json:"special_day"`   // 0。作用尚不明确
	DisplayNum  int                                `json:"display_num"`   // 0。作用尚不明确
	CntPrivType int                                `json:"cnt_priv_type"` // 0。作用尚不明确
	List        []GetVideoElectricChargeListItem   `json:"list"`          // 本月充电用户列表。按照充电排名排列
}

// GetVideoElectricChargeList 获取视频充电鸣谢名单。
func (c *Client) GetVideoElectricChargeList(ctx context.Context, param GetVideoElectricChargeListParam) (*GetVideoElectricChargeListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/elec/show"
	)
	return bilibili.Execute[*GetVideoElectricChargeListResult](ctx, c.Client, method, url, param)
}

// GetReceivedElectricChargeListParam 是查询我收到的充电列表的请求参数。
type GetReceivedElectricChargeListParam struct {
	CurrentPage int    `json:"currentPage"`                                   // 页数
	PageSize    int    `json:"pageSize"`                                      // 分页大小。取值范围[1,50]
	CustomerID  int    `json:"customerId"`                                    // (?)。目前为固定值：10026
	BeginTime   string `json:"beginTime,omitempty" request:"query,omitempty"` // 开始日期。yyyy-MM-dd
	EndTime     string `json:"endTime,omitempty" request:"query,omitempty"`   // 结束日期。yyyy-MM-dd
}

// GetReceivedElectricChargeListPage 是收到的充电列表的分页信息。
type GetReceivedElectricChargeListPage struct {
	CurrentPage int `json:"currentPage"` // 当前页数
	PageSize    int `json:"pageSize"`    // 当前分页大小
	TotalCount  int `json:"totalCount"`  // 记录总数
	TotalPage   int `json:"totalPage"`   // 总页数
}

// GetReceivedElectricChargeListItem 是收到的充电列表 result 数组中的对象。
type GetReceivedElectricChargeListItem struct {
	Mid               int     `json:"mid"`               // 充电人mid
	Name              string  `json:"name"`              // 充电人昵称
	Avatar            string  `json:"avatar"`            // 充电人头像
	OriginalThirdCoin int     `json:"originalThirdCoin"` // 原始B币数
	Brokerage         float64 `json:"brokerage"`         // 实际收到的贝壳数
	Remark            string  `json:"remark"`            // 充电渠道。Web/安卓/iOS
	Ctime             string  `json:"ctime"`             // 充电时间。yyyy-MM-dd HH:mm:ss
}

// GetReceivedElectricChargeListConfigItem 是收到的充电列表 config 数组中的对象，字段总是为 null。
type GetReceivedElectricChargeListConfigItem struct {
	Mid               *int     `json:"mid"`
	Name              *string  `json:"name"`
	Avatar            *string  `json:"avatar"`
	OriginalThirdCoin *int     `json:"originalThirdCoin"`
	Brokerage         *float64 `json:"brokerage"`
	Remark            *string  `json:"remark"`
	Ctime             *string  `json:"ctime"`
}

// GetReceivedElectricChargeListResult 是查询我收到的充电列表的 data 对象。
type GetReceivedElectricChargeListResult struct {
	Page   GetReceivedElectricChargeListPage         `json:"page"`   // 分页信息
	Result []GetReceivedElectricChargeListItem       `json:"result"` // 充电信息本体
	Config []GetReceivedElectricChargeListConfigItem `json:"config"` // (?)。数组元素字段总是为 null
}

// GetReceivedElectricChargeList 查询我收到的充电列表。
func (c *Client) GetReceivedElectricChargeList(ctx context.Context, param GetReceivedElectricChargeListParam) (*GetReceivedElectricChargeListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://pay.bilibili.com/bk/brokerage/listForCustomerRechargeRecord"
	)
	return bilibili.Execute[*GetReceivedElectricChargeListResult](ctx, c.Client, method, url, param)
}

// GetElectricChargeHistoryParam 是查询历史充电数据的请求参数。
type GetElectricChargeHistoryParam struct {
	Pn int `json:"pn,omitempty" request:"query,omitempty"` // 页数
	Ps int `json:"ps,omitempty" request:"query,omitempty"` // 分页大小。取值范围[1,20]
}

// GetElectricChargeHistoryItem 是历史充电数据 list 数组中的对象。
type GetElectricChargeHistoryItem struct {
	Aid     int    `json:"aid"`      // 0
	Bvid    string `json:"bvid"`     // 空
	ElecNum int    `json:"elec_num"` // 充电电池数
	Title   string `json:"title"`    // 空
	Uname   string `json:"uname"`    // 空
	Avatar  string `json:"avatar"`   // 空
	Ctime   string `json:"ctime"`    // 充电时间。yyyy-MM-dd HH:mm:ss
}

// GetElectricChargeHistoryPager 是历史充电数据的分页信息。
type GetElectricChargeHistoryPager struct {
	Current int `json:"current"` // 当前页数
	Size    int `json:"size"`    // 当前分页大小
	Total   int `json:"total"`   // 记录总数
}

// GetElectricChargeHistoryResult 是查询历史充电数据的 data 对象。
type GetElectricChargeHistoryResult struct {
	List  []GetElectricChargeHistoryItem `json:"list"`  // 充电信息本体
	Pager GetElectricChargeHistoryPager  `json:"pager"` // 分页信息
}

// GetElectricChargeHistory 查询历史充电数据。
func (c *Client) GetElectricChargeHistory(ctx context.Context, param GetElectricChargeHistoryParam) (*GetElectricChargeHistoryResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/h5/elec/rank/recent"
	)
	return bilibili.Execute[*GetElectricChargeHistoryResult](ctx, c.Client, method, url, param)
}
