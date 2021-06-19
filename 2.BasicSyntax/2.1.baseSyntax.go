package main

import "fmt"

// 包内变量
var a = 1
var b = 2
var c = "巴拉巴拉"

var d, e, f = 4, 5, "6"

var (
	g = 7
	h = 8
	i = "9"
)

func variableZeroValue() {
	var age int
	var str string
	fmt.Printf("%d %q\n", age, str)
}

func variableInitialValue() {
	var age, height = 25, 175
	var str = "今天学习go语言"
	// 变量定义了一定要被使用
	fmt.Println(age, height, str)
}

func variableTypeDeduction() {
	// 编译器自动判断类型
	var age, height, str = 25, 175, "今天学习go语言"
	fmt.Println(age, height, str)
}

func variableShorter() {
	// 只能在函数内使用
	age, height, str := 25, 175, "今天很开心"
	fmt.Println(age, height, str)
}

func main() {
	fmt.Println("hello go world!")
	fmt.Println(a, b, c, d, e, f, g, h, i)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
