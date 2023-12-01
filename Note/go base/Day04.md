# 数组

## 数组的定义

![](https://pic.imgdb.cn/item/655309abc458853aef540c27.jpg)

~~~go
package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)
	var grid [4][5]int
	fmt.Println(grid)
}
~~~

## 数组的遍历

<img src="https://pic.imgdb.cn/item/65530ab8c458853aef573956.jpg" style="zoom:50%;" />

~~~go
	//数组的一般遍历方式
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	//通过range关键字遍历
	for _, v := range arr3 {
		fmt.Println(v)
	}
~~~

* 可通过_省略变量
* 不仅range，任何地方都可以通过_省略变量
* 如果只要下标可以只写成  **for i := range numbers**

## why range

* 意义明确 美观
* C++没有类似的能力
* java/python :只能for each Value ，不能同时获取i，v

## 数组是值类型

值类型和引用的类型的区别要注意

<img src="https://pic.imgdb.cn/item/65531023c458853aef686081.jpg" style="zoom:50%;" />

# 切片（Slice）

<img src="https://pic.imgdb.cn/item/655310a9c458853aef69d875.jpg" style="zoom:50%;" />

~~~go
package main

import "fmt"

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
}
~~~

<img src="https://pic.imgdb.cn/item/65531435c458853aef743726.jpg" style="zoom:50%;" />

<img src="https://pic.imgdb.cn/item/65531516c458853aef773bba.jpg" style="zoom:50%;" />

## Slice的扩展

<img src="https://pic.imgdb.cn/item/65531b6dc458853aef88ee1c.jpg" style="zoom:50%;" />

## Slice的实现

<img src="https://pic.imgdb.cn/item/65531bc8c458853aef89e7d4.jpg" style="zoom:50%;" />

<img src="https://pic.imgdb.cn/item/65531c2dc458853aef8b0640.jpg" style="zoom:50%;" />

 

以上两部分的实现代码

~~~go
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
}
~~~

## 向Slice添加元素

<img src="https://pic.imgdb.cn/item/65531f7ec458853aef94a511.jpg" style="zoom:50%;" />

## Slice的其他操作

### Slice的创建

~~~go
	var s []int //zero value for slice is nil
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32) //参数含义 指定类型 指定长度 指定cap长度
~~~

### Slice的Copy

~~~go
	fmt.Println("Copying Slice")
	copy(s2, s1)
	printSlice(s2)
~~~

### Slice的删除

~~~go
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
~~~

# Map（键值对）

<img src="https://pic.imgdb.cn/item/65532551c458853aefa421b0.jpg" style="zoom:50%;" />

## Map的创建

~~~go
	m := map[string]string{
		"mame":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
~~~

## Map的遍历

~~~go
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
~~~

* map在里面是无序的（hashmap）

## Map的值的获取

~~~go
	fmt.Println("Getting values")
	coursename, ok := m["course"]
	fmt.Println(coursename, ok)
	if caursename, ok := m["caurse"]; ok {
		fmt.Println(caursename)
	} else {
		fmt.Println("Key dose not exist")
	}
~~~

* 通过键值对获取
* 要是没有值得到的是空串
* 如果值是空值可以通过再加一个参数进行值的获取

## Map的删除

~~~go
	fmt.Println("Deleting values ")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
~~~

<img src="https://pic.imgdb.cn/item/65532aabc458853aefb21567.jpg" style="zoom:50%;" />

 

<img src="https://pic.imgdb.cn/item/65532ad0c458853aefb26f11.jpg" style="zoom:50%;" />



<img src="https://pic.imgdb.cn/item/65532b17c458853aefb322a2.jpg" style="zoom:50%;" />

## Map的一些简单问题

![](https://pic.imgdb.cn/item/65532c40c458853aefb67ebd.jpg)

~~~go
package main

import "fmt"

// 寻找最长的不含有重复字符的子串
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("v"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("asdfghjk"))
}
~~~

# 字符和字符串的处理

~~~go
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
~~~

![](https://pic.imgdb.cn/item/65533780c458853aefd2b091.jpg)

![](https://pic.imgdb.cn/item/6553384fc458853aefd4789e.jpg)

## 其余的可以详见strings库
