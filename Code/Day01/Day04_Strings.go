package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "成功的秘诀就是每天都比别人多努力一点" //UTF-8
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("%d %X\n", i, ch)
	}
	fmt.Println(
		"RuneCountInString:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("(%d %c) ", size, ch)
	}
	fmt.Println()
}
