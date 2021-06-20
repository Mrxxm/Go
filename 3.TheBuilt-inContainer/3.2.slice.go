package main

import "fmt"

// slice本身没有数据，是对底层array的一个view

// 参数传递slice
func changeArr(arr []int) {
	arr[0] = 100
}

func sliceExtends() {
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr1[2:6]
	s2 := s1[3:5]
	fmt.Println("arr1=", arr1)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

// 向slice中添加元素
func sliceExtendsAdd() {
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr1[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("arr1=", arr1)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	// s4 、s5 view another array view
	// 添加元素时如果超过cap，系统会重新分配更大的底层数组
	// 由于值传递的关系，必须接收append的返回值(append操作可能会使slice的len和cap变大，需要新的slice接收)
	// s = append(s, val)
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5=%v, len(s5)=%d, cap(s5)=%d\n", s5, len(s5), cap(s5))

}

func printSlice(slice []int) {
	fmt.Printf("slice=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
}

// 创建slice
func createSlice() {
	var slice1 []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		slice1 = append(slice1, 2 * i + 1)
		printSlice(slice1)
	}
	fmt.Println(slice1)
}

func update(arr []int) {
	arr[0] = 100
}

func copyingSlice() {
	// 1.
	s1 := []int{2, 4, 6, 8}
	// 2.
	s2 := make([]int, 16)

	printSlice(s1)
	printSlice(s2)
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
}

func deletingSlice() {
	// 1.
	s1 := []int{2, 4, 6, 8}

	// 2.
	s2 := make([]int, 16)
	copy(s2, s1)

	printSlice(s2)
	// 将s2中的元素8删除
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
}

func poppingFromFrontOrBack() {
	// 1.
	s1 := []int{2, 4, 6, 8}

	fmt.Println("Popping from front")
	front := s1[0]
	s1 = s1[1:]
	fmt.Println("front =", front)
	fmt.Println("s1 =", s1)

	fmt.Println("Popping from back")
	back := s1[len(s1) - 1]
	s1 = s1[:len(s1) - 1]
	fmt.Println("back =", back)
	fmt.Println("s1 =", s1)
}

func main() {

	// 1.slice引用传递
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	//update(s)
	//fmt.Println("s =", s)
	//fmt.Println("arr =", arr)

	// 2.Reslice
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} // 初始化数组
	//s2 := arr[:]
	//fmt.Println("s2 =", s2)
	//s2 = s2[:5]
	//fmt.Println("s2 =", s2)
	//s2 = s2[2:]
	//fmt.Println("s2 =", s2)

	// 3.slice扩展
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	//fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	//fmt.Println(arr[2:5])
	//fmt.Println(arr[:])
	//changeArr(arr[3:4])
	//fmt.Println(arr)

	//sliceExtends()
	// 4.slice添加元素
	//sliceExtendsAdd()

	// 5.创建slice
	//createSlice()
	// 1.
	//s1 := []int{2, 4, 6, 8}
	//
	//// 2.
	//s2 := make([]int, 16)
	//
	//// 3.
	//s3 := make([]int, 10, 32)
	//printSlice(s1)
	//printSlice(s2)
	//printSlice(s3)

	// 6.Copying slice
	//copyingSlice()

	// 7.Deleting slice
	//deletingSlice()

	// 8.popping
	poppingFromFrontOrBack()
}
