// 黑暗角落：String 本质、不可变、UTF-8，以及 len 与 range 的区别
package main

import "fmt"

// 演示 string 零值判断
func showStringZeroValue() {
	var s string
	fmt.Println("string 零值:", s == "")
	fmt.Println("string 零值长度:", len(s))
	// s = nil  // 编译错误：不能赋 nil 给 string
}

// 演示 len、下标与 range 的区别（UTF-8）
func showUTF8Behavior() {
	s := "touché你好"
	fmt.Println("字符串:", s)

	// len 返回字节数（不是字符数！）
	fmt.Printf("len(s) = %d (字节数)\n", len(s)) // touché 占 7 字节，你好占 6 字节，共 13 字节

	// 按字节遍历（中文会乱码）
	fmt.Print("按字节遍历: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()

	// range 按 rune 遍历（正确）
	fmt.Print("按 rune 遍历: ")
	for _, r := range s {
		fmt.Printf("%c", r)
	}
	fmt.Println()

	// 统计 rune 数量
	fmt.Printf("rune count: %d\n", len([]rune(s)))
}

// 演示 string 不可变
func showStringImmutability() {
	s := "hello"
	// s[0] = 'H' // 编译错误：不能给 s[0] 赋值

	// 要修改需要用 []byte 或 []rune
	b := []byte(s)
	b[0] = 'H'
	s = string(b)
	fmt.Println("修改后:", s)
}

func main() {
	showStringZeroValue()
	showUTF8Behavior()
	showStringImmutability()
}
