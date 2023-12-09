package main

import (
	"Thought/tree"
	"fmt"
)

// 内嵌
type myTreeNode struct {
	*tree.Node //内嵌
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}
func (myNode *myTreeNode) Traverse() {
	//	效果类似重载
	fmt.Println("This Traverse func")
}
func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	//new 内建函数
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreatNode(2)
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	var pRoot *tree.Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	root.Traverse()
	myroot := myTreeNode{&root}
	myroot.postOrder()
	myroot.Node.Traverse()

	nodeCount := 0
	root.TraverseFun(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)
}
