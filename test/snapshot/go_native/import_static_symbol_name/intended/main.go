package main

import (
	"snapshot/hxrt"
	"strings"
)

var Main_hxrt *string = hxrt.StringFromLiteral("runtime")

func main() {
	convert := hxrt.NamedFunction(Main_strings, "Main.strings")
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(Main_strings(hxrt.StringFromLiteral("x")), func() *string {
		var hx_callable_1 *hxrt.Function[func(*string) *string] = convert
		var hx_argument_2 *string = hxrt.StringFromLiteral("y")
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()), plain(Main_hxrt)))
	hxrt.Println(v)
}

func plain(value *string) *string {
	return hxrt.StringConcatStringPtr(value, hxrt.StringFromLiteral("!"))
}

func Main_strings(value *string) *string {
	return hxrt.StdString(strings.ToUpper(*hxrt.StdString(value)))
}
