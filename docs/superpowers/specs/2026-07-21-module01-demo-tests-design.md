# module01 课堂作业 Demo 测试设计

## 目标

为 `module01_basics` 的每组 solution 增加一个可直接在 GoLand 中点击运行的 Demo 测试，用 `testing.T.Logf` 展示典型输入、操作过程和结果，帮助课堂讲解和学生观察程序行为。

## 范围

覆盖以下 6 组 solution：

1. `blocks/01_go_basics/lab/solution`
2. `blocks/02_collections/lab/solution`
3. `blocks/03_modeling/lab/solution`
4. `blocks/04_functions_testing/lab/solution`
5. `integrated_lab/scorebook/solution`
6. `homework/task_manager/teacher/solution/taskmanager`

不修改 `starter`，不改变现有断言测试的行为。

## 设计

每组新增独立的 `demo_test.go`（文本统计已有的 `TestAnalyzeDemo` 保留并作为该组 Demo），Demo 测试使用真实的业务流程：

- 成绩等级：展示多个分数到等级的映射。
- 文本统计：展示输入文本、字节数、字符数、单词数和词频。
- 学生模型：展示创建学生、改名、更新成绩和快照结果。
- 函数与闭包：展示筛选结果和审计事件顺序。
- Scorebook：展示添加学生、更新成绩、平均分和等级统计。
- Task manager：展示添加任务、完成任务、删除任务后的列表。

Demo 只使用 `t.Logf` 输出可读信息；遇到 API 错误时使用 `t.Fatal`，确保演示本身仍是有效测试。输出不依赖 map 的遍历顺序，必要时使用固定输入或稳定字段展示。

## 使用方式

在 GoLand 中直接点击 `demo_test.go` 中 Demo 函数旁的运行图标，或右键对应 solution 目录运行测试。命令行等价方式为：

```bash
go test -v ./path/to/solution
```

## 验收标准

- 6 组 solution 各有一个可运行的 Demo 测试。
- 所有 Demo 都能在 GoLand 的测试窗口显示过程日志。
- 原有测试和新增 Demo 全部通过。
- 不修改 starter 文件和用户已有的无关改动。
