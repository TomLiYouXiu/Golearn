package main

import "fmt"

// 交换两个变量
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
