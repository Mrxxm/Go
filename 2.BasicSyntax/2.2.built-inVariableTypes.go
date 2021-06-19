package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 复数的模
func complexNumber() {
	var c = 3 + 4i
	// 复数的模
	fmt.Println(cmplx.Abs(c))
}


// 欧拉公式
func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}


// 类型转化
func triangle() {
	var a, b float64 = 3, 4
	var c int
	// 开平方根
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(c)
}


func main() {

	complexNumber()
	euler()

	triangle()
}
