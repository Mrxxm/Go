package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"

	"Go/6_grpc/grpc_token_test/proto"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Errorf(codes.Unauthenticated, "无Token认证信息")
		}

		var (
			appid  string
			appkey string
		)

		if val, ok := md["appid"]; ok {
			appid = val[0]
		}

		if val, ok := md["appkey"]; ok {
			appkey = val[0]
		}

		if appid != "101011" || appkey != "xx" {
			return resp, status.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
		}

		// 继续处理请求
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	s := grpc.NewServer(opts...)
	ser := &Server{}
	proto.RegisterGreeterServer(s, ser)
	s.Serve(lis)
}
