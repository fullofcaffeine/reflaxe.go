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

func (self *sys__thread__Mutex) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__Mutex() *sys__thread__Mutex {
	return hx_init_sys__thread__Mutex(hx_alloc_sys__thread__Mutex(nil))
}

func hx_alloc_sys__thread__Mutex(owner any) *sys__thread__Mutex {
	self := &sys__thread__Mutex{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__Mutex)
	}
	return self
}

func hx_init_sys__thread__Mutex(self *sys__thread__Mutex) *sys__thread__Mutex {
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
