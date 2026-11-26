package main

import "snapshot/hxrt"

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("args="), hxrt.StdString((Sys_args() != nil))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("direct="), Sys_command(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("exit 3")))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("shell="), Sys_command(hxrt.StringFromLiteral("exit 4"), nil)))
	hxrt.Println(v_2)
}
