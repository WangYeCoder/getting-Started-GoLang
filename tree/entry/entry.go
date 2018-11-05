package main

import (
	"../../tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNodeFun *myTreeNode) postOrder() {
	if myNodeFun == nil {
		return
	}
	myTreeNode{myNodeFun.node.Left}.postOrder()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 4}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	root.Right.Left.PrintMethods()

	fmt.Println(nodes)

	root.PrintMethods()
	root.SetValue(3)

	pRoot := &root
	pRoot.PrintMethods()
	pRoot.SetValue(123123)

	fmt.Println(pRoot)

	//root.Traves()

}
