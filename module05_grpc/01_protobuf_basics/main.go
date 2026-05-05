// 01 Protobuf 基础：JSON vs Protobuf 序列化对比
package main

import (
	"encoding/json"
	"fmt"

	pb "iotestgo/module05_grpc/01_protobuf_basics/examplepb"

	"google.golang.org/protobuf/proto"
)

// 手动构造的 User 结构体（纯 Go），用于 JSON 序列化对比
type JSONUser struct {
	ID       int32             `json:"id"`
	Name     string            `json:"name"`
	Age      int32             `json:"age"`
	Gender   int32             `json:"gender"`
	Tags     []string          `json:"tags"`
	Address  *JSONAddress      `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type JSONAddress struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip_code"`
}

func main() {
	fmt.Println("=== Protobuf vs JSON 序列化对比 ===")
	fmt.Println()

	// ========== 构造相同的 User 数据 ==========
	// Go JSON 版本
	jsonUser := JSONUser{
		ID:     1,
		Name:   "张三",
		Age:    25,
		Gender: 1,
		Tags:   []string{"go", "grpc", "protobuf"},
		Address: &JSONAddress{
			Street: "中关村大街1号",
			City:   "北京",
			Zip:    "100080",
		},
		Metadata: map[string]string{
			"department": "engineering",
			"level":      "senior",
		},
	}

	// Protobuf 版本（使用生成的代码）
	protoUser := &pb.User{
		Id:     1,
		Name:   "张三",
		Age:    25,
		Gender: pb.Gender_MALE,
		Tags:   []string{"go", "grpc", "protobuf"},
		Address: &pb.Address{
			Street: "中关村大街1号",
			City:   "北京",
			ZipCode: "100080",
		},
		Metadata: map[string]string{
			"department": "engineering",
			"level":      "senior",
		},
	}

	// ========== JSON 序列化 ==========
	jsonData, _ := json.MarshalIndent(jsonUser, "", "  ")
	fmt.Println("--- JSON 序列化 ---")
	fmt.Println(string(jsonData))
	fmt.Printf("JSON 大小: %d bytes\n\n", len(jsonData))

	// ========== Protobuf 序列化 ==========
	protoData, _ := proto.Marshal(protoUser)
	fmt.Println("--- Protobuf 序列化 ---")
	fmt.Printf("Protobuf 大小: %d bytes\n", len(protoData))
	fmt.Printf("Protobuf 二进制 (hex): %x\n\n", protoData)

	// ========== Protobuf 反序列化 ==========
	decoded := &pb.User{}
	proto.Unmarshal(protoData, decoded)
	fmt.Println("--- Protobuf 反序列化验证 ---")
	fmt.Printf("  Name: %s, Age: %d, Tags: %v\n", decoded.GetName(), decoded.GetAge(), decoded.GetTags())

	// ========== 对比总结 ==========
	fmt.Println()
	fmt.Println("=== 核心差异 ===对照")
	sizeRatio := float64(len(jsonData)) / float64(len(protoData))
	fmt.Printf("  JSON:     %d bytes (可读文本，包含字段名)\n", len(jsonData))
	fmt.Printf("  Protobuf: %d bytes (二进制，用字段编号替代字段名)\n", len(protoData))
	fmt.Printf("  JSON 是 Protobuf 的 %.1fx 大\n\n", sizeRatio)

	fmt.Println("Protobuf 优势：")
	fmt.Println("  1. 体积小 3-10x：字段名用编号(varint)代替，不传输 \"name\" \"age\" 等字符串")
	fmt.Println("  2. 解析快 20-100x：二进制直接映射到内存，无需词法/语法分析")
	fmt.Println("  3. 强类型：enum 确保 Gender 只能是 0/1/2，JSON 可能传 \"male\"/\"MALE\"/0")
	fmt.Println("  4. 向后兼容：加字段只需新编号，老客户端忽略未知字段")
	fmt.Println("  5. 字段编号 1-15 占 1 字节，16-2047 占 2 字节 → 高频字段用 1-15")

	fmt.Println()
	fmt.Println("JSON 优势：")
	fmt.Println("  1. 人类可读：浏览器/curl 直接查看，调试方便")
	fmt.Println("  2. 通用标准：所有语言内置支持，无需代码生成")
	fmt.Println("  3. Web 原生：浏览器 fetch 直接解析")
	fmt.Println()
	fmt.Println("推荐：微服务间通信用 Protobuf（gRPC），对外 API 用 JSON（REST）")
}
