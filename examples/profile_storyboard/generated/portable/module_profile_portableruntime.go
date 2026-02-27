package main

import "examples_profile_storyboard_portable/hxrt"

type I_profile__PortableRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(metrics *profile__StorySignalMetrics) *string
	supportsVelocityHint() bool
	velocityPerSprint() int
	riskThreshold() int
}

type profile__PortableRuntime struct {
	__hx_this I_profile__PortableRuntime
}

func New_profile__PortableRuntime() *profile__PortableRuntime {
	self := &profile__PortableRuntime{}
	self.__hx_this = self
	return self
}

func (self *profile__PortableRuntime) profileId() *string {
	return hxrt.StringFromLiteral("portable")
}

func (self *profile__PortableRuntime) decorateTitle(title *string) *string {
	return title
}

func (self *profile__PortableRuntime) highlightTag(tag *string) *string {
	return tag
}

func (self *profile__PortableRuntime) extraSignal(metrics *profile__StorySignalMetrics) *string {
	return hxrt.StringFromLiteral("interop_lane=off,optimizer=stable,policy_gate=off")
}

func (self *profile__PortableRuntime) supportsVelocityHint() bool {
	return false
}

func (self *profile__PortableRuntime) velocityPerSprint() int {
	return 5
}

func (self *profile__PortableRuntime) riskThreshold() int {
	return 5
}
