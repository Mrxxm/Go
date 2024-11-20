package main

import (
	"LearnGo/tree"
	"fmt"
)

type myTreeNode struct {
	node *treeDemo.Node
}

// 后续遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	myTreeNodeLeft := myTreeNode{myNode.node.Left}
	myTreeNodeLeft.postOrder()
	myTreeModeRight := myTreeNode{myNode.node.Right}
	myTreeModeRight.postOrder()
	myNode.node.Print()

}

func main() {
	initTree := treeDemo.InitStruct()
	initTree.Traverse()
	fmt.Println()
	myInitTree := myTreeNode{&initTree}
	myInitTree.postOrder()

}
