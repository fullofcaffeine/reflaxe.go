package main

import "examples_profile_storyboard_metal/hxrt"

type I_profile__MetalRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(metrics *profile__StorySignalMetrics) *string
	supportsVelocityHint() bool
	velocityPerSprint() int
	riskThreshold() int
}

type profile__MetalRuntime struct {
	__hx_this I_profile__MetalRuntime
}

func New_profile__MetalRuntime() *profile__MetalRuntime {
	self := &profile__MetalRuntime{}
	self.__hx_this = self
	return self
}

func (self *profile__MetalRuntime) profileId() *string {
	return hxrt.StringFromLiteral("metal")
}

func (self *profile__MetalRuntime) decorateTitle(title *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("[strict] "), title)
}

func (self *profile__MetalRuntime) highlightTag(tag *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("metal-"), tag)
}

func (self *profile__MetalRuntime) extraSignal(metrics *profile__StorySignalMetrics) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("interop_lane=typed+strict,high_value="), metrics.highValue), hxrt.StringFromLiteral(",open_high_value=")), metrics.openHighValue), hxrt.StringFromLiteral(",policy_gate=on"))
}

func (self *profile__MetalRuntime) supportsVelocityHint() bool {
	return true
}

func (self *profile__MetalRuntime) velocityPerSprint() int {
	return 9
}

func (self *profile__MetalRuntime) riskThreshold() int {
	return 4
}
