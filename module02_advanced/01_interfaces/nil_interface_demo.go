// 黑暗角落：Interface typed-nil — 接口的 nil 判断陷阱
package main

import (
	"fmt"
	"reflect"
)

// 接口包含两部分：动态类型 + 动态值
// 只有两者都 nil 时，接口才 == nil

type Greeter interface {
	Greet()
}

type G struct{}

func (g *G) Greet() {
	fmt.Println("hi")
}

// 错误示例：返回带类型的 nil 指针
func NewGreeterWrong(ok bool) Greeter {
	var g *G = nil
	if !ok {
		return g // BUG: 返回的接口 != nil（因为类型信息是 *G）
	}
	return &G{}
}

// 正确示例：直接 return nil
func NewGreeterRight(ok bool) Greeter {
	if !ok {
		return nil // 正确：接口的 nil
	}
	return &G{}
}

func showTypedNil() {
	// 基础演示
	var gi Greeter
	fmt.Println("零值接口:", gi == nil) // true

	var p *G
	gi = p
	fmt.Println("赋值 nil 指针后:", gi == nil) // false！接口 != nil！

	// 为什么？
	fmt.Println()
	fmt.Println("接口值结构:")
	fmt.Printf("  零值接口: type=%v, value=%v\n", reflect.TypeOf(gi), reflect.ValueOf(gi))
	gi = p
	fmt.Printf("  nil 赋值后: type=%v, value=%v\n", reflect.TypeOf(gi), reflect.ValueOf(gi))
	fmt.Println("  → 虽然 value 是 nil，但 type 是 *G，所以接口 != nil")
}

func showBestPractice() {
	g1 := NewGreeterWrong(false)
	fmt.Println("NewGreeterWrong(false) == nil:", g1 == nil) // false! (BUG)

	g2 := NewGreeterRight(false)
	fmt.Println("NewGreeterRight(false) == nil:", g2 == nil) // true (正确)

	// 规约：返回接口时，没有值就直接 return nil
	//       不要 return 一个带类型的 nil 指针
}

func RunNilInterfaceDemo() {
	fmt.Println("=== Interface Typed-Nil 黑暗角落 ===")
	fmt.Println()

	showTypedNil()
	fmt.Println()
	showBestPractice()
	fmt.Println()
	fmt.Println("规约：返回接口类型时，没有值就直接 return nil")
	fmt.Println("      不要返回一个带类型的 nil 指针")
	fmt.Println()
	fmt.Println("常见 Bug 模式：")
	fmt.Println("  func GetUser() (*User, error) { ... }   // 返回具体类型，OK")
	fmt.Println("  func GetRepo() RepoInterface { ... }    // 返回接口，注意！")
	fmt.Println("     → 内部不要定义 var r *RepoImpl = nil; return r")
	fmt.Println("     → 应该 return nil")
}
