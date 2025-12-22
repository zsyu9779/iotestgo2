# Module 03: Gin Web 开发

本模块专注于使用 Gin 框架构建 Web 应用程序和 RESTful API。

## 目录结构

### 01_net_basics/
- **udp_client.go**, **udp_server.go**: Socket 编程 (UDP/TCP)
- **http_handlers.go**: HTTP 处理器
- 学习内容：Socket 通信原理, net/http 标准库, Handler 接口

### 02_gin_intro/
- **main.go**: Gin 框架入门
- 学习内容：Gin 引擎初始化、基本路由定义、JSON 响应

### 03_binding_viper/
- **main.go**: 数据绑定与配置
- 学习内容：ShouldBindJSON, Viper 配置文件读取与管理

### 04_middleware_jwt/
- **main.go**: 中间件与认证
- 学习内容：自定义中间件、JWT (JSON Web Token) 生成与解析

### 05_logging_zap/
- **main.go**: 日志系统
- 学习内容：集成 Uber Zap 高性能日志库、结构化日志

### 06_api_design/
- **main.go**: API 设计规范
- 学习内容：RESTful API 设计原则、路由分组、版本控制

### 07_testing_httptest/
- **handler.go**, **handler_test.go**: HTTP 测试
- 学习内容：httptest 包使用、模拟 HTTP 请求与响应测试

### 08_perf_context/
- **main.go**: 性能与上下文
- 学习内容：Gin Context 深入、性能优化技巧

### project_user_center/
- **internal/**, **pkg/**, **main.go**: 用户中心微服务
- 学习内容：综合项目实战，包含用户注册、登录、鉴权等完整功能，采用分层架构

## 学习目标

1. 理解 Socket 与 HTTP 协议基础
2. 熟练掌握 Gin 框架的核心功能
3. 能够处理请求数据绑定与验证
4. 掌握中间件机制与 JWT 认证
5. 学会使用结构化日志与配置管理
6. 能够设计符合 RESTful 规范的 API
7. 掌握 Web 服务的单元测试方法
8. 具备构建完整 Web 微服务的能力
