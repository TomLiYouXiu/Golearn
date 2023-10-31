# GO的学习

配置go环境和goland

~~~
ps：一些小技巧
1.在goland中的setting中的tools里有个File Watchers中设置goimports，在保存代码时会自动进行格式化较为方便
~~~

# Basic

## 变量的定义

~~~go
package main

import "fmt"

// 但是此时不可以使用:=定义变量
// 如果需要定义多个变量时可以使用var()
var (
	aa = 33
	ss = "kkk"
)

// go和其他的语言不太一样，进行初始化的时候会自动设置初值 int=0 string为空字符串
func variableZeroValue() {
	var a int
	var s string
	//Printf 类似于C语言的格式化输出
	fmt.Printf("%d %q\n", a, s)
}

// go语言中定义了变量一定要用到要不然会出错
func variableInitiateValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 变量中的类型也可以省略掉 跟python类似 甚至可以在一行中定义不同的类型
func variableTypeDeduction() {
	var a, b, s, c = 3, 4, "abc", true
	fmt.Println(a, b, s, c)
}

// 也可以省略var用：代替（在定义变量时，再次进行更改时不再需要：）
func variableShrter() {
	a, b, s, c := 3, 4, "abc", true
	b = 5
	fmt.Println(a, b, s, c)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitiateValue()
	variableTypeDeduction()
	variableShrter()
	fmt.Println(aa, ss)
}
~~~



