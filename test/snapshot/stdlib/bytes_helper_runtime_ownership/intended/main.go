package main

import "snapshot/hxrt"

func main() {
	parsed := haxe__io__Bytes_ofHex(hxrt.StringFromLiteral("0fDA"))
	var v any = any(parsed.__hx_this.toHex())
	hxrt.Println(v)
	buffer := New_haxe__io__BytesBuffer()
	buffer.b = hxrt.BytesBufferAddByte(buffer.b, 260)
	if 2 > parsed.length {
		hxrt.Throw(haxe__Exception_thrown(haxe__io__Error_OutsideBounds))
	}
	buffer.b = hxrt.BytesBufferAddSlice(buffer.b, parsed.__hx_this.getData(), 1, 1)
	var encoding *haxe__io__Encoding = nil
	source := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("Z"), encoding)
	buffer.b = hxrt.BytesBufferAdd(buffer.b, source.__hx_this.getData())
	out := buffer.__hx_this.getBytes()
	var v_1 any = any(out.__hx_this.toHex())
	hxrt.Println(v_1)
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
