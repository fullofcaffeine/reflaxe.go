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

func main() {
	method := hxrt.StringFromLiteral("POST")
	status := 200
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("method="), method)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("status="), status)))
	base := New_haxe__http__HttpBase(hxrt.StringFromLiteral("http://example.com"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("responseData="), hxrt.StdString(base.__hx_this.get_responseData())))
	hxrt.Println(v)
	hxrt.TryCatch(func() {
		base.__hx_this.request(false)
		hxrt.Println(any(hxrt.StringFromLiteral("request=no_throw")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var e any = hx_payload_3
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("requestType="), Type_getClassName(Type_getClass(e))))
		hxrt.Println(v_1)
	})
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_6 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_6.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_6.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_6.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_6.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_6.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_6.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_6.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_6.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_17 := hx_structural_source_6
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_17.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_18 := hx_structural_source_6
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_18.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_19 := hx_structural_source_6
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_19.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_20 := hx_structural_source_6
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_20.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_21 := hx_structural_source_6
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_21.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_22 := hx_structural_source_6
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_22.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_23 := hx_structural_source_6
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_23.__hx_this.request(post)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_6.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_6.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_6.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_6.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_24 := hx_structural_source_6
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_24.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_25 := hx_structural_source_6
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_25.__hx_this.success(data)
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_26 := hx_structural_source_6
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_26.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_6.url = func(hx_value_7 any) *string {
						if hx_value_7 == nil {
							var hx_zero_8 *string
							return hx_zero_8
						}
						return hx_value_7.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_6.responseAsString = func(hx_value_9 any) *string {
						if hx_value_9 == nil {
							var hx_zero_10 *string
							return hx_zero_10
						}
						return hx_value_9.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_6.postData = func(hx_value_11 any) *string {
						if hx_value_11 == nil {
							var hx_zero_12 *string
							return hx_zero_12
						}
						return hx_value_11.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_6.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_6.headers = func(hx_value_13 any) *hxrt.Array {
						if hx_value_13 == nil {
							var hx_zero_14 *hxrt.Array
							return hx_zero_14
						}
						return hx_value_13.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_6.params = func(hx_value_15 any) *hxrt.Array {
						if hx_value_15 == nil {
							var hx_zero_16 *hxrt.Array
							return hx_zero_16
						}
						return hx_value_15.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_6.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_6.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_6.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_6.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_27 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_27, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_27.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_27.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_27.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_27.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_30 := hx_structural_source_27
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_30.__hx_this.get_message()
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_31 := hx_structural_source_27
						return hxrt.BoundFunction(func() any {
							return hx_receiver_31.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_32 := hx_structural_source_27
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_32.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_33 := hx_structural_source_27
						return hxrt.BoundFunction(func() any {
							return hx_receiver_33.__hx_this.get_native()
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_34 := hx_structural_source_27
						return hxrt.BoundFunction(func() any {
							return hx_receiver_34.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_35 := hx_structural_source_27
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_35.__hx_this.toString()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_36 := hx_structural_source_27
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_36.__hx_this.details()
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_27.savedMessage = func(hx_value_28 any) *string {
						if hx_value_28 == nil {
							var hx_zero_29 *string
							return hx_zero_29
						}
						return hx_value_28.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_27.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_27.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_27.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_37 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_37, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_37.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_37.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_37.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_37.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_37.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_46 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_46.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_46.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_47 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_47.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_47.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_48 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
							hx_receiver_48.__hx_this.blit(pos, src, srcpos, len_local)
						}, func() any {
							return hx_receiver_48.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_49 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, len_local int, value int) {
							hx_receiver_49.__hx_this.fill(pos, len_local, value)
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_50 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
							return hx_receiver_50.__hx_this.sub(pos, len_local)
						}, func() any {
							return hx_receiver_50.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_51 := hx_structural_source_37
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_51.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_52 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_52.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_53 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_53.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_54 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_54.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_55 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_55.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_56 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_56.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_57 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_57.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_58 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_58.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_59 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_59.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_60 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_60.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_61 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_61.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_62 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_62.__hx_this.getString(pos, len_local, encoding)
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_63 := hx_structural_source_37
						return hxrt.BoundFunction(func(pos int, len_local int) *string {
							return hx_receiver_63.__hx_this.readString(pos, len_local)
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_64 := hx_structural_source_37
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_64.__hx_this.toString()
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_65 := hx_structural_source_37
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_65.__hx_this.toHex()
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_66 := hx_structural_source_37
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_66.__hx_this.getData()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_67 := hx_structural_source_37
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_67.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_37.b = func(hx_value_38 any) []int {
						if hx_value_38 == nil {
							var hx_zero_39 []int
							return hx_zero_39
						}
						return hx_value_38.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_37.__hx_raw = func(hx_value_40 any) *hxrt.ByteView {
						if hx_value_40 == nil {
							var hx_zero_41 *hxrt.ByteView
							return hx_zero_41
						}
						return hx_value_40.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_37.__hx_rawValid = func(hx_value_42 any) bool {
						if hx_value_42 == nil {
							var hx_zero_43 bool
							return hx_zero_43
						}
						return hx_value_42.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_37.__hx_dataExposed = func(hx_value_44 any) bool {
						if hx_value_44 == nil {
							var hx_zero_45 bool
							return hx_zero_45
						}
						return hx_value_44.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_68 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_68, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_68.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_69 := hx_structural_source_68
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_69.__hx_this.toString()
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_68.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_68.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_68.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_68.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_72 := hx_structural_source_68
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_72.__hx_this.get_message()
						}, func() any {
							return hx_receiver_72.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_73 := hx_structural_source_68
						return hxrt.BoundFunction(func() any {
							return hx_receiver_73.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_73.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_74 := hx_structural_source_68
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_74.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_75 := hx_structural_source_68
						return hxrt.BoundFunction(func() any {
							return hx_receiver_75.__hx_this.get_native()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_76 := hx_structural_source_68
						return hxrt.BoundFunction(func() any {
							return hx_receiver_76.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_77 := hx_structural_source_68
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_77.__hx_this.details()
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_68.savedMessage = func(hx_value_70 any) *string {
						if hx_value_70 == nil {
							var hx_zero_71 *string
							return hx_zero_71
						}
						return hx_value_70.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_68.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_68.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_68.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_78 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_78, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_78.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_79 := hx_structural_source_78
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_79.__hx_this.toString()
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_78.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_78.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_78.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_78.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_82 := hx_structural_source_78
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_82.__hx_this.get_message()
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_83 := hx_structural_source_78
						return hxrt.BoundFunction(func() any {
							return hx_receiver_83.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_83.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_84 := hx_structural_source_78
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_84.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_85 := hx_structural_source_78
						return hxrt.BoundFunction(func() any {
							return hx_receiver_85.__hx_this.get_native()
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_86 := hx_structural_source_78
						return hxrt.BoundFunction(func() any {
							return hx_receiver_86.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_87 := hx_structural_source_78
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_87.__hx_this.details()
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_78.savedMessage = func(hx_value_80 any) *string {
						if hx_value_80 == nil {
							var hx_zero_81 *string
							return hx_zero_81
						}
						return hx_value_80.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_78.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_78.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_78.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_88 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_88, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_88.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_89 := hx_structural_source_88
						return hxrt.BoundFunction(func() any {
							return hx_receiver_89.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_88.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_88.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_88.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_88.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_92 := hx_structural_source_88
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_92.__hx_this.get_message()
						}, func() any {
							return hx_receiver_92.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_93 := hx_structural_source_88
						return hxrt.BoundFunction(func() any {
							return hx_receiver_93.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_93.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_94 := hx_structural_source_88
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_94.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_94.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_95 := hx_structural_source_88
						return hxrt.BoundFunction(func() any {
							return hx_receiver_95.__hx_this.get_native()
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_96 := hx_structural_source_88
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_96.__hx_this.toString()
						}, func() any {
							return hx_receiver_96.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_97 := hx_structural_source_88
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_97.__hx_this.details()
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_88.savedMessage = func(hx_value_90 any) *string {
						if hx_value_90 == nil {
							var hx_zero_91 *string
							return hx_zero_91
						}
						return hx_value_90.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_88.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_88.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_88.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

type Type struct {
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

func hxrt_typeCreateEmpty__haxe__exceptions__NotImplementedException() *haxe__exceptions__NotImplementedException {
	instance := &haxe__exceptions__NotImplementedException{}
	instance.haxe__exceptions__PosException = &haxe__exceptions__PosException{}
	instance.haxe__exceptions__PosException.haxe__Exception = &haxe__Exception{}
	instance.haxe__exceptions__PosException.__hx_this = instance
	instance.haxe__exceptions__PosException.haxe__Exception.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__exceptions__PosException() *haxe__exceptions__PosException {
	instance := &haxe__exceptions__PosException{}
	instance.haxe__Exception = &haxe__Exception{}
	instance.haxe__Exception.__hx_this = instance
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__http__HttpBase() *haxe__http__HttpBase {
	instance := &haxe__http__HttpBase{}
	instance.__hx_this = instance
	return instance
}

func hxrt_typeCreateEmpty__haxe__io__Bytes() *haxe__io__Bytes {
	instance := &haxe__io__Bytes{}
	instance.__hx_this = instance
	return instance
}

type ValueType struct {
	tag    int
	params []any
}

func (self *ValueType) HxEnumValue() (string, []any) {
	return [9]string{"TNull", "TInt", "TFloat", "TBool", "TObject", "TFunction", "TClass", "TEnum", "TUnknown"}[self.tag], self.params
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
	case "Main":
		return nil, false
	case "StringBuf":
		return nil, false
	case "haxe.Exception":
		return hxrt_typeCallAny(New_haxe__Exception, args)
	case "haxe.Int64Helper":
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
	case "haxe.exceptions.NotImplementedException":
		return hxrt_typeCallAny(New_haxe__exceptions__NotImplementedException, args)
	case "haxe.exceptions.PosException":
		return hxrt_typeCallAny(New_haxe__exceptions__PosException, args)
	case "haxe.http.HttpBase":
		return hxrt_typeCallAny(New_haxe__http__HttpBase, args)
	case "haxe.io.Bytes":
		return hxrt_typeCallAny(New_haxe__io__Bytes, args)
	case "haxe.io.FPHelper":
		return nil, false
	default:
		return nil, false
	}
}

func hxrt_typeCreateClassEmptyInstance(className string) (any, bool) {
	switch className {
	case "haxe.Exception":
		return hxrt_typeCreateEmpty__haxe__Exception(), true
	case "haxe.ValueException":
		return hxrt_typeCreateEmpty__haxe__ValueException(), true
	case "haxe._Int64.___Int64":
		return hxrt_typeCreateEmpty__haxe___Int64_____Int64(), true
	case "haxe.exceptions.NotImplementedException":
		return hxrt_typeCreateEmpty__haxe__exceptions__NotImplementedException(), true
	case "haxe.exceptions.PosException":
		return hxrt_typeCreateEmpty__haxe__exceptions__PosException(), true
	case "haxe.http.HttpBase":
		return hxrt_typeCreateEmpty__haxe__http__HttpBase(), true
	case "haxe.io.Bytes":
		return hxrt_typeCreateEmpty__haxe__io__Bytes(), true
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
	case "haxe.io.Encoding":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_UTF8, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Encoding_RawNative, true
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "UTF8":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_UTF8, true
		case "RawNative":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Encoding_RawNative, true
		default:
			return nil, false
		}
	case "haxe.io.Error":
		if useIndex {
			switch constructorIndex {
			case 0:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Blocked, true
			case 1:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_Overflow, true
			case 2:
				if len(args) != 0 {
					return nil, false
				}
				return haxe__io__Error_OutsideBounds, true
			case 3:
				if len(args) != 1 {
					return nil, false
				}
				return hxrt_typeCallAny(haxe__io__Error_Custom, args)
			default:
				return nil, false
			}
		}
		switch constructorName {
		case "Blocked":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Blocked, true
		case "Overflow":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_Overflow, true
		case "OutsideBounds":
			if len(args) != 0 {
				return nil, false
			}
			return haxe__io__Error_OutsideBounds, true
		case "Custom":
			if len(args) != 1 {
				return nil, false
			}
			return hxrt_typeCallAny(haxe__io__Error_Custom, args)
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
	case *haxe__exceptions__NotImplementedException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.NotImplementedException")}
	case *haxe__exceptions__PosException:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case *haxe__http__HttpBase:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.http.HttpBase")}
	case *haxe__io__Bytes:
		if value == nil {
			return nil
		}
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.io.Bytes")}
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
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Encoding")}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.io.Error")}
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
	case "Main":
		return nil
	case "StringBuf":
		return nil
	case "haxe.Exception":
		return nil
	case "haxe.Int64Helper":
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
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}
	case "haxe.http.HttpBase":
		return nil
	case "haxe.io.Bytes":
		return nil
	case "haxe.io.FPHelper":
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
	case "Main":
		return hxrt.NewArray(hxrt.StringFromLiteral("main"))
	case "StringBuf":
		return hxrt.NewArray()
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("catchValue"), hxrt.StringFromLiteral("caught"), hxrt.StringFromLiteral("thrown"))
	case "haxe.Int64Helper":
		return hxrt.NewArray()
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
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray()
	case "haxe.exceptions.PosException":
		return hxrt.NewArray()
	case "haxe.http.HttpBase":
		return hxrt.NewArray()
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_fromNativeView"), hxrt.StringFromLiteral("alloc"), hxrt.StringFromLiteral("fastGet"), hxrt.StringFromLiteral("ofData"), hxrt.StringFromLiteral("ofHex"), hxrt.StringFromLiteral("ofString"), hxrt.StringFromLiteral("rawNativeUsesUtf16LE"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray(hxrt.StringFromLiteral("doubleToI64"), hxrt.StringFromLiteral("floatToI32"), hxrt.StringFromLiteral("i32ToFloat"), hxrt.StringFromLiteral("i64ToDouble"))
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
	case "Main":
		return hxrt.NewArray()
	case "StringBuf":
		return hxrt.NewArray()
	case "haxe.Exception":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.Int64Helper":
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
	case "haxe.exceptions.NotImplementedException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.exceptions.PosException":
		return hxrt.NewArray(hxrt.StringFromLiteral("details"), hxrt.StringFromLiteral("get_message"), hxrt.StringFromLiteral("get_native"), hxrt.StringFromLiteral("get_previous"), hxrt.StringFromLiteral("get_stack"), hxrt.StringFromLiteral("message"), hxrt.StringFromLiteral("native"), hxrt.StringFromLiteral("posInfos"), hxrt.StringFromLiteral("previous"), hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("stack"), hxrt.StringFromLiteral("toString"), hxrt.StringFromLiteral("unwrap"))
	case "haxe.http.HttpBase":
		return hxrt.NewArray(hxrt.StringFromLiteral("addHeader"), hxrt.StringFromLiteral("addParameter"), hxrt.StringFromLiteral("emptyOnData"), hxrt.StringFromLiteral("get_responseData"), hxrt.StringFromLiteral("hasOnData"), hxrt.StringFromLiteral("headers"), hxrt.StringFromLiteral("onBytes"), hxrt.StringFromLiteral("onData"), hxrt.StringFromLiteral("onError"), hxrt.StringFromLiteral("onStatus"), hxrt.StringFromLiteral("params"), hxrt.StringFromLiteral("postBytes"), hxrt.StringFromLiteral("postData"), hxrt.StringFromLiteral("request"), hxrt.StringFromLiteral("responseAsString"), hxrt.StringFromLiteral("responseBytes"), hxrt.StringFromLiteral("responseData"), hxrt.StringFromLiteral("setHeader"), hxrt.StringFromLiteral("setParameter"), hxrt.StringFromLiteral("setPostBytes"), hxrt.StringFromLiteral("setPostData"), hxrt.StringFromLiteral("success"), hxrt.StringFromLiteral("url"))
	case "haxe.io.Bytes":
		return hxrt.NewArray(hxrt.StringFromLiteral("__hx_dataExposed"), hxrt.StringFromLiteral("__hx_nativeView"), hxrt.StringFromLiteral("__hx_raw"), hxrt.StringFromLiteral("__hx_rawValid"), hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("blit"), hxrt.StringFromLiteral("compare"), hxrt.StringFromLiteral("fill"), hxrt.StringFromLiteral("get"), hxrt.StringFromLiteral("getData"), hxrt.StringFromLiteral("getDouble"), hxrt.StringFromLiteral("getFloat"), hxrt.StringFromLiteral("getInt32"), hxrt.StringFromLiteral("getInt64"), hxrt.StringFromLiteral("getString"), hxrt.StringFromLiteral("getUInt16"), hxrt.StringFromLiteral("length"), hxrt.StringFromLiteral("readString"), hxrt.StringFromLiteral("set"), hxrt.StringFromLiteral("setDouble"), hxrt.StringFromLiteral("setFloat"), hxrt.StringFromLiteral("setInt32"), hxrt.StringFromLiteral("setInt64"), hxrt.StringFromLiteral("setUInt16"), hxrt.StringFromLiteral("sub"), hxrt.StringFromLiteral("toHex"), hxrt.StringFromLiteral("toString"))
	case "haxe.io.FPHelper":
		return hxrt.NewArray()
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
	case "Main":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "StringBuf":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Exception":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.Int64Helper":
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
	case "haxe.exceptions.NotImplementedException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.exceptions.PosException":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.http.HttpBase":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Bytes":
		return &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.FPHelper":
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
	case "haxe.io.Encoding":
		return &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}
	case "haxe.io.Error":
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
	case *haxe__io__Encoding:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("UTF8")
		case 1:
			return hxrt.StringFromLiteral("RawNative")
		default:
			return nil
		}
	case *haxe__io__Error:
		if value == nil {
			return nil
		}
		switch value.tag {
		case 0:
			return hxrt.StringFromLiteral("Blocked")
		case 1:
			return hxrt.StringFromLiteral("Overflow")
		case 2:
			return hxrt.StringFromLiteral("OutsideBounds")
		case 3:
			return hxrt.StringFromLiteral("Custom")
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
	case *haxe__io__Encoding:
		if value == nil {
			return -1
		}
		return value.tag
	case *haxe__io__Error:
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
	case "haxe.io.Encoding":
		return hxrt.NewArray(hxrt.StringFromLiteral("UTF8"), hxrt.StringFromLiteral("RawNative"))
	case "haxe.io.Error":
		return hxrt.NewArray(hxrt.StringFromLiteral("Blocked"), hxrt.StringFromLiteral("Overflow"), hxrt.StringFromLiteral("OutsideBounds"), hxrt.StringFromLiteral("Custom"))
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
	case *haxe__io__Encoding:
		if value == nil || value.params == nil {
			return hxrt.NewArray()
		}
		return hxrt.NewArray(value.params...)
	case *haxe__io__Error:
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
	case "haxe.io.Encoding":
		return hxrt.NewArray(haxe__io__Encoding_UTF8, haxe__io__Encoding_RawNative)
	case "haxe.io.Error":
		return hxrt.NewArray(haxe__io__Error_Blocked, haxe__io__Error_Overflow, haxe__io__Error_OutsideBounds)
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

func hx_recover_callable_9da3ad96(value any) *hxrt.Function[func(*string)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*string)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5f1b802(value any) *hxrt.Function[func(*haxe__io__Bytes)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_d8be58e9(value any) *hxrt.Function[func(int)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int)]:
		return source
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
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
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
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

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
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
