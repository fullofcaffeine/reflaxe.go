package main

import "snapshot/hxrt"

func main() {
	server := New_sys__net__UdpSocket()
	client := New_sys__net__UdpSocket()
	var failure any = nil
	hxrt.TryCatch(func() {
		server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
		bound := server.__hx_this.host()
		if (bound == nil) || (func(hx_value_4 any) int {
			if hx_value_4 == nil {
				var hx_zero_5 int
				return hx_zero_5
			}
			return hx_value_4.(int)
		}(hxrt.AnonymousGet(bound, "port")) <= 0) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("missing bound udp port")))
		}
		server.__hx_this.setBlocking(true)
		client.__hx_this.setBroadcast(true)
		sent := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("udp-ping"), nil)
		target := New_sys__net__Address()
		target.host = New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")).ip
		target.port = func(hx_value_6 any) int {
			if hx_value_6 == nil {
				var hx_zero_7 int
				return hx_zero_7
			}
			return hx_value_6.(int)
		}(hxrt.AnonymousGet(bound, "port"))
		wrote := client.__hx_this.sendTo(sent, 0, sent.length, target)
		recv := haxe__io__Bytes_alloc(32)
		remote := New_sys__net__Address()
		read := server.__hx_this.readFrom(recv, 0, recv.length, remote)
		var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bound.host="), hx_recover_class_0(hxrt.AnonymousGet(bound, "host")).__hx_this.toString()))
		hxrt.Println(v)
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bound.port.positive="), hxrt.StdString((func(hx_value_8 any) int {
			if hx_value_8 == nil {
				var hx_zero_9 int
				return hx_zero_9
			}
			return hx_value_8.(int)
		}(hxrt.AnonymousGet(bound, "port")) > 0))))
		hxrt.Println(v_1)
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("wrote="), wrote)))
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("read="), read)))
		var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("payload="), recv.__hx_this.sub(0, read).__hx_this.toString()))
		hxrt.Println(v_2)
		var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remote.port.positive="), hxrt.StdString((remote.port > 0))))
		hxrt.Println(v_3)
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("remote.host="), remote.__hx_this.getHost().__hx_this.toString()))
		hxrt.Println(v_4)
		empty := haxe__io__Bytes_alloc(0)
		emptyWrote := client.__hx_this.sendTo(empty, 0, empty.length, target)
		emptyRecv := haxe__io__Bytes_alloc(1)
		emptyRemote := New_sys__net__Address()
		emptyRead := server.__hx_this.readFrom(emptyRecv, 0, emptyRecv.length, emptyRemote)
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("empty.wrote="), emptyWrote)))
		hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("empty.read="), emptyRead)))
		var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("empty.remote.port.positive="), hxrt.StdString((emptyRemote.port > 0))))
		hxrt.Println(v_5)
		var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("empty.remote.host="), emptyRemote.__hx_this.getHost().__hx_this.toString()))
		hxrt.Println(v_6)
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var error any = hx_payload_3
		failure = error
	})
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failure) {
		hxrt.Throw(haxe__Exception_thrown(failure))
	}
}

func safeClose(socket *sys__net__UdpSocket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_10 any) {
		hx_exception_11 := haxe__Exception_caught(hx_caught_10)
		hx_payload_12 := haxe__Exception_catchValue(hx_caught_10)
		_ = hx_exception_11
		var hx_tmp any = hx_payload_12
		_ = hx_tmp
	})
}

func hx_recover_class_0(value any) *sys__net__Host {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Host:
		return source
	default:
		return value.(*sys__net__Host)
	}
}

func hx_recover_class_1(value any) *sys__net__Socket {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__UdpSocket:
		return source.sys__net__Socket
	case *sys__net__Socket:
		return source
	default:
		return value.(*sys__net__Socket)
	}
}

func hx_recover_class_2(value any) *haxe__io__Eof {
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

func hx_recover_class_3(value any) *haxe__ValueException {
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

func hx_recover_class_4(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
