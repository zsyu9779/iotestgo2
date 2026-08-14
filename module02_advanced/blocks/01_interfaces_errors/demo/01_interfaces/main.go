package main

import (
	"fmt"
	"reflect"
)

// 1. 定义接口。
type Animal interface {
	Speak() string
}

type Named interface {
	DisplayName() string
}

// NamedAnimal 演示小接口组合。
type NamedAnimal interface {
	Animal
	Named
}

// 2. 通过方法集隐式实现接口。
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) DisplayName() string { return d.Name }

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) DisplayName() string { return c.Name }

// 3. 使用接口实现多态。
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func Describe(a NamedAnimal) {
	fmt.Printf("%s says %s\n", a.DisplayName(), a.Speak())
}

type Renamer interface {
	Rename(string)
}

type Profile struct {
	Name string
}

// Rename 使用指针接收者，因此只有 *Profile 实现 Renamer。
func (p *Profile) Rename(name string) { p.Name = name }

// Greeter 演示接口中的 typed-nil 陷阱：接口同时包含动态类型和动态值。
type Greeter interface {
	Greet()
}

type G struct{}

func (g *G) Greet() {
	fmt.Println("hi")
}

// 错误示例：返回带类型的 nil 指针。
func NewGreeterWrong(ok bool) Greeter {
	var g *G
	if !ok {
		return g // 返回的接口 != nil，因为接口中保存了 *G 类型信息。
	}
	return &G{}
}

// 正确示例：没有值时直接返回 nil 接口。
func NewGreeterRight(ok bool) Greeter {
	if !ok {
		return nil
	}
	return &G{}
}

func showTypedNil() {
	var gi Greeter
	fmt.Println("零值接口:", gi == nil) // true

	var p *G
	gi = p
	fmt.Println("赋值 nil 指针后:", gi == nil) // false

	fmt.Println()
	fmt.Println("接口值结构:")
	fmt.Printf("  nil 赋值后: type=%v, value=%v\n", reflect.TypeOf(gi), reflect.ValueOf(gi))
	fmt.Println("  → 虽然 value 是 nil，但 type 是 *G，所以接口 != nil")
}

func showBestPractice() {
	g1 := NewGreeterWrong(false)
	fmt.Println("NewGreeterWrong(false) == nil:", g1 == nil) // false

	g2 := NewGreeterRight(false)
	fmt.Println("NewGreeterRight(false) == nil:", g2 == nil) // true
}

func RunNilInterfaceDemo() {
	fmt.Println("=== Interface Typed-Nil 黑暗角落 ===")
	fmt.Println("=============================")

	showTypedNil()
	fmt.Println()
	showBestPractice()
}

func main() {
	d := Dog{Name: "Buddy"}
	c := Cat{Name: "Whiskers"}

	MakeSound(d)
	MakeSound(c)
	Describe(d)

	// 4. 空接口和类型断言。
	var any interface{} = "I am a string"

	// 类型断言。
	str, ok := any.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}

	// 类型 switch。
	switch v := any.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
	//
	// 断言失败时 ok=false，结果是目标类型的零值，不会 panic。
	number, numberOK := any.(int)
	fmt.Printf("Failed assertion: value=%d, ok=%v\n", number, numberOK)

	// 方法集矩阵：Profile 有值接收者方法集；*Profile 还包含指针接收者方法。
	profile := Profile{Name: "before"}
	var renamer Renamer = &profile
	renamer.Rename("after")
	fmt.Println("Pointer method set:", profile.Name)

	RunNilInterfaceDemo()
}
