# 哔哩哔哩 API Go 客户端

[![Go](https://github.com/rongyuio/bilibili/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/rongyuio/bilibili/actions/workflows/golangci-lint.yml)
[![GoFmt](https://github.com/rongyuio/bilibili/actions/workflows/gofmt.yml/badge.svg)](https://github.com/rongyuio/bilibili/actions/workflows/gofmt.yml)
[![Latest tag](https://img.shields.io/github/v/tag/rongyuio/bilibili?label=latest&sort=semver)](https://github.com/rongyuio/bilibili/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/rongyuio/bilibili.svg)](https://pkg.go.dev/github.com/rongyuio/bilibili)
[![License: AGPL-3.0](https://img.shields.io/badge/License-AGPL--3.0-blue.svg)](LICENSE)

基于 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 继续维护的 Go 客户端，封装 Bilibili API，提供 Cookie 管理、WBI 签名、context 取消和结构化错误定位。

模块路径为 `github.com/rongyuio/bilibili`，要求 Go 1.26 或更新版本，依赖以 [go.mod](go.mod) 为准。接口可能随服务端变化。

- [维护状态](#维护状态)
- [特性](#特性)
- [快速开始](#快速开始)
- [常用接口](#常用接口)
- [自定义请求](#自定义请求)
- [错误处理](#错误处理)
- [详细文档](#详细文档)
- [开发与贡献](#开发与贡献)
- [License](#license)
- [声明](#声明)

## 维护状态

- 本仓库是 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 的 fork。上游**已归档**：README 标注 Deprecated、默认分支改名为 `deprecated`、issues/discussions/wiki 全部关闭，不再接受改动与合并。其所依据的接口文档仓库 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect) 也已归档。
- 本 fork **继续维护**：跟进新增接口，已实现接口在发现问题时修复。以自用为主，不承诺固定的支持范围与响应时间。
- 本仓库保留 fork 关联；`master` 已开启分支保护（改动须经 PR 合入，禁止强制推送与删除）。
- 版本处于 **v0 阶段**，允许在次版本号内引入破坏性变更。升级前请阅读[迁移指南](docs/migration.md)。

## 特性

- **Cookie 管理**：浏览器 Cookie 导入、游客初始化、扫码 / 密码 / 短信登录，以及 Cookie 的保存与恢复
- **WBI 签名**：按需对 query 参数签名，密钥自动缓存并按周期刷新
- **统一请求链路**：全部内置接口共用同一套 Cookie 快照、参数编码、context 取消与错误处理
- **结构化错误诊断**：`HTTPError` / `Error` / `ParamError` / `DecodeError` 均支持 `errors.As`，解码失败时给出 JSON 路径与偏移
- **容错解码**：单个字段的 JSON 类型与模型不符时只丢弃该字段并经 `SetDroppedFieldHandler` 上报，不连累整条响应
- **逃生通道**：尚未封装的接口可通过 `Client.Do` 复用同一套能力

## 快速开始

### 安装

在已有 Go 项目的目录中执行：

```bash
go get github.com/rongyuio/bilibili@latest
```

v0 阶段 API 仍可能调整，升级前请阅读[迁移指南](docs/migration.md)与[版本策略](docs/versioning.md)。需要开发分支代码时可使用 `@master`，Go 会记录对应提交的伪版本。

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

`GetHomePage` 访问一次 B 站首页，可补全 `buvid3` 等设备字段；`RefreshWebCookie(ctx, refreshToken)` 一条龙完成官方 Cookie 刷新（不需要刷新时返回 nil），刷新成功后务必持久化返回值中的新 `RefreshToken`，否则之后无法再次刷新。

`New()` 不联网；需要游客 Cookie 时用 `NewAnonymousClient(ctx)`。扫码登录用 `GetQRCode` 取码、`qr.Print()` 展示、再用 `LoginWithQRCode` 轮询结果，登录成功后客户端自动保存 Cookie；`ctx` 应预留人工确认所需时间。密码与短信登录、Cookie 的保存与恢复、Resty 接管见 [认证与会话](docs/authentication.md)。

## 常用接口

以下示例聚焦调用方式，省略了错误判断；完整写法见上文的「创建客户端并调用接口」。

### 视频与空间动态

视频详情可使用快速开始中的 `GetVideoInfo`。空间动态按页获取：

```go
page, err := client.GetUserSpaceDynamic(ctx, bilibili.GetUserSpaceDynamicParam{
    HostMid:        mid,
    TimezoneOffset: -480,
    Features:       "itemOpusStyle",
})
for _, item := range page.Items {
    log.Println(item.IDStr.String(), item.Modules.ModuleAuthor.Name)
}
```

`mid` 是调用方提供的 UID 字符串。需要继续读取时，根据 `page.HasMore` 将 `page.Offset` 传入下一次调用。调用方应检查取消及分页进度，避免空 offset 或重复 offset 导致循环。

上报观看时长使用 `ReportVideoWatchTime`，封装 `/x/click-interface/web/heartbeat`。它需要登录态：`mid` 取自 Cookie `DedeUserID`，CSRF 取自 `bili_jct`，两者缺失都会返回错误；请求启用 WBI 签名。观看时长与播放进度完全由调用方指定，库不设默认值、不自动重试、也不限制调用间隔。`Cid` 与 `VideoDuration` 可先用 `GetVideoInfo` 取得：

```go
info, err := client.GetVideoInfo(ctx, bilibili.VideoParam{Bvid: bvid})
err = client.ReportVideoWatchTime(ctx, bilibili.ReportVideoWatchTimeParam{
    Bvid:          bvid,
    Cid:           info.Cid,
    Realtime:      600,           // 本次上报的观看时长（秒）
    PlayedTime:    info.Duration, // 播放进度（秒）
    VideoDuration: info.Duration, // 视频总时长（秒）
})
```

参数校验在发出请求前完成：`cid`、`realtime`、`video_duration` 必须大于 0，`aid` 与 `bvid` 任选一个（`bvid` 需为 12 位）；`played_time` 传 `VideoPlayedComplete`（-1）表示已看完。清晰度由库固定为 `VideoQuality720P`，如需其他清晰度见 [video_stream_model.go](video_stream_model.go) 中的 `VideoQuality*` 常量。

### 话题动态列表

`GetTopicFeed(ctx, param)` 封装 `/x/polymer/web-dynamic/v1/feed/topic`，返回一页响应的 `data`，复用统一 Cookie、context、参数编码和错误处理。参数全部位于 query，不额外启用 WBI 签名或 CSRF，也不自动重试或翻页；`WebLocation` 留空时由库填入默认值。

```go
result, err := client.GetTopicFeed(ctx, bilibili.GetTopicFeedParam{
    TopicID:     topicId,
    SortBy:      bilibili.TopicFeedSortByLatest,
    PageSize:    20,
    Features:    "itemOpusStyle,listOnlyfans,opusBigCover,onlyfansVote,decorationCard",
})
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
log.Printf("本页普通勋章数: %d，总页数: %d", len(panel.List), panel.PageInfo.TotalPage)
```

勋章面板按 `PageInfo.TotalPage` 控制分页，不单独依赖 `HasMore`。`ReportLiveLike` 上报直播点赞；账号、房间和数量由调用方指定，CSRF 由库填入。方法、参数见 [live.go](live.go)，响应模型见 [live_model.go](live_model.go)。

### 活动抽奖与动态抽奖信息

`GetActivityLotteryTimes` 获取当前账号的活动剩余抽奖次数；`DoActivityLottery` 执行活动抽奖；`GetDynamicLotteryInfo` 查询用户抽奖动态的抽奖信息。参数与响应模型均位于 [lottery.go](lottery.go)，查询结果直接对应 `data`；`x-bili-device-req-json` 由库按 `WebLocation` 生成，无需调用方传入。

```go
result, err := client.GetActivityLotteryTimes(ctx, bilibili.GetActivityLotteryTimesParam{Sid: sid})
log.Printf("剩余抽奖次数: %d", result.Times)
```

`DoActivityLottery` 返回 `error`，不解析中奖明细；需要明细可使用 `Client.Do`。活动循环和动态删除决策由调用方实现。

### 直播间列表与天选时刻

`GetLiveWebRoomList` 获取直播间列表（网页端新接口，WBI 签名）。旧的 second/getList 接口已被风控拦截（-352），请勿再使用。天选时刻的房间集中在互动玩法下的天选分区：

```go
list, err := client.GetLiveWebRoomList(ctx, bilibili.GetLiveWebRoomListParam{
    ParentAreaID: 15, AreaID: 1457, Page: 1, // 1457 为天选聚集分区
})
for _, module := range list.RoomList {
    for _, room := range module.List {
        check, err := client.CheckLiveAnchorLottery(ctx, bilibili.CheckLiveAnchorLotteryParam{RoomID: room.RoomID})
        // check.Status==1 且 GiftPrice==0 且 RequireType 为 0/1 时可免消耗参与：
        // client.JoinLiveAnchorLottery(ctx, bilibili.JoinLiveAnchorLotteryParam{...})
    }
}
```

`GetLiveHotRankList` 获取首页人气榜，条目中的 `LotStatus`/`RedPocketStatus` 可低成本初筛有抽奖或红包活动的房间。

### Web 端观看时长上报

当前网页播放器已把观看时长上报迁移至 `data.bilivideo.com` 域的新端点，对应本库的 `ReportWebWatchEnter`（进房）与 `ReportWebWatchHeartBeat`（心跳）。流程：进房响应下发 `STKY`（心跳密钥，每跳轮换）、`SID`（会话 id）与 `HBIL`（间隔秒数）；之后按间隔循环发送心跳，每跳使用上一响应轮换出的 `STKY`。

心跳表单中的 `Csn` 字段由官方 skynet wasm 对其余字段的 JSON 计算签名，本库不内置该算法，也不打包 wasm 文件：调用方需自行下载 wasm（地址随页面版本变化，可从直播间页面资源中解析）并用 wasm 运行时（如 wazero）调用其 `skynet` 导出函数，输入为表单 JSON 字符串，输出为签名串。同一账号可同时向多个直播间发送心跳，每个直播间的会话状态（STKY/SID/QID 序号）互相独立。

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

尚未封装的接口可以用 `Client.Do`，复用客户端的 Cookie、网络配置、签名与解码流程：

```go
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
```

示例需导入 `github.com/rongyuio/bilibili`、`context`、`net/http`、`net/url`。`out` 接收 `data`，传 `nil` 只检查状态及业务错误；WBI 只签 query，CSRF 由调用方按接口要求提供；`Form` 与 `JSON` 互斥。参数规则与边界见 [请求与错误处理](docs/request.md)。

## 错误处理

错误类型按 `errors.Is` / `errors.As` 设计：网络故障、取消与超时保留原始错误链，`*HTTPError`、`Error`、`*ParamError`、`*DecodeError` 可依次提取。

```go
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
```

`ParamError` 提供参数类型、字段与位置；`DecodeError` 提供 JSON 路径、Go 字段及预期类型。完整示例见 [请求与错误处理](docs/request.md)；记录错误时不要输出 Cookie、凭证或完整响应。

单个字段的 JSON 类型与模型不符时，该字段被丢弃（保持零值）并上报，其余字段照常解码；`DecodeError` 只在这类容错也失败时返回。需要感知漂移时注册回调：

```go
client.SetDroppedFieldHandler(func(field bilibili.DroppedField) {
    log.Printf("接口=%s JSON路径=%s 预期=%s 实际=%s",
        field.Endpoint, field.JSONPath, field.Expected, field.Actual)
})
```

回调在同一条响应内可能被多次调用，且可能在并发请求中执行。规则边界见 [请求与错误处理](docs/request.md#单个字段类型不符时的容错)。

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

相对路径以调用项目的 `go.mod` 为基准，请替换为实际路径。`replace` 生效后实际使用的是本地代码，`require` 里的版本号不会被解析；`v0.0.0` 只是简化的占位写法（`go mod tidy` 写成的是等价形式的伪版本 `v0.0.0-00010101000000-000000000000`）。若该依赖本就存在于 `go.mod`，保留原版本号、只添加 `replace` 即可。导入仍使用完整模块路径。

### 验证与贡献

- `go build ./...`：编译，不执行程序。
- `go vet ./...`：静态检查。
- `gofmt -s -w <文件.go>`：格式化修改的文件。
- `golangci-lint run`：使用仓库 v2 配置。
- `go test ./...`：运行单元测试；CI 使用 `go test -race -v ./...`。

测试均为纯逻辑单元测试，**不会**访问真实 API，也不需要凭证；CI 在向 `master` 的 push 与 PR 上运行上述全部检查。编译与测试通过不等于实机或并发行为已验证——`test/` 下的账号工具需自行运行。

贡献规则见 [CONTRIBUTING.md](.github/CONTRIBUTING.md)；安全问题请按 [SECURITY.md](.github/SECURITY.md) 私下报告，参与讨论请遵守 [CODE_OF_CONDUCT.md](.github/CODE_OF_CONDUCT.md)。反馈问题请说明方法、错误类型和必要的脱敏信息，不提交凭证。被忽略的 `test/` 为本地工具，不随仓库分发。

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

## License

本项目以 **AGPL-3.0** 授权，详见 [LICENSE](LICENSE)。作为 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 的衍生作品，需继续以 AGPL 授权并保留原作者署名；以网络服务方式对外提供时，需按要求向使用者提供源代码。

## 声明

1. 本项目基于 [SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)
   中描述的接口编写。请尊重该项目作者的努力，遵循该项目的开源要求，禁止一切商业使用。
2. **请勿滥用，本项目仅用于学习和测试！利用本项目提供的接口、文档等造成不良影响及后果与本人无关。**
3. 由于本项目的特殊性，可能随时停止开发或删档
4. 本项目为开源项目，不接受任何形式的催单和索取行为，更不容许存在付费内容

PS：目前，B站调用接口时强制使用 `https` 协议
