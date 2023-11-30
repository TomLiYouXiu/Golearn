package tree

import (
	"fmt"
)

// 自定义结构体
type Node struct {
	Value       int
	Left, Right *Node
}

// 给结构体定义方法 (接收者)
func (node Node) Print() {
	fmt.Println(node.Value)
}

// 第二种写法 只是方法在调用时会有不同
func print(node Node) {
	fmt.Println(node.Value)
}

// 一般是传值要是想改变的话就需要改为指针传地址
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. " +
			"Ignore.")
		return
	}
	node.Value = value
}

// 工厂函数
func CreatNode(vlaue int) *Node {
	return &Node{Value: vlaue}
}
