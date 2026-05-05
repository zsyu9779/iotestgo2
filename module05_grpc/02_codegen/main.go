// 02 Protobuf 代码生成：protoc 编译流程 + 生成文件解读
package main

import (
	"fmt"
	"reflect"

	pb "iotestgo/module05_grpc/02_codegen/hellopb"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("=== Protobuf 代码生成指南 ===")
	fmt.Println()

	// 1. 安装指引
	fmt.Println("--- 1. 安装 protoc 编译器 ---")
	fmt.Println("  macOS:  brew install protobuf")
	fmt.Println("  Linux:  apt install protobuf-compiler")
	fmt.Println()
	fmt.Println("--- 2. 安装 Go 插件 ---")
	fmt.Println("  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest")
	fmt.Println("  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest")
	fmt.Println()
	fmt.Println("--- 3. 生成命令 ---")
	fmt.Println("  protoc --proto_path=. \\")
	fmt.Println("    --go_out=hellopb --go_opt=module=iotestgo/module05_grpc/02_codegen/hellopb \\")
	fmt.Println("    --go-grpc_out=hellopb --go-grpc_opt=module=iotestgo/module05_grpc/02_codegen/hellopb \\")
	fmt.Println("    hello.proto")
	fmt.Println()

	// 2. 演示生成的结构体
	req := &pb.HelloRequest{Name: "Gopher"}
	fmt.Println("--- 生成的结构体 ---")
	fmt.Printf("  HelloRequest{Name: %q}  →  %T\n", req.GetName(), req)
	fmt.Printf("  HelloRequest 实现了 proto.Message 接口\n\n")

	// 3. 演示序列化
	data, _ := proto.Marshal(req)
	fmt.Println("--- 序列化测试 ---")
	fmt.Printf("  HelloRequest{Name:\"Gopher\"} 序列化 → %d bytes\n", len(data))
	fmt.Printf("  二进制: %x\n\n", data)

	// 4. 反序列化
	decoded := &pb.HelloRequest{}
	proto.Unmarshal(data, decoded)
	fmt.Printf("  反序列化 → Name=%q\n\n", decoded.GetName())

	// 5. 检查生成的接口
	fmt.Println("--- GreeterServer 接口（服务端需实现） ---")
	serverType := reflect.TypeOf((*pb.GreeterServer)(nil)).Elem()
	for i := 0; i < serverType.NumMethod(); i++ {
		m := serverType.Method(i)
		fmt.Printf("  %s%v\n", m.Name, m.Type)
	}
	fmt.Println()

	fmt.Println("--- GreeterClient 接口（客户端调用） ---")
	clientType := reflect.TypeOf((*pb.GreeterClient)(nil)).Elem()
	for i := 0; i < clientType.NumMethod(); i++ {
		m := clientType.Method(i)
		fmt.Printf("  %s%v\n", m.Name, m.Type)
	}
	fmt.Println()

	// 6. 文件解读
	fmt.Println("--- 生成文件解读 ---")
	fmt.Println("  hello.pb.go:")
	fmt.Println("    - HelloRequest / HelloResponse 结构体")
	fmt.Println("    - 字段 getter (GetName, GetMessage)")
	fmt.Println("    - proto.Marshal / proto.Unmarshal 实现")
	fmt.Println()
	fmt.Println("  hello_grpc.pb.go:")
	fmt.Println("    - GreeterClient 接口 + 客户端实现")
	fmt.Println("    - GreeterServer 接口（服务端实现它）")
	fmt.Println("    - RegisterGreeterServer 注册函数")
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: protoc --java_out=. hello.proto")
	fmt.Println("        生成 HelloOuterClass.java（Builder 模式，巨大单文件）")
	fmt.Println("  Go:   protoc --go_out=. hello.proto")
	fmt.Println("        生成轻量 struct + 接口，符合 Go 简洁哲学")
	fmt.Println("  Go 不需要 Builder：直接用 &HelloRequest{Name: \"xxx\"}")
}
