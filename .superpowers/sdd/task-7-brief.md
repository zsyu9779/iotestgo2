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

