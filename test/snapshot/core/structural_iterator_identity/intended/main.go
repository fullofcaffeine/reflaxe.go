package main

import "snapshot/hxrt"

type I_Cursor interface {
	hasNext() bool
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

func (self *Cursor) hasNext() bool {
	return false
}

func (self *Cursor) next() int {
	return 7
}

func main() {
	source := New_Cursor()
	left := func(hx_structural_source_1 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_2 := hx_structural_source_1
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_2.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_2.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_3 := hx_structural_source_1
					return hxrt.BoundFunction(func() int {
						return hx_receiver_3.__hx_this.next()
					}, func() any {
						return hx_receiver_3.__hx_this
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
	right := func(hx_structural_source_4 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_4, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_5 := hx_structural_source_4
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_5.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_5.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_6 := hx_structural_source_4
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
	}(source)
	var v any = any(same(left, right))
	hxrt.Println(v)
	var v_1 any = any(same(left, source))
	hxrt.Println(v_1)
	var v_2 any = any(same(source, left))
	hxrt.Println(v_2)
	hxrt.Println(any(hxrt.HaxeEqual(left, right)))
	narrow := left
	var v_3 any = any(same(narrow, source))
	hxrt.Println(v_3)
	var v_4 any = any(func() int {
		var hx_callable_7 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(narrow, "next"))
		return hxrt.FunctionCall(hx_callable_7)()
	}())
	hxrt.Println(v_4)
	var erased any = left
	restored := hx_structural_recover(hx_structural_recover(erased))
	var v_5 any = any(same(restored, source))
	hxrt.Println(v_5)
	var v_6 any = any(func() int {
		var hx_callable_8 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(restored, "next"))
		return hxrt.FunctionCall(hx_callable_8)()
	}())
	hxrt.Println(v_6)
	a := hx_recover_callable_0(hxrt.AnonymousGet(left, "next"))
	b := hx_recover_callable_0(hxrt.AnonymousGet(left, "next"))
	hxrt.Println(any(hxrt.HaxeEqual(a, b)))
	var v_7 any = any(hxrt.ReflectCompareMethods(any(a), any(b)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.ReflectCompareMethods(any(a), any(func() *hxrt.Function[func() int] {
		hx_receiver_9 := source
		return hxrt.BoundFunction(func() int {
			return hx_receiver_9.__hx_this.next()
		}, func() any {
			return hx_receiver_9.__hx_this
		}, "next")
	}())))
	hxrt.Println(v_8)
	absent := func(hx_structural_source_10 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_10, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_11 := hx_structural_source_10
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_11.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_11.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_12 := hx_structural_source_10
					return hxrt.BoundFunction(func() int {
						return hx_receiver_12.__hx_this.next()
					}, func() any {
						return hx_receiver_12.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(missing())
	hxrt.Println(any((absent == nil)))
}

func missing() *Cursor {
	return nil
}

func same(left any, right any) bool {
	return hxrt.HaxeEqual(left, right)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Cursor:
		return func(hx_structural_source_1 *Cursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_2.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_3 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_3.__hx_this.next()
						}, func() any {
							return hx_receiver_3.__hx_this
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
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
