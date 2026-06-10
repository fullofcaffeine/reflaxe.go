package main

import "snapshot/hxrt"

type I_sys__thread__ElasticThreadPool interface {
	get_isShutdown() bool
	run(task func())
	shutdown()
	get_threadsCount() int
}

type sys__thread__ElasticThreadPool struct {
	__hx_this       I_sys__thread__ElasticThreadPool
	threadsCount    int
	maxThreadsCount int
	isShutdown      bool
	_isShutdown     bool
	pool            []*sys__thread__ElasticThreadPoolWorker
	queue           *sys__thread__Deque
	mutex           *sys__thread__Mutex
	threadTimeout   float64
}

func New_sys__thread__ElasticThreadPool(maxThreadsCount int, threadTimeout float64) *sys__thread__ElasticThreadPool {
	self := &sys__thread__ElasticThreadPool{}
	self.__hx_this = self
	self.mutex = New_sys__thread__Mutex()
	self.queue = New_sys__thread__Deque()
	self.pool = []*sys__thread__ElasticThreadPoolWorker{}
	self._isShutdown = false
	if maxThreadsCount < 1 {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("ElasticThreadPool needs maxThreadsCount to be at least 1."), nil, nil))
	}
	self.maxThreadsCount = maxThreadsCount
	self.threadTimeout = threadTimeout
	return self
}

func (self *sys__thread__ElasticThreadPool) get_isShutdown() bool {
	return self._isShutdown
}

func (self *sys__thread__ElasticThreadPool) run(task func()) {
	if self._isShutdown {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task is rejected. Thread pool is shut down."), nil, nil))
	}
	if task == nil {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task to run must not be null."), nil, nil))
	}
	self.mutex.acquire()
	submitted := false
	var deadWorker *sys__thread__ElasticThreadPoolWorker = nil
	_g := 0
	_g1 := self.pool
	for _g < len(_g1) {
		worker := _g1[_g]
		_g = int(int32((_g + 1)))
		if (deadWorker == nil) && worker.dead {
			deadWorker = worker
		}
		if worker.task == nil {
			submitted = true
			worker.wakeup(task)
			break
		}
	}
	if !submitted {
		if deadWorker != nil {
			deadWorker.wakeup(task)
		} else {
			if len(self.pool) < self.maxThreadsCount {
				worker_1 := New_sys__thread__ElasticThreadPoolWorker(self.queue, self.threadTimeout)
				hx_arr_44 := self.pool
				hx_arr_44 = append(hx_arr_44, worker_1)
				self.pool = hx_arr_44
				worker_1.wakeup(task)
			} else {
				self.queue.add(task)
			}
		}
	}
	self.mutex.release()
}

func (self *sys__thread__ElasticThreadPool) shutdown() {
	if self._isShutdown {
		return
	}
	self.mutex.acquire()
	self._isShutdown = true
	_g := 0
	_g1 := self.pool
	for _g < len(_g1) {
		worker := _g1[_g]
		_g = int(int32((_g + 1)))
		worker.shutdown()
	}
	self.mutex.release()
}

func (self *sys__thread__ElasticThreadPool) get_threadsCount() int {
	result := 0
	_g := 0
	_g1 := self.pool
	for _g < len(_g1) {
		worker := _g1[_g]
		_g = int(int32((_g + 1)))
		if !worker.dead {
			result = int(int32((result + 1)))
		}
	}
	return result
}
