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
	self := &Cursor{}
	self.__hx_this = self
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
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_3 := hx_structural_source_2
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_3.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_3.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_4 := hx_structural_source_2
					return hxrt.BoundFunction(func() int {
						return hx_receiver_4.__hx_this.next()
					}, func() any {
						return hx_receiver_4.__hx_this
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
	var v any = Reflect_field(view, hxrt.StringFromLiteral("tag"))
	hxrt.Println(v)
	var v_1 any = any(Reflect_hasField(view, hxrt.StringFromLiteral("tag")))
	hxrt.Println(v_1)
	names := Reflect_fields(view)
	hxrt.ArraySort(names, func(hx_comparator_10 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_10, func(hx_cmp_left_8 any, hx_cmp_right_9 any) int {
			return hxrt.FunctionCall(hx_comparator_10)(func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hx_cmp_left_8), func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hx_cmp_right_9))
		})
	}(func(hx_callback_source_7 *hxrt.Function[func(any, any) int]) *hxrt.Function[func(*string, *string) int] {
		return hxrt.AdaptFunction(hx_callback_source_7, func(hx_erased_callback_arg_5 *string, hx_erased_callback_arg_6 *string) int {
			return hxrt.FunctionCall(hx_callback_source_7)(hx_erased_callback_arg_5, hx_erased_callback_arg_6)
		})
	}(hxrt.NamedFunction(Reflect_compare, "Reflect.compare"))))
	var v_2 any = any(hxrt.StringJoinAny(names.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StdString(view))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringEqualStringPtr(Type_getClassName(Type_getClass(view)), Type_getClassName(Type_getClass(source))))
	hxrt.Println(v_4)
	var v_5 any = any(func(hx_value any) bool {
		switch hx_carrier := hx_value.(type) {
		case *Cursor:
			if hx_carrier == nil {
				return false
			}
			return true
		default:
			return false
		}
	}(hxrt.ObjectSource(hxrt.ObjectSource(view))))
	hxrt.Println(v_5)
	keys := New_haxe__ds__ObjectMap()
	keys.__hx_this.set(source, 11)
	var v_6 any = any(func(hx_value_15 any) any {
		if hx_value_15 == nil {
			return nil
		}
		return hx_value_15.(int)
	}(keys.__hx_this.get(view)))
	hxrt.Println(v_6)
	var v_7 any = any(func() bool {
		var f1 any = Reflect_field(view, hxrt.StringFromLiteral("next"))
		return hxrt.ReflectCompareMethods(f1, any(func() *hxrt.Function[func() int] {
			hx_receiver_16 := source
			return hxrt.BoundFunction(func() int {
				return hx_receiver_16.__hx_this.next()
			}, func() any {
				return hx_receiver_16.__hx_this
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
	copied := hxrt.AnonymousFromValue(Reflect_copy(view))
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
	hx_obj_18 := map[string]any{}
	hx_obj_18["tag"] = hxrt.StringFromLiteral("extra")
	hx_obj_18["hasNext"] = hxrt.NewFunction(func() bool {
		return false
	})
	hx_obj_18["next"] = callback
	rich := hxrt.AnonymousFromMap(hx_obj_18)
	projected := rich
	var v_12 any = any(same(rich, projected))
	hxrt.Println(v_12)
	var v_13 any = Reflect_field(projected, hxrt.StringFromLiteral("tag"))
	hxrt.Println(v_13)
	clone := hxrt.AnonymousFromValue(Reflect_copy(projected))
	var v_14 any = any(same(clone, projected))
	hxrt.Println(v_14)
	hxrt.Println(any(hxrt.HaxeEqual(func(hx_obj_20 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_21 := hxrt.AnonymousGet(hx_obj_20, "next")
		if hx_field_21 == nil {
			var hx_zero_22 *hxrt.Function[func() int]
			return hx_zero_22
		}
		return hx_field_21.(*hxrt.Function[func() int])
	}(clone), callback)))
	func() *hxrt.Function[func() int] {
		hx_structural_target_23 := rich
		hx_structural_value_25 := hxrt.NewFunction(func() int {
			return 22
		})
		hxrt.AnonymousSet(hx_structural_target_23, "next", hx_structural_value_25)
		return hx_structural_value_25
	}()
	var v_15 any = any(func() int {
		var hx_callable_29 *hxrt.Function[func() int] = func(hx_obj_26 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_27 := hxrt.AnonymousGet(hx_obj_26, "next")
			if hx_field_27 == nil {
				var hx_zero_28 *hxrt.Function[func() int]
				return hx_zero_28
			}
			return hx_field_27.(*hxrt.Function[func() int])
		}(projected)
		return hxrt.FunctionCall(hx_callable_29)()
	}())
	hxrt.Println(v_15)
	var v_16 any = any(func() int {
		var hx_callable_33 *hxrt.Function[func() int] = func(hx_obj_30 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_31 := hxrt.AnonymousGet(hx_obj_30, "next")
			if hx_field_31 == nil {
				var hx_zero_32 *hxrt.Function[func() int]
				return hx_zero_32
			}
			return hx_field_31.(*hxrt.Function[func() int])
		}(clone)
		return hxrt.FunctionCall(hx_callable_33)()
	}())
	hxrt.Println(v_16)
	values := hxrt.NewArray(1, 2)
	cursor := func(hx_structural_source_34 *haxe__iterators__ArrayIterator) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_34, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_35 := hx_structural_source_34
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_35.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_35.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func(hx_callback_source_37 *hxrt.Function[func() any]) *hxrt.Function[func() int] {
					return hxrt.AdaptFunction(hx_callback_source_37, func() int {
						return func(hx_value_38 any) int {
							if hx_value_38 == nil {
								var hx_zero_39 int
								return hx_zero_39
							}
							return hx_value_38.(int)
						}(hxrt.FunctionCall(hx_callback_source_37)())
					})
				}(func() *hxrt.Function[func() any] {
					hx_receiver_36 := hx_structural_source_34
					return hxrt.BoundFunction(func() any {
						return hx_receiver_36.__hx_this.next()
					}, func() any {
						return hx_receiver_36.__hx_this
					}, "next")
				}()), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_haxe__iterators__ArrayIterator(values))
	a := func(hx_obj_40 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_41 := hxrt.AnonymousGet(hx_obj_40, "next")
		if hx_field_41 == nil {
			var hx_zero_42 *hxrt.Function[func() int]
			return hx_zero_42
		}
		return hx_field_41.(*hxrt.Function[func() int])
	}(cursor)
	b := func(hx_obj_43 hxrt.Anonymous) *hxrt.Function[func() int] {
		hx_field_44 := hxrt.AnonymousGet(hx_obj_43, "next")
		if hx_field_44 == nil {
			var hx_zero_45 *hxrt.Function[func() int]
			return hx_zero_45
		}
		return hx_field_44.(*hxrt.Function[func() int])
	}(cursor)
	hxrt.Println(any(hxrt.HaxeEqual(a, b)))
	var v_17 any = any(hxrt.ReflectCompareMethods(any(a), any(b)))
	hxrt.Println(v_17)
	var v_18 any = any(func() int {
		var hx_callable_49 *hxrt.Function[func() int] = func(hx_obj_46 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_47 := hxrt.AnonymousGet(hx_obj_46, "next")
			if hx_field_47 == nil {
				var hx_zero_48 *hxrt.Function[func() int]
				return hx_zero_48
			}
			return hx_field_47.(*hxrt.Function[func() int])
		}(cursor)
		return hxrt.FunctionCall(hx_callable_49)()
	}())
	hxrt.Println(v_18)
	cursorFields := Reflect_fields(cursor)
	hxrt.ArraySort(cursorFields, func(hx_comparator_55 *hxrt.Function[func(*string, *string) int]) *hxrt.Function[func(any, any) int] {
		return hxrt.AdaptFunction(hx_comparator_55, func(hx_cmp_left_53 any, hx_cmp_right_54 any) int {
			return hxrt.FunctionCall(hx_comparator_55)(func(hx_value_56 any) *string {
				if hx_value_56 == nil {
					var hx_zero_57 *string
					return hx_zero_57
				}
				return hx_value_56.(*string)
			}(hx_cmp_left_53), func(hx_value_58 any) *string {
				if hx_value_58 == nil {
					var hx_zero_59 *string
					return hx_zero_59
				}
				return hx_value_58.(*string)
			}(hx_cmp_right_54))
		})
	}(func(hx_callback_source_52 *hxrt.Function[func(any, any) int]) *hxrt.Function[func(*string, *string) int] {
		return hxrt.AdaptFunction(hx_callback_source_52, func(hx_erased_callback_arg_50 *string, hx_erased_callback_arg_51 *string) int {
			return hxrt.FunctionCall(hx_callback_source_52)(hx_erased_callback_arg_50, hx_erased_callback_arg_51)
		})
	}(hxrt.NamedFunction(Reflect_compare, "Reflect.compare"))))
	var v_19 any = any(hxrt.StringJoinAny(cursorFields.Values(), hxrt.StringFromLiteral(",")))
	hxrt.Println(v_19)
	hx_array_target_60 := values
	hx_array_index_61 := 1
	hx_array_target_60.Set(hx_array_index_61, 9)
	var v_20 any = any(func() int {
		var hx_callable_65 *hxrt.Function[func() int] = func(hx_obj_62 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_63 := hxrt.AnonymousGet(hx_obj_62, "next")
			if hx_field_63 == nil {
				var hx_zero_64 *hxrt.Function[func() int]
				return hx_zero_64
			}
			return hx_field_63.(*hxrt.Function[func() int])
		}(cursor)
		return hxrt.FunctionCall(hx_callable_65)()
	}())
	hxrt.Println(v_20)
}

func same(left any, right any) bool {
	return hxrt.HaxeEqual(left, right)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Cursor:
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
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_1.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() int] {
			hx_receiver_2 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_2.__hx_this.next()
			}, func() any {
				return hx_receiver_2.__hx_this
			}, "next")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_3 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_3.__hx_this.toString()
			}, func() any {
				return hx_receiver_3.__hx_this
			}, "toString")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ds__ObjectMap(value *haxe__ds__ObjectMap, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "clear":
		return func() *hxrt.Function[func()] {
			hx_receiver_4 := value
			return hxrt.BoundFunction(func() {
				hx_receiver_4.__hx_this.clear()
			}, func() any {
				return hx_receiver_4.__hx_this
			}, "clear")
		}()
	case "copy":
		return func() *hxrt.Function[func() *haxe__ds__ObjectMap] {
			hx_receiver_5 := value
			return hxrt.BoundFunction(func() *haxe__ds__ObjectMap {
				return hx_receiver_5.__hx_this.copy()
			}, func() any {
				return hx_receiver_5.__hx_this
			}, "copy")
		}()
	case "copyIMap":
		return func() *hxrt.Function[func() haxe__IMap] {
			hx_receiver_6 := value
			return hxrt.BoundFunction(func() haxe__IMap {
				return hx_receiver_6.__hx_this.copyIMap()
			}, func() any {
				return hx_receiver_6.__hx_this
			}, "copyIMap")
		}()
	case "exists":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_7 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_7.__hx_this.exists(key)
			}, func() any {
				return hx_receiver_7.__hx_this
			}, "exists")
		}()
	case "existsIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_8 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_8.__hx_this.existsIMap(key)
			}, func() any {
				return hx_receiver_8.__hx_this
			}, "existsIMap")
		}()
	case "get":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_9 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_9.__hx_this.get(key)
			}, func() any {
				return hx_receiver_9.__hx_this
			}, "get")
		}()
	case "getIMap":
		return func() *hxrt.Function[func(any) any] {
			hx_receiver_10 := value
			return hxrt.BoundFunction(func(key any) any {
				return hx_receiver_10.__hx_this.getIMap(key)
			}, func() any {
				return hx_receiver_10.__hx_this
			}, "getIMap")
		}()
	case "iterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_11 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_11.__hx_this.iterator()
			}, func() any {
				return hx_receiver_11.__hx_this
			}, "iterator")
		}()
	case "keyValueIterator":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_12 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_12.__hx_this.keyValueIterator()
			}, func() any {
				return hx_receiver_12.__hx_this
			}, "keyValueIterator")
		}()
	case "keys":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_13 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_13.__hx_this.keys()
			}, func() any {
				return hx_receiver_13.__hx_this
			}, "keys")
		}()
	case "remove":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_14 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_14.__hx_this.remove(key)
			}, func() any {
				return hx_receiver_14.__hx_this
			}, "remove")
		}()
	case "removeIMap":
		return func() *hxrt.Function[func(any) bool] {
			hx_receiver_15 := value
			return hxrt.BoundFunction(func(key any) bool {
				return hx_receiver_15.__hx_this.removeIMap(key)
			}, func() any {
				return hx_receiver_15.__hx_this
			}, "removeIMap")
		}()
	case "set":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_16 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_16.__hx_this.set(key, value)
			}, func() any {
				return hx_receiver_16.__hx_this
			}, "set")
		}()
	case "setIMap":
		return func() *hxrt.Function[func(any, any)] {
			hx_receiver_17 := value
			return hxrt.BoundFunction(func(key any, value any) {
				hx_receiver_17.__hx_this.setIMap(key, value)
			}, func() any {
				return hx_receiver_17.__hx_this
			}, "setIMap")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_18 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_18.__hx_this.toString()
			}, func() any {
				return hx_receiver_18.__hx_this
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
			hx_receiver_19 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_19.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_19.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() any] {
			hx_receiver_20 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_20.__hx_this.next()
			}, func() any {
				return hx_receiver_20.__hx_this
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
			hx_receiver_21 := value
			return hxrt.BoundFunction(func() bool {
				return hx_receiver_21.__hx_this.hasNext()
			}, func() any {
				return hx_receiver_21.__hx_this
			}, "hasNext")
		}()
	case "next":
		return func() *hxrt.Function[func() hxrt.Anonymous] {
			hx_receiver_22 := value
			return hxrt.BoundFunction(func() hxrt.Anonymous {
				return hx_receiver_22.__hx_this.next()
			}, func() any {
				return hx_receiver_22.__hx_this
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
	case "haxe.Int64Helper":
		return nil, false
	case "haxe.Json":
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
	case "haxe.Int64Helper":
		return nil
	case "haxe.Json":
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
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
		return hxrt.NewArray(hxrt.StringFromLiteral("parse"), hxrt.StringFromLiteral("stringify"))
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
	case "haxe.Int64Helper":
		return hxrt.NewArray()
	case "haxe.Json":
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
	case "haxe.Int64Helper":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Json":
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
	return false
}
