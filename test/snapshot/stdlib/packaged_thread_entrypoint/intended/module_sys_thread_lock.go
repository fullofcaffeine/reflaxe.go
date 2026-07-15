package main

import "snapshot/hxrt"

type I_sys__thread__Lock interface {
	wait(timeout any) bool
	release()
}

type sys__thread__Lock struct {
	__hx_this I_sys__thread__Lock
	__h       *hxrt.LockHandle
}

func New_sys__thread__Lock() *sys__thread__Lock {
	self := &sys__thread__Lock{}
	self.__hx_this = self
	self.__h = hxrt.ThreadLockNew()
	return self
}

func (self *sys__thread__Lock) wait(timeout any) bool {
	if timeout == nil {
		return hxrt.ThreadLockWait(self.__h)
	}
	return hxrt.ThreadLockWaitTimeout(self.__h, timeout.(float64))
}

func (self *sys__thread__Lock) release() {
	hxrt.ThreadLockRelease(self.__h)
}
