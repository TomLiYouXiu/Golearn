package fib

// 定义一个生成器函数，用于生成斐波那契数列
func Fibonacci() func() int {
	a, b := 0, 1        // 初始化斐波那契数列的第一个和第二个数
	return func() int { // 返回一个函数，用于生成斐波那契数列的下一个数
		a, b = b, a+b // 更新斐波那契数列的值
		return a      // 返回斐波那契数列的当前值
	}
}
