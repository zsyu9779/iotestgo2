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
