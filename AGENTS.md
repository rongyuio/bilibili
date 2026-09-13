# 仓库贡献指南

## 项目结构

本项目是 Bilibili API 的 Go 客户端库，根目录使用单个 `bilibili` 包。接口按业务放在 `video.go`、`user.go`、`live.go` 等文件，共享类型放在 `type.go`；新增接口沿用对应业务分类，不按 controller/service/repository 拆层。

- `client.go`、`wbi.go`：客户端配置、Cookie 与 WBI 签名。
- `request.go`、`params.go`：请求执行与参数编码。
- `response.go`、`decode_diagnostic.go`、`errors.go`：响应解码与错误定位。
- `number.go`：数值／字符串兼容类型。
- `tools/`：Markdown 表格转 Go 结构体工具；`video_zone.csv`：视频分区数据。
- 根目录 `*_test.go`：已有辅助函数测试；被 Git 忽略的 `test/`：本地实用工具，可能操作真实账号。

## 开发与验证

Go 版本以 `go.mod` 为准。在根目录执行：

- `go build ./...`：编译库、生成器及本地工具，不运行程序。
- `go vet ./...`：静态检查。
- `gofmt -s -w <文件.go>`：仅格式化本次修改的文件，使用标准制表符缩进。
- `golangci-lint run`：执行 v2 配置；排除测试和 `tools/`。工具缺失时如实说明，不宣称通过。

默认以编译和静态检查验证，不要求新增测试。任务明确要求时才运行测试；已有测试使用标准库 `testing`、`TestXxx` 命名，无覆盖率门槛。不得为了验证擅自运行 `test/` 工具或访问真实 API。CI 现有测试配置不因此自动修改。

## API 与编码约定

遵循 `.github/CONTRIBUTING.md` 的业务命名，例如 `GetLiveAreaList`、`GetLiveAreaListParam`、`GetLiveAreaListResult`。保留 JSON/request 标签及接口要求的参数位置。

内置接口复用 `execute`；未封装接口使用 `Client.Do(ctx, Request, out)`，通过 `WBI: true` 启用签名。`out` 接收 `data`，不是完整业务响应。不要重新导出客户端签名器；`Resty()` 仅作底层出口，其请求不会自动获得统一解码。

客户端请求必须通过 `newRequest` 获取 Cookie 快照，再经 `sendRaw` 合并响应；CSRF 使用同一快照，不直接读写 `resty.Cookies`。Cookie 按名称合并，读写复制；禁止重新启用底层 Jar。`NewWithClient` 接管 Resty，已有 Jar 会话需构造前显式导出。`NewAnonymousClient(ctx)` 返回 `(*Client, error)`，必须处理错误。

普通请求可并发，登录、主动刷新、账号切换和配置变更串行进行；不要复制已使用的 Client。底层 Resty 直接请求不纳入会话管理或并发保证。未做运行或 race 验证时，不得把编译通过表述为已证明并发正确。

业务错误使用 `Error`，解码失败使用可解包的 `DecodeError`。禁止静默吞错或将无效值转换为零；日志不要输出 Cookie、凭证或完整响应。不要批量将字段改成 `json.Number`：数字字符串与任意字符串应按实际语义区分，缺少响应依据的字段先记录疑点。

## Git 与协作

按可审查阶段提交，使用 `feat(request): …`、`fix(decode): …`、`refactor(api): …`、`docs: …` 等简短中文说明。提交前检查暂存差异，仅纳入本次改动；保留用户原有修改，不强制跟踪被忽略的工具、配置或产物。

不同重构目标使用独立分支，分支内分阶段提交。后续阶段依赖前一阶段时，从其最新提交创建分支；前一阶段未合并前，以它作为审查基线，不直接从旧 master 开始。

PR 说明问题、行为变化、相关 issue/API 文档、验证结果及限制。公开方法或字段类型的破坏性变更必须说明迁移方式；仅在任务明确允许时实施。不要擅自升级依赖、改模块路径或发布版本。
