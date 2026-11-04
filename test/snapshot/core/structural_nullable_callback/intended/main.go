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

func (self *Box) read(value any) any {
	return nil
}

func main() {
	view := func(hx_structural_source_1 *Box) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "read":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func(value any) any {
						return hx_receiver_2.__hx_this.read(value)
					}, func() any {
						return hx_receiver_2.__hx_this
					}, "read")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Box())
	var v any = any((func() any {
		var hx_callable_3 *hxrt.Function[func(int) any] = hx_recover_callable_0(hxrt.AnonymousGet(view, "read"))
		var hx_argument_4 int = 1
		return hxrt.FunctionCall(hx_callable_3)(hx_argument_4)
	}() == nil))
	hxrt.Println(v)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(int) any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) any {
			return hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
