package parser

import (
	"go-crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "芜湖小啊妹")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
