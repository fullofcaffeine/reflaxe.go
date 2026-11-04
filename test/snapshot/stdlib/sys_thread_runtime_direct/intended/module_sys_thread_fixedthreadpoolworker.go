package main

import "snapshot/hxrt"

type I_sys__thread__FixedThreadPoolWorker interface {
	loop()
}

type sys__thread__FixedThreadPoolWorker struct {
	__hx_this I_sys__thread__FixedThreadPoolWorker
	queue     *sys__thread__Deque
}

func (self *sys__thread__FixedThreadPoolWorker) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__FixedThreadPoolWorker(queue *sys__thread__Deque) *sys__thread__FixedThreadPoolWorker {
	return hx_init_sys__thread__FixedThreadPoolWorker(hx_alloc_sys__thread__FixedThreadPoolWorker(nil), queue)
}

func hx_alloc_sys__thread__FixedThreadPoolWorker(owner any) *sys__thread__FixedThreadPoolWorker {
	self := &sys__thread__FixedThreadPoolWorker{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__FixedThreadPoolWorker)
	}
	return self
}

func hx_init_sys__thread__FixedThreadPoolWorker(self *sys__thread__FixedThreadPoolWorker, queue *sys__thread__Deque) *sys__thread__FixedThreadPoolWorker {
	self.queue = queue
	sys__thread__Thread_create(func() *hxrt.Function[func()] {
		hx_receiver_1 := self
		return hxrt.BoundFunction(func() {
			hx_receiver_1.__hx_this.loop()
		}, func() any {
			return hx_receiver_1.__hx_this
		}, "loop")
	}())
	return self
}

func (self *sys__thread__FixedThreadPoolWorker) loop() {
	hx_try_return_2 := false
	hxrt.TryCatch(func() {
		for true {
			task := hx_recover_callable_1(self.queue.__hx_this.pop(true))
			if task != nil {
				func() {
					var hx_callable_7 *hxrt.Function[func()] = task
					hxrt.FunctionCall(hx_callable_7)()
				}()
			}
		}
	}, func(hx_caught_3 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_3)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_3)
		if hx_is_nominal_2(hx_exception_5) {
			var hx_tmp *sys__thread__FixedThreadPoolShutdownException = hx_recover_class_2(hx_exception_5)
			_ = hx_tmp
			hx_try_return_2 = true
			return
		} else {
			var err any = hx_payload_6
			sys__thread__Thread_create(func() *hxrt.Function[func()] {
				hx_receiver_4 := self
				return hxrt.BoundFunction(func() {
					hx_receiver_4.__hx_this.loop()
				}, func() any {
					return hx_receiver_4.__hx_this
				}, "loop")
			}())
			hxrt.Throw(haxe__Exception_thrown(err))
		}
	})
	if hx_try_return_2 {
		return
	}
}
