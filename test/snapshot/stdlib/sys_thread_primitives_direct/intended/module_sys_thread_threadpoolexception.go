package main

import "snapshot/hxrt"

type I_sys__thread__ThreadPoolException interface {
}

type sys__thread__ThreadPoolException struct {
	__hx_this      I_sys__thread__ThreadPoolException
	__hx_exception *hxrt.ExceptionValue
}

func New_sys__thread__ThreadPoolException(message *string, previous *hxrt.ExceptionValue, native any) *sys__thread__ThreadPoolException {
	self := &sys__thread__ThreadPoolException{}
	self.__hx_exception = hxrt.BindException(self, message, previous, native)
	self.__hx_this = self
	return self
}

func (self *sys__thread__ThreadPoolException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}
