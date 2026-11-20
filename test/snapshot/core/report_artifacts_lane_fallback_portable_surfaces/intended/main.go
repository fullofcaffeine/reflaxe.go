package main

import (
	"reflect"
	"snapshot/hxrt"
)

func main() {
	LaneWorker_produce()
	NonLaneWorker_produce()
	func() {
		var hx_callable_2 *hxrt.Function[func(any, hxrt.Anonymous)] = haxe__Log_trace
		var hx_argument_3 any = hxrt.StringFromLiteral("fallback-report-portable-surfaces")
		var hx_argument_4 hxrt.Anonymous = func() hxrt.Anonymous {
			hx_obj_1 := map[string]any{}
			hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
			hx_obj_1["lineNumber"] = 5
			hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
			hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
			return hxrt.AnonymousFromMap(hx_obj_1)
		}()
		hxrt.FunctionCall(hx_callable_2)(hx_argument_3, hx_argument_4)
	}()
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *go___Chan:
		return func(hx_structural_source_1 *go___Chan) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "__hx_native":
					return hxrt.AnonymousField{Value: hx_structural_source_1.__hx_native, Found: true}
				case "__hx_setBuffer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_2 := hx_structural_source_1
						return hxrt.BoundFunction(func(buffer int) {
							hx_receiver_2.__hx_this.__hx_setBuffer(buffer)
						}, func() any {
							return hx_receiver_2.__hx_this
						}, "__hx_setBuffer")
					}(), Found: true}
				case "send":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_3 := hx_structural_source_1
						return hxrt.BoundFunction(func(value any) {
							hx_receiver_3.__hx_this.send(value)
						}, func() any {
							return hx_receiver_3.__hx_this
						}, "send")
					}(), Found: true}
				case "recv":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_4 := hx_structural_source_1
						return hxrt.BoundFunction(func() any {
							return hx_receiver_4.__hx_this.recv()
						}, func() any {
							return hx_receiver_4.__hx_this
						}, "recv")
					}(), Found: true}
				case "trySend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_5 := hx_structural_source_1
						return hxrt.BoundFunction(func(value any) bool {
							return hx_receiver_5.__hx_this.trySend(value)
						}, func() any {
							return hx_receiver_5.__hx_this
						}, "trySend")
					}(), Found: true}
				case "tryRecv":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *go___Result] {
						hx_receiver_6 := hx_structural_source_1
						return hxrt.BoundFunction(func() *go___Result {
							return hx_receiver_6.__hx_this.tryRecv()
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "tryRecv")
					}(), Found: true}
				case "recvOr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_7 := hx_structural_source_1
						return hxrt.BoundFunction(func(defaultValue any) any {
							return hx_receiver_7.__hx_this.recvOr(defaultValue)
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "recvOr")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_8 := hx_structural_source_1
						return hxrt.BoundFunction(func() {
							hx_receiver_8.__hx_this.close()
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "close")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "__hx_native":
					hx_structural_source_1.__hx_native = incoming
					return true
				}
				return false
			})
		}(source)
	case *go___Error:
		return func(hx_structural_source_9 *go___Error) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_9, func(name string) hxrt.AnonymousField {
				switch name {
				case "message":
					return hxrt.AnonymousField{Value: hx_structural_source_9.message, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_10 := hx_structural_source_9
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_10.__hx_this.toString()
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "toString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *go___Map:
		return func(hx_structural_source_11 *go___Map) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_11, func(name string) hxrt.AnonymousField {
				switch name {
				case "inner":
					return hxrt.AnonymousField{Value: hx_structural_source_11.inner, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_12 := hx_structural_source_11
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_12.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_13 := hx_structural_source_11
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_13.__hx_this.get(key)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_14 := hx_structural_source_11
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_14.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "exists")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *go___Result:
		return func(hx_structural_source_15 *go___Result) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_15, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_15.value, Found: true}
				case "errorValue":
					return hxrt.AnonymousField{Value: hx_structural_source_15.errorValue, Found: true}
				case "isOk":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_16 := hx_structural_source_15
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_16.__hx_this.isOk()
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "isOk")
					}(), Found: true}
				case "isErr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_17 := hx_structural_source_15
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_17.__hx_this.isErr()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "isErr")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_18 := hx_structural_source_15
						return hxrt.BoundFunction(func() any {
							return hx_receiver_18.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "error":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_19 := hx_structural_source_15
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_19.__hx_this.error()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "error")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *go___Slice:
		return func(hx_structural_source_20 *go___Slice) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_20, func(name string) hxrt.AnonymousField {
				switch name {
				case "data":
					return hxrt.AnonymousField{Value: hx_structural_source_20.data, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_21 := hx_structural_source_20
						return hxrt.BoundFunction(func() int {
							return hx_receiver_21.__hx_this.get_length()
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "get_length")
					}(), Found: true}
				case "push":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_22 := hx_structural_source_20
						return hxrt.BoundFunction(func(value any) {
							hx_receiver_22.__hx_this.push(value)
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "push")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) any] {
						hx_receiver_23 := hx_structural_source_20
						return hxrt.BoundFunction(func(index int) any {
							return hx_receiver_23.__hx_this.get(index)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, any)] {
						hx_receiver_24 := hx_structural_source_20
						return hxrt.BoundFunction(func(index int, value any) {
							hx_receiver_24.__hx_this.set(index, value)
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "set")
					}(), Found: true}
				case "toArray":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Array] {
						hx_receiver_25 := hx_structural_source_20
						return hxrt.BoundFunction(func() *hxrt.Array {
							return hx_receiver_25.__hx_this.toArray()
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "toArray")
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
		return func(hx_structural_source_26 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_26, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_26.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_26.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_26.high = func(hx_value_27 any) int {
						if hx_value_27 == nil {
							var hx_zero_28 int
							return hx_zero_28
						}
						return hx_value_27.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_26.low = func(hx_value_29 any) int {
						if hx_value_29 == nil {
							var hx_zero_30 int
							return hx_zero_30
						}
						return hx_value_29.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_31 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_31, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_31.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_31.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_34 := hx_structural_source_31
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_34.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_35 := hx_structural_source_31
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_35.__hx_this.next()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_31.map_ = func(hx_value_32 any) haxe__IMap {
						if hx_value_32 == nil {
							var hx_zero_33 haxe__IMap
							return hx_zero_33
						}
						return hx_value_32.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_31.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds___List__GoListKeyValueIterator:
		return func(hx_structural_source_36 *haxe__ds___List__GoListKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_36, func(name string) hxrt.AnonymousField {
				switch name {
				case "items":
					return hxrt.AnonymousField{Value: hx_structural_source_36.items, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_36.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_39 := hx_structural_source_36
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_39.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_40 := hx_structural_source_36
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_40.__hx_this.next()
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_36.index = func(hx_value_37 any) int {
						if hx_value_37 == nil {
							var hx_zero_38 int
							return hx_zero_38
						}
						return hx_value_37.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds___List__GoListIterator:
		return func(hx_structural_source_41 *haxe__ds___List__GoListIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_41, func(name string) hxrt.AnonymousField {
				switch name {
				case "items":
					return hxrt.AnonymousField{Value: hx_structural_source_41.items, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_41.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_44 := hx_structural_source_41
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_44.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_45 := hx_structural_source_41
						return hxrt.BoundFunction(func() any {
							return hx_receiver_45.__hx_this.next()
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_41.index = func(hx_value_42 any) int {
						if hx_value_42 == nil {
							var hx_zero_43 int
							return hx_zero_43
						}
						return hx_value_42.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_46 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_46, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_46.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_49 := hx_structural_source_46
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_49.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_50 := hx_structural_source_46
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_50.__hx_this.get(key)
						}, func() any {
							return hx_receiver_50.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_51 := hx_structural_source_46
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_51.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_52 := hx_structural_source_46
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_52.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_53 := hx_structural_source_46
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_53.__hx_this.keys()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_54 := hx_structural_source_46
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_54.__hx_this.iterator()
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_55 := hx_structural_source_46
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_55.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_56 := hx_structural_source_46
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_56.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_57 := hx_structural_source_46
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_57.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_58 := hx_structural_source_46
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_58.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_59 := hx_structural_source_46
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_59.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_60 := hx_structural_source_46
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_60.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_61 := hx_structural_source_46
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_61.__hx_this.copy()
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_62 := hx_structural_source_46
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_62.__hx_this.toString()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_63 := hx_structural_source_46
						return hxrt.BoundFunction(func() {
							hx_receiver_63.__hx_this.clear()
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_46.h = func(hx_value_47 any) *hxrt.StringMapCell {
						if hx_value_47 == nil {
							var hx_zero_48 *hxrt.StringMapCell
							return hx_zero_48
						}
						return hx_value_47.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__List:
		return func(hx_structural_source_64 *haxe__ds__List) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_64, func(name string) hxrt.AnonymousField {
				switch name {
				case "items":
					return hxrt.AnonymousField{Value: hx_structural_source_64.items, Found: true}
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_64.length, Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_67 := hx_structural_source_64
						return hxrt.BoundFunction(func(item any) {
							hx_receiver_67.__hx_this.add(item)
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "add")
					}(), Found: true}
				case "push":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_68 := hx_structural_source_64
						return hxrt.BoundFunction(func(item any) {
							hx_receiver_68.__hx_this.push(item)
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "push")
					}(), Found: true}
				case "first":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_69 := hx_structural_source_64
						return hxrt.BoundFunction(func() any {
							return hx_receiver_69.__hx_this.first()
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "first")
					}(), Found: true}
				case "last":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_70 := hx_structural_source_64
						return hxrt.BoundFunction(func() any {
							return hx_receiver_70.__hx_this.last()
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "last")
					}(), Found: true}
				case "pop":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_71 := hx_structural_source_64
						return hxrt.BoundFunction(func() any {
							return hx_receiver_71.__hx_this.pop()
						}, func() any {
							return hx_receiver_71.__hx_this
						}, "pop")
					}(), Found: true}
				case "isEmpty":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_72 := hx_structural_source_64
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_72.__hx_this.isEmpty()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "isEmpty")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_73 := hx_structural_source_64
						return hxrt.BoundFunction(func() {
							hx_receiver_73.__hx_this.clear()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "clear")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_74 := hx_structural_source_64
						return hxrt.BoundFunction(func(value any) bool {
							return hx_receiver_74.__hx_this.remove(value)
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "remove")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds___List__GoListIterator] {
						hx_receiver_75 := hx_structural_source_64
						return hxrt.BoundFunction(func() *haxe__ds___List__GoListIterator {
							return hx_receiver_75.__hx_this.iterator()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator] {
						hx_receiver_76 := hx_structural_source_64
						return hxrt.BoundFunction(func() *haxe__ds___List__GoListKeyValueIterator {
							return hx_receiver_76.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_77 := hx_structural_source_64
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_77.__hx_this.toString()
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "toString")
					}(), Found: true}
				case "join":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *string] {
						hx_receiver_78 := hx_structural_source_64
						return hxrt.BoundFunction(func(separator *string) *string {
							return hx_receiver_78.__hx_this.join(separator)
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "join")
					}(), Found: true}
				case "filter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List] {
						hx_receiver_79 := hx_structural_source_64
						return hxrt.BoundFunction(func(predicate *hxrt.Function[func(any) bool]) *haxe__ds__List {
							return hx_receiver_79.__hx_this.filter(predicate)
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "filter")
					}(), Found: true}
				case "map":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List] {
						hx_receiver_80 := hx_structural_source_64
						return hxrt.BoundFunction(func(transform *hxrt.Function[func(any) any]) *haxe__ds__List {
							return hx_receiver_80.__hx_this.map_(transform)
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "map")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "items":
					hx_structural_source_64.items = func(hx_value_65 any) *hxrt.Array {
						if hx_value_65 == nil {
							var hx_zero_66 *hxrt.Array
							return hx_zero_66
						}
						return hx_value_65.(*hxrt.Array)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_81 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_81, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_81.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_84 := hx_structural_source_81
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_84.__hx_this.get_message()
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_85 := hx_structural_source_81
						return hxrt.BoundFunction(func() any {
							return hx_receiver_85.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_86 := hx_structural_source_81
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_86.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_87 := hx_structural_source_81
						return hxrt.BoundFunction(func() any {
							return hx_receiver_87.__hx_this.get_native()
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_88 := hx_structural_source_81
						return hxrt.BoundFunction(func() any {
							return hx_receiver_88.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_89 := hx_structural_source_81
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_89.__hx_this.toString()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_90 := hx_structural_source_81
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_90.__hx_this.details()
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_81.savedMessage = func(hx_value_82 any) *string {
						if hx_value_82 == nil {
							var hx_zero_83 *string
							return hx_zero_83
						}
						return hx_value_82.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_81.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_81.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_81.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_91 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_91, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_91.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_92 := hx_structural_source_91
						return hxrt.BoundFunction(func() any {
							return hx_receiver_92.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_92.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_91.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_91.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_91.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_91.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_95 := hx_structural_source_91
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_95.__hx_this.get_message()
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_96 := hx_structural_source_91
						return hxrt.BoundFunction(func() any {
							return hx_receiver_96.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_96.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_97 := hx_structural_source_91
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_97.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_98 := hx_structural_source_91
						return hxrt.BoundFunction(func() any {
							return hx_receiver_98.__hx_this.get_native()
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_99 := hx_structural_source_91
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_99.__hx_this.toString()
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_100 := hx_structural_source_91
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_100.__hx_this.details()
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_91.savedMessage = func(hx_value_93 any) *string {
						if hx_value_93 == nil {
							var hx_zero_94 *string
							return hx_zero_94
						}
						return hx_value_93.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_91.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_91.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_91.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func go__concurrency_makeChan(buffer int) any {
	if buffer > 0 {
		return make(chan any, buffer)
	}
	return make(chan any)
}

func go__concurrency_send(channel any, value any) {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return
	}
	sendValue := reflect.ValueOf(value)
	if !sendValue.IsValid() {
		sendValue = reflect.Zero(chanValue.Type().Elem())
	} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {
		if sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {
			sendValue = sendValue.Convert(chanValue.Type().Elem())
		} else {
			return
		}
	}
	chanValue.Send(sendValue)
}

func go__concurrency_trySend(channel any, value any) bool {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return false
	}
	sendValue := reflect.ValueOf(value)
	if !sendValue.IsValid() {
		sendValue = reflect.Zero(chanValue.Type().Elem())
	} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {
		if sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {
			sendValue = sendValue.Convert(chanValue.Type().Elem())
		} else {
			return false
		}
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectSend, Chan: chanValue, Send: sendValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, _, _ := reflect.Select(cases)
	return chosen == 0
}

func go__concurrency_recv(channel any) any {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return nil
	}
	recvValue, _ := chanValue.Recv()
	if !recvValue.IsValid() {
		return nil
	}
	return recvValue.Interface()
}

func go__concurrency_recvOr(channel any, defaultValue any) any {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return defaultValue
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chanValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, recvValue, received := reflect.Select(cases)
	if chosen == 0 {
		if !received {
			return defaultValue
		}
		return recvValue.Interface()
	}
	return defaultValue
}

func go__concurrency_tryRecv(channel any) *go___Result {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
	}
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chanValue},
		{Dir: reflect.SelectDefault},
	}
	chosen, recvValue, received := reflect.Select(cases)
	if chosen == 0 {
		if !received {
			return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("closed")))
		}
		return New_go___Result(recvValue.Interface(), nil)
	}
	return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral("empty")))
}

func go__concurrency_close(channel any) {
	chanValue := reflect.ValueOf(channel)
	if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {
		return
	}
	chanValue.Close()
}

func go__concurrency_spawn(fn func()) {
	go fn()
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
	case *hxrt.Function[func() *go___Result]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			hxrt.FunctionCall(source)()
			return hx_structural_recover(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() hxrt.Anonymous {
			return hx_structural_recover(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
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
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *go___Result]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(hxrt.FunctionCall(source)())
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
	case *hxrt.Function[func() *go___Result]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.Array]:
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
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
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
	case *hxrt.Function[func() *string]:
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
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *go___Result]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_19 any) bool {
				if hx_value_19 == nil {
					var hx_zero_20 bool
					return hx_zero_20
				}
				return hx_value_19.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_21 any) bool {
				if hx_value_21 == nil {
					var hx_zero_22 bool
					return hx_zero_22
				}
				return hx_value_21.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_23 any) bool {
				if hx_value_23 == nil {
					var hx_zero_24 bool
					return hx_zero_24
				}
				return hx_value_23.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_go___Result(value any) *go___Result {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *go___Result:
		return source
	default:
		return value.(*go___Result)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_recover_class_haxe__ds__StringMap(value any) *haxe__ds__StringMap {
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

func hx_recover_class_haxe__ds___List__GoListIterator(value any) *haxe__ds___List__GoListIterator {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds___List__GoListIterator:
		return source
	default:
		return value.(*haxe__ds___List__GoListIterator)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
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

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
