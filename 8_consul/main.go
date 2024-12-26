package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	// 1.初始化配置
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	// 2.创建一个consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 3.注册服务&生成注册对象&生成检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.15.21:8022/health",
		Timeout:                        "5s",  // 超时时间
		Interval:                       "5s",  // 健康检查间隔
		DeregisterCriticalServiceAfter: "10s", // 多久后注销服务
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	return nil
}

// 获取所有服务
func AllServices() {
	// 1.初始化配置
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	// 2.创建一个consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 3.获取所有服务
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for k, _ := range data {
		fmt.Println(k)
	}
}

// 过滤服务
func FilterServices() {
	// 1.初始化配置
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	// 2.创建一个consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 3.获取所有服务
	data, err := client.Agent().ServicesWithFilter(`Service == "shop-api"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	// 启动调用接口
	//err := Register("127.0.0.1", 8022, "shop-api", []string{"shop", "api"}, "shop-api")
	//if err != nil {
	//	panic(err)
	//}
	//AllServices()
	FilterServices()
}
