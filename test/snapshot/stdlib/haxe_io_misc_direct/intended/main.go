package main

import "snapshot/hxrt"

func errorTag(err *haxe__io__Error) *string {
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

func main() {
	stringInput := New_haxe__io__StringInput(hxrt.StringFromLiteral("abc"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("stringInput="), stringInput.__hx_this.readByte()), hxrt.StringFromLiteral(",")), readTwo(stringInput.haxe__io__BytesInput.haxe__io__Input)))
	hxrt.Println(v)
	buffered := New_haxe__io__BufferInput(New_haxe__io__StringInput(hxrt.StringFromLiteral("wxyz")).haxe__io__BytesInput.haxe__io__Input, haxe__io__Bytes_alloc(2), 0, 0)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("bufferInput="), buffered.__hx_this.readByte()), hxrt.StringFromLiteral(",")), readTwo(buffered.haxe__io__Input)))
	hxrt.Println(v_1)
	source := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("ab"), nil)
	data := source.__hx_this.getData()
	data[1] = 90
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bytesData="), haxe__io__Bytes_ofData(data).__hx_this.toString()))
	hxrt.Println(v_2)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("encoding="), func() *string {
		_g := haxe__io__Encoding_RawNative
		var hx_switch_2 *string
		switch _g.tag {
		case 0:
			hx_switch_2 = hxrt.StringFromLiteral("utf8")
		case 1:
			hx_switch_2 = hxrt.StringFromLiteral("raw")
		}
		return hx_switch_2
	}())))
	hxrt.TryCatch(func() {
		hxrt.Throw(New_haxe__io__Eof())
	}, func(hx_caught_3 any) {
		switch hx_typed_4 := hx_caught_3.(type) {
		case *haxe__io__Eof:
			e := hx_typed_4
			var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("eof="), hxrt.StdString(e)))
			hxrt.Println(v_3)
		default:
			hxrt.Throw(hx_caught_3)
		}
	})
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__io__Error_Custom(hxrt.StringFromLiteral("boom")))
	}, func(hx_caught_5 any) {
		switch hx_typed_6 := hx_caught_5.(type) {
		case *haxe__io__Error:
			e_1 := hx_typed_6
			var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error="), errorTag(e_1)))
			hxrt.Println(v_4)
		default:
			hxrt.Throw(hx_caught_5)
		}
	})
	var mime any = any(hxrt.StringFromLiteral("application/json"))
	var scheme any = any(hxrt.StringFromLiteral("https"))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("mime="), mime)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("scheme="), scheme)))
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("fp.i32ToFloat="), haxe__io__FPHelper_i32ToFloat(1065353216)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("fp.floatToI32="), haxe__io__FPHelper_floatToI32(1.5)))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("fp.i64ToDouble="), haxe__io__FPHelper_i64ToDouble(0, 1072693248)))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("fp.doubleToI64="), func() *string {
		x := haxe__io__FPHelper_doubleToI64(1.0)
		return haxe___Int64__Int64_Impl__toString(x)
	}()))
	hxrt.Println(v_8)
}

func readTwo(input *haxe__io__Input) *string {
	out := haxe__io__Bytes_alloc(2)
	input.__hx_this.readBytes(out, 0, 2)
	return out.__hx_this.toString()
}
