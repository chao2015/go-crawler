package persist

import (
	"log"

	"context"

	"crawler/engine"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Item Saver: got item "+"#%d: %v",
				itemCount, item)

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error"+"saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	// 1. 安装docker
	// 2. 安装elasticsearch: docker run -d -p 9200:9200 elasticsearch
	// 3. 安装elastic client: go get -v gopkg.in/olivere/elastic.v5

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	// index:数据库名 type:表名
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	//fmt.Printf("%+v", resp)

	return nil
}
