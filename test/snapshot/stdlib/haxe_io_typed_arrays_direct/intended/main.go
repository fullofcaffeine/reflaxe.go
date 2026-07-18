package main

import "snapshot/hxrt"

func errTag(err *haxe__io__Error) *string {
	var hx_switch_1 *string
	switch err.tag {
	case 0:
		hx_switch_1 = hxrt.StringFromLiteral("blocked")
	case 1:
		hx_switch_1 = hxrt.StringFromLiteral("overflow")
	case 2:
		hx_switch_1 = hxrt.StringFromLiteral("outside")
	case 3:
		var _g any = err.params[0]
		var v any = _g
		hx_switch_1 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom:"), hxrt.StdString(v))
	}
	return hx_switch_1
}

func fmtFloat(v float64) *string {
	return hxrt.StdString((float64(hxrt.MathRoundInt((v * float64(100)))) / float64(100)))
}

func main() {
	bytes := haxe__io__Bytes_alloc(12)
	_g := 0
	_g1 := bytes.length
	for _g < _g1 {
		hx_post_2 := _g
		_g = int(int32((_g + 1)))
		i := hx_post_2
		bytes.b[i] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	view := haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, 2, 6)
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("view="), view.byteOffset), hxrt.StringFromLiteral(",")), view.byteLength), hxrt.StringFromLiteral(",")), view.bytes.length))
	hxrt.Println(v)
	a := view.__hx_this.sub(1, 3)
	sub := a
	a_1 := view.__hx_this.subarray(2, 5)
	subarray := a_1
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("view.sub="), sub.byteOffset), hxrt.StringFromLiteral(",")), sub.byteLength))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("view.subarray="), subarray.byteOffset), hxrt.StringFromLiteral(",")), subarray.byteLength))
	hxrt.Println(v_2)
	u8 := haxe__io___UInt8Array__UInt8Array_Impl__fromBytes(bytes, 1, 4)
	if 1 < u8.byteLength {
		_this := u8.bytes
		pos := int(int32((hxrt.Int32Wrap(1) + hxrt.Int32Wrap(u8.byteOffset))))
		_this.b[pos] = 99
		_this.__hx_rawValid = false
		_ = 99
	} else {
		_ = 0
	}
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("u8="), u8.byteLength), hxrt.StringFromLiteral(",")), u8.bytes.b[u8.byteOffset]), hxrt.StringFromLiteral(",")), u8.bytes.b[int(int32((hxrt.Int32Wrap(1)+hxrt.Int32Wrap(u8.byteOffset))))]), hxrt.StringFromLiteral(",")), bytes.b[2]))
	hxrt.Println(v_3)
	u16 := haxe__io___UInt16Array__UInt16Array_Impl__fromArray(hxrt.NewArray(4660, 255, 51966), 0, 3)
	var scaledLength any = 4
	u16sub := haxe__io___UInt16Array__UInt16Array_Impl__fromData(u16.__hx_this.sub(2, scaledLength))
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("u16="), int(int32((hxrt.Int32Wrap(u16.byteLength)>>uint(1))))), hxrt.StringFromLiteral(",")), func() int {
		_this_1 := u16.bytes
		pos_1 := u16.byteOffset
		return int(int32((hxrt.Int32Wrap(_this_1.b[pos_1]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_1.b[int(int32((hxrt.Int32Wrap(pos_1)+hxrt.Int32Wrap(1))))]) << uint(8))))))))
	}()), hxrt.StringFromLiteral(",")), int(int32((hxrt.Int32Wrap(u16sub.byteLength)>>uint(1))))), hxrt.StringFromLiteral(",")), func() int {
		_this_2 := u16sub.bytes
		pos_2 := u16sub.byteOffset
		return int(int32((hxrt.Int32Wrap(_this_2.b[pos_2]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_2.b[int(int32((hxrt.Int32Wrap(pos_2)+hxrt.Int32Wrap(1))))]) << uint(8))))))))
	}()))
	hxrt.Println(v_4)
	u32 := haxe__io___UInt32Array__UInt32Array_Impl__fromArray(hxrt.NewArray(1, 2, 3), 0, 3)
	var scaledBegin any = 4
	var scaledEnd any = 12
	u32view := haxe__io___UInt32Array__UInt32Array_Impl__fromData(u32.__hx_this.subarray(scaledBegin, scaledEnd))
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("u32="), int(int32((hxrt.Int32Wrap(u32.byteLength)>>uint(2))))), hxrt.StringFromLiteral(",")), func() *string {
		this1 := func() int {
			_this_3 := u32.bytes
			pos_3 := int(int32((hxrt.Int32Wrap(8) + hxrt.Int32Wrap(u32.byteOffset))))
			return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_3.b[pos_3]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_3.b[int(int32((hxrt.Int32Wrap(pos_3)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_3.b[int(int32((hxrt.Int32Wrap(pos_3)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_3.b[int(int32((hxrt.Int32Wrap(pos_3)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
		}()
		return hxrt.StdString(func() float64 {
			int := this1
			var hx_if_3 float64
			if int < 0 {
				hx_if_3 = (4294967296.0 + float64(int))
			} else {
				hx_if_3 = (float64(int) + 0.0)
			}
			return hx_if_3
		}())
	}()), hxrt.StringFromLiteral(",")), int(int32((hxrt.Int32Wrap(u32view.byteLength)>>uint(2))))), hxrt.StringFromLiteral(",")), func() *string {
		this1_1 := func() int {
			_this_4 := u32view.bytes
			pos_4 := u32view.byteOffset
			return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_4.b[pos_4]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_4.b[int(int32((hxrt.Int32Wrap(pos_4)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_4.b[int(int32((hxrt.Int32Wrap(pos_4)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_4.b[int(int32((hxrt.Int32Wrap(pos_4)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
		}()
		return hxrt.StdString(func() float64 {
			int_1 := this1_1
			var hx_if_4 float64
			if int_1 < 0 {
				hx_if_4 = (4294967296.0 + float64(int_1))
			} else {
				hx_if_4 = (float64(int_1) + 0.0)
			}
			return hx_if_4
		}())
	}()))
	hxrt.Println(v_5)
	i32 := haxe__io___Int32Array__Int32Array_Impl__fromArray(hxrt.NewArray(-7, 42), 0, 2)
	if 1 < int(int32((hxrt.Int32Wrap(i32.byteLength) >> uint(2)))) {
		_this_5 := i32.bytes
		pos_5 := int(int32((hxrt.Int32Wrap(4) + hxrt.Int32Wrap(i32.byteOffset))))
		_this_5.b[pos_5] = 156
		_this_5.__hx_rawValid = false
		_this_5.b[int(int32((hxrt.Int32Wrap(pos_5) + hxrt.Int32Wrap(1))))] = 255
		_this_5.__hx_rawValid = false
		_this_5.b[int(int32((hxrt.Int32Wrap(pos_5) + hxrt.Int32Wrap(2))))] = 255
		_this_5.__hx_rawValid = false
		_this_5.b[int(int32((hxrt.Int32Wrap(pos_5) + hxrt.Int32Wrap(3))))] = 255
		_this_5.__hx_rawValid = false
		_ = -100
	} else {
		_ = 0
	}
	var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("i32="), int(int32((hxrt.Int32Wrap(i32.byteLength)>>uint(2))))), hxrt.StringFromLiteral(",")), func() int {
		_this_6 := i32.bytes
		pos_6 := i32.byteOffset
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_6.b[pos_6]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_6.b[int(int32((hxrt.Int32Wrap(pos_6)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_6.b[int(int32((hxrt.Int32Wrap(pos_6)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_6.b[int(int32((hxrt.Int32Wrap(pos_6)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	}()), hxrt.StringFromLiteral(",")), func() int {
		_this_7 := i32.bytes
		pos_7 := int(int32((hxrt.Int32Wrap(4) + hxrt.Int32Wrap(i32.byteOffset))))
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_7.b[pos_7]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_7.b[int(int32((hxrt.Int32Wrap(pos_7)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_7.b[int(int32((hxrt.Int32Wrap(pos_7)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_7.b[int(int32((hxrt.Int32Wrap(pos_7)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	}()))
	hxrt.Println(v_6)
	var this1_2 *haxe__io__ArrayBufferViewImpl
	size := 8
	var this1_4 *haxe__io__ArrayBufferViewImpl
	this1_4 = New_haxe__io__ArrayBufferViewImpl(haxe__io__Bytes_alloc(size), 0, size)
	this1_3 := this1_4
	this1_2 = this1_3
	f32 := this1_2
	if 0 < int(int32((hxrt.Int32Wrap(f32.byteLength) >> uint(2)))) {
		_this_8 := f32.bytes
		pos_8 := f32.byteOffset
		value := haxe__io__FPHelper_floatToI32(1.25)
		_this_8.b[pos_8] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		_this_8.__hx_rawValid = false
		_this_8.b[int(int32((hxrt.Int32Wrap(pos_8) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this_8.__hx_rawValid = false
		_this_8.b[int(int32((hxrt.Int32Wrap(pos_8) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this_8.__hx_rawValid = false
		_this_8.b[int(int32((hxrt.Int32Wrap(pos_8) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		_this_8.__hx_rawValid = false
		_ = 1.25
	} else {
		_ = 0
	}
	if 1 < int(int32((hxrt.Int32Wrap(f32.byteLength) >> uint(2)))) {
		_this_9 := f32.bytes
		pos_9 := int(int32((hxrt.Int32Wrap(4) + hxrt.Int32Wrap(f32.byteOffset))))
		value_1 := haxe__io__FPHelper_floatToI32(2.5)
		_this_9.b[pos_9] = int(int32((hxrt.Int32Wrap(value_1) & hxrt.Int32Wrap(255))))
		_this_9.__hx_rawValid = false
		_this_9.b[int(int32((hxrt.Int32Wrap(pos_9) + hxrt.Int32Wrap(1))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(8))))) & hxrt.Int32Wrap(255))))
		_this_9.__hx_rawValid = false
		_this_9.b[int(int32((hxrt.Int32Wrap(pos_9) + hxrt.Int32Wrap(2))))] = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value_1) >> uint(16))))) & hxrt.Int32Wrap(255))))
		_this_9.__hx_rawValid = false
		_this_9.b[int(int32((hxrt.Int32Wrap(pos_9) + hxrt.Int32Wrap(3))))] = int(int32((hxrt.Int32Wrap(int(int32(int32((uint32(hxrt.Int32Wrap(value_1)) >> uint(24)))))) & hxrt.Int32Wrap(255))))
		_this_9.__hx_rawValid = false
		_ = 2.5
	} else {
		_ = 0
	}
	var scaledBegin_1 any = 4
	var scaledEnd_1 any = 8
	f32sub := haxe__io___Float32Array__Float32Array_Impl__fromData(f32.__hx_this.subarray(scaledBegin_1, scaledEnd_1))
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("f32="), int(int32((hxrt.Int32Wrap(f32.byteLength)>>uint(2))))), hxrt.StringFromLiteral(",")), fmtFloat(haxe__io__FPHelper_i32ToFloat(func() int {
		_this_10 := f32.bytes
		pos_10 := f32.byteOffset
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_10.b[pos_10]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_10.b[int(int32((hxrt.Int32Wrap(pos_10)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_10.b[int(int32((hxrt.Int32Wrap(pos_10)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_10.b[int(int32((hxrt.Int32Wrap(pos_10)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	}()))), hxrt.StringFromLiteral(",")), fmtFloat(haxe__io__FPHelper_i32ToFloat(func() int {
		_this_11 := f32sub.bytes
		pos_11 := f32sub.byteOffset
		return int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_11.b[pos_11]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_11.b[int(int32((hxrt.Int32Wrap(pos_11)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_11.b[int(int32((hxrt.Int32Wrap(pos_11)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_11.b[int(int32((hxrt.Int32Wrap(pos_11)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
	}()))))
	hxrt.Println(v_7)
	f64 := haxe__io___Float64Array__Float64Array_Impl__fromArray(hxrt.NewArray(3.5, -1.25), 0, 2)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("f64="), int(int32((hxrt.Int32Wrap(f64.byteLength)>>uint(3))))), hxrt.StringFromLiteral(",")), fmtFloat(func() float64 {
		pos_12 := f64.byteOffset
		_this_12 := f64.bytes
		low := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_12.b[pos_12]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_12.b[int(int32((hxrt.Int32Wrap(pos_12)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_12.b[int(int32((hxrt.Int32Wrap(pos_12)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_12.b[int(int32((hxrt.Int32Wrap(pos_12)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
		_this_13 := f64.bytes
		pos_13 := int(int32((hxrt.Int32Wrap(pos_12) + hxrt.Int32Wrap(4))))
		high := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_13.b[pos_13]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_13.b[int(int32((hxrt.Int32Wrap(pos_13)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_13.b[int(int32((hxrt.Int32Wrap(pos_13)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_13.b[int(int32((hxrt.Int32Wrap(pos_13)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
		return haxe__io__FPHelper_i64ToDouble(low, high)
	}())), hxrt.StringFromLiteral(",")), fmtFloat(func() float64 {
		pos_14 := int(int32((hxrt.Int32Wrap(8) + hxrt.Int32Wrap(f64.byteOffset))))
		_this_14 := f64.bytes
		low_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_14.b[pos_14]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_14.b[int(int32((hxrt.Int32Wrap(pos_14)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_14.b[int(int32((hxrt.Int32Wrap(pos_14)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_14.b[int(int32((hxrt.Int32Wrap(pos_14)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
		_this_15 := f64.bytes
		pos_15 := int(int32((hxrt.Int32Wrap(pos_14) + hxrt.Int32Wrap(4))))
		high_1 := int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_15.b[pos_15]) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_15.b[int(int32((hxrt.Int32Wrap(pos_15)+hxrt.Int32Wrap(1))))]) << uint(8))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_15.b[int(int32((hxrt.Int32Wrap(pos_15)+hxrt.Int32Wrap(2))))]) << uint(16))))))))) | hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(_this_15.b[int(int32((hxrt.Int32Wrap(pos_15)+hxrt.Int32Wrap(3))))]) << uint(24))))))))
		return haxe__io__FPHelper_i64ToDouble(low_1, high_1)
	}())))
	hxrt.Println(v_8)
	hxrt.TryCatch(func() {
		haxe__io___ArrayBufferView__ArrayBufferView_Impl__fromBytes(bytes, 9, 4)
		hxrt.Println(any(hxrt.StringFromLiteral("bounds=miss")))
	}, func(hx_caught_5 any) {
		switch hx_typed_6 := hx_caught_5.(type) {
		case *haxe__io__Error:
			err := hx_typed_6
			var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bounds="), errTag(err)))
			hxrt.Println(v_9)
		default:
			hxrt.Throw(hx_caught_5)
		}
	})
}
