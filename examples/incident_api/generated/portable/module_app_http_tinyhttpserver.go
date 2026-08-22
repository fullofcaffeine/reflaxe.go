package main

import "examples_incident_api_portable/hxrt"

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

func New_app__http__TinyHttpServer(api *app__core__IncidentApi, host *string, port int) *app__http__TinyHttpServer {
	self := &app__http__TinyHttpServer{}
	self.__hx_this = self
	self.api = api
	self.host = host
	self.server = New_sys__net__Socket()
	self.server.__hx_this.bind(New_sys__net__Host(host), port)
	self.server.__hx_this.listen(16)
	bound := self.server.__hx_this.host()
	var hx_if_4 int
	if bound == nil {
		hx_if_4 = port
	} else {
		hx_if_4 = func(hx_obj_1 map[string]any) int {
			hx_field_2 := hx_obj_1["port"]
			if hx_field_2 == nil {
				var hx_zero_3 int
				return hx_zero_3
			}
			return hx_field_2.(int)
		}(bound)
	}
	self.port = hx_if_4
	return self
}

func (self *app__http__TinyHttpServer) serveOnce() {
	var peer *sys__net__Socket = nil
	hxrt.TryCatch(func() {
		peer = self.server.__hx_this.accept()
		request := self.readRequest(peer)
		response := self.api.handle(request)
		self.writeResponse(peer, response)
	}, func(hx_caught_5 any) {
		error := hxrt.ExceptionCaught(hx_caught_5)
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
		hx_tmp := hxrt.ExceptionCaught(hx_caught_7)
		_ = hx_tmp
	})
}

func (self *app__http__TinyHttpServer) readRequest(peer *sys__net__Socket) *app__http__HttpRequest {
	line := peer.input.__hx_this.readLine()
	first := hxrt.ArrayFromValues(func(hx_sort_src_9 []*string) []any {
		hx_sort_out_11 := make([]any, 0, len(hx_sort_src_9))
		for _, hx_sort_item_10 := range hx_sort_src_9 {
			hx_sort_out_11 = append(hx_sort_out_11, hx_sort_item_10)
		}
		return hx_sort_out_11
	}(hxrt.StringSplitStringPtr(line, hxrt.StringFromLiteral(" "))))
	var hx_if_14 *string
	if first.Len() > 0 {
		hx_if_14 = hxrt.StdString(func(hx_value_12 any) *string {
			if hx_value_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_value_12.(*string)
		}(first.Get(0)))
	} else {
		hx_if_14 = hxrt.StringFromLiteral("GET")
	}
	method := hx_if_14
	var hx_if_17 *string
	if first.Len() > 1 {
		hx_if_17 = hxrt.StdString(func(hx_value_15 any) *string {
			if hx_value_15 == nil {
				var hx_zero_16 *string
				return hx_zero_16
			}
			return hx_value_15.(*string)
		}(first.Get(1)))
	} else {
		hx_if_17 = hxrt.StringFromLiteral("/")
	}
	path := hx_if_17
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
			var hx_if_18 int
			if parsed == nil {
				hx_if_18 = 0
			} else {
				hx_if_18 = parsed.(int)
			}
			contentLength = hx_if_18
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
	}, func(hx_caught_19 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_19)
		_ = hx_tmp
	})
}

func app__http__TinyHttpServer_reason(status int) *string {
	var hx_switch_21 *string
	switch status {
	case 200:
		hx_switch_21 = hxrt.StringFromLiteral("OK")
	case 201:
		hx_switch_21 = hxrt.StringFromLiteral("Created")
	case 400:
		hx_switch_21 = hxrt.StringFromLiteral("Bad Request")
	case 404:
		hx_switch_21 = hxrt.StringFromLiteral("Not Found")
	case 500:
		hx_switch_21 = hxrt.StringFromLiteral("Internal Server Error")
	default:
		hx_switch_21 = hxrt.StringFromLiteral("OK")
	}
	return hx_switch_21
}
