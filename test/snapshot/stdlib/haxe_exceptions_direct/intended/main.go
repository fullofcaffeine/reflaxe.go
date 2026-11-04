package main

import "snapshot/hxrt"

func main() {
	hx_obj_1 := map[string]any{}
	hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
	hx_obj_1["lineNumber"] = 7
	hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
	hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
	pos := hxrt.AnonymousFromMap(hx_obj_1)
	posError := New_haxe__exceptions__PosException(hxrt.StringFromLiteral("boom"), nil, pos)
	argError := New_haxe__exceptions__ArgumentException(hxrt.StringFromLiteral("count"), nil, nil, pos)
	notImpl := New_haxe__exceptions__NotImplementedException(nil, nil, pos)
	var v any = any(posError.__hx_this.get_message())
	hxrt.Println(v)
	var v_1 any = any(argError.argument)
	hxrt.Println(v_1)
	var v_2 any = any(notImpl.__hx_this.get_message())
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
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__ArgumentException:
		return source.haxe__exceptions__PosException.haxe__Exception
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
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__ArgumentException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
