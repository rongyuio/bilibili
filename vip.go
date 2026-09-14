package bilibili

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// 大会员相关接口。响应模型见 vip_model.go。

type ReceiveVipPrivilegeParam struct {
	Type int `json:"type"` // 兑换类型。1：B币券。2：会员购优惠券。3：漫画福利券。4：会员购包邮券。5：漫画商城优惠券。6：装扮体验卡。7：课堂优惠券
}

// ReceiveVipPrivilege 兑换大会员卡券，1：B币券，2：会员购优惠券，3：漫画福利券，4：会员购包邮券，5：漫画商城优惠券
func (c *Client) ReceiveVipPrivilege(ctx context.Context, param ReceiveVipPrivilegeParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/vip/privilege/receive"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// SignVipScore 大积分签到
func (c *Client) SignVipScore(ctx context.Context) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/pgc/activity/score/task/sign"
	)
	_, err := execute[any](ctx, c, method, url, nil, fillCsrf(c))
	return err
}

// GetVipPrivilege 卡券状态查询
func (c *Client) GetVipPrivilege(ctx context.Context) (*VipPrivilege, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vip/privilege/my"
	)
	return execute[*VipPrivilege](ctx, c, method, url, nil)
}

type GetVipCenterInfoParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token
	Platform  string `json:"platform,omitempty" request:"query,omitempty"`   // 平台表示。web端：web。安卓APP：android
	MobiApp   string `json:"mobi_app,omitempty" request:"query,omitempty"`   // APP 名称。安卓APP：android
	Build     int    `json:"build,omitempty" request:"query,omitempty"`      // 构建 id
}

// GetVipCenterInfo 获取大会员中心信息
func (c *Client) GetVipCenterInfo(ctx context.Context, param GetVipCenterInfoParam) (*VipCenterInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/vip/privilege/my"
	)
	return execute[*VipCenterInfo](ctx, c, method, url, param)
}
