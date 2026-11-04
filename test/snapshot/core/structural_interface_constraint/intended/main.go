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
	self := &Cursor{}
	self.__hx_this = self
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
		var hx_callable_6 *hxrt.Function[func() int] = func(hx_obj_3 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_4 := hxrt.AnonymousGet(hx_obj_3, "next")
			if hx_field_4 == nil {
				var hx_zero_5 *hxrt.Function[func() int]
				return hx_zero_5
			}
			return hx_field_4.(*hxrt.Function[func() int])
		}(narrow)
		return hxrt.FunctionCall(hx_callable_6)()
	}())
	hxrt.Println(v)
	var v_1 any = any(read(func() hxrt.Anonymous {
		hx_obj_7 := map[string]any{}
		hx_obj_7["next"] = hxrt.NewFunction(func() int {
			return 8
		})
		return hxrt.AnonymousFromMap(hx_obj_7)
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(read(func(hx_structural_source_8 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_8, func(name string) hxrt.AnonymousField {
			switch name {
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_9 := hx_structural_source_8
					return hxrt.BoundFunction(func() int {
						return hx_receiver_9.__hx_this.next()
					}, func() any {
						return hx_receiver_9.__hx_this
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
	recovered := hx_structural_recover_0(hx_structural_recover_0(erased))
	var v_3 any = any(func() int {
		var hx_callable_13 *hxrt.Function[func() int] = func(hx_obj_10 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_11 := hxrt.AnonymousGet(hx_obj_10, "next")
			if hx_field_11 == nil {
				var hx_zero_12 *hxrt.Function[func() int]
				return hx_zero_12
			}
			return hx_field_11.(*hxrt.Function[func() int])
		}(recovered)
		return hxrt.FunctionCall(hx_callable_13)()
	}())
	hxrt.Println(v_3)
	hxrt.Println(any(hxrt.HaxeEqual(any(recovered), erased)))
	var missing any = nil
	nullView := hx_structural_recover_0(hx_structural_recover_0(missing))
	hxrt.Println(any((nullView == nil)))
	mutable := New_MutableCursor()
	var mutableErased any = mutable
	live := hx_structural_recover_0(hx_structural_recover_0(mutableErased))
	saved := func(hx_obj_14 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_15 := hxrt.AnonymousGet(hx_obj_14, "next")
		if hx_field_15 == nil {
			var hx_zero_16 *hxrt.Function[func() int]
			return hx_zero_16
		}
		return hx_field_15.(*hxrt.Function[func() int])
	}(live)
	mutable.next = hxrt.NewFunction(func() int {
		return 19
	})
	var v_4 any = any(func() int {
		var hx_callable_17 *hxrt.Function[func() int] = saved
		return hxrt.FunctionCall(hx_callable_17)()
	}())
	hxrt.Println(v_4)
	var v_5 any = any(func() int {
		var hx_callable_21 *hxrt.Function[func() int] = func(hx_obj_18 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_19 := hxrt.AnonymousGet(hx_obj_18, "next")
			if hx_field_19 == nil {
				var hx_zero_20 *hxrt.Function[func() int]
				return hx_zero_20
			}
			return hx_field_19.(*hxrt.Function[func() int])
		}(live)
		return hxrt.FunctionCall(hx_callable_21)()
	}())
	hxrt.Println(v_5)
	hxrt.Println(any(hxrt.HaxeEqual(any(live), mutableErased)))
}

func read(value hxrt.Anonymous) int {
	return func() int {
		var hx_callable_25 *hxrt.Function[func() int] = func(hx_obj_22 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_23 := hxrt.AnonymousGet(hx_obj_22, "next")
			if hx_field_23 == nil {
				var hx_zero_24 *hxrt.Function[func() int]
				return hx_zero_24
			}
			return hx_field_23.(*hxrt.Function[func() int])
		}(value)
		return hxrt.FunctionCall(hx_callable_25)()
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
	self := &MutableCursor{}
	self.__hx_this = self
	self.next = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func() int {
		return 11
	}), func() any {
		return self.__hx_this
	}, "next")
	return self
}

func hx_structural_recover_0(value any) hxrt.Anonymous {
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
					hx_structural_source_3.next = func(hx_value_4 any) *hxrt.Function[func() int] {
						if hx_value_4 == nil {
							var hx_zero_5 *hxrt.Function[func() int]
							return hx_zero_5
						}
						return hx_value_4.(*hxrt.Function[func() int])
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
