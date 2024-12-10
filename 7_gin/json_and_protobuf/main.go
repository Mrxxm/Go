package main

import (
	"Go/7_gin/json_and_protobuf/protobuf"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	routeGroup := router.Group("/gin_test")
	{
		routeGroup.GET("/more_json", moreJson)
		routeGroup.GET("/some_protobuf", someProtobuf)

	}

	router.Run(":8081")
}

func someProtobuf(context *gin.Context) {
	teacher := &protobuf.Teacher{
		Name:   "xxm",
		Course: []string{"math", "english"},
	}

	context.ProtoBuf(http.StatusOK, teacher)
}

func moreJson(context *gin.Context) {
	var message struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	message.Name = "xxm"
	message.Message = "hello"
	message.Number = 123

	context.JSON(http.StatusOK, message)
}
