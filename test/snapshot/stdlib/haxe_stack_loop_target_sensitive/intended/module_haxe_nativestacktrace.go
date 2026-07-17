package main

import "snapshot/hxrt"

func haxe__NativeStackTrace_callStack() any {
	stack := hxrt.NewArray()
	return stack
}

func haxe__NativeStackTrace_exceptionStack() any {
	stack := hxrt.NewArray()
	return stack
}

func haxe__NativeStackTrace_saveStack(_exception any) {
}

func haxe__NativeStackTrace_toHaxe(nativeStackTrace any, skip int) *hxrt.Array {
	if !func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(any(nativeStackTrace)) {
		return hxrt.NewArray()
	}
	stack := func(hx_value_12 any) *hxrt.Array {
		if hx_value_12 == nil {
			var hx_zero_13 *hxrt.Array
			return hx_zero_13
		}
		return hx_value_12.(*hxrt.Array)
	}(nativeStackTrace)
	if skip <= 0 {
		return stack.Copy()
	}
	out := hxrt.NewArray()
	_g := skip
	_g1 := stack.Len()
	for _g < _g1 {
		hx_post_14 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_14
		out.Push(stack.Get(index))
	}
	return out
}
