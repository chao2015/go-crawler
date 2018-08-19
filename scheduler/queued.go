package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 装chan的chan
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				// worker和request同时有闲置的时，才给active...赋值，否则均为初始值nil
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			// 独立的两件事用select分别处理，它们发生的前后不定
			select {
			case r := <-s.requestChan:
				// 有request来到request队列，则加入
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				// 有worker来到worker队列，则加入
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				// 当activeWorker是初始值nil，则不会执行该select case
				// 当有worker，执行该case，则把已经使用的worker和request从队列里拿掉
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
