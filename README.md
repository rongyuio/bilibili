# 哔哩哔哩 API Go 客户端

本仓库是基于 CuteReimu/bilibili 继续维护的 fork，封装 Bilibili API，并提供 Cookie 管理、WBI 签名、context 取消和结构化错误定位。当前模块名为 `bilibili`，要求 Go 1.27，具体依赖以 [go.mod](go.mod) 为准。

本文描述当前仓库代码，不将上游版本或构建状态作为本 fork 的发布信息。上游安装命令与来源链接见[上游历史参考](#上游历史参考)。接口可能随服务端变化，需要结合对应接口的实际响应维护。

- [快速开始](#快速开始)
- [常用接口](#常用接口)
- [自定义请求](#自定义请求)
- [错误处理](#错误处理)
- [迁移说明](#迁移说明)
- [开发与贡献](#开发与贡献)
- [声明](#声明)
- [上游历史参考](#上游历史参考)

## 快速开始

### 接入当前 fork

在本仓库内部使用 `import "bilibili"`。外部项目可在自己的 `go.mod` 中添加本地依赖；将 replace 路径替换为实际仓库路径：

```go.mod
require bilibili v0.0.0

replace bilibili => D:/Project/bilibili
```

`v0.0.0` 在这里配合本地 replace 使用，不表示已有对应发布版本。外部项目同样需要满足 Go 1.27 要求；相对路径以该项目的 go.mod 所在目录为基准。

### 创建客户端并调用接口

下面是一个完整程序，按需在自己的项目运行：

```go
package main

import (
    "context"
    "log"
    "time"

    "bilibili"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    client := bilibili.New()
    info, err := client.GetVideoInfo(ctx, bilibili.VideoParam{Bvid: "BV1L9Uoa9EUx"})
    if err != nil {
        log.Printf("获取视频失败: %v", err)
        return
    }
    log.Println(info.Title)
}
```

所有可能联网的方法都以 `ctx context.Context` 为首参。HTTP 处理函数直接传 `r.Context()`，批量任务可从 `signal.NotifyContext(context.Background(), os.Interrupt)` 派生 context，并调用返回的停止函数释放资源。需要缩短期限时由调用方使用 `context.WithTimeout`。

库在准备请求前拒绝 nil 或已结束的 context，将其传递到 HTTP 请求及 WBI 密钥刷新，不创建后台替代请求或额外重试；已有 Resty 超时可能更早结束请求。任务应在取消后停止后续操作和等待，取消不保证服务端撤销已收到的写操作。

以下片段置于调用方函数中，复用已创建的 `ctx` 和 `client`，其它参数由调用方提供；用到的标准库符号需按示例导入。

### 游客初始化

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

### 首次登录

登录前可使用游客或已有 Cookie 客户端；登录、主动刷新会话和手动修改 Cookie 必须在普通请求之外串行进行。涉及扫码或人工验证时，应为 context 设置足够的有效期。

#### 方法一：扫码登录

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
if err != nil {
    log.Printf("登录请求失败: %v", err)
    return
}
if result.Status == 0 {
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
log.Printf("地区代码: %+v", countryCrownResult)
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

### 保存与恢复 Cookie

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

### Resty 配置与接管

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

## 常用接口

### 视频与空间动态

视频详情可使用快速开始中的 `GetVideoInfo`。空间动态按页获取：

```go
page, err := client.GetUserSpaceDynamic(ctx, bilibili.GetUserSpaceDynamicParam{
    HostMid:        mid,
    TimezoneOffset: -480,
    Features:       "itemOpusStyle",
})
if err != nil {
    log.Printf("获取空间动态失败: %v", err)
    return
}
for _, item := range page.Items {
    log.Println(item.IdStr.String(), item.Modules.ModuleAuthor.Name)
}
```

`mid` 是调用方提供的 UID 字符串。需要继续读取时，根据 `page.HasMore` 将 `page.Offset` 传入下一次调用。调用方应检查取消及分页进度，避免空 offset 或重复 offset 导致循环。

### 话题动态列表

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

示例中的 `ctx`、`client`、`topicId` 由调用方提供。库不填入工具专用默认值；`SortBy`、`PageSize`、`Offset`、`Features`、`WebLocation` 为零值时不发送。需要下一页时，根据 `result.TopicCardList.HasMore`，将 `result.TopicCardList.Offset` 传入下一次调用；调用方应检查取消、空 offset 或 offset 未变化，避免无进展循环。示例及

`test/dynamicItem.txt` 是空间动态 `GetUserSpaceDynamic` 的记录，不作为话题字段类型的证据。话题字段的全部返回形式尚未确认。

### 直播勋章与点赞

`GetLiveMedalWall` 查询指定用户的勋章墙，`GetLiveFansMedalPanel` 查询当前账号的一页勋章面板，`GetLiveActivatedMedalInfo` 查询已激活勋章及任务信息。返回值直接对应 `data`，分页由调用方控制。

```go
panel, err := client.GetLiveFansMedalPanel(ctx, bilibili.GetLiveFansMedalPanelParam{
    Page: 1, PageSize: 10,
})
if err != nil {
    log.Printf("获取勋章面板失败: %v", err)
    return
}
log.Printf("本页普通勋章数: %d，总页数: %d", len(panel.List), panel.PageInfo.TotalPage)
```

查询任务时显式指定平台和页面位置，例如 `GetLiveActivatedMedalInfoParam{Platform: "pc", RoomId: roomID, TargetId: anchorID, WebLocation: "0.0"}`。`ReportLiveLike(ctx, ReportLiveLikeParam{ClickTime: count, RoomId: roomID, AnchorId: anchorID, Uid: myUID})` 上报点赞并返回 `error`，调用方必须处理错误；数量由调用方指定，库不调度批量任务或自动重试。

任务查询的 CSRF 位于 query，点赞的 CSRF 与业务参数位于 URL 编码表单；两者均由库从本次请求的 Cookie 快照填入，无须手动传递。三个查询结果及嵌套命名类型见 [live_medal_model.go](live_medal_model.go)。模型沿用本地工具已声明字段，勋章墙与面板保留独立结构；尚未通过真实 API 验证全部字段类型。

本地工具迁移时，将原来的 `resp.Data.List` / `resp.Data.TaskInfo` 改为 `result.List` / `result.TaskInfo`，删除对应手写请求和重复响应结构。此次仅新增库 API，不改变既有公开方法或字段类型。移动端直播心跳不在本次封装范围内。

### 工具方法

以下网络方法均需传入 context，并先处理错误再使用结果：

| 方法 | 用途 |
| --- | --- |
| `client.UnwrapShortUrl(ctx, shortURL)` | 解析短链接，返回目标类型和标识 |
| `client.Now(ctx)` | 获取服务器时间 |
| `client.GetZoneLocation(ctx)` | 查询 IP 所属地理位置 |
| `client.GetRegionDailyCount(ctx)` | 获取分区当日投稿数 |
| `bilibili.Av2Bv(aid)` / `bilibili.Bv2Av(bvid)` | 纯计算转换，不需要 context |

其它接口按业务位于 `video.go`、`user.go`、`live.go` 等文件，可通过方法和参数注释查阅。内置请求是否省略参数由 `request` 标签决定，不能只根据 JSON 标签判断。

## 自定义请求

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

### 内置接口的参数标签

内置接口先完成参数编码，再向请求应用 query、请求体和头部；编码失败不写入部分参数，也不会继续签名或发送请求。以下标签规则用于内置接口的参数结构体，不用于 `Client.Do` 的 JSON 对象。

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

query 可以与一种请求体并存，Content-Type 由请求体类型决定。实际参与编码的字段若同时声明多个位置，或混用 JSON 与 multipart，返回错误；已被省略的字段不参与冲突判断。multipart 由 Resty 构造请求体和 boundary，不使用普通 map 加请求头模拟。当前业务参数没有使用 JSON/multipart 标签，图片上传的专用实现保持原样。

## 错误处理

### 分类与判断

标准业务响应先判断 `code`，非零时返回业务错误；成功后再解码 `data`，失败时不写入部分结果。

HTTP 状态不符合接口要求时，库返回可通过 `errors.As` 提取的 `*HTTPError`，包含 `Method`、`Endpoint` 和 `StatusCode`。库生成的 Endpoint 去掉查询参数、用户信息和片段；该错误不保存请求头、Cookie 或响应体。

```go
if err != nil {
    var he *bilibili.HTTPError
    var be bilibili.Error
    switch {
    case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
        return
    case errors.As(err, &he):
        log.Printf("接口=%s %s HTTP状态=%d", he.Method, he.Endpoint, he.StatusCode)
    case errors.As(err, &be):
        log.Printf("业务错误码=%d", be.Code)
    default:
        // 参数与解码失败仍可继续用 errors.As 提取 ParamError、DecodeError。
        log.Print("请求失败")
    }
}
```

示例需要导入 `bilibili`、`context`、`errors`、`log`。HTTP 错误按各接口原有规则判定：普通请求、游客初始化和 WBI 要求 200；短链接要求 302；Cookie 刷新页面接受 2xx。HTTP 状态错误不自动触发重试。网络故障、取消和超时保留原始错误链，不转换成 HTTPError；HTML 解析失败、缺少有效 Cookie、无效密钥等也保持独立错误。

`Message` 仍保留服务端消息，接口地址的清理不代表任意底层错误或服务端消息都经过脱敏。结构化记录优先选择接口、状态码和业务码；不要额外输出凭证、请求头或完整响应。

### 定位参数错误

```go
var pe *bilibili.ParamError
if errors.As(err, &pe) {
    log.Printf("参数类型=%s Go字段=%s 参数名=%s 位置=%s",
        pe.RootType, pe.GoField, pe.Parameter, pe.Location)
}
```

`GoField` 可为 `Ids[2]`；整体参数类型错误的字段、参数名和位置为空。JSON 编码错误定位到顶层参数字段，底层 `*json.MarshalerError` 等错误可继续解包；不会重复执行自定义编码器来探测内部路径。内置请求的错误外层还包含 HTTP 方法及去除查询参数的接口地址。

参数转换失败现在明确返回错误，不再静默变为空字符串。正常错误文本不包含参数值或原始错误文本；`ParamError.Err` 保留底层错误，可能含有原始值，不要直接写入日志。`Client.Do` 的手工 JSON 编码错误不转换为 `ParamError`。

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

## 迁移说明

### 入口、签名与字段

| 旧用法 | 新用法 |
| --- | --- |
| `client.Wbi`、`FillWbiHandler(...)` | `Client.Do(ctx, Request{WBI: true, ...}, &data)`；签名器由客户端管理 |
| 自定义接口 `SetResult(&response)` | `Do(..., &response.Data)`，并处理返回的业务错误 |
| `VideoStatusNumber.View` 为 `json.Number` | 改为 `NumberOrString`，仍提供 `String()`、`Int64()`、`Float64()` |
| 根据错误字符串或堆栈定位 | `errors.As(err, &decodeError)` 获取结构化位置 |

`NumberOrString` 保留数字、字符串（包括 `"--"`）与 `null`，`Kind()` 返回 `number`、`string` 或 `null`。数值转换会显式返回错误，JSON 再编码保留原始类别；零值表示 `null`。原有直接赋值 `json.Number` 或强制转换为字符串的代码需要迁移到这些方法；需要构造值时可使用 `json.Unmarshal`。

其余字段不批量改型。`json.Number` 支持数字及数字字符串，不支持任意文本。动态数据中的两处 `Following` 已根据实际响应改为 `json.Number`，迁移方式见下文。消息参数 `SendPrivateMessageParam.Content` 也包含多种语义，后续应结合实际响应／请求证据定点处理，暂不推测改型。

### Context 签名

相对旧版，这是破坏性签名变更：所有可能联网的 `Client` 方法统一增加首参 `ctx context.Context`，例如 `client.GetVideoInfo(ctx, param)`、`client.GetMyUserSpaceDetail(ctx)`。本文调用示例中的 `ctx` 均由调用方提供。`Do` 和 `NewAnonymousClient` 已有的 context 签名不变；Cookie 读写、配置和纯计算方法不变。

独立使用 WBI 时，改为 `wbi.GetKeys(ctx)`、`wbi.GetMixinKey(ctx)`、`wbi.SignQuery(ctx, query, ts)`、`wbi.SignMap(ctx, payload, ts)`，因为签名可能触发密钥刷新。没有新增 `XxxContext` 或无 context 的兼容包装；调用方接口声明、方法表达式和回调类型也需同步修改。

### 会话与参数编码

`NewAnonymousClient()` 迁移为 `NewAnonymousClient(ctx) (*Client, error)`；`NewWithClient` 接管 Resty 并关闭 Jar，已有 Jar 会话需要构造前显式导出。Cookie 快照和配置约束见[快速开始](#快速开始)，不要继续通过底层 Cookies 或默认 Cookie 请求头管理会话。

非 nil 的非结构体参数指针现在返回 `ParamError`，不再触发 panic；标签值中的等号完整保留。参数转换失败会返回错误而不是静默变为空字符串。JSON/multipart 混用或字段位置冲突会报错；依赖 Resty 内部 Body 为普通 map 的代码需适配已编码 JSON 字节或 multipart 构造方式。

### 错误包装

普通业务失败在原有 `Error{Code, Message}` 外补充 HTTP 方法和安全接口地址。WBI 非零业务码且缺少密钥时也采用此包装；非零码但存在密钥时仍沿用原有流程并继续检查密钥合法性。`Error` 的字段和自身错误文本不变，但请求返回错误的最外层类型及整体文本发生变化。将 `err.(bilibili.Error)` 或错误字符串匹配迁移为示例中的 `var be bilibili.Error; errors.As(err, &be)`，不要改为指针类型目标。

### 动态命名类型

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
log.Printf("关注状态原文=%s 数值=%d", rawStatus, status)
```

原动态对应字段为 `item.Orig.Modules.ModuleAuthor.Following`，读取方式相同。头像尺寸仍保留 `float64`，样本包含小数；原动态 `LikeIcon` 在该样本中全部为 `null`，不能据此判断其内部 `Id` 类型，因此未改动。调试记录不随 Git 提交分发。

## 开发与贡献

贡献约定见 [AGENTS.md](AGENTS.md)，命名规则见 [CONTRIBUTING.md](.github/CONTRIBUTING.md)。接口和模型按业务组织，共享请求、参数、响应和错误处理仍在同一个 `bilibili` 包内。问题记录应注明对应方法、错误类型及必要的脱敏响应片段，不提交 Cookie、凭证或完整调试记录。

- `go build ./...`：编译库、生成器及本地工具，不执行程序。
- `go vet ./...`：静态检查。
- `gofmt -s -w <文件.go>`：仅格式化修改的 Go 文件。
- `golangci-lint run`：使用仓库 v2 配置；工具未安装时应如实记录。

已有辅助测试使用标准库 `testing`。默认只做编译和静态检查，任务明确要求时才运行测试；不擅自运行本地工具或访问真实 API，不自动修改现有 CI 测试配置。此前代码重构通过编译、vet 和相关静态比对，但未做实机或 race 验证；不能据此宣称所有响应、取消、multipart 或并发场景均已验证。

被 Git 忽略的 `test/` 是可能操作真实账号的本地工具，不随库提交分发。`watchVideo` 已使用 `GetTopicFeed`，抽奖和视频心跳仍使用自定义请求；部分批量工具已贯通取消，配置加载和独立 HTTP 工具不在相同保证范围内。

## 声明

1. 本项目遵守 AGPL 开源协议。
2. 本项目基于 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)
   中描述的接口编写。请尊重该项目作者的努力，遵循该项目的开源要求，禁止一切商业使用。
3. **请勿滥用，本项目仅用于学习和测试！利用本项目提供的接口、文档等造成不良影响及后果与本人无关。**
4. 由于本项目的特殊性，可能随时停止开发或删档
5. 本项目为开源项目，不接受任何形式的催单和索取行为，更不容许存在付费内容

PS：目前，B站调用接口时强制使用 `https` 协议

## 上游历史参考

本 fork 源自 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili)。以下版本、徽章、统计和链接都指向上游，可能已不可访问，不代表本 fork 的发布版本、Go 要求或构建结果。

上游文档曾将 v2.1+ 标为需要 Go 1.23 以上，v2.0.0 标为支持 Go 1.19 以上。历史安装命令如下，仅用于获取上游代码，不用于安装本文描述的 fork：

```bash
go get -u github.com/CuteReimu/bilibili/v2
go get -u github.com/CuteReimu/bilibili/v2@v2.0.0
```

历史导入路径为 `github.com/CuteReimu/bilibili/v2`；更早版本见[上游 v1](https://github.com/CuteReimu/bilibili/tree/v1)。上游的 [issue 入口](https://github.com/CuteReimu/bilibili/issues/new/choose)和[贡献页面](https://github.com/CuteReimu/bilibili/contribute)仅作来源记录。

[![](https://img.shields.io/github/v/tag/CuteReimu/bilibili?label=release "最新版本")](https://github.com/CuteReimu/bilibili/tags)
![](https://img.shields.io/github/go-mod/go-version/CuteReimu/bilibili "语言")
[![](https://img.shields.io/github/stars/CuteReimu/bilibili?style=flat&color=yellow)](#star-history "stars")
[![](https://img.shields.io/github/actions/workflow/status/CuteReimu/bilibili/golangci-lint.yml?branch=master)](https://github.com/CuteReimu/bilibili/actions/workflows/golangci-lint.yml "代码分析")
[![](https://img.shields.io/github/contributors/CuteReimu/bilibili)](https://github.com/CuteReimu/bilibili/graphs/contributors "贡献者")
[![](https://img.shields.io/github/license/CuteReimu/bilibili)](https://github.com/CuteReimu/bilibili/blob/master/LICENSE "许可协议")

### Star History

<a href="https://star-history.com/#CuteReimu/bilibili&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=CuteReimu/bilibili&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=CuteReimu/bilibili&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=CuteReimu/bilibili&type=Date" />
 </picture>
</a>
