package main

import "snapshot/hxrt"

type I_Box interface {
	keep(f *hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]
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

func (self *Box) keep(f *hxrt.Function[func(any) any]) *hxrt.Function[func(any) any] {
	return f
}

func main() {
	view := func(hx_structural_source_1 *Box) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "keep":
				return hxrt.AnonymousField{Value: func(hx_callback_source_8 *hxrt.Function[func(*hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]]) *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]] {
					return hxrt.AdaptFunction(hx_callback_source_8, func(hx_erased_callback_arg_3 *hxrt.Function[func(int) int]) *hxrt.Function[func(int) int] {
						return func(hx_callback_source_10 *hxrt.Function[func(any) any]) *hxrt.Function[func(int) int] {
							return hxrt.AdaptFunction(hx_callback_source_10, func(hx_erased_callback_arg_9 int) int {
								return func(hx_value_11 any) int {
									if hx_value_11 == nil {
										var hx_zero_12 int
										return hx_zero_12
									}
									return hx_value_11.(int)
								}(hxrt.FunctionCall(hx_callback_source_10)(hx_erased_callback_arg_9))
							})
						}(hxrt.FunctionCall(hx_callback_source_8)(func(hx_callback_source_7 *hxrt.Function[func(int) int]) *hxrt.Function[func(any) any] {
							return hxrt.AdaptFunction(hx_callback_source_7, func(hx_erased_callback_arg_4 any) any {
								return hxrt.FunctionCall(hx_callback_source_7)(func(hx_value_5 any) int {
									if hx_value_5 == nil {
										var hx_zero_6 int
										return hx_zero_6
									}
									return hx_value_5.(int)
								}(hx_erased_callback_arg_4))
							})
						}(hx_erased_callback_arg_3)))
					})
				}(func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func(f *hxrt.Function[func(any) any]) *hxrt.Function[func(any) any] {
						return hx_receiver_2.__hx_this.keep(f)
					}, func() any {
						return hx_receiver_2.__hx_this
					}, "keep")
				}()), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Box())
	f := func() *hxrt.Function[func(int) int] {
		var hx_callable_16 *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]] = func(hx_obj_13 hxrt.Anonymous) *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]] {
			hx_field_14 := hxrt.AnonymousGet(hx_obj_13, "keep")
			if hx_field_14 == nil {
				var hx_zero_15 *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]]
				return hx_zero_15
			}
			return hx_field_14.(*hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]])
		}(view)
		var hx_argument_17 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(x int) int {
			return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(1)))
		})
		return hxrt.FunctionCall(hx_callable_16)(hx_argument_17)
	}()
	var v any = any(func() int {
		var hx_callable_18 *hxrt.Function[func(int) int] = f
		var hx_argument_19 int = 2
		return hxrt.FunctionCall(hx_callable_18)(hx_argument_19)
	}())
	hxrt.Println(v)
}
