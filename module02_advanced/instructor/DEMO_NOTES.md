# Module 02 Demo Notes

## Block 1：接口、错误与恢复边界

- 先运行 `go run ./module02_advanced/blocks/01_interfaces_errors/demo/01_interfaces`，展示隐式实现、接口组合、类型断言和方法集矩阵。
- 再单独说明 typed-nil：接口同时包含动态类型和动态值。
- 使用 `go run ./module02_advanced/blocks/01_interfaces_errors/demo/02_errors_defer` 展示 `%w`、`errors.Is` 与 `errors.As`；让学员从少量结果反推错误链，而不是逐行阅读打印讲稿。
- 强调 `panic/recover` 是边界兜底，不是业务错误处理机制；本 Block 不重复讲解 defer 的执行顺序与资源释放。

## Block 2：Goroutine 与 Channel

- 先演示主函数提前退出，再使用 WaitGroup 修复生命周期。
- 让学员观察发送方关闭 channel、接收方使用 range 的完整协议。
- 解释 nil channel、关闭 channel 写入和 select 超时的区别。
- 强调多个就绪 case 的选择不确定；带 `default` 的循环必须避免忙等。

## Block 3：Context 与并发安全

- 用 Context 取消长任务，要求任务主动监听 `Done`。
- 使用 `go test -race` 验证无锁计数器和修复版本的差异。
- 对比 Mutex、RWMutex、Atomic、Once 和 sync.Map 的适用边界。

## Block 4：Testing 与 Reflection

- 先让公开测试 RED，再实现最小行为使其 GREEN。
- 用 Benchmark 说明“可测量后再优化”。
- Reflection 演示 Type、Value、Elem、CanSet、CanInterface 和值/指针方法集，并说明 panic 边界。

## 教学简化声明

当前 Demo 中的固定数据、Sleep、打印输出和内存 channel 主要用于降低课堂认知负担。它们不是生产系统的完整设计；生产实现需要处理错误、取消、资源释放、日志和可观测性。
