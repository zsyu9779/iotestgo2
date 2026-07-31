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

