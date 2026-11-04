package main

import "snapshot/hxrt"

type I_haxe__Timer interface {
	stop()
}

type haxe__Timer struct {
	__hx_this    I_haxe__Timer
	thread       *sys__thread__Thread
	eventHandler any
	run          *hxrt.Function[func()]
}

func (self *haxe__Timer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__Timer(time_ms int) *haxe__Timer {
	return hx_init_haxe__Timer(hx_alloc_haxe__Timer(nil), time_ms)
}

func hx_alloc_haxe__Timer(owner any) *haxe__Timer {
	self := &haxe__Timer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__Timer)
	}
	return self
}

func hx_init_haxe__Timer(self *haxe__Timer, time_ms int) *haxe__Timer {
	_gthis := self
	self.run = hxrt.NewFunction(func() {
	})
	self.thread = sys__thread__Thread_current()
	self.eventHandler = self.thread.__hx_this.get_events().__hx_this.repeat(hxrt.NewFunction(func() {
		func() {
			var hx_callable_1 *hxrt.Function[func()] = _gthis.run
			hxrt.FunctionCall(hx_callable_1)()
		}()
	}), time_ms)
	return self
}

func (self *haxe__Timer) stop() {
	if !hxrt.HaxeEqual(self.eventHandler, any(0)) {
		self.thread.__hx_this.get_events().__hx_this.cancel(self.eventHandler)
		self.eventHandler = any(0)
	}
}

func haxe__Timer_delay(f *hxrt.Function[func()], time_ms int) *haxe__Timer {
	timer := New_haxe__Timer(time_ms)
	timer.run = hxrt.NewFunction(func() {
		timer.__hx_this.stop()
		func() {
			var hx_callable_2 *hxrt.Function[func()] = f
			hxrt.FunctionCall(hx_callable_2)()
		}()
	})
	return timer
}

func haxe__Timer_measure(f *hxrt.Function[func() any], pos hxrt.Anonymous) any {
	t0 := hxrt.ThreadNowSeconds()
	var result any = func() any {
		var hx_callable_3 *hxrt.Function[func() any] = f
		return hxrt.FunctionCall(hx_callable_3)()
	}()
	func() {
		var hx_callable_4 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_5 any = hxrt.StringConcatAny((hxrt.ThreadNowSeconds() - t0), hxrt.StringFromLiteral("s"))
		var hx_argument_6 hxrt.Anonymous = pos
		hxrt.FunctionCall(hx_callable_4)(hx_argument_5, hx_argument_6)
	}()
	return result
}

func haxe__Timer_stamp() float64 {
	return hxrt.ThreadNowSeconds()
}
