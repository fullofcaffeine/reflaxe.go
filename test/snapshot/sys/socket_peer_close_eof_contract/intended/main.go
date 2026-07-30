package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	server := New_sys__net__Socket()
	server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.__hx_this.listen(1)
	port := func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["port"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(server.__hx_this.host())
	serverResult := New_sys__thread__Deque()
	sys__thread__Thread_create(func() {
		var peer *sys__net__Socket = nil
		result := hxrt.StringFromLiteral("closed")
		hxrt.TryCatch(func() {
			peer = server.__hx_this.accept()
			peer.output.__hx_this.writeString(hxrt.StringFromLiteral("xy"), nil)
			peer.output.__hx_this.flush()
		}, func(hx_caught_4 any) {
			error := hx_caught_4
			result = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error:"), hxrt.StdString(error))
		})
		safeClose(peer)
		serverResult.__hx_this.add(result)
	})
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
		switch hx_typed_9 := hx_caught_8.(type) {
		case *haxe__io__Eof:
			hx_tmp := hx_typed_9
			_ = hx_tmp
			reachedEof = true
		default:
			error := hx_caught_8
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
	}, func(hx_caught_10 any) {
		hx_tmp := hx_caught_10
		_ = hx_tmp
	})
}
