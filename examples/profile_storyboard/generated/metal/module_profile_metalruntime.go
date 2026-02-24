package main

import "examples_profile_storyboard_metal/hxrt"

type I_profile__MetalRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(cards *haxe__ds__List) *string
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

func (self *profile__MetalRuntime) extraSignal(cards *haxe__ds__List) *string {
	highValue := 0
	_ = highValue
	openHighValue := 0
	_ = openHighValue
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_27 any) *domain__StoryCard {
			if hx_value_27 == nil {
				var hx_zero_28 *domain__StoryCard
				return hx_zero_28
			}
			return hx_value_27.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		if card.points >= 5 {
			highValue = int(int32((highValue + 1)))
			if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
				openHighValue = int(int32((openHighValue + 1)))
			}
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("interop_lane=typed+strict,high_value="), highValue), hxrt.StringFromLiteral(",open_high_value=")), openHighValue), hxrt.StringFromLiteral(",policy_gate=on"))
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
