package main

import "snapshot/hxrt"

type I_sys__thread__Thread interface {
	get_events() *sys__thread__EventLoop
	sendMessage(msg any)
}

type sys__thread__Thread struct {
	__hx_this I_sys__thread__Thread
	__id      int
	events    *sys__thread__EventLoop
}

func (self *sys__thread__Thread) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__Thread(id int) *sys__thread__Thread {
	return hx_init_sys__thread__Thread(hx_alloc_sys__thread__Thread(nil), id)
}

func hx_alloc_sys__thread__Thread(owner any) *sys__thread__Thread {
	self := &sys__thread__Thread{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__Thread)
	}
	return self
}

func hx_init_sys__thread__Thread(self *sys__thread__Thread, id int) *sys__thread__Thread {
	self.__id = id
	return self
}

func (self *sys__thread__Thread) get_events() *sys__thread__EventLoop {
	if !hxrt.ThreadHasEventLoop(self.__id) {
		hxrt.Throw(haxe__Exception_thrown(New_sys__thread__NoEventLoopException(hxrt.StringFromLiteral("Event loop is not available. Refer to sys.thread.Thread.runWithEventLoop."), nil)))
	}
	return sys__thread__EventLoop___fromHandle(hxrt.ThreadEvents(self.__id))
}

func (self *sys__thread__Thread) sendMessage(msg any) {
	hxrt.ThreadSendMessage(self.__id, msg)
}

func sys__thread__Thread_create(job *hxrt.Function[func()]) *sys__thread__Thread {
	return New_sys__thread__Thread(hxrt.ThreadSpawn(hxrt.FunctionNative(job), hxrt.FunctionNative(hxrt.NamedFunction(sys__thread__Thread_reportUncaught, "sys.thread.Thread.reportUncaught"))))
}

func sys__thread__Thread_createWithEventLoop(job *hxrt.Function[func()]) *sys__thread__Thread {
	return New_sys__thread__Thread(hxrt.ThreadSpawnWithEventLoop(hxrt.FunctionNative(job), hxrt.FunctionNative(hxrt.NamedFunction(sys__thread__Thread_reportUncaught, "sys.thread.Thread.reportUncaught"))))
}

func sys__thread__Thread_current() *sys__thread__Thread {
	return New_sys__thread__Thread(hxrt.ThreadCurrentId())
}

func sys__thread__Thread_processEvents() {
	current := sys__thread__Thread_current()
	if hxrt.ThreadHasEventLoop(current.__id) {
		current.__hx_this.get_events().__hx_this.progress()
	}
}

func sys__thread__Thread_readMessage(block bool) any {
	return hxrt.ThreadReadMessage(block)
}

func sys__thread__Thread_reportUncaught(value any) {
	exception := haxe__Exception_caught(value)
	hxrt.ThreadWriteStderr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Uncaught exception "), exception.__hx_this.get_message()), hxrt.StringFromLiteral("\n")))
}

func sys__thread__Thread_runWithEventLoop(job *hxrt.Function[func()]) {
	hxrt.ThreadRunWithEventLoop(hxrt.FunctionNative(job))
}
