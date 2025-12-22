package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

// 1. 正则表达式 (Regexp)
func demoRegexp() {
	fmt.Println("--- 1. Regexp Demo ---")
	
	// 基础匹配
	text := "My email is test@example.com and backup is backup@test.org"
	emailReg := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	
	emails := emailReg.FindAllString(text, -1)
	fmt.Printf("Found emails: %v\n", emails)

	// 分组提取 (Submatch)
	// 提取: (数字).(数字).(数字)
	verReg := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	versionStr := "App version 1.2.3 released"
	
	matches := verReg.FindStringSubmatch(versionStr)
	if len(matches) > 0 {
		fmt.Printf("Full match: %s\n", matches[0])
		fmt.Printf("Major: %s, Minor: %s, Patch: %s\n", matches[1], matches[2], matches[3])
	}
}

// 2. JSON 处理 (Encoding/JSON)
type User struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Roles     []string `json:"roles,omitempty"` // omitempty: 如果为空则不生成该字段
	Password  string   `json:"-"`               // -: 忽略该字段
	CreatedAt int64    `json:"created_at"`
}

func demoJSON() {
	fmt.Println("\n--- 2. JSON Demo ---")

	// Struct -> JSON
	user := User{
		ID:        101,
		Username:  "gopher",
		Roles:     []string{"admin", "editor"},
		Password:  "secret",
		CreatedAt: time.Now().Unix(),
	}
	
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("JSON Output:")
	fmt.Println(string(jsonData))

	// JSON -> Struct
	jsonStr := `{"id": 102, "username": "guest", "created_at": 1678888888}`
	var user2 User
	if err := json.Unmarshal([]byte(jsonStr), &user2); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Struct Parsed: %+v\n", user2)
}

// 3. Base64 编码
func demoBase64() {
	fmt.Println("\n--- 3. Base64 Demo ---")
	
	msg := "Hello, Go Standard Library!"
	
	// Encode
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Printf("Original: %s\n", msg)
	fmt.Printf("Base64:   %s\n", encoded)
	
	// Decode
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Printf("Decoded:  %s\n", string(decoded))
}

// 4. SHA1 哈希 (Crypto)
func demoHash() {
	fmt.Println("\n--- 4. SHA1 Hash Demo ---")
	
	data := "password123"
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("SHA1: %x\n", bs)
}

// 5. 时间处理 (Time)
func demoTime() {
	fmt.Println("\n--- 5. Time Demo ---")
	
	now := time.Now()
	fmt.Printf("Current Time: %v\n", now.Format("2006-01-02 15:04:05")) // 固定 Layout
	
	// 时间戳转换
	timestamp := now.Unix()
	fmt.Printf("Timestamp: %d\n", timestamp)
	
	tFromTs := time.Unix(timestamp, 0)
	fmt.Printf("Time from TS: %v\n", tFromTs.Format(time.RFC3339))
	
	// 时间计算
	tomorrow := now.Add(24 * time.Hour)
	fmt.Printf("Tomorrow: %v\n", tomorrow.Format("2006-01-02"))
	
	// 解析字符串
	layout := "2006-01-02"
	str := "2023-12-25"
	t, _ := time.Parse(layout, str)
	fmt.Printf("Parsed Time: %v\n", t)
}

func main() {
	demoRegexp()
	demoJSON()
	demoBase64()
	demoHash()
	demoTime()
}
