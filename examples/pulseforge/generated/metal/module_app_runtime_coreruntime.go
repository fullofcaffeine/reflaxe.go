package main

import "examples_pulseforge_metal/hxrt"

type I_app__runtime__CoreRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	parse(frames []*app__core__PulseIngressFrame, workerCount int) []*app__core__PulseEvent
	enrich(events []*app__core__PulseEvent, workerCount int) []*app__core__PulseEnrichedEvent
	stageScore(parsed []*app__core__PulseEvent, enriched []*app__core__PulseEnrichedEvent, alerts []*app__core__PulseAlert, backpressureEvents int) int
}

type app__runtime__CoreRuntime struct {
	__hx_this I_app__runtime__CoreRuntime
}

func New_app__runtime__CoreRuntime() *app__runtime__CoreRuntime {
	self := &app__runtime__CoreRuntime{}
	self.__hx_this = self
	return self
}

func (self *app__runtime__CoreRuntime) profileId() *string {
	return hxrt.StringFromLiteral("metal")
}

func (self *app__runtime__CoreRuntime) variantId() *string {
	return hxrt.StringFromLiteral("core")
}

func (self *app__runtime__CoreRuntime) capabilityId() *string {
	return hxrt.StringFromLiteral("core_loop")
}

func (self *app__runtime__CoreRuntime) parse(frames []*app__core__PulseIngressFrame, workerCount int) []*app__core__PulseEvent {
	parsed := []*app__core__PulseEvent{}
	_g := 0
	for _g < len(frames) {
		frame := frames[_g]
		_g = int(int32((_g + 1)))
		parsed = append(parsed, app__core__PulseCodec_parse(frame))
	}
	return parsed
}

func (self *app__runtime__CoreRuntime) enrich(events []*app__core__PulseEvent, workerCount int) []*app__core__PulseEnrichedEvent {
	enriched := []*app__core__PulseEnrichedEvent{}
	_g := 0
	for _g < len(events) {
		event := events[_g]
		_g = int(int32((_g + 1)))
		enriched = append(enriched, app__core__PulseCodec_enrich(event))
	}
	return enriched
}

func (self *app__runtime__CoreRuntime) stageScore(parsed []*app__core__PulseEvent, enriched []*app__core__PulseEnrichedEvent, alerts []*app__core__PulseAlert, backpressureEvents int) int {
	score := 0
	_g := 0
	for _g < len(enriched) {
		entry := enriched[_g]
		_g = int(int32((_g + 1)))
		score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(entry.weightedValue))))
	}
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(len(alerts)) * hxrt.Int32Wrap(5))))))))
	score = int(int32((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(backpressureEvents) * hxrt.Int32Wrap(2))))))))
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(len(parsed)))))
	return score
}
