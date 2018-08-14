package parser

import (
	"go-crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 城市列表解析器
// input: utf-8编码的文本
// output: Request{URL, 对应Parser}列表，Item列表。见type ParseResult struct{}
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				// Url对应采用ParseCity解析器来分析
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
