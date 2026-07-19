# 新老教学项目逐项对比与改进建议

## 总体取舍

| 范围 | 对比结果 | 采用策略 |
| --- | --- | --- |
| Module 1：Go 基础 | 老项目语言细节更全面；新项目课程结构、Lab、Quiz 更好 | 保留新结构，系统迁入老项目的语义 Case |
| Module 2：Go 进阶 | 老项目在接口、反射、并发、I/O、运行时实验上更细；新项目已有工程化分章和综合项目 | 保留新结构，按专题补齐老项目细节并淘汰过时写法 |
| Module 3：net 包 | 老项目覆盖 TCP、UDP、广播、组播、HTTP Client/Server、TLS、HTTP/2、上传、WebSocket，明显更全面 | 重建 `01_net_basics`，以老项目覆盖面为基线，以现代工程实践重写 |
| Module 3：Gin | 新项目的路由、配置、鉴权、日志、API 设计、测试、分层项目更好 | 保留新项目，仅补生产级边界 |
| Module 4：GORM | 老项目没有等价体系；新项目组织方式明显更好 | 不从老项目迁移，继续工程化加固 |
| Module 5：gRPC | 老项目没有等价体系；新项目从 IDL 到综合项目的路径更合理 | 不从老项目迁移，继续工程化加固 |
| Module 6：go-zero | 老项目没有等价体系；新项目双教学线更合理 | 不从老项目迁移，继续工程化加固 |

## 验证原则：不是全绿，而是结果符合教学预期

| 内容类型 | 正确状态 | 处理方式 |
| --- | --- | --- |
| Solution、工程项目、公共库 | 必须通过编译、测试、vet/race 等对应验收 | 纳入默认 `make verify` 和 CI |
| Starter 学员练习 | 课前可以 RED，学员完成后应 GREEN | 使用 exercise build tag 或独立目录；文档明确预期失败测试名 |
| 编译失败教学 Case | 必须编译失败，并且失败原因与课程目标一致 | 放入 `testdata/compile_fail` 或独立 fixtures；由脚本断言非零退出码和关键诊断文本 |
| panic/fatal/deadlock 教学 Case | 必须以预期方式失败，不能拖死主测试进程 | 使用子进程、超时和期望退出码；不直接混入普通 `go test ./...` |
| 版本差异 Case | 在指定 Go 版本下呈现对应结果 | 标注版本，必要时用容器或 CI matrix 验证 |

默认验收拆成两条：

```text
make verify                 # 工程基线必须通过
make teaching-failures      # 验证“应该失败”的 Case 确实按预期失败
```

`teaching-failures` 成功不代表示例编译通过，而代表命令失败、错误类型正确、诊断信息能支撑教学目标。不能为了追求 `go test ./...` 全绿而删掉这些例子，也不能让故意错误的代码污染默认构建路径。

### 优先保留的受控失败 Case

| 模块 | 错误写法 | 教学目标 |
| --- | --- | --- |
| Module 1 | 包级使用 `:=` | 短变量声明只能出现在函数体内 |
| Module 1 | `x := 1` 后同作用域再次写 `x := 2` | `:=` 左侧至少要有一个新变量 |
| Module 1 | 未使用局部变量或 import | Go 把无效代码当作编译错误处理 |
| Module 1 | 常量赋值溢出目标整数类型 | 区分编译期常量溢出和运行时整数环绕 |
| Module 1 | `type UserID int` 与 int 直接混用 | 定义类型不是别名，必须显式转换 |
| Module 1 | Slice、Map、Function 使用 `==` 比较 | 理解不可比较类型和仅能与 nil 比较的边界 |
| Module 1 | `scores["Alice"].Value = 1`，Map value 为 Struct | Map index 结果不可寻址，需要读改写回 |
| Module 1 | 最后一个 case 使用 `fallthrough` | `fallthrough` 只能进入紧邻的后续 case |
| Module 2 | `T` 赋给只由 `*T` 实现的接口 | 理解方法集和接收者选择 |
| Module 2 | 向 `<-chan T` 发送或从 `chan<- T` 接收 | 单向 Channel 是编译期约束 |
| Module 2 | `close` 一个 receive-only Channel | 关闭权属于发送方 |
| Module 2 | 对不可设置的 reflect.Value 调用 Set | 区分编译期安全与反射运行时 panic |
| Module 3 | 忽略 `net.Conn` 消息边界，假定一次 Read 等于一条消息 | 这是运行时协议错误，使用测试稳定复现拆包/粘包 |

每个失败 Case 只保留最小代码、预期诊断关键词、讲解问题和修正版本，不和正常 Demo 混在同一个 package。

---

## Module 1：Go 基础

### 1. 变量、常量与类型

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| `var`、类型推断、`:=` | 有多种声明形式和作用域内短声明 | 有基础示例 | 保留新项目顺序；补充 `:=` 左侧至少一个新变量、不能用于包级作用域 |
| 多变量与 `_` | `var a, _, c = 1, 2, 3`，直观展示丢弃值 | 缺失 | 合入 `02_vars_types`，并与多返回值连接 |
| 零值 | 老项目零散存在 | 新项目已系统展示标量、Struct、nil Slice、nil Map | 采用新项目版本 |
| 常量表达式 | 展示 `len(constString)` 可作为常量，普通变量不行 | 只有简单常量 | 补入 `02_vars_types`，讲清“编译期可求值” |
| `iota` 简单枚举 | 有 | 有 | 保留新项目状态枚举 |
| ConstSpec 隐式赋值 | `k3=100; k4; k5=iota`，能观察完整表达式复用 | 原先遗漏，现已补一段 | 保留输出；注释改为“省略 ConstSpec 时复用上一行完整表达式，iota 仍按行递增”，不能表述为“iota 被打断” |
| `iota` 新块重置 | 老项目有第二个 const 块 | 缺失 | 在当前 Case 后加第二个短块，输出 `0, 1` |
| 自定义类型 | `type size int` | 缺失 | 补 `type Size int`，并与 `type Size = int` 对比“定义类型”和“别名” |
| 显式类型转换 | 数字转换、截断、字符串数字互转 | 只有 `int -> float64` | 补 `float64 -> int` 截断；字符串解析放错误处理部分并保留 error |
| 位运算 | `&`、`|`、`^`、`&^` 有二进制对照 | 缺失 | 放 Module 1 选讲，不进入核心验收 |
| 整数溢出 | 老项目未系统讲 | 新项目有 int32 环绕输出 | 保留，但明确“运行时整数运算溢出”和“常量溢出编译失败”的区别 |

### 2. 流程控制

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| if 初始化语句 | 有，并展示外层同名变量不受影响 | 有初始化语句，但作用域对比弱 | 增加同名外层变量输出，讲清 shadowing 和作用域 |
| `for` 三种形式 | 无限循环、条件循环、三段式齐全 | 条件式、三段式为主 | 在 `03_control_funcs` 补极短无限循环 + `break` |
| `break` / `continue` | 普通和标签式都有 | 只有简单 `continue` | 增加同一循环内的可观察对比；标签式放选讲 |
| 多值 `case` | 老项目原例不突出 | 新项目现已补 `case "admin", "owner"` | 保留在主流程 |
| 默认不贯穿 | 有 | 有 | 保留新项目 Java 对比和实际输出 |
| `fallthrough` | 有真实执行 | 现已补真实执行 | 保留；强调无条件进入下一 case，不重新判断条件，业务代码不建议使用 |
| 无表达式 switch | 有 `switch init; {}` | 有 `switch {}` | 补“带初始化语句的无表达式 switch”，与 if 初始化形成对应 |
| 标签式 break/continue | 双层循环 Case 完整 | 只有文档提及 | 合并为一个 Bonus Case，不占核心时间 |
| `goto` | 有 | 只口头提及 | 保持识别级，不迁入核心 Demo |

### 3. String 与转换

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| 拼接、长度、下标 | 有 | 有 | 采用新项目 UTF-8 版本 |
| byte/rune | 有 `[]rune` 遍历 | 新项目讲得更清楚 | 采用新项目版本 |
| `strings` 常用 API | 老项目覆盖非常广 | 新项目只选常用项 | 不照抄 API 清单；围绕文本清洗任务补 `Split`、`Join`、`Replace`、`Index` |
| 空字符串边界 | 有 `Contains("", "")`、`Count(s, "")` 等 | 缺失 | 放 String dark corner |
| `[]byte` / string | 老项目有双向转换 | 新项目未突出 | 合入 String Demo，说明 byte 语义和分配成本 |
| `strconv` | Atoi、ParseInt、ParseUint、进制转换齐全 | 缺失 | 保留最有工程价值的 Atoi/ParseInt，并且必须处理 error；进制转换放选讲 |

### 4. Array、Slice、Map

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| Array 值语义 | 有 range 副本与 index 修改 | 已覆盖且更清楚 | 采用新项目 |
| `[...]T` 推导 | 有 | 有 | 保留 |
| Slice len/cap | 有 | 有 | 保留新项目 |
| Slice 共享底层数组 | 有相邻切片互相影响 | 已覆盖 | 保留新项目 |
| append 扩容分离 | 有，但解释不够严谨 | 已覆盖容量前后两种结果 | 采用新项目 |
| `copy` | 老项目展示 dst/src 长度不同时的结果 | 缺失 | 合入现有 Slice Demo，输出复制数量和两个 Slice |
| nil Slice / nil Map | 老项目有但注释存在误导 | 新项目讲得更准确 | 采用新项目 |
| 嵌套 Map 初始化 | 老项目有 `map[int]map[int]string` | 缺失 | 补入 Map Demo，强调每一层都要初始化 |
| Map Struct value 修改 | 老项目没有形成最佳对照 | 新项目已有读改写回和指针 value | 采用新项目 |
| Map 无序遍历 | 新项目更明确 | 已覆盖 | 采用新项目 |
| 自定义排序 | 老项目实现 `sort.Interface` | 新项目只涉及基础 sort | Module 2 补“旧式 sort.Interface 与现代 sort.Slice/slices.SortFunc”对比 |
| unsafe 指针算术 | 老项目有 | 新项目没有 | 不进主线；最多放危险示例，明确普通 Go 指针不能算术 |

### 5. Struct、Pointer、Method

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| Struct 字面量 | 匿名 Struct、嵌套 Struct、匿名字段都很全 | 具名 Struct 为主 | 核心保留新项目；匿名 Struct 放 JSON/临时 DTO 场景讲 |
| 值传递与指针传递 | 有直接修改对照 | 新项目已有 | 采用新项目 |
| 指针自动解引用 | 老项目展示 `p.field` 与 `(*p).field` | 新项目未突出 | 加两行对照即可 |
| 值/指针接收者 | 老项目细节更多 | 新项目已有主语义 | 增加“值接收者内改字段不影响调用方”的输出 |
| Embedding | 老项目案例更多 | 新项目表达更准确 | 采用新项目“组合而非继承”的表述 |
| 方法表达式 | 老项目有 `(*T).Method(&x, ...)` | 缺失 | 放 Module 2 方法集选讲 |

### 6. Functions、Defer、Error、Testing

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| 多返回值、命名返回值 | 老项目更全 | 新项目覆盖多返回值 | 补一个命名返回值 Case，但明确裸 return 的可读性边界 |
| Variadic 与 Slice 展开 | 老项目有 | 新项目有 | 采用新项目 |
| 函数类型、高阶函数 | 老项目探索很深 | 新项目工程表达更清楚 | 采用新项目 `Combiner`/filter；不迁移“函数指针”叙述 |
| 闭包状态 | 两边都有 | 新项目更清楚 | 采用新项目 |
| defer LIFO | 两边都有 | 新项目更清楚 | 采用新项目 |
| defer 参数求值 | 老项目有循环闭包差异 | 新项目已有参数立即求值与闭包晚求值 | 采用新项目，并关联 range dark corner |
| error | 老项目有自定义 Error 和 panic/recover | 新项目有 Lab 契约 | Module 1 保留 `(value, error)`；自定义 error、wrap、Is/As 放 Module 2 |
| 测试 | 老项目基本没有体系 | 新项目有 Starter/Solution、Quiz、综合 Lab、Homework | 完整采用新项目 |

### Module 1 落地顺序

1. 修正当前 `iota` 注释并补“新 const 块重置”。
2. 在 `02_vars_types` 补 `_`、常量表达式、自定义类型与别名、转换截断。
3. 在 `03_control_funcs` 补无限循环、break/continue 对照、switch 初始化。
4. 在 String Demo 补 `[]byte`、Atoi/ParseInt 的 error 路径。
5. 在 Collections Demo 补 `copy` 和嵌套 Map 初始化。
6. 不新增碎片化目录；全部合入现有主题 Demo 或 dark corners。

---

## Module 2：Go 进阶

### 1. Interface

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| 隐式实现 | 有 person/student 行为 | 有 Animal/Dog/Cat | 采用新项目简单入口 |
| 接口组合 | 有多个接口 | README 声明有，代码较基础 | 补小接口组合，强调“接受接口、返回具体类型” |
| 类型断言 | 有 comma-ok | 基础示例有 | 增加失败路径和 comma-ok |
| type switch | 有完整 Case | 缺失或不突出 | 必补，直接迁移语义但重写命名 |
| typed nil interface | 老项目未系统讲 | 新项目已有专门 Demo | 采用新项目，属于高价值工程 Case |
| 方法集 | 老项目通过反射间接展示 | 新项目未系统串联 | 补 `T` / `*T` 实现接口的矩阵，与接收者章节连接 |

### 2. Error、Defer、Panic

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| 自定义 error | 有 | 有 | 采用新项目 |
| errors wrap / Is / As | 老项目缺少现代写法 | 新项目已有 | 采用新项目，并增加多层 wrap 断言测试 |
| panic/recover | 老项目过程解释更细 | 新项目已有边界 Demo | 保留新项目，补“recover 只在 defer 中有效”和“不能跨 goroutine recover”测试 |
| `log.Fatal` 与 defer | 老项目未突出 | 新项目已有 | 采用新项目，工程价值高 |

### 3. Goroutine、Channel、Context

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| goroutine 启动 | 有 | 有 | 采用新项目 WaitGroup，淘汰 Sleep 等待完成的写法 |
| range 变量捕获 | 老项目受旧版本语义影响 | Module 1 dark corner 有版本说明 | 保留版本对照，不混用旧结论 |
| 无缓冲/缓冲 Channel | 老项目 Case 更多 | 新项目有基础 | 补阻塞时序图和容量耗尽行为 |
| 单向 Channel | 老项目有 `chan<-` / `<-chan` | 新项目 producer 已使用 | 明确函数签名约束，保留 |
| close + range | 老项目有 | 新项目有 | 保留新项目 |
| select | 老项目有随机就绪分支、default 非阻塞 | 新项目有 timeout | 补“多个 case 同时就绪随机选择”和 default 忙等风险 |
| nil Channel | 老项目缺失 | 新项目未突出 | 必补：nil Channel 收发永久阻塞，select 中可用于动态禁用 case |
| closed Channel | 老项目只展示 range | 新项目未系统讲 | 必补：接收零值/ok=false，重复 close 和向已关闭 Channel 发送会 panic |
| Context | 老项目没有 | 新项目有 cancel/timeout | 采用新项目；补取消传播、defer cancel、不要滥用 Value |

### 4. 并发安全

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| WaitGroup | 有 | 有 | 采用新项目，并讲 Add 必须先于 goroutine 启动 |
| Mutex | 有 | 有 | 采用新项目；验证 race detector |
| RWMutex | 老项目有读写耗时对照 | 新项目已有更清楚的专门 Demo | 采用新项目 |
| Atomic | 有无限循环计数 | 新项目有有限任务 | 采用新项目，避免不可控 goroutine |
| Once | 老项目有标准和自实现 | 新项目缺失 | 补 `sync.Once` 工程 Case；自实现只做源码阅读选讲 |
| sync.Map | 老项目缺失 | 新项目有 | 采用新项目，并强调适用场景而非默认替代 map+lock |
| 并发 Map fatal | 老项目缺失 | 新项目已有注释式危险 Case | 保留；用独立进程或讲师 Demo，避免破坏测试进程 |

### 5. Testing、OS、File、Reflection、Runtime、Stdlib

| 内容 | 老项目 | 新项目 | 改进建议 |
| --- | --- | --- | --- |
| 单测/表驱动/子测试/并行测试 | 老项目缺失 | 新项目齐全 | 采用新项目 |
| Benchmark/Profile | 老项目缺失 | 新项目已有 | 采用新项目；增加基准结果解释而非只会运行 |
| OS signal/exec/env | 老项目内容更散 | 新项目整合更好 | 采用新项目；补 CommandContext 超时和退出码 |
| File I/O | 老项目 API 覆盖更细 | 新项目已重写为临时目录和现代 API | 采用新项目；补 Scanner token limit、短写入、fsync/atomic rename 作为工程选讲 |
| Reflection | 老项目字段、方法、可写性、Kind 覆盖更全 | 新项目只覆盖主路径 | 以老项目为清单补齐 `Elem`、exported/unexported、method set、CanSet/CanInterface、panic 边界 |
| Runtime | 老项目只有 NumCPU/GOMAXPROCS 等基础 | 新项目还有 Caller | 采用新项目；GOMAXPROCS 不再教成“必须手动设为 CPU 数” |
| sort/time/json/regexp/hash/base64 | 老项目 API 更广 | 新项目按工程常用项整合 | 采用新项目结构；从老项目补边界 Case，不做 API 大全 |
| embed/generate | 老项目没有 | 新项目有 | 完整采用新项目 |
| 综合项目 | 老项目没有 | 新项目有并发日志分析器 | 完整采用新项目，并让其覆盖 Channel 关闭、Context 取消、错误聚合和 benchmark |

### Module 2 落地顺序

1. Interface 补 type switch 和方法集矩阵。
2. Channel 补 nil/closed Channel、select 随机与 default 风险。
3. Concurrency 补 `sync.Once`，并把 `go test -race` 纳入验收。
4. Reflection 按老项目覆盖面补齐可见性、可写性和方法集。
5. File/OS 只补现代工程边界，不回迁硬编码路径和 `ioutil`。

---

## Module 3：net 包与 Gin

### 1. net 包逐项对比

| 内容 | 老项目 | 新项目 `01_net_basics` | 改进建议 |
| --- | --- | --- | --- |
| TCP Server | `ResolveTCPAddr`、`ListenTCP`、Accept、每连接 goroutine | 只有测试内的 TCP listener | 必须增加可运行 TCP echo server |
| TCP Client | `DialTCP` 和读写 | 缺失 | 必须增加 client，支持超时、退出和半关闭 |
| TCP 消息边界 | 老项目按一次 Read 处理，未讲粘包/拆包 | 有简单行协议 helper | 采用新项目行协议方向；补 partial read、最大帧长、EOF、deadline |
| UDP 单播 | client/server 齐全 | client/server 齐全 | 采用新项目并增加 Context/退出机制和 deadline |
| UDP 广播 | 有 | 缺失 | 放 net 选讲，说明局域网和系统权限限制 |
| UDP 组播 | 有 | 缺失 | 放 net 选讲，说明组地址、网卡和部署环境限制 |
| `net.Conn` 通用接口 | 老项目分 TCP/UDP 具体 API | 新项目 helper 使用 `net.Conn` | 加一节统一讲 Reader/Writer/Closer、LocalAddr/RemoteAddr、Deadline |
| HTTP HandlerFunc | 有 | 有 | 采用新项目 |
| 自定义 Handler | 老项目有多种 `ServeHTTP` | 新项目主要用 HandlerFunc | 增加一个实现 `http.Handler` 的 Struct，连接 Interface 知识 |
| ServeMux | 老项目有默认 mux、自建 mux、静态文件 | 新项目有自建 mux | 补默认 mux 风险、路由匹配、静态文件服务 |
| 自定义 `http.Server` | 老项目有 Addr/Handler | 新项目直接 ListenAndServe | 必须改为显式 Server，并配置 ReadHeaderTimeout/ReadTimeout/WriteTimeout/IdleTimeout |
| HTTP Client GET/POST/Form | 老项目齐全 | 缺失 | 增加 `http.Client` 专题，不使用无超时默认 Client 作为生产示例 |
| 自定义 Request | 老项目有 Header/Do | 缺失 | 补 RequestWithContext、Header、状态码和 body 关闭/读取 |
| TLS Server/Client | 老项目有 HTTPS | 缺失 | 增加本地证书实验；禁止迁移 `InsecureSkipVerify: true` 为推荐写法 |
| HTTP/2 | 老项目有 `x/net/http2` server/client | 缺失 | 迁移概念，不照搬旧 API；优先讲标准库自动协商和 httptest TLS |
| 文件上传 | 老项目有 multipart client/server | 缺失 | 增加大小限制、临时文件清理、文件名安全和内容类型校验 |
| WebSocket | 老项目有 x/net/websocket echo | 缺失 | 放 Gin 之后选讲；使用当前维护的库和 Context/心跳/断线清理 |
| HTTP 测试 | 老项目没有 | 新项目已有 httptest | 采用新项目，增加状态码/Header/body 精确断言 |
| 请求取消 | 老项目没有 | 新项目 Slow Handler 有 | 采用新项目，测试客户端取消后服务端停止工作 |
| 优雅关闭 | 老项目没有 | 缺失 | 必补 `Server.Shutdown(ctx)`、signal、连接排空 |
| 中间件状态码 | 老项目没有 | 新 Logging 固定记录 200 | 修复：包装 ResponseWriter 捕获真实状态码和响应字节数 |

### 2. `01_net_basics` 重构建议

保持一个专题目录，但内部按可运行程序组织：

```text
01_net_basics/
├── tcp/
│   ├── server/
│   ├── client/
│   └── framing/
├── udp/
│   ├── unicast/
│   └── bonus_broadcast_multicast/
├── http/
│   ├── server/
│   ├── client/
│   ├── tls/
│   └── upload/
└── tests/
```

教学顺序：TCP 字节流与消息边界 → UDP 数据报 → `net.Conn` 抽象 → HTTP Server → HTTP Client → TLS → httptest → Graceful Shutdown。HTTP/2、广播、组播、上传、WebSocket 放选讲，不挤压 Gin 主线。

### 3. Gin 部分

| 内容 | 对比结果 | 改进建议 |
| --- | --- | --- |
| Gin 入门与路由 | 新项目明显更好 | 保留 |
| Binding/Viper | 老项目无体系 | 增加配置优先级、必填校验、自定义 validator、敏感配置禁止打印 |
| Middleware/JWT | 新项目更好 | 增加 token 过期、签名算法校验、claims 类型、401/403 区分 |
| Zap 日志 | 新项目更好 | 增加 request-id、真实状态码、panic stack、敏感字段脱敏 |
| API Design | 新项目更好 | 增加统一错误模型、分页、幂等键、版本兼容和 OpenAPI 契约测试 |
| httptest | 新项目更好 | 增加中间件顺序、异常路径、取消和并发测试 |
| Gin Context | 新项目更好 | 明确不可跨 goroutine 直接使用原 Context，必要时 Copy；下游统一传 request context |
| 用户中心项目 | 新项目分层更好 | 增加依赖注入、repository 接口、错误映射、配置启动校验、优雅关闭 |

### Module 3 落地顺序

1. 补齐 TCP server/client 和 framing。
2. 增加现代 HTTP Client、自定义 Server timeout、Graceful Shutdown。
3. 修复 Logging 固定状态码问题。
4. 增加 TLS、上传；广播/组播、HTTP/2、WebSocket 放选讲。
5. Gin 继续沿新项目结构加固，不回退到老项目的单文件 Demo 组织。

---

## Module 4：GORM

老项目无等价内容，新项目结构全部保留。

| 新项目章节 | 当前优势 | 改进建议 |
| --- | --- | --- |
| Setup | 有连接和连接池入口 | DSN 只从配置读取；增加 Ping、连接失败分类、关闭底层 sql.DB |
| Models/Relations | 覆盖一对一、一对多、多对多 | 增加唯一约束、外键策略、软删除、时间字段、DTO 与持久化模型分离 |
| CRUD | 路径清楚 | 所有操作检查 Error/RowsAffected；区分 not found、冲突和内部错误；使用 WithContext |
| Query/Preload | 已有 N+1 和 Hook Demo | 增加投影、分页稳定排序、批量查询、Preload 条件、Hook 副作用边界 |
| Migrations | 有 AutoMigrate 和手动迁移概念 | 引入版本化迁移工具；说明 AutoMigrate 不等于生产迁移方案 |
| Transactions | 有事务和保存点 | 增加 Context 超时、嵌套调用传递 tx、幂等、死锁重试策略 |
| Testing | 有 MySQL 集成测试和 sqlmock | 测试容器/独立 schema；明确 sqlmock 只验证交互，不替代真实数据库语义测试 |
| Raw SQL | 已有专题 | 增加参数绑定、扫描 NULL、动态 IN、Explain、SQL 注入边界 |
| Blog API | 有 Handler/Service/Repository 分层 | 增加接口隔离、统一错误映射、事务归属、集成测试和启动/关闭生命周期 |

优先级：错误处理与 Context > 生产迁移 > 事务边界 > 测试隔离 > 查询性能。

---

## Module 5：gRPC

老项目无等价内容，新项目结构全部保留。

| 新项目章节 | 当前优势 | 改进建议 |
| --- | --- | --- |
| Protobuf | 从 IDL 开始 | 增加 field number 兼容规则、reserved、optional、oneof、timestamp/duration |
| Codegen | 有生成链路 | 固定工具版本；生成结果可复现；增加 buf lint/breaking 或等价检查 |
| Unary RPC | server/client 齐全 | 默认 deadline、连接生命周期、keepalive、消息大小限制、graceful stop |
| Streaming | 三种流齐全 | 增加 EOF、半关闭、背压、客户端取消、并发 Send/Recv 约束 |
| Interceptor | 日志与 recovery | 增加 request-id、deadline、metrics、auth 链顺序和 status 映射 |
| Metadata/Auth | 有 Bearer Token | 增加 TLS/mTLS；token 不硬编码；认证与授权分离 |
| Error Handling | 有 codes 和 details | 建立领域错误到 status code 的固定映射，并测试 details |
| Gateway | 有 HTTP bridge | 增加 HTTP 错误映射、Header 转发、CORS、OpenAPI、流式限制说明 |
| Distributed Compute | 有 engine/auth/test | 增加负载上限、取消传播、worker 背压、部分失败策略、benchmark/race test |

优先级：IDL 兼容性 > deadline/cancel > TLS/auth > 流背压 > 可观测性和稳定测试。

---

## Module 6：go-zero

老项目无等价内容，新项目“双教学线”结构全部保留：小型概念 Demo 用于第一次理解，`project_ecommerce_standard` 用于工程实践。

| 新项目章节 | 当前优势 | 改进建议 |
| --- | --- | --- |
| Intro/goctl | 有框架全景 | 固定 goctl 版本并验证生成结果；明确哪些代码可改、哪些应重新生成 |
| API Service | 有 `.api -> handler -> logic` | 增加统一错误模型、参数校验、配置启动失败、优雅关闭 |
| RPC Service | 有 API 调 RPC | 强制 deadline、错误码透传/转换、重试边界、幂等性 |
| Etcd | 有注册发现概念 | 增加 lease、断线恢复、服务摘除、启动依赖失败和本地测试策略 |
| MySQL/Cache | 有 Cache-Aside | 增加缓存穿透/击穿/雪崩、负缓存、双删风险、事务与缓存不一致 |
| Message Queue | 有异步消息 | 增加 at-least-once、重复消费、幂等键、重试/死信、顺序性 |
| Observability | 有 metrics/tracing | 统一 trace-id、日志/指标/链路关联；增加 RED 指标和告警阈值 |
| K8s | 有 Dockerfile/YAML | 增加非 root、健康检查、资源限制、配置/Secret、滚动发布和优雅终止 |
| 标准电商项目 | 有真实 API/RPC 分层 | 去除内存 stub；补订单状态机、库存幂等、补偿策略、契约和集成测试 |

优先级：消除教学 stub 与标准工程的边界模糊 > 超时/错误/幂等 > 缓存与 MQ 一致性 > 可观测性 > 部署可靠性。

---

## 最终执行清单

### 第一批：语言与并发完整性

1. Module 1 补常量隐式赋值、自定义类型、转换、控制流和集合细节。
2. Module 2 补方法集、type switch、Channel 状态语义、Once、反射边界。
3. 所有 Case 保留“预测输出 → 运行 → 解释 → 工程建议”的教学节奏。

### 第二批：net 包重建

1. TCP server/client/framing。
2. HTTP Client/Server timeout、TLS、上传、优雅关闭。
3. UDP 广播/组播、HTTP/2、WebSocket 作为选讲。
4. 用 httptest、临时端口和 Context 取代硬编码地址与 Sleep。

### 第三批：工程模块加固

1. Module 4 加固 Context、迁移、事务和数据库测试。
2. Module 5 加固 IDL 兼容、deadline、TLS、流控和错误契约。
3. Module 6 加固幂等、一致性、可观测性、部署与真实依赖。
4. Module 4–6 不迁移老项目组织方式，只接受能改善工程实践的具体 Case。
