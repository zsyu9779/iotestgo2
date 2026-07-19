# Module 01 内容补充与教案同步实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在不改变 Module 01 四个 Block、Lab、综合练习和作业主结构的前提下，补齐老教学项目中高教学密度的 Go 基础语义 Case，并同步完善讲师 Demo Notes、Runbook、学员 README 和测评材料。

**Architecture:** 继续以 `module01_basics/blocks` 为课程主线，不为每个知识点新建可运行 Demo 目录；新增内容合入现有主题 Demo。正常工程代码、Solution 和教师答案保持 GREEN；Starter 保留预期 RED；编译错误与 panic 示例放入 `teaching_failures/testdata`，通过独立脚本验证“按预期失败”。每个代码任务同时修改对应讲师小抄，避免代码与教案漂移。

**Tech Stack:** Go 1.25 根模块、Go 1.16 Task Manager 独立模块、Bash、Make、Markdown、Go `testing`、`go vet`、`gofmt`。

## Global Constraints

- 本计划只修改 `module01_basics`、根 `Makefile` 和本实施计划；不得修改 `module02_advanced` 及之后模块。
- 保留用户当前未提交的 `module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go` 中 `consA` 到 `consG` 示例，不得覆盖或丢失。
- `consE` 的解释必须是“省略 ConstSpec 时复用上一行完整表达式”；不得表述为“iota 被打断”，因为 `iota` 仍按声明行递增。
- 不新增碎片化的正常 Demo 目录；新增可运行知识点必须合入现有 `02_vars_types`、`03_control_funcs`、`05_strings_basics`、Collections 或 Modeling Demo。
- 故意编译失败或 panic 的 Case 必须位于 `module01_basics/teaching_failures/testdata`，确保 `go test ./module01_basics/...` 不会扫描这些目录。
- `make module01-verify` 验证正常工程基线；`make module01-teaching-failures` 验证预期失败；`make module01-audit` 同时执行两类验收。
- 每个知识点必须在 `module01_basics/instructor/DEMO_NOTES.md` 写明：提问、预期结果、准确解释、常见误解和可裁级别。
- `module01_basics/instructor/RUNBOOK.md` 的总时间仍为 09:30–16:00，净教学与实操时间仍为 310 分钟，学员动手预算仍为 185 分钟。
- Task Manager 的两个独立 `go.mod` 继续保持 `go 1.16`；不得把根模块 Go 版本语法迁入其学员包或教师答案。
- 不以“所有测试全过”为课程唯一目标；Starter RED、编译失败和 panic Case 的成功标准是结果符合文档约定。

---

## 文件结构与职责

### 修改

- `module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go`：变量、空白标识符、常量表达式、iota、定义类型/别名、转换和零值。
- `module01_basics/blocks/01_go_basics/demo/03_control_funcs/main.go`：if 初始化作用域、三种 for、break/continue、switch 初始化、多值 case、无表达式 switch 和 fallthrough。
- `module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go`：String/byte/rune、`[]byte` 往返、`strconv` 成功与错误路径。
- `module01_basics/blocks/02_collections/demo/04_arrays_slices/main.go`：`copy(dst, src)` 的方向、复制数量和长度限制。
- `module01_basics/blocks/02_collections/demo/05_maps_strings/main.go`：嵌套 Map 的逐层初始化。
- `module01_basics/blocks/03_modeling/demo/06_pointers/main.go`：Struct 指针字段自动解引用语法糖。
- `module01_basics/blocks/03_modeling/demo/09_copy_and_receivers/main.go`：值接收者内部修改不影响调用方。
- `module01_basics/blocks/01_go_basics/README.md`、`02_collections/README.md`、`03_modeling/README.md`：同步学员可见语义和运行入口。
- `module01_basics/instructor/DEMO_NOTES.md`：讲师逐知识点小抄。
- `module01_basics/instructor/RUNBOOK.md`：时间盒、核心/深挖/可裁顺序和预期失败入口。
- `module01_basics/assessments/exit_quiz.md`、`answer_key.md`：增加 iota、`copy` 和编译失败诊断。
- `module01_basics/README.md`：增加两类验收命令说明。
- `Makefile`：增加内容契约、教学失败和总审计目标。

### 新建

- `module01_basics/instructor/scripts/verify_demo_contracts.sh`：运行现有 Demo 并断言关键教学输出。
- `module01_basics/teaching_failures/README.md`：说明预期失败的用途和运行方式。
- `module01_basics/teaching_failures/verify.sh`：验证编译失败与 panic 诊断。
- `module01_basics/teaching_failures/testdata/compile_fail/package_short_decl/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/no_new_variable/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/defined_type_assignment/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/slice_comparison/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/map_struct_field/main.go`
- `module01_basics/teaching_failures/testdata/compile_fail/final_fallthrough/main.go`
- `module01_basics/teaching_failures/testdata/runtime_fail/nil_map_write/main.go`
- `module01_basics/teaching_failures/testdata/runtime_fail/nil_pointer/main.go`

---

### Task 1：补齐变量、常量、iota、定义类型与转换

**Files:**
- Modify: `module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go:8-76`
- Create: `module01_basics/instructor/scripts/verify_demo_contracts.sh`
- Modify: `module01_basics/blocks/01_go_basics/README.md:35-63`
- Modify: `module01_basics/instructor/DEMO_NOTES.md:24-60`
- Modify: `module01_basics/instructor/RUNBOOK.md:62-70`
- Modify: `Makefile:1-3,52-58`

**Interfaces:**
- Consumes: 用户已添加的 `consA`、`consB`、`consC`、`consD`、`consE`、`consF`、`consG` 常量及其输出。
- Produces: `make module01-demo-contracts`；稳定输出 `blank identifier`、`const expression length`、`iota edge cases`、`iota reset`、`defined type conversion`、`truncated float`。

- [ ] **Step 1: 先创建会失败的 Demo 输出契约脚本**

创建 `module01_basics/instructor/scripts/verify_demo_contracts.sh`：

```bash
#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd "$script_dir/../../.." && pwd)"
cd "$repo_root"

assert_contains() {
    local output="$1"
    local expected="$2"
    if ! grep -Fq "$expected" <<<"$output"; then
        printf 'missing expected output: %s\n' "$expected" >&2
        exit 1
    fi
}

vars_output="$(go run ./module01_basics/blocks/01_go_basics/demo/02_vars_types)"
assert_contains "$vars_output" 'blank identifier: 1 3'
assert_contains "$vars_output" 'const expression length: 3'
assert_contains "$vars_output" 'iota edge cases: 0 1 2 250 250 5 6'
assert_contains "$vars_output" 'iota reset: 0 1'
assert_contains "$vars_output" 'defined type conversion: 12 alias assignment: 12'
assert_contains "$vars_output" 'truncated float: 1'

printf 'module01 demo contracts: PASS\n'
```

在 `Makefile` 的 `.PHONY` 增加 `module01-demo-contracts`，并增加：

```make
module01-demo-contracts:
	bash module01_basics/instructor/scripts/verify_demo_contracts.sh
```

Run:

```bash
make module01-demo-contracts
```

Expected: FAIL，第一条缺失输出为 `blank identifier: 1 3`。

- [ ] **Step 2: 在 `02_vars_types` 增加空白标识符和常量表达式**

在 `Config` 后增加定义类型和别名：

```go
type UserID int
type UserIDAlias = int
```

在基础变量输出后加入：

```go
left, _, right := 1, 2, 3
fmt.Println("blank identifier:", left, right)
```

在 `const pi` 后加入：

```go
const message = "abc"
const messageLength = len(message)
fmt.Println("const expression length:", messageLength)
```

Expected teaching result: `_` 丢弃中间值且不能被读取；`len(message)` 在 `message` 为常量时是编译期常量。

- [ ] **Step 3: 保留用户 iota Case 并修正解释**

将当前 `consA` 到 `consG` 块整理为：

```go
// 仅供参考，不想挨打请勿模仿：省略声明会复用上一行完整表达式。
const (
	consA = iota // 0
	consB        // 1，复用 = iota
	consC        // 2，复用 = iota
	consD = 250  // iota 仍递增到 3，但当前表达式为 250
	consE        // 250，复用上一行完整表达式
	consF = iota // 5，显式恢复使用当前行的 iota
	consG        // 6，复用 = iota
)
fmt.Println("iota edge cases:", consA, consB, consC, consD, consE, consF, consG)

const (
	resetA = iota
	resetB
)
fmt.Println("iota reset:", resetA, resetB)
```

Expected output:

```text
iota edge cases: 0 1 2 250 250 5 6
iota reset: 0 1
```

- [ ] **Step 4: 增加定义类型、别名和转换截断**

在现有显式转换部分加入：

```go
var rawID int = 12
var userID UserID = UserID(rawID)
var aliasID UserIDAlias = rawID
fmt.Println("defined type conversion:", userID, "alias assignment:", aliasID)

truncated := int(1.9)
fmt.Println("truncated float:", truncated)
```

Expected teaching result: `UserID` 是新的定义类型，需要显式转换；`UserIDAlias` 与 `int` 是同一类型；浮点转整数截去小数部分，不做四舍五入。

- [ ] **Step 5: 同步 Block 1 README、讲师小抄和 Runbook**

在 `Block 1 README` 的变量语义段明确写入：

```markdown
- `_` 是空白标识符，用于显式丢弃值，它不是可读取变量。
- const 块中省略表达式会复用上一行完整 ConstSpec；`iota` 仍按声明行递增。
- 每个新的 const 块都会让 `iota` 从 0 重新开始。
- `type UserID int` 定义新类型；`type UserID = int` 声明别名。
- 浮点数转整数会截去小数部分。
```

在 `DEMO_NOTES` 的 `02_vars_types` 投影提示加入下表内容：

```markdown
| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| `left, _, right` | `_` 能否打印？ | 只能打印 1 和 3 | `_` 丢弃值，不绑定可读变量 | 把 `_` 当普通变量 | 核心 |
| `consD/consE/consF` | E、F 分别是多少？ | 250、5 | E 复用 `= 250`；iota 未停止，F 使用当前行号 5 | “iota 被打断后重新从 0 开始” | 核心 |
| 第二个 const 块 | resetA 是多少？ | 0 | 每个 const 块独立重置 iota | iota 在整个包连续计数 | 深挖 |
| `UserID` / alias | 哪个需要转换？ | UserID 需要，alias 不需要 | 定义类型与别名语义不同 | 两种 type 写法完全一样 | 深挖 |
| `int(1.9)` | 输出 1 还是 2？ | 1 | 转换截断小数，不四舍五入 | 数值转换会自动舍入 | 核心 |
```

将 Runbook Block 1 的可裁内容改为：位运算、第二个 iota 重置块和类型别名可裁；用户补充的隐式 ConstSpec/iota Case 不可裁。

- [ ] **Step 6: 格式化并验证 Task 1**

Run:

```bash
gofmt -w module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go
chmod +x module01_basics/instructor/scripts/verify_demo_contracts.sh
make module01-demo-contracts
go test ./module01_basics/blocks/01_go_basics/...
git diff --check
```

Expected: 契约脚本输出 `module01 demo contracts: PASS`；Block 1 正常包测试 PASS；无格式错误。

- [ ] **Step 7: 提交变量与常量补充**

```bash
git add Makefile module01_basics/blocks/01_go_basics/demo/02_vars_types/main.go module01_basics/blocks/01_go_basics/README.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md module01_basics/instructor/scripts/verify_demo_contracts.sh
git commit -m "feat: deepen module01 variables and constants teaching"
```

---

### Task 2：补齐控制流与 String 转换语义

**Files:**
- Modify: `module01_basics/blocks/01_go_basics/demo/03_control_funcs/main.go:5-62`
- Modify: `module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go:3-30`
- Modify: `module01_basics/instructor/scripts/verify_demo_contracts.sh`
- Modify: `module01_basics/blocks/01_go_basics/README.md`
- Modify: `module01_basics/instructor/DEMO_NOTES.md`
- Modify: `module01_basics/instructor/RUNBOOK.md`

**Interfaces:**
- Consumes: Task 1 的 `module01-demo-contracts` 脚本和 Block 1 文档结构。
- Produces: 可观察的 if 初始化作用域、break/continue、无限 for、switch 初始化、多值 case、fallthrough、`[]byte` 往返和 `strconv.Atoi` 错误路径。

- [ ] **Step 1: 扩展输出契约并确认失败**

在 `verify_demo_contracts.sh` 的最终 PASS 前加入：

```bash
control_output="$(go run ./module01_basics/blocks/01_go_basics/demo/03_control_funcs)"
assert_contains "$control_output" 'outer score: 50'
assert_contains "$control_output" 'switch init: owner has full access'
assert_contains "$control_output" 'loop body: 0'
assert_contains "$control_output" 'loop body: 2'
assert_contains "$control_output" 'infinite for stopped at: 3'
assert_contains "$control_output" 'Admin branch reached by fallthrough'

string_output="$(go run ./module01_basics/blocks/01_go_basics/demo/05_strings_basics)"
assert_contains "$string_output" 'bytes round trip: Go语言'
assert_contains "$string_output" 'Atoi success: 42'
assert_contains "$string_output" 'Atoi error: true'
```

Run: `make module01-demo-contracts`

Expected: FAIL，缺失 `outer score: 50`。

- [ ] **Step 2: 修改 if 初始化和循环 Case**

用以下代码替换 `03_control_funcs` 中现有 if 和 for 演示：

```go
outerScore := 50
if score := 85; score >= 60 {
	fmt.Println("Passed with score:", score)
}
fmt.Println("outer score:", outerScore)

for i := 0; i < 5; i++ {
	if i == 1 {
		continue
	}
	if i == 3 {
		break
	}
	fmt.Println("loop body:", i)
}

attempts := 0
for {
	attempts++
	if attempts == 3 {
		break
	}
}
fmt.Println("infinite for stopped at:", attempts)
```

Expected output includes `loop body: 0` and `loop body: 2`，不包含 `loop body: 1` 或 `loop body: 3`。

- [ ] **Step 3: 用带初始化的 switch 承载多值 case**

将第一个 role switch 替换为：

```go
role := "owner"
switch currentRole := role; currentRole {
case "admin", "owner":
	fmt.Println("switch init:", currentRole, "has full access")
case "user", "viewer":
	fmt.Println("switch init:", currentRole, "has read only access")
default:
	fmt.Println("switch init:", currentRole, "is denied")
}
```

保留现有独立 fallthrough switch 和无表达式 switch。讲解时明确：`currentRole` 只存在于当前 switch；多值 case 表示 OR；fallthrough 不检查下一 case 的表达式。

- [ ] **Step 4: 补 String 与 byte、strconv 的工程路径**

在 `05_strings_basics/main.go` import 中加入 `strconv`，并在 rune 输出后加入：

```go
bytes := []byte("Go语言")
roundTrip := string(bytes)
fmt.Println("bytes round trip:", roundTrip)

number, err := strconv.Atoi("42")
if err != nil {
	fmt.Println("Atoi unexpected error:", err)
} else {
	fmt.Println("Atoi success:", number)
}

_, err = strconv.Atoi("forty-two")
fmt.Println("Atoi error:", err != nil)
```

Expected teaching result: String/`[]byte` 可以显式转换；解析外部字符串必须处理 error，不能使用 `_` 丢弃失败。

- [ ] **Step 5: 同步教案**

在 `DEMO_NOTES` 增加以下讲师提示：

```markdown
| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| if 初始化 | if 外打印哪个 score？ | 外层仍为 50 | 初始化变量只在 if/else 作用域内，外层变量未改变 | 认为短声明修改了外层变量 | 核心 |
| continue / break | 1 和 3 会不会进入循环体？ | 1 被跳过，3 终止循环 | continue 跳过本轮，break 结束当前循环 | 两者都只是跳过本轮 | 核心 |
| 无限 for | 如何安全退出？ | attempts 到 3 时 break | `for {}` 是合法无限循环，需要显式退出条件 | Go 必须写 while | 核心 |
| switch 初始化 | currentRole 在 switch 外可用吗？ | 不可用 | 初始化变量作用域属于 switch | 与普通赋值作用域相同 | 深挖 |
| 多值 case | owner 匹配哪个分支？ | 第一个分支 | 逗号分隔值表示 OR | 一个 case 只能有一个值 | 核心 |
| fallthrough | 下一 case 条件会重算吗？ | 不会 | 无条件进入紧邻 case，业务代码通常不建议使用 | 相当于继续匹配条件 | 深挖 |
| Atoi 失败 | 错误能否忽略？ | err 非 nil | 外部输入解析属于可预期失败，应显式处理 | 转换失败自动得到 0 且安全 | 核心 |
```

在 Block 1 README 同步上述规则；Runbook 保持 15 分钟 Demo 时间，标签、goto、switch 初始化作用域细节和 String 转换成功路径可裁，break/continue、多值 case、fallthrough 语义和 Atoi 错误路径保留。

- [ ] **Step 6: 验证并提交**

```bash
gofmt -w module01_basics/blocks/01_go_basics/demo/03_control_funcs/main.go module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go
make module01-demo-contracts
go test ./module01_basics/blocks/01_go_basics/...
git diff --check
git add module01_basics/blocks/01_go_basics/demo/03_control_funcs/main.go module01_basics/blocks/01_go_basics/demo/05_strings_basics/main.go module01_basics/blocks/01_go_basics/README.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md module01_basics/instructor/scripts/verify_demo_contracts.sh
git commit -m "feat: deepen module01 control flow and string teaching"
```

Expected: Demo contract PASS，Block 1 包测试 PASS。

---

### Task 3：补齐 Slice copy 与嵌套 Map 初始化

**Files:**
- Modify: `module01_basics/blocks/02_collections/demo/04_arrays_slices/main.go:5-43`
- Modify: `module01_basics/blocks/02_collections/demo/05_maps_strings/main.go:8-58`
- Modify: `module01_basics/instructor/scripts/verify_demo_contracts.sh`
- Modify: `module01_basics/blocks/02_collections/README.md`
- Modify: `module01_basics/instructor/DEMO_NOTES.md`
- Modify: `module01_basics/instructor/RUNBOOK.md`

**Interfaces:**
- Consumes: 现有 Array/Slice/Map/String Demo 和输出契约脚本。
- Produces: `copy count=3 dst=[10 20 30 0 0]` 与 `nested map value: ready` 两个稳定教学结果。

- [ ] **Step 1: 先增加 Collections 输出断言**

在 `verify_demo_contracts.sh` 的 PASS 前加入：

```bash
slice_output="$(go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices)"
assert_contains "$slice_output" 'copy count=3 dst=[10 20 30 0 0] src=[10 20 30]'

map_output="$(go run ./module01_basics/blocks/02_collections/demo/05_maps_strings)"
assert_contains "$map_output" 'nested map value: ready'
```

Run: `make module01-demo-contracts`

Expected: FAIL，缺失 `copy count=3`。

- [ ] **Step 2: 在 Array/Slice Demo 增加 copy 对照**

在 `dynamicSlice` 输出后加入：

```go
source := []int{10, 20, 30}
destination := make([]int, 5)
copied := copy(destination, source)
fmt.Printf("copy count=%d dst=%v src=%v\n", copied, destination, source)
```

Expected teaching result: 参数顺序是 `copy(dst, src)`；返回复制数量 `3`；destination 剩余位置保持零值。

- [ ] **Step 3: 在 Map Demo 增加逐层初始化**

在普通 Map 删除/遍历后加入：

```go
nested := make(map[int]map[int]string)
nested[1] = make(map[int]string)
nested[1][2] = "ready"
fmt.Println("nested map value:", nested[1][2])
```

Expected teaching result: 只初始化外层 Map 不足以写入内层 key；每一层 Map 都必须初始化。

- [ ] **Step 4: 同步 Block 2 教案**

在 `DEMO_NOTES` 增加：

```markdown
| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| `copy(destination, source)` | 返回值和剩余两个位置是什么？ | 3，剩余为 0 | 复制数量为 min(len(dst), len(src)) | copy 会自动扩展 dst | 核心 |
| 嵌套 Map | 只 make 外层能否直接写内层？ | 不能，会 panic | 每层 Map 都要初始化 | 外层 make 会递归初始化 | 深挖 |
```

在 Block 2 README 增加 `copy` 方向、数量和嵌套 Map 初始化；Runbook 将 `copy` 保留为核心，嵌套 Map 归为深挖可裁。

- [ ] **Step 5: 验证并提交**

```bash
gofmt -w module01_basics/blocks/02_collections/demo/04_arrays_slices/main.go module01_basics/blocks/02_collections/demo/05_maps_strings/main.go
make module01-demo-contracts
go test ./module01_basics/blocks/02_collections/...
git diff --check
git add module01_basics/blocks/02_collections/demo/04_arrays_slices/main.go module01_basics/blocks/02_collections/demo/05_maps_strings/main.go module01_basics/blocks/02_collections/README.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md module01_basics/instructor/scripts/verify_demo_contracts.sh
git commit -m "feat: add module01 collection edge cases"
```

---

### Task 4：补齐指针字段语法糖和值接收者修改语义

**Files:**
- Modify: `module01_basics/blocks/03_modeling/demo/06_pointers/main.go:5-48`
- Modify: `module01_basics/blocks/03_modeling/demo/09_copy_and_receivers/main.go:11-48`
- Modify: `module01_basics/instructor/scripts/verify_demo_contracts.sh`
- Modify: `module01_basics/blocks/03_modeling/README.md`
- Modify: `module01_basics/instructor/DEMO_NOTES.md`
- Modify: `module01_basics/instructor/RUNBOOK.md`

**Interfaces:**
- Consumes: 现有 `Student.Label`、`Student.Rename` 和指针基础 Demo。
- Produces: 自动解引用与显式解引用等价输出；值接收者内部修改不影响调用方的直接输出。

- [ ] **Step 1: 增加 Modeling 输出断言**

在 `verify_demo_contracts.sh` 的 PASS 前加入：

```bash
pointer_output="$(go run ./module01_basics/blocks/03_modeling/demo/06_pointers)"
assert_contains "$pointer_output" 'pointer field sugar: 3'

receiver_output="$(go run ./module01_basics/blocks/03_modeling/demo/09_copy_and_receivers)"
assert_contains "$receiver_output" 'value receiver mutation keeps: Alice'
```

Run: `make module01-demo-contracts`

Expected: FAIL，缺失 `pointer field sugar: 3`。

- [ ] **Step 2: 增加 Struct 指针字段自动解引用**

在 `06_pointers/main.go` 包级加入：

```go
type Counter struct {
	Value int
}
```

在 nil pointer 演示后加入：

```go
counter := &Counter{Value: 1}
counter.Value = 2
(*counter).Value = 3
fmt.Println("pointer field sugar:", counter.Value)
```

Expected teaching result: `counter.Value` 是 Go 提供的字段访问语法糖，与 `(*counter).Value` 指向同一字段；这不意味着指针被复制成 Struct。

- [ ] **Step 3: 增加值接收者修改不生效的直接 Case**

在 `09_copy_and_receivers/main.go` 加入：

```go
func (s Student) RenameCopy(name string) {
	s.Name = name
}
```

在 `student` 创建后、指针接收者 `Rename` 调用前加入：

```go
student.RenameCopy("Nobody")
fmt.Println("value receiver mutation keeps:", student.Name)
```

Expected output: `value receiver mutation keeps: Alice`。

- [ ] **Step 4: 同步 Block 3 教案**

在 `DEMO_NOTES` 增加：

```markdown
| 暂停点 | 先问 | 预期结果 | 准确解释 | 常见误解 | 级别 |
| --- | --- | --- | --- | --- | --- |
| `counter.Value` / `(*counter).Value` | 是否修改同一字段？ | 最终为 3 | Go 自动解引用字段访问 | `p.field` 会复制对象 | 深挖 |
| `RenameCopy` | student.Name 会变吗？ | 仍为 Alice | 值接收者得到 Struct 副本 | 方法天然拥有 Java this 引用语义 | 核心 |
| `Rename` | 为什么变为 Bob？ | 指针接收者修改同一对象 | 传入的仍是值，只是该值为地址 | Go 存在“引用传递” | 核心 |
```

同步 Block 3 README；Runbook 保持值/指针接收者为核心，自动解引用语法糖为深挖可裁。

- [ ] **Step 5: 验证并提交**

```bash
gofmt -w module01_basics/blocks/03_modeling/demo/06_pointers/main.go module01_basics/blocks/03_modeling/demo/09_copy_and_receivers/main.go
make module01-demo-contracts
go test ./module01_basics/blocks/03_modeling/...
git diff --check
git add module01_basics/blocks/03_modeling/demo/06_pointers/main.go module01_basics/blocks/03_modeling/demo/09_copy_and_receivers/main.go module01_basics/blocks/03_modeling/README.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md module01_basics/instructor/scripts/verify_demo_contracts.sh
git commit -m "feat: clarify module01 pointer and receiver semantics"
```

---

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

### Task 6：同步 Exit Quiz、答案和讲师诊断语言

**Files:**
- Modify: `module01_basics/assessments/exit_quiz.md:1-180`
- Modify: `module01_basics/assessments/answer_key.md:20-90`
- Modify: `module01_basics/instructor/DEMO_NOTES.md`
- Modify: `module01_basics/instructor/RUNBOOK.md`

**Interfaces:**
- Consumes: Tasks 1–5 的实际输出和教学失败诊断。
- Produces: 17 题 Exit Quiz；每题唯一答案和对应讲师诊断；5 分钟核心题与选做题边界。

- [ ] **Step 1: 增加 iota 行为题**

在 Exit Quiz 增加第 15 题：

````markdown
## 15. ConstSpec 与 iota

```go
const (
    a = iota
    b
    c = 100
    d
    e = iota
)
```

`a, b, c, d, e` 分别是什么？

- A. `0, 1, 100, 100, 4`
- B. `0, 1, 100, 3, 4`
- C. `0, 1, 100, 100, 0`
- D. `1, 2, 100, 100, 5`
````

答案 A。诊断语言：选 B 表示没有理解省略声明复用完整表达式；选 C 表示误以为同一 const 块内 iota 被重置。

- [ ] **Step 2: 增加 copy 行为题**

增加第 16 题：

````markdown
## 16. Slice copy

```go
src := []int{1, 2, 3}
dst := make([]int, 5)
n := copy(dst, src)
```

哪个结果正确？

- A. `n=5, dst=[1 2 3 0 0]`
- B. `n=3, dst=[1 2 3 0 0]`
- C. `n=3, dst=[0 0 1 2 3]`
- D. copy 会自动把 dst 长度改成 3
````

答案 B。诊断语言：错误选择说明学员未掌握 `copy(dst, src)` 方向或 `min(len(dst), len(src))`。

- [ ] **Step 3: 增加编译失败诊断题**

增加第 17 题：

````markdown
## 17. Map Struct Value

```go
type Item struct{ Value int }
items := map[string]Item{"one": {Value: 1}}
items["one"].Value = 2
```

这段代码的结果是什么？

- A. Value 被改为 2
- B. 运行时 panic
- C. 编译失败，因为 Map index 得到的 Struct value 不可寻址
- D. 编译成功但赋值被静默忽略
````

答案 C。诊断语言：选 A 表示把 Map value 当成可寻址对象；选 B 表示不能区分编译错误与运行时 panic。

- [ ] **Step 4: 保持 5 分钟时间盒**

将标题改为“核心 14 题 + 选做 3 题”。Runbook 规定课堂 5 分钟先完成 1–14 题；15–17 题用于提前完成者或课后 Code Review，不挤占 Task Manager 启动时间。

- [ ] **Step 5: 校验题号、答案并提交**

```bash
rg -n '^## (15|16|17)\.' module01_basics/assessments/exit_quiz.md
rg -n '^\| (15|16|17) \|' module01_basics/assessments/answer_key.md
git diff --check
git add module01_basics/assessments/exit_quiz.md module01_basics/assessments/answer_key.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md
git commit -m "docs: extend module01 semantic assessments"
```

Expected: Quiz 和答案均各出现 15、16、17 三个题号；无缺号或重复号。

---

### Task 7：把 Demo Notes 完善为可直接讲课的小抄

**Files:**
- Modify: `module01_basics/instructor/DEMO_NOTES.md:1-185`
- Modify: `module01_basics/instructor/RUNBOOK.md:1-160`
- Modify: `module01_basics/instructor/RUBRIC.md`
- Modify: `module01_basics/README.md`

**Interfaces:**
- Consumes: Tasks 1–6 的全部 Demo 输出、失败 Case 和测评题。
- Produces: 每个知识点都有现场提问、预期结果、准确解释、误解、救援话术和裁剪优先级的讲师入口。

- [ ] **Step 1: 为每个 Block 增加“不得讲错”清单**

在 DEMO_NOTES 四个 Block 分别加入：

```markdown
### 不得讲错

- Go 所有参数传递都是值传递；传指针时复制的是地址值。
- nil Slice 可以 append；nil Map 可以读但不能写。
- iota 始终按 ConstSpec 行递增；表达式省略与 iota 计数是两个维度。
- String 下标和切片按 byte；range 解码 rune，index 仍是 byte offset。
- append 可能复用也可能更换底层数组；不能承诺固定扩容倍数。
- Map 遍历顺序不是稳定契约。
- 值接收者操作副本；指针接收者用于修改同一对象，但仍是值传递。
- defer 普通参数在注册时求值；闭包体读取发生在执行时。
```

实际写入时按所属 Block 分散这些条目，不在每个 Block 重复全部八条。

- [ ] **Step 2: 为每个 Demo 增加精确运行结果**

每个 Demo 命令后增加“必须看到”的关键行，不复制全部输出：

```markdown
- `02_vars_types`：`0 1 2 250 250 5 6`、`iota reset: 0 1`、零值行。
- `03_control_funcs`：`owner has full access`、fallthrough 的两行、循环只打印 0 和 2。
- `05_strings_basics`：`A你` 为 4 bytes、range 两次、Atoi 错误为 true。
- `04_arrays_slices`：Array 副本不变、子 Slice 共享、copy 返回 3。
- `05_maps_strings`：comma-ok 区分零值和不存在、嵌套 Map 输出 ready。
- `06_pointers`：值参数不变、指针参数改变、字段语法糖最终为 3。
- `09_copy_and_receivers`：RenameCopy 后仍为 Alice、指针 Rename 后为 Bob。
- `10_function_forms`：两个 counter 状态独立。
- `11_defer_edges`：LIFO、普通参数读取 1、闭包读取 2。
```

- [ ] **Step 3: 增加讲师救援话术**

在每个知识点表后使用统一格式：

```markdown
**一级提示：**让学员只预测当前一行输出。

**二级提示：**指出当前操作的是值、地址、底层数组、Map key 或闭包捕获变量中的哪一个。

**三级提示：**运行最小 Case，只解释 actual/want 或编译器第一条诊断，不直接展示 Solution。
```

对于受控失败 Case，三级提示改为展示修正版本的最小差异，不修改失败 fixture 本身。

- [ ] **Step 4: 校准 Runbook 裁剪顺序**

Runbook 必须明确以下不可裁内容：

```markdown
- Block 1：隐式 ConstSpec/iota、多值 case、默认不贯穿、fallthrough 语义、零值、Atoi 错误路径。
- Block 2：Array/Slice 值与共享语义、copy、comma-ok、byte/rune。
- Block 3：Go 只有值传递、值/指针接收者、校验后修改、Snapshot 隔离。
- Block 4：函数值、闭包捕获、defer 顺序、读取失败测试。
```

可裁顺序固定为：位运算 → iota 重置块 → 类型别名 → switch 初始化作用域 → 嵌套 Map → 指针字段语法糖 → 多 defer 扩展。不得用删学员动手时间来容纳深挖内容。

- [ ] **Step 5: 更新 Rubric 和 Module 01 总览**

Rubric 的“解释与迁移”增加两项可观察证据：

```markdown
- 能区分“编译失败、运行时 panic、测试断言失败”三类反馈，并说明下一步动作。
- 能先预测关键输出，再用语言规则解释结果，而不是只复述运行结果。
```

Module 01 README 增加 `make module01-demo-contracts`、`make module01-teaching-failures`、`make module01-audit` 的用途和预期。

- [ ] **Step 6: 做教案覆盖审计并提交**

```bash
rg -n 'iota|空白标识符|定义类型|break|continue|fallthrough|Atoi|copy\(|嵌套 Map|值接收者|指针接收者|编译失败|panic' module01_basics/instructor/DEMO_NOTES.md
rg -n '不可裁|可裁顺序|185|310' module01_basics/instructor/RUNBOOK.md
git diff --check
git add module01_basics/README.md module01_basics/instructor/DEMO_NOTES.md module01_basics/instructor/RUNBOOK.md module01_basics/instructor/RUBRIC.md
git commit -m "docs: turn module01 notes into complete instructor cheat sheet"
```

Expected: 每个搜索词至少命中一条讲师说明；Runbook 仍包含 185 分钟动手和 310 分钟净课程时间。

---

### Task 8：完整验收 Module 01，证明 GREEN、RED 与预期失败均正确

**Files:**
- Modify only if verification reveals a concrete defect: files already listed in Tasks 1–7

**Interfaces:**
- Consumes: `module01-verify`、`module01-demo-contracts`、`module01-teaching-failures`、四个 Starter 和教师答案。
- Produces: 可重复的 Module 01 完成证据；不触碰 Module 02 及以后代码。

- [ ] **Step 1: 验证格式、正常代码和教师答案 GREEN**

```bash
make module01-verify
```

Expected:

- Module 01 所有正常 Go 文件通过 gofmt 检查。
- `go vet ./module01_basics/...` PASS。
- `go test ./module01_basics/...` PASS。
- Task Manager teacher solution grade PASS。

- [ ] **Step 2: 验证 Demo 教学输出契约**

```bash
make module01-demo-contracts
```

Expected: `module01 demo contracts: PASS`。

- [ ] **Step 3: 验证故意错误按预期失败**

```bash
make module01-teaching-failures
```

Expected: 八个 Case 都打印 `expected failure`，最后打印 `module01 teaching failures: PASS`。

- [ ] **Step 4: 验证 Starter 仍保持正确 RED**

分别运行：

```bash
go test -tags=exercise ./module01_basics/blocks/01_go_basics/lab/starter
go test -tags=exercise ./module01_basics/blocks/02_collections/lab/starter
go test -tags=exercise ./module01_basics/blocks/03_modeling/lab/starter
go test -tags=exercise ./module01_basics/blocks/04_functions_testing/lab/starter
```

Expected: 未完成的 Starter 返回非零，失败来自约定行为或断言，不是编译错误、缺依赖或路径错误。记录每个包第一条失败测试名到验收日志；不得把 Starter 改成假 GREEN。

- [ ] **Step 5: 验证总审计和范围边界**

```bash
make module01-audit
git diff --check
git status --short
git diff --name-only HEAD~7..HEAD | rg '^module02_' && exit 1 || true
```

Expected: `module01-audit` PASS；无空白错误；提交范围不包含 `module02_advanced` 或之后模块。

- [ ] **Step 6: 最终提交验收修正**

只有 Step 1–5 暴露具体缺陷并产生修正时才创建本提交：

```bash
git add Makefile module01_basics
git commit -m "test: complete module01 curriculum audit"
```

若没有修正，保留 Tasks 1–7 的提交序列，不创建空提交。

---

## 自检结果

### 需求覆盖

- 只补 Module 01：Tasks 1–8 均限定 `module01_basics` 和根 Makefile；范围检查禁止 Module 02 变更。
- 语言细节补充：Task 1 覆盖变量/常量/iota/类型；Task 2 覆盖控制流/String；Task 3 覆盖 Collections；Task 4 覆盖 Modeling。
- 讲师小抄同步：每个实现 Task 都修改 Demo Notes，Task 7 做完整教案审计。
- 不盲目追求全绿：Task 5 建立受控失败路径，Task 8 分别验证 GREEN、RED 和 expected failure。
- 保留用户代码：Task 1 明确消费并修正现有 `consA` 到 `consG`，不覆盖该片段。
- 不碎片化 Demo：所有正常语义合入现有 Demo；只有故意错误进入 `testdata` fixtures。

### 占位符检查

本计划没有未定义实现、空白步骤或待定文件；每个代码变更都给出具体代码、命令和预期结果。

### 接口一致性

- Make 目标统一为 `module01-demo-contracts`、`module01-teaching-failures`、`module01-audit`。
- 输出契约脚本统一使用 `assert_contains`。
- 教学失败脚本统一使用 `expect_failure command_kind package_path expected_pattern`。
- 讲师文档统一使用“先问、预期结果、准确解释、常见误解、级别”五类信息。
