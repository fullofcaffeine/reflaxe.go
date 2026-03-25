package main

func haxe__NativeStackTrace_callStack() any {
	return []any{}
}

func haxe__NativeStackTrace_exceptionStack() any {
	return []any{}
}

func haxe__NativeStackTrace_saveStack(_exception any) {
}

func haxe__NativeStackTrace_toHaxe(nativeStackTrace any, skip int) []*haxe__StackItem {
	if !func(hx_value any) bool {
		switch hx_value.(type) {
		case []*haxe___Int64_____Int64:
			return true
		case []*string:
			return true
		case []any:
			return true
		case []bool:
			return true
		case []float64:
			return true
		case []int:
			return true
		default:
			return false
		}
	}(any(nativeStackTrace)) {
		return []*haxe__StackItem{}
	}
	stack := func(hx_value_8 any) []*haxe__StackItem {
		if hx_value_8 == nil {
			var hx_zero_9 []*haxe__StackItem
			return hx_zero_9
		}
		return hx_value_8.([]*haxe__StackItem)
	}(nativeStackTrace)
	if skip <= 0 {
		return func(src []*haxe__StackItem) []*haxe__StackItem {
			out := append([]*haxe__StackItem{}, src...)
			return out
		}(stack)
	}
	out := []*haxe__StackItem{}
	_g := skip
	_g1 := len(stack)
	for _g < _g1 {
		hx_post_10 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_10
		hx_arr_11 := out
		hx_arr_11 = append(hx_arr_11, stack[index])
		out = hx_arr_11
	}
	return out
}
