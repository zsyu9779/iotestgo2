# Enterprise Backend Extension Track Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add optional enterprise backend extension materials that broaden undergraduate students' view beyond Go syntax and framework usage, without overloading the required course.

**Architecture:** Create a separate extension track under `docs/enterprise_backend_extensions/` with optional topics, short demos, selection rules, and classroom-ready exercises. Each topic is independent and can be taught in 30-60 minutes depending on student absorption speed. The track references existing Go modules when possible and adds small companion demos only when a concept cannot be taught well from current code.

**Tech Stack:** Markdown, Go standard library, Gin, GORM, gRPC, go-zero, Docker Compose, Prometheus, OpenTelemetry concepts, Redis/MySQL concepts, CI/CD concepts, security basics.

---

## Design Principles

1. **Optional means optional.** These topics should not become hidden required homework.
2. **Use current project as anchor.** Every enterprise concept should point back to an existing module or capstone project.
3. **One concept, one demo.** Avoid giant "enterprise platform" examples.
4. **Teach tradeoffs, not tool worship.** Students should learn why a pattern exists, when it helps, and when it is unnecessary.
5. **Prefer diagrams and failure stories.** Enterprise backend topics often become abstract; anchor them in request flow, failure flow, and data flow.

## Extension Topic Matrix

| Topic | When to Teach | Anchor Module | Demo Form | Difficulty |
|---|---|---|---|---|
| API design and compatibility | After Gin | Module 03 | versioned REST routes + response envelope | Low |
| Authentication and password security | After JWT | Module 03 | bcrypt + JWT secret config | Low |
| Database schema and transactions | After GORM | Module 04 | order transaction sketch | Medium |
| Cache consistency | After GORM or go-zero cache | Module 06 | Cache-Aside timeline | Medium |
| Message queues and eventual consistency | After go-zero MQ | Module 06 | order.created event flow | Medium |
| Observability | After go-zero observability | Module 06 | metrics/logs/traces map | Medium |
| Resilience | After gRPC/go-zero | Module 05/06 | timeout/retry/circuit breaker lab | Medium |
| API Gateway and edge concerns | After gRPC Gateway | Module 05/06 | gateway responsibilities diagram | Medium |
| Deployment and release flow | End of course | Module 06 | Docker -> Compose -> K8s -> CI | Medium |
| Performance and profiling | After testing/benchmark | Module 02 | benchmark + pprof story | Medium |
| Domain modeling and service boundaries | End of course | Module 03/06 | user/order/product boundary exercise | High |
| Data governance and privacy | End of course | Module 03/04 | password, PII, audit fields | Medium |

---

## Target File Structure

**Create:**
- `docs/enterprise_backend_extensions/README.md`
- `docs/enterprise_backend_extensions/topic_01_api_design.md`
- `docs/enterprise_backend_extensions/topic_02_auth_security.md`
- `docs/enterprise_backend_extensions/topic_03_database_transactions.md`
- `docs/enterprise_backend_extensions/topic_04_cache_consistency.md`
- `docs/enterprise_backend_extensions/topic_05_message_queue.md`
- `docs/enterprise_backend_extensions/topic_06_observability.md`
- `docs/enterprise_backend_extensions/topic_07_resilience.md`
- `docs/enterprise_backend_extensions/topic_08_gateway_edge.md`
- `docs/enterprise_backend_extensions/topic_09_deployment_release.md`
- `docs/enterprise_backend_extensions/topic_10_performance_pprof.md`
- `docs/enterprise_backend_extensions/topic_11_service_boundaries.md`
- `docs/enterprise_backend_extensions/topic_12_data_privacy.md`
- `docs/enterprise_backend_extensions/selection_guide.md`
- `docs/enterprise_backend_extensions/diagrams/README.md`

**Optionally Create Small Demos:**
- `module07_enterprise_extensions/01_api_compatibility/main.go`
- `module07_enterprise_extensions/02_resilience/main.go`
- `module07_enterprise_extensions/03_pprof/main.go`

This `module07_enterprise_extensions` directory is optional. Create it only after the docs are written and reviewed.

---

### Task 1: Create Extension Track Landing Page

**Files:**
- Create: `docs/enterprise_backend_extensions/README.md`

- [ ] **Step 1: Write landing page**

```markdown
# 企业级后端开发视野拓展专题

本目录是 Go 后端课程的可选拓展材料。它不替代主线课程，也不是每个班都必须讲完。教师可以根据学生吸收速度，从中选择 2-5 个专题穿插到 Module 03-06 或期末项目阶段。

## 使用方式

| 节奏 | 推荐选择 |
|---|---|
| 学生刚能跑通代码 | API 设计、密码安全、数据库事务 |
| 学生能理解 Web + DB | 缓存一致性、消息队列、可观测性 |
| 学生能理解 RPC/微服务 | 韧性设计、网关边界、服务边界 |
| 期末项目准备 | 部署发布、性能分析、数据隐私 |

## 专题列表

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

## 教学原则

- 每次只讲一个企业问题，不堆工具名。
- 每个专题都要回答：为什么需要、什么时候不需要、学生能做一个什么小实验。
- 如果学生基础不稳，优先回到主线代码，不强行拓展。
```

- [ ] **Step 2: Verify**

```bash
rg -n "专题列表|教学原则|使用方式" docs/enterprise_backend_extensions/README.md
```

Expected: all three anchors exist.

- [ ] **Step 3: Commit**

```bash
git add docs/enterprise_backend_extensions/README.md
git commit -m "docs: add enterprise backend extension track"
```

---

### Task 2: Add Selection Guide for Different Student Speeds

**Files:**
- Create: `docs/enterprise_backend_extensions/selection_guide.md`

- [ ] **Step 1: Write guide**

```markdown
# 拓展专题选择指南

## A 档：学生吸收较慢，只能补 2 个专题

1. API 设计与兼容性
2. 认证、密码与密钥管理

理由：这两个专题和日常 Web 项目最贴近，能立即提升作业质量。

## B 档：学生能顺利完成 Gin + GORM，可补 4 个专题

1. API 设计与兼容性
2. 数据库建模与事务边界
3. 缓存一致性
4. 可观测性

理由：这条线从单体服务质量出发，不要求学生完全理解微服务。

## C 档：学生能理解 gRPC + go-zero，可补 6 个专题

1. 消息队列与最终一致性
2. 韧性设计
3. API Gateway 与边缘层职责
4. 部署、发布与回滚
5. 性能分析与容量意识
6. 服务边界与领域建模

理由：这条线面向微服务系统思维，适合作为期末项目的高级要求。

## 课堂判断信号

| 现象 | 教师选择 |
|---|---|
| 学生频繁卡在语法/运行命令 | 不拓展，回到主线 |
| 学生能解释 handler/service/repository | 加 API 设计、认证 |
| 学生能解释事务和 Preload | 加缓存、MQ |
| 学生能解释 API -> RPC 调用链 | 加韧性、网关、服务边界 |
| 学生主动问线上系统怎么排查问题 | 加可观测性、性能、发布 |
```

- [ ] **Step 2: Commit**

```bash
git add docs/enterprise_backend_extensions/selection_guide.md
git commit -m "docs: add extension topic selection guide"
```

---

### Task 3: Create API Design and Compatibility Topic

**Files:**
- Create: `docs/enterprise_backend_extensions/topic_01_api_design.md`

- [ ] **Step 1: Write topic**

```markdown
# Topic 01: API 设计与兼容性

## 适合插入位置

Module 03 Gin API 设计之后。

## 要解决的问题

学生通常会把接口理解为“能返回 JSON 就行”。企业项目更关心接口是否可演进、可排查、可被前端/客户端稳定使用。

## 核心概念

- 路由版本：`/api/v1/users`
- 统一错误响应：`code/message/request_id`
- 幂等性：重复请求是否安全
- 兼容性：新增字段通常安全，删除/改名字段通常危险
- 分页：`page/page_size` 或 `cursor/limit`

## 课堂 Demo

基于 `module03_web_gin/06_api_design/main.go`，把一个随意接口改成版本化接口：

```json
{
  "code": "OK",
  "message": "success",
  "request_id": "req-20260614-001",
  "data": {
    "user_id": 1,
    "username": "gopher"
  }
}
```

## 练习

让学生设计 `GET /api/v1/posts` 的响应，要求包含：

- 分页信息
- 文章列表
- 请求追踪 ID
- 空列表时的响应

## 讨论题

为什么“直接返回数据库模型”在小作业里很方便，但在长期维护中容易出问题？
```

- [ ] **Step 2: Commit**

```bash
git add docs/enterprise_backend_extensions/topic_01_api_design.md
git commit -m "docs: add api design extension topic"
```

---

### Task 4: Create Authentication and Security Topic

**Files:**
- Create: `docs/enterprise_backend_extensions/topic_02_auth_security.md`

- [ ] **Step 1: Write topic**

```markdown
# Topic 02: 认证、密码与密钥管理

## 适合插入位置

Module 03 JWT 中间件之后。

## 要解决的问题

当前课程项目中有明文密码和硬编码 JWT secret，这是教学简化。拓展课要让学生知道真实项目为什么不能这么做。

## 核心概念

- 密码不加密保存，而是保存不可逆 hash。
- JWT secret 不写死在代码里，从环境变量或配置系统读取。
- Access Token 有过期时间。
- 鉴权和认证不同：认证回答“你是谁”，鉴权回答“你能做什么”。
- 不要在日志里打印密码、token、身份证号、手机号等敏感信息。

## 课堂 Demo

把 `project_user_center` 的注册逻辑从明文密码改成 bcrypt hash：

```go
hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
if err != nil {
	return nil, err
}
user.Password = string(hashed)
```

登录时：

```go
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	return "", ErrInvalidCredentials
}
```

## 练习

让学生找出当前项目中所有硬编码 secret/token/password，并按“教学可接受 / 必须改造”分类。

## 讨论题

为什么生产系统里“忘记密码”通常不能把原密码发给用户？
```

- [ ] **Step 2: Commit**

```bash
git add docs/enterprise_backend_extensions/topic_02_auth_security.md
git commit -m "docs: add auth security extension topic"
```

---

### Task 5: Create Database and Transaction Topic

**Files:**
- Create: `docs/enterprise_backend_extensions/topic_03_database_transactions.md`

- [ ] **Step 1: Write topic**

```markdown
# Topic 03: 数据库建模与事务边界

## 适合插入位置

Module 04 GORM 事务之后。

## 要解决的问题

学生容易把事务理解成“出错就回滚”。企业项目更关心事务边界：哪些数据必须一起成功，哪些可以异步补偿。

## 核心概念

- 主键、唯一索引、外键、普通索引
- 事务 ACID 的课堂版解释
- 事务边界不应跨 HTTP 请求
- 长事务会占用连接和锁
- 订单创建通常需要订单表、订单项表、库存记录配合

## 课堂 Demo

基于 `module04_gorm/project_blog_api/internal/service/post_service.go` 的事务风格，讲订单事务：

1. 创建订单主表
2. 创建订单项
3. 扣减库存
4. 任一步失败则回滚

## 练习

让学生画出博客系统 `posts/comments/tags` 的表结构，并标注：

- 哪些字段需要索引
- 删除文章时哪些数据要一起处理
- 哪些操作应该放在事务内

## 讨论题

为什么“所有操作都放进一个大事务”不是好设计？
```

- [ ] **Step 2: Commit**

```bash
git add docs/enterprise_backend_extensions/topic_03_database_transactions.md
git commit -m "docs: add database transaction extension topic"
```

---

### Task 6: Create Cache, MQ, and Observability Topics

**Files:**
- Create: `docs/enterprise_backend_extensions/topic_04_cache_consistency.md`
- Create: `docs/enterprise_backend_extensions/topic_05_message_queue.md`
- Create: `docs/enterprise_backend_extensions/topic_06_observability.md`

- [ ] **Step 1: Write cache topic**

```markdown
# Topic 04: 缓存一致性

## 适合插入位置

Module 06 `05_mysql_cache` 之后。

## 核心问题

缓存不是为了“显得高级”，而是为了减少数据库压力和降低读取延迟。代价是数据一致性变复杂。

## 必讲流程

Cache-Aside 读：

1. 读缓存
2. 未命中读数据库
3. 写回缓存
4. 返回数据

Cache-Aside 写：

1. 更新数据库
2. 删除缓存
3. 下次读取时重新加载

## 练习

给学生三个场景，让他们判断该不该加缓存：

- 课程公告列表
- 用户登录密码校验
- 热门商品详情页
```

- [ ] **Step 2: Write MQ topic**

```markdown
# Topic 05: 消息队列与最终一致性

## 适合插入位置

Module 06 `06_message_queue` 之后。

## 核心问题

消息队列用于解耦、削峰、异步处理。它不能神奇地让系统更简单；它把同步复杂度换成了异步一致性复杂度。

## 订单事件例子

`order.created` 事件被三个消费者处理：

- 库存服务扣库存
- 通知服务发消息
- 数据分析服务记录事件

## 必讲风险

- 消息重复
- 消息丢失
- 消费失败重试
- 死信队列
- 幂等消费

## 练习

让学生设计 `order.created` 消息体，要求包含 `event_id`、`order_id`、`user_id`、`created_at`、`items`。
```

- [ ] **Step 3: Write observability topic**

```markdown
# Topic 06: 可观测性：日志、指标、链路追踪

## 适合插入位置

Module 06 `07_observability` 之后。

## 核心问题

上线后的系统不能靠猜。可观测性回答三类问题：

- 发生了什么：日志
- 现在健康吗：指标
- 慢在哪里：链路追踪

## 三大支柱

| 类型 | 适合回答 | 例子 |
|---|---|---|
| Logs | 单次事件细节 | 某个订单创建失败的错误 |
| Metrics | 趋势和告警 | QPS、P95、错误率 |
| Traces | 跨服务路径 | API -> RPC -> DB 花了多久 |

## 练习

给 `order-api` 的一次请求设计 5 个日志字段：

- request_id
- user_id
- order_id
- latency_ms
- error_code
```

- [ ] **Step 4: Commit**

```bash
git add docs/enterprise_backend_extensions/topic_04_cache_consistency.md docs/enterprise_backend_extensions/topic_05_message_queue.md docs/enterprise_backend_extensions/topic_06_observability.md
git commit -m "docs: add cache mq observability extension topics"
```

---

### Task 7: Create Resilience, Gateway, Deployment, and Performance Topics

**Files:**
- Create: `docs/enterprise_backend_extensions/topic_07_resilience.md`
- Create: `docs/enterprise_backend_extensions/topic_08_gateway_edge.md`
- Create: `docs/enterprise_backend_extensions/topic_09_deployment_release.md`
- Create: `docs/enterprise_backend_extensions/topic_10_performance_pprof.md`

- [ ] **Step 1: Write resilience topic**

```markdown
# Topic 07: 韧性设计：超时、重试、熔断、限流

## 适合插入位置

Module 05 metadata/interceptor 或 Module 06 go-zero resilience 之后。

## 核心问题

分布式系统一定会部分失败。韧性设计不是让失败消失，而是让失败被限制、被观察、被恢复。

## 四个概念

- Timeout：不要无限等。
- Retry：只对可重试错误重试，必须有次数上限。
- Circuit Breaker：下游持续失败时快速失败。
- Rate Limit：保护系统不被流量打穿。

## 练习

给 API -> RPC 调用设计策略：

- 超时：3 秒
- 重试：最多 2 次，只重试 `Unavailable`
- 熔断：连续失败率超过阈值后打开
- 限流：每秒最多 100 请求
```

- [ ] **Step 2: Write gateway topic**

```markdown
# Topic 08: API Gateway 与边缘层职责

## 适合插入位置

Module 05 grpc-gateway 或 Module 06 网关部署之后。

## 核心问题

Gateway 不是业务逻辑垃圾桶。它适合处理跨服务统一问题，不适合塞满订单、用户、商品规则。

## Gateway 适合做

- TLS 终止
- 鉴权前置
- 路由
- 限流
- 灰度
- 请求 ID 注入
- 基础审计

## Gateway 不适合做

- 订单价格计算
- 库存扣减
- 用户状态变更
- 数据库事务

## 练习

让学生把 10 个功能分到 Gateway / API Service / RPC Service 三层。
```

- [ ] **Step 3: Write deployment topic**

```markdown
# Topic 09: 部署、发布与回滚

## 适合插入位置

Module 06 K8s 部署之后。

## 核心问题

写完代码不是结束。企业后端还要考虑怎么构建、怎么发布、怎么回滚、怎么观察发布是否成功。

## 最小发布链路

1. Git commit
2. CI 运行测试
3. 构建 Docker image
4. 推送镜像仓库
5. 更新部署
6. 健康检查
7. 失败回滚

## 练习

让学生给 `user-center` 设计健康检查：

- `/healthz`: 进程是否活着
- `/readyz`: 数据库/Redis 是否可用
```

- [ ] **Step 4: Write performance topic**

```markdown
# Topic 10: 性能分析与容量意识

## 适合插入位置

Module 02 benchmark 或 Module 06 结束后。

## 核心问题

性能优化不是猜。先测量，再定位，再优化。

## 工具顺序

1. Benchmark：函数级性能
2. pprof：CPU/内存热点
3. 压测：接口级吞吐和延迟
4. 指标：线上持续观察

## 练习

比较字符串拼接：

```bash
go test -bench=. ./module01_basics/05_maps_strings
```

让学生解释为什么 `strings.Builder` 通常比循环 `+=` 更适合大量拼接。
```

- [ ] **Step 5: Commit**

```bash
git add docs/enterprise_backend_extensions/topic_07_resilience.md docs/enterprise_backend_extensions/topic_08_gateway_edge.md docs/enterprise_backend_extensions/topic_09_deployment_release.md docs/enterprise_backend_extensions/topic_10_performance_pprof.md
git commit -m "docs: add resilience gateway deployment performance topics"
```

---

### Task 8: Create Service Boundary and Data Privacy Topics

**Files:**
- Create: `docs/enterprise_backend_extensions/topic_11_service_boundaries.md`
- Create: `docs/enterprise_backend_extensions/topic_12_data_privacy.md`

- [ ] **Step 1: Write service boundary topic**

```markdown
# Topic 11: 服务边界与领域建模

## 适合插入位置

Module 06 电商综合项目之后。

## 核心问题

微服务不是“每张表一个服务”。服务边界应该围绕业务能力，而不是围绕文件夹或数据库表。

## 电商例子

- User Service：用户资料、状态、认证关联信息
- Product Service：商品信息、库存读写
- Order Service：订单生命周期、订单项、金额快照
- Payment Service：支付单、支付状态、回调

## 练习

给出 12 个功能，让学生分配服务边界：

1. 修改用户名
2. 查询商品价格
3. 创建订单
4. 扣减库存
5. 支付回调
6. 取消订单
7. 发送短信
8. 查询订单列表
9. 修改商品标题
10. 退款
11. 用户封禁
12. 统计每日销售额
```

- [ ] **Step 2: Write data privacy topic**

```markdown
# Topic 12: 数据隐私、审计与合规意识

## 适合插入位置

课程末尾或期末项目前。

## 核心问题

后端系统处理的不只是数据结构，也是用户数据。学生需要早一点建立隐私和审计意识。

## 必讲点

- 最小化收集：不需要就不收集
- 脱敏展示：手机号、邮箱、身份证号
- 敏感字段不进日志
- 管理员操作要审计
- 数据删除和数据备份存在冲突，需要制度设计

## 练习

指出下面日志的问题：

```text
login failed username=alice password=123456 token=eyJ...
```

要求学生改成：

```text
login failed username=alice reason=invalid_credentials request_id=req-001
```
```

- [ ] **Step 3: Commit**

```bash
git add docs/enterprise_backend_extensions/topic_11_service_boundaries.md docs/enterprise_backend_extensions/topic_12_data_privacy.md
git commit -m "docs: add service boundary and privacy topics"
```

---

### Task 9: Add Optional Demo Module Only After Docs Are Stable

**Files:**
- Create: `module07_enterprise_extensions/README.md`
- Create: `module07_enterprise_extensions/01_api_compatibility/main.go`
- Create: `module07_enterprise_extensions/02_resilience/main.go`
- Create: `module07_enterprise_extensions/03_pprof/main.go`

- [ ] **Step 1: Create module README**

```markdown
# Module 07: 企业级后端拓展演示

本模块是可选演示，不属于主线必讲内容。只有当学生已经能顺利理解 Module 03-06 时才使用。

## Demo

1. API compatibility: 统一响应、版本路由、request_id
2. Resilience: timeout、retry、rate limit 的最小模拟
3. pprof: CPU 热点与 benchmark 对照
```

- [ ] **Step 2: Add API compatibility demo**

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/api/v1/users/1", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{
			Code:      "OK",
			Message:   "success",
			RequestID: "req-" + time.Now().Format("20060102150405"),
			Data: map[string]interface{}{
				"user_id":  1,
				"username": "gopher",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	log.Println("listening on :8891")
	log.Fatal(http.ListenAndServe(":8891", nil))
}
```

- [ ] **Step 3: Add resilience demo**

Use standard library only:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func callDownstream(ctx context.Context, latency time.Duration) error {
	select {
	case <-time.After(latency):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	err := callDownstream(ctx, 500*time.Millisecond)
	fmt.Printf("downstream result: %v\n", err)
}
```

- [ ] **Step 4: Add pprof demo**

Create a CPU-heavy Fibonacci endpoint only for local demonstration and document that this is intentionally inefficient.

- [ ] **Step 5: Verify optional demos**

```bash
go run ./module07_enterprise_extensions/01_api_compatibility
go run ./module07_enterprise_extensions/02_resilience
```

Expected: first starts HTTP server, second prints `context deadline exceeded`.

- [ ] **Step 6: Commit**

```bash
git add module07_enterprise_extensions
git commit -m "feat: add optional enterprise extension demos"
```

---

## Self-Review

- **Spec coverage:** Covers optional extension strategy, selection by student speed, twelve enterprise backend topics, and optional demos.
- **占位检查:** No unresolved marker content remains.
- **Type consistency:** Topic numbering and filenames are consistent from README through task list.
- **Scope control:** The extension track is intentionally docs-first. Code demos are optional and isolated in `module07_enterprise_extensions` so they do not distract from the Go backend core modules.
