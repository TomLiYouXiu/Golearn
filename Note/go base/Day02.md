#  内建变量类型

![](https://pic.imgdb.cn/item/65474f11c458853aefd7701b.jpg)

## 关于复数的代码

~~~go
// 欧拉公式的验证
func euler() {
	////复数的定义
	//c := 3 + 4i
	////复数的库
	//cmplx.Abs(c)
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
	//cmplx.Pow(math.E, 1i*math.Pi) + 1)
}
~~~

# 强制类型转换

![](https://pic.imgdb.cn/item/654752ecc458853aefe4ae9c.jpg)

~~~go
// 强制转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
~~~

# 常量的定义

![](https://pic.imgdb.cn/item/654755dcc458853aefefa8e8.jpg)

~~~go
// 常量的定义
func consts() {
	const filename string = "abc.txt"
	//类型也可以不规定 也可定义在全剧局 也可以使用组定义
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
~~~

## 枚举的使用

~~~go
// 使用枚举类型
func enums() {
	const (
		// iota 表示自增 _表示占位
		cpp = iota
		_
		python
		goland
		php
	)
	const (
		//iota也可以参加运算
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, php, python, goland)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
~~~

![](https://pic.imgdb.cn/item/65475806c458853aeff757d7.jpg)

# 条件语句

## if

![](https://pic.imgdb.cn/item/65475a48c458853aefff7614.jpg)

~~~go
func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
~~~

## switch

![](https://pic.imgdb.cn/item/65475adac458853aef01f421.jpg)

~~~go
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
~~~

写法优化

~~~go
func grade(score int) string {
	g := ""
	switch {
	case score<0||score>100:
		panic(fmt.Sprintf(
			"Wrong score:%d",score))
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
~~~

## for

~~~go
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		//strconv.Itoa 整形转字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 打开文件并且按行读取
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//部分条件可以省略
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	//类似while
	for {
		fmt.Println("abc")
	}
}
~~~

在go中for类似while

 ![](https://pic.imgdb.cn/item/65483931c458853aef64df54.jpg)

