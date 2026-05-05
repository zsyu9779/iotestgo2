// 07 gRPC 错误处理客户端：用不同 name 触发不同错误码
package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	pb "iotestgo/module05_grpc/07_error_handling/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient("localhost:50055",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	fmt.Println("=== gRPC 错误处理演示 ===")
	fmt.Println()

	// 测试各种场景
	testCases := []struct {
		name        string
		timeout     time.Duration
		description string
	}{
		{"Gopher", 2 * time.Second, "正常请求"},
		{"", 2 * time.Second, "空参数 → InvalidArgument"},
		{"notfound", 2 * time.Second, "资源不存在 → NotFound"},
		{"exists", 2 * time.Second, "资源冲突 → AlreadyExists"},
		{"timeout", 200 * time.Millisecond, "超时处理 → DeadlineExceeded"},
	}

	for _, tc := range testCases {
		fmt.Printf("--- %s (name=%q, timeout=%v) ---\n", tc.description, tc.name, tc.timeout)

		ctx, cancel := context.WithTimeout(context.Background(), tc.timeout)
		defer cancel()

		resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: tc.name})
		if err != nil {
			st := status.Convert(err)
			printGRPCError(st)
		} else {
			fmt.Printf("  ✓ 成功: %s\n", resp.GetMessage())
		}
		fmt.Println()
	}

	fmt.Println("=== 常用 gRPC Status Codes ===")
	codesList := []struct {
		code codes.Code
		desc string
		when string
	}{
		{codes.OK, "成功", "RPC 正常完成"},
		{codes.InvalidArgument, "参数错误", "客户端传入无效参数"},
		{codes.NotFound, "资源不存在", "请求的资源未找到"},
		{codes.AlreadyExists, "资源冲突", "创建已存在的资源"},
		{codes.PermissionDenied, "无权限", "已认证但无权限"},
		{codes.Unauthenticated, "未认证", "缺少或无效的凭证"},
		{codes.ResourceExhausted, "资源耗尽", "限流/配额超限"},
		{codes.DeadlineExceeded, "超时", "操作超过 deadline"},
		{codes.Unimplemented, "未实现", "方法未实现"},
		{codes.Internal, "内部错误", "服务端异常（如 panic）"},
		{codes.Unavailable, "服务不可用", "服务临时不可用"},
	}
	for _, c := range codesList {
		fmt.Printf("  %-25s %-3d  %-12s  %s\n", c.code.String(), c.code, c.desc, c.when)
	}
}

func printGRPCError(st *status.Status) {
	fmt.Printf("  ✗ gRPC 错误:\n")
	fmt.Printf("    Code:    %s (%d)\n", st.Code(), st.Code())
	fmt.Printf("    Message: %s\n", st.Message())

	// 尝试提取详细错误信息
	for _, detail := range st.Details() {
		msg := fmt.Sprintf("%v", detail)
		msg = strings.TrimPrefix(msg, "[")
		msg = strings.TrimSuffix(msg, "]")
		fmt.Printf("    Detail:  %s\n", msg)
	}
}
