package main

import "examples_pulseforge_metal/hxrt"

type I_app__core__PulseIngestResult interface {
}

type app__core__PulseIngestResult struct {
	__hx_this          I_app__core__PulseIngestResult
	receivedCount      int
	acceptedFrames     *hxrt.Array
	backpressureEvents int
}

func (self *app__core__PulseIngestResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__PulseIngestResult(receivedCount int, acceptedFrames *hxrt.Array, backpressureEvents int) *app__core__PulseIngestResult {
	return hx_init_app__core__PulseIngestResult(hx_alloc_app__core__PulseIngestResult(nil), receivedCount, acceptedFrames, backpressureEvents)
}

func hx_alloc_app__core__PulseIngestResult(owner any) *app__core__PulseIngestResult {
	self := &app__core__PulseIngestResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__PulseIngestResult)
	}
	return self
}

func hx_init_app__core__PulseIngestResult(self *app__core__PulseIngestResult, receivedCount int, acceptedFrames *hxrt.Array, backpressureEvents int) *app__core__PulseIngestResult {
	self.receivedCount = receivedCount
	self.acceptedFrames = acceptedFrames
	self.backpressureEvents = backpressureEvents
	return self
}
