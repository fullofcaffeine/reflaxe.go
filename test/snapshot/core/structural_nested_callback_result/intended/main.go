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
	return hx_init_Box(hx_alloc_Box(nil))
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

func hx_init_Box(self *Box) *Box {
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
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func(f *hxrt.Function[func(any) any]) *hxrt.Function[func(any) any] {
						return hx_receiver_2.__hx_this.keep(f)
					}, func() any {
						return hx_receiver_2.__hx_this
					}, "keep")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Box())
	f := func() *hxrt.Function[func(int) int] {
		var hx_callable_3 *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]] = hx_recover_callable_0(hxrt.AnonymousGet(view, "keep"))
		var hx_argument_4 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(x int) int {
			return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(1)))
		})
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}()
	var v any = any(func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = f
		var hx_argument_6 int = 2
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}())
	hxrt.Println(v)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(int) int]) *hxrt.Function[func(int) int]]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(int) int]) *hxrt.Function[func(int) int] {
			return hx_recover_callable_2(hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(int) int]) *hxrt.Function[func(int) int] {
			return hx_recover_callable_2(hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(any(arg_0))))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func(any) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(any(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(any) any]) *hxrt.Function[func(any) any]]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(any(hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0))))
		})
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
