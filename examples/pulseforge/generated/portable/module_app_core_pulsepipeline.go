package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulsePipeline interface {
	run(events []*app__core__PulseEvent) *app__core__PulseReport
}

type app__core__PulsePipeline struct {
	__hx_this I_app__core__PulsePipeline
	runtime   app__runtime__PulseRuntime
}

func New_app__core__PulsePipeline(runtime app__runtime__PulseRuntime) *app__core__PulsePipeline {
	self := &app__core__PulsePipeline{}
	self.__hx_this = self
	self.runtime = runtime
	return self
}

func (self *app__core__PulsePipeline) run(events []*app__core__PulseEvent) *app__core__PulseReport {
	alertCount := 0
	_ = alertCount
	totalValue := 0
	_ = totalValue
	_g := 0
	for _g < len(events) {
		event := events[_g]
		_ = event
		_g = int(int32((_g + 1)))
		totalValue = int(int32((hxrt.Int32Wrap(totalValue) + hxrt.Int32Wrap(event.value))))
		if event.value >= 8 {
			alertCount = int(int32((alertCount + 1)))
		}
	}
	runtimeScore := self.runtime.processScore(events)
	return New_app__core__PulseReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), len(events), totalValue, alertCount, runtimeScore)
}
