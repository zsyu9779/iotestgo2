# Task 6 Report：同步 Exit Quiz、答案和讲师诊断语言

## 改动

- 将 Exit Quiz 标题改为“核心 14 题 + 选做 3 题”，增加第 15 题 ConstSpec/iota、第 16 题 Slice copy 和第 17 题 Map Struct value。
- 将教师答案扩展至 17 题：15=A、16=B、17=C；每题均记录任务简报规定的错误选项诊断。
- 在 Demo Notes 增加 Exit Quiz 诊断表，提供三道选做题的讲师追问和救援语言。
- 在 Runbook 固定 15:54–15:59 为核心题 1–14；15–17 只供提前完成者或课后 Code Review，且不挤占 Task Manager 启动时间。

## 验证

- Quiz 题号和 Exit 答案表均按顺序唯一覆盖 1–17。
- `rg -n '^## (15|16|17)\\.' module01_basics/assessments/exit_quiz.md`：三道选做题均存在。
- Exit 答案表中第 15–17 题分别为 A、B、C，且含对应诊断。
- Runbook、Demo Notes 和 Quiz 均明确核心题 1–14 的 5 分钟时间盒，以及 15–17 的提前完成者/课后 Code Review 定位。
- `make module01-demo-contracts`：PASS。
- `git diff --check`：PASS。

## 疑虑

- 无。`.superpowers/sdd/task-6-report.md` 按任务要求保留为未提交的交接记录；提交中不包含 `.superpowers`。

## 提交

- `a7b2557 docs: extend module01 semantic assessments`
