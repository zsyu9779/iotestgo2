// Metadata 与认证：通过 metadata 在 gRPC 调用中传递认证 token 和自定义信息
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func main() {
	fmt.Println("=== 06 Metadata & Auth ===")
	fmt.Println()
	fmt.Println("Metadata 是 gRPC 的请求头（类似 HTTP Header），key-value 结构")
	fmt.Println()

	// ========== 服务端：提取 Metadata ==========
	fmt.Println("--- 服务端：从 Metadata 提取 Token ---")
	fmt.Println()
	fmt.Println(`func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    // 从 context 中提取 metadata
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
    }

    // 获取 Authorization header
    authHeader := md.Get("authorization")
    if len(authHeader) == 0 {
        return nil, status.Errorf(codes.Unauthenticated, "missing authorization")
    }

    // 验证 Bearer Token
    token := authHeader[0]
    if !strings.HasPrefix(token, "Bearer ") {
        return nil, status.Errorf(codes.Unauthenticated, "invalid authorization format")
    }
    token = strings.TrimPrefix(token, "Bearer ")

    if token != "valid-token" {
        return nil, status.Errorf(codes.PermissionDenied, "invalid token")
    }

    // 将用户信息存入 context 供后续业务方法使用
    ctx = context.WithValue(ctx, "user_id", "12345")
    return handler(ctx, req)
}`)
	fmt.Println()

	// ========== 客户端：发送 Metadata ==========
	fmt.Println("--- 客户端：发送 Token ---")
	fmt.Println()
	fmt.Println("// 方案 1：使用 AppendToOutgoingContext")
	fmt.Println(`md := metadata.Pairs("authorization", "Bearer valid-token")
ctx := metadata.NewOutgoingContext(context.Background(), md)
resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Gopher"})`)
	fmt.Println()
	fmt.Println("// 方案 2：使用 grpc.WithPerRPCCredentials（自动为每个 RPC 添加）")
	fmt.Println(`type TokenAuth struct { Token string }
func (t *TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
    return map[string]string{"authorization": "Bearer " + t.Token}, nil
}
func (t *TokenAuth) RequireTransportSecurity() bool { return false }

conn, _ := grpc.Dial("localhost:50051",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    grpc.WithPerRPCCredentials(&TokenAuth{Token: "valid-token"}),
)`)
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: ClientInterceptor 中操作 Metadata.Key")
	fmt.Println("  Go:   metadata.FromIncomingContext / NewOutgoingContext")
	fmt.Println("  两者都支持 Per-RPC Credential 机制")

	_ = context.Background
	_ = metadata.New
}
