package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v len=%d,cap=%d\n",
		s, len(s), cap(s))
}
func main() {
	fmt.Printf("Creating Slice")
	var s []int //zero value for slice is nil
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32) //参数含义 指定类型 指定长度 指定cap长度
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i)
	}
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	//删除下标为3的元素
	s2 = append(s2[:3], s2[4:]...) //s2[4:]...将后面的所有元素进行拼接
	printSlice(s2)
	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)
	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
