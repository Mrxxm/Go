package treeDemo

import "fmt"

type Node struct {
	Left, Right *Node
	Value int
}

func InitStruct() {
	// 1.初始化struct
	// 方式一
	var root Node

	root = Node{Value:2}
	root.Left = &Node{} // Left 是指针，所以传递给Left是变量地址，需要加&号
	root.Left.Right = CreateNode(3)
	root.Right = &Node{nil, nil, 4}
	root.Right.Left = new(Node)
	root.Right.Left.SetValue2(5)

	root.Traverse()
}

// 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value} // 返回局部变量的地址 给全局使用 C++中程序会挂
}

// (node Node) - 接收者
func (node Node) Print() {
	fmt.Print(node.Value)
}

// 值传递
func (node Node) setValue1 (value int) {
	node.Value = value
}

// 引用传递
func (node *Node) SetValue2 (value int) {
	node.Value = value
}

// 中序遍历
func (node *Node) Traverse() {
	if (node == nil) {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
