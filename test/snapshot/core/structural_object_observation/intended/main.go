package main

import (
	"reflect"
	"snapshot/hxrt"
)

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Cursor interface {
	hasNext() bool
	next() int
	toString() *string
}

type Cursor struct {
	__hx_this I_Cursor
	tag       *string
	position  int
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
	self.position = 0
	self.tag = hxrt.StringFromLiteral("cursor")
	return self
}

func (self *Cursor) hasNext() bool {
	return false
}

func (self *Cursor) next() int {
	hx_post_1 := self.position
	self.position = int(int32((self.position + 1)))
	return hx_post_1
}

func (self *Cursor) toString() *string {
	return self.tag
}

func (self *Cursor) String() string {
	return *self.toString()
}

func main() {
	source := New_Cursor()
	view := func(hx_structural_source_2 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "tag":
				return hxrt.AnonymousField{Value: hx_structural_source_2.tag, Found: true}
			case "position":
				return hxrt.AnonymousField{Value: hx_structural_source_2.position, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_7 := hx_structural_source_2
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_7.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_7.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_8 := hx_structural_source_2
					return hxrt.BoundFunction(func() int {
						return hx_receiver_8.__hx_this.next()
					}, func() any {
						return hx_receiver_8.__hx_this
					}, "next")
				}(), Found: true}
			case "toString":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
					hx_receiver_9 := hx_structural_source_2
					return hxrt.BoundFunction(func() *string {
						return hx_receiver_9.__hx_this.toString()
					}, func() any {
						return hx_receiver_9.__hx_this
					}, "toString")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "tag":
				hx_structural_source_2.tag = func(hx_value_3 any) *string {
					if hx_value_3 == nil {
						var hx_zero_4 *string
						return hx_zero_4
					}
					return hx_value_3.(*string)
				}(incoming)
				return true
			case "position":
				hx_structural_source_2.position = func(hx_value_5 any) int {
					if hx_value_5 == nil {
						var hx_zero_6 int
						return hx_zero_6
					}
					return hx_value_5.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(source)
	var v any = Reflect_field(view, hxrt.StringFromLiteral("tag"))
	hxrt.Println(v)
	var v_1 any = any(Reflect_hasField(view, hxrt.StringFromLiteral("tag")))
	hxrt.Println(v_1)
	names := Reflect_fields(view)
	hxrt.ArraySort(names, func(hx_comparator_12 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_12, func(hx_cmp_left_10 any, hx_cmp_right_11 any) int {
			return hxrt.FunctionCall(hx_comparator_12)(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hx_cmp_left_10), func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hx_cmp_right_11))
		})
	}(hx_recover_callable_0(hxrt.NamedFunction(Reflect_compare, "Reflect.compare"))))
	var v_2 any = any(hxrt.StringJoinAny(names.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StdString(view))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringEqualStringPtr(Type_getClassName(Type_getClass(view)), Type_getClassName(Type_getClass(source))))
	hxrt.Println(v_4)
	var v_5 any = any(hx_is_nominal_0(view))
	hxrt.Println(v_5)
	keys := New_haxe__ds__ObjectMap()
	keys.__hx_this.set(source, 11)
	var v_6 any = any(func(hx_value_17 any) any {
		if hx_value_17 == nil {
			return nil
		}
		return hx_value_17.(int)
	}(keys.__hx_this.get(view)))
	hxrt.Println(v_6)
	var v_7 any = any(func() bool {
		var f1 any = Reflect_field(view, hxrt.StringFromLiteral("next"))
		return hxrt.ReflectCompareMethods(f1, any(func() *hxrt.Function[func() int] {
			hx_receiver_18 := source
			return hxrt.BoundFunction(func() int {
				return hx_receiver_18.__hx_this.next()
			}, func() any {
				return hx_receiver_18.__hx_this
			}, "next")
		}()))
	}())
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringEqualStringPtr(func() *string {
		var space *string = nil
		return hxrt.StdString(hxrt.JsonStringify(any(view), space))
	}(), func() *string {
		var space_1 *string = nil
		return hxrt.StdString(hxrt.JsonStringify(any(source), space_1))
	}()))
	hxrt.Println(v_8)
	copied := hx_structural_recover(Reflect_copy(view))
	var v_9 any = any(same(copied, source))
	hxrt.Println(v_9)
	var v_10 any = Reflect_field(copied, hxrt.StringFromLiteral("tag"))
	hxrt.Println(v_10)
	Reflect_setField(copied, hxrt.StringFromLiteral("tag"), hxrt.StringFromLiteral("copy"))
	var v_11 any = any(source.tag)
	hxrt.Println(v_11)
	callback := hxrt.NewFunction(func() int {
		return 11
	})
	hx_obj_19 := map[string]any{}
	hx_obj_19["tag"] = hxrt.StringFromLiteral("extra")
	hx_obj_19["hasNext"] = hxrt.NewFunction(func() bool {
		return false
	})
	hx_obj_19["next"] = callback
	rich := hxrt.AnonymousFromMap(hx_obj_19)
	projected := rich
	var v_12 any = any(same(rich, projected))
	hxrt.Println(v_12)
	var v_13 any = Reflect_field(projected, hxrt.StringFromLiteral("tag"))
	hxrt.Println(v_13)
	clone := hx_structural_recover(Reflect_copy(projected))
	var v_14 any = any(same(clone, projected))
	hxrt.Println(v_14)
	hxrt.Println(any(hxrt.HaxeEqual(hx_recover_callable_1(hxrt.AnonymousGet(clone, "next")), callback)))
	func() *hxrt.Function[func() int] {
		hx_structural_target_20 := rich
		hx_structural_value_22 := hxrt.NewFunction(func() int {
			return 22
		})
		hxrt.AnonymousSet(hx_structural_target_20, "next", hx_structural_value_22)
		return hx_structural_value_22
	}()
	var v_15 any = any(func() int {
		var hx_callable_23 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(projected, "next"))
		return hxrt.FunctionCall(hx_callable_23)()
	}())
	hxrt.Println(v_15)
	var v_16 any = any(func() int {
		var hx_callable_24 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(clone, "next"))
		return hxrt.FunctionCall(hx_callable_24)()
	}())
	hxrt.Println(v_16)
	values := hxrt.NewArray(1, 2)
	cursor := func(hx_structural_source_25 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_25, func(name string) hxrt.AnonymousField {
			switch name {
			case "array":
				return hxrt.AnonymousField{Value: hx_structural_source_25.array, Found: true}
			case "current":
				return hxrt.AnonymousField{Value: hx_structural_source_25.current, Found: true}
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_28 := hx_structural_source_25
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_28.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_28.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
					hx_receiver_29 := hx_structural_source_25
					return hxrt.BoundFunction(func() any {
						return hx_receiver_29.__hx_this.next()
					}, func() any {
						return hx_receiver_29.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "current":
				hx_structural_source_25.current = func(hx_value_26 any) int {
					if hx_value_26 == nil {
						var hx_zero_27 int
						return hx_zero_27
					}
					return hx_value_26.(int)
				}(incoming)
				return true
			}
			return false
		})
	}(New_haxe__iterators__ArrayIterator(values))
	a := hx_recover_callable_1(hxrt.AnonymousGet(cursor, "next"))
	b := hx_recover_callable_1(hxrt.AnonymousGet(cursor, "next"))
	hxrt.Println(any(hxrt.HaxeEqual(a, b)))
	var v_17 any = any(hxrt.ReflectCompareMethods(any(a), any(b)))
	hxrt.Println(v_17)
	var v_18 any = any(func() int {
		var hx_callable_30 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(cursor, "next"))
		return hxrt.FunctionCall(hx_callable_30)()
	}())
	hxrt.Println(v_18)
	cursorFields := Reflect_fields(cursor)
	hxrt.ArraySort(cursorFields, func(hx_comparator_33 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_33, func(hx_cmp_left_31 any, hx_cmp_right_32 any) int {
			return hxrt.FunctionCall(hx_comparator_33)(func(hx_value_34 any) *string {
				if hx_value_34 == nil {
					var hx_zero_35 *string
					return hx_zero_35
				}
				return hx_value_34.(*string)
			}(hx_cmp_left_31), func(hx_value_36 any) *string {
				if hx_value_36 == nil {
					var hx_zero_37 *string
					return hx_zero_37
				}
				return hx_value_36.(*string)
			}(hx_cmp_right_32))
		})
	}(hx_recover_callable_0(hxrt.NamedFunction(Reflect_compare, "Reflect.compare"))))
	var v_19 any = any(hxrt.StringJoinAny(cursorFields.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_19)
	hx_array_target_38 := values
	hx_array_index_39 := 1
	hx_array_target_38.Set(hx_array_index_39, 9)
	var v_20 any = any(func() int {
		var hx_callable_40 *hxrt.Function[func() int] = hx_recover_callable_1(hxrt.AnonymousGet(cursor, "next"))
		return hxrt.FunctionCall(hx_callable_40)()
	}())
	hxrt.Println(v_20)
}

func same(left any, right any) bool {
	return hxrt.HaxeEqual(left, right)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Cursor:
		return func(hx_structural_source_132 *Cursor) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_132, func(name string) hxrt.AnonymousField {
				switch name {
				case "tag":
					return hxrt.AnonymousField{Value: hx_structural_source_132.tag, Found: true}
				case "position":
					return hxrt.AnonymousField{Value: hx_structural_source_132.position, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_137 := hx_structural_source_132
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_137.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_138 := hx_structural_source_132
						return hxrt.BoundFunction(func() int {
							return hx_receiver_138.__hx_this.next()
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "next")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_139 := hx_structural_source_132
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_139.__hx_this.toString()
						}, func() any {
							return hx_receiver_139.__hx_this
						}, "toString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "tag":
					hx_structural_source_132.tag = func(hx_value_133 any) *string {
						if hx_value_133 == nil {
							var hx_zero_134 *string
							return hx_zero_134
						}
						return hx_value_133.(*string)
					}(incoming)
					return true
				case "position":
					hx_structural_source_132.position = func(hx_value_135 any) int {
						if hx_value_135 == nil {
							var hx_zero_136 int
							return hx_zero_136
						}
						return hx_value_135.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_140 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_140, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_140.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_140.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_140.high = func(hx_value_141 any) int {
						if hx_value_141 == nil {
							var hx_zero_142 int
							return hx_zero_142
						}
						return hx_value_141.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_140.low = func(hx_value_143 any) int {
						if hx_value_143 == nil {
							var hx_zero_144 int
							return hx_zero_144
						}
						return hx_value_143.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_145 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_145, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_145.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_145.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_148 := hx_structural_source_145
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_148.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_148.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_149 := hx_structural_source_145
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_149.__hx_this.next()
						}, func() any {
							return hx_receiver_149.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_145.map_ = func(hx_value_146 any) haxe__IMap {
						if hx_value_146 == nil {
							var hx_zero_147 haxe__IMap
							return hx_zero_147
						}
						return hx_value_146.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_145.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__ArrayIterator:
		return func(hx_structural_source_150 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_150, func(name string) hxrt.AnonymousField {
				switch name {
				case "array":
					return hxrt.AnonymousField{Value: hx_structural_source_150.array, Found: true}
				case "current":
					return hxrt.AnonymousField{Value: hx_structural_source_150.current, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_153 := hx_structural_source_150
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_153.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_153.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_154 := hx_structural_source_150
						return hxrt.BoundFunction(func() any {
							return hx_receiver_154.__hx_this.next()
						}, func() any {
							return hx_receiver_154.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "current":
					hx_structural_source_150.current = func(hx_value_151 any) int {
						if hx_value_151 == nil {
							var hx_zero_152 int
							return hx_zero_152
						}
						return hx_value_151.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__ObjectMap:
		return func(hx_structural_source_155 *haxe__ds__ObjectMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_155, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_155.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_158 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_158.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_158.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_159 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_159.__hx_this.get(key)
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_160 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_160.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_161 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_161.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_161.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_162 := hx_structural_source_155
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_162.__hx_this.keys()
						}, func() any {
							return hx_receiver_162.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_163 := hx_structural_source_155
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_163.__hx_this.iterator()
						}, func() any {
							return hx_receiver_163.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_164 := hx_structural_source_155
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_164.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_164.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_165 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_165.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_165.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_166 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_166.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_166.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_167 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_167.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_167.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_168 := hx_structural_source_155
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_168.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_168.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_169 := hx_structural_source_155
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_169.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_169.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__ObjectMap] {
						hx_receiver_170 := hx_structural_source_155
						return hxrt.BoundFunction(func() *haxe__ds__ObjectMap {
							return hx_receiver_170.__hx_this.copy()
						}, func() any {
							return hx_receiver_170.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_171 := hx_structural_source_155
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_171.__hx_this.toString()
						}, func() any {
							return hx_receiver_171.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_172 := hx_structural_source_155
						return hxrt.BoundFunction(func() {
							hx_receiver_172.__hx_this.clear()
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_155.h = func(hx_value_156 any) *hxrt.ObjectMapCell {
						if hx_value_156 == nil {
							var hx_zero_157 *hxrt.ObjectMapCell
							return hx_zero_157
						}
						return hx_value_156.(*hxrt.ObjectMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_173 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_173, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_173.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_173.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_173.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_173.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_176 := hx_structural_source_173
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_176.__hx_this.get_message()
						}, func() any {
							return hx_receiver_176.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_177 := hx_structural_source_173
						return hxrt.BoundFunction(func() any {
							return hx_receiver_177.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_177.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_178 := hx_structural_source_173
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_178.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_178.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_179 := hx_structural_source_173
						return hxrt.BoundFunction(func() any {
							return hx_receiver_179.__hx_this.get_native()
						}, func() any {
							return hx_receiver_179.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_180 := hx_structural_source_173
						return hxrt.BoundFunction(func() any {
							return hx_receiver_180.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_180.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_181 := hx_structural_source_173
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_181.__hx_this.toString()
						}, func() any {
							return hx_receiver_181.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_182 := hx_structural_source_173
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_182.__hx_this.details()
						}, func() any {
							return hx_receiver_182.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_173.savedMessage = func(hx_value_174 any) *string {
						if hx_value_174 == nil {
							var hx_zero_175 *string
							return hx_zero_175
						}
						return hx_value_174.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_173.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_173.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_173.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_183 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_183, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_183.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_184 := hx_structural_source_183
						return hxrt.BoundFunction(func() any {
							return hx_receiver_184.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_184.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_183.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_183.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_183.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_183.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_187 := hx_structural_source_183
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_187.__hx_this.get_message()
						}, func() any {
							return hx_receiver_187.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_188 := hx_structural_source_183
						return hxrt.BoundFunction(func() any {
							return hx_receiver_188.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_188.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_189 := hx_structural_source_183
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_189.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_189.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_190 := hx_structural_source_183
						return hxrt.BoundFunction(func() any {
							return hx_receiver_190.__hx_this.get_native()
						}, func() any {
							return hx_receiver_190.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_191 := hx_structural_source_183
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_191.__hx_this.toString()
						}, func() any {
							return hx_receiver_191.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_192 := hx_structural_source_183
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_192.__hx_this.details()
						}, func() any {
							return hx_receiver_192.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_183.savedMessage = func(hx_value_185 any) *string {
						if hx_value_185 == nil {
							var hx_zero_186 *string
							return hx_zero_186
						}
						return hx_value_185.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_183.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_183.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_183.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Cursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Cursor:
		return hxrt__generated_method_field__Cursor(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_method_field__haxe__ds__ObjectMap(value, key)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_method_field__haxe__iterators__ArrayIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Cursor(value *Cursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_193 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_193.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_193.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_194 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_194.__hx_this.next()
			}, func() any {
				return hx_receiver_194.__hx_this
			}, "next")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_195 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_195.__hx_this.toString()
			}, func() any {
				return hx_receiver_195.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_196 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_196.__hx_this.details()
			}, func() any {
				return hx_receiver_196.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_197 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_197.__hx_this.get_message()
			}, func() any {
				return hx_receiver_197.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_198 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_198.__hx_this.get_native()
			}, func() any {
				return hx_receiver_198.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_199 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_199.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_199.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_200 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_200.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_200.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_201 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_201.__hx_this.toString()
			}, func() any {
				return hx_receiver_201.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_202 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_202.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_202.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_203 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_203.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_203.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_method_field__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_204 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_204.__hx_this.clear()
			}, func() any {
				return hx_receiver_204.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__ObjectMap] {
			hx_receiver_205 := value
			return hxrt.BoundFunction(func() *haxe__ds__ObjectMap {
				return hx_receiver_205.__hx_this.copy()
			}, func() any {
				return hx_receiver_205.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_206 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_206.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_206.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_207 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_207.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_207.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_208 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_208.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_208.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_209 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_209.__hx_this.get(key)
			}, func() any {
				return hx_receiver_209.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_210 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_210.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_210.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_211 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_211.__hx_this.iterator()
			}, func() any {
				return hx_receiver_211.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_212 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_212.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_212.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_213 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_213.__hx_this.keys()
			}, func() any {
				return hx_receiver_213.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_214 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_214.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_214.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_215 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_215.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_215.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_216 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_216.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_216.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_217 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_217.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_217.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_218 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_218.__hx_this.toString()
			}, func() any {
				return hx_receiver_218.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_219 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_219.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_219.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() any] {
			hx_receiver_220 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_220.__hx_this.next()
			}, func() any {
				return hx_receiver_220.__hx_this
			}, "next")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hasNext":
		return func() *hxrt.Function[func() bool] {
			hx_receiver_221 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_221.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_221.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_222 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_222.__hx_this.next()
			}, func() any {
				return hx_receiver_222.__hx_this
			}, "next")
		}()
	}
	return nil
}

type Type struct {
}

func hxrt_typeCreateEmpty__Cursor() *Cursor {
	instance := &Cursor{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__Exception() *haxe__Exception {
	instance := &haxe__Exception{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ValueException() *haxe__ValueException {
	instance := &haxe__ValueException{}
	instance.haxe__Exception = &haxe__Exception{}
	instance.haxe__Exception.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe___Int64_____Int64() *haxe___Int64_____Int64 {
	instance := &haxe___Int64_____Int64{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	instance := &haxe__ds__ObjectMap{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__ArrayIterator() *haxe__iterators__ArrayIterator {
	instance := &haxe__iterators__ArrayIterator{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator() *haxe__iterators__MapKeyValueIterator {
	instance := &haxe__iterators__MapKeyValueIterator{}
	instance.__hx_this = instance
	return instance
}

type ValueType struct {
	tag    int
	params []any
}

var ValueType_TNull *ValueType = &ValueType{tag: 0, params: []any{}}

var ValueType_TInt *ValueType = &ValueType{tag: 1, params: []any{}}

var ValueType_TFloat *ValueType = &ValueType{tag: 2, params: []any{}}

var ValueType_TBool *ValueType = &ValueType{tag: 3, params: []any{}}

var ValueType_TObject *ValueType = &ValueType{tag: 4, params: []any{}}

var ValueType_TFunction *ValueType = &ValueType{tag: 5, params: []any{}}

var ValueType_TUnknown *ValueType = &ValueType{tag: 8, params: []any{}}

func ValueType_TClass(c any) *ValueType {
	return &ValueType{tag: 6, params: []any{c}}
}

func ValueType_TEnum(e any) *ValueType {
	return &ValueType{tag: 7, params: []any{e}}
}

func hxrt_typeCallAny(callable any, args []any) (any, bool) {
	result := any(nil)
	ok := false
	defer func() {
		if recover() != nil {
			result = nil
			ok = false
		}
	}()
	if callable == nil {
		return nil, false
	}
	fn := reflect.ValueOf(callable)
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, false
	}
	fnType := fn.Type()
	if fnType.NumIn() != len(args) {
		return nil, false
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		paramType := fnType.In(i)
		arg := args[i]
		if arg == nil {
			in[i] = reflect.Zero(paramType)
			continue
		}
		v := reflect.ValueOf(arg)
		if v.IsValid() && v.Type().AssignableTo(paramType) {
			in[i] = v
			continue
		}
		if v.IsValid() && v.Type().ConvertibleTo(paramType) {
			in[i] = v.Convert(paramType)
			continue
		}
		if paramType.Kind() == reflect.Interface && v.IsValid() {
			in[i] = v
			continue
		}
		return nil, false
	}
	out := fn.Call(in)
	if len(out) == 0 {
		return nil, true
	}
	first := out[0]
	if !first.IsValid() {
		return nil, true
	}
	result = first.Interface()
	ok = true
	return result, ok
}

func hxrt_typeArrayValues(value *hxrt.Array) []any {
	if value == nil {
		return []any{}
	}
	return value.Values()
}

func hxrt_typeResolvedClassName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeClassValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeClassValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeResolvedEnumName(value any) (string, bool) {
	switch current := value.(type) {
	case *hxrt__TypeEnumValue:
		if current == nil || current.name == nil {
			return "", false
		}
		return *current.name, true
	case hxrt__TypeEnumValue:
		if current.name == nil {
			return "", false
		}
		return *current.name, true
	case string:
		return current, true
	case *string:
		if current == nil {
			return "", false
		}
		return *current, true
	default:
		return "", false
	}
}

func hxrt_typeCreateClassInstance(className string, args []any) (any, bool) {
	switch className {
	case "Cursor":
		return hxrt_typeCallAny(New_Cursor, args)
	case "Main":
		return nil, false
	case "Reflect":
		return nil, false
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Json":
		return nil, false
	case "haxe.ValueException":
		return hxrt_typeCallAny(New_haxe__ValueException, args)
	case "haxe._CallStack.CallStack_Impl_":
		return nil, false
	case "haxe._Int32.Int32_Impl_":
		return nil, false
	case "haxe._Int64.Int64_Impl_":
		return nil, false
	case "haxe._Int64.___Int64":
		return hxrt_typeCallAny(New_haxe___Int64_____Int64, args)
	case "haxe.ds.ObjectMap":
		return hxrt_typeCallAny(New_haxe__ds__ObjectMap, args)
	case "haxe.iterators.ArrayIterator":
		return hxrt_typeCallAny(New_haxe__iterators__ArrayIterator, args)
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCallAny(New_haxe__iterators__MapKeyValueIterator, args)
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "Cursor":
		return hxrt_typeCreateEmpty__Cursor(), true
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	case "haxe.ds.ObjectMap":
		return hxrt_typeCreateEmpty__haxe__ds__ObjectMap(), true
	case "haxe.iterators.ArrayIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__ArrayIterator(), true
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt_typeCreateEmpty__haxe__iterators__MapKeyValueIterator(), true
	default:
		return nil, false
	}
}

func hxrt_typeCreateEnumInstance(enumName string, constructorName string, constructorIndex int, useIndex bool, args []any) (any, bool) {
	switch enumName {
	case "ValueType":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TNull, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TInt, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFloat, true
			case 3:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TBool, true
			case 4:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TObject, true
			case 5:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TFunction, true
			case 6:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TClass, args)
			case 7:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(ValueType_TEnum, args)
			case 8:
				if len(args) != 0 {
					return nil, false
				}
				return ValueType_TUnknown, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "TNull":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TNull, true
		case "TInt":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TInt, true
		case "TFloat":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFloat, true
		case "TBool":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TBool, true
		case "TObject":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TObject, true
		case "TFunction":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TFunction, true
		case "TClass":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TClass, args)
		case "TEnum":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(ValueType_TEnum, args)
		case "TUnknown":
			if len(args) != 0 {
				return nil, false
			}
			return ValueType_TUnknown, true
		default:
			return nil, false
		}
	case "haxe.StackItem":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__StackItem_CFunction, true
			case 1:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Module, args)
			case 2:
				if len(args) != 4 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
			case 3:
				if len(args) != 2 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_Method, args)
			case 4:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "CFunction":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__StackItem_CFunction, true
		case "Module":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Module, args)
		case "FilePos":
			if len(args) != 4 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_FilePos, args)
		case "Method":
			if len(args) != 2 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_Method, args)
		case "LocalFunction":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__StackItem_LocalFunction, args)
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

func Type_getClass(o any) any {
	o = hxrt.ObjectSource(o)
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeClassValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeClassValue:
		copyValue := value
		return &copyValue
	case *hxrt.Array:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")}
	case *Cursor:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Cursor")}
	case *haxe__Exception:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case *haxe__ValueException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ValueException")}
	case *haxe___Int64_____Int64:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe._Int64.___Int64")}
	case *haxe__ds__ObjectMap:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.ds.ObjectMap")}
	case *haxe__iterators__ArrayIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.ArrayIterator")}
	case *haxe__iterators__MapKeyValueIterator:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.iterators.MapKeyValueIterator")}
	default:
		return nil
	}
}

func Type_getEnum(o any) any {
	if hxrt.AnyEqualsNull(o) {
		return nil
	}
	switch value := o.(type) {
	case *hxrt__TypeEnumValue:
		if value == nil {
			return nil
		}
		return value
	case hxrt__TypeEnumValue:
		copyValue := value
		return &copyValue
	case *ValueType:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("ValueType")}
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.StackItem")}
	default:
		return nil
	}
}

func Type_getSuperClass(c any) any {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	switch className {
	case "Cursor":
		return nil
	case "Main":
		return nil
	case "Reflect":
		return nil
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
		return nil
	case "haxe.Json":
		return nil
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe._CallStack.CallStack_Impl_":
		return nil
	case "haxe._Int32.Int32_Impl_":
		return nil
	case "haxe._Int64.Int64_Impl_":
		return nil
	case "haxe._Int64.___Int64":
		return nil
	case "haxe.ds.ObjectMap":
		return nil
	case "haxe.iterators.ArrayIterator":
		return nil
	case "haxe.iterators.MapKeyValueIterator":
		return nil
	default:
		return nil
	}
}

func Type_getClassName(c any) *string {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(className)
}

func Type_getClassFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Cursor":
		return hxrt.NewArray()
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"), hxrt.StringFromLiteral("same"))
	case "Reflect":
		return hxrt.NewArray(hxrt.StringFromLiteral("callMethod"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("compareMethods"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("deleteField"), hxrt.StringFromLiteral("field"), hxrt.StringFromLiteral("fields"), hxrt.StringFromLiteral("getProperty"), hxrt.StringFromLiteral("hasField"), hxrt.StringFromLiteral("isEnumValue"), hxrt.StringFromLiteral("isFunction"), hxrt.StringFromLiteral("isObject"), hxrt.StringFromLiteral("makeVarArgs"), hxrt.StringFromLiteral("setField"), hxrt.StringFromLiteral("setProperty"))
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray(hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("stringify"))
	case "haxe.ValueException":
		return hxrt.NewArray()
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray(hxrt.StringFromLiteral("asArray"), hxrt.StringFromLiteral("callStack"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("equalItems"), hxrt.StringFromLiteral("exceptionStack"), hxrt.StringFromLiteral("exceptionToString"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("get_length"), hxrt.StringFromLiteral("itemToString"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("subtract"), hxrt.StringFromLiteral("toString"))
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray()
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray()
	case "haxe.iterators.ArrayIterator":
		return hxrt.NewArray()
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray()
	default:
		return hxrt.NewArray()
	}
}

func Type_getInstanceFields(c any) *hxrt.Array {
	className, ok := hxrt_typeResolvedClassName(c)
	if !ok {
		return hxrt.NewArray()
	}
	switch className {
	case "Cursor":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"), hxrt.StringFromLiteral("position"), hxrt.StringFromLiteral("tag"), hxrt.StringFromLiteral("toString"))
	case "Main":
		return hxrt.NewArray()
	case "Reflect":
		return hxrt.NewArray()
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray()
	case "haxe.ValueException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"), hxrt.StringFromLiteral("value"))
	case "haxe._CallStack.CallStack_Impl_":
		return hxrt.NewArray()
	case "haxe._Int32.Int32_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.Int64_Impl_":
		return hxrt.NewArray()
	case "haxe._Int64.___Int64":
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case "haxe.ds.ObjectMap":
		return hxrt.NewArray(hxrt.StringFromLiteral("clear"), hxrt.StringFromLiteral("copy"), hxrt.StringFromLiteral("copyIMap"), hxrt.StringFromLiteral("exists"), hxrt.StringFromLiteral("existsIMap"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getIMap"), hxrt.StringFromLiteral("h"), hxrt.StringFromLiteral("iterator"), hxrt.StringFromLiteral("keyValueIterator"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("remove"), hxrt.StringFromLiteral("removeIMap"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setIMap"), hxrt.StringFromLiteral("toString"))
	case "haxe.iterators.ArrayIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("array"), hxrt.StringFromLiteral("current"), hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("next"))
	case "haxe.iterators.MapKeyValueIterator":
		return hxrt.NewArray(hxrt.StringFromLiteral("hasNext"), hxrt.StringFromLiteral("keys"), hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("next"))
	default:
		return hxrt.NewArray()
	}
}

func Type_getEnumName(e any) *string {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	return hxrt.StringFromLiteral(enumName)
}

func Type_resolveClass(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "Cursor":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "Reflect":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Json":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ValueException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._CallStack.CallStack_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int32.Int32_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.Int64_Impl_":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe._Int64.___Int64":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.ds.ObjectMap":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.ArrayIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.iterators.MapKeyValueIterator":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_resolveEnum(name *string) any {
	if name == nil {
		return nil
	}
	rawName := *hxrt.StdString(name)
	switch rawName {
	case "ValueType":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.StackItem":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	default:
		return nil
	}
}

func Type_createInstance(cl any, args *hxrt.Array) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))
	if !ok {
		return nil
	}
	return instance
}

func Type_createEmptyInstance(cl any) any {
	className, ok := hxrt_typeResolvedClassName(cl)
	if !ok {
		return nil
	}
	instance, ok := hxrt_typeCreateClassEmptyInstance(className)
	if !ok {
		return nil
	}
	return instance
}

func Type_createEnum(e any, constr *string, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	constructorName := ""
	if constr != nil {
		constructorName = *hxrt.StdString(constr)
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_createEnumIndex(e any, index int, params *hxrt.Array) any {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return nil
	}
	enumValue, ok := hxrt_typeCreateEnumInstance(enumName, "", index, true, hxrt_typeArrayValues(params))
	if !ok {
		return nil
	}
	return enumValue
}

func Type_enumConstructor(e any) *string {
	if hxrt.AnyEqualsNull(e) {
		return nil
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("TNull")
		case 1:
			return hxrt.StringFromLiteral("TInt")
		case 2:
			return hxrt.StringFromLiteral("TFloat")
		case 3:
			return hxrt.StringFromLiteral("TBool")
		case 4:
			return hxrt.StringFromLiteral("TObject")
		case 5:
			return hxrt.StringFromLiteral("TFunction")
		case 6:
			return hxrt.StringFromLiteral("TClass")
		case 7:
			return hxrt.StringFromLiteral("TEnum")
		case 8:
			return hxrt.StringFromLiteral("TUnknown")
		default:
			return nil
		}
	case *haxe__StackItem:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("CFunction")
		case 1:
			return hxrt.StringFromLiteral("Module")
		case 2:
			return hxrt.StringFromLiteral("FilePos")
		case 3:
			return hxrt.StringFromLiteral("Method")
		case 4:
			return hxrt.StringFromLiteral("LocalFunction")
		default:
			return nil
		}
	default:
		return nil
	}
}

func Type_enumIndex(e any) int {
	if hxrt.AnyEqualsNull(e) {
		return -1
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__StackItem:
		if value == nil {
			return -1
		}
		return value.tag
	default:
		return -1
	}
}

func Type_getEnumConstructs(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(hxrt.StringFromLiteral("TNull"), hxrt.StringFromLiteral("TInt"), hxrt.StringFromLiteral("TFloat"), hxrt.StringFromLiteral("TBool"), hxrt.StringFromLiteral("TObject"), hxrt.StringFromLiteral("TFunction"), hxrt.StringFromLiteral("TClass"), hxrt.StringFromLiteral("TEnum"), hxrt.StringFromLiteral("TUnknown"))
	case "haxe.StackItem":
		return hxrt.NewArray(hxrt.StringFromLiteral("CFunction"), hxrt.StringFromLiteral("Module"), hxrt.StringFromLiteral("FilePos"), hxrt.StringFromLiteral("Method"), hxrt.StringFromLiteral("LocalFunction"))
	default:
		return hxrt.NewArray()
	}
}

func Type_enumParameters(e any) *hxrt.Array {
	if hxrt.AnyEqualsNull(e) {
		return hxrt.NewArray()
	}
	switch value := e.(type) {
	case *ValueType:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__StackItem:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	default:
		return hxrt.NewArray()
	}
}

func Type_allEnums(e any) *hxrt.Array {
	enumName, ok := hxrt_typeResolvedEnumName(e)
	if !ok {
		return hxrt.NewArray()
	}
	switch enumName {
	case "ValueType":
		return hxrt.NewArray(ValueType_TNull, ValueType_TInt, ValueType_TFloat, ValueType_TBool, ValueType_TObject, ValueType_TFunction, ValueType_TUnknown)
	case "haxe.StackItem":
		return hxrt.NewArray(haxe__StackItem_CFunction)
	default:
		return hxrt.NewArray()
	}
}

func Type_typeof(v any) *ValueType {
	v = hxrt.ObjectSource(v)
	if hxrt.ReflectIsFunction(v) {
		return ValueType_TFunction
	}
	if hxrt.AnyEqualsNull(v) {
		return ValueType_TNull
	}
	switch v.(type) {
	case *hxrt__TypeClassValue, hxrt__TypeClassValue, *hxrt__TypeEnumValue, hxrt__TypeEnumValue:
		return ValueType_TObject
	}
	if enumValue := Type_getEnum(v); enumValue != nil {
		return ValueType_TEnum(enumValue)
	}
	if classValue := Type_getClass(v); classValue != nil {
		return ValueType_TClass(classValue)
	}
	switch v.(type) {
	case bool:
		return ValueType_TBool
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return ValueType_TInt
	case float32, float64:
		return ValueType_TFloat
	case string, *string:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")})
	case *hxrt.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	}
	ref := reflect.ValueOf(v)
	if !ref.IsValid() {
		return ValueType_TNull
	}
	switch ref.Kind() {
	case reflect.Slice, reflect.Array:
		return ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral("Array")})
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:
		return ValueType_TObject
	default:
		return ValueType_TUnknown
	}
}

func Type_enumEq(a any, b any) bool {
	return reflect.DeepEqual(a, b)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Cursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Cursor:
		return hxrt__generated_field_lookup__Cursor(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_lookup__haxe__ds__ObjectMap(value, key)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_field_lookup__haxe__iterators__ArrayIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Cursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Cursor:
		return hxrt__generated_field_has__Cursor(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_has__haxe__ds__ObjectMap(value, key)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_field_has__haxe__iterators__ArrayIterator(value, key)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Cursor:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Cursor:
		return hxrt__generated_field_set__Cursor(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	case *haxe__ds__ObjectMap:
		return hxrt__generated_field_set__haxe__ds__ObjectMap(value, key, incoming)
	case *haxe__iterators__ArrayIterator:
		return hxrt__generated_field_set__haxe__iterators__ArrayIterator(value, key, incoming)
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Cursor:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ds__ObjectMap:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__ArrayIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__iterators__MapKeyValueIterator:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Cursor:
		return hxrt.NewArray(hxrt.StringFromLiteral("tag"), hxrt.StringFromLiteral("position"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	case *haxe__ds__ObjectMap:
		return hxrt.NewArray(hxrt.StringFromLiteral("h"))
	case *haxe__iterators__ArrayIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("array"), hxrt.StringFromLiteral("current"))
	case *haxe__iterators__MapKeyValueIterator:
		return hxrt.NewArray(hxrt.StringFromLiteral("map"), hxrt.StringFromLiteral("keys"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Cursor(value *Cursor, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "position":
		return hxrt.ReadStoredField(value.position)
	case "tag":
		return hxrt.ReadStoredField(value.tag)
	}
	return nil
}

func hxrt__generated_field_has__Cursor(value *Cursor, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "position":
		return true
	case "tag":
		return true
	}
	return false
}

func hxrt__generated_field_set__Cursor(value *Cursor, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "position":
		if incoming == nil {
			var zero int
			value.position = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.position = typed
			return true
		default:
			return false
		}
	case "tag":
		if incoming == nil {
			var zero *string
			value.tag = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.tag = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "h":
		return hxrt.ReadStoredField(value.h)
	}
	return nil
}

func hxrt__generated_field_has__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "h":
		if incoming == nil {
			var zero *hxrt.ObjectMapCell
			value.h = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.ObjectMapCell:
			value.h = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "array":
		return hxrt.ReadStoredField(value.array)
	case "current":
		return hxrt.ReadStoredField(value.current)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "array":
		return true
	case "current":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__ArrayIterator(value *haxe__iterators__ArrayIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "array":
		if incoming == nil {
			var zero *hxrt.Array
			value.array = zero
			return true
		}
		switch typed := incoming.(type) {
		case *hxrt.Array:
			value.array = typed
			return true
		default:
			return false
		}
	case "current":
		if incoming == nil {
			var zero int
			value.current = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.current = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "keys":
		return hxrt.ReadStoredField(value.keys)
	case "map":
		return hxrt.ReadStoredField(value.map_)
	}
	return nil
}

func hxrt__generated_field_has__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		return true
	case "map":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__iterators__MapKeyValueIterator(value *haxe__iterators__MapKeyValueIterator, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "keys":
		if incoming == nil {
			var zero hxrt.Anonymous
			value.keys = zero
			return true
		}
		switch typed := incoming.(type) {
		case hxrt.Anonymous:
			value.keys = typed
			return true
		default:
			return false
		}
	case "map":
		if incoming == nil {
			var zero haxe__IMap
			value.map_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case haxe__IMap:
			value.map_ = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(any, any) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		return hxrt.FunctionCall(source)(argument_0, argument_1)
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func(any, any)]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		var argument_1 any = nil
		if 1 < len(arguments) {
			argument_1 = arguments[1]
		}
		hxrt.FunctionCall(source)(argument_0, argument_1)
		return nil
	case *hxrt.Function[func(any) any]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func(any) bool]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(argument_0)
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func()]:
		hxrt.FunctionCall(source)()
		return nil
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func(*string, *string) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*string, *string) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(any, any) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			return hxrt.FunctionCall(source)(arg_0, arg_1)
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			return func(hx_value_223 any) int {
				if hx_value_223 == nil {
					var hx_zero_224 int
					return hx_zero_224
				}
				return hx_value_223.(int)
			}(any(hxrt.FunctionCall(source)(arg_0, arg_1)))
		})
	case *hxrt.Function[func(any, any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string, arg_1 *string) int {
			hxrt.FunctionCall(source)(arg_0, arg_1)
			return func(hx_value_225 any) int {
				if hx_value_225 == nil {
					var hx_zero_226 int
					return hx_zero_226
				}
				return hx_value_225.(int)
			}(any(nil))
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
			return func(hx_value_227 any) int {
				if hx_value_227 == nil {
					var hx_zero_228 int
					return hx_zero_228
				}
				return hx_value_227.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_229 any) int {
				if hx_value_229 == nil {
					var hx_zero_230 int
					return hx_zero_230
				}
				return hx_value_229.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_231 any) int {
				if hx_value_231 == nil {
					var hx_zero_232 int
					return hx_zero_232
				}
				return hx_value_231.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_233 any) int {
				if hx_value_233 == nil {
					var hx_zero_234 int
					return hx_zero_234
				}
				return hx_value_233.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_235 any) int {
				if hx_value_235 == nil {
					var hx_zero_236 int
					return hx_zero_236
				}
				return hx_value_235.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_237 any) int {
				if hx_value_237 == nil {
					var hx_zero_238 int
					return hx_zero_238
				}
				return hx_value_237.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_239 any) int {
				if hx_value_239 == nil {
					var hx_zero_240 int
					return hx_zero_240
				}
				return hx_value_239.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() int {
			hxrt.FunctionCall(source)()
			return func(hx_value_241 any) int {
				if hx_value_241 == nil {
					var hx_zero_242 int
					return hx_zero_242
				}
				return hx_value_241.(int)
			}(any(nil))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_243 any) int {
				if hx_value_243 == nil {
					var hx_zero_244 int
					return hx_zero_244
				}
				return hx_value_243.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() bool] {
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
			return func(hx_value_245 any) bool {
				if hx_value_245 == nil {
					var hx_zero_246 bool
					return hx_zero_246
				}
				return hx_value_245.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_247 any) bool {
				if hx_value_247 == nil {
					var hx_zero_248 bool
					return hx_zero_248
				}
				return hx_value_247.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_249 any) bool {
				if hx_value_249 == nil {
					var hx_zero_250 bool
					return hx_zero_250
				}
				return hx_value_249.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_251 any) bool {
				if hx_value_251 == nil {
					var hx_zero_252 bool
					return hx_zero_252
				}
				return hx_value_251.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_253 any) bool {
				if hx_value_253 == nil {
					var hx_zero_254 bool
					return hx_zero_254
				}
				return hx_value_253.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_255 any) bool {
				if hx_value_255 == nil {
					var hx_zero_256 bool
					return hx_zero_256
				}
				return hx_value_255.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_257 any) bool {
				if hx_value_257 == nil {
					var hx_zero_258 bool
					return hx_zero_258
				}
				return hx_value_257.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_259 any) bool {
				if hx_value_259 == nil {
					var hx_zero_260 bool
					return hx_zero_260
				}
				return hx_value_259.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_261 any) bool {
				if hx_value_261 == nil {
					var hx_zero_262 bool
					return hx_zero_262
				}
				return hx_value_261.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__ObjectMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ds__ObjectMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__ObjectMap:
		return source
	default:
		return value.(*haxe__ds__ObjectMap)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_2(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Cursor:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
