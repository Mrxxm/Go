package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
1.访问方式：http://127.0.0.1:8080/add?a=1&b=2
2.返回格式： json {"data": 3}
3.CallID的问题：r.URL.Path
*/
func handler(w http.ResponseWriter, r *http.Request) {
	// 1.页面返回输出
	fmt.Fprintf(w, "Hello, World!\n")

	// 2.打印路径
	// 解析参数
	_ = r.ParseForm()
	fmt.Println("path: ", r.URL.Path)

	// 3.获取参数
	// type Values map[string][]string - 映射类型，key是字符串，value是字符串切片
	intA, _ := strconv.Atoi(r.Form["a"][0])
	intB, _ := strconv.Atoi(r.Form["b"][0])
	fmt.Printf("a: %d, b: %d \n", intA, intB)

	// 4.返回json
	w.Header().Set("Content-Type", "application/json")
	jData, _ := json.Marshal(map[string]int{
		"data": intA + intB,
	})
	_, _ = w.Write(jData)
}

func main() {

	// 设置路径
	http.HandleFunc("/add", handler)
	// 服务启动打印标记
	fmt.Println("Starting server on :8080")
	// 监听端口
	http.ListenAndServe(":8080", nil)

}
