package main

import "snapshot/hxrt"

type I_sys__thread__ElasticThreadPoolWorker interface {
	start()
	loop()
}

type sys__thread__ElasticThreadPoolWorker struct {
	__hx_this I_sys__thread__ElasticThreadPoolWorker
	task      func()
	dead      bool
	owner     *sys__thread__ElasticThreadPool
	available *sys__thread__Lock
	timeout   float64
}

func New_sys__thread__ElasticThreadPoolWorker(owner *sys__thread__ElasticThreadPool, available *sys__thread__Lock, timeout float64) *sys__thread__ElasticThreadPoolWorker {
	self := &sys__thread__ElasticThreadPoolWorker{}
	self.__hx_this = self
	self.dead = true
	self.task = nil
	self.owner = owner
	self.available = available
	self.timeout = timeout
	return self
}

func (self *sys__thread__ElasticThreadPoolWorker) start() {
	self.dead = false
	self.task = nil
	sys__thread__Thread_create(self.loop)
}

func (self *sys__thread__ElasticThreadPoolWorker) loop() {
	for true {
		woke := self.available.wait(self.timeout)
		if !self.owner.workerResolveWait(self, woke) {
			return
		}
		_g := self.task
		if _g == nil {
			continue
		} else {
			fn := _g
			hxrt.TryCatch(func() {
				fn()
			}, func(hx_caught_31 any) {
				err := hx_caught_31
				self.owner.workerTaskFailed(self)
				hxrt.Throw(err)
			})
			self.owner.workerTaskFinished(self)
		}
	}
}
