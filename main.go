package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	// 1. SimpleEngine
	// 网络利用率：50-70k/s
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 2. ConcurrentEngine
	// 网络利用率：WorkerCount=10，500-700k/s；WorkerCount=100，3M/s
	// 设置了time.Tick流量限制后，大概1.5M/s
	// 一共获取了18274个item，一共(18274-470城市)/2=8902个用户，每个城市首页8902/470=18.94个用户
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
