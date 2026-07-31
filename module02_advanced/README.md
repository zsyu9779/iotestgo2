# Module 02：Go 进阶——抽象、并发与工程化

Module 02 面向已完成 Module 01 的学员，以四个连续 Block 建立接口与错误、Goroutine 与 Channel、Context 与并发安全、Testing 与 Reflection 能力，最后完成可取消的并发日志分析器。

> 教学说明：固定数据、短暂 Sleep、打印输出和内存 Channel 用于课堂理解，不等同于生产实现。生产代码还需完整的资源管理、日志、监控和部署约束。

> 教师答案披露：当前仓库包含课堂 Solution、作业教师答案、测评答案和讲师材料。发布学员仓库时只复制允许公开的 Demo 与 `homework/file_scanner/student_pack/`，不要把本仓库直接当作无答案学员包。

## 学习结果

完成本模块后，你能够：

- 使用隐式接口、方法集、类型断言和错误包装表达行为契约；
- 用 Goroutine、WaitGroup、Channel 和明确的关闭责任组织 worker pool；
- 使用 Context 实现取消和超时，并等待并发任务完整退出；
- 使用 Mutex、RWMutex、Atomic、Once 和 race detector 处理共享状态；
- 编写表格驱动测试和 Benchmark，并安全使用 Reflection；
- 构建一个可取消、可测试、无竞态的并发流水线。

## 真实日程

| 时间 | 内容 | 学习产出 |
|---|---|---|
| 09:30–09:50 | 环境检查、Entry Quiz | 能运行 Module 02 验收命令 |
| 09:50–10:40 | Block 1：接口、错误与恢复边界 | 完成接口与错误链练习 |
| 10:40–10:50 | 短休息 | — |
| 10:50–12:00 | Block 2：Goroutine 与 Channel | 完成保持顺序的 worker pool |
| 12:00–13:00 | 午休 | — |
| 13:00–14:05 | Block 3：Context 与并发安全 | 完成可取消且无竞态的计数器 |
| 14:05–14:50 | Block 4：Testing 与 Reflection | 完成表格测试和反射练习 |
| 14:50–15:00 | 短休息 | — |
| 15:00–15:40 | 并发日志分析器综合 Lab | 通过分阶段项目验收 |
| 15:40–16:00 | Homework 启动、Exit Quiz | 跑通 Starter RED 和提交流程 |

净课堂时间为 310 分钟，学员动手时间为 180 分钟。

## 课前

从仓库根目录运行：

```bash
go version
go env GOMOD
make module02-demo-contracts
make module02-lab-01
```

完成 [Entry Quiz](assessments/entry_quiz.md)，它只用于调整课堂节奏。

## 课中：四个 Block

| 阶段 | 导航 | Starter 练习命令 | Solution 验收 |
|---|---|---|---|
| Block 1 | [接口、错误与恢复边界](blocks/01_interfaces_errors/README.md) | `go test -tags=exercise ./module02_advanced/blocks/01_interfaces_errors/lab/starter` | `make module02-lab-01` |
| Block 2 | [Goroutine 与 Channel](blocks/02_goroutines_channels/README.md) | `go test -tags=exercise ./module02_advanced/blocks/02_goroutines_channels/lab/starter` | `make module02-lab-02` |
| Block 3 | [Context 与并发安全](blocks/03_context_concurrency/README.md) | `go test -tags=exercise -race ./module02_advanced/blocks/03_context_concurrency/lab/starter` | `make module02-lab-03` |
| Block 4 | [Testing 与 Reflection](blocks/04_testing_reflection/README.md) | `go test -tags=exercise ./module02_advanced/blocks/04_testing_reflection/lab/starter` | `make module02-lab-04` |
| 综合 Lab | [并发日志分析器](integrated_lab/log_analyzer/README.md) | `go test -tags=exercise -race ./module02_advanced/integrated_lab/log_analyzer/starter` | `make module02-integrated-lab` |

讲师入口是 [instructor/RUNBOOK.md](instructor/RUNBOOK.md)，Demo 精确步骤见 [instructor/DEMO_NOTES.md](instructor/DEMO_NOTES.md)。

## 课后作业

[可取消的并发文件扫描器](homework/file_scanner/student_pack/README.md) 是独立 Go Module，只使用标准库。进入 `student_pack` 后运行：

```bash
make grade
```

初始测试因 `ErrNotImplemented` 失败是预期 RED。教师答案使用：

```bash
make module02-homework-solution
```

## 测评与 Bonus

- [Exit Quiz](assessments/exit_quiz.md)
- [Bonus 总览](bonus/README.md)
- [受控失败案例](teaching_failures/README.md)

OS、File I/O、Runtime、标准库工具、Embed、Generate、pprof 和 DeepEqual 不占用核心 310 分钟。

## 仓库级验收

```bash
make module02-verify
make module02-demo-contracts
make module02-teaching-failures
make module02-audit
```

`module02-verify` 检查格式、Vet、普通测试、race、构建和教师 Homework Solution。故意 panic、死锁、竞态和 runtime fatal 的程序只由 `module02-teaching-failures` 隔离运行。
