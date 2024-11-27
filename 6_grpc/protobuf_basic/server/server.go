package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

// HelloRequest、HelloReply为自动生成
func (s *Server) SayHello(ctx context.Context, request *proto_bak.HelloRequest) (*proto_bak.HelloReply, error) {
	return &proto_bak.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	// 1.实例化一个server
	g := grpc.NewServer()
	// 2.注册处理逻辑handler(RegisterGreeterServer为自动生成)
	proto_bak.RegisterGreeterServer(g, &Server{})
	// 3.启动服务
	listener, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	_ = g.Serve(listener)
}
