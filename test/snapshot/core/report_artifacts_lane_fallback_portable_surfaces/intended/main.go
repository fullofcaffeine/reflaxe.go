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
		return func(hx_structural_source_161 *go___Chan) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_161, func(name string) hxrt.AnonymousField {
				switch name {
				case "__hx_native":
					return hxrt.AnonymousField{Value: hx_structural_source_161.__hx_native, Found: true}
				case "__hx_setBuffer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_162 := hx_structural_source_161
						return hxrt.BoundFunction(func(buffer int) {
							hx_receiver_162.__hx_this.__hx_setBuffer(buffer)
						}, func() any {
							return hx_receiver_162.__hx_this
						}, "__hx_setBuffer")
					}(), Found: true}
				case "send":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_163 := hx_structural_source_161
						return hxrt.BoundFunction(func(value any) {
							hx_receiver_163.__hx_this.send(value)
						}, func() any {
							return hx_receiver_163.__hx_this
						}, "send")
					}(), Found: true}
				case "recv":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_164 := hx_structural_source_161
						return hxrt.BoundFunction(func() any {
							return hx_receiver_164.__hx_this.recv()
						}, func() any {
							return hx_receiver_164.__hx_this
						}, "recv")
					}(), Found: true}
				case "trySend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_165 := hx_structural_source_161
						return hxrt.BoundFunction(func(value any) bool {
							return hx_receiver_165.__hx_this.trySend(value)
						}, func() any {
							return hx_receiver_165.__hx_this
						}, "trySend")
					}(), Found: true}
				case "tryRecv":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *go___Result] {
						hx_receiver_166 := hx_structural_source_161
						return hxrt.BoundFunction(func() *go___Result {
							return hx_receiver_166.__hx_this.tryRecv()
						}, func() any {
							return hx_receiver_166.__hx_this
						}, "tryRecv")
					}(), Found: true}
				case "recvOr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_167 := hx_structural_source_161
						return hxrt.BoundFunction(func(defaultValue any) any {
							return hx_receiver_167.__hx_this.recvOr(defaultValue)
						}, func() any {
							return hx_receiver_167.__hx_this
						}, "recvOr")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_168 := hx_structural_source_161
						return hxrt.BoundFunction(func() {
							hx_receiver_168.__hx_this.close()
						}, func() any {
							return hx_receiver_168.__hx_this
						}, "close")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "__hx_native":
					hx_structural_source_161.__hx_native = incoming
					return true
				}
				return false
			})
		}(source)
	case *go___Error:
		return func(hx_structural_source_169 *go___Error) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_169, func(name string) hxrt.AnonymousField {
				switch name {
				case "message":
					return hxrt.AnonymousField{Value: hx_structural_source_169.message, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_170 := hx_structural_source_169
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_170.__hx_this.toString()
						}, func() any {
							return hx_receiver_170.__hx_this
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
		return func(hx_structural_source_171 *go___Map) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_171, func(name string) hxrt.AnonymousField {
				switch name {
				case "inner":
					return hxrt.AnonymousField{Value: hx_structural_source_171.inner, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_172 := hx_structural_source_171
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_172.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_173 := hx_structural_source_171
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_173.__hx_this.get(key)
						}, func() any {
							return hx_receiver_173.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_174 := hx_structural_source_171
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_174.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_174.__hx_this
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
		return func(hx_structural_source_175 *go___Result) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_175, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_175.value, Found: true}
				case "errorValue":
					return hxrt.AnonymousField{Value: hx_structural_source_175.errorValue, Found: true}
				case "isOk":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_176 := hx_structural_source_175
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_176.__hx_this.isOk()
						}, func() any {
							return hx_receiver_176.__hx_this
						}, "isOk")
					}(), Found: true}
				case "isErr":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_177 := hx_structural_source_175
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_177.__hx_this.isErr()
						}, func() any {
							return hx_receiver_177.__hx_this
						}, "isErr")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_178 := hx_structural_source_175
						return hxrt.BoundFunction(func() any {
							return hx_receiver_178.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_178.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "error":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_179 := hx_structural_source_175
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_179.__hx_this.error()
						}, func() any {
							return hx_receiver_179.__hx_this
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
		return func(hx_structural_source_180 *go___Slice) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_180, func(name string) hxrt.AnonymousField {
				switch name {
				case "data":
					return hxrt.AnonymousField{Value: hx_structural_source_180.data, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_181 := hx_structural_source_180
						return hxrt.BoundFunction(func() int {
							return hx_receiver_181.__hx_this.get_length()
						}, func() any {
							return hx_receiver_181.__hx_this
						}, "get_length")
					}(), Found: true}
				case "push":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_182 := hx_structural_source_180
						return hxrt.BoundFunction(func(value any) {
							hx_receiver_182.__hx_this.push(value)
						}, func() any {
							return hx_receiver_182.__hx_this
						}, "push")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) any] {
						hx_receiver_183 := hx_structural_source_180
						return hxrt.BoundFunction(func(index int) any {
							return hx_receiver_183.__hx_this.get(index)
						}, func() any {
							return hx_receiver_183.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, any)] {
						hx_receiver_184 := hx_structural_source_180
						return hxrt.BoundFunction(func(index int, value any) {
							hx_receiver_184.__hx_this.set(index, value)
						}, func() any {
							return hx_receiver_184.__hx_this
						}, "set")
					}(), Found: true}
				case "toArray":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Array] {
						hx_receiver_185 := hx_structural_source_180
						return hxrt.BoundFunction(func() *hxrt.Array {
							return hx_receiver_185.__hx_this.toArray()
						}, func() any {
							return hx_receiver_185.__hx_this
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
		return func(hx_structural_source_186 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_186, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_186.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_186.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_186.high = func(hx_value_187 any) int {
						if hx_value_187 == nil {
							var hx_zero_188 int
							return hx_zero_188
						}
						return hx_value_187.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_186.low = func(hx_value_189 any) int {
						if hx_value_189 == nil {
							var hx_zero_190 int
							return hx_zero_190
						}
						return hx_value_189.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_191 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_191, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_191.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_191.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_194 := hx_structural_source_191
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_194.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_194.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_195 := hx_structural_source_191
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_195.__hx_this.next()
						}, func() any {
							return hx_receiver_195.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_191.map_ = func(hx_value_192 any) haxe__IMap {
						if hx_value_192 == nil {
							var hx_zero_193 haxe__IMap
							return hx_zero_193
						}
						return hx_value_192.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_191.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds___List__GoListKeyValueIterator:
		return func(hx_structural_source_196 *haxe__ds___List__GoListKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_196, func(name string) hxrt.AnonymousField {
				switch name {
				case "items":
					return hxrt.AnonymousField{Value: hx_structural_source_196.items, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_196.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_199 := hx_structural_source_196
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_199.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_199.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_200 := hx_structural_source_196
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_200.__hx_this.next()
						}, func() any {
							return hx_receiver_200.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_196.index = func(hx_value_197 any) int {
						if hx_value_197 == nil {
							var hx_zero_198 int
							return hx_zero_198
						}
						return hx_value_197.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds___List__GoListIterator:
		return func(hx_structural_source_201 *haxe__ds___List__GoListIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_201, func(name string) hxrt.AnonymousField {
				switch name {
				case "items":
					return hxrt.AnonymousField{Value: hx_structural_source_201.items, Found: true}
				case "index":
					return hxrt.AnonymousField{Value: hx_structural_source_201.index, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_204 := hx_structural_source_201
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_204.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_204.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_205 := hx_structural_source_201
						return hxrt.BoundFunction(func() any {
							return hx_receiver_205.__hx_this.next()
						}, func() any {
							return hx_receiver_205.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "index":
					hx_structural_source_201.index = func(hx_value_202 any) int {
						if hx_value_202 == nil {
							var hx_zero_203 int
							return hx_zero_203
						}
						return hx_value_202.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_206 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_206, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_206.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_209 := hx_structural_source_206
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_209.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_209.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_210 := hx_structural_source_206
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_210.__hx_this.get(key)
						}, func() any {
							return hx_receiver_210.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_211 := hx_structural_source_206
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_211.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_211.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_212 := hx_structural_source_206
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_212.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_212.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_213 := hx_structural_source_206
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_213.__hx_this.keys()
						}, func() any {
							return hx_receiver_213.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_214 := hx_structural_source_206
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_214.__hx_this.iterator()
						}, func() any {
							return hx_receiver_214.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_215 := hx_structural_source_206
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_215.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_215.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_216 := hx_structural_source_206
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_216.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_216.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_217 := hx_structural_source_206
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_217.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_217.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_218 := hx_structural_source_206
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_218.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_218.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_219 := hx_structural_source_206
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_219.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_219.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_220 := hx_structural_source_206
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_220.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_220.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_221 := hx_structural_source_206
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_221.__hx_this.copy()
						}, func() any {
							return hx_receiver_221.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_222 := hx_structural_source_206
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_222.__hx_this.toString()
						}, func() any {
							return hx_receiver_222.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_223 := hx_structural_source_206
						return hxrt.BoundFunction(func() {
							hx_receiver_223.__hx_this.clear()
						}, func() any {
							return hx_receiver_223.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_206.h = func(hx_value_207 any) *hxrt.StringMapCell {
						if hx_value_207 == nil {
							var hx_zero_208 *hxrt.StringMapCell
							return hx_zero_208
						}
						return hx_value_207.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__List:
		return func(hx_structural_source_224 *haxe__ds__List) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_224, func(name string) hxrt.AnonymousField {
				switch name {
				case "items":
					return hxrt.AnonymousField{Value: hx_structural_source_224.items, Found: true}
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_224.length, Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_227 := hx_structural_source_224
						return hxrt.BoundFunction(func(item any) {
							hx_receiver_227.__hx_this.add(item)
						}, func() any {
							return hx_receiver_227.__hx_this
						}, "add")
					}(), Found: true}
				case "push":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_228 := hx_structural_source_224
						return hxrt.BoundFunction(func(item any) {
							hx_receiver_228.__hx_this.push(item)
						}, func() any {
							return hx_receiver_228.__hx_this
						}, "push")
					}(), Found: true}
				case "first":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_229 := hx_structural_source_224
						return hxrt.BoundFunction(func() any {
							return hx_receiver_229.__hx_this.first()
						}, func() any {
							return hx_receiver_229.__hx_this
						}, "first")
					}(), Found: true}
				case "last":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_230 := hx_structural_source_224
						return hxrt.BoundFunction(func() any {
							return hx_receiver_230.__hx_this.last()
						}, func() any {
							return hx_receiver_230.__hx_this
						}, "last")
					}(), Found: true}
				case "pop":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_231 := hx_structural_source_224
						return hxrt.BoundFunction(func() any {
							return hx_receiver_231.__hx_this.pop()
						}, func() any {
							return hx_receiver_231.__hx_this
						}, "pop")
					}(), Found: true}
				case "isEmpty":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_232 := hx_structural_source_224
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_232.__hx_this.isEmpty()
						}, func() any {
							return hx_receiver_232.__hx_this
						}, "isEmpty")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_233 := hx_structural_source_224
						return hxrt.BoundFunction(func() {
							hx_receiver_233.__hx_this.clear()
						}, func() any {
							return hx_receiver_233.__hx_this
						}, "clear")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_234 := hx_structural_source_224
						return hxrt.BoundFunction(func(value any) bool {
							return hx_receiver_234.__hx_this.remove(value)
						}, func() any {
							return hx_receiver_234.__hx_this
						}, "remove")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds___List__GoListIterator] {
						hx_receiver_235 := hx_structural_source_224
						return hxrt.BoundFunction(func() *haxe__ds___List__GoListIterator {
							return hx_receiver_235.__hx_this.iterator()
						}, func() any {
							return hx_receiver_235.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator] {
						hx_receiver_236 := hx_structural_source_224
						return hxrt.BoundFunction(func() *haxe__ds___List__GoListKeyValueIterator {
							return hx_receiver_236.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_236.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_237 := hx_structural_source_224
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_237.__hx_this.toString()
						}, func() any {
							return hx_receiver_237.__hx_this
						}, "toString")
					}(), Found: true}
				case "join":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *string] {
						hx_receiver_238 := hx_structural_source_224
						return hxrt.BoundFunction(func(separator *string) *string {
							return hx_receiver_238.__hx_this.join(separator)
						}, func() any {
							return hx_receiver_238.__hx_this
						}, "join")
					}(), Found: true}
				case "filter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(any) bool]) *haxe__ds__List] {
						hx_receiver_239 := hx_structural_source_224
						return hxrt.BoundFunction(func(predicate *hxrt.Function[func(any) bool]) *haxe__ds__List {
							return hx_receiver_239.__hx_this.filter(predicate)
						}, func() any {
							return hx_receiver_239.__hx_this
						}, "filter")
					}(), Found: true}
				case "map":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.Function[func(any) any]) *haxe__ds__List] {
						hx_receiver_240 := hx_structural_source_224
						return hxrt.BoundFunction(func(transform *hxrt.Function[func(any) any]) *haxe__ds__List {
							return hx_receiver_240.__hx_this.map_(transform)
						}, func() any {
							return hx_receiver_240.__hx_this
						}, "map")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "items":
					hx_structural_source_224.items = func(hx_value_225 any) *hxrt.Array {
						if hx_value_225 == nil {
							var hx_zero_226 *hxrt.Array
							return hx_zero_226
						}
						return hx_value_225.(*hxrt.Array)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_241 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_241, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_241.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_241.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_241.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_241.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_244 := hx_structural_source_241
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_244.__hx_this.get_message()
						}, func() any {
							return hx_receiver_244.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_245 := hx_structural_source_241
						return hxrt.BoundFunction(func() any {
							return hx_receiver_245.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_245.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_246 := hx_structural_source_241
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_246.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_246.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_247 := hx_structural_source_241
						return hxrt.BoundFunction(func() any {
							return hx_receiver_247.__hx_this.get_native()
						}, func() any {
							return hx_receiver_247.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_248 := hx_structural_source_241
						return hxrt.BoundFunction(func() any {
							return hx_receiver_248.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_248.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_249 := hx_structural_source_241
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_249.__hx_this.toString()
						}, func() any {
							return hx_receiver_249.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_250 := hx_structural_source_241
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_250.__hx_this.details()
						}, func() any {
							return hx_receiver_250.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_241.savedMessage = func(hx_value_242 any) *string {
						if hx_value_242 == nil {
							var hx_zero_243 *string
							return hx_zero_243
						}
						return hx_value_242.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_241.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_241.savedPrevious = hx_recover_class_4(incoming)
					return true
				case "savedNative":
					hx_structural_source_241.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_251 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_251, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_251.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_252 := hx_structural_source_251
						return hxrt.BoundFunction(func() any {
							return hx_receiver_252.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_252.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_251.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_251.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_251.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_251.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_255 := hx_structural_source_251
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_255.__hx_this.get_message()
						}, func() any {
							return hx_receiver_255.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_256 := hx_structural_source_251
						return hxrt.BoundFunction(func() any {
							return hx_receiver_256.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_256.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_257 := hx_structural_source_251
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_257.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_257.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_258 := hx_structural_source_251
						return hxrt.BoundFunction(func() any {
							return hx_receiver_258.__hx_this.get_native()
						}, func() any {
							return hx_receiver_258.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_259 := hx_structural_source_251
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_259.__hx_this.toString()
						}, func() any {
							return hx_receiver_259.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_260 := hx_structural_source_251
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_260.__hx_this.details()
						}, func() any {
							return hx_receiver_260.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_251.savedMessage = func(hx_value_253 any) *string {
						if hx_value_253 == nil {
							var hx_zero_254 *string
							return hx_zero_254
						}
						return hx_value_253.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_251.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_251.savedPrevious = hx_recover_class_4(incoming)
					return true
				case "savedNative":
					hx_structural_source_251.savedNative = incoming
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
			return func(hx_value_261 any) bool {
				if hx_value_261 == nil {
					var hx_zero_262 bool
					return hx_zero_262
				}
				return hx_value_261.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_263 any) bool {
				if hx_value_263 == nil {
					var hx_zero_264 bool
					return hx_zero_264
				}
				return hx_value_263.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_265 any) bool {
				if hx_value_265 == nil {
					var hx_zero_266 bool
					return hx_zero_266
				}
				return hx_value_265.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *go___Result]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_267 any) bool {
				if hx_value_267 == nil {
					var hx_zero_268 bool
					return hx_zero_268
				}
				return hx_value_267.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_269 any) bool {
				if hx_value_269 == nil {
					var hx_zero_270 bool
					return hx_zero_270
				}
				return hx_value_269.(bool)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_271 any) bool {
				if hx_value_271 == nil {
					var hx_zero_272 bool
					return hx_zero_272
				}
				return hx_value_271.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_273 any) bool {
				if hx_value_273 == nil {
					var hx_zero_274 bool
					return hx_zero_274
				}
				return hx_value_273.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_275 any) bool {
				if hx_value_275 == nil {
					var hx_zero_276 bool
					return hx_zero_276
				}
				return hx_value_275.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_277 any) bool {
				if hx_value_277 == nil {
					var hx_zero_278 bool
					return hx_zero_278
				}
				return hx_value_277.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_279 any) bool {
				if hx_value_279 == nil {
					var hx_zero_280 bool
					return hx_zero_280
				}
				return hx_value_279.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_281 any) bool {
				if hx_value_281 == nil {
					var hx_zero_282 bool
					return hx_zero_282
				}
				return hx_value_281.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_283 any) bool {
				if hx_value_283 == nil {
					var hx_zero_284 bool
					return hx_zero_284
				}
				return hx_value_283.(bool)
			}(hxrt.FunctionCall(source)())
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
			return func(hx_value_285 any) *string {
				if hx_value_285 == nil {
					var hx_zero_286 *string
					return hx_zero_286
				}
				return hx_value_285.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_287 any) *string {
				if hx_value_287 == nil {
					var hx_zero_288 *string
					return hx_zero_288
				}
				return hx_value_287.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_289 any) *string {
				if hx_value_289 == nil {
					var hx_zero_290 *string
					return hx_zero_290
				}
				return hx_value_289.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *go___Result]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_291 any) *string {
				if hx_value_291 == nil {
					var hx_zero_292 *string
					return hx_zero_292
				}
				return hx_value_291.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_293 any) *string {
				if hx_value_293 == nil {
					var hx_zero_294 *string
					return hx_zero_294
				}
				return hx_value_293.(*string)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_295 any) *string {
				if hx_value_295 == nil {
					var hx_zero_296 *string
					return hx_zero_296
				}
				return hx_value_295.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.Array]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_297 any) *string {
				if hx_value_297 == nil {
					var hx_zero_298 *string
					return hx_zero_298
				}
				return hx_value_297.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_299 any) *string {
				if hx_value_299 == nil {
					var hx_zero_300 *string
					return hx_zero_300
				}
				return hx_value_299.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_301 any) *string {
				if hx_value_301 == nil {
					var hx_zero_302 *string
					return hx_zero_302
				}
				return hx_value_301.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_303 any) *string {
				if hx_value_303 == nil {
					var hx_zero_304 *string
					return hx_zero_304
				}
				return hx_value_303.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds___List__GoListKeyValueIterator]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_305 any) *string {
				if hx_value_305 == nil {
					var hx_zero_306 *string
					return hx_zero_306
				}
				return hx_value_305.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_307 any) *string {
				if hx_value_307 == nil {
					var hx_zero_308 *string
					return hx_zero_308
				}
				return hx_value_307.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *go___Result {
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

func hx_recover_class_1(value any) *haxe__ds__StringMap {
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

func hx_recover_class_2(value any) *haxe__ds___List__GoListIterator {
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

func hx_recover_class_3(value any) *haxe__ValueException {
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

func hx_recover_class_4(value any) *haxe__Exception {
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
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
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
