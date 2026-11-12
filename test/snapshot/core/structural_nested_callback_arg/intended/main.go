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
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(any) any], any) any] {
					hx_receiver_4 := hx_structural_source_3
					return hxrt.BoundFunction(func(f *hxrt.Function[func(any) any], value any) any {
						return hx_receiver_4.__hx_this.apply(f, value)
					}, func() any {
						return hx_receiver_4.__hx_this
					}, "apply")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Box())
	var v any = any(func() int {
		var hx_callable_5 *hxrt.Function[func(*hxrt.Function[func(int) int], int) int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "apply"))
		var hx_argument_6 *hxrt.Function[func(int) int] = hxrt.NewFunction(func(x int) int {
			return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(1)))
		})
		var hx_argument_7 int = 2
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6, hx_argument_7)
	}())
	hxrt.Println(v)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(*hxrt.Function[func(int) int], int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(int) int], int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*hxrt.Function[func(any) any], any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *hxrt.Function[func(int) int], arg_1 int) int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(hxrt.FunctionCall(source)(hx_recover_callable_1(arg_0), arg_1))
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
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) any {
			return hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
