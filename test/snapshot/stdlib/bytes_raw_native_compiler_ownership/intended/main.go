package main

import "snapshot/hxrt"

func main() {
	sample := hxrt.StringFromLiteral("hé")
	raw := haxe__io__Bytes_ofString(sample, haxe__io__Encoding_RawNative)
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("raw.len="), raw.length))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.hex="), raw.__hx_this.toHex()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.get="), raw.__hx_this.getString(0, raw.length, haxe__io__Encoding_RawNative)))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.base64.before="), haxe__crypto__Base64_encode(raw, true)))
	hxrt.Println(v_3)
	raw.b[0] = 65
	raw.__hx_rawValid = false
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.hex.after="), raw.__hx_this.toHex()))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.get.after="), raw.__hx_this.getString(0, raw.length, haxe__io__Encoding_RawNative)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("raw.base64.after="), haxe__crypto__Base64_encode(raw, true)))
	hxrt.Println(v_6)
	output := New_haxe__io__BytesOutput()
	output.__hx_this.writeString(sample, haxe__io__Encoding_RawNative)
	written := output.__hx_this.getBytes()
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("out.hex="), written.__hx_this.toHex()))
	hxrt.Println(v_7)
	input := New_haxe__io__BytesInput(written, nil, nil)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("in.raw="), input.__hx_this.readString(written.length, haxe__io__Encoding_RawNative)))
	hxrt.Println(v_8)
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
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
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
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
