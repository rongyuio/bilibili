<div align="center">

# 哔哩哔哩-API-Go版本

[![](https://img.shields.io/github/v/tag/CuteReimu/bilibili?label=release "最新版本")](https://github.com/CuteReimu/bilibili/tags)
![](https://img.shields.io/github/go-mod/go-version/CuteReimu/bilibili "语言")
[![](https://img.shields.io/github/stars/CuteReimu/bilibili?style=flat&color=yellow)](#star-history "stars")
[![](https://img.shields.io/github/actions/workflow/status/CuteReimu/bilibili/golangci-lint.yml?branch=master)](https://github.com/CuteReimu/bilibili/actions/workflows/golangci-lint.yml "代码分析")
[![](https://img.shields.io/github/contributors/CuteReimu/bilibili)](https://github.com/CuteReimu/bilibili/graphs/contributors "贡献者")
[![](https://img.shields.io/github/license/CuteReimu/bilibili)](https://github.com/CuteReimu/bilibili/blob/master/LICENSE "许可协议")
</div>

本项目是基于Go语言编写的哔哩哔哩API调用。目前常用的接口已经基本完成。

**本项目不会编写单元测试代码**。一则因为各项数据会频繁变动，难以写成固定的结果；二则因为每次单元测试都要大量请求B站API，会对其产生不必要的压力。
如果你发现有**接口bug**或者**有你需要但是本库尚未实现的接口**，可以[提交issue](https://github.com/CuteReimu/bilibili/issues/new/choose)或者[提交pull request](.github/CONTRIBUTING.md)。
如果因为B站修改了接口导致接口突然不可用，不一定能够及时更新，很大程度上需要依赖各位的告知。

> [!IMPORTANT]
> 现在是v2.1+版本，鉴于`golang.org/x`下面的很多库都已经强制要求Go1.23以上了，我们也同步进行了更新。
> 
> 如果想使用v2.0版本（支持Go1.19及以上），请执行`go get -u github.com/CuteReimu/bilibili/v2@v2.0.0`获取旧版本。
> 
> [如果还想使用更早的版本可以点击这里跳转](https://github.com/CuteReimu/bilibili/tree/v1)。

**如果你觉得本项目对你有帮助，点亮右上角的↗ :star: 不迷路**

## 声明

1. 本项目遵守 AGPL 开源协议。
2. 本项目基于 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)
   中描述的接口编写。请尊重该项目作者的努力，遵循该项目的开源要求，禁止一切商业使用。
3. **请勿滥用，本项目仅用于学习和测试！利用本项目提供的接口、文档等造成不良影响及后果与本人无关。**
4. 由于本项目的特殊性，可能随时停止开发或删档
5. 本项目为开源项目，不接受任何形式的催单和索取行为，更不容许存在付费内容

PS：目前，B站调用接口时强制使用 `https` 协议

## 快速开始

### 安装

```bash
go get -u github.com/CuteReimu/bilibili/v2 # 定期执行可以更新最新版本
```

在项目中引用即可使用

```go
import "github.com/CuteReimu/bilibili/v2"

var client = bilibili.New()
```

### 首次登录

> [!TIP]
> 下文为了篇幅更短，示例中把很多显而易见的`err`校验忽略成了`_`，实际使用请自行校验`err`。

#### 方法一：扫码登录

首先获取二维码：

```go
qrCode, _ := client.GetQRCode()
buf, _ := qrCode.Encode()
img, _ := png.Decode(buf) // 或者写入文件 os.WriteFile("qrcode.png", buf, 0644)
// 也可以调用 qrCode.Print() 将二维码打印在控制台
```

扫码并确认成功后，发送登录请求：

```go
result, err := client.LoginWithQRCode(bilibili.LoginWithQRCodeParam{
    QrcodeKey: qrCode.QrcodeKey,
})
if err == nil && result.Code == 0 {
    log.Println("登录成功")
}
```

#### 方法二：账号密码登录

首先获取人机验证参数：

```go
captchaResult, _ := client.Captcha()
```

将`captchaResult`中的`gt`和`challenge`值保存下来，自行使用 [手动验证器](https://kuresaru.github.io/geetest-validator/) 进行人机验证，并获得`validate`和`seccode`。然后使用账号密码进行登录即可：

```go
result, err := client.LoginWithPassword(bilibili.LoginWithPasswordParam{
    Username:  userName,
    Password:  password,
    Token:     captchaResult.Token,
    Challenge: captchaResult.Geetest.Challenge,
    Validate:  validate,
    Seccode:   seccode,
})
if err == nil && result.Status == 0 {
    log.Println("登录成功")
}
```

#### 方法三：使用短信验证码登录（不推荐）

首先用上述方法二相同的方式获取人机验证参数并进行人机验证。然后获取国际地区代码：

```go
countryCrownResult, _ := client.GetCountryCrown()
```

当然，如果你已经确定`cid`的值，这一步可以跳过。中国大陆的`cid`就是`86`。

然后发送短信验证码：*（[这个接口大概率返回86103错误](https://github.com/SocialSisterYi/bilibili-API-collect/issues/756)）*

```go
sendSMSResult, _ := client.SendSMS(bilibili.SendSMSParam{
    Cid:       cid,
    Tel:       tel,
    Source:    "main_web",
    Token:     captchaResult.Token,
    Challenge: captchaResult.Geetest.Challenge,
    Validate:  validate,
    Seccode:   seccode,
})
```

然后就可以使用手机验证码登录了：

```go
result, err := client.LoginWithSMS(bilibili.LoginWithSMSParam{
    Cid:        cid,
    Tel:        tel,
    Code:       123456, // 短信验证码
    Source:     "main_web",
    CaptchaKey: sendSMSResult.CaptchaKey,
})
if err == nil && result.Status == 0 {
    log.Println("登录成功")
}
```

### 储存Cookies

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
> - `GetCookiesString`和`SetCookiesString`使用的字符串是`"cookie1=xxx; expires=xxx; domain=xxx.com; path=/\ncookie2=xxx; expires=xxx; domain=xxx.com; path=/"`，包含过期时间、domain等一些其它信息，以`"\n"`分隔多个cookie
> - `SetRawCookies`使用的字符串是`"cookie1=xxx; cookie2=xxx"`，只包含key=value，以`"; "`分隔多个cookie，这和在浏览器F12里复制的一样
>
> 请注意不要混用。

### 其它接口

你可以很方便的调用其它接口，以下举个例子：

```go
videoInfo, err := client.GetVideoInfo(bilibili.VideoParam{
    Aid: 12345678,
})
```

参数中非必填字段你可以不填（可以通过是否有`omitempty`来判断这个字段是否为非必填字段）。

方法都是按照对应功能的英文翻译命名的，因此你可以方便地使用IDE找到想要的方法，配合注释便能够知道如何使用。

### 对B站返回的错误码进行处理

因为B站的返回内容是这样的格式：

```json
{
   "code": 0,
   "message": "错误信息",
   "data": {}
}
```

而我们这个库的接口只会返回`data`数据和一个`error`，若`code`为`0`则`error`为`nil`，否则我们并不会把`code`和`message`字段直接返回。

在一般情况下，调用者不太需要关心`code`和`message`字段，只需要关心是否有`error`即可。
但如果你实在需要`code`和`message`字段，我们也提供了一个办法：

```go
videoInfo, err := client.GetVideoInfo(bilibili.VideoParam{
    Aid: 12345678,
})
if err != nil {
    var e bilibili.Error
    if errors.As(err, &e) { // B站返回的错误
        log.Printf("错误码: %d, 错误信息: %s", e.Code, e.Message)
    } else { // 其它错误
        log.Printf("%+v", err)
    }
}
```

> [!TIP]
> 通过 `errors.As` 检查业务错误和解码错误，通过 `errors.Is` 检查取消、超时等底层错误。新的统一请求流程不保证所有错误都携带堆栈；不要依赖错误字符串进行判断。

### 可能用到的工具接口

```go
// 解析短连接
typ, id, err := client.UnwrapShortUrl("https://b23.tv/xxxxxx")

// 获取服务器当前时间
now, err := client.Now()

// av号转bv号
bvid := bilibili.AvToBv(111298867365120)

// bv号转av号
aid := bilibili.BvToAv("BV1L9Uoa9EUx")

// 通过ip确定地理位置
zoneLocation, err := client.GetZoneLocation()

// 获取分区当日投稿稿件数
regionDailyCount, err := client.GetRegionDailyCount()
```

### 设置*resty.Client的一些参数

调用`client.Resty()`就可以获取到`*resty.Client`，然后自行操作即可。**但是不要做一些离谱的操作**~~（比如把Cookies删了）~~

```go
client.Resty().SetTimeout(20 * time.Second) // 设置超时时间
client.Resty().SetLogger(logger) // 自定义logger
```

## 自定义接口与重构迁移

当前工作区的模块名是 `bilibili`，Go 版本以 `go.mod` 为准（目前为 1.27）。本文前面的上游安装路径和旧版本说明保留作历史参考，本次重构不调整模块路径或发布版本。

### 调用尚未封装的接口

使用 `Client.Do` 复用客户端的 Cookie、网络配置、签名和解码流程。以下函数只演示调用方式，需由调用方提供 context 和客户端：

```go
func loadAccount(ctx context.Context, client *bilibili.Client, mid string) error {
    var result struct {
        Mid  int64  `json:"mid"`
        Name string `json:"name"`
    }
    err := client.Do(ctx, bilibili.Request{
        Method: http.MethodGet,
        URL:    "https://api.bilibili.com/x/space/wbi/acc/info",
        Query:  url.Values{"mid": {mid}},
        WBI:    true,
    }, &result)
    if err != nil {
        return err
    }
    fmt.Println(result.Name)
    return nil
}
```

示例所需导入为 `bilibili`、`context`、`fmt`、`net/http`、`net/url`。

- `out` 接收响应的 `data`，不要再次包裹 `code/message/data`。传 `nil` 只检查 HTTP 状态及业务错误；传 `*json.RawMessage` 保留原始 `data`。
- `Query` 可与 `Form` 或 `JSON` 并用，但 `Form` 与 `JSON` 互斥。`JSON` 按标准库规则编码，包括字符串值；`Headers` 使用 `http.Header`。
- URL 自带查询参数会参与请求，同名键以 `Request.Query` 为准。WBI 只签查询参数，不签表单；签名请求的每个查询键必须只有一个值。CSRF 由调用方按接口要求放在查询或表单中。
- context 会传到 HTTP 请求和 WBI 密钥刷新；旧的内置 API 方法签名保持不变。此入口不额外启用重试；通过 Resty 自行配置的重试策略仍然有效。
- `Resty()` 保留给特殊请求，但直接调用它不会自动签名、共享 Client 的 Cookie 存储或获得 `DecodeError`。普通 Client 请求支持并发；登录、主动刷新登录态、账号切换、手动修改 Cookie 和配置必须在请求之外串行执行。自定义中间件的并发安全由调用方负责。

### 定位反序列化失败

```go
var de *bilibili.DecodeError
if errors.As(err, &de) {
    log.Printf("接口=%s 类型=%s Go字段=%s JSON路径=%s 预期=%s 实际=%s 偏移=%d 精确=%t",
        de.Endpoint, de.RootType, de.GoField, de.JSONPath,
        de.Expected, de.Actual, de.Offset, de.Exact)
}
```

`JSONPath` 包含数组下标，例如 `$.data.items[3].modules.module_author.mid`；匿名结构通过根类型与 Go 字段链定位。`Offset` 从响应体第一个字节起按 1 计数，零表示不可用。原始错误保留在错误链中，包括可通过 `errors.As` 提取的 `*json.UnmarshalTypeError`。

正常响应只使用标准库解码；失败时才额外扫描 JSON。自定义解码器不重复执行，最多定位到其字段边界，多个不确定边界退回共同父级并标记 `Exact=false`；语法错误只报告可用偏移量。诊断记录首个能确认的不匹配，不保证枚举所有问题。错误文本不包含字段值、完整响应或查询参数；JSON 路径中的 map 键仍来自响应，分享日志前请注意这一点。

先判断业务 `code`，再解码 `data`，避免业务错误被结果类型不匹配掩盖。失败时不向 `out` 写入部分结果，不自动将无效值转成零。`out=nil` 时不会检查 `data` 的字段类型。

### 字段与调用迁移

| 旧用法 | 新用法 |
| --- | --- |
| `client.Wbi`、`FillWbiHandler(...)` | `Client.Do(ctx, Request{WBI: true, ...}, &data)`；签名器由客户端管理 |
| 自定义接口 `SetResult(&response)` | `Do(..., &response.Data)`，并处理返回的业务错误 |
| `VideoStatusNumber.View` 为 `json.Number` | 改为 `NumberOrString`，仍提供 `String()`、`Int64()`、`Float64()` |
| 根据错误字符串或堆栈定位 | `errors.As(err, &decodeError)` 获取结构化位置 |

`NumberOrString` 保留数字、字符串（包括 `"--"`）与 `null`，`Kind()` 返回 `number`、`string` 或 `null`。数值转换会显式返回错误，JSON 再编码保留原始类别；零值表示 `null`。原有直接赋值 `json.Number` 或强制转换为字符串的代码需要迁移到这些方法；需要构造值时可使用 `json.Unmarshal`。

其余字段不批量改型。`json.Number` 支持数字及数字字符串，不支持任意文本。动态数据中的 `Following` 在历史提交和本地修改之间存在 `bool/json.Number` 差异；消息参数 `SendPrivateMessageParam.Content` 也包含多种语义，后续应结合实际响应／请求证据定点处理，本次不推测改型。

### 维护与验证

请求、参数、响应、诊断分别位于同包内的独立文件；内置接口继续按业务分类组织。本地 `test/` 工具已迁移到新入口，但该目录被 Git 忽略，不随库提交分发。

本次仅使用 `go build ./...` 和 `go vet ./...` 验证，不新增或运行测试、不调用实机 API。并发行为未经运行或 race 检查验证，编译通过不能证明并发正确性。后续若进行测试，应先获得明确任务要求。

## 客户端会话状态迁移

### Cookie 所有权

`Client` 现在独立保存 Cookie，读写都会复制 Cookie 及其 `Unparsed` 切片。修改 `GetCookies()` 返回的切片或对象不会改变客户端；需要更新时应在请求结束后调用 `SetCookie` / `SetCookies`。

仍按 Cookie 名称合并，不实现域名、路径或 Secure 匹配。同名响应 Cookie 以最后完成合并的响应为准；HTTP 或业务失败响应也可更新 Cookie。`MaxAge < 0` 删除同名项；正 `MaxAge` 优先于 `Expires`，导入时转换成绝对到期时间并清零 `MaxAge`，读取或重新导入快照不会续期。过期 Cookie 不进入新请求。传 nil 项会被忽略，`SetCookies(nil)` 不表示清空会话；切换账号推荐创建新的 Client。

每次请求只取一次 Cookie 快照，内置 CSRF 和直播签名使用这份快照；自动合并响应期间不持锁等待网络。一次批量操作应复用已配置的客户端，但不要与登录或手动替换会话并行，也不要复制已使用的 Client。

### 接管自定义 Resty

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

### 游客初始化

旧的 `NewAnonymousClient()` 改为接收 context 并返回错误：

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

nil context、取消、网络故障、非 HTTP 200 或没有有效 Cookie 都会返回错误。初始化复用默认配置；短链接仍要求 HTTP 302，刷新口令页面保持 HTML 解析，WBI 保留非零业务码但存在有效密钥的特殊处理。

参数编码另外修复了非空、非结构体指针导致的 panic，并保留标签值中的等号；nil 参数与 nil 指针继续视为未传参。现有内置接口的方法签名、参数位置及请求体编码规则不变。

## Star History

<a href="https://star-history.com/#CuteReimu/bilibili&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=CuteReimu/bilibili&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=CuteReimu/bilibili&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=CuteReimu/bilibili&type=Date" />
 </picture>
</a>

## 如何为仓库做贡献？

不知道在哪些方面可以做贡献？[点击这里看看吧！](https://github.com/CuteReimu/bilibili/contribute)

命名规范和编码风格请参考[CONTRIBUTING.md](.github/CONTRIBUTING.md)
