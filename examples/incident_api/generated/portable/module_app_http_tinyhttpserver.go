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
	var hx_if_75 int
	if bound == nil {
		hx_if_75 = port
	} else {
		hx_if_75 = func(hx_obj_72 map[string]any) int {
			hx_field_73 := hx_obj_72["port"]
			if hx_field_73 == nil {
				var hx_zero_74 int
				return hx_zero_74
			}
			return hx_field_73.(int)
		}(bound)
	}
	self.port = hx_if_75
	return self
}

func (self *app__http__TinyHttpServer) serveOnce() {
	var peer *sys__net__Socket = nil
	hxrt.TryCatch(func() {
		peer = self.server.__hx_this.accept()
		request := self.readRequest(peer)
		response := self.api.handle(request)
		self.writeResponse(peer, response)
	}, func(hx_caught_76 any) {
		error := hxrt.ExceptionCaught(hx_caught_76)
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
	}, func(hx_caught_78 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_78)
		_ = hx_tmp
	})
}

func (self *app__http__TinyHttpServer) readRequest(peer *sys__net__Socket) *app__http__HttpRequest {
	line := peer.input.__hx_this.readLine()
	first := hxrt.ArrayFromValues(func(hx_sort_src_80 []*string) []any {
		hx_sort_out_82 := make([]any, 0, len(hx_sort_src_80))
		for _, hx_sort_item_81 := range hx_sort_src_80 {
			hx_sort_out_82 = append(hx_sort_out_82, hx_sort_item_81)
		}
		return hx_sort_out_82
	}(hxrt.StringSplitStringPtr(line, hxrt.StringFromLiteral(" "))))
	var hx_if_85 *string
	if first.Len() > 0 {
		hx_if_85 = hxrt.StdString(func(hx_value_83 any) *string {
			if hx_value_83 == nil {
				var hx_zero_84 *string
				return hx_zero_84
			}
			return hx_value_83.(*string)
		}(first.Get(0)))
	} else {
		hx_if_85 = hxrt.StringFromLiteral("GET")
	}
	method := hx_if_85
	var hx_if_88 *string
	if first.Len() > 1 {
		hx_if_88 = hxrt.StdString(func(hx_value_86 any) *string {
			if hx_value_86 == nil {
				var hx_zero_87 *string
				return hx_zero_87
			}
			return hx_value_86.(*string)
		}(first.Get(1)))
	} else {
		hx_if_88 = hxrt.StringFromLiteral("/")
	}
	path := hx_if_88
	contentLength := 0
	for true {
		header := peer.input.__hx_this.readLine()
		if hxrt.StringEqualStringPtr(header, hxrt.StringFromLiteral("")) {
			break
		}
		lower := hxrt.StringToLowerCaseStringPtr(header)
		if StringTools_startsWith(lower, hxrt.StringFromLiteral("content-length:")) {
			rawLength := StringTools_trim(hxrt.StringSubstrStringPtr(header, hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("content-length:")), 0, false))
			var parsed any = hxrt.StdParseInt(rawLength)
			var hx_if_89 int
			if parsed == nil {
				hx_if_89 = 0
			} else {
				hx_if_89 = parsed.(int)
			}
			contentLength = hx_if_89
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
	}, func(hx_caught_90 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_90)
		_ = hx_tmp
	})
}

func app__http__TinyHttpServer_reason(status int) *string {
	var hx_switch_92 *string
	switch status {
	case 200:
		hx_switch_92 = hxrt.StringFromLiteral("OK")
	case 201:
		hx_switch_92 = hxrt.StringFromLiteral("Created")
	case 400:
		hx_switch_92 = hxrt.StringFromLiteral("Bad Request")
	case 404:
		hx_switch_92 = hxrt.StringFromLiteral("Not Found")
	case 500:
		hx_switch_92 = hxrt.StringFromLiteral("Internal Server Error")
	default:
		hx_switch_92 = hxrt.StringFromLiteral("OK")
	}
	return hx_switch_92
}
