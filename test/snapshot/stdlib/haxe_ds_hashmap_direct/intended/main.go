package main

import "snapshot/hxrt"

type I_Key interface {
	hashCode() int
}

type Key struct {
	__hx_this I_Key
	id        int
}

func (self *Key) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Key(id int) *Key {
	return hx_init_Key(hx_alloc_Key(nil), id)
}

func hx_alloc_Key(owner any) *Key {
	self := &Key{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Key)
	}
	return self
}

func hx_init_Key(self *Key, id int) *Key {
	self.id = id
	return self
}

func (self *Key) hashCode() int {
	return self.id
}

func (self *Key) HxClassName() string {
	return "Key"
}

func main() {
	var map_valuesByHash *haxe__ds__IntMap
	var map_keysByHash *haxe__ds__IntMap
	map_keysByHash = New_haxe__ds__IntMap()
	map_valuesByHash = New_haxe__ds__IntMap()
	key3 := New_Key(3)
	key4 := New_Key(4)
	hash := key3.hashCode()
	map_keysByHash.__hx_this.set(hash, key3)
	map_valuesByHash.__hx_this.set(hash, hxrt.StringFromLiteral("three"))
	hash_1 := key4.hashCode()
	map_keysByHash.__hx_this.set(hash_1, key4)
	map_valuesByHash.__hx_this.set(hash_1, hxrt.StringFromLiteral("four"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("three="), func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(map_valuesByHash.__hx_this.get(key3.hashCode()))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("four="), func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(map_valuesByHash.__hx_this.get(key4.hashCode()))))
	hxrt.Println(v_1)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Key:
		return func(hx_structural_source_1 *Key) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "id":
					return hxrt.AnonymousField{Value: hx_structural_source_1.id, Found: true}
				case "hashCode":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_2.__hx_this.hashCode()
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "hashCode")
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
		return func(hx_structural_source_3 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_3, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_3.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_3.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_3.high = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_3.low = func(hx_value_6 any) int {
						if hx_value_6 == nil {
							var hx_zero_7 int
							return hx_zero_7
						}
						return hx_value_6.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__HashMap:
		return func(hx_structural_source_8 *haxe__ds__HashMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_8, func(name string) hxrt.AnonymousField {
				switch name {
				case "keysByHash":
					return hxrt.AnonymousField{Value: hx_structural_source_8.keysByHash, Found: true}
				case "valuesByHash":
					return hxrt.AnonymousField{Value: hx_structural_source_8.valuesByHash, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, any)] {
						hx_receiver_9 := hx_structural_source_8
						return hxrt.BoundFunction(func(k hxrt.Anonymous, v any) {
							hx_receiver_9.__hx_this.set(k, v)
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous) any] {
						hx_receiver_10 := hx_structural_source_8
						return hxrt.BoundFunction(func(k hxrt.Anonymous) any {
							return hx_receiver_10.__hx_this.get(k)
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous) bool] {
						hx_receiver_11 := hx_structural_source_8
						return hxrt.BoundFunction(func(k hxrt.Anonymous) bool {
							return hx_receiver_11.__hx_this.exists(k)
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous) bool] {
						hx_receiver_12 := hx_structural_source_8
						return hxrt.BoundFunction(func(k hxrt.Anonymous) bool {
							return hx_receiver_12.__hx_this.remove(k)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_13 := hx_structural_source_8
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_13.__hx_this.keys()
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_14 := hx_structural_source_8
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_14.__hx_this.iterator()
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator] {
						hx_receiver_15 := hx_structural_source_8
						return hxrt.BoundFunction(func() *haxe__iterators__HashMapKeyValueIterator {
							return hx_receiver_15.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__HashMap] {
						hx_receiver_16 := hx_structural_source_8
						return hxrt.BoundFunction(func() *haxe__ds__HashMap {
							return hx_receiver_16.__hx_this.copy()
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "copy")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_17 := hx_structural_source_8
						return hxrt.BoundFunction(func() {
							hx_receiver_17.__hx_this.clear()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__iterators__HashMapKeyValueIterator:
		return func(hx_structural_source_18 *haxe__iterators__HashMapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_18, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_18.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_18.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_19 := hx_structural_source_18
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_19.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_20 := hx_structural_source_18
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_20.__hx_this.next()
						}, func() any {
							return hx_receiver_20.__hx_this
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
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_21 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_21, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_21.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_21.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_24 := hx_structural_source_21
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_24.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_25 := hx_structural_source_21
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_25.__hx_this.next()
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_21.map_ = func(hx_value_22 any) haxe__IMap {
						if hx_value_22 == nil {
							var hx_zero_23 haxe__IMap
							return hx_zero_23
						}
						return hx_value_22.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_21.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__IntMap:
		return func(hx_structural_source_26 *haxe__ds__IntMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_26, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_26.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, any)] {
						hx_receiver_29 := hx_structural_source_26
						return hxrt.BoundFunction(func(key int, value any) {
							hx_receiver_29.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) any] {
						hx_receiver_30 := hx_structural_source_26
						return hxrt.BoundFunction(func(key int) any {
							return hx_receiver_30.__hx_this.get(key)
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) bool] {
						hx_receiver_31 := hx_structural_source_26
						return hxrt.BoundFunction(func(key int) bool {
							return hx_receiver_31.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) bool] {
						hx_receiver_32 := hx_structural_source_26
						return hxrt.BoundFunction(func(key int) bool {
							return hx_receiver_32.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_33 := hx_structural_source_26
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_33.__hx_this.keys()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_34 := hx_structural_source_26
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_34.__hx_this.iterator()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_35 := hx_structural_source_26
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_35.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_36 := hx_structural_source_26
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_36.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_37 := hx_structural_source_26
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_37.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_38 := hx_structural_source_26
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_38.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_39 := hx_structural_source_26
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_39.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_40 := hx_structural_source_26
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_40.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__IntMap] {
						hx_receiver_41 := hx_structural_source_26
						return hxrt.BoundFunction(func() *haxe__ds__IntMap {
							return hx_receiver_41.__hx_this.copy()
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_42 := hx_structural_source_26
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_42.__hx_this.toString()
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_43 := hx_structural_source_26
						return hxrt.BoundFunction(func() {
							hx_receiver_43.__hx_this.clear()
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_26.h = func(hx_value_27 any) *hxrt.IntMapCell {
						if hx_value_27 == nil {
							var hx_zero_28 *hxrt.IntMapCell
							return hx_zero_28
						}
						return hx_value_27.(*hxrt.IntMapCell)
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

func hx_recover_callable_7838dc94(value any) *hxrt.Function[func() hxrt.Anonymous] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() hxrt.Anonymous]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__HashMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			hxrt.FunctionCall(source)()
			return hx_structural_recover(nil)
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__HashMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
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
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__HashMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(nil)
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_e273723c(value any) *hxrt.Function[func() int] {
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
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__HashMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() int {
			hxrt.FunctionCall(source)()
			return func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(nil)
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_13 any) int {
				if hx_value_13 == nil {
					var hx_zero_14 int
					return hx_zero_14
				}
				return hx_value_13.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_15 any) int {
				if hx_value_15 == nil {
					var hx_zero_16 int
					return hx_zero_16
				}
				return hx_value_15.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_17 any) int {
				if hx_value_17 == nil {
					var hx_zero_18 int
					return hx_zero_18
				}
				return hx_value_17.(int)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_haxe__ds__IntMap(value any) *haxe__ds__IntMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__IntMap:
		return source
	default:
		return value.(*haxe__ds__IntMap)
	}
}
