package main

import "examples_tui_todo_metal/hxrt"

type I_profile__MetalRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(metrics *profile__TodoRuntimeMetrics) *string
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

func (self *profile__MetalRuntime) normalizeTitle(title *string) *string {
	return title
}

func (self *profile__MetalRuntime) normalizeTag(tag *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("metal-"), tag)
}

func (self *profile__MetalRuntime) supportsBatchAdd() bool {
	return true
}

func (self *profile__MetalRuntime) supportsDiagnostics() bool {
	return true
}

func (self *profile__MetalRuntime) diagnostics(metrics *profile__TodoRuntimeMetrics) *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("p1="), metrics.p1), hxrt.StringFromLiteral(",completed=")), metrics.done)
}
