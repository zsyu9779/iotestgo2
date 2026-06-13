# iotestgo2 项目调研报告

生成日期：2026-06-14
项目路径：`/Users/zhangshiyu/iotestgo2`
调研对象：计算机专业本科生 Go 后端讲课演示项目

---

## 1. 执行摘要

`iotestgo2` 现在已经从“Go 后端示例集合”提升为一套较完整的本科课堂演示项目。项目主线覆盖 Go 语言基础、并发与工程化、Gin Web、GORM、gRPC、go-zero 微服务，并新增了企业级后端视野拓展专题。它不适合作为生产项目模板，但非常适合作为课堂讲解、课后练习、综合实践和期末项目选题的材料库。

本轮改造后，项目的主要变化是：

- 明确课程对象是计算机专业本科生，Java 对比只作为连接已有知识的辅助脚手架。
- 增加统一课程入口、运行手册、Makefile 验证目标和 Module 03-06 教案。
- 补齐 gRPC 综合项目的 testable engine、stream service、auth interceptor 和 annotation-based grpc-gateway。
- 增加标准 go-zero 电商工程，形成真实 `order-api -> order-rpc -> user-rpc + product-rpc` 调用链。
- 增加 go-zero 标准工程的 Docker Compose、Prometheus metrics 配置和教学路线说明。
- 增加 12 个企业级后端拓展专题和 3 个可选 Module 07 小 demo。

总体判断：项目已经适合正式授课试运行。剩余风险主要不在“能不能讲”，而在“课堂环境是否统一”和“哪些内容应该选讲”。建议第一次正式授课时采用主线必讲 + 拓展选讲的方式，不要试图一次讲完所有高级专题。

---

## 2. 当前规模与组成

| 指标 | 当前数量 |
|---|---:|
| 仓库文件总数 | 276 |
| Go 源码文件 | 186 |
| Markdown 文档 | 39 |
| 测试文件 | 25 |
| Proto 文件 | 13 |

### 2.1 顶层课程资产

- `README.md`：项目定位、模块总览、快速开始和教学边界。
- `COURSE_RUNBOOK.md`：课前检查、Go 版本要求、课堂运行顺序和故障处理。
- `Makefile`：统一 `fmt-check`、`test-basic`、`test-race` 和常见 demo 运行命令。
- `.env.example`：课堂共享默认配置样例。
- `Golang_Backend_Training_Syllabus.md`：46 课时课程大纲。

### 2.2 教师备课材料

`docs/` 下已包含：

- Module 01-06 lesson plan。
- Module 01/02 讲台小抄。
- Go 暗角补充材料。
- 企业级后端拓展专题。
- 本调研报告 Markdown 与 HTML 版本。

### 2.3 课程模块

| 模块 | 定位 | 当前状态 |
|---|---|---|
| `module01_basics` | Go 基础、集合、结构体、泛型、CLI 项目 | 稳定，适合作为入门主线 |
| `module02_advanced` | 接口、错误、并发、context、测试、runtime | 稳定，测试和 race 示例较完整 |
| `module03_web_gin` | Gin、JWT、中间件、httptest、用户中心 | 可授课，需明确安全简化 |
| `module04_gorm` | GORM、事务、Preload、Raw SQL、博客 API | 可授课，依赖 MySQL 环境 |
| `module05_grpc` | Protobuf、gRPC、stream、metadata、gateway | 已补强，适合作为 RPC 主线 |
| `module06_gozero` | go-zero、API/RPC、Etcd、缓存、MQ、观测、部署 | 已增加概念版 + 标准工程双轨 |
| `module07_enterprise_extensions` | 企业级后端可选小 demo | 新增，适合拓展课或期末项目前 |

---

## 3. 课程定位分析

这个项目不是“Java 工程师转 Go”培训，而是面向计算机专业本科生的 Go 后端课程。课程中保留 Java 对比是合理的，因为很多学生在学校主语言环境中接触过 Java。它的作用是降低迁移成本，而不是把 Go 概念强行翻译成 Java 概念。

推荐课堂表达：

- 可以说“如果你学过 Java，可以这样类比”。
- 不建议说“Go 就是 Java 的某某替代品”。
- 对接口、错误处理、并发、context、结构体组合、微服务边界等内容，要强调 Go 自身的设计取舍。

当前课程主线比较健康：

1. Module 01/02 先建立语言和工程基本功。
2. Module 03/04 进入单体 Web + DB。
3. Module 05 讲服务间强类型通信。
4. Module 06 讲微服务框架和工程化生态。
5. Module 07 和 `docs/enterprise_backend_extensions` 用作企业级视野拓展。

---

## 4. 模块调研结论

### 4.1 Module 01：Go 语言基础

优势：

- 小节颗粒度清晰，适合一节课一个目录。
- CLI Task Manager 能自然串联 slice、map、struct、pointer、method。
- 已包含 string/map/range 等暗角材料和 benchmark。

注意点：

- `project_task_manager` 仍有业务逻辑与 IO 混合的问题，可作为“可测试设计”重构练习。
- 泛型、高阶函数等内容适合按学生吸收速度选讲，不宜压在第一轮主线里。

### 4.2 Module 02：高级特性与工程化

优势：

- 接口、错误、defer、goroutine、channel、context、锁和测试覆盖完整。
- `project_log_analyzer` 有 pipeline 测试和 race 验证入口。
- `Makefile` 已提供 `test-race`，适合课堂现场演示竞态检测。

注意点：

- 日志分析器主要是模拟数据，不是真实大文件读取。可在作业中扩展为真实文件版。
- typed nil、defer 参数预计算、range 变量等暗角适合第二遍复盘讲，不建议第一遍全部展开。

### 4.3 Module 03：Gin Web

优势：

- 从 `net/http` 到 Gin 的路径自然。
- 用户中心项目采用 handler/service/repository/model 分层，适合本科生理解 Web 后端结构。
- 有 JWT、中间件、Zap、httptest 等常用能力。

风险与教学口径：

- 明文密码、硬编码 JWT secret、内存仓库都是教学简化，不是生产写法。
- `docs/enterprise_backend_extensions/topic_02_auth_security.md` 已提供密码 hash、secret 管理和敏感日志的拓展专题，可在 JWT 后插入。

### 4.4 Module 04：GORM

优势：

- 覆盖连接池、模型关系、CRUD、Preload、Hooks、Migration、Transaction、Raw SQL。
- 博客 API 有 service/repository 分层和事务测试。
- 可自然引出“事务边界”和“数据库建模”两个企业级专题。

风险与教学口径：

- GORM 示例依赖 MySQL，课前必须确认数据库或 Docker 环境。
- 硬编码 DSN 适合课堂简化，但正式工程要迁移到配置。
- 多对多 Tag、级联删除、sqlmock 严格断言仍可作为后续增强点。

### 4.5 Module 05：gRPC

本模块已经显著补强。

新增或改进点：

- `project_distributed_compute/internal/engine`：计算逻辑独立并有单元测试。
- `project_distributed_compute/internal/server`：stream service 可测试，并避免并发 goroutine 直接写同一个 stream。
- `project_distributed_compute/internal/auth`：基于 metadata 的 stream auth interceptor。
- `08_grpc_gateway`：从手写 HTTP bridge 改为 proto annotation + grpc-gateway 生成代码。
- `module05_grpc/README.md`：补充建议讲课顺序和生成代码规则。

教学价值：

- 可以先讲小 demo，再讲综合项目。
- 综合项目可用于展示“算法核心、transport、auth、并发、测试”如何拆开。
- grpc-gateway 可以真实展示 `.proto` 作为 HTTP/gRPC 双协议契约来源。

仍需注意：

- `protoc`、`protoc-gen-go`、`protoc-gen-go-grpc`、`protoc-gen-grpc-gateway` 必须课前装好。
- 生成文件不要手改，已经在 README 中明确。

### 4.6 Module 06：go-zero

本模块现在形成双轨结构。

Track A：概念演示

- `01_gozero_intro` 到 `08_k8s_deploy` 继续用于第一遍讲概念。
- 适合讲 go-zero 的术语、生态、注册发现、缓存、MQ、观测和部署。

Track B：标准工程

- 新增 `project_ecommerce_standard`。
- 使用 goctl 风格目录。
- 提供 API/RPC/proto/config/service context/logic 结构。
- 形成真实调用链：`HTTP Client -> order-api -> order-rpc -> user-rpc + product-rpc`。
- Product RPC 维护内存库存并支持预留扣减。
- User RPC 对缺失用户返回 `NotFound`，禁用用户返回业务无效。
- Prometheus metrics 使用独立端口 `19100-19103`，避免误抓业务端口。

教学价值：

- 可以把概念版作为“看懂 go-zero 做什么”，把标准版作为“知道真实工程长什么样”。
- 适合期末项目基础，也适合讲微服务调用链、服务发现、配置、错误语义、库存状态和观测端口。

仍需注意：

- 本机没有 Docker，无法在当前环境运行 `docker compose up` 做容器验收。
- 标准工程当前仍使用内存数据，适合教学，不是生产电商系统。
- 真实环境还需要数据库持久化、幂等、事务/补偿、分布式锁或库存服务更严谨设计。

### 4.7 Module 07：企业级拓展 demo

新增可选 demo：

- `01_api_compatibility`：版本路由、统一响应、request_id。
- `02_resilience`：timeout、retry、rate limit 的最小模拟，并有测试。
- `03_pprof`：故意低效 Fibonacci endpoint + benchmark。

教学价值：

- 小而完整，适合 10-20 分钟插入课堂。
- 不引入第三方依赖，避免学生被工具链打断。
- 与 `docs/enterprise_backend_extensions` 的专题一一呼应。

---

## 5. 企业级拓展方向设计

拓展专题位于 `docs/enterprise_backend_extensions/`，定位是“看学生吸收速度选讲”，不是隐藏必修作业。

### 5.1 12 个专题

1. API 设计与兼容性
2. 认证、密码与密钥管理
3. 数据库建模与事务边界
4. 缓存一致性
5. 消息队列与最终一致性
6. 可观测性：日志、指标、链路追踪
7. 韧性设计：超时、重试、熔断、限流
8. API Gateway 与边缘层职责
9. 部署、发布与回滚
10. 性能分析与容量意识
11. 服务边界与领域建模
12. 数据隐私、审计与合规意识

### 5.2 推荐选讲策略

| 班级状态 | 推荐专题 |
|---|---|
| 学生刚能跑通代码 | API 设计、认证安全 |
| 学生能完成 Gin + GORM | 数据库事务、缓存一致性、可观测性 |
| 学生能理解 gRPC + go-zero | MQ、韧性、Gateway、服务边界 |
| 期末项目前 | 部署发布、性能分析、数据隐私 |

### 5.3 教学提醒

企业级专题容易讲成工具名堆叠。建议每个专题都按三问组织：

1. 为什么需要这个问题？
2. 什么时候不需要？
3. 学生能做一个什么小实验？

---

## 6. 验证记录

已通过的本地验证：

```bash
GOFLAGS=-mod=readonly make test-basic
GOFLAGS=-mod=readonly make test-race
make fmt-check
GOFLAGS=-mod=readonly go test ./module05_grpc/project_distributed_compute/internal/... ./module05_grpc/08_grpc_gateway ./module06_gozero/project_ecommerce_standard/...
go test ./module07_enterprise_extensions/...
go test -bench=. ./module07_enterprise_extensions/03_pprof
```

手动验证过：

- `module05_grpc/08_grpc_gateway` 的 HTTP `POST /v1/hello` 能经 grpc-gateway 返回 gRPC 响应。
- `module07_enterprise_extensions/01_api_compatibility` 能启动并返回统一 JSON envelope。
- `module07_enterprise_extensions/02_resilience` 输出 timeout、retry、rate limit 三段结果。

未完成的本机验证：

- Docker Compose 验证未运行，因为当前机器 `docker` 命令不存在：`zsh:1: command not found: docker`。
- 因此 Module 06 的 Etcd/MySQL/Redis/Prometheus 容器链路需要在有 Docker 的课堂机或 CI 环境中再验收。

---

## 7. 当前剩余风险

| 风险 | 影响 | 建议 |
|---|---|---|
| Docker 未安装 | Module 04/06 容器化演示无法本机验证 | 课前统一 Docker Desktop 或替代云环境 |
| 外部工具链依赖 | gRPC/go-zero 生成链路可能现场卡住 | 课前运行 `COURSE_RUNBOOK.md` 中检查项 |
| 安全简化较多 | 学生可能误以为明文密码/硬编码 token 可生产使用 | 在 Module 03 后插入认证安全专题 |
| 后半部分测试仍偏少 | API/RPC 集成行为主要靠局部测试 | 后续可增加 order-api/order-rpc mock 或集成测试 |
| 内容丰富度较高 | 46 课时内容密度大 | 使用 `selection_guide.md` 控制选讲范围 |

当前工作树还有一个未提交改动：`module05_grpc/05_interceptors/main.go`。本轮提交和报告未纳入该文件，需由原修改者决定是否保留、提交或回滚。

---

## 8. 结论与建议

项目已经具备正式授课试运行条件。最适合的使用方式是：

1. 按 Module 01-06 走主线，确保学生能跑、能改、能解释。
2. 对 Module 05/06 采用“概念小 demo -> 标准综合项目”的两段式讲法。
3. 企业级拓展只选 2-5 个专题穿插讲，不把全部专题变成必修。
4. 每次课前运行 `make fmt-check`、`make test-basic` 和必要的模块专项测试。
5. 第一次完整授课后，根据学生卡点把拓展专题重新排序，而不是继续增加新技术点。

一句话评价：`iotestgo2` 当前已经是一套结构清晰、可演示、可扩展的本科 Go 后端课程项目；它的下一阶段重点不是继续堆技术，而是用稳定课堂环境和精选作业把学习闭环打磨出来。
