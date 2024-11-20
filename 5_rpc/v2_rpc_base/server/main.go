package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "hello, " + request

	return nil
}

func main() {
	// 1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3.启动服务
	conn, _ := listener.Accept() // 当一个新的连接进来的时候

	rpc.ServeConn(conn) // 只能处理一次连接

	// rpc调用中解决了两个问题：1.call id 2.序列化和反序列化
}
