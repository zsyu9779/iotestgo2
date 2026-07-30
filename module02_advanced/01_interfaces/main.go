package main

import "fmt"

// 1. 定义接口。
type Animal interface {
	Speak() string
}

// 2. 通过方法集隐式实现接口。
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// 3. 使用接口实现多态。
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	d := Dog{Name: "Buddy"}
	c := Cat{Name: "Whiskers"}

	MakeSound(d)
	MakeSound(c)

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
}
