# Module 06: go-zero 微服务开发 - 教师备课教案

**适用对象**: 已掌握 Gin、GORM、gRPC 基本概念，并准备理解工程化微服务框架的本科生
**总课时**: 预计 8 课时
**教学目标**: 让学生理解 go-zero 的 API/RPC 代码生成、服务发现、缓存、消息、可观测性、部署和综合链路。

> 教学简化: 当前模块同时包含单文件概念示例和电商综合项目。讲课时应明确哪些用于解释框架概念，哪些用于接近真实工程结构；课堂示例不是完整生产脚手架。

## Track A: Concept Demo

使用当前单文件示例解释 API、RPC、Etcd、缓存、MQ、可观测性。

## Track B: Framework Practice

使用 goctl 生成标准目录，演示真实 go-zero 工程结构。

## 每节课固定结构

1. 一句话定位
2. 运行一个最小 Demo
3. 解释核心机制
4. 修改一个小功能
5. 指出一个生产边界

## 第 1 课: go-zero 全景

**源码路径**: `module06_gozero/01_gozero_intro`
**演示命令**:

```bash
go run ./module06_gozero/01_gozero_intro
```

**讲解重点**:
- go-zero 的 API、RPC、Model、Gateway 基本版图
- 为什么企业框架强调代码生成和约束
- 单文件示例与真实目录结构的区别

**练习**:
- 让学生画出“HTTP 请求进入 API，再调用 RPC”的链路图。

**生产边界**:
- 框架不能替代架构设计，服务拆分仍要从业务边界出发。

## 第 2 课: API 服务

**源码路径**: `module06_gozero/02_api_service`
**演示命令**:

```bash
go run ./module06_gozero/02_api_service
```

**讲解重点**:
- `.api` 文件表达接口契约
- Handler、Logic、Config 的职责
- goctl 生成代码后哪些文件应该手写

**练习**:
- 在 `.api` 中增加一个查询用户详情接口，并补充对应 logic。

**生产边界**:
- 真实 API 要配合鉴权、限流、参数校验、错误码和文档生成。

## 第 3 课: RPC 服务

**源码路径**: `module06_gozero/03_rpc_service`
**演示命令**:

```bash
go run ./module06_gozero/03_rpc_service/userrpc
go run ./module06_gozero/03_rpc_service/userapi
```

**讲解重点**:
- `.proto` 到 zRPC 服务的生成流程
- API 调 RPC 的调用边界
- RPC client 放在哪里、如何被 logic 使用

**练习**:
- 给 User RPC 增加按 ID 查询方法，并在 API 层暴露。

**生产边界**:
- 当前示例可用模拟调用解释链路，真实项目应使用生成 client、服务发现和超时配置。

## 第 4 课: Etcd 服务发现

**源码路径**: `module06_gozero/04_etcd_discovery`
**演示命令**:

```bash
go run ./module06_gozero/04_etcd_discovery
```

**讲解重点**:
- 服务注册、租约、心跳
- 客户端如何发现服务实例
- 服务上下线对调用方的影响

**练习**:
- 模拟服务下线，观察服务列表变化。

**生产边界**:
- 真实环境要考虑多实例、网络抖动、灰度发布和配置隔离。

## 第 5 课: MySQL 与缓存

**源码路径**: `module06_gozero/05_mysql_cache`
**演示命令**:

```bash
go run ./module06_gozero/05_mysql_cache
```

**讲解重点**:
- sqlx、Model 生成和缓存封装
- Cache-Aside 基本模式
- 缓存穿透、击穿、雪崩的概念入口

**练习**:
- 为用户查询增加未命中日志，观察缓存路径。

**生产边界**:
- 真实缓存需要 TTL、空值缓存、并发保护和一致性策略。

## 第 6 课: 消息队列

**源码路径**: `module06_gozero/06_message_queue`
**演示命令**:

```bash
go run ./module06_gozero/06_message_queue
```

**讲解重点**:
- 同步调用和异步消息的区别
- Producer、Consumer、Topic 的基本概念
- 订单创建后发消息的业务动机

**练习**:
- 增加一个“发送邮件通知”的模拟 consumer。

**生产边界**:
- 真实 MQ 要处理幂等、重试、死信、顺序性和消息积压。

## 第 7 课: 可观测性与部署

**源码路径**: `module06_gozero/07_observability`, `module06_gozero/08_k8s_deploy`
**演示命令**:

```bash
go run ./module06_gozero/07_observability
go run ./module06_gozero/08_k8s_deploy
```

**讲解重点**:
- metrics、logs、traces 三类信号
- Prometheus、Grafana、Jaeger 在链路中的位置
- Dockerfile、Deployment、Service 的基础角色

**练习**:
- 给一个业务接口增加计数指标，并说明 Grafana 应该如何展示。

**生产边界**:
- 部署不是只写 YAML，还包括配置、密钥、资源限制、健康检查和发布策略。

## 第 8 课: 电商综合项目

**源码路径**: `module06_gozero/project_ecommerce`
**演示命令**:

```bash
make ecommerce-up
go run ./module06_gozero/project_ecommerce/user-rpc
go run ./module06_gozero/project_ecommerce/order-rpc
go run ./module06_gozero/project_ecommerce/order-api
```

**讲解重点**:
- Order-API、Order-RPC、User-RPC 的服务关系
- API 请求如何经过课堂模拟的 RPC 调用链；数据库、缓存和消息作为本模块前面单独示例解释
- 从概念 Demo 过渡到综合工程的思维方式

**练习**:
- 为订单接口补充一个订单状态字段，并让 API 响应返回该字段。

**生产边界**:
- 综合项目仍是课堂骨架，真实电商系统还需要支付、库存一致性、审计、风控、压测和灾备。
