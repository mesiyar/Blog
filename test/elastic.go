package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200"
var indexName = "test_go_els"
var ctx = context.Background()

func init() {
	// 创建ES client用于后续操作ES
	var err error
	client, err = elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		// Handle error
		log.Fatal("连接失败: %v\n", err)
	} else {
		log.Println("连接成功")
	}

}

type user struct {
	Username string
	Age      int
	Profile  string
}

func createUser() {
	user := user{
		Username: "张三",
		Age:      10,
		Profile:  "我就是我 颜色不一样的花火",
	}
	pst, err := client.Index().Index(indexName).BodyJson(user).Do(ctx)
	if err != nil {
		log.Fatal("post data error ", err)
	}
	log.Println("post success", pst)
}

func main() {
	//createUser()
	getUses()
}

func getUses() {
	terquery := elastic.NewTermQuery("username", "eddie")
	result, err := client.Search().Index(indexName).Query(terquery).Pretty(true).Do(ctx)
	if err != nil {
		log.Fatal("get data error ", err)
	}
	log.Println(result.TookInMillis)
	log.Println(result.Hits.TotalHits.Value)
	if result.Hits != nil {
		log.Printf("total hits %d \n", result.Hits.TotalHits.Value)
		// 解析数据
		for _, hit := range result.Hits.Hits {
			var u user
			if err := json.Unmarshal(hit.Source, &u); err != nil {
				log.Fatal("unmarshal data error", err)
			}
			log.Printf("Id %s data is %+v", hit.Id,u)

		}
	}
}
