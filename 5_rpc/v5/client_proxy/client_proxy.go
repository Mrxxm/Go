package client_proxy

import (
	"Go/5.rpc/v5/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol, address string) *HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("connect error")
	}
	return &HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	return c.Client.Call(handler.HelloServiceName+".Hello", request, reply)
}
