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
		return func(hx_structural_source_71 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_71, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_71.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_71.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_71.high = func(hx_value_72 any) int {
						if hx_value_72 == nil {
							var hx_zero_73 int
							return hx_zero_73
						}
						return hx_value_72.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_71.low = func(hx_value_74 any) int {
						if hx_value_74 == nil {
							var hx_zero_75 int
							return hx_zero_75
						}
						return hx_value_74.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_76 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_76, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_76.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_76.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_81 := hx_structural_source_76
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_81.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_82 := hx_structural_source_76
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_82.__hx_this.next()
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_76.offset = func(hx_value_77 any) int {
						if hx_value_77 == nil {
							var hx_zero_78 int
							return hx_zero_78
						}
						return hx_value_77.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_76.s = func(hx_value_79 any) *string {
						if hx_value_79 == nil {
							var hx_zero_80 *string
							return hx_zero_80
						}
						return hx_value_79.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_83 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_83, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_83.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_83.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_88 := hx_structural_source_83
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_88.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_89 := hx_structural_source_83
						return hxrt.BoundFunction(func() int {
							return hx_receiver_89.__hx_this.next()
						}, func() any {
							return hx_receiver_89.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_83.offset = func(hx_value_84 any) int {
						if hx_value_84 == nil {
							var hx_zero_85 int
							return hx_zero_85
						}
						return hx_value_84.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_83.s = func(hx_value_86 any) *string {
						if hx_value_86 == nil {
							var hx_zero_87 *string
							return hx_zero_87
						}
						return hx_value_86.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_90 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_90, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_90.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_90.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_90.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_90.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_90.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_99 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_99.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_100 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_100.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_101 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_101.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_101.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_102 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_102.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_103 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_103.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_104 := hx_structural_source_90
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_104.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_105 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_105.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_106 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_106.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_107 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_107.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_108 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_108.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_109 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_109.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_110 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_110.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_110.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_111 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_111.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_111.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_112 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_112.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_112.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_113 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_113.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_113.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_114 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_114.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_114.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_115 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_115.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_115.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_116 := hx_structural_source_90
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_116.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_117 := hx_structural_source_90
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_117.__hx_this.toString()
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_118 := hx_structural_source_90
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_118.__hx_this.toHex()
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_119 := hx_structural_source_90
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_119.__hx_this.getData()
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_120 := hx_structural_source_90
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_120.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_90.b = func(hx_value_91 any) []int {
						if hx_value_91 == nil {
							var hx_zero_92 []int
							return hx_zero_92
						}
						return hx_value_91.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_90.__hx_raw = func(hx_value_93 any) *hxrt.ByteView {
						if hx_value_93 == nil {
							var hx_zero_94 *hxrt.ByteView
							return hx_zero_94
						}
						return hx_value_93.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_90.__hx_rawValid = func(hx_value_95 any) bool {
						if hx_value_95 == nil {
							var hx_zero_96 bool
							return hx_zero_96
						}
						return hx_value_95.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_90.__hx_dataExposed = func(hx_value_97 any) bool {
						if hx_value_97 == nil {
							var hx_zero_98 bool
							return hx_zero_98
						}
						return hx_value_97.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_121 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_121, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_121.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_121.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_121.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_121.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_124 := hx_structural_source_121
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_124.__hx_this.get_message()
						}, func() any {
							return hx_receiver_124.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_125 := hx_structural_source_121
						return hxrt.BoundFunction(func() any {
							return hx_receiver_125.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_125.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_126 := hx_structural_source_121
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_126.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_126.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_127 := hx_structural_source_121
						return hxrt.BoundFunction(func() any {
							return hx_receiver_127.__hx_this.get_native()
						}, func() any {
							return hx_receiver_127.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_128 := hx_structural_source_121
						return hxrt.BoundFunction(func() any {
							return hx_receiver_128.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_128.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_129 := hx_structural_source_121
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_129.__hx_this.toString()
						}, func() any {
							return hx_receiver_129.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_130 := hx_structural_source_121
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_130.__hx_this.details()
						}, func() any {
							return hx_receiver_130.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_121.savedMessage = func(hx_value_122 any) *string {
						if hx_value_122 == nil {
							var hx_zero_123 *string
							return hx_zero_123
						}
						return hx_value_122.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_121.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_121.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_121.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_131 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_131, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_131.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_132 := hx_structural_source_131
						return hxrt.BoundFunction(func() any {
							return hx_receiver_132.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_131.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_131.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_131.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_131.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_135 := hx_structural_source_131
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_135.__hx_this.get_message()
						}, func() any {
							return hx_receiver_135.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_136 := hx_structural_source_131
						return hxrt.BoundFunction(func() any {
							return hx_receiver_136.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_136.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_137 := hx_structural_source_131
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_137.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_138 := hx_structural_source_131
						return hxrt.BoundFunction(func() any {
							return hx_receiver_138.__hx_this.get_native()
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_139 := hx_structural_source_131
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_139.__hx_this.toString()
						}, func() any {
							return hx_receiver_139.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_140 := hx_structural_source_131
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_140.__hx_this.details()
						}, func() any {
							return hx_receiver_140.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_131.savedMessage = func(hx_value_133 any) *string {
						if hx_value_133 == nil {
							var hx_zero_134 *string
							return hx_zero_134
						}
						return hx_value_133.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_131.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_131.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_131.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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
