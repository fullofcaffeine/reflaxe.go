package main

import (
	"log"
	"snapshot/hxrt"
	"time"
)

func main() {
	now := time.Now()
	unix := func() *hxrt.Function[func() int] {
		hx_extern_method_1 := now.Unix
		if hx_extern_method_1 == nil {
			return nil
		}
		return hxrt.NewFunction(func() int {
			return int(hx_extern_method_1())
		})
	}()
	var v any = any((func() int {
		var hx_callable_2 *hxrt.Function[func() int] = unix
		return hxrt.FunctionCall(hx_callable_2)()
	}() > 0))
	hxrt.Println(v)
	logger := log.Default()
	originalPrefix := hxrt.StdString(logger.Prefix())
	setPrefix := func() *hxrt.Function[func(*string)] {
		hx_extern_method_3 := logger.SetPrefix
		if hx_extern_method_3 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_4 *string) {
			hx_extern_method_3(*hxrt.StdString(hx_extern_arg_4))
		})
	}()
	getPrefix := func() *hxrt.Function[func() *string] {
		hx_extern_method_5 := logger.Prefix
		if hx_extern_method_5 == nil {
			return nil
		}
		return hxrt.NewFunction(func() *string {
			return hxrt.StdString(hx_extern_method_5())
		})
	}()
	func() {
		var hx_callable_6 *hxrt.Function[func(*string)] = setPrefix
		var hx_argument_7 *string = hxrt.StringFromLiteral("hx:")
		hxrt.FunctionCall(hx_callable_6)(hx_argument_7)
	}()
	var v_1 any = any(hxrt.StringEqualStringPtr(func() *string {
		var hx_callable_8 *hxrt.Function[func() *string] = getPrefix
		return hxrt.FunctionCall(hx_callable_8)()
	}(), hxrt.StringFromLiteral("hx:")))
	hxrt.Println(v_1)
	logger.SetPrefix(*hxrt.StdString(originalPrefix))
}
