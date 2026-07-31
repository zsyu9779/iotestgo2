# Module01 Task 7 完成报告

## 修改文件

- `module01_basics/instructor/DEMO_NOTES.md`
- `module01_basics/instructor/RUNBOOK.md`
- `module01_basics/instructor/RUBRIC.md`
- `module01_basics/README.md`

本报告位于 `.superpowers/`，按任务要求不纳入 Git 提交。

## 覆盖范围

- 四个 Block 均补充“不得讲错”清单，并按知识归属分散说明：ConstSpec/iota、byte/rune 与 nil 集合、Slice/Map 语义、值传递与接收者、defer 求值。
- 对任务指定的九个 Demo 写明现场必须看到的关键行：iota/零值、控制流与 fallthrough、UTF-8/Atoi、Array/Slice/copy、Map/comma-ok/嵌套 Map、指针和值/指针接收者、闭包状态和 defer LIFO/求值时机。
- 四个知识点表均增加统一的一级、二级、三级救援话术；受控失败的三级提示改为只展示正确写法的最小差异，不改 fixture、不展示 Solution。
- Runbook 明确四个 Block 的不可裁内容、固定可裁顺序，以及不占用 185 分钟学员动手时间的约束；保留 310 分钟净课程时间。
- Rubric 的“解释与迁移”增加三类反馈分类与“先预测、再按语言规则解释”两项可观察证据。
- Module01 README 说明 `module01-demo-contracts`、`module01-teaching-failures`、`module01-audit` 的用途和预期成功标记。

## 核对命令与结果

```bash
rg -n 'iota|空白标识符|定义类型|break|continue|fallthrough|Atoi|copy\(|嵌套 Map|值接收者|指针接收者|编译失败|panic' module01_basics/instructor/DEMO_NOTES.md
rg -n '不可裁|可裁顺序|185|310' module01_basics/instructor/RUNBOOK.md
git diff --check
make module01-audit
```

结果：两次 `rg` 均命中所有要求的讲师说明与时间/裁剪约束；`git diff --check` 无输出且退出 0；`make module01-audit` 退出 0，完成 `go vet ./module01_basics/...`、`go test ./module01_basics/...`、Task Manager 教师答案验收、Demo 输出契约，以及全部 8 个受控失败 Case（末行分别为 `module01 demo contracts: PASS` 与 `module01 teaching failures: PASS`）。

另外逐个运行了任务指定的九个 Demo，记录的“必须看到”行与当前实际输出一致。

## 疑虑

无阻塞疑虑。受控失败 Case 的正确性仍依赖 Go 编译器诊断正则；当前环境已由 `make module01-teaching-failures` 验证匹配。

## 提交

已提交四份 Module01 文档：`55107df docs: turn module01 notes into complete instructor cheat sheet`。`.superpowers/` 未暂存、未提交。

## 审查修复（后续提交）

- 更正 Module01 README：教学失败的每个 fixture 预期非零且诊断匹配，但聚合命令在全部验证成功时返回 0。
- 运行并记录五个原先无明确锚点的课堂 Demo 的实际稳定输出，并在 Demo Notes 增加对应救援话术：`06_slice_map_edges`、`07_string_utf8_edges`、`07_structs_methods`、`08_struct_zero_values`、`09_advanced_functions`。
- 未修改示例代码；`make module01-audit` 与 `git diff --check` 已复验通过。
- 已提交文档修复：`82f9fd3 docs: clarify module01 teaching verification`；`.superpowers/` 仍未暂存、未提交。

## 聚合退出状态修复（后续提交）

- 将 DEMO_NOTES、README 和 RUNBOOK 统一为准确语义：每个 fixture 预期非零且诊断匹配；全部 fixture 匹配时，聚合命令 `make module01-teaching-failures` 返回 0。
- 已重新运行 `make module01-audit`、`git diff --check`，并检索三份文档确认不再有相反表述。
- 已提交最小文档修复：`f4e24bc docs: clarify teaching failure exit status`；`.superpowers/` 仍未暂存、未提交。

## 五字段覆盖修复（后续提交）

- Demo Notes 新增 String/`[]byte` 往返的五字段讲师条目，预期输出与 `05_strings_basics` 的 `bytes round trip: Go语言` 一致。
- 受控失败部分逐项覆盖 package 短声明、同作用域无新变量、定义类型赋值、Slice 比较、Map Struct 字段、最后一个 `fallthrough`、nil Map 写入和 nil 指针解引用；每项都含提问、预期诊断、准确解释、常见误解和级别，诊断与 `teaching_failures/verify.sh` 的实际匹配模式一致。
- `.superpowers/sdd/task-8-report.md` 的范围审计基线已改为 `86715a8..3f783f8`，不纳入提交。
- `make module01-audit` 和 `git diff --check` 已通过；已提交 `68fc3aa docs: complete module01 failure teaching notes`，仅包含 `module01_basics/instructor/DEMO_NOTES.md`。
