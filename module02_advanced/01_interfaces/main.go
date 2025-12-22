package main

import "fmt"

// 1. Interface Definition
// Java: interface Animal { void speak(); }
type Animal interface {
	Speak() string
}

// 2. Implementation (Implicit)
// Java: class Dog implements Animal { ... }
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

// 3. Polymorphism
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	d := Dog{Name: "Buddy"}
	c := Cat{Name: "Whiskers"}

	MakeSound(d)
	MakeSound(c)

	// 4. Empty Interface & Type Assertion
	var any interface{} = "I am a string"

	// Type Assertion
	// Java: if (obj instanceof String) { String s = (String) obj; }
	str, ok := any.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}

	// Type Switch
	switch v := any.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}
