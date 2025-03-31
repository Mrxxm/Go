package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

type EsGoods struct {
}

func (EsGoods) GetMapping() string {
	goodsMapping := `
	{
		"mappings" : {
			"properties" : {
				"brands_id" : {
					"type" : "integer"
				},
				"category_id" : {
					"type" : "integer"
				},
				"click_num" : {
					"type" : "integer"
				},
				"fav_num" : {
					"type" : "integer"
				},
				"id" : {
					"type" : "integer"
				},
				"is_hot" : {
					"type" : "boolean"
				},
				"is_new" : {
					"type" : "boolean"
				},
				"market_price" : {
					"type" : "float"
				},
				"name" : {
					"type" : "text",
					"analyzer":"ik_max_word"
				},
				"goods_brief" : {
					"type" : "text",
					"analyzer":"ik_max_word"
				},
				"on_sale" : {
					"type" : "boolean"
				},
				"ship_free" : {
					"type" : "boolean"
				},
				"shop_price" : {
					"type" : "float"
				},
				"sold_num" : {
					"type" : "long"
				}
			}
		}
	}`
	return goodsMapping
}

type Account struct {
	AccountNumber int32  `json:"account_number"`
	FirstName     string `json:"firstname"`
}

func main() {
	// Create a client
	host := "http://192.168.15.21:9200"
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags) // 添加日志
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		// Handle error
		panic(err)
	}

	// 1.查询数据
	q := elastic.NewMatchQuery("address", "street")
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	fmt.Printf("搜索结果数量：%d/n", total)
	for _, value := range result.Hits.Hits {
		// 将字符串转成struct对象
		account := Account{}
		_ = json.Unmarshal(value.Source, &account)
		fmt.Println(account)

		//if jsonData, err := value.Source.MarshalJSON(); err == nil {
		//	fmt.Println(string(jsonData))
		//}
	}

	// 2.添加数据
	acc := Account{AccountNumber: 14268, FirstName: "xxm"}
	put1, err := client.Index().
		Index("account").
		BodyJson(acc).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed account %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	// Use the client
	_ = client

	// 3.创建索引
	createIndex, err := client.CreateIndex("goods").BodyString(EsGoods{}.GetMapping()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !createIndex.Acknowledged {
		fmt.Println("索引已经存在")
	}
}
