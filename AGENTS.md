# 仓库贡献指南

## 项目结构

本项目是 Bilibili API 的 Go 客户端库，根目录使用单个 `bilibili` 包。接口按业务放在 `video.go`、`user.go`、`live.go` 等文件，共享类型放在 `type.go`；新增接口沿用对应业务分类，不按 controller/service/repository 拆层。

- `client.go`、`wbi.go`：客户端配置、Cookie 与 WBI 签名。
- `request.go`、`params.go`：请求执行与参数编码。
- `response.go`、`decode_diagnostic.go`、`errors.go`：响应解码与错误定位。
- `number.go`：数值／字符串兼容类型。
- `dynamic_model.go`：`DynamicItem`、`DynamicInfo` 及主要动态模块；动态请求方法和参数仍在 `dynamic.go`。
- `live_medal_model.go`：勋章墙、勋章面板和已激活勋章任务模型；请求方法、参数及点赞上报位于 `live.go`。勋章墙与面板保留独立字段类型，不按名称相似合并。
- `topic_model.go`：话题动态列表的完整已声明模型；`GetTopicFeed` 请求方法和参数位于 `dynamic.go`。
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

所有可能联网的公开方法以 `ctx context.Context` 为首参，包括独立 WBI 的取密钥与签名方法；纯计算和配置方法不加 context。不保留无 context 的兼容包装。`execute` 向 HTTP 请求及密钥刷新传递同一 context，准备请求前拒绝 nil 或已取消的 context，不在库内回退到 `context.Background()`。取消和超时错误保持可用 `errors.Is` 判断。

批量工具在入口使用 `signal.NotifyContext`，向辅助函数贯通任务 context；取消后停止翻页及后续账号操作，等待使用 timer/select。先处理查询错误再访问结果，不忽略写操作错误。取消不表示服务端已撤销执行，不自动重试写操作。

客户端请求必须通过 `newRequest` 获取 Cookie 快照，再经 `sendRaw` 合并响应；CSRF 使用同一快照，不直接读写 `resty.Cookies`。Cookie 按名称合并，读写复制；禁止重新启用底层 Jar。`NewWithClient` 接管 Resty，已有 Jar 会话需构造前显式导出。`NewAnonymousClient(ctx)` 返回 `(*Client, error)`，必须处理错误。

普通请求可并发，登录、主动刷新、账号切换和配置变更串行进行；不要复制已使用的 Client。底层 Resty 直接请求不纳入会话管理或并发保证。未做运行或 race 验证时，不得把编译通过表述为已证明并发正确。

业务错误使用 `Error`，解码失败使用可解包的 `DecodeError`。禁止静默吞错或将无效值转换为零；日志不要输出 Cookie、凭证或完整响应。不要批量将字段改成 `json.Number`：数字字符串与任意字符串应按实际语义区分，缺少响应依据的字段先记录疑点。

HTTP 状态不符合接口要求时使用 `*HTTPError`；保留各接口的成功状态规则，不将网络故障、取消或解析失败归类为 HTTP 状态错误。业务错误在原有值类型 `Error` 外包装方法和安全接口地址，通过 `errors.As` 提取；取消和超时通过 `errors.Is` 判断，不直接断言最外层类型或匹配错误字符串。WBI 非零业务码但存在可用密钥的特殊流程保持不变。

参数先由 `encodeParams` 编码完成，再统一应用到请求。默认进入 query；`request` 标签决定省略、默认值及位置，不将 `json:",omitempty"` 当作请求省略规则。query 可与 JSON 或 multipart 并存，JSON 与 multipart 互斥；不得由字段顺序决定 Content-Type。multipart 交给 Resty 生成 boundary，不手工伪造请求体。

转换失败使用 `ParamError`，通过 `errors.As` 读取根类型、Go 字段（可含切片下标）、参数名及位置。请求层用 `%w` 补充方法和去除查询参数的 URL。错误文本不包含参数值；原始 `Err` 可能包含敏感值，不直接记录。JSON 自定义编码器只执行一次，失败定位到对应顶层参数字段，不承诺完整的嵌套字段路径。

动态模型使用公开命名类型，保留字段顺序、JSON 标签和指针层级。`DynamicOriginalItem` 与外层动态采用独立模型，不合并为递归类型；仅共享完全一致的 `DynamicArchive`、`DynamicDraw`。不要为复用而统一两套作者、富文本或主体类型。嵌套字段的类型身份改变时，应说明复合字面量、赋值和反射代码的迁移方式。

话题模型沿用本地工具的字段定义，不直接替换为 `DynamicItem`；话题与空间动态仅在结构完全一致时复用模块。`test/dynamicItem.txt` 来自 `GetUserSpaceDynamic`，不能作为话题字段改型的证据。`GetTopicFeed` 只获取一页，由调用方显式传入分页和扩展参数。

## Git 与协作

按可审查阶段提交，使用 `feat(request): …`、`fix(decode): …`、`refactor(api): …`、`docs: …` 等简短中文说明。提交前检查暂存差异，仅纳入本次改动；保留用户原有修改，不强制跟踪被忽略的工具、配置或产物。

不同重构目标使用独立分支，分支内分阶段提交。后续阶段依赖前一阶段时，从其最新提交创建分支；前一阶段未合并前，以它作为审查基线，不直接从旧 master 开始。

PR 说明问题、行为变化、相关 issue/API 文档、验证结果及限制。公开方法或字段类型的破坏性变更必须说明迁移方式；仅在任务明确允许时实施。不要擅自升级依赖、改模块路径或发布版本。
