package main

import "snapshot/hxrt"

type I_sys__thread__Mutex interface {
	acquire()
	tryAcquire() bool
	release()
}

type sys__thread__Mutex struct {
	__hx_this I_sys__thread__Mutex
	__h       *hxrt.MutexHandle
}

func New_sys__thread__Mutex() *sys__thread__Mutex {
	self := &sys__thread__Mutex{}
	self.__hx_this = self
	self.__h = hxrt.ThreadMutexNew()
	return self
}

func (self *sys__thread__Mutex) acquire() {
	hxrt.ThreadMutexAcquire(self.__h)
}

func (self *sys__thread__Mutex) tryAcquire() bool {
	return hxrt.ThreadMutexTryAcquire(self.__h)
}

func (self *sys__thread__Mutex) release() {
	hxrt.ThreadMutexRelease(self.__h)
}
