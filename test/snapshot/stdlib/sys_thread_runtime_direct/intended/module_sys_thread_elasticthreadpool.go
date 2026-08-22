package main

import "snapshot/hxrt"

type I_sys__thread__ElasticThreadPool interface {
	get_isShutdown() bool
	get_threadsCount() int
	run(task func())
	shutdown()
	startWorkerLocked()
	workerResolveWait(worker *sys__thread__ElasticThreadPoolWorker, woke bool) bool
	retireWorkerLocked(worker *sys__thread__ElasticThreadPoolWorker)
	workerTaskFinished(worker *sys__thread__ElasticThreadPoolWorker)
	workerTaskFailed(worker *sys__thread__ElasticThreadPoolWorker)
}

type sys__thread__ElasticThreadPool struct {
	__hx_this       I_sys__thread__ElasticThreadPool
	threadsCount    int
	maxThreadsCount int
	isShutdown      bool
	_isShutdown     bool
	liveWorkers     int
	pendingTasks    int
	pool            *hxrt.Array
	queue           *sys__thread__Deque
	available       *sys__thread__Lock
	mutex           *sys__thread__Mutex
	threadTimeout   float64
}

func New_sys__thread__ElasticThreadPool(maxThreadsCount int, threadTimeout float64) *sys__thread__ElasticThreadPool {
	self := &sys__thread__ElasticThreadPool{}
	self.__hx_this = self
	self.mutex = New_sys__thread__Mutex()
	self.available = New_sys__thread__Lock()
	self.queue = New_sys__thread__Deque()
	self.pool = hxrt.NewArray()
	self.pendingTasks = 0
	self.liveWorkers = 0
	self._isShutdown = false
	if maxThreadsCount < 1 {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("ElasticThreadPool needs maxThreadsCount to be at least 1."), nil, nil))
	}
	self.maxThreadsCount = maxThreadsCount
	self.threadTimeout = threadTimeout
	return self
}

func (self *sys__thread__ElasticThreadPool) get_isShutdown() bool {
	self.mutex.__hx_this.acquire()
	result := self._isShutdown
	self.mutex.__hx_this.release()
	return result
}

func (self *sys__thread__ElasticThreadPool) get_threadsCount() int {
	self.mutex.__hx_this.acquire()
	result := self.liveWorkers
	self.mutex.__hx_this.release()
	return result
}

func (self *sys__thread__ElasticThreadPool) run(task func()) {
	self.mutex.__hx_this.acquire()
	if self._isShutdown {
		self.mutex.__hx_this.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task is rejected. Thread pool is shut down."), nil, nil))
	}
	if task == nil {
		self.mutex.__hx_this.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task to run must not be null."), nil, nil))
	}
	self.pendingTasks = int(int32((self.pendingTasks + 1)))
	self.queue.__hx_this.add(task)
	self.available.__hx_this.release()
	if (self.pendingTasks > self.liveWorkers) && (self.liveWorkers < self.maxThreadsCount) {
		self.__hx_this.startWorkerLocked()
	}
	self.mutex.__hx_this.release()
}

func (self *sys__thread__ElasticThreadPool) shutdown() {
	self.mutex.__hx_this.acquire()
	if self._isShutdown {
		self.mutex.__hx_this.release()
		return
	}
	self._isShutdown = true
	_g := 0
	_g1 := self.liveWorkers
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_1
		_ = hx_tmp
		self.available.__hx_this.release()
	}
	self.mutex.__hx_this.release()
}

func (self *sys__thread__ElasticThreadPool) startWorkerLocked() {
	var selected *sys__thread__ElasticThreadPoolWorker = nil
	_g := 0
	_g1 := self.pool
	for _g < _g1.Len() {
		worker := func(hx_value_2 any) *sys__thread__ElasticThreadPoolWorker {
			if hx_value_2 == nil {
				var hx_zero_3 *sys__thread__ElasticThreadPoolWorker
				return hx_zero_3
			}
			return hx_value_2.(*sys__thread__ElasticThreadPoolWorker)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if worker.dead {
			selected = worker
			break
		}
	}
	if selected == nil {
		selected = New_sys__thread__ElasticThreadPoolWorker(self, self.available, self.threadTimeout)
		hx_arr_4 := self.pool
		hx_arr_4.Push(selected)
	}
	self.liveWorkers = int(int32((self.liveWorkers + 1)))
	selected.__hx_this.start()
}

func (self *sys__thread__ElasticThreadPool) workerResolveWait(worker *sys__thread__ElasticThreadPoolWorker, woke bool) bool {
	self.mutex.__hx_this.acquire()
	hasToken := woke
	if !hasToken {
		hasToken = self.available.__hx_this.wait(0)
	}
	if hasToken {
		nextTask := func(hx_value_5 any) func() {
			if hx_value_5 == nil {
				var hx_zero_6 func()
				return hx_zero_6
			}
			return hx_value_5.(func())
		}(self.queue.__hx_this.pop(false))
		if nextTask != nil {
			worker.task = nextTask
			self.mutex.__hx_this.release()
			return true
		}
		if !self._isShutdown {
			self.mutex.__hx_this.release()
			return true
		}
	}
	self.__hx_this.retireWorkerLocked(worker)
	self.mutex.__hx_this.release()
	return false
}

func (self *sys__thread__ElasticThreadPool) retireWorkerLocked(worker *sys__thread__ElasticThreadPoolWorker) {
	if !worker.dead {
		worker.dead = true
		worker.task = nil
		self.liveWorkers = int(int32((self.liveWorkers - 1)))
	}
}

func (self *sys__thread__ElasticThreadPool) workerTaskFinished(worker *sys__thread__ElasticThreadPoolWorker) {
	self.mutex.__hx_this.acquire()
	worker.task = nil
	self.pendingTasks = int(int32((self.pendingTasks - 1)))
	self.mutex.__hx_this.release()
}

func (self *sys__thread__ElasticThreadPool) workerTaskFailed(worker *sys__thread__ElasticThreadPoolWorker) {
	self.mutex.__hx_this.acquire()
	worker.task = nil
	self.pendingTasks = int(int32((self.pendingTasks - 1)))
	self.__hx_this.retireWorkerLocked(worker)
	if (self.pendingTasks > self.liveWorkers) && (self.liveWorkers < self.maxThreadsCount) {
		self.__hx_this.startWorkerLocked()
	}
	self.mutex.__hx_this.release()
}
