package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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

// 强制转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量的定义
func consts() {
	const filename string = "abc.txt"
	//类型也可以不规定 也可定义在全剧局 也可以使用组定义
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

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

func main() {
	euler()
	triangle()
	consts()
	enums()
}
