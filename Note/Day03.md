# 函数

与其他语言相比go还可以返回多个类型的返回值，不需要列表或者再次定义对象，较为简约

~~~go
// 正常函数
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}

}

// 可以返回多个类型
// 13 / 4 = 3 ... 1
// 如果此时想要单独的调用某个参数可以使用_替代不想调用的那个参数
func div(a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	//return q, r
	return a / b, a % b

}
~~~

![](https://pic.imgdb.cn/item/65484105c458853aef74fa0a.jpg)

## 在函数中调用另一个函数

~~~go
func apply(op func(int, int) int, a, b int) int {
	//获得当前函数的指针
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func main(){
    fmt.Println(apply(pow, 3, 4))
}
~~~

也可以使用匿名函数调用

~~~go
func apply(op func(int, int) int, a, b int) int {
	//获得当前函数的指针
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}
func main(){
//匿名函数调用
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))
}
~~~

![](https://pic.imgdb.cn/item/65484581c458853aef7e8715.jpg)

## 可变参数列表

~~~go
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
~~~

![](https://pic.imgdb.cn/item/6548468ec458853aef80ddfe.jpg)

# 指针

在go中指针不能运算

## 参数传递

值传递，引用传递（地址传递）

但是go语言只有值传递一种方式，但是通过指针可以代替引用传递（节省内存空间）

~~~go
// 交换两个变量
func swap(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
~~~

另一种交换思路

~~~go
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
~~~

