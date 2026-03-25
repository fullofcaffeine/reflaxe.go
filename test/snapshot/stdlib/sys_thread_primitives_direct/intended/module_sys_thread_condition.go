package main

import "snapshot/hxrt"

type I_sys__thread__Condition interface {
	acquire()
	tryAcquire() bool
	release()
	wait()
	signal()
	broadcast()
}

type sys__thread__Condition struct {
	__hx_this I_sys__thread__Condition
	__h       *hxrt.ConditionHandle
}

func New_sys__thread__Condition() *sys__thread__Condition {
	self := &sys__thread__Condition{}
	self.__hx_this = self
	self.__h = hxrt.ThreadConditionNew()
	return self
}

func (self *sys__thread__Condition) acquire() {
	hxrt.ThreadConditionAcquire(self.__h)
}

func (self *sys__thread__Condition) tryAcquire() bool {
	return hxrt.ThreadConditionTryAcquire(self.__h)
}

func (self *sys__thread__Condition) release() {
	hxrt.ThreadConditionRelease(self.__h)
}

func (self *sys__thread__Condition) wait() {
	hxrt.ThreadConditionWait(self.__h)
}

func (self *sys__thread__Condition) signal() {
	hxrt.ThreadConditionSignal(self.__h)
}

func (self *sys__thread__Condition) broadcast() {
	hxrt.ThreadConditionBroadcast(self.__h)
}
