# Module 01 标准课程模板设计

## 1. 背景与目标

`iotestgo2` 用于面向有 Java 经验的大学生开展六周 Go 后端课程。每周授课一天，Module 01 对应第一周。

实际授课时间为：

- 上午：09:30–12:00
- 午休：12:00–13:00，不计入授课
- 下午：13:00–16:00
- 上午、下午各安排一次 10 分钟短休息
- 净教学与实操时间：约 310 分钟

本设计把 Module 01 从“按知识点排列的完成版示例集合”重组为可复用的课程模板，同时满足讲师授课、学员实操、课后作业和自动验收四种需要。

本轮目标：

1. 建立可复制到 Module 02–06 的统一教学结构。
2. 让课堂每 30–45 分钟产生一个可运行、可验收的结果。
3. 将学员动手时间提高到净授课时间的 50% 以上。
4. 将 Task Manager 定位为一周课后实战作业，不挤占课堂主体时间。
5. 提供可独立复制到 Gitee 学员仓库的作业包。
6. 让本地验收和 Gitee CI 使用同一套规则。

## 2. 设计原则

### 2.1 以学习旅程为主线

课堂不再依次讲十个独立目录，而是组织为四个连续 Block。每个 Block 都为下一阶段准备能力，最终由综合课堂 Lab 串联。

### 2.2 讲练闭环

每个 Block 使用相同流程：

```text
学习目标 → Java 对比 → 讲师 Demo → 学员 Lab
        → 自动验收 → 错误复盘 → 迁移问题
```

### 2.3 课堂与作业分离

课堂综合 Lab 使用 Scorebook 领域。Task Manager 使用相同能力但不同领域，要求学员在一周内独立迁移，而不是照抄课堂答案。

### 2.4 教师仓库与学员仓库分离

当前 GitHub 仓库保存完整教学内容、Starter、公开测试、Solution 和讲师材料。面向学员的 Gitee 仓库由教师按课程进度渐进披露，只复制本周允许公开的 Demo 和 `student_pack`。

### 2.5 验收规则平台无关

作业验收的唯一事实来源是 `scripts/grade.sh`。本地 `make grade` 与 Gitee 手动流水线都调用该脚本，避免两套规则漂移。

## 3. Module 01 目标结构

```text
module01_basics/
├── README.md
├── instructor/
│   ├── RUNBOOK.md
│   ├── DEMO_NOTES.md
│   └── RUBRIC.md
├── blocks/
│   ├── 01_go_basics/
│   ├── 02_collections/
│   ├── 03_modeling/
│   └── 04_functions_testing/
│       ├── README.md
│       ├── demo/
│       └── lab/
│           ├── README.md
│           ├── starter/
│           └── solution/
├── integrated_lab/
│   └── scorebook/
│       ├── README.md
│       ├── starter/
│       └── solution/
├── homework/
│   └── task_manager/
│       ├── student_pack/
│       └── teacher/
├── assessments/
│   ├── entry_quiz.md
│   ├── exit_quiz.md
│   └── answer_key.md
└── bonus/
    ├── generics/
    ├── data_structures/
    └── dark_corners/
```

每个 Block 的 `README.md` 固定包含：

1. 学习结果
2. 时间盒
3. 前置知识
4. Java 对比
5. 讲师 Demo 步骤
6. 学员任务
7. 验收命令
8. 常见错误
9. 三级提示
10. 复盘问题
11. Bonus

## 4. 课堂安排

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

净教学与实操时间为 310 分钟。课堂核心不要求 Generics、链表或复杂函数式模式；这些内容进入 Bonus。

## 5. 四个教学 Block

### 5.1 Block 1：Go Basics

核心内容：

- `go run`、`go build`、`go test` 的角色
- `package main`、`import`、`func main`
- 变量、常量、`iota`、零值、显式类型转换
- `if`、`for`、`switch`
- 函数、多返回值和基础输入校验

课堂 Lab：实现成绩等级计算器。Starter 可编译；学员完成范围校验和等级映射后通过验收。

### 5.2 Block 2：Collections

核心内容：

- Array 与 Slice 的值语义差异
- `len`、`cap`、`append`、共享底层数组
- Map 初始化、comma-ok、删除和无序性
- String、byte、rune 与 UTF-8

课堂 Lab：对中英文文本进行字符数、单词数和词频统计，并观察 Slice 扩容和 Map 无序输出。

### 5.3 Block 3：Modeling

核心内容：

- Go 只有值传递
- Pointer 的修改语义
- Struct、构造函数惯例和 Method
- 值接收器与指针接收器的选择
- 组合优先于继承

课堂 Lab：为 Scorebook 建立 `Student`、`Scorebook` 和状态修改方法。

### 5.4 Block 4：Functions & Testing

核心内容：

- 函数值、匿名函数和闭包的基本用途
- `defer` 的资源收尾语义
- `testing` 包入门
- 读取测试失败信息并完成修复

课堂 Lab：为 Scorebook 增加统计函数和基础测试。表格驱动测试仅做初步体验，系统讲解留给 Module 02。

## 6. 综合课堂 Lab：Scorebook

Scorebook 用于验证学员能否组合当天知识，但不提供 Task Manager 的直接答案。

最低行为：

- 添加学生
- 按 ID 查找学生
- 更新学生成绩
- 计算平均分
- 按等级统计人数
- 为核心行为编写测试

Starter 提供可编译的类型和方法骨架；Solution 提供完整实现和测试。课堂验收只关注行为，不要求复杂分层、接口注入或文件持久化。

## 7. Task Manager 课后作业

### 7.1 定位

Task Manager 是 Module 01 结束后的一周实战作业。课堂最后 20 分钟只负责：

1. 演示目标行为。
2. 跑通 Starter 和本地验收命令。
3. 解释 Git 提交方式、里程碑和评分规则。

下一周开课时安排简短 Code Review，形成“布置—完成—反馈”的闭环。

### 7.2 作业目录

```text
homework/task_manager/
├── student_pack/
│   ├── README.md
│   ├── go.mod
│   ├── Makefile
│   ├── cmd/taskmanager/main.go
│   ├── taskmanager/
│   │   ├── task.go
│   │   ├── manager.go
│   │   └── manager_test.go
│   ├── scripts/grade.sh
│   └── .workflow/GradePipeline.yml
└── teacher/
    ├── solution/
    ├── RUBRIC.md
    ├── RELEASE_CHECKLIST.md
    └── TROUBLESHOOTING.md
```

`student_pack` 是一个独立 Go Module，可直接把目录内容复制到 Gitee 学员仓库根目录。它只使用 Go 标准库，不依赖外部模块下载。

### 7.3 必做行为

- 新增任务并分配递增 ID
- 列出任务
- 根据 ID 标记完成
- 根据 ID 删除任务
- 对空标题、未知 ID 等输入返回明确错误
- CLI 支持新增、列表、完成、删除和退出
- 核心业务测试通过

文件持久化、复杂 CLI 框架、第三方库和接口注入不属于 Week 1 必做范围。

### 7.4 一周里程碑

1. 建立 `Task` 和 `Manager` 模型。
2. 完成 Add、List、Complete、Delete。
3. 完成 CLI 输入循环和错误提示。
4. 修复全部测试并整理提交记录。

### 7.5 评分建议

| 维度 | 分值 |
|---|---:|
| 必做行为正确 | 60 |
| 测试通过并补充有效测试 | 20 |
| `gofmt`、命名和代码组织 | 10 |
| README 设计说明与提交记录 | 10 |

## 8. 本地与 Gitee 验收

### 8.1 单一验收入口

学员本地运行：

```bash
make grade
```

`make grade` 调用 `scripts/grade.sh`，依次执行：

1. 检查 `gofmt`。
2. 执行 `go vet ./...`。
3. 执行 `go test ./...`。
4. 执行 `go build ./...`。
5. 输出中文通过或失败摘要。

脚本在任一步失败时返回非零退出码。失败信息指出失败阶段，但不暴露实现答案。

### 8.2 Gitee 手动流水线

学生把作业提交到各自的开发分支。Gitee 流水线不配置 push 或 Pull Request 自动触发。教师在 Gitee 选择学生分支后手动运行 `GradePipeline`，流水线只调用 `scripts/grade.sh`。

根据 Gitee 官方说明，流水线描述文件位于仓库根目录的 `.workflow` 目录。本实现提供可复制的 `GradePipeline.yml`，但评分规则不写入 YAML，避免平台配置和本地规则漂移。

参考：

- [Gitee Go 快速入门](https://gitee.com/help/articles/4357)
- [Gitee Go 操作指南](https://gitee.com/help/categories/72)

## 9. 仓库级验证

仓库根目录新增统一入口：

```bash
make module01-verify
make module01-lab-01
make module01-lab-02
make module01-lab-03
make module01-lab-04
make module01-integrated-lab
make module01-homework-solution
```

完整仓库默认验证只运行 Demo、课堂 Lab Solution、Scorebook Solution 和教师版 Task Manager Solution。`student_pack` 使用独立 `go.mod`，其未完成 Starter 不进入根模块的默认 `go test ./...`。

## 10. 现有内容迁移

| 现有内容 | 新位置 |
|---|---|
| `01_hello`、`02_vars_types`、`03_control_funcs` | `blocks/01_go_basics/` |
| `04_arrays_slices`、`05_maps_strings` | `blocks/02_collections/` |
| `06_pointers`、`07_structs_methods` | `blocks/03_modeling/` |
| `09_advanced_functions` 核心内容 | `blocks/04_functions_testing/` |
| `08_data_structures` | `bonus/data_structures/` |
| `10_generics_intro` | `bonus/generics/` |
| range、map、string 深层陷阱 | `bonus/dark_corners/` |
| `project_task_manager` | `homework/task_manager/` |

旧目录在内容迁移完成后移除，避免两套课程结构并存。Git 历史通过移动文件保留。现有 Module 01 教案和讲课小抄中的有效内容分别吸收到 `RUNBOOK.md`、`DEMO_NOTES.md` 和 Block README；原文档随后改为指向新入口，消除课程大纲、教案和源码之间的 8/9/10 节不一致。

## 11. 通用课程模板文档

实现阶段新增 `docs/course-module-standard.md`，凝练所有 Module 可复用的规则：

- 标准目录
- 时间盒和课堂节奏
- Block README 模板
- Demo、Lab、Homework 的职责
- Starter/Solution 分离规则
- 自动验收规则
- 教师仓库到学员仓库的披露流程
- Gitee 手动 CI 适配原则
- Module 完成定义

Module 02–06 本轮不改代码，只在该文档中记录后续迁移原则。

## 12. 完成定义

Module 01 达到以下条件才算完成：

1. 新目录结构落地，旧内容完成迁移且无重复入口。
2. 四个 Block 都具备 Demo、Starter、Solution、Lab README 和验收方式。
3. Scorebook 综合 Lab 可独立运行和测试。
4. Task Manager `student_pack` 可复制为独立仓库，并且只使用标准库。
5. Task Manager 教师 Solution 通过格式、vet、测试和构建检查。
6. Gitee 手动流水线调用与本地相同的 `grade.sh`。
7. 讲师 Runbook 精确覆盖 09:30–16:00，并包含两次短休息和一小时午休。
8. Entry Quiz、Exit Quiz、评分标准和发布检查表齐全。
9. 总课程大纲与 Module 01 新结构一致。
10. `docs/course-module-standard.md` 可作为 Module 02–06 的改造模板。

## 13. 本轮边界

本轮不包含：

- 重构 Module 02–06 源码
- 创建或管理实际 Gitee 学员仓库
- 自动批量导出或同步 GitHub/Gitee 仓库
- 隐藏测试、反作弊或相似度检测
- 引入第三方 CLI、测试或评分依赖
- 将文件持久化列为 Week 1 必做项

