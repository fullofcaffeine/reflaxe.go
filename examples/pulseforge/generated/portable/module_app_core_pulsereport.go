package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulseReport interface {
	lines() []*string
}

type app__core__PulseReport struct {
	__hx_this    I_app__core__PulseReport
	profile      *string
	variant      *string
	capability   *string
	ingestCount  int
	totalValue   int
	alertCount   int
	runtimeScore int
}

func New_app__core__PulseReport(profile *string, variant *string, capability *string, ingestCount int, totalValue int, alertCount int, runtimeScore int) *app__core__PulseReport {
	self := &app__core__PulseReport{}
	self.__hx_this = self
	self.profile = profile
	self.variant = variant
	self.capability = capability
	self.ingestCount = ingestCount
	self.totalValue = totalValue
	self.alertCount = alertCount
	self.runtimeScore = runtimeScore
	return self
}

func (self *app__core__PulseReport) lines() []*string {
	return []*string{hxrt.StringConcatAny(hxrt.StringFromLiteral("pulseforge.profile="), self.profile), hxrt.StringConcatAny(hxrt.StringFromLiteral("pulseforge.variant="), self.variant), hxrt.StringConcatAny(hxrt.StringFromLiteral("runtime.capability="), self.capability), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingest.events="), self.ingestCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("pipeline.total="), self.totalValue), hxrt.StringConcatAny(hxrt.StringFromLiteral("pipeline.alerts="), self.alertCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("runtime.score="), self.runtimeScore)}
}
