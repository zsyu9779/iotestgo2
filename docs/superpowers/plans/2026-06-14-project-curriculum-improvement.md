# Undergraduate Go Backend Course Improvement Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn `iotestgo2` from a useful classroom demo collection into a repeatable undergraduate Go backend course project with reliable setup, runnable demos, clear teacher materials, and minimum verification loops.

**Architecture:** Keep the current six-module teaching structure. Add a thin course-delivery layer around it: root setup guide, Makefile, environment examples, lesson indexes, generated-code/binary hygiene, and focused tests for each capstone project. Do not rewrite the whole repository into a production monorepo; preserve small per-lesson examples because that is what makes the project teachable.

**Tech Stack:** Go, Gin, GORM, gRPC, grpc-gateway, go-zero, Docker Compose, Markdown, shell scripts, `go test`, `go test -race`, `grpcurl`, `curl`.

---

## Scope Decision

This plan covers the project-wide course delivery layer. It deliberately does not implement the deeper gRPC/go-zero redesign or enterprise extension track; those are handled in:

- `docs/superpowers/plans/2026-06-14-grpc-gozero-deep-improvement.md`
- `docs/superpowers/plans/2026-06-14-enterprise-backend-extension-track.md`

## Current Findings Driving This Plan

- The project target audience is computer science undergraduates. Java comparison should stay as a bridge to students' existing coursework, not as the course identity.
- `go.mod` currently uses `go 1.25.0`; local `go1.20.6` cannot parse it, so environment setup must be explicit.
- `docs/` is ignored by `.gitignore`, yet course documents live there. That blocks normal version control of the most valuable teaching artifacts.
- Some generated or compiled artifacts are mixed into source directories, especially executable binaries inside `module05_grpc/08_grpc_gateway/proto`.
- There is no root `README.md`, `Makefile`, or single "how to run the course" guide.
- Module 01/02 have strong teacher docs; Module 03-06 mostly rely on README and comments.
- Tests exist, but there is no consistent command that tells a teacher "the classroom demos are healthy enough."

## Target File Structure

**Create:**
- `README.md` - root course overview, target students, prerequisites, module map, quickstart.
- `COURSE_RUNBOOK.md` - teacher-facing runbook: before class, per-module commands, troubleshooting.
- `.env.example` - shared local environment defaults for MySQL, Redis, Etcd, JWT secrets.
- `Makefile` - stable commands for setup, tests, formatting, proto generation, selected demos.
- `docs/module03_web_gin_lesson_plan.md` - teacher plan for Gin module.
- `docs/module04_gorm_lesson_plan.md` - teacher plan for GORM module.
- `docs/module05_grpc_lesson_plan.md` - teacher plan for gRPC module.
- `docs/module06_gozero_lesson_plan.md` - teacher plan for go-zero module.
- `docs/course_quality_checklist.md` - pre-class verification checklist.

**Modify:**
- `.gitignore` - stop ignoring all `docs`; ignore generated binaries and local runtime artifacts.
- `go.mod` - standardize the course Go version policy.
- `Golang_Backend_Training_Syllabus.md` - clarify undergraduate audience and Java-as-reference framing.
- `module01_basics/README.md` - include `10_generics_intro`.
- Module README files - add runnable commands, prerequisites, and "teaching simplification" notes.

**Remove or Regenerate:**
- Remove compiled Mach-O binaries from `module05_grpc/08_grpc_gateway/proto/`.
- Keep generated `.pb.go` files if they are intentionally committed for students without protoc.

---

### Task 1: Fix Documentation Version Control and Source Hygiene

**Files:**
- Modify: `.gitignore`
- Delete: compiled binaries under `module05_grpc/08_grpc_gateway/proto/`
- Verify: `git status --short --ignored docs`

- [ ] **Step 1: Update `.gitignore`**

Replace the broad `docs` ignore rule with targeted local artifact ignores:

```gitignore
# IDE / local tooling
.idea/
.trae/

# Build outputs
bin/
*.test
*.out
coverage.out

# Local compiled demo binaries. Source directories should contain .go/.proto/.md only.
module05_grpc/08_grpc_gateway/proto/client
module05_grpc/08_grpc_gateway/proto/server
module05_grpc/08_grpc_gateway/proto/userapi
module05_grpc/08_grpc_gateway/proto/userrpc
module05_grpc/08_grpc_gateway/proto/user-rpc
module05_grpc/08_grpc_gateway/proto/order-api
module05_grpc/08_grpc_gateway/proto/order-rpc
module05_grpc/08_grpc_gateway/proto/04_etcd_discovery
module05_grpc/08_grpc_gateway/proto/05_mysql_cache
module05_grpc/08_grpc_gateway/proto/06_message_queue
module05_grpc/08_grpc_gateway/proto/07_observability
module05_grpc/08_grpc_gateway/proto/08_k8s_deploy

# Local environment
.env
```

- [ ] **Step 2: Delete compiled binaries**

Run:

```bash
rm module05_grpc/08_grpc_gateway/proto/client
rm module05_grpc/08_grpc_gateway/proto/server
rm module05_grpc/08_grpc_gateway/proto/userapi
rm module05_grpc/08_grpc_gateway/proto/userrpc
rm module05_grpc/08_grpc_gateway/proto/user-rpc
rm module05_grpc/08_grpc_gateway/proto/order-api
rm module05_grpc/08_grpc_gateway/proto/order-rpc
rm module05_grpc/08_grpc_gateway/proto/04_etcd_discovery
rm module05_grpc/08_grpc_gateway/proto/05_mysql_cache
rm module05_grpc/08_grpc_gateway/proto/06_message_queue
rm module05_grpc/08_grpc_gateway/proto/07_observability
rm module05_grpc/08_grpc_gateway/proto/08_k8s_deploy
```

Expected: commands remove only Mach-O executable files. They must not remove `.go`, `.proto`, `.sh`, or `.md` files.

- [ ] **Step 3: Verify docs are visible to Git**

Run:

```bash
git status --short --ignored docs
```

Expected: `docs/*.md`, `docs/*.html`, and `docs/superpowers/plans/*.md` are no longer shown with `!!`.

- [ ] **Step 4: Commit**

```bash
git add .gitignore docs
git add -u module05_grpc/08_grpc_gateway/proto
git commit -m "chore: clean course docs and generated artifacts"
```

---

### Task 2: Standardize Go Toolchain Policy

**Files:**
- Modify: `go.mod`
- Create: `.env.example`
- Create: `README.md`
- Create: `COURSE_RUNBOOK.md`

- [ ] **Step 1: Decide and encode the course Go version**

For a 2026 undergraduate classroom, use the installed lab-machine version consistently. If the teaching machines can install current Go, the original target was:

```go
module iotestgo

go 1.25

toolchain go1.25.0
```

If the school lab machines are pinned below Go 1.25, set `go 1.23` and downgrade dependencies only after testing. Do not leave `go 1.25.0` without a matching `toolchain` directive, because older Go versions fail before they can explain the required upgrade.

Implementation note after verification on Go 1.25.11: `go mod tidy` normalizes this repository to `go 1.25.0` and removes `toolchain go1.25.0`. The final project guidance therefore treats Go 1.25.x as the classroom requirement and records the exact `go.mod` shape produced by the current Go toolchain.

- [ ] **Step 2: Add `.env.example`**

```dotenv
MYSQL_DSN=root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local
ECOMMERCE_MYSQL_DSN=root:root123@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local
REDIS_ADDR=127.0.0.1:6379
ETCD_ENDPOINTS=127.0.0.1:2379
JWT_SECRET=classroom-demo-secret-change-me
GRPC_AUTH_TOKEN=valid-token-12345
```

- [ ] **Step 3: Create root `README.md`**

```markdown
# iotestgo2

面向计算机专业本科生的 Go 后端课程演示项目。

## 课程定位

本项目用于课堂演示和课后练习，不是生产项目模板。Java 对比只用于连接学生已有课程经验，核心目标是建立 Go 后端开发的概念地图和动手能力。

## 模块

| 模块 | 内容 | 运行依赖 |
|---|---|---|
| module01_basics | Go 基础、指针、结构体、泛型、CLI 项目 | Go |
| module02_advanced | 接口、错误、并发、测试、文件 IO、反射 | Go |
| module03_web_gin | Gin、JWT、中间件、httptest、用户中心 | Go |
| module04_gorm | GORM、事务、Preload、Raw SQL、博客 API | Go + MySQL |
| module05_grpc | Protobuf、gRPC、stream、metadata、gateway | Go + protoc + grpcurl |
| module06_gozero | go-zero、API/RPC、Etcd、缓存、MQ、观测、部署 | Go + Docker |

## 快速开始

```bash
go version
make test-basic
make run-user-center
```

## 教学边界

部分示例刻意简化：明文密码、内存存储、硬编码 token、模拟消息队列和模拟 RPC 都只用于解释概念。正式工程需要替换为安全配置、持久化存储、真实 RPC client、真实消息系统和可观测性方案。
```

- [ ] **Step 4: Create `COURSE_RUNBOOK.md`**

```markdown
# Course Runbook

## 课前 10 分钟检查

```bash
go version
make fmt-check
make test-basic
```

## 外部依赖

- Module 04 需要 MySQL。
- Module 05 需要 `protoc`、`protoc-gen-go`、`protoc-gen-go-grpc`，Gateway 需要 `protoc-gen-grpc-gateway`。
- Module 06 的电商综合项目需要 Docker Compose 启动 MySQL、Redis、Etcd、Prometheus、Grafana。

## 课堂运行顺序

1. 基础语法：`make run-basic-hello`
2. 并发：`make run-log-analyzer`
3. Gin 用户中心：`make run-user-center`
4. GORM 博客 API：`make run-blog-api`
5. gRPC 双向流：`make run-compute-server`，另开终端 `make run-compute-client`
6. go-zero 电商链路：`make ecommerce-up` 后按 README 启动服务。

## 现场故障处理

- `invalid go version`: 检查 Go 版本是否符合 `go.mod` 与 `toolchain`。
- `connection refused`: 先确认对应 server 是否启动，端口是否被占用。
- `protoc: command not found`: 使用课程安装页安装 `protoc` 和 Go 插件。
- MySQL 连接失败：检查 Docker 容器、端口、DSN。
```

- [ ] **Step 5: Verify**

Run:

```bash
go version
go env GOVERSION GOMOD
```

Expected: Go version matches the course README. If the local environment still uses Go 1.20.6, the runbook must explicitly say this machine needs upgrade before full verification.

- [ ] **Step 6: Commit**

```bash
git add go.mod .env.example README.md COURSE_RUNBOOK.md
git commit -m "docs: add course setup and toolchain policy"
```

---

### Task 3: Add Root Makefile for Repeatable Classroom Commands

**Files:**
- Create: `Makefile`

- [ ] **Step 1: Create `Makefile`**

```makefile
.PHONY: help fmt-check test-basic test-race run-basic-hello run-task-manager run-log-analyzer run-user-center run-blog-api run-compute-server run-compute-client ecommerce-up ecommerce-down

help:
	@echo "Targets:"
	@echo "  fmt-check           Check Go formatting"
	@echo "  test-basic          Run tests that do not require external services"
	@echo "  test-race           Run selected race detector tests"
	@echo "  run-basic-hello     Run module01 hello"
	@echo "  run-log-analyzer    Run module02 log analyzer"
	@echo "  run-user-center     Run Gin user center"
	@echo "  run-blog-api        Run GORM blog API"
	@echo "  run-compute-server  Run gRPC compute server"
	@echo "  run-compute-client  Run gRPC compute client"
	@echo "  ecommerce-up        Start ecommerce infra"
	@echo "  ecommerce-down      Stop ecommerce infra"

fmt-check:
	@test -z "$$(gofmt -l $$(find . -name '*.go' -not -path './.git/*'))"

test-basic:
	go test ./module01_basics/... ./module02_advanced/... ./module03_web_gin/01_net_basics ./module03_web_gin/07_testing_httptest ./module03_web_gin/project_user_center/internal/service

test-race:
	go test -race ./module02_advanced/project_log_analyzer

run-basic-hello:
	go run ./module01_basics/01_hello

run-task-manager:
	go run ./module01_basics/project_task_manager

run-log-analyzer:
	go run ./module02_advanced/project_log_analyzer

run-user-center:
	go run ./module03_web_gin/project_user_center

run-blog-api:
	go run ./module04_gorm/project_blog_api

run-compute-server:
	go run ./module05_grpc/project_distributed_compute/server

run-compute-client:
	go run ./module05_grpc/project_distributed_compute/client

ecommerce-up:
	cd module06_gozero/project_ecommerce && docker compose up -d

ecommerce-down:
	cd module06_gozero/project_ecommerce && docker compose down
```

- [ ] **Step 2: Verify make help**

Run:

```bash
make help
```

Expected: target list prints without invoking Go.

- [ ] **Step 3: Verify formatting target**

Run:

```bash
make fmt-check
```

Expected: exits 0 after all Go files are formatted. If it prints files, run `gofmt -w` on those files and rerun.

- [ ] **Step 4: Commit**

```bash
git add Makefile
git commit -m "chore: add classroom make targets"
```

---

### Task 4: Align Audience Framing Across Course Documents

**Files:**
- Modify: `Golang_Backend_Training_Syllabus.md`
- Modify: `module01_basics/README.md`
- Modify: `module02_advanced/README.md`
- Modify: `module03_web_gin/README.md`
- Modify: `module04_gorm/README.md`
- Modify: `module05_grpc/README.md`
- Modify: `module06_gozero/README.md`

- [ ] **Step 1: Update syllabus overview**

Replace the opening metadata with:

```markdown
## 课程概览
*   **目标学生：** 计算机专业本科生，已学习过至少一门编程语言；学校课程以 Java 为主，因此部分章节使用 Java 对比帮助理解。
*   **课时安排：** 总计 46 课时（每节课 1-1.5 小时）
*   **教学风格：** 高密度、快节奏、概念对比、重实战
*   **核心目标：** 建立 Go 后端开发的完整认知地图，掌握 Go 语言特性、Web API、数据库访问、RPC 通信和微服务工程的基本实践。
```

- [ ] **Step 2: Add teaching simplification note to every module README**

Add this paragraph near the top of each module README:

```markdown
> 教学说明：本模块代码优先服务课堂理解，部分实现会使用内存数据、固定 token、简化错误处理或本地模拟组件。讲课时应明确这些是教学简化，不是生产写法。
```

- [ ] **Step 3: Fix Module 01 README module list**

Add:

```markdown
### 10_generics_intro/
- **main.go**: 泛型入门
- 学习内容：类型参数、泛型函数、泛型结构体、约束的基本概念
```

- [ ] **Step 4: Verify wording**

Run:

```bash
rg -n "Java 工程师|转 Go|训练营|针对 Java 开发者" README.md Golang_Backend_Training_Syllabus.md module*/README.md docs/*.md
```

Expected: no result except historical analysis files where the phrase is explicitly described as an old title.

- [ ] **Step 5: Commit**

```bash
git add Golang_Backend_Training_Syllabus.md module*/README.md README.md
git commit -m "docs: align course audience with undergraduate teaching"
```

---

### Task 5: Complete Teacher Lesson Plans for Modules 03-06

**Files:**
- Create: `docs/module03_web_gin_lesson_plan.md`
- Create: `docs/module04_gorm_lesson_plan.md`
- Create: `docs/module05_grpc_lesson_plan.md`
- Create: `docs/module06_gozero_lesson_plan.md`

- [ ] **Step 1: Create Module 03 lesson plan**

Use this structure:

```markdown
# Module 03: Gin Web 开发 - 教师备课教案

**适用对象**: 已掌握 Go 基础语法和基本并发概念的本科生
**总课时**: 预计 8 课时
**教学目标**: 让学生理解 HTTP 服务、中间件、认证、配置、日志和 Web 项目分层。

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
go run ./module03_web_gin/01_net_basics -mode server -protocol http
```

**讲解重点**:
- TCP/UDP 与 HTTP 的层级关系
- `net/http` 的 Handler 接口
- 为什么框架本质上仍然站在标准库之上

**练习**:
- 给 HTTP handler 增加 `/healthz` 路由，返回 `{"status":"ok"}`。
```

Then continue for lessons 2-8 using the existing README topics.

- [ ] **Step 2: Create Module 04 lesson plan**

Include exact demo commands:

```bash
docker compose -f module06_gozero/project_ecommerce/docker-compose.yml up -d mysql
go run ./module04_gorm/01_setup
go test -v ./module04_gorm/07_testing_mysql
```

Add explicit warning that hard-coded DSN is classroom-only and `.env.example` is preferred.

- [ ] **Step 3: Create Module 05 lesson plan**

Include a "generated code map":

```markdown
| 文件 | 角色 | 是否手写 |
|---|---|---|
| `*.proto` | 接口定义 | 手写 |
| `gen.sh` | 生成命令 | 手写 |
| `*.pb.go` | message 序列化代码 | 生成 |
| `*_grpc.pb.go` | client/server stub | 生成 |
| `server/main.go` | 服务实现 | 手写 |
| `client/main.go` | 调用演示 | 手写 |
```

- [ ] **Step 4: Create Module 06 lesson plan**

Split it into two teaching tracks:

```markdown
## Track A: Concept Demo
使用当前单文件示例解释 API、RPC、Etcd、缓存、MQ、可观测性。

## Track B: Framework Practice
使用 goctl 生成标准目录，演示真实 go-zero 工程结构。
```

- [ ] **Step 5: Verify docs are complete**

Run:

```bash
rg -n "适用对象|教学目标|演示命令|练习|教学简化" docs/module03_web_gin_lesson_plan.md docs/module04_gorm_lesson_plan.md docs/module05_grpc_lesson_plan.md docs/module06_gozero_lesson_plan.md
```

Expected: every file contains these teaching anchors.

- [ ] **Step 6: Commit**

```bash
git add docs/module03_web_gin_lesson_plan.md docs/module04_gorm_lesson_plan.md docs/module05_grpc_lesson_plan.md docs/module06_gozero_lesson_plan.md
git commit -m "docs: add lesson plans for backend modules"
```

---

### Task 6: Establish Minimum Verification for Capstone Projects

**Files:**
- Modify: `module01_basics/project_task_manager/main.go`
- Modify: `module01_basics/project_task_manager/task_manager_test.go`
- Modify: `module02_advanced/project_log_analyzer/main.go`
- Create: `module02_advanced/project_log_analyzer/pipeline_test.go`
- Create: `module03_web_gin/project_user_center/internal/handler/user_handler_test.go`
- Create: `module04_gorm/project_blog_api/internal/service/post_service_test.go`

- [ ] **Step 1: Refactor Task Manager to return values**

Add a pure method:

```go
func (tm *TaskManager) Snapshot() []Task {
	result := make([]Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		result = append(result, *task)
	}
	return result
}
```

- [ ] **Step 2: Test Task Manager snapshot**

```go
func TestTaskManagerSnapshotIsCopy(t *testing.T) {
	tm := NewTaskManager()
	tm.Add("read grpc chapter")

	snapshot := tm.Snapshot()
	if len(snapshot) != 1 {
		t.Fatalf("expected 1 task, got %d", len(snapshot))
	}
	snapshot[0].Completed = true

	if tm.tasks[0].Completed {
		t.Fatal("snapshot mutation changed internal task state")
	}
}
```

- [ ] **Step 3: Add deterministic log analyzer path**

Add:

```go
func CountErrors(entries []LogEntry, numProcessors int) int {
	logsCh := make(chan LogEntry, len(entries))
	errorsCh := make(chan LogEntry, len(entries))

	var wg sync.WaitGroup
	for i := 1; i <= numProcessors; i++ {
		wg.Add(1)
		go LogProcessor(i, logsCh, errorsCh, &wg)
	}

	for _, entry := range entries {
		logsCh <- entry
	}
	close(logsCh)
	wg.Wait()
	close(errorsCh)

	count := 0
	for range errorsCh {
		count++
	}
	return count
}
```

- [ ] **Step 4: Test deterministic log analyzer**

```go
func TestCountErrors(t *testing.T) {
	entries := []LogEntry{
		{ID: 1, Level: "INFO"},
		{ID: 2, Level: "ERROR"},
		{ID: 3, Level: "WARN"},
		{ID: 4, Level: "ERROR"},
	}
	got := CountErrors(entries, 2)
	if got != 2 {
		t.Fatalf("expected 2 errors, got %d", got)
	}
}
```

- [ ] **Step 5: Add Web handler test**

Test `POST /register` and `POST /login` with `httptest`, using the in-memory repository.

- [ ] **Step 6: Add Blog service test with sqlmock**

Create a real assertion-based test for `CreatePostWithComment` using `sqlmock`, replacing the current print-only style where needed.

- [ ] **Step 7: Run verification**

```bash
make test-basic
make test-race
```

Expected: all selected tests pass. If external dependencies are missing, the command must skip those packages rather than fail unpredictably.

- [ ] **Step 8: Commit**

```bash
git add module01_basics/project_task_manager module02_advanced/project_log_analyzer module03_web_gin/project_user_center module04_gorm/project_blog_api Makefile
git commit -m "test: add capstone verification coverage"
```

---

## Self-Review

- **Spec coverage:** Covers project-wide setup, docs visibility, audience framing, course runbook, lesson plans, and minimum tests. gRPC/go-zero deep redesign and enterprise extensions are intentionally split into separate plans.
- **占位检查:** No unresolved marker instructions remain.
- **Type consistency:** Go snippets use existing `TaskManager`, `LogEntry`, `LogProcessor`, and package-local test style.
- **Execution risk:** `go.mod` requires a deliberate Go version decision. Do that before running `go test`.
