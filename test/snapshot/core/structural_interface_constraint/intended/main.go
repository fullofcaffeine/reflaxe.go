package main

import "snapshot/hxrt"

type I_Cursor interface {
	next() int
}

type Cursor struct {
	__hx_this I_Cursor
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
	return self
}

func (self *Cursor) next() int {
	return 3
}

type HasNext interface {
	next() int
}

func main() {
	source := New_Cursor()
	var nominal HasNext = source
	narrow := func(hx_structural_source_1 HasNext) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func() int {
						return hxrt.FunctionReceiver(hx_receiver_2).(HasNext).next()
					}, func() any {
						return hxrt.FunctionReceiver(hx_receiver_2)
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(nominal)
	var v any = any(func() int {
		var hx_callable_3 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(narrow, "next"))
		return hxrt.FunctionCall(hx_callable_3)()
	}())
	hxrt.Println(v)
	var v_1 any = any(read(func() hxrt.Anonymous {
		hx_obj_4 := map[string]any{}
		hx_obj_4["next"] = hxrt.NewFunction(func() int {
			return 8
		})
		return hxrt.AnonymousFromMap(hx_obj_4)
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(read(func(hx_structural_source_5 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_5, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_6 := hx_structural_source_5
					return hxrt.BoundFunction(func() int {
						return hx_receiver_6.__hx_this.next()
					}, func() any {
						return hx_receiver_6.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(source)))
	hxrt.Println(v_2)
	var erased any = source
	recovered := hx_structural_recover(hx_structural_recover(erased))
	var v_3 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(recovered, "next"))
		return hxrt.FunctionCall(hx_callable_7)()
	}())
	hxrt.Println(v_3)
	hxrt.Println(any(hxrt.HaxeEqual(any(recovered), erased)))
	var missing any = nil
	nullView := hx_structural_recover(hx_structural_recover(missing))
	hxrt.Println(any((nullView == nil)))
	mutable := New_MutableCursor()
	var mutableErased any = mutable
	live := hx_structural_recover(hx_structural_recover(mutableErased))
	saved := hx_recover_callable_0(hxrt.AnonymousGet(live, "next"))
	mutable.next = hxrt.NewFunction(func() int {
		return 19
	})
	var v_4 any = any(func() int {
		var hx_callable_8 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_8)()
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() int {
		var hx_callable_9 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(live, "next"))
		return hxrt.FunctionCall(hx_callable_9)()
	}())
	hxrt.Println(v_5)
	hxrt.Println(any(hxrt.HaxeEqual(any(live), mutableErased)))
}

func read(value hxrt.Anonymous) int {
	return func() int {
		var hx_callable_10 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(value, "next"))
		return hxrt.FunctionCall(hx_callable_10)()
	}()
}

type I_MutableCursor interface {
}

type MutableCursor struct {
	__hx_this I_MutableCursor
	next      *hxrt.Function[func() int]
}

func (self *MutableCursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_MutableCursor() *MutableCursor {
	return hx_init_MutableCursor(hx_alloc_MutableCursor(nil))
}

func hx_alloc_MutableCursor(owner any) *MutableCursor {
	self := &MutableCursor{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_MutableCursor)
	}
	return self
}

func hx_init_MutableCursor(self *MutableCursor) *MutableCursor {
	self.next = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() int {
		return 11
	}), func() any {
		return self.__hx_this
	}, "next")
	return self
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Cursor:
		return func(hx_structural_source_1 *Cursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_2.__hx_this.next()
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *MutableCursor:
		return func(hx_structural_source_3 *MutableCursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
				switch name {
				case "next":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_3.next), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "next":
					hx_structural_source_3.next = hx_recover_callable_0(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_4 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_4, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_4.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_4.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_4.high = func(hx_value_5 any) int {
						if hx_value_5 == nil {
							var hx_zero_6 int
							return hx_zero_6
						}
						return hx_value_5.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_4.low = func(hx_value_7 any) int {
						if hx_value_7 == nil {
							var hx_zero_8 int
							return hx_zero_8
						}
						return hx_value_7.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
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
