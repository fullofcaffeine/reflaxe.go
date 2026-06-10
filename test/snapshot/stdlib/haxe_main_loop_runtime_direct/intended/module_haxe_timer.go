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
	self.eventHandler = self.thread.get_events().repeat(func() {
		_gthis.run()
	}, time_ms)
	return self
}

func (self *haxe__Timer) stop() {
	if self.eventHandler != any(0) {
		self.thread.get_events().cancel(self.eventHandler)
		self.eventHandler = any(0)
	}
}

func haxe__Timer_delay(f func(), time_ms int) *haxe__Timer {
	timer := New_haxe__Timer(time_ms)
	timer.run = func() {
		timer.stop()
		f()
	}
	return timer
}

func haxe__Timer_measure(f func() any, pos map[string]any) any {
	t0 := hxrt.ThreadNowSeconds()
	var result any = f()
	hxrt.Println(hxrt.StringConcatAny((hxrt.ThreadNowSeconds() - t0), hxrt.StringFromLiteral("s")))
	return result
}

func haxe__Timer_stamp() float64 {
	return hxrt.ThreadNowSeconds()
}
