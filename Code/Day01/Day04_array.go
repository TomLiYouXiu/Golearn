package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100 //但是值没有发生变化
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int
	fmt.Println(grid)
	fmt.Println("printArray(arr3)")
	printArray(&arr3)
	fmt.Println("printArray(arr1)")
	printArray(&arr1)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr2, arr3)

}
