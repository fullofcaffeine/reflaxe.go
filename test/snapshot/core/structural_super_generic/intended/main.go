package main

import "snapshot/hxrt"

type I_Base interface {
	read() any
}

type Base struct {
	__hx_this I_Base
	it        hxrt.Anonymous
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base(it hxrt.Anonymous) *Base {
	self := &Base{}
	self.__hx_this = self
	self.it = it
	return self
}

func (self *Base) read() any {
	return func() any {
		var hx_callable_4 *hxrt.Function[func() any] = func(hx_obj_1 hxrt.Anonymous) *hxrt.Function[func() any] {
			hx_field_2 := hxrt.AnonymousGet(hx_obj_1, "next")
			if hx_field_2 == nil {
				var hx_zero_3 *hxrt.Function[func() any]
				return hx_zero_3
			}
			return hx_field_2.(*hxrt.Function[func() any])
		}(self.it)
		return hxrt.FunctionCall(hx_callable_4)()
	}()
}

type I_Child interface {
	read() any
}

type Child struct {
	*Base
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child(it hxrt.Anonymous) *Child {
	self := &Child{}
	self.Base = New_Base(func(hx_structural_origin_5 hxrt.Anonymous) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_origin_5, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_9 *hxrt.Function[func() int]) *hxrt.Function[func() any] {
					return hxrt.AdaptFunction(hx_callback_source_9, func() any {
						return hxrt.FunctionCall(hx_callback_source_9)()
					})
				}(func(hx_obj_6 hxrt.Anonymous) *hxrt.Function[func() int] {
					hx_field_7 := hxrt.AnonymousGet(hx_obj_6, "next")
					if hx_field_7 == nil {
						var hx_zero_8 *hxrt.Function[func() int]
						return hx_zero_8
					}
					return hx_field_7.(*hxrt.Function[func() int])
				}(hx_structural_origin_5)), Found: true}
			}
			return hxrt.AnonymousLookup(hx_structural_origin_5, name)
		}, func(name string, incoming any) bool {
			switch name {
			case "next":
				hxrt.AnonymousSet(hx_structural_origin_5, "next", func(hx_callback_source_12 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
					return hxrt.AdaptFunction(hx_callback_source_12, func() int {
						return func(hx_value_13 any) int {
							if hx_value_13 == nil {
								var hx_zero_14 int
								return hx_zero_14
							}
							return hx_value_13.(int)
						}(hxrt.FunctionCall(hx_callback_source_12)())
					})
				}(func(hx_value_10 any) *hxrt.Function[func() any] {
					if hx_value_10 == nil {
						var hx_zero_11 *hxrt.Function[func() any]
						return hx_zero_11
					}
					return hx_value_10.(*hxrt.Function[func() any])
				}(incoming)))
				return true
			}
			hxrt.AnonymousSet(hx_structural_origin_5, name, incoming)
			return true
		})
	}(it))
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func main() {
	hx_obj_15 := map[string]any{}
	hx_obj_15["hasNext"] = hxrt.NewFunction(func() bool {
		return true
	})
	hx_obj_15["next"] = hxrt.NewFunction(func() int {
		return 3
	})
	source := hxrt.AnonymousFromMap(hx_obj_15)
	child := New_Child(source)
	var v any = any(func(hx_value_16 any) int {
		if hx_value_16 == nil {
			var hx_zero_17 int
			return hx_zero_17
		}
		return hx_value_16.(int)
	}(child.read()))
	hxrt.Println(v)
	func() *hxrt.Function[func() int] {
		hx_structural_target_18 := source
		hx_structural_value_20 := hxrt.NewFunction(func() int {
			return 9
		})
		hxrt.AnonymousSet(hx_structural_target_18, "next", hx_structural_value_20)
		return hx_structural_value_20
	}()
	var v_1 any = any(func(hx_value_21 any) int {
		if hx_value_21 == nil {
			var hx_zero_22 int
			return hx_zero_22
		}
		return hx_value_21.(int)
	}(child.read()))
	hxrt.Println(v_1)
}
