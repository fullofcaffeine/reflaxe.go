package main

import "examples_fluxproxy_portable/hxrt"

type I_app__core__FluxIngestResult interface {
}

type app__core__FluxIngestResult struct {
	__hx_this          I_app__core__FluxIngestResult
	receivedCount      int
	acceptedRequests   *hxrt.Array
	backpressureEvents int
}

func (self *app__core__FluxIngestResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__FluxIngestResult(receivedCount int, acceptedRequests *hxrt.Array, backpressureEvents int) *app__core__FluxIngestResult {
	return hx_init_app__core__FluxIngestResult(hx_alloc_app__core__FluxIngestResult(nil), receivedCount, acceptedRequests, backpressureEvents)
}

func hx_alloc_app__core__FluxIngestResult(owner any) *app__core__FluxIngestResult {
	self := &app__core__FluxIngestResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__FluxIngestResult)
	}
	return self
}

func hx_init_app__core__FluxIngestResult(self *app__core__FluxIngestResult, receivedCount int, acceptedRequests *hxrt.Array, backpressureEvents int) *app__core__FluxIngestResult {
	self.receivedCount = receivedCount
	self.acceptedRequests = acceptedRequests
	self.backpressureEvents = backpressureEvents
	return self
}
