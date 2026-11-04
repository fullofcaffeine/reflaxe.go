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
	return hx_init_Main(hx_alloc_Main(nil))
}

func hx_alloc_Main(owner any) *Main {
	self := &Main{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Main)
	}
	return self
}

func hx_init_Main(self *Main) *Main {
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
				hx_structural_source_1.next = hx_recover_callable_0(incoming)
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
	c := hx_recover_callable_0(hxrt.AnonymousGet(view, "next"))
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
		var hx_callable_2 *hxrt.Function[func() int] = a
		return hxrt.FunctionCall(hx_callable_2)()
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() int {
		var hx_callable_3 *hxrt.Function[func() int] = d
		return hxrt.FunctionCall(hx_callable_3)()
	}())
	hxrt.Println(v_5)
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
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
