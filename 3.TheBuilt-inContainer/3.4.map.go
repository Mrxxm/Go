package main

import "fmt"

func mapDemo() {
	m := map[string]string {
		"name"    :  "ccmounse",
		"course"  :  "goloang",
		"site"    :  "kenrou",
		"quality" :  "good",
	}
	fmt.Println("map = ", m)

	// 定义空map
	m2 := make(map[string]int)  // m2 == empty map
	var m3 map[string]int       // m3 == nil
	fmt.Println("map2 = ", m2)
	fmt.Println("map3 = ", m3)
}

func traversingMap() {
	m := map[string]string {
		"name"    :  "ccmounse",
		"course"  :  "goloang",
		"site"    :  "kenrou",
		"quality" :  "good",
	}
	for k, v := range m {
		fmt.Printf("k = %s, v = %s", k, v)
		fmt.Println()
	}
}

func operateMap() {
	m := map[string]string {
		"name"    :  "ccmounse",
		"course"  :  "goloang",
		"site"    :  "kenrou",
		"quality" :  "good",
	}
	courseName := m["course"]
	fmt.Println("course name = ", courseName)
	cseName := m["cse"]                      // key值填写错误
	fmt.Println("cse name = ", cseName)

	// 判断key是否存在
	courseName1, ok := m["course"]
	fmt.Println("course name1 = ", courseName1, ok)
	cseName1, ok := m["cse"]                      // key值填写错误
	fmt.Println("cse name1 = ", cseName1, ok)

	// 一般写法
	if cseName2, ok := m["cse"]; ok {
		fmt.Println("cse name2 = ", cseName2)
	} else {
		fmt.Println("key does not exist!")
	}
}

func deletingElements() {
	m := map[string]string {
		"name"    :  "ccmounse",
		"course"  :  "goloang",
		"site"    :  "kenrou",
		"quality" :  "good",
	}

	fmt.Println("Deleting elements")
	name, ok := m["name"]
	fmt.Printf("name = %s, ok = %v", name, ok)

	fmt.Println();

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("name = %s, ok = %v", name, ok)
}

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI := 0
		ok := false
		if lastI, ok = lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		//fmt.Printf("i = %v, ch = %v, lastI = %v, ok = %v, start = %v, lastOccurred = %v",
		//	i, ch, lastI, ok, start, lastOccurred)
		//fmt.Println();

		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	// 1.map定义
	//mapDemo()
	// 2.map遍历
	//traversingMap()
	// 3.map操作
	//operateMap()
	// 4.删除元素
	//deletingElements()
	// 5.寻找最长不含有重复字符的子串
	fmt.Println(lengthOfLongestSubstring("abcabcdabc"))
}
