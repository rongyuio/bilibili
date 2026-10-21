# 迁移指南

[返回 README](../README.md)

本文代码片段中的 `ctx`、`root` 及业务参数由调用方提供。根包导入路径为 `github.com/rongyuio/bilibili`；业务接口从对应子包导入，例如 `github.com/rongyuio/bilibili/video`、`github.com/rongyuio/bilibili/login`。

## 包结构：根包与业务子包

根包 `bilibili` 只保留基础设施：

- `Client` 的构造与会话：`New()`、`NewAnonymousClient(ctx)`、`NewWithClient(restyClient)`；
- 统一请求链路与参数编码：`Client.Do`、`Request`、`request` 结构体标签，以及供子包复用的 `Execute`、`ExecuteRaw`、`FillCsrf`、`FillParam` 等导出设施；
- 错误类型：`HTTPError`、`Error`、`ParamError`、`DecodeError`；
- Cookie 管理：`SetRawCookies`、`GetCookiesString`、`SetCookiesString`、`GetCookies`、`SetCookie`、`SetCookies`；
- WBI 签名：`WBI`、`NewDefaultWbi()`；
- 解码容错：`DroppedField`、`SetDroppedFieldHandler`；
- 跨业务域复用的基础类型：`NumberOrString`、`Size` 等。

业务接口按业务域拆分到 34 个子包，完整清单见 [README 的「包结构」](../README.md#包结构)。

每个子包提供 `New(*bilibili.Client) *Client`，返回嵌入了根包 `Client` 的调用入口。所有子包共享同一个根包客户端的 Cookie、WBI 缓存与错误处理，不会新建会话：

```go
root := bilibili.New()
v := video.New(root)
info, err := v.GetVideoInfo(ctx, video.GetVideoInfoParam{Bvid: "BV1L9Uoa9EUx"})
if err != nil {
    log.Printf("获取视频失败: %v", err)
    return
}
log.Println(info.Title)
```

每个接口的请求参数与响应结果都在所属子包内独立定义，接口之间不共享返回体；即使字段相同，也各自定义。参数类型通常命名为 `XxxParam`，结果类型通常命名为 `XxxResult`；没有参数时传 `nil`。可能为 `null` 的响应字段使用指针，避免 `json.Unmarshal` 失败。

## 从扁平包迁移

旧版把所有接口方法直接定义在根包 `*bilibili.Client` 上，`*Param` 与 `*Result` 类型也位于根包。迁移按以下步骤进行：

1. 为用到的业务域导入子包。
2. 用 `root := bilibili.New()` 创建根客户端，再用 `<domain>.New(root)` 构造业务入口。
3. 把 `client.Xxx(...)` 改为业务入口上的 `Xxx(...)`，把 `bilibili.XxxParam` / `bilibili.XxxResult` 改为 `<domain>.XxxParam` / `<domain>.XxxResult`。
4. Cookie、WBI、错误处理与 Resty 配置仍通过 `root` 操作。

常见对应关系：

| 旧写法（扁平包） | 新写法（子包） |
| --- | --- |
| `client.GetQRCode(ctx)` | `l := login.New(root); l.GetQRCode(ctx)` |
| `bilibili.LoginWithQRCodeParam` | `login.QRCodeLoginParam` |
| `client.LoginWithPassword(ctx, bilibili.LoginWithPasswordParam{...})` | `l.LoginWithPassword(ctx, login.PasswordLoginParam{...})` |
| `bilibili.SendSMSParam` / `bilibili.LoginWithSMSParam` | `login.SendSMSParam` / `login.SMSLoginParam` |
| `client.GetVideoInfo(ctx, bilibili.VideoParam{...})` | `v := video.New(root); v.GetVideoInfo(ctx, video.GetVideoInfoParam{...})` |
| `client.GetMyUserSpaceDetail(ctx)` | `u := user.New(root); u.GetMyUserSpaceDetail(ctx)` |
| `client.SetRawCookies(...)`、`client.GetCookiesString()`、`client.SetCookiesString(...)`、`client.SetCookies(...)` | `root.SetRawCookies(...)`、`root.GetCookiesString()`、`root.SetCookiesString(...)`、`root.SetCookies(...)` |
| `client.Resty()` | `root.Resty()` |
| `client.SetDroppedFieldHandler(...)` | `root.SetDroppedFieldHandler(...)` |
| `bilibili.New()`、`bilibili.NewAnonymousClient(ctx)`、`bilibili.NewWithClient(...)` | 不变 |

表中 Cookie、`Resty()` 与 `SetDroppedFieldHandler` 三行涉及的方法仍在 `*bilibili.Client` 上，签名没有变化；变化只是示例中的接收者变量从旧版的 `client` 改为新写法中的 `root`。

登录相关的完整示例见[认证与会话](authentication.md)。

## 模块路径迁移

原来使用本地模块 `bilibili` 的项目，需要将导入路径统一替换为 `github.com/rongyuio/bilibili`，并更新 `go.mod` 中对应的 `require` 和 `replace`。远程接入请移除旧的本地替换，再执行快速开始中的安装命令；继续本地联调则按 [README 的本地开发说明](../README.md#本地开发)配置。

模块路径迁移与包拆分是两件事：模块路径决定从哪个仓库取代码，子包决定导入哪个业务域。模块路径变更不改变子包名称，子包导入路径统一为 `github.com/rongyuio/bilibili/<domain>`。

## Context 签名

所有可能联网的方法都以 `ctx context.Context` 为首参，子包方法沿用同一约定；根包的 `Do` 与 `NewAnonymousClient` 一直使用该签名，未发生变化：

```go
v.GetVideoInfo(ctx, video.GetVideoInfoParam{Bvid: "BV1L9Uoa9EUx"})
u.GetMyUserSpaceDetail(ctx)
```

Cookie 读写、配置和纯计算方法不变。

独立使用 WBI 时，通过 `wbi := bilibili.NewDefaultWbi()` 构造签名器，再调用 `wbi.GetKeys(ctx)`、`wbi.GetMixinKey(ctx)`、`wbi.SignQuery(ctx, query, ts)`、`wbi.SignMap(ctx, payload, ts)`，因为签名可能触发密钥刷新。没有新增 `XxxContext` 或无 context 的兼容包装；调用方接口声明、方法表达式和回调类型也需同步修改。

## 会话与参数编码

`NewAnonymousClient(ctx)` 的签名是 `(ctx context.Context) (*Client, error)`，没有变化；`NewWithClient` 接管 Resty 并关闭 Jar，已有 Jar 会话需要在构造前显式导出。Cookie 快照和配置约束见[认证与会话](authentication.md)，不要继续通过底层 Cookies 或默认 Cookie 请求头管理会话。

非 nil 的非结构体参数指针返回 `ParamError`，不再触发 panic；标签值中的等号完整保留。参数转换失败会返回错误而不是静默变为空字符串。JSON/multipart 混用或字段位置冲突会报错；依赖 Resty 内部 Body 为普通 map 的代码需适配已编码 JSON 字节或 multipart 构造方式。子包方法同样使用根包的 `request` 标签规则，完整说明见[请求与错误处理](request.md)。

## 错误分类

错误类型都定义在根包，子包方法返回的也是这些类型：

- HTTP 非 200：`*bilibili.HTTPError`，包含 `Method`、`Endpoint`、`StatusCode`；`Endpoint` 已去掉 query、fragment 与 user 信息。
- 响应信封中的业务码非零：`bilibili.Error`（值类型），包含 `Code`、`Message`。
- 参数编码失败：`*bilibili.ParamError`。
- 解码失败：`*bilibili.DecodeError`，包含 JSON 路径、Go 字段、预期/实际类型与偏移。
- 网络故障、取消与超时：保留原始错误链，不转换成 `HTTPError`。

将 `err.(bilibili.Error)` 或错误字符串匹配迁移为 `errors.As`：

```go
var be bilibili.Error
if errors.As(err, &be) {
    log.Printf("业务错误: %d %s", be.Code, be.Message)
}
```

`Error` 是值类型，`errors.As` 的目标必须是非指针的 `bilibili.Error`，不要写成 `*bilibili.Error`。其余类型按需提取，例如 `var he *bilibili.HTTPError`、`var pe *bilibili.ParamError`、`var de *bilibili.DecodeError`。

## NumberOrString

`bilibili.NumberOrString` 保留数字、字符串（包括 `"--"`）、布尔值与 `null`，`Kind()` 返回 `number`、`string`、`boolean` 或 `null`。数值转换会显式返回错误，布尔与 `"--"` 都会失败；JSON 再编码保留原始类别；零值表示 `null`。

该类型用于线上类型不稳定的字段，例如 `video.GetVideoStatusNumberResult.View`（播放次数可能是数字或 `"--"`）和动态作者模型中的 `Following`。原有直接赋值 `json.Number` 或强制转换为字符串的代码需要迁移到 `Kind()`、`String()`、`Int64()`、`Float64()`；需要构造值时可使用 `json.Unmarshal`。

其余字段不批量改型。`json.Number` 支持数字及数字字符串，不支持任意文本。

## 动态关注状态字段

动态接口的作者模型（如 `dynamic.GetDynamicFeedAllModuleAuthor`、`dynamic.GetUserSpaceDynamicModuleAuthor` 及其原动态对应类型）中的 `Following` 字段为 `bilibili.NumberOrString`。线上返回的 JSON 类型不固定：未登录为 `null`，登录态为布尔，部分样本为数字 `1`/`2`。本库不定义状态常量，也不把非零值解释为“已关注”。

判断原始类别用 `Kind()`，取值用 `String()`；只有确认是数字时才调用 `Int64()` 并处理错误：

```go
switch following := item.Modules.ModuleAuthor.Following; following.Kind() {
case "boolean":
    log.Printf("关注状态=%s", following.String())
case "number":
    status, err := following.Int64()
    if err != nil {
        log.Print("关注状态无法转换为整数")
        return
    }
    log.Printf("关注状态数值=%d", status)
default:
    log.Printf("关注状态缺失，原文=%q", following.String())
}
```

`item` 是当前接口的条目类型（如 `dynamic.GetDynamicFeedAllItem`）；原动态对应字段为 `item.Orig.Modules.ModuleAuthor.Following`。

## 解码容错

单个字段的 JSON 类型与模型不符时，库会丢弃该字段（保持零值）并上报，其余字段照常解码，调用返回成功；只有 `data` 根节点类型不匹配、JSON 语法错误、map 键无法转换等情况才返回 `*DecodeError`。完整规则与 `SetDroppedFieldHandler` 回调见[请求与错误处理](request.md#单个字段类型不符时的容错)。

## 旧模型名的处理

当前版本按接口独立命名返回体，根包不再提供 `DynamicItem`、`AudioOrVideo`、`SearchRespResult` 这类跨接口共享模型。使用旧模型名的代码需要改为对应接口的返回类型，例如：

| 旧模型 | 当前类型 |
| --- | --- |
| `DynamicItem` | 各动态接口的条目类型，如 `dynamic.GetDynamicFeedAllItem`、`dynamic.GetUserSpaceDynamicItem` |
| `AudioOrVideo` | `video.GetVideoStreamAudioOrVideo` |
| `SearchRespResult` | `search.IntegratedSearchResultItem` |

具体类型名以子包源码和方法签名为准。

## APP 端接口清理

以下 19 个接口均为 APP 域名（`app.bilibili.com`、`app.biliapi.com`、`app.biliapi.net`、`club.bilibili.com`），由接口补齐时从文档引入，文档中没有独立实测记录。库不提供 appkey 签名，调用方必须自行计算 `appkey` / `sign` 才能使用，实际不可直接调用，现已全部删除：

| 已删除 | 替代 |
| --- | --- |
| `appwidget.GetSplashBrandList` | 无（APP 开屏） |
| `appwidget.GetSplashList` | 无 |
| `appwidget.GetVersionInfo` | 无 |
| `appwidget.GetVersionUpgrade` | 无 |
| `clientinfo.GetResourceIP` | `clientinfo.GetIPInfoNew` / `GetWebRoomIPInfo` |
| `dynamic.GetRecommendTopic` | 无 |
| `dynamic.GetTopicDetails` | 无 |
| `garb.GetGarbColorInfoFromApp` | `garb.GetGarbColorInfo` |
| `garb.GetGarbSkin` | 无 |
| `login.GetAppAccountMyInfo` | `login.GetNavInfo` / `GetHomeUserInfo` |
| `search.GetSearchHotListApp` | `search.GetSearchHotListWeb` |
| `user.GetUserFollowings2` | `user.GetUserFollowings` / `GetUserFollowings3` |
| `user.GetUserSpaceArchiveCursor` | 无 |
| `user.GetUserSpaceCoinVideoListApp` | `user.GetUserSpaceCoinVideoList` |
| `user.GetUserSpaceLikeVideoListApp` | `user.GetUserSpaceLikeVideoList` |
| `video.GetShortVideoList` | 无 |
| `video.GetStoryVideoList` | 无 |
| `video.GetVideoOnlineInfoApp` | `video.GetVideoOnlineInfo` |
| `video.GetVideoSnapshotApp` | `video.GetVideoSnapshot` |

`appwidget` 子包保留 `GetAppVersion`（走 `api.bilibili.com`，不需要 APP 签名）。web 接口上原有的可选 `access_key` 参数不受影响，不传时仍走 Cookie。

## 文件组织

每个业务子包内：

- `client.go` 定义 `Client` 与 `New(*bilibili.Client) *Client`；
- 按接口或子领域命名的 `.go` 文件定义方法与 `*Param`；
- 响应模型放在同目录的 `*_model.go` 或 `model_*.go` 中，由该接口独立定义，例如 `video/model_info.go`、`video/model_stream.go`、`dynamic/model_dynamic_feed_all.go`、`live/live_model.go`。

根包只保留跨业务域的基础设施与基础类型；新增业务接口应放入对应子包，不要向根包添加业务方法。

## 最低 Go 版本

最低支持版本从 Go 1.26 提升到 **Go 1.27.3**。代码没有改用新语法，提升是为了让 `crypto/x509` 带上 CVE-2026-111943 的修复（受损中间证书在 `Certificate.Verify` 时未被拒绝），本库的 HTTPS 请求依赖标准库的证书链验证。

使用更旧工具链时，`go` 命令（`GOTOOLCHAIN=auto`）会自动下载 1.27.3；显式设置 `GOTOOLCHAIN=local` 的调用方需要手动升级 Go。

## 版本策略

本模块处于 v0 阶段：破坏性变更随次版本号发布，不需要升主版本，也不需要改模块路径。发布 `v1.0.0` 之后再出现破坏性变更，才需要升主版本并把模块路径改为 `github.com/rongyuio/bilibili/v2` 这类形式。

包的拆分、接口在子包之间的移动、`*Param` / `*Result` 的改名都属于破坏性变更。完整规则见[版本策略](versioning.md)。
