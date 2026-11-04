package main

import "snapshot/hxrt"

type I_Main interface {
}

type Main struct {
	__hx_this I_Main
	next      *hxrt.Function[func() int]
}

func (self *Main) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Main() *Main {
	self := &Main{}
	self.__hx_this = self
	self.next = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() int {
		return 1
	}), func() any {
		return self.__hx_this
	}, "next")
	return self
}

func main() {
	source := New_Main()
	view := func(hx_structural_source_1 *Main) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_1.next), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "next":
				hx_structural_source_1.next = func(hx_value_2 any) *hxrt.Function[func() int] {
					if hx_value_2 == nil {
						var hx_zero_3 *hxrt.Function[func() int]
						return hx_zero_3
					}
					return hx_value_2.(*hxrt.Function[func() int])
				}(incoming)
				return true
			}
			return false
		})
	}(source)
	a := hxrt.ReadDynamicFunction(source.next)
	b := hxrt.ReadDynamicFunction(source.next)
	hxrt.Println(any(hxrt.HaxeEqual(a, b)))
	var v any = any(hxrt.ReflectCompareMethods(any(a), any(b)))
	hxrt.Println(v)
	c := func(hx_obj_4 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_5 := hxrt.AnonymousGet(hx_obj_4, "next")
		if hx_field_5 == nil {
			var hx_zero_6 *hxrt.Function[func() int]
			return hx_zero_6
		}
		return hx_field_5.(*hxrt.Function[func() int])
	}(view)
	hxrt.Println(any(hxrt.HaxeEqual(a, c)))
	var v_1 any = any(hxrt.ReflectCompareMethods(any(a), any(c)))
	hxrt.Println(v_1)
	replacement := hxrt.NewFunction(func() int {
		return 2
	})
	source.next = replacement
	d := hxrt.ReadDynamicFunction(source.next)
	hxrt.Println(any(hxrt.HaxeEqual(d, replacement)))
	var v_2 any = any(hxrt.HaxeEqual(d, hxrt.ReadDynamicFunction(source.next)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.ReflectCompareMethods(any(d), any(replacement)))
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func() int] = a
		return hxrt.FunctionCall(hx_callable_7)()
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() int {
		var hx_callable_8 *hxrt.Function[func() int] = d
		return hxrt.FunctionCall(hx_callable_8)()
	}())
	hxrt.Println(v_5)
}
