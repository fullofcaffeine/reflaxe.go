package main

import "snapshot/hxrt"

func main() {
	host := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	server := New_sys__net__Socket()
	server.__hx_this.bind(host, 0)
	port := func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["port"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(server.__hx_this.host())
	beforeListen := New_sys__net__Socket()
	beforeListen.__hx_this.setTimeout(0.02)
	connectedBeforeListen := true
	hxrt.TryCatch(func() {
		beforeListen.__hx_this.connect(host, port)
	}, func(hx_caught_4 any) {
		hx_tmp := hx_caught_4
		_ = hx_tmp
		connectedBeforeListen = false
	})
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("connectedBeforeListen="), hxrt.StdString(connectedBeforeListen)))
	hxrt.Println(v)
	safeClose(beforeListen)
	if connectedBeforeListen {
		safeClose(server)
		return
	}
	server.__hx_this.listen(1)
	server.__hx_this.listen(2)
	client := New_sys__net__Socket()
	var accepted *sys__net__Socket = nil
	var failed any = nil
	hxrt.TryCatch(func() {
		client.__hx_this.connect(host, port)
		accepted = server.__hx_this.accept()
		client.output.__hx_this.writeString(hxrt.StringFromLiteral("ping\n"), nil)
		client.output.__hx_this.flush()
		request := accepted.input.__hx_this.readLine()
		accepted.output.__hx_this.writeString(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("pong:"), request), hxrt.StringFromLiteral("\n")), nil)
		accepted.output.__hx_this.flush()
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("roundTrip="), client.input.__hx_this.readLine()))
		hxrt.Println(v_1)
	}, func(hx_caught_6 any) {
		error := hx_caught_6
		failed = error
	})
	safeClose(accepted)
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failed) {
		hxrt.Throw(failed)
	}
}

func safeClose(socket *sys__net__Socket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_8 any) {
		hx_tmp := hx_caught_8
		_ = hx_tmp
	})
}
