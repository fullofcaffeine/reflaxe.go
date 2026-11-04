package main

import "snapshot/hxrt"

type I_sys__thread__ElasticThreadPoolWorker interface {
	start()
	loop()
}

type sys__thread__ElasticThreadPoolWorker struct {
	__hx_this I_sys__thread__ElasticThreadPoolWorker
	task      *hxrt.Function[func()]
	dead      bool
	owner     *sys__thread__ElasticThreadPool
	available *sys__thread__Lock
	timeout   float64
}

func (self *sys__thread__ElasticThreadPoolWorker) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__ElasticThreadPoolWorker(owner *sys__thread__ElasticThreadPool, available *sys__thread__Lock, timeout float64) *sys__thread__ElasticThreadPoolWorker {
	return hx_init_sys__thread__ElasticThreadPoolWorker(hx_alloc_sys__thread__ElasticThreadPoolWorker(nil), owner, available, timeout)
}

func hx_alloc_sys__thread__ElasticThreadPoolWorker(owner any) *sys__thread__ElasticThreadPoolWorker {
	self := &sys__thread__ElasticThreadPoolWorker{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__ElasticThreadPoolWorker)
	}
	return self
}

func hx_init_sys__thread__ElasticThreadPoolWorker(self *sys__thread__ElasticThreadPoolWorker, owner *sys__thread__ElasticThreadPool, available *sys__thread__Lock, timeout float64) *sys__thread__ElasticThreadPoolWorker {
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
	sys__thread__Thread_create(func() *hxrt.Function[func()] {
		hx_receiver_1 := self
		return hxrt.BoundFunction(func() {
			hx_receiver_1.__hx_this.loop()
		}, func() any {
			return hx_receiver_1.__hx_this
		}, "loop")
	}())
}

func (self *sys__thread__ElasticThreadPoolWorker) loop() {
	for true {
		woke := self.available.__hx_this.wait(self.timeout)
		if !self.owner.__hx_this.workerResolveWait(self, woke) {
			return
		}
		_g := self.task
		if _g == nil {
			continue
		} else {
			fn := _g
			hxrt.TryCatch(func() {
				func() {
					var hx_callable_5 *hxrt.Function[func()] = fn
					hxrt.FunctionCall(hx_callable_5)()
				}()
			}, func(hx_caught_2 any) {
				hx_exception_3 := haxe__Exception_caught(hx_caught_2)
				hx_payload_4 := haxe__Exception_catchValue(hx_caught_2)
				_ = hx_exception_3
				var err any = hx_payload_4
				self.owner.__hx_this.workerTaskFailed(self)
				hxrt.Throw(haxe__Exception_thrown(err))
			})
			self.owner.__hx_this.workerTaskFinished(self)
		}
	}
}
