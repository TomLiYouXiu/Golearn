package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

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

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func main() {
	if result, err := eval(13, 4, "."); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	//q, r := div(13, 3)
	fmt.Println(
		div(13, 3))
	//匿名函数调用
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3))
}
