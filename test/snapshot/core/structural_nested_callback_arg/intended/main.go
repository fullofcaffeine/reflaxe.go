package main

import "snapshot/hxrt"

type I_Box interface {
	apply(f *hxrt.Function[func(any) any], value any) any
}

type Box struct {
	__hx_this I_Box
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box() *Box {
	self := &Box{}
	self.__hx_this = self
	return self
}

func (self *Box) apply(f *hxrt.Function[func(any) any], value any) any {
	return func() any {
		var hx_callable_1 *hxrt.Function[func(any) any] = f
		var hx_argument_2 any = value
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

func main() {
	view := func(hx_structural_source_3 *Box) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
			switch name {
			case "apply":
				return hxrt.AnonymousField{Value: func(hx_callback_source_11 *hxrt.Function[func(*hxrt.Function[func(any) any], any) any]) *hxrt.Function[func(*hxrt.Function[func(int) int], int) int] {
					return hxrt.AdaptFunction(hx_callback_source_11, func(hx_erased_callback_arg_5 *hxrt.Function[func(int) int], hx_erased_callback_arg_10 int) int {
						return func(hx_value_12 any) int {
							if hx_value_12 == nil {
								var hx_zero_13 int
								return hx_zero_13
							}
							return hx_value_12.(int)
						}(hxrt.FunctionCall(hx_callback_source_11)(func(hx_callback_source_9 *hxrt.Function[func(int) int]) *hxrt.Function[func(any) any] {
							return hxrt.AdaptFunction(hx_callback_source_9, func(hx_erased_callback_arg_6 any) any {
								return hxrt.FunctionCall(hx_callback_source_9)(func(hx_value_7 any) int {
									if hx_value_7 == nil {
										var hx_zero_8 int
										return hx_zero_8
									}
									return hx_value_7.(int)
								}(hx_erased_callback_arg_6))
							})
						}(hx_erased_callback_arg_5), hx_erased_callback_arg_10))
					})
				}(func() *hxrt.Function[func(*hxrt.Function[func(any) any], any) any] {
					hx_receiver_4 := hx_structural_source_3
					return hxrt.BoundFunction(func(f *hxrt.Function[func(any) any], value any) any {
						return hx_receiver_4.__hx_this.apply(f, value)
					}, func() any {
						return hx_receiver_4.__hx_this
					}, "apply")
				}()), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Box())
	var v any = any(func() int {
		var hx_callable_17 *hxrt.Function[func(*hxrt.Function[func(int) int], int) int] = func(hx_obj_14 hxrt.Anonymous) *hxrt.Function[func(*hxrt.Function[func(int) int], int) int] {
			hx_field_15 := hxrt.AnonymousGet(hx_obj_14, "apply")
			if hx_field_15 == nil {
				var hx_zero_16 *hxrt.Function[func(*hxrt.Function[func(int) int], int) int]
				return hx_zero_16
			}
			return hx_field_15.(*hxrt.Function[func(*hxrt.Function[func(int) int], int) int])
		}(view)
		var hx_argument_18 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(x int) int {
			return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(1)))
		})
		var hx_argument_19 int = 2
		return hxrt.FunctionCall(hx_callable_17)(hx_argument_18, hx_argument_19)
	}())
	hxrt.Println(v)
}
