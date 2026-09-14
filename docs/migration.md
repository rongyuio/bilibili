# 迁移指南

[返回 README](../README.md)

本文代码片段中的 `ctx`、`client` 及业务参数由调用方提供；客户端导入路径为 `github.com/rongyuio/bilibili`。

## 模块路径迁移

原来使用本地模块 `bilibili` 的项目，需要将导入路径统一替换为 `github.com/rongyuio/bilibili`，并更新 `go.mod` 中对应的 `require` 和 `replace`。远程接入请移除旧的本地替换，再执行快速开始中的安装命令；继续本地联调则按 [README 的本地开发说明](../README.md#本地开发)配置。包名和现有 API 签名不因这次模块路径迁移而改变。

## 入口、签名与字段

| 旧用法 | 新用法 |
| --- | --- |
| `client.Wbi`、`FillWbiHandler(...)` | `Client.Do(ctx, Request{WBI: true, ...}, &data)`；签名器由客户端管理 |
| 自定义接口 `SetResult(&response)` | `Do(..., &response.Data)`，并处理返回的业务错误 |
| `VideoStatusNumber.View` 为 `json.Number` | 改为 `NumberOrString`，仍提供 `String()`、`Int64()`、`Float64()` |
| 根据错误字符串或堆栈定位 | `errors.As(err, &decodeError)` 获取结构化位置 |

`NumberOrString` 保留数字、字符串（包括 `"--"`）与 `null`，`Kind()` 返回 `number`、`string` 或 `null`。数值转换会显式返回错误，JSON 再编码保留原始类别；零值表示 `null`。原有直接赋值 `json.Number` 或强制转换为字符串的代码需要迁移到这些方法；需要构造值时可使用 `json.Unmarshal`。

其余字段不批量改型。`json.Number` 支持数字及数字字符串，不支持任意文本。动态数据中的两处 `Following` 已根据实际响应改为 `json.Number`，迁移方式见下文。

## Context 签名

相对旧版，这是破坏性签名变更：所有可能联网的 `Client` 方法统一增加首参 `ctx context.Context`，例如 `client.GetVideoInfo(ctx, param)`、`client.GetMyUserSpaceDetail(ctx)`。本文调用示例中的 `ctx` 均由调用方提供。`Do` 和 `NewAnonymousClient` 已有的 context 签名不变；Cookie 读写、配置和纯计算方法不变。

独立使用 WBI 时，改为 `wbi.GetKeys(ctx)`、`wbi.GetMixinKey(ctx)`、`wbi.SignQuery(ctx, query, ts)`、`wbi.SignMap(ctx, payload, ts)`，因为签名可能触发密钥刷新。没有新增 `XxxContext` 或无 context 的兼容包装；调用方接口声明、方法表达式和回调类型也需同步修改。

## 会话与参数编码

`NewAnonymousClient()` 迁移为 `NewAnonymousClient(ctx) (*Client, error)`；`NewWithClient` 接管 Resty 并关闭 Jar，已有 Jar 会话需要构造前显式导出。Cookie 快照和配置约束见[认证与会话](authentication.md)，不要继续通过底层 Cookies 或默认 Cookie 请求头管理会话。

非 nil 的非结构体参数指针现在返回 `ParamError`，不再触发 panic；标签值中的等号完整保留。参数转换失败会返回错误而不是静默变为空字符串。JSON/multipart 混用或字段位置冲突会报错；依赖 Resty 内部 Body 为普通 map 的代码需适配已编码 JSON 字节或 multipart 构造方式。

## 错误包装

普通业务失败在原有 `Error{Code, Message}` 外补充 HTTP 方法和安全接口地址。WBI 非零业务码且缺少密钥时也采用此包装；非零码但存在密钥时仍沿用原有流程并继续检查密钥合法性。`Error` 的字段和自身错误文本不变，但请求返回错误的最外层类型及整体文本发生变化。将 `err.(bilibili.Error)` 或错误字符串匹配迁移为示例中的 `var be bilibili.Error; errors.As(err, &be)`，不要改为指针类型目标。

## 动态命名类型

`DynamicItem`、`DynamicInfo` 和主要模块定义放在 `dynamic_model.go`，接口方法与请求参数放在 `dynamic.go`；话题接口 `GetTopicFeed`、参数放在 `topic.go`，模型放在 `topic_model.go`。仅就模型提取而言，网络方法签名、返回根类型以及 `item.Modules.ModuleAuthor.Name` 等字段访问路径不变。

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

两套模型的差异完整保留：外层 `Major` 是指针，原动态 `Major` 是值；外层 `Basic.LikeIcon.Id` 为 `json.Number`，原动态对应字段为 `int`。原动态头像、作者和富文本也有不同字段，不能直接复用外层模块。除头像 `fallback_layers` 渲染树外，更深层的匿名结构已在[第二轮重构](#第二轮重构方法模型拆分与去重)中提取为具名类型。

模型提取本身没有修改 JSON 标签、字段顺序、叶子类型或指针／切片结构，没有新增自定义反序列化；`DecodeError` 仍按既有规则遍历模型并报告 JSON 路径和 Go 字段。随后单独修复了两处 `Following` 类型，见下文。

## Following 数值状态修复

`DynamicModuleAuthor.Following` 和 `DynamicOriginalModuleAuthor.Following` 从 `bool` 改为 `json.Number`。该字段返回数值状态，暂不定义未经确认的状态常量，也不将非零值解释为“已关注”。

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

原动态对应字段为 `item.Orig.Modules.ModuleAuthor.Following`，读取方式相同。头像尺寸使用 `float64`，支持小数。

## 第二轮重构：方法/模型拆分与去重

本轮把响应模型从接口文件中拆出，并为原本内联的匿名结构体命名，同时合并了字段完全一致的重复类型。网络方法签名、请求参数位置、JSON 标签、字段顺序与指针／切片层级都没有变化；下列变更只影响嵌套字段的 **Go 类型身份**。

### 文件组织

接口方法与 `*Param` 保留在各业务文件，响应模型移动到同名 `*_model.go`（同包，导入方式不变）：`article_model.go`、`comment_model.go`、`fav_model.go`、`history_model.go`、`login_model.go`、`message_model.go`、`topic_model.go`、`vip_model.go`，以及既有的 `dynamic_model.go`、`live_model.go`、`user_model.go`、`video_model.go`。

### 合并为别名的类型

| 旧类型 | 现状 |
| --- | --- |
| `SpaceVip` | `type SpaceVip = CardVip`，字段完全一致 |
| `LiveMedalWallItemUinfoMedal`、`LiveFansMedalPanelItemUinfoMedal` | `type ... = LiveUinfoMedal`，字段完全一致 |

别名赋值与字段访问无需修改；依赖类型名称的反射代码需要改为新类型名。

### 新增的具名类型

原本内联的匿名结构体现在有了类型名（节选）：

| 位置 | 命名类型 |
| --- | --- |
| `DynamicItemBasic.LikeIcon` / `DynamicOriginalBasic.LikeIcon` | `DynamicLikeIcon` / `DynamicOriginalLikeIcon` |
| `DynamicModuleAuthor.Vip`、原动态同名字段 | `DynamicAuthorVip`（外层与原动态字段一致，合并） |
| `DynamicModuleAuthor.Pendant` / 原动态 `Pendant` | `DynamicPendant` / `DynamicOriginalPendant`（`pid` 类型不同，保留两个） |
| 外层与原动态头像 `container_size` | `AvatarContainerSize` |
| 外层与原动态描述 `rich_text_nodes` | `DynamicRichTextNode` / `DynamicOriginalRichTextNode` |
| 外层与原动态 `emoji` | `DynamicEmoji`（字段一致，共用） |
| `DynamicModuleMore.ThreePointItems` | `DynamicThreePointItem` |
| `DynamicModuleStat` 的评论/转发与点赞 | `DynamicStat` / `DynamicLikeStat` |
| `DynamicArchive.Badge` / `Stat` | `DynamicArchiveBadge` / `DynamicArchiveStat` |
| `DynamicDraw.Items` | `DynamicDrawItem` |
| `DynamicRepostDetail.Items` 及其 `desc`、`origin`、`previous`、`display` | `DynamicRepostItem` / `DynamicRepostDesc` / `DynamicRepostOrigin` / `DynamicRepostPrevious` / `DynamicRepostDisplay` 等 |
| `DynamicLikeList.ItemLikes[].UserInfo` | `DynamicLikeUserInfo`，其中 `Vip`/`Pendant`/`LevelInfo` 复用 `DynamicUserVip`/`DynamicUserPendant`/`DynamicUserLevelInfo` |
| `DynamicUpList.Items[].UserProfile` | `DynamicUpUserProfile`、`DynamicUpUserInfo`、`DynamicUpVip`、`DynamicUpVipLabel`、`DynamicUpPendant`、`DynamicUpLevelInfo` |
| `DynamicPortal.MyInfo` / `UpList` | `DynamicPortalMyInfo`（`LevelInfo` 为 `DynamicPortalLevelInfo`） / `DynamicPortalUp` |
| `TopicModuleAuthor.OfficialVerify` | 复用共享的 `OfficialVerify` |
| `TopicModuleAuthor.Pendant` / `Vip` | `TopicPendant` / `TopicVip`（`Label` 复用共享的 `Label`） |
| `TopicDynamicBasic.LikeIcon` | `TopicLikeIcon` |
| `TopicArchive.Badge` / `Stat` | `TopicArchiveBadge` / `TopicArchiveStat` |
| `TopicModuleStat` | `TopicStat` / `TopicLikeStat` |
| `TopicSortByConf.AllSortBy` | `[]TopicSortByItem` |
| `AllFavourFolderInfo.List` | `[]AllFavourFolderItem` |
| `FavourInfo` / `FavourList.Medias` 的 `upper`、`cnt_info`、`ugc` | `FavourUpper` / `FavourResourceCntInfo` / `FavourUgc` |
| `FavourList.Info` | `FavourFolderDetail`，其 `upper`/`cnt_info` 复用 `Upper`/`CntInfo` |
| `SelfFavourList.MediaListResponse` | `SelfFavourMediaListResponse`，`list` 为 `[]SelfFavourMediaItem` |
| `Elec.ShowInfo` | `ElecShowInfo` |
| `UserSpaceDetail.UserHonourInfo` / `Series` | `UserHonourInfo` / `UserSeries` |
| `PrivateMessageList.SessionList[].LastMsg` / `AccountInfo` | `PrivateMessageSession` / `PrivateMessageLastMsg` / `PrivateMessageAccountInfo` |

普通字段读取无需迁移；手写匿名结构赋值、嵌套复合字面量、函数或接口声明以及依赖类型名称的反射代码需要检查。

### 刻意保留的差异类型

以下类型字段集或字段命名不同，继续保留为独立类型，不做合并：`VipUserVip`、`MyVip`、`UserCardVip`、`StaffVip`（会员变体）；`VipLabel` 与 `Label`（会员铭牌）；`OfficialVerify` 与 `Official`；`Owner`、`Author` 与各业务作者类型；`DynamicModuleStat` 与 `TopicModuleStat`；`DynamicArchive` 与 `TopicArchive`；各业务分页类型（`CommentsPage`、`UserVideoPage`、`CollectionPage`、`ZoneVideoPage`、`LiveFansMedalPanelPageInfo`）。

### 保留的匿名结构

动态头像 `fallback_layers` 是随接口版本漂移的一次性渲染配置树，层级极深且无复用价值，继续保留为内联匿名结构体。

### 请求框架收敛

本轮复核了仍然直接使用 `newRequest`/`sendRaw` 的接口，确认它们都有正当理由，保持现状：短链解析 `UnwrapShortUrl`（只读 302 `Location`）、`GetWebCookieRefreshCsrf`（返回 HTML）、`RefreshCookie`（CSRF 允许留空并回退到 Cookie 快照）、`StartLive`（签名依赖 Cookie 快照）、`UploadDynamicBfs`（multipart）、`NewAnonymousClient`（初始化匿名 Cookie）。其余接口的参数编码、CSRF 与错误处理都继续走 `execute`/`encodeParams`/`fillCsrf`。

### 文档

仓库根目录的 `AGENTS.md` 已删除，贡献规则以 [CONTRIBUTING.md](../.github/CONTRIBUTING.md) 为准。
