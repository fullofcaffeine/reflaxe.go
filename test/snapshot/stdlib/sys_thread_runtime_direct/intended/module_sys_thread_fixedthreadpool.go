package main

import "snapshot/hxrt"

type I_sys__thread__FixedThreadPool interface {
	get_threadsCount() int
	get_isShutdown() bool
	run(task func())
	shutdown()
}

type sys__thread__FixedThreadPool struct {
	__hx_this    I_sys__thread__FixedThreadPool
	threadsCount int
	isShutdown   bool
	_isShutdown  bool
	pool         []*sys__thread__FixedThreadPoolWorker
	queue        *sys__thread__Deque
	mutex        *sys__thread__Mutex
}

func New_sys__thread__FixedThreadPool(threadsCount int) *sys__thread__FixedThreadPool {
	self := &sys__thread__FixedThreadPool{}
	self.__hx_this = self
	self.mutex = New_sys__thread__Mutex()
	self.queue = New_sys__thread__Deque()
	self._isShutdown = false
	if threadsCount < 1 {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("FixedThreadPool needs threadsCount to be at least 1."), nil, nil))
	}
	workers := []*sys__thread__FixedThreadPoolWorker{}
	_g := 0
	_g1 := threadsCount
	for _g < _g1 {
		hx_post_33 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_33
		_ = hx_tmp
		workers = append(workers, New_sys__thread__FixedThreadPoolWorker(self.queue))
	}
	self.pool = workers
	return self
}

func (self *sys__thread__FixedThreadPool) get_threadsCount() int {
	return len(self.pool)
}

func (self *sys__thread__FixedThreadPool) get_isShutdown() bool {
	self.mutex.acquire()
	result := self._isShutdown
	self.mutex.release()
	return result
}

func (self *sys__thread__FixedThreadPool) run(task func()) {
	self.mutex.acquire()
	if self._isShutdown {
		self.mutex.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task is rejected. Thread pool is shut down."), nil, nil))
	}
	if task == nil {
		self.mutex.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task to run must not be null."), nil, nil))
	}
	self.queue.add(task)
	self.mutex.release()
}

func (self *sys__thread__FixedThreadPool) shutdown() {
	self.mutex.acquire()
	if self._isShutdown {
		self.mutex.release()
		return
	}
	self._isShutdown = true
	_g := 0
	_g1 := self.pool
	for _g < len(_g1) {
		hx_tmp := _g1[_g]
		_ = hx_tmp
		_g = int(int32((_g + 1)))
		self.queue.add(sys__thread__FixedThreadPool_shutdownTask)
	}
	self.mutex.release()
}

func sys__thread__FixedThreadPool_shutdownTask() {
	hxrt.Throw(New_sys__thread__FixedThreadPoolShutdownException(hxrt.StringFromLiteral(""), nil, nil))
}
