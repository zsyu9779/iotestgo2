package main

import "fmt"

type Address struct {
	City string
}

type User struct {
	ID      int
	Name    string
	Enabled bool
	Address Address
}

func main() {
	var zero User
	fmt.Printf("struct zero value: %#v\n", zero)

	user := User{ID: 1, Name: "Alice", Enabled: true}
	copy := user
	copy.Name = "Bob"
	fmt.Printf("struct value copy: original=%#v copy=%#v\n", user, copy)

	mutateValue(user)
	fmt.Println("after value parameter:", user.Name)
	mutatePointer(&user)
	fmt.Println("after pointer parameter:", user.Name)

	var pointer *User
	fmt.Println("nil pointer:", pointer == nil)
	if pointer != nil {
		fmt.Println(pointer.Name)
	}
}

func mutateValue(user User) {
	user.Name = "value copy"
}

func mutatePointer(user *User) {
	user.Name = "shared value"
}
