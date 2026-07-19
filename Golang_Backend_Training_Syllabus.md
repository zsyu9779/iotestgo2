# Go 后端开发课程教学大纲

## 课程概览
*   **目标学生：** 计算机专业本科生，已学习过至少一门编程语言；学校课程以 Java 为主，因此部分章节使用 Java 对比帮助理解。
*   **课时安排：** 总计 46 课时（每节课 1-1.5 小时）
*   **教学风格：** 高密度、快节奏、概念对比、重实战
*   **核心目标：** 建立 Go 后端开发的完整认知地图，掌握 Go 语言特性、Web API、数据库访问、RPC 通信和微服务工程的基本实践。

## 模块 1: Go 语言基础与内存模型（一日四 Block 工作坊 + 一周作业）
> **焦点：** 在 09:30–16:00 的一日工作坊中建立 Go 思维，通过连续实操组合语法、集合、建模、函数和测试；Task Manager 在课后一周独立完成。

*   **Block 1: Go Basics**
    *   Go 工具链、程序入口、变量与零值、控制流、函数和显式错误返回
    *   **Java 对比：** JDK vs Go SDK，Class 入口 vs `package main`，Exception vs Error 返回值
    *   实战：带输入校验的成绩等级计算器
*   **Block 2: Collections**
    *   Array/Slice 值与共享语义、Map comma-ok、String/byte/rune 与 UTF-8
    *   **Java 对比：** ArrayList vs Slice，HashMap vs Map，String vs UTF-8 字节序列
    *   实战：中英文文本统计和词频汇总
*   **Block 3: Modeling**
    *   Go 只有值传递、Pointer、Struct、构造函数惯例、值/指针接收器和组合
    *   **Java 对比：** Class/继承 vs Struct/组合，`this` vs Receiver
    *   实战：Scorebook 核心模型和状态修改
*   **Block 4: Functions & Testing**
    *   函数值、闭包、`defer` 和 `testing` 包；复杂函数式模式留作 Bonus
    *   **Java 对比：** Lambda/函数式接口 vs Go 函数值，`finally` vs `defer`
    *   实战：统计函数、审计顺序和基础测试
*   **🏆 课堂综合 Lab: Scorebook**
    *   组合集合、建模、方法与测试，产生可运行、可验收的课堂成果。
*   **🏆 一周课后作业: 命令行 Task Manager**
    *   功能：新增、列出、完成和删除任务，并处理空标题与未知 ID。
    *   要求：仅使用 Go 标准库；本地 `make grade` 与教师手动运行的 Gitee 流水线调用同一验收脚本。

## 模块 2: Go 高级特性与工程化 (9 节课)
> **焦点：** 掌握 Go 的杀手级特性——并发，以及接口和工程标准。

*   **课 1: 接口 (Interface) 与 鸭子类型**
    *   隐式实现, 空接口 `interface{}`, 类型断言, Type Switch
    *   **Java 对比：** `implements` 关键字 vs 隐式满足, 侵入式 vs 非侵入式设计
*   **课 2: 错误处理与 Defer**
    *   `error` 接口模式, `panic` & `recover` 机制
    *   `defer` 执行顺序与陷阱 (参数预计算)
    *   **Java 对比：** Try-Catch-Finally vs Defer-Panic-Recover
*   **课 3: 并发基石 - Goroutine**
    *   MPG 调度模型简介, Goroutine 轻量级原理
    *   `sync.WaitGroup` 等待组
    *   **Java 对比：** Thread vs Goroutine, 内存占用对比
*   **课 4: 通信顺序进程 - Channel**
    *   无缓冲 vs 有缓冲 Channel, 单向 Channel
    *   `select` 多路复用, 超时控制模式
    *   **Java 对比：** BlockingQueue vs Channel, 共享内存通信 vs 通信共享内存
*   **课 5: Context 上下文管理 (新增重点)**
    *   `context.Context` 核心作用：超时控制、取消信号、请求范围值传递
    *   `WithCancel`, `WithTimeout`, `WithValue`
    *   实战：父子 Goroutine 的生命周期管理
*   **课 6: 并发安全与锁**
    *   `sync.Mutex` vs `sync.RWMutex`, `atomic` 包
    *   Race Detector (`go run -race`) 检测竞态条件
*   **课 7: 单元测试与基准测试**
    *   `testing` 包, 表格驱动测试 (Table-driven tests)
    *   Benchmark 基准测试与性能分析 (pprof) 初探
*   **课 8: 系统交互与文件操作 (新增)**
    *   文件 I/O: `os`, `io`, `bufio` 核心包, 文件读写与拷贝
    *   系统交互: 进程信号处理 (`os/signal`), 执行外部命令
    *   **Java 对比：** `java.io`/`NIO` vs Go `io` Reader/Writer 接口
*   **课 9: 反射与运行时控制 (新增)**
    *   反射机制: `reflect.Type` vs `reflect.Value`, 动态修改值
    *   Runtime: `GOMAXPROCS` 调优, `Gosched` 调度控制
    *   **Java 对比：** `java.lang.reflect` vs `reflect` 包 (Go 反射性能开销更大)
*   **🏆 综合实践任务 2: 并发日志分析器**
    *   功能：启动多个 Goroutine 读取大文件，通过 Channel 汇总错误日志，使用 Context 控制整体超时。
    *   要求：包含完整的单元测试，并通过 Race 检测。

## 模块 3: Web 开发与 Gin 框架 (8 节课)
> **焦点：** 快速构建高性能 RESTful API，理解 Go Web 的中间件模型。

*   **课 1: 网络编程基础 (Socket & HTTP)**
    *   Socket 编程: TCP/UDP 通信原理, `net` 包实战
    *   HTTP 标准库: `net/http` 快速启动 Server, Handler 接口
    *   **Java 对比：** `java.net.Socket` vs `net` 包, Servlet vs Go Handler
*   **课 2: Gin 框架入门**
    *   Gin 优势, Router 路由组, 参数解析 (Uri, Query, Form)
    *   **Java 对比：** Spring MVC vs Gin Router
*   **课 3: 模型绑定与验证**
    *   `ShouldBindJSON`, `ShouldBindQuery`
    *   Go-playground/validator 校验库集成
*   **课 4: 中间件 (Middleware) 深度解析**
    *   洋葱模型, `Next()` 与 `Abort()`
    *   实战：编写耗时统计中间件、CORS 中间件
    *   **Java 对比：** Spring AOP/Filter vs Gin Middleware
*   **课 5: 配置管理与 Viper**
    *   使用 Viper 读取 YAML/Env 配置
    *   配置的热加载机制
*   **课 6: 日志管理与 Zap**
    *   结构化日志的重要性
    *   Zap 高性能日志库集成与配置 (Rotation, Level)
*   **课 7: JWT 认证实战**
    *   JWT 原理, 生成与解析 Token
    *   编写 AuthMiddleware 实现路由保护
*   **课 8: 优雅重启与 Swagger 文档**
    *   `context` 在优雅关机中的应用
    *   Swag 自动生成 API 文档
*   **🏆 综合实践任务 3: 简易电商用户中心 API**
    *   功能：用户注册/登录 (JWT)、个人信息修改、头像上传。
    *   要求：集成 Viper、Zap，输出 Swagger 文档，标准分层结构 (Controller/Service/Model)。

## 模块 4: 数据持久化与 GORM (7 节课)
> **焦点：** 数据操作与模型映射，强调 Go 中的 SQL 最佳实践。

*   **课 1: GORM 连接与配置**
    *   MySQL 驱动, 连接池配置 (`SetMaxOpenConns`)
    *   **Java 对比：** Hibernate/MyBatis vs GORM
*   **课 2: 模型定义与自动迁移**
    *   Model Struct, GORM Tags, `AutoMigrate`
    *   逻辑删除 (`gorm.DeletedAt`)
*   **课 3: CRUD 核心操作**
    *   Create (批量插入), First vs Find, Updates
    *   零值更新问题与解决方案 (使用 Map 或 指针)
*   **课 4: 高级查询与钩子 (Hooks)**
    *   `Where`, `Joins`, `Preload` (预加载解决 N+1)
    *   BeforeCreate/AfterSave 钩子函数的使用
*   **课 5: 事务管理 (Transaction)**
    *   闭包事务模式 `Transaction(func...)`
    *   嵌套事务与手动控制
*   **课 6: 原生 SQL 与 SQL Builder**
    *   `Raw`, `Exec` 使用场景
    *   何时应该放弃 ORM 使用原生 SQL (性能场景)
*   **课 7: GORM 单元测试**
    *   使用 `go-sqlmock` 模拟数据库连接
    *   编写不依赖真实 DB 的 Service 层测试
*   **🏆 综合实践任务 4: 博客系统核心数据层**
    *   功能：实现文章与标签的多对多关系、文章评论的一对多关系。
    *   要求：包含复杂的预加载查询，完整的事务控制（如删除文章同时删除关联数据）。

## 模块 5: gRPC 与 RPC 通信 (8 节课)
> **焦点：** 掌握强类型、高性能的微服务通信协议。

*   **课 1: RPC 理论与 Protobuf**
    *   RPC vs REST, Protocol Buffers 语法 (`.proto`)
    *   **Java 对比：** Serialization (Java Native) vs Protobuf
*   **课 2: Go 代码生成与 gRPC 存根**
    *   `protoc` 编译器安装与插件 (`protoc-gen-go`, `protoc-gen-go-grpc`)
    *   生成的 `.pb.go` 文件解读
*   **课 3: Unary RPC (简单模式) 实现**
    *   服务端 Handler 实现, 客户端连接 (`grpc.Dial`)
    *   Context 在 RPC 中的超时传递
*   **课 4: Streaming RPC (流模式) 实战**
    *   服务端流, 客户端流, 双向流 (聊天室案例)
    *   流模式下的错误处理与 EOF
*   **课 5: gRPC 拦截器 (Interceptor)**
    *   Unary 与 Stream 拦截器
    *   实战：实现 RPC 日志记录与 Panic 捕获
*   **课 6: Metadata 元数据传递**
    *   Metadata 读写 (类比 HTTP Header)
    *   基于 Token 的 RPC 认证实现
*   **课 7: gRPC 错误处理与状态码**
    *   `status` 包与 `codes`
    *   自定义错误详情 (Error Details)
*   **课 8: gRPC 与 Gateway**
    *   `grpc-gateway` 简介：同时提供 HTTP 和 gRPC 接口
    *   Docker 部署 gRPC 服务
*   **🏆 综合实践任务 5: 分布式计算服务**
    *   功能：客户端发送大量数据流，服务端进行实时计算并流式返回结果。
    *   要求：实现双向流，自定义认证拦截器。

## 模块 6: go-zero 微服务架构与生态集成 (8 节课)
> **焦点：** 站在巨人的肩膀上，使用工业级框架 go-zero 串联分布式生态。**本模块提供预配置好的 Docker-Compose 环境包，聚焦于“使用”而非“运维”。**

*   **课 1: 微服务架构与 go-zero 全景**
    *   单体 vs 微服务, go-zero 架构图 (极简依赖)
    *   `goctl` 神器安装与代码生成理念
*   **课 2: API 服务工程化**
    *   `.api` 定义语言, 生成 API 代码
    *   集成 Logic 层, Error 处理模型 (`xerr`)
*   **课 3: RPC 服务与 zRPC**
    *   `.proto` 结合 `goctl` 生成 RPC 服务
    *   实战：API 服务调用 RPC 服务 (服务拆分)
*   **课 4: 服务发现与注册中心 (Etcd)**
    *   Etcd 核心概念 (Key-Value, Lease)
    *   **实战：** 配置 go-zero 自动注册到 Etcd, 观察服务发现过程
*   **课 5: 数据库集群与缓存系统**
    *   `sqlx` 包装器与 `Model` 生成
    *   **核心特性：** 内置的缓存一致性模式 (Cache-Aside Pattern) 讲解与配置
    *   集成 MySQL/PostgreSQL (二选一演示)
*   **课 6: 消息队列与异步处理**
    *   集成 Kafka/RabbitMQ (使用 go-queue)
    *   实战：订单创建后的异步消息通知
*   **课 7: 可观测性 (监控与链路追踪)**
    *   **监控：** 配置 Prometheus 抓取 go-zero 指标, Grafana 看板展示
    *   **追踪：** 集成 Jaeger/Zipkin, 查看 API -> RPC -> DB 的全链路耗时
*   **课 8: 网关与 K8s 部署入门**
    *   API Gateway 概念 (APISIX/Envoy) 路由到 go-zero
    *   编写 Dockerfile, 生成 K8s Deployment/Service YAML (`goctl kube`)
*   **🏆 综合实践任务 6: 完整的微服务电商系统**
    *   **架构：** Order-API -> Order-RPC -> User-RPC / Product-RPC
    *   **技术栈：** go-zero, Etcd, MySQL, Redis, Prometheus.
    *   **要求：** 使用 `goctl` 生成大部分代码，手写业务逻辑；演示高并发下的服务限流与熔断（go-zero 内置功能）。
