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
	return hx_init_Base(hx_alloc_Base(nil), it)
}

func hx_alloc_Base(owner any) *Base {
	self := &Base{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Base)
	}
	return self
}

func hx_init_Base(self *Base, it hxrt.Anonymous) *Base {
	self.it = it
	return self
}

func (self *Base) read() any {
	return func() any {
		var hx_callable_1 *hxrt.Function[func() any] = hx_recover_callable_0(hxrt.AnonymousGet(self.it, "next"))
		return hxrt.FunctionCall(hx_callable_1)()
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
	return hx_init_Child(hx_alloc_Child(nil), it)
}

func hx_alloc_Child(owner any) *Child {
	self := &Child{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Child)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Child(self *Child, it hxrt.Anonymous) *Child {
	hx_init_Base(self.Base, it)
	return self
}

func main() {
	hx_obj_2 := map[string]any{}
	hx_obj_2["hasNext"] = hxrt.NewFunction(func() bool {
		return true
	})
	hx_obj_2["next"] = hxrt.NewFunction(func() int {
		return 3
	})
	source := hxrt.AnonymousFromMap(hx_obj_2)
	child := New_Child(source)
	var v any = any(func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(child.read()))
	hxrt.Println(v)
	func() *hxrt.Function[func() int] {
		hx_structural_target_5 := source
		hx_structural_value_7 := hxrt.NewFunction(func() int {
			return 9
		})
		hxrt.AnonymousSet(hx_structural_target_5, "next", hx_structural_value_7)
		return hx_structural_value_7
	}()
	var v_1 any = any(func(hx_value_8 any) int {
		if hx_value_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_value_8.(int)
	}(child.read()))
	hxrt.Println(v_1)
}

func hx_recover_callable_0(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
