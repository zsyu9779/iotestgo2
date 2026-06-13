# Module 03: Gin Web 开发 - 教师备课教案

**适用对象**: 已掌握 Go 基础语法和基本并发概念的本科生
**总课时**: 预计 8 课时
**教学目标**: 让学生理解 HTTP 服务、中间件、认证、配置、日志和 Web 项目分层。

> 教学简化: 本模块部分示例使用内存数据、固定 token、硬编码配置或简化错误响应。讲课时要明确这些是为了突出 Web 后端机制，不是生产写法。

## 每节课固定结构

1. 一句话定位
2. 运行一个最小 Demo
3. 解释核心机制
4. 修改一个小功能
5. 指出一个生产边界

## 第 1 课: 网络编程基础

**源码路径**: `module03_web_gin/01_net_basics`
**演示命令**:

```bash
go test -v ./module03_web_gin/01_net_basics
go run ./module03_web_gin/01_net_basics -mode=server -proto=http
```

**讲解重点**:
- TCP/UDP 与 HTTP 的层级关系
- `net/http` 的 Handler 接口
- 为什么框架本质上仍然站在标准库之上

**练习**:
- 给 HTTP handler 增加 `/healthz` 路由，返回 `{"status":"ok"}`。

**生产边界**:
- 真实服务需要超时、优雅关闭、限流和统一错误响应，课堂示例先聚焦协议路径。

## 第 2 课: Gin 入门与路由

**源码路径**: `module03_web_gin/02_gin_intro`
**演示命令**:

```bash
go run ./module03_web_gin/02_gin_intro
```

**讲解重点**:
- `gin.Default()`、路由注册、JSON 响应
- 路由参数、查询参数和请求上下文
- Gin 与 `net/http` 的关系

**练习**:
- 增加 `GET /users/:id`，返回路径参数和查询参数。

**生产边界**:
- 示例不包含统一响应结构、鉴权和日志追踪，真实 API 需要提前约定接口规范。

## 第 3 课: 数据绑定与配置

**源码路径**: `module03_web_gin/03_binding_viper`
**演示命令**:

```bash
go run ./module03_web_gin/03_binding_viper
```

**讲解重点**:
- `ShouldBindJSON` 的绑定和校验
- Viper 读取配置的基本流程
- 配置默认值、环境变量和本地文件的职责边界

**练习**:
- 给注册请求增加邮箱字段校验，并把端口改成从配置读取。

**生产边界**:
- 真实项目不要把敏感配置提交到仓库，应使用 `.env`、密钥系统或部署平台配置。

## 第 4 课: 中间件与 JWT

**源码路径**: `module03_web_gin/04_middleware_jwt`
**演示命令**:

```bash
go run ./module03_web_gin/04_middleware_jwt
```

**讲解重点**:
- Gin middleware 的洋葱模型
- JWT 的 claim、签名和过期时间
- 认证信息如何从请求进入业务 handler

**练习**:
- 给受保护接口增加角色字段检查。

**生产边界**:
- 示例中的密钥和 token 策略仅用于课堂，真实系统必须使用安全密钥、刷新机制和撤销策略。

## 第 5 课: 结构化日志

**源码路径**: `module03_web_gin/05_logging_zap`
**演示命令**:

```bash
go run ./module03_web_gin/05_logging_zap
```

**讲解重点**:
- 为什么结构化日志比普通字符串更适合服务端
- request id、耗时、状态码、错误字段
- 日志中间件和业务日志的分工

**练习**:
- 在响应中加入 request id，并在日志中打印同一个字段。

**生产边界**:
- 真实系统需要日志采集、脱敏、采样和按环境区分日志级别。

## 第 6 课: API 设计规范

**源码路径**: `module03_web_gin/06_api_design`
**演示命令**:

```bash
go run ./module03_web_gin/06_api_design
```

**讲解重点**:
- RESTful 资源建模
- 路由分组和版本管理
- Swagger/OpenAPI 文档的作用

**练习**:
- 为文章资源补充 `GET /api/v1/posts/:id` 和 `DELETE /api/v1/posts/:id`。

**生产边界**:
- API 文档必须和实现同步，真实团队需要自动生成、契约评审和兼容性策略。

## 第 7 课: HTTP 测试

**源码路径**: `module03_web_gin/07_testing_httptest`
**演示命令**:

```bash
go test -v ./module03_web_gin/07_testing_httptest
```

**讲解重点**:
- `httptest.NewRecorder` 和构造请求
- Handler 级测试和端到端测试的区别
- 表格驱动测试覆盖多输入

**练习**:
- 为错误请求增加一个测试用例，断言状态码和响应 body。

**生产边界**:
- 真实服务还需要 repository mock、认证上下文、超时和错误路径覆盖。

## 第 8 课: 用户中心综合项目

**源码路径**: `module03_web_gin/project_user_center`
**演示命令**:

```bash
make run-user-center
```

**讲解重点**:
- handler、service、repository 的分层
- 注册、登录、鉴权的请求链路
- 小项目如何从单文件示例过渡到分层结构

**练习**:
- 为用户资料接口增加一个仅登录可访问的 `/profile` 路由。

**生产边界**:
- 项目使用内存仓储和简化密码处理，真实用户系统必须使用哈希密码、数据库事务、审计日志和风控策略。
