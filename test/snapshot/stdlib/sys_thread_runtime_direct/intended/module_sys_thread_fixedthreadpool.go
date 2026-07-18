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
	pool         *hxrt.Array
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
	workers := hxrt.NewArray()
	_g := 0
	_g1 := threadsCount
	for _g < _g1 {
		hx_post_23 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_23
		_ = hx_tmp
		workers.Push(New_sys__thread__FixedThreadPoolWorker(self.queue))
	}
	self.pool = workers
	return self
}

func (self *sys__thread__FixedThreadPool) get_threadsCount() int {
	return self.pool.Len()
}

func (self *sys__thread__FixedThreadPool) get_isShutdown() bool {
	self.mutex.__hx_this.acquire()
	result := self._isShutdown
	self.mutex.__hx_this.release()
	return result
}

func (self *sys__thread__FixedThreadPool) run(task func()) {
	self.mutex.__hx_this.acquire()
	if self._isShutdown {
		self.mutex.__hx_this.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task is rejected. Thread pool is shut down."), nil, nil))
	}
	if task == nil {
		self.mutex.__hx_this.release()
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("Task to run must not be null."), nil, nil))
	}
	self.queue.__hx_this.add(task)
	self.mutex.__hx_this.release()
}

func (self *sys__thread__FixedThreadPool) shutdown() {
	self.mutex.__hx_this.acquire()
	if self._isShutdown {
		self.mutex.__hx_this.release()
		return
	}
	self._isShutdown = true
	_g := 0
	_g1 := self.pool
	for _g < _g1.Len() {
		hx_tmp := func(hx_value_25 any) *sys__thread__FixedThreadPoolWorker {
			if hx_value_25 == nil {
				var hx_zero_26 *sys__thread__FixedThreadPoolWorker
				return hx_zero_26
			}
			return hx_value_25.(*sys__thread__FixedThreadPoolWorker)
		}(_g1.Get(_g))
		_ = hx_tmp
		_g = int(int32((_g + 1)))
		self.queue.__hx_this.add(sys__thread__FixedThreadPool_shutdownTask)
	}
	self.mutex.__hx_this.release()
}

func sys__thread__FixedThreadPool_shutdownTask() {
	hxrt.Throw(New_sys__thread__FixedThreadPoolShutdownException(hxrt.StringFromLiteral(""), nil, nil))
}
