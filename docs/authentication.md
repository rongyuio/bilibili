# 认证与会话

[返回 README](../README.md)

本文代码片段中的 `ctx`、`client` 及业务参数由调用方提供；客户端导入路径为 `github.com/rongyuio/bilibili`。

## 游客初始化

`New()` 只创建客户端；需要从首页获取游客 Cookie 时使用 `NewAnonymousClient(ctx)`，并检查初始化错误：

```go
ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
defer cancel()
client, err := bilibili.NewAnonymousClient(ctx)
if err != nil {
    log.Printf("游客初始化失败: %v", err)
    return
}
// 后续使用 client；构造失败时不要继续调用其方法。
```

nil context、取消、网络故障、非 HTTP 200 或没有有效 Cookie 都会返回错误。初始化保留首页所需的浏览器请求头；构造失败后不要继续使用返回的客户端。

## 首次登录

登录前可使用游客或已有 Cookie 客户端；登录、主动刷新会话和手动修改 Cookie 必须在普通请求之外串行进行。涉及扫码或人工验证时，应为 context 设置足够的有效期。

### 方法一：扫码登录

首先获取二维码：

```go
qrCode, err := client.GetQRCode(ctx)
if err != nil {
    log.Printf("获取二维码失败: %v", err)
    return
}
qrCode.Print() // 在控制台显示；需要 PNG 字节时调用 Encode() 并处理其返回错误。
```

显示二维码后调用登录方法，它会轮询等待扫码确认或终止状态；先检查请求错误，再检查返回的扫码状态：

```go
result, err := client.LoginWithQRCode(ctx, bilibili.LoginWithQRCodeParam{
    QrcodeKey: qrCode.QrcodeKey,
})
if err != nil {
    log.Printf("扫码登录失败: %v", err)
    return
}
if result.Code == 0 {
    log.Println("登录成功")
} else {
    log.Printf("扫码登录未完成，状态码: %d", result.Code)
}
```

### 方法二：账号密码登录

首先获取人机验证参数：

```go
captchaResult, err := client.Captcha(ctx)
if err != nil {
    log.Printf("获取验证参数失败: %v", err)
    return
}
```

将`captchaResult`中的`gt`和`challenge`值保存下来，自行使用 [手动验证器](https://kuresaru.github.io/geetest-validator/) 进行人机验证，并获得`validate`和`seccode`。然后使用账号密码进行登录即可：

```go
result, err := client.LoginWithPassword(ctx, bilibili.LoginWithPasswordParam{
    Username:  userName,
    Password:  password,
    Token:     captchaResult.Token,
    Challenge: captchaResult.Geetest.Challenge,
    Validate:  validate,
    Seccode:   seccode,
})
if err != nil {
    log.Printf("登录请求失败: %v", err)
    return
}
if result.Status == 0 {
    log.Println("登录成功")
}
```

### 方法三：使用短信验证码登录（不推荐）

首先用上述方法二相同的方式获取人机验证参数并进行人机验证。然后获取国际地区代码：

```go
countryCrownResult, err := client.GetCountryCrown(ctx)
if err != nil {
    log.Printf("获取地区代码失败: %v", err)
    return
}
log.Printf("地区代码: %+v", countryCrownResult)
```

当然，如果你已经确定`cid`的值，这一步可以跳过。中国大陆的`cid`就是`86`。

然后发送短信验证码：*（上游曾记录 [86103 错误](https://github.com/SocialSisterYi/bilibili-API-collect/issues/756)，当前可用性需以实际响应为准）*

```go
sendSMSResult, err := client.SendSMS(ctx, bilibili.SendSMSParam{
    Cid:       cid,
    Tel:       tel,
    Source:    "main_web",
    Token:     captchaResult.Token,
    Challenge: captchaResult.Geetest.Challenge,
    Validate:  validate,
    Seccode:   seccode,
})
if err != nil {
    log.Printf("发送短信失败: %v", err)
    return
}
```

发送短信成功后，使用验证码登录：

```go
result, err := client.LoginWithSMS(ctx, bilibili.LoginWithSMSParam{
    Cid:        cid,
    Tel:        tel,
    Code:       123456, // 短信验证码
    Source:     "main_web",
    CaptchaKey: sendSMSResult.CaptchaKey,
})
if err != nil {
    log.Printf("登录请求失败: %v", err)
    return
}
if result.Status == 0 {
    log.Println("登录成功")
}
```

## 保存与恢复 Cookie

使用上述任意方式登录成功后，Cookies值就已经设置好了。你可以保存Cookies值方便下次启动程序时不需要重新登录。

```go
// 获取cookiesString，自行存储，方便下次启动程序时不需要重新登录
cookiesString := client.GetCookiesString()

// 下次启动时，把存储的cookiesString设置进来，就不需要登录操作了
client.SetCookiesString(cookiesString)

// 如果你是从浏览器request的header中直接复制出来的cookies，则改为调用SetRawCookies
client.SetRawCookies("cookie1=xxx; cookie2=xxx")
```

> [!NOTE]
>
> - `GetCookiesString`和`SetCookiesString`使用的字符串是`"cookie1=xxx; expires=xxx; domain=xxx.com; path=/\ncookie2=xxx; expires=xxx; domain=xxx.com; path=/"`，包含过期时间、domain等一些其它信息，以`"\n"`分隔多个cookie
> - `SetRawCookies`使用的字符串是`"cookie1=xxx; cookie2=xxx"`，只包含key=value，以`"; "`分隔多个cookie，这和在浏览器F12里复制的一样
>
> 请注意不要混用。

`Client` 独立保存 Cookie，读写都会复制 Cookie 及其 `Unparsed` 切片。修改 `GetCookies()` 返回的切片或对象不会改变客户端；需要更新时应在请求结束后调用 `SetCookie` / `SetCookies`。

仍按 Cookie 名称合并，不实现域名、路径或 Secure 匹配。同名响应 Cookie 以最后完成合并的响应为准；HTTP 或业务失败响应也可更新 Cookie。`MaxAge < 0` 删除同名项；正 `MaxAge` 优先于 `Expires`，导入时转换成绝对到期时间并清零 `MaxAge`，读取或重新导入快照不会续期。过期 Cookie 不进入新请求。传 nil 项会被忽略，`SetCookies(nil)` 不表示清空会话；切换账号推荐创建新的 Client。

每次请求只取一次 Cookie 快照，内置 CSRF 和直播签名使用这份快照；自动合并响应期间不持锁等待网络。一次批量操作应复用已配置的客户端，但不要与登录或手动替换会话并行，也不要复制已使用的 Client。

## Resty 配置与接管

配置应在请求开始前完成，例如 `client.Resty().SetTimeout(20*time.Second)` 或 `SetLogger(logger)`。不要复制已使用的 Client。普通请求支持并发，登录、会话切换和配置变更需串行；自定义中间件的并发安全由调用方负责。

直接使用 `Resty()` 发送请求不会自动签名、共享 Client Cookie 存储或获得统一错误处理。

`NewWithClient` 接管传入 Resty 的独占使用权，自动复制其显式 `Cookies` 并清空原切片，同时关闭 HTTP CookieJar，避免两套会话来源。传 nil 与 `New()` 等效。不要将同一 Resty 再交给另一个 Client，也不要重新启用 Jar、设置 `Resty().Cookies` 或用默认 Cookie 请求头管理会话。

CookieJar 没有通用的全量导出方法。对于已经登录过的外部 Resty，必须在构造前从已知 URL 提取需要的 Cookie，再显式导入，例如：

```go
// existingResty 是调用方已配置的 *resty.Client；此处不发送请求。
target := &url.URL{Scheme: "https", Host: "api.bilibili.com", Path: "/"}
var imported []*http.Cookie
if jar := existingResty.GetClient().Jar; jar != nil {
    imported = jar.Cookies(target)
}
client := bilibili.NewWithClient(existingResty)
client.SetCookies(imported)
```

这只能取出该 URL 对应的 Cookie，不保留 Jar 的完整作用域和过期元数据。不同 URL 导出的同名项仍按名称合并。底层直接请求需要调用方自行提供 Cookie，其响应也不会自动进入 Client 存储。
