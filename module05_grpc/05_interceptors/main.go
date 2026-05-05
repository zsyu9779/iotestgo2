// Interceptor 拦截器：gRPC 的"中间件"
//
// 作用：日志、认证、限流、panic 恢复、链路追踪、指标收集
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("=== 05 Interceptor 拦截器 ===")
	fmt.Println()

	// ========== Unary Interceptor ==========
	fmt.Println("--- Unary Interceptor ---")
	showUnaryInterceptor()
	fmt.Println()

	// ========== Stream Interceptor ==========
	fmt.Println("--- Stream Interceptor ---")
	showStreamInterceptor()
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: ServerInterceptor 接口, 用 ServerInterceptors.intercept() 包装")
	fmt.Println("  Go:   grpc.UnaryInterceptor() / grpc.StreamInterceptor() 函数式选项")
	fmt.Println("  Go 的拦截器是函数，更轻量；Java 的拦截器是类，更 OOP")
}

func showUnaryInterceptor() {
	// Unary Interceptor 签名：
	// func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)

	fmt.Println("// 日志 + 计时拦截器：")
	fmt.Println(`func loggingInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    start := time.Now()

    // 前置处理：打印请求
    log.Printf("[REQ] %s: %v", info.FullMethod, req)

    // 调用实际处理方法
    resp, err := handler(ctx, req)

    // 后置处理：打印耗时
    log.Printf("[RESP] %s: %v (elapsed: %v)", info.FullMethod, err, time.Since(start))

    return resp, err
}`)
	fmt.Println()

	fmt.Println("// Panic 恢复拦截器：")
	fmt.Println(`func recoveryInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (resp interface{}, err error) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("[PANIC] %s: %v", info.FullMethod, r)
            err = status.Errorf(codes.Internal, "internal error")
        }
    }()
    return handler(ctx, req)
}`)
	fmt.Println()

	fmt.Println("// 链式组合多个拦截器：")
	fmt.Println("s := grpc.NewServer(")
	fmt.Println("    grpc.ChainUnaryInterceptor(")
	fmt.Println("        recoveryInterceptor,  // 最外层：恢复 panic")
	fmt.Println("        loggingInterceptor,   // 第二层：记录日志")
	fmt.Println("        authInterceptor,      // 最内层：认证校验")
	fmt.Println("    ),")
	fmt.Println(")")

	_ = context.Background
	_ = time.Now
	_ = log.Println
	_ = grpc.NewServer
}

func showStreamInterceptor() {
	fmt.Println("// Stream Interceptor 签名：")
	fmt.Println("// func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error")
	fmt.Println()
	fmt.Println("// 包装 ServerStream 添加日志：")
	fmt.Println(`type wrappedStream struct {
    grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
    err := w.ServerStream.RecvMsg(m)
    log.Printf("[STREAM RECV] %v, err=%v", m, err)
    return err
}

func (w *wrappedStream) SendMsg(m interface{}) error {
    err := w.ServerStream.SendMsg(m)
    log.Printf("[STREAM SEND] %v, err=%v", m, err)
    return err
}`)

	_ = fmt.Println
}
