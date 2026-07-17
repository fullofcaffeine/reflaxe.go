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

func New_app__core__PulseIngestResult(receivedCount int, acceptedFrames *hxrt.Array, backpressureEvents int) *app__core__PulseIngestResult {
	self := &app__core__PulseIngestResult{}
	self.__hx_this = self
	self.receivedCount = receivedCount
	self.acceptedFrames = acceptedFrames
	self.backpressureEvents = backpressureEvents
	return self
}
