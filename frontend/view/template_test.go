package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template.Must(template.ParseFiles("template.html"))

	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1552811555",
		Type: "zhenai",
		Id:   "1552811555",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	//err = template.Execute(os.Stdout, page)
	err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}
}
