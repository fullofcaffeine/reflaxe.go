package main

import "snapshot/hxrt"

func main() {
	entries := New_haxe__ds__StringMap()
	entries.__hx_this.set(hxrt.StringFromLiteral("first"), func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["name"] = hxrt.StringFromLiteral("alpha")
		hx_obj_1["count"] = 3
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())
	_g := hxrt.NewArray()
	entry := entries.__hx_this.iterator()
	for func() bool {
		var hx_callable_2 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(entry, "hasNext"))
		return hxrt.FunctionCall(hx_callable_2)()
	}() {
		entry_1 := func() hxrt.Anonymous {
			var hx_callable_3 *hxrt.Function[func() hxrt.Anonymous] = hx_recover_callable_1(hxrt.AnonymousGet(entry, "next"))
			return hxrt.FunctionCall(hx_callable_3)()
		}()
		_g.Push(entry_1)
	}
	copied := _g
	entry_2 := hx_structural_recover(copied.Get(0))
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(""), func(hx_value_5 any) *string {
		if hx_value_5 == nil {
			var hx_zero_6 *string
			return hx_zero_6
		}
		return hx_value_5.(*string)
	}(hxrt.AnonymousGet(entry_2, "name"))), hxrt.StringFromLiteral(":")), func(hx_value_7 any) int {
		if hx_value_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_value_7.(int)
	}(hxrt.AnonymousGet(entry_2, "count"))))
	hxrt.Println(v)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_53 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_53, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_53.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_53.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_53.high = func(hx_value_54 any) int {
						if hx_value_54 == nil {
							var hx_zero_55 int
							return hx_zero_55
						}
						return hx_value_54.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_53.low = func(hx_value_56 any) int {
						if hx_value_56 == nil {
							var hx_zero_57 int
							return hx_zero_57
						}
						return hx_value_56.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_58 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_58, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_58.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_58.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_61 := hx_structural_source_58
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_61.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_62 := hx_structural_source_58
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_62.__hx_this.next()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_58.map_ = func(hx_value_59 any) haxe__IMap {
						if hx_value_59 == nil {
							var hx_zero_60 haxe__IMap
							return hx_zero_60
						}
						return hx_value_59.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_58.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_63 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_63, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_63.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_66 := hx_structural_source_63
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_66.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_67 := hx_structural_source_63
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_67.__hx_this.get(key)
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_68 := hx_structural_source_63
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_68.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_69 := hx_structural_source_63
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_69.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_70 := hx_structural_source_63
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_70.__hx_this.keys()
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_71 := hx_structural_source_63
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_71.__hx_this.iterator()
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_72 := hx_structural_source_63
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_72.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_73 := hx_structural_source_63
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_73.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_74 := hx_structural_source_63
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_74.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_75 := hx_structural_source_63
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_75.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_76 := hx_structural_source_63
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_76.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_77 := hx_structural_source_63
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_77.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_78 := hx_structural_source_63
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_78.__hx_this.copy()
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_79 := hx_structural_source_63
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_79.__hx_this.toString()
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_80 := hx_structural_source_63
						return hxrt.BoundFunction(func() {
							hx_receiver_80.__hx_this.clear()
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_63.h = func(hx_value_64 any) *hxrt.StringMapCell {
						if hx_value_64 == nil {
							var hx_zero_65 *hxrt.StringMapCell
							return hx_zero_65
						}
						return hx_value_64.(*hxrt.StringMapCell)
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_81 any) bool {
				if hx_value_81 == nil {
					var hx_zero_82 bool
					return hx_zero_82
				}
				return hx_value_81.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_83 any) bool {
				if hx_value_83 == nil {
					var hx_zero_84 bool
					return hx_zero_84
				}
				return hx_value_83.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_85 any) bool {
				if hx_value_85 == nil {
					var hx_zero_86 bool
					return hx_zero_86
				}
				return hx_value_85.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_87 any) bool {
				if hx_value_87 == nil {
					var hx_zero_88 bool
					return hx_zero_88
				}
				return hx_value_87.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_89 any) bool {
				if hx_value_89 == nil {
					var hx_zero_90 bool
					return hx_zero_90
				}
				return hx_value_89.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_91 any) bool {
				if hx_value_91 == nil {
					var hx_zero_92 bool
					return hx_zero_92
				}
				return hx_value_91.(bool)
			}(nil)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() hxrt.Anonymous] {
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
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return source
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			hxrt.FunctionCall(source)()
			return hx_structural_recover(nil)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_93 any) *string {
				if hx_value_93 == nil {
					var hx_zero_94 *string
					return hx_zero_94
				}
				return hx_value_93.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_95 any) *string {
				if hx_value_95 == nil {
					var hx_zero_96 *string
					return hx_zero_96
				}
				return hx_value_95.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_97 any) *string {
				if hx_value_97 == nil {
					var hx_zero_98 *string
					return hx_zero_98
				}
				return hx_value_97.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_99 any) *string {
				if hx_value_99 == nil {
					var hx_zero_100 *string
					return hx_zero_100
				}
				return hx_value_99.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_101 any) *string {
				if hx_value_101 == nil {
					var hx_zero_102 *string
					return hx_zero_102
				}
				return hx_value_101.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_103 any) *string {
				if hx_value_103 == nil {
					var hx_zero_104 *string
					return hx_zero_104
				}
				return hx_value_103.(*string)
			}(nil)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}
