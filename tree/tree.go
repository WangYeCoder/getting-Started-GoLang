package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//工厂函数 返回局部变量地址
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) Traves() {
	node.Left.Traves()
	node.PrintMethods()
	node.Right.Traves()
}

func (node Node) PrintMethods() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}
