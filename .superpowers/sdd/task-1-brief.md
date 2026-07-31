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

