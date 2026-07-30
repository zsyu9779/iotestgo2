package main

import (
	"fmt"
	"reflect"
)

// 本示例介绍 Go Reflection，对应原项目中的 myref 部分。

type Person struct {
	ID   int
	Name string
}

type Student struct {
	Person // Embedded struct
	Age    int
	Score  float64
}

func (s Student) Study() {
	fmt.Printf("%s is studying...\n", s.Name)
}

func (s *Student) SetScore(score float64) {
	s.Score = score
}

func main() {
	fmt.Println("=== 1. Inspecting Types and Values ===")
	s := Student{Person: Person{ID: 1, Name: "Alice"}, Age: 20, Score: 90.5}
	inspectStruct(s)

	fmt.Println("\n=== 2. Modifying Values via Reflection ===")
	x := 100
	fmt.Printf("Before: %d\n", x)
	modifyValue(&x, 200)
	fmt.Printf("After: %d\n", x)

	fmt.Println("\n=== 3. Calling Methods via Reflection ===")
	callMethod(s, "Study")
}

func inspectStruct(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())

	if t.Kind() == reflect.Struct {
		fmt.Printf("Fields (%d):\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			val := v.Field(i)
			fmt.Printf("  Name: %-10s Type: %-10s Value: %v\n", field.Name, field.Type, val)
		}
	}
}

func modifyValue(ptr interface{}, newValue int) {
	v := reflect.ValueOf(ptr)

	// 检查是否为指针。
	if v.Kind() != reflect.Ptr {
		fmt.Println("Expected a pointer")
		return
	}

	// 获取指针指向的值。
	elem := v.Elem()

	// 检查值是否可设置。
	if elem.CanSet() && elem.Kind() == reflect.Int {
		elem.SetInt(int64(newValue))
	}
}

func callMethod(i interface{}, methodName string) {
	v := reflect.ValueOf(i)
	method := v.MethodByName(methodName)

	if method.IsValid() {
		method.Call(nil) // 无参数调用方法。
	} else {
		fmt.Println("Method not found")
	}
}
