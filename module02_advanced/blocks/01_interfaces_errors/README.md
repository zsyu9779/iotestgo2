# Block 1：接口、错误与恢复边界

## 学习结果

完成后能够使用隐式接口、类型断言和错误包装编写可读的边界代码，使用 `io.Reader` / `io.Writer` 完成基础文件读写，并理解 `recover` 必须由 `defer` 在边界处调用。

## 时间盒

50 分钟，其中学员动手不少于 25 分钟。

## 前置知识

Module 01 的函数、结构体、指针和测试基础。

## Java 对比

Go 接口描述行为，不需要 `implements`；错误通常是返回值，不使用异常作为普通业务流程。

## 讲师 Demo

```bash
go run ./module02_advanced/blocks/01_interfaces_errors/demo/01_interfaces
go run ./module02_advanced/blocks/01_interfaces_errors/demo/02_errors_defer
go run ./module02_advanced/blocks/01_interfaces_errors/demo/03_file_io
```

## 学员任务

实现 `Shape` 接口、`ParsePort` 错误返回，并用 `errors.As` 检查包装链。

## 验收命令

```bash
go test -tags=exercise ./module02_advanced/blocks/01_interfaces_errors/lab/starter
```

## 常见错误

- 把带类型的 nil 指针直接返回给接口；
- 用 panic 表达普通输入错误；
- 忘记 `%w` 导致错误链断裂；
- 试图在普通业务流程中使用 panic/recover。

## 三级提示

1. 先写出接口方法集和函数返回签名。
2. 检查返回接口时是否直接返回 nil。
3. 用 `errors.Is` 或 `errors.As` 验证链上的原始错误。

## 复盘问题

- 哪些错误应该由调用方处理？
- recover 应该放在库代码还是边界代码？

`defer` 的求值、执行顺序和资源释放已在 Module 01 学过；本 Block 只保留其配合 `recover` 的边界用法。
