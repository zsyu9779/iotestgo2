# Block 4：Testing 与 Reflection

## 学习结果

完成后能够编写表格驱动测试、Benchmark，并理解反射的 Type、Value、CanSet 和方法调用。

## 时间盒

45 分钟，其中学员动手不少于 25 分钟。

## 前置知识

前面三个 Block 和 Module 01 测试基础。

## Java 对比

Go testing 不需要额外测试框架；Reflection 可类比 Java Reflection，但必须主动处理类型和可设置性。

## 讲师 Demo

```bash
go test -v ./module02_advanced/blocks/04_testing_reflection/demo/07_testing
go run ./module02_advanced/blocks/04_testing_reflection/demo/10_reflection
```

## 学员任务

用表格测试驱动 `ReadFieldName`，区分字段不存在、字段类型错误和合法 String 字段。

## 验收命令

```bash
go test -tags=exercise ./module02_advanced/blocks/04_testing_reflection/lab/starter
```

## 常见错误

- 测试只覆盖正常路径；
- Benchmark 中没有使用 `b.N`；
- 反射值不可设置时直接 Set；
- 把反射当作普通业务代码的首选方案。

## 三级提示

1. 先列出输入、输出和边界条件。
2. 用表格和子测试组织用例。
3. 修改前检查 Kind、Type 和 CanSet。

## 复盘问题

- 测试如何证明 Context 取消后 goroutine 已退出？
- 反射给框架带来了什么能力和代价？

## Bonus

阅读 `bonus/testing_tools` 中的 DeepEqual 与 pprof 示例。
