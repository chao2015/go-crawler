package engine

import (
	"log"
)

type SimpleEngine struct{}

// input: Request
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := e.worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...) // ...就是把slice的内容展开加进去

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
