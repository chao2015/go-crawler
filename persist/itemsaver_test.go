package persist

import (
	"crawler/model"
	"testing"

	"context"

	"encoding/json"

	"crawler/engine"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	//expected := model.Profile{
	//	Name:          "芜湖小啊妹",
	//	Gender:        "女",
	//	Age:           30,
	//	Height:        163,
	//	Weight:        0,
	//	Income:        "5001-8000元",
	//	Marriage:      "离异",
	//	Education:     "高中及以下",
	//	Occupation:    "销售专员",
	//	NativePlace:   "安徽芜湖",
	//	Constellation: "狮子座",
	//	House:         "和家人同住",
	//	Car:           "未购车",
	//}

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1552811555",
		Type: "zhenai",
		Id:   "1552811555",
		Payload: model.Profile{
			Name:          "芜湖小啊妹",
			Gender:        "女",
			Age:           30,
			Height:        163,
			Weight:        0,
			Income:        "5001-8000元",
			Marriage:      "离异",
			Education:     "高中及以下",
			Occupation:    "销售专员",
			NativePlace:   "安徽芜湖",
			Constellation: "狮子座",
			House:         "和家人同住",
			Car:           "未购车",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	// 解决从数据库读取的json字符串转为Item中的Payload interface{}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
