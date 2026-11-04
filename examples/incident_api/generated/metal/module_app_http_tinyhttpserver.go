package main

import "examples_incident_api_metal/hxrt"

type I_app__http__TinyHttpServer interface {
	serveOnce()
	close()
	readRequest(peer *sys__net__Socket) *app__http__HttpRequest
	readBody(peer *sys__net__Socket, length int) *string
	writeResponse(peer *sys__net__Socket, response *app__http__HttpResponse)
}

type app__http__TinyHttpServer struct {
	__hx_this I_app__http__TinyHttpServer
	api       *app__core__IncidentApi
	server    *sys__net__Socket
	host      *string
	port      int
}

func (self *app__http__TinyHttpServer) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__http__TinyHttpServer(api *app__core__IncidentApi, host *string, port int) *app__http__TinyHttpServer {
	return hx_init_app__http__TinyHttpServer(hx_alloc_app__http__TinyHttpServer(nil), api, host, port)
}

func hx_alloc_app__http__TinyHttpServer(owner any) *app__http__TinyHttpServer {
	self := &app__http__TinyHttpServer{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__http__TinyHttpServer)
	}
	return self
}

func hx_init_app__http__TinyHttpServer(self *app__http__TinyHttpServer, api *app__core__IncidentApi, host *string, port int) *app__http__TinyHttpServer {
	self.api = api
	self.host = host
	self.server = New_sys__net__Socket()
	self.server.__hx_this.bind(New_sys__net__Host(host), port)
	self.server.__hx_this.listen(16)
	bound := self.server.__hx_this.host()
	var hx_if_3 int
	if bound == nil {
		hx_if_3 = port
	} else {
		hx_if_3 = func(hx_value_1 any) int {
			if hx_value_1 == nil {
				var hx_zero_2 int
				return hx_zero_2
			}
			return hx_value_1.(int)
		}(hxrt.AnonymousGet(bound, "port"))
	}
	self.port = hx_if_3
	return self
}

func (self *app__http__TinyHttpServer) serveOnce() {
	var peer *sys__net__Socket = nil
	hxrt.TryCatch(func() {
		peer = self.server.__hx_this.accept()
		request := self.readRequest(peer)
		response := self.api.handle(request)
		self.writeResponse(peer, response)
	}, func(hx_caught_4 any) {
		hx_exception_5 := haxe__Exception_caught(hx_caught_4)
		hx_payload_6 := haxe__Exception_catchValue(hx_caught_4)
		_ = hx_payload_6
		var error *haxe__Exception = hx_recover_class_0(hx_exception_5)
		_ = error
		if peer != nil {
			self.writeResponse(peer, app__http__HttpResponse_json(500, hxrt.StringFromLiteral("{\"error\":\"server_error\"}")))
		}
	})
	app__http__TinyHttpServer_closePeer(peer)
}

func (self *app__http__TinyHttpServer) close() {
	hxrt.TryCatch(func() {
		self.server.__hx_this.close()
	}, func(hx_caught_7 any) {
		hx_exception_8 := haxe__Exception_caught(hx_caught_7)
		hx_payload_9 := haxe__Exception_catchValue(hx_caught_7)
		_ = hx_payload_9
		var hx_tmp *haxe__Exception = hx_recover_class_0(hx_exception_8)
		_ = hx_tmp
	})
}

func (self *app__http__TinyHttpServer) readRequest(peer *sys__net__Socket) *app__http__HttpRequest {
	line := peer.input.__hx_this.readLine()
	first := hxrt.ArrayFromValues(func(hx_sort_src_10 []*string) []any {
		hx_sort_out_12 := make([]any, 0, len(hx_sort_src_10))
		for _, hx_sort_item_11 := range hx_sort_src_10 {
			hx_sort_out_12 = append(hx_sort_out_12, hx_sort_item_11)
		}
		return hx_sort_out_12
	}(hxrt.StringSplitStringPtr(line, hxrt.StringFromLiteral(" "))))
	var hx_if_15 *string
	if first.Len() > 0 {
		hx_if_15 = hxrt.StdString(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(first.Get(0)))
	} else {
		hx_if_15 = hxrt.StringFromLiteral("GET")
	}
	method := hx_if_15
	var hx_if_18 *string
	if first.Len() > 1 {
		hx_if_18 = hxrt.StdString(func(hx_value_16 any) *string {
			if hx_value_16 == nil {
				var hx_zero_17 *string
				return hx_zero_17
			}
			return hx_value_16.(*string)
		}(first.Get(1)))
	} else {
		hx_if_18 = hxrt.StringFromLiteral("/")
	}
	path := hx_if_18
	contentLength := 0
	for true {
		header := peer.input.__hx_this.readLine()
		if hxrt.StringEqualStringPtr(header, hxrt.StringFromLiteral("")) {
			break
		}
		lower := hxrt.StringToLowerCaseStringPtr(header)
		if StringTools_startsWith(lower, hxrt.StringFromLiteral("content-length:")) {
			rawLength := StringTools_trim(hxrt.StringSubstrStringPtr(header, hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("content-length:")), 0, false))
			var parsed any = Std_parseInt(rawLength)
			var hx_if_19 int
			if parsed == nil {
				hx_if_19 = 0
			} else {
				hx_if_19 = parsed.(int)
			}
			contentLength = hx_if_19
		}
	}
	return New_app__http__HttpRequest(method, path, self.readBody(peer, contentLength))
}

func (self *app__http__TinyHttpServer) readBody(peer *sys__net__Socket, length int) *string {
	if length <= 0 {
		return hxrt.StringFromLiteral("")
	}
	out := New_haxe__io__BytesBuffer()
	i := 0
	for i < length {
		value := peer.input.__hx_this.readByte()
		out.b = hxrt.BytesBufferAddByte(out.b, value)
		i = int(int32((i + 1)))
	}
	return out.__hx_this.getBytes().__hx_this.toString()
}

func (self *app__http__TinyHttpServer) writeResponse(peer *sys__net__Socket, response *app__http__HttpResponse) {
	bodyBytes := haxe__io__Bytes_ofString(response.body, nil)
	head := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("HTTP/1.1 "), response.status), hxrt.StringFromLiteral(" ")), app__http__TinyHttpServer_reason(response.status)), hxrt.StringFromLiteral("\r\n")), hxrt.StringFromLiteral("Content-Type: application/json\r\n")), hxrt.StringFromLiteral("Content-Length: ")), bodyBytes.length), hxrt.StringFromLiteral("\r\n")), hxrt.StringFromLiteral("Connection: close\r\n")), hxrt.StringFromLiteral("\r\n"))
	peer.output.__hx_this.writeString(hxrt.StringConcatStringPtr(head, response.body), nil)
	peer.output.__hx_this.flush()
}

func app__http__TinyHttpServer_closePeer(peer *sys__net__Socket) {
	if peer == nil {
		return
	}
	hxrt.TryCatch(func() {
		peer.__hx_this.close()
	}, func(hx_caught_20 any) {
		hx_exception_21 := haxe__Exception_caught(hx_caught_20)
		hx_payload_22 := haxe__Exception_catchValue(hx_caught_20)
		_ = hx_payload_22
		var hx_tmp *haxe__Exception = hx_recover_class_0(hx_exception_21)
		_ = hx_tmp
	})
}

func app__http__TinyHttpServer_reason(status int) *string {
	var hx_switch_23 *string
	switch status {
	case 200:
		hx_switch_23 = hxrt.StringFromLiteral("OK")
	case 201:
		hx_switch_23 = hxrt.StringFromLiteral("Created")
	case 400:
		hx_switch_23 = hxrt.StringFromLiteral("Bad Request")
	case 404:
		hx_switch_23 = hxrt.StringFromLiteral("Not Found")
	case 500:
		hx_switch_23 = hxrt.StringFromLiteral("Internal Server Error")
	default:
		hx_switch_23 = hxrt.StringFromLiteral("OK")
	}
	return hx_switch_23
}
