package main

import (
	"Go/6_grpc/proto/protobuf"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {

	// 1.protobuf
	req := protobuf.HelloRequest{
		Name:    "xxm",
		Age:     18,
		Courses: []string{"go", "python", "java"},
	}
	marshal, err := proto.Marshal(&req)
	if err != nil {
		return
	}

	fmt.Println(len(marshal))

	// 2.json
	jsonStruct := Hello{
		Name:    "xxm",
		Age:     18,
		Courses: []string{"go", "python", "java"},
	}
	jsonRes, _ := json.Marshal(jsonStruct)

	fmt.Println(len(jsonRes))

	// 3.反解析Protobuf
	newResult := protobuf.HelloRequest{}
	_ = proto.Unmarshal(marshal, &newResult)
	fmt.Println(newResult.Name, newResult.Age, newResult.Courses)
}
