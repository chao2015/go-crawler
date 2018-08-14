package engine

import (
	"go-crawler/fetcher"
	"log"
)

// input: Request
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error "+"fetching url %s: %v",
				r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...) // ...就是把slice的内容展开加进去

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
