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

func (self *sys__thread__Semaphore) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__Semaphore(value int) *sys__thread__Semaphore {
	return hx_init_sys__thread__Semaphore(hx_alloc_sys__thread__Semaphore(nil), value)
}

func hx_alloc_sys__thread__Semaphore(owner any) *sys__thread__Semaphore {
	self := &sys__thread__Semaphore{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__Semaphore)
	}
	return self
}

func hx_init_sys__thread__Semaphore(self *sys__thread__Semaphore, value int) *sys__thread__Semaphore {
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
	return hxrt.ThreadSemaphoreTryAcquireTimeout(self.__h, timeout.(float64))
}

func (self *sys__thread__Semaphore) release() {
	hxrt.ThreadSemaphoreRelease(self.__h)
}
