package main

import (
	"fmt"
)

// 切片的值改变之后原本的值也会改 因为切片是数组的视图
// 如果一个数组要改变值可以变为切片 array--》Slice arr[]--->arr[:]
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 := [...]int{1, 2, 3}
	//切片
	fmt.Println("s := arr[2:6]", arr[2:6])
	fmt.Println("s := arr[:6]", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1=", s1)
	s2 := arr[:]
	fmt.Println("s2=", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println(arr1)
	fmt.Println("Affter updateSlice(arr1)")
	updateSlice(arr1[:])
	fmt.Println(arr1)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice ")
	arr[0], arr[2] = 0, 2
	fmt.Println(arr)
	s1 = arr[2:6]
	s2 = s1[3:5] //虽然取不到值从s1上，但是底层的arr还是存在的所以最后取到的值依然是arr的值
	fmt.Printf("s1=%v , len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v , len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
}
