package main

import "examples_fluxproxy_metal/hxrt"

type I_app__core__FluxIngestResult interface {
}

type app__core__FluxIngestResult struct {
	__hx_this          I_app__core__FluxIngestResult
	receivedCount      int
	acceptedRequests   *hxrt.Array
	backpressureEvents int
}

func New_app__core__FluxIngestResult(receivedCount int, acceptedRequests *hxrt.Array, backpressureEvents int) *app__core__FluxIngestResult {
	self := &app__core__FluxIngestResult{}
	self.__hx_this = self
	self.receivedCount = receivedCount
	self.acceptedRequests = acceptedRequests
	self.backpressureEvents = backpressureEvents
	return self
}
