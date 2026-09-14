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

以下网络调用都需要传入 `ctx`：HTTP 处理函数使用 `r.Context()`，独立任务可使用 `context.WithTimeout` 或 `signal.NotifyContext`。完整示例见下文“Context 迁移与任务取消”。

### 首次登录

> [!TIP]
> 下文为了篇幅更短，示例中把很多显而易见的`err`校验忽略成了`_`，实际使用请自行校验`err`。

#### 方法一：扫码登录

首先获取二维码：

```go
qrCode, err := client.GetQRCode(ctx)
if err != nil {
    log.Printf("获取二维码失败: %v", err)
    return
}
buf, _ := qrCode.Encode()
img, _ := png.Decode(buf) // 或者写入文件 os.WriteFile("qrcode.png", buf, 0644)
// 也可以调用 qrCode.Print() 将二维码打印在控制台
```

扫码并确认成功后，发送登录请求：

```go
result, err := client.LoginWithQRCode(ctx, bilibili.LoginWithQRCodeParam{
    QrcodeKey: qrCode.QrcodeKey,
})
if err == nil && result.Code == 0 {
    log.Println("登录成功")
}
```

#### 方法二：账号密码登录

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
if err == nil && result.Status == 0 {
    log.Println("登录成功")
}
```

#### 方法三：使用短信验证码登录（不推荐）

首先用上述方法二相同的方式获取人机验证参数并进行人机验证。然后获取国际地区代码：

```go
countryCrownResult, err := client.GetCountryCrown(ctx)
if err != nil {
    log.Printf("获取地区代码失败: %v", err)
    return
}
```

当然，如果你已经确定`cid`的值，这一步可以跳过。中国大陆的`cid`就是`86`。

然后发送短信验证码：*（[这个接口大概率返回86103错误](https://github.com/SocialSisterYi/bilibili-API-collect/issues/756)）*

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
```

发送短信后先检查 `err`，失败（包括取消或超时）时结束任务，不访问 `sendSMSResult`。成功后就可以使用手机验证码登录：

```go
result, err := client.LoginWithSMS(ctx, bilibili.LoginWithSMSParam{
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
videoInfo, err := client.GetVideoInfo(ctx, bilibili.VideoParam{
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
videoInfo, err := client.GetVideoInfo(ctx, bilibili.VideoParam{
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
typ, id, err := client.UnwrapShortUrl(ctx, "https://b23.tv/xxxxxx")

// 获取服务器当前时间
now, err := client.Now(ctx)

// av号转bv号
bvid := bilibili.AvToBv(111298867365120)

// bv号转av号
aid := bilibili.BvToAv("BV1L9Uoa9EUx")

// 通过ip确定地理位置
zoneLocation, err := client.GetZoneLocation(ctx)

// 获取分区当日投稿稿件数
regionDailyCount, err := client.GetRegionDailyCount(ctx)
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
- context 会传到 HTTP 请求和 WBI 密钥刷新；所有内置网络 API 同样以 context 为首参。此入口不额外启用重试；通过 Resty 自行配置的重试策略仍然有效。
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

其余字段不批量改型。`json.Number` 支持数字及数字字符串，不支持任意文本。动态数据中的两处 `Following` 已根据实际响应改为 `json.Number`，迁移方式见下文。消息参数 `SendPrivateMessageParam.Content` 也包含多种语义，后续应结合实际响应／请求证据定点处理，本次不推测改型。

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

参数编码另外修复了非空、非结构体指针导致的 panic，并保留标签值中的等号；nil 参数与 nil 指针继续视为未传参。内置接口的请求参数位置及请求体编码规则不变。

## Context 迁移与任务取消

这是一次破坏性签名变更：所有可能联网的 `Client` 方法统一增加首参 `ctx context.Context`，例如 `client.GetVideoInfo(ctx, param)`、`client.GetMyUserSpaceDetail(ctx)`。本文调用示例中的 `ctx` 均由调用方提供。`Do` 和 `NewAnonymousClient` 已有的 context 签名不变；Cookie 读写、配置和纯计算方法不变。

独立使用 WBI 时，改为 `wbi.GetKeys(ctx)`、`wbi.GetMixinKey(ctx)`、`wbi.SignQuery(ctx, query, ts)`、`wbi.SignMap(ctx, payload, ts)`，因为签名可能触发密钥刷新。没有新增 `XxxContext` 或无 context 的兼容包装；调用方接口声明、方法表达式和回调类型也需同步修改。

批量任务应从入口创建可取消 context，并贯通辅助函数：

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()
ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
defer cancel()

for _, param := range params { // params 为调用方的视频参数列表
    if ctx.Err() != nil {
        return
    }
    info, err := client.GetVideoInfo(ctx, param)
    if err != nil {
        if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
            return
        }
        log.Printf("获取视频失败: %v", err)
        continue
    }
    log.Println(info.Title)
}
```

示例需要 `context`、`os`、`os/signal`、`time`、`errors`、`log`。HTTP 服务中直接使用 `r.Context()`；需要缩短期限时由调用方派生 context。库在准备请求前拒绝 nil 或已结束的 context，并将取消、超时保留在错误链中，不创建后台替代请求或增加重试；已有 Resty 请求超时仍可更早结束请求。

本地批量工具已接入中断信号，取消后停止翻页、后续账号操作及等待；配置加载和独立 HTTP 工具不在本次迁移范围。取消不保证服务端撤销已经收到的写操作，不应据此自动重试。`test/` 被忽略，这些适配不随 Git 提交分发。此次仅通过编译和静态检查，未验证实机取消行为。

## 参数编码与错误定位

内置接口先完成参数编码，再向请求应用 query、请求体和头部；编码失败不写入部分参数，也不会继续签名或发送请求。公开网络方法签名及现有业务参数位置不变，`Client.Do` 的 `Query`、`Form`、`JSON` 仍按原有方式使用。以下标签规则用于内置接口的参数结构体，不用于 `Client.Do` 的 JSON 对象。

| 规则 | 行为 |
| --- | --- |
| 未指定位置／`request:"query"` | 放入 URL 查询参数，POST 方法也不自动改为表单 |
| `request:"json"` | 作为 JSON 请求体字段，保留标准库编码语义 |
| `request:"form-data"` | 转换为 multipart 文本字段，由 Resty 构造请求体和 boundary |
| `request:"-"` | 跳过字段；未导出字段同样跳过 |
| `request:"field=name"` | 优先使用该名称，其次取 JSON 标签名，最后使用原有 snake_case 规则 |
| `request:"omitempty,default=1"` | 零值优先省略；未指定省略时才使用字符串默认值 |

保留历史零值语义：非 nil 指针即使指向零值也不算零值；空但非 nil 的切片不算零值。query 切片仍按元素转换后用逗号拼接，nil 切片为 `""`，其默认值继续不参与拼接；multipart 文本切片使用相同规则。JSON 中的 `default=1` 仍是字符串 `"1"`，不会按 Go 字段类型转换成数字。

只有 `request` 标签控制请求省略；`json:"-"` 不等于 `request:"-"`，`json:",omitempty"` 也不新增请求省略行为。nil 参数或 nil 指针继续视为未传参；非 nil 参数仅接受结构体或一层结构体指针。没有实际参与编码的字段时，请求保持原状。

query 可以与一种请求体并存，不再因字段排列覆盖 Content-Type。实际参与编码的字段若同时声明多个位置，或混用 JSON 与 multipart，返回错误；已被省略的字段不参与冲突判断。multipart 不再通过普通 map 加请求头模拟，因此依赖该内部表示的代码需要调整。当前业务参数没有使用 JSON/multipart 标签，图片上传的专用实现保持原样。

```go
var pe *bilibili.ParamError
if errors.As(err, &pe) {
    log.Printf("参数类型=%s Go字段=%s 参数名=%s 位置=%s",
        pe.RootType, pe.GoField, pe.Parameter, pe.Location)
}
```

`GoField` 可为 `Ids[2]`；整体参数类型错误的字段、参数名和位置为空。JSON 编码错误定位到顶层参数字段，底层 `*json.MarshalerError` 等错误可继续解包；不会重复执行自定义编码器来探测内部路径。内置请求的错误外层还包含 HTTP 方法及去除查询参数的接口地址。

参数转换失败现在明确返回错误，不再静默变为空字符串。正常错误文本不包含参数值或原始错误文本；`ParamError.Err` 保留底层错误，可能含有原始值，不要直接写入日志。`Client.Do` 的手工 JSON 编码错误不转换为 `ParamError`。

本阶段仅通过 `go build ./...` 和 `go vet ./...`，已有测试源码按新的 JSON/multipart 内部表示及错误类型适配，未新增或执行测试、未访问真实 API，multipart 的线上行为尚未验证。

## 动态响应模型迁移

`DynamicItem`、`DynamicInfo` 和主要模块定义移至同包的 `dynamic_model.go`，接口调用与参数仍在 `dynamic.go`。导入路径、网络方法签名、返回根类型以及 `item.Modules.ModuleAuthor.Name` 等字段访问路径不变。

| 字段 | 命名类型 |
| --- | --- |
| `DynamicItem.Basic` / `Modules` | `DynamicItemBasic` / `DynamicItemModules` |
| 外层作者 / 头像 | `DynamicModuleAuthor` / `DynamicAuthorAvatar` |
| 外层内容 / 描述 / 主体 | `DynamicModuleDynamic` / `*DynamicDescription` / `*DynamicMajor` |
| 更多操作 / 统计 | `DynamicModuleMore` / `DynamicModuleStat` |
| `DynamicItem.Orig` | `DynamicOriginalItem`，保留值类型，不是递归动态 |
| 原动态 Basic / Modules | `DynamicOriginalBasic` / `DynamicOriginalModules` |
| 原动态作者 / 头像 | `DynamicOriginalModuleAuthor` / `DynamicOriginalAuthorAvatar` |
| 原动态内容 / 描述 / 主体 | `DynamicOriginalModuleDynamic` / `*DynamicOriginalDescription` / `DynamicOriginalMajor` |
| 两套主体中的视频 / 图文 | `DynamicArchive` / `DynamicDraw` |

这是嵌套字段 Go 类型身份的变更，使用命名类型而不是匿名结构别名。普通字段读取无需迁移；手写匿名结构赋值、嵌套复合字面量、函数或接口声明以及依赖类型名称的反射代码需要检查。可改用新类型构造模块，例如：

```go
item := bilibili.DynamicItem{
    Modules: bilibili.DynamicItemModules{
        ModuleAuthor: bilibili.DynamicModuleAuthor{Name: "作者"},
    },
}
item.Modules.ModuleDynamic.Major = &bilibili.DynamicMajor{
    Archive: bilibili.DynamicArchive{Bvid: "BV1L9Uoa9EUx"},
}
item.Orig.Modules.ModuleDynamic.Major = bilibili.DynamicOriginalMajor{
    Archive: bilibili.DynamicArchive{Bvid: "BV1L9Uoa9EUx"},
}
```

两套模型的差异完整保留：外层 `Major` 是指针，原动态 `Major` 是值；外层 `Basic.LikeIcon.Id` 为 `json.Number`，原动态对应字段为 `int`。原动态头像、作者和富文本也有不同字段，不能直接复用外层模块。更深层的小型匿名结构暂不提取。

模型提取本身没有修改 JSON 标签、字段顺序、叶子类型或指针／切片结构，没有新增自定义反序列化；`DecodeError` 仍按既有规则遍历模型并报告 JSON 路径和 Go 字段。随后单独修复了两处 `Following` 类型，见下文。

静态展开命名类型后，已分别确认提交版本和本地版本与各自重构前结构一致，并通过 `go build ./...`、`go vet ./...`。未新增或运行测试、未访问真实 API；这些检查不能替代真实响应兼容性验证。动态模型阶段不新增话题接口，也不修改 `watchVideo` 的独立话题响应结构。

### Following 数值状态修复

`DynamicModuleAuthor.Following` 和 `DynamicOriginalModuleAuthor.Following` 从 `bool` 改为 `json.Number`。一份成功响应包含 12 条动态，外层作者该字段均为数字 `2`，原动态作者均为数字 `1`，原来的布尔类型会导致解码失败。这份样本尚不能确定全部状态含义或所有可能返回形式，因此暂不定义状态常量，也不将非零值解释为“已关注”。

字段不能再直接用于布尔条件；显示原始状态使用 `String()`，需要数值时调用 `Int64()` 并处理错误。不要忽略转换失败，也不要把失败转换成零：

```go
rawStatus := item.Modules.ModuleAuthor.Following.String()
status, err := item.Modules.ModuleAuthor.Following.Int64()
if err != nil {
    log.Print("关注状态无法转换为整数")
    return
}
// 按调用方已确认的状态定义处理 status。
_ = rawStatus
_ = status
```

原动态对应字段为 `item.Orig.Modules.ModuleAuthor.Following`，读取方式相同。头像尺寸仍保留 `float64`，样本包含小数；原动态 `LikeIcon` 在该样本中全部为 `null`，不能据此判断其内部 `Id` 类型，因此未改动。调试记录不随 Git 提交分发。

## 话题动态列表

`GetTopicFeed(ctx, param)` 封装 `/x/polymer/web-dynamic/v1/feed/topic`，返回一页响应的 `data`，复用统一 Cookie、context、参数编码和错误处理。参数全部位于 query，不额外启用 WBI 签名或 CSRF，也不自动重试或翻页。

```go
result, err := client.GetTopicFeed(ctx, bilibili.GetTopicFeedParam{
    TopicId:     topicId,
    SortBy:      3,
    PageSize:    20,
    Features:    "itemOpusStyle,listOnlyfans,opusBigCover,onlyfansVote,decorationCard",
    WebLocation: "0.0",
})
if err != nil {
    log.Printf("获取话题动态失败: %v", err)
    return
}
for _, item := range result.TopicCardList.Items {
    if item.DynamicCardItem.Type == "DYNAMIC_TYPE_AV" {
        log.Println(item.DynamicCardItem.Modules.ModuleDynamic.Major.Archive.Bvid)
    }
}
```

示例中的 `ctx`、`client`、`topicId` 由调用方提供。库不填入工具专用默认值；`SortBy`、`PageSize`、`Offset`、`Features`、`WebLocation` 为零值时不发送。需要下一页时，根据 `result.TopicCardList.HasMore`，将 `result.TopicCardList.Offset` 传入下一次调用；调用方应检查取消、空 offset 或 offset 未变化，避免无进展循环。示例及本地 `watchVideo` 仍只处理一页。

`topic_model.go` 保留原工具完整的已声明字段，并将卡片、作者、头像、内容和统计提取为命名类型；这不代表已经覆盖服务端所有字段。仅更多操作模块复用 `DynamicModuleMore`。话题的 `Following`、内容描述等仍为 `any`，计数保留 `int`，头像尺寸保留 `float64`，没有按空间动态模型推断改型。

本地 `watchVideo` 已从 `Client.Do` 和 `topicResp.Data` 迁移至此方法及 `GetTopicFeedResult`，移除了被替代的 `topicResp`、`topicData`；抽奖和视频心跳调用不变。`test/` 被 Git 忽略，工具迁移不随提交分发。

验证包含模型静态展开比对、请求参数检查以及 `go build ./...`、`go vet ./...`，未运行测试或访问真实 API。`test/dynamicItem.txt` 是 `GetUserSpaceDynamic` 返回的空间动态记录，对应 `DynamicItem`，不作为话题模型的验证样本；话题字段的全部返回形式尚未确认。

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
