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
	var hx_if_77 int
	if bound == nil {
		hx_if_77 = port
	} else {
		hx_if_77 = func(hx_obj_74 map[string]any) int {
			hx_field_75 := hx_obj_74["port"]
			if hx_field_75 == nil {
				var hx_zero_76 int
				return hx_zero_76
			}
			return hx_field_75.(int)
		}(bound)
	}
	self.port = hx_if_77
	return self
}

func (self *app__http__TinyHttpServer) serveOnce() {
	var peer *sys__net__Socket = nil
	hxrt.TryCatch(func() {
		peer = self.server.__hx_this.accept()
		request := self.readRequest(peer)
		response := self.api.handle(request)
		self.writeResponse(peer, response)
	}, func(hx_caught_78 any) {
		error := hxrt.ExceptionCaught(hx_caught_78)
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
	}, func(hx_caught_80 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_80)
		_ = hx_tmp
	})
}

func (self *app__http__TinyHttpServer) readRequest(peer *sys__net__Socket) *app__http__HttpRequest {
	line := peer.input.__hx_this.readLine()
	first := hxrt.ArrayFromValues(func(hx_sort_src_82 []*string) []any {
		hx_sort_out_84 := make([]any, 0, len(hx_sort_src_82))
		for _, hx_sort_item_83 := range hx_sort_src_82 {
			hx_sort_out_84 = append(hx_sort_out_84, hx_sort_item_83)
		}
		return hx_sort_out_84
	}(hxrt.StringSplitStringPtr(line, hxrt.StringFromLiteral(" "))))
	var hx_if_87 *string
	if first.Len() > 0 {
		hx_if_87 = hxrt.StdString(func(hx_value_85 any) *string {
			if hx_value_85 == nil {
				var hx_zero_86 *string
				return hx_zero_86
			}
			return hx_value_85.(*string)
		}(first.Get(0)))
	} else {
		hx_if_87 = hxrt.StringFromLiteral("GET")
	}
	method := hx_if_87
	var hx_if_90 *string
	if first.Len() > 1 {
		hx_if_90 = hxrt.StdString(func(hx_value_88 any) *string {
			if hx_value_88 == nil {
				var hx_zero_89 *string
				return hx_zero_89
			}
			return hx_value_88.(*string)
		}(first.Get(1)))
	} else {
		hx_if_90 = hxrt.StringFromLiteral("/")
	}
	path := hx_if_90
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
			var hx_if_91 int
			if parsed == nil {
				hx_if_91 = 0
			} else {
				hx_if_91 = parsed.(int)
			}
			contentLength = hx_if_91
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
	}, func(hx_caught_92 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_92)
		_ = hx_tmp
	})
}

func app__http__TinyHttpServer_reason(status int) *string {
	var hx_switch_94 *string
	switch status {
	case 200:
		hx_switch_94 = hxrt.StringFromLiteral("OK")
	case 201:
		hx_switch_94 = hxrt.StringFromLiteral("Created")
	case 400:
		hx_switch_94 = hxrt.StringFromLiteral("Bad Request")
	case 404:
		hx_switch_94 = hxrt.StringFromLiteral("Not Found")
	case 500:
		hx_switch_94 = hxrt.StringFromLiteral("Internal Server Error")
	default:
		hx_switch_94 = hxrt.StringFromLiteral("OK")
	}
	return hx_switch_94
}
