package main

import "examples_tui_todo_portable/hxrt"

type I_profile__PortableRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	diagnostics(metrics *profile__TodoRuntimeMetrics) *string
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

func (self *profile__PortableRuntime) normalizeTitle(title *string) *string {
	return title
}

func (self *profile__PortableRuntime) normalizeTag(tag *string) *string {
	return tag
}

func (self *profile__PortableRuntime) diagnostics(metrics *profile__TodoRuntimeMetrics) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("p1="), metrics.p1), hxrt.StringFromLiteral(",completed=")), metrics.done)
}
