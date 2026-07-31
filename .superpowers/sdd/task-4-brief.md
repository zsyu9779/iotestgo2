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

