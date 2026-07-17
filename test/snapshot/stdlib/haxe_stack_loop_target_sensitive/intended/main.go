package main

import "snapshot/hxrt"

func main() {
	stack := haxe___CallStack__CallStack_Impl__callStack()
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("call.len="), stack.Len()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call.str="), haxe___CallStack__CallStack_Impl__toString(stack)))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("call.copy.len="), stack.Copy().Len()))
	hxrt.Println(v_2)
	hxrt.TryCatch(func() {
		hxrt.Throw(hxrt.StringFromLiteral("boom"))
	}, func(hx_caught_1 any) {
		error := hx_caught_1
		_ = error
		exceptionStack := haxe___CallStack__CallStack_Impl__exceptionStack(false)
		var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("exc.len="), exceptionStack.Len()))
		hxrt.Println(v_3)
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exc.str="), haxe___CallStack__CallStack_Impl__toString(exceptionStack)))
		hxrt.Println(v_4)
	})
	var nativeCall any = haxe__NativeStackTrace_callStack()
	nativeHaxe := haxe__NativeStackTrace_toHaxe(nativeCall, 0)
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("native.len="), nativeHaxe.Len()))
	hxrt.Println(v_5)
}
