package main

import (
	"Go/10_nacos/config"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "192.168.15.21",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "f26b61d5-6074-44a2-8892-0c0d7e14db15", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel: "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// json转换映射成结构体
	jsonBytesContent := []byte(content)
	serverConfig := config.ServerConfig{}
	err = json.Unmarshal(jsonBytesContent, &serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	time.Sleep(3000 * time.Second)

}
