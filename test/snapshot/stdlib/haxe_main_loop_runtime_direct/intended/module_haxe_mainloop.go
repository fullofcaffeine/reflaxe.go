package main

import "snapshot/hxrt"

func haxe__MainLoop___remove(event *haxe__MainEvent) {
	next := hxrt.NewArray()
	_g := 0
	_g1 := haxe__MainLoop_pending
	for _g < _g1.Len() {
		candidate := hx_recover_class_0(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if candidate != event {
			next.Push(candidate)
		}
	}
	haxe__MainLoop_pending = next
}

func haxe__MainLoop_add(f *hxrt.Function[func()], priority int) *haxe__MainEvent {
	if f == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Event function is null")))
	}
	event := New_haxe__MainEvent(f, priority)
	hx_arr_2 := haxe__MainLoop_pending
	hx_arr_2.Push(event)
	event.__hx_this.delayNow()
	return event
}

func haxe__MainLoop_addThread(f *hxrt.Function[func()]) {
	haxe__EntryPoint_addThread(f)
}

func haxe__MainLoop_get_threadCount() int {
	return haxe__EntryPoint_threadCount
}

func haxe__MainLoop_hasEvents() bool {
	_g := 0
	_g1 := haxe__MainLoop_pending
	for _g < _g1.Len() {
		event := hx_recover_class_0(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if event.isBlocking {
			return true
		}
	}
	return false
}

var haxe__MainLoop_pending *hxrt.Array = hxrt.NewArray()

func haxe__MainLoop_runInMainThread(f *hxrt.Function[func()]) {
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
	f          *hxrt.Function[func()]
	timer      *haxe__Timer
	active     bool
	isBlocking bool
	nextRun    float64
	priority   int
}

func (self *haxe__MainEvent) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__MainEvent(f *hxrt.Function[func()], priority int) *haxe__MainEvent {
	return hx_init_haxe__MainEvent(hx_alloc_haxe__MainEvent(nil), f, priority)
}

func hx_alloc_haxe__MainEvent(owner any) *haxe__MainEvent {
	self := &haxe__MainEvent{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__MainEvent)
	}
	return self
}

func hx_init_haxe__MainEvent(self *haxe__MainEvent, f *hxrt.Function[func()], priority int) *haxe__MainEvent {
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
		self.timer.__hx_this.stop()
		self.timer = nil
	}
	self.__hx_this.schedule()
}

func (self *haxe__MainEvent) call() {
	if self.f != nil {
		func() {
			var hx_callable_3 *hxrt.Function[func()] = self.f
			hxrt.FunctionCall(hx_callable_3)()
		}()
	}
}

func (self *haxe__MainEvent) stop() {
	if !self.active {
		return
	}
	self.active = false
	self.f = nil
	if self.timer != nil {
		self.timer.__hx_this.stop()
		self.timer = nil
	}
	haxe__MainLoop___remove(self)
}

func (self *haxe__MainEvent) delayNow() {
	self.nextRun = -1.0
	if self.timer != nil {
		self.timer.__hx_this.stop()
		self.timer = nil
	}
	self.__hx_this.schedule()
}

func (self *haxe__MainEvent) schedule() {
	if !self.active {
		return
	}
	wait := (self.nextRun - hxrt.ThreadNowSeconds())
	if wait <= 0 {
		haxe__EntryPoint_runInMainThread(func() *hxrt.Function[func()] {
			hx_receiver_4 := self
			return hxrt.BoundFunction(func() {
				hx_receiver_4.__hx_this.dispatch()
			}, func() any {
				return hx_receiver_4.__hx_this
			}, "dispatch")
		}())
	} else {
		ms := hxrt.MathCeilInt((wait * float64(1000)))
		if ms < 1 {
			ms = 1
		}
		self.timer = haxe__Timer_delay(func() *hxrt.Function[func()] {
			hx_receiver_5 := self
			return hxrt.BoundFunction(func() {
				hx_receiver_5.__hx_this.dispatch()
			}, func() any {
				return hx_receiver_5.__hx_this
			}, "dispatch")
		}(), ms)
	}
}

func (self *haxe__MainEvent) dispatch() {
	if !self.active || (self.f == nil) {
		return
	}
	wait := (self.nextRun - hxrt.ThreadNowSeconds())
	if wait > 0 {
		self.__hx_this.schedule()
		return
	}
	if self.f != nil {
		func() {
			var hx_callable_6 *hxrt.Function[func()] = self.f
			hxrt.FunctionCall(hx_callable_6)()
		}()
	}
	if self.active && (self.f != nil) {
		self.__hx_this.schedule()
	}
}
