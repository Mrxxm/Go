package main

import (
	"Go/6_grpc/grpc_test/proto"
	"context"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

// HelloRequest、HelloReply为自动生成
func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	// 1.实例化一个server
	g := grpc.NewServer()
	// 2.注册处理逻辑handler(RegisterGreeterServer为自动生成)
	proto.RegisterGreeterServer(g, &Server{})
	// 3.启动服务
	listener, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	_ = g.Serve(listener)
}
