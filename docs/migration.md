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

## 第三轮重构：缺陷修复、类型合并与命名统一

本轮修复了若干真实缺陷，合并字段完全一致的重复类型，补齐剩余模型文件，并统一命名与文件职责。请求参数位置、JSON 标签、字段顺序与指针/切片层级除下列明确列出的项以外均未变化。

### 缺陷修复

| 项 | 说明 |
| --- | --- |
| `HistoryList.Uri` | 该字段此前被误写入 `Covers` 的注释中，实际从未解析；现已补回 `Uri string \`json:"uri"\``（剧集/直播的重定向 url）。属于**新增可解析字段**，不影响原有字段。 |
| `IntergratedSearch` | 方法名拼写修正为 `IntegratedSearch`。**破坏性变更**，请同步修改调用处；无兼容包装。 |
| `SeachRespResult` | 类型名拼写修正为 `SearchRespResult`。**破坏性变更**。 |
| `SearchRespTopTList` | 类型名拼写修正为 `SearchRespTopList`。**破坏性变更**。`SearchRespData.TopTList` 字段名与 `top_tlist` 标签不变。 |

### 合并为别名的类型

字段名、JSON 标签、类型三者完全一致的类型已合并为一处，旧名以别名承接，字段访问与复合字面量无需修改：

| 旧类型 | 现状 |
| --- | --- |
| `GetUserFollowersResult`、`GetUserFollowingsResult`、`SearchUserFollowingsResult`、`GetSameFollowingsResult` | `type ... = RelationUserPage`（`list`/`re_version`/`total`） |
| `GetWhispersResult`、`GetFriendsResult`、`GetBlacksResult` | `type ... = RelationUserList`（`list`/`re_version`） |
| `FavourUpper` | `type FavourUpper = Owner`（`mid`/`name`/`face`） |
| `RecommendPendant`、`RecommendCard` | `type ... = RecommendItem`（`id`/`name`/`image`/`jump_url`） |

`video_model.go` 文件头与 `type.go` 中「Owner 不与其他作者类型合并」的旧注释已同步更新。依赖类型名称的反射代码需要改为新类型名。

### 新增具名类型

`GetTopicFeedResult.RelatedTopics` 由空匿名结构改为具名类型 `TopicRelatedTopics`（仍为空结构体，维持“忽略未知字段”的解码行为）。普通字段读取无需迁移。

### `AudioOrVideo` 字段命名规范化

该接口对同一份数据同时返回 camelCase 与 snake_case 两种键，因此每种键各保留一个字段：规范化命名对应 camelCase 键，以 `Snake` 结尾的字段对应 snake_case 键。**JSON 标签全部原样保留**。

| 旧字段名 | 新字段名 | JSON 标签 |
| --- | --- | --- |
| `Baseurl` | `BaseURL` | `baseUrl` |
| `BaseUrl` | `BaseURLSnake` | `base_url` |
| `Backupurl` | `BackupURL` | `backupUrl` |
| `BackupUrl` | `BackupURLSnake` | `backup_url` |
| `Mimetype` | `MimeType` | `mimeType` |
| `MimeType` | `MimeTypeSnake` | `mime_type` |
| `Framerate` | `FrameRate` | `frameRate` |
| `FrameRate` | `FrameRateSnake` | `frame_rate` |
| `Startwithsap` | `StartWithSap` | `startWithSap` |
| `StartWithSap` | `StartWithSapSnake` | `start_with_sap` |
| `Segmentbase` | `SegmentBase` | `SegmentBase` |
| `SegmentBase` | `SegmentBaseSnake` | `segment_base` |

这是**破坏性字段改名**，读取 DASH 流的调用方需要按上表调整字段名；未删除任何字段，两种响应形态都仍然被解析。

### 文件组织

- 类型归位：`Notice` 从 `live.go` 上移到共享的 `type.go`（它同时被 `CommentsDetail` 与 `StartLiveResult` 引用）。
- 新增模型文件：`emote_model.go`、`search_model.go`、`video_ranking_model.go`、`lottery_model.go`；直播响应模型移入既有的 `live_model.go`。
- 职责拆分：`wbi.go` 拆出 `wbi_storage.go`（`Storage` 接口与 `MemoryStorage`）；`decode_diagnostic.go` 拆出 `decode_fields.go`（字段匹配与诊断类型名）。
- `cookie.go` 的 `type ( ... )` 分组声明改为逐个声明，与其它文件保持一致。

同包内的文件搬移不影响导入方式；`wbi.go` 与 `decode_diagnostic.go` 拆分后对外签名、错误语义与 `DecodeError` 可解包性均未变化。

## 第四轮重构：初始缩写规范化、清理与模型文件细分

### Go 初始缩写命名规范化（破坏性）

按 Go 惯例把初始缩写统一为大写，规则为：`Id`→`ID`、`Ids`→`IDs`、`Uid`→`UID`、`Uids`→`UIDs`、`Url`→`URL`、`Urls`→`URLs`、`Uri`→`URI`、`Uuid`→`UUID`、`Json`→`JSON`。规则只作用于标识符，**JSON 标签一个都没有改动**（改名前后的 2857 个标签集合完全一致），请求与响应数据不受影响。

共涉及 107 个不同标识符、29 个源码文件。常见对照：

| 旧名 | 新名 |
| --- | --- |
| `Id`、`Ids` | `ID`、`IDs` |
| `DynamicId`、`DynamicIdStr` | `DynamicID`、`DynamicIDStr` |
| `RoomId`、`TargetId`、`AreaId` | `RoomID`、`TargetID`、`AreaID` |
| `Uid`、`Uids`、`SenderUid`、`UidType` | `UID`、`UIDs`、`SenderUID`、`UIDType` |
| `MediaIdParam`、`FavourId`、`GetFavourIds` | `MediaIDParam`、`FavourID`、`GetFavourIDs` |
| `Url`、`Urls`、`JumpUrl`、`AvatarSubscriptUrl` | `URL`、`URLs`、`JumpURL`、`AvatarSubscriptURL` |
| `UnwrapShortUrl` | `UnwrapShortURL` |
| `Uri`、`BackdropUri`、`HistoryList.Uri` | `URI`、`BackdropURI`、`HistoryList.URI` |
| `Json`、`ExtendJson` | `JSON`、`ExtendJSON` |
| `deviceId`（包内变量） | `deviceID` |

迁移方式：调用方按同一规则批量改名即可，无兼容别名（字段与方法名无法用别名承接）。上一轮新增的 `HistoryList.Uri` 在本轮已改名为 `HistoryList.URI`。

**刻意未纳入**的缩写字：`Mid`、`Aid`、`Cid`、`Tid`、`Rid`、`Fid`、`Pid`、`Oid`、`Bvid`、`Vmid`、`Vip`、`Nft`、`Md5`。它们不在 Go 官方初始缩写表内（属 B 站专有缩写或商品名），改动收益低而破坏面大。若后续要一并调整，应作为单独的大版本变更处理。

被忽略的本地工具 `test/` 也已同步改名，以保持 `go build ./...` 通过。

### 死代码清理（破坏性）

| 删除项 | 说明 |
| --- | --- |
| `ResultData` | `search_model.go` 中的空结构体，全仓零引用 |
| `VideoSubtitles` | `video_model.go` 中全仓零引用的类型；字幕列表请使用 `[]VideoSubtitle` |

### 机械一致性

- `live_model.go`、`topic_model.go` 中的 10 处 `interface{}` 统一为 `any`。
- `emote_model.go` 的 `User_panel_packages`、`All_packages` 改为 `UserPanelPackages`、`AllPackages`（JSON 标签 `user_panel_packages`、`all_packages` 保持不变）。
- 清理 `fav_model.go` 封面字段注释中的游离制表符、`search.go` 参数注释中的制表符，并把 `article.go` 的两条 `import` 语句合并为一个 import 块。

### 文件组织

`dynamic_model.go` 原先 775 行、79 个类型，现按子领域拆为 5 个文件（同包，导入方式不变）：

| 文件 | 内容 |
| --- | --- |
| `dynamic_model.go` | 整体动态（`DynamicItem` 及其各 `Module*`）、`AvatarContainerSize` |
| `dynamic_original_model.go` | 转发动态内嵌的原动态（`DynamicOriginal*`、`DynamicDecorate*`） |
| `dynamic_repost_model.go` | 转发列表与转发用户、`@` 搜索（`DynamicRepost*`、`DynamicUserVip` 等） |
| `dynamic_interact_model.go` | 点赞列表、直播中关注者、更新 UP 主（`DynamicLike*`、`DynamicLiveUser*`、`DynamicUp*`） |
| `dynamic_portal_model.go` | 动态卡片、门户与发布接口（`DynamicCard`、`DynamicPortal*`、`CreateDynamicResult` 等） |

### 版本影响

导出标识符与字段改名、以及删除导出类型，按语义化版本约定都属于**破坏性变更**。本模块目前只有 `v0.1.0`，处于 **v0 阶段**：v0 允许在次版本号内引入破坏性变更，因此上述改动随 **v0.2.0** 发布即可，**不需要**升主版本，也**不需要**改模块路径。只有发布 `v1.0.0` 之后再出现破坏性变更，才需要升主版本并把模块路径改为 `github.com/rongyuio/bilibili/v2` 这类形式。
