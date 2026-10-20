package login

// 会员中心相关响应模型。各接口的返回体独立定义，互不复用。

// GetDailyRewardStatusResult 是每日奖励完成情况的 data 对象。
type GetDailyRewardStatusResult struct {
	Login        bool `json:"login"`         // 每日登录。false：未完成。true：已完成。完成奖励 5 经验
	Watch        bool `json:"watch"`         // 每日观看。false：未完成。true：已完成。完成奖励 5 经验
	Coins        int  `json:"coins"`         // 每日投币所奖励的经验。上限为 50。该值更新存在延迟
	Share        bool `json:"share"`         // 每日分享。false：未完成。true：已完成。完成奖励 5 经验
	Email        bool `json:"email"`         // 绑定邮箱。false：未完成。true：已完成。首次完成奖励 20 经验
	Tel          bool `json:"tel"`           // 绑定手机号。false：未完成。true：已完成。首次完成奖励 100 经验
	SafeQuestion bool `json:"safe_question"` // 设置密保问题。false：未完成。true：已完成。首次完成奖励 30 经验
	IdentifyCard bool `json:"identify_card"` // 实名认证。false：未完成。true：已完成。首次完成奖励 50 经验
}

// GetVipStatusResult 是大会员状态的 data 对象。
type GetVipStatusResult struct {
	Mid        int64 `json:"mid"`          // 我的 mid
	VipType    int   `json:"vip_type"`     // 大会员类型。0：无。1：月度。2：年度
	VipStatus  int   `json:"vip_status"`   // 大会员状态。1：正常。2：由于 IP 地址更换过于频繁，服务被冻结。3：账号风险过高，大会员功能已被锁定
	VipDueDate int64 `json:"vip_due_date"` // 大会员到期时间。毫秒时间戳
	VipPayType int   `json:"vip_pay_type"` // 是否已购买大会员。0：未购买。1：已购买
	ThemeType  int   `json:"theme_type"`   // 作用尚不明确
}

// GetAccountSecurityInfoResult 是账号安全情况的 data 对象。
type GetAccountSecurityInfoResult struct {
	AccountInfo  GetAccountSecurityInfoBindInfo     `json:"account_info"`  // 账号绑定信息
	AccountSafe  GetAccountSecurityInfoPasswordInfo `json:"account_safe"`  // 密码安全信息
	AccountSns   GetAccountSecurityInfoSnsInfo      `json:"account_sns"`   // 互联登录绑定信息
	AccountOther GetAccountSecurityInfoOtherInfo    `json:"account_other"` // 其他信息
}

// GetAccountSecurityInfoBindInfo 是账号安全情况中的账号绑定信息。
type GetAccountSecurityInfoBindInfo struct {
	HideTel           string `json:"hide_tel"`           // 绑定的手机号。星号隐藏部分信息
	HideMail          string `json:"hide_mail"`          // 绑定的邮箱。星号隐藏部分信息
	BindTel           bool   `json:"bind_tel"`           // 是否绑定手机号。false：未绑定。true：已绑定
	BindMail          bool   `json:"bind_mail"`          // 是否绑定邮箱。false：未绑定。true：已绑定
	TelVerify         bool   `json:"tel_verify"`         // 是否验证手机号。false：未验证。true：已验证
	MailVerify        bool   `json:"mail_verify"`        // 是否验证邮箱。false：未验证。true：已验证
	UnneededCheck     bool   `json:"unneeded_check"`     // 是否未设置密码。false：已设置。true：未设置
	RealnameCertified bool   `json:"realname_certified"` // 是否实名认证。false：未认证。true：已认证
}

// GetAccountSecurityInfoPasswordInfo 是账号安全情况中的密码安全信息。
type GetAccountSecurityInfoPasswordInfo struct {
	Score    int  `json:"Score"`     // 账号安全等级。0-100。已弃用？
	ScoreNew int  `json:"score_new"` // 新版账号安全等级。0-100
	PwdLevel int  `json:"pwd_level"` // 当前密码强度等级。1：弱。2：中。3：强
	Security bool `json:"security"`  // 当前密码是否安全。false：不安全。true：安全
}

// GetAccountSecurityInfoSnsInfo 是账号安全情况中的互联登录绑定信息。
type GetAccountSecurityInfoSnsInfo struct {
	WeiboBind  int `json:"weibo_bind"`  // 是否绑定微博。0：未绑定。1：已绑定
	QqBind     int `json:"qq_bind"`     // 是否绑定 QQ。0：未绑定。1：已绑定
	WechatBind int `json:"wechat_bind"` // 是否绑定微信。0：未绑定。1：已绑定
}

// GetAccountSecurityInfoOtherInfo 是账号安全情况中的账号其他信息。
type GetAccountSecurityInfoOtherInfo struct {
	SkipVerify bool `json:"skipVerify"` // 作用尚不明确
}

// GetRealnameStatusResult 是账号实名认证状态的 data 对象。
type GetRealnameStatusResult struct {
	Status int `json:"status"` // 实名认证状态。0：未认证。1：已认证
}

// GetRealnameApplyStatusResult 是实名认证详细信息的 data 对象。
type GetRealnameApplyStatusResult struct {
	Status            int                                      `json:"status"`              // 认证状态。1：已认证。3：未认证
	Remark            string                                   `json:"remark"`              // 驳回信息。默认为空
	Realname          string                                   `json:"realname"`            // 实名姓名。星号隐藏完全信息
	Card              string                                   `json:"card"`                // 证件号码。星号隐藏部分信息
	CardType          int                                      `json:"card_type"`           // 证件类型代码。0：身份证。2：港澳居民来往内地通行证。3：台湾居民来往大陆通行证。4：护照（中国签发）。5：外国人永久居留证。6：其他国家或地区身份证明
	LockOriginal      bool                                     `json:"lock_original"`       // 是否锁定原始信息。实测为 false
	UserFieldSettings []GetRealnameApplyStatusUserFieldSetting `json:"user_field_settings"` // 字段设置列表
}

// GetRealnameApplyStatusUserFieldSetting 是实名认证详细信息中的一项字段设置。
type GetRealnameApplyStatusUserFieldSetting struct {
	Field       string `json:"field"`        // 字段名。实测为 handle
	Visible     bool   `json:"visible"`      // 是否可见。实测为 false
	Editable    bool   `json:"editable"`     // 是否可编辑。实测为 false
	RuleMessage string `json:"rule_message"` // 规则提示。实测为空
}

// GetCoinLogResult 是硬币变化情况的 data 对象。
type GetCoinLogResult struct {
	List  []GetCoinLogItem `json:"list"`  // 变化记录条目列表
	Count int              `json:"count"` // 变化记录条目数
}

// GetCoinLogItem 是硬币变化情况中的一条记录。
type GetCoinLogItem struct {
	Time   string  `json:"time"`   // 变化时间。YYYY-MM-DD HH:MM:SS
	Delta  float64 `json:"delta"`  // 变化量。正值为收入，负值为支出
	Reason string  `json:"reason"` // 变化说明
}

// GetExpLogResult 是最近一周经验记录的 data 对象。
type GetExpLogResult struct {
	List  []GetExpLogItem `json:"list"`  // 经验记录条目列表
	Count int             `json:"count"` // 经验记录条目数
}

// GetExpLogItem 是经验记录中的一条记录。
type GetExpLogItem struct {
	Delta  int    `json:"delta"`  // 经验值
	Time   string `json:"time"`   // 记录时间。格式为 yyyy-MM-dd HH:mm:ss
	Reason string `json:"reason"` // 记录原因
}

// GetMoralLogResult 是最近一周节操记录的 data 对象。
type GetMoralLogResult struct {
	Moral int               `json:"moral"` // 节操值
	List  []GetMoralLogItem `json:"list"`  // 记录条目列表
	Count int               `json:"count"` // 记录条目数
}

// GetMoralLogItem 是节操记录中的一条记录。
type GetMoralLogItem struct {
	Time   string  `json:"time"`   // 时间
	Delta  float64 `json:"delta"`  // 变化
	Origin string  `json:"origin"` // 原因
	Reason string  `json:"reason"` // 说明
}
