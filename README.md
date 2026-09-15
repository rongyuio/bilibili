# 哔哩哔哩 API Go 客户端

基于 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 继续维护的 Go 客户端，封装 Bilibili API，提供 Cookie 管理、WBI 签名、context 取消和结构化错误定位。

模块路径为 `github.com/rongyuio/bilibili`，要求 Go 1.26 或更新版本，依赖以 [go.mod](go.mod) 为准。接口可能随服务端变化。

- [维护状态](#维护状态)
- [快速开始](#快速开始)
- [常用接口](#常用接口)
- [自定义请求](#自定义请求)
- [错误处理](#错误处理)
- [详细文档](#详细文档)
- [开发与贡献](#开发与贡献)
- [声明](#声明)

## 维护状态

- 本仓库是 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 的 fork。上游**已归档**：README 标注 Deprecated、默认分支改名为 `deprecated`、issues/discussions/wiki 全部关闭，不再接受改动与合并。其停止维护的原因是所依据的接口文档仓库 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect) 也已归档。
- 本 fork **继续维护**：跟进新增接口，已实现接口在发现问题时修复。以自用为主，不承诺固定的支持范围与响应时间。
- 本仓库保留 fork 关联；`master` 已开启分支保护（改动须经 PR 合入，禁止强制推送与删除）。
- 版本处于 **v0 阶段**，允许在次版本号内引入破坏性变更。升级前请阅读[迁移指南](docs/migration.md)。

## 快速开始

### 安装

在已有 Go 项目的目录中执行：

```bash
go get github.com/rongyuio/bilibili@v0.3.0
```

`v0.3.0` 相对 `v0.2.0` 包含第五轮重构：字段命名修正（45 处，JSON 标签不变）、重复模型合并为别名、头像渲染树具名化与文件进一步细分，并把 golangci-lint 告警清零、CI 门禁转为阻断。更早的 `v0.2.0` 已含第三、四轮重构（初始缩写规范化 `Id`→`ID`、`Url`→`URL` 等、字段命名缺陷修正、参数名推导兼容修复）。v0 阶段 API 仍可能调整，升级前请阅读[迁移指南](docs/migration.md)。需要开发分支代码时可使用 `@master`，Go 会记录对应提交的伪版本。

### 创建客户端并调用接口

下面是一个完整程序，按需在自己的项目运行：

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/rongyuio/bilibili"
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

所有网络方法以 `ctx context.Context` 为首参，调用后先处理错误再使用结果。普通请求可并发，登录、会话切换及配置修改需串行；不要复制已使用的 Client。

### Cookie 与扫码登录

已有浏览器 Cookie 时，在发送请求前导入：

```go
client := bilibili.New()
client.SetRawCookies(cookieHeader) // cookieHeader 由调用方读取，不写入代码或日志。
```

`New()` 不联网。需要游客 Cookie 时使用 `NewAnonymousClient(ctx)` 并处理返回错误。扫码登录流程如下，`ctx` 应预留人工确认所需时间：

```go
qr, err := client.GetQRCode(ctx)
if err != nil {
    log.Printf("获取二维码失败: %v", err)
    return
}
qr.Print()
result, err := client.LoginWithQRCode(ctx, bilibili.LoginWithQRCodeParam{QrcodeKey: qr.QrcodeKey})
if err != nil {
    log.Printf("扫码登录失败: %v", err)
    return
}
if result.Code != 0 {
    log.Printf("扫码登录未完成，状态码: %d", result.Code)
    return
}
log.Println("登录成功")
```

登录成功后客户端自动保存 Cookie。保存与恢复、密码和短信登录、Resty 接管见 [认证与会话](docs/authentication.md)。

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
    log.Println(item.IDStr.String(), item.Modules.ModuleAuthor.Name)
}
```

`mid` 是调用方提供的 UID 字符串。需要继续读取时，根据 `page.HasMore` 将 `page.Offset` 传入下一次调用。调用方应检查取消及分页进度，避免空 offset 或重复 offset 导致循环。

上报观看时长使用 `ReportVideoWatchTime`，封装 `/x/click-interface/web/heartbeat`。它需要登录态：`mid` 取自 Cookie `DedeUserID`，CSRF 取自 `bili_jct`，两者缺失都会返回错误；请求启用 WBI 签名。观看时长与播放进度完全由调用方指定，库不设默认值、不自动重试、也不限制调用间隔。`Cid` 与 `VideoDuration` 可先用 `GetVideoInfo` 取得：

```go
info, err := client.GetVideoInfo(ctx, bilibili.VideoParam{Bvid: bvid})
if err != nil {
    log.Printf("获取视频信息失败: %v", err)
    return
}
err = client.ReportVideoWatchTime(ctx, bilibili.ReportVideoWatchTimeParam{
    Bvid:          bvid,
    Cid:           info.Cid,
    Realtime:      600,           // 本次上报的观看时长（秒）
    PlayedTime:    info.Duration, // 播放进度（秒）
    VideoDuration: info.Duration, // 视频总时长（秒）
})
if err != nil {
    log.Printf("上报观看时长失败: %v", err)
}
```

参数校验在发出请求前完成：`cid`、`realtime`、`video_duration` 必须大于 0，`played_time` 不能为负，`aid` 与 `bvid` 任选一个（`bvid` 需为 12 位）。

### 话题动态列表

`GetTopicFeed(ctx, param)` 封装 `/x/polymer/web-dynamic/v1/feed/topic`，返回一页响应的 `data`，复用统一 Cookie、context、参数编码和错误处理。参数全部位于 query，不额外启用 WBI 签名或 CSRF，也不自动重试或翻页。

```go
result, err := client.GetTopicFeed(ctx, bilibili.GetTopicFeedParam{
    TopicID:     topicId,
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

下一页将返回的 `TopicCardList.Offset` 传入参数，根据 `HasMore` 判断是否结束；检查取消、空游标和重复游标。

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

勋章面板按 `PageInfo.TotalPage` 控制分页，不单独依赖 `HasMore`。`ReportLiveLike` 上报直播点赞；账号、房间和数量由调用方指定，CSRF 由库填入。方法、参数见 [live.go](live.go)，响应模型见 [live_model.go](live_model.go)。

### 活动抽奖与动态抽奖信息

`GetActivityLotteryTimes` 获取当前账号的活动剩余抽奖次数；`DoActivityLottery` 执行活动抽奖；`GetDynamicLotteryInfo` 查询用户抽奖动态的抽奖信息。参数与响应模型均位于 [lottery.go](lottery.go)，查询结果直接对应 `data`。

```go
result, err := client.GetActivityLotteryTimes(ctx, bilibili.GetActivityLotteryTimesParam{Sid: sid})
if err != nil {
    log.Printf("查询活动抽奖次数失败: %v", err)
    return
}
log.Printf("剩余抽奖次数: %d", result.Times)
```

`DoActivityLottery` 返回 `error`，不解析中奖明细；需要明细可使用 `Client.Do`。活动循环和动态删除决策由调用方实现。

### 工具方法

网络方法需传入 context；AV/BV 转换为纯计算：

| 方法 | 用途 |
| --- | --- |
| `client.UnwrapShortURL(ctx, shortURL)` | 解析短链接，返回目标类型和标识 |
| `client.Now(ctx)` | 获取服务器时间 |
| `client.GetZoneLocation(ctx)` | 查询 IP 所属地理位置 |
| `client.GetRegionDailyCount(ctx)` | 获取分区当日投稿数 |
| `bilibili.Av2Bv(aid)` / `bilibili.Bv2Av(bvid)` | 纯计算转换，不需要 context |

其它接口按业务位于 `video.go`、`user.go`、`live.go` 等文件，响应模型单独放在对应的 `*_model.go`（如 `video_model.go`、`user_model.go`），可通过方法和参数注释查阅。内置请求是否省略参数由 `request` 标签决定，不能只根据 JSON 标签判断。

## 自定义请求

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

示例需导入 `github.com/rongyuio/bilibili`、`context`、`fmt`、`net/http`、`net/url`。`out` 接收 `data`，传 `nil` 只检查状态及业务错误；WBI 只签 query，CSRF 由调用方按接口要求提供。`Form` 与 `JSON` 互斥。参数规则与边界见 [请求与错误处理](docs/request.md)。

## 错误处理

通过标准库 `errors.Is` 判断取消和超时，通过 `errors.As` 提取错误类型：

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

`ParamError` 提供参数类型、字段和位置；`DecodeError` 提供 JSON 路径、Go 字段及预期类型。完整示例见 [请求与错误处理](docs/request.md)。记录错误时不要输出 Cookie、凭证或完整响应。

## 详细文档

| 文档 | 内容 |
| --- | --- |
| [认证与会话](docs/authentication.md) | 游客、扫码、密码与短信登录，Cookie 保存及 Resty 接管 |
| [请求与错误处理](docs/request.md) | 自定义请求、参数标签、错误分类与解码定位 |
| [迁移指南](docs/migration.md) | 模块路径、context 签名、会话规则和模型类型变更 |
| [版本策略](docs/versioning.md) | 版本号格式与禁用形式、递增规则、v0 与 v1 之后的兼容性约定 |

## 开发与贡献

### 本地开发

先克隆仓库。外部项目需要使用本地修改时，在调用项目的 `go.mod` 中配置：

```go.mod
require github.com/rongyuio/bilibili v0.0.0

replace github.com/rongyuio/bilibili => ../bilibili
```

相对路径以调用项目的 `go.mod` 为基准，请替换为实际路径。`v0.0.0` 是本地替换占位版本；已有依赖可以保留原版本，仅添加 `replace`。导入仍使用完整模块路径。

### 验证与贡献

- `go build ./...`：编译，不执行程序。
- `go vet ./...`：静态检查。
- `gofmt -s -w <文件.go>`：格式化修改的文件。
- `golangci-lint run`：使用仓库 v2 配置。

默认以编译和静态检查验证，不运行真实 API 或本地账号工具；已有测试使用标准库 `testing`，按任务要求执行。编译通过不等于实机或并发行为已验证。

贡献规则见 [CONTRIBUTING.md](.github/CONTRIBUTING.md)。反馈问题请说明方法、错误类型和必要的脱敏信息，不提交凭证。被忽略的 `test/` 为本地工具，不随仓库分发。

### 发版

维护者推送合法的语义化版本 tag（形如 `v<major>.<minor>.<patch>`）即触发 [release.yml](.github/workflows/release.yml)，无需手动创建 Release：

```bash
git switch master && git pull
git tag v0.4.0
git push origin v0.4.0
```

工作流依次执行：校验 tag 名符合语义化版本格式，并校验该 tag 指向的提交位于 `master`（任一不满足即拒绝发布）；运行 `go test -race ./...` 与 `go build ./...`；最后用 `gh release create --generate-notes` 创建 GitHub Release，发布说明的分类规则见 [.github/release.yml](.github/release.yml)。

tag 名含 `-` 后缀（如 `v0.4.0-rc.1`）时发布为 Pre-release，否则标记为 Latest。重复推送同一 tag 会先删除已有 Release 再重建。

版本号格式、递增规则与 v0 阶段的兼容性约定见[版本策略](docs/versioning.md)。

## 声明

1. 本项目遵守 **AGPL-3.0** 开源协议，fork 自 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili)。衍生作品需继续以 AGPL 授权并保留原作者署名；以网络服务方式对外提供时，需按要求向使用者提供源代码。
2. 本项目基于 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)
   中描述的接口编写。请尊重该项目作者的努力，遵循该项目的开源要求，禁止一切商业使用。
3. **请勿滥用，本项目仅用于学习和测试！利用本项目提供的接口、文档等造成不良影响及后果与本人无关。**
4. 由于本项目的特殊性，可能随时停止开发或删档
5. 本项目为开源项目，不接受任何形式的催单和索取行为，更不容许存在付费内容

PS：目前，B站调用接口时强制使用 `https` 协议
