package main

import "snapshot/hxrt"

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20leaf"))
	sink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, sink.haxe__io__Output, nil, nil)
	values := http.__hx_this.getResponseHeaderValues(hxrt.StringFromLiteral("Content-Type"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("headers="), func() *string {
		var hx_if_3 *string
		if (values != nil) && (values.Len() > 0) {
			hx_if_3 = hxrt.StdString(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(values.Get(0)))
		} else {
			hx_if_3 = hxrt.StringFromLiteral("null")
		}
		return hx_if_3
	}()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct="), sys__Http_requestUrl(hxrt.StringFromLiteral("data:text/plain,direct%20ok"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_2)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_4 := map[string]any{}
		hx_obj_4["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_4["port"] = 3128
		hx_obj_5 := map[string]any{}
		hx_obj_5["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_5["pass"] = hxrt.StringFromLiteral("tiger")
		hx_obj_4["auth"] = hx_obj_5
		return hx_obj_4
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_3)
}
