package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFun(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}
func (node *Node) TraverseFun(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFun(f)
	f(node)
	node.Right.TraverseFun(f)
}
