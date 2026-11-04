package main

import "snapshot/hxrt"

func main() {
	error := New_haxe__ValueException(hxrt.StringFromLiteral("boom"), nil, nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("msg="), error.__hx_this.get_message()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("value="), func() *string {
		var hx_if_1 *string
		if hxrt.AnyEqualsNull(error.value) {
			hx_if_1 = hxrt.StringFromLiteral("null")
		} else {
			var this1 any = error.value
			hx_if_1 = hxrt.StdString(this1)
		}
		return hx_if_1
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("str="), error.__hx_this.toString()))
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
	case *haxe__ValueException:
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
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
