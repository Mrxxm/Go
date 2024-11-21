package main

import (
	"Go/6_grpc/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "xxm"})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
