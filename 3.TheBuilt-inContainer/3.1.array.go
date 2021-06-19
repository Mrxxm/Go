package main

import "fmt"

// 数组 值类型

func main() {
	var arr1 [5]int         // 定义五个int的数组
	arr2 := [3]int{1, 3, 5} // := 需要给出初始值
	arr3 := [...]int{}      // 自动计算长度
	var arr4 [4][5]int      // 二维数组
	fmt.Println(arr1, arr2, arr3, arr4)

	// for循环遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// range循环遍历
	for i := range arr2 {
		fmt.Println(arr2[i])
	}

	// i是下标 v是数组的值
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	// 可通过_省略变量
	// 不仅range，任何地方都可通过_省略变量
	for _, value := range arr2 {
		fmt.Println(value)
	}
 }