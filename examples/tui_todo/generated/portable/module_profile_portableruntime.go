package main

import "examples_tui_todo_portable/hxrt"

type I_profile__PortableRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
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

func (self *profile__PortableRuntime) supportsBatchAdd() bool {
	return false
}

func (self *profile__PortableRuntime) supportsDiagnostics() bool {
	return false
}

func (self *profile__PortableRuntime) diagnostics(items *haxe__ds__List) *string {
	return hxrt.StringFromLiteral("off")
}
