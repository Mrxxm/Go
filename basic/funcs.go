package main

import (
	"fmt"
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

func main() {
	fmt.Println(div1(13, 3))
	fmt.Println(div2(13, 3))
	fmt.Println(div3(13, 3))
	q, _ := div3(13, 3)
	fmt.Println(q)
}