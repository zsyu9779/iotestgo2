# Go 语言复杂对象配置模式对比指南

在 Go 语言开发中，当一个对象（Struct）拥有大量配置项，且大部分配置项都有默认值时，如何优雅地进行初始化是一个常见的问题。

本文将对比三种常见的配置模式：
1. **柯里化 (Currying)** - _(学术派，实战中不推荐)_
2. **建造者模式 (Builder Pattern)** - _(经典 OOP 模式)_
3. **函数式选项模式 (Functional Options Pattern)** - _(Go 语言最佳实践)_

---

## 1. 柯里化 (Currying)

这是函数式编程中的概念，通过多层函数嵌套，每次传递一个参数并返回一个新的函数，直到所有参数传完。

### 代码示例
```go
func ConfigureService(timeout time.Duration) func(retries int) ServiceConfig {
    return func(retries int) ServiceConfig {
        return ServiceConfig{
            Timeout:    timeout,
            MaxRetries: retries,
        }
    }
}

// 调用
cfg := ConfigureService(5 * time.Second)(3)
```

### 优缺点分析
| 特性 | 评价 | 说明 |
| :--- | :--- | :--- |
| **可读性** | ❌ 低 | 函数签名层层嵌套（如 `func() func() func()`），非常难以阅读和理解。 |
| **灵活性** | ❌ 低 | **强制顺序**。调用者必须按照定义的顺序传递参数，无法跳过。 |
| **扩展性** | ❌ 差 | 增加一个配置项需要修改整个调用链的函数签名，破坏性极大。 |
| **适用场景** | ⚠️ 极少 | 仅适用于参数极少且顺序逻辑强相关的特殊算法场景，**不适合**做对象配置。 |

---

## 2. 建造者模式 (Builder Pattern)

源自 Java 等面向对象语言的经典模式。通过引入一个辅助的 `Builder` 对象来分步设置参数，最后调用 `Build()` 方法生成最终对象。

### 代码示例
```go
type ServerBuilder struct {
    server Server
}

func (b *ServerBuilder) WithTimeout(t time.Duration) *ServerBuilder {
    b.server.Timeout = t
    return b
}

func (b *ServerBuilder) Build() Server {
    return b.server
}

// 调用
server := NewBuilder().
    WithTimeout(5 * time.Second).
    WithRetries(3).
    Build()
```

### 优缺点分析
| 特性 | 评价 | 说明 |
| :--- | :--- | :--- |
| **可读性** | ✅ 高 | 链式调用（Method Chaining）读起来像自然语言句子，非常流畅。 |
| **灵活性** | ✅ 高 | 参数可选，无顺序限制。 |
| **代码量** | ⚠️ 中 | 需要额外定义一个 Builder 结构体和一堆 Set/With 方法，略显啰嗦。 |
| **错误处理** | ⚠️ 一般 | 如果配置过程可能出错（返回 error），链式调用处理起来会比较别扭。 |
| **适用场景** | ✅ 常用 | 适合需要复杂校验或分步骤构建对象的场景。 |

---

## 3. 函数式选项模式 (Functional Options Pattern) 🔥

这是 Go 语言社区公认的**最佳实践**（Idiomatic Go）。利用 Go 的函数是一等公民特性，将配置项抽象为“修改配置对象的函数”。

### 代码示例
```go
// 1. 定义选项函数类型
type Option func(*Server)

// 2. 返回选项闭包
func WithTimeout(t time.Duration) Option {
    return func(s *Server) {
        s.Timeout = t
    }
}

// 3. 构造函数接受变长参数
func NewServer(opts ...Option) *Server {
    // 设置默认值
    s := &Server{
        Timeout: 1 * time.Second,
    }
    // 应用所有选项
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// 调用
server := NewServer(
    WithTimeout(5 * time.Second),
    WithRetries(3),
)
```

### 优缺点分析
| 特性 | 评价 | 说明 |
| :--- | :--- | :--- |
| **可读性** | ✅ 高 | 调用代码非常干净，像是在声明配置。 |
| **扩展性** | ✅ 极佳 | 新增配置只需增加一个新的 `WithXxx` 函数，无需修改 `NewServer` 或其他已有代码。 |
| **默认值** | ✅ 完美 | 构造函数内部先初始化默认值，再应用用户选项，逻辑最自然。 |
| **安全性** | ✅ 高 | 可以将具体的配置字段私有化，只暴露 Option 函数，防止外部直接修改。 |
| **适用场景** | 🏆 首选 | 绝大多数 Go 库（如 gRPC, Context, Zap 等）都采用此模式。 |

---

## 总结与建议

1. **坚决避免**在配置场景使用 **柯里化**。它会把简单的赋值变成复杂的函数调用链。
2. 如果你的对象非常复杂，且构建过程涉及很多步骤校验，可以考虑 **建造者模式**。
3. **在 99% 的 Go 项目中，请直接使用 函数式选项模式**。它最符合 Go 的设计哲学：简单、显式、组合优于继承。
