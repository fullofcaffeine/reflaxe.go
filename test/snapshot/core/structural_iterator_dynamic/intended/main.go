package main

import "snapshot/hxrt"

type I_Cursor interface {
}

type Cursor struct {
	__hx_this I_Cursor
	index     int
	hasNext   *hxrt.Function[func() bool]
	next      *hxrt.Function[func() int]
}

func (self *Cursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Cursor() *Cursor {
	return hx_init_Cursor(hx_alloc_Cursor(nil))
}

func hx_alloc_Cursor(owner any) *Cursor {
	self := &Cursor{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Cursor)
	}
	return self
}

func hx_init_Cursor(self *Cursor) *Cursor {
	self.hasNext = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() bool {
		return (self.index < 1)
	}), func() any {
		return self.__hx_this
	}, "hasNext")
	self.next = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() int {
		hx_post_1 := self.index
		self.index = int(int32((self.index + 1)))
		return hx_post_1
	}), func() any {
		return self.__hx_this
	}, "next")
	self.index = 0
	return self
}

func main() {
	source := New_Cursor()
	view := func(hx_structural_source_2 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "index":
				return hxrt.AnonymousField{Value: hx_structural_source_2.index, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_2.hasNext), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_2.next), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "index":
				hx_structural_source_2.index = func(hx_value_3 any) int {
					if hx_value_3 == nil {
						var hx_zero_4 int
						return hx_zero_4
					}
					return hx_value_3.(int)
				}(incoming)
				return true
			case "hasNext":
				hx_structural_source_2.hasNext = hx_recover_callable_0(incoming)
				return true
			case "next":
				hx_structural_source_2.next = hx_recover_callable_1(incoming)
				return true
			}
			return false
		})
	}(source)
	var v any = any(func() bool {
		var hx_callable_5 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(view, "hasNext"))
		return hxrt.FunctionCall(hx_callable_5)()
	}())
	hxrt.Println(v)
	var v_1 any = any(func() int {
		var hx_callable_6 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "next"))
		return hxrt.FunctionCall(hx_callable_6)()
	}())
	hxrt.Println(v_1)
	var v_2 any = any(func() bool {
		var hx_callable_7 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(view, "hasNext"))
		return hxrt.FunctionCall(hx_callable_7)()
	}())
	hxrt.Println(v_2)
	saved := hx_recover_callable_1(hxrt.AnonymousGet(view, "next"))
	hxrt.Println(any(hxrt.HaxeEqual(saved, hx_recover_callable_1(hxrt.AnonymousGet(view, "next")))))
	source.next = hxrt.NewFunction(func() int {
		return 23
	})
	hxrt.Println(any(hxrt.HaxeEqual(saved, hx_recover_callable_1(hxrt.AnonymousGet(view, "next")))))
	var v_3 any = any(func() int {
		var hx_callable_8 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_8)()
	}())
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		var hx_callable_9 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(view, "next"))
		return hxrt.FunctionCall(hx_callable_9)()
	}())
	hxrt.Println(v_4)
}

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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() int] {
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
			return func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
