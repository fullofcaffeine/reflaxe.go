package main

import "snapshot/hxrt"

func main() {
	host := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	server := New_sys__net__Socket()
	server.__hx_this.bind(host, 0)
	port := func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(hxrt.AnonymousGet(server.__hx_this.host(), "port"))
	beforeListen := New_sys__net__Socket()
	beforeListen.__hx_this.setTimeout(0.02)
	connectedBeforeListen := true
	hxrt.TryCatch(func() {
		beforeListen.__hx_this.connect(host, port)
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_4
		var hx_tmp any = hx_payload_5
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
		hx_exception_7 := haxe__Exception_caught(hx_caught_6)
		hx_payload_8 := haxe__Exception_catchValue(hx_caught_6)
		_ = hx_exception_7
		var error any = hx_payload_8
		failed = error
	})
	safeClose(accepted)
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failed) {
		hxrt.Throw(haxe__Exception_thrown(failed))
	}
}

func safeClose(socket *sys__net__Socket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_9 any) {
		hx_exception_10 := haxe__Exception_caught(hx_caught_9)
		hx_payload_11 := haxe__Exception_catchValue(hx_caught_9)
		_ = hx_exception_10
		var hx_tmp any = hx_payload_11
		_ = hx_tmp
	})
}

func hx_recover_class_0(value any) *sys__net__Socket {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Socket:
		return source
	default:
		return value.(*sys__net__Socket)
	}
}

func hx_recover_class_1(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_2(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_3(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
