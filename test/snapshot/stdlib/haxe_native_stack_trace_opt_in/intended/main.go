package main

import "snapshot/hxrt"

func main() {
	nested()
}

func nested() {
	stack := haxe___CallStack__CallStack_Impl__callStack()
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call.nonEmpty="), hxrt.StdString((len(stack) > 0))))
	rendered := haxe___CallStack__CallStack_Impl__toString(stack)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call.renderedNonEmpty="), hxrt.StdString(!hxrt.StringEqualStringPtr(rendered, hxrt.StringFromLiteral("")))))
	var native any = haxe__NativeStackTrace_callStack()
	nativeHaxe := haxe__NativeStackTrace_toHaxe(native, 0)
	skipped := haxe__NativeStackTrace_toHaxe(native, 1)
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native.nonEmpty="), hxrt.StdString((len(nativeHaxe) > 0))))
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("native.skipOk="), hxrt.StdString((len(skipped) <= len(nativeHaxe)))))
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("native.bogus.len="), len(haxe__NativeStackTrace_toHaxe(hxrt.StringFromLiteral("not a stack"), 0))))
}
