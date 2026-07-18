package main

import "snapshot/hxrt"

func main() {
	server := New_sys__net__UdpSocket()
	client := New_sys__net__UdpSocket()
	var failure any = nil
	hxrt.TryCatch(func() {
		server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
		bound := server.__hx_this.host()
		if (bound == nil) || (func(hx_obj_3 map[string]any) int {
			hx_field_4 := hx_obj_3["port"]
			if hx_field_4 == nil {
				var hx_zero_5 int
				return hx_zero_5
			}
			return hx_field_4.(int)
		}(bound) <= 0) {
			hxrt.Throw(hxrt.StringFromLiteral("missing bound udp port"))
		}
		server.__hx_this.setBlocking(true)
		client.__hx_this.setBroadcast(true)
		sent := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("udp-ping"), nil)
		target := New_sys__net__Address()
		target.host = New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")).ip
		target.port = func(hx_obj_6 map[string]any) int {
			hx_field_7 := hx_obj_6["port"]
			if hx_field_7 == nil {
				var hx_zero_8 int
				return hx_zero_8
			}
			return hx_field_7.(int)
		}(bound)
		wrote := client.__hx_this.sendTo(sent, 0, sent.length, target)
		recv := haxe__io__Bytes_alloc(32)
		remote := New_sys__net__Address()
		read := server.__hx_this.readFrom(recv, 0, recv.length, remote)
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bound.host="), func(hx_obj_9 map[string]any) *sys__net__Host {
			hx_field_10 := hx_obj_9["host"]
			if hx_field_10 == nil {
				var hx_zero_11 *sys__net__Host
				return hx_zero_11
			}
			return hx_field_10.(*sys__net__Host)
		}(bound).__hx_this.toString()))
		hxrt.Println(v)
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bound.port.positive="), hxrt.StdString((func(hx_obj_12 map[string]any) int {
			hx_field_13 := hx_obj_12["port"]
			if hx_field_13 == nil {
				var hx_zero_14 int
				return hx_zero_14
			}
			return hx_field_13.(int)
		}(bound) > 0))))
		hxrt.Println(v_1)
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("wrote="), wrote)))
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("read="), read)))
		var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("payload="), recv.__hx_this.sub(0, read).__hx_this.toString()))
		hxrt.Println(v_2)
		var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remote.port.positive="), hxrt.StdString((remote.port > 0))))
		hxrt.Println(v_3)
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remote.host="), remote.__hx_this.getHost().__hx_this.toString()))
		hxrt.Println(v_4)
	}, func(hx_caught_1 any) {
		error := hx_caught_1
		failure = error
	})
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failure) {
		hxrt.Throw(failure)
	}
}

func safeClose(socket *sys__net__UdpSocket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_15 any) {
		hx_tmp := hx_caught_15
		_ = hx_tmp
	})
}
