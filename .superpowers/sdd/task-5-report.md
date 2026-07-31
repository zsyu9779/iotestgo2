# Module01 Task 5 执行报告

## 提交

- Commit: `546e0f7 feat: add controlled failure lessons for module01`
- Follow-up commit: `3ff7057 fix: exclude controlled failures from fmt check`
- `.superpowers/` 仅保留为本地任务记录，未纳入提交。

## 提交文件

- `Makefile`
- `module01_basics/README.md`
- `module01_basics/instructor/DEMO_NOTES.md`
- `module01_basics/instructor/RUNBOOK.md`
- `module01_basics/teaching_failures/README.md`
- `module01_basics/teaching_failures/verify.sh`
- `module01_basics/teaching_failures/testdata/compile_fail/package_short_decl/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/no_new_variable/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/defined_type_assignment/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/slice_comparison/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/map_struct_field/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/final_fallthrough/main.go`
- `module01_basics/teaching_failures/testdata/runtime_fail/nil_map_write/main.go`
- `module01_basics/teaching_failures/testdata/runtime_fail/nil_pointer/main.go`

## 先失败后通过验证

1. RED：先只创建 `verify.sh`，不创建任何教学 Case；运行 `bash module01_basics/teaching_failures/verify.sh` 以非零退出。首个 Case 收到的是 `directory not found`，不匹配要求的编译诊断，证明脚本不会把任意非零状态误判为教学成功。
2. GREEN：创建全部八个隔离 Case 后，`make module01-teaching-failures` 通过；脚本逐项输出 `expected failure`，最终输出 `module01 teaching failures: PASS`。
3. RED：完整审计首次显示 `module01-verify` 的 gofmt 扫描解析了故意无效的 `testdata` 源码，出现 `expected declaration, found value`。使用断言命令确认该噪声存在：`if make module01-verify 2>&1 | grep -q 'expected declaration, found value'; then exit 1; fi`。
4. GREEN：Makefile 的正常格式扫描显式 prune `module01_basics/teaching_failures/testdata`。同一断言转为成功；随后 `make module01-audit`、`git diff --check` 均通过。

## 最终验证

- `make module01-audit`：通过。包含 `module01-verify`（gofmt、vet、正常 `go test ./module01_basics/...`、Task Manager 教师答案）、`module01-demo-contracts` 和全部八个受控失败 Case。
- `git diff --check`：通过。
- 正常 `go test ./module01_basics/...`：通过；故意无法编译和会 panic 的源码仍全部位于 `teaching_failures/testdata/`，未进入正常 Demo 或正常 go test 包。

## 诊断匹配策略

- 每个 Case 都要求命令以非零状态退出；若意外成功，验证立即失败。
- 随后用 `grep -E` 匹配每个知识点的稳定诊断片段，而不是仅凭退出码。
- 对 Go 版本可能有表述差异的诊断使用受控备选模式：包级短声明接受 `outside function body|syntax error`，Slice 比较接受 `slice can only be compared to nil|invalid operation.*left == right`，nil 指针接受 `nil pointer dereference|invalid memory address`。
- 其余模式针对核心语义片段，例如 `no new variables`、`cannot assign to struct field`、`assignment to entry in nil map`，从而拒绝目录、权限、依赖或工具链问题造成的非零结果。

## 疑虑与处理

- 根 `module01-verify` 原有的 gofmt 扫描会遍历 `testdata`，对故意无效源码输出解析错误；已在 Makefile 中仅为该格式扫描排除该隔离目录，保持正常基线输出干净。正常 Vet/Test 本身遵循 Go 的 `testdata` 约定，不扫描这些 Case。
- 后续审查发现通用 `fmt-check` 同样使用 `git ls-files '*.go'`，会扫描隔离 Case；此前命令替换还会吞掉 `gofmt` 的非零状态。已在 follow-up commit 中排除同一 `testdata` 路径，并使用临时文件保留 `gofmt -l` 输出：任何其余已跟踪 Go 文件解析或格式化失败都会令 `fmt-check` 非零退出并打印清晰错误。修复前已复现 `expected declaration, found value` 且目标错误地返回 0；修复后 `make fmt-check`、`make module01-audit` 和 `git diff --check` 均通过。
- 编译器诊断可能随 Go 版本微调。脚本只匹配稳定且与教学点直接相关的片段；如未来 Go 版本更改措辞，应在保留“非零 + 语义匹配”原则下谨慎更新正则。
