// go:generate 代码生成简介
//
// go generate 是 Go 内置的代码生成工具
// 通过源码中的特殊注释触发外部工具生成代码
//
// 运行：go generate ./...
//
// 常用生成工具：
//   - stringer: 为枚举类型生成 String() 方法
//   - mockgen: 生成 mock 实现
//   - goctl: go-zero 代码生成
//   - protoc: protobuf 代码生成
package main

import "fmt"

//go:generate echo "Running go generate..."

// 1. stringer：为枚举生成 String() 方法
//go:generate go run golang.org/x/tools/cmd/stringer -type=Status -output=status_string.go

type Status int

const (
	Pending Status = iota
	Active
	Inactive
	Deleted
)

// 2. mockgen：生成接口的 mock 实现
//go:generate mockgen -source=$GOFILE -destination=mock_repo.go -package=main

type UserRepository interface {
	FindByID(id int) (*User, error)
	Create(name string) (*User, error)
}

type User struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("=== go:generate 代码生成 ===")
	fmt.Println()

	fmt.Println("//go:generate 注释用法：")
	fmt.Println("  //go:generate <command> <args>")
	fmt.Println()
	fmt.Println("常用生成工具：")
	fmt.Println("  1. stringer  - 为枚举生成 String() 方法")
	fmt.Println("     //go:generate go run golang.org/x/tools/cmd/stringer -type=Status")
	fmt.Println()
	fmt.Println("  2. mockgen   - 生成接口 mock 实现")
	fmt.Println("     //go:generate mockgen -source=$GOFILE -destination=mock_repo.go")
	fmt.Println()
	fmt.Println("  3. protoc    - protobuf 代码生成")
	fmt.Println("     //go:generate protoc --go_out=. hello.proto")
	fmt.Println()
	fmt.Println("  4. goctl     - go-zero 代码生成")
	fmt.Println("     //go:generate goctl api go -api user.api")
	fmt.Println()
	fmt.Println("运行方式：")
	fmt.Println("  go generate ./...              # 递归所有包")
	fmt.Println("  go generate ./module02_advanced/14_generate/  # 指定包")
	fmt.Println("  go generate -n ./...           # 预览（不执行）")
	fmt.Println("  go generate -x ./...           # 打印执行命令")
	fmt.Println()
	fmt.Println("变量：")
	fmt.Println("  $GOFILE     - 当前文件名")
	fmt.Println("  $GOPACKAGE  - 当前包名")
	fmt.Println("  $GOARCH     - 目标架构")
	fmt.Println("  $GOOS       - 目标操作系统")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Java: Lombok / MapStruct 注解处理器")
	fmt.Println("  Go:   go:generate + go generate 命令")
	fmt.Println("  Go 的生成更灵活（任意命令），但不与编译器集成")
}
