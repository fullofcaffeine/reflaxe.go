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
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("stringInput="), stringInput.__hx_this.readByte()), hxrt.StringFromLiteral(",")), readTwo(func(hx_value *haxe__io__StringInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__BytesInput.haxe__io__Input
	}(stringInput))))
	hxrt.Println(v)
	buffered := New_haxe__io__BufferInput(func(hx_value *haxe__io__StringInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__BytesInput.haxe__io__Input
	}(New_haxe__io__StringInput(hxrt.StringFromLiteral("wxyz"))), haxe__io__Bytes_alloc(2), 0, 0)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("bufferInput="), buffered.__hx_this.readByte()), hxrt.StringFromLiteral(",")), readTwo(func(hx_value *haxe__io__BufferInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Input
	}(buffered))))
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
		hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		if hx_is_nominal_0(hx_payload_5) {
			var e *haxe__io__Eof = hx_recover_class_0(hx_payload_5)
			var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("eof="), hxrt.StdString(e)))
			hxrt.Println(v_3)
		} else {
			hxrt.Throw(hx_caught_3)
		}
	})
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(haxe__io__Error_Custom(hxrt.StringFromLiteral("boom"))))
	}, func(hx_caught_6 any) {
		hx_exception_7 := haxe__Exception_caught(hx_caught_6)
		hx_payload_8 := haxe__Exception_catchValue(hx_caught_6)
		_ = hx_exception_7
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *haxe__io__Error:
				return true
			default:
				return false
			}
		}(hx_payload_8) {
			var e_1 *haxe__io__Error = func(hx_value_9 any) *haxe__io__Error {
				if hx_value_9 == nil {
					var hx_zero_10 *haxe__io__Error
					return hx_zero_10
				}
				return hx_value_9.(*haxe__io__Error)
			}(hx_payload_8)
			var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error="), errorTag(e_1)))
			hxrt.Println(v_4)
		} else {
			hxrt.Throw(hx_caught_6)
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

func hx_recover_class_0(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
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

func hx_recover_class_2(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
