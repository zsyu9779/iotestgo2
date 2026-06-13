# Module 05: gRPC 开发 - 教师备课教案

**适用对象**: 已理解 HTTP API、接口分层和基础并发的本科生
**总课时**: 预计 8 课时
**教学目标**: 让学生理解 Protobuf 契约、gRPC 调用模型、流式通信、metadata、拦截器、错误语义和 Gateway 暴露方式。

> 教学简化: 本模块使用本地进程和简化 token 展示 RPC 概念。正式工程需要服务发现、TLS、真实认证、超时治理、可观测性和版本兼容策略。

## 生成代码地图

| 文件 | 角色 | 是否手写 |
|---|---|---|
| `*.proto` | 接口定义 | 手写 |
| `gen.sh` | 生成命令 | 手写 |
| `*.pb.go` | message 序列化代码 | 生成 |
| `*_grpc.pb.go` | client/server stub | 生成 |
| `server/main.go` | 服务实现 | 手写 |
| `client/main.go` | 调用演示 | 手写 |

## 每节课固定结构

1. 一句话定位
2. 运行一个最小 Demo
3. 解释核心机制
4. 修改一个小功能
5. 指出一个生产边界

## 第 1 课: Protobuf 基础

**源码路径**: `module05_grpc/01_protobuf_basics`
**演示命令**:

```bash
go run ./module05_grpc/01_protobuf_basics
```

**讲解重点**:
- `.proto`、message、enum、field number
- Protobuf 与 JSON 的差异
- 生成代码为什么不适合手改

**练习**:
- 给 message 增加一个可选字段，重新生成并观察 Go 结构体变化。

**生产边界**:
- field number 一旦发布不能随意复用，接口演进需要兼容性意识。

## 第 2 课: 代码生成

**源码路径**: `module05_grpc/02_codegen`
**演示命令**:

```bash
(cd module05_grpc/02_codegen && bash ./gen.sh)
go run ./module05_grpc/02_codegen
```

**讲解重点**:
- `protoc` 与 Go 插件的职责
- `pb.go` 和 `grpc.pb.go` 分别生成什么
- client stub 和 server interface 的含义

**练习**:
- 修改服务方法名，观察编译错误如何指向调用方和实现方。

**生产边界**:
- 真实项目通常把 proto 契约独立管理，并使用 CI 检查生成代码是否同步。

## 第 3 课: Unary RPC

**源码路径**: `module05_grpc/03_unary_rpc`
**演示命令**:

```bash
go run ./module05_grpc/03_unary_rpc/server
go run ./module05_grpc/03_unary_rpc/client
```

**讲解重点**:
- 一次请求一次响应的 RPC 模型
- server 注册、client Dial、context 超时
- RPC 错误和普通 Go error 的关系

**练习**:
- 给请求增加 `name` 长度校验，并返回合适错误。

**生产边界**:
- 真实客户端需要连接复用、超时、重试和负载均衡策略。

## 第 4 课: Streaming RPC

**源码路径**: `module05_grpc/04_streaming_rpc`
**演示命令**:

```bash
go run ./module05_grpc/04_streaming_rpc/server
go run ./module05_grpc/04_streaming_rpc/client
```

**讲解重点**:
- Server streaming、Client streaming、Bidirectional streaming
- `io.EOF` 与流结束
- 流式 RPC 和 channel 思维的类比

**练习**:
- 给双向聊天消息增加序号字段，并在客户端打印。

**生产边界**:
- 流式连接要处理背压、取消、心跳和异常断开。

## 第 5 课: Interceptor

**源码路径**: `module05_grpc/05_interceptors`
**演示命令**:

```bash
go run ./module05_grpc/05_interceptors/server
go run ./module05_grpc/05_interceptors/client
```

**讲解重点**:
- Unary interceptor 和 Stream interceptor
- 日志、panic 恢复、耗时统计
- 横切逻辑为什么不应散落在 handler 中

**练习**:
- 在 interceptor 中打印 method name 和耗时。

**生产边界**:
- 真实系统还会把 tracing、metrics、auth、rate limit 放进统一拦截链。

## 第 6 课: Metadata 与认证

**源码路径**: `module05_grpc/06_metadata_auth`
**演示命令**:

```bash
go run ./module05_grpc/06_metadata_auth/server
go run ./module05_grpc/06_metadata_auth/client
```

**讲解重点**:
- Metadata 与 HTTP header 的类比
- client 如何附加 token
- server 如何读取并拒绝非法请求

**练习**:
- 增加一个 `x-request-id` metadata，并在服务端日志中打印。

**生产边界**:
- 课堂 token 是固定字符串，真实服务应结合 TLS、JWT/OAuth2 或内部身份体系。

## 第 7 课: 错误处理

**源码路径**: `module05_grpc/07_error_handling`
**演示命令**:

```bash
go run ./module05_grpc/07_error_handling/server
go run ./module05_grpc/07_error_handling/client
```

**讲解重点**:
- `status` 和 `codes`
- 客户端如何分辨 NotFound、InvalidArgument、Internal
- 错误详情和业务错误码的边界

**练习**:
- 为参数错误返回 `codes.InvalidArgument`，并在客户端做分支处理。

**生产边界**:
- 错误语义是契约的一部分，需要文档化并保持兼容。

## 第 8 课: Gateway 与分布式计算项目

**源码路径**: `module05_grpc/08_grpc_gateway`, `module05_grpc/project_distributed_compute`
**演示命令**:

```bash
make run-compute-server
make run-compute-client
```

**讲解重点**:
- gRPC-Gateway 如何把 HTTP 请求映射到 RPC
- 双协议暴露的适用场景
- 分布式计算项目里的流式输入、实时计算和流式输出

**练习**:
- 给计算请求增加一个任务名称字段，并贯穿 client/server 输出。

**生产边界**:
- Gateway 不是简单“自动 REST”，真实网关还涉及鉴权、限流、CORS、OpenAPI 和兼容性。
