package main

import "snapshot/hxrt"

type I_Holder interface {
	read() hxrt.Anonymous
	erase() any
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

func (self *Holder) erase() any {
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
	read := func() *hxrt.Function[func() hxrt.Anonymous] {
		hx_receiver_2 := holder
		return hxrt.BoundFunction(func() hxrt.Anonymous {
			return hx_receiver_2.__hx_this.read()
		}, func() any {
			return hx_receiver_2.__hx_this
		}, "read")
	}()
	var v any = any(func() int {
		var hx_callable_4 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(func() hxrt.Anonymous {
			var hx_callable_3 *hxrt.Function[func() hxrt.Anonymous] = read
			return hxrt.FunctionCall(hx_callable_3)()
		}(), "next"))
		return hxrt.FunctionCall(hx_callable_4)()
	}())
	hxrt.Println(v)
	func() *hxrt.Function[func() int] {
		hx_structural_target_5 := source
		hx_structural_value_7 := hxrt.NewFunction(func() int {
			return 9
		})
		hxrt.AnonymousSet(hx_structural_target_5, "next", hx_structural_value_7)
		return hx_structural_value_7
	}()
	var v_1 any = any(func() int {
		var hx_callable_9 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(func() hxrt.Anonymous {
			var hx_callable_8 *hxrt.Function[func() hxrt.Anonymous] = read
			return hxrt.FunctionCall(hx_callable_8)()
		}(), "next"))
		return hxrt.FunctionCall(hx_callable_9)()
	}())
	hxrt.Println(v_1)
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
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
