package engine

import (
	"go-crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	// 把in放到SimpleScheduler里
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 往Scheduler里发任务
	for _, r := range seeds {
		// 1. 提交r到SimpleScheduler中的channel中(也就是in中)
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		// 4. 从out中接收result
		result := <-out
		for _, item := range result.Items {
			itemCount++
			log.Printf("Got item: #%d: %v", itemCount, item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// 2. 从in中收到request任务
			request := <-in
			// 收到任务后，worker开始工作
			result, err := worker(request)
			if err != nil {
				continue
			}
			// 3. result送到out中
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
