package view

import (
	"os"
	"testing"

	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

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

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}

	// TODO: verify contents in template.test.html
}
