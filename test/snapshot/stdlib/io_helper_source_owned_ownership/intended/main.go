package main

import "snapshot/hxrt"

func main() {
	lines := New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("first\r\nsecond"), nil), nil, nil)
	var v any = any(lines.__hx_this.readLine())
	hxrt.Println(v)
	var v_1 any = any(lines.__hx_this.readLine())
	hxrt.Println(v_1)
	replay := New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("012345"), nil), nil, nil)
	buf := haxe__io__Bytes_alloc(4)
	replay.__hx_this.readFullBytes(buf, 0, 4)
	var v_2 any = any(buf.__hx_this.toString())
	hxrt.Println(v_2)
	copyOut := New_haxe__io__BytesOutput()
	copyOut.__hx_this.writeInput(func(hx_value *haxe__io__BytesInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Input
	}(New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("xy"), nil), nil, nil)), nil)
	var v_3 any = any(copyOut.__hx_this.getBytes().__hx_this.toString())
	hxrt.Println(v_3)
	all := New_haxe__io__BytesInput(haxe__io__Bytes_ofString(hxrt.StringFromLiteral("zz"), nil), nil, nil).__hx_this.readAll(nil)
	var v_4 any = any(all.__hx_this.toString())
	hxrt.Println(v_4)
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
