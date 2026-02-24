package main

import "examples_pulseforge_portable/hxrt"

type I_app__runtime__CoreRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	processScore(events []*app__core__PulseEvent) int
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

func (self *app__runtime__CoreRuntime) processScore(events []*app__core__PulseEvent) int {
	score := 0
	_ = score
	_g := 0
	_ = _g
	for _g < len(events) {
		event := events[_g]
		_ = event
		_g = int(int32((_g + 1)))
		score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(event.value))))
		if event.value >= 8 {
			score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(3))))
		}
	}
	return score
}
