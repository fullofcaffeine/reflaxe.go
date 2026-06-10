package main

import "snapshot/hxrt"

type I_sys__thread__EventLoop interface {
	repeat(event func(), intervalMs int) any
	cancel(eventHandler any)
	promise()
	run(event func())
	runPromised(event func())
	progress() *sys__thread__NextEventTime
	wait(timeout any) bool
	loop()
}

type sys__thread__EventLoop struct {
	__hx_this I_sys__thread__EventLoop
	__h       *hxrt.EventLoopHandle
}

func New_sys__thread__EventLoop() *sys__thread__EventLoop {
	self := &sys__thread__EventLoop{}
	self.__hx_this = self
	self.__h = hxrt.ThreadEventLoopNew()
	return self
}

func (self *sys__thread__EventLoop) repeat(event func(), intervalMs int) any {
	return any(hxrt.ThreadEventLoopRepeat(self.__h, event, intervalMs))
}

func (self *sys__thread__EventLoop) cancel(eventHandler any) {
	hxrt.ThreadEventLoopCancel(self.__h, hxrt.IntFromNullableAny(func(hx_value_2 any) int {
		if hx_value_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_value_2.(int)
	}(eventHandler)))
}

func (self *sys__thread__EventLoop) promise() {
	hxrt.ThreadEventLoopPromise(self.__h)
}

func (self *sys__thread__EventLoop) run(event func()) {
	hxrt.ThreadEventLoopRun(self.__h, event)
}

func (self *sys__thread__EventLoop) runPromised(event func()) {
	hxrt.ThreadEventLoopRunPromised(self.__h, event)
}

func (self *sys__thread__EventLoop) progress() *sys__thread__NextEventTime {
	result := hxrt.ThreadEventLoopProgress(self.__h)
	_g := result.Kind
	var hx_switch_4 *sys__thread__NextEventTime
	switch _g {
	case 0:
		hx_switch_4 = sys__thread__NextEventTime_Now
	case 1:
		hx_switch_4 = sys__thread__NextEventTime_Never
	case 2:
		var hx_if_5 *sys__thread__NextEventTime
		if result.Time < 0 {
			hx_if_5 = sys__thread__NextEventTime_AnyTime(nil)
		} else {
			hx_if_5 = sys__thread__NextEventTime_AnyTime(result.Time)
		}
		hx_switch_4 = hx_if_5
	case 3:
		hx_switch_4 = sys__thread__NextEventTime_At(result.Time)
	default:
		hx_switch_4 = sys__thread__NextEventTime_Never
	}
	return hx_switch_4
}

func (self *sys__thread__EventLoop) wait(timeout any) bool {
	if timeout == nil {
		return hxrt.ThreadEventLoopWait(self.__h)
	}
	return hxrt.ThreadEventLoopWaitTimeout(self.__h, timeout.(float64))
}

func (self *sys__thread__EventLoop) loop() {
	hxrt.ThreadEventLoopLoop(self.__h)
}

func sys__thread__EventLoop___fromHandle(handle *hxrt.EventLoopHandle) *sys__thread__EventLoop {
	loop := New_sys__thread__EventLoop()
	loop.__h = handle
	return loop
}

type sys__thread__NextEventTime struct {
	tag    int
	params []any
}

var sys__thread__NextEventTime_Now *sys__thread__NextEventTime = &sys__thread__NextEventTime{tag: 0}

var sys__thread__NextEventTime_Never *sys__thread__NextEventTime = &sys__thread__NextEventTime{tag: 1}

func sys__thread__NextEventTime_AnyTime(time any) *sys__thread__NextEventTime {
	enumValue := &sys__thread__NextEventTime{tag: 2}
	enumValue.params = []any{time}
	return enumValue
}

func sys__thread__NextEventTime_At(time float64) *sys__thread__NextEventTime {
	enumValue := &sys__thread__NextEventTime{tag: 3}
	enumValue.params = []any{time}
	return enumValue
}
