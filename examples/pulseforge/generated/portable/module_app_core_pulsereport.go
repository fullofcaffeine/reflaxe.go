package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulseReport interface {
	lines() []*string
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
