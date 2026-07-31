# Module 02 测评答案

## Entry Quiz

1. 不需要；方法集满足接口即可隐式实现。
2. 不成立。接口保存了动态类型 `*T`，因此接口值本身不为 nil。
3. 无缓冲 Channel 的发送与接收方完成同步交接时完成。
4. 在启动 goroutine 前调用，避免 `Wait` 与 `Add` 并发造成生命周期错误。
5. 把 Context 作为第一个参数显式传给下游，并在阻塞点监听 `Done`。
6. 多 goroutine 未同步访问同一内存导致的数据竞态。

## Exit Quiz

1. 接口由动态类型和动态值组成；typed nil 仍携带非 nil 动态类型。
2. 发送方知道不会再产生值；接收方擅自关闭可能让仍在发送的 goroutine panic。
3. `errors.Is` 判断链上某个错误值，`errors.As` 提取链上某种错误类型。
4. 取消只是信号；任务必须主动响应，并由调用方等待退出，才能证明没有泄漏。
5. 普通互斥场景优先 Mutex；只有测量证明读多写少时才考虑 RWMutex，sync.Map 只适合特定访问模式。
6. 非指针值通常不可设置；修改前必须取得可寻址的 `Elem` 并检查 `CanSet`。
7. 结果 Channel 仍可能被 worker 写入；必须等待全部 worker 后由协调者关闭。
8. nil Channel 的收发永久阻塞，因此可以把某个 case 动态禁用。
9. 它保证给定 Once 实例关联的函数至多成功进入一次；不替代错误重试、Context 或其他状态同步设计。
10. 未导出字段可能可以被反射观察，但调用 `Interface` 会 panic；必须先检查 `CanInterface`。
