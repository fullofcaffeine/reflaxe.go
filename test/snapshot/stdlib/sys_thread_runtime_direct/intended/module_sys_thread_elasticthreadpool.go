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
	self.mutex.acquire()
	result := self._isShutdown
	self.mutex.release()
	return result
}

func (self *sys__thread__ElasticThreadPool) get_threadsCount() int {
	self.mutex.acquire()
	result := self.liveWorkers
	self.mutex.release()
	return result
}

func (self *sys__thread__ElasticThreadPool) run(task func()) {
	self.mutex.acquire()
	if self._isShutdown {
		self.mutex.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task is rejected. Thread pool is shut down."), nil, nil))
	}
	if task == nil {
		self.mutex.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task to run must not be null."), nil, nil))
	}
	self.pendingTasks = int(int32((self.pendingTasks + 1)))
	self.queue.add(task)
	self.available.release()
	if (self.pendingTasks > self.liveWorkers) && (self.liveWorkers < self.maxThreadsCount) {
		self.startWorkerLocked()
	}
	self.mutex.release()
}

func (self *sys__thread__ElasticThreadPool) shutdown() {
	self.mutex.acquire()
	if self._isShutdown {
		self.mutex.release()
		return
	}
	self._isShutdown = true
	_g := 0
	_g1 := self.liveWorkers
	for _g < _g1 {
		hx_post_33 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_33
		_ = hx_tmp
		self.available.release()
	}
	self.mutex.release()
}

func (self *sys__thread__ElasticThreadPool) startWorkerLocked() {
	var selected *sys__thread__ElasticThreadPoolWorker = nil
	_g := 0
	_g1 := self.pool
	for _g < _g1.Len() {
		worker := func(hx_value_34 any) *sys__thread__ElasticThreadPoolWorker {
			if hx_value_34 == nil {
				var hx_zero_35 *sys__thread__ElasticThreadPoolWorker
				return hx_zero_35
			}
			return hx_value_34.(*sys__thread__ElasticThreadPoolWorker)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if worker.dead {
			selected = worker
			break
		}
	}
	if selected == nil {
		selected = New_sys__thread__ElasticThreadPoolWorker(self, self.available, self.threadTimeout)
		hx_arr_36 := self.pool
		hx_arr_36.Push(selected)
	}
	self.liveWorkers = int(int32((self.liveWorkers + 1)))
	selected.start()
}

func (self *sys__thread__ElasticThreadPool) workerResolveWait(worker *sys__thread__ElasticThreadPoolWorker, woke bool) bool {
	self.mutex.acquire()
	hasToken := woke
	if !hasToken {
		hasToken = self.available.wait(0)
	}
	if hasToken {
		nextTask := func(hx_value_37 any) func() {
			if hx_value_37 == nil {
				var hx_zero_38 func()
				return hx_zero_38
			}
			return hx_value_37.(func())
		}(self.queue.pop(false))
		if nextTask != nil {
			worker.task = nextTask
			self.mutex.release()
			return true
		}
		if !self._isShutdown {
			self.mutex.release()
			return true
		}
	}
	self.retireWorkerLocked(worker)
	self.mutex.release()
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
	self.mutex.acquire()
	worker.task = nil
	self.pendingTasks = int(int32((self.pendingTasks - 1)))
	self.mutex.release()
}

func (self *sys__thread__ElasticThreadPool) workerTaskFailed(worker *sys__thread__ElasticThreadPoolWorker) {
	self.mutex.acquire()
	worker.task = nil
	self.pendingTasks = int(int32((self.pendingTasks - 1)))
	self.retireWorkerLocked(worker)
	if (self.pendingTasks > self.liveWorkers) && (self.liveWorkers < self.maxThreadsCount) {
		self.startWorkerLocked()
	}
	self.mutex.release()
}
