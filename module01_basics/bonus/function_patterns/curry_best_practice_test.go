package main

import (
	"fmt"
	"strings"
	"testing"
)

// ---------------------------------------------------------
// 场景 1: 中间件 (Middleware) / 装饰器模式
// ---------------------------------------------------------
// 这是 Go Web 框架（如 Gin, Echo）中最经典的应用。
// 本质上是：func(NextHandler) -> CurrentHandler

type Handler func(string) string

// 基础处理函数
func HelloHandler(name string) string {
	return "Hello, " + name
}

// 柯里化：日志中间件
// 它接受一个 logger 前缀，返回一个“装饰器函数”
// 装饰器函数接受一个 Handler，返回一个新的 Handler
func WithLogging(prefix string) func(Handler) Handler {
	return func(next Handler) Handler {
		return func(name string) string {
			// 前置逻辑
			logMsg := fmt.Sprintf("[%s] Before handling %s", prefix, name)
			fmt.Println(logMsg)

			// 调用原有逻辑
			result := next(name)

			// 后置逻辑
			fmt.Println(fmt.Sprintf("[%s] After handling", prefix))
			return result
		}
	}
}

func TestMiddleware(t *testing.T) {
	// 1. 创建中间件 (柯里化第一层：配置中间件)
	authLogMiddleware := WithLogging("AUTH")
	requestLogMiddleware := WithLogging("REQUEST")

	// 2. 组装链条 (柯里化第二层：包装处理函数)
	// 相当于：WithLogging("AUTH")(WithLogging("REQUEST")(HelloHandler))
	handler := authLogMiddleware(requestLogMiddleware(HelloHandler))

	// 3. 执行
	result := handler("Alice")
	if !strings.Contains(result, "Hello, Alice") {
		t.Errorf("Handler result incorrect: %s", result)
	}
}

// ---------------------------------------------------------
// 场景 2: 延迟计算 / 模板生成
// ---------------------------------------------------------
// 固定一部分公共参数，生成专门的工具函数

func URLBuilder(baseURL string) func(path string) string {
	return func(path string) string {
		return fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), strings.TrimLeft(path, "/"))
	}
}

func TestURLBuilder(t *testing.T) {
	// 固定 API 根地址
	apiV1 := URLBuilder("https://api.example.com/v1")
	apiV2 := URLBuilder("https://api.example.com/v2")

	// 此时生成具体的 URL 变得非常简洁
	url1 := apiV1("users/list")
	url2 := apiV2("orders/create")

	if url1 != "https://api.example.com/v1/users/list" {
		t.Errorf("URL build failed: %s", url1)
	}
	if url2 != "https://api.example.com/v2/orders/create" {
		t.Errorf("URL build failed: %s", url2)
	}
}
