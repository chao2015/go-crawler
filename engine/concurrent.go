package engine

import (
	"crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// 2.1版修正内容：把统一使用的in chan，替换为Scheduler在worker和request队列里进行调度，每次调度有对应的channel
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// simple是所有worker共用channel，queued是每个worker一个channel，到底怎么用问Scheduler.WorkerChan()
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 往Scheduler里发任务
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		// 1. 提交r到SimpleScheduler中的channel中(也就是in中)
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		// 4. 从out中接收result
		result := <-out
		for _, item := range result.Items {
			//itemCount++
			//log.Printf("Got item: #%d: %v", itemCount, item)

			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// 整个Scheduler作为参数太重了，把要用到的WorkerReady()分离出来
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
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
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
