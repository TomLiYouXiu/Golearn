package main // 定义包名

import ( // 导入所需的包
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 定义一个生成器函数，用于生成斐波那契数列
func fibonacci() intGen {
	a, b := 0, 1 // 初始化斐波那契数列的第一个和第二个数
	return func() int { // 返回一个函数，用于生成斐波那契数列的下一个数
		a, b = b, a+b // 更新斐波那契数列的值
		return a      // 返回斐波那契数列的当前值
	}
}

// 定义一个接口类型，用于实现生成器函数
type intGen func() int

// 定义一个函数，用于处理生成器函数的读取操作
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()       // 获取生成器函数的下一个值
	if next > 10000 { // 如果下一个值大于10000，表示生成器已经完成
		return 0, io.EOF // 返回0和EOF错误
	}
	s := fmt.Sprintf("%d\n", next) // 格式化字符串，包含下一个值和换行符
	//TODO：incorrect if p is too small！

	return strings.NewReader(s).Read(p) // 使用格式化后的字符串创建一个Reader对象，并读取到p中
}

// 定义一个函数，用于打印文件内容
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader) // 创建一个Scanner对象，用于逐行读取文件内容
	for scanner.Scan() {                // 逐行读取文件内容
		fmt.Println(scanner.Text()) // 打印每一行的内容
	}
}

// 定义主函数，用于启动程序
func main() {
	f := fibonacci()     // 获取斐波那契生成器函数
	printFileContents(f) // 打印斐波那契数列的内容
}
