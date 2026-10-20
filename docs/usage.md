# 使用示例

[返回 README](../README.md)

本文代码片段中的 `ctx`、`root`、`mid`、`bvid` 等由调用方提供。`root` 是 `bilibili.New()` 创建的根客户端，子包入口按需构造，例如 `v := video.New(root)`、`d := dynamic.New(root)`、`l := live.New(root)`。示例省略错误判断的写法见 README 的「快速开始」。

## 登录与 Cookie

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

`LoginWithQRCode` 会轮询到扫码成功或二维码失效，`ctx` 应预留人工确认所需时间。密码与短信登录、Cookie 的保存与恢复、Resty 接管见 [认证与会话](authentication.md)。

## 常用接口

以下示例聚焦调用方式，省略了错误判断；完整写法见 README 的「创建客户端并调用接口」。

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

参数校验在发出请求前完成：`cid`、`realtime`、`video_duration` 必须大于 0，`aid` 与 `bvid` 任选一个（`bvid` 需为 12 位）；`played_time` 传 `video.VideoPlayedComplete`（-1）表示已看完。上报的清晰度固定为 `video.VideoQuality720P`，调用方无法指定；需要选择清晰度请用 `v.GetVideoStream`（参数见 `video.GetVideoStreamParam.Qn`，常量见 [video/quality.go](../video/quality.go)），其响应模型见 [video/model_stream.go](../video/model_stream.go)。

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

勋章面板按 `PageInfo.TotalPage` 控制分页，不单独依赖 `HasMore`。`l.ReportLiveLike` 上报直播点赞；账号、房间和数量由调用方指定，CSRF 由库填入。方法、参数见 [live/live.go](../live/live.go)，响应模型见 [live/live_model.go](../live/live_model.go)。

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

`lot.GetActivityLotteryTimes` 获取当前账号的活动剩余抽奖次数；`lot.DoActivityLottery` 执行活动抽奖；`lot.GetDynamicLotteryInfo` 查询用户抽奖动态的抽奖信息。参数与响应模型均位于 [activity/activity.go](../activity/activity.go) 与 [lottery/lottery.go](../lottery/lottery.go)，查询结果直接对应 `data`；仅 `GetDynamicLotteryInfo` 需要 `x-bili-device-req-json`，由库按 `WebLocation` 生成，无需调用方传入。

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

当前网页播放器已把观看时长上报迁移至 `data.bilivideo.com` 域的新端点，对应 `live` 子包的 `l.ReportWebWatchEnter`（进房）与 `l.ReportWebWatchHeartBeat`（心跳），实现见 [live/trace.go](../live/trace.go)。流程：进房响应下发 `STKY`（心跳密钥，每跳轮换）、`SID`（会话 id）与 `HBIL`（间隔秒数）；之后按间隔循环发送心跳，每跳使用上一响应轮换出的 `STKY`。

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

示例需导入 `github.com/rongyuio/bilibili`、`net/http`、`net/url`；`ctx` 与 `root` 由调用方提供。`out` 接收 `data`，传 `nil` 只检查状态及业务错误；WBI 只签 query，CSRF 由调用方按接口要求提供；`Form` 与 `JSON` 互斥。参数规则与边界见 [请求与错误处理](request.md)。

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

`ParamError` 提供参数类型、字段与位置；`DecodeError` 提供 JSON 路径、Go 字段及预期类型。完整示例见 [请求与错误处理](request.md)；记录错误时不要输出 Cookie、凭证或完整响应。

单个字段的 JSON 类型与模型不符时，该字段被丢弃（保持零值）并上报，其余字段照常解码；`DecodeError` 只在这类容错也失败时返回。需要感知漂移时注册回调：

```go
root.SetDroppedFieldHandler(func(field bilibili.DroppedField) {
    log.Printf("接口=%s JSON路径=%s 预期=%s 实际=%s",
        field.Endpoint, field.JSONPath, field.Expected, field.Actual)
})
```

回调在同一条响应内可能被多次调用，且可能在并发请求中执行。规则边界见 [请求与错误处理](request.md#单个字段类型不符时的容错)。
