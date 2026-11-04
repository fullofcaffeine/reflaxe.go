package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulseAlert interface {
}

type app__core__PulseAlert struct {
	__hx_this     I_app__core__PulseAlert
	eventId       int
	source        *string
	region        *string
	severity      int
	weightedValue int
	reason        *string
}

func (self *app__core__PulseAlert) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__PulseAlert(eventId int, source *string, region *string, severity int, weightedValue int, reason *string) *app__core__PulseAlert {
	return hx_init_app__core__PulseAlert(hx_alloc_app__core__PulseAlert(nil), eventId, source, region, severity, weightedValue, reason)
}

func hx_alloc_app__core__PulseAlert(owner any) *app__core__PulseAlert {
	self := &app__core__PulseAlert{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__PulseAlert)
	}
	return self
}

func hx_init_app__core__PulseAlert(self *app__core__PulseAlert, eventId int, source *string, region *string, severity int, weightedValue int, reason *string) *app__core__PulseAlert {
	self.eventId = eventId
	self.source = source
	self.region = region
	self.severity = severity
	self.weightedValue = weightedValue
	self.reason = reason
	return self
}

func app__core__PulseAlert_fromEnriched(entry *app__core__PulseEnrichedEvent) *app__core__PulseAlert {
	var hx_if_1 *string
	if entry.severity >= 3 {
		hx_if_1 = hxrt.StringFromLiteral("critical")
	} else {
		hx_if_1 = hxrt.StringFromLiteral("warning")
	}
	label := hx_if_1
	return New_app__core__PulseAlert(entry.event.id, entry.event.source, entry.event.region, entry.severity, entry.weightedValue, label)
}
