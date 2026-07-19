# Module 01 Bonus

这里集中存放不占用当天 310 分钟核心课程的扩展材料。请先完成四个 Block、Scorebook 综合 Lab 和 Task Manager 作业启动，再按兴趣选择：

- [数据结构](data_structures/main.go)：用链表和栈继续练习 Struct、Pointer 与边界处理。
- [Generics](generics/main.go)：体验类型参数；它不属于兼容 Go 1.16 的 Week 1 作业要求。
- Dark Corners：[range](dark_corners/range/main.go)、[map](dark_corners/map/main.go) 与 [string/UTF-8](dark_corners/string/main.go)。
- Function Patterns：[完整演示](function_patterns/main.go)、[配置模式说明](function_patterns/configuration_patterns.md)、[柯里化真实场景](function_patterns/curry_best_practice_test.go) 与 [Builder/Functional Options 对比](function_patterns/patterns_comparison_test.go)。

从仓库根目录独立运行各程序：

```bash
go run ./module01_basics/bonus/data_structures
go run ./module01_basics/bonus/generics
go run ./module01_basics/bonus/dark_corners/range
go run ./module01_basics/bonus/dark_corners/map
go run ./module01_basics/bonus/dark_corners/string
go run ./module01_basics/bonus/function_patterns
```

一次验证全部 Bonus 包：

```bash
go test ./module01_basics/bonus/...
```
