package main

import "net/http"

type I_GoHTTPRequestResult interface {
}

type GoHTTPRequestResult struct {
	__hx_this I_GoHTTPRequestResult
	request   *http.Request
	error     *go___Error
}

func (self *GoHTTPRequestResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_GoHTTPRequestResult(request *http.Request, error *go___Error) *GoHTTPRequestResult {
	return hx_init_GoHTTPRequestResult(hx_alloc_GoHTTPRequestResult(nil), request, error)
}

func hx_alloc_GoHTTPRequestResult(owner any) *GoHTTPRequestResult {
	self := &GoHTTPRequestResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_GoHTTPRequestResult)
	}
	return self
}

func hx_init_GoHTTPRequestResult(self *GoHTTPRequestResult, request *http.Request, error *go___Error) *GoHTTPRequestResult {
	self.request = request
	self.error = error
	return self
}
