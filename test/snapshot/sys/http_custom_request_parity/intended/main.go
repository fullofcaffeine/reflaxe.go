package main

import "snapshot/hxrt"

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20from%20haxe.go"))
	sink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, sink.haxe__io__Output, nil, nil)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("custom="), sink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v)
	values := http.__hx_this.getResponseHeaderValues(hxrt.StringFromLiteral("Content-Type"))
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("headers="), func() int {
		var hx_if_1 int
		if values == nil {
			hx_if_1 = -1
		} else {
			hx_if_1 = values.Len()
		}
		return hx_if_1
	}()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("header0="), func() *string {
		var hx_if_4 *string
		if (values != nil) && (values.Len() > 0) {
			hx_if_4 = hxrt.StdString(func(hx_value_2 any) *string {
				if hx_value_2 == nil {
					var hx_zero_3 *string
					return hx_zero_3
				}
				return hx_value_2.(*string)
			}(values.Get(0)))
		} else {
			hx_if_4 = hxrt.StringFromLiteral("none")
		}
		return hx_if_4
	}()))
	hxrt.Println(v_2)
	putSink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, putSink.haxe__io__Output, nil, hxrt.StringFromLiteral("PUT"))
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("method="), putSink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_3)
	upload := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,ignored"))
	upload.__hx_this.setParameter(hxrt.StringFromLiteral("token"), hxrt.StringFromLiteral("42"))
	upload.__hx_this.fileTransfer(hxrt.StringFromLiteral("asset"), hxrt.StringFromLiteral("demo.txt"), func(hx_value_5 any) *haxe__io__Input {
		if hx_value_5 == nil {
			var hx_zero_6 *haxe__io__Input
			return hx_zero_6
		}
		return hx_value_5.(*haxe__io__Input)
	}(nil), 4, hxrt.StringFromLiteral("text/plain"))
	uploadSink := New_haxe__io__BytesOutput()
	upload.__hx_this.customRequest(true, uploadSink.haxe__io__Output, nil, nil)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("upload="), uploadSink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_4)
}
