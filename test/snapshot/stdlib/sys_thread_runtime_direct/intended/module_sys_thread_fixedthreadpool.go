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
}

func New_sys__thread__FixedThreadPool(threadsCount int) *sys__thread__FixedThreadPool {
	self := &sys__thread__FixedThreadPool{}
	self.__hx_this = self
	self.queue = New_sys__thread__Deque()
	self._isShutdown = false
	if threadsCount < 1 {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("FixedThreadPool needs threadsCount to be at least 1."), nil, nil))
	}
	workers := []*sys__thread__FixedThreadPoolWorker{}
	_g := 0
	_g1 := threadsCount
	for _g < _g1 {
		hx_post_32 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_32
		_ = hx_tmp
		hx_arr_33 := workers
		hx_arr_33 = append(hx_arr_33, New_sys__thread__FixedThreadPoolWorker(self.queue))
		workers = hx_arr_33
	}
	self.pool = workers
	return self
}

func (self *sys__thread__FixedThreadPool) get_threadsCount() int {
	return len(self.pool)
}

func (self *sys__thread__FixedThreadPool) get_isShutdown() bool {
	return self._isShutdown
}

func (self *sys__thread__FixedThreadPool) run(task func()) {
	if self._isShutdown {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task is rejected. Thread pool is shut down."), nil, nil))
	}
	if task == nil {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task to run must not be null."), nil, nil))
	}
	self.queue.add(task)
}

func (self *sys__thread__FixedThreadPool) shutdown() {
	if self._isShutdown {
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
}

func sys__thread__FixedThreadPool_shutdownTask() {
	hxrt.Throw(New_sys__thread__FixedThreadPoolShutdownException(hxrt.StringFromLiteral(""), nil, nil))
}
