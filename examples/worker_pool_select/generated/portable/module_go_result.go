package main

import "examples_worker_pool_select_portable/hxrt"

type I_go___Result interface {
	isOk() bool
	isErr() bool
	unwrap() any
	error() *string
}

type go___Result struct {
	__hx_this  I_go___Result
	value      any
	errorValue *go___Error
}

func New_go___Result(value any, errorValue *go___Error) *go___Result {
	self := &go___Result{}
	self.__hx_this = self
	self.value = value
	self.errorValue = errorValue
	return self
}

func (self *go___Result) isOk() bool {
	return (self.errorValue == nil)
}

func (self *go___Result) isErr() bool {
	return (self.errorValue != nil)
}

func (self *go___Result) unwrap() any {
	if self.errorValue != nil {
		hxrt.Throw(self.errorValue.__hx_this.toString())
		var hx_throw_zero_11 any
		return hx_throw_zero_11
	}
	return self.value
}

func (self *go___Result) error() *string {
	if self.errorValue == nil {
		return nil
	}
	return self.errorValue.__hx_this.toString()
}

func go___Result_failure(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go___Result_ok(value any) *go___Result {
	return New_go___Result(value, nil)
}
