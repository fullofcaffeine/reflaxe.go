package main

import "snapshot/hxrt"

func haxe__MainLoop___remove(event *haxe__MainEvent) {
	next := []*haxe__MainEvent{}
	_g := 0
	_g1 := haxe__MainLoop_pending
	for _g < len(_g1) {
		candidate := _g1[_g]
		_g = int(int32((_g + 1)))
		if candidate != event {
			next = append(next, candidate)
		}
	}
	haxe__MainLoop_pending = next
}

func haxe__MainLoop_add(f func(), priority int) *haxe__MainEvent {
	if f == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Event function is null"))
		var hx_throw_zero_9 *haxe__MainEvent
		return hx_throw_zero_9
	}
	event := New_haxe__MainEvent(f, priority)
	hx_arr_10 := haxe__MainLoop_pending
	hx_arr_10 = append(hx_arr_10, event)
	haxe__MainLoop_pending = hx_arr_10
	event.delayNow()
	return event
}

func haxe__MainLoop_addThread(f func()) {
	haxe__EntryPoint_addThread(f)
}

func haxe__MainLoop_get_threadCount() int {
	return haxe__EntryPoint_threadCount
}

func haxe__MainLoop_hasEvents() bool {
	_g := 0
	_g1 := haxe__MainLoop_pending
	for _g < len(_g1) {
		event := _g1[_g]
		_g = int(int32((_g + 1)))
		if event.isBlocking {
			return true
		}
	}
	return false
}

var haxe__MainLoop_pending []*haxe__MainEvent = []*haxe__MainEvent{}

func haxe__MainLoop_runInMainThread(f func()) {
	haxe__EntryPoint_runInMainThread(f)
}

var haxe__MainLoop_threadCount int

type I_haxe__MainEvent interface {
	delay(t float64)
	call()
	stop()
	delayNow()
	schedule()
	dispatch()
}

type haxe__MainEvent struct {
	__hx_this  I_haxe__MainEvent
	f          func()
	timer      *haxe__Timer
	active     bool
	isBlocking bool
	nextRun    float64
	priority   int
}

func New_haxe__MainEvent(f func(), priority int) *haxe__MainEvent {
	self := &haxe__MainEvent{}
	self.__hx_this = self
	self.isBlocking = true
	self.active = true
	self.f = f
	self.priority = priority
	self.nextRun = -1.0
	return self
}

func (self *haxe__MainEvent) delay(t float64) {
	self.nextRun = (hxrt.ThreadNowSeconds() + t)
	if self.timer != nil {
		self.timer.stop()
		self.timer = nil
	}
	self.schedule()
}

func (self *haxe__MainEvent) call() {
	if self.f != nil {
		self.f()
	}
}

func (self *haxe__MainEvent) stop() {
	if !self.active {
		return
	}
	self.active = false
	self.f = nil
	if self.timer != nil {
		self.timer.stop()
		self.timer = nil
	}
	haxe__MainLoop___remove(self)
}

func (self *haxe__MainEvent) delayNow() {
	self.nextRun = -1.0
	if self.timer != nil {
		self.timer.stop()
		self.timer = nil
	}
	self.schedule()
}

func (self *haxe__MainEvent) schedule() {
	if !self.active {
		return
	}
	wait := (self.nextRun - hxrt.ThreadNowSeconds())
	if wait <= 0 {
		haxe__EntryPoint_runInMainThread(self.dispatch)
	} else {
		ms := Math_ceil((wait * float64(1000)))
		if ms < 1 {
			ms = 1
		}
		self.timer = haxe__Timer_delay(self.dispatch, ms)
	}
}

func (self *haxe__MainEvent) dispatch() {
	if !self.active || (self.f == nil) {
		return
	}
	wait := (self.nextRun - hxrt.ThreadNowSeconds())
	if wait > 0 {
		self.schedule()
		return
	}
	if self.f != nil {
		self.f()
	}
	if self.active && (self.f != nil) {
		self.schedule()
	}
}
