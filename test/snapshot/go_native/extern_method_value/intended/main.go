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
	var v any = any((func(hx_fn *hxrt.Function[func() int]) int {
		return hxrt.FunctionCall(hx_fn)()
	}(unix) > 0))
	hxrt.Println(v)
	logger := log.Default()
	originalPrefix := hxrt.StdString(logger.Prefix())
	setPrefix := func() *hxrt.Function[func(*string)] {
		hx_extern_method_2 := logger.SetPrefix
		if hx_extern_method_2 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_3 *string) {
			hx_extern_method_2(*hxrt.StdString(hx_extern_arg_3))
		})
	}()
	getPrefix := func() *hxrt.Function[func() *string] {
		hx_extern_method_4 := logger.Prefix
		if hx_extern_method_4 == nil {
			return nil
		}
		return hxrt.NewFunction(func() *string {
			return hxrt.StdString(hx_extern_method_4())
		})
	}()
	func(hx_fn *hxrt.Function[func(*string)], hx_arg_0 *string) {
		hxrt.FunctionCall(hx_fn)(hx_arg_0)
	}(setPrefix, hxrt.StringFromLiteral("hx:"))
	var v_1 any = any(hxrt.StringEqualStringPtr(func(hx_fn *hxrt.Function[func() *string]) *string {
		return hxrt.FunctionCall(hx_fn)()
	}(getPrefix), hxrt.StringFromLiteral("hx:")))
	hxrt.Println(v_1)
	logger.SetPrefix(*hxrt.StdString(originalPrefix))
}
