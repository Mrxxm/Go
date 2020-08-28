package main
// 2.4循环语句

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// 转二进制 5/2 101
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	} 
	return result
}

// 读取文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if (err != nil) {
		panic(err)
	} 

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
	) 
	printFile("abc.txt")
}