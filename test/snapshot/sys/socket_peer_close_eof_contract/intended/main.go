package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	server := New_sys__net__Socket()
	server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.__hx_this.listen(1)
	port := func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(hxrt.AnonymousGet(server.__hx_this.host(), "port"))
	serverResult := New_sys__thread__Deque()
	sys__thread__Thread_create(hxrt.NewFunction(func() {
		var peer *sys__net__Socket = nil
		result := hxrt.StringFromLiteral("closed")
		hxrt.TryCatch(func() {
			peer = server.__hx_this.accept()
			peer.output.__hx_this.writeString(hxrt.StringFromLiteral("xy"), nil)
			peer.output.__hx_this.flush()
		}, func(hx_caught_3 any) {
			hx_exception_4 := haxe__Exception_caught(hx_caught_3)
			hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
			_ = hx_exception_4
			var error any = hx_payload_5
			result = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error:"), hxrt.StdString(error))
		})
		safeClose(peer)
		serverResult.__hx_this.add(result)
	}))
	client := New_sys__net__Socket()
	client.__hx_this.connect(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), port)
	bytes := haxe__io__Bytes_alloc(8)
	count := client.input.__hx_this.readBytes(bytes, 0, bytes.length)
	peerState := func(hx_value_6 any) *string {
		if hx_value_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_value_6.(*string)
	}(serverResult.__hx_this.pop(true))
	reachedEof := false
	unexpected := hxrt.StringFromLiteral("")
	hxrt.TryCatch(func() {
		client.input.__hx_this.readByte()
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		_ = hx_exception_9
		if hx_is_nominal_0(hx_payload_10) {
			var hx_tmp *haxe__io__Eof = hx_recover_class_0(hx_payload_10)
			_ = hx_tmp
			reachedEof = true
		} else {
			var error any = hx_payload_10
			unexpected = hxrt.StdString(error)
		}
	})
	var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("partial="), bytes.__hx_this.getString(0, count, nil)), hxrt.StringFromLiteral(":")), count))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server="), peerState)))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("eof="), hxrt.StdString(reachedEof)))
	hxrt.Println(v_1)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unexpected="), unexpected)))
	safeClose(client)
	safeClose(server)
}

func safeClose(socket *sys__net__Socket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_11 any) {
		hx_exception_12 := haxe__Exception_caught(hx_caught_11)
		hx_payload_13 := haxe__Exception_catchValue(hx_caught_11)
		_ = hx_exception_12
		var hx_tmp any = hx_payload_13
		_ = hx_tmp
	})
}

func hx_recover_class_0(value any) *haxe__io__Eof {
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

func hx_recover_class_1(value any) *sys__net__Socket {
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
	case *sys__thread__NoEventLoopException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_4(value any) *haxe__ds___List__GoListIterator {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds___List__GoListIterator:
		return source
	default:
		return value.(*haxe__ds___List__GoListIterator)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__NoEventLoopException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}
