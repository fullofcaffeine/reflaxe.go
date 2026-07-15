package main

import "snapshot/hxrt"

func main() {
	hx_obj_1 := map[string]any{}
	hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
	hx_obj_1["lineNumber"] = 7
	hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
	hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
	pos := hx_obj_1
	posError := New_haxe__exceptions__PosException(hxrt.StringFromLiteral("boom"), nil, pos)
	argError := New_haxe__exceptions__ArgumentException(hxrt.StringFromLiteral("count"), nil, nil, pos)
	notImpl := New_haxe__exceptions__NotImplementedException(nil, nil, pos)
	var v any = any(hxrt.ExceptionMessage(posError))
	hxrt.Println(v)
	var v_1 any = any(argError.argument)
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.ExceptionMessage(notImpl))
	hxrt.Println(v_2)
}
