package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Account struct {
	AccountNumber int32  `json:"account_number"`
	FirstName     string `json:"firstname"`
}

func main() {
	// Create a client
	host := "http://192.168.15.21:9200"
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
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
}
