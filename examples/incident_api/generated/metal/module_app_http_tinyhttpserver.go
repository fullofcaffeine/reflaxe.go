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

func New_app__http__TinyHttpServer(api *app__core__IncidentApi, host *string, port int) *app__http__TinyHttpServer {
	self := &app__http__TinyHttpServer{}
	self.__hx_this = self
	self.api = api
	self.host = host
	self.server = New_sys__net__Socket()
	self.server.bind(New_sys__net__Host(host), port)
	self.server.listen(16)
	bound := self.server.host()
	var hx_if_43 int
	if bound == nil {
		hx_if_43 = port
	} else {
		hx_if_43 = func(hx_obj_40 map[string]any) int {
			hx_field_41 := hx_obj_40["port"]
			if hx_field_41 == nil {
				var hx_zero_42 int
				return hx_zero_42
			}
			return hx_field_41.(int)
		}(bound)
	}
	self.port = hx_if_43
	return self
}

func (self *app__http__TinyHttpServer) serveOnce() {
	var peer *sys__net__Socket = nil
	hxrt.TryCatch(func() {
		peer = self.server.accept()
		request := self.readRequest(peer)
		response := self.api.handle(request)
		self.writeResponse(peer, response)
	}, func(hx_caught_44 any) {
		error := hxrt.ExceptionCaught(hx_caught_44)
		_ = error
		if peer != nil {
			self.writeResponse(peer, app__http__HttpResponse_json(500, hxrt.StringFromLiteral("{\"error\":\"server_error\"}")))
		}
	})
	app__http__TinyHttpServer_closePeer(peer)
}

func (self *app__http__TinyHttpServer) close() {
	hxrt.TryCatch(func() {
		self.server.close()
	}, func(hx_caught_46 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_46)
		_ = hx_tmp
	})
}

func (self *app__http__TinyHttpServer) readRequest(peer *sys__net__Socket) *app__http__HttpRequest {
	line := peer.input.readLine()
	first := hxrt.StringSplitStringPtr(line, hxrt.StringFromLiteral(" "))
	var hx_if_48 *string
	if len(first) > 0 {
		hx_if_48 = first[0]
	} else {
		hx_if_48 = hxrt.StringFromLiteral("GET")
	}
	method := hx_if_48
	var hx_if_49 *string
	if len(first) > 1 {
		hx_if_49 = first[1]
	} else {
		hx_if_49 = hxrt.StringFromLiteral("/")
	}
	path := hx_if_49
	contentLength := 0
	for true {
		header := peer.input.readLine()
		if hxrt.StringEqualStringPtr(header, hxrt.StringFromLiteral("")) {
			break
		}
		lower := hxrt.StringToLowerCaseStringPtr(header)
		if StringTools_startsWith(lower, hxrt.StringFromLiteral("content-length:")) {
			rawLength := StringTools_trim(hxrt.StringSubstrStringPtr(header, hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("content-length:")), 0, false))
			var parsed any = hxrt.StdParseInt(rawLength)
			var hx_if_50 int
			if parsed == nil {
				hx_if_50 = 0
			} else {
				hx_if_50 = parsed.(int)
			}
			contentLength = hx_if_50
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
		byte := peer.input.readByte()
		hx_arr_51 := out.b
		hx_arr_51 = append(hx_arr_51, (byte & 255))
		out.b = hx_arr_51
		i = int(int32((i + 1)))
	}
	return out.getBytes().toString()
}

func (self *app__http__TinyHttpServer) writeResponse(peer *sys__net__Socket, response *app__http__HttpResponse) {
	bodyBytes := haxe__io__Bytes_ofString(response.body)
	head := hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("HTTP/1.1 "), response.status), hxrt.StringFromLiteral(" ")), app__http__TinyHttpServer_reason(response.status)), hxrt.StringFromLiteral("\r\n")), hxrt.StringFromLiteral("Content-Type: application/json\r\n")), hxrt.StringFromLiteral("Content-Length: ")), bodyBytes.length), hxrt.StringFromLiteral("\r\n")), hxrt.StringFromLiteral("Connection: close\r\n")), hxrt.StringFromLiteral("\r\n"))
	peer.output.writeString(hxrt.StringConcatStringPtr(head, response.body))
	peer.output.flush()
}

func app__http__TinyHttpServer_closePeer(peer *sys__net__Socket) {
	if peer == nil {
		return
	}
	hxrt.TryCatch(func() {
		peer.close()
	}, func(hx_caught_52 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_52)
		_ = hx_tmp
	})
}

func app__http__TinyHttpServer_reason(status int) *string {
	var hx_switch_54 *string
	switch status {
	case 200:
		hx_switch_54 = hxrt.StringFromLiteral("OK")
	case 201:
		hx_switch_54 = hxrt.StringFromLiteral("Created")
	case 400:
		hx_switch_54 = hxrt.StringFromLiteral("Bad Request")
	case 404:
		hx_switch_54 = hxrt.StringFromLiteral("Not Found")
	case 500:
		hx_switch_54 = hxrt.StringFromLiteral("Internal Server Error")
	default:
		hx_switch_54 = hxrt.StringFromLiteral("OK")
	}
	return hx_switch_54
}
