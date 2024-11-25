package main

import (
	"Go/6_grpc/stream_grpc_test/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamReqData, serverStr proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = serverStr.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}
func (s *server) PutStream(clientStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := clientStr.Recv(); err == nil {
			fmt.Println("client:" + a.Data)
		} else {
			break
		}
	}

	return nil
}
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			if a, err := allStr.Recv(); err == nil {
				fmt.Println("收到客户消息:" + a.Data)
			} else {
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{
				Data: fmt.Sprintf("服务端发送%v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	// 1.实例化一个server
	g := grpc.NewServer()
	// 2.注册处理逻辑handler(RegisterGreeterServer为自动生成)
	proto.RegisterGreeterServer(g, &server{})
	// 3.启动服务
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	_ = g.Serve(listener)
}
