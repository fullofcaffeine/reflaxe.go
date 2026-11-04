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
		return func(hx_structural_source_80 *Key) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_80, func(name string) hxrt.AnonymousField {
				switch name {
				case "id":
					return hxrt.AnonymousField{Value: hx_structural_source_80.id, Found: true}
				case "hashCode":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_81 := hx_structural_source_80
						return hxrt.BoundFunction(func() int {
							return hx_receiver_81.__hx_this.hashCode()
						}, func() any {
							return hx_receiver_81.__hx_this
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
		return func(hx_structural_source_82 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_82, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_82.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_82.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_82.high = func(hx_value_83 any) int {
						if hx_value_83 == nil {
							var hx_zero_84 int
							return hx_zero_84
						}
						return hx_value_83.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_82.low = func(hx_value_85 any) int {
						if hx_value_85 == nil {
							var hx_zero_86 int
							return hx_zero_86
						}
						return hx_value_85.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__HashMap:
		return func(hx_structural_source_87 *haxe__ds__HashMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_87, func(name string) hxrt.AnonymousField {
				switch name {
				case "keysByHash":
					return hxrt.AnonymousField{Value: hx_structural_source_87.keysByHash, Found: true}
				case "valuesByHash":
					return hxrt.AnonymousField{Value: hx_structural_source_87.valuesByHash, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous, any)] {
						hx_receiver_88 := hx_structural_source_87
						return hxrt.BoundFunction(func(k hxrt.Anonymous, v any) {
							hx_receiver_88.__hx_this.set(k, v)
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous) any] {
						hx_receiver_89 := hx_structural_source_87
						return hxrt.BoundFunction(func(k hxrt.Anonymous) any {
							return hx_receiver_89.__hx_this.get(k)
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous) bool] {
						hx_receiver_90 := hx_structural_source_87
						return hxrt.BoundFunction(func(k hxrt.Anonymous) bool {
							return hx_receiver_90.__hx_this.exists(k)
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(hxrt.Anonymous) bool] {
						hx_receiver_91 := hx_structural_source_87
						return hxrt.BoundFunction(func(k hxrt.Anonymous) bool {
							return hx_receiver_91.__hx_this.remove(k)
						}, func() any {
							return hx_receiver_91.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_92 := hx_structural_source_87
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_92.__hx_this.keys()
						}, func() any {
							return hx_receiver_92.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_93 := hx_structural_source_87
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_93.__hx_this.iterator()
						}, func() any {
							return hx_receiver_93.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator] {
						hx_receiver_94 := hx_structural_source_87
						return hxrt.BoundFunction(func() *haxe__iterators__HashMapKeyValueIterator {
							return hx_receiver_94.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_94.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__HashMap] {
						hx_receiver_95 := hx_structural_source_87
						return hxrt.BoundFunction(func() *haxe__ds__HashMap {
							return hx_receiver_95.__hx_this.copy()
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "copy")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_96 := hx_structural_source_87
						return hxrt.BoundFunction(func() {
							hx_receiver_96.__hx_this.clear()
						}, func() any {
							return hx_receiver_96.__hx_this
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
		return func(hx_structural_source_97 *haxe__iterators__HashMapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_97, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_97.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_97.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_98 := hx_structural_source_97
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_98.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_99 := hx_structural_source_97
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_99.__hx_this.next()
						}, func() any {
							return hx_receiver_99.__hx_this
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
		return func(hx_structural_source_100 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_100, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_100.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_100.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_103 := hx_structural_source_100
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_103.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_104 := hx_structural_source_100
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_104.__hx_this.next()
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_100.map_ = func(hx_value_101 any) haxe__IMap {
						if hx_value_101 == nil {
							var hx_zero_102 haxe__IMap
							return hx_zero_102
						}
						return hx_value_101.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_100.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__IntMap:
		return func(hx_structural_source_105 *haxe__ds__IntMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_105, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_105.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, any)] {
						hx_receiver_108 := hx_structural_source_105
						return hxrt.BoundFunction(func(key int, value any) {
							hx_receiver_108.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) any] {
						hx_receiver_109 := hx_structural_source_105
						return hxrt.BoundFunction(func(key int) any {
							return hx_receiver_109.__hx_this.get(key)
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) bool] {
						hx_receiver_110 := hx_structural_source_105
						return hxrt.BoundFunction(func(key int) bool {
							return hx_receiver_110.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_110.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) bool] {
						hx_receiver_111 := hx_structural_source_105
						return hxrt.BoundFunction(func(key int) bool {
							return hx_receiver_111.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_111.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_112 := hx_structural_source_105
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_112.__hx_this.keys()
						}, func() any {
							return hx_receiver_112.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_113 := hx_structural_source_105
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_113.__hx_this.iterator()
						}, func() any {
							return hx_receiver_113.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_114 := hx_structural_source_105
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_114.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_114.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_115 := hx_structural_source_105
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_115.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_115.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_116 := hx_structural_source_105
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_116.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_117 := hx_structural_source_105
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_117.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_118 := hx_structural_source_105
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_118.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_119 := hx_structural_source_105
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_119.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__IntMap] {
						hx_receiver_120 := hx_structural_source_105
						return hxrt.BoundFunction(func() *haxe__ds__IntMap {
							return hx_receiver_120.__hx_this.copy()
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_121 := hx_structural_source_105
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_121.__hx_this.toString()
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_122 := hx_structural_source_105
						return hxrt.BoundFunction(func() {
							hx_receiver_122.__hx_this.clear()
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_105.h = func(hx_value_106 any) *hxrt.IntMapCell {
						if hx_value_106 == nil {
							var hx_zero_107 *hxrt.IntMapCell
							return hx_zero_107
						}
						return hx_value_106.(*hxrt.IntMapCell)
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
			return func(hx_value_123 any) int {
				if hx_value_123 == nil {
					var hx_zero_124 int
					return hx_zero_124
				}
				return hx_value_123.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_125 any) int {
				if hx_value_125 == nil {
					var hx_zero_126 int
					return hx_zero_126
				}
				return hx_value_125.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_127 any) int {
				if hx_value_127 == nil {
					var hx_zero_128 int
					return hx_zero_128
				}
				return hx_value_127.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return source
	case *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_129 any) int {
				if hx_value_129 == nil {
					var hx_zero_130 int
					return hx_zero_130
				}
				return hx_value_129.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__HashMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_131 any) int {
				if hx_value_131 == nil {
					var hx_zero_132 int
					return hx_zero_132
				}
				return hx_value_131.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() int {
			hxrt.FunctionCall(source)()
			return func(hx_value_133 any) int {
				if hx_value_133 == nil {
					var hx_zero_134 int
					return hx_zero_134
				}
				return hx_value_133.(int)
			}(any(nil))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_135 any) int {
				if hx_value_135 == nil {
					var hx_zero_136 int
					return hx_zero_136
				}
				return hx_value_135.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_137 any) int {
				if hx_value_137 == nil {
					var hx_zero_138 int
					return hx_zero_138
				}
				return hx_value_137.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_139 any) int {
				if hx_value_139 == nil {
					var hx_zero_140 int
					return hx_zero_140
				}
				return hx_value_139.(int)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() bool] {
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
			return func(hx_value_141 any) bool {
				if hx_value_141 == nil {
					var hx_zero_142 bool
					return hx_zero_142
				}
				return hx_value_141.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_143 any) bool {
				if hx_value_143 == nil {
					var hx_zero_144 bool
					return hx_zero_144
				}
				return hx_value_143.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_145 any) bool {
				if hx_value_145 == nil {
					var hx_zero_146 bool
					return hx_zero_146
				}
				return hx_value_145.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__iterators__HashMapKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_147 any) bool {
				if hx_value_147 == nil {
					var hx_zero_148 bool
					return hx_zero_148
				}
				return hx_value_147.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__HashMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_149 any) bool {
				if hx_value_149 == nil {
					var hx_zero_150 bool
					return hx_zero_150
				}
				return hx_value_149.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_151 any) bool {
				if hx_value_151 == nil {
					var hx_zero_152 bool
					return hx_zero_152
				}
				return hx_value_151.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_153 any) bool {
				if hx_value_153 == nil {
					var hx_zero_154 bool
					return hx_zero_154
				}
				return hx_value_153.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__IntMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_155 any) bool {
				if hx_value_155 == nil {
					var hx_zero_156 bool
					return hx_zero_156
				}
				return hx_value_155.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_157 any) bool {
				if hx_value_157 == nil {
					var hx_zero_158 bool
					return hx_zero_158
				}
				return hx_value_157.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() hxrt.Anonymous] {
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

func hx_recover_class_0(value any) *haxe__ds__IntMap {
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
