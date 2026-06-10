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

func haxe__NativeStackTrace_toHaxe(nativeStackTrace any, skip int) []*haxe__StackItem {
	if hxrt.AnyEqualsNull(nativeStackTrace) || !hxrt.NativeStackIsFrameSlice(nativeStackTrace) {
		return []*haxe__StackItem{}
	}
	frames := func(hx_value_7 any) []*hxrt.StackFrame {
		if hx_value_7 == nil {
			var hx_zero_8 []*hxrt.StackFrame
			return hx_zero_8
		}
		return hx_value_7.([]*hxrt.StackFrame)
	}(nativeStackTrace)
	out := []*haxe__StackItem{}
	var hx_if_9 int
	if skip < 0 {
		hx_if_9 = 0
	} else {
		hx_if_9 = skip
	}
	start := hx_if_9
	_g := start
	_g1 := len(frames)
	for _g < _g1 {
		hx_post_10 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_10
		frame := frames[index]
		out = append(out, haxe__StackItem_FilePos(haxe__StackItem_Method(nil, frame.Function), frame.File, frame.Line, 0))
	}
	return out
}
