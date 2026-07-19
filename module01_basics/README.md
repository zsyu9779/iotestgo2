# Module 01：从 Java 到 Go 的一日工作坊

Module 01 面向已能阅读 Java 的学员，用四个连续 Block 建立 Go 语法、集合、建模、函数与测试的基础，再通过 Scorebook 综合实验组合这些能力。Task Manager 是课后一周作业，课堂最后 20 分钟只做启动和验收演示。

> **教师答案披露：**当前 GitHub 仓库是教师仓库，包含所有课堂 `solution/`、Task Manager `teacher/solution/`、测验答案和讲师材料。学员仓库应由教师按课程进度单独发布；请不要把这个 GitHub 仓库当作无答案的学员包。

## 学习结果

完成本模块后，你能够：

- 用 `go run`、`go test`、`go vet` 和 `gofmt` 完成最小开发反馈循环。
- 说明 Array 的值语义、Slice 的共享行为、Map 的 comma-ok 查询以及 byte/rune 的区别。
- 用 Struct、指针、方法和值副本维护小型业务模型的不变量。
- 传递函数值，用闭包保存配置，用 `defer` 表达返回前的收尾动作。
- 从失败的 Go 测试定位当前行为差距，并用小步修改完成 Scorebook。

## 真实日程

| 时间 | 内容 | 学习产出 |
|---|---|---|
| 09:30–09:50 | 开场、环境检查、Entry Quiz | Go 工具链可用，讲师掌握班级起点 |
| 09:50–10:35 | Block 1：Go Basics | 完成带输入校验的成绩等级计算器 |
| 10:35–10:45 | 短休息 | — |
| 10:45–12:00 | Block 2：Collections | 完成中英文文本统计和词频汇总 |
| 12:00–13:00 | 午休 | — |
| 13:00–14:05 | Block 3：Modeling | 完成 Scorebook 的核心模型和状态修改 |
| 14:05–14:50 | Block 4：Functions & Testing | 增加统计函数、defer 使用和基础测试 |
| 14:50–15:00 | 短休息 | — |
| 15:00–15:40 | Scorebook 综合 Lab | 组合集合、建模、方法和测试 |
| 15:40–16:00 | Task Manager 作业启动、Exit Quiz | 跑通 Starter、本地验收和 Git 提交流程 |

授课时间为 09:30–12:00 和 13:00–16:00；午休不计入课堂。扣除两次 10 分钟短休息后，净教学与实操时间是 310 分钟。

## 课前：确认反馈环境

1. 在仓库根目录运行 `go version` 和 `go env GOMOD`，确认 Go 可用且 `GOMOD` 指向本仓库。
2. 运行 `make module01-lab-01`，确认可编译并执行 Go 测试。
3. 在不查资料的情况下完成 [Entry Quiz](assessments/entry_quiz.md)；它用来调整课堂节奏，不计作业分。

## 课中：按 Block 完成 RED–GREEN–REFACTOR

每个练习都从仓库根目录运行。带 `-tags=exercise` 的 Starter 命令是学员实操入口；课前的失败是预期 RED，完成实现后应变为 GREEN。Solution 命令用于讲师验收和对照。

| 阶段 | 导航 | Starter 练习命令 | Solution 验收命令 |
|---|---|---|---|
| Block 1：Go Basics | [Block](blocks/01_go_basics/README.md) · [Lab 指南](blocks/01_go_basics/lab/README.md) | `go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter` | `make module01-lab-01` |
| Block 2：Collections | [Block](blocks/02_collections/README.md) · [Lab 指南](blocks/02_collections/lab/README.md) | `go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter` | `make module01-lab-02` |
| Block 3：Modeling | [Block](blocks/03_modeling/README.md) · [Lab 指南](blocks/03_modeling/lab/README.md) | `go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter` | `make module01-lab-03` |
| Block 4：Functions & Testing | [Block](blocks/04_functions_testing/README.md) · [Lab 指南](blocks/04_functions_testing/lab/README.md) | `go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter` | `make module01-lab-04` |
| Scorebook 综合 Lab | [实验与分阶段检查点](integrated_lab/scorebook/README.md) | `go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter` | `make module01-integrated-lab` |

讲师 Demo 的精确步骤在各 Block 页面中；所有 Demo 也可从根目录依次运行：

```bash
go run ./module01_basics/blocks/01_go_basics/demo/01_hello
go run ./module01_basics/blocks/01_go_basics/demo/02_vars_types
go run ./module01_basics/blocks/01_go_basics/demo/03_control_funcs
go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices
go run ./module01_basics/blocks/02_collections/demo/05_maps_strings
go run ./module01_basics/blocks/03_modeling/demo/06_pointers
go run ./module01_basics/blocks/03_modeling/demo/07_structs_methods
go run ./module01_basics/blocks/04_functions_testing/demo/09_advanced_functions
go test ./module01_basics/bonus/function_patterns
```

课堂结束前完成 [Exit Quiz](assessments/exit_quiz.md)，把不确定的题号记入课后复习清单。

## 课后：独立迁移到 Task Manager

Task Manager 是课后一周作业，课堂内不实现答案。先阅读 [学员作业说明](homework/task_manager/student_pack/README.md)，再在独立学员包内完成行为：

```bash
cd module01_basics/homework/task_manager/student_pack
make grade
```

未完成的 Starter 应在验收脚本的测试步骤因 `ErrNotImplemented` 失败；这是作业起点，不是环境故障。完成后同一命令应依次通过 gofmt、Vet、测试和构建。讲师在仓库根目录用下列命令验证教师答案：

```bash
make module01-homework-solution
```

详细作业分数见 [Task Manager 评分表](homework/task_manager/teacher/RUBRIC.md)；当前 GitHub 仓库中的 `teacher/` 不得复制到学员仓库。

## Bonus：核心课程之后再探索

Bonus 不占用当天的 310 分钟核心课程：

- [Bonus 总览与运行命令](bonus/README.md)
- [数据结构示例](bonus/data_structures/main.go)
- [Generics 入门](bonus/generics/main.go)
- [range 变量、取地址与闭包陷阱](bonus/dark_corners/range/main.go)
- [Map 深层行为](bonus/dark_corners/map/main.go) 与 [String/UTF-8 深层行为](bonus/dark_corners/string/main.go)
- [函数配置模式](bonus/function_patterns/configuration_patterns.md)、[柯里化场景](bonus/function_patterns/curry_best_practice_test.go) 与 [Builder/Functional Options 对比](bonus/function_patterns/patterns_comparison_test.go)

## 仓库级验收

从仓库根目录运行：

```bash
make module01-verify
```

该命令检查 Module 01 的 gofmt，对根 Go Module 中的 Demo、课堂包和 Solution 执行 Vet 与测试，并通过统一脚本验证 Task Manager 教师答案。`student_pack` 是独立 Go Module，其故意未完成的测试不进入根验收。
