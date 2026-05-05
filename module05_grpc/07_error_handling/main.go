// gRPC 错误处理：使用 status 包和 codes 返回结构化错误
package main

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("=== 07 gRPC 错误处理 ===")
	fmt.Println()

	fmt.Println("--- 基本错误返回 ---")
	fmt.Println()
	fmt.Println("// 返回标准 gRPC 错误：")
	fmt.Println(`if req.GetName() == "" {
    return nil, status.Errorf(codes.InvalidArgument, "name is required")
}`)
	fmt.Println()
	fmt.Println("// 客户端判断错误码：")
	fmt.Println(`resp, err := client.SayHello(ctx, req)
if err != nil {
    st := status.Convert(err)
    switch st.Code() {
    case codes.InvalidArgument:
        fmt.Println("参数错误:", st.Message())
    case codes.NotFound:
        fmt.Println("资源不存在")
    case codes.DeadlineExceeded:
        fmt.Println("请求超时")
    default:
        fmt.Println("未知错误:", st.Code())
    }
}`)

	fmt.Println()
	fmt.Println("--- 常用 gRPC Status Codes ---")
	codesList := []struct {
		code codes.Code
		desc string
	}{
		{codes.OK, "成功"},
		{codes.InvalidArgument, "客户端参数错误"},
		{codes.NotFound, "资源不存在"},
		{codes.AlreadyExists, "资源已存在（冲突）"},
		{codes.PermissionDenied, "无权限"},
		{codes.Unauthenticated, "未认证"},
		{codes.ResourceExhausted, "资源耗尽（限流）"},
		{codes.DeadlineExceeded, "超时"},
		{codes.Unimplemented, "方法未实现"},
		{codes.Internal, "服务端内部错误"},
		{codes.Unavailable, "服务不可用"},
	}

	for _, c := range codesList {
		fmt.Printf("  %-25s = %d  // %s\n", c.code.String(), c.code, c.desc)
	}

	fmt.Println()
	fmt.Println("--- 自定义 Error Detail ---")
	fmt.Println("// 使用 google.golang.org/genproto/googleapis/rpc/errdetails")
	fmt.Println(`import "google.golang.org/genproto/googleapis/rpc/errdetails"

func badRequest(msg string) error {
    st := status.New(codes.InvalidArgument, msg)
    // 附加 BadRequest 详情，携带具体哪个字段出错
    br := &errdetails.BadRequest{
        FieldViolations: []*errdetails.BadRequest_FieldViolation{
            {Field: "name", Description: "name is required"},
        },
    }
    st, _ = st.WithDetails(br)
    return st.Err()
}`)

	fmt.Println()
	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: io.grpc.Status 类 + StatusException")
	fmt.Println("  Go:   google.golang.org/grpc/status 包")
	fmt.Println("  两者都用相同的 code 枚举值（gRPC 标准）")

	_ = status.New
	_ = codes.OK
}
