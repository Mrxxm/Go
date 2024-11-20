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

	root = TreeNode{Value:2}
	root.Left = &TreeNode{} // Left 是指针，所以传递给Left是变量地址，需要加&号
	root.Left.Right = createNode(3)
	root.Right = &TreeNode{nil, nil, 4}
	root.Right.Left = new(TreeNode)
	root.Right.Left.setValue2(5)

	root.traverse()
}

// 工厂函数
func createNode(value int) *TreeNode {
	return &TreeNode{Value: value} // 返回局部变量的地址 给全局使用 C++中程序会挂
}

// (node TreeNode) - 接收者
func (node TreeNode) print() {
	fmt.Print(node.Value)
}

func print(node TreeNode) {
	fmt.Print(node.Value)
}

// 值传递
func (node TreeNode) setValue1 (value int) {
	node.Value = value
}

// 引用传递
func (node *TreeNode) setValue2 (value int) {
	node.Value = value
}

// 中序遍历
func (node *TreeNode) traverse() {
	if (node == nil) {
		return
	}
	node.Left.traverse()
	node.print()
	node.Right.traverse()
}

func main() {
	initStruct()

}
