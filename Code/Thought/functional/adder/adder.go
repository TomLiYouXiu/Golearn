package main // 定义包名

import "fmt" // 导入fmt包，用于输出数据

// adder函数定义了一个匿名函数，用于实现一个累加器功能
func adder() func(int) int {
	sum := 0                     // 定义一个变量sum，用于存储累加和
	return func(value int) int { // 返回一个匿名函数，用于实现累加功能
		sum += value // 累加value值到sum中
		return sum   // 返回累加后的sum值
	}
}

// 定义一个接口iAdder，用于表示一个累加器函数
type iAdder func(int) (int, iAdder)

// adder2函数定义了一个匿名函数，用于实现另一个累加器功能
func adder2(base int) iAdder {
	return func(value int) (int, iAdder) { // 返回一个匿名函数，用于实现累加功能
		return base + value, adder2(base + value) // 返回累加后的值和一个新的累加器函数
	}
}

// main函数是程序的入口点
func main() {
	a := adder2(0)            // 创建一个累加器函数a，初始值为0
	for i := 0; i < 10; i++ { // 循环10次
		var s int      // 定义一个变量s，用于存储累加结果
		s, a = a(i)    // 使用累加器函数a，计算第i次的累加结果，并将结果存储到变量s中
		fmt.Println(s) // 输出累加结果
	}
}
