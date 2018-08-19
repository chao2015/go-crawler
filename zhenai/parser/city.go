package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

// 城市解析器
// 得到各个城市首页上的用户名称和用户详情页url
func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				// ParserFunc: ParseProfile
				// 为了满足type Request struct中的定义，改用匿名函数：
				ParserFunc: ProfileParser(string(m[2])),
			})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		// /qishi之类的网页都已失效
		if strings.Contains(string(m[1]), "qishi") {
			continue
		}
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
