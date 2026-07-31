# Module 02 讲师 Runbook（09:30–16:00）

本 Runbook 是 Module 02 的唯一讲师入口。课前同时打开 [Demo Notes](DEMO_NOTES.md)、[课堂评分表](RUBRIC.md) 和 [测评答案](../assessments/answer_key.md)。所有命令默认从仓库根目录运行。

## 授课边界

- 面向已完成 Module 01、具备 Java 阅读经验的学员。
- 核心课堂净时间为 310 分钟，学员动手时间不少于 155 分钟。
- 允许 Demo 使用固定数据、Sleep 和内存对象；讲解时必须标注这是教学简化。
- Reflection 保留在主流程；Embed、Generate、OS、文件 I/O、runtime 和标准库扩展进入 Bonus。

## 学员动手预算

| 阶段 | 动手分钟 |
|---|---:|
| 环境与 Entry Quiz | 10 |
| Block 1 | 25 |
| Block 2 | 40 |
| Block 3 | 35 |
| Block 4 | 25 |
| 综合 Lab | 35 |
| Homework 启动与 Exit Quiz | 10 |
| **合计** | **180 / 310（58.1%）** |

延误时先裁剪 Bonus 和扩展讲解，不压缩 Starter、综合 Lab、Quiz 或两次休息。

## 全日日程

| 时间 | 内容 | 可观察产出 |
|---|---|---|
| 09:30–09:50 | 环境检查、Entry Quiz | 能运行 Module 02 验收命令 |
| 09:50–10:40 | Block 1：接口、错误与恢复边界 | 完成接口与错误链练习 |
| 10:40–10:50 | 短休息 | — |
| 10:50–12:00 | Block 2：Goroutine 与 Channel | 完成可关闭的 worker pipeline |
| 12:00–13:00 | 午休 | — |
| 13:00–14:05 | Block 3：Context 与并发安全 | 完成可取消且无竞态的计数器 |
| 14:05–14:50 | Block 4：Testing 与 Reflection | 完成表格测试、Benchmark 和反射练习 |
| 14:50–15:00 | 短休息 | — |
| 15:00–15:40 | 并发日志分析器综合 Lab | 通过分阶段项目验收 |
| 15:40–16:00 | Homework、Exit Quiz | 跑通作业 Starter 和提交流程 |

## 讲授节奏

每个 Block 均按以下顺序推进：

1. 学习结果和 Java 对比；
2. 最小 Demo；
3. 学员 Starter RED；
4. 分阶段 Lab；
5. 自动验收；
6. 常见错误复盘；
7. 迁移问题。

四个 Block 的 Starter 必须使用显式 `-tags=exercise` 运行；默认仓库测试只包含可运行 Demo 和教师 Solution。

## 课堂验收

```bash
make module02-demo-contracts
make module02-integrated-lab
make module02-verify
```

故意 panic、死锁和竞态的示例不进入默认验收，只能通过显式的教学失败命令运行。

## 分段讲授入口

- 09:50：运行 Block 1 的接口与错误 Demo，25 分钟完成 `Shape` 和 `ParsePort` Starter。
- 10:50：运行 Block 2 的 Goroutine 与 Channel Demo，40 分钟完成 `CollectSquares`。
- 13:00：运行 Block 3 的 Context 与并发安全 Demo，35 分钟内用 `-race` 验证计数器。
- 14:05：运行 Block 4 的测试与 Reflection Demo，25 分钟完成表格测试驱动的字段读取。
- 15:00：严格按 [综合 Lab 六个检查点](../integrated_lab/log_analyzer/README.md)推进。
- 15:40：只演示 [文件扫描器 Starter](../homework/file_scanner/student_pack/README.md) 的 RED、验收和提交流程，不现场实现答案。

## 受控失败边界

课堂最多选择两个失败案例投影。必须运行 `make module02-teaching-failures`，不能复制 panic、死锁、竞态或并发 Map fatal 到正常 Demo 进程。
