package garb

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 装扮/收藏集相关接口。各接口的返回体独立定义，互不复用。

// SearchGarbParam 是搜索装扮/收藏集的请求参数。
type SearchGarbParam struct {
	KeyWord string `json:"key_word,omitempty" request:"query,omitempty"` // 关键词
	Ps      int    `json:"ps,omitempty" request:"query,omitempty"`       // 每页返回数据的最大值。默认为20
	Pn      int    `json:"pn,omitempty" request:"query,omitempty"`       // 当前页数。默认为1
}

// SearchGarbResult 是搜索装扮/收藏集的 data 对象。
type SearchGarbResult struct {
	List  []SearchGarbItem `json:"list"`  // 返回数据。若无则为 null
	Ps    int              `json:"ps"`    // 每页返回数据的最大值。默认为20
	Pn    int              `json:"pn"`    // 当前页数。默认为1
	Total int              `json:"total"` // 查询到数据的总个数
}

// SearchGarbItem 是搜索装扮/收藏集列表中的一项。
type SearchGarbItem struct {
	ItemID          int64                     `json:"item_id"`          // 装扮对应的id。收藏集均为0，需要在 properties 获取
	Name            string                    `json:"name"`             // 装扮/收藏集名称
	GroupID         int64                     `json:"group_id"`         // ?分类的id
	GroupName       string                    `json:"group_name"`       // ?分类的名称
	PartID          int                       `json:"part_id"`          // 类型id。装扮为6，收藏集为0
	State           string                    `json:"state"`            // 当前状态。默认为 active
	Properties      *SearchGarbItemProperties `json:"properties"`       // 装扮/收藏集具体属性
	CurrentActivity *SearchGarbItemActivity   `json:"current_activity"` // 当前活动。实测为 null
	NextActivity    *SearchGarbItemActivity   `json:"next_activity"`    // 下次活动。文档标注为 int，实测为对象或 null
	CurrentSources  any                       `json:"current_sources"`  // ?活动开始时间。实测为 null
	FinishSources   any                       `json:"finish_sources"`   // ?活动结束时间。实测为 null
	SaleLeftTime    int64                     `json:"sale_left_time"`   // ?销售剩余时间
	SaleTimeEnd     int64                     `json:"sale_time_end"`    // ?销售结束时间
	SaleSurplus     int                       `json:"sale_surplus"`     // 未知
	SaleCountDesc   string                    `json:"sale_count_desc"`  // 销售量显示文本
	TotalCountDesc  string                    `json:"total_count_desc"` // 总量显示文本
	Tag             string                    `json:"tag"`              // 当前状态标签
	JumpLink        string                    `json:"jump_link"`        // 跳转链接
	SalesMode       int                       `json:"sales_mode"`       // ?销售状态
	DynamicTag      *SearchGarbItemDynamicTag `json:"dynamic_tag"`      // 动态标签
	TrackingInfo    string                    `json:"tracking_info"`    // 埋点信息。实测为空
}

// SearchGarbItemProperties 是搜索结果列表项中的 properties 对象。
// 该对象同时承载装扮与收藏集的字段，字段均为字符串。
type SearchGarbItemProperties struct {
	// 共有字段
	ImageCover       string `json:"image_cover"`         // 图片封面
	SaleBpForeverRaw string `json:"sale_bp_forever_raw"` // 价格。以0.01B币为单位
	Type             string `json:"type"`                // 类型。收藏集为 dlc_act，装扮为 ip

	// 装扮特有
	Desc                    string `json:"desc"`                       // 介绍文本。仅装扮
	FanDesc                 string `json:"fan_desc"`                   // 装扮名称。仅装扮
	FanID                   string `json:"fan_id"`                     // 装扮id。仅装扮
	FanItemIDs              string `json:"fan_item_ids"`               // 未知。仅装扮
	FanMid                  string `json:"fan_mid"`                    // 该up的uid。仅装扮
	FanNoColor              string `json:"fan_no_color"`               // 十六进制颜色。仅装扮
	FanRecommendDesc        string `json:"fan_recommend_desc"`         // 装备说明。仅装扮
	FanRecommendJumpType    string `json:"fan_recommend_jump_type"`    // 跳转类型。仅装扮
	FanRecommendJumpValue   string `json:"fan_recommend_jump_value"`   // 跳转值。仅装扮
	FanShareImage           string `json:"fan_share_image"`            // 分享时的背景图。仅装扮
	GrayRule                string `json:"gray_rule"`                  // 未知。仅装扮
	GrayRuleType            string `json:"gray_rule_type"`             // 未知。仅装扮
	ImageCoverColor         string `json:"image_cover_color"`          // ?图片封面纯色背景。仅装扮
	ImageCoverLong          string `json:"image_cover_long"`           // 长封面图。仅装扮
	ImageDesc               string `json:"image_desc"`                 // 介绍图片。仅装扮
	IsHide                  string `json:"is_hide"`                    // 是否隐藏。仅装扮
	ItemIDCard              string `json:"item_id_card"`               // 装扮背景卡片id。仅装扮
	ItemIDEmoji             string `json:"item_id_emoji"`              // 装扮表情包id。仅装扮
	ItemIDPendant           string `json:"item_id_pendant"`            // 装扮头像挂件id。仅装扮
	ItemIDThumbup           string `json:"item_id_thumbup"`            // 装扮点赞动画id。仅装扮
	OpenPlatformVipDiscount string `json:"open_platform_vip_discount"` // 是否有大会员减免。仅装扮
	OwnerUID                string `json:"owner_uid"`                  // UID。默认为虚拟主播衍生品小货架，仅装扮
	RankInvestorShow        string `json:"rank_investor_show"`         // 未知。仅装扮
	RealnameAuth            string `json:"realname_auth"`              // ?是否需要实名认证。仅装扮
	SaleBpPmRaw             string `json:"sale_bp_pm_raw"`             // 该装扮基础套餐价格。以0.01B币为单位，仅装扮
	SaleBuyNumLimit         string `json:"sale_buy_num_limit"`         // 购买限额。仅装扮
	SaleQuantity            string `json:"sale_quantity"`              // 该装扮粉丝专属套餐限额。仅装扮
	SaleQuantityLimit       string `json:"sale_quantity_limit"`        // 该装扮粉丝专属套餐是否限额。仅装扮
	SaleRegionIPLimit       string `json:"sale_region_ip_limit"`       // 该装扮限制购买地区。仅装扮
	SaleReserveSwitch       string `json:"sale_reserve_switch"`        // 未知。仅装扮
	SaleTimeBegin           string `json:"sale_time_begin"`            // 开始售卖时的时间戳。仅装扮
	SaleType                string `json:"sale_type"`                  // 售卖类型。默认为 pay，仅装扮
	SuitCardType            string `json:"suit_card_type"`             // 装扮卡片类型。仅装扮

	// 收藏集特有
	BookAmount             string `json:"book_amount"`               // 购买总数。仅收藏集
	DlcActID               string `json:"dlc_act_id"`                // 收藏集活动id。仅收藏集
	DlcActStatus           string `json:"dlc_act_status"`            // 收藏集活动状态。仅收藏集
	DlcIsFree              string `json:"dlc_is_free"`               // 收藏集抽奖是否免费。仅收藏集
	DlcLotteryID           string `json:"dlc_lottery_id"`            // 收藏集抽奖id。仅收藏集
	DlcLotterySaleQuantity string `json:"dlc_lottery_sale_quantity"` // 购买总数。仅收藏集
	DlcLotteryType         string `json:"dlc_lottery_type"`          // ?抽奖类型。仅收藏集
	DlcSaleEndTime         string `json:"dlc_sale_end_time"`         // 收藏集抽奖结束时间。仅收藏集
	DlcSaleMode            string `json:"dlc_sale_mode"`             // 未知。仅收藏集
	DlcSaleStartTime       string `json:"dlc_sale_start_time"`       // 收藏集抽奖开始时间。仅收藏集
	DlcSurplusStock        string `json:"dlc_surplus_stock"`         // 未知。仅收藏集
}

// SearchGarbItemActivity 是搜索结果列表项的活动信息，current_activity 与 next_activity 共用。
type SearchGarbItemActivity struct {
	Type           string `json:"type"`             // 当前永久价格活动类型。装扮一般为 open_platform_vip_discount，收藏集一般为 first_draw_discount
	TimeLimit      bool   `json:"time_limit"`       // 是否存在时间限制
	TimeLeft       int64  `json:"time_left"`        // 剩余时间
	Tag            string `json:"tag"`              // 显示标签
	PriceBpForever int    `json:"price_bp_forever"` // 永久价格。以0.01B币为单位
	PriceBpMonth   int    `json:"price_bp_month"`   // 一个月的价格。以0.01B币为单位
	TypeMonth      string `json:"type_month"`       // 当前一个月的价格活动类型。仅装扮
	TagMonth       string `json:"tag_month"`        // 显示标签。仅装扮
	TimeLimitMonth bool   `json:"time_limit_month"` // 是否存在时间限制。仅装扮
	TimeLeftMonth  int64  `json:"time_left_month"`  // 剩余时间。仅装扮
}

// SearchGarbItemDynamicTag 是搜索结果列表项中的动态标签。
type SearchGarbItemDynamicTag struct {
	Text    string `json:"text"`     // 标签文字。实测为 出框头像
	BgColor string `json:"bg_color"` // 标签背景颜色。实测为 #00AEEC
	BgImg   string `json:"bg_img"`   // 标签背景图片。实测为图片链接
}

// SearchGarb 搜索装扮/收藏集。
func (c *Client) SearchGarb(ctx context.Context, param SearchGarbParam) (*SearchGarbResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/garb/v2/mall/home/search"
	)
	return bilibili.Execute[*SearchGarbResult](ctx, c.Client, method, url, param)
}

// GetLotteryHomeDetailParam 是获取收藏集信息的请求参数。
type GetLotteryHomeDetailParam struct {
	ActID     int `json:"act_id" request:"query"`                         // 收藏集活动id
	LotteryID int `json:"lottery_id,omitempty" request:"query,omitempty"` // 收藏集抽奖id。缺了不返回数据
}

// GetLotteryHomeDetailResult 是收藏集信息的 data 对象。
type GetLotteryHomeDetailResult struct {
	LotteryID        int                              `json:"lottery_id"`        // 收藏集抽奖id
	Name             string                           `json:"name"`              // 收藏集抽奖名称
	ItemList         []GetLotteryHomeDetailItem       `json:"item_list"`         // 可抽出的物品列表
	CollectList      *GetLotteryHomeDetailCollectList `json:"collect_list"`      // 收集品列表
	ButtonBubble     any                              `json:"button_bubble"`     // 未知
	GuideInfo        any                              `json:"guide_info"`        // 未知
	IsBooked         int                              `json:"is_booked"`         // 未知
	TotalBookCnt     int                              `json:"total_book_cnt"`    // 未知
	IsFission        int                              `json:"is_fission"`        // 未知
	PhysicalExchange int                              `json:"physical_exchange"` // 未知
}

// GetLotteryHomeDetailItem 是收藏集可抽出的物品。
type GetLotteryHomeDetailItem struct {
	ItemType int                           `json:"item_type"` // 物品类型。目前只拿到1，其他值未知
	CardInfo *GetLotteryHomeDetailCardInfo `json:"card_info"` // 卡片信息
}

// GetLotteryHomeDetailCardInfo 是收藏集物品对应的卡片信息。
type GetLotteryHomeDetailCardInfo struct {
	CardTypeID            int      `json:"card_type_id"`            // 该卡片id
	CardName              string   `json:"card_name"`               // 该卡片名称
	CardImg               string   `json:"card_img"`                // 该卡片图片。无水印
	CardType              int      `json:"card_type"`               // 未知
	VideoList             []string `json:"video_list"`              // 该卡片动态视频。无水印
	IsPhysicalOrientation int      `json:"is_physical_orientation"` // 该卡片旋转方向
	CardScarcity          int      `json:"card_scarcity"`           // 该卡片稀有度
	IsMute                int      `json:"is_mute"`                 // 该卡片是否静音
	Width                 int      `json:"width"`                   // 该卡片像素宽度
	Height                int      `json:"height"`                  // 该卡片像素高度
	CardExtText           string   `json:"card_ext_text"`           // ?该卡片文件名字符串
	CardImgDownload       string   `json:"card_img_download"`       // 该卡片图片。有水印
	VideoListDownload     []string `json:"video_list_download"`     // 该卡片动态视频。有水印
	SubtitlesURL          any      `json:"subtitles_url"`           // 类型未知。示例为空字符串
	Play                  any      `json:"play"`                    // 未知
	Tag                   any      `json:"tag"`                     // 未知
	CardSubType           int      `json:"card_sub_type"`           // 未知
	IsNewTag              int      `json:"is_new_tag"`              // 未知
	IsUpTag               int      `json:"is_up_tag"`               // 未知
	IsLimitedCard         int      `json:"is_limited_card"`         // 未知
	StockInfo             any      `json:"stock_info"`              // 未知
}

// GetLotteryHomeDetailCollectList 是收藏集的收集品列表。
type GetLotteryHomeDetailCollectList struct {
	CollectInfos []GetLotteryHomeDetailCollectInfo `json:"collect_infos"` // 收集品信息
	CollectChain any                               `json:"collect_chain"` // 未知
}

// GetLotteryHomeDetailCollectInfo 是一种收集品及其兑换条件。
type GetLotteryHomeDetailCollectInfo struct {
	CollectID               int                                  `json:"collect_id"`                 // 收集品id
	StartTime               int64                                `json:"start_time"`                 // 开始时间
	EndTime                 int64                                `json:"end_time"`                   // 结束时间
	RedeemText              string                               `json:"redeem_text"`                // 兑换条件
	RedeemItemType          int                                  `json:"redeem_item_type"`           // 兑换物类型
	RedeemItemID            string                               `json:"redeem_item_id"`             // 兑换物id
	RedeemItemName          string                               `json:"redeem_item_name"`           // 兑换物名称
	RedeemItemImage         string                               `json:"redeem_item_image"`          // 兑换物预览图片
	OwnedItemAmount         int                                  `json:"owned_item_amount"`          // 拥有的数量
	RequireItemAmount       int                                  `json:"require_item_amount"`        // 需要的数量
	HasRedeemedCnt          int                                  `json:"has_redeemed_cnt"`           // 兑换次数
	EffectiveForever        int                                  `json:"effective_forever"`          // 是否永久有效
	RedeemItemImageDownload string                               `json:"redeem_item_image_download"` // 未知
	CardItem                *GetLotteryHomeDetailCollectCardItem `json:"card_item"`                  // 卡片物品。有时为 null
	JumpURL                 string                               `json:"jump_url"`                   // ?跳转链接
	RedeemCondType          string                               `json:"redeem_cond_type"`           // 当前兑换状态
	RemainStock             int                                  `json:"remain_stock"`               // 当前库存
	TotalStock              int                                  `json:"total_stock"`                // 总库存
	LotteryID               int                                  `json:"lottery_id"`                 // 抽奖id
	RewardTag               string                               `json:"reward_tag"`                 // 奖励显示标签
	RedeemDetailImage       string                               `json:"redeem_detail_image"`        // 兑换详情图片
	RedeemDetailVideos      any                                  `json:"redeem_detail_videos"`       // 未知
	Sort                    int                                  `json:"sort"`                       // 排序
	RedeemItemsOptional     any                                  `json:"redeem_items_optional"`      // 未知
	UnlockCondition         *GetLotteryHomeDetailUnlockCondition `json:"unlock_condition"`           // 解锁条件
}

// GetLotteryHomeDetailCollectCardItem 是收集品中的卡片物品。实测各字段均为 null。
type GetLotteryHomeDetailCollectCardItem struct {
	CardTypeInfo  any `json:"card_type_info"`  // 未知
	CardAssetInfo any `json:"card_asset_info"` // 未知
	Play          any `json:"play"`            // 未知
	Tag           any `json:"tag"`             // 未知
}

// GetLotteryHomeDetailUnlockCondition 是收集品的解锁条件。
type GetLotteryHomeDetailUnlockCondition struct {
	Unlocked         bool  `json:"unlocked"`          // 是否解锁
	LockType         int   `json:"lock_type"`         // 解锁类型
	ExpireAt         int64 `json:"expire_at"`         // 过期于
	UnlockedAt       int64 `json:"unlocked_at"`       // 解锁于
	UnlockThreshold  int   `json:"unlock_threshold"`  // ?解锁起点
	CurrentThreshold int   `json:"current_threshold"` // ?当前起点
}

// GetLotteryHomeDetail 获取收藏集信息。
func (c *Client) GetLotteryHomeDetail(ctx context.Context, param GetLotteryHomeDetailParam) (*GetLotteryHomeDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vas/dlc_act/lottery_home_detail"
	)
	return bilibili.Execute[*GetLotteryHomeDetailResult](ctx, c.Client, method, url, param)
}

// GetGarbSuitDetailParam 是获取主题装扮信息的请求参数。
type GetGarbSuitDetailParam struct {
	Buvid  string `json:"buvid,omitempty" request:"query,omitempty"`   // 设备唯一标识
	From   string `json:"from,omitempty" request:"query,omitempty"`    // 来源页面
	FromID int    `json:"from_id,omitempty" request:"query,omitempty"` // 来源页面id
	ItemID int    `json:"item_id" request:"query"`                     // 装扮id
	Part   string `json:"part,omitempty" request:"query,omitempty"`    // ?分类
}

// GetGarbSuitDetailResult 是主题装扮信息的 data 对象。
type GetGarbSuitDetailResult struct {
	ItemID           int64                              `json:"item_id"`           // 装扮id
	Name             string                             `json:"name"`              // 装扮名称
	GroupID          int64                              `json:"group_id"`          // 分组id
	GroupName        string                             `json:"group_name"`        // 分组名称
	PartID           int                                `json:"part_id"`           // 分类id
	State            string                             `json:"state"`             // 状态
	Properties       *GetGarbSuitDetailProperties       `json:"properties"`        // 装扮具体属性
	CurrentActivity  any                                `json:"current_activity"`  // 当前活动
	NextActivity     *GetGarbSuitDetailActivity         `json:"next_activity"`     // 下一个活动
	CurrentSources   any                                `json:"current_sources"`   // 作用尚不明确
	FinishSources    any                                `json:"finish_sources"`    // 作用尚不明确
	SaleLeftTime     int64                              `json:"sale_left_time"`    // 作用尚不明确
	SaleTimeEnd      int64                              `json:"sale_time_end"`     // 作用尚不明确
	SaleSurplus      int                                `json:"sale_surplus"`      // 商品剩余数量
	SaleCountDesc    string                             `json:"sale_count_desc"`   // 促销销量说明
	TotalCountDesc   string                             `json:"total_count_desc"`  // 总销量说明
	Tag              string                             `json:"tag"`               // 标签
	JumpLink         string                             `json:"jump_link"`         // 跳转链接
	SalesMode        int                                `json:"sales_mode"`        // 促销模式
	SuitItems        *GetGarbSuitDetailSuitItems        `json:"suit_items"`        // 装扮具体内容
	FanUser          *GetGarbSuitDetailFanUser          `json:"fan_user"`          // 装扮来源用户
	UnlockItems      any                                `json:"unlock_items"`      // 未解锁装扮
	ActivityEntrance *GetGarbSuitDetailActivityEntrance `json:"activity_entrance"` // 活动入口
	TrackingInfo     string                             `json:"tracking_info"`     // 埋点信息。实测为空
	UserData         *GetGarbSuitDetailUserData         `json:"user_data"`         // 用户数据
}

// GetGarbSuitDetailProperties 是主题装扮的 properties 对象。字段均为字符串，
// 其中若干布尔值以 "true"/"false" 字符串形式返回。
type GetGarbSuitDetailProperties struct {
	Desc                    string `json:"desc"`                       // 说明
	FanDesc                 string `json:"fan_desc"`                   // 用户说明
	FanID                   string `json:"fan_id"`                     // 装扮id。有时是数字文本，有时是普通文本
	FanItemIDs              string `json:"fan_item_ids"`               // 装扮id列表
	FanMid                  string `json:"fan_mid"`                    // 用户mid
	FanNoColor              string `json:"fan_no_color"`               // 颜色16进制字符串。作用尚不明确
	FanRecommendDesc        string `json:"fan_recommend_desc"`         // 用户推荐说明
	FanRecommendJumpType    string `json:"fan_recommend_jump_type"`    // 跳转类型
	FanRecommendJumpValue   string `json:"fan_recommend_jump_value"`   // 跳转的值。一般为该装扮所有者的个人空间链接
	FanShareImage           string `json:"fan_share_image"`            // 分享图片
	GrayRule                string `json:"gray_rule"`                  // 布尔型转换的字符串。作用尚不明确
	GrayRuleType            string `json:"gray_rule_type"`             // 作用尚不明确
	ImageCover              string `json:"image_cover"`                // 图片封面链接
	ImageCoverColor         string `json:"image_cover_color"`          // 图片封面颜色
	IsHide                  string `json:"is_hide"`                    // 是否隐藏。布尔型转换的字符串
	ItemIDCard              string `json:"item_id_card"`               // 动态卡片id
	ItemIDEmoji             string `json:"item_id_emoji"`              // 表情包id
	ItemIDThumbup           string `json:"item_id_thumbup"`            // 动态点赞特效id
	OpenPlatformVipDiscount string `json:"open_platform_vip_discount"` // 是否开启平台VIP折扣
	OwnerUID                string `json:"owner_uid"`                  // 装扮所有者的用户uid
	RankInvestorShow        string `json:"rank_investor_show"`         // ?显示投资者排名。布尔型转换的字符串
	RealnameAuth            string `json:"realname_auth"`              // 布尔型转换的字符串。作用尚不明确
	SaleBpForeverRaw        string `json:"sale_bp_forever_raw"`        // 永久价格原始值
	SaleBpPmRaw             string `json:"sale_bp_pm_raw"`             // 月度价格原始值
	SaleBuyNumLimit         string `json:"sale_buy_num_limit"`         // 促销限制数量
	SaleQuantity            string `json:"sale_quantity"`              // 促销质量。整数型转换的字符串
	SaleQuantityLimit       string `json:"sale_quantity_limit"`        // ?是否限制某些质量装扮的销售。布尔型转换的字符串
	SaleRegionIPLimit       string `json:"sale_region_ip_limit"`       // 促销限制地区
	SaleReserveSwitch       string `json:"sale_reserve_switch"`        // 布尔型转换的字符串。作用尚不明确
	SaleTimeBegin           string `json:"sale_time_begin"`            // 促销开始时间的时间戳
	SaleType                string `json:"sale_type"`                  // 促销类型
	SuitCardType            string `json:"suit_card_type"`             // 装扮卡片类型
	Type                    string `json:"type"`                       // 类型。作用尚不明确
	SaleSkuID1              string `json:"sale_sku_id_1"`              // 售卖SKU id 1。实测为 421932
	SaleSkuID2              string `json:"sale_sku_id_2"`              // 售卖SKU id 2。实测为 421931
	UserVasOrder            string `json:"user_vas_order"`             // 用户增值服务订单。布尔型转换的字符串，实测为 "true"
}

// GetGarbSuitDetailActivity 是主题装扮的活动信息，current_activity 与 next_activity 共用。
type GetGarbSuitDetailActivity struct {
	Type           string `json:"type"`             // 当前永久价格活动类型。装扮一般为 open_platform_vip_discount，收藏集一般为 first_draw_discount
	TimeLimit      bool   `json:"time_limit"`       // 是否存在时间限制
	TimeLeft       int64  `json:"time_left"`        // 剩余时间
	Tag            string `json:"tag"`              // 显示标签
	PriceBpForever int    `json:"price_bp_forever"` // 永久价格。以0.01B币为单位
	PriceBpMonth   int    `json:"price_bp_month"`   // 一个月的价格。以0.01B币为单位
	TypeMonth      string `json:"type_month"`       // 当前一个月的价格活动类型。仅装扮
	TagMonth       string `json:"tag_month"`        // 显示标签。仅装扮
	TimeLimitMonth bool   `json:"time_limit_month"` // 是否存在时间限制。仅装扮
	TimeLeftMonth  int64  `json:"time_left_month"`  // 剩余时间。仅装扮
}

// GetGarbSuitDetailSuitItems 是主题装扮包含的各类子装扮。文档可能不全，后续会继续补充。
type GetGarbSuitDetailSuitItems struct {
	Card         []GetGarbSuitDetailSuitItem `json:"card"`          // 动态卡片
	EmojiPackage []GetGarbSuitDetailSuitItem `json:"emoji_package"` // 表情包
	CardBg       []GetGarbSuitDetailSuitItem `json:"card_bg"`       // 专属评论装扮
	Thumbup      []GetGarbSuitDetailSuitItem `json:"thumbup"`       // 动态点赞特效
	Loading      []GetGarbSuitDetailSuitItem `json:"loading"`       // 专属加载动画
	PlayIcon     []GetGarbSuitDetailSuitItem `json:"play_icon"`     // 专属进度条
	Skin         []GetGarbSuitDetailSuitItem `json:"skin"`          // 专属个性主题
	SpaceBg      []GetGarbSuitDetailSuitItem `json:"space_bg"`      // 专属空间海报
}

// GetGarbSuitDetailSuitItem 是 GetGarbSuitDetailSuitItems 中各数组的通用对象，不同类别的差异集中在 properties。
type GetGarbSuitDetailSuitItem struct {
	ItemID          int64                            `json:"item_id"`          // 装扮id
	Name            string                           `json:"name"`             // 装扮名称
	State           string                           `json:"state"`            // 状态
	TabID           int                              `json:"tab_id"`           // 分栏id
	SuitItemID      int64                            `json:"suit_item_id"`     // 所属装扮的id
	Properties      *GetGarbSuitDetailItemProperties `json:"properties"`       // 装扮具体属性
	CurrentActivity any                              `json:"current_activity"` // 当前活动
	NextActivity    *GetGarbSuitDetailActivity       `json:"next_activity"`    // 下一个活动
	CurrentSources  any                              `json:"current_sources"`  // 作用尚不明确
	FinishSources   any                              `json:"finish_sources"`   // 作用尚不明确
	SaleLeftTime    int64                            `json:"sale_left_time"`   // 作用尚不明确
	SaleTimeEnd     int64                            `json:"sale_time_end"`    // 作用尚不明确
	SaleSurplus     int                              `json:"sale_surplus"`     // 商品剩余数量
	Items           []GetGarbSuitDetailSubItem       `json:"items"`            // 装扮的具体内容。文档标注为 str，实测为对象数组
}

// GetGarbSuitDetailSubItem 是表情包等子装扮中的单个元素。
type GetGarbSuitDetailSubItem struct {
	ItemID          int64                            `json:"item_id"`          // 装扮id
	Name            string                           `json:"name"`             // 装扮名称
	State           string                           `json:"state"`            // 状态
	TabID           int                              `json:"tab_id"`           // 分栏id
	SuitItemID      int64                            `json:"suit_item_id"`     // 所属装扮的id
	Properties      *GetGarbSuitDetailItemProperties `json:"properties"`       // 装扮具体属性
	CurrentActivity any                              `json:"current_activity"` // 当前活动
	NextActivity    any                              `json:"next_activity"`    // 下一个活动
	CurrentSources  any                              `json:"current_sources"`  // 作用尚不明确
	FinishSources   any                              `json:"finish_sources"`   // 作用尚不明确
	SaleLeftTime    int64                            `json:"sale_left_time"`   // 作用尚不明确
	SaleTimeEnd     int64                            `json:"sale_time_end"`    // 作用尚不明确
	SaleSurplus     int                              `json:"sale_surplus"`     // 商品剩余数量
}

// GetGarbSuitDetailItemProperties 是子装扮的 properties 对象。
// 各类子装扮的额外字段合并于此，字段均为字符串。
type GetGarbSuitDetailItemProperties struct {
	// 共有字段
	GrayRule          string `json:"gray_rule"`           // 布尔型转换的字符串。作用尚不明确
	GrayRuleType      string `json:"gray_rule_type"`      // 作用尚不明确
	RealnameAuth      string `json:"realname_auth"`       // 布尔型转换的字符串。作用尚不明确
	SaleType          string `json:"sale_type"`           // 促销类型
	Image             string `json:"image"`               // 图片
	ImagePreviewSmall string `json:"image_preview_small"` // 预览图

	// 表情包额外字段
	Addable              string `json:"addable"`                 // 布尔型转换的字符串。作用尚不明确
	Biz                  string `json:"biz"`                     // 作用尚不明确
	IsSymbol             string `json:"is_symbol"`               // 布尔型转换的字符串。作用尚不明确
	Permanent            string `json:"permanent"`               // 是否永久
	Preview              string `json:"preview"`                 // 布尔型转换的字符串。作用尚不明确
	RecentlyUsed         string `json:"recently_used"`           // 布尔型转换的字符串。作用尚不明确
	Recommend            string `json:"recommend"`               // 是否推荐
	RefMid               string `json:"ref_mid"`                 // 关联用户mid
	Removable            string `json:"removable"`               // 是否可移除
	SettingPannelNotShow string `json:"setting_pannel_not_show"` // 布尔型转换的字符串。作用尚不明确
	Size                 string `json:"size"`                    // 尺寸
	Sortable             string `json:"sortable"`                // 排序类型
	ItemIDs              string `json:"item_ids"`                // 表情包id列表。逗号分隔的id字符串

	// 加载动画额外字段
	LoadingFrameURL string `json:"loading_frame_url"` // 进度条动画的其中一帧
	LoadingURL      string `json:"loading_url"`       // 进度条动画

	// 进度条额外字段
	DragLeftPng     string `json:"drag_left_png"`     // 进度条向左拖动时的图片
	DragRightPng    string `json:"drag_right_png"`    // 进度条向右拖动时的图片
	MiddlePng       string `json:"middle_png"`        // 进度条暂停时的图片
	SquaredImage    string `json:"squared_image"`     // 效果图
	StaticIconImage string `json:"static_icon_image"` // 静态图标

	// 个性主题额外字段
	Color                    string `json:"color"`                         // 前景色。十六进制颜色值
	ColorMode                string `json:"color_mode"`                    // 颜色模式。light：亮色，dark：暗色
	ColorSecondPage          string `json:"color_second_page"`             // 背景色。十六进制颜色值
	HeadBg                   string `json:"head_bg"`                       // 首页顶部图片
	HeadMyselfMp4Play        string `json:"head_myself_mp4_play"`          // 个人空间顶部视频动画的播放类型
	HeadMyselfSquaredBg      string `json:"head_myself_squared_bg"`        // 个人空间顶部图片
	HeadTabBg                string `json:"head_tab_bg"`                   // 首页顶部标签栏背景图
	ImageCover               string `json:"image_cover"`                   // 封面图
	ImagePreview             string `json:"image_preview"`                 // 预览图。实测为图片链接
	PackageMd5               string `json:"package_md5"`                   // 装扮图包的md5值
	PackageURL               string `json:"package_url"`                   // 装扮图包的压缩包链接
	SkinMode                 string `json:"skin_mode"`                     // 皮肤模式
	TailBg                   string `json:"tail_bg"`                       // 首页底部图片
	TailColor                string `json:"tail_color"`                    // 首页底部颜色
	TailColorSelected        string `json:"tail_color_selected"`           // 首页底部被选中时的颜色
	TailIconAni              string `json:"tail_icon_ani"`                 // 首页底部是否播放动画
	TailIconAniMode          string `json:"tail_icon_ani_mode"`            // 首页底部动画的播放类型
	TailIconChannel          string `json:"tail_icon_channel"`             // 首页底部“动态”按钮图片
	TailIconDynamic          string `json:"tail_icon_dynamic"`             // 首页底部“发布动态”按钮图片
	TailIconMain             string `json:"tail_icon_main"`                // 首页底部“首页”按钮图片
	TailIconMode             string `json:"tail_icon_mode"`                // 首页底部图标模式
	TailIconMyself           string `json:"tail_icon_myself"`              // 首页底部“我的”按钮图片
	TailIconPubBtnBg         string `json:"tail_icon_pub_btn_bg"`          // 首页底部“发布动态”按钮图片
	TailIconSelectedChannel  string `json:"tail_icon_selected_channel"`    // 首页底部“动态”按钮被选中时的图片
	TailIconSelectedDynamic  string `json:"tail_icon_selected_dynamic"`    // 首页底部“发布动态”按钮被选中时的图片
	TailIconSelectedMain     string `json:"tail_icon_selected_main"`       // 首页底部“首页”按钮被选中时的图片
	TailIconSelectedMyself   string `json:"tail_icon_selected_myself"`     // 首页底部“我的”按钮被选中时的图片
	TailIconSelectedPubBtnBg string `json:"tail_icon_selected_pub_btn_bg"` // 首页底部“发布动态”按钮被选中时的图片
	TailIconSelectedShop     string `json:"tail_icon_selected_shop"`       // 首页底部“会员购”按钮被选中时的图片
	TailIconShop             string `json:"tail_icon_shop"`                // 首页底部“会员购”按钮图片

	// 空间海报额外字段
	FanNoColor      string `json:"fan_no_color"`     // 无粉丝专属颜色
	FanNoImage      string `json:"fan_no_image"`     // 无粉丝专属图片。实测为图片链接
	Image1Landscape string `json:"image1_landscape"` // 第一张空间海报
	Image1Portrait  string `json:"image1_portrait"`  // 第一张空间海报（纵向）
	Image2Landscape string `json:"image2_landscape"` // 第二张空间海报。横版
	Image2Portrait  string `json:"image2_portrait"`  // 第二张空间海报（纵向）。纵版
	Image3Landscape string `json:"image3_landscape"` // 第三张空间海报。横版
	Image3Portrait  string `json:"image3_portrait"`  // 第三张空间海报（纵向）。纵版

	// 点赞特效额外字段
	ImageAni    string `json:"image_ani"`     // 动态点赞特效动画。实测为图片链接
	ImageAniCut string `json:"image_ani_cut"` // 动态点赞特效动画（裁剪）。实测为图片链接

	// 子装扮额外字段
	Associate string `json:"associate"` // 是否关联。布尔型转换的字符串，实测为 "false"

	// 版本
	Ver string `json:"ver"` // 版本。时间戳
}

// GetGarbSuitDetailFanUser 是装扮来源用户。
type GetGarbSuitDetailFanUser struct {
	Mid      int64  `json:"mid"`      // 用户mid
	Nickname string `json:"nickname"` // 用户昵称
	Avatar   string `json:"avatar"`   // 用户头像
}

// GetGarbSuitDetailActivityEntrance 是装扮的活动入口。
type GetGarbSuitDetailActivityEntrance struct {
	ID         int64  `json:"id"`          // 活动id
	ItemID     int64  `json:"item_id"`     // 装扮id
	Title      string `json:"title"`       // 活动标题
	ImageCover string `json:"image_cover"` // 活动封面
	JumpLink   string `json:"jump_link"`   // 跳转链接
}

// GetGarbSuitDetailUserData 是主题装扮信息中的用户数据。
type GetGarbSuitDetailUserData struct {
	IsNeverPaid  bool `json:"is_never_paid"`  // 是否从未付费
	NeedMinorAck bool `json:"need_minor_ack"` // 是否需要未成年确认
}

// GetGarbSuitDetail 获取主题装扮信息。
func (c *Client) GetGarbSuitDetail(ctx context.Context, param GetGarbSuitDetailParam) (*GetGarbSuitDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/garb/v2/mall/suit/detail"
	)
	return bilibili.Execute[*GetGarbSuitDetailResult](ctx, c.Client, method, url, param)
}

// GetGarbSuitListParam 是获取主题装扮列表的请求参数。
type GetGarbSuitListParam struct {
	GroupID  int    `json:"group_id,omitempty" request:"query,omitempty"`  // 分组id。0：装扮，22：头像挂件，5：动态卡片
	Location string `json:"location,omitempty" request:"query,omitempty"`  // 作用尚不明确
	PartID   int    `json:"part_id" request:"query"`                       // 分类id。6：装扮，1：头像挂件，2：动态卡片
	Pn       int    `json:"pn,omitempty" request:"query,omitempty"`        // 页码。不填为1
	Ps       int    `json:"ps,omitempty" request:"query,omitempty"`        // 每页的数据数量。不填为20，最大也为20
	SortType int    `json:"sort_type,omitempty" request:"query,omitempty"` // 排序方式。0：默认，1：按销量，2：按最新上架时间
	UserInfo string `json:"user_info,omitempty" request:"query,omitempty"` // 用户信息。json对象，包含 buvid 与 buvid3
}

// GetGarbSuitListResult 是主题装扮列表的 data 对象。
type GetGarbSuitListResult struct {
	Page       *GetGarbSuitListPage  `json:"page"`        // 分页信息
	List       []GetGarbSuitListItem `json:"list"`        // 装扮列表
	OffsetInfo string                `json:"offset_info"` // 补偿信息。作用尚不明确
	Group      string                `json:"group"`       // 分组。作用尚不明确
}

// GetGarbSuitListPage 是主题装扮列表的分页信息。
type GetGarbSuitListPage struct {
	Total int `json:"total"` // 装扮总数
	Pn    int `json:"pn"`    // 页码
	Ps    int `json:"ps"`    // 每页的数据数量
}

// GetGarbSuitListItem 是主题装扮列表中的一项。
type GetGarbSuitListItem struct {
	ItemID          int64                          `json:"item_id"`          // 装扮对应的id。收藏集均为0，需要在 properties 获取
	Name            string                         `json:"name"`             // 装扮/收藏集名称
	GroupID         int64                          `json:"group_id"`         // ?分类的id
	GroupName       string                         `json:"group_name"`       // ?分类的名称
	PartID          int                            `json:"part_id"`          // 类型id。装扮为6，收藏集为0
	State           string                         `json:"state"`            // 当前状态。默认为 active
	Properties      *GetGarbSuitListItemProperties `json:"properties"`       // 装扮/收藏集具体属性
	CurrentActivity *GetGarbSuitListItemActivity   `json:"current_activity"` // 当前活动。实测为 null
	NextActivity    *GetGarbSuitListItemActivity   `json:"next_activity"`    // 下次活动。文档标注为 int，实测为对象或 null
	CurrentSources  any                            `json:"current_sources"`  // ?活动开始时间。实测为 null
	FinishSources   any                            `json:"finish_sources"`   // ?活动结束时间。实测为 null
	SaleLeftTime    int64                          `json:"sale_left_time"`   // ?销售剩余时间
	SaleTimeEnd     int64                          `json:"sale_time_end"`    // ?销售结束时间
	SaleSurplus     int                            `json:"sale_surplus"`     // 未知
	SaleCountDesc   string                         `json:"sale_count_desc"`  // 销售量显示文本
	TotalCountDesc  string                         `json:"total_count_desc"` // 总量显示文本
	Tag             string                         `json:"tag"`              // 当前状态标签
	JumpLink        string                         `json:"jump_link"`        // 跳转链接
	SalesMode       int                            `json:"sales_mode"`       // ?销售状态
	DynamicTag      *GetGarbSuitListDynamicTag     `json:"dynamic_tag"`      // 动态标签
	TrackingInfo    string                         `json:"tracking_info"`    // 埋点信息。实测为空
}

// GetGarbSuitListItemProperties 是主题装扮列表项中的 properties 对象。
// 该对象同时承载装扮与收藏集的字段，字段均为字符串。
type GetGarbSuitListItemProperties struct {
	// 共有字段
	ImageCover       string `json:"image_cover"`         // 图片封面
	SaleBpForeverRaw string `json:"sale_bp_forever_raw"` // 价格。以0.01B币为单位
	Type             string `json:"type"`                // 类型。收藏集为 dlc_act，装扮为 ip

	// 装扮特有
	Desc                    string `json:"desc"`                       // 介绍文本。仅装扮
	FanDesc                 string `json:"fan_desc"`                   // 装扮名称。仅装扮
	FanID                   string `json:"fan_id"`                     // 装扮id。仅装扮
	FanItemIDs              string `json:"fan_item_ids"`               // 未知。仅装扮
	FanMid                  string `json:"fan_mid"`                    // 该up的uid。仅装扮
	FanNoColor              string `json:"fan_no_color"`               // 十六进制颜色。仅装扮
	FanRecommendDesc        string `json:"fan_recommend_desc"`         // 装备说明。仅装扮
	FanRecommendJumpType    string `json:"fan_recommend_jump_type"`    // 跳转类型。仅装扮
	FanRecommendJumpValue   string `json:"fan_recommend_jump_value"`   // 跳转值。仅装扮
	FanShareImage           string `json:"fan_share_image"`            // 分享时的背景图。仅装扮
	GrayRule                string `json:"gray_rule"`                  // 未知。仅装扮
	GrayRuleType            string `json:"gray_rule_type"`             // 未知。仅装扮
	ImageCoverColor         string `json:"image_cover_color"`          // ?图片封面纯色背景。仅装扮
	ImageCoverLong          string `json:"image_cover_long"`           // 长封面图。仅装扮
	ImageDesc               string `json:"image_desc"`                 // 介绍图片。仅装扮
	IsHide                  string `json:"is_hide"`                    // 是否隐藏。仅装扮
	ItemIDCard              string `json:"item_id_card"`               // 装扮背景卡片id。仅装扮
	ItemIDEmoji             string `json:"item_id_emoji"`              // 装扮表情包id。仅装扮
	ItemIDPendant           string `json:"item_id_pendant"`            // 装扮头像挂件id。仅装扮
	ItemIDThumbup           string `json:"item_id_thumbup"`            // 装扮点赞动画id。仅装扮
	OpenPlatformVipDiscount string `json:"open_platform_vip_discount"` // 是否有大会员减免。仅装扮
	OwnerUID                string `json:"owner_uid"`                  // UID。默认为虚拟主播衍生品小货架，仅装扮
	RankInvestorShow        string `json:"rank_investor_show"`         // 未知。仅装扮
	RealnameAuth            string `json:"realname_auth"`              // ?是否需要实名认证。仅装扮
	SaleBpPmRaw             string `json:"sale_bp_pm_raw"`             // 该装扮基础套餐价格。以0.01B币为单位，仅装扮
	SaleBuyNumLimit         string `json:"sale_buy_num_limit"`         // 购买限额。仅装扮
	SaleQuantity            string `json:"sale_quantity"`              // 该装扮粉丝专属套餐限额。仅装扮
	SaleQuantityLimit       string `json:"sale_quantity_limit"`        // 该装扮粉丝专属套餐是否限额。仅装扮
	SaleRegionIPLimit       string `json:"sale_region_ip_limit"`       // 该装扮限制购买地区。仅装扮
	SaleReserveSwitch       string `json:"sale_reserve_switch"`        // 未知。仅装扮
	SaleTimeBegin           string `json:"sale_time_begin"`            // 开始售卖时的时间戳。仅装扮
	SaleType                string `json:"sale_type"`                  // 售卖类型。默认为 pay，仅装扮
	SuitCardType            string `json:"suit_card_type"`             // 装扮卡片类型。仅装扮

	// 收藏集特有
	BookAmount             string `json:"book_amount"`               // 购买总数。仅收藏集
	DlcActID               string `json:"dlc_act_id"`                // 收藏集活动id。仅收藏集
	DlcActStatus           string `json:"dlc_act_status"`            // 收藏集活动状态。仅收藏集
	DlcIsFree              string `json:"dlc_is_free"`               // 收藏集抽奖是否免费。仅收藏集
	DlcLotteryID           string `json:"dlc_lottery_id"`            // 收藏集抽奖id。仅收藏集
	DlcLotterySaleQuantity string `json:"dlc_lottery_sale_quantity"` // 购买总数。仅收藏集
	DlcLotteryType         string `json:"dlc_lottery_type"`          // ?抽奖类型。仅收藏集
	DlcSaleEndTime         string `json:"dlc_sale_end_time"`         // 收藏集抽奖结束时间。仅收藏集
	DlcSaleMode            string `json:"dlc_sale_mode"`             // 未知。仅收藏集
	DlcSaleStartTime       string `json:"dlc_sale_start_time"`       // 收藏集抽奖开始时间。仅收藏集
	DlcSurplusStock        string `json:"dlc_surplus_stock"`         // 未知。仅收藏集
}

// GetGarbSuitListItemActivity 是主题装扮列表项的活动信息，current_activity 与 next_activity 共用。
type GetGarbSuitListItemActivity struct {
	Type           string `json:"type"`             // 当前永久价格活动类型。装扮一般为 open_platform_vip_discount，收藏集一般为 first_draw_discount
	TimeLimit      bool   `json:"time_limit"`       // 是否存在时间限制
	TimeLeft       int64  `json:"time_left"`        // 剩余时间
	Tag            string `json:"tag"`              // 显示标签
	PriceBpForever int    `json:"price_bp_forever"` // 永久价格。以0.01B币为单位
	PriceBpMonth   int    `json:"price_bp_month"`   // 一个月的价格。以0.01B币为单位
	TypeMonth      string `json:"type_month"`       // 当前一个月的价格活动类型。仅装扮
	TagMonth       string `json:"tag_month"`        // 显示标签。仅装扮
	TimeLimitMonth bool   `json:"time_limit_month"` // 是否存在时间限制。仅装扮
	TimeLeftMonth  int64  `json:"time_left_month"`  // 剩余时间。仅装扮
}

// GetGarbSuitListDynamicTag 是主题装扮列表项中的动态标签。
type GetGarbSuitListDynamicTag struct {
	Text    string `json:"text"`     // 标签文字。实测为 出框头像
	BgColor string `json:"bg_color"` // 标签背景颜色。实测为 #00AEEC
	BgImg   string `json:"bg_img"`   // 标签背景图片。实测为图片链接
}

// GetGarbSuitList 获取主题装扮列表。
func (c *Client) GetGarbSuitList(ctx context.Context, param GetGarbSuitListParam) (*GetGarbSuitListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/garb/v2/mall/partition/item/list"
	)
	return bilibili.Execute[*GetGarbSuitListResult](ctx, c.Client, method, url, param)
}

// GetLotteryActListParam 是获取收藏集列表的请求参数。
type GetLotteryActListParam struct {
	Scene int `json:"scene,omitempty" request:"query,omitempty"` // 作用尚不明确。默认为1，不填则获取到空数据
	Site  int `json:"site,omitempty" request:"query,omitempty"`  // 位置。建议填上，会影响到后面的json数据
}

// GetLotteryActListResult 是收藏集列表的 data 对象。
type GetLotteryActListResult struct {
	List   []GetLotteryActListItem `json:"list"`    // 收藏集列表
	IsMore bool                    `json:"is_more"` // 是否还有足够的收藏集
	Site   int                     `json:"site"`    // 下一页的位置。为请求参数 site 的基础上再加20
}

// GetLotteryActListItem 是收藏集列表中的一项。
type GetLotteryActListItem struct {
	ActID       int    `json:"act_id"`       // 收藏集id
	ActName     string `json:"act_name"`     // 收藏集名称
	ActPic      string `json:"act_pic"`      // 收藏集封面图片
	SalePrice   int    `json:"sale_price"`   // 收藏集价格。以0.01B币为单位
	ActDesc     string `json:"act_desc"`     // 收藏集说明
	Tag         string `json:"tag"`          // 标签
	LotteryID   int    `json:"lottery_id"`   // 抽奖id
	LotteryType int    `json:"lottery_type"` // 抽奖类型
	ActLink     string `json:"act_link"`     // 收藏集链接
}

// GetLotteryActList 获取收藏集列表。
func (c *Client) GetLotteryActList(ctx context.Context, param GetLotteryActListParam) (*GetLotteryActListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vas/dlc_act/act/list"
	)
	return bilibili.Execute[*GetLotteryActListResult](ctx, c.Client, method, url, param)
}
