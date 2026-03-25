package main

import "snapshot/hxrt"

type I_sys__thread__Semaphore interface {
	acquire()
	tryAcquire(timeout any) bool
	release()
}

type sys__thread__Semaphore struct {
	__hx_this I_sys__thread__Semaphore
	__h       *hxrt.SemaphoreHandle
}

func New_sys__thread__Semaphore(value int) *sys__thread__Semaphore {
	self := &sys__thread__Semaphore{}
	self.__hx_this = self
	self.__h = hxrt.ThreadSemaphoreNew(value)
	return self
}

func (self *sys__thread__Semaphore) acquire() {
	hxrt.ThreadSemaphoreAcquire(self.__h)
}

func (self *sys__thread__Semaphore) tryAcquire(timeout any) bool {
	if timeout == nil {
		return hxrt.ThreadSemaphoreTryAcquire(self.__h)
	}
	return hxrt.ThreadSemaphoreTryAcquireTimeoutAny(self.__h, timeout.(float64))
}

func (self *sys__thread__Semaphore) release() {
	hxrt.ThreadSemaphoreRelease(self.__h)
}
