package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"

	"crawler/persist"
)

func main() {
	// 1. SimpleEngine
	// 网络利用率：50-70k/s
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 2. ConcurrentEngine
	// 网络利用率：WorkerCount=10，500-700k/s；WorkerCount=100，3.2M/s
	// 设置了time.Tick流量限制后，大概1.4M/s
	// 一共获取了18274个item，一共(18274-470城市)/2=8902个用户，每个城市首页8902/470=18.94个用户
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 2.1 ConcurrentEngine
	// 2.0版本每个worker用一个goroutine；2.1版本每个worker共用一个队列调度
	// 效率差不多，但是可以方便对调度进行控制，实现负载均衡
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 2.2 针对某个城市测试
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})

	// 2.2 用户去重
	// 2.3 存储items
	// 2.4 代码重构
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	// version 2.5 前端显示：crawler/frontend/starter.go
}
