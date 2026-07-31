# Module 02 Exit Quiz

1. 为什么 typed-nil 接口不等于 nil？
2. 为什么通常由发送方关闭 channel？
3. `errors.Is` 和 `errors.As` 的区别是什么？
4. Context 取消后，为什么不能只依赖函数返回而不确认 goroutine 已结束？
5. 什么时候应该优先使用 Mutex，而不是 RWMutex 或 sync.Map？
6. Reflection 修改值时为什么通常需要传指针并检查 `CanSet`？
7. 综合日志分析器中，为什么错误结果需要在消费者结束后再关闭？
8. nil Channel 在 `select` 中有什么用途？
9. `sync.Once` 保证什么，不保证什么？
10. `CanInterface` 为什么对未导出字段重要？
