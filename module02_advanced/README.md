# Module 02：Go 进阶——抽象、并发与工程化

Module 02 面向已经完成 Module 01、能够阅读 Java 基础代码的学员。课程主线是：

> 接口与错误建模 → Goroutine 与 Channel → Context 与并发安全 → Testing 与 Reflection → 综合流水线。

代码中的固定数据、Sleep、打印输出和内存对象主要用于课堂理解，不能直接视为生产实现。生产代码还需要完整的错误处理、资源管理、日志、监控和部署约束。

## 学习结果

完成本模块后，你能够：

- 使用隐式接口、类型断言和错误包装表达行为契约；
- 使用 Goroutine、WaitGroup、Channel 和 close 组织并发任务；
- 使用 Context 实现取消、超时和请求生命周期传播；
- 使用 Mutex、RWMutex、Atomic 和 race detector 处理共享状态；
- 编写表格驱动测试、Benchmark，并使用 Reflection 读取和修改值；
- 完成一个可取消、可测试、无竞态的并发日志分析流水线。

## 课程入口

- [讲师 Runbook](instructor/RUNBOOK.md)
- [讲师 Demo Notes](instructor/DEMO_NOTES.md)
- [课堂评分表](instructor/RUBRIC.md)
- [Entry Quiz](assessments/entry_quiz.md)
- [Exit Quiz](assessments/exit_quiz.md)
- [Homework](homework/README.md)

## 全日路径

| 阶段 | 主题 | 入口 |
|---|---|---|
| Block 1 | 接口、错误与 defer | [Block README](blocks/01_interfaces_errors/README.md) |
| Block 2 | Goroutine 与 Channel | [Block README](blocks/02_goroutines_channels/README.md) |
| Block 3 | Context 与并发安全 | [Block README](blocks/03_context_concurrency/README.md) |
| Block 4 | Testing 与 Reflection | [Block README](blocks/04_testing_reflection/README.md) |
| 综合 Lab | 并发日志分析器 | [Block README](blocks/05_integrated_lab/README.md) |

## 常用验收命令

从仓库根目录运行：

```bash
make module02-demo-contracts
make module02-integrated-lab
make module02-verify
```

Module 02 的默认验收不会运行故意 panic、死锁或竞态的教学失败案例。故意失败只能通过显式教学命令调用。

## 课后作业

完成 [可取消的并发文件扫描器](homework/README.md)。作业要求使用标准库、Context、Worker、表格测试和 race detector，不直接复制综合 Lab 的实现。

## Bonus

以下内容不占用核心 310 分钟：

- [Bonus 总览](bonus/README.md)
- `embed`：编译时嵌入静态资源；
- `go:generate`：代码生成与外部工具依赖；
- OS 命令、信号处理和跨平台边界；
- 文件 I/O、runtime 控制和标准库工具；
- pprof、DeepEqual 和并发 map dark corners。
