package main

import "fmt"

// 1. Struct 定义
type User struct {
	ID    int
	Name  string
	Email string
}

// 2. 方法
// 值接收者：创建副本（适合小型 Struct 或只读操作）
func (u User) String() string {
	return fmt.Sprintf("User[ID=%d, Name=%s]", u.ID, u.Name)
}

// 指针接收者：修改原 Struct（最常见）
func (u *User) UpdateName(newName string) {
	u.Name = newName
}

// Embedding（通过组合提升字段和方法）
type Admin struct {
	User  // 匿名字段
	Level int
}

func main() {
	u := User{ID: 1, Name: "John", Email: "john@example.com"}
	fmt.Println(u.String())

	u.UpdateName("John Doe")
	fmt.Println("Updated:", u.Name)

	// 使用 Embedding
	admin := Admin{
		User:  User{ID: 2, Name: "Admin", Email: "admin@corp.com"},
		Level: 1,
	}
	// 可以直接访问 User 的字段
	fmt.Printf("Admin Name: %s, Level: %d\n", admin.Name, admin.Level)
}
