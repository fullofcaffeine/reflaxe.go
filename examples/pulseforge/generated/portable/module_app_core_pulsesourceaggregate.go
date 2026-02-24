package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulseSourceAggregate interface {
	record(entry *app__core__PulseEnrichedEvent)
	summaryToken() *string
}

type app__core__PulseSourceAggregate struct {
	__hx_this     I_app__core__PulseSourceAggregate
	source        *string
	count         int
	totalValue    int
	totalWeighted int
	maxValue      int
	maxSeverity   int
}

func New_app__core__PulseSourceAggregate(source *string) *app__core__PulseSourceAggregate {
	self := &app__core__PulseSourceAggregate{}
	self.__hx_this = self
	self.source = source
	self.count = 0
	self.totalValue = 0
	self.totalWeighted = 0
	self.maxValue = 0
	self.maxSeverity = 0
	return self
}

func (self *app__core__PulseSourceAggregate) record(entry *app__core__PulseEnrichedEvent) {
	self.count = int(int32((self.count + 1)))
	self.totalValue = int(int32((hxrt.Int32Wrap(self.totalValue) + hxrt.Int32Wrap(entry.event.value))))
	self.totalWeighted = int(int32((hxrt.Int32Wrap(self.totalWeighted) + hxrt.Int32Wrap(entry.weightedValue))))
	if entry.event.value > self.maxValue {
		self.maxValue = entry.event.value
	}
	if entry.severity > self.maxSeverity {
		self.maxSeverity = entry.severity
	}
}

func (self *app__core__PulseSourceAggregate) summaryToken() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(self.source, hxrt.StringFromLiteral(":")), self.count), hxrt.StringFromLiteral("/")), self.totalValue), hxrt.StringFromLiteral("/")), self.totalWeighted), hxrt.StringFromLiteral("/sev")), self.maxSeverity)
}
