package main

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Value int
}

func initStruct() {
	// 1.初始化struct
	// 方式一
	var root TreeNode

	root = TreeNode{Value:3}
	root.Left = &TreeNode{} // Left 是指针，所以传递给Left是变量地址，需要加&号
	root.Right = &TreeNode{nil, nil, 5}
	root.Right.Left = new(TreeNode)

	nodes := []TreeNode {
		{Value :  5},
		{nil, &root, 6},
	}
	fmt.Println(nodes)
}

func main() {
	initStruct()

}
