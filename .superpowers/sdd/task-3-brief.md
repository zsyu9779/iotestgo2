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

