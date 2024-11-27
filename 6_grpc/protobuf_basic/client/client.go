package main

import (
	"Go/6_grpc/protobuf_basic/proto"
	proto_bak "Go/6_grpc/protobuf_basic/proto-bak"
	"context"
	"fmt"
	//"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()

	_ = proto.Pong{
		Id: "31",
		G:  proto.Gender_FEMALE,
		Mp: map[string]string{
			"name": "xxm",
		},
		AddTime: timestamppb.New(time.Now()),
	}

	c := proto_bak.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto_bak.HelloRequest{})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
