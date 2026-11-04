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
	return hx_init_Concrete(hx_alloc_Concrete(nil))
}

func hx_alloc_Concrete(owner any) *Concrete {
	self := &Concrete{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Concrete)
	}
	return self
}

func hx_init_Concrete(self *Concrete) *Concrete {
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
		var hx_callable_5 *hxrt.Function[func(int) int] = hx_recover_callable_0(hxrt.AnonymousGet(live, "f"))
		var hx_argument_6 int = 4
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}())
	hxrt.Println(v)
	saved := hx_recover_callable_0(hxrt.AnonymousGet(live, "f"))
	var v_1 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func(int) int] = saved
		var hx_argument_8 int = 5
		return hxrt.FunctionCall(hx_callable_7)(hx_argument_8)
	}())
	hxrt.Println(v_1)
}

type Middle interface {
	f(x any) any
}

type Parent interface {
	f(x any) any
}

func hx_recover_callable_0(value any) *hxrt.Function[func(int) int] {
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
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
