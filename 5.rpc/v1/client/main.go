package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

// 5.rpc 远程调用，如何像本地调用一样

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	url := fmt.Sprintf("http://127.0.0.1:8080/add?a=%d&b=%d", a, b)

	res, _ := req.Get(url)
	body, _ := res.Body()
	fmt.Println(string(body))

	resData := ResponseData{}
	err := json.Unmarshal(body, &resData)

	if err != nil {
		fmt.Println(err.Error())
	}

	return resData.Data
}

func main() {

	fmt.Println(Add(3, 4))
}
