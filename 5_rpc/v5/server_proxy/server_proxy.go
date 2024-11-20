package server_proxy

import (
	"Go/5.rpc/v5/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

// 传递接口 鸭子类型 实现解耦
func RegisterHelloService(svc HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, svc)
}
