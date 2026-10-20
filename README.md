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
- [包结构](#包结构)
- [快速开始](#快速开始)
- [常用接口](#常用接口)
- [自定义请求](#自定义请求)
- [错误处理](#错误处理)
- [详细文档](#详细文档)
- [开发与贡献](#开发与贡献)
- [License](#license)
- [声明](#声明)

## 维护状态

- 本仓库是 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 的 fork。上游已被作者删除（该地址现返回 404）；删除前已归档：README 标注 Deprecated、默认分支改名为 `deprecated`、issues/discussions/wiki 全部关闭，不再接受改动与合并。其所依据的接口文档仓库也已停止维护。
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

## 包结构

根包 `bilibili` 提供基础设施：`Client` 的构造与会话、统一请求链路、参数编码、错误类型、解码诊断、Cookie 管理、WBI 签名，以及供子包复用的 `Execute` / `ExecuteRaw` / `FillCsrf` 等导出设施。根包的 `type.go` 只保留真正跨业务域的基础类型。

业务接口按域拆分到子包。每个子包提供 `New(*bilibili.Client) *Client`，返回嵌入了根包 `Client` 的调用入口，因此同一份 Cookie、WBI 缓存与错误处理在所有子包之间共享：

- `activity` —— 活动主题与活动列表
- `appwidget` —— APP 组件
- `article` —— 专栏文章与文集
- `audio` —— 音频、歌单与音频排行
- `bangumi` —— 番剧、影视、时间表与视频流
- `blackroom` —— 小黑屋与风纪仲裁
- `broadcast` —— 广播服务器
- `cheese` —— 课程与课程视频流
- `clientinfo` —— 客户端 IP 归属地
- `comment` —— 评论列表与评论操作
- `creativecenter` —— 创作中心、稿件数据与投稿辅助
- `customerservice` —— 客服心跳
- `danmaku` —— 视频弹幕发送、查询与历史弹幕
- `dynamic` —— 动态流、空间动态、话题、发布与转发
- `electric` —— 充电
- `emoji` —— 表情包与表情列表
- `fav` —— 收藏夹与收藏内容
- `garb` —— 装扮、装扮套装与装扮活动
- `history` —— 历史记录与稍后再看
- `live` —— 直播信息、勋章、天选、弹幕与观看心跳
- `login` —— 扫码 / 密码 / 短信登录与会话
- `lottery` —— 活动抽奖与动态抽奖信息
- `message` —— 私信、消息与会话设置
- `miscellaneous` —— 短链解析、服务器时间、IP 归属、分区投稿数与 AV/BV 转换
- `newbieexam` —— 新用户答题
- `note` —— 视频笔记
- `opus` —— 图文与用户空间图文
- `search` —— 搜索、热词与搜索建议
- `user` —— 用户信息、关系、空间与粉丝
- `video` —— 视频信息、播放地址、互动、合集与观看上报
- `videoranking` —— 视频排行、每周必看与分区排行
- `vip` —— 大会员权益、签到与积分
- `wallet` —— 钱包余额
- `webwidget` —— 网页顶栏与分区横幅

各接口的返回体按接口独立定义，接口之间不共享模型；即使字段相同，也在各自的 `model_*.go` 或 `<主题>_model.go` 中单独定义。

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
    "github.com/rongyuio/bilibili/video"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    root := bilibili.New()
    v := video.New(root)
    info, err := v.GetVideoInfo(ctx, video.GetVideoInfoParam{Bvid: "BV1L9Uoa9EUx"})
    if err != nil {
        log.Printf("获取视频失败: %v", err)
        return
    }
    log.Println(info.Title)
}
```

`bilibili.New()` 只负责会话与网络配置；`video.New(root)` 用同一个根包客户端构造业务子包入口，不会新建会话。所有网络方法以 `ctx context.Context` 为首参，调用后先处理错误再使用结果。普通请求可并发，登录、会话切换及配置修改需串行；不要复制已使用的 `Client`。

### Cookie 与扫码登录

已有浏览器 Cookie 时，在发送请求前导入：

```go
root := bilibili.New()
root.SetRawCookies(cookieHeader) // cookieHeader 由调用方读取，不写入代码或日志。
```

`root.GetHomePage(ctx)` 访问一次 B 站首页，可补全 `buvid3` 等设备字段；`root.RefreshWebCookie(ctx, refreshToken)` 一条龙完成官方 Cookie 刷新（不需要刷新时返回 nil），刷新成功后务必持久化返回值中的新 `RefreshToken`，否则之后无法再次刷新。

`bilibili.New()` 不联网；需要游客 Cookie 时用 `bilibili.NewAnonymousClient(ctx)`。扫码登录在 `login` 子包：

```go
root := bilibili.New()
l := login.New(root)

qr, err := l.GetQRCode(ctx)
if err != nil {
    log.Fatal(err)
}
qr.Print() // 在终端展示二维码

result, err := l.LoginWithQRCode(ctx, login.QRCodeLoginParam{QrcodeKey: qr.QrcodeKey})
if err != nil {
    log.Fatal(err)
}
if result.Code != 0 {
    log.Printf("登录未完成: %s", result.Message)
}
// 登录成功后 Cookie 已自动合并进 root，后续子包共用同一会话。
```

`LoginWithQRCode` 会轮询到扫码成功或二维码失效，`ctx` 应预留人工确认所需时间。密码与短信登录、Cookie 的保存与恢复、Resty 接管见 [认证与会话](docs/authentication.md)。

## 常用接口

以下示例聚焦调用方式，省略了错误判断；完整写法见上文的「创建客户端并调用接口」。示例中的 `root` 均由 `bilibili.New()` 创建，子包入口按需构造，例如 `v := video.New(root)`、`d := dynamic.New(root)`、`l := live.New(root)`。

### 视频与空间动态

视频详情可使用快速开始中的 `GetVideoInfo`。空间动态按页获取：

```go
v := video.New(root)
d := dynamic.New(root)

page, err := d.GetUserSpaceDynamic(ctx, dynamic.GetUserSpaceDynamicParam{
    HostMid:        mid,
    TimezoneOffset: -480,
    Features:       "itemOpusStyle",
})
for _, item := range page.Items {
    log.Println(item.IDStr.String(), item.Modules.ModuleAuthor.Name)
}
```

`mid` 是调用方提供的 UID 字符串。需要继续读取时，根据 `page.HasMore` 将 `page.Offset` 传入下一次调用。调用方应检查取消及分页进度，避免空 offset 或重复 offset 导致循环。

上报观看时长使用 `v.ReportVideoWatchTime`，封装 `/x/click-interface/web/heartbeat`。它需要登录态：`mid` 取自 Cookie `DedeUserID`，CSRF 取自 `bili_jct`，两者缺失都会返回错误；请求启用 WBI 签名。观看时长与播放进度完全由调用方指定，库不设默认值、不自动重试、也不限制调用间隔。`Cid` 与 `Duration` 可先用 `GetVideoInfo` 取得：

```go
info, err := v.GetVideoInfo(ctx, video.GetVideoInfoParam{Bvid: bvid})
err = v.ReportVideoWatchTime(ctx, video.ReportVideoWatchTimeParam{
    Bvid:          bvid,
    Cid:           info.Cid,
    Realtime:      600,           // 本次上报的观看时长（秒）
    PlayedTime:    600,           // 播放进度（秒），与 Realtime 一致
    VideoDuration: info.Duration, // 视频总时长（秒）
})
```

参数校验在发出请求前完成：`cid`、`realtime`、`video_duration` 必须大于 0，`aid` 与 `bvid` 任选一个（`bvid` 需为 12 位）；`played_time` 传 `video.VideoPlayedComplete`（-1）表示已看完。上报的清晰度固定为 `video.VideoQuality720P`，调用方无法指定；需要选择清晰度请用 `v.GetVideoStream`（参数见 `video.GetVideoStreamParam.Qn`，常量见 [video/quality.go](video/quality.go)），其响应模型见 [video/model_stream.go](video/model_stream.go)。

### 关注动态流与转发

`GetDynamicFeedAll` 封装 `/x/polymer/web-dynamic/v1/feed/all`（动态首页的关注流）。它的 `data` 与空间动态形态相近，但返回体独立定义为 `dynamic.GetDynamicFeedAllResult`，不与空间动态共享类型。`FeedType` 留空即 `all`，`Features` 与 `WebLocation` 留空时由库填入网页端默认值：

```go
feed, err := d.GetDynamicFeedAll(ctx, dynamic.GetDynamicFeedAllParam{})
for _, item := range feed.Items {
    log.Println(item.IDStr.String(), item.Modules.ModuleAuthor.Name)
}
```

`d.RepostDynamic` 转发一条动态（网页版「发表动态」接口的 `scene=4`）。它需要登录态：`DedeUserID` 与 `bili_jct` 都取自 Cookie，缺失会返回错误；不需要 WBI 签名，CSRF 走 URL 参数：

```go
res, err := d.RepostDynamic(ctx, dynamic.RepostDynamicParam{
    DynamicID: 755402172521250838, // 被转发的原动态 ID
    Content:   "转发抽奖",           // 可为空，空即纯转发
})
log.Println(res.DynIDStr)
```

⚠️ 被转发的原动态 ID 放在**外层的 `web_repost_src.dyn_id_str`**，而不是 `dyn_req.repost_src` —— 网页端构造完 `dyn_req` 之后会把 `repost_src` 删掉。上游文档只有 `scene=2`（图文）的示例，照它拼 `repost_src` 服务端不认。

转发是**不可逆的对外动作**：库不判重、不限制频率、也不自动重试，这些都留给调用方决定。

### 话题动态列表

`d.GetTopicFeed(ctx, param)` 封装 `/x/polymer/web-dynamic/v1/feed/topic`，返回一页响应的 `data`，复用统一 Cookie、context、参数编码和错误处理。参数全部位于 query，不额外启用 WBI 签名或 CSRF，也不自动重试或翻页；`WebLocation` 留空时由库填入默认值。

```go
result, err := d.GetTopicFeed(ctx, dynamic.GetTopicFeedParam{
    TopicID:  topicID,
    SortBy:   dynamic.TopicFeedSortByLatest,
    PageSize: 20,
    Features: "itemOpusStyle,listOnlyfans,opusBigCover,onlyfansVote,decorationCard",
})
for _, item := range result.TopicCardList.Items {
    if item.DynamicCardItem.Type == "DYNAMIC_TYPE_AV" {
        log.Println(item.DynamicCardItem.Modules.ModuleDynamic.Major.Archive.Bvid)
    }
}
```

下一页将返回的 `TopicCardList.Offset` 传入参数，根据 `HasMore` 判断是否结束；检查取消、空游标和重复游标。

### 直播勋章与点赞

`l.GetLiveMedalWall` 查询指定用户的勋章墙，`l.GetLiveFansMedalPanel` 查询当前账号的一页勋章面板，`l.GetLiveActivatedMedalInfo` 查询已激活勋章及任务信息。返回值直接对应 `data`，分页由调用方控制。

```go
l := live.New(root)

panel, err := l.GetLiveFansMedalPanel(ctx, live.GetLiveFansMedalPanelParam{
    Page: 1, PageSize: 10,
})
log.Printf("本页普通勋章数: %d，总页数: %d", len(panel.List), panel.PageInfo.TotalPage)
```

勋章面板按 `PageInfo.TotalPage` 控制分页，不单独依赖 `HasMore`。`l.ReportLiveLike` 上报直播点赞；账号、房间和数量由调用方指定，CSRF 由库填入。方法、参数见 [live/live.go](live/live.go)，响应模型见 [live/live_model.go](live/live_model.go)。

### 活动抽奖与动态抽奖信息

活动主题与活动列表在 `activity` 子包，抽奖相关接口在 `lottery` 子包：

```go
a := activity.New(root)
subject, err := a.GetActivitySubjectInfo(ctx, activity.GetActivitySubjectInfoParam{Sid: 123})
log.Println(subject.ID)

lot := lottery.New(root)
result, err := lot.GetActivityLotteryTimes(ctx, lottery.GetActivityLotteryTimesParam{Sid: "123"})
log.Printf("剩余抽奖次数: %d", result.Times)
```

`lot.GetActivityLotteryTimes` 获取当前账号的活动剩余抽奖次数；`lot.DoActivityLottery` 执行活动抽奖；`lot.GetDynamicLotteryInfo` 查询用户抽奖动态的抽奖信息。参数与响应模型均位于 [activity/activity.go](activity/activity.go) 与 [lottery/lottery.go](lottery/lottery.go)，查询结果直接对应 `data`；仅 `GetDynamicLotteryInfo` 需要 `x-bili-device-req-json`，由库按 `WebLocation` 生成，无需调用方传入。

`lot.DoActivityLottery` 返回 `error`，不解析中奖明细；需要明细可使用 `root.Do`。活动循环和动态删除决策由调用方实现。

### 直播间列表与天选时刻

`l.GetLiveWebRoomList` 获取直播间列表（网页端新接口，WBI 签名）。旧的 second/getList 接口已被风控拦截（-352），请勿再使用。天选时刻的房间集中在互动玩法下的天选分区：

```go
list, err := l.GetLiveWebRoomList(ctx, live.GetLiveWebRoomListParam{
    ParentAreaID: 15, AreaID: 1457, Page: 1, // 1457 为天选聚集分区
})
for _, module := range list.RoomList {
    for _, room := range module.List {
        check, _ := l.CheckLiveAnchorLottery(ctx, live.CheckLiveAnchorLotteryParam{RoomID: room.RoomID})
        if check.Status == 1 && check.GiftPrice == 0 {
            // RequireType 为 0/1 时可免消耗参与：
            // l.JoinLiveAnchorLottery(ctx, live.JoinLiveAnchorLotteryParam{...})
        }
    }
}
```

`l.GetLiveHotRankList` 获取首页人气榜，条目中的 `LotStatus`/`RedPocketStatus` 可低成本初筛有抽奖或红包活动的房间。

`l.SendLiveDanmaku` 向直播间发送一条文本弹幕，CSRF 自动从 Cookie 的 `bili_jct` 填入。服务端对弹幕有频率与长度限制（普通用户单条 20 字符），库不代为限流，需要批量发送时由调用方控制节奏：

```go
_, err := l.SendLiveDanmaku(ctx, live.SendLiveDanmakuParam{
    RoomID: 23058, Msg: "打卡",
})
```

### Web 端观看时长上报

当前网页播放器已把观看时长上报迁移至 `data.bilivideo.com` 域的新端点，对应 `live` 子包的 `l.ReportWebWatchEnter`（进房）与 `l.ReportWebWatchHeartBeat`（心跳），实现见 [live/trace.go](live/trace.go)。流程：进房响应下发 `STKY`（心跳密钥，每跳轮换）、`SID`（会话 id）与 `HBIL`（间隔秒数）；之后按间隔循环发送心跳，每跳使用上一响应轮换出的 `STKY`。

心跳表单中的 `Csn` 字段由官方 skynet wasm 对其余字段的 JSON 计算签名，本库不内置该算法，也不打包 wasm 文件：调用方需自行下载 wasm（地址随页面版本变化，可从直播间页面资源中解析）并用 wasm 运行时（如 wazero）调用其 `skynet` 导出函数，输入为表单 JSON 字符串，输出为签名串。同一账号可同时向多个直播间发送心跳，每个直播间的会话状态（STKY/SID/QID 序号）互相独立。

### 工具方法

网络方法需传入 context；AV/BV 转换为纯计算。构造入口：

```go
m := miscellaneous.New(root)
```

| 子包 | 方法 | 用途 |
| --- | --- | --- |
| `miscellaneous` | `m.UnwrapShortURL(ctx, shortURL)` | 解析短链接，返回目标类型和标识 |
| `miscellaneous` | `m.Now(ctx)` | 获取服务器时间 |
| `miscellaneous` | `m.GetZoneLocation(ctx)` | 查询 IP 所属地理位置 |
| `miscellaneous` | `m.GetRegionDailyCount(ctx)` | 获取分区当日投稿数 |
| `miscellaneous` | `miscellaneous.Av2Bv(aid)` / `miscellaneous.Bv2Av(bvid)` | 纯计算转换，不需要 context |

其它接口按业务位于 `video/`、`user/`、`live/`、`dynamic/` 等子包，每个子包内的响应模型与接口定义放在同一目录（如 `video/model_info.go`、`live/live_model.go`），可通过方法和参数注释查阅。内置请求是否省略参数由 `request` 标签决定，不能只根据 JSON 标签判断。

## 自定义请求

尚未封装的接口可以用 `root.Do`，复用客户端的 Cookie、网络配置、签名与解码流程：

```go
var result struct {
    Mid  int64  `json:"mid"`
    Name string `json:"name"`
}
err := root.Do(ctx, bilibili.Request{
    Method: http.MethodGet,
    URL:    "https://api.bilibili.com/x/space/wbi/acc/info",
    Query:  url.Values{"mid": {mid}},
    WBI:    true,
}, &result)
```

示例需导入 `github.com/rongyuio/bilibili`、`net/http`、`net/url`；`ctx` 与 `root` 由调用方提供。`out` 接收 `data`，传 `nil` 只检查状态及业务错误；WBI 只签 query，CSRF 由调用方按接口要求提供；`Form` 与 `JSON` 互斥。参数规则与边界见 [请求与错误处理](docs/request.md)。

## 错误处理

错误类型按 `errors.Is` / `errors.As` 设计：网络故障、取消与超时保留原始错误链，`*HTTPError`、`Error`、`*ParamError`、`*DecodeError` 可依次提取。子包接口返回的也是这些根包类型。

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
root.SetDroppedFieldHandler(func(field bilibili.DroppedField) {
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

- `make check`：一次跑完 `fmtcheck + vet + race + build + lint`，与 CI 口径一致。
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

本项目以 **AGPL-3.0** 授权，详见 [LICENSE](LICENSE)。作为上游 `CuteReimu/bilibili` 的衍生作品，保留原作者署名。

## 声明

1. 本项目基于上游接口文档中描述的接口编写，请尊重上游文档作者的努力。
2. 请勿滥用；因使用本项目提供的接口、文档等造成的后果由使用者自行承担。
3. 本项目以自用为主，可能随时停止维护。
