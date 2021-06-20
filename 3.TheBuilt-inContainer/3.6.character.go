package main

import (
	"fmt"
	"unicode/utf8"
)

// 国际版 byte -> rune
func lengthOfLongestSubstring1(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
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

// 中文占3字节
func charDemo() {
	// step:1
	s := "Yes我爱Go" // utf-8可变长
	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	// step:2
	for i, ch := range s { // ch is a rune (string -> utf-8解码 -> 每个字符 -> 转unicode -> 最后放到rune类型中)
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	// step:3
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("ch = %c, size = %v", ch, size)
		fmt.Println()
		bytes = bytes[size:]

	}

	// step:4 转rune
	runes := []rune(s)
	for i, ch := range runes {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}

func otherChar() {
	
}

func main() {
	//charDemo()
	//lengthOfLongestSubstring1("abc")
}
