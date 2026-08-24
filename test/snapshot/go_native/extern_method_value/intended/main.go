package main

import (
	"log"
	"snapshot/hxrt"
	"time"
)

func main() {
	now := time.Now()
	unix := now.Unix
	var v any = any((unix() > 0))
	hxrt.Println(v)
	logger := log.Default()
	originalPrefix := hxrt.StdString(logger.Prefix())
	setPrefix := func() func(*string) {
		hx_extern_method_1 := logger.SetPrefix
		return func(hx_extern_arg_2 *string) {
			hx_extern_method_1(*hxrt.StdString(hx_extern_arg_2))
		}
	}()
	getPrefix := func() func() *string {
		hx_extern_method_3 := logger.Prefix
		return func() *string {
			return hxrt.StdString(hx_extern_method_3())
		}
	}()
	setPrefix(hxrt.StringFromLiteral("hx:"))
	var v_1 any = any(hxrt.StringEqualStringPtr(getPrefix(), hxrt.StringFromLiteral("hx:")))
	hxrt.Println(v_1)
	logger.SetPrefix(*hxrt.StdString(originalPrefix))
}
