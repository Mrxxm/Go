package main

import (
	"fmt"
	"io/ioutil"
)

// if 
func bounded(v int) bool {
	if v > 100 {
		return true
	} else if v < 0 {
		return false
	} else {
		return false	
	}
} 

// switch
func eval(a, b int, op string) int {
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
			panic("unsupported opeartor:" + op)
	}
	return result
}

func main() {
	// if 
	fmt.Println(bounded(101)) 

	// 读取文件信息
	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// switch
	fmt.Println(eval(2, 4, "+"))
}