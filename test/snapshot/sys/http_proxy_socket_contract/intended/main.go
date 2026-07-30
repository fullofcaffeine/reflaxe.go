package main

import "snapshot/hxrt"

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_1 := map[string]any{}
		hx_obj_1["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_1["port"] = 3128
		hx_obj_2 := map[string]any{}
		hx_obj_2["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_2["pass"] = hxrt.StringFromLiteral("tiger")
		hx_obj_1["auth"] = hx_obj_2
		return hx_obj_1
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_1)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_3 := map[string]any{}
		hx_obj_3["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_3["port"] = 80
		hx_obj_4 := map[string]any{}
		hx_obj_4["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_4["pass"] = nil
		hx_obj_3["auth"] = hx_obj_4
		return hx_obj_3
	}()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy2="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_2)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_5 := map[string]any{}
		hx_obj_5["host"] = hxrt.StringFromLiteral("proxy.local:9000")
		hx_obj_5["port"] = 3128
		hx_obj_5["auth"] = nil
		return hx_obj_5
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy3="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_3)
	sys__Http_PROXY = func() map[string]any {
		hx_obj_6 := map[string]any{}
		hx_obj_6["host"] = nil
		hx_obj_6["port"] = 3128
		hx_obj_6["auth"] = nil
		return hx_obj_6
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy4="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_4)
	sys__Http_PROXY = nil
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,body"))
	sink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, sink.haxe__io__Output, New_sys__net__Socket(), hxrt.StringFromLiteral("PATCH"))
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("methodSock="), sink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_5)
}
