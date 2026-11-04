package main

import "snapshot/hxrt"

type I_Holder interface {
	read() hxrt.Anonymous
}

type Holder struct {
	__hx_this I_Holder
	value     hxrt.Anonymous
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value hxrt.Anonymous) *Holder {
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

func hx_init_Holder(self *Holder, value hxrt.Anonymous) *Holder {
	self.value = value
	return self
}

func (self *Holder) read() hxrt.Anonymous {
	return self.value
}

func main() {
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return true
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() int {
		return 3
	})
	source := hxrt.AnonymousFromMap(hx_obj_1)
	holder := New_Holder(source)
	view := holder.read()
	saved := hx_recover_callable_0(hxrt.AnonymousGet(view, "next"))
	var v any = any(same(view, source))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.HaxeEqual(saved, hx_recover_callable_0(hxrt.AnonymousGet(view, "next")))))
	var v_1 any = any(func() int {
		var hx_callable_2 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_2)()
	}())
	hxrt.Println(v_1)
	func() *hxrt.Function[func() int] {
		hx_structural_target_3 := source
		hx_structural_value_5 := hxrt.NewFunction(func() int {
			return 8
		})
		hxrt.AnonymousSet(hx_structural_target_3, "next", hx_structural_value_5)
		return hx_structural_value_5
	}()
	var v_2 any = any(func() int {
		var hx_callable_6 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(view, "next"))
		return hxrt.FunctionCall(hx_callable_6)()
	}())
	hxrt.Println(v_2)
	var v_3 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_7)()
	}())
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.HaxeEqual(hx_recover_callable_0(hxrt.AnonymousGet(view, "next")), hx_recover_callable_0(hxrt.AnonymousGet(source, "next"))))
	hxrt.Println(v_4)
	var missing hxrt.Anonymous = nil
	var v_5 any = any((New_Holder(missing).read() == nil))
	hxrt.Println(v_5)
	hx_obj_8 := map[string]any{}
	hx_obj_8["hasNext"] = hxrt.NewFunction(func() bool {
		return true
	})
	hx_obj_8["next"] = hxrt.NewFunction(func() any {
		return nil
	})
	nullable := hxrt.AnonymousFromMap(hx_obj_8)
	nullableView := New_Holder(nullable).read()
	var v_6 any = any((func() any {
		var hx_callable_9 *hxrt.Function[func() any] = hx_recover_callable_1(hxrt.AnonymousGet(nullableView, "next"))
		return hxrt.FunctionCall(hx_callable_9)()
	}() == nil))
	hxrt.Println(v_6)
	func() *hxrt.Function[func() any] {
		hx_structural_target_10 := nullable
		hx_structural_value_12 := hxrt.NewFunction(func() any {
			return 5
		})
		hxrt.AnonymousSet(hx_structural_target_10, "next", hx_structural_value_12)
		return hx_structural_value_12
	}()
	var v_7 any = any(func() any {
		var hx_callable_13 *hxrt.Function[func() any] = hx_recover_callable_1(hxrt.AnonymousGet(nullableView, "next"))
		return hxrt.FunctionCall(hx_callable_13)()
	}())
	hxrt.Println(v_7)
}

func same(left any, right any) bool {
	return hxrt.HaxeEqual(left, right)
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
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
