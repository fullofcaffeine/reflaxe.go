package main

import "snapshot/hxrt"

type Child interface {
	f(x int) int
}

type I_Concrete interface {
	f(x int) int
}

type Concrete struct {
	__hx_this I_Concrete
}

func (self *Concrete) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Concrete() *Concrete {
	self := &Concrete{}
	self.__hx_this = self
	return self
}

func (self *Concrete) f(x int) int {
	return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(1)))
}

func main() {
	var nominal Child = nil
	view := func(hx_structural_source_1 Child) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "f":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func(hx_arg_0 int) int {
						return hxrt.FunctionReceiver(hx_receiver_2).(Child).f(hx_arg_0)
					}, func() any {
						return hxrt.FunctionReceiver(hx_receiver_2)
					}, "f")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(nominal)
	hxrt.Println(any((view == nil)))
	var instance Child = New_Concrete()
	live := func(hx_structural_source_3 Child) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
			switch name {
			case "f":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
					hx_receiver_4 := hx_structural_source_3
					return hxrt.BoundFunction(func(hx_arg_0 int) int {
						return hxrt.FunctionReceiver(hx_receiver_4).(Child).f(hx_arg_0)
					}, func() any {
						return hxrt.FunctionReceiver(hx_receiver_4)
					}, "f")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(instance)
	var v any = any(func() int {
		var hx_callable_8 *hxrt.Function[func(int) int] = func(hx_obj_5 hxrt.Anonymous) *hxrt.Function[func(int) int] {
			hx_field_6 := hxrt.AnonymousGet(hx_obj_5, "f")
			if hx_field_6 == nil {
				var hx_zero_7 *hxrt.Function[func(int) int]
				return hx_zero_7
			}
			return hx_field_6.(*hxrt.Function[func(int) int])
		}(live)
		var hx_argument_9 int = 4
		return hxrt.FunctionCall(hx_callable_8)(hx_argument_9)
	}())
	hxrt.Println(v)
	saved := func(hx_obj_10 hxrt.Anonymous) *hxrt.Function[func(int) int] {
		hx_field_11 := hxrt.AnonymousGet(hx_obj_10, "f")
		if hx_field_11 == nil {
			var hx_zero_12 *hxrt.Function[func(int) int]
			return hx_zero_12
		}
		return hx_field_11.(*hxrt.Function[func(int) int])
	}(live)
	var v_1 any = any(func() int {
		var hx_callable_13 *hxrt.Function[func(int) int] = saved
		var hx_argument_14 int = 5
		return hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
	}())
	hxrt.Println(v_1)
}

type Middle interface {
	f(x any) any
}

type Parent interface {
	f(x any) any
}
