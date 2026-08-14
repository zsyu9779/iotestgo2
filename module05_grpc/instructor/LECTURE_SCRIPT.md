# Module 05 gRPC 授课逐字讲稿

> 适用对象：已经学过 Go 基础、并发、Gin 与 HTTP API 的学生  
> 建议课时：1 天，净课堂时间 310 分钟  
> 使用方式：普通文字可直接照读；`【操作】` 是教师屏幕操作；`【预期】` 是结果检查；`【提问】` 后停 3–5 秒；`【板书】` 写到白板；`【学员操作】` 是学生动手时间。

## 一、课前准备

### 1. 环境检查

上课前在仓库根目录执行：

```bash
go version
protoc --version
which protoc-gen-go
which protoc-gen-go-grpc
go test ./module05_grpc/...
```

Gateway 生成演示还需要：

```bash
which protoc-gen-grpc-gateway
```

如果缺少工具，按需安装：

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

生成文件已经提交到仓库。现场即使 `protoc` 或插件安装失败，也可以跳过重新生成，继续运行 Demo。

### 2. 终端布局

准备两个终端：

- 终端 A：长期运行服务端。
- 终端 B：运行客户端、测试和代码生成命令。

每完成一个服务端 Demo，都在终端 A 按 `Ctrl+C` 停止进程，再进入下一节。端口如下：

| Demo | 端口 |
|---|---:|
| Unary RPC | 50051 |
| Streaming RPC | 50052 |
| Interceptor | 50053 |
| Metadata/Auth | 50054 |
| Error Handling | 50055 |
| Gateway | gRPC 50051、HTTP 8080 |
| Distributed Compute | 50056 |

### 3. 当天时间表

| 时间 | 内容 | 学员动手 |
|---|---|---:|
| 09:30–10:05 | Protobuf 基础 | 16 分钟 |
| 10:05–10:45 | 代码生成 | 20 分钟 |
| 10:45–10:55 | 休息 | — |
| 10:55–11:30 | Unary RPC | 16 分钟 |
| 11:30–12:00 | Streaming RPC | 13 分钟 |
| 12:00–13:00 | 午休 | — |
| 13:00–13:40 | Interceptor | 20 分钟 |
| 13:40–14:20 | Metadata 与认证 | 20 分钟 |
| 14:20–14:30 | 休息 | — |
| 14:30–15:10 | 错误处理 | 20 分钟 |
| 15:10–16:00 | Gateway 与综合项目 | 30 分钟 |

学员操作共约 155 分钟，占净课堂时间的 50%。时间紧张时优先压缩教师走读，不压缩学生运行和验收环节。

---

## 二、开场（09:30，3 分钟）

【讲师逐字说】

“大家上午好。前面我们已经用 Gin 做过 HTTP API。今天我们换一个场景：假设用户服务、订单服务和库存服务都部署在内部网络里，它们之间每天要调用很多次。我们仍然可以传 JSON，但现在我们更关心三件事：接口契约能不能被机器检查，传输能不能更紧凑，以及一条连接上能不能持续收发数据。”

“今天的主角就是 gRPC。请大家先记住一句话：gRPC 是远程调用框架，Protobuf 是我们今天使用的接口描述和序列化方案。两者经常一起出现，但不是同一个东西。”

“今天每一段都遵循同一个节奏：先看到可运行结果，再读关键代码，再做一个小改动。生成的 `*.pb.go`、`*_grpc.pb.go` 和 `*.gw.pb.go` 只读不手改。要改契约，就改 `.proto`，然后重新生成。”

【板书】

```text
.proto 契约
   ↓ protoc + plugins
Go message + client stub + server interface
   ↓
HTTP/2 上的 RPC 调用
```

【提问】“如果客户端和服务端对同一个字段的类型理解不同，靠口头约定能不能稳定发现问题？”

【收束】“这正是 IDL，也就是接口描述语言，要解决的第一类问题。”

---

## 三、第 1 课：Protobuf 基础（09:33–10:05）

### 1. 一句话定位与先跑结果（5 分钟）

【讲师逐字说】

“Protobuf 的核心不是‘二进制比 JSON 神奇’，而是先写一份机器可读、跨语言的稳定契约，再由工具生成各语言代码。”

【操作】打开：

- `module05_grpc/01_protobuf_basics/example.proto`
- `module05_grpc/01_protobuf_basics/main.go`

然后运行：

```bash
go run ./module05_grpc/01_protobuf_basics
```

【预期】屏幕上依次出现 JSON、Protobuf 十六进制、反序列化结果和本次样例的字节数对比。

【讲师逐字说】

“先看结果，不急着看实现。同一个用户对象，JSON 输出可读文本，Protobuf 输出二进制。当前样例中 Protobuf 更小，但请注意：这个倍数只对当前数据成立，不是所有业务都固定小三倍、十倍。性能也必须用自己的消息结构和负载做基准测试。”

### 2. 读 `.proto`（9 分钟）

【操作】回到 `example.proto`，从上向下圈出 `syntax`、`package`、`go_package`、`message`、字段编号、`service`。

【讲师逐字说】

“第一行 `syntax = "proto3"` 指定语法版本。`package example` 是 Protobuf 命名空间。`go_package` 决定生成的 Go 包导入路径。它们名字相近，但职责不同。”

“`message` 可以类比 Java DTO 或 Go struct。比如 `User` 里，`int32 id = 1` 的 1 不是数组下标，也不是默认值，它叫 field number，字段编号。线上编码靠编号识别字段，所以发布后不能随意改，也不能删除后拿给完全不同的字段复用。”

“`repeated string tags` 在 Go 中会生成切片。嵌套的 `Address` 会生成指针字段。`map<string, string>` 会生成 Go map。枚举的 0 值应表示未知或未指定，因为 proto3 未赋值时会得到零值。”

“最后的 `service UserService` 描述远程能力。`GetUser` 接收一个请求，返回一个响应。这一节只认识契约，真正的 RPC 调用第三节再做。”

【板书】

```text
字段名可重命名（仍需谨慎）
字段号一旦发布不可随意改变或复用
新增字段 → 使用新的字段号
删除字段 → 最好 reserved 原编号和原字段名
```

【提问】“为什么 enum 的第一个值一般是 UNKNOWN，而不是 MALE？”

【参考收束】“因为未设置字段时会得到 0。让 0 表示未知，比让未填写的信息被误判成某个真实业务值更安全。”

### 3. 对照生成 Go 代码（3 分钟）

【操作】打开 `examplepb/example.pb.go`，只搜索并展示：

```text
type User struct
func (x *User) GetName()
```

【讲师逐字说】

“这里只做定位，不逐行读生成文件。我们确认 `message User` 变成了 Go 的 `User` 结构体，字段还有安全 getter。文件头通常写着 Code generated，请不要手改。手改的结果要么下一次生成被覆盖，要么造成契约和实现不一致。”

### 4. 学员练习：新增字段（15 分钟）

【讲师逐字说】

“现在请大家给 `User` 增加邮箱字段。要求使用一个从未使用过的字段编号 8，字段名是 `email`，类型是 `string`。改完运行生成脚本，再在 `main.go` 构造 `protoUser` 时填入邮箱，并在反序列化结果中打印出来。”

【学员操作】

```proto
string email = 8;
```

```bash
(cd module05_grpc/01_protobuf_basics && bash ./gen.sh)
go run ./module05_grpc/01_protobuf_basics
```

【验收】

- `example.pb.go` 中出现 `Email string` 和 `GetEmail()`。
- 反序列化后能打印原邮箱。
- 学员没有直接编辑生成文件。

【三级提示】

1. 先只改 `.proto`，不要动 `examplepb` 目录。
2. 字段写在 `message User` 内，编号用 8。
3. 生成后在 `pb.User{...}` 中加 `Email: "gopher@example.com"`，输出用 `decoded.GetEmail()`。

### 5. 本节收口（2 分钟）

【讲师逐字说】

“这一节请带走三句话：第一，`.proto` 是契约源头；第二，线上真正识别字段的是编号；第三，生成代码只读不手改。下一节我们打开生成流水线，看看一份契约如何变成消息代码和 RPC 桩代码。”

---

## 四、第 2 课：代码生成（10:05–10:45）

### 1. 展示生成命令（7 分钟）

【操作】打开：

- `module05_grpc/02_codegen/hello.proto`
- `module05_grpc/02_codegen/gen.sh`

【讲师逐字说】

“这里有三个参与者。`protoc` 是编译器和调度者；`protoc-gen-go` 负责 message 的 Go 代码；`protoc-gen-go-grpc` 负责 gRPC client 和 server 代码。`protoc` 本身不会凭空知道 Go 的 gRPC 写法，它通过插件完成生成。”

【操作】在终端 B 执行：

```bash
(cd module05_grpc/02_codegen && bash ./gen.sh)
go run ./module05_grpc/02_codegen
```

【预期】生成命令成功，并打印 `HelloRequest` 序列化结果、`GreeterServer` 与 `GreeterClient` 接口的方法。

【讲师逐字说】

“生成后出现两个关键文件。`hello.pb.go` 主要是消息、枚举、描述信息和序列化支持；`hello_grpc.pb.go` 主要是客户端接口、客户端实现、服务端接口、注册函数和调用处理器。”

### 2. 生成代码地图（8 分钟）

【操作】在编辑器全局搜索以下符号，并按顺序跳转：

```text
type HelloRequest struct
type GreeterClient interface
func NewGreeterClient
type GreeterServer interface
type UnimplementedGreeterServer
func RegisterGreeterServer
```

【讲师逐字说】

“客户端拿到 `GreeterClient`，像调用本地方法一样调用 `SayHello`；底层的序列化、HTTP/2 帧和响应反序列化由生成代码与 gRPC 库完成。”

“服务端要满足 `GreeterServer` 接口。我们通常嵌入 `UnimplementedGreeterServer`，这样未来契约新增方法时，旧实现有明确的未实现行为，也能满足生成代码要求。”

“`RegisterGreeterServer` 把我们的实现和 gRPC Server 绑定起来。请把它理解成路由注册，但路由不是手写 URL，而是由 package、service 和 method 共同形成。”

【板书】

```text
hello.proto
 ├─ protoc-gen-go      → hello.pb.go
 └─ protoc-gen-go-grpc → hello_grpc.pb.go

client stub → 编码 → HTTP/2 → 解码 → server implementation
```

### 3. 学员练习：契约驱动的编译错误（20 分钟）

【讲师逐字说】

“下面故意制造一次安全的破坏。把 proto 中的 RPC 方法 `SayHello` 改名为 `Greet`，重新生成，然后先不要改 `main.go`，直接运行。观察编译器如何告诉我们旧调用点已经失效。随后把引用改为新名字，恢复编译。”

【学员操作】

```proto
rpc Greet(HelloRequest) returns (HelloResponse);
```

```bash
(cd module05_grpc/02_codegen && bash ./gen.sh)
go run ./module05_grpc/02_codegen
```

【预期】`main.go` 搜索旧接口方法或相关引用时暴露不一致；修复后可再次运行。

【提醒】当前 `main.go` 主要反射打印接口，没有直接调用 `SayHello`，因此编译不一定失败。若现场没有出现编译错误，让学生再临时加入下面这一行后运行：

```go
_ = pb.GreeterClient.SayHello
```

随后删除临时代码，或把契约恢复为 `SayHello` 并重新生成，保证后续课程使用原始名字。

【讲师逐字说】

“这次错误是好事。契约变更以后，受影响的调用方在编译期暴露，比请求上线后才返回 404 或解析失败更早、更便宜。”

### 4. 本节收口（5 分钟）

【讲师逐字说】

“真实项目通常把 proto 作为独立契约管理，在 CI 中检查生成文件有没有同步，还会检查 breaking change。今天我们的规则简单一点：只改 proto，通过脚本生成，把生成文件和契约一起提交。”

“请大家把 `SayHello` 和生成文件恢复到仓库原始状态。休息十分钟。回来以后，我们真正启动第一个 gRPC 服务。”

---

## 五、第 3 课：Unary RPC（10:55–11:30）

### 1. 启动服务与调用（7 分钟）

【讲师逐字说】

“Unary RPC 是最像普通函数调用的一种：一个请求，对应一个响应。区别是函数的另一端在另一个进程，失败类型也多了网络、超时和服务不可用。”

【操作】终端 A：

```bash
go run ./module05_grpc/03_unary_rpc/server
```

【操作】终端 B：

```bash
go run ./module05_grpc/03_unary_rpc/client Alice
```

【预期】客户端显示 `Hello, Alice! (from gRPC server)`，服务端保持运行。

### 2. 服务端调用链（7 分钟）

【操作】打开 `03_unary_rpc/server/main.go`，依次定位 `net.Listen`、`grpc.NewServer`、`RegisterGreeterServer`、`Serve`、`SayHello`。

【讲师逐字说】

“服务端启动可以压缩成四步：监听 TCP 端口，创建 gRPC Server，注册服务实现，开始 Serve。真正的业务入口是 `SayHello`。”

“`SayHello` 的 `context.Context` 不是装饰品。客户端取消、deadline、链路 metadata 都会沿调用传播。当前代码先检查 `ctx.Err()`，然后读请求、模拟业务、构造响应。”

“`UnimplementedGreeterServer` 是生成代码提供的默认实现。我们嵌入它，再实现自己关心的方法。”

### 3. 客户端调用链（6 分钟）

【操作】打开 `03_unary_rpc/client/main.go`，依次定位 `grpc.NewClient`、`NewGreeterClient`、`WithTimeout`、`SayHello`。

【讲师逐字说】

“`grpc.NewClient` 返回一个可复用的 `ClientConn`。这里的名字容易让人误以为此刻已经完成了所有网络握手；实际上连接通常按需建立。我们应该长期复用这个对象，不要每次业务请求都创建和关闭。”

“`NewGreeterClient` 创建类型安全的 stub。`context.WithTimeout` 给这一次 RPC 三秒预算。`client.SayHello` 看起来像本地方法，但失败时必须处理远程调用语义。”

【提问】“如果服务端处理需要五秒，而客户端 deadline 是三秒，谁应该停止工作？”

【收束】“客户端三秒后不再等待，服务端也应该监听 context 尽快停止无价值的计算。”

### 4. 学员练习：参数校验（15 分钟）

【讲师逐字说】

“请给 `name` 增加长度校验：去掉首尾空格后少于两个字符，返回 `codes.InvalidArgument`；合法时继续响应。客户端不要 `Fatal`，而是把错误转换成 status，并打印 code 和 message。”

【学员操作目标】服务端需要用到：

```go
strings.TrimSpace
status.Errorf
codes.InvalidArgument
```

【验收命令】

```bash
go run ./module05_grpc/03_unary_rpc/client A
go run ./module05_grpc/03_unary_rpc/client Alice
```

【预期】第一个调用得到 `InvalidArgument`，第二个成功。

【讲评逐字说】

“参数错误不是服务崩了，也不是资源不存在。选择准确错误码，就是在维护接口契约。第七节我们会系统处理 status code。”

【操作】结束本节时在终端 A 按 `Ctrl+C`。

---

## 六、第 4 课：Streaming RPC（11:30–12:00）

### 1. 先看三种流（10 分钟）

【讲师逐字说】

“Unary 是一进一出。现在把请求端和响应端分别允许多条消息，就得到三种流：一进多出、多进一出、多进多出。”

【板书】

```text
Server streaming: request 1  → response N
Client streaming: request N  → response 1
Bidirectional:    request N  ↔ response N（双方独立）
```

【操作】打开 `04_streaming_rpc/proto/chat.proto`，圈出三个 `stream` 的位置。

【操作】终端 A：

```bash
go run ./module05_grpc/04_streaming_rpc/server
```

【操作】终端 B：

```bash
go run ./module05_grpc/04_streaming_rpc/client
```

【预期】先收到 5 条服务端推送，再发送 3 条并收到汇总，最后双向发送并收到 3 条回显。

### 2. 读 `Send`、`Recv` 与 EOF（8 分钟）

【讲师逐字说】

“流式代码最重要的不是背 API，而是理解生命周期。`Send` 发送一条消息，`Recv` 接收一条消息。`io.EOF` 在这里通常表示对端正常结束了发送，不等于系统故障。”

【操作】在客户端定位：

- 服务端流的 `for` + `stream.Recv()`。
- 客户端流的 `stream.Send()` + `CloseAndRecv()`。
- 双向流的接收 goroutine + 主 goroutine 发送 + `CloseSend()`。

【讲师逐字说】

“`CloseSend` 只表示‘我不再发了’，不表示‘我也不收了’。双向流里，发送和接收互相独立，所以常见写法是一个 goroutine 接收，另一个执行流发送。它与 channel 有相似的连续数据思维，但不要把远程流当成本地 channel：远程流还会断网、超时、有背压和跨进程序列化。”

### 3. 学员练习：消息序号（12 分钟）

【讲师逐字说】

“请给 `ChatMessage` 增加 `int64 sequence = 4`。重新生成后，双向聊天发送时写入 1、2、3，服务端回显时保留这个序号，客户端打印序号。”

【学员操作】

```bash
(cd module05_grpc/04_streaming_rpc/proto && bash ./gen.sh)
```

【验收】输出能对应显示 1、2、3；原有三个流场景仍能运行。

【讲师收口】

“流式 RPC 的四个生产关键词是：取消、背压、心跳、断线处理。今天的 Demo 是有限消息流，还没有覆盖这些治理问题。现在停止服务端，午休后看如何把日志、恢复和认证从业务方法里抽走。”

【操作】终端 A 按 `Ctrl+C`。

---

## 七、第 5 课：Interceptor（13:00–13:40）

### 1. 问题导入与运行（8 分钟）

【讲师逐字说】

“假设有二十个 RPC 方法，每个方法都要记录耗时、检查认证、捕获 panic。如果把这些代码复制进二十个 handler，会发生什么？业务代码被横切逻辑淹没，规则也很难保持一致。Interceptor 就是 gRPC 调用链上的统一切入点。”

【操作】终端 A：

```bash
go run ./module05_grpc/05_interceptors/server
```

【操作】终端 B：

```bash
go run ./module05_grpc/05_interceptors/client
```

【预期】

- 正常 Unary 调用成功。
- `name=panic` 被转换为 `Internal`，服务进程没有退出。
- Server streaming 返回三条消息。
- 服务端终端出现 method、耗时、panic 和 stream send 日志。

### 2. 洋葱调用链（10 分钟）

【操作】打开 `05_interceptors/server/main.go`，定位 `ChainUnaryInterceptor`。

【板书】

```text
request
  → recovery before
    → logging before
      → handler
    ← logging after
  ← recovery after
response / error
```

【讲师逐字说】

“链的顺序有语义。这里 recovery 在外层，logging 在内层。正常调用依次进入，再反向退出。如果 handler panic，panic 先越过 logging，外层 recovery 捕获并转换为 `Internal`。因此当前实现中，panic 场景未必能走完 logging 的后置耗时日志。这就是拦截器顺序会改变可观测行为的例子。”

“Unary interceptor 的核心是调用 `handler(ctx, req)`。调用之前是前置逻辑，调用之后是后置逻辑。忘记调用 handler，真正的业务方法就永远不会执行。”

### 3. Stream interceptor（4 分钟）

【操作】定位 `streamLoggingInterceptor` 和 `loggingServerStream`。

【讲师逐字说】

“流式拦截器既可以包住整个流的开始和结束，也可以包装 `ServerStream`，重写 `RecvMsg` 和 `SendMsg`，观察每条消息。生产环境要谨慎记录消息体，避免泄露 token、手机号或大对象，并控制日志量。”

### 4. 学员练习：请求标识与耗时（18 分钟）

【讲师逐字说】

“请在 Unary 日志拦截器中统一输出三项：完整 method name、耗时、最终 status code。无论成功还是失败都要有一条结束日志。完成后分别跑正常调用和 panic 调用，比较日志。”

【提示】

```go
status.Code(err)
info.FullMethod
time.Since(start)
```

【验收】

- 正常调用日志 code 为 `OK`。
- panic 调用客户端收到 `Internal`。
- 学生能解释为什么拦截器顺序影响日志是否完整。

【讲师收口】

“Interceptor 适合统一日志、metrics、tracing、auth、限流和恢复；它不应该变成塞所有业务规则的万能抽屉。下一节我们把认证放进拦截器，并看看凭证怎样随调用传递。”

【操作】终端 A 按 `Ctrl+C`。

---

## 八、第 6 课：Metadata 与认证（13:40–14:20）

### 1. 概念与运行（8 分钟）

【讲师逐字说】

“Metadata 可以类比 HTTP header：它不是核心业务 message 的字段，但携带 token、request ID、trace 信息等调用上下文。客户端写入 outgoing context，经过传输后，服务端从 incoming context 读取。”

【板书】

```text
client outgoing metadata
        ↓ network
server incoming metadata
```

【操作】终端 A：

```bash
go run ./module05_grpc/06_metadata_auth/server
```

【操作】终端 B：

```bash
go run ./module05_grpc/06_metadata_auth/client
```

【预期】单次 metadata 和 `PerRPCCredentials` 两种调用成功；无 token 调用得到 `Unauthenticated`。

### 2. 认证链路逐行讲解（12 分钟）

【操作】先打开客户端 `06_metadata_auth/client/main.go`。

【讲师逐字说】

“第一种方式用 `metadata.NewOutgoingContext`，只影响这个 context 派生出的调用，适合 request ID 或临时字段。第二种方式实现 `PerRPCCredentials`，连接发起每个 RPC 时自动附加凭证，更适合统一 token。”

“Demo 中 `RequireTransportSecurity` 返回 false，是为了本地明文演示。生产环境承载 token 时应使用 TLS，并让凭证要求安全传输。”

【操作】再打开服务端 `06_metadata_auth/server/main.go`，沿 `FromIncomingContext`、`md.Get`、Bearer 校验、`context.WithValue` 讲解。

【讲师逐字说】

“认证和授权不要混为一谈。没有凭证或凭证无法证明身份，通常是 `Unauthenticated`；身份已经确认，但没有执行某动作的权限，通常是 `PermissionDenied`。当前 Demo 对错误 token 返回 `PermissionDenied`，是课堂简化；真实系统还要考虑是否隐藏具体失败原因。”

“这里用字符串作为 context key 便于课堂阅读。生产代码建议定义私有 key 类型，避免不同包使用同名字符串发生碰撞。”

### 3. 学员练习：`x-request-id`（20 分钟）

【讲师逐字说】

“请在客户端每次调用时额外加入 `x-request-id`，值可以写成 `req-001`。服务端认证拦截器读出它并打印。如果没有 request ID，也允许业务继续执行，但日志打印 `missing`。”

【学员操作参考入口】

```go
metadata.Pairs(
    "authorization", "Bearer valid-token-12345",
    "x-request-id", "req-001",
)
```

【验收】

- 带 token、带 request ID：调用成功，服务端日志能关联 request ID。
- 带 token、不带 request ID：调用仍成功。
- 不带 token：仍返回 `Unauthenticated`。

【追问】“request ID 应该放进 `HelloRequest` 吗？”

【参考收束】“如果它只是通用链路信息，不属于问候业务本身，放 metadata 更合适。如果它决定业务幂等或成为领域标识，则需要按契约认真建模，不能机械地全部塞进 header。”

【操作】终端 A 按 `Ctrl+C`。

---

## 九、第 7 课：错误处理（14:30–15:10）

### 1. 先看五种结果（8 分钟）

【讲师逐字说】

“普通 Go error 只有文字时，客户端很容易靠字符串做脆弱判断。gRPC 的 `status` 由 code、message 和可选 details 组成。code 是跨语言的机器可判断语义。”

【操作】终端 A：

```bash
go run ./module05_grpc/07_error_handling/server
```

【操作】终端 B：

```bash
go run ./module05_grpc/07_error_handling/client
```

【预期】依次看到：

- `Gopher`：成功。
- 空字符串：`InvalidArgument`。
- `notfound`：`NotFound`。
- `exists`：`AlreadyExists`。
- `timeout`：`DeadlineExceeded`。

### 2. 服务端构造、客户端分支（10 分钟）

【操作】打开服务端 switch，再打开客户端的 `status.Convert(err)`。

【讲师逐字说】

“服务端用 `status.Errorf(codes.NotFound, ...)` 构造语义化错误。客户端用 `status.Convert` 或 `status.Code` 读取。不要依赖英文错误消息做分支，因为文案会变，code 才是稳定契约的一部分。”

“`DeadlineExceeded` 也值得注意。它可能是服务端主动返回，也可能由客户端 context deadline 触发。客户端看到超时，并不能证明服务端一定没有完成写操作。因此涉及支付、创建订单等非幂等操作时，重试必须配合幂等键或查询确认。”

【板书】

```text
InvalidArgument   参数格式/值不合法
NotFound          目标资源不存在
AlreadyExists     创建冲突
Unauthenticated   身份未确认
PermissionDenied  身份已确认但无权
ResourceExhausted 限流/配额
DeadlineExceeded  超时
Unavailable       临时不可用
Internal          内部异常
```

### 3. 受控提醒：panic 场景（2 分钟）

【讲师逐字说】

“这个错误处理 Demo 的服务端没有 recovery interceptor，所以不要在课堂主流程传 `panic`。如果传入，服务进程可能退出。上一节已经演示了正确边界：panic 恢复属于服务端统一防护，而不是普通业务错误分支。”

### 4. 学员练习：客户端按 code 决策（20 分钟）

【讲师逐字说】

“请把客户端从‘统一打印错误’改成按 code 给出动作。`InvalidArgument` 提示检查输入；`NotFound` 提示不重试；`DeadlineExceeded` 提示结果未知、先查询再决定；`Unavailable` 才提示可退避重试；其他错误走兜底。”

【目标结构】

```go
switch st.Code() {
case codes.InvalidArgument:
    // 提示修正输入
case codes.NotFound:
    // 不重试
case codes.DeadlineExceeded:
    // 结果可能未知
case codes.Unavailable:
    // 可考虑指数退避重试
default:
    // 统一兜底
}
```

【验收】五个现有场景都输出不同且合理的客户端动作建议。

【讲师收口】

“错误码不是服务端随手选的日志标签，它会直接决定客户端重试、提示和补偿动作，所以也是接口契约。下一节我们解决最后两个问题：浏览器只会方便地使用 HTTP/JSON 时怎么办，以及如何把今天的能力组合成一个双向流项目。”

【操作】终端 A 按 `Ctrl+C`。

---

## 十、第 8 课：Gateway 与分布式计算项目（15:10–16:00）

### A. gRPC-Gateway（15 分钟）

#### 1. 运行双协议入口（6 分钟）

【讲师逐字说】

“内部服务适合类型安全的 gRPC，但浏览器、第三方合作方或运维调试经常更习惯 HTTP/JSON。gRPC-Gateway 根据 proto 注解生成 HTTP 到 gRPC 的转换层，让同一份业务能力拥有两种入口。”

【操作】打开 `08_grpc_gateway/proto/hello.proto`，圈出：

```proto
option (google.api.http) = {
  post: "/v1/hello"
  body: "*"
};
```

【操作】终端 A：

```bash
go run ./module05_grpc/08_grpc_gateway
```

【操作】终端 B：

```bash
curl -X POST http://localhost:8080/v1/hello \
  -H 'Content-Type: application/json' \
  -d '{"name":"Gopher"}'
```

如果已安装 `grpcurl`，再运行：

```bash
grpcurl -plaintext \
  -d '{"name":"Gopher"}' \
  localhost:50051 hello.Greeter/SayHello
```

【预期】两个入口都得到同一类问候响应。

#### 2. 解释转换链（5 分钟）

【板书】

```text
HTTP/JSON client → :8080 Gateway → gRPC :50051 → Greeter service
gRPC client      ───────────────→ gRPC :50051 → Greeter service
```

【讲师逐字说】

“当前程序虽然在同一进程里启动两个 server，但 Gateway 仍通过 `localhost:50051` 调用 gRPC 服务。`hello.pb.gw.go` 是注解生成的桥接代码。我们维护的是 proto 注解和业务实现，不手写这层重复转换。”

“Gateway 不等于‘免费得到完美 REST API’。真实网关还要处理鉴权、限流、CORS、OpenAPI、错误映射、字段兼容与公网安全。”

#### 3. 快速练习（4 分钟）

【学员操作】把请求中的 `name` 换成自己的名字，先用 curl 调用，再观察服务端实现只执行一次。若时间允许，给 HTTP 注解增加一个 GET 映射作为 `additional_bindings`，重新生成并验证。

【操作】完成后终端 A 按 `Ctrl+C`，释放 50051 和 8080。

### B. 综合项目：分布式计算（35 分钟）

#### 1. 先跑完整结果（6 分钟）

【讲师逐字说】

“最后一个 Demo 不再引入新 API，而是组合今天的能力：Protobuf 契约、双向流、metadata token、stream interceptor、worker pool、status 和可测试的计算引擎。”

【操作】终端 A：

```bash
make run-compute-server
```

【操作】终端 B：

```bash
make run-compute-client
```

【预期】客户端发送 7 个任务，逐条收到 `sum`、`avg`、`max`、`min`、`stddev`、`median` 等结果。由于 4 个 worker 并发执行，结果顺序不保证与发送顺序一致。

【提问】“输出顺序变化是 bug 吗？”

【收束】“如果契约只要求每个任务最终有结果，就不是 bug；客户端通过 `task_id` 关联。如果业务要求严格有序，就必须显式设计排序或序号，不能靠运行时碰巧有序。”

#### 2. 从契约到实现走读（10 分钟）

【操作】按以下顺序打开文件：

1. `project_distributed_compute/proto/compute.proto`
2. `client/main.go`
3. `internal/auth/auth.go`
4. `internal/server/service.go`
5. `internal/engine/engine.go`

【讲师逐字说】

“契约中 `Process(stream ComputeTask) returns (stream ComputeResult)` 表示双向流。每个任务带 `task_id`、数字集合和操作名，每个结果也带 `task_id`，这让乱序结果仍可关联。”

“客户端把 Bearer token 放入 outgoing metadata，创建带 30 秒 deadline 的 context。接收在 goroutine 中持续 `Recv`，主 goroutine 持续 `Send`。发送完调用 `CloseSend`，但接收仍继续。”

“服务端先经过 stream auth interceptor。认证通过才进入 `Process`。`Process` 建立任务 channel、结果 channel 和 4 个 worker。接收端把任务放进队列，worker 调用纯计算引擎，再由单独发送 goroutine写回 stream。”

“为什么只有一个发送 goroutine？除了结构更容易推理，也避免多个 worker 直接竞争同一个 stream 的发送时序。为什么计算引擎单独放包里？因为纯函数比网络流更容易做表格测试。”

#### 3. 用测试证明边界（5 分钟）

【操作】终端 B：

```bash
go test ./module05_grpc/project_distributed_compute/internal/...
```

【预期】auth、engine、server 三个包测试通过。

【讲师逐字说】

“测试不需要真的启动端口。认证测试构造 incoming metadata；引擎测试直接验证纯函数；服务测试用 fake stream 驱动 `Recv` 和收集 `Send`。把网络边界隔开以后，核心逻辑仍然可以快速、稳定地测试。”

#### 4. 学员练习：任务名称贯穿全链路（12 分钟）

【讲师逐字说】

“请给 `ComputeTask` 增加 `string task_name = 4`，给 `ComputeResult` 增加 `string task_name = 6`。重新生成；客户端构造任务时填名称；服务端结果复制名称；客户端输出名称。不要修改现有字段编号。”

【学员操作】

```bash
(cd module05_grpc/project_distributed_compute/proto && bash ./gen.sh)
go test ./module05_grpc/project_distributed_compute/internal/...
```

然后重新运行服务端和客户端验收。

【验收】

- proto 新字段使用新编号，没有改动旧编号。
- 生成文件来自脚本，没有手改。
- 成功和失败结果都保留 task name。
- 测试通过，客户端输出任务名称。

【三级提示】

1. 先改请求和响应 message，再运行 `proto/gen.sh`。
2. 在 `computeTask` 构造 `ComputeResult` 时复制 `task.GetTaskName()`。
3. 在客户端任务字面量赋值，并在 `Printf` 中增加一列。

#### 5. 课程总收口（2 分钟）

【讲师逐字说】

“今天我们完成了一条完整链路：用 proto 定义契约，用插件生成 Go 类型和 stub，用 Unary 完成一次请求响应，用 Streaming 完成持续收发，用 Interceptor 承载横切逻辑，用 Metadata 传递身份，用 Status Code 表达机器可判断的失败，再通过 Gateway 暴露 HTTP/JSON。”

“最后请记住生产边界：内部通信也要 TLS；ClientConn 要复用；每次调用要有 deadline；流要处理取消、背压和断线；错误码和字段编号都是契约；生成代码不能手改；认证、日志和 tracing 应形成统一拦截链。”

“下课前请在自己的纸上写一句回答：如果明天要给这套计算服务增加一个新操作，你会改哪些手写文件，哪些生成文件绝对不直接改？写完就可以下课。”

【参考答案】

“修改 proto（若消息结构需要变化）、生成脚本产物、`engine` 的 operation 与计算逻辑、客户端示例和测试；`*.pb.go`、`*_grpc.pb.go` 不直接修改。”

【操作】终端 A 按 `Ctrl+C`。

---

## 十一、现场故障速查

### 1. `bind: address already in use`

【讲师说】“上一节服务端还在占用端口。先回到终端 A 按 `Ctrl+C`，确认停止后再启动。”

只读定位命令：

```bash
lsof -nP -iTCP:50051 -sTCP:LISTEN
```

不要在不确认 PID 的情况下批量杀进程。

### 2. 客户端报 `Unavailable` 或 connection refused

检查服务端是否已启动、端口是否对应。各 Demo 的客户端不能混用，因为 proto package、service 和端口不同。

### 3. `protoc: command not found`

跳过生成环节，使用仓库中已提交的生成文件继续讲课。课后再完成环境安装。

### 4. `protoc-gen-go: program not found or is not executable`

确认安装插件，并确保 Go bin 在 PATH 中：

```bash
go env GOPATH
```

临时可在当前终端补充：

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

### 5. Gateway 找不到 `google/api/annotations.proto`

先执行：

```bash
go mod download
```

若仍失败，跳过现场重新生成，直接使用仓库内现成的 `hello.pb.gw.go` 运行 Gateway。

### 6. `grpcurl: command not found`

`grpcurl` 是可选调试工具，不影响 Go 客户端 Demo。Gateway 仍可用 curl 验证。

### 7. 流式客户端一直不退出

依次检查：客户端是否调用 `CloseSend`/`CloseAndRecv`；服务端收到 `io.EOF` 后是否结束 worker 并返回；发送 goroutine 是否关闭结果 channel；context 是否设置 deadline。

### 8. 生成后出现大量编译错误

先确认 `.proto` 的 `go_package` 和 `gen.sh` 中的 module 参数匹配，再检查是否只生成了 message 文件而漏掉 `--go-grpc_out`。不要通过手改生成文件消除错误。

---

## 十二、课后最小验收清单

```bash
go test ./module05_grpc/...
go run ./module05_grpc/01_protobuf_basics
go run ./module05_grpc/02_codegen
```

需要双终端验收的服务：

```text
03 server + client
04 server + client
05 server + client
06 server + client
07 server + client
08 Gateway + curl
project server + client
```

教师确认学生能够口头回答：

1. `.proto`、`*.pb.go`、`*_grpc.pb.go` 各自职责是什么？
2. field number 为什么不能随意复用？
3. Unary、三种 Streaming 分别是什么形状？
4. `CloseSend` 和 `io.EOF` 各表示什么？
5. Interceptor 的顺序为什么重要？
6. Metadata 与业务 message 如何划分？
7. `Unauthenticated` 和 `PermissionDenied` 有何区别？
8. 为什么客户端不能靠错误字符串决定重试？
9. Gateway 解决了什么问题，又没有解决什么问题？
10. 综合项目为什么用 `task_id` 关联乱序结果？
