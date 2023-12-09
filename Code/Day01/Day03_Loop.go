package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//将数字转换为二进制
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
// 无限循环
func forever() {
	//类似while
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5), //101
		convertToBin(13),
		convertToBin(0)) // 1011 --> 1101

	printFile("abc.txt")
	forever()
}