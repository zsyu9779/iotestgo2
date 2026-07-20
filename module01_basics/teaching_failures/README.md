# Module 01 受控失败教学 Case

这些 Case 用来演示 Go 的编译诊断和运行时 panic。它们全部位于 `testdata/`，因此正常的 `go test ./module01_basics/...` 不会扫描或执行它们。不要把这些源码移入正常 Demo 或正常测试包。

从仓库根目录运行：

```bash
make module01-teaching-failures
```

通过的定义不是命令成功退出：每个 Case 必须以非零状态失败，且输出匹配下表诊断。这样才能区分预期教学失败与环境、路径或工具链故障。

| Case 路径 | 单独运行命令 | 预期诊断 | 正确写法 | 核心 |
| --- | --- | --- | --- | --- |
| `testdata/compile_fail/package_short_decl` | `go test ./module01_basics/teaching_failures/testdata/compile_fail/package_short_decl` | `outside function body` 或 `syntax error` | 在函数体中使用 `value := 1`，包级使用 `var value = 1` | 否 |
| `testdata/compile_fail/no_new_variable` | `go test ./module01_basics/teaching_failures/testdata/compile_fail/no_new_variable` | `no new variables` | 已有变量改用 `value = 2` | 是 |
| `testdata/compile_fail/defined_type_assignment` | `go test ./module01_basics/teaching_failures/testdata/compile_fail/defined_type_assignment` | `cannot use raw` | `var id UserID = UserID(raw)` | 否 |
| `testdata/compile_fail/slice_comparison` | `go test ./module01_basics/teaching_failures/testdata/compile_fail/slice_comparison` | `slice can only be compared to nil` 或 `invalid operation` | 比较元素、使用 `slices.Equal`，或只与 `nil` 比较 | 否 |
| `testdata/compile_fail/map_struct_field` | `go test ./module01_basics/teaching_failures/testdata/compile_fail/map_struct_field` | `cannot assign to struct field` | 取出值、修改后写回：`item := items["one"]; item.Value = 2; items["one"] = item` | 是 |
| `testdata/compile_fail/final_fallthrough` | `go test ./module01_basics/teaching_failures/testdata/compile_fail/final_fallthrough` | `cannot fallthrough final case` | 删除最后一个 `fallthrough`，或添加可落入的后续 case | 否 |
| `testdata/runtime_fail/nil_map_write` | `go run ./module01_basics/teaching_failures/testdata/runtime_fail/nil_map_write` | `assignment to entry in nil map` | `scores := make(map[string]int)` 后再写入 | 是 |
| `testdata/runtime_fail/nil_pointer` | `go run ./module01_basics/teaching_failures/testdata/runtime_fail/nil_pointer` | `nil pointer dereference` 或 `invalid memory address` | 先分配并赋地址：`value := new(int); *value = 1` | 否 |

课堂只投影三个核心 Case：同作用域无新变量的 `:=`、Map 的 Struct 字段不可直接赋值，以及 nil Map 写入 panic。其余 Case 是讲师根据班级情况调用的小抄，不逐个占用课堂时间。
