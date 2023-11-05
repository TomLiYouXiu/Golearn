package main

import (
	"fmt"
	"io/ioutil"
)

// switch跟其他语言相比没有break 其实是方便的
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("对不起没有找到这个符号：" + op)
	}
	return result
}
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(eval(3, 4, "+"))
	fmt.Println(grade(95))
}
