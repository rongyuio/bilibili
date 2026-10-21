# 哔哩哔哩 API Go 客户端

[![Go](https://github.com/rongyuio/bilibili/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/rongyuio/bilibili/actions/workflows/golangci-lint.yml)
[![GoFmt](https://github.com/rongyuio/bilibili/actions/workflows/gofmt.yml/badge.svg)](https://github.com/rongyuio/bilibili/actions/workflows/gofmt.yml)
[![Latest tag](https://img.shields.io/github/v/tag/rongyuio/bilibili?label=latest&sort=semver)](https://github.com/rongyuio/bilibili/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/rongyuio/bilibili.svg)](https://pkg.go.dev/github.com/rongyuio/bilibili)
[![License: AGPL-3.0](https://img.shields.io/badge/License-AGPL--3.0-blue.svg)](LICENSE)

基于 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 继续维护的 Go 客户端，封装 Bilibili API，提供 Cookie 管理、WBI 签名、context 取消和结构化错误定位。

模块路径为 `github.com/rongyuio/bilibili`，要求 Go 1.27.3 或更新版本，依赖以 [go.mod](go.mod) 为准。接口可能随服务端变化。

- [维护状态](#维护状态)
- [特性](#特性)
- [包结构](#包结构)
- [快速开始](#快速开始)
- [详细文档](#详细文档)
- [开发与贡献](#开发与贡献)
- [License](#license)
- [声明](#声明)

## 维护状态

- 本仓库是 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 的 fork。上游已被作者删除（该地址现返回 404）；删除前已归档：README 标注 Deprecated、默认分支改名为 `deprecated`、issues/discussions/wiki 全部关闭，不再接受改动与合并。其所依据的接口文档仓库也已停止维护。
- 本 fork **继续维护**：跟进新增接口，已实现接口在发现问题时修复。以自用为主，不承诺固定的支持范围与响应时间。
- B 站服务端可能随时调整接口与字段。生产环境请对关键接口自行验证，并处理业务错误与字段漂移。
- 本仓库保留 fork 关联；`master` 已开启分支保护（改动须经 PR 合入，禁止强制推送与删除）。
- 版本处于 **v0 阶段**，允许在次版本号内引入破坏性变更。升级前请阅读[迁移指南](docs/migration.md)。

## 特性

- **Cookie 管理**：浏览器 Cookie 导入、游客初始化、扫码 / 密码 / 短信登录，以及 Cookie 的保存与恢复
- **WBI 签名**：按需对 query 参数签名，密钥自动缓存并按周期刷新
- **统一请求链路**：全部内置接口共用同一套 Cookie 快照、参数编码、context 取消与错误处理
- **结构化错误诊断**：`HTTPError` / `Error` / `ParamError` / `DecodeError` 均支持 `errors.As`，解码失败时给出 JSON 路径与偏移
- **容错解码**：单个字段的 JSON 类型与模型不符时只丢弃该字段并经 `SetDroppedFieldHandler` 上报，不连累整条响应。这是兼容服务端局部漂移的降级解码，不是无损解码；被丢弃的字段保持零值。
- **逃生通道**：尚未封装的接口可通过 `Client.Do` 复用同一套能力

## 包结构

根包 `bilibili` 提供基础设施：`Client` 的构造与会话、统一请求链路、参数编码、错误类型、解码诊断、Cookie 管理、WBI 签名，以及供子包复用的 `Execute` / `ExecuteRaw` / `FillCsrf` 等导出设施。根包的 `type.go` 只保留真正跨业务域的基础类型。

业务接口按域拆分到子包。每个子包提供 `New(*bilibili.Client) *Client`，返回嵌入了根包 `Client` 的调用入口，因此同一份 Cookie、WBI 缓存与错误处理在所有子包之间共享：

- `activity` —— 活动主题与活动列表
- `appwidget` —— 客户端版本下载地址
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

跨域调用可以使用聚合包 `bili`：`bili.New()` 创建根客户端，并把全部业务域挂在同一个对象上，`client.Video`、`client.User` 等字段直接可用。聚合包只做包装，不持有独立会话；字段由 `go generate ./...` 维护，新增子包后需重新生成。只需要单个域时，仍可直接用根包配合子包构造。

## 快速开始

### 安装

在已有 Go 项目的目录中执行：

```bash
go get github.com/rongyuio/bilibili@latest
```

v0 阶段 API 仍可能调整，升级前请阅读[迁移指南](docs/migration.md)与[版本策略](docs/versioning.md)。需要开发分支代码时可使用 `@master`，Go 会记录对应提交的伪版本。

Go 1.27.3 是当前最低工具链门槛：代码本身只需要 Go 1.26 的 `errors.AsType`，提高到 1.27.3 是为了带上 `crypto/x509` 对 CVE-2026-111943 的修复，与接口稳定性无关。

| 变更 | 版本 |
| --- | --- |
| 修复 bug，不改导出 API | patch |
| 新增接口或能力 | minor |
| v0 阶段的破坏性变更 | minor，需阅读迁移指南 |
| 保证兼容 | v1 之后 |

### 创建客户端并调用接口

下面是一个完整程序，按需在自己的项目运行：

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/rongyuio/bilibili/bili"
    "github.com/rongyuio/bilibili/video"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

    client := bili.New()
    info, err := client.Video.GetVideoInfo(ctx, video.GetVideoInfoParam{Bvid: "BV1L9Uoa9EUx"})
    if err != nil {
        log.Printf("获取视频失败: %v", err)
        return
    }
    log.Println(info.Title)
}
```

`bili.New()` 创建根客户端并构造聚合入口，所有业务域共享同一个会话。只需要单个域时，也可以直接用 `bilibili.New()` 配合 `video.New(root)`。所有网络方法以 `ctx context.Context` 为首参，调用后先处理错误再使用结果。普通请求可并发，登录、会话切换及配置修改需串行；不要复制已使用的 `Client`。

### 按任务选择入口

聚合入口把全部业务域挂在同一个对象上。常用域：

| 任务 | 入口 |
| --- | --- |
| 视频信息、播放地址、观看上报 | `client.Video` |
| 用户信息、关系、空间 | `client.User` |
| 动态流、发布、转发 | `client.Dynamic` |
| 直播房间、勋章、弹幕 | `client.Live` |
| 登录 | `client.Login` |
| 收藏与历史 | `client.Fav` / `client.History` |
| 搜索 | `client.Search` |
| 尚未封装的接口 | `client.Do` |

只需要单个域时，直接用根包与子包，避免把全部业务包拉进依赖图：

```go
root := bilibili.New()
v := video.New(root)
```

### 客户端构造选择

| 场景 | 用法 |
| --- | --- |
| 默认，跨域调用 | `bili.New()` |
| 已有根客户端 | `bili.NewWith(root)` |
| 需要游客 Cookie | `bili.NewAnonymous(ctx)` |
| 只用单个域 | `root := bilibili.New()` 后接 `video.New(root)` |
| 接管已有 Resty 客户端 | `bilibili.NewWithClient(restyClient)` |

### 登录与 Cookie

已有浏览器 Cookie 时，在发送请求前导入：

```go
client := bili.New()
client.SetRawCookies(cookieHeader) // cookieHeader 由调用方读取，不写入代码或日志。
```

`bili.New()` 不联网；需要游客 Cookie 时用 `bili.NewAnonymous(ctx)`。扫码 / 密码 / 短信登录、Cookie 保存与恢复、Resty 接管见[认证与会话](docs/authentication.md)，完整调用示例见[使用示例](docs/usage.md#登录与-cookie)。

### 并发与会话

- 普通请求可以并发。
- 登录、Cookie 修改、`SetDroppedFieldHandler` 与 Resty 配置变更必须在请求之外串行执行。
- `Client` 首次使用后不可复制；需要独立会话时创建新的客户端。
- 每次请求只取一次 Cookie 快照，响应 Cookie 自动合并回客户端。

## 详细文档

| 文档 | 内容 |
| --- | --- |
| [使用示例](docs/usage.md) | 常用接口、登录与 Cookie、自定义请求、错误处理 |
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

推送合法的语义化版本 tag（形如 `v<major>.<minor>.<patch>`）即触发 [release.yml](.github/workflows/release.yml)，自动校验、测试、构建并创建 GitHub Release。完整流程与 tag 格式见 [CONTRIBUTING.md](.github/CONTRIBUTING.md) 的「关于发版」与[版本策略](docs/versioning.md)。

## License

本项目以 **AGPL-3.0** 授权，详见 [LICENSE](LICENSE)。作为上游 `CuteReimu/bilibili` 的衍生作品，保留原作者署名。

## 声明

1. 本项目基于上游接口文档中描述的接口编写，请尊重上游文档作者的努力。
2. 请勿滥用；因使用本项目提供的接口、文档等造成的后果由使用者自行承担。
3. 本项目以自用为主，可能随时停止维护。
