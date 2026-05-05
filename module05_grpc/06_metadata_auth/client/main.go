// 客户端：演示 metadata 传递 Token
package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("=== 06 Metadata Auth 客户端 ===")
	fmt.Println()
	fmt.Println("客户端发送 Token 的三种方式：")
	fmt.Println()
	fmt.Println("1. metadata.NewOutgoingContext（单次 RPC 使用）")
	fmt.Println("   ctx := metadata.AppendToOutgoingContext(ctx, \"authorization\", \"Bearer xxx\")")
	fmt.Println()
	fmt.Println("2. grpc.WithPerRPCCredentials（全局自动附加）")
	fmt.Println("   实现 credentials.PerRPCCredentials 接口")
	fmt.Println()
	fmt.Println("3. grpc.WithUnaryInterceptor（客户端拦截器统一注入）")
	fmt.Println("   func clientAuthInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {")
	fmt.Println("       ctx = metadata.AppendToOutgoingContext(ctx, \"authorization\", \"Bearer xxx\")")
	fmt.Println("       return invoker(ctx, method, req, reply, cc, opts...)")
	fmt.Println("   }")

	_ = context.Background
	_ = fmt.Println
}
