package main

import "snapshot/hxrt"

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	value     any
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(value any) *Box {
	return hx_init_Box(hx_alloc_Box(nil), value)
}

func hx_alloc_Box(owner any) *Box {
	self := &Box{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Box)
	}
	return self
}

func hx_init_Box(self *Box, value any) *Box {
	self.value = value
	return self
}

func main() {
	parts := hxrt.NewArray(hxrt.StringFromLiteral("ab"), hxrt.StringFromLiteral("c"))
	raw := hxrt.StringJoinAny(parts.Values(), hxrt.StringFromLiteral(""))
	var this1 *string
	this1 = raw
	a := this1
	raw_1 := hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("a"), hxrt.StringSubstrStringPtr(hxrt.StringFromLiteral("bc"), 0, 0, false))
	var this1_1 *string
	this1_1 = raw_1
	b := this1_1
	var this1_2 *string
	this1_2 = hxrt.StringFromLiteral("abd")
	other := this1_2
	hxrt.Println(any(hxrt.StringEqualStringPtr(a, b)))
	hxrt.Println(any(!hxrt.StringEqualStringPtr(a, b)))
	hxrt.Println(any(hxrt.StringEqualStringPtr(a, other)))
	tokens := hxrt.NewArray(a, other)
	var v any = any(func() int {
		hx_indexof_target_1 := tokens
		hx_indexof_value_2 := b
		var hx_indexof_start_input_3 any = 0
		hx_indexof_start_4 := hxrt.IntFromNullableAny(hx_indexof_start_input_3)
		hx_indexof_length_5 := hx_indexof_target_1.Len()
		if hx_indexof_start_4 < 0 {
			hx_indexof_start_4 = (hx_indexof_length_5 + hx_indexof_start_4)
		}
		if hx_indexof_start_4 < 0 {
			hx_indexof_start_4 = 0
		}
		hx_indexof_index_6 := hx_indexof_start_4
		for hx_indexof_index_6 < hx_indexof_length_5 {
			hx_indexof_element_7 := func(hx_value_8 any) *string {
				if hx_value_8 == nil {
					var hx_zero_9 *string
					return hx_zero_9
				}
				return hx_value_8.(*string)
			}(hx_indexof_target_1.Get(hx_indexof_index_6))
			if hxrt.StringEqualStringPtr(hx_indexof_element_7, hx_indexof_value_2) {
				return hx_indexof_index_6
			}
			hx_indexof_index_6 = (hx_indexof_index_6 + 1)
		}
		return -1
	}())
	hxrt.Println(v)
	var v_1 any = any(func() bool {
		var hx_remove_value_11 any = b
		hx_remove_index_12 := 0
		for hx_remove_index_12 < tokens.Len() {
			hx_remove_element_13 := tokens.Get(hx_remove_index_12)
			if hxrt.HaxeEqual(hx_remove_element_13, hx_remove_value_11) {
				tokens.RemoveAt(hx_remove_index_12)
				return true
			}
			hx_remove_index_12 = (hx_remove_index_12 + 1)
		}
		return false
	}())
	hxrt.Println(v_1)
	var v_2 any = any(tokens.Len())
	hxrt.Println(v_2)
	var v_3 any = any(same(a, b))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringEqualStringPtr(func(hx_value_14 any) *string {
		if hx_value_14 == nil {
			var hx_zero_15 *string
			return hx_zero_15
		}
		return hx_value_14.(*string)
	}(New_Box(a).value), b))
	hxrt.Println(v_4)
	maybe := b
	hxrt.Println(any(hxrt.StringEqualStringPtr(a, maybe)))
	var absent *string = nil
	hxrt.Println(any((absent == nil)))
	hxrt.Println(any(hxrt.StringEqualStringPtr(a, absent)))
	var record_token *string
	record_token = a
	var v_5 any = any(hxrt.StringEqualStringPtr(record_token, b))
	hxrt.Println(v_5)
	read := hxrt.NewFunction(func() *string {
		return a
	})
	var v_6 any = any(hxrt.StringEqualStringPtr(func() *string {
		var hx_callable_16 *hxrt.Function[func() *string] = read
		return hxrt.FunctionCall(hx_callable_16)()
	}(), b))
	hxrt.Println(v_6)
	var dynamicToken any = a
	hxrt.Println(any(hxrt.HaxeEqual(dynamicToken, b)))
}

func same(left any, right any) bool {
	return hxrt.HaxeEqual(left, right)
}

func _Main__Token_Impl___new(raw *string) *string {
	var this1 *string
	this1 = raw
	return this1
}
