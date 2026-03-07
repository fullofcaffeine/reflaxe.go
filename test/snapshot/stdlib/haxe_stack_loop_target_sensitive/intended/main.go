package main

import "snapshot/hxrt"

func main() {
	stack := haxe___CallStack__CallStack_Impl__callStack()
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("call.len="), len(stack)))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call.str="), haxe___CallStack__CallStack_Impl__toString(stack)))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("call.copy.len="), len(func(src []*haxe__StackItem) []*haxe__StackItem {
		out := append([]*haxe__StackItem{}, src...)
		return out
	}(stack))))
	hxrt.TryCatch(func() {
		hxrt.Throw(hxrt.StringFromLiteral("boom"))
	}, func(hx_caught_1 any) {
		error := hx_caught_1
		_ = error
		exceptionStack := haxe___CallStack__CallStack_Impl__exceptionStack(false)
		hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("exc.len="), len(exceptionStack)))
		hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exc.str="), haxe___CallStack__CallStack_Impl__toString(exceptionStack)))
	})
	var nativeCall any = haxe__NativeStackTrace_callStack()
	nativeHaxe := haxe__NativeStackTrace_toHaxe(nativeCall, 0)
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("native.len="), len(nativeHaxe)))
}
