# Module 03: Web 开发与 Gin 框架

本模块包含 Go Web 开发的基础知识和 Gin 框架的使用，涵盖从基础 HTTP 到完整 API 开发的各个方面。

## 目录结构

### 01_http_net/
- **main.go**: 标准库 net/http 基础
- 学习内容：http.Handler、http.HandlerFunc、基本路由处理

### 01_net_basics/
- **http_handlers.go**: HTTP 处理器和中间件
- **protocol_line.go**: HTTP 协议基础
- **server_test.go**: 服务器测试
- 学习内容：中间件编写、HTTP 协议理解、测试编写

### 02_gin_intro/
- **main.go**: Gin 框架入门
- 学习内容：Gin 路由、中间件、JSON 响应

### 03_binding_viper/
- **main.go**: 数据绑定和配置管理
- 学习内容：Gin 绑定、Viper 配置管理、环境变量

### 04_middleware_jwt/
- **main.go**: 中间件和 JWT 认证
- 学习内容：自定义中间件、JWT 令牌、认证流程

### 05_logging_zap/
- **main.go**: 日志记录
- 学习内容：Zap 日志库、结构化日志、日志级别

### 06_api_design/
- **main.go**: API 设计
- 学习内容：RESTful API 设计、版本控制、错误处理

### 07_testing_httptest/
- **handler.go**: HTTP 处理器
- **handler_test.go**: HTTP 测试
- 学习内容：httptest 包、集成测试、模拟请求

### 08_perf_context/
- **main.go**: 性能优化和上下文
- 学习内容：性能分析、上下文传递、优化技巧

### project_user_center/
完整的用户中心项目，包含：
- **internal/handler/user_handler.go**: 用户处理器
- **internal/middleware/middleware.go**: 中间件
- **internal/model/user.go**: 用户模型
- **internal/repository/user_repo.go**: 数据存储
- **internal/service/user_service.go**: 业务逻辑
- **internal/service/user_service_test.go**: 服务测试
- **pkg/utils/jwt.go**: JWT 工具
- **main.go**: 主程序

## 学习目标

1. 掌握标准库 net/http 的使用
2. 熟练使用 Gin Web 框架
3. 理解中间件设计和实现
4. 掌握数据绑定和配置管理
5. 实现 JWT 认证和授权
6. 使用结构化日志记录
7. 设计良好的 RESTful API
8. 编写全面的测试用例
9. 完成一个完整的用户管理系统

## 运行方式

每个目录下的程序都可以通过以下命令运行：
```bash
cd 目录名
go run main.go
```

对于项目运行：
```bash
cd project_user_center/
go run main.go
```

运行测试：
```bash
cd 07_testing_httptest/
go test -v

cd project_user_center/internal/service/
go test -v
```

## 依赖安装

需要安装 Gin 和其他依赖：
```bash
go mod tidy
```