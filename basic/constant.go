package main
// 2.3常量与枚举

import (
	"fmt"
	"math"
)

func consts() {
	// 未指定常量类型时，是文本类型
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int 
	c = int(math.Sqrt(a * a + b * b))

	fmt.Println(filename, a, b)
	fmt.Println(c)
}

func enums() {
	// const(
	// 	cpp = 0
	// 	java = 1
	// 	python = 2
	// 	golang = 3
	// )
	// 自增枚举
	const(
		cpp = iota
		java 
		python 
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func main() {
	consts()
	enums()
}