package main

import "net/http"

type I_GoHTTPRequestResult interface {
}

type GoHTTPRequestResult struct {
	__hx_this I_GoHTTPRequestResult
	request   *http.Request
	error     *go___Error
}

func New_GoHTTPRequestResult(request *http.Request, error *go___Error) *GoHTTPRequestResult {
	self := &GoHTTPRequestResult{}
	self.__hx_this = self
	self.request = request
	self.error = error
	return self
}
