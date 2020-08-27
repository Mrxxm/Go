package main

import "fmt"

// 包内变量
var a, b, c = 1, 2, "3"
// 使用括号
var (
	d = 4
	e = 5
	f = "6"
)

func variableZeroValue() {
	var age int
	var str string
	fmt.Printf("%d %q\n", age, str)
}

func variableInitialValue() {
	var age, height int = 25, 175
	var str string = "今天学习go语言"
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
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, c, d, e, f)
}