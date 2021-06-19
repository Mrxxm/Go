package main

import "fmt"

/**
 * 1.指针不能运算
 * 2.值传递 引用传递
 * 3.go语言只有值传递一种方式
 * 4.
 */

 func swap(a, b int) (int, int) {
 	b, a = a, b
 	return a, b
 }

 func swap2(a, b *int) {
 	*b, *a = *a, *b
 }

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	fmt.Println(swap(1, 2))
	c, d := 3, 4
	swap2(&c, &d)
	fmt.Println(c, d)
}