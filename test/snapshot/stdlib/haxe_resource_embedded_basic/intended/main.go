package main

import "snapshot/hxrt"

func main() {
	names := haxe__Resource_listNames()
	var v any = any(names.Len())
	hxrt.Println(v)
	var hx_if_1 any
	if names.Len() > 0 {
		hx_if_1 = names.Get(0)
	} else {
		hx_if_1 = hxrt.StringFromLiteral("<none>")
	}
	var v_1 any = hx_if_1
	hxrt.Println(v_1)
	var v_2 any = any(StringTools_replace(haxe__Resource_getString(hxrt.StringFromLiteral("greet")), hxrt.StringFromLiteral("\n"), hxrt.StringFromLiteral("\\n")))
	hxrt.Println(v_2)
	bytes := haxe__Resource_getBytes(hxrt.StringFromLiteral("greet"))
	var hx_if_2 any
	if bytes == nil {
		hx_if_2 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_2 = bytes.length
	}
	var v_3 any = hx_if_2
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringEqualStringPtr(haxe__Resource_getString(hxrt.StringFromLiteral("missing")), nil))
	hxrt.Println(v_4)
	var v_5 any = any((haxe__Resource_getBytes(hxrt.StringFromLiteral("missing")) == nil))
	hxrt.Println(v_5)
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
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_6 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_6.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_6.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_11 := hx_structural_source_6
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_11.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_12 := hx_structural_source_6
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_12.__hx_this.next()
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_6.offset = func(hx_value_7 any) int {
						if hx_value_7 == nil {
							var hx_zero_8 int
							return hx_zero_8
						}
						return hx_value_7.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_6.s = func(hx_value_9 any) *string {
						if hx_value_9 == nil {
							var hx_zero_10 *string
							return hx_zero_10
						}
						return hx_value_9.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_13 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_13, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_13.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_13.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_18 := hx_structural_source_13
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_18.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_19 := hx_structural_source_13
						return hxrt.BoundFunction(func() int {
							return hx_receiver_19.__hx_this.next()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_13.offset = func(hx_value_14 any) int {
						if hx_value_14 == nil {
							var hx_zero_15 int
							return hx_zero_15
						}
						return hx_value_14.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_13.s = func(hx_value_16 any) *string {
						if hx_value_16 == nil {
							var hx_zero_17 *string
							return hx_zero_17
						}
						return hx_value_16.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_20 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_20, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_20.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_20.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_20.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_20.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_20.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_29 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_29.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_29.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_30 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_30.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_31 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_31.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_32 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_32.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_33 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_33.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_34 := hx_structural_source_20
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_34.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_35 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_35.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_36 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_36.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_37 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_37.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_38 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_38.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_39 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_39.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_40 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_40.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_41 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_41.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_42 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_42.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_43 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_43.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_44 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_44.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_45 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_45.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_46 := hx_structural_source_20
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_46.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_46.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_47 := hx_structural_source_20
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_47.__hx_this.toString()
						}, func() any {
							return hx_receiver_47.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_48 := hx_structural_source_20
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_48.__hx_this.toHex()
						}, func() any {
							return hx_receiver_48.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_49 := hx_structural_source_20
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_49.__hx_this.getData()
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_50 := hx_structural_source_20
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_50.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_50.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_20.b = func(hx_value_21 any) []int {
						if hx_value_21 == nil {
							var hx_zero_22 []int
							return hx_zero_22
						}
						return hx_value_21.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_20.__hx_raw = func(hx_value_23 any) *hxrt.ByteView {
						if hx_value_23 == nil {
							var hx_zero_24 *hxrt.ByteView
							return hx_zero_24
						}
						return hx_value_23.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_20.__hx_rawValid = func(hx_value_25 any) bool {
						if hx_value_25 == nil {
							var hx_zero_26 bool
							return hx_zero_26
						}
						return hx_value_25.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_20.__hx_dataExposed = func(hx_value_27 any) bool {
						if hx_value_27 == nil {
							var hx_zero_28 bool
							return hx_zero_28
						}
						return hx_value_27.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_51 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_51, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_51.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_51.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_51.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_51.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_54 := hx_structural_source_51
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_54.__hx_this.get_message()
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_55 := hx_structural_source_51
						return hxrt.BoundFunction(func() any {
							return hx_receiver_55.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_56 := hx_structural_source_51
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_56.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_57 := hx_structural_source_51
						return hxrt.BoundFunction(func() any {
							return hx_receiver_57.__hx_this.get_native()
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_58 := hx_structural_source_51
						return hxrt.BoundFunction(func() any {
							return hx_receiver_58.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_59 := hx_structural_source_51
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_59.__hx_this.toString()
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_60 := hx_structural_source_51
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_60.__hx_this.details()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_51.savedMessage = func(hx_value_52 any) *string {
						if hx_value_52 == nil {
							var hx_zero_53 *string
							return hx_zero_53
						}
						return hx_value_52.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_51.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_51.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_51.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_61 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_61, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_61.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_62 := hx_structural_source_61
						return hxrt.BoundFunction(func() any {
							return hx_receiver_62.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_61.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_61.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_61.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_61.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_65 := hx_structural_source_61
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_65.__hx_this.get_message()
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_66 := hx_structural_source_61
						return hxrt.BoundFunction(func() any {
							return hx_receiver_66.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_67 := hx_structural_source_61
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_67.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_68 := hx_structural_source_61
						return hxrt.BoundFunction(func() any {
							return hx_receiver_68.__hx_this.get_native()
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_69 := hx_structural_source_61
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_69.__hx_this.toString()
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_70 := hx_structural_source_61
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_70.__hx_this.details()
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_61.savedMessage = func(hx_value_63 any) *string {
						if hx_value_63 == nil {
							var hx_zero_64 *string
							return hx_zero_64
						}
						return hx_value_63.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_61.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_61.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_61.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
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
