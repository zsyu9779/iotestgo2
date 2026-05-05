package main

import "fmt"

func main() {
	fmt.Println("=== Protobuf 代码生成指南 ===")
	fmt.Println()
	fmt.Println("1. 安装 protoc 编译器:")
	fmt.Println("   macOS:  brew install protobuf")
	fmt.Println("   Linux:  apt install protobuf-compiler 或从 GitHub 下载")
	fmt.Println()
	fmt.Println("2. 安装 Go 插件:")
	fmt.Println("   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest")
	fmt.Println("   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest")
	fmt.Println()
	fmt.Println("3. 确保 $GOPATH/bin 在 PATH 中:")
	fmt.Println("   export PATH=$PATH:$(go env GOPATH)/bin")
	fmt.Println()
	fmt.Println("4. 生成代码:")
	fmt.Println("   protoc --go_out=. --go_opt=paths=source_relative \\")
	fmt.Println("          --go-grpc_out=. --go-grpc_opt=paths=source_relative \\")
	fmt.Println("          hello.proto")
	fmt.Println()
	fmt.Println("=== 生成文件解读 ===")
	fmt.Println()
	fmt.Println("hello.pb.go (由 protoc-gen-go 生成):")
	fmt.Println("  - HelloRequest / HelloResponse 结构体定义")
	fmt.Println("  - 字段的 getter 方法")
	fmt.Println("  - proto.Message 接口实现（序列化/反序列化）")
	fmt.Println()
	fmt.Println("hello_grpc.pb.go (由 protoc-gen-go-grpc 生成):")
	fmt.Println("  - GreeterClient 接口 + 实现（客户端调用）")
	fmt.Println("  - GreeterServer 接口（服务端需实现）")
	fmt.Println("  - RegisterGreeterServer() 注册函数")
	fmt.Println()
	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: protoc --java_out=. hello.proto")
	fmt.Println("        生成一个巨大的 Hello.java，包含 Builder 模式")
	fmt.Println("  Go:   生成轻量的结构体 + 接口，更符合 Go 简洁哲学")

	// 实际项目中，generated code 会在同目录下
	// 这里仅展示结构，不包含生成的 .pb.go 文件（需手动运行 protoc）
}
