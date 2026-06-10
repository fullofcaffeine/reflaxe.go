package main

import "snapshot/hxrt"

type I_sys__thread__ElasticThreadPoolWorker interface {
	wakeup(task func())
	shutdown()
	start()
	loop()
}

type sys__thread__ElasticThreadPoolWorker struct {
	__hx_this  I_sys__thread__ElasticThreadPoolWorker
	task       func()
	dead       bool
	deathMutex *sys__thread__Mutex
	waiter     *sys__thread__Lock
	queue      *sys__thread__Deque
	timeout    float64
	isShutdown bool
}

func New_sys__thread__ElasticThreadPoolWorker(queue *sys__thread__Deque, timeout float64) *sys__thread__ElasticThreadPoolWorker {
	self := &sys__thread__ElasticThreadPoolWorker{}
	self.__hx_this = self
	self.isShutdown = false
	self.waiter = New_sys__thread__Lock()
	self.deathMutex = New_sys__thread__Mutex()
	self.dead = false
	self.queue = queue
	self.timeout = timeout
	self.start()
	return self
}

func (self *sys__thread__ElasticThreadPoolWorker) wakeup(task func()) {
	self.deathMutex.acquire()
	if self.dead {
		self.start()
	}
	self.task = task
	self.waiter.release()
	self.deathMutex.release()
}

func (self *sys__thread__ElasticThreadPoolWorker) shutdown() {
	self.isShutdown = true
	self.waiter.release()
}

func (self *sys__thread__ElasticThreadPoolWorker) start() {
	self.dead = false
	sys__thread__Thread_create(self.loop)
}

func (self *sys__thread__ElasticThreadPoolWorker) loop() {
	hxrt.TryCatch(func() {
		for self.waiter.wait(self.timeout) {
			_g := self.task
			if _g == nil {
				if self.isShutdown {
					break
				}
			} else {
				fn := _g
				fn()
				for true {
					_g_1 := func(hx_value_42 any) func() {
						if hx_value_42 == nil {
							var hx_zero_43 func()
							return hx_zero_43
						}
						return hx_value_42.(func())
					}(self.queue.pop(false))
					if _g_1 == nil {
						break
					} else {
						queued := _g_1
						queued()
					}
				}
				self.task = nil
			}
		}
		self.deathMutex.acquire()
		if self.task != nil {
			self.start()
		} else {
			self.dead = true
		}
		self.deathMutex.release()
	}, func(hx_caught_40 any) {
		err := hx_caught_40
		self.task = nil
		self.start()
		hxrt.Throw(err)
	})
}
