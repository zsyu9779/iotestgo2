# Module 01 Task 2 Report

## 修改文件

- `module01_basics/blocks/01_go_basics/demo/03_control_funcs/main.go`
- `module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go`
- `module01_basics/instructor/scripts/verify_demo_contracts.sh`
- `module01_basics/blocks/01_go_basics/README.md`
- `module01_basics/instructor/DEMO_NOTES.md`
- `module01_basics/instructor/RUNBOOK.md`

## RED 证据

在仅扩展 `verify_demo_contracts.sh` 后执行 `make module01-demo-contracts`，退出码为 2，输出：

```text
bash module01_basics/instructor/scripts/verify_demo_contracts.sh
missing expected output: outer score: 50
make: *** [module01-demo-contracts] Error 1
```

这证明新控制流输出契约在实现前失败，且失败原因符合简报预期。

## GREEN 证据

实现 if 初始化作用域、continue/break、无限 for、switch 初始化/多值 case、fallthrough 保留、String/`[]byte` 往返与 `strconv.Atoi` 成功/失败路径后，`make module01-demo-contracts` 输出 `module01 demo contracts: PASS`。

## 命令输出摘要

- `gofmt -w module01_basics/blocks/01_go_basics/demo/03_control_funcs/main.go module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go`：成功，无输出。
- `make module01-demo-contracts`：成功，`module01 demo contracts: PASS`。
- `go test ./module01_basics/blocks/01_go_basics/...`：成功；四个 Demo 与 starter 无测试文件，`lab/solution` 为 `ok`。
- `git diff --check`：成功，无输出。

## Commit

`b8083da feat: deepen module01 control flow and string teaching`

## 疑虑

无。报告位于 `.superpowers/sdd/`，按要求不纳入提交。
