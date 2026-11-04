package main

import "snapshot/hxrt"

type I_Box interface {
	read(value any) any
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

func (self *Box) read(value any) any {
	return nil
}

func main() {
	view := func(hx_structural_source_1 *Box) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "read":
				return hxrt.AnonymousField{Value: func(hx_callback_source_4 *hxrt.Function[func(any) any]) *hxrt.Function[func(int) any] {
					return hxrt.AdaptFunction(hx_callback_source_4, func(hx_erased_callback_arg_3 int) any {
						return hxrt.FunctionCall(hx_callback_source_4)(hx_erased_callback_arg_3)
					})
				}(func() *hxrt.Function[func(any) any] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func(value any) any {
						return hx_receiver_2.__hx_this.read(value)
					}, func() any {
						return hx_receiver_2.__hx_this
					}, "read")
				}()), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Box())
	var v any = any((func() any {
		var hx_callable_8 *hxrt.Function[func(int) any] = func(hx_obj_5 hxrt.Anonymous) *hxrt.Function[func(int) any] {
			hx_field_6 := hxrt.AnonymousGet(hx_obj_5, "read")
			if hx_field_6 == nil {
				var hx_zero_7 *hxrt.Function[func(int) any]
				return hx_zero_7
			}
			return hx_field_6.(*hxrt.Function[func(int) any])
		}(view)
		var hx_argument_9 int = 1
		return hxrt.FunctionCall(hx_callable_8)(hx_argument_9)
	}() == nil))
	hxrt.Println(v)
}
