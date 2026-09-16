# 请求与错误处理

[返回 README](../README.md)

本文代码片段中的 `ctx`、`client` 及业务参数由调用方提供；客户端导入路径为 `github.com/rongyuio/bilibili`。

## 调用尚未封装的接口

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

示例所需导入为 `github.com/rongyuio/bilibili`、`context`、`fmt`、`net/http`、`net/url`。

- `out` 接收响应的 `data`，不要再次包裹 `code/message/data`。传 `nil` 只检查 HTTP 状态及业务错误；传 `*json.RawMessage` 保留原始 `data`。
- `Query` 可与 `Form` 或 `JSON` 并用，但 `Form` 与 `JSON` 互斥。`JSON` 按标准库规则编码，包括字符串值；`Headers` 使用 `http.Header`。
- URL 自带查询参数会参与请求，同名键以 `Request.Query` 为准。WBI 只签查询参数，不签表单；签名请求的每个查询键必须只有一个值。CSRF 由调用方按接口要求放在查询或表单中。
- context 会传到 HTTP 请求和 WBI 密钥刷新；所有内置网络 API 同样以 context 为首参。此入口不额外启用重试；通过 Resty 自行配置的重试策略仍然有效。

## 内置接口的参数标签

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


## 分类与判断

标准业务响应先判断 `code`，非零时返回业务错误；成功后再解码 `data`，解码失败时不写入部分结果。单个字段的 JSON 类型与模型不符属于例外：该字段被丢弃并上报，其余字段照常写入，见[单个字段类型不符时的容错](#单个字段类型不符时的容错)。

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

示例需要导入 `github.com/rongyuio/bilibili`、`context`、`errors`、`log`。HTTP 错误按各接口原有规则判定：普通请求、游客初始化和 WBI 要求 200；短链接要求 302；Cookie 刷新页面接受 2xx。HTTP 状态错误不自动触发重试。网络故障、取消和超时保留原始错误链，不转换成 HTTPError；HTML 解析失败、缺少有效 Cookie、无效密钥等也保持独立错误。

`Message` 仍保留服务端消息，接口地址的清理不代表任意底层错误或服务端消息都经过脱敏。结构化记录优先选择接口、状态码和业务码；不要额外输出凭证、请求头或完整响应。

## 定位参数错误

```go
var pe *bilibili.ParamError
if errors.As(err, &pe) {
    log.Printf("参数类型=%s Go字段=%s 参数名=%s 位置=%s",
        pe.RootType, pe.GoField, pe.Parameter, pe.Location)
}
```

`GoField` 可为 `IDs[2]`；整体参数类型错误的字段、参数名和位置为空。JSON 编码错误定位到顶层参数字段，底层 `*json.MarshalerError` 等错误可继续解包；不会重复执行自定义编码器来探测内部路径。内置请求的错误外层还包含 HTTP 方法及去除查询参数的接口地址。

参数转换失败现在明确返回错误，不再静默变为空字符串。正常错误文本不包含参数值或原始错误文本；`ParamError.Err` 保留底层错误，可能含有原始值，不要直接写入日志。`Client.Do` 的手工 JSON 编码错误不转换为 `ParamError`。

## 定位反序列化失败

```go
var de *bilibili.DecodeError
if errors.As(err, &de) {
    log.Printf("接口=%s 类型=%s Go字段=%s JSON路径=%s 预期=%s 实际=%s 偏移=%d 精确=%t",
        de.Endpoint, de.RootType, de.GoField, de.JSONPath,
        de.Expected, de.Actual, de.Offset, de.Exact)
}
```

`JSONPath` 包含数组下标，例如 `$.data.items[3].modules.module_author.mid`；匿名结构通过根类型与 Go 字段链定位。`Offset` 从响应体第一个字节起按 1 计数，零表示不可用。原始错误保留在错误链中，包括可通过 `errors.As` 提取的 `*json.UnmarshalTypeError`。

正常响应只使用标准库解码，不额外遍历 JSON；解码失败后才会扫描：先尝试剪掉类型不兼容的字段重试，重试仍失败才生成诊断。诊断基于**未剪枝的原始响应体**，路径与偏移因此始终对应原始字节。自定义解码器不重复执行，最多定位到其字段边界，多个不确定边界退回共同父级并标记 `Exact=false`；语法错误只报告可用偏移量。诊断记录首个能确认的不匹配，不保证枚举所有问题。错误文本不包含字段值、完整响应或查询参数；JSON 路径中的 map 键仍来自响应，分享日志前请注意这一点。

先判断业务 `code`，再解码 `data`，避免业务错误被结果类型不匹配掩盖。解码失败时不向 `out` 写入部分结果。容错丢弃的字段保持零值，但会显式上报，不是静默转零。`out=nil` 时不会检查 `data` 的字段类型。

## 单个字段类型不符时的容错

服务端偶尔会改变某个字段的 JSON 类型（例如动态的 `following` 在未登录时返回 `null`、登录态返回布尔、更早的样本返回数字）。这类漂移只会丢弃该字段，不会作废整条响应：

- 严格解码失败后，库把 JSON 类型与 Go 类型不兼容的叶子替换为 `null` 再解码一次。
- 根节点（`data` 本身）不参与容错，接口整体形态变化仍然返回 `DecodeError`。
- 实现了自定义 `UnmarshalJSON` / `UnmarshalText` 的类型是不透明边界，其内部失败不被容错。
- map 键无法转换、JSON 语法错误等无法靠丢弃字段修复的情况仍然报错。
- 容错成功时调用返回 `nil`，被丢弃的字段保持零值。

需要感知漂移时注册回调：

```go
client.SetDroppedFieldHandler(func(field bilibili.DroppedField) {
    log.Printf("接口=%s 类型=%s JSON路径=%s 预期=%s 实际=%s 偏移=%d",
        field.Endpoint, field.RootType, field.JSONPath,
        field.Expected, field.Actual, field.Offset)
})
```

`DroppedField` 的字段与 `DecodeError` 对齐（`Method`、`Endpoint`、`RootType`、`GoField`、`JSONPath`、`Expected`、`Actual`、`Offset`），不含字段值。同一条响应可能触发多次调用，回调在请求的 goroutine 中执行，必须并发安全；不同请求之间也可能并发。传入 `nil` 只关闭上报，容错仍然开启。该设置与 Cookie、会话切换一样，必须在请求之外调用。

此前依赖「漂移字段会返回 `DecodeError`」做告警的代码，请改为在回调中记录。
