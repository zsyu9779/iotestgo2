# Module 01 讲师 Runbook（09:30–16:00）

本 Runbook 是 Module 01 的单一现场授课入口。课前同时打开 [Demo Notes](DEMO_NOTES.md)、[课堂评分语言](RUBRIC.md) 和 [测评答案](../assessments/answer_key.md)。所有命令默认从仓库根目录运行。

## 授课边界与时间算术

- 上午 09:30–12:00，下午 13:00–16:00，共 330 分钟。
- 午休 12:00–13:00 不计入上述 330 分钟。
- 短休息为 10:35–10:45 和 14:50–15:00，共 20 分钟。
- 净教学与实操时间：`150 + 180 - 20 = 310` 分钟。
- Task Manager 仅在 15:40–16:00 启动作业；不把实现、直播解答或额外项目塞入课堂主体。

## 全日日程（不合并任何时间行）

| 时间 | 内容 | 可观察产出 |
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

## 学员动手预算

下表只计入学员亲自操作工具链、编辑代码、运行测试或完成 Quiz 的分钟；观看 Demo、听讲、全班讨论和休息均不计入。

| Time box | Learner hands-on minutes |
|---|---:|
| Opening and Entry Quiz | 10 |
| Block 1 | 25 |
| Block 2 | 40 |
| Block 3 | 35 |
| Block 4 | 25 |
| Scorebook integrated lab | 40 |
| Homework kickoff and Exit Quiz | 10 |
| **Total** | **185 / 310 (59.7%)** |

`185 ÷ 310 = 0.59677…`，四舍五入为 59.7%。延误时先删除各段标明的可裁内容，不挪用这 185 分钟中的编码、测试和 Quiz 时间。

## 09:30–09:50：开场、环境检查、Entry Quiz

- **目标：**确认学员能在仓库根目录运行 Go，并识别班级最需要纠正的 Java-to-Go 默认假设。
- **讲师动作：**09:30–09:35 说明一日产出；09:35–09:40 投影 `go version`、`go env GOMOD` 和 `make module01-lab-01`；09:40–09:45 安静答 Quiz；09:45–09:50 按题号举手统计。
- **学员动作（10 分钟动手）：**独立运行三个环境命令 5 分钟，完成 [Entry Quiz](../assessments/entry_quiz.md) 5 分钟。
- **可观察检查点：**`GOMOD` 指向当前仓库且 Block 1 solution 测试 PASS；讲师记录每题错误人数。
- **常见延误：**学员不在仓库根目录、Go 不在 `PATH` 或 IDE 打开了错误目录。
- **可裁内容：**精确删除开场自我介绍和课程背景扩展；不删环境命令、Quiz 或错题统计。

## 09:50–10:35：Block 1 — Go Basics

- **目标：**使用条件、函数、多返回值和显式 `error` 实现有边界校验的 `Grade`。
- **讲师动作：**09:50–10:05 按 [Demo Notes](DEMO_NOTES.md#block-1go-basics15-分钟) 展示程序入口、类型推断、`for`、`switch` 和 `(value, error)`；10:05–10:30 巡视结对实现；10:30–10:35 用一个边界失败做复盘。
- **学员动作（25 分钟动手）：**在 `blocks/01_go_basics/lab/starter` 运行 RED，实现范围校验和 A–F 映射，重复运行全包测试。
- **可观察检查点：**`go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter` PASS，且学员能解释 60、90、-1 和 101 的路径。
- **常见延误：**从低到高写过宽分支，或返回临时文本错误而不是 `ErrScoreOutOfRange`。
- **可裁内容：**精确删除 `iota` 的扩展例子和 Bonus 表驱动测试讨论；保留错误返回与边界测试。

## 10:35–10:45：短休息

准时停止投影和编码，10:45 直接从 Array 副本对比开始。休息不是可裁内容，不用它补 Block 1。

## 10:45–12:00：Block 2 — Collections

- **目标：**通过中英文文本统计验证 Array 值语义、Slice 共享、Map comma-ok 与 UTF-8 byte/rune 行为。
- **讲师动作：**10:45–11:10 运行两个 Collections Demo 并先让学员预测输出；11:10–11:50 巡视 `Analyze`实现；11:50–12:00 对比一个中文字符的 byte 和 rune 计数。
- **学员动作（40 分钟动手）：**实现 `Analyze`，用 `strings.Fields`、`strings.ToLower` 和 `utf8.RuneCountInString` 完成计数，反复运行练习测试。
- **可观察检查点：**`go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter` PASS；`"Go go 你好"` 的结果为 12 bytes、8 runes、3 words 和 `go:2`。
- **常见延误：**把 `len(text)` 当作字符数，未初始化 Map，或自行删标点而改变了分词契约。
- **可裁内容：**精确删除 Slice 容量增长策略推演、Map 深层陷阱和 String benchmark；保留 Slice 共享、comma-ok 和 byte/rune。

## 12:00–13:00：午休

关闭上午问题停车场的现场讨论，只留下文字记录。13:00 按新的 65 分钟时间盒开始，不用午休补课。

## 13:00–14:05：Block 3 — Modeling

- **目标：**用 Struct、构造函数惯例、指针接收者和值快照建立保持不变量的 `Student`。
- **讲师动作：**13:00–13:20 用值传递与指针修改 Demo 对比 Java 引用；13:20–13:55 巡视建模练习；13:55–14:05 用“修改 Snapshot”的测试复盘接收者选择。
- **学员动作（35 分钟动手）：**实现 `New`、`Rename`、`UpdateScore` 和 `Snapshot`，先校验再修改状态，逐步使测试变绿。
- **可观察检查点：**`go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter` PASS；修改 Snapshot 不影响原对象，失败的 Rename/Update 不改原状态。
- **常见延误：**修改方法使用值接收者，或先赋值再校验，导致变更丢失或无效状态泄漏。
- **可裁内容：**精确删除 Struct embedding 对比和 nil pointer 扩展讨论；保留 Go 只有值传递、接收者选择和校验前置。

## 14:05–14:50：Block 4 — Functions & Testing

- **目标：**把函数作为值传递，用闭包捕获配置，用 `defer` 安排可观察的收尾顺序，并能阅读具名测试失败。
- **讲师动作：**14:05–14:20 只演示函数值、闭包、`defer` 与 `-run`；14:20–14:45 让学员依次处理两个具名测试；14:45–14:50 对照三个 audit 事件。
- **学员动作（25 分钟动手）：**先使 `TestFilterWithClosure` 通过，再使 `TestWithAuditRecordsEndAfterOperation` 通过，最后运行全包测试。
- **可观察检查点：**`go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter` PASS，audit 事件严格为 `start:average`、`operation`、`end:average`。
- **常见延误：**`Filter` 硬编码阈值，`AtLeast` 使用 `>`，或在 `operation` 之后才注册 `defer`。
- **可裁内容：**精确删除柯里化、函数组合、Functional Options 和多 `defer` LIFO 扩展；保留单一闭包、单一 `defer` 和失败输出阅读。

## 14:50–15:00：短休息

准时停止。15:00 所有学员回到 Scorebook Starter；休息不是可裁内容，不用它补 Block 4。

## 15:00–15:40：Scorebook 综合 Lab

- **目标：**在新的 Scorebook 问题中独立组合当天所学，不复制任何 Block solution。
- **讲师动作：**在 10、20、30、40 分钟节点只播报 [Scorebook Lab](../integrated_lab/scorebook/README.md) 当前检查点和停止规则；个别辅导时只询问“第一个失败断言是什么”。
- **学员动作（40 分钟动手）：**0–10 分钟做模型与 Add，10–20 分钟做 Find/UpdateScore，20–30 分钟做 Average，30–40 分钟做 CountByGrade 和全量测试。
- **可观察检查点：**每 10 分钟指定的具名测试输出可见；最终 `go test -tags=exercise ./module01_basics/integrated_lab/scorebook/starter` PASS，或学员保留了第一个未通过断言的完整输出。
- **常见延误：**无效 Add 也消耗 ID，返回值泄漏内部指针，或 Average 在转换 `float64` 之前先做整数除法。
- **可裁内容：**只可删除最后的全班 solution 讲解，并改为课后链接；不删任何一个 10 分钟编码检查点。

## 15:40–16:00：Task Manager 作业启动、Exit Quiz

- **目标：**让学员知道一周作业的范围、预期 RED、唯一验收入口和分支交付方式，不在现场实现 Task Manager。
- **讲师动作：**15:40–15:44 演示作业目标行为；15:44–15:49 指导学员跑 `make grade`；15:49–15:54 说明一周里程碑、开发分支和 Gitee 手动验收；15:54–15:59 主持 Exit Quiz；15:59–16:00 收取离场票。
- **学员动作（10 分钟动手）：**在 `student_pack` 运行 `make grade` 5 分钟并标记预期失败阶段，然后完成 [Exit Quiz](../assessments/exit_quiz.md) 5 分钟。
- **可观察检查点：**Starter 在 gofmt 和 Vet 后的测试步骤因 `ErrNotImplemented` 行为性失败，而不是编译、权限或依赖失败；每人交一个最不确定题号。
- **常见延误：**学员在仓库根目录运行作业 `make grade`，或讲师开始直播编写 `Add`。
- **可裁内容：**精确删除 CLI 所有命令的逐一现场操作和 Gitee UI 屏幕导览；保留 Starter RED、`make grade`、分支交付规则和 Exit Quiz。

## 课后收尾

1. 用 [answer key](../assessments/answer_key.md) 统计 Entry/Exit 题号，按诊断语言记录下次开场要回收的误解。
2. 运行 `make module01-verify`，确认课堂 Solution 与 Task Manager 教师答案都经过验收。
3. 发布时只复制 `homework/task_manager/student_pack/` 的内容，严格按 [发布检查清单](../homework/task_manager/teacher/RELEASE_CHECKLIST.md) 排除教师答案。
4. 下一次课开场先安排 10 分钟 Task Manager 提交与 Code Review 反馈：展示两类匿名共性问题、让学员修正一处代码，并说明评分反馈；该活动属于下一次课，不计入本日 09:30–16:00 的 310 分钟。
