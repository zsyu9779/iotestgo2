# Module 01 Task 8 验收报告

执行日期：2026-07-20（Asia/Shanghai）

## 发现与最小修复

发现 Task 2 review 的 Minor 尚未落实：
`module01_basics/instructor/scripts/verify_demo_contracts.sh` 只断言
`loop body: 0` 与 `loop body: 2` 存在，未反向断言 `loop body: 1` 和
`loop body: 3` 不存在。

最小修复：新增 `assert_not_contains`，并为上述两条不应出现的循环输出增加
反向断言。没有修改 Module 02 或之后模块，也没有修改 Starter。

## 命令与退出结果

| 命令 | 退出结果 | 结果摘要 |
| --- | --- | --- |
| `git status --short` | 0 | 初始仅有未跟踪的 `.superpowers/`；无既有已暂存改动。 |
| `git log --oneline -10` | 0 | 确认 Tasks 1–7 提交序列及 HEAD `f4e24bc`。 |
| `git rev-parse --show-toplevel` | 0 | 工作区为 `/Users/zhangshiyu/iotestgo2`。 |
| `git branch --show-current` | 0 | 当前分支为 `main`；按用户要求直接在当前工作区操作。 |
| `make module01-verify` | 0 | gofmt、`go vet ./module01_basics/...`、`go test ./module01_basics/...` 及 Task Manager 教师答案验收均通过。 |
| `make module01-demo-contracts` | 0 | 输出 `module01 demo contracts: PASS`；含新增的两个反向循环输出断言。 |
| `make module01-teaching-failures` | 0 | 八个 Case 均输出 `expected failure`，最后输出 `module01 teaching failures: PASS`。 |
| `go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter` | 1（预期） | 首个失败：`TestGrade/excellent_lower_bound`，为 `Grade(90)` 行为断言差异。 |
| `go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter` | 1（预期） | 首个失败：`TestAnalyze`，为 Bytes 行为断言差异。 |
| `go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter` | 1（预期） | 首个失败：`TestStudentLifecycle`，为构造结果行为断言差异。 |
| `go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter` | 1（预期） | 首个失败：`TestFilterPreservesInputOrderForAdditionalCase`，为过滤行为断言差异。 |
| `make module01-audit` | 0 | 汇总 GREEN、输出契约和八个受控失败验证均通过。 |
| `git diff --check` | 0 | 无空白错误。 |
| `git diff --name-only 86715a8..3f783f8 \| rg '^module02_' && exit 1 \|\| true` | 0 | 明确基线范围检查命令正常完成。 |
| `git diff --name-only 86715a8..3f783f8 \| rg -q '^module02_'` | 1（预期） | 无 `module02_` 路径；额外的严格检查标记为 `MODULE02_SCOPE_CHECK=PASS`。 |
| `git diff --cached --name-only` | 0 | 审计阶段索引为空；`.superpowers/` 未暂存。 |
| `make module01-audit`（提交后复验） | 0 | GREEN、Demo 契约与八个受控失败再次全部通过。 |
| `git diff --check`（提交后复验） | 0 | 无空白错误。 |
| `git status --short`（提交后复验） | 0 | 仅 `.superpowers/` 未跟踪。 |
| `git diff --name-only 86715a8..3f783f8 \| rg '^module02_' && exit 1 \|\| true`（提交后复验） | 0 | 明确基线范围检查命令正常完成。 |
| `git diff --name-only 86715a8..3f783f8 \| rg -q '^module02_'`（提交后复验） | 1（预期） | 严格检查确认不存在 Module 02 路径。 |
| `git show --stat --oneline --summary HEAD` | 0 | 本次提交仅涉及一个 Module 01 输出契约脚本。 |
| `git diff --cached --name-only`（提交后复验） | 0 | 索引为空；`.superpowers/` 未被暂存。 |

## Starter RED 判定

四个 Starter 均因未完成实现导致的测试断言/行为差异返回非零；输出中没有编译错误、模块依赖缺失或包路径错误。因此这些 RED 均符合任务约定，未将 Starter 改为 GREEN。

## 范围与疑虑

- 修复范围仅为 `module01_basics/instructor/scripts/verify_demo_contracts.sh`。
- `.superpowers/` 中的本报告和进度账本保持未跟踪，不会提交。
- 无未解决疑虑；有意失败的 Starter 与教学失败 fixture 均已按其各自契约验证。

## 提交

- 提交：`3f783f8 test: complete module01 curriculum audit`。
- 提交内容仅为 `module01_basics/instructor/scripts/verify_demo_contracts.sh` 的 11 行新增反向断言辅助逻辑；`.superpowers/` 未被暂存或提交。
