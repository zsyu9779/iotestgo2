# Module 02 Demo Notes

## Block 1：接口、错误与恢复边界

- 先运行 `go run ./module02_advanced/blocks/01_interfaces_errors/demo/01_interfaces`，展示隐式实现、接口组合、类型断言和方法集矩阵。
- 再单独说明 typed-nil：接口同时包含动态类型和动态值。
- 使用 `go run ./module02_advanced/blocks/01_interfaces_errors/demo/02_errors_defer` 展示 `%w`、`errors.Is` 与 `errors.As`；让学员从少量结果反推错误链，而不是逐行阅读打印讲稿。
- 强调 `panic/recover` 是边界兜底，不是业务错误处理机制；本 Block 不重复讲解 defer 的执行顺序与资源释放。
- 使用 `go run ./module02_advanced/blocks/01_interfaces_errors/demo/03_file_io` 连接 `io.Reader` / `io.Writer`、缓冲读写和文件错误边界；输入与输出文件均在 `03_file_io/main.go` 同级目录，可直接打开对照。

## Block 2：Goroutine 与 Channel

- 先演示主函数提前退出，再使用 WaitGroup 修复生命周期。
- 运行 `go run ./module02_advanced/blocks/02_goroutines_channels/demo/04_channels`，从 `GenerateNatural` 开始逐级画出 `Filter(2) -> Filter(3) -> Filter(5)`。
- 精讲 `prime := <-ch` 为什么得到素数，以及 `ch = PrimeFilter(ch, prime)` 如何动态扩展并发流水线。
- 用 `25`、`27`、`29` 追踪不同数字经过过滤器的路径，解释无缓冲 channel 的同步和背压。
- 明确这不是欧拉线性筛：每个数字可能经过多个过滤器；每发现一个素数还会新增一个 goroutine 和 channel。
- 课堂版本依赖进程退出回收无限流水线；迁移到长期运行服务时必须增加取消和完整退出协议。

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
