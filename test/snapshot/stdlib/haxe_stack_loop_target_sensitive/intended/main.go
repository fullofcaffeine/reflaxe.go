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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var error any = hx_payload_3
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
