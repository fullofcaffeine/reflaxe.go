package main

import "snapshot/hxrt"

func countAsFloat(count int) float64 {
	return float64(count)
}

func halfReal(value float64) float64 {
	return (value / float64(2))
}

func main() {
	var seven any = 7
	fromDynamic := hxrt.DynamicToFloatSlot(seven)
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("real from dynamic: "), halfReal(fromDynamic)))
	hxrt.Println(v)
	three := 3
	fromLocal := float64(three)
	fromLiteral := 7.0
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("real from int: "), halfReal(fromLocal)), hxrt.StringFromLiteral(" ")), halfReal(fromLiteral)))
	hxrt.Println(v_1)
	track := New_Track()
	track.length = float64(three)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("real field: "), halfReal(track.length)))
	hxrt.Println(v_2)
	count := 5
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("count as float: "), (countAsFloat(count) / float64(2))))
	hxrt.Println(v_3)
	level := 3.0
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("level as float: "), (level / float64(2)))))
	var maybe any = 1.5
	sure := func(hx_value_1 any) float64 {
		if hx_value_1 == nil {
			var hx_zero_2 float64
			return hx_zero_2
		}
		return hx_value_1.(float64)
	}(maybe)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("typed float: "), (sure / float64(2)))))
}

type I_Track interface {
}

type Track struct {
	__hx_this I_Track
	length    float64
}

func (self *Track) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Track() *Track {
	return hx_init_Track(hx_alloc_Track(nil))
}

func hx_alloc_Track(owner any) *Track {
	self := &Track{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Track)
	}
	return self
}

func hx_init_Track(self *Track) *Track {
	return self
}

func (self *Track) HxClassName() string {
	return "Track"
}
