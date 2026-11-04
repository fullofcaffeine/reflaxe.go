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
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
