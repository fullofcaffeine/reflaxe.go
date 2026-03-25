package main

import "snapshot/hxrt"

type I_sys__thread__FixedThreadPoolShutdownException interface {
}

type sys__thread__FixedThreadPoolShutdownException struct {
	__hx_this      I_sys__thread__FixedThreadPoolShutdownException
	__hx_exception *hxrt.ExceptionValue
}

func New_sys__thread__FixedThreadPoolShutdownException(message *string, previous *hxrt.ExceptionValue, native any) *sys__thread__FixedThreadPoolShutdownException {
	self := &sys__thread__FixedThreadPoolShutdownException{}
	self.__hx_exception = hxrt.BindException(self, message, previous, native)
	self.__hx_this = self
	return self
}

func (self *sys__thread__FixedThreadPoolShutdownException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}
