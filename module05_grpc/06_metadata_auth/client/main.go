// 06 Metadata & Auth 客户端：演示三种传 token 方式
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "iotestgo/module05_grpc/06_metadata_auth/proto/authpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient("localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	fmt.Println("=== Metadata & Auth 客户端演示 ===")
	fmt.Println()

	// ========== 方法 1：metadata.NewOutgoingContext（单次调用传 token）==========
	fmt.Println("--- 方法 1: metadata.NewOutgoingContext（单次调用）---")
	callWithOutgoingContext(conn)
	fmt.Println()

	// ========== 方法 2：PerRPCCredentials（每次自动附加）==========
	fmt.Println("--- 方法 2: PerRPCCredentials（自动附加到每个 RPC）---")
	callWithPerRPCCredentials()
	fmt.Println()

	// ========== 方法 3：无 token（触发认证失败）==========
	fmt.Println("--- 方法 3: 无 token（期望认证失败）---")
	callWithoutToken(conn)
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: ClientInterceptor → Metadata.Key.of(\"Authorization\", \"Bearer xxx\")")
	fmt.Println("  Go:   metadata.NewOutgoingContext / grpc.WithPerRPCCredentials")
}

func callWithOutgoingContext(conn *grpc.ClientConn) {
	client := pb.NewGreeterClient(conn)

	md := metadata.Pairs("authorization", "Bearer valid-token-12345")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "OutgoingCtx"})
	if err != nil {
		st := status.Convert(err)
		fmt.Printf("  失败: [%s] %s\n", st.Code(), st.Message())
	} else {
		fmt.Printf("  成功: %s\n", resp.GetMessage())
	}
}

// TokenAuth 实现 credentials.PerRPCCredentials 接口
type TokenAuth struct {
	Token string
}

func (t *TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.Token,
	}, nil
}

func (t *TokenAuth) RequireTransportSecurity() bool {
	return false
}

func callWithPerRPCCredentials() {
	conn, err := grpc.NewClient("localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&TokenAuth{Token: "valid-token-12345"}),
	)
	if err != nil {
		log.Printf("连接失败: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "PerRPCCaller"})
	if err != nil {
		st := status.Convert(err)
		fmt.Printf("  失败: [%s] %s\n", st.Code(), st.Message())
	} else {
		fmt.Printf("  成功: %s\n", resp.GetMessage())
	}
}

func callWithoutToken(conn *grpc.ClientConn) {
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.SayHello(ctx, &pb.HelloRequest{Name: "NoToken"})
	if err != nil {
		st := status.Convert(err)
		fmt.Printf("  符合预期: [%s] %s\n", st.Code(), st.Message())
	}
}
