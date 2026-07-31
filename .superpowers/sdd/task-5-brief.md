### Task 5：建立受控编译失败与 panic 教学路径

**Files:**
- Create: `module01_basics/teaching_failures/README.md`
- Create: `module01_basics/teaching_failures/verify.sh`
- Create: `module01_basics/teaching_failures/testdata/compile_fail/*/main.go`
- Create: `module01_basics/teaching_failures/testdata/runtime_fail/*/main.go`
- Modify: `Makefile`
- Modify: `module01_basics/instructor/DEMO_NOTES.md`
- Modify: `module01_basics/instructor/RUNBOOK.md`
- Modify: `module01_basics/README.md`

**Interfaces:**
- Consumes: 根 Go Module 和 `testdata` 不被 `go test ./...` 默认扫描的 Go 工具约定。
- Produces: `make module01-teaching-failures`；八个“按预期失败”的独立 Case；`make module01-audit`。

- [ ] **Step 1: 创建编译失败 Case**

`package_short_decl/main.go`：

```go
package packageshortdecl

value := 1

var _ = value
```

`no_new_variable/main.go`：

```go
package nonewvariable

func demonstrate() {
	value := 1
	value := 2
	_ = value
}
```

`defined_type_assignment/main.go`：

```go
package definedtypeassignment

type UserID int

func demonstrate() {
	var raw int = 1
	var id UserID = raw
	_ = id
}
```

`slice_comparison/main.go`：

```go
package slicecomparison

func demonstrate() {
	left := []int{1}
	right := []int{1}
	_ = left == right
}
```

`map_struct_field/main.go`：

```go
package mapstructfield

type Item struct {
	Value int
}

func demonstrate() {
	items := map[string]Item{"one": {Value: 1}}
	items["one"].Value = 2
}
```

`final_fallthrough/main.go`：

```go
package finalfallthrough

func demonstrate(value int) {
	switch value {
	case 1:
		fallthrough
	}
}
```

- [ ] **Step 2: 创建 runtime failure Case**

`nil_map_write/main.go`：

```go
package main

func main() {
	var scores map[string]int
	scores["Alice"] = 95
}
```

`nil_pointer/main.go`：

```go
package main

func main() {
	var value *int
	*value = 1
}
```

- [ ] **Step 3: 创建预期失败验证脚本**

创建 `module01_basics/teaching_failures/verify.sh`：

```bash
#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd "$script_dir/../.." && pwd)"
cd "$repo_root"

expect_failure() {
    local command_kind="$1"
    local package_path="$2"
    local expected_pattern="$3"
    local output
    local status

    set +e
    if [[ "$command_kind" == "test" ]]; then
        output="$(go test "$package_path" 2>&1)"
        status=$?
    else
        output="$(go run "$package_path" 2>&1)"
        status=$?
    fi
    set -e

    if [[ $status -eq 0 ]]; then
        printf 'expected failure but command passed: %s %s\n' "$command_kind" "$package_path" >&2
        exit 1
    fi
    if ! grep -Eq "$expected_pattern" <<<"$output"; then
        printf 'failure diagnostic mismatch: %s\n%s\n' "$package_path" "$output" >&2
        exit 1
    fi
    printf 'expected failure: %s\n' "$package_path"
}

expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/package_short_decl 'outside function body|syntax error'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/no_new_variable 'no new variables'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/defined_type_assignment 'cannot use raw'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/slice_comparison 'slice can only be compared to nil|invalid operation.*left == right'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/map_struct_field 'cannot assign to struct field'
expect_failure test ./module01_basics/teaching_failures/testdata/compile_fail/final_fallthrough 'cannot fallthrough final case'
expect_failure run ./module01_basics/teaching_failures/testdata/runtime_fail/nil_map_write 'assignment to entry in nil map'
expect_failure run ./module01_basics/teaching_failures/testdata/runtime_fail/nil_pointer 'nil pointer dereference|invalid memory address'

printf 'module01 teaching failures: PASS\n'
```

- [ ] **Step 4: 增加 Make 入口并保持两类验证分离**

在 `.PHONY` 增加 `module01-teaching-failures module01-audit`，并增加：

```make
module01-teaching-failures:
	bash module01_basics/teaching_failures/verify.sh

module01-audit: module01-verify module01-demo-contracts module01-teaching-failures
```

不得把故意错误源码移出 `testdata`，不得让 `go test ./module01_basics/...` 因这些 Case 失败。

- [ ] **Step 5: 编写教学失败 README 与讲师小抄**

`teaching_failures/README.md` 必须列出：Case 路径、运行命令、预期诊断、正确写法、是否核心。`DEMO_NOTES` 的课前命令增加：

```bash
make module01-demo-contracts
make module01-teaching-failures
```

课堂只选择三个核心失败：同作用域无新变量的 `:=`、Map Struct 字段不可直接赋值、nil Map 写入 panic。其余 Case 作为讲师按班级情况调用的小抄，不逐个占用课堂时间。

Runbook 明确：预期失败不是环境失败；只有命令非零且诊断匹配才算教学 Case 通过。

- [ ] **Step 6: 验证并提交**

```bash
chmod +x module01_basics/teaching_failures/verify.sh
make module01-teaching-failures
go test ./module01_basics/...
make module01-demo-contracts
git diff --check
git add Makefile module01_basics/README.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md module01_basics/teaching_failures
git commit -m "feat: add controlled failure lessons for module01"
```

Expected: `module01 teaching failures: PASS`；正常 `go test ./module01_basics/...` PASS。

---

