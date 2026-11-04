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
	self := &Cursor{}
	self.__hx_this = self
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
		var hx_callable_10 *hxrt.Function[func() int] = func(hx_obj_7 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_8 := hxrt.AnonymousGet(hx_obj_7, "next")
			if hx_field_8 == nil {
				var hx_zero_9 *hxrt.Function[func() int]
				return hx_zero_9
			}
			return hx_field_8.(*hxrt.Function[func() int])
		}(narrow)
		return hxrt.FunctionCall(hx_callable_10)()
	}())
	hxrt.Println(v_4)
	var erased any = left
	restored := hxrt.AnonymousFromValue(hxrt.AnonymousFromValue(erased))
	var v_5 any = any(same(restored, source))
	hxrt.Println(v_5)
	var v_6 any = any(func() int {
		var hx_callable_14 *hxrt.Function[func() int] = func(hx_obj_11 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_12 := hxrt.AnonymousGet(hx_obj_11, "next")
			if hx_field_12 == nil {
				var hx_zero_13 *hxrt.Function[func() int]
				return hx_zero_13
			}
			return hx_field_12.(*hxrt.Function[func() int])
		}(restored)
		return hxrt.FunctionCall(hx_callable_14)()
	}())
	hxrt.Println(v_6)
	a := func(hx_obj_15 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_16 := hxrt.AnonymousGet(hx_obj_15, "next")
		if hx_field_16 == nil {
			var hx_zero_17 *hxrt.Function[func() int]
			return hx_zero_17
		}
		return hx_field_16.(*hxrt.Function[func() int])
	}(left)
	b := func(hx_obj_18 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_19 := hxrt.AnonymousGet(hx_obj_18, "next")
		if hx_field_19 == nil {
			var hx_zero_20 *hxrt.Function[func() int]
			return hx_zero_20
		}
		return hx_field_19.(*hxrt.Function[func() int])
	}(left)
	hxrt.Println(any(hxrt.HaxeEqual(a, b)))
	var v_7 any = any(hxrt.ReflectCompareMethods(any(a), any(b)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.ReflectCompareMethods(any(a), any(func() *hxrt.Function[func() int] {
		hx_receiver_21 := source
		return hxrt.BoundFunction(func() int {
			return hx_receiver_21.__hx_this.next()
		}, func() any {
			return hx_receiver_21.__hx_this
		}, "next")
	}())))
	hxrt.Println(v_8)
	absent := func(hx_structural_source_22 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_22, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_23 := hx_structural_source_22
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_23.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_23.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_24 := hx_structural_source_22
					return hxrt.BoundFunction(func() int {
						return hx_receiver_24.__hx_this.next()
					}, func() any {
						return hx_receiver_24.__hx_this
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
