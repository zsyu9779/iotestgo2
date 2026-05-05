# Module 06: go-zero 微服务开发

本模块专注于使用 go-zero 框架构建生产级微服务应用。

## 目录结构

### 01_gozero_intro/
- go-zero 架构全景、极简依赖哲学、goctl 安装与理念

### 02_api_service/
- .api 定义语言 → goctl 生成 → Logic 层手写业务

### 03_rpc_service/
- .proto → goctl 生成 RPC → API 调用 RPC（第一个微服务拆分）

### 04_etcd_discovery/
- Key-Value + Lease、go-zero 自动注册、观察服务上下线

### 05_mysql_cache/
- sqlx 包装器、Model 生成、内置缓存一致性模式（Cache-Aside）

### 06_message_queue/
- go-queue 抽象、Kafka 集成、异步消息实战

### 07_observability/
- Prometheus 指标暴露、Grafana 看板、Jaeger 链路追踪

### 08_k8s_deploy/
- Dockerfile 编写、goctl kube 生成 YAML、Deployment/Service

### project_ecommerce/
- 电商微服务实战：Order-API → Order-RPC → User-RPC / Product-RPC

## 前置条件

需要准备 Docker-Compose 环境包（MySQL + Redis + Etcd + Prometheus），学员一键启动：

```bash
cd module06_gozero/project_ecommerce
docker-compose up -d
```

## 学习目标

1. 理解 go-zero 架构设计理念
2. 掌握 .api 定义语言与 goctl 代码生成
3. 熟练掌握 API 服务与 RPC 服务开发
4. 理解 Etcd 服务注册与发现机制
5. 掌握数据库缓存一致性模式
6. 能够集成消息队列实现异步处理
7. 具备微服务可观测性建设能力
8. 了解微服务容器化部署流程

## Java 对比

| 概念 | go-zero | Spring Cloud |
|------|---------|--------------|
| 服务定义 | .api 文件 | OpenAPI/Swagger |
| RPC | zRPC（基于 gRPC） | Feign / gRPC |
| 服务发现 | Etcd | Eureka / Nacos |
| 配置中心 | 本地文件 + K8s ConfigMap | Spring Cloud Config / Nacos |
| 网关 | go-zero Gateway | Spring Cloud Gateway |
| 熔断限流 | 内置 breaker / limiter | Resilience4j / Sentinel |
| ORM | sqlx | MyBatis / JPA |
