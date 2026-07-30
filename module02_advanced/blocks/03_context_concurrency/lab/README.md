# Block 3 Lab

运行公开测试：

```bash
go test -race ./module02_advanced/blocks/03_context_concurrency/lab
```

任务是让长任务响应 Context 取消，并使用互斥锁保护共享计数器。
