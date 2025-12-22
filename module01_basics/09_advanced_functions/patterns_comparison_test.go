package main

import (
	"testing"
	"time"
)

// ---------------------------------------------------------
// 1. 原始结构 (为了对比，我们重新定义一个类似的结构)
// ---------------------------------------------------------
type Server struct {
	Timeout    time.Duration
	MaxRetries int
	LogLevel   string
}

// ---------------------------------------------------------
// 2. 建造者模式 (Builder Pattern)
// 优点：可读性高，链式调用，无顺序限制
// 缺点：需要额外的 Builder 结构体，代码量稍多
// ---------------------------------------------------------

type ServerBuilder struct {
	server Server
}

func NewServerBuilder() *ServerBuilder {
	// 设置默认值
	return &ServerBuilder{
		server: Server{
			Timeout:    1 * time.Second,
			MaxRetries: 1,
			LogLevel:   "INFO",
		},
	}
}

func (b *ServerBuilder) WithTimeout(t time.Duration) *ServerBuilder {
	b.server.Timeout = t
	return b
}

func (b *ServerBuilder) WithRetries(r int) *ServerBuilder {
	b.server.MaxRetries = r
	return b
}

func (b *ServerBuilder) WithLogLevel(l string) *ServerBuilder {
	b.server.LogLevel = l
	return b
}

func (b *ServerBuilder) Build() Server {
	return b.server
}

// ---------------------------------------------------------
// 3. 选项模式 (Functional Options Pattern) - Go 语言中最推荐的写法
// 优点：扩展性极强，API 干净，支持默认值，无顺序限制
// ---------------------------------------------------------

type Option func(*Server)

func WithTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func WithRetries(r int) Option {
	return func(s *Server) {
		s.MaxRetries = r
	}
}

func WithLogLevel(l string) Option {
	return func(s *Server) {
		s.LogLevel = l
	}
}

func NewServer(opts ...Option) Server {
	// 1. 初始化默认值
	s := Server{
		Timeout:    1 * time.Second,
		MaxRetries: 1,
		LogLevel:   "INFO",
	}

	// 2. 应用所有选项
	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// ---------------------------------------------------------
// 测试对比
// ---------------------------------------------------------

func TestBuilderPattern(t *testing.T) {
	// 建造者模式：清晰，像在写句子
	server := NewServerBuilder().
		WithTimeout(5 * time.Second).
		WithLogLevel("DEBUG").
		WithRetries(3). // 顺序可以随意
		Build()

	if server.MaxRetries != 3 {
		t.Errorf("Builder pattern failed")
	}
}

func TestFunctionalOptionsPattern(t *testing.T) {
	// 选项模式：Go 社区标准写法
	// 看起来非常像构造函数，但参数是可选的
	server := NewServer(
		WithLogLevel("ERROR"),
		WithRetries(5), // 也可以随意顺序
		// WithTimeout 没传，就用默认值
	)

	if server.LogLevel != "ERROR" || server.MaxRetries != 5 {
		t.Errorf("Functional Options pattern failed")
	}
	if server.Timeout != 1*time.Second { // 验证默认值
		t.Errorf("Default value failed")
	}
}
