package main

import "snapshot/hxrt"

type I_haxe__Timer interface {
	stop()
}

type haxe__Timer struct {
	__hx_this    I_haxe__Timer
	thread       *sys__thread__Thread
	eventHandler any
	run          func()
}

func New_haxe__Timer(time_ms int) *haxe__Timer {
	self := &haxe__Timer{}
	self.__hx_this = self
	_gthis := self
	self.run = func() {
	}
	self.thread = sys__thread__Thread_current()
	self.eventHandler = self.thread.__hx_this.get_events().__hx_this.repeat(func() {
		_gthis.run()
	}, time_ms)
	return self
}

func (self *haxe__Timer) stop() {
	if !hxrt.HaxeEqual(self.eventHandler, any(0)) {
		self.thread.__hx_this.get_events().__hx_this.cancel(self.eventHandler)
		self.eventHandler = any(0)
	}
}

func haxe__Timer_delay(f func(), time_ms int) *haxe__Timer {
	timer := New_haxe__Timer(time_ms)
	timer.run = func() {
		timer.__hx_this.stop()
		f()
	}
	return timer
}

func haxe__Timer_measure(f func() any, pos map[string]any) any {
	t0 := hxrt.ThreadNowSeconds()
	var result any = f()
	func(hx_fn func(any, map[string]any), hx_arg_0 any, hx_arg_1 map[string]any) {
		if hx_fn == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Invalid operation: null function"))
			return
		}
		hx_fn(hx_arg_0, hx_arg_1)
	}(haxe__Log_trace, hxrt.StringConcatAny((hxrt.ThreadNowSeconds()-t0), hxrt.StringFromLiteral("s")), pos)
	return result
}

func haxe__Timer_stamp() float64 {
	return hxrt.ThreadNowSeconds()
}
