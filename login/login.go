package login

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	qrcodeterminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
	"github.com/skip2/go-qrcode"
)

// 登录相关接口：验证码、密码 / 短信 / 扫码登录与我的信息。
// 响应模型见 login_model.go。各接口的返回体独立定义，互不复用。

var (
	errPublicKeyDecode = errors.New("failed to decode public key")
	errPublicKeyType   = errors.New("rsa public key type error")
)

// Captcha 申请验证码参数
func (c *Client) Captcha(ctx context.Context) (*CaptchaResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/x/passport-login/captcha"
	)
	return bilibili.Execute[*CaptchaResult](ctx, c.Client, method, url, nil, bilibili.FillParam("source", "main_web"))
}

// encryptPassword 用 B 站下发的 PEM 公钥加密「盐值 + 明文密码」。
func encryptPassword(publicKey, data string) (string, error) {
	// pem解码
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errPublicKeyDecode
	}
	// x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("解析 RSA 公钥: %w", err)
	}
	pk, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return "", errPublicKeyType
	}
	// 加密。B 站登录接口要求 PKCS#1 v1.5 填充，不能改用 OAEP。
	//nolint:staticcheck
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pk, []byte(data))
	if err != nil {
		return "", fmt.Errorf("加密密码: %w", err)
	}
	// base64
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// PasswordLoginParam 是账号密码登录的请求参数。
type PasswordLoginParam struct {
	Username  string `json:"username"`                                   // 用户登录账号。手机号或邮箱地址
	Password  string `json:"password"`                                   // 参数传入原密码，下文会自动转为加密后的带盐密码
	Keep      int    `json:"keep"`                                       // 0
	Token     string `json:"token"`                                      // 登录 API token。使用 Captcha() 方法获取
	Challenge string `json:"challenge"`                                  // 极验 challenge。使用 Captcha() 方法获取
	Validate  string `json:"validate"`                                   // 极验 result。极验验证后得到
	Seccode   string `json:"seccode"`                                    // 极验 result +jordan。极验验证后得到
	GoURL     string `json:"go_url,omitempty" request:"query,omitempty"` // 跳转 url。默认为 https://www.bilibili.com
	Source    string `json:"source,omitempty" request:"query,omitempty"` // 登录来源。main_web：独立登录页。main_mini：小窗登录
}

// LoginWithPassword 账号密码登录，其中validate, seccode字段需要在极验人机验证后获取
func (c *Client) LoginWithPassword(ctx context.Context, param PasswordLoginParam) (*PasswordLoginResult, error) {
	// 获取密码盐值
	getKey, err := c.GetLoginKey(ctx)
	if err != nil {
		return nil, err
	}

	// 密码加盐
	encryptPwd, err := encryptPassword(getKey.Key, getKey.Hash+param.Password)
	if err != nil {
		return nil, err
	}
	param.Password = encryptPwd

	// 登录操作(web端)
	const (
		method = resty.MethodPost
		url    = "https://passport.bilibili.com/x/passport-login/web/login"
	)
	return bilibili.Execute[*PasswordLoginResult](ctx, c.Client, method, url, param)
}

// GetCountryCrown 获取国际冠字码
func (c *Client) GetCountryCrown(ctx context.Context) (*GetCountryCrownResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/web/generic/country/list"
	)
	return bilibili.Execute[*GetCountryCrownResult](ctx, c.Client, method, url, nil)
}

// SendSMSParam 是发送短信验证码的请求参数。
type SendSMSParam struct {
	Cid       int    `json:"cid"`       // 国际冠字码。使用 GetCountryCrown() 方法获取
	Tel       int    `json:"tel"`       // 手机号码
	Source    string `json:"source"`    // 登录来源。main_web：独立登录页。main_mini：小窗登录
	Token     string `json:"token"`     // 登录 API token。使用 Captcha() 方法获取
	Challenge string `json:"challenge"` // 极验 challenge。使用 Captcha() 方法获取
	Validate  string `json:"validate"`  // 极验 result。极验验证后得到
	Seccode   string `json:"seccode"`   // 极验 result +jordan。极验验证后得到
}

// SendSMS 发送短信验证码
func (c *Client) SendSMS(ctx context.Context, param SendSMSParam) (*SendSMSResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://passport.bilibili.com/x/passport-login/web/sms/send"
	)
	return bilibili.Execute[*SendSMSResult](ctx, c.Client, method, url, param)
}

// SMSLoginParam 是短信验证码登录的请求参数。
type SMSLoginParam struct {
	Cid        int    `json:"cid"`                                        // 国际冠字码。可以从 GetCountryCrown() 获取
	Tel        int    `json:"tel"`                                        // 手机号码
	Code       int    `json:"code"`                                       // 短信验证码。timeout 为 5min
	Source     string `json:"source"`                                     // 登录来源。main_web：独立登录页。main_mini：小窗登录
	CaptchaKey string `json:"captcha_key"`                                // 短信登录 token。从 SendSMS() 请求成功后返回
	GoURL      string `json:"go_url,omitempty" request:"query,omitempty"` // 跳转url。默认为 https://www.bilibili.com
	Keep       bool   `json:"keep,omitempty" request:"query,omitempty"`   // 是否记住登录。true：记住登录。false：不记住登录
}

// LoginWithSMS 使用短信验证码登录
func (c *Client) LoginWithSMS(ctx context.Context, param SMSLoginParam) (*SMSLoginResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://passport.bilibili.com/x/passport-login/web/login/sms"
	)
	return bilibili.Execute[*SMSLoginResult](ctx, c.Client, method, url, param)
}

// Encode a QRCode and return a raw PNG image.
func (result *GetQRCodeResult) Encode() ([]byte, error) {
	return qrcode.Encode(result.URL, qrcode.Medium, 256)
}

// Print the QRCode in the console
func (result *GetQRCodeResult) Print() {
	front, back := qrcodeterminal.ConsoleColors.BrightBlack, qrcodeterminal.ConsoleColors.BrightWhite
	qrcodeterminal.New2(front, back, qrcodeterminal.QRCodeRecoveryLevels.Low).Get(result.URL).Print()
}

// GetQRCode 申请二维码
func (c *Client) GetQRCode(ctx context.Context) (*GetQRCodeResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	)
	return bilibili.Execute[*GetQRCodeResult](ctx, c.Client, method, url, nil)
}

// QRCodeLoginParam 是扫码登录的请求参数。
type QRCodeLoginParam struct {
	QrcodeKey string `json:"qrcode_key"` // 扫码登录秘钥
}

// LoginWithQRCode 使用扫码登录。
//
// 该方法会阻塞直到扫码成功或者已经无法扫码。
func (c *Client) LoginWithQRCode(ctx context.Context, param QRCodeLoginParam) (*QRCodeLoginResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
	)
	for {
		result, err := bilibili.Execute[*QRCodeLoginResult](ctx, c.Client, method, url, param)
		if err != nil {
			return nil, err
		}
		if result.Code != 86090 && result.Code != 86101 {
			// 86090：二维码已扫码未确认
			// 86101：未扫码
			return result, nil
		}
		time.Sleep(3 * time.Second) // 主站 3s 一次请求
	}
}

// GetAccountInformation 获取我的信息
func (c *Client) GetAccountInformation(ctx context.Context) (*GetAccountInformationResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/member/web/account"
	)
	return bilibili.Execute[*GetAccountInformationResult](ctx, c.Client, method, url, nil)
}
