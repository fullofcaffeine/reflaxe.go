package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	lock := New_sys__thread__Lock()
	lock.__hx_this.release()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("lock.release_before_wait="), hxrt.StdString(lock.__hx_this.wait(nil))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("lock.timeout_empty="), hxrt.StdString(lock.__hx_this.wait(0.0))))
	hxrt.Println(v_1)
	mutex := New_sys__thread__Mutex()
	mutex.__hx_this.acquire()
	mutex.__hx_this.acquire()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mutex.try_reentrant="), hxrt.StdString(mutex.__hx_this.tryAcquire())))
	hxrt.Println(v_2)
	mutex.__hx_this.release()
	mutex.__hx_this.release()
	mutex.__hx_this.release()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mutex.try_after_release="), hxrt.StdString(mutex.__hx_this.tryAcquire())))
	hxrt.Println(v_3)
	mutex.__hx_this.release()
	condition := New_sys__thread__Condition()
	condition.__hx_this.acquire()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("condition.try_reentrant="), hxrt.StdString(condition.__hx_this.tryAcquire())))
	hxrt.Println(v_4)
	condition.__hx_this.release()
	condition.__hx_this.release()
	sem := New_sys__thread__Semaphore(1)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sem.try_first="), hxrt.StdString(sem.__hx_this.tryAcquire(nil))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sem.try_empty="), hxrt.StdString(sem.__hx_this.tryAcquire(0.0))))
	hxrt.Println(v_6)
	sem.__hx_this.release()
	sem.__hx_this.acquire()
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("sem.try_after_acquire="), hxrt.StdString(sem.__hx_this.tryAcquire(0.0))))
	hxrt.Println(v_7)
	loop := New_sys__thread__EventLoop()
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("loop.wait_empty="), hxrt.StdString(loop.__hx_this.wait(0.0))))
	hxrt.Println(v_8)
	deque := New_sys__thread__Deque()
	deque.__hx_this.add(hxrt.StringFromLiteral("tail"))
	deque.__hx_this.push(hxrt.StringFromLiteral("head"))
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("deque.pop1="), func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(deque.__hx_this.pop(false))))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("deque.pop2="), func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(deque.__hx_this.pop(false))))
	hxrt.Println(v_10)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("deque.pop3="), hxrt.StdString(func(hx_value_5 any) *string {
		if hx_value_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_value_5.(*string)
	}(deque.__hx_this.pop(false)))))
	hxrt.Println(v_11)
	tls := New_sys__thread__Tls()
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.initial="), hxrt.StdString(func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(tls.__hx_this.get_value()))))
	hxrt.Println(v_12)
	func(hx_value_9 any) *string {
		if hx_value_9 == nil {
			var hx_zero_10 *string
			return hx_zero_10
		}
		return hx_value_9.(*string)
	}(tls.__hx_this.set_value(hxrt.StringFromLiteral("worker")))
	var v_13 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.after_set="), hxrt.StdString(func(hx_value_11 any) *string {
		if hx_value_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_value_11.(*string)
	}(tls.__hx_this.get_value()))))
	hxrt.Println(v_13)
	func(hx_value_13 any) *string {
		if hx_value_13 == nil {
			var hx_zero_14 *string
			return hx_zero_14
		}
		return hx_value_13.(*string)
	}(tls.__hx_this.set_value(nil))
	var v_14 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.after_clear="), hxrt.StdString(func(hx_value_15 any) *string {
		if hx_value_15 == nil {
			var hx_zero_16 *string
			return hx_zero_16
		}
		return hx_value_15.(*string)
	}(tls.__hx_this.get_value()))))
	hxrt.Println(v_14)
	noLoop := New_sys__thread__NoEventLoopException(hxrt.StringFromLiteral("Event loop is not available. Refer to sys.thread.Thread.runWithEventLoop."), nil)
	var v_15 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("noLoop.msg="), hxrt.ExceptionMessage(noLoop)))
	hxrt.Println(v_15)
	pool := New__Main__DummyPool()
	pool.run(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("pool.task=ran")))
	})
	var v_16 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("pool.runs="), pool.runCount()))
	hxrt.Println(v_16)
	pool.shutdown()
	var v_17 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pool.shutdown="), hxrt.StdString(pool.get_isShutdown())))
	hxrt.Println(v_17)
	hxrt.TryCatch(func() {
		pool.run(func() {
		})
	}, func(hx_caught_17 any) {
		switch hx_typed_18 := hx_caught_17.(type) {
		case *sys__thread__ThreadPoolException:
			err := hx_typed_18
			var v_18 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pool.error="), hxrt.ExceptionMessage(err)))
			hxrt.Println(v_18)
		default:
			hxrt.Throw(hx_caught_17)
		}
	})
}

type I__Main__DummyPool interface {
	get_threadsCount() int
	get_isShutdown() bool
	run(task func())
	shutdown()
	runCount() int
}

type _Main__DummyPool struct {
	__hx_this    I__Main__DummyPool
	threadsCount int
	isShutdown   bool
	_isShutdown  bool
	runs         int
}

func New__Main__DummyPool() *_Main__DummyPool {
	self := &_Main__DummyPool{}
	self.__hx_this = self
	self.runs = 0
	self._isShutdown = false
	return self
}

func (self *_Main__DummyPool) get_threadsCount() int {
	return 0
}

func (self *_Main__DummyPool) get_isShutdown() bool {
	return self._isShutdown
}

func (self *_Main__DummyPool) run(task func()) {
	if self._isShutdown {
		hxrt.Throw(New_sys__thread__ThreadPoolException(hxrt.StringFromLiteral("shutdown"), nil, nil))
	}
	self.runs = int(int32((self.runs + 1)))
	task()
}

func (self *_Main__DummyPool) shutdown() {
	self._isShutdown = true
}

func (self *_Main__DummyPool) runCount() int {
	return self.runs
}
