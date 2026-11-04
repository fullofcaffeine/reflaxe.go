package main

import "snapshot/hxrt"

type I_Cursor interface {
}

type Cursor struct {
	__hx_this I_Cursor
	index     int
	hasNext   *hxrt.Function[func() bool]
	next      *hxrt.Function[func() int]
}

func (self *Cursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Cursor() *Cursor {
	self := &Cursor{}
	self.__hx_this = self
	self.hasNext = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() bool {
		return (self.index < 1)
	}), func() any {
		return self.__hx_this
	}, "hasNext")
	self.next = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() int {
		hx_post_1 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_1
	}), func() any {
		return self.__hx_this
	}, "next")
	self.index = 0
	return self
}

func main() {
	source := New_Cursor()
	view := func(hx_structural_source_2 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_2.hasNext), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_2.next), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "hasNext":
				hx_structural_source_2.hasNext = func(hx_value_3 any) *hxrt.Function[func() bool] {
					if hx_value_3 == nil {
						var hx_zero_4 *hxrt.Function[func() bool]
						return hx_zero_4
					}
					return hx_value_3.(*hxrt.Function[func() bool])
				}(incoming)
				return true
			case "next":
				hx_structural_source_2.next = func(hx_value_5 any) *hxrt.Function[func() int] {
					if hx_value_5 == nil {
						var hx_zero_6 *hxrt.Function[func() int]
						return hx_zero_6
					}
					return hx_value_5.(*hxrt.Function[func() int])
				}(incoming)
				return true
			}
			return false
		})
	}(source)
	var v any = any(func() bool {
		var hx_callable_10 *hxrt.Function[func() bool] = func(hx_obj_7 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_8 := hxrt.AnonymousGet(hx_obj_7, "hasNext")
			if hx_field_8 == nil {
				var hx_zero_9 *hxrt.Function[func() bool]
				return hx_zero_9
			}
			return hx_field_8.(*hxrt.Function[func() bool])
		}(view)
		return hxrt.FunctionCall(hx_callable_10)()
	}())
	hxrt.Println(v)
	var v_1 any = any(func() int {
		var hx_callable_14 *hxrt.Function[func() int] = func(hx_obj_11 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_12 := hxrt.AnonymousGet(hx_obj_11, "next")
			if hx_field_12 == nil {
				var hx_zero_13 *hxrt.Function[func() int]
				return hx_zero_13
			}
			return hx_field_12.(*hxrt.Function[func() int])
		}(view)
		return hxrt.FunctionCall(hx_callable_14)()
	}())
	hxrt.Println(v_1)
	var v_2 any = any(func() bool {
		var hx_callable_18 *hxrt.Function[func() bool] = func(hx_obj_15 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_16 := hxrt.AnonymousGet(hx_obj_15, "hasNext")
			if hx_field_16 == nil {
				var hx_zero_17 *hxrt.Function[func() bool]
				return hx_zero_17
			}
			return hx_field_16.(*hxrt.Function[func() bool])
		}(view)
		return hxrt.FunctionCall(hx_callable_18)()
	}())
	hxrt.Println(v_2)
	saved := func(hx_obj_19 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_20 := hxrt.AnonymousGet(hx_obj_19, "next")
		if hx_field_20 == nil {
			var hx_zero_21 *hxrt.Function[func() int]
			return hx_zero_21
		}
		return hx_field_20.(*hxrt.Function[func() int])
	}(view)
	hxrt.Println(any(hxrt.HaxeEqual(saved, func(hx_obj_22 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_23 := hxrt.AnonymousGet(hx_obj_22, "next")
		if hx_field_23 == nil {
			var hx_zero_24 *hxrt.Function[func() int]
			return hx_zero_24
		}
		return hx_field_23.(*hxrt.Function[func() int])
	}(view))))
	source.next = hxrt.NewFunction(func() int {
		return 23
	})
	hxrt.Println(any(hxrt.HaxeEqual(saved, func(hx_obj_25 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_26 := hxrt.AnonymousGet(hx_obj_25, "next")
		if hx_field_26 == nil {
			var hx_zero_27 *hxrt.Function[func() int]
			return hx_zero_27
		}
		return hx_field_26.(*hxrt.Function[func() int])
	}(view))))
	var v_3 any = any(func() int {
		var hx_callable_28 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_28)()
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		var hx_callable_32 *hxrt.Function[func() int] = func(hx_obj_29 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_30 := hxrt.AnonymousGet(hx_obj_29, "next")
			if hx_field_30 == nil {
				var hx_zero_31 *hxrt.Function[func() int]
				return hx_zero_31
			}
			return hx_field_30.(*hxrt.Function[func() int])
		}(view)
		return hxrt.FunctionCall(hx_callable_32)()
	}())
	hxrt.Println(v_4)
}
