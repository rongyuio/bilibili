# AGENTS.md

本文件为在此仓库中工作的编码助手提供指引。

## 概述

封装 Bilibili HTTP API 的 Go 客户端库。模块路径为 `github.com/rongyuio/bilibili`（导入路径与目录名不一致），要求 Go 1.26+（`go.mod` 的 `go` 指令是最低支持版本，提升需说明理由，见 `docs/versioning.md`）。本仓库是已归档的 `CuteReimu/bilibili` 的持续维护 fork；上游与接口文档仓库 `SocialSisterYi/bilibili-API-collect` 均已归档。版本处于 v0 阶段，允许在次版本号内引入破坏性变更（见 `docs/migration.md`）。

所有代码都位于仓库根目录的 `bilibili` 包中（扁平包结构，除被忽略的 `test/` 和 `tools/` 生成器外没有子包）。文档与代码注释使用中文；修改既有文档和注释时请保持这一风格。

## 分支与提交

`master` 已开启分支保护（禁止强推与删除，要求走 PR），**不要直接向 `master` 提交或推送**，即使当前账号有权限绕过保护。改动一律从最新 `master` 切分支、在分支上提交、推送后开 PR 合入，分支名用 `<type>/<简短描述>`。若不慎已在本地 `master` 提交且尚未推送，先 `git branch <新分支>` 保住提交，再 `git reset --hard origin/master`，**不要强推**。细节见 `.github/CONTRIBUTING.md` 的「分支与提交流程」。

## 常用命令

```bash
go build ./...                    # 编译
go vet ./...                      # 静态检查
gofmt -s -w <file.go>             # 格式化（CI 要求 gofmt -s -l . 输出为空）
golangci-lint run                 # 使用 .golangci.yml（v2 schema，lint 版本 v2.13.2）
go test ./...                     # 运行全部测试
go test -run TestName ./...       # 运行单个测试
go test -run 'TestCookie|TestSignMap' ./...   # 按正则运行一组测试
go test -race -v ./...            # 输出详细结果（CI 使用的形式）
```

`go run tools/gen_struct.go` 可将粘贴的 Markdown 字段表格转换成 Go 结构体定义。详见 `tools/README.md`。

测试**不会访问真实 API，也不需要凭证**。多数只依赖标准库 `testing`；`util_test.go` 用 `resty.New().R()` 构造请求对象来验证 `withParams` 的编码结果，`dynamic_test.go` 的少数用例用假 transport（`SetTransport`）把整条请求链路走完 —— 覆盖 URL / 方法 / CSRF 位置 / body 形状，那些是 handler 单测够不到的部分（URL 与方法是写死在 `execute` 调用里的常量）。**两者都不发出真实请求。** `test/` 存放本地账号相关脚本，已被 gitignore，且是**独立的 Go 模块**（自带 `go.mod`，用 `replace` 指向本仓库），其依赖与 Go 版本要求都不影响主模块；`.golangci.yml` 的 `exclusions.paths` 同时排除了 `test/` 与 `tools/`。

CI（`.github/workflows/`）在向 `master` 的 push / PR 时运行两个 workflow：`gofmt.yml` 要求 `gofmt -s -l .` 输出为空；`golangci-lint.yml`（workflow 名为 `Go`）依次运行 golangci-lint、`go test -race -v ./...`、`go build -v ./...`。三个 workflow 的 Go 版本均由 `go-version-file: go.mod` 决定，因此改 `go.mod` 的 `go` 指令即可切换验证环境；该指令表示最低支持版本（当前 1.26），提升需在 PR 中说明理由。`master` 已开启分支保护，改动通过 PR 合入。

发版由 `release.yml` 处理：推送 `v*` 标签后，先校验标签名符合语义化版本（`v<major>.<minor>.<patch>`）且该标签指向的提交位于 `master`（任一不满足即拒绝发布），再运行 `go test -race ./...` 与 `go build ./...`，最后执行 `gh release create --generate-notes` 创建 GitHub Release，无需手动操作。

## 文件组织约定

包按 API 业务域划分，与上游 `bilibili-API-collect` 的文档结构保持一致：

- `<domain>.go` —— 方法与 `*Param` / `*Result` 类型（例如 `video.go`、`user.go`、`live.go`、`dynamic.go`、`fav.go`、`login.go`）。
- `<domain>_model.go` —— 该业务域的响应模型结构体（例如 `video_model.go`、`user_model.go`）。
- `type.go` —— 跨多个文件复用的共享模型类型。
- `video_zone.csv` —— 分区查询所用的内置/参考数据。

新增接口时，放入与上游文档分类对应的文件；其响应模型放入同域的 `*_model.go`。

## 请求链路（核心架构）

所有内置接口调用都走同一条链路，主要定义在 `request.go`：

```
Client.<Method>(ctx, param)
  -> execute[Out](ctx, c, method, endpoint, in, handlers...)   // request.go
     -> c.newRequest(ctx)                                      // 只取一次 Cookie 快照
     -> executeRequest[Out]
        -> withParams(r, in)                                   // params.go：结构体标签 -> query/json/multipart
        -> handlers...                                         // fillCsrf / fillWbi / 各接口定制处理
        -> c.send(r, method, endpoint, &out)
           -> c.sendRaw                                         // 执行请求；即使出错也合并响应 Cookie
           -> 非 200 时返回 newHTTPError
           -> decodeResponse                                    // response.go
```

修改这条链路时需保持的关键约束：

- **context 永远作为首参。** 所有网络方法都以 `ctx context.Context` 为第一个参数。`checkContext` 在任何工作开始前拒绝 nil / 已结束的 context。context 也会传入 WBI 密钥刷新。
- **每次请求只取一次 Cookie 快照。** `newRequest` 通过 `GetCookies()` 深拷贝一次快照；CSRF 和各接口签名都读取这份快照。响应 Cookie 在 `sendRaw` 中合并回客户端（HTTP / 业务错误时同样合并）。`Client` 首次使用后不可复制；普通请求可并发，但登录 / 会话切换 / 配置变更必须在请求之外串行进行。
- **解码到全新的值。** `decodeResponse`（response.go）先检查 `code` 信封，非零时返回 `Error`，再把 `data` 严格反序列化到一个全新值，因此 `out` 绝不会被部分写入。严格解码失败时，`decode_tolerate.go` 会把 JSON kind 与 Go 类型不兼容的叶子替换成 `null`，再解码到**另一个全新值**重试；重试成功则写入 `out`，并通过 `DroppedFieldHandler` 上报被丢弃的字段。根节点不参与剪枝（`data` 整体形态变化仍然报错），实现了自定义 unmarshaler 的类型是不透明边界。容错重试仍失败时才用**未剪枝的原始字节**生成诊断。WBI 密钥请求是例外：它走 `decodeWBIResponse`，忽略非零业务码且不容错，因为 nav 可能在返回可用签名密钥的同时给出非零 `code`。`out == nil` 表示"只检查状态码与业务码"。
- **`execute[Out]` 对返回类型泛型**；`Out` 通常是 `*SomeResult` 或 `any`。

### 参数编码（`params.go`）

`withParams` 通过反射处理参数结构体。位置与省略行为**只**由 `request` 结构体标签控制，与 `json` 标签无关：

| 标签 | 行为 |
| --- | --- |
| （无）／`request:"query"` | URL 查询参数（POST 不会自动转为表单） |
| `request:"json"` | JSON 请求体字段 |
| `request:"form-data"` | multipart 文本字段 |
| `request:"-"` | 跳过该字段（未导出字段同样总是跳过） |
| `request:"field=name"` | 显式指定参数名；否则取 `json` 标签名；再否则用 Go 字段名的 snake_case |
| `request:"omitempty,default=1"` | 零值省略；未省略时使用字符串默认值 |

参数名推导（`toSnakeCase`）把 Go 初始缩写（`ID`、`URL`、`UID`、`API` 等）整体视作一个词，因此 `IDs` -> `ids`、`DynamicID` -> `dynamic_id`。同一字段声明多个位置，或混用 JSON 与 multipart，都会返回错误（`ParamError`）。请求体中只允许一种编码（json **或** form-data）；query 可以与一种请求体并存。

### WBI 签名（`wbi.go`）

`WBI` 从 `/x/web-interface/nav` 获取 `img_url`/`sub_url`，把派生密钥缓存到可替换的 `Storage`（`wbi_storage.go`，默认 `MemoryStorage`），并且**只对 query 参数**签名：追加 `wts` + `w_rid`（排序、清洗后的 query 拼接 mixin key 的 MD5）。密钥在 `updateCheckerInterval`（默认 60 分钟）后刷新。一个可取消的 channel（`wbi.refresh`）用于串行化刷新，且不会产生后台请求。通过 `Request.WBI` 或 `fillWbi()` 处理函数逐请求启用。注意已记录的坑：WBI 请求不能带 `Referer` 头 —— `fillWbi` 会显式清空它。query 的每个键必须且只能有一个值。

### CSRF

CSRF 从不作为用户填写的字段。处理函数调用 `csrfValue(r)`，它从请求 Cookie 快照读取 `bili_jct`，缺失时返回稳定的"登录过期"错误。`fillCsrf` 同时写入 `csrf` 与 `csrf_token`；各接口的处理函数按该接口要求把值放入 query 或表单（两种写法可参考 `live.go`）。

## 错误与解码诊断

错误类型定义在 `errors.go`，为 `errors.Is` / `errors.As` 设计：

- `*HTTPError` —— `Method`、`Endpoint`、`StatusCode`。Endpoint 经 `safeEndpoint` 清洗（去掉 query/fragment/user）。不包含响应体、请求头或 Cookie。
- `Error` —— 响应信封中的业务错误（`code`/`message`）。
- `*ParamError` —— `RootType`、`GoField`（可能含切片下标，如 `IDs[2]`）、`Parameter`、`Location`，以及包装的 `Err`（`Unwrap`）。`Error()` 刻意不输出参数值和原始错误文本。
- `*DecodeError` —— `Method`、`Endpoint`、`RootType`、`GoField`、`JSONPath`（例如 `$.data.items[3].modules.module_author.mid`）、`Expected`、`Actual`、从 1 开始计数的 `Offset`、`Exact`，以及包装的 `Err`。

网络故障、取消和超时保留原始错误链（不会转换成 `HTTPError`）。`decode_diagnostic.go` / `decode_fields.go` 实现失败后的 JSON 路径 / 偏移定位；它**仅在**解码失败后运行，绝不重复执行自定义 unmarshaler，当只能定位到边界时标记 `Exact=false`，且始终基于未剪枝的原始字节，因此 `JSONPath` 与 `Offset` 不受容错影响。`DecodeError` 只在容错也失败时返回；被容错丢弃的字段不是错误，改用 `DroppedField` 上报（`decode_tolerate.go`，经 `SetDroppedFieldHandler` 注册），同样不含字段值。不要把 `ParamError.Err`、响应值、Cookie 或凭证泄漏到日志中。

⚠️ **新增的错误要用包级静态变量**（`var errXxx = errors.New(...)`），别在函数里现写。golangci-lint 开着 `err113`（不许定义动态错误）与 `staticcheck` 的 `ST1005`（错误串不能以大写字母开头），而**这两条规则只认标准库的 `errors.New`** —— 库里既有的那些内联写法走的是 `github.com/pkg/errors`，所以历史代码从没被扫到。**新代码别照抄它们**：2026-09-26 加转发接口时，同一个 `errors.New` 连撞这两条。中文错误串尤其容易中 `ST1005`（「B站…」的 B 是拉丁大写）。

## 客户端构造与会话

- `New()` —— 离线创建，不联网；只构造带 B 站风格默认请求头和 20 秒超时的 Resty 客户端，并使用 `NoRedirectPolicy`。
- `NewAnonymousClient(ctx)` —— 从首页获取游客 Cookie；非 200 或 Cookie 为空时返回错误。构造失败后不要继续使用该客户端。
- `NewWithClient(restyClient)` —— 取得**独占所有权**：把显式 `Cookies` 复制到 `Client` 存储，清空 `restyClient.Cookies`，并关闭 CookieJar，保证只有一份会话来源。
- Cookie 辅助方法：`SetRawCookies`（浏览器 `k=v; k=v` 格式）、`GetCookiesString`/`SetCookiesString`（以换行分隔的 `Cookie.String()` 形式）。不要混用这两种格式。同名 Cookie 仅按名称合并（不匹配 domain/path）；正的 `MaxAge` 会被归一化为绝对 `Expires` 并清零 `MaxAge`，因此快照不会续期。
- `Resty()` 暴露底层客户端以供配置，但直接使用 Resty 发请求会跳过签名、Client Cookie 存储与统一错误处理。
- 认证流程（扫码 / 密码 / 短信）、Cookie 持久化与 Resty 接管详见 `docs/authentication.md`。

## 命名规范（见 `.github/CONTRIBUTING.md`）

- 方法名逐词翻译中文接口描述（例如 `GetLiveAreaList`）。术语遵循 `bilibili-API-collect`（专栏文集为 `Articles`，专栏文章为 `Article`）。
- 参数结构体：`XxxParam`；结果结构体：`XxxResult`。对于 `Get*` 类方法，结果类型可省略 `Get` 和 `Result`（例如 `LiveAreaList`）。没有参数就不要建空结构体，传 `nil`。
- `csrf` 不作为调用方传入的参数。
- 能在多文件间复用的子类型尽量复用；共享类型放入 `type.go`。注意通过前缀避免文件之间的命名冲突。
- 可能为 `null` 的字段应使用指针，避免 `json.Unmarshal` 失败。

## 行尾

`.gitattributes` 强制 `*.go` 使用 LF。不要让 `core.autocrlf` 检出 CRLF —— 否则 gofmt CI 会把每个文件都报为未格式化。

## 被忽略的目录

`.gitignore` 忽略了 `test/`、`.idea/`，以及编码助手在本地生成的计划稿目录（具体路径见 `.gitignore`）。该目录下自动生成的内容不要提交；其余可共享的助手配置（如 `settings.json`、`rules/` 等）仍可入库供团队共享。

## 文档

- `README.md` —— 概览、特性、快速开始、各业务域用法示例、贡献、验证与发版。
- `docs/authentication.md` —— 游客初始化、扫码 / 密码 / 短信登录、Cookie 保存与恢复、Resty 接管。
- `docs/request.md` —— `Client.Do` 逃生通道、`request` 标签规则、错误分类、解码诊断。
- `docs/migration.md` —— 模块路径、context 签名、会话规则，以及 v0 各轮重构中的模型与字段重命名。
- `docs/versioning.md` —— tag 与模块版本的格式约束（含禁止 build metadata）、递增规则、v0 与 v1 之后的兼容性约定。
