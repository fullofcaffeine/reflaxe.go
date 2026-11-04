package main

import "snapshot/hxrt"

type I_Base interface {
	ping() int
}

type Base struct {
	__hx_this I_Base
	n         int
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base(value any) *Base {
	return hx_init_Base(hx_alloc_Base(nil), value)
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

func hx_init_Base(self *Base, value any) *Base {
	self.n = 3
	observed = self.__hx_this.ping()
	return self
}

func (self *Base) ping() int {
	return self.n
}

type I_Child interface {
	ping() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child() *Child {
	return hx_init_Child(hx_alloc_Child(nil))
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

func hx_init_Child(self *Child) *Child {
	hx_init_Base(self.Base, remember(self))
	return self
}

func (self *Child) ping() int {
	return int((hxrt.Int32Wrap(self.n) + hxrt.Int32Wrap(2)))
}

var keys *haxe__ds__ObjectMap = New_haxe__ds__ObjectMap()

func main() {
	child := New_Child()
	var v any = any(hxrt.HaxeEqual(saved, child))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(keys.__hx_this.exists(child)))
	hxrt.Println(v_1)
	var v_2 any = any(func(hx_value_3 any) any {
		if hx_value_3 == nil {
			return nil
		}
		return hx_value_3.(int)
	}(keys.__hx_this.get(child)))
	hxrt.Println(v_2)
	var v_3 any = any(observed)
	hxrt.Println(v_3)
}

var observed int

func remember(value any) any {
	saved = value
	keys.__hx_this.set(value, 7)
	return value
}

var saved any

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
