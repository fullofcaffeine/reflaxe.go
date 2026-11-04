package main

import "snapshot/hxrt"

func main() {
	bytes := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("abc"), nil)
	bytes.b[1] = 122
	bytes.__hx_rawValid = false
	var v any = any(bytes.__hx_this.toString())
	hxrt.Println(v)
	var v_1 any = any(bytes.length)
	hxrt.Println(v_1)
	buffer := New_haxe__io__BytesBuffer()
	var encoding *haxe__io__Encoding = nil
	source := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("Hi"), encoding)
	buffer.b = hxrt.BytesBufferAdd(buffer.b, source.__hx_this.getData())
	buffer.b = hxrt.BytesBufferAddByte(buffer.b, 33)
	out := buffer.__hx_this.getBytes()
	var v_2 any = any(out.__hx_this.toString())
	hxrt.Println(v_2)
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
