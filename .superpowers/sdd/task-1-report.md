# Module 01 Task 1 Report

## 修改文件

- `Makefile`
- `module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go`
- `module01_basics/blocks/01_go_basics/README.md`
- `module01_basics/instructor/DEMO_NOTES.md`
- `module01_basics/instructor/RUNBOOK.md`
- `module01_basics/instructor/scripts/verify_demo_contracts.sh`

## TDD 与验证

- `make module01-demo-contracts`（RED）：按预期失败，缺少 `blank identifier: 1 3`。
- `gofmt -w module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go`：完成格式化。
- `chmod +x module01_basics/instructor/scripts/verify_demo_contracts.sh`：设置验证脚本可执行。
- `make module01-demo-contracts`：PASS，输出 `module01 demo contracts: PASS`。
- `go test ./module01_basics/blocks/01_go_basics/...`：PASS；Demo 包无测试，lab/solution 为 `ok`。
- `git diff --check`：PASS，无空白错误。

## Commit

`a3de6e3 feat: deepen module01 variables and constants teaching`

## 已知疑虑

- 简报中的 `truncated := int(1.9)` 不能在 Go 中编译：未类型化浮点常量必须可精确表示为目标整数类型。为保留指定的教学结果，演示使用 `floatValue := 1.9` 后执行 `int(floatValue)`，输出仍为 `truncated float: 1`。
- 本报告按要求保留在 `.superpowers/sdd/`，未纳入任务提交。

## 审查修复

- 修正 `DEMO_NOTES.md` 的浮点转整数暂停点，使其与实际可编译演示 `floatValue := 1.9; int(floatValue)` 一致，并说明 `int(1.9)` 因 1.9 不能精确表示为 int 而编译失败。
- 补充 `const messageLength = len(message)` 的提问、预期结果、准确解释、常见误解和级别，明确当 `message` 为常量时 `len(message)` 是编译期常量。
- `make module01-demo-contracts`：PASS，输出 `module01 demo contracts: PASS`。
- `go test ./module01_basics/blocks/01_go_basics/...`：PASS；Demo 包无测试，lab/solution 为 `ok`。
- `git diff --check`：PASS，无空白错误。
- 审查修复提交：`69853aa docs: correct module01 conversion teaching notes`。
