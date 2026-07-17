package main

import "snapshot/hxrt"

func main() {
	nested()
}

func nested() {
	stack := haxe___CallStack__CallStack_Impl__callStack()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call.nonEmpty="), hxrt.StdString((stack.Len() > 0))))
	hxrt.Println(v)
	rendered := haxe___CallStack__CallStack_Impl__toString(stack)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call.renderedNonEmpty="), hxrt.StdString(!hxrt.StringEqualStringPtr(rendered, hxrt.StringFromLiteral("")))))
	hxrt.Println(v_1)
	var native any = haxe__NativeStackTrace_callStack()
	nativeHaxe := haxe__NativeStackTrace_toHaxe(native, 0)
	skipped := haxe__NativeStackTrace_toHaxe(native, 1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native.nonEmpty="), hxrt.StdString((nativeHaxe.Len() > 0))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native.skipOk="), hxrt.StdString((skipped.Len() <= nativeHaxe.Len()))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("native.bogus.len="), haxe__NativeStackTrace_toHaxe(hxrt.StringFromLiteral("not a stack"), 0).Len()))
	hxrt.Println(v_4)
}
