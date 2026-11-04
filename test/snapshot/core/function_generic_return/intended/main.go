package main

import "snapshot/hxrt"

type I_Holder interface {
	reader() *hxrt.Function[func() any]
	savedReader() *hxrt.Function[func() any]
	echo() *hxrt.Function[func(any) any]
	missing() *hxrt.Function[func() any]
}

type Holder struct {
	__hx_this I_Holder
	value     any
	saved     *hxrt.Function[func() any]
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value any) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), value)
}

func hx_alloc_Holder(owner any) *Holder {
	self := &Holder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Holder)
	}
	return self
}

func hx_init_Holder(self *Holder, value any) *Holder {
	self.value = value
	self.saved = hxrt.NewFunction(func() any {
		return value
	})
	return self
}

func (self *Holder) reader() *hxrt.Function[func() any] {
	_gthis := self
	return hxrt.NewFunction(func() any {
		return _gthis.value
	})
}

func (self *Holder) savedReader() *hxrt.Function[func() any] {
	return self.saved
}

func (self *Holder) echo() *hxrt.Function[func(any) any] {
	return hxrt.NewFunction(func(value any) any {
		return value
	})
}

func (self *Holder) missing() *hxrt.Function[func() any] {
	return nil
}

func main() {
	read := hx_recover_callable_0(New_Holder(19).reader())
	alias := read
	var v any = any(func() int {
		var hx_callable_1 *hxrt.Function[func() int] = read
		return hxrt.FunctionCall(hx_callable_1)()
	}())
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(read), any(alias)))
	hxrt.Println(v_1)
	holder := New_Holder(23)
	var v_2 any = any(func() int {
		var hx_callable_2 *hxrt.Function[func(int) int] = hx_recover_callable_1(holder.echo())
		var hx_argument_3 int = 29
		return hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		var hx_callable_4 *hxrt.Function[func() int] = hx_recover_callable_0(holder.savedReader())
		return hxrt.FunctionCall(hx_callable_4)()
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() bool {
		var f1 any = any(hx_recover_callable_0(holder.savedReader()))
		var f2 any = any(hx_recover_callable_0(holder.savedReader()))
		return hxrt.ReflectCompareMethods(f1, f2)
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() bool {
		var f1_1 any = any(hx_recover_callable_0(holder.reader()))
		var f2_1 any = any(hx_recover_callable_0(holder.reader()))
		return hxrt.ReflectCompareMethods(f1_1, f2_1)
	}())
	hxrt.Println(v_5)
	var v_6 any = any((hx_recover_callable_0(holder.missing()) == nil))
	hxrt.Println(v_6)
	var v_7 any = any(func() *string {
		var hx_callable_5 *hxrt.Function[func() *string] = hx_recover_callable_2(New_Holder(hxrt.StringFromLiteral("text")).reader())
		return hxrt.FunctionCall(hx_callable_5)()
	}())
	hxrt.Println(v_7)
	var v_8 any = any((func() any {
		var hx_callable_6 *hxrt.Function[func() any] = New_Holder(nil).reader()
		return hxrt.FunctionCall(hx_callable_6)()
	}() == nil))
	hxrt.Println(v_8)
}

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func(int) int] {
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
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
