package main

import "examples_pulseforge_portable/hxrt"

type I_app__runtime__CoreRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	parse(frames *hxrt.Array, workerCount int) *hxrt.Array
	enrich(events *hxrt.Array, workerCount int) *hxrt.Array
	stageScore(parsed *hxrt.Array, enriched *hxrt.Array, alerts *hxrt.Array, backpressureEvents int) int
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
	return hxrt.StringFromLiteral("portable")
}

func (self *app__runtime__CoreRuntime) variantId() *string {
	return hxrt.StringFromLiteral("core")
}

func (self *app__runtime__CoreRuntime) capabilityId() *string {
	return hxrt.StringFromLiteral("core_loop")
}

func (self *app__runtime__CoreRuntime) parse(frames *hxrt.Array, workerCount int) *hxrt.Array {
	parsed := hxrt.NewArray()
	_g := 0
	for _g < frames.Len() {
		frame := func(hx_value_1 any) *app__core__PulseIngressFrame {
			if hx_value_1 == nil {
				var hx_zero_2 *app__core__PulseIngressFrame
				return hx_zero_2
			}
			return hx_value_1.(*app__core__PulseIngressFrame)
		}(frames.Get(_g))
		_g = int(int32((_g + 1)))
		parsed.Push(app__core__PulseCodec_parse(frame))
	}
	return parsed
}

func (self *app__runtime__CoreRuntime) enrich(events *hxrt.Array, workerCount int) *hxrt.Array {
	enriched := hxrt.NewArray()
	_g := 0
	for _g < events.Len() {
		event := func(hx_value_4 any) *app__core__PulseEvent {
			if hx_value_4 == nil {
				var hx_zero_5 *app__core__PulseEvent
				return hx_zero_5
			}
			return hx_value_4.(*app__core__PulseEvent)
		}(events.Get(_g))
		_g = int(int32((_g + 1)))
		enriched.Push(app__core__PulseCodec_enrich(event))
	}
	return enriched
}

func (self *app__runtime__CoreRuntime) stageScore(parsed *hxrt.Array, enriched *hxrt.Array, alerts *hxrt.Array, backpressureEvents int) int {
	score := 0
	_g := 0
	for _g < enriched.Len() {
		entry := func(hx_value_7 any) *app__core__PulseEnrichedEvent {
			if hx_value_7 == nil {
				var hx_zero_8 *app__core__PulseEnrichedEvent
				return hx_zero_8
			}
			return hx_value_7.(*app__core__PulseEnrichedEvent)
		}(enriched.Get(_g))
		_g = int(int32((_g + 1)))
		score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(entry.weightedValue))))
	}
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(alerts.Len()) * hxrt.Int32Wrap(5))))))))
	score = int(int32((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(backpressureEvents) * hxrt.Int32Wrap(2))))))))
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(parsed.Len()))))
	return score
}
