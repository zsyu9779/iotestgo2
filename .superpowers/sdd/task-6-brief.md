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

