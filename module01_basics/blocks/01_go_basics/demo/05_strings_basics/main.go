package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := " Go 语言 "
	fmt.Println("concat:", text+" backend")
	fmt.Println("trim:", strings.TrimSpace(text))
	fmt.Println("contains Go:", strings.Contains(text, "Go"))
	fmt.Println("fields:", strings.Fields(text))
	fmt.Println("lower:", strings.ToLower("Go Backend"))

	text = "A你"
	fmt.Printf("%q bytes=%d\n", text, len(text))
	fmt.Printf("byte at index 1=%d\n", text[1])

	fmt.Println("range over string:")
	for index, r := range text {
		fmt.Printf("index=%d rune=%c code=%d\n", index, r, r)
	}

	// String indexes and slices use bytes; use range or []rune for characters.
	runes := []rune(text)
	fmt.Printf("runes=%v first-rune=%c\n", runes, runes[0])

	bytes := []byte("Go语言")
	roundTrip := string(bytes)
	fmt.Println("bytes round trip:", roundTrip)

	number, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Atoi unexpected error:", err)
	} else {
		fmt.Println("Atoi success:", number)
	}

	_, err = strconv.Atoi("forty-two")
	fmt.Println("Atoi error:", err != nil)
}
