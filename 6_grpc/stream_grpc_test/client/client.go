package main

import (
	"Go/6_grpc/stream_grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()

	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	r, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "xxm"})

	for {
		a, err := r.Recv()
		if err != nil {
			break
		}
		fmt.Println("server:" + a.Data)
	}

	// 客户端流模式
	r2, _ := c.PutStream(context.Background())

	for i := 0; i < 10; i++ {
		_ = r2.Send(&proto.StreamReqData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(time.Second)
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			if a, err := allStr.Recv(); err == nil {
				fmt.Println("收到服务端消息:" + a.Data)
			} else {
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{
				Data: fmt.Sprintf("客户端发送消息%v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
