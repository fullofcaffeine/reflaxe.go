package main

import "examples_tui_todo_metal/hxrt"

type I_profile__MetalRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
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

func (self *profile__MetalRuntime) diagnostics(items *haxe__ds__List) *string {
	p1 := 0
	_ = p1
	completed := 0
	_ = completed
	count := items.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_35 any) *model__TodoItem {
			if hx_value_35 == nil {
				var hx_zero_36 *model__TodoItem
				return hx_zero_36
			}
			return hx_value_35.(*model__TodoItem)
		}(items.pop())
		if value == nil {
			break
		}
		item := value
		if item.priority == 1 {
			p1 = int(int32((p1 + 1)))
		}
		if item.done {
			completed = int(int32((completed + 1)))
		}
		items.add(item)
		i = int(int32((i + 1)))
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("p1="), p1), hxrt.StringFromLiteral(",completed=")), completed)
}
