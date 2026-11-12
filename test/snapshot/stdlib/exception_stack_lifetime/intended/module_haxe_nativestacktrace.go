package main

import "snapshot/hxrt"

func haxe__NativeStackTrace_callStack() any {
	return hxrt.NativeStackCapture(1)
}

func haxe__NativeStackTrace_exceptionStack() any {
	return hxrt.NativeStackCapture(1)
}

func haxe__NativeStackTrace_saveStack(_exception any) {
}

func haxe__NativeStackTrace_toHaxe(nativeStackTrace any, skip int) *hxrt.Array {
	if hxrt.AnyEqualsNull(nativeStackTrace) || !hxrt.NativeStackIsFrameSlice(nativeStackTrace) {
		return hxrt.NewArray()
	}
	frames := func(hx_value_3 any) []*hxrt.StackFrame {
		if hx_value_3 == nil {
			var hx_zero_4 []*hxrt.StackFrame
			return hx_zero_4
		}
		return hx_value_3.([]*hxrt.StackFrame)
	}(func(hx_value_1 any) []*hxrt.StackFrame {
		if hx_value_1 == nil {
			var hx_zero_2 []*hxrt.StackFrame
			return hx_zero_2
		}
		return hx_value_1.([]*hxrt.StackFrame)
	}(nativeStackTrace))
	out := hxrt.NewArray()
	var hx_if_5 int
	if skip < 0 {
		hx_if_5 = 0
	} else {
		hx_if_5 = skip
	}
	start := hx_if_5
	_g := start
	_g1 := len(frames)
	for _g < _g1 {
		hx_post_6 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_6
		frame := frames[index]
		out.Push(haxe__StackItem_FilePos(haxe__StackItem_Method(nil, frame.Function), frame.File, frame.Line, 0))
	}
	return out
}
