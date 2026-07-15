package main

import "snapshot/hxrt"

type I_sys__thread__NoEventLoopException interface {
}

type sys__thread__NoEventLoopException struct {
	__hx_this      I_sys__thread__NoEventLoopException
	__hx_exception *hxrt.ExceptionValue
}

func New_sys__thread__NoEventLoopException(msg *string, previous *hxrt.ExceptionValue) *sys__thread__NoEventLoopException {
	self := &sys__thread__NoEventLoopException{}
	self.__hx_exception = hxrt.BindException(self, msg, previous, nil)
	self.__hx_this = self
	return self
}

func (self *sys__thread__NoEventLoopException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}
