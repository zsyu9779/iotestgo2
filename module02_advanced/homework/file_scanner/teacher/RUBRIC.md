# 文件扫描器作业评分表

| 维度 | 分值 |
|---|---:|
| 参数和错误契约 | 15 |
| 文件遍历与匹配统计 | 25 |
| 固定 worker pool 与关闭顺序 | 25 |
| Context 取消与无 goroutine 泄漏 | 15 |
| 表格测试和 race detector | 10 |
| 格式、命名、README 与提交记录 | 10 |

出现死锁、数据竞态或通过启动无限 goroutine 规避固定 worker 要求时，并发部分不得分。
