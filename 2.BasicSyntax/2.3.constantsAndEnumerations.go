package main

import (
	"math"
	"fmt"
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
	// iota表示自增枚举
	const(
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	consts()
	enums()
}