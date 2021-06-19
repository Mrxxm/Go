package main

import (
	"fmt"
	"runtime"
	"reflect"
)

func div1(a, b int) (int, int) {
	return a / b, a % b
}

// 与上面等同
func div2(a, b int) (q, r int) {
	return a / b, a % b
}

// 与上面等同
func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func div4(a, b int) int {
	var c = a * b
	return c
}

func eval1(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
	return result, nil
}

func apply(funcName func(int, int) int, a, b int) int{
	fmt.Printf("Calling %s with %d, %d\n", runtime.FuncForPC(reflect.ValueOf(funcName).Pointer()).Name(), a, b)
	return funcName(a, b)
}

func sumArgs(values ...int) int {
	var sum = 0
	for i := range values {
		fmt.Println(values)
		sum += values[i]
	}
	return sum
}

func main() {
	fmt.Println(div1(1, 2))
	fmt.Println(div2(1, 2))
	fmt.Println(div3(1, 2))
	fmt.Println(div4(1, 2))
	if result, err := eval1(2, 3, "t"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	apply(func(x, y int) int {
		return x * y
	}, 5, 6)
	fmt.Println(sumArgs(1, 5, 7, 9))
	var arr = [5]int{1, 2, 3}
	fmt.Println(arr)
}