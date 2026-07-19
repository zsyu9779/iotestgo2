package main

import "fmt"

// 1. Struct Definition
type User struct {
	ID    int
	Name  string
	Email string
}

// 2. Methods
// Value Receiver: creates a copy (good for small structs or read-only)
func (u User) String() string {
	return fmt.Sprintf("User[ID=%d, Name=%s]", u.ID, u.Name)
}

// Pointer Receiver: modifies the original struct (most common)
func (u *User) UpdateName(newName string) {
	u.Name = newName
}

// Embedding (Inheritance-like)
type Admin struct {
	User  // Anonymous field
	Level int
}

func main() {
	u := User{ID: 1, Name: "John", Email: "john@example.com"}
	fmt.Println(u.String())

	u.UpdateName("John Doe")
	fmt.Println("Updated:", u.Name)

	// Embedding usage
	admin := Admin{
		User:  User{ID: 2, Name: "Admin", Email: "admin@corp.com"},
		Level: 1,
	}
	// Can access User fields directly
	fmt.Printf("Admin Name: %s, Level: %d\n", admin.Name, admin.Level)
}
