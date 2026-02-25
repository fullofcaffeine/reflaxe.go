package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulseReport interface {
	lines() []*string
	profileId() *string
	variantId() *string
	capabilityId() *string
	ingestReceivedCount() int
	ingestAcceptedCount() int
	ingestBackpressureCount() int
	alertEventCount() int
	alertEventDigest() *string
	score() int
	render() *string
}

type app__core__PulseReport struct {
	__hx_this              I_app__core__PulseReport
	profile                *string
	variant                *string
	capability             *string
	ingestReceived         int
	ingestAccepted         int
	backpressureEvents     int
	parseCount             int
	enrichCount            int
	aggregateSourceCount   int
	aggregateTotalValue    int
	aggregateWeightedTotal int
	aggregateDigest        *string
	alertCount             int
	alertDigest            *string
	runtimeScore           int
}

func New_app__core__PulseReport(profile *string, variant *string, capability *string, ingestReceived int, ingestAccepted int, backpressureEvents int, parseCount int, enrichCount int, aggregateSourceCount int, aggregateTotalValue int, aggregateWeightedTotal int, aggregateDigest *string, alertCount int, alertDigest *string, runtimeScore int) *app__core__PulseReport {
	self := &app__core__PulseReport{}
	self.__hx_this = self
	self.profile = profile
	self.variant = variant
	self.capability = capability
	self.ingestReceived = ingestReceived
	self.ingestAccepted = ingestAccepted
	self.backpressureEvents = backpressureEvents
	self.parseCount = parseCount
	self.enrichCount = enrichCount
	self.aggregateSourceCount = aggregateSourceCount
	self.aggregateTotalValue = aggregateTotalValue
	self.aggregateWeightedTotal = aggregateWeightedTotal
	self.aggregateDigest = aggregateDigest
	self.alertCount = alertCount
	self.alertDigest = alertDigest
	self.runtimeScore = runtimeScore
	return self
}

func (self *app__core__PulseReport) lines() []*string {
	return []*string{hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pulseforge.profile="), self.profile), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pulseforge.variant="), self.variant), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("runtime.capability="), self.capability), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingest.received="), self.ingestReceived), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingest.accepted="), self.ingestAccepted), hxrt.StringConcatAny(hxrt.StringFromLiteral("ingest.backpressure="), self.backpressureEvents), hxrt.StringConcatAny(hxrt.StringFromLiteral("parse.events="), self.parseCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("enrich.events="), self.enrichCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("aggregate.sources="), self.aggregateSourceCount), hxrt.StringConcatAny(hxrt.StringFromLiteral("aggregate.total="), self.aggregateTotalValue), hxrt.StringConcatAny(hxrt.StringFromLiteral("aggregate.weighted_total="), self.aggregateWeightedTotal), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("aggregate.summary="), self.aggregateDigest), hxrt.StringConcatAny(hxrt.StringFromLiteral("alert.count="), self.alertCount), hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("alert.events="), self.alertDigest), hxrt.StringConcatAny(hxrt.StringFromLiteral("runtime.score="), self.runtimeScore)}
}

func (self *app__core__PulseReport) profileId() *string {
	return self.profile
}

func (self *app__core__PulseReport) variantId() *string {
	return self.variant
}

func (self *app__core__PulseReport) capabilityId() *string {
	return self.capability
}

func (self *app__core__PulseReport) ingestReceivedCount() int {
	return self.ingestReceived
}

func (self *app__core__PulseReport) ingestAcceptedCount() int {
	return self.ingestAccepted
}

func (self *app__core__PulseReport) ingestBackpressureCount() int {
	return self.backpressureEvents
}

func (self *app__core__PulseReport) alertEventCount() int {
	return self.alertCount
}

func (self *app__core__PulseReport) alertEventDigest() *string {
	return self.alertDigest
}

func (self *app__core__PulseReport) score() int {
	return self.runtimeScore
}

func (self *app__core__PulseReport) render() *string {
	out := hxrt.StringFromLiteral("")
	values := self.lines()
	i := 0
	for i < len(values) {
		if i > 0 {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral("\n"))
		}
		out = hxrt.StringConcatStringPtr(out, values[i])
		i = int(int32((i + 1)))
	}
	return out
}
