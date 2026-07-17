package main

import "snapshot/hxrt"

type I_sys__thread__FixedThreadPoolWorker interface {
	loop()
}

type sys__thread__FixedThreadPoolWorker struct {
	__hx_this I_sys__thread__FixedThreadPoolWorker
	queue     *sys__thread__Deque
}

func New_sys__thread__FixedThreadPoolWorker(queue *sys__thread__Deque) *sys__thread__FixedThreadPoolWorker {
	self := &sys__thread__FixedThreadPoolWorker{}
	self.__hx_this = self
	self.queue = queue
	sys__thread__Thread_create(self.loop)
	return self
}

func (self *sys__thread__FixedThreadPoolWorker) loop() {
	hx_try_return_27 := false
	hxrt.TryCatch(func() {
		for true {
			task := func(hx_value_30 any) func() {
				if hx_value_30 == nil {
					var hx_zero_31 func()
					return hx_zero_31
				}
				return hx_value_30.(func())
			}(self.queue.pop(true))
			if task != nil {
				task()
			}
		}
	}, func(hx_caught_28 any) {
		switch hx_typed_29 := hx_caught_28.(type) {
		case *sys__thread__FixedThreadPoolShutdownException:
			hx_tmp := hx_typed_29
			_ = hx_tmp
			hx_try_return_27 = true
			return
		default:
			err := hx_caught_28
			sys__thread__Thread_create(self.loop)
			hxrt.Throw(err)
		}
	})
	if hx_try_return_27 {
		return
	}
}
