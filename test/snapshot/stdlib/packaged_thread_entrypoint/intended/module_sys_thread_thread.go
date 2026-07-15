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

func New_sys__thread__Thread(id int) *sys__thread__Thread {
	self := &sys__thread__Thread{}
	self.__hx_this = self
	self.__id = id
	return self
}

func (self *sys__thread__Thread) get_events() *sys__thread__EventLoop {
	if !hxrt.ThreadHasEventLoop(self.__id) {
		hxrt.Throw(New_sys__thread__NoEventLoopException(hxrt.StringFromLiteral("Event loop is not available. Refer to sys.thread.Thread.runWithEventLoop."), nil))
		var hx_throw_zero_1 *sys__thread__EventLoop
		return hx_throw_zero_1
	}
	return sys__thread__EventLoop___fromHandle(hxrt.ThreadEvents(self.__id))
}

func (self *sys__thread__Thread) sendMessage(msg any) {
	hxrt.ThreadSendMessage(self.__id, msg)
}

func sys__thread__Thread_create(job func()) *sys__thread__Thread {
	return New_sys__thread__Thread(hxrt.ThreadSpawn(job))
}

func sys__thread__Thread_createWithEventLoop(job func()) *sys__thread__Thread {
	return New_sys__thread__Thread(hxrt.ThreadSpawnWithEventLoop(job))
}

func sys__thread__Thread_current() *sys__thread__Thread {
	return New_sys__thread__Thread(hxrt.ThreadCurrentId())
}

func sys__thread__Thread_processEvents() {
	current := sys__thread__Thread_current()
	if hxrt.ThreadHasEventLoop(current.__id) {
		current.get_events().progress()
	}
}

func sys__thread__Thread_readMessage(block bool) any {
	return hxrt.ThreadReadMessage(block)
}

func sys__thread__Thread_runWithEventLoop(job func()) {
	hxrt.ThreadRunWithEventLoop(job)
}
