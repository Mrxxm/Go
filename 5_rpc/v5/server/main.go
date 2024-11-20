package main

import (
	"Go/5_rpc/v5/handler"
	"Go/5_rpc/v5/server_proxy"
	"net"
	"net/rpc"
)

func main() {

	// 1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑handler
	_ = server_proxy.RegisterHelloService(&handler.HelloService{})
	// 3.启动服务
	for {
		conn, _ := listener.Accept() // 当一个新的连接进来的时候
		go rpc.ServeConn(conn)       // 多个请求过来时，需要等待处理，改成协程即可解决
	}
}
