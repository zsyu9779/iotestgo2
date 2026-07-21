package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := " Go, Go, 你好 "
	fmt.Println("contains:", strings.Contains(text, "你好"))
	fmt.Println("index:", strings.Index(text, "Go"))
	fmt.Println("replace:", strings.Replace(text, "Go", "Rust", 1))
	fmt.Println("split:", strings.Split("a-b-c", "-"))
	fmt.Println("join:", strings.Join([]string{"a", "b", "c"}, ","))
	fmt.Println("equal fold:", strings.EqualFold("Go", "go"))
	fmt.Printf("trim=%q fields=%q lower=%q\n", strings.TrimSpace(text), strings.Fields(text), strings.ToLower(text))

	text = "A你"
	fmt.Printf("text=%q bytes=%d runes=%d\n", text, len(text), utf8.RuneCountInString(text))
	fmt.Println("byte values:")
	for index := 0; index < len(text); index++ {
		fmt.Printf("index=%d byte=%d\n", index, text[index])
	}
	fmt.Println("rune values:")
	for index, r := range text {
		fmt.Printf("byte-index=%d rune=%c\n", index, r)
	}

	runes := []rune(text)
	fmt.Printf("rune slice=%v second-rune=%c\n", runes, runes[1])
	// String 切片使用字节偏移；按字符处理时使用 []rune。
}
