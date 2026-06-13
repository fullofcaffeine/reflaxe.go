package main

import "examples_incident_api_portable/hxrt"

type I_app__core__IncidentRequestException interface {
}

type app__core__IncidentRequestException struct {
	__hx_this      I_app__core__IncidentRequestException
	code           *string
	__hx_exception *hxrt.ExceptionValue
}

func New_app__core__IncidentRequestException(code *string) *app__core__IncidentRequestException {
	self := &app__core__IncidentRequestException{}
	self.__hx_exception = hxrt.BindException(self, code, nil, nil)
	self.__hx_this = self
	self.code = code
	return self
}

func (self *app__core__IncidentRequestException) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}
