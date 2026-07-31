# 7 天 Go 后端课程教学计划与讲师规划指南

## 一、课程定位

本课程面向已有编程基础的学员，默认学员已经理解变量、函数、流程控制、集合、基础面向对象或模块化编程等通用编程概念。

课程目标不是从零教授编程，而是帮助学员快速从已有语言经验迁移到 Go，并通过连续项目实践建立 Go 后端开发的完整认知地图。

课程覆盖：

- Go 语言核心特性
- 并发编程与工程化基础
- Gin Web API 开发
- GORM 数据持久化
- Protobuf 与 gRPC
- go-zero 微服务框架
- 认证、事务、缓存、网关、可观测性等企业级后端概念入口

课程节奏偏快，强调“概念快速建立 + 项目驱动理解 + 工程链路串联”。

## 二、整体课程安排

课程共 7 天，每天按 9:30-16:00 安排。名义上每天约 6.5 小时，但实际授课时需要预留环境问题、提问、午休、代码调试和课堂练习时间，因此每天建议按 5-6 个教学块规划。

| 天数 | 主题 | 当天主线成果 |
|---|---|---|
| Day 1 | Go 快速迁移与核心语法 | 能读写基础 Go 代码，完成小型命令行程序 |
| Day 2 | Go 进阶、并发与工程化 | 理解接口、错误、并发、context、测试，完成并发日志分析器 |
| Day 3 | Web 开发与 Gin | 完成简易用户中心 API |
| Day 4 | GORM 与数据持久化 | 完成博客或用户中心数据层 |
| Day 5 | Protobuf 与 gRPC | 实现 gRPC 服务与客户端，理解服务契约 |
| Day 6 | go-zero 与微服务工程 | 跑通电商 API 到 RPC 的调用链 |
| Day 7 | 综合项目、企业级专题与答疑验收 | 串联完整后端工程认知图 |

## 三、对外教学计划

### Day 1：Go 快速迁移与核心语法

主题：从已有编程经验迁移到 Go。

#### 教学目标

学员能够读懂并编写基础 Go 程序，理解 Go 与 Java、C、Python 等语言在类型、函数、集合、指针和结构体上的关键差异。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-10:30 | Go 环境、`go run`、`go build`、`go mod`、包与工程结构 |
| 10:30-11:30 | 变量、常量、零值、类型转换、控制流、函数、多返回值 |
| 11:30-12:30 | 数组、slice、map、string、rune、range 常见问题 |
| 13:30-14:30 | 指针、`nil`、`new` / `make`、值传递、内存基本认知 |
| 14:30-15:30 | struct、method、tag、简单 `defer`、error/interface 入门 |
| 15:30-16:00 | 小练习：CLI Task Manager / 单词统计 / 简单数据建模 |

#### 当天产出

学员能够完成一个小型命令行程序，使用 slice、map、struct、method 组织数据。

### Day 2：Go 进阶、并发与工程化

主题：写出可测试、可并发、可维护的 Go 代码。

#### 教学目标

学员理解 Go 的接口、错误处理、defer/panic/recover、goroutine、channel、context 和基础测试方法。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-09:50 | 环境检查、Entry Quiz |
| 09:50-10:40 | Block 1：interface、错误包装与 defer |
| 10:40-10:50 | 短休息 |
| 10:50-12:00 | Block 2：goroutine、WaitGroup、channel、select |
| 12:00-13:00 | 午休 |
| 13:00-14:05 | Block 3：context、Mutex/RWMutex/atomic、race detector |
| 14:05-14:50 | Block 4：testing、benchmark 与 Reflection |
| 14:50-15:00 | 短休息 |
| 15:00-15:40 | 并发日志分析器综合 Lab |
| 15:40-16:00 | 文件扫描器作业启动、Exit Quiz |

#### 当天产出

学员能够完成一个可取消、可测试、无数据竞态的并发日志分析器，并能独立启动并发文件扫描器作业。

### Day 3：Web 开发与 Gin

主题：从 HTTP 到可维护 Web API。

#### 教学目标

学员理解 HTTP 服务模型、Gin 路由、中间件、请求绑定、JWT 认证、配置和日志。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-10:30 | `net/http`、Handler、TCP/HTTP 基本关系 |
| 10:30-11:30 | Gin 入门、路由、路由组、路径参数、查询参数 |
| 11:30-12:30 | JSON binding、参数校验、统一响应结构 |
| 13:30-14:30 | middleware 洋葱模型、日志中间件、CORS、请求链路 |
| 14:30-15:30 | JWT 登录认证、Viper 配置、Zap 日志 |
| 15:30-16:00 | 用户中心项目：注册、登录、鉴权接口串联 |

#### 当天产出

学员能够完成一个简易用户中心 API，理解 handler / service / repository / model 分层。

### Day 4：GORM 与数据持久化

主题：把 Web API 接入数据库。

#### 教学目标

学员掌握 GORM 的模型定义、CRUD、关联关系、事务、高级查询和测试基本方法。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-10:30 | MySQL 连接、DSN、连接池、GORM 初始化 |
| 10:30-11:30 | Model、tag、AutoMigrate、一对一/一对多/多对多 |
| 11:30-12:30 | CRUD、批量操作、零值更新问题 |
| 13:30-14:30 | Where、Joins、Preload、N+1 查询问题 |
| 14:30-15:30 | Transaction、事务边界、Raw SQL |
| 15:30-16:00 | 博客 API / 用户中心接数据库项目串联 |

#### 当天产出

学员能够完成一个带数据库的博客或用户中心数据层，理解 repository 和事务边界。

### Day 5：Protobuf 与 gRPC

主题：服务间强类型通信。

#### 教学目标

学员理解 RPC、Protobuf、gRPC 代码生成、Unary RPC、Streaming RPC、metadata、interceptor 和错误语义。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-10:30 | RPC vs REST、Protobuf 语法、field number、兼容性 |
| 10:30-11:30 | protoc、生成代码、pb.go / grpc.pb.go 解读 |
| 11:30-12:30 | Unary RPC：server、client、context timeout |
| 13:30-14:30 | Streaming RPC：server stream、client stream、bidi stream |
| 14:30-15:30 | metadata auth、interceptor、status/codes |
| 15:30-16:00 | grpc-gateway 简介 + 分布式计算项目串联 |

#### 当天产出

学员能够实现一个 gRPC 服务和客户端，理解 `.proto` 如何成为服务契约。

### Day 6：go-zero 与微服务工程

主题：用框架组织真实微服务项目。

#### 教学目标

学员理解 go-zero 的 API/RPC 代码生成、服务发现、缓存、消息队列、可观测性和电商调用链。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-10:30 | go-zero 全景、与 Gin/gRPC 原生开发对比 |
| 10:30-11:30 | `.api` 文件、API 服务生成、handler/logic/config |
| 11:30-12:30 | `.proto` + zRPC、API 调 RPC |
| 13:30-14:30 | Etcd 服务注册发现、服务调用链 |
| 14:30-15:30 | MySQL/cache、MQ、metrics、Prometheus 简介 |
| 15:30-16:00 | 电商项目：order-api -> order-rpc -> user/product-rpc |

#### 当天产出

学员能够看懂 go-zero 标准工程结构，并跑通一个简化电商微服务调用链。

### Day 7：综合项目、企业级专题与答疑验收

主题：把知识点串成后端工程认知地图。

#### 教学目标

学员能够从整体上理解一个 Go 后端系统的分层、调用链、数据流、错误处理、认证、事务和微服务边界。

#### 课程安排

| 时间 | 内容 |
|---|---|
| 09:30-10:30 | 全课程主线复盘：Go -> Gin -> GORM -> gRPC -> go-zero |
| 10:30-11:30 | 综合项目功能补充：订单状态、库存校验、分页、鉴权 |
| 11:30-12:30 | 项目调试、测试、接口联调、常见错误排查 |
| 13:30-14:30 | 企业级专题：API 设计、认证安全、事务边界 |
| 14:30-15:30 | 企业级专题：缓存一致性、限流熔断、可观测性、部署 |
| 15:30-16:00 | 学员展示、答疑、后续学习路线 |

#### 当天产出

学员形成完整 Go 后端开发认知图，并具备继续扩展期末项目或实战项目的基础。

## 四、讲师自用规划指南

### 1. 总体节奏原则

这批学员有编程基础，所以不要按零基础语法课讲。

Day 1 的核心不是“变量是什么”，而是“Go 和学员以前写过的语言有什么不一样”。

讲课时可以多使用以下表达：

- “这个概念你们在 Java 或其他语言里见过，但 Go 的做法不同。”
- “这里不是语法难，而是语义容易误解。”
- “这个点后面 Gin、GORM、gRPC 会反复出现，现在先建立直觉。”

每天都要落到一个可运行成果，主线项目不能砍，扩展知识可以砍。

### 2. 每天的主线抓手

| 天数 | 讲师需要抓住的主线 |
|---|---|
| Day 1 | 让学生能读懂 Go 代码 |
| Day 2 | 让学生能写出不乱飞的并发代码 |
| Day 3 | 让学生知道 HTTP 请求怎么进业务代码 |
| Day 4 | 让学生知道业务对象怎么落数据库 |
| Day 5 | 让学生知道服务之间怎么强类型通信 |
| Day 6 | 让学生知道真实微服务工程长什么样 |
| Day 7 | 把所有模块串成一个后端系统 |

## 五、每日讲课小抄

### Day 1 小抄：Go 快速迁移

#### 必讲

- `package main`
- `go run` / `go build`
- `go mod`
- 零值
- 显式类型转换
- 多返回值
- slice 的 `len` / `cap` / `append`
- map 的 comma-ok
- string / rune / byte
- 指针不能运算
- `new` vs `make`
- struct + method
- 简单 error / defer / interface

#### 可砍

- `goto`
- 泛型深入
- GC 深入
- 逃逸分析深入
- 函数式编程深入
- 文件 IO 全套
- 反射

#### 讲课提醒

Day 1 最容易讲爆的是 slice、指针、struct。

如果时间紧，优先保 slice 和 struct。指针讲到“能改值、避免拷贝、nil 会 panic”即可。

### Day 2 小抄：进阶与并发

#### 必讲

- interface 隐式实现
- typed nil
- `error` 是接口
- `fmt.Errorf("%w")`
- `errors.Is` / `errors.As`
- defer 参数预计算
- panic/recover 只用于兜底
- goroutine 生命周期
- WaitGroup
- channel 阻塞
- select
- context cancel/timeout
- mutex/race detector
- testing 表格驱动

#### 可砍

- GMP 深入
- runtime 调度细节
- reflect 深入
- atomic 底层实现
- pprof 深入

#### 讲课提醒

并发这天不要讲成操作系统课。

让学生记住这条主线即可：

```text
goroutine 负责并发执行
channel 负责通信
context 负责生命周期
mutex 负责共享内存保护
```

### Day 3 小抄：Gin

#### 必讲

- `net/http` Handler
- Gin router
- Gin `Context`
- bind
- middleware
- JWT
- config
- log
- handler / service / repository 分层

#### 可砍

- TCP/UDP 深入
- Swagger 深入
- CORS 所有细节
- 优雅重启源码级分析

#### 讲课提醒

这天要反复画请求链路：

```text
HTTP Request
-> Router
-> Middleware
-> Handler
-> Service
-> Repository
-> Response
```

学生只要吃透这条线，后面 GORM 和 go-zero 都好讲。

### Day 4 小抄：GORM

#### 必讲

- DSN
- model tag
- AutoMigrate 的边界
- CRUD
- 零值更新坑
- Preload
- Joins
- transaction
- repository 层
- 数据库测试分类

#### 可砍

- gorm gen
- Hooks 深入
- 复杂多对多
- sqlmock 严格匹配细节
- 迁移工具完整体系

#### 讲课提醒

Day 4 不要只教 API 调用，要强调：

```text
ORM 不是魔法，它只是帮你把 struct 和 SQL 映射起来。
```

多让学生看 SQL 日志，效果会更好。

### Day 5 小抄：gRPC

#### 必讲

- `.proto` 是契约
- field number 不能乱改
- protoc 生成代码
- server interface
- client stub
- unary
- streaming
- metadata
- interceptor
- status/codes

#### 可砍

- Protobuf wire format
- HTTP/2 帧细节
- grpc-gateway 深入
- 复杂错误详情
- TLS 证书实操

#### 讲课提醒

gRPC 这天容易被工具链打断。

课前一定确认：

- `protoc`
- `protoc-gen-go`
- `protoc-gen-go-grpc`
- `grpcurl`

如果工具链不稳，就少现场生成，多展示已有生成文件。

### Day 6 小抄：go-zero

#### 必讲

- 为什么需要框架约束
- `.api` 到 API 服务
- `.proto` 到 RPC 服务
- API 调 RPC
- Etcd 是服务发现
- MySQL/cache 是数据访问能力
- MQ 是异步解耦
- metrics 是观测入口
- 电商调用链

#### 可砍

- APISIX / Envoy 深入
- DTM 分布式事务深入
- Jaeger 实操
- K8s 深入
- go-zero 源码
- 完整压测

#### 讲课提醒

Day 6 不是“微服务百科全书”。

你要让学生知道真实工程结构长什么样：

```text
api/order
rpc/order
rpc/user
rpc/product
config
logic
svc
proto
```

只要他们能看懂目录、配置和调用链，这天就成功了。

### Day 7 小抄：综合收束

#### 推荐优先专题

如果学生吸收一般，讲：

- API 设计
- 认证安全
- 数据库事务边界

如果学生吸收较好，追加：

- 缓存一致性
- 限流熔断
- 可观测性

如果学生很强，再讲：

- 网关
- 部署发布
- 服务边界
- 分布式事务

#### 讲课提醒

Day 7 不要继续狂塞新技术。

这天的价值是帮学生把脑子里的碎片拼起来。

可以反复问学生：

- 一个请求从哪里进来？
- 认证在哪里做？
- 参数在哪里校验？
- 业务逻辑在哪里？
- 数据在哪里落库？
- 服务之间怎么通信？
- 失败了怎么返回？
- 超时了谁负责取消？
- 如何知道线上哪里慢？

## 六、每天固定讲课结构

建议每天都按这个节奏组织：

1. 10 分钟：今天解决什么问题
2. 40 分钟：核心概念
3. 60 分钟：最小 demo
4. 60 分钟：项目代码串联
5. 60 分钟：学生动手改一个功能
6. 30 分钟：复盘 + 生产边界

不要让课程变成“API 文档朗读会”。

每个模块都要落到一个问题：

- Day 1：怎么写 Go？
- Day 2：怎么写并发？
- Day 3：怎么写 API？
- Day 4：怎么接数据库？
- Day 5：服务怎么通信？
- Day 6：工程怎么组织？
- Day 7：系统怎么串起来？

## 七、控时原则

每个知识点按三档准备：

### A 档：必须讲清

学生不懂就没法继续。

### B 档：讲个直觉

学生先知道有这个东西，后面项目里再见。

### C 档：只作为扩展

有时间讲，没时间砍掉也不影响主线。

最重要的是：

```text
主线项目不能砍，扩展知识可以砍。
```

## 八、最终课程目标

这套课最好的最终形态，不是让学生记住所有 Go 语法细节，而是让他们脑子里形成这张图：

```text
Go 基础语法
-> 并发与工程化
-> Gin Web API
-> GORM 数据库
-> gRPC 服务通信
-> go-zero 微服务工程
-> 企业级后端系统意识
```

7 天结束后，学生应该感觉自己不是“学了一堆 Go 知识点”，而是“知道一个 Go 后端项目是怎么长出来的”。
