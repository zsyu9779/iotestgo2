package main

import (
	"encoding/json"
	"fmt"
)

// 手动构造的 User 结构体（模拟 .proto 生成的结构体），用于对比 Protobuf 与 JSON
type User struct {
	ID       int32             `json:"id"`
	Name     string            `json:"name"`
	Age      int32             `json:"age"`
	Gender   int32             `json:"gender"` // 0=UNKNOWN, 1=MALE, 2=FEMALE
	Tags     []string          `json:"tags"`
	Address  *Address          `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip_code"`
}

func main() {
	u := User{
		ID:   1,
		Name: "张三",
		Age:  25,
		Tags: []string{"go", "grpc"},
		Address: &Address{
			Street: "中关村大街1号",
			City:   "北京",
			Zip:    "100080",
		},
		Metadata: map[string]string{
			"department": "engineering",
		},
	}

	// JSON 序列化
	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("=== JSON 序列化结果 ===")
	fmt.Println(string(data))
	fmt.Printf("JSON 大小: %d bytes\n", len(data))

	// 对比要点：
	// 1. JSON 有字段名开销，Protobuf 用字段编号（varint），更紧凑
	// 2. Protobuf 有类型信息，JSON 全是字符串/数字
	// 3. Protobuf 向后兼容性更好（字段编号不变即可）
	// 4. JSON 可读性好，适合调试和 Web 浏览器
	// 5. Protobuf 二进制不可读，但体积小 3-10 倍，解析快 20-100 倍

	fmt.Println("\n=== Protobuf 核心语法速览 ===")
	fmt.Println("1. syntax = \"proto3\";              // 声明使用 proto3")
	fmt.Println("2. message User { ... }             // 定义消息类型")
	fmt.Println("3. enum Gender { ... }              // 定义枚举")
	fmt.Println("4. repeated string tags = 5;        // 数组字段")
	fmt.Println("5. map<string, string> meta = 7;    // map 字段")
	fmt.Println("6. service UserService { ... }      // 定义 RPC 服务")
	fmt.Println("7. rpc GetUser(...) returns (...);  // 定义 RPC 方法")
	fmt.Println()
	fmt.Println("字段编号 1-15 用 1 字节编码，16-2047 用 2 字节，高频字段用 1-15")
}
