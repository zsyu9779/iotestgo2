# Module 02 受控失败案例

这些程序故意 panic、死锁、产生竞态或触发 runtime fatal，只能通过统一验证脚本隔离运行：

```bash
make module02-teaching-failures
```

成功含义是：每个子程序以非零状态退出，并且诊断与预期原因匹配；聚合脚本本身最终返回 0。不要直接把这些目录加入正常 Demo 或课堂 Solution。
